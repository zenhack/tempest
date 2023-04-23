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
	HTTP HTTPConfig
	SMTP SMTPConfig
}

type HTTPConfig struct {
	RootDomain        string // Main Tempest domain name
	Port              string
	TLSPort           string
	CertFile, KeyFile string
	DefaultTLS        bool
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
		Host:     settings.GetString("SMTP_HOST"),
		Port:     strconv.Itoa(int(settings.GetUint16("SMTP_PORT"))),
		Username: settings.GetString("SMTP_USERNAME"),
		Password: settings.GetString("SMTP_PASSWORD"),
	}
}

func HTTPConfigFromSettings(lg *slog.Logger) HTTPConfig {
	baseURLStr := settings.GetString("BASE_URL")
	baseURL := util.Must(url.Parse(baseURLStr))
	if baseURL.Scheme != "http" && baseURL.Scheme != "https" {
		logging.Panic(lg, "parsing BASE_URL: must use http(s) scheme")
	}
	if baseURL.Path != "" {
		logging.Panic(lg, "parsing BASE_URL: must not have a path")
	}
	cfg := HTTPConfig{
		DefaultTLS: baseURL.Scheme == "https",
		RootDomain: baseURL.Host,
		Port:       settings.GetString("HTTP_PORT"),
		TLSPort:    settings.GetString("HTTPS_PORT"),
		CertFile:   settings.GetString("HTTPS_CERT_FILE"),
		KeyFile:    settings.GetString("HTTPS_KEY_FILE"),
	}
	return cfg
}

func ConfigFromSettings(lg *slog.Logger) Config {
	return Config{
		HTTP: HTTPConfigFromSettings(lg),
		SMTP: SMTPConfigFromSettings(),
	}
}
