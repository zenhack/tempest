package hacksession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	email "zenhack.net/go/sandstorm/capnp/email"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	ip "zenhack.net/go/sandstorm/capnp/ip"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type HackSessionContext struct{ Client capnp.Client }

func (c HackSessionContext) GetPublicId(ctx context.Context, params func(HackSessionContext_getPublicId_Params) error, opts ...capnp.CallOption) HackSessionContext_getPublicId_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getPublicId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      0,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getPublicId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getPublicId_Params{Struct: s}) }
	}
	return HackSessionContext_getPublicId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) HttpGet(ctx context.Context, params func(HackSessionContext_httpGet_Params) error, opts ...capnp.CallOption) HackSessionContext_httpGet_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_httpGet_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      1,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "httpGet",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_httpGet_Params{Struct: s}) }
	}
	return HackSessionContext_httpGet_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetUserAddress(ctx context.Context, params func(HackSessionContext_getUserAddress_Params) error, opts ...capnp.CallOption) email.EmailAddress_Promise {
	if c.Client == nil {
		return email.EmailAddress_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      2,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUserAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getUserAddress_Params{Struct: s}) }
	}
	return email.EmailAddress_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGenerateApiToken(ctx context.Context, params func(HackSessionContext_obsoleteGenerateApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGenerateApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      3,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGenerateApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(HackSessionContext_obsoleteGenerateApiToken_Params{Struct: s})
		}
	}
	return HackSessionContext_obsoleteGenerateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteListApiTokens(ctx context.Context, params func(HackSessionContext_obsoleteListApiTokens_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteListApiTokens_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteListApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteListApiTokens",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteListApiTokens_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteListApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteRevokeApiToken(ctx context.Context, params func(HackSessionContext_obsoleteRevokeApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteRevokeApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteRevokeApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteRevokeApiToken_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteRevokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGetIpNetwork(ctx context.Context, params func(HackSessionContext_obsoleteGetIpNetwork_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGetIpNetwork_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGetIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      6,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpNetwork",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteGetIpNetwork_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteGetIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGetIpInterface(ctx context.Context, params func(HackSessionContext_obsoleteGetIpInterface_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGetIpInterface_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGetIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpInterface",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteGetIpInterface_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteGetIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetUiViewForEndpoint(ctx context.Context, params func(HackSessionContext_getUiViewForEndpoint_Params) error, opts ...capnp.CallOption) HackSessionContext_getUiViewForEndpoint_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getUiViewForEndpoint_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      8,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUiViewForEndpoint",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getUiViewForEndpoint_Params{Struct: s}) }
	}
	return HackSessionContext_getUiViewForEndpoint_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetSharedPermissions(ctx context.Context, params func(grain.SessionContext_getSharedPermissions_Params) error, opts ...capnp.CallOption) grain.SessionContext_getSharedPermissions_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_getSharedPermissions_Params{Struct: s}) }
	}
	return grain.SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) TieToUser(ctx context.Context, params func(grain.SessionContext_tieToUser_Params) error, opts ...capnp.CallOption) grain.SessionContext_tieToUser_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_tieToUser_Params{Struct: s}) }
	}
	return grain.SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Offer(ctx context.Context, params func(grain.SessionContext_offer_Params) error, opts ...capnp.CallOption) grain.SessionContext_offer_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_offer_Params{Struct: s}) }
	}
	return grain.SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Request(ctx context.Context, params func(grain.SessionContext_request_Params) error, opts ...capnp.CallOption) grain.SessionContext_request_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_request_Params{Struct: s}) }
	}
	return grain.SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) FulfillRequest(ctx context.Context, params func(grain.SessionContext_fulfillRequest_Params) error, opts ...capnp.CallOption) grain.SessionContext_fulfillRequest_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_fulfillRequest_Params{Struct: s}) }
	}
	return grain.SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Close(ctx context.Context, params func(grain.SessionContext_close_Params) error, opts ...capnp.CallOption) grain.SessionContext_close_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_close_Params{Struct: s}) }
	}
	return grain.SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) OpenView(ctx context.Context, params func(grain.SessionContext_openView_Params) error, opts ...capnp.CallOption) grain.SessionContext_openView_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_openView_Params{Struct: s}) }
	}
	return grain.SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ClaimRequest(ctx context.Context, params func(grain.SessionContext_claimRequest_Params) error, opts ...capnp.CallOption) grain.SessionContext_claimRequest_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_claimRequest_Params{Struct: s}) }
	}
	return grain.SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Activity(ctx context.Context, params func(grain.SessionContext_activity_Params) error, opts ...capnp.CallOption) grain.SessionContext_activity_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_activity_Params{Struct: s}) }
	}
	return grain.SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Send(ctx context.Context, params func(email.EmailSendPort_send_Params) error, opts ...capnp.CallOption) email.EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_send_Params{Struct: s}) }
	}
	return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) HintAddress(ctx context.Context, params func(email.EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) email.EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type HackSessionContext_Server interface {
	GetPublicId(HackSessionContext_getPublicId) error

	HttpGet(HackSessionContext_httpGet) error

	GetUserAddress(HackSessionContext_getUserAddress) error

	ObsoleteGenerateApiToken(HackSessionContext_obsoleteGenerateApiToken) error

	ObsoleteListApiTokens(HackSessionContext_obsoleteListApiTokens) error

	ObsoleteRevokeApiToken(HackSessionContext_obsoleteRevokeApiToken) error

	ObsoleteGetIpNetwork(HackSessionContext_obsoleteGetIpNetwork) error

	ObsoleteGetIpInterface(HackSessionContext_obsoleteGetIpInterface) error

	GetUiViewForEndpoint(HackSessionContext_getUiViewForEndpoint) error

	GetSharedPermissions(grain.SessionContext_getSharedPermissions) error

	TieToUser(grain.SessionContext_tieToUser) error

	Offer(grain.SessionContext_offer) error

	Request(grain.SessionContext_request) error

	FulfillRequest(grain.SessionContext_fulfillRequest) error

	Close(grain.SessionContext_close) error

	OpenView(grain.SessionContext_openView) error

	ClaimRequest(grain.SessionContext_claimRequest) error

	Activity(grain.SessionContext_activity) error

	Send(email.EmailSendPort_send) error

	HintAddress(email.EmailSendPort_hintAddress) error
}

func HackSessionContext_ServerToClient(s HackSessionContext_Server) HackSessionContext {
	c, _ := s.(server.Closer)
	return HackSessionContext{Client: server.New(HackSessionContext_Methods(nil, s), c)}
}

func HackSessionContext_Methods(methods []server.Method, s HackSessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 20)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      0,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getPublicId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getPublicId{c, opts, HackSessionContext_getPublicId_Params{Struct: p}, HackSessionContext_getPublicId_Results{Struct: r}}
			return s.GetPublicId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      1,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "httpGet",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_httpGet{c, opts, HackSessionContext_httpGet_Params{Struct: p}, HackSessionContext_httpGet_Results{Struct: r}}
			return s.HttpGet(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      2,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUserAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getUserAddress{c, opts, HackSessionContext_getUserAddress_Params{Struct: p}, email.EmailAddress{Struct: r}}
			return s.GetUserAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      3,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGenerateApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGenerateApiToken{c, opts, HackSessionContext_obsoleteGenerateApiToken_Params{Struct: p}, HackSessionContext_obsoleteGenerateApiToken_Results{Struct: r}}
			return s.ObsoleteGenerateApiToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteListApiTokens",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteListApiTokens{c, opts, HackSessionContext_obsoleteListApiTokens_Params{Struct: p}, HackSessionContext_obsoleteListApiTokens_Results{Struct: r}}
			return s.ObsoleteListApiTokens(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteRevokeApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteRevokeApiToken{c, opts, HackSessionContext_obsoleteRevokeApiToken_Params{Struct: p}, HackSessionContext_obsoleteRevokeApiToken_Results{Struct: r}}
			return s.ObsoleteRevokeApiToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      6,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpNetwork",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGetIpNetwork{c, opts, HackSessionContext_obsoleteGetIpNetwork_Params{Struct: p}, HackSessionContext_obsoleteGetIpNetwork_Results{Struct: r}}
			return s.ObsoleteGetIpNetwork(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpInterface",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGetIpInterface{c, opts, HackSessionContext_obsoleteGetIpInterface_Params{Struct: p}, HackSessionContext_obsoleteGetIpInterface_Results{Struct: r}}
			return s.ObsoleteGetIpInterface(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      8,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUiViewForEndpoint",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getUiViewForEndpoint{c, opts, HackSessionContext_getUiViewForEndpoint_Params{Struct: p}, HackSessionContext_getUiViewForEndpoint_Results{Struct: r}}
			return s.GetUiViewForEndpoint(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_getSharedPermissions{c, opts, grain.SessionContext_getSharedPermissions_Params{Struct: p}, grain.SessionContext_getSharedPermissions_Results{Struct: r}}
			return s.GetSharedPermissions(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_tieToUser{c, opts, grain.SessionContext_tieToUser_Params{Struct: p}, grain.SessionContext_tieToUser_Results{Struct: r}}
			return s.TieToUser(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_offer{c, opts, grain.SessionContext_offer_Params{Struct: p}, grain.SessionContext_offer_Results{Struct: r}}
			return s.Offer(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_request{c, opts, grain.SessionContext_request_Params{Struct: p}, grain.SessionContext_request_Results{Struct: r}}
			return s.Request(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_fulfillRequest{c, opts, grain.SessionContext_fulfillRequest_Params{Struct: p}, grain.SessionContext_fulfillRequest_Results{Struct: r}}
			return s.FulfillRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_close{c, opts, grain.SessionContext_close_Params{Struct: p}, grain.SessionContext_close_Results{Struct: r}}
			return s.Close(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_openView{c, opts, grain.SessionContext_openView_Params{Struct: p}, grain.SessionContext_openView_Results{Struct: r}}
			return s.OpenView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_claimRequest{c, opts, grain.SessionContext_claimRequest_Params{Struct: p}, grain.SessionContext_claimRequest_Results{Struct: r}}
			return s.ClaimRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_activity{c, opts, grain.SessionContext_activity_Params{Struct: p}, grain.SessionContext_activity_Results{Struct: r}}
			return s.Activity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_send{c, opts, email.EmailSendPort_send_Params{Struct: p}, email.EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_hintAddress{c, opts, email.EmailSendPort_hintAddress_Params{Struct: p}, email.EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// HackSessionContext_getPublicId holds the arguments for a server call to HackSessionContext.getPublicId.
type HackSessionContext_getPublicId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getPublicId_Params
	Results HackSessionContext_getPublicId_Results
}

// HackSessionContext_httpGet holds the arguments for a server call to HackSessionContext.httpGet.
type HackSessionContext_httpGet struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_httpGet_Params
	Results HackSessionContext_httpGet_Results
}

// HackSessionContext_getUserAddress holds the arguments for a server call to HackSessionContext.getUserAddress.
type HackSessionContext_getUserAddress struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getUserAddress_Params
	Results email.EmailAddress
}

// HackSessionContext_obsoleteGenerateApiToken holds the arguments for a server call to HackSessionContext.obsoleteGenerateApiToken.
type HackSessionContext_obsoleteGenerateApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGenerateApiToken_Params
	Results HackSessionContext_obsoleteGenerateApiToken_Results
}

// HackSessionContext_obsoleteListApiTokens holds the arguments for a server call to HackSessionContext.obsoleteListApiTokens.
type HackSessionContext_obsoleteListApiTokens struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteListApiTokens_Params
	Results HackSessionContext_obsoleteListApiTokens_Results
}

// HackSessionContext_obsoleteRevokeApiToken holds the arguments for a server call to HackSessionContext.obsoleteRevokeApiToken.
type HackSessionContext_obsoleteRevokeApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteRevokeApiToken_Params
	Results HackSessionContext_obsoleteRevokeApiToken_Results
}

// HackSessionContext_obsoleteGetIpNetwork holds the arguments for a server call to HackSessionContext.obsoleteGetIpNetwork.
type HackSessionContext_obsoleteGetIpNetwork struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGetIpNetwork_Params
	Results HackSessionContext_obsoleteGetIpNetwork_Results
}

// HackSessionContext_obsoleteGetIpInterface holds the arguments for a server call to HackSessionContext.obsoleteGetIpInterface.
type HackSessionContext_obsoleteGetIpInterface struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGetIpInterface_Params
	Results HackSessionContext_obsoleteGetIpInterface_Results
}

// HackSessionContext_getUiViewForEndpoint holds the arguments for a server call to HackSessionContext.getUiViewForEndpoint.
type HackSessionContext_getUiViewForEndpoint struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getUiViewForEndpoint_Params
	Results HackSessionContext_getUiViewForEndpoint_Results
}

type HackSessionContext_TokenInfo struct{ capnp.Struct }

func NewHackSessionContext_TokenInfo(s *capnp.Segment) (HackSessionContext_TokenInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_TokenInfo{st}, err
}

func NewRootHackSessionContext_TokenInfo(s *capnp.Segment) (HackSessionContext_TokenInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_TokenInfo{st}, err
}

func ReadRootHackSessionContext_TokenInfo(msg *capnp.Message) (HackSessionContext_TokenInfo, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_TokenInfo{root.Struct()}, err
}

func (s HackSessionContext_TokenInfo) String() string {
	str, _ := text.Marshal(0xf910658ae8c6674d, s.Struct)
	return str
}

func (s HackSessionContext_TokenInfo) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_TokenInfo) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_TokenInfo) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_TokenInfo) Petname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_TokenInfo) HasPetname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_TokenInfo) SetPetname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_TokenInfo) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(2)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s HackSessionContext_TokenInfo) HasUserInfo() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_TokenInfo) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// HackSessionContext_TokenInfo_List is a list of HackSessionContext_TokenInfo.
type HackSessionContext_TokenInfo_List struct{ capnp.List }

