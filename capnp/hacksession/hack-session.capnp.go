package hacksession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	email "zenhack.net/go/sandstorm/capnp/email"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	ip "zenhack.net/go/sandstorm/capnp/ip"
	capnp "zombiezen.com/go/capnproto2"
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

func (c HackSessionContext) GetIpNetwork(ctx context.Context, params func(HackSessionContext_getIpNetwork_Params) error, opts ...capnp.CallOption) HackSessionContext_getIpNetwork_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      6,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getIpNetwork",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getIpNetwork_Params{Struct: s}) }
	}
	return HackSessionContext_getIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c HackSessionContext) GetIpInterface(ctx context.Context, params func(HackSessionContext_getIpInterface_Params) error, opts ...capnp.CallOption) HackSessionContext_getIpInterface_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getIpInterface",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getIpInterface_Params{Struct: s}) }
	}
	return HackSessionContext_getIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
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
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
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

	GetIpNetwork(HackSessionContext_getIpNetwork) error

	GetIpInterface(HackSessionContext_getIpInterface) error

	GetUiViewForEndpoint(HackSessionContext_getUiViewForEndpoint) error

	GetSharedPermissions(grain.SessionContext_getSharedPermissions) error

	TieToUser(grain.SessionContext_tieToUser) error

	Offer(grain.SessionContext_offer) error

	Request(grain.SessionContext_request) error

	FulfillRequest(grain.SessionContext_fulfillRequest) error

	Close(grain.SessionContext_close) error

	OpenView(grain.SessionContext_openView) error

	Send(email.EmailSendPort_send) error

	HintAddress(email.EmailSendPort_hintAddress) error
}

func HackSessionContext_ServerToClient(s HackSessionContext_Server) HackSessionContext {
	c, _ := s.(server.Closer)
	return HackSessionContext{Client: server.New(HackSessionContext_Methods(nil, s), c)}
}

func HackSessionContext_Methods(methods []server.Method, s HackSessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 18)
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
			MethodName:    "getIpNetwork",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getIpNetwork{c, opts, HackSessionContext_getIpNetwork_Params{Struct: p}, HackSessionContext_getIpNetwork_Results{Struct: r}}
			return s.GetIpNetwork(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getIpInterface",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getIpInterface{c, opts, HackSessionContext_getIpInterface_Params{Struct: p}, HackSessionContext_getIpInterface_Results{Struct: r}}
			return s.GetIpInterface(call)
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

// HackSessionContext_getIpNetwork holds the arguments for a server call to HackSessionContext.getIpNetwork.
type HackSessionContext_getIpNetwork struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getIpNetwork_Params
	Results HackSessionContext_getIpNetwork_Results
}

// HackSessionContext_getIpInterface holds the arguments for a server call to HackSessionContext.getIpInterface.
type HackSessionContext_getIpInterface struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getIpInterface_Params
	Results HackSessionContext_getIpInterface_Results
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
	if err != nil {
		return HackSessionContext_TokenInfo{}, err
	}
	return HackSessionContext_TokenInfo{st}, nil
}

func NewRootHackSessionContext_TokenInfo(s *capnp.Segment) (HackSessionContext_TokenInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HackSessionContext_TokenInfo{}, err
	}
	return HackSessionContext_TokenInfo{st}, nil
}

func ReadRootHackSessionContext_TokenInfo(msg *capnp.Message) (HackSessionContext_TokenInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_TokenInfo{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_TokenInfo{st}, nil
}

func (s HackSessionContext_TokenInfo) TokenId() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_TokenInfo) SetTokenId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HackSessionContext_TokenInfo) Petname() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_TokenInfo) SetPetname(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s HackSessionContext_TokenInfo) UserInfo() (grain.UserInfo, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return grain.UserInfo{}, err
	}

	ss := capnp.ToStruct(p)

	return grain.UserInfo{Struct: ss}, nil
}

func (s HackSessionContext_TokenInfo) SetUserInfo(v grain.UserInfo) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewUserInfo sets the userInfo field to a newly
// allocated grain.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_TokenInfo) NewUserInfo() (grain.UserInfo, error) {

	ss, err := grain.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return grain.UserInfo{}, err
	}
	err = s.Struct.SetPointer(2, ss)
	return ss, err
}

