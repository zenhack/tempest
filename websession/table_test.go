package websession

import (
	"bytes"
	"fmt"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"zenhack.net/go/sandstorm/capnp/grain"
	util_capnp "zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/util"
	"zenhack.net/go/sandstorm/websession/pogs"
	"zombiezen.com/go/capnproto2/pogs"
)

type testCase struct {
	name     string
	handler  http.Handler
	request  testRequest
	response testResponse
}

type testRequest interface {
	Call(ctx context.Context, ws websession.WebSession) (testResponse, error)
}

type testResponse struct {
	resp       websession_pogs.Response
	streamBody []byte
}

type GetHeadReq websession_pogs.Get_args
type PostReq websession_pogs.Post_args

func doCall(wsCtx *websession_pogs.Context, f func() (websession.WebSession_Response, error)) (testResponse, error) {
	r, w := io.Pipe()
	buf := &bytes.Buffer{}
	done := make(chan struct{})
	go func() {
		io.Copy(buf, r)
		done <- struct{}{}
	}()
	wsCtx.ResponseStream = util_capnp.ByteStream_ServerToClient(
		&util.WriteCloserByteStream{w},
	)
	resp, err := f()
	if err != nil {
		return testResponse{}, err
	}
	goResp := testResponse{
		resp: websession_pogs.Response{},
	}
	err = pogs.Extract(&goResp.resp, websession.WebSession_Response_TypeID, resp.Struct)
	<-done
	goResp.streamBody = buf.Bytes()
	return goResp, err
}

func (req GetHeadReq) Call(ctx context.Context, ws websession.WebSession) (testResponse, error) {
	return doCall(&req.Context, func() (websession.WebSession_Response, error) {
		return ws.Get(ctx, func(p websession.WebSession_get_Params) error {
			return pogs.Insert(websession.WebSession_get_Params_TypeID, p.Struct, req)
		}).Struct()
	})
}

func (req PostReq) Call(ctx context.Context, ws websession.WebSession) (testResponse, error) {
	return doCall(&req.Context, func() (websession.WebSession_Response, error) {
		return ws.Post(ctx, func(p websession.WebSession_post_Params) error {
			return pogs.Insert(websession.WebSession_post_Params_TypeID, p.Struct, req)
		}).Struct()
	})
}

var testCases = []testCase{
	{
		name:    "Simple 200 OK",
		request: GetHeadReq{Path: "/", IgnoreBody: false},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		}),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte{},
		},
	},
	{
		name:    "Hello, World!",
		request: GetHeadReq{Path: "/", IgnoreBody: false},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Hello, World!"))
		}),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte("Hello, World!"),
		},
	},
	{
		name:    "Basic routing",
		request: GetHeadReq{Path: "/foo", IgnoreBody: false},
		handler: func() http.Handler {
			mux := http.NewServeMux()
			mux.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte("Visited Foo"))
			})
			mux.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte("Visited Bar"))
			})
			return mux
		}(),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte("Visited Foo"),
		},
	},
	{
		name:    "GET vs HEAD (GET)",
		request: GetHeadReq{Path: "/", IgnoreBody: false},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(req.Method))
		}),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte("GET"),
		},
	},
	{
		name:    "GET vs HEAD (HEAD)",
		request: GetHeadReq{Path: "/", IgnoreBody: true},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(req.Method))
		}),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte("HEAD"),
		},
	},
	{
		name: "POST",
		request: PostReq{
			Path: "/post-test",
			Content: websession_pogs.PContent{
				MimeType: "text/plain",
				Content:  []byte("The Data."),
				Encoding: "identity",
			},
		},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/post-test" {
				w.Write([]byte("bad path"))
				return
			}
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				return
				w.Write([]byte("ReadAll error: " + err.Error()))
			}
			contentType := req.Header.Get("Content-Type")
			if contentType != "text/plain" {
				fmt.Printf("Unexpected Content-Type: %q", contentType)
				return
			}
			contentEncoding := req.Header.Get("Content-Encoding")
			if contentEncoding != "identity" {
				fmt.Printf("Unexpected Content-Encoding: %q", contentEncoding)
			}
			if string(body) != "The Data." {
				fmt.Fprintf(w, "Unexpected request body: %q", body)
				return
			}
			w.Write([]byte("ok"))
		}),
		response: testResponse{
			resp:       mkOkResponse(),
			streamBody: []byte("ok"),
		},
	},
}

// Return a bare-bones capnp response object.
func mkOkResponse() websession_pogs.Response {
	return websession_pogs.Response{
		Which: websession.WebSession_Response_Which_content,
		Content: websession_pogs.Response_content{
			StatusCode: websession.WebSession_Response_SuccessCode_ok,
			Body: websession_pogs.Response_content_body{
				Which:  websession.WebSession_Response_content_body_Which_stream,
				Stream: util_capnp.Handle_ServerToClient(struct{}{}),
			},
		},
	}

}

func TestTable(t *testing.T) {
	for _, v := range testCases {
		ctx := context.TODO()
		handlerUiView := grain.UiView_ServerToClient(FromHandler(v.handler))
		results, err := handlerUiView.NewSession(
			ctx,
			func(p grain.UiView_newSession_Params) error {
				// TODO: We'll need to pass in a SessionContext if
				// we're to test any related functionality.
				return nil
			}).Struct()
		uiSession := results.Session()
		handlerWS := websession.WebSession{Client: uiSession.Client}
		if err != nil {
			t.Errorf("Error in NewSession in table test case %q: %v",
				v.name, err)
			continue
		}
		resp, err := v.request.Call(ctx, handlerWS)
		if err != nil {
			t.Errorf("Error in v.reqeust.Call in table test case %q: %v",
				v.name, err)
			continue
		}
		if !responseEq(v.response, resp) {
			t.Errorf("Unexpected response in table test case %q", v.name)
			pretty.Ldiff(t, resp, v.response)
			continue
		}
	}
}

func responseEq(expected, actual testResponse) bool {
	// reflect.DeepEqual does what we want for *most* of the data,
	// but not the interfaces. So, we set those to nil before the check,
	// then restore them.
	eStream := &expected.resp.Content.Body.Stream
	aStream := &actual.resp.Content.Body.Stream
	eClient := eStream.Client
	aClient := aStream.Client
	eStream.Client = nil
	aStream.Client = nil
	defer func() {
		eStream.Client = eClient
		aStream.Client = aClient
	}()

	return reflect.DeepEqual(expected, actual) &&
		// either the clients should both be nil, or neither of them
		// should be:
		(eClient == nil) == (aClient == nil)
}
