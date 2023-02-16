package servermain

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/apex/log"
	"zenhack.net/go/tempest/internal/database"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util"
)

func defaultTo(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func Main() {
	rootDomain := defaultTo(os.Getenv("ROOT_DOMAIN"), "local.sandstorm.io:8000")
	var port string
	if strings.Contains(rootDomain, ":") {
		var err error
		_, port, err = net.SplitHostPort(rootDomain)
		util.Chkfatal(err)
	} else {
		port = "80"
	}
	listenAddr := ":" + port

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
