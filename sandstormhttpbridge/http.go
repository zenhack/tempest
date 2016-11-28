package sandstormhttpbridge

import (
	"golang.org/x/net/context"
	"net"
	"net/http"
	grain_capnp "zenhack.net/go/sandstorm/capnp/grain"
	capnp "zenhack.net/go/sandstorm/capnp/sandstormhttpbridge"
	"zombiezen.com/go/capnproto2/rpc"
)

// Like http.ListenAndServe, except that the http.ResponseWriters passed to the
// handler will always implement grain.GetSessionContext, using bridge to retrieve the session
// context.
func ListenAndServe(bridge capnp.SandstormHttpBridge, addr string, handler http.Handler) error {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	handler = proxiedHandler{
		handler: handler,
		bridge:  bridge,
	}
	return http.ListenAndServe(addr, handler)
}

// Establish an rpc session with sandstorm-http-bridge.
func Dial(ctx context.Context) (capnp.SandstormHttpBridge, error) {
	conn, err := net.Dial("unix", "/tmp/sandstorm-api")
	if err != nil {
		return capnp.SandstormHttpBridge{}, err
	}
	transport := rpc.StreamTransport(conn)
	client := rpc.NewConn(transport).Bootstrap(ctx)
	return capnp.SandstormHttpBridge{Client: client}, nil
}

type proxiedHandler struct {
	handler http.Handler
	bridge  capnp.SandstormHttpBridge
}

func (h proxiedHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handler.ServeHTTP(httpBridgeResponseWriter{
		ResponseWriter: w,
		requestId:      req.Header.Get("X-Sandstorm-Session-Id"),
		bridge:         h.bridge,
	}, req)
}

// A responseWriter which gets the websesion context from sandstorm-http-bridge.
type httpBridgeResponseWriter struct {
	http.ResponseWriter
	requestId string
	bridge    capnp.SandstormHttpBridge
}

func (w httpBridgeResponseWriter) GetSessionContext(ctx context.Context) grain_capnp.SessionContext {
	return w.bridge.GetSessionContext(ctx,
		func(p capnp.SandstormHttpBridge_getSessionContext_Params) error {
			p.SetId(w.requestId)
			return nil
		}).Context()
}
