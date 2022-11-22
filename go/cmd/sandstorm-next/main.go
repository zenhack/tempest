package main

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
	"zenhack.net/go/sandstorm-next/go/internal/container"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/webui/embed"
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

func userSessionId(store sessions.Store, req *http.Request) string {
	session, _ := store.Get(req, "user-session")
	sessionId, ok := session.Values["session-id"]
	if !ok {
		return ""
	}
	ret, ok := sessionId.(string)
	if !ok {
		return ""
	}
	return ret
}

func main() {
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

	r.Host(rootDomain).Path("/_capnp-api").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			sid := userSessionId(sessionStore, req)
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
				db:            db,
				userSessionId: sid,
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
