package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"capnproto.org/go/capnp/v3"
	httpcp "zenhack.net/go/sandstorm-next/capnp/http"
	"zenhack.net/go/sandstorm/exp/util/bytestream"
	"zenhack.net/go/util/exn"
)

type httpBridge struct {
	portNo       int
	roundTripper http.RoundTripper
	serverReady  bool
}

func (b *httpBridge) ensureServerReady() error {
	if b.serverReady {
		return nil
	}
	conn, err := exponentialBackoff(func() (net.Conn, error) {
		return net.Dial("tcp", b.netAddr())
	})
	if err != nil {
		return err
	}
	conn.Close()
	b.serverReady = true
	return nil
}

func (b *httpBridge) netAddr() string {
	return net.JoinHostPort("127.0.0.1", strconv.Itoa(b.portNo))
}

func (b *httpBridge) Request(ctx context.Context, p httpcp.Server_request) error {
	b.ensureServerReady()

	p.Go()
	return exn.Try0(func(throw func(error)) {
		// First, copy/translate the parameters into an http.Request:
		var req http.Request

		args := p.Args()
		results, err := p.AllocResults()
		throw(err)

		cpReq, err := args.Request()
		throw(err)

		path, err := cpReq.Path()
		throw(err)
		req.URL, err = url.ParseRequestURI(path)
		throw(err)
		req.URL.Host = fmt.Sprintf("127.0.0.1:%v", b.portNo)
		req.URL.Scheme = "http"

		method, err := cpReq.Method()
		throw(err)
		req.Method = method

		headers, err := cpReq.Headers()
		throw(err)
		req.Header = make(http.Header, headers.Len())
		for i := 0; i < headers.Len(); i++ {
			h := headers.At(i)
			key, err := h.Key()
			throw(err)
			val, err := h.Value()
			throw(err)
			req.Header.Set(key, val)
		}
		req.Host = req.Header.Get("Host")

		responder := args.Responder().AddRef()

		r, w := bytestream.Pipe()
		results.SetRequestBody(w)
		req.Body = r

		// Ok, request is all set up. Fork off a goroutine to actually send it and
		// copy the response back, so that the caller can start pushing data into
		// requestBody
		go func() {
			defer responder.Release()
			resp, err := b.roundTripper.RoundTrip((&req).WithContext(context.TODO()))
			var (
				fut httpcp.Responder_respond_Results_Future
				rel capnp.ReleaseFunc
			)
			if err != nil {
				// Push an error response to the caller:
				fut, rel = responder.Respond(context.TODO(),
					func(p httpcp.Responder_respond_Params) error {
						p.SetStatus(500)
						return nil
					})
			} else {
				defer resp.Body.Close()

				// Now push the response back to our caller:
				fut, rel = responder.Respond(context.TODO(), func(p httpcp.Responder_respond_Params) error {
					return exn.Try0(func(throw func(error)) {
						p.SetStatus(uint16(resp.StatusCode))
						totalHeaders := 0
						for _, vs := range resp.Header {
							totalHeaders += len(vs)
						}
						headers, err := p.NewHeaders(int32(totalHeaders))
						throw(err)
						i := 0
						for k, vs := range resp.Header {
							for _, v := range vs {
								h := headers.At(i)
								i++
								throw(h.SetKey(k))
								throw(h.SetValue(v))
							}
						}
					})
				})
			}
			defer rel()
			sink := fut.Sink()
			responseBody := bytestream.ToWriteCloser(context.TODO(), sink)
			defer responseBody.Close()
			if err != nil {
				fmt.Fprintf(
					responseBody,
					"Internal Server Error: %v\n",
					err,
				)
			} else {
				n, err := io.Copy(responseBody, resp.Body)
				if err != nil {
					log.Printf("Error copying response body after %v bytes: %v", n, err)
				}
			}
		}()
	})
}

// Try calling f at exponentially increasing intervals until either it returns a nil error,
// or the length of the interval exceeds 30 seconds.
func exponentialBackoff[T any](f func() (T, error)) (val T, err error) {
	delay := time.Millisecond
	limit := 30 * time.Second
	for {
		val, err = f()
		if err == nil || delay > limit {
			return
		}
		log.Printf("Error %v\n; trying again in %v\n", err, delay)
		time.Sleep(delay)
		delay *= 2
	}
}
