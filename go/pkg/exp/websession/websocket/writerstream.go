package websocket

import (
	"context"
	"io"

	websession "zenhack.net/go/tempest/capnp/web-session"
)

// WriterStream implements websession.WebSocketStream_Server by writing data to W.
// Note: the raw websocket traffic is written, rather than the high-level websocket
// messages (i.e. we do not interpret headers & individual message boundaries).
type WriterStream struct {
	W io.Writer
}

func (w WriterStream) SendBytes(ctx context.Context, p websession.WebSocketStream_sendBytes) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	data, err := p.Args().Msg()
	if err != nil {
		return err
	}
	_, err = w.W.Write(data)
	return err
}
