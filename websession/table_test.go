package websession

import (
	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"net/http"
	"testing"
	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/websession/pogs"
	"zombiezen.com/go/capnproto2/pogs"
)

type testCase struct {
	name     string
	handler  http.Handler
	request  testRequest
	response websession_pogs.Response
}

type testRequest interface {
	Call(ctx context.Context, ws websession.WebSession) (websession_pogs.Response, error)
}

type GetHeadReq websession_pogs.Get_args

func (req GetHeadReq) Call(ctx context.Context, ws websession.WebSession) (websession_pogs.Response, error) {
	resp, err := ws.Get(ctx, func(p websession.WebSession_get_Params) error {
		return pogs.Insert(websession.WebSession_get_Params_TypeID, p.Struct, req)
	}).Struct()
	if err != nil {
		return websession_pogs.Response{}, err
	}
	goResp := websession_pogs.Response{}
	err = pogs.Extract(&goResp, websession.WebSession_Response_TypeID, resp.Struct)
	return goResp, err
}

var testCases = []testCase{
	{
		name: "Simple 200 OK",
		request: GetHeadReq{
			Path:       "/",
			IgnoreBody: false,
		},
		handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		}),
		response: websession_pogs.Response{
			Which: websession.WebSession_Response_Which_content,
			Content: websession_pogs.Response_content{
				StatusCode: websession.WebSession_Response_SuccessCode_ok,
				Body: websession_pogs.Response_content_body{
					Which:  websession.WebSession_Response_content_body_Which_stream,
					Stream: util.Handle_ServerToClient(struct{}{}),
				},
			},
		},
	},
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

func responseEq(expected, actual websession_pogs.Response) bool {
	return true
}
