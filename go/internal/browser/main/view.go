package browsermain

import (
	"crypto/rand"
	"encoding/hex"

	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
)

func (m Model) View(msgs chan<- Msg) vdom.VNode {
	var grainNodes []vdom.VNode
	for k, v := range m.Grains {
		grainNodes = append(grainNodes, viewGrain(msgs, k, v))
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
				vb.H("p", nil, nil, vb.T("Topbar")),
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
			vb.A{"href": "#"},
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
