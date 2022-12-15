package servermain

import (
	"context"
	"crypto/rand"
	"log"
	"net/http"
	"os"
	"sync"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
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

func SetAppHeaders(w http.ResponseWriter) {
}

type ContainerSet struct {
	mu                  sync.Mutex
	db                  database.DB
	containersByGrainId map[string]*container.Container
}

func (cset *ContainerSet) Get(ctx context.Context, grainId string) (*container.Container, error) {
	cset.mu.Lock()
	defer cset.mu.Unlock()
	c, ok := cset.containersByGrainId[grainId]
	if ok {
		return c, nil
	}
	c, err := container.Start(ctx, cset.db, grainId)
	if err == nil {
		cset.containersByGrainId[grainId] = c
	}
	return c, err
}

func Main() {
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
			err := session.ReadCookie(sessionStore, req, &sess)
			if err == nil {
				c, err := containers.Get(ctx, sess.GrainId)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Println("Opening grain: ", err)
					return
				}
				ServeApp(c, w, req)
				return
			}

			// See if we can transfer the token from query params to
			// cookie:

			if req.URL.Path != "/_sandstorm-init" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			query := req.URL.Query()
			err = sess.Unseal(sessionStore, session.Payload{
				CookieName: sess.CookieName(),
				Data:       query.Get("sandstorm-sid"),
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println("error unsealing: ", err)
			}
			session.WriteCookie(sessionStore, req, w, sess)
			w.Header().Set("Location", query.Get("path"))
			// FIXME: sanity check this is the right redirect:
			w.WriteHeader(http.StatusSeeOther)
			return

			// TODO(perf): when doing the redirect,
			// Use http/2 push to avoid a round trip.
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
				log.Println(err)
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
	log.Printf("Listening on %v", listenAddr)
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
