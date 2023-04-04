package servermain

import (
	"context"
	"net/http"
	"strings"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"github.com/gobwas/ws"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/grain"
	websession "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/internal/capnp/system"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/container"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/embed"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util/orerr"
	"zenhack.net/go/util/sync/mutex"
	"zenhack.net/go/util/thunk"
	websocketcapnp "zenhack.net/go/websocket-capnp"
)

type webSessionParams struct {
	BasePath            string
	UserAgent           string
	AcceptableLanguages []string
}

func (p *webSessionParams) FromRequest(req *http.Request) {
	p.BasePath = "http"
	if req.TLS != nil {
		p.BasePath += "s"
	}
	p.BasePath += "://" + req.Host
	p.UserAgent = req.Header.Get("User-Agent")
	p.AcceptableLanguages = strings.Split(
		req.Header.Get("Accept-Language"),
		",",
	)
}

func (p *webSessionParams) Insert(into websession.Params) error {
	return pogs.Insert(websession.Params_TypeID, capnp.Struct(into), p)
}

// A server encapsulates the state of a running server.
type server struct {
	cfg          Config
	log          *slog.Logger
	db           database.DB
	sessionStore session.Store
	state        mutex.Mutex[serverState]
}

// Server state that requires synchronization when accessed by multiple goroutines;
// this is factored out so we can put a lock around it.
type serverState struct {
	grainSessions map[grainSessionKey]grainSession
	containers    ContainerSet
}

func newServer(cfg Config, lg *slog.Logger, db database.DB, sessionStore session.Store) *server {
	return &server{
		cfg:          cfg,
		log:          lg,
		db:           db,
		sessionStore: sessionStore,
		state: mutex.New[serverState](serverState{
			containers: ContainerSet{
				containersByGrainID: make(map[types.GrainID]container.Container),
			},
			grainSessions: make(map[grainSessionKey]grainSession),
		}),
	}
}

type grainSessionKey struct {
	userSessionID string
	grainID       types.GrainID

	// Things that go in WebSession.Params.
	basePath            string
	userAgent           string
	acceptableLanguages string
}

type grainSession struct {
	webSession *thunk.T[orerr.T[websession.WebSession]]
}

func (s grainSession) Release() {
	sess, err := s.webSession.Force().Get()
	if err == nil {
		sess.Release()
	}
}

