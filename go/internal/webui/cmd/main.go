package main

import (
	"syscall/js"

	"zenhack.net/go/sandstorm-next/go/internal/webui/render"
	renderjs "zenhack.net/go/sandstorm-next/go/internal/webui/render/js"
)

func view(r render.Renderer) {
	host := js.Global().Get("window").Get("location").Get("host").String()
	r.E("p", render.A{"class": "foo"}, func(r render.Renderer) {
		r.T("Hello, World!")
	})
	r.E("iframe", render.A{
		"src":   "//grain." + host,
		"style": "width: 100%; height: 100%;",
	}, nil)
}

func main() {
	body := js.Global().
		Get("document").
		Call("getElementsByTagName", "body").
		Index(0)
	view(renderjs.New(body))
}
