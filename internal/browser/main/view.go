package browsermain

import (
	"crypto/rand"
	"encoding/hex"

	"zenhack.net/go/util/maps"
	"zenhack.net/go/util/slices"
	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
)

func msgEvent(msgs chan<- Msg, msg Msg) vdom.EventHandler {
	ret := func(vdom.Event) any {
		msgs <- msg
		return nil
	}
	return &ret
}

func (m Model) View(msgs chan<- Msg) vdom.VNode {
	var (
		content vdom.VNode
	)
	session, loginReady := m.LoginSession.Get()
	if !loginReady {
		content = vb.T("Loading...")
	} else if session.Err() != nil {
		// TODO: deferrentiate between disconnects/failures. Or maybe just
		// tweak the API to return all this info in-band?
		content = viewLoginForm()
	} else {
		switch m.CurrentFocus {
		case FocusGrainList:
			kvs := maps.Items(m.Grains)
			slices.SortOn(kvs, func(kv maps.KV[ID[Grain], Grain]) string {
				return kv.Value.Title
			})
			var grainNodes []vdom.VNode
			for _, kv := range kvs {
				grainNodes = append(
					grainNodes,
					viewGrain(msgs, kv.Key, kv.Value),
				)
			}
			content = vb.H("ul", nil, nil, grainNodes...)
		case FocusOpenGrain:
			if m.FocusedGrain == "" {
				content = vb.T("Placeholder; select a grain.")
			} else {
				content = viewGrainIframe(
					m.ServerAddr,
					m.FocusedGrain,
					m.Grains[m.FocusedGrain],
					m.OpenGrains[m.FocusedGrain],
				)
			}
		default:
			panic("Unknown focus value")
		}
	}
	keys := maps.Keys(m.OpenGrains)
	slices.SortOn(keys, func(k ID[Grain]) string {
		return m.Grains[k].Title
	})
	var activeGrainNodes []vdom.VNode
	for _, k := range keys {
		activeGrainNodes = append(
			activeGrainNodes,
			viewOpenGrain(msgs, k, m.Grains[k], m.FocusedGrain == k),
		)
	}
	return vb.H("body", nil, nil,
		vb.H("div", vb.A{"class": "main-ui"}, nil,
			vb.H("div", vb.A{"class": "main-ui__main"}, nil,
				vb.H("div", vb.A{"class": "main-ui__sidebar"}, nil,
					vb.H("h1", nil, nil,
						vb.H("a", vb.A{"href": "/"}, nil, vb.T("Tempest")),
					),
					vb.H("nav", nil, nil, vb.H("ul", vb.A{"class": "nav-links"}, nil,
						vb.H("li", vb.A{"class": "nav-link"}, nil,
							vb.H("a",
								vb.A{"href": "#/grains"},
								vb.E{"click": msgEvent(
									msgs,
									ChangeFocus{FocusGrainList},
								)},
								vb.T("Grains"),
							),
						),
					)),
					vb.H("h2", nil, nil, vb.T("Grains")),
					vb.H("nav", nil, nil,
						vb.H("ul", vb.A{"class": "nav-links"}, nil, activeGrainNodes...),
					),
				),
				content,
			),
		),
	)
}

func newDomainNonce() string {
	var buf [16]byte
	rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

func viewLoginForm() vdom.VNode {
	return vb.H("form", vb.A{"action": "/login/dev", "method": "post"}, nil,
		vb.H("label", vb.A{"for": "name"}, nil,
			vb.T("Dev account login"),
		),
		vb.H("input", vb.A{
			"name":        "name",
			"placeholder": "e.g. Alice Dev Admin",
		}, nil),
		vb.H("button", vb.A{"type": "submit"}, nil, vb.T("Submit")),
	)
}

func viewOpenGrain(msgs chan<- Msg, id ID[Grain], grain Grain, isFocused bool) vdom.VNode {
	onClick := msgEvent(msgs, FocusGrain{Id: id})
	classes := "nav-link"
	if isFocused {
		classes += " nav-link--focused"
	}
	return vb.H("li", vb.A{"class": classes}, vb.E{"click": onClick},
		vb.H("a",
			vb.A{"href": "#/grain/" + string(id)},
			vb.E{"click": onClick},
			vb.T(grain.Title),
		),
	)
}

func viewGrain(msgs chan<- Msg, id ID[Grain], grain Grain) vdom.VNode {
	// XXX
	return viewOpenGrain(msgs, id, grain, false)
}

func viewGrainIframe(addr ServerAddr, id ID[Grain], grain Grain, open OpenGrain) vdom.VNode {
	grainUrl := addr.Subdomain("ui-" + open.DomainNonce)
	qv := grainUrl.Query()
	qv.Set("sandstorm-sid", grain.SessionToken)
	qv.Set("path", "/")
	grainUrl.Path = "/_sandstorm-init"
	grainUrl.RawQuery = qv.Encode()
	return vb.H("iframe", vb.A{"src": grainUrl.String(),
		"class": "main-ui__grain-iframe",
	}, nil)
}
