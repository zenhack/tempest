// Package websession facilitates using http.Handlers with web-session.
//
// Overview:
//
// - HandlerUiView implements new*Session for you; you just need to supply
//   an http.Handler and an implementation of GetViewInfo.
// - GetSessionData allows http.Handlers to access session information.
// - SessionData contains the information passed to the new*Session methods.
package websession
