package sandstormhttpbridge

import (
	"context"
	"errors"
	"net"
	"os"
	"time"

	bridge "zenhack.net/go/tempest/capnp/sandstorm-http-bridge"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
)

// Connect to the API socket, using exponential backoff to wait for the bridge to
// start listening.
//
// TODO: it would be nice if the bridge would set this up in a way such that we could
// assume the socket is already there on start; see about sending a patch upstream.
func connectSocket() (net.Conn, error) {
	conn, err := tryConnectSocket()
	delay := time.Second / 100
	for delay < time.Second && errors.Is(err, os.ErrNotExist) {
		time.Sleep(delay)
		conn, err = tryConnectSocket()
		delay *= 2
	}
	return conn, err
}

func tryConnectSocket() (net.Conn, error) {
	return net.Dial("unix", "/tmp/sandstorm-api")
}

func connectBridge(ctx context.Context, hooks capnp.Client) (ret bridge.SandstormHttpBridge, err error) {
	conn, err := connectSocket()
	if err != nil {
		return
	}
	var options *rpc.Options
	if (hooks != capnp.Client{}) {
		options = &rpc.Options{
			BootstrapClient: hooks,
		}
	}

	transport := rpc.NewStreamTransport(conn)
	ret = bridge.SandstormHttpBridge(rpc.NewConn(transport, options).Bootstrap(ctx))
	return
}

func ConnectWithHooks(ctx context.Context, hooks bridge.AppHooks) (bridge.SandstormHttpBridge, error) {
	return connectBridge(ctx, capnp.Client(hooks))
}

func Connect(ctx context.Context) (bridge.SandstormHttpBridge, error) {
	return connectBridge(ctx, capnp.Client{})
}
