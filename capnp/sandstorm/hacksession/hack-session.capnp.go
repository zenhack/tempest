package hacksession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	email "zenhack.net/go/sandstorm/capnp/sandstorm/email"
	grain "zenhack.net/go/sandstorm/capnp/sandstorm/grain"
	ip "zenhack.net/go/sandstorm/capnp/sandstorm/ip"
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

func (c HackSessionContext) GenerateApiToken(ctx context.Context, params func(HackSessionContext_generateApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_generateApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_generateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      3,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "generateApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_generateApiToken_Params{Struct: s}) }
	}
	return HackSessionContext_generateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c HackSessionContext) ListApiTokens(ctx context.Context, params func(HackSessionContext_listApiTokens_Params) error, opts ...capnp.CallOption) HackSessionContext_listApiTokens_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_listApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "listApiTokens",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_listApiTokens_Params{Struct: s}) }
	}
	return HackSessionContext_listApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c HackSessionContext) RevokeApiToken(ctx context.Context, params func(HackSessionContext_revokeApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_revokeApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_revokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "revokeApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_revokeApiToken_Params{Struct: s}) }
	}
	return HackSessionContext_revokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
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

func (c HackSessionContext) Provide(ctx context.Context, params func(grain.SessionContext_provide_Params) error, opts ...capnp.CallOption) grain.SessionContext_provide_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_provide_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "provide",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_provide_Params{Struct: s}) }
	}
	return grain.SessionContext_provide_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
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

type HackSessionContext_Server interface {
	GetPublicId(HackSessionContext_getPublicId) error

	HttpGet(HackSessionContext_httpGet) error

	GetUserAddress(HackSessionContext_getUserAddress) error

	GenerateApiToken(HackSessionContext_generateApiToken) error

	ListApiTokens(HackSessionContext_listApiTokens) error

	RevokeApiToken(HackSessionContext_revokeApiToken) error

	GetIpNetwork(HackSessionContext_getIpNetwork) error

	GetIpInterface(HackSessionContext_getIpInterface) error

	GetUiViewForEndpoint(HackSessionContext_getUiViewForEndpoint) error

	GetSharedPermissions(grain.SessionContext_getSharedPermissions) error

	TieToUser(grain.SessionContext_tieToUser) error

	Offer(grain.SessionContext_offer) error

	Request(grain.SessionContext_request) error

	Provide(grain.SessionContext_provide) error

	Close(grain.SessionContext_close) error

	OpenView(grain.SessionContext_openView) error

	Send(email.EmailSendPort_send) error
}

func HackSessionContext_ServerToClient(s HackSessionContext_Server) HackSessionContext {
	c, _ := s.(server.Closer)
	return HackSessionContext{Client: server.New(HackSessionContext_Methods(nil, s), c)}
}

func HackSessionContext_Methods(methods []server.Method, s HackSessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 17)
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
			MethodName:    "generateApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_generateApiToken{c, opts, HackSessionContext_generateApiToken_Params{Struct: p}, HackSessionContext_generateApiToken_Results{Struct: r}}
			return s.GenerateApiToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "listApiTokens",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_listApiTokens{c, opts, HackSessionContext_listApiTokens_Params{Struct: p}, HackSessionContext_listApiTokens_Results{Struct: r}}
			return s.ListApiTokens(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "revokeApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_revokeApiToken{c, opts, HackSessionContext_revokeApiToken_Params{Struct: p}, HackSessionContext_revokeApiToken_Results{Struct: r}}
			return s.RevokeApiToken(call)
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
			MethodName:    "provide",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_provide{c, opts, grain.SessionContext_provide_Params{Struct: p}, grain.SessionContext_provide_Results{Struct: r}}
			return s.Provide(call)
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

// HackSessionContext_generateApiToken holds the arguments for a server call to HackSessionContext.generateApiToken.
type HackSessionContext_generateApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_generateApiToken_Params
	Results HackSessionContext_generateApiToken_Results
}

// HackSessionContext_listApiTokens holds the arguments for a server call to HackSessionContext.listApiTokens.
type HackSessionContext_listApiTokens struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_listApiTokens_Params
	Results HackSessionContext_listApiTokens_Results
}

// HackSessionContext_revokeApiToken holds the arguments for a server call to HackSessionContext.revokeApiToken.
type HackSessionContext_revokeApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_revokeApiToken_Params
	Results HackSessionContext_revokeApiToken_Results
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_TokenInfo{}, err
	}
	return HackSessionContext_TokenInfo{root.Struct()}, nil
}

