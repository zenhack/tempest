package websession

import (
	"bufio"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

type responseWriter struct {
	hijack   chan<- struct{}
	hijacked bool

	StatusCode int
	header     http.Header
	respond    chan<- *http.Response

	body    io.WriteCloser
	request *http.Request
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if !w.hijacked && w.StatusCode == 0 {
		w.WriteHeader(200)
	}
	return w.body.Write(p)
}

func (w *responseWriter) WriteHeader(status int) {
	w.StatusCode = status
	response := &http.Response{
		StatusCode: w.StatusCode,
		Header:     w.Header(),
	}

	if w.hijacked {
		response.Write(w.body)
	} else {
		w.respond <- response
	}
}
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	w.hijacked = true
	w.hijack <- struct{}{}
	clientConn, serverConn := net.Pipe()
	go func() {
		io.Copy(clientConn, w.request.Body)
		w.request.Body.Close()
	}()
	go func() {
		io.Copy(w.body, clientConn)
		w.body.Close()
	}()
	bufrw := &bufio.ReadWriter{
		Reader: bufio.NewReader(serverConn),
		Writer: bufio.NewWriter(serverConn),
	}
	return serverConn, bufrw, nil
}

// Run handler.ServeHTTP, and call buildResponse when the status and
// headers are ready. The handler is run in a separate gorotuine,
// and may continue writing data to the body during and after the
// call to buildResponse.
func runHandler(handler http.Handler, req *http.Request, buildResponse func(*http.Response)) {
	done := make(chan struct{}, 1)
	hijack := make(chan struct{}, 1)
	respond := make(chan *http.Response)

	respPipeReader, respPipeWriter := io.Pipe()

	w := &responseWriter{
		StatusCode: 0,
		header:     make(map[string][]string),
		hijack:     hijack,
		respond:    respond,
		body:       respPipeWriter,
		request:    req,
	}

	go func() {
		handler.ServeHTTP(w, req)
		done <- struct{}{}
		if !w.hijacked {
			respPipeWriter.Close()
		}
	}()
	select {
	case <-hijack:
		resp, err := http.ReadResponse(bufio.NewReader(respPipeReader), req)
		if err != nil {
			buildResponse(&http.Response{
				StatusCode: 500,
				Header:     make(map[string][]string),
				Body: ioutil.NopCloser(strings.NewReader(
					"Internal Server Error",
				)),
			})
			return
		}
		buildResponse(resp)
	case resp := <-respond:
		resp.Body = respPipeReader
		buildResponse(resp)
	case <-done:
		buildResponse(&http.Response{
			StatusCode: 200,
			Header:     w.Header(),
			Body:       respPipeReader,
		})
	}
}
