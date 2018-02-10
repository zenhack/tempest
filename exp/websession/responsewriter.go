package websession

// Implement http.ResponseWriter on top of WebSession.

import (
	"io"
	"log"
	"net/http"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

// ResponseWriter used for WebSession methods that just return a WebSession_Response,
// e.g. get, put, and post, but not postStreaming or openWebSocket.
type basicResponseWriter struct {
	// Status to send. 0 until WriteHeader() has been called.
	statusCode int

	// handle which cancels the request's Context.
	cancel util.Handle

	header     http.Header
	bodyWriter io.Writer
	response   websession.WebSession_Response
}

func (w *basicResponseWriter) Header() http.Header {
	return w.header
}

func (w *basicResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		panic("WriteHeader called twice!")
	}
	w.statusCode = statusCode
	switch statusCode {
	case 200:
		w.response.SetContent()
		content := w.response.Content()
		content.SetStatusCode(websession.WebSession_Response_SuccessCode_ok)

		body := content.Body()
		body.SetStream(w.cancel)

		// A bit of a hack; afaik there's no public
		// function for parsing a cookie, but we can
		// abuse http.Response to do it.
		cookies := (&http.Response{Header: w.header}).Cookies()

		setCookies, err := w.response.NewSetCookies(int32(len(cookies)))
		if err != nil {
			// FIXME figure out the appropriate way to handle this.
			panic("ERROR: error allocating space for SetCookies.")
		}

		for i, cookie := range cookies {
			// Note that sandstorm sets all cookies as secure.

			setCookie := setCookies.At(i)
			setCookie.SetName(cookie.Name)
			setCookie.SetValue(cookie.Value)
			if cookie.Path != "" {
				setCookie.SetPath(cookie.Path)
			}

			if cookie.MaxAge != 0 && !cookie.Expires.IsZero() {
				log.Println("Warning: both cookie.MaxAge and cookie.Expires " +
					"were set; preferring MaxAge.")
				// TODO: do we want to be more pedantic about this?
			}

			if cookie.MaxAge < 0 {
				setCookie.Expires().SetRelative(0)
			} else if cookie.MaxAge > 0 {
				setCookie.Expires().SetRelative(uint64(cookie.MaxAge))
			} else if !cookie.Expires.IsZero() {
				setCookie.Expires().SetAbsolute(cookie.Expires.Unix())
			} else {
				setCookie.Expires().SetNone()
			}
		}

		// TODO:
		//
		// * setCookies
		// * cachePolicy
		// * encoding
		// * language
		// * mimeType
		// * eTag
		// * disposition
	// TODO: other status codes.
	default:
		panic("Not implemented")
	}

	// TODO:
	//
	// * additionalHeaders
}

func (w *basicResponseWriter) Write(data []byte) (n int, err error) {
	if w.statusCode == 0 {
		w.WriteHeader(200)
	}
	return w.bodyWriter.Write(data)
}
