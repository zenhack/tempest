package websession

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/exp/util/bytestream"
	"zenhack.net/go/sandstorm/exp/util/handle"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/server"
)

var specialRequestHeaders = map[string]struct{}{
	"Accept":          {},
	"Accept-Encoding": {},
	"Cookie":          {},
	"If-Match":        {},
	"If-None-Match":   {},
}

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
func (h *handlerWebSession) initRequest(ctx context.Context, params commonParams) (*basicResponseWriter, *http.Request, error) {
	ctx, cancel := handle.WithCancel(ctx)
	path, err := params.Path()
	if err != nil {
		return nil, nil, err
	}

	wsCtx, err := params.Context()
	if err != nil {
		return nil, nil, err
	}

	// Sandstorm gives us a path no leading slash, but Go's http library
	// expects one:
	parsedUrl, err := url.ParseRequestURI("/" + path)
	if err != nil {
		return nil, nil, err
	}

	ctx = context.WithValue(ctx, sessionDataKey, h.sessionData)
	req := &http.Request{
		Header: http.Header{},
		URL:    parsedUrl,
	}

	req = req.WithContext(ctx)
	err = copyContextInfo(req, wsCtx)
	if err != nil {
		return nil, nil, err
	}

	// Set this as a default, so it's never nil.
	req.Body = http.NoBody

	w := &basicResponseWriter{
		statusCode:     0,
		header:         http.Header{},
		cancel:         cancel,
		responseStream: wsCtx.ResponseStream(),
	}
	w.bodyWriter = bytestream.ToWriteCloser(req.Context(), w.responseStream)

	return w, req, nil
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
	acceptHeaders := make([]string, accept.Len())
	for i := range acceptHeaders {
		str, err := formatAccept(accept.At(i))
		if err != nil {
			return err
		}
		acceptHeaders[i] = str
	}
	req.Header["Accept"] = acceptHeaders

	acceptEncoding, err := wsCtx.AcceptEncoding()
	if err != nil {
		return err
	}
	acceptEncodingHeaders := make([]string, acceptEncoding.Len())
	for i := range acceptEncodingHeaders {
		encoding, err := acceptEncoding.At(i).ContentCoding()
		if err != nil {
			return err
		}
		acceptEncodingHeaders[i] = fmt.Sprintf(
			"%s;q=%v",
			encoding,
			acceptEncoding.At(i).QValue())
	}
	req.Header["Accept-Encoding"] = acceptEncodingHeaders

	eTagPrecondition := wsCtx.ETagPrecondition()
	switch eTagPrecondition.Which() {
	case websession.WebSession_Context_eTagPrecondition_Which_none:
	case websession.WebSession_Context_eTagPrecondition_Which_exists:
		req.Header.Set("If-Match", "*")
	case websession.WebSession_Context_eTagPrecondition_Which_doesntExist:
		req.Header.Set("If-None-Match", "*")
	case websession.WebSession_Context_eTagPrecondition_Which_matchesOneOf:
		etags, err := eTagPrecondition.MatchesOneOf()
		if err != nil {
			return err
		}
		etagString, err := formatETags(etags)
		if err != nil {
			return err
		}
		req.Header.Set("If-Match", etagString)
	case websession.WebSession_Context_eTagPrecondition_Which_matchesNoneOf:
		etags, err := eTagPrecondition.MatchesNoneOf()
		if err != nil {
			return err
		}
		etagString, err := formatETags(etags)
		if err != nil {
			return err
		}
		req.Header.Set("If-None-Match", etagString)
	}

	additionalHeaders, err := wsCtx.AdditionalHeaders()
	if err != nil {
		return err
	}
	for i := 0; i < additionalHeaders.Len(); i++ {
		hdr := additionalHeaders.At(i)
		name, err := hdr.Name()
		if err != nil {
			return err
		}
		value, err := hdr.Value()
		if err != nil {
			return err
		}
		_, ok := specialRequestHeaders[name]
		if ok {
			log.Printf("Warning: special request header %q in "+
				"websession additionalHeaders field", name)
		}
		req.Header.Set(name, value)
	}

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

func formatETags(etags websession.WebSession_ETag_List) (string, error) {
	etagStrings := make([]string, etags.Len())
	for i := range etagStrings {
		etag := etags.At(i)
		value, err := etag.Value()
		if err != nil {
			return "", err
		}

		// Sandstorm strips off the quotes, so we have to add them back:
		etagStrings[i] = `"` + value + `"`

		if etag.Weak() {
			etagStrings[i] = "W/" + etagStrings[i]
		}
	}
	return strings.Join(etagStrings, ", "), nil
}

// Common logic for the request methods that return a websession.WebSession_Request.
//
// `ctx` should be the context from the capnp method argument.
//
// `params` should be the capnp Params for the method.
//
// `response` should be the Response object to store the result in.
//
// `customize` is a function which will be called just before ServeHTTP. It
// is responsible for setting the HTTP Method on the request, as well as any
// other capnp method specific data.
func (h *handlerWebSession) handleCommon(
	ctx context.Context,
	params commonParams,
	response websession.WebSession_Response,
	customize func(*http.Request) error,
) error {
	w, req, err := h.initRequest(ctx, params)
	if err != nil {
		return err
	}
	w.response = response

	err = customize(req)
	if err != nil {
		return err
	}

	h.handler.ServeHTTP(w, req)
	if w.statusCode == 0 {
		w.WriteHeader(200)
	}
	return w.finishResponse(req.Context())
}

// PostContent/PutContent
type pContent interface {
	MimeType() (string, error)
	Content() ([]byte, error)
	HasEncoding() bool
	Encoding() (string, error)
}

func copyPContent(req *http.Request, content pContent) error {
	mimeType, err := content.MimeType()
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", mimeType)

	data, err := content.Content()
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if content.HasEncoding() {
		encoding, err := content.Encoding()
		if err != nil {
			return err
		}
		req.Header.Set("Content-Encoding", encoding)
	}
	return nil
}

// Common logic for capnp methods which take a `pContent` argument.
//
// `ctx`, `params`, and `response` are the same as in `handleCommon`.
//
// `content` is the `content` parameter.
// `method` is the HTTP method to set.
func (h *handlerWebSession) handlePContent(
	ctx context.Context,
	params commonParams,
	response websession.WebSession_Response,
	content pContent,
	method string,
) error {
	return h.handleCommon(ctx, params, response, func(req *http.Request) error {
		err := copyPContent(req, content)
		if err != nil {
			return err
		}
		req.Method = method
		return nil
	})
}

func (h *handlerWebSession) Shutdown() {
	// Release caps from sessionData
	wsCtx := h.sessionData.Context()
	wsCtx.Client.Release()
	switch h.sessionData.Which() {
	case SessionData_Which_offer:
		capability, err := h.sessionData.Offer().Offer()
		if err == nil {
			client := capability.Interface().Client()
			if client != nil {
				client.Release()
			}
		}
	}
}

//// Actual WebSession methods ////

func (h *handlerWebSession) Get(ctx context.Context, p websession.WebSession_get) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return h.handleCommon(ctx, p.Args(), results, func(req *http.Request) error {
		if p.Args().IgnoreBody() {
			req.Method = "HEAD"
		} else {
			req.Method = "GET"
		}
		return nil
	})
}

