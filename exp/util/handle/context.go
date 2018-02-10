// Package handle provides utilites for working with sandstorm's util.Handle
package handle

import (
	"context"

	"zenhack.net/go/sandstorm/capnp/util"
)

// WithCancel is like context.WithCancel, except that it returns a Handle
// instead of a context.CancelFunc. When the handle is dropped, the context
// is cancelled.
func WithCancel(ctx context.Context) (context.Context, util.Handle) {
	ctx, cancel := context.WithCancel(ctx)
	return ctx, util.Handle_ServerToClient(&cancelHandle{cancel})
}

type cancelHandle struct {
	cancel context.CancelFunc
}

func (h *cancelHandle) Close() error {
	h.cancel()
	return nil
}
