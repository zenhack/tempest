package sandstormhttpbridge

import (
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
	"time"
	grain_capnp "zenhack.net/go/sandstorm/capnp/grain"
	capnp "zenhack.net/go/sandstorm/capnp/sandstormhttpbridge"
	"zombiezen.com/go/capnproto2/rpc"
)

// Like http.ListenAndServe, except that the http.ResponseWriters passed to the
// handler will always implement grain.GetSessionContext, using bridge to retrieve the session
// context.
//
// If bridge is nil, ListenAndServe will connect to the bridge listening at
// /tmp/sandstorm-api. In this case, `ctx` will be used for the capnproto
// rpc session (otherwise it is ignored).
func ListenAndServe(ctx context.Context,
	bridge *capnp.SandstormHttpBridge, addr string, handler http.Handler) error {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	log.Println("listeneing")

	if bridge == nil {
		log.Println("bridge nil")
		// sandstorm-http-bridge won't start listening for capnp connections
		// until it successfully connects to the web server, so we need to
		// accept a connection before we can do anything else:
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		log.Println("Got an HTTP connection")
		conn.Close()

		bridge = &capnp.SandstormHttpBridge{}

		for i := 0; i < 100; i++ {
			// Even after accepting the HTTP connection,
			// there's a race condition where we try to
			// connect to the socket before
			// sandstorm-http-bridge actually starts
			// listening. To give it ample time to get
			// started, we retry on failures for a while.
			*bridge, err = Dial(ctx)
			if err == nil {
				break
			}
			time.Sleep(time.Second / 5)
		}
		if err != nil {
			return err
		}
	}
	if handler == nil {
		handler = http.DefaultServeMux
	}
	handler = proxiedHandler{
		handler: handler,
		bridge:  *bridge,
	}
	return http.Serve(l, handler)
}

// Establish an rpc session with sandstorm-http-bridge. Note that this will
// fail if sandstorm-http-bridge has not yet detected an HTTP server from
// the sandstorm app; see:
//
// https://github.com/sandstorm-io/sandstorm/issues/2779
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