// HackSessionContext_TokenInfo_List is a list of HackSessionContext_TokenInfo.
type HackSessionContext_TokenInfo_List struct{ capnp.List }

// NewHackSessionContext_TokenInfo creates a new list of HackSessionContext_TokenInfo.
func NewHackSessionContext_TokenInfo_List(s *capnp.Segment, sz int32) (HackSessionContext_TokenInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return HackSessionContext_TokenInfo_List{}, err
	}
	return HackSessionContext_TokenInfo_List{l}, nil
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

func (p HackSessionContext_TokenInfo_Promise) UserInfo() grain.UserInfo_Promise {
	return grain.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type HackSessionContext_getPublicId_Params struct{ capnp.Struct }

func NewHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getPublicId_Params{}, err
	}
	return HackSessionContext_getPublicId_Params{st}, nil
}

func NewRootHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getPublicId_Params{}, err
	}
	return HackSessionContext_getPublicId_Params{st}, nil
}

func ReadRootHackSessionContext_getPublicId_Params(msg *capnp.Message) (HackSessionContext_getPublicId_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getPublicId_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getPublicId_Params{st}, nil
}

// HackSessionContext_getPublicId_Params_List is a list of HackSessionContext_getPublicId_Params.
type HackSessionContext_getPublicId_Params_List struct{ capnp.List }

// NewHackSessionContext_getPublicId_Params creates a new list of HackSessionContext_getPublicId_Params.
func NewHackSessionContext_getPublicId_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getPublicId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_getPublicId_Params_List{}, err
	}
	return HackSessionContext_getPublicId_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_getPublicId_Results{}, err
	}
	return HackSessionContext_getPublicId_Results{st}, nil
}

func NewRootHackSessionContext_getPublicId_Results(s *capnp.Segment) (HackSessionContext_getPublicId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return HackSessionContext_getPublicId_Results{}, err
	}
	return HackSessionContext_getPublicId_Results{st}, nil
}

func ReadRootHackSessionContext_getPublicId_Results(msg *capnp.Message) (HackSessionContext_getPublicId_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getPublicId_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getPublicId_Results{st}, nil
}

func (s HackSessionContext_getPublicId_Results) PublicId() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_getPublicId_Results) SetPublicId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HackSessionContext_getPublicId_Results) Hostname() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_getPublicId_Results) SetHostname(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s HackSessionContext_getPublicId_Results) AutoUrl() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_getPublicId_Results) SetAutoUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
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
	if err != nil {
		return HackSessionContext_getPublicId_Results_List{}, err
	}
	return HackSessionContext_getPublicId_Results_List{l}, nil
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
	if err != nil {
		return HackSessionContext_httpGet_Params{}, err
	}
	return HackSessionContext_httpGet_Params{st}, nil
}

func NewRootHackSessionContext_httpGet_Params(s *capnp.Segment) (HackSessionContext_httpGet_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_httpGet_Params{}, err
	}
	return HackSessionContext_httpGet_Params{st}, nil
}

func ReadRootHackSessionContext_httpGet_Params(msg *capnp.Message) (HackSessionContext_httpGet_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_httpGet_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_httpGet_Params{st}, nil
}

func (s HackSessionContext_httpGet_Params) Url() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_httpGet_Params) SetUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// HackSessionContext_httpGet_Params_List is a list of HackSessionContext_httpGet_Params.
type HackSessionContext_httpGet_Params_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Params creates a new list of HackSessionContext_httpGet_Params.
func NewHackSessionContext_httpGet_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_httpGet_Params_List{}, err
	}
	return HackSessionContext_httpGet_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_httpGet_Results{}, err
	}
	return HackSessionContext_httpGet_Results{st}, nil
}

func NewRootHackSessionContext_httpGet_Results(s *capnp.Segment) (HackSessionContext_httpGet_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return HackSessionContext_httpGet_Results{}, err
	}
	return HackSessionContext_httpGet_Results{st}, nil
}

