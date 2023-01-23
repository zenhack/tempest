package servermain

import (
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

	log.SetLevel(log.DebugLevel)
	lg := log.Log
	db := util.Must(database.Open())
	sessionStore := session.NewStore(util.Must(session.GetKeys()))
	srv := newServer(rootDomain, lg, db, sessionStore)
	defer srv.Release()

	http.Handle("/", srv.Handler())
	lg.Infof("Listening on %v", listenAddr)
	httpSrv := &http.Server{Addr: listenAddr}
	go monitorSignals(httpSrv)
	panic(httpSrv.ListenAndServe())
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
