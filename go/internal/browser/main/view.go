package browsermain

import (
	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
)

func (m Model) View() vdom.VNode {
	var grainNodes []vdom.VNode
	for k, v := range m.Grains {
		grainNodes = append(grainNodes, viewGrain(k, v))
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
				vb.H("iframe", vb.A{
					"src":   "//grain." + m.Host,
					"class": "main-ui__grain-iframe",
				}, nil),
			),
		),
	)
}

func viewGrain(id string, grain Grain) vdom.VNode {
	return vb.H("li", nil, nil, vb.T(grain.Title+" ("+id+")"))
}
