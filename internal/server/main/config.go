package servermain

import (
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
	rootDomain string // Main Tempest domain name

	smtp SMTPConfig
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (c SMTPConfig) getAuth() smtp.Auth {
	return smtp.PlainAuth(
		c.Username,
		c.Username,
		c.Password,
		c.Host,
	)
}

func (c SMTPConfig) SendMail(to []string, msg []byte) error {
	return smtp.SendMail(
		net.JoinHostPort(c.Host, c.Port),
		c.getAuth(),
		c.Username,
		to,
		msg,
	)
}

func SMTPConfigFromEnv() SMTPConfig {
	return SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
}

func ConfigFromEnv(lg *slog.Logger) Config {
	baseURLStr := defaultTo(os.Getenv("BASE_URL"), "http://local.sandstorm.io:8000")
	baseURL := util.Must(url.Parse(baseURLStr))
	if baseURL.Scheme != "http" {
		logging.Panic(lg, "parsing BASE_URL: must use http scheme")
	}
	if baseURL.Path != "" {
		logging.Panic(lg, "parsing BASE_URL: must not have a path")
	}
	cfg := Config{
		rootDomain: baseURL.Host,
		smtp:       SMTPConfigFromEnv(),
	}
	var err error
	_, cfg.listenPort, err = net.SplitHostPort(cfg.rootDomain)
	util.Chkfatal(err)
	if cfg.listenPort == "" {
		cfg.listenPort = "80"
	}
	return cfg
}
