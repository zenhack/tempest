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

/*
		appAddr, ok := apps[appName]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
			switch req.Method {
			case "GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS":
			default:
				badRequest(w, "Bad Method")
				return
			}
			var newUrl url.URL
			newUrl.Scheme = "http"
			newUrl.Host = "127.0.0.1" + appAddr
			newUrl.Path = req.URL.Path
			newUrl.RawPath = req.URL.RawPath
			newUrl.RawQuery = req.URL.RawQuery
			newReq := http.Request{
				Method: req.Method,
				URL:    &newUrl,
				Header: http.Header{},
				Body:   req.Body,
			}
			resp, err := http.DefaultTransport.RoundTrip(&newReq)
			if err != nil {
				log.Printf("Error connecting to app %q: %q", appName, err)
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error connecting to app."))
				return
			}
			switch resp.StatusCode {
			case 200, 201, 202, 024, 206, 207, 301, 302, 303, 304, 307, 308:
			case 400, 403, 404, 405, 406, 409, 410, 412, 413, 414, 415, 418, 422:
			case 500:
			default:
				switch {
				case resp.StatusCode >= 400 && resp.StatusCode < 500:
					badRequest(w, "Bad Request")
				default:
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
				}
				return
			}

			for header, sanitizer := range responseHeaderWhiteList {
				val, ok := resp.Header[header]
				if ok {
					w.Header()[header] = sanitizer(val)
				}
			}

			w.WriteHeader(resp.StatusCode)
			// TODO: Limit according to Content-Length & set a timeout.
			io.Copy(w, resp.Body)
}
*/
