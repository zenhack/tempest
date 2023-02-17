package browsermain

import (
	"net/url"
	"syscall/js"

	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/capnp/util"
	"zenhack.net/go/util/maybe"
	"zenhack.net/go/util/orerr"
)

func initModel() Model {
	loc := js.Global().Get("window").Get("location")
	return Model{
		CurrentFocus: InitialFocus,
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

type Focus int

const (
	FocusGrainList Focus = iota
	FocusOpenGrain
)

const InitialFocus = FocusGrainList

type Model struct {
	ServerAddr   ServerAddr
	CurrentFocus Focus
	FocusedGrain ID[Grain] // ID for the currently focused grain

	Grains     map[ID[Grain]]Grain
	OpenGrains map[ID[Grain]]OpenGrain

	// TODO(cleanup): factor this out into its own type and clean up related code.
	GrainDomOrder struct {
		IDGen idgen
		Nodes []ID[Grain]
	}

	LoginSession maybe.T[orerr.T[external.LoginSession]]
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

	DomIndex int
}
