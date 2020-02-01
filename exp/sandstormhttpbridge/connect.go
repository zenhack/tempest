package sandstormhttpbridge

import (
	"context"
	"net"

	bridge "zenhack.net/go/sandstorm/capnp/sandstormhttpbridge"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/rpc"
)

func connect(ctx context.Context, hooks *capnp.Client) (ret bridge.SandstormHttpBridge, err error) {
	conn, err := net.Dial("unix", "/tmp/sandstorm-api")
	if err != nil {
		return
	}
	var options *rpc.Options
	if hooks != nil {
		options = &rpc.Options{
			BootstrapClient: hooks,
		}
	}

	transport := rpc.NewStreamTransport(conn)
	ret.Client = rpc.NewConn(transport, options).Bootstrap(ctx)
	return
}

func ConnectWithHooks(ctx context.Context, hooks bridge.AppHooks) (bridge.SandstormHttpBridge, error) {
	return connect(ctx, hooks.Client)
}

func Connect(ctx context.Context) (bridge.SandstormHttpBridge, error) {
	return connect(ctx, nil)
}