func ReadRootHackSessionContext_httpGet_Results(msg *capnp.Message) (HackSessionContext_httpGet_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_httpGet_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_httpGet_Results{st}, nil
}

func (s HackSessionContext_httpGet_Results) MimeType() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_httpGet_Results) SetMimeType(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HackSessionContext_httpGet_Results) Content() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s HackSessionContext_httpGet_Results) SetContent(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
}

// HackSessionContext_httpGet_Results_List is a list of HackSessionContext_httpGet_Results.
type HackSessionContext_httpGet_Results_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Results creates a new list of HackSessionContext_httpGet_Results.
func NewHackSessionContext_httpGet_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return HackSessionContext_httpGet_Results_List{}, err
	}
	return HackSessionContext_httpGet_Results_List{l}, nil
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
	if err != nil {
		return HackSessionContext_getUserAddress_Params{}, err
	}
	return HackSessionContext_getUserAddress_Params{st}, nil
}

func NewRootHackSessionContext_getUserAddress_Params(s *capnp.Segment) (HackSessionContext_getUserAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getUserAddress_Params{}, err
	}
	return HackSessionContext_getUserAddress_Params{st}, nil
}

func ReadRootHackSessionContext_getUserAddress_Params(msg *capnp.Message) (HackSessionContext_getUserAddress_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getUserAddress_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getUserAddress_Params{st}, nil
}

// HackSessionContext_getUserAddress_Params_List is a list of HackSessionContext_getUserAddress_Params.
type HackSessionContext_getUserAddress_Params_List struct{ capnp.List }

// NewHackSessionContext_getUserAddress_Params creates a new list of HackSessionContext_getUserAddress_Params.
func NewHackSessionContext_getUserAddress_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUserAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_getUserAddress_Params_List{}, err
	}
	return HackSessionContext_getUserAddress_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Params{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, nil
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Params{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, nil
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) Petname() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetPetname(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) UserInfo() (grain.UserInfo, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return grain.UserInfo{}, err
	}

	ss := capnp.ToStruct(p)

	return grain.UserInfo{Struct: ss}, nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetUserInfo(v grain.UserInfo) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewUserInfo sets the userInfo field to a newly
// allocated grain.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_obsoleteGenerateApiToken_Params) NewUserInfo() (grain.UserInfo, error) {

	ss, err := grain.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return grain.UserInfo{}, err
	}
	err = s.Struct.SetPointer(1, ss)
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
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Params_List{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Params_List{l}, nil
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

func (p HackSessionContext_obsoleteGenerateApiToken_Params_Promise) UserInfo() grain.UserInfo_Promise {
	return grain.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HackSessionContext_obsoleteGenerateApiToken_Results struct{ capnp.Struct }

func NewHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, nil
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, nil
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) Token() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetToken(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) EndpointUrl() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetEndpointUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenId() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetTokenId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

// HackSessionContext_obsoleteGenerateApiToken_Results_List is a list of HackSessionContext_obsoleteGenerateApiToken_Results.
type HackSessionContext_obsoleteGenerateApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGenerateApiToken_Results creates a new list of HackSessionContext_obsoleteGenerateApiToken_Results.
func NewHackSessionContext_obsoleteGenerateApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGenerateApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results_List{}, err
	}
	return HackSessionContext_obsoleteGenerateApiToken_Results_List{l}, nil
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
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Params{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Params{st}, nil
}