func (s *server) Handler() http.Handler {
	r := mux.NewRouter()

	r.Host("ui-{subdomain:[a-zA-Z0-9]+}." + s.cfg.rootDomain).
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.GrainSession

			query := req.URL.Query()
			querySid := query.Has("sandstorm-sid")
			readCookieErr := session.ReadCookie(s.sessionStore, req, &sess)

			switch {
			case querySid && req.URL.Path == "/_sandstorm-init":
				// Transfer the token from query params to cookie:
				err := sess.Unseal(s.sessionStore, session.Payload{
					CookieName: sess.CookieName(),
					Data:       query.Get("sandstorm-sid"),
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					s.log.Debug("Access to grain UI denied.",
						"error", err,
						"reason", "unsealing sandstorm-sid failed",
					)
				}
				session.WriteCookie(s.sessionStore, req, w, sess)
				http.Redirect(w, req, query.Get("path"), http.StatusSeeOther)
				// TODO(perf): when doing the redirect,
				// Use http/2 push to avoid a round trip.
			case querySid:
				w.WriteHeader(http.StatusUnauthorized)
				s.log.Debug("Access to grain UI denied",
					"url path", req.URL.Path,
					"reason", []string{
						"sandstorm-sid query parameter is present",
						"path is not /_sandstorm-init",
					},
				)
			case readCookieErr != nil:
				w.WriteHeader(http.StatusUnauthorized)
				s.log.Debug("Access to grain UI denied",
					"error", readCookieErr,
					"url", req.URL,
					"reason", []string{
						"no grain session cookie",
						"no sandstorm-sid query parameter",
					},
				)
			default:
				var wsp webSessionParams
				wsp.FromRequest(req)
				session, err := s.getWebSession(req.Context(), wsp, sess)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					s.log.Error(
						"Could not get web session reference", err,
						"grainID", sess.GrainID,
						"params", wsp,
					)
					return
				}
				defer session.Release()
				ServeApp(session, w, req, s.cfg.rootDomain)
			}
		})

	r.Host(s.cfg.rootDomain).Path("/login/dev").Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(`<!doctype html>
			<html>
			<head>
			<meta charset="utf-8" />
			<title>Login dev account</title>
			</head>
			<body>
			<form action="/login/dev" method="post">
				<input name="name">
				<button type="submit">Submit</button>
			</form>
			</body>
			</html>
			`))
		})

	r.Host(s.cfg.rootDomain).Path("/login/dev").Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			sess.Credential.Type = "dev"
			sess.Credential.ScopedID = req.FormValue("name")
			sess.SessionID = session.GenSessionID()
			session.WriteCookie(s.sessionStore, req, w, sess)
			http.Redirect(w, req, "/", http.StatusSeeOther)
			// TODO:
			// - Check if the credential is already linked to
			//   an account.
			//   - If so, check if it is usable for login
			//   - If not, create one.
		})

	r.Host(s.cfg.rootDomain).Path("/login/email/{token}").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token := mux.Vars(req)["token"]
			tx, err := s.db.Begin()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				s.log.Error("failed to open database transaction", err)
				return
			}
			defer tx.Rollback()
			key := database.SturdyRefKey{
				Token:     []byte(token),
				OwnerType: "external",
				Owner:     "",
			}
			ref, err := tx.RestoreSturdyRef(key)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("No such token (maybe expired?)"))
				s.log.Debug("failed to restore token",
					"error", err,
				)
				return
			}
			if err = tx.DeleteSturdyRef(key); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				s.log.Error("deleting sturdyref: ", err)
				return
			}
			if err = tx.Commit(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				s.log.Error("restoring email token: commit", err)
				return
			}
			oid := system.SystemObjectId(ref.ObjectID)
			if oid.Which() != system.SystemObjectId_Which_emailLoginToken {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("token has the wrong type"))
				return
			}
			addr, err := oid.EmailLoginToken()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				s.log.Error("reading email address from token", err)
				return
			}

			sess := session.UserSession{
				SessionID: session.GenSessionID(),
				Credential: types.Credential{
					Type:     types.EmailCredential,
					ScopedID: addr,
				},
			}
			session.WriteCookie(s.sessionStore, req, w, sess)
			http.Redirect(w, req, "/", http.StatusSeeOther)
		})

	r.Host(s.cfg.rootDomain).Path("/_capnp-api").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			err := session.ReadCookie(s.sessionStore, req, &sess)
			if err != nil {
				s.log.Debug("Failed to read session cookie; treating as anonymous",
					"error", err,
				)
				// Don't rely on ReadCookie leaving the zero value in place:
				sess = session.UserSession{}
			}
			codec, err := websocketcapnp.UpgradeHTTP(
				ws.HTTPUpgrader{
					Protocol: func(s string) bool {
						return s == "capnp-rpc"
					},
				}, req, w)
			if err != nil {
				s.log.Error("Failed to upgrade http connection", err)
				return
			}
			transport := transport.New(codec)
			defer transport.Close()
			bootstrap := externalApiImpl{
				server:       s,
				userSession:  sess,
				sessionStore: s.sessionStore,
			}
			rpcConn := rpc.NewConn(transport, &rpc.Options{
				BootstrapClient: capnp.Client(external.ExternalApi_ServerToClient(bootstrap)),
				ErrorReporter:   logErrorReporter{log: s.log},
			})
			<-rpcConn.Done()
		})

	r.Host(s.cfg.rootDomain).Handler(http.FileServer(http.FS(embed.Content)))

	return r
}

