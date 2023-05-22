package browsermain

import (
	"net/url"
	"syscall/js"

	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/browser/intl"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/maybe"
	"zenhack.net/go/util/orerr"
	"zenhack.net/go/util/slices/poolslice"
)

type Model struct {
	L10N intl.L10N

	ServerAddr   ServerAddr
	CurrentFocus Focus
	FocusedGrain types.GrainID // ID for the currently focused grain

	// Errors to display to the user. As the UI evolves, we will
	// probably want more out of this, but YAGNI.
	Errors []error

	Grains     map[types.GrainID]Grain
	OpenGrains map[types.GrainID]OpenGrain
	Packages   map[types.ID[external.Package]]external.Package

	// Keeps track of the order we need to display grain iframes in.
	// Grain iframes must never change order or be detached from the
	// DOM, or they will reload the page within them, losing state.
	// So we keep them in a stable order, rendering empty slots
	// as dummyNode, and hiding everything but the active grain
	// with CSS (display: none).
	GrainDomOrder poolslice.PoolSlice[types.GrainID]

	API           external.ExternalApi
	LoginSessions maybe.Maybe[orerr.OrErr[Sessions]]

	LoginForm LoginForm
}

type Sessions struct {
	Visitor external.VisitorSession
	User    external.UserSession
}

type ServerAddr struct {
	Host string
	TLS  bool
}

type Focus int

const (
	FocusGrainList Focus = iota
	FocusApps
	FocusOpenGrain
	FocusShareGrain

	InitialFocus = FocusGrainList
)

type Grain struct {
	Title        string
	SessionToken string
	Subdomain    string
	Controller   external.UiView_Controller
}

type OpenGrain struct {
	DomIndex int
}

func initModel(api external.ExternalApi) Model {
	loc := js.Global().Get("window").Get("location")
	return Model{
		CurrentFocus: InitialFocus,
		ServerAddr: ServerAddr{
			TLS:  loc.Get("protocol").String() == "https:",
			Host: loc.Get("host").String(),
		},
		Grains:     make(map[types.GrainID]Grain),
		OpenGrains: make(map[types.GrainID]OpenGrain),
		Packages:   make(map[types.ID[external.Package]]external.Package),
		API:        api,
	}
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
