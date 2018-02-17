package websession

// Implement http.ResponseWriter on top of WebSession.

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

var (
	successCodeMap = map[int]websession.WebSession_Response_SuccessCode{
		200: websession.WebSession_Response_SuccessCode_ok,
		201: websession.WebSession_Response_SuccessCode_created,
		202: websession.WebSession_Response_SuccessCode_accepted,
		204: websession.WebSession_Response_SuccessCode_noContent,
		206: websession.WebSession_Response_SuccessCode_partialContent,
		207: websession.WebSession_Response_SuccessCode_multiStatus,
		304: websession.WebSession_Response_SuccessCode_notModified,
	}

	clientErrorCodeMap = map[int]websession.WebSession_Response_ClientErrorCode{
		400: websession.WebSession_Response_ClientErrorCode_badRequest,
		403: websession.WebSession_Response_ClientErrorCode_forbidden,
		404: websession.WebSession_Response_ClientErrorCode_notFound,
		405: websession.WebSession_Response_ClientErrorCode_methodNotAllowed,
		406: websession.WebSession_Response_ClientErrorCode_notAcceptable,
		409: websession.WebSession_Response_ClientErrorCode_conflict,
		410: websession.WebSession_Response_ClientErrorCode_gone,
		412: websession.WebSession_Response_ClientErrorCode_preconditionFailed,
		413: websession.WebSession_Response_ClientErrorCode_requestEntityTooLarge,
		414: websession.WebSession_Response_ClientErrorCode_requestUriTooLong,
		415: websession.WebSession_Response_ClientErrorCode_unsupportedMediaType,
		418: websession.WebSession_Response_ClientErrorCode_imATeapot,
		422: websession.WebSession_Response_ClientErrorCode_unprocessableEntity,
	}
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

// Set fields common to all possible websession responses.
func (w *basicResponseWriter) writeHeaderCommon() {
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

	// TODO: cachePolicy
}

func (w *basicResponseWriter) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		panic("WriteHeader called twice!")
	}
	w.statusCode = statusCode
	w.writeHeaderCommon()
	switch statusCode {
	case 200, 201, 202:
		w.response.SetContent()
		content := w.response.Content()
		content.SetStatusCode(successCodeMap[statusCode])

		body := content.Body()
		body.SetStream(w.cancel)

		// TODO:
		//
		// * encoding
		// * language
		// * mimeType
		// * eTag
		// * disposition
	// TODO: other status codes.
	default:
		panic(fmt.Sprintf("Status not implemented: %d", statusCode))
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