func (s HackSessionContext_TokenInfo) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_TokenInfo) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_TokenInfo) HasPetname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_TokenInfo) SetPetname(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_TokenInfo) UserInfo() (grain.UserInfo, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return grain.UserInfo{}, err
	}

	return grain.UserInfo{Struct: p.Struct()}, nil

}

func (s HackSessionContext_TokenInfo) HasUserInfo() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) SetUserInfo(v grain.UserInfo) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated grain.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_TokenInfo) NewUserInfo() (grain.UserInfo, error) {

	ss, err := grain.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return grain.UserInfo{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getPublicId_Params{}, err
	}
	return HackSessionContext_getPublicId_Params{root.Struct()}, nil
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getPublicId_Results{}, err
	}
	return HackSessionContext_getPublicId_Results{root.Struct()}, nil
}

func (s HackSessionContext_getPublicId_Results) PublicId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_getPublicId_Results) HasPublicId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) PublicIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_getPublicId_Results) HasHostname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_getPublicId_Results) HasAutoUrl() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) AutoUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_httpGet_Params{}, err
	}
	return HackSessionContext_httpGet_Params{root.Struct()}, nil
}

func (s HackSessionContext_httpGet_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_httpGet_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_httpGet_Results{}, err
	}
	return HackSessionContext_httpGet_Results{root.Struct()}, nil
}

func (s HackSessionContext_httpGet_Results) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_httpGet_Results) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Results) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getUserAddress_Params{}, err
	}
	return HackSessionContext_getUserAddress_Params{root.Struct()}, nil
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

type HackSessionContext_generateApiToken_Params struct{ capnp.Struct }

func NewHackSessionContext_generateApiToken_Params(s *capnp.Segment) (HackSessionContext_generateApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return HackSessionContext_generateApiToken_Params{}, err
	}
	return HackSessionContext_generateApiToken_Params{st}, nil
}

func NewRootHackSessionContext_generateApiToken_Params(s *capnp.Segment) (HackSessionContext_generateApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return HackSessionContext_generateApiToken_Params{}, err
	}
	return HackSessionContext_generateApiToken_Params{st}, nil
}

func ReadRootHackSessionContext_generateApiToken_Params(msg *capnp.Message) (HackSessionContext_generateApiToken_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_generateApiToken_Params{}, err
	}
	return HackSessionContext_generateApiToken_Params{root.Struct()}, nil
}

func (s HackSessionContext_generateApiToken_Params) Petname() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_generateApiToken_Params) HasPetname() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_generateApiToken_Params) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_generateApiToken_Params) SetPetname(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_generateApiToken_Params) UserInfo() (grain.UserInfo, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return grain.UserInfo{}, err
	}

	return grain.UserInfo{Struct: p.Struct()}, nil

}

func (s HackSessionContext_generateApiToken_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_generateApiToken_Params) SetUserInfo(v grain.UserInfo) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated grain.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_generateApiToken_Params) NewUserInfo() (grain.UserInfo, error) {

	ss, err := grain.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return grain.UserInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s HackSessionContext_generateApiToken_Params) Expires() uint64 {
	return s.Struct.Uint64(0)
}

func (s HackSessionContext_generateApiToken_Params) SetExpires(v uint64) {

	s.Struct.SetUint64(0, v)
}

// HackSessionContext_generateApiToken_Params_List is a list of HackSessionContext_generateApiToken_Params.
type HackSessionContext_generateApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_generateApiToken_Params creates a new list of HackSessionContext_generateApiToken_Params.
func NewHackSessionContext_generateApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_generateApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return HackSessionContext_generateApiToken_Params_List{}, err
	}
	return HackSessionContext_generateApiToken_Params_List{l}, nil
}

func (s HackSessionContext_generateApiToken_Params_List) At(i int) HackSessionContext_generateApiToken_Params {
	return HackSessionContext_generateApiToken_Params{s.List.Struct(i)}
}
func (s HackSessionContext_generateApiToken_Params_List) Set(i int, v HackSessionContext_generateApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_generateApiToken_Params_Promise is a wrapper for a HackSessionContext_generateApiToken_Params promised by a client call.
type HackSessionContext_generateApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_generateApiToken_Params_Promise) Struct() (HackSessionContext_generateApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_generateApiToken_Params{s}, err
}

