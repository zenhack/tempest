package servermain

import (
	"context"
	"crypto/rand"
	"net/http"
	"sync"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"zenhack.net/go/sandstorm-next/capnp/external"
	httpcp "zenhack.net/go/sandstorm-next/capnp/http"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm-next/go/internal/server/embed"
	"zenhack.net/go/sandstorm-next/go/internal/server/session"
	websocketcapnp "zenhack.net/go/websocket-capnp"
)

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
		containersByGrainId: make(map[string]*container.Container),
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
	// TODO: switch over to websession
	webServer httpcp.Server
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
					srv := httpcp.Server(c.Bootstrap.AddRef())
					gs = grainSession{
						webServer: srv,
					}
					s.lk.grainSessions[key] = gs
				}
				srv := gs.webServer.AddRef()
				defer srv.Release()
				unlock()
				ServeApp(s.log, srv, w, req)
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