// NewHackSessionContext_TokenInfo creates a new list of HackSessionContext_TokenInfo.
func NewHackSessionContext_TokenInfo_List(s *capnp.Segment, sz int32) (HackSessionContext_TokenInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HackSessionContext_TokenInfo_List{l}, err
}

func (s HackSessionContext_TokenInfo_List) At(i int) HackSessionContext_TokenInfo {
	return HackSessionContext_TokenInfo{s.List.Struct(i)}
}

func (s HackSessionContext_TokenInfo_List) Set(i int, v HackSessionContext_TokenInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_TokenInfo_Promise is a wrapper for a HackSessionContext_TokenInfo promised by a client call.
type HackSessionContext_TokenInfo_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_TokenInfo_Promise) Struct() (HackSessionContext_TokenInfo, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_TokenInfo{s}, err
}

func (p HackSessionContext_TokenInfo_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type HackSessionContext_getPublicId_Params struct{ capnp.Struct }

func NewHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getPublicId_Params{st}, err
}

func NewRootHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getPublicId_Params{st}, err
}

func ReadRootHackSessionContext_getPublicId_Params(msg *capnp.Message) (HackSessionContext_getPublicId_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getPublicId_Params{root.Struct()}, err
}

func (s HackSessionContext_getPublicId_Params) String() string {
	str, _ := text.Marshal(0xe706d835e9cd64af, s.Struct)
	return str
}

