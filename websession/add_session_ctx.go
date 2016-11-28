package websession

import (
	"bufio"
	"net"
	"net/http"
	grain_capnp "zenhack.net/go/sandstorm/capnp/grain"
)

type hijackResponseWriter struct {
	sessionResponseWriter
}

func (w hijackResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.sessionResponseWriter.ResponseWriter.(http.Hijacker).Hijack()
}

type sessionResponseWriter struct {
	http.ResponseWriter
	sessionCtx grain_capnp.SessionContext
}

func (w sessionResponseWriter) GetSessionContext() grain_capnp.SessionContext {
	return w.sessionCtx
}

func WithSessionContext(handler http.Handler, sessionCtx grain_capnp.SessionContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ret := sessionResponseWriter{
			ResponseWriter: w,
			sessionCtx:     sessionCtx,
		}
		if _, ok := w.(http.Hijacker); ok {
			handler.ServeHTTP(hijackResponseWriter{ret}, req)
		} else {
			handler.ServeHTTP(ret, req)
		}
	})
}
