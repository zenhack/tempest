package browsermain

import (
	"crypto/rand"
	"encoding/hex"

	"zenhack.net/go/util/maps"
	"zenhack.net/go/util/slices"
	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
)

func (m Model) View(msgs chan<- Msg) vdom.VNode {
	var grainNodes []vdom.VNode
	items := maps.Items(m.Grains)
	slices.SortOn(items, func(kv maps.KV[ID[Grain], Grain]) string {
		return kv.Value.Title
	})
	for _, kv := range items {
		grainNodes = append(grainNodes, viewGrain(msgs, kv.Key, kv.Value))
	}
	var content vdom.VNode
	if m.FocusedGrain == "" {
		content = vb.T("some text")
	} else {
		content = viewGrainIframe(
			m.ServerAddr,
			m.FocusedGrain,
			m.Grains[m.FocusedGrain],
			m.OpenGrains[m.FocusedGrain],
		)
	}
	return vb.H("body", nil, nil,
		vb.H("div", vb.A{"class": "main-ui"}, nil,
			vb.H("div", vb.A{"class": "main-ui__topbar"}, nil,
				vb.H("a", vb.A{"href": "/"}, nil, vb.T("Sandstorm")),
				vb.H("nav", nil, nil,
					vb.H("ul", nil, nil,
						vb.H("li", nil, nil,
							vb.H("a",
								vb.A{"href": "/login/dev"}, nil,
								vb.T("Login with dev account"),
							),
						),
					),
				),
			),
			vb.H("div", vb.A{"class": "main-ui__main"}, nil,
				vb.H("div", vb.A{"class": "main-ui__sidebar"}, nil,
					vb.H("p", nil, nil, vb.T("Sidebar")),
					vb.H("ul", nil, nil, grainNodes...),
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

func viewGrain(msgs chan<- Msg, id ID[Grain], grain Grain) vdom.VNode {
	onClick := func(vdom.Event) any {
		msgs <- func(m Model) Model {
			m.FocusedGrain = id
			_, ok := m.OpenGrains[id]
			if !ok {
				m.OpenGrains[id] = OpenGrain{
					DomainNonce: newDomainNonce(),
				}
			}
			return m
		}
		return nil
	}
	return vb.H("li", nil, nil,
		vb.H("a",
			vb.A{"href": "#/grain/" + string(id)},
			vb.E{"click": &onClick},
			vb.T(grain.Title),
		),
	)
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
