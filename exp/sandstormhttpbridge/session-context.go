package sandstormhttpbridge

import (
	"net/http"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/capnp/sandstormhttpbridge"
)

func GetSessionContext(
	bridge sandstormhttpbridge.SandstormHttpBridge,
	req *http.Request,
) grain.SessionContext {
	res, _ := bridge.GetSessionContext(
		req.Context(),
		func(p sandstormhttpbridge.SandstormHttpBridge_getSessionContext_Params) error {
			p.SetId(req.Header.Get("X-Sandstorm-Session-Id"))
			return nil
		})
	return res.Context()
}
