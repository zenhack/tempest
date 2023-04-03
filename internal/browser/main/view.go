package browsermain

import (
	"crypto/rand"
	"encoding/hex"

	"zenhack.net/go/tempest/internal/browser/intl"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/maps"
	"zenhack.net/go/util/slices"
	"zenhack.net/go/vdom"
	"zenhack.net/go/vdom/builder"
	"zenhack.net/go/vdom/events"
	"zenhack.net/go/vdom/tea"
)

var (
	h = builder.H
)

type (
	a = builder.A
	e = builder.E
)

func t(l10n intl.L10N, f intl.L10NString, args ...string) vdom.VNode {
	return builder.T(l10n.Fmt(f, args...))
}

var dummyNode = h("div", a{"class": "dummy-node"}, nil)

func (m Model) View(ms tea.MessageSender[Model]) vdom.VNode {
	content := dummyNode
	session, loginReady := m.LoginSession.Get()
	if !loginReady {
		content = t(m.L10N, "Loading...")
	} else if session.Err() != nil {
		// TODO: deferrentiate between disconnects/failures. Or maybe just
		// tweak the API to return all this info in-band?
		content = viewLoginForm(m.L10N, ms)
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
					viewGrain(ms, kv.Key, kv.Value),
				)
			}
			content = h("ul", nil, nil, grainNodes...)
		case FocusApps:
			content = m.viewApps(ms)
		case FocusOpenGrain:
			if m.FocusedGrain == "" {
				content = t(m.L10N, "Placeholder; select a grain.")
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
			viewOpenGrain(m.L10N, ms, k, m.Grains[k], m.FocusedGrain == k),
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
						e{"click": ms.Event(ChangeFocus{InitialFocus})},
						t(m.L10N, "Tempest"),
					),
				),
				h("nav", nil, nil, h("ul", a{"class": "nav-links"}, nil,
					h("li", a{"class": "nav-link"}, nil,
						h("a",
							a{"href": "#/apps"},
							e{"click": ms.Event(ChangeFocus{FocusApps})},
							t(m.L10N, "Apps"),
						),
					),
					h("li", a{"class": "nav-link"}, nil,
						h("a",
							a{"href": "#/grains"},
							e{"click": ms.Event(
								ChangeFocus{FocusGrainList},
							)},
							t(m.L10N, "Grains"),
						),
					),
				)),
				h("h2", nil, nil, t(m.L10N, "Grains")),
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
			// TODO: figure out how translating the error should work.
			h("div", a{"class": "error-notice"}, nil, builder.T(m.Error.Error())),
		)
	}

	return h("body", nil, nil,
		h("div", a{"class": "main-ui"}, nil, mainUiNodes...),
	)
}

func (m Model) viewApps(ms tea.MessageSender[Model]) vdom.VNode {
	var appItems []vdom.VNode
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
		actions, err := manifest.Actions()
		if err != nil {
			println("actions: " + err.Error())
			continue
		}
		var links []vdom.VNode
		for i := 0; i < actions.Len(); i++ {
			action := actions.At(i)
			nounPhrasel10n, err := action.NounPhrase()
			if err != nil {
				println("nounPhrase: " + err.Error())
				continue
			}
			nounPhrase, err := nounPhrasel10n.DefaultText()
			if err != nil {
				println("nounPhrase.defaultText: " + err.Error())
				continue
			}

			links = append(
				links,
				h("li", nil, nil, h("a",
					a{"href": "#/grain/new"},
					e{
						"click": ms.Event(SpawnGrain{
							Index: i,
							PkgID: id,
						}),
					},
					// TODO: figure out how translation
					// should work for app-provided strings.
					builder.T(nounPhrase),
				)),
			)
		}
		appItems = append(
			appItems,
			h("li", nil, nil,
				// TODO: figure out how translation
				// should work for app-provided strings.
				builder.T(title),
				h("ul", nil, nil, links...),
			),
		)
	}
	return h("ul", nil, nil, appItems...)
}

func newDomainNonce() string {
	var buf [16]byte
	rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

func viewLoginForm(l10n intl.L10N, ms tea.MessageSender[Model]) vdom.VNode {
	return h("div", nil, nil,
		h("form", a{"action": "/login/dev", "method": "post"}, nil,
			h("label", a{"for": "name"}, nil,
				t(l10n, "Dev account login"),
			),
			h("input", a{
				"name":        "name",
				"placeholder": "e.g. Alice Dev Admin",
			}, nil),
			h("button", a{"type": "submit"}, nil, t(l10n, "Submit")),
		),
		h("div", nil, nil,
			h("label", a{"for": "address"}, nil,
				t(l10n, "Email address for login"),
			),
			h("input", a{
				"name":        "address",
				"placeholder": "e.g. alice@example.com",
			}, e{
				"input": events.OnInput(func(value string) {
					ms.Send(EditEmailLogin{NewValue: value})
				}),
			}),
			h("button", nil, nil, t(l10n, "Submit")),
		),
	)
}

func viewOpenGrain(l10n intl.L10N, ms tea.MessageSender[Model], id types.GrainID, grain Grain, isFocused bool) vdom.VNode {
	focusGrain := ms.Event(FocusGrain{ID: id})
	classes := "open-grain-tab"
	if isFocused {
		classes += " open-grain-tab--focused"
	}
	return h("li", a{"class": classes}, nil,
		h("a",
			a{"href": "#/grain/" + string(id)},
			e{"click": focusGrain},
			builder.T(grain.Title),
		),
		h("button",
			a{"class": "close-button"},
			e{"click": ms.Event(CloseGrain{ID: id})},
			t(l10n, "Close Grain"),
		),
	)
}

func viewGrain(ms tea.MessageSender[Model], id types.GrainID, grain Grain) vdom.VNode {
	return h("li", a{"class": "nav-link"}, nil,
		h("a",
			a{"href": "#/grain/" + string(id)},
			e{"click": ms.Event(FocusGrain{ID: id})},
			builder.T(grain.Title),
		),
	)
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
