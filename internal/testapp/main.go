package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/websession/handler"

	"zombiezen.com/go/capnproto2/rpc"
)

type UiView struct {
	*handler.HandlerUiView
}

func (*UiView) GetViewInfo(grain.UiView_getViewInfo) error { return nil }

func main() {
	file := os.NewFile(3, "<sandstorm rpc socket @ fd #3>")
	conn, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	transport := rpc.StreamTransport(conn)
	rpc.NewConn(transport, rpc.MainInterface(grain.UiView_ServerToClient(&UiView{
		HandlerUiView: &handler.HandlerUiView{
			Handler: http.DefaultServeMux,
		},
	}).Client))
	<-context.Background().Done()
}
