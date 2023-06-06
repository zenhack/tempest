package browsermain

import (
	"context"
	"syscall/js"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"zenhack.net/go/tea"
	"zenhack.net/go/tea/vdom"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/orerr"
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

// navigateMessage returns a Navigate that communicates a navigation event to the current
// value of window.location.
func navigateMessage() Navigate {
	loc := js.Global().Get("location")
	return Navigate{
		Path:     loc.Get("pathname").String(),
		Fragment: loc.Get("hash").String(),
	}
}

func Main() {
	ctx := context.Background()

	body := vdom.DomNode{
		Value: js.Global().
			Get("document").
			Call("getElementsByTagName", "body").
			Index(0),
	}
	apiPromise, apiResolver := capnp.NewLocalPromise[external.ExternalApi]()
	model := initModel(apiPromise)
	cmd := navigateMessage().Update(&model)
	app := tea.NewApp(model)
	if cmd != nil {
		go cmd(ctx, app.SendMessage)
	}
	js.Global().Call("addEventListener", "hashchange",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			app.SendMessage(navigateMessage())
			return nil
		}))
	go app.Run(ctx, body)

	conn, api := getCapnpApi(ctx)
	defer conn.Close()
	apiResolver.Fulfill(api)

	sessionsFut, rel := api.GetSessions(ctx, nil)
	defer rel()
	viewsFut, rel := sessionsFut.Visitor().Views(ctx, nil)
	defer rel()

	// FIXME: we're blocking on viewsFut for now due to what
	// seems like a go-capnp bug. Figure out what's going on
	// there and then drop this for improved latency.
	viewsRes, err := viewsFut.Struct()
	if err != nil {
		app.SendMessage(NewError{Err: err})
	} else {
		go func() {
			syncFut, rel := viewsRes.Views().Sync(ctx, func(p collection.Puller_sync_Params) error {
				p.SetInto(collection.Pusher_ServerToClient(pusher[types.GrainID, external.UiView]{
					sendMsg: app.SendMessage,
					hooks:   grainPusher{},
				}))
				return nil
			})
			defer rel()
			if _, err := syncFut.Struct(); err != nil {
				app.SendMessage(NewError{Err: err})
			}
		}()
	}

	res, err := sessionsFut.Struct()
	if err != nil {
		app.SendMessage(LoginSessionResult{Result: orerr.New(Sessions{}, err)})
	} else {
		app.SendMessage(LoginSessionResult{
			Result: orerr.New(Sessions{
				Visitor: res.Visitor().AddRef(),
				User:    res.User().AddRef(),
			}, nil),
		})
	}
	<-ctx.Done()
}
