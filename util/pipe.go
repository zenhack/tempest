package util

import (
	"io"

	util_capnp "zenhack.net/go/sandstorm/capnp/util"
)

// ByteStreamPipe is like io.Pipe, except that the write end is a ByteStream.
//
// The ByteStream's ExpectSize method is a noop.
//
// Once Done is called on the ByteStream, reads will return io.EOF.
//
// If all references to the ByteStream are dropped before Done is called, reads
// will return io.ErrUnexpectedEOF.
func ByteStreamPipe() (*io.PipeReader, util_capnp.ByteStream) {
	r, w := io.Pipe()
	return r, util_capnp.ByteStream_ServerToClient(&byteStreamPipeWriter{
		w:        w,
		isClosed: false,
	})
}

// The type that powers ByteStreamPipe; see the comments there for an overview.
type byteStreamPipeWriter struct {
	w        *io.PipeWriter
	isClosed bool
}

func (w *byteStreamPipeWriter) ExpectSize(p util_capnp.ByteStream_expectSize) error {
	if w.isClosed {
		return io.ErrClosedPipe
	}
	return nil
}

func (w *byteStreamPipeWriter) Write(p util_capnp.ByteStream_write) error {
	if w.isClosed {
		return io.ErrClosedPipe
	}
	data, err := p.Params.Data()
	if err != nil {
		w.w.CloseWithError(err)
		w.isClosed = true
		return err
	}
	_, err = w.w.Write(data)
	return err
}

func (w *byteStreamPipeWriter) Done(p util_capnp.ByteStream_done) error {
	if w.isClosed {
		return io.ErrClosedPipe
	}
	w.isClosed = true
	return w.w.Close()
}

func (w *byteStreamPipeWriter) Close() error {
	if !w.isClosed {
		w.w.CloseWithError(io.ErrUnexpectedEOF)
		w.isClosed = true
	}
	return nil
}
