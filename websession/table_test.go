package websession

import (
	"bytes"
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"reflect"
	"testing"
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

func (req GetHeadReq) Call(ctx context.Context, ws websession.WebSession) (testResponse, error) {
	r, w := io.Pipe()
	buf := &bytes.Buffer{}
	done := make(chan struct{})
	go func() {
		io.Copy(buf, r)
		done <- struct{}{}
	}()
	req.Context.ResponseStream = util_capnp.ByteStream_ServerToClient(
		&util.WriteCloserByteStream{w},
	)
	resp, err := ws.Get(ctx, func(p websession.WebSession_get_Params) error {
		return pogs.Insert(websession.WebSession_get_Params_TypeID, p.Struct, req)
	}).Struct()
	if err != nil {
		return testResponse{}, err
	}
	goResp := testResponse{
		resp: websession_pogs.Response{},
	}
	err = pogs.Extract(&goResp.resp, websession.WebSession_Response_TypeID, resp.Struct)
	r.Close()
	<-done
	goResp.streamBody = buf.Bytes()
	return goResp, err
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
		handlerWS := websession.WebSession_ServerToClient(FromHandler(ctx, v.handler))
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
