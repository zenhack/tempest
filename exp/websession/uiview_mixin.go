package websession

import (
	"net/http"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/capnp/identity"
	"zenhack.net/go/sandstorm/capnp/websession"
	"zombiezen.com/go/capnproto2"
)

// A HandlerUiView implements the new*Session methods of the grain.UiView_Server,
// by wrapping an http.Handler.
//
// When then handler is invoked, its request argument will contain a SessionData
// corresponding to the UiView method's arguments. The SessionData may be extracted
// with GetSessionData.
//
// A user of this package should embed one of these inside of another struct which
// implements GetViewInfo.
type HandlerUiView struct {
	Handler http.Handler
}

func (v *HandlerUiView) NewSession(p grain.UiView_newSession) error {
	sessionData, err := sessionDataCommon(p.Params)
	if err != nil {
		return err
	}
	sessionData.SetNormal()
	p.Results.SetSession(v.makeUiSession(sessionData))
	return nil
}

func (v *HandlerUiView) NewRequestSession(p grain.UiView_newRequestSession) error {
	sessionData, err := sessionDataCommon(p.Params)
	if err != nil {
		return err
	}
	sessionData.SetRequest()
	sessionData.SetSessionType(p.Params.SessionType())
	requestInfo, err := p.Params.RequestInfo()
	if err != nil {
		return err
	}
	sessionData.Request().SetRequestInfo(requestInfo)
	p.Results.SetSession(v.makeUiSession(sessionData))
	return nil
}

func (v *HandlerUiView) NewOfferSession(p grain.UiView_newOfferSession) error {
	sessionData, err := sessionDataCommon(p.Params)
	if err != nil {
		return err
	}
	sessionData.SetOffer()
	sessionData.SetSessionType(p.Params.SessionType())
	offer, err := p.Params.Offer()
	if err != nil {
		return err
	}
	sessionData.Offer().SetOffer(offer)
	descriptor, err := p.Params.Descriptor()
	if err != nil {
		return err
	}
	sessionData.Offer().SetDescriptor(descriptor)
	p.Results.SetSession(v.makeUiSession(sessionData))
	return nil
}

// Return a new UiSession client using our Handler and the provided SessionData.
func (v *HandlerUiView) makeUiSession(data SessionData) grain.UiSession {
	client := websession.WebSession_ServerToClient(&handlerWebSession{
		sessionData: data,
		handler:     v.Handler,
	})
	return grain.UiSession{Client: client.Client}
}

// Extract the data common to all new*Session methods from args, and return it as
// a SessionData. The caller will want to fill in the remaining fields for their
// specific method.
//
// Returns any error encountered in the process.
func sessionDataCommon(args sessionArgs) (data SessionData, err error) {
	data, err = NewSessionData(args.Segment())

	if err != nil {
		return
	}

	data.SetContext(args.Context())
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
	SessionParams() (capnp.Pointer, error)
	TabId() ([]byte, error)

	// Params always has this, and it's a convienient place to allocate the
	// SessionData:
	Segment() *capnp.Segment
}
