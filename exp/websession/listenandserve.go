package websession

import (
	"context"
	"io/ioutil"
	"net/http"

	graincp "zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/grain"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/server"
)

// Convienence wrapper for launching an http.Handler talking directly to the sandstorm
// api.
//
// If the handler implements MainView or UiView directly, those will be used rather
// than wrapping the handler.
//
// Unless the handler itself implements UiView, The getViewInfo method will pull its
// data from the sandstorm manifest.
func ListenAndServe(ctx context.Context, h http.Handler, policy *server.Policy) error {
	var uiView graincp.UiView

	if ctx == nil {
		ctx = context.Background()
	}

	if h == nil {
		h = http.DefaultServeMux
	}

	if mainSrv, ok := h.(graincp.MainView_Server); ok {
		mainClnt := graincp.MainView_ServerToClient(mainSrv, policy)
		uiView.Client = mainClnt.Client
	} else if uiSrv, ok := h.(graincp.UiView_Server); ok {
		uiView = graincp.UiView_ServerToClient(uiSrv, policy)
	} else if mostlyUiSrv, ok := h.(mostlyUiView); ok {
		uiSrv := manifestViewInfo{mostlyUiSrv}
		uiView = graincp.UiView_ServerToClient(uiSrv, policy)
	} else {
		hui := &HandlerUiView{Handler: h}
		uiSrv := manifestViewInfo{hui}
		uiView = graincp.UiView_ServerToClient(uiSrv, policy)
	}
	_, err := grain.ConnectAPI(ctx, uiView)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

// UiView_Server, sans GetViewInfo.
type mostlyUiView interface {
	NewSession(context.Context, graincp.UiView_newSession) error
	NewRequestSession(context.Context, graincp.UiView_newRequestSession) error
	NewOfferSession(context.Context, graincp.UiView_newOfferSession) error
}

// Implements UiView_Server. `getViewInfo` fetches its result from
// sandstorm-manifest. The rest of the methods are delegated to
// the mostlyUiView field.
type manifestViewInfo struct {
	mostlyUiView
}

func (v manifestViewInfo) GetViewInfo(context.Context, graincp.UiView_getViewInfo) error {
	buf, err := ioutil.ReadFile("/sandstorm-manifest")
	if err != nil {
		return err
	}
	_, err = capnp.Unmarshal(buf)
	if err != nil {
		return err
	}
	return capnp.Unimplemented("TODO")
}