// HackSessionContext_getPublicId_Params_List is a list of HackSessionContext_getPublicId_Params.
type HackSessionContext_getPublicId_Params_List struct{ capnp.List }

// NewHackSessionContext_getPublicId_Params creates a new list of HackSessionContext_getPublicId_Params.
func NewHackSessionContext_getPublicId_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getPublicId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_getPublicId_Params_List{l}, err
}

func (s HackSessionContext_getPublicId_Params_List) At(i int) HackSessionContext_getPublicId_Params {
	return HackSessionContext_getPublicId_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getPublicId_Params_List) Set(i int, v HackSessionContext_getPublicId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getPublicId_Params_Promise is a wrapper for a HackSessionContext_getPublicId_Params promised by a client call.
type HackSessionContext_getPublicId_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getPublicId_Params_Promise) Struct() (HackSessionContext_getPublicId_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getPublicId_Params{s}, err
}

type HackSessionContext_getPublicId_Results struct{ capnp.Struct }

func NewHackSessionContext_getPublicId_Results(s *capnp.Segment) (HackSessionContext_getPublicId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return HackSessionContext_getPublicId_Results{st}, err
}

func NewRootHackSessionContext_getPublicId_Results(s *capnp.Segment) (HackSessionContext_getPublicId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return HackSessionContext_getPublicId_Results{st}, err
}

func ReadRootHackSessionContext_getPublicId_Results(msg *capnp.Message) (HackSessionContext_getPublicId_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getPublicId_Results{root.Struct()}, err
}

func (s HackSessionContext_getPublicId_Results) String() string {
	str, _ := text.Marshal(0xe96193c522f6c05d, s.Struct)
	return str
}

func (s HackSessionContext_getPublicId_Results) PublicId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasPublicId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) PublicIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetPublicId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) Hostname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasHostname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetHostname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) AutoUrl() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasAutoUrl() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) AutoUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetAutoUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) IsDemoUser() bool {
	return s.Struct.Bit(0)
}

func (s HackSessionContext_getPublicId_Results) SetIsDemoUser(v bool) {
	s.Struct.SetBit(0, v)
}

// HackSessionContext_getPublicId_Results_List is a list of HackSessionContext_getPublicId_Results.
type HackSessionContext_getPublicId_Results_List struct{ capnp.List }

// NewHackSessionContext_getPublicId_Results creates a new list of HackSessionContext_getPublicId_Results.
func NewHackSessionContext_getPublicId_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getPublicId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return HackSessionContext_getPublicId_Results_List{l}, err
}

func (s HackSessionContext_getPublicId_Results_List) At(i int) HackSessionContext_getPublicId_Results {
	return HackSessionContext_getPublicId_Results{s.List.Struct(i)}
}

func (s HackSessionContext_getPublicId_Results_List) Set(i int, v HackSessionContext_getPublicId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getPublicId_Results_Promise is a wrapper for a HackSessionContext_getPublicId_Results promised by a client call.
type HackSessionContext_getPublicId_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getPublicId_Results_Promise) Struct() (HackSessionContext_getPublicId_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getPublicId_Results{s}, err
}

type HackSessionContext_httpGet_Params struct{ capnp.Struct }

func NewHackSessionContext_httpGet_Params(s *capnp.Segment) (HackSessionContext_httpGet_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_httpGet_Params{st}, err
}

func NewRootHackSessionContext_httpGet_Params(s *capnp.Segment) (HackSessionContext_httpGet_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_httpGet_Params{st}, err
}

func ReadRootHackSessionContext_httpGet_Params(msg *capnp.Message) (HackSessionContext_httpGet_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_httpGet_Params{root.Struct()}, err
}

func (s HackSessionContext_httpGet_Params) String() string {
	str, _ := text.Marshal(0xe54437a7b8f52843, s.Struct)
	return str
}

func (s HackSessionContext_httpGet_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_httpGet_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_httpGet_Params) SetUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_httpGet_Params_List is a list of HackSessionContext_httpGet_Params.
type HackSessionContext_httpGet_Params_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Params creates a new list of HackSessionContext_httpGet_Params.
func NewHackSessionContext_httpGet_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_httpGet_Params_List{l}, err
}

func (s HackSessionContext_httpGet_Params_List) At(i int) HackSessionContext_httpGet_Params {
	return HackSessionContext_httpGet_Params{s.List.Struct(i)}
}

func (s HackSessionContext_httpGet_Params_List) Set(i int, v HackSessionContext_httpGet_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_httpGet_Params_Promise is a wrapper for a HackSessionContext_httpGet_Params promised by a client call.
type HackSessionContext_httpGet_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_httpGet_Params_Promise) Struct() (HackSessionContext_httpGet_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_httpGet_Params{s}, err
}