func (p HackSessionContext_generateApiToken_Params_Promise) UserInfo() grain.UserInfo_Promise {
	return grain.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HackSessionContext_generateApiToken_Results struct{ capnp.Struct }

func NewHackSessionContext_generateApiToken_Results(s *capnp.Segment) (HackSessionContext_generateApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HackSessionContext_generateApiToken_Results{}, err
	}
	return HackSessionContext_generateApiToken_Results{st}, nil
}

func NewRootHackSessionContext_generateApiToken_Results(s *capnp.Segment) (HackSessionContext_generateApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return HackSessionContext_generateApiToken_Results{}, err
	}
	return HackSessionContext_generateApiToken_Results{st}, nil
}

func ReadRootHackSessionContext_generateApiToken_Results(msg *capnp.Message) (HackSessionContext_generateApiToken_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_generateApiToken_Results{}, err
	}
	return HackSessionContext_generateApiToken_Results{root.Struct()}, nil
}

func (s HackSessionContext_generateApiToken_Results) Token() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_generateApiToken_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_generateApiToken_Results) TokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_generateApiToken_Results) SetToken(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_generateApiToken_Results) EndpointUrl() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_generateApiToken_Results) HasEndpointUrl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_generateApiToken_Results) EndpointUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_generateApiToken_Results) SetEndpointUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_generateApiToken_Results) TokenId() (string, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_generateApiToken_Results) HasTokenId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_generateApiToken_Results) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_generateApiToken_Results) SetTokenId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// HackSessionContext_generateApiToken_Results_List is a list of HackSessionContext_generateApiToken_Results.
type HackSessionContext_generateApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_generateApiToken_Results creates a new list of HackSessionContext_generateApiToken_Results.
func NewHackSessionContext_generateApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_generateApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return HackSessionContext_generateApiToken_Results_List{}, err
	}
	return HackSessionContext_generateApiToken_Results_List{l}, nil
}

func (s HackSessionContext_generateApiToken_Results_List) At(i int) HackSessionContext_generateApiToken_Results {
	return HackSessionContext_generateApiToken_Results{s.List.Struct(i)}
}
func (s HackSessionContext_generateApiToken_Results_List) Set(i int, v HackSessionContext_generateApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_generateApiToken_Results_Promise is a wrapper for a HackSessionContext_generateApiToken_Results promised by a client call.
type HackSessionContext_generateApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_generateApiToken_Results_Promise) Struct() (HackSessionContext_generateApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_generateApiToken_Results{s}, err
}

type HackSessionContext_listApiTokens_Params struct{ capnp.Struct }

func NewHackSessionContext_listApiTokens_Params(s *capnp.Segment) (HackSessionContext_listApiTokens_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_listApiTokens_Params{}, err
	}
	return HackSessionContext_listApiTokens_Params{st}, nil
}

func NewRootHackSessionContext_listApiTokens_Params(s *capnp.Segment) (HackSessionContext_listApiTokens_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_listApiTokens_Params{}, err
	}
	return HackSessionContext_listApiTokens_Params{st}, nil
}

func ReadRootHackSessionContext_listApiTokens_Params(msg *capnp.Message) (HackSessionContext_listApiTokens_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_listApiTokens_Params{}, err
	}
	return HackSessionContext_listApiTokens_Params{root.Struct()}, nil
}

// HackSessionContext_listApiTokens_Params_List is a list of HackSessionContext_listApiTokens_Params.
type HackSessionContext_listApiTokens_Params_List struct{ capnp.List }

// NewHackSessionContext_listApiTokens_Params creates a new list of HackSessionContext_listApiTokens_Params.
func NewHackSessionContext_listApiTokens_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_listApiTokens_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_listApiTokens_Params_List{}, err
	}
	return HackSessionContext_listApiTokens_Params_List{l}, nil
}

func (s HackSessionContext_listApiTokens_Params_List) At(i int) HackSessionContext_listApiTokens_Params {
	return HackSessionContext_listApiTokens_Params{s.List.Struct(i)}
}
func (s HackSessionContext_listApiTokens_Params_List) Set(i int, v HackSessionContext_listApiTokens_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_listApiTokens_Params_Promise is a wrapper for a HackSessionContext_listApiTokens_Params promised by a client call.
type HackSessionContext_listApiTokens_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_listApiTokens_Params_Promise) Struct() (HackSessionContext_listApiTokens_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_listApiTokens_Params{s}, err
}

