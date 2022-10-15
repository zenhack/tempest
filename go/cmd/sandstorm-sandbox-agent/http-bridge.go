package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	httpcp "zenhack.net/go/sandstorm-next/capnp/http"
	"zenhack.net/go/sandstorm/exp/util/bytestream"
	"zenhack.net/go/util/exn"
)

type httpBridge struct {
	portNo       int
	roundTripper http.RoundTripper
}

func (b httpBridge) Request(ctx context.Context, p httpcp.Server_request) error {
	p.Ack()
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
			resp, err := b.roundTripper.RoundTrip(&req)
			if err != nil {
				// TODO: handle errors by sending an error response.
				panic(err)
			}
			defer resp.Body.Close()

			// Now push the response back to our caller:
			fut, rel := responder.Respond(context.TODO(), func(p httpcp.Responder_respond_Params) error {
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
			defer rel()
			sink := fut.Sink()
			responseBody := bytestream.ToWriteCloser(context.TODO(), sink)
			n, err := io.Copy(responseBody, resp.Body)
			if err != nil {
				log.Printf("Error copying response body after %v bytes: %v", n, err)
			}
		}()
	})
}
