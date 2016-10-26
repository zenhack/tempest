package util

import (
	"golang.org/x/net/context"
	"io"
	capnp "zenhack.net/go/sandstorm/capnp/util"
)

// A WriteCloserByteStream wraps an io.WriteCloser to satisfy the ByteStream
// interface.
type WriteCloserByteStream struct {
	WC io.WriteCloser
}

// Write calls Write on the WriteCloser
func (bs *WriteCloserByteStream) Write(args capnp.ByteStream_write) error {
	data, err := args.Params.Data()
	if err != nil {
		return nil
	}
	_, err = bs.WC.Write(data)
	return err
}

// Done calls Close on the WriteCloser
func (bs *WriteCloserByteStream) Done(args capnp.ByteStream_done) error {
	return bs.WC.Close()
}

// ExpectSize is a No-Op, provided to satisfy the ByteStream interface.
func (bs *WriteCloserByteStream) ExpectSize(args capnp.ByteStream_expectSize) error {
	return nil
}

// A ByteStreamWriteCloser wraps a ByteStream to satisfy the io.WriteCloser
// interface.
type ByteStreamWriteCloser struct {
	Ctx context.Context
	Obj capnp.ByteStream
}

// Write calls Write on the ByteStream
func (bs ByteStreamWriteCloser) Write(p []byte) (n int, err error) {
	bs.Obj.Write(bs.Ctx, func(args capnp.ByteStream_write_Params) error {
		args.SetData(p)
		return nil
	})
	return len(p), err
}

// Close calls Done on the ByteStream
func (bs ByteStreamWriteCloser) Close() error {
	bs.Obj.Done(bs.Ctx, func(args capnp.ByteStream_done_Params) error {
		return nil
	})
	return nil
}
