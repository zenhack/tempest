package browsermain

import (
	"net/url"
	"syscall/js"

	"zenhack.net/go/sandstorm/capnp/util"
)

func initModel() Model {
	loc := js.Global().Get("window").Get("location")
	return Model{
		ServerAddr: ServerAddr{
			TLS:  loc.Get("protocol").String() == "https:",
			Host: loc.Get("host").String(),
		},
		Grains:     make(map[ID[Grain]]Grain),
		OpenGrains: make(map[ID[Grain]]OpenGrain),
	}
}

type ID[T any] string

type ServerAddr struct {
	Host string
	TLS  bool
}

func (sa ServerAddr) Root() url.URL {
	ret := url.URL{
		Host: sa.Host,
	}
	if sa.TLS {
		ret.Scheme = "https"
	} else {
		ret.Scheme = "http"
	}
	return ret
}

func (sa ServerAddr) Subdomain(s string) url.URL {
	root := sa.Root()
	root.Host = s + "." + sa.Host
	return root
}

type Model struct {
	ServerAddr   ServerAddr
	FocusedGrain ID[Grain] // ID for the currently focused grain

	Grains     map[ID[Grain]]Grain
	OpenGrains map[ID[Grain]]OpenGrain
}

type Grain struct {
	Title        string
	SessionToken string
	Handle       util.Handle
}

type OpenGrain struct {
	// The random part of the subdomain for this grain, i.e. the grain is
	// served at http(s)://ui-${DomainNonce}.sandstorm.example.org/
	DomainNonce string
}
