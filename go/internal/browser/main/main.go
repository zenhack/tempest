package browsermain

import (
	"context"
	"syscall/js"

	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/vdom"
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

func Main() {
	ctx := context.Background()

	body := vdom.DomNode{
		Value: js.Global().
			Get("document").
			Call("getElementsByTagName", "body").
			Index(0),
	}
	up := vdom.NewUpdater(body)
	defer up.Close()
	uiMsgs := make(chan Msg)
	go func() {
		m := initModel()
		for {
			up.Update(m.View(uiMsgs))
			f := <-uiMsgs
			m = f(m)
		}
	}()

	conn, api := getCapnpApi(ctx)
	defer conn.Close()
	fut, rel := api.GetLoginSession(ctx, nil)
	defer rel()
	_, rel = fut.Session().ListGrains(ctx, func(p external.LoginSession_listGrains_Params) error {
		p.SetInto(collection.Pusher_ServerToClient(grainPusher{
			uiMsgs: uiMsgs,
		}))
		return nil
	})
	defer rel()
	_, err := fut.Struct()
	if err != nil {
		println("getLoginSesion(): " + err.Error())
	}

	<-conn.Done()
}
