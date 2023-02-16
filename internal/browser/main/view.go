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
		session, ok := m.LoginSession.Get()
		if !ok {
			content = vb.T("Loading...")
		} else if session.Err() != nil {
			// TODO: deferrentiate between disconnects/failures. Or maybe just
			// tweak the API to return all this info in-band?
			content = viewLoginForm()
		} else {
			content = vb.T("Placeholder; select a grain.")
		}
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
			vb.H("div", vb.A{"class": "main-ui__main"}, nil,
				vb.H("div", vb.A{"class": "main-ui__sidebar"}, nil,
					vb.H("a", vb.A{"href": "/"}, nil, vb.T("Tempest")),
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

func viewGrain(msgs chan<- Msg, id ID[Grain], grain Grain) vdom.VNode {
	onClick := func(vdom.Event) any {
		msgs <- FocusGrain{Id: id}
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
