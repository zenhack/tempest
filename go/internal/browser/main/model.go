package browsermain

import (
	"syscall/js"

	"zenhack.net/go/sandstorm/capnp/util"
)

func initModel() Model {
	return Model{
		Host:   js.Global().Get("window").Get("location").Get("host").String(),
		Grains: make(map[string]Grain),
	}
}

type Model struct {
	Host string

	Grains map[string]Grain
}

type Grain struct {
	Title        string
	SessionToken string
	Handle       util.Handle
}
