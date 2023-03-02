package servermain

import (
	"context"
	"errors"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/database"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util/exn"
)

var ErrNotLoggedIn = errors.New("You are not logged in.")

type externalApiImpl struct {
	db           database.DB
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
		db:           api.db,
		userSession:  api.userSession,
		sessionStore: api.sessionStore,
	}))
	return nil
}

func (api externalApiImpl) Restore(ctx context.Context, p external.ExternalApi_restore) error {
	return capnp.Unimplemented("ExternalApi.restore() is unimplemented.")
}

type loginSessionImpl struct {
	db           database.DB
	userSession  session.UserSession
	sessionStore session.Store
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
		tx, err := s.db.Begin()
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
		tx, err := s.db.Begin()
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
