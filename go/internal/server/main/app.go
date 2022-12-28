package servermain

import (
	"net/http"

	websessioncp "zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/exp/websession"
)

func ServeApp(webSession websessioncp.WebSession, w http.ResponseWriter, req *http.Request) {
	websession.Handler{
		Session: webSession,
	}.ServeHTTP(w, req)
}
