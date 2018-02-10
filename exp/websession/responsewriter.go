package websession

// Implement http.ResponseWriter on top of WebSession.

import (
	"io"
	"net/http"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

// ResponseWriter used for WebSession methods that just return a WebSession_Response,
// e.g. get, put, and post, but not postStreaming or openWebSocket.
type basicResponseWriter struct {
	// Status to send. 0 until WriteHeader() has been called.
	statusCode int

	// handle which cancels the request's Context.
	cancel util.Handle

	header     http.Header
	bodyWriter io.Writer
	response   websession.WebSession_Response
}

func (w *basicResponseWriter) Header() http.Header {
	return w.header
}

func (w *basicResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		panic("WriteHeader called twice!")
	}
	w.statusCode = statusCode
	switch statusCode {
	case 200:
		w.response.SetContent()
		content := w.response.Content()
		content.SetStatusCode(websession.WebSession_Response_SuccessCode_ok)

		body := content.Body()
		body.SetStream(w.cancel)

		// TODO:
		//
		// * encoding
		// * language
		// * mimeType
		// * eTag
		// * body
		// * disposition
	// TODO: other status codes.
	default:
		panic("Not implemented")
	}

	// TODO:
	//
	// * additionalHeaders
}

func (w *basicResponseWriter) Write(data []byte) (n int, err error) {
	if w.statusCode == 0 {
		w.WriteHeader(200)
	}
	return w.bodyWriter.Write(data)
}
