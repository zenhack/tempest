package render

type A map[string]string

type Renderer interface {
	E(tag string, attrs A, body func(Renderer))
	T(text string)
}
