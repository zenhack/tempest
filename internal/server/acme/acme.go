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

func ConfigFromSettings() (Config, error) {
	return exn.Try(func(throw exn.Thrower) Config {
		providerName := settings.GetString("ACME_DNS_PROVIDER")
		if providerName == "" {
			throw(ErrNoProvider)
		}
		provider, err := dns.NewDNSChallengeProviderByName(providerName)
		throw(err)
		return Config{
			User: User{
				Email: settings.GetString("ACME_EMAIL"),
			},
			Directory: settings.GetString("ACME_DIRECTORY_URL"),
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
