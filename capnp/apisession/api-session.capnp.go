package apisession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	ip "zenhack.net/go/sandstorm/capnp/ip"
	websession "zenhack.net/go/sandstorm/capnp/websession"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	return methods
}

type ApiSession_Params struct{ capnp.Struct }

func NewApiSession_Params(s *capnp.Segment) (ApiSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ApiSession_Params{st}, err
}

func NewRootApiSession_Params(s *capnp.Segment) (ApiSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ApiSession_Params{st}, err
}

func ReadRootApiSession_Params(msg *capnp.Message) (ApiSession_Params, error) {
	root, err := msg.RootPtr()
	return ApiSession_Params{root.Struct()}, err
}

func (s ApiSession_Params) String() string {
	str, _ := text.Marshal(0xfee82d8d4c4ff597, s.Struct)
	return str
}

func (s ApiSession_Params) RemoteAddress() (ip.IpAddress, error) {
	p, err := s.Struct.Ptr(0)
	return ip.IpAddress{Struct: p.Struct()}, err
}

func (s ApiSession_Params) HasRemoteAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiSession_Params) SetRemoteAddress(v ip.IpAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRemoteAddress sets the remoteAddress field to a newly
// allocated ip.IpAddress struct, preferring placement in s's segment.
func (s ApiSession_Params) NewRemoteAddress() (ip.IpAddress, error) {
	ss, err := ip.NewIpAddress(s.Struct.Segment())
	if err != nil {
		return ip.IpAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// ApiSession_Params_List is a list of ApiSession_Params.
type ApiSession_Params_List struct{ capnp.List }

// NewApiSession_Params creates a new list of ApiSession_Params.
func NewApiSession_Params_List(s *capnp.Segment, sz int32) (ApiSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ApiSession_Params_List{l}, err
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
	return ApiSession_PowerboxTag{st}, err
}

func NewRootApiSession_PowerboxTag(s *capnp.Segment) (ApiSession_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ApiSession_PowerboxTag{st}, err
}

func ReadRootApiSession_PowerboxTag(msg *capnp.Message) (ApiSession_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return ApiSession_PowerboxTag{root.Struct()}, err
}

func (s ApiSession_PowerboxTag) String() string {
	str, _ := text.Marshal(0x93e1f7ca567dee32, s.Struct)
	return str
}

func (s ApiSession_PowerboxTag) CanonicalUrl() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiSession_PowerboxTag) HasCanonicalUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiSession_PowerboxTag) CanonicalUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiSession_PowerboxTag) SetCanonicalUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiSession_PowerboxTag) OauthScopes() (ApiSession_PowerboxTag_OAuthScope_List, error) {
	p, err := s.Struct.Ptr(1)
	return ApiSession_PowerboxTag_OAuthScope_List{List: p.List()}, err
}

func (s ApiSession_PowerboxTag) HasOauthScopes() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ApiSession_PowerboxTag) SetOauthScopes(v ApiSession_PowerboxTag_OAuthScope_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewOauthScopes sets the oauthScopes field to a newly
// allocated ApiSession_PowerboxTag_OAuthScope_List, preferring placement in s's segment.
func (s ApiSession_PowerboxTag) NewOauthScopes(n int32) (ApiSession_PowerboxTag_OAuthScope_List, error) {
	l, err := NewApiSession_PowerboxTag_OAuthScope_List(s.Struct.Segment(), n)
	if err != nil {
		return ApiSession_PowerboxTag_OAuthScope_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s ApiSession_PowerboxTag) Authentication() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s ApiSession_PowerboxTag) HasAuthentication() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ApiSession_PowerboxTag) AuthenticationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s ApiSession_PowerboxTag) SetAuthentication(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// ApiSession_PowerboxTag_List is a list of ApiSession_PowerboxTag.
type ApiSession_PowerboxTag_List struct{ capnp.List }

// NewApiSession_PowerboxTag creates a new list of ApiSession_PowerboxTag.
func NewApiSession_PowerboxTag_List(s *capnp.Segment, sz int32) (ApiSession_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return ApiSession_PowerboxTag_List{l}, err
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
	return ApiSession_PowerboxTag_OAuthScope{st}, err
}

func NewRootApiSession_PowerboxTag_OAuthScope(s *capnp.Segment) (ApiSession_PowerboxTag_OAuthScope, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ApiSession_PowerboxTag_OAuthScope{st}, err
}

func ReadRootApiSession_PowerboxTag_OAuthScope(msg *capnp.Message) (ApiSession_PowerboxTag_OAuthScope, error) {
	root, err := msg.RootPtr()
	return ApiSession_PowerboxTag_OAuthScope{root.Struct()}, err
}

func (s ApiSession_PowerboxTag_OAuthScope) String() string {
	str, _ := text.Marshal(0xbb95f31093b1fd88, s.Struct)
	return str
}

func (s ApiSession_PowerboxTag_OAuthScope) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiSession_PowerboxTag_OAuthScope) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiSession_PowerboxTag_OAuthScope) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiSession_PowerboxTag_OAuthScope) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// ApiSession_PowerboxTag_OAuthScope_List is a list of ApiSession_PowerboxTag_OAuthScope.
type ApiSession_PowerboxTag_OAuthScope_List struct{ capnp.List }

