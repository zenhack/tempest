package websession

import (
	"fmt"
	"log"
	"mime"
	"net/http"

	"zenhack.net/go/sandstorm/capnp/websession"
)

// Our UiSession implementation; this implements WebSession, and holds both the
// SessionData from the new*Session call and the http.Handler to invoke.
type handlerWebSession struct {
	sessionData SessionData
	handler     http.Handler
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
	// responseStream
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

func (h *handlerWebSession) Get(p websession.WebSession_get) error {
	req := &http.Request{
		Header: http.Header{},
	}

	if p.Params.IgnoreBody() {
		req.Method = "HEAD"
	} else {
		req.Method = "GET"
	}

	wsCtx, err := p.Params.Context()
	if err != nil {
		return err
	}

	err = copyContextInfo(req, wsCtx)
	if err != nil {
		return err
	}

	log.Println(req)
	return nil
}

// Stubs for unimplemented websession methods:

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
