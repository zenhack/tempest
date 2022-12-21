package servermain

import (
	"context"
	"crypto/rand"
	"net/http"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm-next/go/internal/server/embed"
	"zenhack.net/go/sandstorm-next/go/internal/server/session"
	"zenhack.net/go/util"
	websocketcapnp "zenhack.net/go/websocket-capnp"
)

var (
	rootDomain = defaultTo(os.Getenv("ROOT_DOMAIN"), "local.sandstorm.io")
	listenAddr = defaultTo(os.Getenv("LISTEN_ADDR"), ":8000")
)

func defaultTo(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func Main() {
	lg := log.Log

	db := util.Must(database.Open())
	ctx := context.Background()

	containers := &ContainerSet{
		db:                  db,
		containersByGrainId: make(map[string]*container.Container),
	}

	sessionStore := session.NewStore(util.Must(session.GetKeys()))

	r := mux.NewRouter()

	r.Host("ui-{subdomain:[a-zA-Z0-9]+}." + rootDomain).
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.GrainSession

			query := req.URL.Query()
			querySid := query.Has("sandstorm-sid")
			readCookieErr := session.ReadCookie(sessionStore, req, &sess)

			switch {
			case querySid && req.URL.Path == "/_sandstorm-init":
				// Transfer the token from query params to cookie:
				err := sess.Unseal(sessionStore, session.Payload{
					CookieName: sess.CookieName(),
					Data:       query.Get("sandstorm-sid"),
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					lg.WithFields(log.Fields{
						"error":  err,
						"reason": "unsealing sandstorm-sid failed",
					}).Debug("Access to grain UI denied.")
				}
				session.WriteCookie(sessionStore, req, w, sess)
				w.Header().Set("Location", query.Get("path"))
				w.WriteHeader(http.StatusSeeOther)
				// TODO(perf): when doing the redirect,
				// Use http/2 push to avoid a round trip.
			case querySid:
				w.WriteHeader(http.StatusUnauthorized)
				lg.WithFields(log.Fields{
					"url path": req.URL.Path,
					"reason": []string{
						"sandstorm-sid query parameter is present",
						"path is not /_sandstorm-init",
					},
				}).Debug("Access to grain UI denied")
			case readCookieErr != nil:
				w.WriteHeader(http.StatusUnauthorized)
				lg.WithFields(log.Fields{
					"error": readCookieErr,
					"url":   req.URL,
					"reason": []string{
						"no grain session cookie",
						"no sandstorm-sid query parameter",
					},
				}).Debug("Access to grain UI denied")
			default:
				c, err := containers.Get(ctx, sess.GrainId)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					lg.WithFields(log.Fields{
						"grainId": sess.GrainId,
						"error":   err,
					}).Error("Failed to open grain")
					return
				}
				ServeApp(lg, c, w, req)
			}
		})

	r.Host(rootDomain).Path("/login/dev").Methods("GET").
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

	r.Host(rootDomain).Path("/login/dev").Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			sess.Credential.Type = "dev"
			sess.Credential.ScopedId = req.FormValue("name")
			var buf [32]byte
			_, err := rand.Read(buf[:])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				lg.WithField("error", err).Error("crypto/rand.Read() failed")
				return
			}
			sess.SessionId = buf[:]
			session.WriteCookie(sessionStore, req, w, sess)
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
			// TODO:
			// - Check if the credential is already linked to
			//   an account.
			//   - If so, check if it is usable for login
			//   - If not, create one.
		})

	r.Host(rootDomain).Path("/_capnp-api").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var sess session.UserSession
			err := session.ReadCookie(sessionStore, req, &sess)
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
				db:           db,
				userSession:  sess,
				sessionStore: sessionStore,
			}
			rpcConn := rpc.NewConn(transport, &rpc.Options{
				BootstrapClient: capnp.Client(external.ExternalApi_ServerToClient(bootstrap)),
			})
			<-rpcConn.Done()
		})

	r.Host(rootDomain).Handler(http.FileServer(http.FS(embed.Content)))

	http.Handle("/", r)
	lg.Infof("Listening on %v", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func badRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`<!doctype html>
		<html>
			<head>
				<meta charset="utf-8" />
				<title>Bad Request</title>
			</head>
			<body>
				<p>Bad Request: ` + msg + `</p>
			</body>
		</html>`))
}