type HackSessionContext_httpGet_Results struct{ capnp.Struct }

func NewHackSessionContext_httpGet_Results(s *capnp.Segment) (HackSessionContext_httpGet_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return HackSessionContext_httpGet_Results{st}, err
}

func NewRootHackSessionContext_httpGet_Results(s *capnp.Segment) (HackSessionContext_httpGet_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return HackSessionContext_httpGet_Results{st}, err
}

func ReadRootHackSessionContext_httpGet_Results(msg *capnp.Message) (HackSessionContext_httpGet_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_httpGet_Results{root.Struct()}, err
}

func (s HackSessionContext_httpGet_Results) String() string {
	str, _ := text.Marshal(0xb44df810894a44f4, s.Struct)
	return str
}

func (s HackSessionContext_httpGet_Results) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_httpGet_Results) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Results) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_httpGet_Results) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_httpGet_Results) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s HackSessionContext_httpGet_Results) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Results) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

// HackSessionContext_httpGet_Results_List is a list of HackSessionContext_httpGet_Results.
type HackSessionContext_httpGet_Results_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Results creates a new list of HackSessionContext_httpGet_Results.
func NewHackSessionContext_httpGet_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return HackSessionContext_httpGet_Results_List{l}, err
}

func (s HackSessionContext_httpGet_Results_List) At(i int) HackSessionContext_httpGet_Results {
	return HackSessionContext_httpGet_Results{s.List.Struct(i)}
}

func (s HackSessionContext_httpGet_Results_List) Set(i int, v HackSessionContext_httpGet_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_httpGet_Results_Promise is a wrapper for a HackSessionContext_httpGet_Results promised by a client call.
type HackSessionContext_httpGet_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_httpGet_Results_Promise) Struct() (HackSessionContext_httpGet_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_httpGet_Results{s}, err
}

type HackSessionContext_getUserAddress_Params struct{ capnp.Struct }

func NewHackSessionContext_getUserAddress_Params(s *capnp.Segment) (HackSessionContext_getUserAddress_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getUserAddress_Params{st}, err
}

func NewRootHackSessionContext_getUserAddress_Params(s *capnp.Segment) (HackSessionContext_getUserAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getUserAddress_Params{st}, err
}

func ReadRootHackSessionContext_getUserAddress_Params(msg *capnp.Message) (HackSessionContext_getUserAddress_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUserAddress_Params{root.Struct()}, err
}

func (s HackSessionContext_getUserAddress_Params) String() string {
	str, _ := text.Marshal(0xa15e445037d1834c, s.Struct)
	return str
}

// HackSessionContext_getUserAddress_Params_List is a list of HackSessionContext_getUserAddress_Params.
type HackSessionContext_getUserAddress_Params_List struct{ capnp.List }

// NewHackSessionContext_getUserAddress_Params creates a new list of HackSessionContext_getUserAddress_Params.
func NewHackSessionContext_getUserAddress_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUserAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_getUserAddress_Params_List{l}, err
}

func (s HackSessionContext_getUserAddress_Params_List) At(i int) HackSessionContext_getUserAddress_Params {
	return HackSessionContext_getUserAddress_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getUserAddress_Params_List) Set(i int, v HackSessionContext_getUserAddress_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUserAddress_Params_Promise is a wrapper for a HackSessionContext_getUserAddress_Params promised by a client call.
type HackSessionContext_getUserAddress_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUserAddress_Params_Promise) Struct() (HackSessionContext_getUserAddress_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUserAddress_Params{s}, err
}

type HackSessionContext_obsoleteGenerateApiToken_Params struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGenerateApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGenerateApiToken_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) String() string {
	str, _ := text.Marshal(0x837afa61d880beb6, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) Petname() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) HasPetname() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetPetname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(1)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_obsoleteGenerateApiToken_Params) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) Expires() uint64 {
	return s.Struct.Uint64(0)
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetExpires(v uint64) {
	s.Struct.SetUint64(0, v)
}

// HackSessionContext_obsoleteGenerateApiToken_Params_List is a list of HackSessionContext_obsoleteGenerateApiToken_Params.
type HackSessionContext_obsoleteGenerateApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGenerateApiToken_Params creates a new list of HackSessionContext_obsoleteGenerateApiToken_Params.
func NewHackSessionContext_obsoleteGenerateApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGenerateApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return HackSessionContext_obsoleteGenerateApiToken_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params_List) At(i int) HackSessionContext_obsoleteGenerateApiToken_Params {
	return HackSessionContext_obsoleteGenerateApiToken_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params_List) Set(i int, v HackSessionContext_obsoleteGenerateApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGenerateApiToken_Params_Promise is a wrapper for a HackSessionContext_obsoleteGenerateApiToken_Params promised by a client call.
type HackSessionContext_obsoleteGenerateApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGenerateApiToken_Params_Promise) Struct() (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGenerateApiToken_Params{s}, err
}

func (p HackSessionContext_obsoleteGenerateApiToken_Params_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HackSessionContext_obsoleteGenerateApiToken_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGenerateApiToken_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) String() string {
	str, _ := text.Marshal(0xc9973f31a90887f9, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) Token() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) EndpointUrl() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasEndpointUrl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) EndpointUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetEndpointUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasTokenId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// HackSessionContext_obsoleteGenerateApiToken_Results_List is a list of HackSessionContext_obsoleteGenerateApiToken_Results.
type HackSessionContext_obsoleteGenerateApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGenerateApiToken_Results creates a new list of HackSessionContext_obsoleteGenerateApiToken_Results.
func NewHackSessionContext_obsoleteGenerateApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGenerateApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HackSessionContext_obsoleteGenerateApiToken_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results_List) At(i int) HackSessionContext_obsoleteGenerateApiToken_Results {
	return HackSessionContext_obsoleteGenerateApiToken_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results_List) Set(i int, v HackSessionContext_obsoleteGenerateApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGenerateApiToken_Results_Promise is a wrapper for a HackSessionContext_obsoleteGenerateApiToken_Results promised by a client call.
type HackSessionContext_obsoleteGenerateApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGenerateApiToken_Results_Promise) Struct() (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGenerateApiToken_Results{s}, err
}

type HackSessionContext_obsoleteListApiTokens_Params struct{ capnp.Struct }

func NewHackSessionContext_obsoleteListApiTokens_Params(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteListApiTokens_Params{st}, err
}

