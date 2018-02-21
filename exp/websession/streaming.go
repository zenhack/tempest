package websession

import (
	"net/http"

	"zombiezen.com/go/capnproto2/server"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

type requestStream struct {
	util.ByteStream_Server
	responseChan chan websession.WebSession_Response
	errChan      chan error
}

// http.ResponseWriter for the putStreaming & postStreaming methods.
type streamingResponseWriter struct {
	basic        *basicResponseWriter
	responseChan chan websession.WebSession_Response
}

func (w *streamingResponseWriter) WriteHeader(status int) {
	w.basic.WriteHeader(status)
	if w.basic.bodyBuffer == nil {
		// The handler is streaming the response.
		w.responseChan <- w.basic.response
	} else {
		// The handler has picked a status that requires buffering
		// (4xx or 500). We'll wait until ServeHTTP returns.
	}
}

func (w *streamingResponseWriter) Write(p []byte) (int, error) {
	if w.basic.statusCode == 0 {
		w.WriteHeader(200)
	}
	return w.basic.Write(p)
}

func (w *streamingResponseWriter) Header() http.Header {
	return w.basic.Header()
}

func (rs *requestStream) GetResponse(p websession.WebSession_RequestStream_getResponse) error {
	server.Ack(p.Options)
	select {
	case p.Results = <-rs.responseChan:
		return nil
	case <-p.Ctx.Done():
		return p.Ctx.Err()
	}
}
