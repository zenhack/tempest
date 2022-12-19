package websession

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

type handler struct {
	session websession.WebSession
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.doGet(w, req, true)
	case "HEAD":
		h.doGet(w, req, false)
	default:
		panic("TODO")
	}
}

func (h handler) doGet(w http.ResponseWriter, req *http.Request, ignoreBody bool) {
	responseStreamServer := newResponseStreamImpl(w)
	responseStreamClient := util.ByteStream_ServerToClient(responseStreamServer)

	respFut, rel := h.session.Get(req.Context(), func(p websession.WebSession_get_Params) error {
		if err := p.SetPath(req.RequestURI); err != nil {
			return err
		}
		wsCtx, err := p.Context()
		if err != nil {
			return err
		}

		if err = populateContext(wsCtx, req, responseStreamClient); err != nil {
			return err
		}

		p.SetIgnoreBody(ignoreBody)
		return nil
	})
	defer rel()
	resp, err := respFut.Struct()
	if err != nil {
		replyErr(w, err)
		return
	}
	translateResponse(
		req.Context(),
		w,
		resp,
		responseStreamServer,
	)
}

func translateResponse(
	ctx context.Context,
	w http.ResponseWriter,
	resp websession.WebSession_Response,
	responseStream *responseStreamImpl,
) {
	status, err := responseStatus(resp)
	if err != nil {
		replyErr(w, err)
		close(responseStream.ready)
		return
	}

	if resp.Which() == websession.WebSession_Response_Which_content {
		content := resp.Content()
		body := content.Body()
		if body.Which() == websession.WebSession_Response_content_body_Which_stream {
			defer body.Stream().Release()
			responseStream.used = true

			select {
			case <-ctx.Done():
				return
			case size, ok := <-responseStream.size:
				if ok {
					w.Header().Set("Content-Length", strconv.FormatUint(size, 10))
				}
			}
			if err := populateResponseHeaders(w.Header(), resp); err != nil {
				replyErr(w, err)
				return
			}
			w.WriteHeader(status)
			close(responseStream.ready)
			select {
			case <-ctx.Done():
			case <-responseStream.done:
			case <-responseStream.shutdown:
			}
			return
		}
	}

	close(responseStream.ready)
	if err := populateResponseHeaders(w.Header(), resp); err != nil {
		replyErr(w, err)
		return
	}
	data, err := responseBodyBytes(resp)
	if err != nil {
		replyErr(w, err)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

func responseStatus(resp websession.WebSession_Response) (int, error) {
	switch resp.Which() {
	case websession.WebSession_Response_Which_content:
		successCode := resp.Content().StatusCode()
		status, ok := successCodeStatuses[successCode]
		if ok {
			return status, nil
		}
		return 0, fmt.Errorf("Unknown success code enumerant: %v", successCode)
	case websession.WebSession_Response_Which_noContent:
	case websession.WebSession_Response_Which_preconditionFailed:
	case websession.WebSession_Response_Which_redirect:
	case websession.WebSession_Response_Which_clientError:
		errorCode := resp.ClientError().StatusCode()
		status, ok := clientErrorCodeStatuses[errorCode]
		if ok {
			return status, nil
		}
		return 0, fmt.Errorf("Unknown error code enumerant: %v", errorCode)
	case websession.WebSession_Response_Which_serverError:
	default:
		return 0, fmt.Errorf("Unknown response variant: %v", resp.Which())
	}
	panic("TODO")
}

func responseBodyBytes(resp websession.WebSession_Response) ([]byte, error) {
	panic("TODO")
}

func populateResponseHeaders(h http.Header, r websession.WebSession_Response) error {
	// TODO: setCookies
	// TODO: cachePolicy
	// TODO: union
	// TODO: additionalHeaders
	panic("TODO")
}

// replyErr responds to the request with a 500 status and the given error.
func replyErr(w http.ResponseWriter, err error) {
	hdr := w.Header()

	// Delete anty header we may have set when trying to build a proper response.
	for k, _ := range hdr {
		delete(hdr, k)
	}

	body := []byte(err.Error())

	hdr.Set("Content-Type", "text/plain")
	hdr.Set("Content-Length", strconv.Itoa(len(body)))

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// populateContext populates a websession context based on the request, using the supplied
// value for the responseStream field. The reference to responseStream is stolen.
func populateContext(wsCtx websession.WebSession_Context, req *http.Request, responseStream util.ByteStream) error {
	// TODO: cookies
	// TODO: responseStream
	// TODO: accept
	// TODO: acceptEncoding
	// TODO: eTagPrecondition
	// TODO: additionalHeaders
	panic("TODO")
}
