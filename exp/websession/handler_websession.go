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
	"zenhack.net/go/sandstorm/internal/errors"
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

//// Actual WebSession methods ////

func (h *handlerWebSession) Get(p websession.WebSession_get) error {
	return h.handleCommon(p.Ctx, p.Params, p.Results, func(req *http.Request) error {
		if p.Params.IgnoreBody() {
			req.Method = "HEAD"
		} else {
			req.Method = "GET"
		}
		return nil
	})
}

func (h *handlerWebSession) Post(p websession.WebSession_post) error {
	content, err := p.Params.Content()
	if err != nil {
		return err
	}
	return h.handlePContent(p.Ctx, p.Params, p.Results, content, "POST")
}

func (h *handlerWebSession) Put(p websession.WebSession_put) error {
	content, err := p.Params.Content()
	if err != nil {
		return err
	}
	return h.handlePContent(p.Ctx, p.Params, p.Results, content, "PUT")
}

func (h *handlerWebSession) Delete(p websession.WebSession_delete) error {
	return h.handleCommon(p.Ctx, p.Params, p.Results, func(req *http.Request) error {
		req.Method = "DELETE"
		return nil
	})
}

func (h *handlerWebSession) Patch(p websession.WebSession_patch) error {
	content, err := p.Params.Content()
	if err != nil {
		return err
	}
	return h.handlePContent(p.Ctx, p.Params, p.Results, content, "PATCH")
}
func (h *handlerWebSession) PostStreaming(p websession.WebSession_postStreaming) error {
	reqR, reqW := bytestream.PipeServer()
	reqStream := &requestStream{
		ByteStream_Server: reqW,
		responseChan:      make(chan websession.WebSession_Response, 1),
		errChan:           make(chan error, 1),
	}
	p.Results.SetStream(websession.WebSession_RequestStream_ServerToClient(reqStream))
	response, err := websession.NewWebSession_Response(p.Params.Segment())
	if err != nil {
		panic("Error allocating response: " + err.Error())
	}
	go func() {
		// It's not clear to me(zenhack) what context we should use here;
		// we can't use p.Ctx because that will be canceled when
		// postStreaming returns.
		basicW, req, err := h.initRequest(context.TODO(), p.Params)
		if err != nil {
			reqStream.errChan <- err
			return
		}

		req.Method = "POST"
		// TODO: copy mimeType & encoding.

		basicW.response = response
		w := &streamingResponseWriter{
			basic:        basicW,
			responseChan: reqStream.responseChan,
		}
		req.Body = reqR
		h.handler.ServeHTTP(w, req)
		if w.basic.statusCode == 0 {
			w.WriteHeader(200)
		}
		w.basic.finishResponse(req.Context())
	}()
	return nil
}

//// Stubs for unimplemented WebSession methods ////

func (*handlerWebSession) PutStreaming(p websession.WebSession_putStreaming) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) OpenWebSocket(p websession.WebSession_openWebSocket) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Propfind(p websession.WebSession_propfind) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Proppatch(p websession.WebSession_proppatch) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Mkcol(p websession.WebSession_mkcol) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Copy(p websession.WebSession_copy) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Move(p websession.WebSession_move) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Lock(p websession.WebSession_lock) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Unlock(p websession.WebSession_unlock) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Acl(p websession.WebSession_acl) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (*handlerWebSession) Report(p websession.WebSession_report) error {
	return errors.UnImplementedExn(p.Results.Segment())
}

func (h *handlerWebSession) Options(p websession.WebSession_options) error {
	return errors.UnImplementedExn(p.Results.Segment())
}
