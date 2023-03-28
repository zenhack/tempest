package servermain

import (
	"errors"
	"net"
	"net/smtp"
	"net/url"
	"os"

	"golang.org/x/exp/slog"
	"zenhack.net/go/tempest/internal/server/logging"
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

	smtpAddr string
	smtpAuth smtp.Auth
	smtpFrom string
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
		cfg.smtpAddr = os.Getenv("SMTP_ADDR")
		cfg.smtpAuth = smtpAuthFromEnv()
		cfg.smtpFrom = os.Getenv("SMTP_FROM")
	}
	return cfg
}

func smtpAuthFromEnv() smtp.Auth {
	return smtp.PlainAuth(
		os.Getenv("SMTP_AUTH_IDENTITY"),
		os.Getenv("SMTP_AUTH_USERNAME"),
		os.Getenv("SMTP_AUTH_PASSWORD"),
		os.Getenv("SMTP_AUTH_HOST"),
	)
}