func NewRootHackSessionContext_obsoleteListApiTokens_Params(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteListApiTokens_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Params(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteListApiTokens_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Params) String() string {
	str, _ := text.Marshal(0xfe777c71f871f413, s.Struct)
	return str
}

// HackSessionContext_obsoleteListApiTokens_Params_List is a list of HackSessionContext_obsoleteListApiTokens_Params.
type HackSessionContext_obsoleteListApiTokens_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Params creates a new list of HackSessionContext_obsoleteListApiTokens_Params.
func NewHackSessionContext_obsoleteListApiTokens_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteListApiTokens_Params_List{l}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Params_List) At(i int) HackSessionContext_obsoleteListApiTokens_Params {
	return HackSessionContext_obsoleteListApiTokens_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteListApiTokens_Params_List) Set(i int, v HackSessionContext_obsoleteListApiTokens_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteListApiTokens_Params_Promise is a wrapper for a HackSessionContext_obsoleteListApiTokens_Params promised by a client call.
type HackSessionContext_obsoleteListApiTokens_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteListApiTokens_Params_Promise) Struct() (HackSessionContext_obsoleteListApiTokens_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteListApiTokens_Params{s}, err
}

type HackSessionContext_obsoleteListApiTokens_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteListApiTokens_Results(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteListApiTokens_Results{st}, err
}

func NewRootHackSessionContext_obsoleteListApiTokens_Results(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteListApiTokens_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Results(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteListApiTokens_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results) String() string {
	str, _ := text.Marshal(0xe9e4890dae20b03c, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteListApiTokens_Results) Tokens() (HackSessionContext_TokenInfo_List, error) {
	p, err := s.Struct.Ptr(0)
	return HackSessionContext_TokenInfo_List{List: p.List()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results) HasTokens() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteListApiTokens_Results) SetTokens(v HackSessionContext_TokenInfo_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTokens sets the tokens field to a newly
// allocated HackSessionContext_TokenInfo_List, preferring placement in s's segment.
func (s HackSessionContext_obsoleteListApiTokens_Results) NewTokens(n int32) (HackSessionContext_TokenInfo_List, error) {
	l, err := NewHackSessionContext_TokenInfo_List(s.Struct.Segment(), n)
	if err != nil {
		return HackSessionContext_TokenInfo_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HackSessionContext_obsoleteListApiTokens_Results_List is a list of HackSessionContext_obsoleteListApiTokens_Results.
type HackSessionContext_obsoleteListApiTokens_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Results creates a new list of HackSessionContext_obsoleteListApiTokens_Results.
func NewHackSessionContext_obsoleteListApiTokens_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteListApiTokens_Results_List{l}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results_List) At(i int) HackSessionContext_obsoleteListApiTokens_Results {
	return HackSessionContext_obsoleteListApiTokens_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteListApiTokens_Results_List) Set(i int, v HackSessionContext_obsoleteListApiTokens_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteListApiTokens_Results_Promise is a wrapper for a HackSessionContext_obsoleteListApiTokens_Results promised by a client call.
type HackSessionContext_obsoleteListApiTokens_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteListApiTokens_Results_Promise) Struct() (HackSessionContext_obsoleteListApiTokens_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteListApiTokens_Results{s}, err
}

type HackSessionContext_obsoleteRevokeApiToken_Params struct{ capnp.Struct }

func NewHackSessionContext_obsoleteRevokeApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, err
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteRevokeApiToken_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) String() string {
	str, _ := text.Marshal(0x845e8081686d8a0f, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_obsoleteRevokeApiToken_Params_List is a list of HackSessionContext_obsoleteRevokeApiToken_Params.
type HackSessionContext_obsoleteRevokeApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Params creates a new list of HackSessionContext_obsoleteRevokeApiToken_Params.
func NewHackSessionContext_obsoleteRevokeApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteRevokeApiToken_Params_List{l}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params_List) At(i int) HackSessionContext_obsoleteRevokeApiToken_Params {
	return HackSessionContext_obsoleteRevokeApiToken_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params_List) Set(i int, v HackSessionContext_obsoleteRevokeApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteRevokeApiToken_Params_Promise is a wrapper for a HackSessionContext_obsoleteRevokeApiToken_Params promised by a client call.
type HackSessionContext_obsoleteRevokeApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteRevokeApiToken_Params_Promise) Struct() (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteRevokeApiToken_Params{s}, err
}

type HackSessionContext_obsoleteRevokeApiToken_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteRevokeApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, err
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteRevokeApiToken_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results) String() string {
	str, _ := text.Marshal(0xaea67098dc479ce6, s.Struct)
	return str
}

// HackSessionContext_obsoleteRevokeApiToken_Results_List is a list of HackSessionContext_obsoleteRevokeApiToken_Results.
type HackSessionContext_obsoleteRevokeApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Results creates a new list of HackSessionContext_obsoleteRevokeApiToken_Results.
func NewHackSessionContext_obsoleteRevokeApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteRevokeApiToken_Results_List{l}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results_List) At(i int) HackSessionContext_obsoleteRevokeApiToken_Results {
	return HackSessionContext_obsoleteRevokeApiToken_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results_List) Set(i int, v HackSessionContext_obsoleteRevokeApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteRevokeApiToken_Results_Promise is a wrapper for a HackSessionContext_obsoleteRevokeApiToken_Results promised by a client call.
type HackSessionContext_obsoleteRevokeApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteRevokeApiToken_Results_Promise) Struct() (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteRevokeApiToken_Results{s}, err
}

type HackSessionContext_obsoleteGetIpNetwork_Params struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGetIpNetwork_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpNetwork_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpNetwork_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpNetwork_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpNetwork_Params(msg *capnp.Message) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpNetwork_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params) String() string {
	str, _ := text.Marshal(0x97f9d7a51ffe0741, s.Struct)
	return str
}

// HackSessionContext_obsoleteGetIpNetwork_Params_List is a list of HackSessionContext_obsoleteGetIpNetwork_Params.
type HackSessionContext_obsoleteGetIpNetwork_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpNetwork_Params creates a new list of HackSessionContext_obsoleteGetIpNetwork_Params.
func NewHackSessionContext_obsoleteGetIpNetwork_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpNetwork_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteGetIpNetwork_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params_List) At(i int) HackSessionContext_obsoleteGetIpNetwork_Params {
	return HackSessionContext_obsoleteGetIpNetwork_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params_List) Set(i int, v HackSessionContext_obsoleteGetIpNetwork_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpNetwork_Params_Promise is a wrapper for a HackSessionContext_obsoleteGetIpNetwork_Params promised by a client call.
type HackSessionContext_obsoleteGetIpNetwork_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpNetwork_Params_Promise) Struct() (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpNetwork_Params{s}, err
}