func NewRootHackSessionContext_obsoleteListApiTokens_Params(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Params{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Params{st}, nil
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Params(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteListApiTokens_Params{st}, nil
}

// HackSessionContext_obsoleteListApiTokens_Params_List is a list of HackSessionContext_obsoleteListApiTokens_Params.
type HackSessionContext_obsoleteListApiTokens_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Params creates a new list of HackSessionContext_obsoleteListApiTokens_Params.
func NewHackSessionContext_obsoleteListApiTokens_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Params_List{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Results{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Results{st}, nil
}

func NewRootHackSessionContext_obsoleteListApiTokens_Results(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Results{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Results{st}, nil
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Results(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteListApiTokens_Results{st}, nil
}

func (s HackSessionContext_obsoleteListApiTokens_Results) Tokens() (HackSessionContext_TokenInfo_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return HackSessionContext_TokenInfo_List{}, err
	}

	l := capnp.ToList(p)

	return HackSessionContext_TokenInfo_List{List: l}, nil
}

func (s HackSessionContext_obsoleteListApiTokens_Results) SetTokens(v HackSessionContext_TokenInfo_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// HackSessionContext_obsoleteListApiTokens_Results_List is a list of HackSessionContext_obsoleteListApiTokens_Results.
type HackSessionContext_obsoleteListApiTokens_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Results creates a new list of HackSessionContext_obsoleteListApiTokens_Results.
func NewHackSessionContext_obsoleteListApiTokens_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_obsoleteListApiTokens_Results_List{}, err
	}
	return HackSessionContext_obsoleteListApiTokens_Results_List{l}, nil
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
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Params{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, nil
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Params{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, nil
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, nil
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) TokenId() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) SetTokenId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// HackSessionContext_obsoleteRevokeApiToken_Params_List is a list of HackSessionContext_obsoleteRevokeApiToken_Params.
type HackSessionContext_obsoleteRevokeApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Params creates a new list of HackSessionContext_obsoleteRevokeApiToken_Params.
func NewHackSessionContext_obsoleteRevokeApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Params_List{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, nil
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, nil
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, nil
}

// HackSessionContext_obsoleteRevokeApiToken_Results_List is a list of HackSessionContext_obsoleteRevokeApiToken_Results.
type HackSessionContext_obsoleteRevokeApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Results creates a new list of HackSessionContext_obsoleteRevokeApiToken_Results.
func NewHackSessionContext_obsoleteRevokeApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results_List{}, err
	}
	return HackSessionContext_obsoleteRevokeApiToken_Results_List{l}, nil
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

type HackSessionContext_getIpNetwork_Params struct{ capnp.Struct }

func NewHackSessionContext_getIpNetwork_Params(s *capnp.Segment) (HackSessionContext_getIpNetwork_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getIpNetwork_Params{}, err
	}
	return HackSessionContext_getIpNetwork_Params{st}, nil
}

func NewRootHackSessionContext_getIpNetwork_Params(s *capnp.Segment) (HackSessionContext_getIpNetwork_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getIpNetwork_Params{}, err
	}
	return HackSessionContext_getIpNetwork_Params{st}, nil
}

func ReadRootHackSessionContext_getIpNetwork_Params(msg *capnp.Message) (HackSessionContext_getIpNetwork_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getIpNetwork_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getIpNetwork_Params{st}, nil
}

// HackSessionContext_getIpNetwork_Params_List is a list of HackSessionContext_getIpNetwork_Params.
type HackSessionContext_getIpNetwork_Params_List struct{ capnp.List }

// NewHackSessionContext_getIpNetwork_Params creates a new list of HackSessionContext_getIpNetwork_Params.
func NewHackSessionContext_getIpNetwork_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getIpNetwork_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_getIpNetwork_Params_List{}, err
	}
	return HackSessionContext_getIpNetwork_Params_List{l}, nil
}

func (s HackSessionContext_getIpNetwork_Params_List) At(i int) HackSessionContext_getIpNetwork_Params {
	return HackSessionContext_getIpNetwork_Params{s.List.Struct(i)}
}
func (s HackSessionContext_getIpNetwork_Params_List) Set(i int, v HackSessionContext_getIpNetwork_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getIpNetwork_Params_Promise is a wrapper for a HackSessionContext_getIpNetwork_Params promised by a client call.
type HackSessionContext_getIpNetwork_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getIpNetwork_Params_Promise) Struct() (HackSessionContext_getIpNetwork_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getIpNetwork_Params{s}, err
}

type HackSessionContext_getIpNetwork_Results struct{ capnp.Struct }

func NewHackSessionContext_getIpNetwork_Results(s *capnp.Segment) (HackSessionContext_getIpNetwork_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getIpNetwork_Results{}, err
	}
	return HackSessionContext_getIpNetwork_Results{st}, nil
}