type HackSessionContext_listApiTokens_Results struct{ capnp.Struct }

func NewHackSessionContext_listApiTokens_Results(s *capnp.Segment) (HackSessionContext_listApiTokens_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_listApiTokens_Results{}, err
	}
	return HackSessionContext_listApiTokens_Results{st}, nil
}

func NewRootHackSessionContext_listApiTokens_Results(s *capnp.Segment) (HackSessionContext_listApiTokens_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_listApiTokens_Results{}, err
	}
	return HackSessionContext_listApiTokens_Results{st}, nil
}

func ReadRootHackSessionContext_listApiTokens_Results(msg *capnp.Message) (HackSessionContext_listApiTokens_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_listApiTokens_Results{}, err
	}
	return HackSessionContext_listApiTokens_Results{root.Struct()}, nil
}

func (s HackSessionContext_listApiTokens_Results) Tokens() (HackSessionContext_TokenInfo_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return HackSessionContext_TokenInfo_List{}, err
	}

	return HackSessionContext_TokenInfo_List{List: p.List()}, nil

}

func (s HackSessionContext_listApiTokens_Results) HasTokens() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_listApiTokens_Results) SetTokens(v HackSessionContext_TokenInfo_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// HackSessionContext_listApiTokens_Results_List is a list of HackSessionContext_listApiTokens_Results.
type HackSessionContext_listApiTokens_Results_List struct{ capnp.List }

// NewHackSessionContext_listApiTokens_Results creates a new list of HackSessionContext_listApiTokens_Results.
func NewHackSessionContext_listApiTokens_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_listApiTokens_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_listApiTokens_Results_List{}, err
	}
	return HackSessionContext_listApiTokens_Results_List{l}, nil
}

func (s HackSessionContext_listApiTokens_Results_List) At(i int) HackSessionContext_listApiTokens_Results {
	return HackSessionContext_listApiTokens_Results{s.List.Struct(i)}
}
func (s HackSessionContext_listApiTokens_Results_List) Set(i int, v HackSessionContext_listApiTokens_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_listApiTokens_Results_Promise is a wrapper for a HackSessionContext_listApiTokens_Results promised by a client call.
type HackSessionContext_listApiTokens_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_listApiTokens_Results_Promise) Struct() (HackSessionContext_listApiTokens_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_listApiTokens_Results{s}, err
}

type HackSessionContext_revokeApiToken_Params struct{ capnp.Struct }

func NewHackSessionContext_revokeApiToken_Params(s *capnp.Segment) (HackSessionContext_revokeApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_revokeApiToken_Params{}, err
	}
	return HackSessionContext_revokeApiToken_Params{st}, nil
}

func NewRootHackSessionContext_revokeApiToken_Params(s *capnp.Segment) (HackSessionContext_revokeApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return HackSessionContext_revokeApiToken_Params{}, err
	}
	return HackSessionContext_revokeApiToken_Params{st}, nil
}

func ReadRootHackSessionContext_revokeApiToken_Params(msg *capnp.Message) (HackSessionContext_revokeApiToken_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_revokeApiToken_Params{}, err
	}
	return HackSessionContext_revokeApiToken_Params{root.Struct()}, nil
}

func (s HackSessionContext_revokeApiToken_Params) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_revokeApiToken_Params) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_revokeApiToken_Params) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s HackSessionContext_revokeApiToken_Params) SetTokenId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_revokeApiToken_Params_List is a list of HackSessionContext_revokeApiToken_Params.
type HackSessionContext_revokeApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_revokeApiToken_Params creates a new list of HackSessionContext_revokeApiToken_Params.
func NewHackSessionContext_revokeApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_revokeApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return HackSessionContext_revokeApiToken_Params_List{}, err
	}
	return HackSessionContext_revokeApiToken_Params_List{l}, nil
}

func (s HackSessionContext_revokeApiToken_Params_List) At(i int) HackSessionContext_revokeApiToken_Params {
	return HackSessionContext_revokeApiToken_Params{s.List.Struct(i)}
}
func (s HackSessionContext_revokeApiToken_Params_List) Set(i int, v HackSessionContext_revokeApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_revokeApiToken_Params_Promise is a wrapper for a HackSessionContext_revokeApiToken_Params promised by a client call.
type HackSessionContext_revokeApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_revokeApiToken_Params_Promise) Struct() (HackSessionContext_revokeApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_revokeApiToken_Params{s}, err
}

