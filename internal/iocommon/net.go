package iocommon

import (
	"io"
	"net"
	"time"
	"zenhack.net/go/sandstorm/internal/errors"
)

// Implementation of net.Addr using hard-coded values
type HardCodedAddr struct {
	Net, Addr string
}

func (a *HardCodedAddr) Network() string {
	return a.Net
}

func (a *HardCodedAddr) String() string {
	return a.Addr
}

// A net.Conn that wraps an io.ReadWriteCloser.
//
// RemoteAddr and LocalAddr are just getters for the corresponding fields.
//
// XXX: The Set*Deadline methods are noops; This has some serious security
// implications. They return errors at least, but clients should be careful
// to check errors.
type RWCConn struct {
	io.ReadWriteCloser
	Remote net.Addr
	Local  net.Addr
}

func (c *RWCConn) RemoteAddr() net.Addr {
	return c.Remote
}

func (c *RWCConn) LocalAddr() net.Addr {
	return c.Local
}

func (c *RWCConn) SetDeadline(t time.Time) error {
	return errors.NotImplemented
}

func (c *RWCConn) SetReadDeadline(t time.Time) error {
	return errors.NotImplemented
}

func (c *RWCConn) SetWriteDeadline(t time.Time) error {
	return errors.NotImplemented
}
