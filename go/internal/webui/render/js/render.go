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
	e := r.document.Call("createElement", tag)
	if attrs != nil {
		eAttrs := e.Get("attributes")
		for k, v := range attrs {
			attr := r.document.Call("createAttribute", k)
			attr.Set("value", v)
			eAttrs.Call("setNamedItem", attr)
		}
	}
	if body != nil {
		childR := r
		childR.root = e
		body(childR)
	}
	r.root.Call("appendChild", e)
}

func (r jsRenderer) T(text string) {
	n := r.document.Call("createTextNode", text)
	r.root.Call("appendChild", n)
}
