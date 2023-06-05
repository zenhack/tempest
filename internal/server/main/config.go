package servermain

import (
	"net"
	"net/smtp"
	"net/url"
	"os"
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

func SMTPConfigFromSettings(src settings.Source) (cfg SMTPConfig) {
	sh := src.GetString("SMTP_HOST")
	sprt := os.Getenv("SMTP_PORT")
	su := src.GetString("SMTP_USERNAME")
	sp := src.GetString("SMTP_PASSWORD")

	if sh == "" || sprt == "" || su == "" || sp == "" {
		return
	}

	return SMTPConfig{
		Host:     sh,
		Port:     strconv.Itoa(int(src.GetUint16("SMTP_PORT"))),
		Username: su,
		Password: sp,
	}
}

func HTTPConfigFromSettings(lg *slog.Logger, src settings.Source) HTTPConfig {
	baseURLStr := src.GetString("BASE_URL")
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
		Port:       src.GetString("HTTP_PORT"),
		TLSPort:    src.GetString("HTTPS_PORT"),
		CertFile:   src.GetString("HTTPS_CERT_FILE"),
		KeyFile:    src.GetString("HTTPS_KEY_FILE"),
	}
	return cfg
}

func ConfigFromSettings(lg *slog.Logger, src settings.Source) Config {
	return Config{
		HTTP: HTTPConfigFromSettings(lg, src),
		SMTP: SMTPConfigFromSettings(src),
	}
}