func (h *handlerWebSession) Post(ctx context.Context, p websession.WebSession_post) error {
	content, err := p.Args().Content()
	if err != nil {
		return err
	}
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return h.handlePContent(ctx, p.Args(), results, content, "POST")
}

func (h *handlerWebSession) Put(ctx context.Context, p websession.WebSession_put) error {
	content, err := p.Args().Content()
	if err != nil {
		return err
	}
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return h.handlePContent(ctx, p.Args(), results, content, "PUT")
}

func (h *handlerWebSession) Delete(ctx context.Context, p websession.WebSession_delete) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return h.handleCommon(ctx, p.Args(), results, func(req *http.Request) error {
		req.Method = "DELETE"
		return nil
	})
}

func (h *handlerWebSession) Patch(ctx context.Context, p websession.WebSession_patch) error {
	content, err := p.Args().Content()
	if err != nil {
		return err
	}
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return h.handlePContent(ctx, p.Args(), results, content, "PATCH")
}
func (h *handlerWebSession) PostStreaming(ctx context.Context, p websession.WebSession_postStreaming) error {
	reqR, reqW := bytestream.PipeServer()
	reqStream := &requestStream{
		ByteStream_Server: reqW,
		responseChan:      make(chan websession.WebSession_Response, 1),
		errChan:           make(chan error, 1),
	}
	results, err := p.AllocResults()
	if err != nil {
		panic("Error allocating results: " + err.Error())
	}
	reqStreamClient := websession.WebSession_RequestStream_ServerToClient(
		reqStream,
		&server.Policy{
			// These numbers are somewhat arbitrary, but it is critical
			// that MaxConcurrentCalls > 1, since getResponse is called
			// before pushing the bytes for the request body, and we
			// can't actually respond until ServeHTTP returns.
			MaxConcurrentCalls: 10,
			AnswerQueueSize:    10,
		},
	)
	results.SetStream(reqStreamClient)
	response, err := websession.NewWebSession_Response(p.Args().Segment())
	if err != nil {
		panic("Error allocating response: " + err.Error())
	}
	// It's not clear to me(zenhack) what context we should use here;
	// we can't use ctx because that will be canceled when
	// postStreaming returns.
	basicW, req, err := h.initRequest(context.TODO(), p.Args())
	if err != nil {
		return err
	}

	req.Method = "POST"
	// TODO: copy mimeType & encoding.

	basicW.response = response
	w := &streamingResponseWriter{
		basic:        basicW,
		responseChan: reqStream.responseChan,
	}
	req.Body = reqR
	go func() {
		h.handler.ServeHTTP(w, req)
		if w.basic.statusCode == 0 {
			w.WriteHeader(200)
		}
		w.basic.finishResponse(req.Context())
	}()
	return nil
}

//// Stubs for unimplemented WebSession methods ////

func (*handlerWebSession) PutStreaming(context.Context, websession.WebSession_putStreaming) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) OpenWebSocket(context.Context, websession.WebSession_openWebSocket) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Propfind(context.Context, websession.WebSession_propfind) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Proppatch(context.Context, websession.WebSession_proppatch) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Mkcol(context.Context, websession.WebSession_mkcol) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Copy(context.Context, websession.WebSession_copy) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Move(context.Context, websession.WebSession_move) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Lock(context.Context, websession.WebSession_lock) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Unlock(context.Context, websession.WebSession_unlock) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Acl(context.Context, websession.WebSession_acl) error {
	return capnp.Unimplemented("TODO")
}

func (*handlerWebSession) Report(context.Context, websession.WebSession_report) error {
	return capnp.Unimplemented("TODO")
}

func (h *handlerWebSession) Options(context.Context, websession.WebSession_options) error {
	return capnp.Unimplemented("TODO")
}
