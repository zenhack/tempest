package embed

import (
	"embed"
)

//go:embed webui.wasm
//go:embed wasm_exec.js
//go:embed index.html
var Content embed.FS
