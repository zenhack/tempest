package ip

import (
	"context"
	"io"

	"zenhack.net/go/tempest/capnp/ip"
	"zenhack.net/go/tempest/pkg/exp/util/bytestream"

	"capnproto.org/go/capnp/v3"
)

func ConnectTCP(ctx context.Context, port ip.TcpPort) io.ReadWriteCloser {
	fromThem, toUs := bytestream.Pipe()
	res, release := port.Connect(ctx, func(p ip.TcpPort_connect_Params) error {
		p.SetDownstream(toUs)
		return nil
	})
	toThem := res.Upstream()
	return tcpPort{
		toThem:   bytestream.ToWriteCloser(ctx, toThem),
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
	return nil
}
