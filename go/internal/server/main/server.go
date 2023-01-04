package servermain

import (
	"context"
	"crypto/rand"
	"net/http"
	"strings"
	"sync"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/go/internal/database"
	"zenhack.net/go/tempest/go/internal/server/container"
	"zenhack.net/go/tempest/go/internal/server/embed"
	"zenhack.net/go/tempest/go/internal/server/session"
	websocketcapnp "zenhack.net/go/websocket-capnp"
)

type webSessionParams struct {
	BasePath            string
	UserAgent           string
	AcceptableLanguages []string
}

func (p *webSessionParams) FromRequest(req *http.Request) {
	p.BasePath = req.URL.Scheme + "://" + req.URL.Host
	p.UserAgent = req.Header.Get("User-Agent")
	p.AcceptableLanguages = strings.Split(
		req.Header.Get("Accept-Language"),
		",",
	)
}

func (p *webSessionParams) Insert(into websession.WebSession_Params) error {
	return pogs.Insert(websession.WebSession_Params_TypeID, capnp.Struct(into), p)
}

// A server encapsulates the state of a running server.
type server struct {
	rootDomain   string // Main Sandstorm domain name
	log          log.Interface
	db           database.DB
	sessionStore session.Store

	// State that requires synchronization when accessed by multiple goroutines;
	// the mutex must be held when accessin these fields:
	lk struct {
		sync.Mutex
		grainSessions map[grainSessionKey]grainSession
		containers    ContainerSet
	}
}

func newServer(rootDomain string, lg log.Interface, db database.DB, sessionStore session.Store) *server {
	srv := &server{
		rootDomain:   rootDomain,
		log:          lg,
		db:           db,
		sessionStore: sessionStore,
	}
	srv.lk.containers = ContainerSet{
		containersByGrainId: make(map[string]container.Container),
	}
	srv.lk.grainSessions = make(map[grainSessionKey]grainSession)
	return srv
}

type grainSessionKey struct {
	userSessionId string
	grainId       string
	// TODO: we'll want the ability to have more than one session
	// per user session, e.g. for powerbox requests, maybe multiple
	// tabs.
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
				s.lk.Lock()
				var unlocked bool
				unlock := func() { // idempotent wrapper around s.lk.Unlock
					if !unlocked {
						s.lk.Unlock()
						unlocked = true
					}
				}
				defer unlock()

				key := grainSessionKey{
					userSessionId: string(sess.SessionId),
					grainId:       sess.GrainId,
				}
				gs, ok := s.lk.grainSessions[key]
				if !ok {
					c, err := s.lk.containers.Get(context.Background(), s.db, sess.GrainId)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						s.log.WithFields(log.Fields{
							"grainId": sess.GrainId,
							"error":   err,
						}).Error("Failed to open grain")
						return
					}
					mainView := grain.MainView(c.Bootstrap.AddRef())
					defer mainView.Release()
					ctx := req.Context()
					sessionCtx := grain.SessionContext_ServerToClient(sessionCtxImpl{})
					fut, rel := mainView.NewSession(
						ctx,
						func(p grain.UiView_newSession_Params) error {
							p.SetSessionType(websession.WebSession_TypeID)
							p.SetContext(sessionCtx)
							p.SetTabId([]byte("TODO"))
							params, err := websession.NewWebSession_Params(p.Segment())
							if err != nil {
								return err
							}
							var wsp webSessionParams
							wsp.FromRequest(req)
							wsp.Insert(params)
							p.SetSessionParams(params.ToPtr())
							return nil
						})
					defer rel()
					res, err := fut.Struct()
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						s.log.WithFields(log.Fields{
							"grainId": sess.GrainId,
							"error":   err,
						}).Error("UiView.newSession() failed")
						return
					}

					webSession := websession.WebSession(res.Session().AddRef())
					gs = grainSession{
						webSession: webSession,
					}
					s.lk.grainSessions[key] = gs
				}
				session := gs.webSession.AddRef()
				defer session.Release()
				unlock()
				ServeApp(session, w, req)
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
			up := &websocket.Upgrader{
				Subprotocols:      []string{"capnp-rpc"},
				EnableCompression: true,
			}
			wsConn, err := up.Upgrade(w, req, nil)
			if err != nil {
				return
			}
			wsConn.EnableWriteCompression(true)
			transport := websocketcapnp.NewTransport(wsConn)
			defer transport.Close()
			bootstrap := externalApiImpl{
				db:           s.db,
				userSession:  sess,
				sessionStore: s.sessionStore,
			}
			rpcConn := rpc.NewConn(transport, &rpc.Options{
				BootstrapClient: capnp.Client(external.ExternalApi_ServerToClient(bootstrap)),
			})
			<-rpcConn.Done()
		})

	r.Host(s.rootDomain).Handler(http.FileServer(http.FS(embed.Content)))

	return r
}

func (s *server) Release() {
	s.db.Close()
	s.lk.containers.Release()
	for _, sess := range s.lk.grainSessions {
		sess.Release()
	}
	*s = server{}
}
