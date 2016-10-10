package ip

import (
	"golang.org/x/net/context"
	"io"
	"net"
	"time"
	"zenhack.net/go/sandstorm/capnp/ip"
	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/internal/errors"
)

type conn struct {
	// Accessed only by netConn:
	send          util.ByteStream
	recvRead      io.ReadCloser
	local, remote net.Addr

	// Accessed only by bsConn:
	recvWrite io.WriteCloser

	// Accessed by both
	callCtx, closeCtx context.Context
	closeFn           context.CancelFunc
}

type netConn conn
type bsConn conn

func (c *bsConn) Done(args util.ByteStream_done) error {
	println("done")
	c.closeFn()
	println("error: " + c.closeCtx.Err().Error())
	return c.recvWrite.Close()
}

func (c *bsConn) Write(p util.ByteStream_write) error {
	data, err := p.Params.Data()
	if err != nil {
		return err
	}
	_, err = c.recvWrite.Write(data)
	return err
}

func (c *bsConn) ExpectSize(args util.ByteStream_expectSize) error {
	return nil
}

func (c *netConn) Close() error {
	c.closeFn()
	c.send.Done(c.callCtx, func(p util.ByteStream_done_Params) error {
		return nil
	})
	return c.recvRead.Close()
}

func (c *netConn) Read(p []byte) (n int, err error) {
	if c.closeCtx.Err() != nil {
		return 0, io.EOF
	}
	return c.recvRead.Read(p)
}

func (c *netConn) Write(p []byte) (n int, err error) {
	if c.closeCtx.Err() != nil {
		return 0, io.EOF
	}
	c.send.Write(c.callCtx, func(args util.ByteStream_write_Params) error {
		args.SetData(p)
		return nil
	})
	return len(p), nil
}

func (c *netConn) LocalAddr() net.Addr  { return c.local }
func (c *netConn) RemoteAddr() net.Addr { return c.remote }

func (c *netConn) SetDeadline(t time.Time) error      { return errors.NotImplemented }
func (c *netConn) SetWriteDeadline(t time.Time) error { return errors.NotImplemented }
func (c *netConn) SetReadDeadline(t time.Time) error  { return errors.NotImplemented }

func connect(ctx context.Context, port ip.TcpPort, local, remote net.Addr) *netConn {
	pipeRead, pipeWrite := io.Pipe()
	ret := &conn{
		callCtx:  ctx,
		recvRead: pipeRead,
		local:    local,
		remote:   remote,

		recvWrite: pipeWrite,
	}
	ret.closeCtx, ret.closeFn = context.WithCancel(ctx)
	ret.send = port.Connect(ctx, func(p ip.TcpPort_connect_Params) error {
		p.SetDownstream(util.ByteStream_ServerToClient((*bsConn)(ret)))
		return nil
	}).Upstream()
	return (*netConn)(ret)
}
