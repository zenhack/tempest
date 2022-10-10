package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"zenhack.net/go/sandstorm-next/go/internal/container"
	"zenhack.net/go/sandstorm-next/go/internal/util"
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

func main() {
	ctx := context.Background()
	c, err := container.StartDummy(ctx)
	util.Chkfatal(err)
	defer c.Release()

	r := mux.NewRouter()

	r.Host("grain." + rootDomain).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ServeApp(c, w, req)
	})

	http.Handle("/", r)
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
