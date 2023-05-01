package servermain

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/logging"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/tempest/internal/server/settings"
	"zenhack.net/go/util"
)

func Main() {
	initStorage()
	lg := logging.NewLogger()
	cfg := ConfigFromSettings(lg, settings.Environ)
	httpAddr := ":" + cfg.HTTP.Port
	httpsAddr := ":" + cfg.HTTP.TLSPort
	db := util.Must(database.Open())
	sessionStore := session.NewStore(util.Must(session.GetKeys()))
	srv := newServer(cfg, lg, db, sessionStore)
	defer srv.Release()

	if cfg.HTTP.KeyFile != "" {
		fi, err := os.Lstat(cfg.HTTP.KeyFile)
		util.Chkfatal(err)
		mode := fi.Mode()
		if mode&0077 != 0 {
			logging.Panic(lg,
				"https key file has incorrect permisisons; must have no access by group or other",
				"permissions", fmt.Sprintf("%o", mode&0777),
			)
		}
	}

	http.Handle("/", srv.Handler())
	lg.Info("Listening",
		"root-domain", cfg.HTTP.RootDomain,
		"http-addr", httpAddr,
		"https-addr", httpsAddr,
	)
	httpSrv := &http.Server{Addr: httpAddr}
	go monitorSignals(httpSrv)

	// We can't just use util.Chkfatal for the below, becasue
	// they *always* return an error -- we have to check which
	// error:
	checkServerError := func(err error) {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

	if cfg.HTTP.CertFile != "" && cfg.HTTP.KeyFile != "" {
		l, err := net.Listen("tcp", httpsAddr)
		util.Chkfatal(err)
		go func() {
			checkServerError(httpSrv.ServeTLS(
				l,
				cfg.HTTP.CertFile,
				cfg.HTTP.KeyFile,
			))
		}()
	}
	checkServerError(httpSrv.ListenAndServe())
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
