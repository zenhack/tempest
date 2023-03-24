package servermain

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/grain"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/container"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

var ErrNotLoggedIn = errors.New("you are not logged in")

type externalApiImpl struct {
	server       *server
	userSession  session.UserSession
	sessionStore session.Store
}

func (api externalApiImpl) GetLoginSession(ctx context.Context, p external.ExternalApi_getLoginSession) error {
	if api.userSession.Credential.Type == "" {
		return ErrNotLoggedIn
	}
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetSession(external.LoginSession_ServerToClient(loginSessionImpl{
		externalApiImpl: api,
	}))
	return nil
}

func (api externalApiImpl) Restore(ctx context.Context, p external.ExternalApi_restore) error {
	return capnp.Unimplemented("ExternalApi.restore() is unimplemented.")
}

type loginSessionImpl struct {
	externalApiImpl
}

func (loginSessionImpl) UserInfo(context.Context, external.LoginSession_userInfo) error {
	return capnp.Unimplemented("userInfo() not implemented")
}

func (s loginSessionImpl) ListGrains(ctx context.Context, p external.LoginSession_listGrains) error {
	into := p.Args().Into()
	p.Go()
	return exn.Try0(func(throw func(error)) {
		// TODO(cleanup): update our wrapper to support one-off queries without having to
		// create a whole transaction; this is too much boilerplate.
		tx, err := s.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		info, err := tx.GetCredentialUiViews(s.userSession.Credential)
		throw(err)
		throw(tx.Commit())

		throw(into.Clear(ctx, nil))
		for _, uiViewInfo := range info {
			throw(into.Upsert(ctx, func(p collection.Pusher_upsert_Params) error {
				key, err := capnp.NewText(p.Segment(), string(uiViewInfo.Grain.ID))
				throw(err)
				p.SetKey(key.ToPtr())
				g, err := external.NewGrain(p.Segment())
				throw(err)
				g.SetTitle(uiViewInfo.Grain.Title)
				sessionToken, err := session.GrainSession{
					GrainID:   uiViewInfo.Grain.ID,
					SessionID: s.userSession.SessionID,
				}.Seal(s.sessionStore)
				throw(err)
				g.SetSessionToken(sessionToken)
				// TODO: handle
				p.SetValue(g.ToPtr())
				return nil
			}))
		}
		fut, rel := into.Ready(ctx, nil)
		defer rel()
		throw(into.WaitStreaming())
		_, err = fut.Struct()
		throw(err)
	})
}

func (s loginSessionImpl) ListPackages(ctx context.Context, p external.LoginSession_listPackages) error {
	// TODO: too much boilerplate in common with ListGrains; factor some of this out.
	p.Go()
	into := p.Args().Into()
	return exn.Try0(func(throw func(error)) {
		tx, err := s.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		dbPkgs, err := tx.GetCredentialPackages(s.userSession.Credential)
		throw(err)
		throw(tx.Commit())

		throw(into.Clear(ctx, nil))
		for _, dbPkg := range dbPkgs {
			throw(into.Upsert(ctx, func(p collection.Pusher_upsert_Params) error {
				key, err := capnp.NewText(p.Segment(), dbPkg.ID)
				throw(err)
				p.SetKey(key.ToPtr())
				pkg, err := external.NewPackage(p.Segment())
				throw(err)
				throw(pkg.SetManifest(dbPkg.Manifest))
				pkg.SetController(external.Package_Controller_ServerToClient(pkgController{
					loginSessionImpl: s,
					pkg:              dbPkg,
				}))
				// TODO: controller
				p.SetValue(pkg.ToPtr())
				return nil
			}))
		}
		fut, rel := into.Ready(ctx, nil)
		defer rel()
		throw(into.WaitStreaming())
		_, err = fut.Struct()
		throw(err)
	})
}

type pkgController struct {
	loginSessionImpl
	pkg database.Package
}

func newGrainID() types.GrainID {
	// Oversize buffer so we don't have to think too hard about decoded vs. encoded
	// length:
	var buf [64]byte
	_, err := rand.Read(buf[:])
	util.Chkfatal(err)
	return types.GrainID(base64.URLEncoding.EncodeToString(buf[:])[:22])
}

func (pc pkgController) Create(ctx context.Context, p external.Package_Controller_create) error {
	return exn.Try0(func(th exn.Thrower) {
		args := p.Args()
		actionIndex := args.ActionIndex()
		title, err := args.Title()
		exn.WrapThrow(th, "getting title", err)

		actions, err := pc.pkg.Manifest.Actions()
		exn.WrapThrow(th, "getting actions", err)
		if actionIndex >= uint32(actions.Len()) {
			th(errors.New("actionIndex out of bounds"))
		}
		grainID := newGrainID()

		tx, err := pc.server.db.Begin()
		exn.WrapThrow(th, "creating database transaction", err)

		defer tx.Rollback()
		accountID, err := tx.GetCredentialAccount(pc.userSession.Credential)
		exn.WrapThrow(th, "getting account id", err)

		err = os.MkdirAll(
			config.Localstatedir+"/sandstorm/grains/"+string(grainID)+"/sandbox",
			0770,
		)
		exn.WrapThrow(th, "creating grain sandbox directory", err)
		err = tx.AddGrain(database.NewGrain{
			GrainID: grainID,
			PkgID:   pc.pkg.ID,
			Title:   title,
			OwnerID: accountID,
		})
		exn.WrapThrow(th, "creating grain in database", err)

		startArg, err := exn.Try(func(throw exn.Thrower) string {
			_, seg := capnp.NewSingleSegmentMessage(nil)
			launchCmd, err := grainagent.NewRootLaunchCommand(seg)
			throw(err)
			launchCmd.SetInitGrain(actionIndex)
			return base64.StdEncoding.EncodeToString(seg.Data())
		})
		exn.WrapThrow(th, "encoding LaunchCommand", err)

		results, err := p.AllocResults()
		th(err)

		results.SetId(string(grainID))
		g, err := results.NewGrain()
		th(err)
		th(g.SetTitle(title))
		sessionToken, err := session.GrainSession{
			GrainID:   grainID,
			SessionID: pc.userSession.SessionID,
		}.Seal(pc.sessionStore)
		exn.WrapThrow(th, "creating grain session token", err)
		th(g.SetSessionToken(sessionToken))
		// TODO: set Handle.
		exn.WrapThrow(th, "commiting database transaction", tx.Commit())

		// TODO: maybe change container.Command so it can take tx instead of a DB?
		// But probably we shouldn't do the actual spawning in a tx anyway.
		c, err := container.Command{
			Log:     pc.server.log,
			DB:      pc.server.db,
			GrainID: grainID,
			Api:     grain.SandstormApi_ServerToClient(sandstormApiImpl{}),
			Args:    []string{startArg},
		}.Start(context.TODO())
		exn.WrapThrow(th, "starting container", err)
		pc.server.state.With(func(state *serverState) {
			state.containers.containersByGrainID[grainID] = c
		})
	})

}
