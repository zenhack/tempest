package errors

import (
	"errors"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/rpc"
	rpc_capnp "zombiezen.com/go/capnproto2/std/capnp/rpc"
)

var (
	NotImplemented = errors.New("Not implemented")
)

// Return an 'unimplemented' capnp exception allocated in s. If an
// error occurs in generating the exception, that error will be
// returned instead.
func UnImplementedExn(s *capnp.Segment) error {
	exn, err := rpc_capnp.NewException(s)
	if err != nil {
		return err
	}
	exn.SetType(rpc_capnp.Exception_Type_unimplemented)
	return rpc.Exception{exn}
}
