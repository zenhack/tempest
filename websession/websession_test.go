package websession

import (
	"bytes"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"testing"
	"zenhack.net/go/sandstorm/capnp/grain"
	util_capnp "zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/util"
)

// Very basic: set up a handler that writes "Hello!\n" to the response body
// and make sure it comes through on the other side
func TestGET200(t *testing.T) {
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello!\n"))
	})
	webSession := websession.WebSession_ServerToClient(FromHandler(mux))
	buf := &bytes.Buffer{}
	pipeReader, pipeWriter := io.Pipe()
	go webSession.Get(ctx, func(p websession.WebSession_get_Params) error {
		p.SetPath("/")
		wxCtx, err := p.NewContext()
		if err != nil {
			t.Fatal(err)
			return err
		}
		wxCtx.SetResponseStream(
			util_capnp.ByteStream_ServerToClient(
				&util.WriteCloserByteStream{WC: pipeWriter}))
		return nil
	})
	_, err := io.Copy(buf, pipeReader)
	if err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	if output != "Hello!\n" {
		t.Fatalf("Expected \"Hello!\\n\" but got %q", output)
	}
}

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
