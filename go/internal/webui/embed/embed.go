package embed

import (
	_ "embed"
)

//go:embed webui.wasm
var WasmBytes []byte

//go:embed wasm_exec.js
var WasmExecJs []byte
