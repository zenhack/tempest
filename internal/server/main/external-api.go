package servermain

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/exc"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

var ErrNotLoggedIn = errors.New("You are not logged in.")

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
		c := s.userSession.Credential
		info, err := tx.GetCredentialUiViews(c.Type, c.ScopedId)
		throw(err)
		throw(tx.Commit())

		_, rel := into.Clear(ctx, nil)
		releaseFuncs := []capnp.ReleaseFunc{rel}
		for _, uiViewInfo := range info {
			_, rel = into.Upsert(ctx, func(p collection.Pusher_upsert_Params) error {
				key, err := capnp.NewText(p.Segment(), uiViewInfo.Grain.Id)
				throw(err)
				p.SetKey(key.ToPtr())
				g, err := external.NewGrain(p.Segment())
				throw(err)
				g.SetTitle(uiViewInfo.Grain.Title)
				sessionToken, err := session.GrainSession{
					GrainId:   uiViewInfo.Grain.Id,
					SessionId: s.userSession.SessionId,
				}.Seal(s.sessionStore)
				throw(err)
				g.SetSessionToken(sessionToken)
				// TODO: handle
				p.SetValue(g.ToPtr())
				return nil
			})
			releaseFuncs = append(releaseFuncs, rel)
		}
		_, rel = into.Ready(ctx, nil)
		releaseFuncs = append(releaseFuncs, rel)
		for _, rel := range releaseFuncs {
			rel()
		}
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
		c := s.userSession.Credential
		dbPkgs, err := tx.GetCredentialPackages(c.Type, c.ScopedId)
		throw(err)
		throw(tx.Commit())

		_, rel := into.Clear(ctx, nil)
		releaseFuncs := []capnp.ReleaseFunc{rel}
		for _, dbPkg := range dbPkgs {
			_, rel = into.Upsert(ctx, func(p collection.Pusher_upsert_Params) error {
				key, err := capnp.NewText(p.Segment(), dbPkg.Id)
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
			})
			releaseFuncs = append(releaseFuncs, rel)
		}
		_, rel = into.Ready(ctx, nil)
		releaseFuncs = append(releaseFuncs, rel)
		for _, rel := range releaseFuncs {
			rel()
		}
	})
}

type pkgController struct {
	loginSessionImpl
	pkg database.Package
}

func newGrainID() string {
	buf := make([]byte, base64.URLEncoding.DecodedLen(22))
	_, err := rand.Read(buf[:])
	util.Chkfatal(err)
	return base64.URLEncoding.EncodeToString(buf)
}

func (pc pkgController) Create(ctx context.Context, p external.Package_Controller_create) error {
	args := p.Args()
	actionIndex := args.ActionIndex()
	title, err := args.Title()
	if err != nil {
		return err
	}

	actions, err := pc.pkg.Manifest.Actions()
	if err != nil {
		return err
	}
	if actionIndex >= uint32(actions.Len()) {
		return errors.New("actionIndex out of bounds")
	}
	grainID := newGrainID()

	tx, err := pc.server.db.Begin()
	if err != nil {
		return exc.WrapError("Creating database transaction", err)
	}
	defer tx.Rollback()
	accountID, err := tx.GetCredentialAccount(
		pc.userSession.Credential.Type,
		pc.userSession.Credential.ScopedId,
	)
	if err != nil {
		return exc.WrapError("Getting account ID", err)
	}

	err = os.MkdirAll(
		config.Localstatedir+"/sandstorm/grains/"+grainID+"/sandbox",
		0770,
	)
	if err != nil {
		return exc.WrapError("Creating grain sandbox directory", err)
	}

	err = tx.AddGrain(database.NewGrain{
		GrainId: grainID,
		PkgId:   pc.pkg.Id,
		Title:   title,
		OwnerId: accountID,
	})
	if err != nil {
		return exc.WrapError("Creating grain in database", err)
	}

	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetId(grainID)
	g, err := results.NewGrain()
	if err != nil {
		return err
	}
	g.SetTitle(title)
	// FIXME: Start the grain with the right action; as is this will just get launched w/
	// continue when it hits the UI. Will still work for many apps.
	sessionToken, err := session.GrainSession{
		GrainId:   grainID,
		SessionId: pc.userSession.SessionId,
	}.Seal(pc.sessionStore)
	if err != nil {
		return err
	}
	g.SetSessionToken(sessionToken)
	// TODO: set Handle.
	return tx.Commit()

}
