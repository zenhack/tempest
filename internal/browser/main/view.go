package browsermain

import (
	"crypto/rand"
	"encoding/hex"

	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/maps"
	"zenhack.net/go/util/slices"
	"zenhack.net/go/vdom"
	"zenhack.net/go/vdom/builder"
)

var (
	h = builder.H
	t = builder.T
)

type (
	a = builder.A
	e = builder.E
)

var dummyNode = h("div", a{"class": "dummy-node"}, nil)

func (m Model) View(msgEvent func(Msg) vdom.EventHandler) vdom.VNode {
	content := dummyNode
	session, loginReady := m.LoginSession.Get()
	if !loginReady {
		content = t("Loading...")
	} else if session.Err() != nil {
		// TODO: deferrentiate between disconnects/failures. Or maybe just
		// tweak the API to return all this info in-band?
		content = viewLoginForm()
	} else {
		switch m.CurrentFocus {
		case FocusGrainList:
			kvs := maps.Items(m.Grains)
			slices.SortOn(kvs, func(kv maps.KV[types.GrainID, Grain]) string {
				return kv.Value.Title
			})
			var grainNodes []vdom.VNode
			for _, kv := range kvs {
				grainNodes = append(
					grainNodes,
					viewGrain(msgEvent, kv.Key, kv.Value),
				)
			}
			content = h("ul", nil, nil, grainNodes...)
		case FocusApps:
			content = m.viewApps(msgEvent)
		case FocusOpenGrain:
			if m.FocusedGrain == "" {
				content = t("Placeholder; select a grain.")
			}
		default:
			panic("Unknown focus value")
		}
	}
	keys := maps.Keys(m.OpenGrains)
	slices.SortOn(keys, func(k types.GrainID) string {
		return m.Grains[k].Title
	})
	var activeGrainNodes []vdom.VNode
	for _, k := range keys {
		activeGrainNodes = append(
			activeGrainNodes,
			viewOpenGrain(msgEvent, k, m.Grains[k], m.FocusedGrain == k),
		)
	}
	var iframes []vdom.VNode
	for _, id := range m.GrainDomOrder.Items {
		var vnode vdom.VNode
		if id == "" {
			vnode = dummyNode
		} else {
			vnode = viewGrainIframe(m, id)
		}
		iframes = append(iframes, vnode)
	}
	contentNodes := append([]vdom.VNode{content}, iframes...)

	mainUiNodes := []vdom.VNode{
		h("div", a{"class": "main-ui__main"}, nil,
			h("div", a{"class": "main-ui__sidebar"}, nil,
				h("h1", nil, nil,
					h("a",
						a{"href": "#"},
						e{"click": msgEvent(ChangeFocus{InitialFocus})},
						t("Tempest"),
					),
				),
				h("nav", nil, nil, h("ul", a{"class": "nav-links"}, nil,
					h("li", a{"class": "nav-link"}, nil,
						h("a",
							a{"href": "#/apps"},
							e{"click": msgEvent(ChangeFocus{FocusApps})},
							t("Apps"),
						),
					),
					h("li", a{"class": "nav-link"}, nil,
						h("a",
							a{"href": "#/grains"},
							e{"click": msgEvent(
								ChangeFocus{FocusGrainList},
							)},
							t("Grains"),
						),
					),
				)),
				h("h2", nil, nil, t("Grains")),
				h("nav", nil, nil,
					h("ul", a{"class": "nav-links"}, nil, activeGrainNodes...),
				),
			),
			h("div", a{"class": "main-ui__content"}, nil, contentNodes...),
		),
	}

	if m.Error != nil {
		mainUiNodes = append(
			mainUiNodes,
			h("div", a{"class": "error-notice"}, nil, t(m.Error.Error())),
		)
	}

	return h("body", nil, nil,
		h("div", a{"class": "main-ui"}, nil, mainUiNodes...),
	)
}

func (m Model) viewApps(msgEvent func(Msg) vdom.EventHandler) vdom.VNode {
	var items []vdom.VNode
	for id, pkg := range m.Packages {
		manifest, err := pkg.Manifest()
		if err != nil {
			println("manifest: " + err.Error())
			continue
		}
		l10nTitle, err := manifest.AppTitle()
		if err != nil {
			println("appTitle: " + err.Error())
			continue
		}
		title, err := l10nTitle.DefaultText()
		if err != nil {
			println("defaultText: " + err.Error())
			continue
		}
		link := h("a",
			a{"href": "#/grain/new"},
			e{
				"click": msgEvent(SpawnGrain{PkgID: id}),
			},
			t(title),
		)
		items = append(
			items,
			h("li", nil, nil, link),
		)
	}
	return h("ul", nil, nil, items...)
}

func newDomainNonce() string {
	var buf [16]byte
	rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

func viewLoginForm() vdom.VNode {
	return h("form", a{"action": "/login/dev", "method": "post"}, nil,
		h("label", a{"for": "name"}, nil,
			t("Dev account login"),
		),
		h("input", a{
			"name":        "name",
			"placeholder": "e.g. Alice Dev Admin",
		}, nil),
		h("button", a{"type": "submit"}, nil, t("Submit")),
	)
}

func viewOpenGrain(msgEvent func(Msg) vdom.EventHandler, id types.GrainID, grain Grain, isFocused bool) vdom.VNode {
	onClick := msgEvent(FocusGrain{ID: id})
	classes := "nav-link"
	if isFocused {
		classes += " nav-link--focused"
	}
	return h("li", a{"class": classes}, e{"click": onClick},
		h("a",
			a{"href": "#/grain/" + string(id)},
			e{"click": onClick},
			t(grain.Title),
		),
	)
}

func viewGrain(msgEvent func(Msg) vdom.EventHandler, id types.GrainID, grain Grain) vdom.VNode {
	// XXX
	return viewOpenGrain(msgEvent, id, grain, false)
}

func viewGrainIframe(m Model, id types.GrainID) vdom.VNode {
	grain := m.Grains[id]
	open := m.OpenGrains[id]

	grainUrl := m.ServerAddr.Subdomain("ui-" + open.DomainNonce)
	qv := grainUrl.Query()
	qv.Set("sandstorm-sid", grain.SessionToken)
	qv.Set("path", "/")
	grainUrl.Path = "/_sandstorm-init"
	grainUrl.RawQuery = qv.Encode()
	class := "grain-iframe"
	if m.CurrentFocus != FocusOpenGrain || m.FocusedGrain != id {
		class += " grain-iframe--inactive"
	}
	return h("iframe", a{
		"src":   grainUrl.String(),
		"class": class,
	}, nil)
}
