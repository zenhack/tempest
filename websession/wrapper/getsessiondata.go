package handler

import (
	"net/http"
)

// Extract the SessionData from a request.
//
// This allows an http.Handler to access the information passed to new*Session.
// Its argument must be a request made on a session returned by HandlerUiView,
// otherwise it will panic.
func GetSessionData(req *http.Request) SessionData {
	return req.Context().Value(sessionDataKey).(SessionData)
}

// a simple type wrapper for use as a key with context.Value.
type ctxKey int

const (
	// Key used to store the SessionData in a request's context.
	sessionDataKey ctxKey = iota
)