func (s *server) getWebSession(ctx context.Context, wsp webSessionParams, sess session.GrainSession) (websession.WebSession, error) {

	key := grainSessionKey{
		userSessionID: string(sess.SessionID),
		grainID:       sess.GrainID,

		basePath:            wsp.BasePath,
		userAgent:           wsp.UserAgent,
		acceptableLanguages: strings.Join(wsp.AcceptableLanguages, ","),
	}
	webSessionThunk := mutex.With1(&s.state, func(state *serverState) *thunk.T[orerr.T[websession.WebSession]] {
		gs, ok := state.grainSessions[key]
		if ok {
			return gs.webSession
		}
		c, err := state.containers.Get(context.Background(), s.log, s.db, sess.GrainID)
		if err != nil {
			return thunk.Ready(orerr.New(websession.WebSession{}, err))
		}
		webSessionThunk := thunk.Go(func() orerr.T[websession.WebSession] {
			mainView := grain.MainView(c.Bootstrap.AddRef())
			defer mainView.Release()
			sessionCtx := grain.SessionContext_ServerToClient(sessionCtxImpl{})
			// TODO: we shouldn't need to do this for every session we get, only on
			// grain boot.
			viewInfoFut, rel := mainView.GetViewInfo(ctx, nil)
			defer rel()

			viewInfo, err := viewInfoFut.Struct()
			if err != nil {
				return orerr.New(websession.WebSession{}, err)
			}
			tx, err := s.db.Begin()
			if err != nil {
				return orerr.New(websession.WebSession{}, err)
			}
			defer tx.Rollback()
			if err = tx.SetGrainViewInfo(string(sess.GrainID), viewInfo); err != nil {
				return orerr.New(websession.WebSession{}, err)
			}
			if err = tx.Commit(); err != nil {
				return orerr.New(websession.WebSession{}, err)
			}

			viewInfoPermissions, err := viewInfo.Permissions()
			if err != nil {
				return orerr.New(websession.WebSession{}, err)
			}

			newSessionFut, rel := mainView.NewSession(
				ctx,
				func(p grain.UiView_newSession_Params) error {
					userInfo, err := p.NewUserInfo()
					if err != nil {
						return err
					}

					// For now, just give the user all permissions.
					// we'll store & retrieve this info properly
					// later on.
					permissions, err := userInfo.NewPermissions(int32(viewInfoPermissions.Len()))
					if err != nil {
						return err
					}
					for i := 0; i < permissions.Len(); i++ {
						permissions.Set(i, true)
					}

					p.SetSessionType(websession.WebSession_TypeID)
					p.SetContext(sessionCtx)
					p.SetTabId([]byte("TODO"))
					params, err := websession.NewParams(p.Segment())
					if err != nil {
						return err
					}
					wsp.Insert(params)
					p.SetSessionParams(params.ToPtr())
					return nil
				})
			defer rel()
			newSessionRes, err := newSessionFut.Struct()
			if err != nil {
				return orerr.New(websession.WebSession{}, err)
			}

			webSession := websession.WebSession(newSessionRes.Session().AddRef())
			return orerr.New(webSession, nil)
		})

		state.grainSessions[key] = grainSession{
			webSession: webSessionThunk,
		}
		return webSessionThunk
	})
	webSession, err := webSessionThunk.Force().Get()
	return webSession.AddRef(), err
}

func (s *server) Release() {
	s.db.Close()
	s.state.With(func(state *serverState) {
		state.containers.Release()
		for _, sess := range state.grainSessions {
			sess.Release()
		}
	})
}

// Implementation of capnp.ErrorReporter on top of our logger. TODO(cleanup):
// move this somewhere more appropriate.
type logErrorReporter struct {
	log *slog.Logger
}

func (l logErrorReporter) ReportError(err error) {
	l.log.Error("capnp-rpc error", err)
}
