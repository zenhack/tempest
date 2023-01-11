package servermain

import (
	"net/http"

	websessioncp "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/go/pkg/exp/websession"
)

func ServeApp(webSession websessioncp.WebSession, w http.ResponseWriter, req *http.Request) {
	websession.Handler{
		Session: webSession,
	}.ServeHTTP(w, req)
}
