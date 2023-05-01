// Package acme supports obtaining TLS certificates via ACME/Let's Encrypt.
package acme

import (
	"errors"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns"
	"zenhack.net/go/tempest/internal/server/settings"
	"zenhack.net/go/util/exn"
)

var (
	ErrNoProvider = errors.New("no dns01 challenge provider set")
)

func ConfigFromSettings(src settings.Source) (Config, error) {
	return exn.Try(func(throw exn.Thrower) Config {
		providerName := src.GetString("ACME_DNS_PROVIDER")
		if providerName == "" {
			throw(ErrNoProvider)
		}
		// XXX: this will always pull provider-specific options from the
		// environment, even if src is not settings.Environ. Maybe we should
		// pull this function out into a parameter we can dependency-inject:
		provider, err := dns.NewDNSChallengeProviderByName(providerName)
		throw(err)
		return Config{
			User: User{
				Email: src.GetString("ACME_EMAIL"),
			},
			Directory: src.GetString("ACME_DIRECTORY_URL"),
			Provider:  provider,
		}
	})
}

type Config struct {
	User
	Directory string
	Provider  challenge.Provider
}

type User struct {
	Email string
}
