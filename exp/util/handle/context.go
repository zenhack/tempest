// Package handle provides utilites for working with sandstorm's util.Handle
package handle

import (
	"context"

	"zenhack.net/go/sandstorm/capnp/util"

	"zombiezen.com/go/capnproto2/server"
)

// WithCancel is like context.WithCancel, except that it returns a Handle
// instead of a context.CancelFunc. When the handle is dropped, the context
// is cancelled.
func WithCancel(ctx context.Context) (context.Context, util.Handle) {
	ctx, cancel := context.WithCancel(ctx)

	// The server policy doesn't matter, because handles have no methods:
	return ctx, util.Handle_ServerToClient(&cancelHandle{cancel}, &server.Policy{})
}

type cancelHandle struct {
	cancel context.CancelFunc
}

func (h *cancelHandle) Close() error {
	h.cancel()
	return nil
}
