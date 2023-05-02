// Package acme supports obtaining TLS certificates via ACME/Let's Encrypt.
package acme

import (
	"crypto"
	"errors"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns"
	"github.com/go-acme/lego/v4/registration"
	"zenhack.net/go/tempest/internal/server/settings"
	"zenhack.net/go/util/exn"
)

var (
	ErrNoProvider = errors.New("no dns01 challenge provider set")
)

func ConfigFromSettings(src settings.Source) (*Config, error) {
	return exn.Try(func(throw exn.Thrower) *Config {
		providerName := src.GetString("ACME_DNS_PROVIDER")
		if providerName == "" {
			throw(ErrNoProvider)
		}
		// XXX: this will always pull provider-specific options from the
		// environment, even if src is not settings.Environ. Maybe we should
		// pull this function out into a parameter we can dependency-inject:
		provider, err := dns.NewDNSChallengeProviderByName(providerName)
		throw(err)
		return &Config{
			User: User{
				Email: src.GetString("ACME_EMAIL"),
				// TODO: fill in other fields.
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

func (c *Config) ToLego() *lego.Config {
	ret := lego.NewConfig(&c.User)
	ret.CADirURL = c.Directory
	return ret
}

func (c *Config) ToClient() (*lego.Client, error) {
	client, err := lego.NewClient(c.ToLego())
	if err != nil {
		return nil, err
	}
	client.Challenge.SetDNS01Provider(c.Provider)
	return client, nil
}

type User struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRegistration() *registration.Resource {
	return u.Registration
}

func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}
