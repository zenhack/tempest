// Package websession provides helpers for speaking HTTP over the sandstorm API.
//
// The main thing provided by this package HandlerWebSession, which allows
// converting package "net/http"'s Handler type to Sandstorm WebSessions. This
// makes writing web servers pretty much the same between sandstorm and
// non-sandstorm environments.
package websession

import (
	"net/http"
	"net/url"
	"zenhack.net/go/sandstorm/capnp/util"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
)

func FromHandler(h http.Handler) capnp.WebSession_Server {
	return HandlerWebSession{h}
}

type responseWriter struct {
	status   int
	header   http.Header
	stream   util.ByteStream
	response *capnp.WebSession_Response
}

func (r *responseWriter) Header() http.Header {
	return r.header
}

func (r *responseWriter) WriteHeader(status int) {
	r.status = status
	switch {
	case status >= 200 && status <= 202:
		r.response.SetContent()
		// TODO: the subtraction is breaking an abstraction boundary just a bit.
		// should be safe in this case, but let's think about the implications.
		capnpStatus := capnp.WebSession_Response_SuccessCode(status - 200)
		r.response.Content().SetStatusCode(capnpStatus)
		// TODO: Figure out what we should be passing here; do we actually need to do
		// anything? Handle_Server is just interface{}, so we're passing in 0, since it's
		// handy.
		r.response.Content().Body().SetStream(util.Handle_ServerToClient(0))
		// TODO: actually set headers, and handle non-200 responses.
	}
}

func (r *responseWriter) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.WriteHeader(200)
	}
	// TODO: actually implement
	return 0, nil
}

type HandlerWebSession struct {
	http.Handler
}

func (h HandlerWebSession) Get(args capnp.WebSession_get) error {
	var firstErr error
	var err error
	checkErr := func() {
		if firstErr == nil {
			firstErr = err
		}
	}

	// TODO: some of these have a HasFoo() method; should look into the exact semantics of all
	// this and see what the right way to do this is.
	path, err := args.Params.Path()
	checkErr()
	ctx, err := args.Params.Context()
	checkErr()
	// TODO: pull these out, and add them to request below:
	//cookies, err := ctx.Cookies()
	//checkErr()
	//accept, err := ctx.Accept()
	//checkErr()
	responseStream := ctx.ResponseStream()
	if firstErr != nil {
		return firstErr
	}

	// TODO/FIXME: we need to make sure we're mapping parsed/unparsed URL fields to the right
	// thing; mixing this up could have serioius security implications. Need to dig into this
	// a bit more.
	request := http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: path,
		},
	}

	respond := responseWriter{
		header:   make(map[string][]string),
		stream:   responseStream,
		response: &args.Results,
	}

	h.ServeHTTP(&respond, &request)
	return nil
}

func (h HandlerWebSession) Post(capnp.WebSession_post) error {
	return nil
}

func (h HandlerWebSession) Put(capnp.WebSession_put) error {
	return nil
}

func (h HandlerWebSession) Delete(capnp.WebSession_delete) error {
	return nil
}

func (h HandlerWebSession) PostStreaming(capnp.WebSession_postStreaming) error {
	return nil
}

func (h HandlerWebSession) PutStreaming(capnp.WebSession_putStreaming) error {
	return nil
}

func (h HandlerWebSession) OpenWebSocket(capnp.WebSession_openWebSocket) error {
	return nil
}
