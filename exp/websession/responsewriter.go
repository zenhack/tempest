package websession

// Implement http.ResponseWriter on top of WebSession.

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"strconv"

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
//
// TODO: should we implement Hijacker? it may make sense not to; we can't actually
// not send the header.
type basicResponseWriter struct {
	// Status to send. 0 until WriteHeader() has been called.
	statusCode int

	// handle which cancels the request's Context.
	cancel util.Handle

	// The responseStream field in the websession request. This should only
	// be used for the expectSize() method; use bodyWriter for actually
	// writing data.
	responseStream util.ByteStream

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

		encoding := w.header.Get("Content-Encoding")
		if encoding != "" {
			content.SetEncoding(encoding)
		}
		language := w.header.Get("Content-Language")
		if language != "" {
			content.SetLanguage(language)
		}
		mimeType := w.header.Get("Content-Type")
		if mimeType != "" {
			content.SetMimeType(mimeType)
		}

		contentLength := w.header.Get("Content-Length")
		if contentLength != "" {
			length, err := strconv.ParseUint(contentLength, 10, 64)
			if err != nil {
				log.Print("Error parsing Content-Length set by handler:", err)
			} else {
				w.responseStream.ExpectSize(
					context.TODO(),
					func(params util.ByteStream_expectSize_Params) error {
						params.SetSize(length)
						return nil
					})
			}
		}

		// TODO: this doesn't seem to work. We have a failing test.
		contentDisposition := w.header.Get("Content-Disposition")
		if contentDisposition != "" {
			typ, params, err := mime.ParseMediaType(contentDisposition)
			if err != nil {
				log.Print("Error parsing response's Content-Disposition:", err)
			} else if typ == "inline" {
				content.Disposition().SetNormal()
			} else if typ == "attachment" {
				filename, ok := params["filename"]
				if !ok {
					log.Println("Warning: handler specified " +
						"Content-Disposition = attachment with no " +
						"file name.")
				}
				content.Disposition().SetDownload(filename)
			} else {
				log.Println("Unsupported Content-Disposition type:", typ)
			}
		}

		// TODO: eTag
	case 204, 205:
		w.response.SetNoContent()
		noContent := w.response.NoContent()
		noContent.SetShouldResetForm(statusCode == 205)
		// TODO: eTag
	// TODO:
	//
	// * preConditionFailed
	// * redirect
	// * clientError
	// * serverError
	// * unsupported status
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