type HackSessionContext_obsoleteGetIpNetwork_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGetIpNetwork_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpNetwork_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpNetwork_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpNetwork_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpNetwork_Results(msg *capnp.Message) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpNetwork_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) String() string {
	str, _ := text.Marshal(0xa9502e5fdabb8b07, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) Network() ip.IpNetwork {
	p, _ := s.Struct.Ptr(0)
	return ip.IpNetwork{Client: p.Interface().Client()}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) HasNetwork() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) SetNetwork(v ip.IpNetwork) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_obsoleteGetIpNetwork_Results_List is a list of HackSessionContext_obsoleteGetIpNetwork_Results.
type HackSessionContext_obsoleteGetIpNetwork_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpNetwork_Results creates a new list of HackSessionContext_obsoleteGetIpNetwork_Results.
func NewHackSessionContext_obsoleteGetIpNetwork_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpNetwork_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteGetIpNetwork_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results_List) At(i int) HackSessionContext_obsoleteGetIpNetwork_Results {
	return HackSessionContext_obsoleteGetIpNetwork_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results_List) Set(i int, v HackSessionContext_obsoleteGetIpNetwork_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpNetwork_Results_Promise is a wrapper for a HackSessionContext_obsoleteGetIpNetwork_Results promised by a client call.
type HackSessionContext_obsoleteGetIpNetwork_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpNetwork_Results_Promise) Struct() (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpNetwork_Results{s}, err
}

func (p HackSessionContext_obsoleteGetIpNetwork_Results_Promise) Network() ip.IpNetwork {
	return ip.IpNetwork{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_obsoleteGetIpInterface_Params struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGetIpInterface_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpInterface_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpInterface_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpInterface_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpInterface_Params(msg *capnp.Message) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpInterface_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Params) String() string {
	str, _ := text.Marshal(0xecebff1662ba10a1, s.Struct)
	return str
}

// HackSessionContext_obsoleteGetIpInterface_Params_List is a list of HackSessionContext_obsoleteGetIpInterface_Params.
type HackSessionContext_obsoleteGetIpInterface_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpInterface_Params creates a new list of HackSessionContext_obsoleteGetIpInterface_Params.
func NewHackSessionContext_obsoleteGetIpInterface_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpInterface_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteGetIpInterface_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Params_List) At(i int) HackSessionContext_obsoleteGetIpInterface_Params {
	return HackSessionContext_obsoleteGetIpInterface_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpInterface_Params_List) Set(i int, v HackSessionContext_obsoleteGetIpInterface_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpInterface_Params_Promise is a wrapper for a HackSessionContext_obsoleteGetIpInterface_Params promised by a client call.
type HackSessionContext_obsoleteGetIpInterface_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpInterface_Params_Promise) Struct() (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpInterface_Params{s}, err
}

type HackSessionContext_obsoleteGetIpInterface_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGetIpInterface_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpInterface_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpInterface_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpInterface_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpInterface_Results(msg *capnp.Message) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpInterface_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) String() string {
	str, _ := text.Marshal(0xb9147a48c12c807d, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) Interface() ip.IpInterface {
	p, _ := s.Struct.Ptr(0)
	return ip.IpInterface{Client: p.Interface().Client()}
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) HasInterface() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) SetInterface(v ip.IpInterface) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_obsoleteGetIpInterface_Results_List is a list of HackSessionContext_obsoleteGetIpInterface_Results.
type HackSessionContext_obsoleteGetIpInterface_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpInterface_Results creates a new list of HackSessionContext_obsoleteGetIpInterface_Results.
func NewHackSessionContext_obsoleteGetIpInterface_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpInterface_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteGetIpInterface_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Results_List) At(i int) HackSessionContext_obsoleteGetIpInterface_Results {
	return HackSessionContext_obsoleteGetIpInterface_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpInterface_Results_List) Set(i int, v HackSessionContext_obsoleteGetIpInterface_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpInterface_Results_Promise is a wrapper for a HackSessionContext_obsoleteGetIpInterface_Results promised by a client call.
type HackSessionContext_obsoleteGetIpInterface_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpInterface_Results_Promise) Struct() (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpInterface_Results{s}, err
}

func (p HackSessionContext_obsoleteGetIpInterface_Results_Promise) Interface() ip.IpInterface {
	return ip.IpInterface{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_getUiViewForEndpoint_Params struct{ capnp.Struct }

func NewHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Params{st}, err
}

func NewRootHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Params{st}, err
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Params(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUiViewForEndpoint_Params{root.Struct()}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) String() string {
	str, _ := text.Marshal(0xb45bb2d206694fe6, s.Struct)
	return str
}

func (s HackSessionContext_getUiViewForEndpoint_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getUiViewForEndpoint_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) SetUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_getUiViewForEndpoint_Params_List is a list of HackSessionContext_getUiViewForEndpoint_Params.
type HackSessionContext_getUiViewForEndpoint_Params_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Params creates a new list of HackSessionContext_getUiViewForEndpoint_Params.
func NewHackSessionContext_getUiViewForEndpoint_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_getUiViewForEndpoint_Params_List{l}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Params_List) At(i int) HackSessionContext_getUiViewForEndpoint_Params {
	return HackSessionContext_getUiViewForEndpoint_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getUiViewForEndpoint_Params_List) Set(i int, v HackSessionContext_getUiViewForEndpoint_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUiViewForEndpoint_Params_Promise is a wrapper for a HackSessionContext_getUiViewForEndpoint_Params promised by a client call.
type HackSessionContext_getUiViewForEndpoint_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUiViewForEndpoint_Params_Promise) Struct() (HackSessionContext_getUiViewForEndpoint_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUiViewForEndpoint_Params{s}, err
}

type HackSessionContext_getUiViewForEndpoint_Results struct{ capnp.Struct }

func NewHackSessionContext_getUiViewForEndpoint_Results(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Results{st}, err
}

func NewRootHackSessionContext_getUiViewForEndpoint_Results(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Results{st}, err
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Results(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUiViewForEndpoint_Results{root.Struct()}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Results) String() string {
	str, _ := text.Marshal(0xfdc944999b9296df, s.Struct)
	return str
}

func (s HackSessionContext_getUiViewForEndpoint_Results) View() grain.UiView {
	p, _ := s.Struct.Ptr(0)
	return grain.UiView{Client: p.Interface().Client()}
}

func (s HackSessionContext_getUiViewForEndpoint_Results) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getUiViewForEndpoint_Results) SetView(v grain.UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_getUiViewForEndpoint_Results_List is a list of HackSessionContext_getUiViewForEndpoint_Results.
type HackSessionContext_getUiViewForEndpoint_Results_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Results creates a new list of HackSessionContext_getUiViewForEndpoint_Results.
func NewHackSessionContext_getUiViewForEndpoint_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_getUiViewForEndpoint_Results_List{l}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Results_List) At(i int) HackSessionContext_getUiViewForEndpoint_Results {
	return HackSessionContext_getUiViewForEndpoint_Results{s.List.Struct(i)}
}

