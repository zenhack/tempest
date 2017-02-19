package websession

import (
	"golang.org/x/net/context"
	"net/http"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/websession/pogs"
	"zombiezen.com/go/capnproto2/pogs"
)

type testCase struct {
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
		request: GetHeadReq{
			Path:       "/",
			IgnoreBody: false,
		},
	},
}
