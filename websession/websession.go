// Package websession provides helpers for speaking HTTP over the sandstorm API.
//
// The main thing provided by this package HandlerWebSession, which allows
// converting package "net/http"'s Handler type to Sandstorm WebSessions. This
// makes writing web servers pretty much the same between sandstorm and
// non-sandstorm environments.
package websession // import "zenhack.net/go/sandstorm/websession"

// Copyright (c) 2016 Ian Denhardt <ian@zenhack.net>
//
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"bytes"
	"golang.org/x/net/context"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"zenhack.net/go/sandstorm/capnp/grain"
	capnp_util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/internal/errors"
	"zenhack.net/go/sandstorm/util"
)

func FromHandler(ctx context.Context, h http.Handler) HandlerWebSession {
	return HandlerWebSession{ctx, h}
}

type responseWriter struct {
	status     int
	header     http.Header
	ctx        capnp.WebSession_Context
	body       io.WriteCloser
	response   *capnp.WebSession_Response
	webSession HandlerWebSession
}

func (r *responseWriter) Header() http.Header {
	return r.header
}

// value for responesWriter.body where a body is legal, but the capnp schema
// doesn't provide a way to do streaming. In this case we just buffer the
// writes and transmit on close.
//
// Note that closeCallback needs to be set to a function that will supply the
// apropriate field.
type bufBody struct {
	bytes.Buffer
	closeCallback func()
}

func (b *bufBody) Close() error {
	b.closeCallback()
	return nil
}

// io.WriteCloser with both metods NoOps. For responseWriter.body where no body
// is allowed.
type noBody struct {
}

func (b noBody) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (b noBody) Close() error {
	return nil
}

// Set all the headers that are present and accepted by sandstorm. Silently
// drop any headers sandstorm doesn't support.
func (r *responseWriter) setSuccessHeaders() {
	content := r.response.Content()
	if encoding := r.header.Get("Content-Encoding"); encoding != "" {
		content.SetEncoding(encoding)
	}
	if language := r.header.Get("Content-Language"); language != "" {
		content.SetLanguage(language)
	}
	if mimeType := r.header.Get("Content-Type"); mimeType != "" {
		content.SetMimeType(mimeType)
	}
	if disposition := r.header.Get("Content-Disposition"); disposition != "" {
		typ, params, err := mime.ParseMediaType(disposition)
		if err != nil {
			if typ == "attachment" {
				content.Disposition().SetDownload(params["filename"])
			} else {
				content.Disposition().SetNormal()
			}
		}
	}
}

// Trivial wrapper that adds a no-op Close() method to an io.Reader. Used for
// supplying buffers as request bodies.
type readDummyCloser struct {
	io.Reader
}

func (r readDummyCloser) Close() error {
	return nil
}

func (r *responseWriter) WriteHeader(status int) {
	r.status = status
	switch status {
	case 200, 201, 202:
		r.response.SetContent()
		// TODO: the subtraction is breaking an abstraction boundary just a bit.
		// should be safe in this case, but let's think about the implications.
		capnpStatus := capnp.WebSession_Response_SuccessCode(status - 200)
		r.response.Content().SetStatusCode(capnpStatus)
		// TODO: Figure out what we should be passing here; do we actually need to do
		// anything? Handle_Server is just interface{}, so we're passing in 0, since it's
		// handy.
		r.response.Content().Body().SetStream(capnp_util.Handle_ServerToClient(0))
		r.setSuccessHeaders()
		r.body = util.ByteStreamWriteCloser{r.webSession.Ctx, r.ctx.ResponseStream()}
	case 204, 205:
		r.body = noBody{}
		r.response.SetNoContent()
		r.response.NoContent().SetShouldResetForm(status == 205)
	case 301, 302, 303, 307:
		// Redirects. Web-session.capnp talks about a 308, but I haven't found
		// any info about its semantics. "net/http" doesn't deifine a constant
		// for it, so we're just going to say to heck with it and not support
		// it for now.
		r.body = noBody{}
		r.response.SetRedirect()
		r.response.Redirect().SetLocation(r.header.Get("Location"))
		switch status {
		case 301:
			r.response.Redirect().SetIsPermanent(true)
		case 302, 303, 307:
			r.response.Redirect().SetIsPermanent(false)
		}
		switch status {
		case 302, 303:
			r.response.Redirect().SetSwitchToGet(true)
		default:
			r.response.Redirect().SetSwitchToGet(false)
		}
	case 400, 403, 404, 405, 406, 409, 410, 413, 414, 415, 418:
		r.response.SetClientError()
		capnpStatus := capnp.WebSession_Response_ClientErrorCode(status - 400)
		r.response.ClientError().SetStatusCode(capnpStatus)
		buf := &bufBody{}
		buf.closeCallback = func() {
			r.response.ClientError().SetDescriptionHtml(buf.String())
		}
		r.body = buf
	default:
		r.response.SetServerError()
		if status >= 500 && status < 600 {
			// The handler actually returned a 5xx; let them set the body.
			buf := &bufBody{}
			buf.closeCallback = func() {
				r.response.ServerError().SetDescriptionHtml(buf.String())
			}
			r.body = buf
		} else {
			// The client returned some status sandstorm doesn't support. In this
			// case we just junk the body.
			//
			// TODO: Log this somewhere?
			r.body = noBody{}
		}
	}
}

