package util

import (
	"golang.org/x/net/context"
	capnp "zenhack.net/go/sandstorm/capnp/util"
)

type ByteStream struct {
	Ctx context.Context
	Obj capnp.ByteStream
}

func (bs ByteStream) Write(p []byte) (n int, err error) {
	bs.Obj.Write(bs.Ctx, func(args capnp.ByteStream_write_Params) error {
		args.SetData(p)
		return nil
	})
	return len(p), err
}

func (bs ByteStream) Close() error {
	bs.Obj.Done(bs.Ctx, func(args capnp.ByteStream_done_Params) error {
		return nil
	})
	return nil
}
