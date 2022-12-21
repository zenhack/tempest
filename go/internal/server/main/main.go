package servermain

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm-next/go/internal/server/session"
	"zenhack.net/go/util"
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
	srv := &server{
		log: lg,
		db:  db,
		containers: ContainerSet{
			containersByGrainId: make(map[string]*container.Container),
		},
		sessionStore: session.NewStore(util.Must(session.GetKeys())),
	}

	http.Handle("/", srv.Handler())
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
