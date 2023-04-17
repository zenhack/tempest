package servermain

import (
	"net"
	"net/smtp"
	"net/url"
	"strconv"

	"golang.org/x/exp/slog"
	"zenhack.net/go/tempest/internal/server/logging"
	"zenhack.net/go/tempest/internal/server/settings"
	"zenhack.net/go/util"
)

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

func SMTPConfigFromSettings() SMTPConfig {
	return SMTPConfig{
		Host:     settings.GetString("smtp.host"),
		Port:     strconv.Itoa(int(settings.GetUint16("smtp.port"))),
		Username: settings.GetString("smtp.username"),
		Password: settings.GetString("smtp.password"),
	}
}

func ConfigFromSettings(lg *slog.Logger) Config {
	baseURLStr := settings.GetString("base_url")
	baseURL := util.Must(url.Parse(baseURLStr))
	if baseURL.Scheme != "http" {
		logging.Panic(lg, "parsing BASE_URL: must use http scheme")
	}
	if baseURL.Path != "" {
		logging.Panic(lg, "parsing BASE_URL: must not have a path")
	}
	cfg := Config{
		rootDomain: baseURL.Host,
		smtp:       SMTPConfigFromSettings(),
	}
	var err error
	_, cfg.listenPort, err = net.SplitHostPort(cfg.rootDomain)
	util.Chkfatal(err)
	if cfg.listenPort == "" {
		cfg.listenPort = "80"
	}
	return cfg
}
