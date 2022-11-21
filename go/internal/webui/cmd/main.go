package main

import (
	"context"
	"syscall/js"

	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/vdom"
	vb "zenhack.net/go/vdom/builder"
	wscapnpjs "zenhack.net/go/websocket-capnp/js"
)

func getCapnpApi(ctx context.Context) (*rpc.Conn, external.ExternalApi) {
	// Derive the websocket url from window.location:
	location := js.Global().Get("window").Get("location")
	url := "ws"
	if location.Get("protocol").String() == "https:" {
		url += "s"
	}
	url += "://" + location.Get("host").String() + "/_capnp-api"

	codec := wscapnpjs.New(url, []string{"capnp-rpc"})
	trans := transport.New(codec)
	conn := rpc.NewConn(trans, nil)
	bs := external.ExternalApi(conn.Bootstrap(ctx))
	return conn, bs
}

func view() vdom.VNode {
	host := js.Global().Get("window").Get("location").Get("host").String()
	return vb.H("div", vb.A{"class": "main-ui"}, nil,
		vb.H("div", vb.A{"class": "main-ui__topbar"}, nil,
			vb.H("p", nil, nil, vb.T("Topbar")),
		),
		vb.H("div", vb.A{"class": "main-ui__main"}, nil,
			vb.H("div", vb.A{"class": "main-ui__sidebar"}, nil,
				vb.H("p", nil, nil, vb.T("Sidebar")),
			),
			vb.H("iframe", vb.A{
				"src":   "//grain." + host,
				"class": "main-ui__grain-iframe",
			}, nil),
		),
	)
}

func main() {
	ctx := context.Background()
	conn, api := getCapnpApi(ctx)
	defer conn.Close()

	go func() {
		fut, rel := api.GetLoginSession(ctx, nil)
		defer rel()
		_, err := fut.Struct()
		print("getLoginSession(): " + err.Error())
	}()

	body := js.Global().
		Get("document").
		Call("getElementsByTagName", "body").
		Index(0)
	body.Call("appendChild", view().ToDomNode().Value)
	<-conn.Done()
}
