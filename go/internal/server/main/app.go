package servermain

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/apex/log"
	httpcp "zenhack.net/go/sandstorm-next/capnp/http"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
	"zenhack.net/go/sandstorm/exp/util/bytestream"
	"zenhack.net/go/util/exn"
)

func ServeApp(lg log.Interface, c *container.Container, w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	SetAppHeaders(w)
	srv := httpcp.Server(c.Bootstrap)

	fut, rel := srv.Request(ctx, func(p httpcp.Server_request_Params) error {
		return exn.Try0(func(throw func(error)) {
			r, err := p.NewRequest()
			throw(err)

			throw(p.SetResponder(httpcp.Responder_ServerToClient(&responder{
				w:      w,
				cancel: cancel,
			})))
			throw(r.SetPath(req.RequestURI))
			throw(r.SetMethod(req.Method))
			totalHeaders := 0
			for _, vs := range req.Header {
				totalHeaders += len(vs)
			}
			headers, err := r.NewHeaders(int32(totalHeaders))
			throw(err)

			i := 0
			for k, vs := range req.Header {
				for _, v := range vs {
					h := headers.At(i)
					throw(h.SetKey(k))
					throw(h.SetValue(v))
					i++
				}
			}
		})
	})
	defer rel()
	bodyW := bytestream.ToWriteCloser(ctx, fut.RequestBody())
	go func() {
		_, err := io.Copy(bodyW, req.Body)
		if err != nil {
			lg.Infof("Error copying request body: %v", err)
			cancel()
		}
		err = bodyW.Close()
		if err != nil {
			lg.Infof("bodyW.Close(): %v", err)
			cancel()
		}
	}()
	_, err := fut.Struct()
	if err != nil {
		lg.Infof("Error in request(): %v", err)
		cancel()
	}
	<-ctx.Done()
}

type responder struct {
	w      http.ResponseWriter
	cancel context.CancelFunc
}

func (r *responder) Respond(ctx context.Context, p httpcp.Responder_respond) error {
	defer func() {
		r.w = nil
	}()
	if r.w == nil {
		return fmt.Errorf("respond() called twice.")
	}
	args := p.Args()
	results, err := p.AllocResults()
	if err != nil {
		return err
	}

	headers, err := args.Headers()
	if err != nil {
		return err
	}
	outHeaders := r.w.Header()
	for i := 0; i < headers.Len(); i++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		h := headers.At(i)
		k, err := h.Key()
		if err != nil {
			return err
		}
		v, err := h.Value()
		if err != nil {
			return err
		}
		outHeaders.Set(k, v)
	}
	r.w.WriteHeader(int(args.Status()))
	body := bytestream.FromWriteCloser(writeCancelCloser{
		Writer: r.w,
		Cancel: r.cancel,
	})
	results.SetSink(body)
	return nil
}

func (r *responder) Shutdown() {
	if r.w == nil {
		return
	}
	defer r.cancel()
	r.w.WriteHeader(http.StatusInternalServerError)
	r.w.Write([]byte("App bug: responder reference dropped before calling respond()"))
	r.w = nil
}

type writeCancelCloser struct {
	io.Writer
	Cancel context.CancelFunc
}

func (w writeCancelCloser) Close() error {
	w.Cancel()
	return nil
}