type HackSessionContext_revokeApiToken_Results struct{ capnp.Struct }

func NewHackSessionContext_revokeApiToken_Results(s *capnp.Segment) (HackSessionContext_revokeApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_revokeApiToken_Results{}, err
	}
	return HackSessionContext_revokeApiToken_Results{st}, nil
}

func NewRootHackSessionContext_revokeApiToken_Results(s *capnp.Segment) (HackSessionContext_revokeApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return HackSessionContext_revokeApiToken_Results{}, err
	}
	return HackSessionContext_revokeApiToken_Results{st}, nil
}

func ReadRootHackSessionContext_revokeApiToken_Results(msg *capnp.Message) (HackSessionContext_revokeApiToken_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_revokeApiToken_Results{}, err
	}
	return HackSessionContext_revokeApiToken_Results{root.Struct()}, nil
}

// HackSessionContext_revokeApiToken_Results_List is a list of HackSessionContext_revokeApiToken_Results.
type HackSessionContext_revokeApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_revokeApiToken_Results creates a new list of HackSessionContext_revokeApiToken_Results.
func NewHackSessionContext_revokeApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_revokeApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return HackSessionContext_revokeApiToken_Results_List{}, err
	}
	return HackSessionContext_revokeApiToken_Results_List{l}, nil
}

func (s HackSessionContext_revokeApiToken_Results_List) At(i int) HackSessionContext_revokeApiToken_Results {
	return HackSessionContext_revokeApiToken_Results{s.List.Struct(i)}
}
func (s HackSessionContext_revokeApiToken_Results_List) Set(i int, v HackSessionContext_revokeApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_revokeApiToken_Results_Promise is a wrapper for a HackSessionContext_revokeApiToken_Results promised by a client call.
type HackSessionContext_revokeApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_revokeApiToken_Results_Promise) Struct() (HackSessionContext_revokeApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_revokeApiToken_Results{s}, err
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getIpNetwork_Params{}, err
	}
	return HackSessionContext_getIpNetwork_Params{root.Struct()}, nil
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getIpNetwork_Results{}, err
	}
	return HackSessionContext_getIpNetwork_Results{root.Struct()}, nil
}

func (s HackSessionContext_getIpNetwork_Results) Network() ip.IpNetwork {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return ip.IpNetwork{}
	}
	return ip.IpNetwork{Client: p.Interface().Client()}
}

func (s HackSessionContext_getIpNetwork_Results) HasNetwork() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
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
	return s.Struct.SetPtr(0, in.ToPtr())
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getIpInterface_Params{}, err
	}
	return HackSessionContext_getIpInterface_Params{root.Struct()}, nil
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getIpInterface_Results{}, err
	}
	return HackSessionContext_getIpInterface_Results{root.Struct()}, nil
}

func (s HackSessionContext_getIpInterface_Results) Interface() ip.IpInterface {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return ip.IpInterface{}
	}
	return ip.IpInterface{Client: p.Interface().Client()}
}

func (s HackSessionContext_getIpInterface_Results) HasInterface() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
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
	return s.Struct.SetPtr(0, in.ToPtr())
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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Params{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Params{root.Struct()}, nil
}

func (s HackSessionContext_getUiViewForEndpoint_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s HackSessionContext_getUiViewForEndpoint_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getUiViewForEndpoint_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	root, err := msg.RootPtr()
	if err != nil {
		return HackSessionContext_getUiViewForEndpoint_Results{}, err
	}
	return HackSessionContext_getUiViewForEndpoint_Results{root.Struct()}, nil
}

func (s HackSessionContext_getUiViewForEndpoint_Results) View() grain.UiView {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return grain.UiView{}
	}
	return grain.UiView{Client: p.Interface().Client()}
}

func (s HackSessionContext_getUiViewForEndpoint_Results) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
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
	return s.Struct.SetPtr(0, in.ToPtr())
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

type HackEmailSession_Server interface {
	Send(email.EmailSendPort_send) error
}

func HackEmailSession_ServerToClient(s HackEmailSession_Server) HackEmailSession {
	c, _ := s.(server.Closer)
	return HackEmailSession{Client: server.New(HackEmailSession_Methods(nil, s), c)}
}

func HackEmailSession_Methods(methods []server.Method, s HackEmailSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
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

	return methods
}