func (r *responseWriter) Write(p []byte) (int, error) {
	if r.status == 0 {
		r.WriteHeader(200)
	}
	return r.body.Write(p)
}

type HandlerWebSession struct {
	Ctx context.Context
	http.Handler
}

func (h HandlerWebSession) handleRequest(method string, args requestArgs,
	headers map[string][]string,
	body io.ReadCloser,
	wsResponse *capnp.WebSession_Response) error {

	if headers == nil {
		headers = make(map[string][]string)
	}

	path, err := args.Path()
	if err != nil {
		return err
	}
	ctx, err := args.Context()
	if err != nil {
		return err
	}
	// TODO: pull these out, and add them to request below:
	//cookies, err := ctx.Cookies()
	//accept, err := ctx.Accept()

	if !strings.HasPrefix(path, "/") {
		// ServeMux will give us a redirect otherwise, and sandstorm
		// will then give us the relative path again, resulting in an
		// infinite redirect loop.
		//
		// As far as I know sandstorm always gives relative paths, but
		// I haven't found documentation actually saying so (yet).
		path = "/" + path
	}
	request := http.Request{
		Method: method,
		URL: &url.URL{
			Path: path,
		},
		Header: headers,
		Body:   body,
	}

	respond := responseWriter{
		webSession: h,
		ctx:        ctx,
		header:     make(map[string][]string),
		response:   wsResponse,
	}

	h.ServeHTTP(&respond, &request)
	if respond.status == 0 {
		(&respond).WriteHeader(200)
	}
	respond.body.Close()
	return nil
}

type requestArgs interface {
	// Arguments common to all request types
	Path() (string, error)
	Context() (capnp.WebSession_Context, error)
}

func (h HandlerWebSession) Get(args capnp.WebSession_get) error {
	return h.handleRequest("GET", args.Params, nil, nil, &args.Results)
}

func (h HandlerWebSession) Post(args capnp.WebSession_post) error {
	content, err := args.Params.Content()
	if err != nil {
		return err
	}
	payload, err := content.Content()
	if err != nil {
		return err
	}
	mimeType, err := content.MimeType()
	if err != nil {
		return err
	}
	body := readDummyCloser{bytes.NewBuffer(payload)}
	headers := map[string][]string{
		"Content-Type": {mimeType},
	}
	return h.handleRequest("POST", args.Params, headers, body, &args.Results)
}

// Websession stubs:

func (h HandlerWebSession) Put(capnp.WebSession_put) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Delete(capnp.WebSession_delete) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Patch(capnp.WebSession_patch) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) PostStreaming(capnp.WebSession_postStreaming) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) PutStreaming(capnp.WebSession_putStreaming) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) OpenWebSocket(capnp.WebSession_openWebSocket) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Propfind(capnp.WebSession_propfind) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Proppatch(capnp.WebSession_proppatch) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Mkcol(capnp.WebSession_mkcol) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Copy(capnp.WebSession_copy) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Move(capnp.WebSession_move) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Lock(capnp.WebSession_lock) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Unlock(capnp.WebSession_unlock) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Acl(capnp.WebSession_acl) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Report(capnp.WebSession_report) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) Options(capnp.WebSession_options) error {
	return errors.NotImplemented
}

// UiView stubs.

func (h HandlerWebSession) GetViewInfo(p grain.UiView_getViewInfo) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) NewSession(p grain.UiView_newSession) error {
	// TODO: Check params.
	client := capnp.WebSession_ServerToClient(h).Client
	p.Results.SetSession(grain.UiSession{Client: client})
	return nil
}

func (h HandlerWebSession) NewRequestSession(p grain.UiView_newRequestSession) error {
	return errors.NotImplemented
}

func (h HandlerWebSession) NewOfferSession(p grain.UiView_newOfferSession) error {
	return errors.NotImplemented
}
