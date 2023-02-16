package embed

import (
	"embed"
)

//go:embed *.wasm
//go:embed *.js
//go:embed *.html
//go:embed *.css
//go:embed _dev/*
var Content embed.FS
