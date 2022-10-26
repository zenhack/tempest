package embed

import (
	_ "embed"
)

//go:embed webui.wasm
var WasmBytes []byte
