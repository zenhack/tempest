package spk

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed sandstorm-pkgdef.capnp.template
	pkgdefTemplateSrc string

	pkgdefTemplate = template.Must(template.New("pkgdef").Parse(pkgdefTemplateSrc))
)
