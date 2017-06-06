package util

import (
	"context"
)

// A CancelHandle wraps a context.CancelFunc, such that it may be exported
// via capnproto rpc and will be called when the remote ref is dropped.
type CancelHandle context.CancelFunc

func (c CancelHandle) Close() error {
	c()
	return nil
}