func NewRootHackSessionContext_getIpNetwork_Results(s *capnp.Segment) (HackSessionContext_getIpNetwork_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getIpNetwork_Results{}, err
	}
	return HackSessionContext_getIpNetwork_Results{st}, nil
}

func ReadRootHackSessionContext_getIpNetwork_Results(msg *capnp.Message) (HackSessionContext_getIpNetwork_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getIpNetwork_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getIpNetwork_Results{st}, nil
}

func (s HackSessionContext_getIpNetwork_Results) Network() ip.IpNetwork {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return ip.IpNetwork{}
	}
	c := capnp.ToInterface(p).Client()
	return ip.IpNetwork{Client: c}
}

func (s HackSessionContext_getIpNetwork_Results) SetNetwork(v ip.IpNetwork) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(0, in)
}

// HackSessionContext_getIpNetwork_Results_List is a list of HackSessionContext_getIpNetwork_Results.
type HackSessionContext_getIpNetwork_Results_List struct{ capnp.List }

// NewHackSessionContext_getIpNetwork_Results creates a new list of HackSessionContext_getIpNetwork_Results.
func NewHackSessionContext_getIpNetwork_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getIpNetwork_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_getIpNetwork_Results_List{}, err
	}
	return HackSessionContext_getIpNetwork_Results_List{l}, nil
}

func (s HackSessionContext_getIpNetwork_Results_List) At(i int) HackSessionContext_getIpNetwork_Results {
	return HackSessionContext_getIpNetwork_Results{s.List.Struct(i)}
}
func (s HackSessionContext_getIpNetwork_Results_List) Set(i int, v HackSessionContext_getIpNetwork_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getIpNetwork_Results_Promise is a wrapper for a HackSessionContext_getIpNetwork_Results promised by a client call.
type HackSessionContext_getIpNetwork_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getIpNetwork_Results_Promise) Struct() (HackSessionContext_getIpNetwork_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getIpNetwork_Results{s}, err
}

func (p HackSessionContext_getIpNetwork_Results_Promise) Network() ip.IpNetwork {
	return ip.IpNetwork{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_getIpInterface_Params struct{ capnp.Struct }

func NewHackSessionContext_getIpInterface_Params(s *capnp.Segment) (HackSessionContext_getIpInterface_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getIpInterface_Params{}, err
	}
	return HackSessionContext_getIpInterface_Params{st}, nil
}

func NewRootHackSessionContext_getIpInterface_Params(s *capnp.Segment) (HackSessionContext_getIpInterface_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_getIpInterface_Params{}, err
	}
	return HackSessionContext_getIpInterface_Params{st}, nil
}

func ReadRootHackSessionContext_getIpInterface_Params(msg *capnp.Message) (HackSessionContext_getIpInterface_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getIpInterface_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getIpInterface_Params{st}, nil
}

// HackSessionContext_getIpInterface_Params_List is a list of HackSessionContext_getIpInterface_Params.
type HackSessionContext_getIpInterface_Params_List struct{ capnp.List }

// NewHackSessionContext_getIpInterface_Params creates a new list of HackSessionContext_getIpInterface_Params.
func NewHackSessionContext_getIpInterface_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getIpInterface_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_getIpInterface_Params_List{}, err
	}
	return HackSessionContext_getIpInterface_Params_List{l}, nil
}

func (s HackSessionContext_getIpInterface_Params_List) At(i int) HackSessionContext_getIpInterface_Params {
	return HackSessionContext_getIpInterface_Params{s.List.Struct(i)}
}
func (s HackSessionContext_getIpInterface_Params_List) Set(i int, v HackSessionContext_getIpInterface_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getIpInterface_Params_Promise is a wrapper for a HackSessionContext_getIpInterface_Params promised by a client call.
type HackSessionContext_getIpInterface_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getIpInterface_Params_Promise) Struct() (HackSessionContext_getIpInterface_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getIpInterface_Params{s}, err
}

type HackSessionContext_getIpInterface_Results struct{ capnp.Struct }

