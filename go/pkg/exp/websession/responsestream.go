package websession

import (
	"context"
	"errors"
	"net/http"

	"zenhack.net/go/tempest/capnp/util"
)

var (
	errExpectSizeCalledLater = errors.New("expectSize() called after another method")
	errUnused                = errors.New("body = stream not set in in response")
	errDoneAlreadyCalled     = errors.New("done() already called")
)

// Implementation of ByteStream provided as Context.responseStream
type responseStreamImpl struct {
	// Closed when the Response has been returned. ready will block
	// until this has been closed. The schema suggests supplying a promise,
	// but go-capnp doesn't do that yet, so here we are.
	ready chan struct{}

	// The value passed to expectSize will be sent on this channel. The channel
	// will be closed just after the value was sent, or when it is determined that
	// no value will be sent (something other than expectSize() is called). This
	// can happen before ready is closed.
	size chan uint64

	// Is the response stream actually going to be used? I.e. was body = stream set
	// in the returned Response?
	//
	// this must be set to true before ready is closed if the response wants to use
	// the stream; otherwise methods will return errors.
	used bool

	// The ResponseWriter for the request.
	w http.ResponseWriter

	done     chan struct{} // Closed when Done() is called
	shutdown chan struct{} // Closed when all clients are dropped.

	// Indicates if it is too late to call expectSize(), since other methods have been called.
	tooLateForExpectSize bool
	// Indicates if done has already been called
	doneAlreadyCalled bool
}

func newResponseStreamImpl(w http.ResponseWriter) *responseStreamImpl {
	return &responseStreamImpl{
		w: w,

		size:     make(chan uint64, 1),
		ready:    make(chan struct{}),
		done:     make(chan struct{}),
		shutdown: make(chan struct{}),
	}
}

// waitReady waits until it is appropriate to service methods for the response. It returns an
// error if this should not be done.
func (r *responseStreamImpl) waitReady(ctx context.Context) error {
	select {
	case <-r.ready:
		if !r.used {
			return errUnused
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// commitSize closes r.size and records that it is too late to invoke ExpectSize().
// calling commitSize multiple times has no effect.
func (r *responseStreamImpl) commitSize() error {
	if r.tooLateForExpectSize {
		return errExpectSizeCalledLater
	}
	r.tooLateForExpectSize = true
	close(r.size)
	return nil
}

func (r *responseStreamImpl) ExpectSize(ctx context.Context, p util.ByteStream_expectSize) error {
	if !r.tooLateForExpectSize {
		r.size <- p.Args().Size()
	}
	return r.commitSize()
}

func (r *responseStreamImpl) Write(ctx context.Context, p util.ByteStream_write) error {
	r.commitSize()
	if err := r.waitReady(ctx); err != nil {
		return err
	}
	r.tooLateForExpectSize = true
	data, err := p.Args().Data()
	if err != nil {
		return err
	}
	_, err = r.w.Write(data)
	return err
}

func (r *responseStreamImpl) Done(ctx context.Context, _ util.ByteStream_done) error {
	r.commitSize()
	if err := r.waitReady(ctx); err != nil {
		return err
	}
	if r.doneAlreadyCalled {
		return errDoneAlreadyCalled
	}
	r.doneAlreadyCalled = true
	close(r.done)
	return nil
}

func (r *responseStreamImpl) Shutdown() {
	close(r.shutdown)
}
