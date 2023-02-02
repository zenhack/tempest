package servermain

import (
	"context"
	"crypto/rand"
	"net/http"
	"strings"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"github.com/apex/log"
	"github.com/gobwas/ws"
	"github.com/gorilla/mux"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/grain"
	websession "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/go/internal/database"
	"zenhack.net/go/tempest/go/internal/server/container"
	"zenhack.net/go/tempest/go/internal/server/embed"
	"zenhack.net/go/tempest/go/internal/server/session"
	"zenhack.net/go/util/sync/mutex"
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
	rootDomain   string // Main Sandstorm domain name
	log          log.Interface
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

func newServer(rootDomain string, lg log.Interface, db database.DB, sessionStore session.Store) *server {
	return &server{
		rootDomain:   rootDomain,
		log:          lg,
		db:           db,
		sessionStore: sessionStore,
		state: mutex.New[serverState](serverState{
			containers: ContainerSet{
				containersByGrainId: make(map[string]container.Container),
			},
			grainSessions: make(map[grainSessionKey]grainSession),
		}),
	}
}

type grainSessionKey struct {
	userSessionId string
	grainId       string

	// Things that go in WebSession.Params.
	basePath            string
	userAgent           string
	acceptableLanguages string
}

type grainSession struct {
	webSession websession.WebSession
}

func (s grainSession) Release() {
	s.webSession.Release()
}

func (s *server) Handler() http.Handler {
	r := mux.NewRouter()

	r.Host("ui-{subdomain:[a-zA-Z0-9]+}." + s.rootDomain).
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
					s.log.WithFields(log.Fields{
						"error":  err,
						"reason": "unsealing sandstorm-sid failed",
					}).Debug("Access to grain UI denied.")
				}
				session.WriteCookie(s.sessionStore, req, w, sess)
				w.Header().Set("Location", query.Get("path"))
				w.WriteHeader(http.StatusSeeOther)
				// TODO(perf): when doing the redirect,
				// Use http/2 push to avoid a round trip.
			case querySid:
				w.WriteHeader(http.StatusUnauthorized)
				s.log.WithFields(log.Fields{
					"url path": req.URL.Path,
					"reason": []string{
						"sandstorm-sid query parameter is present",
						"path is not /_sandstorm-init",
					},
				}).Debug("Access to grain UI denied")
			case readCookieErr != nil:
				w.WriteHeader(http.StatusUnauthorized)
				s.log.WithFields(log.Fields{
					"error": readCookieErr,
					"url":   req.URL,
					"reason": []string{
						"no grain session cookie",
						"no sandstorm-sid query parameter",
					},
				}).Debug("Access to grain UI denied")
			default:
				var wsp webSessionParams
				wsp.FromRequest(req)
				session, err := s.getWebSession(req.Context(), wsp, sess)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					s.log.WithFields(log.Fields{
						"grainId": sess.GrainId,
						"params":  wsp,
						"error":   err,
					}).Error("Could not get web session reference")
					return
				}
				defer session.Release()
				ServeApp(session, w, req, s.rootDomain)
			}
		})

	r.Host(s.rootDomain).Path("/login/dev").Methods("GET").
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

	r.Host(s.rootDomain).Path("/login/dev").Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			sess.Credential.Type = "dev"
			sess.Credential.ScopedId = req.FormValue("name")
			var buf [32]byte
			_, err := rand.Read(buf[:])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				s.log.WithField("error", err).Error("crypto/rand.Read() failed")
				return
			}
			sess.SessionId = buf[:]
			session.WriteCookie(s.sessionStore, req, w, sess)
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
			// TODO:
			// - Check if the credential is already linked to
			//   an account.
			//   - If so, check if it is usable for login
			//   - If not, create one.
		})

	r.Host(s.rootDomain).Path("/_capnp-api").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			err := session.ReadCookie(s.sessionStore, req, &sess)
			if err != nil {
				s.log.WithField("error", err).
					Debug("Failed to read session cookie; treating as anonymous")
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
				s.log.WithField("error", err).
					Error("Failed to upgrade http connection")
				return
			}
			transport := transport.New(codec)
			defer transport.Close()
			bootstrap := externalApiImpl{
				db:           s.db,
				userSession:  sess,
				sessionStore: s.sessionStore,
			}
			rpcConn := rpc.NewConn(transport, &rpc.Options{
				BootstrapClient: capnp.Client(external.ExternalApi_ServerToClient(bootstrap)),
				ErrorReporter:   logErrorReporter{log: s.log},
			})
			<-rpcConn.Done()
		})

	r.Host(s.rootDomain).Handler(http.FileServer(http.FS(embed.Content)))

	return r
}

func (s *server) getWebSession(ctx context.Context, wsp webSessionParams, sess session.GrainSession) (websession.WebSession, error) {

	key := grainSessionKey{
		userSessionId: string(sess.SessionId),
		grainId:       sess.GrainId,

		basePath:            wsp.BasePath,
		userAgent:           wsp.UserAgent,
		acceptableLanguages: strings.Join(wsp.AcceptableLanguages, ","),
	}
	return mutex.With2(&s.state, func(state *serverState) (websession.WebSession, error) {
		gs, ok := state.grainSessions[key]
		if !ok {
			c, err := state.containers.Get(context.Background(), s.log, s.db, sess.GrainId)
			if err != nil {
				return websession.WebSession{}, err
			}
			mainView := grain.MainView(c.Bootstrap.AddRef())
			defer mainView.Release()
			sessionCtx := grain.SessionContext_ServerToClient(sessionCtxImpl{})
			fut, rel := mainView.NewSession(
				ctx,
				func(p grain.UiView_newSession_Params) error {
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
			// FIXME: Do this outside of With2 somehow (probably by inserting a promise
			// into the map). Otherwise, the grain can block all other sessions from starting...
			defer rel()
			res, err := fut.Struct()
			if err != nil {
				return websession.WebSession{}, err
			}

			webSession := websession.WebSession(res.Session().AddRef())
			gs = grainSession{
				webSession: webSession,
			}
			state.grainSessions[key] = gs
		}
		return gs.webSession.AddRef(), nil
	})
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
	log log.Interface
}

func (l logErrorReporter) ReportError(err error) {
	l.log.WithField("error", err).Error("capnp-rpc error")
}
