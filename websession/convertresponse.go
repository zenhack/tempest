package websession

// This file contains code for converting an http.Response into a
// WebSession_Response. It's existence makes me(zenhack) sad.

import (
	"bytes"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	capnp_util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/util"
)

var (
	successCodeMap = map[int]capnp.WebSession_Response_SuccessCode{
		200: capnp.WebSession_Response_SuccessCode_ok,
		201: capnp.WebSession_Response_SuccessCode_created,
		202: capnp.WebSession_Response_SuccessCode_accepted,
		204: capnp.WebSession_Response_SuccessCode_noContent,
		206: capnp.WebSession_Response_SuccessCode_partialContent,
		207: capnp.WebSession_Response_SuccessCode_multiStatus,
		304: capnp.WebSession_Response_SuccessCode_notModified,
	}

	clientErrorCodeMap = map[int]capnp.WebSession_Response_ClientErrorCode{
		400: capnp.WebSession_Response_ClientErrorCode_badRequest,
		403: capnp.WebSession_Response_ClientErrorCode_forbidden,
		404: capnp.WebSession_Response_ClientErrorCode_notFound,
		405: capnp.WebSession_Response_ClientErrorCode_methodNotAllowed,
		406: capnp.WebSession_Response_ClientErrorCode_notAcceptable,
		409: capnp.WebSession_Response_ClientErrorCode_conflict,
		410: capnp.WebSession_Response_ClientErrorCode_gone,
		412: capnp.WebSession_Response_ClientErrorCode_preconditionFailed,
		413: capnp.WebSession_Response_ClientErrorCode_requestEntityTooLarge,
		414: capnp.WebSession_Response_ClientErrorCode_requestUriTooLong,
		415: capnp.WebSession_Response_ClientErrorCode_unsupportedMediaType,
		418: capnp.WebSession_Response_ClientErrorCode_imATeapot,
		422: capnp.WebSession_Response_ClientErrorCode_unprocessableEntity,
	}
)

// The main workhorse in this file; copies the data from goResponse into
// wsResponse.
//
// A bit of weirdness: we need both a context.Context and a websession context
// internally, but these are completely unrelated things. The arguments appear
// to have a symmetry that isn't real!
func buildCapnpResponse(goCtx context.Context, goResponse *http.Response,
	wsCtx *capnp.WebSession_Context, wsResponse *capnp.WebSession_Response) {

	// Copy goResponse.Body to w, and then close it. We *need* to do this in
	// all cases, even if we're just junking the body (in which case we use
	// ioutil.Discard).
	writeBody := func(w io.Writer) {
		io.Copy(w, goResponse.Body)
		goResponse.Body.Close()
	}

	status := goResponse.StatusCode

	// TODO: handle the fields outside of the union:
	//
	// * setCookies
	// * cachePoliciy
	// * additionalHeaders

	switch status {
	case 200, 201, 202:
		wsResponse.SetContent()
		capnpStatus := successCodeMap[status]
		wsResponse.Content().SetStatusCode(capnpStatus)
		wsResponse.Content().Body().
			SetStream(capnp_util.Handle_ServerToClient(struct{}{}))
		setContentHeaders(wsResponse, goResponse)
		go func() {
			w := util.ByteStreamWriteCloser{
				Ctx: goCtx,
				Obj: wsCtx.ResponseStream(),
			}
			writeBody(w)
			w.Close()
		}()
	case 204, 205:
		wsResponse.SetNoContent()
		wsResponse.NoContent().SetShouldResetForm(status == 205)
		// TODO: handle etags
		go writeBody(ioutil.Discard)
	case 304, 412:
		wsResponse.SetPreconditionFailed()
		// TODO: etags
	case 301, 302, 303, 307:
		// Redirects. Web-session.capnp talks about a 308, but I haven't found
		// any info about its semantics. "net/http" doesn't deifine a constant
		// for it, so we're just going to say to heck with it and not support
		// it for now.
		wsResponse.SetRedirect()
		wsResponse.Redirect().SetLocation(goResponse.Header.Get("Location"))
		switch status {
		case 301:
			wsResponse.Redirect().SetIsPermanent(true)
		case 302, 303, 307:
			wsResponse.Redirect().SetIsPermanent(false)
		}
		switch status {
		case 302, 303:
			wsResponse.Redirect().SetSwitchToGet(true)
		default:
			wsResponse.Redirect().SetSwitchToGet(false)
		}
		writeBody(ioutil.Discard)

	case 400, 403, 404, 405, 406, 409, 410, 413, 414, 415, 418:
		wsResponse.SetClientError()
		capnpStatus := clientErrorCodeMap[status]
		wsResponse.ClientError().SetStatusCode(capnpStatus)

		// We can't stream the response because the capnp schema doesn't support,
		// it, so we buffer it synchronously.
		buf := &bytes.Buffer{}
		writeBody(buf)
		wsResponse.ClientError().SetDescriptionHtml(buf.String())
	default:
		wsResponse.SetServerError()
		if status >= 500 && status < 600 {
			// The handler actually returned a 5xx; let them set the body.
			// As above, to do this we have to buffer and set it synchronously.
			buf := &bytes.Buffer{}
			writeBody(buf)
			wsResponse.ServerError().SetDescriptionHtml(buf.String())
		} else {
			// The client returned some status sandstorm doesn't support. In this
			// case we junk the body and set our own error page.
			wsResponse.ServerError().SetDescriptionHtml(
				"<!DOCTYPE html>\n" +
					"<html>\n" +
					"\t<head>\n" +
					"\t\t<title>Internal Server Error</title>\n" +
					"\t</head>\n" +
					"\t<body>\n" +
					"\t\t<h1>Internal Server Error</h1>\n" +
					"\t</body>\n" +
					"</html>\n")
			go writeBody(ioutil.Discard)
		}
	}
}

// Set all the headers that are present and accepted by sandstorm as part of the
// "content" variant of Response. Ignore any other headers.
//
// TODO:
//
// * ETag
// * Content-Length (call expectSize on body)
func setContentHeaders(wsResponse *capnp.WebSession_Response, goResponse *http.Response) {
	content := wsResponse.Content()
	hdr := goResponse.Header
	if encoding := hdr.Get("Content-Encoding"); encoding != "" {
		content.SetEncoding(encoding)
	}
	if language := hdr.Get("Content-Language"); language != "" {
		content.SetLanguage(language)
	}
	if mimeType := hdr.Get("Content-Type"); mimeType != "" {
		content.SetMimeType(mimeType)
	}
	if disposition := hdr.Get("Content-Disposition"); disposition != "" {
		typ, params, err := mime.ParseMediaType(disposition)
		if err != nil {
			if typ == "attachment" {
				content.Disposition().SetDownload(params["filename"])
			} else {
				content.Disposition().SetNormal()
			}
		}
	}
}
