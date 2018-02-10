package main

import (
	"context"
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
		w.Write([]byte("Hello, World!"))
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
