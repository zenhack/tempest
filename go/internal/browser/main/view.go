package browsermain

import (
	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
)

func view(m Model) vdom.VNode {
	return vb.H("body", nil, nil,
		vb.H("div", vb.A{"class": "main-ui"}, nil,
			vb.H("div", vb.A{"class": "main-ui__topbar"}, nil,
				vb.H("p", nil, nil, vb.T("Topbar")),
			),
			vb.H("div", vb.A{"class": "main-ui__main"}, nil,
				vb.H("div", vb.A{"class": "main-ui__sidebar"}, nil,
					vb.H("p", nil, nil, vb.T("Sidebar")),
				),
				vb.H("iframe", vb.A{
					"src":   "//grain." + m.Host,
					"class": "main-ui__grain-iframe",
				}, nil),
			),
		),
	)
}
