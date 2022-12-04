package browsermain

import (
	"syscall/js"

	"zenhack.net/go/sandstorm/capnp/util"
)

func initModel() Model {
	return Model{
		Host:       js.Global().Get("window").Get("location").Get("host").String(),
		Grains:     make(map[ID[Grain]]Grain),
		OpenGrains: make(map[ID[OpenGrain]]OpenGrain),
	}
}

type ID[T any] string

type Model struct {
	Host string

	Grains     map[ID[Grain]]Grain
	OpenGrains map[ID[OpenGrain]]OpenGrain
}

type Grain struct {
	Title        string
	SessionToken string
	Handle       util.Handle
}

type OpenGrain struct {
}
