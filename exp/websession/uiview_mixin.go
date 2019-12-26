package websession

import (
	"context"
	"net/http"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/capnp/identity"
	"zenhack.net/go/sandstorm/capnp/websession"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/server"
)

// A HandlerUiView implements the new*Session methods of the grain.UiView_Server
// interface, by wrapping an http.Handler.
//
// When the handler is invoked, its request argument will contain a SessionData
// corresponding to the UiView method's arguments. The SessionData may be extracted
// with GetSessionData.
//
// The capabilities in the session data will be live until the sessions' Shutdown()
// methods are called. If they are needed afterwards, the user must add references
// to the clients.
//
// A user of this package should embed one of these inside of another struct which
// implements GetViewInfo.
type HandlerUiView struct {
	http.Handler
}

func (v *HandlerUiView) NewSession(ctx context.Context, p grain.UiView_newSession) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	sessionData, err := sessionDataCommon(results.Segment(), p.Args())
	if err != nil {
		return err
	}
	sessionData.SetNormal()
	results.SetSession(v.makeUiSession(sessionData))
	return nil
}

func (v *HandlerUiView) NewRequestSession(ctx context.Context, p grain.UiView_newRequestSession) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	sessionData, err := sessionDataCommon(results.Segment(), p.Args())
	if err != nil {
		return err
	}
	sessionData.SetRequest()
	sessionData.SetSessionType(p.Args().SessionType())
	requestInfo, err := p.Args().RequestInfo()
	if err != nil {
		return err
	}
	sessionData.Request().SetRequestInfo(requestInfo)
	results.SetSession(v.makeUiSession(sessionData))
	return nil
}

func (v *HandlerUiView) NewOfferSession(ctx context.Context, p grain.UiView_newOfferSession) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	sessionData, err := sessionDataCommon(results.Segment(), p.Args())
	if err != nil {
		return err
	}
	sessionData.SetOffer()
	sessionData.SetSessionType(p.Args().SessionType())
	offer, err := p.Args().Offer()
	if err != nil {
		return err
	}

	offer.Interface().Client().AddRef()

	sessionData.Offer().SetOffer(offer)
	descriptor, err := p.Args().Descriptor()
	if err != nil {
		return err
	}
	sessionData.Offer().SetDescriptor(descriptor)
	results.SetSession(v.makeUiSession(sessionData))
	return nil
}

// Return a new UiSession client using our Handler and the provided SessionData.
func (v *HandlerUiView) makeUiSession(data SessionData) grain.UiSession {
	client := websession.WebSession_ServerToClient(&handlerWebSession{
		sessionData: data,
		handler:     v.Handler,
	}, &server.Policy{})
	return grain.UiSession{Client: client.Client}
}

// Extract the data common to all new*Session methods from args, and return it as
// a SessionData, allocated in seg. The caller will want to fill in the remaining
// fields for their specific method.
//
// This also adds a reference to interface objects common to all variants, so
// that they are actually usable in requests.
//
// Returns any error encountered in the process.
func sessionDataCommon(seg *capnp.Segment, args sessionArgs) (data SessionData, err error) {
	data, err = NewSessionData(seg)
	if err != nil {
		return
	}
	wsCtx := args.Context()
	wsCtx.Client = wsCtx.Client.AddRef()
	data.SetContext(wsCtx)
	userInfo, err := args.UserInfo()
	if err != nil {
		return
	}
	data.SetUserInfo(userInfo)
	sessionParams, err := args.SessionParams()
	if err != nil {
		return
	}
	data.SetSessionParams(sessionParams)
	tabId, err := args.TabId()
	if err != nil {
		return
	}
	data.SetTabId(tabId)
	return
}

// interface to capture the parameters that exist for all of the new*Session methods.
type sessionArgs interface {
	Context() grain.SessionContext
	UserInfo() (identity.UserInfo, error)
	SessionParams() (capnp.Ptr, error)
	TabId() ([]byte, error)
}