func (s HackSessionContext_getUiViewForEndpoint_Results_List) Set(i int, v HackSessionContext_getUiViewForEndpoint_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUiViewForEndpoint_Results_Promise is a wrapper for a HackSessionContext_getUiViewForEndpoint_Results promised by a client call.
type HackSessionContext_getUiViewForEndpoint_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUiViewForEndpoint_Results_Promise) Struct() (HackSessionContext_getUiViewForEndpoint_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUiViewForEndpoint_Results{s}, err
}

func (p HackSessionContext_getUiViewForEndpoint_Results_Promise) View() grain.UiView {
	return grain.UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackEmailSession struct{ Client capnp.Client }

func (c HackEmailSession) Send(ctx context.Context, params func(email.EmailSendPort_send_Params) error, opts ...capnp.CallOption) email.EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_send_Params{Struct: s}) }
	}
	return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackEmailSession) HintAddress(ctx context.Context, params func(email.EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) email.EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type HackEmailSession_Server interface {
	Send(email.EmailSendPort_send) error

	HintAddress(email.EmailSendPort_hintAddress) error
}

func HackEmailSession_ServerToClient(s HackEmailSession_Server) HackEmailSession {
	c, _ := s.(server.Closer)
	return HackEmailSession{Client: server.New(HackEmailSession_Methods(nil, s), c)}
}

func HackEmailSession_Methods(methods []server.Method, s HackEmailSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_send{c, opts, email.EmailSendPort_send_Params{Struct: p}, email.EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_hintAddress{c, opts, email.EmailSendPort_hintAddress_Params{Struct: p}, email.EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

const schema_bf6889795837d1e0 = "x\xda\xacV{l\x14U\x17\xbfg\xa6\xdbi\x13J" +
	";\xdf\x16\xbaKi\x16\x1a\xc2\xc7G\xa0\xf9( B" +
	"\x0c\xbb\x85VX,\xba[\x1e\"\x0a:m/t\xd2" +
	"\xdd\x9deg\x96B\xa3i\x01\x1f\x01\xff0>\xc1\xe0" +
	"?\xfc\x01\xa6D@\x91\x1a\xc1\x076\x06MI4)" +
	"\x06A\x85((b\x151F$\x94\xa4\x8c\xe7\xce\xee" +
	"\xcc^\x17J\xa95\xcd6\x99\xf3\xfc\xdds\xee\xf9\x9d" +
	"\xfb\xffG\xf2\x02\xc24\xd7\xa7\x15\x84,\x19'\xb8r" +
	"\xcdw>l?\xad\\o\xddB\xc2\x13\x01\x08q\x09" +
	"\x12!\xd3\x83%\x95\x02\x01\xb7Rr\x80\x80Y\xb8-" +
	"\xda\xb4\xa9}\xf5\x93D\x9e\xc8\x0c\x80\x19\xf4\x97\x943" +
	"\x83Q\x1e?\x1aTI7|\xbbO\xf5mO\x19\xe4" +
	"0\xfdL\xcf\x7f\x04\x92c\xd6n\xe9\x99\x15\xaa^\xbd" +
	"\x8b\xd3L\xf5\x9c\x07\xd4H\xcf\xbe\xf7\xf5\xa3\x15\xa1\x0e" +
	">\xe8x\xcf\x18\x16t\xb6\x15\xf4\xc7\xd7\x16\x9c\xd9\x11" +
	"\xdf\xb3\x9fs}\xc83\x99\x05\xbdR\xbdhk\xd1\xb5" +
	"\xc5\x87\xd2\xae\x16\xe0\xb0\xe7 0\xc0\x9e\x16\xe6\xfa\x80" +
	"\x9a{\xe2\xe0\xc3\x87\xf8\xd8]\x0c\x10\xb8\xbf\xb2b?" +
	"\xd1>\xa5kak\xf1a\xde\xa0\x8f\x05\x07\xb7\xece" +
	"\x06{r\xee\x9bq\xea\xf3\xce\x8f\x89\xec\x16\xcd\xefz" +
	"f\xad\xd8\xb8\xb5\xe9(!0}\xa6\xb7\x1c\xdc5^" +
	"tpWy%\xfc\xfd\x97\x10\xb3\xef\x99\xbc\x8ei\xfe" +
	"\xed\xdd\xe9p\"\x0b7\xdb;\x87\x85\x0b{Y\x05\x9f" +
	"\xdb9j\xfc\x12_\xed\xb9\x9b\xc2\xfd\xea\xad\x04w\xbf" +
	"\xb7\x04\x1d\xf2\xc7,\x10\xdc\xfbJY\xbc\xf9\x93\xfe|" +
	"\xf7\xf5Y\xd5\x17xx\xbbK\xf7\xb2\x03v\x962x" +
	"\x07\x1a?\xeb\x9dy:\xf7\"W\x9b\x93\xa5\xc7YY" +
	"W}t\xb5\xfc\xd8\x8bJ\xaf\xddM\x0bKw\xe9\x09" +
	"\xe6{\xb6\xf4\"\xfa\xde\xf3\xe6\xb8\xfd\x05[\x7f\xe8\xe5" +
	"\x83\x7f0\xd6\xea\xe6\xc9\xb1,\xf8\xae\xa2#\xf5\xa3\xcd" +
	"_.q\xc1]e\xe5\xac\xf0\x8b\xd7~\xf2\xd36Z" +
	"\xd4\xc74\x993\xa5r\xf4\x8f\xdd\x8c\xc5+c\x85)" +
	"(cg\xfe\xf6\x95\x17v\xbeZ\xdd\xdd\xcf\xe7\xe9(" +
	"\xb3\x1a\xdcU\xc6\xf2\xb8\xaf\xac\xbb\xb6\xee\xf1\x96\x1b\\" +
	"\x9esLo:\x7f\xc4lR\x1a\x9a\xa7\xeaT\x97t" +
	"U\x8bU4(\xf1X|\xceB\x94-\xa1:\x93\xcc" +
	"\xd7b\x06\xdd`Th\xf5\xba\x16\xa1\x06]@c4" +
	"\xa1\x18\xb4*\xae.\xd5\x9ailB\xc8\xa7$\x94\xa8" +
	"\x1e\x1e!\xe6`\x0aL#\xd7\xcc#$\x1c\x10!\\" +
	"+\x80\x0cP\xcc\xaa$\x07\x17\xa1p!\x0a\x97\x0a\x00" +
	"B1\xe0\x7f9\xcc\x0ckQ\xb6B\x80\xb685b" +
	"J\x94\xc2\x08\"\xe0\x0f\xcc\xa4N\x13\xc1\xd8\x1a\x8d`" +
	"\x13\x8b\xcc/\xe4Uo\xff\xd6s\xf8%\xfc\x80\"\x02" +
	"mtC\\MP\x1d\xf2\xd1:\x1f\xad\x87x\x8a:" +
	"\xba\x1e\xc1g\xce\xa0$$v\x86\x1c\xe7\x0c\x05\x0cZ" +
	"\x1eB+Fh\x063\x0a6:\xd0\x86\\2#\x18" +
	"\xbf\x9f\x1a-Z\xa2\x99\xa5R\xa2\xa0;1ro\x1f" +
	"c-5\x96a!\xaa\x1a\x1b\xf1\xb4\xba\x0dtX\x08" +
	"\xea\xa8\x9e\x8c\x88\xc6\x80\xa7\x8d\xa5\xec@6\x03\xe3&" +
	"\xb76?\xb8\xb9\x83\x15]\x1ev\x911oa2b" +
	"d\xc0\xbbn\x1f\xa6\xc90\xe2\x08\x1c\xfd|:\xf3C" +
	"\x806\xde\xff\xb1\xcb4\x09\xf1\xce\xe0n\xd84v\x88" +
	")(\xbc[\x003\xaaF\xe9\xd2\x8dq\xcanO\xba" +
	"mm\x0d,p\xcc\x80\x02\xfc.\xb8\xf3\xe3\xb0\x16\xa8" +
	"\xcbU\xdar\xaf\x96\xa8\x895\xc655f\xd8m\xe4" +
	"kX\x9e\xa9\xa1\x94LD\x86u[\x82(K\xacQ" +
	"\x1a\xa8S5>S\x1df\x1a\x81\x99<xP5m" +
	"I\x80b\xc7\xae\x7f9\xb77\x16\x9a\xf2}v\xc7\xc4" +
	"\xac\xdc5QE\x8d0\x00\x12\xcaC\x00!\xd1\x15\xce" +
	"\x030\xb5\xd1\x9d+.WU^\xc5\xb2\x99\xc7\xdf\xef" +
	"\xae=Z\xb6\xe5\x12\xf9\x17H\xa2\xceOS=\xe4X" +
	"\xa2\xf2V,Q\x9fa\x09Y`4\x81B\x9e&|" +
	"\xd6,:\xb5\xa5\xe9~\x10iY\xa6\xe2\x03\xce\xab8" +
	"\x10z\x9f\x05?\x9c\x83\x15\xc8\x100\xd4\x99\x16x$" +
	"!\x02Zx\x82\xe8\xc2B\xd8k\x01\xec-\xe0\x96\xa1" +
	"\x9e\x08\xee|\x90\xd0\xd9\xde*`\xefO\xb9\x7f\x1e\x11" +
	"\xe4\xdf%\x10\x9c=\x0d\xe7\xe6.\x7f\xfa\xf2\x99}o" +
	"\xc8\x17ZQwV\x02\xd1y \x80\xbd\xe6\xe4\x9e\xe3" +
	"\xa8\xeb\x91 \xc7\xe1p\xb0\xb7\x8a|\xec \xea\xba$" +
	"p9\xcf\x06\xb07\xb9\xdcy\x04uoI 9/" +
	"\x06\xb0\x9f\x01\xf2\xee\xbd\xa8\xdb%A\x9e\xb3\x7f\xc0\xde" +
	"\xd2\xf2\xcb\xcc\xefy\x09r\x9d\xcd\x0e\xf6v\x91\x9fb" +
	"~\x9b$\x13\x07!\x94\xac\x8f\xa8Dj\x086\x06\xa0" +
	"-=\xa0\x010m\x96\"\xfe\x14O\xa1\xc8\xbe\x0aB" +
	"\xf6] \x84\xd3B\xad\xaa\x1bL\xe3c*\xde\x0fl" +
	"\xf6\xf0\xa7\xdcx\x95Mg\x85\x8c\xa7nR\xe0\xec\xf8" +
	"S\xc3\x93F\xc6\x86\x17\xec\xe9-d\xd7%\x00\xd6m" +
	"\xbf\xeb\x8f\x86\x95e\x81\xb9G\x07\xba\xedwHP!" +
	"\xa50\x91\xb5=\x06\xe3\x82\xc1Y\xdf\xaa4\xd69\xc5" +
	"4:\x87j\x08\xbe\x16\xd5\x1bHSE\x0e4\x85Q" +
	"\xe7c\x08-\xc2\x8d\x9d\xca\x84M(4\xb8\xb1[\xc7" +
	"\xc6.\x82\xc2\x0d\xb8\xab\xc5b\x10Q\x96\\\x892\x03" +
	"e\xedH=\xf1t\x1a\x8ec\xcd&M\xb7\xd68\xcf" +
	"\xbbJ\xd2\xd0\xb8\xf14U\xbd\x9aF\xb5e:\x11i" +
	"\x02\xd3c\xf4\xa1\xb3\xa4}o\xackc\x9dS\xca\"" +
	"\xc99\xe9\x16L\x10\xc0o\x11\x82\x0e#\x09\x12\x1d\xbe" +
	"\x1f2C\x8e49r\xd8\x14\x9d\xbd\x94\x07\xb966" +
	"\xa9h\x04\x89w\xd0g\x13\x13V\xa30\xc4uf\xf1" +
	"\"\xee\xdd\x94EvC{G\x0dk\x07\xde\xea!1" +
	"9s\xf1\x0b\xd7\xa3=\xee\xa4\x8b\xe7\xd7\xfe\xbc\xe3\xd4" +
	"\xa1o\xfe\xe1+\xe2\xef\x9df\xd3 F\xf5\xbf\x02\x00" +
	"\x00\xff\xff\xc5\xae\x0c\xf7"

func init() {
	schemas.Register(schema_bf6889795837d1e0,
		0x837afa61d880beb6,
		0x845e8081686d8a0f,
		0x97f9d7a51ffe0741,
		0xa15e445037d1834c,
		0xa9502e5fdabb8b07,
		0xaea67098dc479ce6,
		0xb44df810894a44f4,
		0xb45bb2d206694fe6,
		0xb9147a48c12c807d,
		0xc3b5ced7344b04a6,
		0xc9973f31a90887f9,
		0xe14c1f5321159b8f,
		0xe54437a7b8f52843,
		0xe706d835e9cd64af,
		0xe96193c522f6c05d,
		0xe9e4890dae20b03c,
		0xecebff1662ba10a1,
		0xf910658ae8c6674d,
		0xfdc944999b9296df,
		0xfe777c71f871f413)
}
