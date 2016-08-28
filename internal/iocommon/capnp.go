package iocommon

import (
	"io"
	"zombiezen.com/go/capnproto2"
)

type closersClient struct {
	capnp.Client
	closers []io.Closer
}

func (c *closersClient) Close() error {
	err := c.Client.Close()
	for i := range c.closers {
		newErr := c.closers[i].Close()
		if err == nil {
			err = newErr
		}
	}
	return err
}

// WithClosers returns a capnp.Client that when closed, closes both
// `client` and each io.Closer in closers. Calls to `Call` are passed through
// to the underlying client
//
// The new Client's close method will always call Close on each closer. If
// any closer's Close() returns a non-nil error, the first such error is
// returned.
func WithClosers(client capnp.Client, closers ...io.Closer) capnp.Client {
	if oldClient, ok := client.(*closersClient); ok {
		// Avoid making deep stacks of these by introspecting; if client
		// is already a closersClient, we copy the data and add the new
		// Closers.

		newClosers := make([]io.Closer, len(oldClient.closers)+len(closers))
		copy(newClosers[:len(oldClient.closers)], oldClient.closers)
		copy(newClosers[len(oldClient.closers):], closers)
		return &closersClient{
			Client:  oldClient.Client,
			closers: newClosers,
		}
	} else {
		return &closersClient{
			Client:  client,
			closers: closers,
		}
	}

}
