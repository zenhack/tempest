package util

import (
	"golang.org/x/net/context"
	"io"
	capnp "zenhack.net/go/sandstorm/capnp/util"
)

type WriteCloserByteStream struct {
	WC io.WriteCloser
}

func (bs *WriteCloserByteStream) Write(args capnp.ByteStream_write) error {
	data, err := args.Params.Data()
	if err != nil {
		return nil
	}
	_, err = bs.WC.Write(data)
	return err
}

func (bs *WriteCloserByteStream) Done(args capnp.ByteStream_done) error {
	return bs.WC.Close()
}

func (bs *WriteCloserByteStream) ExpectSize(args capnp.ByteStream_expectSize) error {
	return nil
}

type ByteStreamWriteCloser struct {
	Ctx context.Context
	Obj capnp.ByteStream
}

func (bs ByteStreamWriteCloser) Write(p []byte) (n int, err error) {
	bs.Obj.Write(bs.Ctx, func(args capnp.ByteStream_write_Params) error {
		args.SetData(p)
		return nil
	})
	return len(p), err
}

func (bs ByteStreamWriteCloser) Close() error {
	_, err := bs.Obj.Done(bs.Ctx, func(args capnp.ByteStream_done_Params) error {
		return nil
	}).Struct()
	return err
}
