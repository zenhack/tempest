package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/exp/websession"

	"zombiezen.com/go/capnproto2/rpc"
)

type UiView struct {
	*websession.HandlerUiView
}

func (*UiView) GetViewInfo(grain.UiView_getViewInfo) error { return nil }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`<a href="/static/">main</a>`))
	})

	http.Handle("/static/", http.FileServer(http.Dir("")))

	http.HandleFunc("/echo-request/", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"url":     req.URL,
			"method":  req.Method,
			"headers": req.Header,
			"body":    body,
			"cookies": req.Cookies(),
		})
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/set-cookie", func(w http.ResponseWriter, req *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "bob",
			Value:  "chocolate chip",
			Secure: true,
		})
	})

	http.HandleFunc("/no-op-handler", func(w http.ResponseWriter, req *http.Request) {
		// Make sure stuff works even if we don't do anything in the handler.
	})

	file := os.NewFile(3, "<sandstorm rpc socket @ fd #3>")
	conn, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	transport := rpc.StreamTransport(conn)
	rpc.NewConn(transport, rpc.MainInterface(grain.UiView_ServerToClient(&UiView{
		HandlerUiView: &websession.HandlerUiView{
			Handler: http.DefaultServeMux,
		},
	}).Client))
	<-context.Background().Done()
}
