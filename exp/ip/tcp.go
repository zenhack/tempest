package ip

/*
import (
	"context"
	"io"

	"zenhack.net/go/sandstorm/capnp/ip"
	"zenhack.net/go/sandstorm/exp/util/bytestream"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/server"
)

func ConnectTCP(ctx context.Context, policy *server.Policy, port ip.TcpPort) io.ReadWriteCloser {
	fromThem, toUs := bytestream.Pipe(policy)
	res, release := port.Connect(ctx, func(p ip.TcpPort_connect_Params) error {
		p.SetDownstream(toUs)
		return nil
	})
	toThem := res.Upstream()
	return tcpPort{
		toThem:   bytestream.ToWriteCloser(toThem),
		fromThem: fromThem,
		release:  release,
	}
}

type tcpPort struct {
	closed   bool
	toThem   io.WriteCloser
	fromThem io.ReadCloser
	release  capnp.ReleaseFunc
}

func (port tcpPort) Read(p []byte) (n int, err error)  { return port.fromThem.Read(p) }
func (port tcpPort) Write(p []byte) (n int, err error) { return port.toThem.Write(p) }

func (port tcpPort) Close() error {
	if port.closed {
		return nil
	}
	port.toThem.Close()
	port.fromThem.Close()
	port.release()
	port.closed = true
}
*/
