package apisession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	ip "zenhack.net/go/sandstorm/capnp/ip"
	websession "zenhack.net/go/sandstorm/capnp/websession"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type ApiSession struct{ Client capnp.Client }

func (c ApiSession) Get(ctx context.Context, params func(websession.WebSession_get_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_get_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Post(ctx context.Context, params func(websession.WebSession_post_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_post_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) OpenWebSocket(ctx context.Context, params func(websession.WebSession_openWebSocket_Params) error, opts ...capnp.CallOption) websession.WebSession_openWebSocket_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_openWebSocket_Params{Struct: s}) }
	}
	return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Put(ctx context.Context, params func(websession.WebSession_put_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_put_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Delete(ctx context.Context, params func(websession.WebSession_delete_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_delete_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) PostStreaming(ctx context.Context, params func(websession.WebSession_postStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_postStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_postStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) PutStreaming(ctx context.Context, params func(websession.WebSession_putStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_putStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_putStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Propfind(ctx context.Context, params func(websession.WebSession_propfind_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_propfind_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Proppatch(ctx context.Context, params func(websession.WebSession_proppatch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_proppatch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Mkcol(ctx context.Context, params func(websession.WebSession_mkcol_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_mkcol_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Copy(ctx context.Context, params func(websession.WebSession_copy_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_copy_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Move(ctx context.Context, params func(websession.WebSession_move_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_move_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Lock(ctx context.Context, params func(websession.WebSession_lock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_lock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Unlock(ctx context.Context, params func(websession.WebSession_unlock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_unlock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Acl(ctx context.Context, params func(websession.WebSession_acl_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_acl_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Report(ctx context.Context, params func(websession.WebSession_report_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_report_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Options(ctx context.Context, params func(websession.WebSession_options_Params) error, opts ...capnp.CallOption) websession.WebSession_Options_Promise {
	if c.Client == nil {
		return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_options_Params{Struct: s}) }
	}
	return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Patch(ctx context.Context, params func(websession.WebSession_patch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_patch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ApiSession_Server interface {
	Get(websession.WebSession_get) error

	Post(websession.WebSession_post) error

	OpenWebSocket(websession.WebSession_openWebSocket) error

	Put(websession.WebSession_put) error

	Delete(websession.WebSession_delete) error

	PostStreaming(websession.WebSession_postStreaming) error

	PutStreaming(websession.WebSession_putStreaming) error

	Propfind(websession.WebSession_propfind) error

	Proppatch(websession.WebSession_proppatch) error

	Mkcol(websession.WebSession_mkcol) error

	Copy(websession.WebSession_copy) error

	Move(websession.WebSession_move) error

	Lock(websession.WebSession_lock) error

	Unlock(websession.WebSession_unlock) error

	Acl(websession.WebSession_acl) error

	Report(websession.WebSession_report) error

	Options(websession.WebSession_options) error

	Patch(websession.WebSession_patch) error
}

func ApiSession_ServerToClient(s ApiSession_Server) ApiSession {
	c, _ := s.(server.Closer)
	return ApiSession{Client: server.New(ApiSession_Methods(nil, s), c)}
}

func ApiSession_Methods(methods []server.Method, s ApiSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 18)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_get{c, opts, websession.WebSession_get_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_post{c, opts, websession.WebSession_post_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Post(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_openWebSocket{c, opts, websession.WebSession_openWebSocket_Params{Struct: p}, websession.WebSession_openWebSocket_Results{Struct: r}}
			return s.OpenWebSocket(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_put{c, opts, websession.WebSession_put_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Put(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_delete{c, opts, websession.WebSession_delete_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Delete(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_postStreaming{c, opts, websession.WebSession_postStreaming_Params{Struct: p}, websession.WebSession_postStreaming_Results{Struct: r}}
			return s.PostStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_putStreaming{c, opts, websession.WebSession_putStreaming_Params{Struct: p}, websession.WebSession_putStreaming_Results{Struct: r}}
			return s.PutStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_propfind{c, opts, websession.WebSession_propfind_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Propfind(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_proppatch{c, opts, websession.WebSession_proppatch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Proppatch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_mkcol{c, opts, websession.WebSession_mkcol_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Mkcol(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_copy{c, opts, websession.WebSession_copy_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Copy(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_move{c, opts, websession.WebSession_move_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Move(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_lock{c, opts, websession.WebSession_lock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Lock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_unlock{c, opts, websession.WebSession_unlock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Unlock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_acl{c, opts, websession.WebSession_acl_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Acl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_report{c, opts, websession.WebSession_report_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Report(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_options{c, opts, websession.WebSession_options_Params{Struct: p}, websession.WebSession_Options{Struct: r}}
			return s.Options(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_patch{c, opts, websession.WebSession_patch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Patch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 8},
	})

	return methods
}

type ApiSession_Params struct{ capnp.Struct }

func NewApiSession_Params(s *capnp.Segment) (ApiSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ApiSession_Params{}, err
	}
	return ApiSession_Params{st}, nil
}

func NewRootApiSession_Params(s *capnp.Segment) (ApiSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ApiSession_Params{}, err
	}
	return ApiSession_Params{st}, nil
}

func ReadRootApiSession_Params(msg *capnp.Message) (ApiSession_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return ApiSession_Params{}, err
	}
	st := capnp.ToStruct(root)
	return ApiSession_Params{st}, nil
}

func (s ApiSession_Params) RemoteAddress() (ip.IpAddress, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return ip.IpAddress{}, err
	}

	ss := capnp.ToStruct(p)

	return ip.IpAddress{Struct: ss}, nil
}

func (s ApiSession_Params) SetRemoteAddress(v ip.IpAddress) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewRemoteAddress sets the remoteAddress field to a newly
// allocated ip.IpAddress struct, preferring placement in s's segment.
func (s ApiSession_Params) NewRemoteAddress() (ip.IpAddress, error) {

	ss, err := ip.NewIpAddress(s.Struct.Segment())
	if err != nil {
		return ip.IpAddress{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// ApiSession_Params_List is a list of ApiSession_Params.
type ApiSession_Params_List struct{ capnp.List }

// NewApiSession_Params creates a new list of ApiSession_Params.
func NewApiSession_Params_List(s *capnp.Segment, sz int32) (ApiSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return ApiSession_Params_List{}, err
	}
	return ApiSession_Params_List{l}, nil
}

func (s ApiSession_Params_List) At(i int) ApiSession_Params {
	return ApiSession_Params{s.List.Struct(i)}
}
func (s ApiSession_Params_List) Set(i int, v ApiSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ApiSession_Params_Promise is a wrapper for a ApiSession_Params promised by a client call.
type ApiSession_Params_Promise struct{ *capnp.Pipeline }

func (p ApiSession_Params_Promise) Struct() (ApiSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return ApiSession_Params{s}, err
}

func (p ApiSession_Params_Promise) RemoteAddress() ip.IpAddress_Promise {
	return ip.IpAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type ApiSession_PowerboxTag struct{ capnp.Struct }

func NewApiSession_PowerboxTag(s *capnp.Segment) (ApiSession_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return ApiSession_PowerboxTag{}, err
	}
	return ApiSession_PowerboxTag{st}, nil
}

func NewRootApiSession_PowerboxTag(s *capnp.Segment) (ApiSession_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return ApiSession_PowerboxTag{}, err
	}
	return ApiSession_PowerboxTag{st}, nil
}

func ReadRootApiSession_PowerboxTag(msg *capnp.Message) (ApiSession_PowerboxTag, error) {
	root, err := msg.Root()
	if err != nil {
		return ApiSession_PowerboxTag{}, err
	}
	st := capnp.ToStruct(root)
	return ApiSession_PowerboxTag{st}, nil
}

func (s ApiSession_PowerboxTag) CanonicalUrl() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s ApiSession_PowerboxTag) SetCanonicalUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s ApiSession_PowerboxTag) OauthScopes() (ApiSession_PowerboxTag_OAuthScope_List, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope_List{}, err
	}

	l := capnp.ToList(p)

	return ApiSession_PowerboxTag_OAuthScope_List{List: l}, nil
}

func (s ApiSession_PowerboxTag) SetOauthScopes(v ApiSession_PowerboxTag_OAuthScope_List) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s ApiSession_PowerboxTag) Authentication() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s ApiSession_PowerboxTag) SetAuthentication(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

// ApiSession_PowerboxTag_List is a list of ApiSession_PowerboxTag.
type ApiSession_PowerboxTag_List struct{ capnp.List }

// NewApiSession_PowerboxTag creates a new list of ApiSession_PowerboxTag.
func NewApiSession_PowerboxTag_List(s *capnp.Segment, sz int32) (ApiSession_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return ApiSession_PowerboxTag_List{}, err
	}
	return ApiSession_PowerboxTag_List{l}, nil
}

func (s ApiSession_PowerboxTag_List) At(i int) ApiSession_PowerboxTag {
	return ApiSession_PowerboxTag{s.List.Struct(i)}
}
func (s ApiSession_PowerboxTag_List) Set(i int, v ApiSession_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// ApiSession_PowerboxTag_Promise is a wrapper for a ApiSession_PowerboxTag promised by a client call.
type ApiSession_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p ApiSession_PowerboxTag_Promise) Struct() (ApiSession_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return ApiSession_PowerboxTag{s}, err
}

type ApiSession_PowerboxTag_OAuthScope struct{ capnp.Struct }

func NewApiSession_PowerboxTag_OAuthScope(s *capnp.Segment) (ApiSession_PowerboxTag_OAuthScope, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope{}, err
	}
	return ApiSession_PowerboxTag_OAuthScope{st}, nil
}

func NewRootApiSession_PowerboxTag_OAuthScope(s *capnp.Segment) (ApiSession_PowerboxTag_OAuthScope, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope{}, err
	}
	return ApiSession_PowerboxTag_OAuthScope{st}, nil
}

func ReadRootApiSession_PowerboxTag_OAuthScope(msg *capnp.Message) (ApiSession_PowerboxTag_OAuthScope, error) {
	root, err := msg.Root()
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope{}, err
	}
	st := capnp.ToStruct(root)
	return ApiSession_PowerboxTag_OAuthScope{st}, nil
}

func (s ApiSession_PowerboxTag_OAuthScope) Name() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s ApiSession_PowerboxTag_OAuthScope) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// ApiSession_PowerboxTag_OAuthScope_List is a list of ApiSession_PowerboxTag_OAuthScope.
type ApiSession_PowerboxTag_OAuthScope_List struct{ capnp.List }

// NewApiSession_PowerboxTag_OAuthScope creates a new list of ApiSession_PowerboxTag_OAuthScope.
func NewApiSession_PowerboxTag_OAuthScope_List(s *capnp.Segment, sz int32) (ApiSession_PowerboxTag_OAuthScope_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope_List{}, err
	}
	return ApiSession_PowerboxTag_OAuthScope_List{l}, nil
}

func (s ApiSession_PowerboxTag_OAuthScope_List) At(i int) ApiSession_PowerboxTag_OAuthScope {
	return ApiSession_PowerboxTag_OAuthScope{s.List.Struct(i)}
}
func (s ApiSession_PowerboxTag_OAuthScope_List) Set(i int, v ApiSession_PowerboxTag_OAuthScope) error {
	return s.List.SetStruct(i, v.Struct)
}

// ApiSession_PowerboxTag_OAuthScope_Promise is a wrapper for a ApiSession_PowerboxTag_OAuthScope promised by a client call.
type ApiSession_PowerboxTag_OAuthScope_Promise struct{ *capnp.Pipeline }

func (p ApiSession_PowerboxTag_OAuthScope_Promise) Struct() (ApiSession_PowerboxTag_OAuthScope, error) {
	s, err := p.Pipeline.Struct()
	return ApiSession_PowerboxTag_OAuthScope{s}, err
}
