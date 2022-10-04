package handle

import (
	"context"

	"zenhack.net/go/sandstorm/capnp/util"
)

// CallbackHandle returns a handle which calls cb when the last reference
// to it is dropped.
func CallbackHandle(cb func()) util.Handle {
	return util.Handle_ServerToClient(callbackHandle(cb))
}

type callbackHandle func()

func (callbackHandle) Ping(context.Context, util.Handle_ping) error {
	return nil
}

func (cb callbackHandle) Shutdown() {
	cb()
}
