package websession

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/tj/assert"
	utilcp "zenhack.net/go/tempest/capnp/util"
	websession "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/go/pkg/exp/util/handle"
	"zenhack.net/go/util"
)

func TestGetBody(t *testing.T) {
	t.Parallel()

	expected := "Expected Body"

	cases := []testWebSessionImpl{
		{
			expectedBody:   expected,
			stream:         false,
			callExpectSize: false,
		},
		{
			expectedBody:   expected,
			stream:         true,
			callExpectSize: true,
		},
		{
			expectedBody:   expected,
			stream:         true,
			callExpectSize: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("TestGetBody stream=%v callExpectSize=%v", c.stream, c.callExpectSize),
			func(t *testing.T) {
				t.Parallel()

				rec := doRequest(c, httptest.NewRequest("GET", "/expected-body", nil))

				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, expected, rec.Body.String())
				if c.callExpectSize {
					assert.Equal(t,
						strconv.Itoa(len(expected)),
						rec.Result().Header.Get("Content-Length"),
					)
				} else {
					assert.Equal(t, "", rec.Result().Header.Get("Content-Length"))
				}
			})
	}

}

func TestGetPath(t *testing.T) {
	t.Parallel()
	expected := "path-body/example"

	rec := doRequest(
		testWebSessionImpl{},
		httptest.NewRequest("GET", "/"+expected, nil),
	)
	assert.Equal(t, expected, rec.Body.String())
}

func doRequest(t testWebSessionImpl, req *http.Request) *httptest.ResponseRecorder {
	client := websession.WebSession_ServerToClient(t)
	defer client.Release()
	h := Handler{Session: client}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

type testWebSessionImpl struct {
	expectedBody string

	stream         bool
	callExpectSize bool
}

func (t testWebSessionImpl) Get(ctx context.Context, p websession.WebSession_get) error {
	args := p.Args()
	path, err := args.Path()
	util.Chkfatal(err)
	wsCtx, err := args.Context()
	util.Chkfatal(err)
	response, err := p.AllocResults()
	util.Chkfatal(err)

	var actualBody string
	switch {
	case path == "expected-body":
		actualBody = t.expectedBody
	case strings.HasPrefix(path, "path-body/"):
		actualBody = path
	default:
		panic("Unexpected path: " + path)
	}

	response.SetContent()
	content := response.Content()
	content.SetStatusCode(websession.WebSession_Response_SuccessCode_ok)
	content.SetMimeType("text/plain")

	body := content.Body()
	if t.stream {
		ctx, hndl := handle.WithCancel(context.Background())
		body.SetStream(hndl)
		responseStream := wsCtx.ResponseStream().AddRef()
		go func() {
			if t.callExpectSize {
				_, rel := responseStream.ExpectSize(ctx, func(p utilcp.ByteStream_expectSize_Params) error {
					p.SetSize(uint64(len(actualBody)))
					return nil
				})
				defer rel()
			}
			_, rel := responseStream.Write(ctx, func(p utilcp.ByteStream_write_Params) error {
				return p.SetData([]byte(actualBody))
			})
			defer rel()
			_, rel = responseStream.Done(ctx, nil)
			defer rel()
		}()
	} else {
		body.SetBytes([]byte(actualBody))
	}
	return nil
}

// TODO: implement these
func (testWebSessionImpl) Post(context.Context, websession.WebSession_post) error {
	return errUnimplemented
}
func (testWebSessionImpl) Delete(context.Context, websession.WebSession_delete) error {
	return errUnimplemented
}
func (testWebSessionImpl) Put(context.Context, websession.WebSession_put) error {
	return errUnimplemented
}
func (testWebSessionImpl) Patch(context.Context, websession.WebSession_patch) error {
	return errUnimplemented
}
func (testWebSessionImpl) PostStreaming(context.Context, websession.WebSession_postStreaming) error {
	return errUnimplemented
}
func (testWebSessionImpl) PutStreaming(context.Context, websession.WebSession_putStreaming) error {
	return errUnimplemented
}
func (testWebSessionImpl) OpenWebSocket(context.Context, websession.WebSession_openWebSocket) error {
	return errUnimplemented
}
func (testWebSessionImpl) Propfind(context.Context, websession.WebSession_propfind) error {
	return errUnimplemented
}
func (testWebSessionImpl) Proppatch(context.Context, websession.WebSession_proppatch) error {
	return errUnimplemented
}
func (testWebSessionImpl) Mkcol(context.Context, websession.WebSession_mkcol) error {
	return errUnimplemented
}
func (testWebSessionImpl) Copy(context.Context, websession.WebSession_copy) error {
	return errUnimplemented
}
func (testWebSessionImpl) Move(context.Context, websession.WebSession_move) error {
	return errUnimplemented
}
func (testWebSessionImpl) Lock(context.Context, websession.WebSession_lock) error {
	return errUnimplemented
}
func (testWebSessionImpl) Unlock(context.Context, websession.WebSession_unlock) error {
	return errUnimplemented
}
func (testWebSessionImpl) Acl(context.Context, websession.WebSession_acl) error {
	return errUnimplemented
}
func (testWebSessionImpl) Report(context.Context, websession.WebSession_report) error {
	return errUnimplemented
}
func (testWebSessionImpl) Options(context.Context, websession.WebSession_options) error {
	return errUnimplemented
}

var errUnimplemented = errors.New("Unimplemented")
