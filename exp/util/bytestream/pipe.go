package bytestream

import (
	"io"

	"zenhack.net/go/sandstorm/capnp/util"
)

// Pipe() is like io.Pipe(), except that the write end is a ByteStream.
//
// The ByteStream's ExpectSize method is a noop.
//
// Once Done is called on the ByteStream, reads will return io.EOF.
//
// If all references to the ByteStream are dropped before Done is called,
// further reads will return io.ErrUnexpectedEOF.
func Pipe() (r *io.PipeReader, w util.ByteStream) {
	r, wServer := PipeServer()
	return r, util.ByteStream_ServerToClient(wServer)
}

// PipeServer() is like Pipe(), except that it returns a (ByteStream) server,
// rather than a client.
func PipeServer() (r *io.PipeReader, w util.ByteStream_Server) {
	ioR, ioW := io.Pipe()
	return ioR, &pipeWriter{
		w:        ioW,
		isClosed: false,
	}
}

// Implementation of ByteStream that wraps an io.PipeWriter; returned by Pipe().
type pipeWriter struct {
	w        *io.PipeWriter
	isClosed bool
}

func (w *pipeWriter) ExpectSize(util.ByteStream_expectSize) error {
	// expectSize is a no-op, but we check if the pipe is already closed and
	// report an error accordingly.
	if w.isClosed {
		return io.ErrClosedPipe
	}
	return nil
}

func (w *pipeWriter) Write(p util.ByteStream_write) error {
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

func (w *pipeWriter) Done(p util.ByteStream_done) error {
	if w.isClosed {
		return io.ErrClosedPipe
	}
	w.isClosed = true
	return w.w.Close()
}

func (w *pipeWriter) Close() error {
	if w.isClosed {
		return io.ErrClosedPipe
	}
	w.w.CloseWithError(io.ErrUnexpectedEOF)
	w.isClosed = true
	return nil
}
