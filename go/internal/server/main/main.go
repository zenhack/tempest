package servermain

import (
	"context"
	"crypto/rand"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/sandstorm-next/go/internal/config"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm-next/go/internal/server/embed"
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

func getSessionKey() []byte {
	const path = config.Localstatedir + "/sandstorm/session-key"
	data, err := ioutil.ReadFile(config.Localstatedir + "/sandstorm/session-key")
	if os.IsNotExist(err) {
		data := make([]byte, 32)
		rand.Read(data)
		util.Chkfatal(os.WriteFile(path, data, 0600))
	} else {
		util.Chkfatal(err)
	}
	return data
}

type userSession struct {
	session *sessions.Session
}

func getUserSession(store sessions.Store, req *http.Request) userSession {
	session, _ := store.Get(req, "user-session")
	return userSession{
		session: session,
	}
}

func (s userSession) Id() string {
	sessionId, ok := s.session.Values["session-id"]
	if !ok {
		return ""
	}
	ret, ok := sessionId.(string)
	if !ok {
		return ""
	}
	return ret
}

func (s userSession) SetDevCredential(name string) {
	s.session.Values["dev-credential-name"] = name
}

func (s userSession) Save(req *http.Request, w http.ResponseWriter) {
	s.session.Save(req, w)
}

func Main() {
	db, err := database.Open()
	util.Chkfatal(err)
	ctx := context.Background()
	c, err := container.StartDummy(ctx, db)
	util.Chkfatal(err)
	defer c.Release()

	sessionStore := sessions.NewCookieStore(getSessionKey())

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
			sess := getUserSession(sessionStore, req)
			sess.SetDevCredential(name)
			sess.Save(req, w)
			w.Header().Set("Location", "/")
			w.WriteHeader(http.StatusSeeOther)
		})

	r.Host(rootDomain).Path("/_capnp-api").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			session := getUserSession(sessionStore, req)
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
