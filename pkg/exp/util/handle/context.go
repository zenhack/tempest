// Package handle provides utilites for working with sandstorm's util.Handle
package handle

import (
	"context"

	"zenhack.net/go/tempest/capnp/util"
)

// WithCancel is like context.WithCancel, except that it returns a Handle
// instead of a context.CancelFunc. When the handle is dropped, the context
// is cancelled.
func WithCancel(ctx context.Context) (context.Context, util.Handle) {
	ctx, cancel := context.WithCancel(ctx)
	return ctx, CallbackHandle(cancel)
}
