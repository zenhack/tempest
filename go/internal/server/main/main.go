package servermain

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/apex/log"
	"zenhack.net/go/tempest/go/internal/database"
	"zenhack.net/go/tempest/go/internal/server/session"
	"zenhack.net/go/util"
)

func defaultTo(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func Main() {
	rootDomain := defaultTo(os.Getenv("ROOT_DOMAIN"), "local.sandstorm.io")
	listenAddr := defaultTo(os.Getenv("LISTEN_ADDR"), ":8000")

	lg := log.Log
	db := util.Must(database.Open())
	sessionStore := session.NewStore(util.Must(session.GetKeys()))
	srv := newServer(rootDomain, lg, db, sessionStore)
	defer srv.Release()

	http.Handle("/", srv.Handler())
	lg.Infof("Listening on %v", listenAddr)
	httpSrv := &http.Server{Addr: listenAddr}
	go monitorSignals(httpSrv)
	fmt.Println(httpSrv.ListenAndServe())
	os.Exit(1) // Only reachable if ListenAndServe() errors.
}

func monitorSignals(srv *http.Server) {
	defer srv.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		// Signals that would normally kill us. Instead, stop the server
		// and let main() do shutdown.
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	defer signal.Stop(sigs)
	<-sigs
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
