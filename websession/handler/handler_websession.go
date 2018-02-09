package handler

import (
	"net/http"

	"zenhack.net/go/sandstorm/capnp/websession"
)

// Our UiSession implementation; this implements WebSession, and holds both the
// SessionData from the new*Session call and the http.Handler to invoke.
type handlerWebSession struct {
	sessionData SessionData
	handler     http.Handler
}

// Stubs for unimplemented websession methods:

func (*handlerWebSession) Get(websession.WebSession_get) error       { panic("Not implemented") }
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
