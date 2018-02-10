package websession

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"net/url"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/exp/util/bytestream"
	"zenhack.net/go/sandstorm/exp/util/handle"
)

// Parameters common to all websession request methods
type commonParams interface {
	Path() (string, error)
	Context() (websession.WebSession_Context, error)
}

// Our UiSession implementation; this implements WebSession, and holds both the
// SessionData from the new*Session call and the http.Handler to invoke.
type handlerWebSession struct {
	sessionData SessionData
	handler     http.Handler
}

//// Helpers for common parts of request handling ////

// Intialize a request with the data common to all webession request methods.
//
// The request will have a Context that is derived from ctx, but includes the
// SessionData in its Values.
func (h *handlerWebSession) initRequest(ctx context.Context, params commonParams) (*http.Request, error) {
	path, err := params.Path()
	if err != nil {
		return nil, err
	}

	wsCtx, err := params.Context()
	if err != nil {
		return nil, err
	}

	// Sandstorm gives us a path no leading slash, but Go's http library
	// expects one:
	parsedUrl, err := url.ParseRequestURI("/" + path)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, sessionDataKey, h.sessionData)
	req := &http.Request{
		Header: http.Header{},
		URL:    parsedUrl,
	}

	req = req.WithContext(ctx)
	err = copyContextInfo(req, wsCtx)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Copy the information from the context into the request.
func copyContextInfo(req *http.Request, wsCtx websession.WebSession_Context) error {
	// cookies

	cookies, err := wsCtx.Cookies()
	if err != nil {
		return err
	}
	numCookies := cookies.Len()
	for i := 0; i < numCookies; i++ {
		kv := cookies.At(i)
		key, err := kv.Key()
		if err != nil {
			return err
		}
		val, err := kv.Value()
		if err != nil {
			return err
		}
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}

	// accept

	accept, err := wsCtx.Accept()
	if err != nil {
		return err
	}
	numAccept := accept.Len()
	acceptHeaders := make([]string, numAccept)
	for i := 0; i < numAccept; i++ {
		str, err := formatAccept(accept.At(i))
		if err != nil {
			return err
		}
		acceptHeaders[i] = str
	}
	req.Header["Accept"] = acceptHeaders

	// TODO:
	//
	// acceptEncoding
	// eTagPrecondition
	// additionalHeaders

	return nil
}

// format the argument as expected for the value of the "Accept" header.
func formatAccept(typ websession.WebSession_AcceptedType) (string, error) {
	mimeType, err := typ.MimeType()
	if err != nil {
		return "", err
	}
	param := map[string]string{
		"q": fmt.Sprint(typ.QValue()),
	}
	return mime.FormatMediaType(mimeType, param), nil
}

//// Actual WebSession methods ////

func (h *handlerWebSession) Get(p websession.WebSession_get) error {
	ctx, cancel := handle.WithCancel(p.Ctx)
	req, err := h.initRequest(ctx, p.Params)
	if err != nil {
		return err
	}

	// Just so it's not nil:
	req.Body = http.NoBody

	// TODO: probably factor this out.
	wsCtx, err := p.Params.Context()
	if err != nil {
		return err
	}
	responseStream := wsCtx.ResponseStream()

	if p.Params.IgnoreBody() {
		req.Method = "HEAD"
	} else {
		req.Method = "GET"
	}

	w := &basicResponseWriter{
		statusCode: 0,
		header:     http.Header{},
		cancel:     cancel,
		bodyWriter: bytestream.ToWriteCloser(req.Context(), responseStream),
		response:   p.Results,
	}

	h.handler.ServeHTTP(w, req)
	responseStream.Done(ctx, func(util.ByteStream_done_Params) error {
		return nil
	})
	return nil
}

//// Stubs for unimplemented WebSession methods ////

func (*handlerWebSession) Post(websession.WebSession_post) error     { panic("Not implemented") }
func (*handlerWebSession) Put(websession.WebSession_put) error       { panic("Not implemented") }
func (*handlerWebSession) Delete(websession.WebSession_delete) error { panic("Not implemented") }
func (*handlerWebSession) Patch(websession.WebSession_patch) error   { panic("Not implemented") }
func (*handlerWebSession) PostStreaming(websession.WebSession_postStreaming) error {
	panic("Not implemented")
}
func (*handlerWebSession) PutStreaming(websession.WebSession_putStreaming) error {
	panic("Not implemented")
}
func (*handlerWebSession) OpenWebSocket(websession.WebSession_openWebSocket) error {
	panic("Not implemented")
}
func (*handlerWebSession) Propfind(websession.WebSession_propfind) error   { panic("Not implemented") }
func (*handlerWebSession) Proppatch(websession.WebSession_proppatch) error { panic("Not implemented") }
func (*handlerWebSession) Mkcol(websession.WebSession_mkcol) error         { panic("Not implemented") }
func (*handlerWebSession) Copy(websession.WebSession_copy) error           { panic("Not implemented") }
func (*handlerWebSession) Move(websession.WebSession_move) error           { panic("Not implemented") }
func (*handlerWebSession) Lock(websession.WebSession_lock) error           { panic("Not implemented") }
func (*handlerWebSession) Unlock(websession.WebSession_unlock) error       { panic("Not implemented") }
func (*handlerWebSession) Acl(websession.WebSession_acl) error             { panic("Not implemented") }
func (*handlerWebSession) Report(websession.WebSession_report) error       { panic("Not implemented") }
func (*handlerWebSession) Options(websession.WebSession_options) error     { panic("Not implemented") }
