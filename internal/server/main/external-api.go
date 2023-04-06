package servermain

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
	"strings"
	"time"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/grain"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
	"zenhack.net/go/tempest/internal/capnp/system"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/container"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/tempest/internal/server/tokenutil"
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

func (api externalApiImpl) Authenticator(ctx context.Context, p external.ExternalApi_authenticator) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetAuthenticator(external.Authenticator_ServerToClient(authenticatorImpl{
		api: api,
	}))
	return nil
}

type authenticatorImpl struct {
	api externalApiImpl
}

func (a authenticatorImpl) SendEmailAuthToken(ctx context.Context, p external.Authenticator_sendEmailAuthToken) error {
	return exn.Try0(func(throw exn.Thrower) {
		addr, err := p.Args().Address()
		throw(err)
		db := a.api.server.db
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()

		// FIXME: sanitize addr?

		_, seg := capnp.NewSingleSegmentMessage(nil)
		oid, err := system.NewRootSystemObjectId(seg)
		throw(err)
		throw(oid.SetEmailLoginToken(addr))

		token := base64.RawURLEncoding.EncodeToString(tokenutil.GenToken()[:16])

		_, err = tx.SaveSturdyRef(
			database.SturdyRefKey{
				Token:     []byte(token),
				OwnerType: "external",
				Owner:     "",
			},
			database.SturdyRefValue{
				Expires:  time.Now().Add(10 * time.Minute),
				ObjectID: capnp.Struct(oid),
			},
		)
		throw(err)
		throw(tx.Commit())

		cfg := a.api.server.cfg
		throw(cfg.smtp.SendMail(
			[]string{addr},
			[]byte(strings.Join([]string{
				"To: " + addr,
				"From: " + cfg.smtp.Username,
				"Subject: Email Login Token",
				"",
				"Login in as " + addr + " by visiting:",
				"",
				cfg.rootDomain + "/login/email/" + token,
				"",
				"Or entering " + token + " at the login prompt.",
			}, "\r\n")),
		))
	})
}

type loginSessionImpl struct {
	externalApiImpl
}

func (loginSessionImpl) UserInfo(context.Context, external.LoginSession_userInfo) error {
	return capnp.Unimplemented("userInfo() not implemented")
}

func (s loginSessionImpl) UserSession(ctx context.Context, p external.LoginSession_userSession) error {
	p.Go()
	return exn.Try0(func(throw exn.Thrower) {
		tx, err := s.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		ok, err := tx.CredentialHasRole(s.userSession.Credential, database.RoleUser)
		throw(err)
		if !ok {
			throw(errors.New("Permission denied; caller does not have the 'user' role."))
		}
		throw(tx.Commit())

		results, err := p.AllocResults()
		throw(err)
		results.SetSession(external.UserSession_ServerToClient(userSessionImpl{
			api: s.externalApiImpl,
		}))
	})
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
				g.SetSubdomain(hex.EncodeToString(tokenutil.GenToken()[:16]))
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

type userSessionImpl struct {
	api externalApiImpl
}

func (s userSessionImpl) InstallPackage(ctx context.Context, p external.UserSession_installPackage) error {
	return capnp.Unimplemented("TODO: implement InstallPackage()")
}
