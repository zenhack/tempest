package websession

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/net/context"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
)

// Types to convert between WebSocketStream and WriteCloser:

type WebSocketStreamWriteCloser struct {
	capnp.WebSession_WebSocketStream
	Ctx context.Context
}

type WriteCloserWebSocketStream struct {
	io.WriteCloser
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

func (wc WriteCloserWebSocketStream) SendBytes(p capnp.WebSession_WebSocketStream_sendBytes) error {
	data, err := p.Params.Message()
	if err != nil {
		return err
	}
	_, err = wc.Write(data)
	return err
}

// Generate a random value for the Sec-Websocket-Key header. Sandstorm doesn't
// give this to us, so for existing server-side websocket libraries to work
// we need to provide it ourselves.
func makeWebSocketKey() string {
	rawBytes := make([]byte, 32)
	rand.Read(rawBytes)
	buf := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.StdEncoding, buf)
	enc.Write(rawBytes)
	enc.Close()
	return buf.String()
}

func (h handlerWebSession) OpenWebSocket(p capnp.WebSession_openWebSocket) error {
	path, err := p.Params.Path()
	if err != nil {
		return err
	}
	clientProtoList, err := p.Params.Protocol()
	if err != nil {
		return err
	}

	reqPipeReader, reqPipeWriter := io.Pipe()

	reqBodyCapnpWriter := capnp.WebSession_WebSocketStream_ServerToClient(
		WriteCloserWebSocketStream{reqPipeWriter},
	)

	p.Results.SetServerStream(reqBodyCapnpWriter)

	clientProtoHeader := make([]string, clientProtoList.Len())
	for i := range clientProtoHeader {
		clientProtoHeader[i], err = clientProtoList.At(i)
		if err != nil {
			return nil
		}
	}

	request := http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: makeAbsolute(path),
		},
		Header: map[string][]string{
			"Upgrade":                {"websocket"},
			"Connection":             {"Upgrade"},
			"Sec-Websocket-Key":      {makeWebSocketKey()},
			"Sec-Websocket-Protocol": clientProtoHeader,
			"Sec-Websocket-Version":  {"13"},
			// XXX: we should find something more sensible to put here:
			"Origin": {"http://dummy.example.com"},
		},
		Body: reqPipeReader,
	}

	clientStream := p.Params.ClientStream()
	capnpResponseBody := &WebSocketStreamWriteCloser{
		WebSession_WebSocketStream: clientStream,
	}
	runHandler(h.handler, &request, func(resp *http.Response) {
		// TODO: handle multiple instances of the header
		protoSlice := strings.Split(resp.Header.Get("Sec-Websocket-Protocol"), ",")
		protoList, err := p.Results.NewProtocol(int32(len(protoSlice)))
		if err != nil {
			// TODO: should probably change the type of runHandler so we can return an
			// error
			return
		}
		for i := range protoSlice {
			protoList.Set(i, strings.Trim(protoSlice[i], " "))
		}
		go func() {
			defer resp.Body.Close()
			io.Copy(capnpResponseBody, resp.Body)
		}()
	})
	return nil
}
