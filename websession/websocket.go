package websession

import (
	"bufio"
	"golang.org/x/net/context"
	"io"
	"net"
	"net/http"
	"strings"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/internal/iocommon"
)

type WebSocketStreamWriteCloser struct {
	capnp.WebSession_WebSocketStream
	Ctx context.Context
}

type websocketResponseWriter struct {
	hd          *heading
	hijack      chan<- struct{}
	hijacked    bool
	headingChan chan<- *heading
	body        io.WriteCloser
	request     *http.Request
}

type heading struct {
	status int
	header map[string][]string
}

func (w *websocketResponseWriter) Write(p []byte) (int, error) {
	if !w.hijacked && w.hd.status == 0 {
		w.WriteHeader(200)
	}
	return w.body.Write(p)
}

func (w *websocketResponseWriter) WriteHeader(status int) {
	if !w.hijacked {
		w.hd.status = status
		w.headingChan <- w.hd
		return
	}

	(&http.Response{
		StatusCode: w.hd.status,
		Header:     w.hd.header,
	}).Write(w.body)
}

func (w *websocketResponseWriter) Header() http.Header {
	return w.hd.header
}

func (w *websocketResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	w.hijacked = true
	w.hijack <- struct{}{}
	conn := &iocommon.RWCConn{
		ReadWriteCloser: iocommon.MergedRWC{
			Reader: w.request.Body,
			Writer: w.body,
			Closer: iocommon.MultiCloser(w.body, w.request.Body),
		},
		Local: &iocommon.HardCodedAddr{
			Net:  "capnp",
			Addr: "app",
		},
		Remote: &iocommon.HardCodedAddr{
			Net:  "capnp",
			Addr: "client",
		},
	}
	bufrw := &bufio.ReadWriter{
		Reader: bufio.NewReader(conn),
		Writer: bufio.NewWriter(conn),
	}
	return conn, bufrw, nil
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

func buildWebSocketHeading(p *capnp.WebSession_openWebSocket, hd *heading) error {
	// TODO: handle multiple instances of the header
	protoSlice := strings.Split(http.Header(hd.header).Get("Sec-Websocket-Protocol"), ",")
	protoList, err := p.Results.NewProtocol(int32(len(protoSlice)))
	if err != nil {
		return err
	}
	for i := range protoSlice {
		protoList.Set(i, strings.Trim(protoSlice[i], " "))
	}
	return nil
}
