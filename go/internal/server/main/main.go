package servermain

import (
	"context"
	"log"
	"net/http"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm-next/go/internal/server/embed"
	"zenhack.net/go/sandstorm-next/go/internal/server/usersession"
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

func Main() {
	db := util.Must(database.Open())
	ctx := context.Background()
	c := util.Must(container.StartDummy(ctx, db))
	defer c.Release()

	sessionStore := sessions.NewCookieStore(util.Must(usersession.GetKey()))

	r := mux.NewRouter()

	r.Host("grain." + rootDomain).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ServeApp(c, w, req)
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
			name := req.FormValue("name")
			sess := usersession.Get(sessionStore, req)
			sess.Data.Credential.Type = "dev"
			sess.Data.Credential.ScopedId = name
			sess.Save(req, w)
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
			session := usersession.Get(sessionStore, req)
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
				db:          db,
				userSession: session,
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