func NewHackSessionContext_getIpInterface_Results(s *capnp.Segment) (HackSessionContext_getIpInterface_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getIpInterface_Results{}, err
	}
	return HackSessionContext_getIpInterface_Results{st}, nil
}

func NewRootHackSessionContext_getIpInterface_Results(s *capnp.Segment) (HackSessionContext_getIpInterface_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getIpInterface_Results{}, err
	}
	return HackSessionContext_getIpInterface_Results{st}, nil
}

func ReadRootHackSessionContext_getIpInterface_Results(msg *capnp.Message) (HackSessionContext_getIpInterface_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getIpInterface_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getIpInterface_Results{st}, nil
}

func (s HackSessionContext_getIpInterface_Results) Interface() ip.IpInterface {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return ip.IpInterface{}
	}
	c := capnp.ToInterface(p).Client()
	return ip.IpInterface{Client: c}
}

func (s HackSessionContext_getIpInterface_Results) SetInterface(v ip.IpInterface) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(0, in)
}

// HackSessionContext_getIpInterface_Results_List is a list of HackSessionContext_getIpInterface_Results.
type HackSessionContext_getIpInterface_Results_List struct{ capnp.List }

// NewHackSessionContext_getIpInterface_Results creates a new list of HackSessionContext_getIpInterface_Results.
func NewHackSessionContext_getIpInterface_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getIpInterface_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_getIpInterface_Results_List{}, err
	}
	return HackSessionContext_getIpInterface_Results_List{l}, nil
}

func (s HackSessionContext_getIpInterface_Results_List) At(i int) HackSessionContext_getIpInterface_Results {
	return HackSessionContext_getIpInterface_Results{s.List.Struct(i)}
}
func (s HackSessionContext_getIpInterface_Results_List) Set(i int, v HackSessionContext_getIpInterface_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getIpInterface_Results_Promise is a wrapper for a HackSessionContext_getIpInterface_Results promised by a client call.
type HackSessionContext_getIpInterface_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getIpInterface_Results_Promise) Struct() (HackSessionContext_getIpInterface_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getIpInterface_Results{s}, err
}

func (p HackSessionContext_getIpInterface_Results_Promise) Interface() ip.IpInterface {
	return ip.IpInterface{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_getUiViewForEndpoint_Params struct{ capnp.Struct }

func NewHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Params{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Params{st}, nil
}

func NewRootHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Params{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Params{st}, nil
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Params(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Params{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getUiViewForEndpoint_Params{st}, nil
}

func (s HackSessionContext_getUiViewForEndpoint_Params) Url() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s HackSessionContext_getUiViewForEndpoint_Params) SetUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// HackSessionContext_getUiViewForEndpoint_Params_List is a list of HackSessionContext_getUiViewForEndpoint_Params.
type HackSessionContext_getUiViewForEndpoint_Params_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Params creates a new list of HackSessionContext_getUiViewForEndpoint_Params.
func NewHackSessionContext_getUiViewForEndpoint_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Params_List{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Params_List{l}, nil
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
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Results{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Results{st}, nil
}

func NewRootHackSessionContext_getUiViewForEndpoint_Results(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Results{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Results{st}, nil
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Results(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Results{}, err
	}
	st := capnp.ToStruct(root)
	return HackSessionContext_getUiViewForEndpoint_Results{st}, nil
}

func (s HackSessionContext_getUiViewForEndpoint_Results) View() grain.UiView {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return grain.UiView{}
	}
	c := capnp.ToInterface(p).Client()
	return grain.UiView{Client: c}
}

func (s HackSessionContext_getUiViewForEndpoint_Results) SetView(v grain.UiView) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(0, in)
}

// HackSessionContext_getUiViewForEndpoint_Results_List is a list of HackSessionContext_getUiViewForEndpoint_Results.
type HackSessionContext_getUiViewForEndpoint_Results_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Results creates a new list of HackSessionContext_getUiViewForEndpoint_Results.
func NewHackSessionContext_getUiViewForEndpoint_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Results_List{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Results_List{l}, nil
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
