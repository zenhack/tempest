package websocket

import (
	"context"

	websession "zenhack.net/go/tempest/capnp/web-session"
)

// A StreamWriter implements io.Writer on top of a WebSocketStream.
type StreamWriter struct {
	Stream  websession.WebSocketStream // The stream to write to
	Context context.Context            // Context to use for RPC calls.
}

func (w StreamWriter) Write(data []byte) (int, error) {
	if err := w.Context.Err(); err != nil {
		return 0, err
	}
	_, rel := w.Stream.SendBytes(w.Context, func(p websession.WebSocketStream_sendBytes_Params) error {
		return p.SetMsg(data)
	})
	go rel()
	return len(data), nil
}
