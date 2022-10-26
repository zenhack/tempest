package js

import (
	"syscall/js"

	"zenhack.net/go/sandstorm-next/go/internal/webui/render"
)

func New(element js.Value) render.Renderer {
	return jsRenderer{
		document: js.Global().Get("document"),
		root:     element,
	}
}

type jsRenderer struct {
	document, root js.Value
}

func (r jsRenderer) E(tag string, attrs render.A, body func(r render.Renderer)) {
	if attrs == nil {
		attrs = render.A{}
	}
	e := r.document.Call("createElement", tag, attrs)
	childR := r
	childR.root = e
	body(childR)
	r.root.Call("appendChild", e)
}

func (r jsRenderer) T(text string) {
	n := r.document.Call("createTextNode", text)
	r.root.Call("appendChild", n)
}
