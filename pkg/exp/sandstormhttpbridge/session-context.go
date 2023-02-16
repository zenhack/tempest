package sandstormhttpbridge

import (
	"net/http"

	"zenhack.net/go/tempest/capnp/grain"
	"zenhack.net/go/tempest/capnp/powerbox"
	sandstormhttpbridge "zenhack.net/go/tempest/capnp/sandstorm-http-bridge"
)

type hasId interface {
	SetId(string) error
}

func setId(p hasId, req *http.Request) {
	val := req.Header.Get("X-Sandstorm-Session-Id")
	p.SetId(val)
}

func GetSessionContext(
	bridge sandstormhttpbridge.SandstormHttpBridge,
	req *http.Request,
) grain.SessionContext {
	res, _ := bridge.GetSessionContext(
		req.Context(),
		func(p sandstormhttpbridge.SandstormHttpBridge_getSessionContext_Params) error {
			setId(p, req)
			return nil
		})
	return res.Context()
}

func GetSessionRequest(
	bridge sandstormhttpbridge.SandstormHttpBridge,
	req *http.Request,
) powerbox.PowerboxDescriptor_List {
	resFuture, _ := bridge.GetSessionRequest(
		req.Context(),
		func(p sandstormhttpbridge.SandstormHttpBridge_getSessionRequest_Params) error {
			setId(p, req)
			return nil
		})
	res, err := resFuture.Struct()
	if err != nil {
		panic(err)
	}
	ret, err := res.RequestInfo()
	if err != nil {
		panic(err)
	}
	return ret
}
