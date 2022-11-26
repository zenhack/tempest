package browsermain

import "syscall/js"

func initModel() Model {
	return Model{
		Host: js.Global().Get("window").Get("location").Get("host").String(),
	}
}

type Model struct {
	Host string
}
