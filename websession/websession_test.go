package websession

import (
	"golang.org/x/net/context"
	"net/http"
	"testing"
	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/capnp/websession"
)

// Make sure concurrent connections are handled correctly. Sandstorm will call UiView's NewSession()
// before each request, so we just need to make sure that we actually return a *new* session (or
// client anyhow) each time.
func TestConcurrent(t *testing.T) {
	ctx := context.Background()
	mux := http.NewServeMux()

	theChan := make(chan struct{})
	done := make(chan struct{})

	mux.HandleFunc("/one", func(w http.ResponseWriter, req *http.Request) {
		<-theChan
		w.Write([]byte("One!\n"))
		done <- struct{}{}
	})
	mux.HandleFunc("/two", func(w http.ResponseWriter, req *http.Request) {
		theChan <- struct{}{}
		w.Write([]byte("Two!\n"))
		done <- struct{}{}
	})

	uiView := grain.UiView_ServerToClient(FromHandler(mux))

	get := func(path string) {
		sess := uiView.NewSession(ctx, func(p grain.UiView_newSession_Params) error {
			return nil
		}).Session()

		webSess := websession.WebSession{Client: sess.Client}
		webSess.Get(ctx, func(p websession.WebSession_get_Params) error {
			p.SetPath(path)
			return nil
		})
	}

	go get("/one")
	go get("/two")
	<-done
	<-done
}
