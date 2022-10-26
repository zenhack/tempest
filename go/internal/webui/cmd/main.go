package main

import (
	"syscall/js"

	"zenhack.net/go/sandstorm-next/go/internal/webui/render"
	renderjs "zenhack.net/go/sandstorm-next/go/internal/webui/render/js"
)

func view(r render.Renderer) {
	r.E("p", render.A{"class": "foo"}, func(r render.Renderer) {
		r.T("Hello, World!")
	})
}

func main() {
	body := js.Global().
		Get("document").
		Call("getElementsByTagName", "body").
		Index(0)
	view(renderjs.New(body))
}
