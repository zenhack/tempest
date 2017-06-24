package util

import (
	"io"

	util_capnp "zenhack.net/go/sandstorm/capnp/util"
)

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

func ByteStreamPipe() (*io.PipeReader, util_capnp.ByteStream) {
	r, w := io.Pipe()
	return r, util_capnp.ByteStream_ServerToClient(&byteStreamPipeWriter{
		w:        w,
		isClosed: false,
	})
}
