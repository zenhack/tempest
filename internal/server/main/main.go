package servermain

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/logging"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util"
)

func defaultTo(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

type Config struct {
	listenPort string
	rootDomain string
}

func ConfigFromEnv(lg *slog.Logger) Config {
	baseURLStr := defaultTo(os.Getenv("BASE_URL"), "http://local.sandstorm.io:8000")
	baseURL := util.Must(url.Parse(baseURLStr))
	if baseURL.Scheme != "http" {
		logging.Panic(lg, "parsing BASE_URL", errors.New("must use http: scheme"))
	}
	if baseURL.Path != "" {
		logging.Panic(lg, "parsing BASE_URL", errors.New("must not have a path"))
	}
	cfg := Config{
		rootDomain: baseURL.Host,
	}
	var err error
	_, cfg.listenPort, err = net.SplitHostPort(cfg.rootDomain)
	util.Chkfatal(err)
	if cfg.listenPort == "" {
		cfg.listenPort = "80"
	}
	return cfg
}

func Main() {
	lg := logging.NewLogger()
	cfg := ConfigFromEnv(lg)
	listenAddr := ":" + cfg.listenPort
	db := util.Must(database.Open())
	sessionStore := session.NewStore(util.Must(session.GetKeys()))
	srv := newServer(cfg.rootDomain, lg, db, sessionStore)
	defer srv.Release()

	http.Handle("/", srv.Handler())
	lg.Info("Listening",
		"root-domain", cfg.rootDomain,
		"listen-addr", listenAddr,
	)
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
