package api_session

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	ip "zenhack.net/go/sandstorm/capnp/sandstorm/ip"
	web_session "zenhack.net/go/sandstorm/capnp/sandstorm/web_session"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type ApiSession struct{ Client capnp.Client }

func (c ApiSession) Get(ctx context.Context, params func(web_session.WebSession_get_Params) error, opts ...capnp.CallOption) web_session.WebSession_Response_Promise {
	if c.Client == nil {
		return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_get_Params{Struct: s}) }
	}
	return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Post(ctx context.Context, params func(web_session.WebSession_post_Params) error, opts ...capnp.CallOption) web_session.WebSession_Response_Promise {
	if c.Client == nil {
		return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_post_Params{Struct: s}) }
	}
	return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) OpenWebSocket(ctx context.Context, params func(web_session.WebSession_openWebSocket_Params) error, opts ...capnp.CallOption) web_session.WebSession_openWebSocket_Results_Promise {
	if c.Client == nil {
		return web_session.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_openWebSocket_Params{Struct: s}) }
	}
	return web_session.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Put(ctx context.Context, params func(web_session.WebSession_put_Params) error, opts ...capnp.CallOption) web_session.WebSession_Response_Promise {
	if c.Client == nil {
		return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_put_Params{Struct: s}) }
	}
	return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) Delete(ctx context.Context, params func(web_session.WebSession_delete_Params) error, opts ...capnp.CallOption) web_session.WebSession_Response_Promise {
	if c.Client == nil {
		return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_delete_Params{Struct: s}) }
	}
	return web_session.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) PostStreaming(ctx context.Context, params func(web_session.WebSession_postStreaming_Params) error, opts ...capnp.CallOption) web_session.WebSession_postStreaming_Results_Promise {
	if c.Client == nil {
		return web_session.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_postStreaming_Params{Struct: s}) }
	}
	return web_session.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ApiSession) PutStreaming(ctx context.Context, params func(web_session.WebSession_putStreaming_Params) error, opts ...capnp.CallOption) web_session.WebSession_putStreaming_Results_Promise {
	if c.Client == nil {
		return web_session.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(web_session.WebSession_putStreaming_Params{Struct: s}) }
	}
	return web_session.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ApiSession_Server interface {
	Get(web_session.WebSession_get) error

	Post(web_session.WebSession_post) error

	OpenWebSocket(web_session.WebSession_openWebSocket) error

	Put(web_session.WebSession_put) error

	Delete(web_session.WebSession_delete) error

	PostStreaming(web_session.WebSession_postStreaming) error

	PutStreaming(web_session.WebSession_putStreaming) error
}

func ApiSession_ServerToClient(s ApiSession_Server) ApiSession {
	c, _ := s.(server.Closer)
	return ApiSession{Client: server.New(ApiSession_Methods(nil, s), c)}
}

func ApiSession_Methods(methods []server.Method, s ApiSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 7)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := web_session.WebSession_get{c, opts, web_session.WebSession_get_Params{Struct: p}, web_session.WebSession_Response{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 7},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := web_session.WebSession_post{c, opts, web_session.WebSession_post_Params{Struct: p}, web_session.WebSession_Response{Struct: r}}
			return s.Post(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 7},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := web_session.WebSession_openWebSocket{c, opts, web_session.WebSession_openWebSocket_Params{Struct: p}, web_session.WebSession_openWebSocket_Results{Struct: r}}
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
			call := web_session.WebSession_put{c, opts, web_session.WebSession_put_Params{Struct: p}, web_session.WebSession_Response{Struct: r}}
			return s.Put(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 7},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := web_session.WebSession_delete{c, opts, web_session.WebSession_delete_Params{Struct: p}, web_session.WebSession_Response{Struct: r}}
			return s.Delete(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 7},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := web_session.WebSession_postStreaming{c, opts, web_session.WebSession_postStreaming_Params{Struct: p}, web_session.WebSession_postStreaming_Results{Struct: r}}
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
			call := web_session.WebSession_putStreaming{c, opts, web_session.WebSession_putStreaming_Params{Struct: p}, web_session.WebSession_putStreaming_Results{Struct: r}}
			return s.PutStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
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
	root, err := msg.RootPtr()
	if err != nil {
		return ApiSession_Params{}, err
	}
	return ApiSession_Params{root.Struct()}, nil
}

func (s ApiSession_Params) RemoteAddress() (ip.IpAddress, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return ip.IpAddress{}, err
	}

	return ip.IpAddress{Struct: p.Struct()}, nil

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
