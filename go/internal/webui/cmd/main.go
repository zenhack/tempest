package main

import (
	"syscall/js"

	"zenhack.net/go/sandstorm-next/go/internal/webui/render"
	renderjs "zenhack.net/go/sandstorm-next/go/internal/webui/render/js"
)

func view(r render.Renderer) {
	host := js.Global().Get("window").Get("location").Get("host").String()
	r.E("div", render.A{"class": "main-ui"}, func(r render.Renderer) {
		r.E("div", render.A{"class": "main-ui__topbar"}, func(r render.Renderer) {
			r.E("p", nil, func(r render.Renderer) { r.T("Topbar") })
		})
		r.E("div", render.A{"class": "main-ui__main"}, func(r render.Renderer) {
			r.E("div", render.A{"class": "main-ui__sidebar"}, func(r render.Renderer) {
				r.E("p", nil, func(r render.Renderer) { r.T("Sidebar") })
			})
			r.E("iframe", render.A{
				"src":   "//grain." + host,
				"class": "main-ui__grain-iframe",
			}, nil)
		})
	})
}

func main() {
	body := js.Global().
		Get("document").
		Call("getElementsByTagName", "body").
		Index(0)
	view(renderjs.New(body))
}