// NewApiSession_PowerboxTag_OAuthScope creates a new list of ApiSession_PowerboxTag_OAuthScope.
func NewApiSession_PowerboxTag_OAuthScope_List(s *capnp.Segment, sz int32) (ApiSession_PowerboxTag_OAuthScope_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ApiSession_PowerboxTag_OAuthScope_List{l}, err
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

const schema_eb014c0c3413cbfb = "x\xda\x84\x92\xcfk\x13A\x14\xc7\xdf\x9b\xd9u+&" +
	"\x9b\xac\x1bP\x10\x11\x04\x11\xc5\x16\x7f\"\xf4\x92\x06\xec" +
	"ER\xbaK\xd4\x83\xb7\xc9f\xb0+\xcd\xee\xb0\xbb\xa2" +
	"\x15z\xd6\x8b\x88\x94\"\x1e\x15\x04=I\x8f\"=\xfa" +
	"\x1b\x05\x11\xff\x00\xf1\"\x0a\xbd\x14\xf1 :\xbe\x09M" +
	"\x13\xa2P\xc2@\xf8\xf0\xdew>\xef\xcd\x1e]\xe7S" +
	"\xec\x98}c\x0f@k\x0a\xedm\xfa\xf8\xda\xe2\x857" +
	"??/\x81\xb7\x17\xf5\xcbw\x07^,|Yx\x05" +
	"6w\x00N<.\x9fE\x7f\xb5\xbc\x0b\xc0\x7f]~" +
	"\x02\xa8o\xfe^Y\xaa\xae/?\x03\xef\x10\x0e:m" +
	"4\xc5\x8b\xee\x03\xf4\x97]\xfa\xeb\xdfq\xeb0\x14\xe6" +
	"\xed\xe4\xfa\xd7[\xffd\xa9\x89\xdf\x01\xd0_q\xbf\xf9" +
	"\xab\xeeA*\xfc\xe4:t(_\xdf\xfd1\xdb\xbc5" +
	"\xfe\xf5\xcf\x88F/\xf9\xb9\xbb\x1fM)5|\xa0d" +
	"\xdd\xfbMk\xa1\xe2\xf1\\\xe6\xb9\x15\xa7\xc9D$T" +
	"\xa2&\x1b*n\x111 H\xaf\xca\xac\x9d^;'" +
	".\x01\x84\x16\x0e\xcb\xe3E=\xdb\xb8R\xcc\xb5\xa2\x14" +
	"\xb8\x92a\x89[\x00\x16\x02x\xd3\x97\xa9\xf8\x0c\xc70" +
	"`\xe8!\xd6\xd0\xc0\x996\xc1&\xc19\x82\x8c\xd5\x90" +
	"\x11\x94\xd7\x09v\x08*\x86:\x12I\x9a\xc4\x91\x80\xca" +
	"\xfc\xf9l\x1eK\xc0\xe8\xa0N\xc5\xc6%\x8e\x929\xba" +
	"\x80\x01G\xac\x0eD\x00\x0d\xd4\xa6J&E\x0c\xf5H" +
	"\x14\xe4\xbe\xd9\xdf\x9f\xd0\xdej\xc2\x89\x8dq\x9cT\xc9" +
	"\x001\xb46\x07*\x1f&\xcd1\xd2\xac1\xac$\xa2" +
	"+\xff\x09g\xa3\xe1\x15\x93\x1e\x8e\xd1\xc6\x06\x8f\xb2}" +
	"r\xe8C\xb1\xdb\xf5@d\xa2\x9b\xeb\xbe\x028$\x11" +
	"p\xbb\xb7\xe7\xf7\x8fN\xcd\xdc\xf7\x9c\x87@\xcf\xda\xbf" +
	"\x84\xff\x7f\x02\x919\x143b\x9c\x91q\x89\x8cw\xd3" +
	"b3\xd9M\x0b\xd9\xe8\xc0\xbeNFM\xb4\xbc\x1dG" +
	">\x9e\xbew\xfb\xe9\x9aY^\x15\xf0o\x00\x00\x00\xff" +
	"\xff`g\xdcE"

func init() {
	schemas.Register(schema_eb014c0c3413cbfb,
		0x93e1f7ca567dee32,
		0xbb95f31093b1fd88,
		0xc879e379c625cdc7,
		0xfee82d8d4c4ff597)
}
