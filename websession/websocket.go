package websession

import (
	"golang.org/x/net/context"
	"io"
	//	"bufio"
	//	"net"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
)

type WebSocketStreamWriteCloser struct {
	capnp.WebSession_WebSocketStream
	Ctx context.Context
}

func (ws *WebSocketStreamWriteCloser) Write(p []byte) (int, error) {
	ws.WebSession_WebSocketStream.SendBytes(
		ws.Ctx,
		func(params capnp.WebSession_WebSocketStream_sendBytes_Params) error {
			params.SetMessage(p)
			return nil
		})
	return len(p), nil
}

func (ws *WebSocketStreamWriteCloser) Close() error {
	return ws.WebSession_WebSocketStream.Client.Close()
}

type WriteCloserWebSocketStream struct {
	io.WriteCloser
}

func (wc WriteCloserWebSocketStream) SendBytes(p capnp.WebSession_WebSocketStream_sendBytes) error {
	data, err := p.Params.Message()
	if err != nil {
		return err
	}
	_, err = wc.Write(data)
	return err
}
