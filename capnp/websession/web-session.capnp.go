package websession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	math "math"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

const HttpStatus = uint64(0xaf480a0c6cab8887)

// Constants defined in web-session.capnp.
const (
	HttpStatusAnnotationId = uint64(12630356203439622279)
)

// Constants defined in web-session.capnp.
var (
	WebSession_Context_headerWhitelist  = capnp.TextList{List: capnp.MustUnmarshalRootPtr(x_a8cb0f2f1a756b32[0:192]).List()}
	WebSession_Response_headerWhitelist = capnp.TextList{List: capnp.MustUnmarshalRootPtr(x_a8cb0f2f1a756b32[192:264]).List()}
)

type HttpStatusDescriptor struct{ capnp.Struct }

func NewHttpStatusDescriptor(s *capnp.Segment) (HttpStatusDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return HttpStatusDescriptor{st}, err
}

func NewRootHttpStatusDescriptor(s *capnp.Segment) (HttpStatusDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return HttpStatusDescriptor{st}, err
}

func ReadRootHttpStatusDescriptor(msg *capnp.Message) (HttpStatusDescriptor, error) {
	root, err := msg.RootPtr()
	return HttpStatusDescriptor{root.Struct()}, err
}

func (s HttpStatusDescriptor) String() string {
	str, _ := text.Marshal(0xbc353583a3731ade, s.Struct)
	return str
}

func (s HttpStatusDescriptor) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s HttpStatusDescriptor) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s HttpStatusDescriptor) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HttpStatusDescriptor) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HttpStatusDescriptor) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HttpStatusDescriptor) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HttpStatusDescriptor_List is a list of HttpStatusDescriptor.
type HttpStatusDescriptor_List struct{ capnp.List }

// NewHttpStatusDescriptor creates a new list of HttpStatusDescriptor.
func NewHttpStatusDescriptor_List(s *capnp.Segment, sz int32) (HttpStatusDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return HttpStatusDescriptor_List{l}, err
}

func (s HttpStatusDescriptor_List) At(i int) HttpStatusDescriptor {
	return HttpStatusDescriptor{s.List.Struct(i)}
}

func (s HttpStatusDescriptor_List) Set(i int, v HttpStatusDescriptor) error {
	return s.List.SetStruct(i, v.Struct)
}

// HttpStatusDescriptor_Promise is a wrapper for a HttpStatusDescriptor promised by a client call.
type HttpStatusDescriptor_Promise struct{ *capnp.Pipeline }

func (p HttpStatusDescriptor_Promise) Struct() (HttpStatusDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return HttpStatusDescriptor{s}, err
}

type WebSession struct{ Client capnp.Client }

func (c WebSession) Get(ctx context.Context, params func(WebSession_get_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_get_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Post(ctx context.Context, params func(WebSession_post_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_post_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) OpenWebSocket(ctx context.Context, params func(WebSession_openWebSocket_Params) error, opts ...capnp.CallOption) WebSession_openWebSocket_Results_Promise {
	if c.Client == nil {
		return WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_openWebSocket_Params{Struct: s}) }
	}
	return WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Put(ctx context.Context, params func(WebSession_put_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_put_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Delete(ctx context.Context, params func(WebSession_delete_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_delete_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) PostStreaming(ctx context.Context, params func(WebSession_postStreaming_Params) error, opts ...capnp.CallOption) WebSession_postStreaming_Results_Promise {
	if c.Client == nil {
		return WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_postStreaming_Params{Struct: s}) }
	}
	return WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) PutStreaming(ctx context.Context, params func(WebSession_putStreaming_Params) error, opts ...capnp.CallOption) WebSession_putStreaming_Results_Promise {
	if c.Client == nil {
		return WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_putStreaming_Params{Struct: s}) }
	}
	return WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Propfind(ctx context.Context, params func(WebSession_propfind_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_propfind_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Proppatch(ctx context.Context, params func(WebSession_proppatch_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_proppatch_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Mkcol(ctx context.Context, params func(WebSession_mkcol_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_mkcol_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Copy(ctx context.Context, params func(WebSession_copy_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_copy_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Move(ctx context.Context, params func(WebSession_move_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_move_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Lock(ctx context.Context, params func(WebSession_lock_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_lock_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Unlock(ctx context.Context, params func(WebSession_unlock_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_unlock_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Acl(ctx context.Context, params func(WebSession_acl_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_acl_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Report(ctx context.Context, params func(WebSession_report_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_report_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Options(ctx context.Context, params func(WebSession_options_Params) error, opts ...capnp.CallOption) WebSession_Options_Promise {
	if c.Client == nil {
		return WebSession_Options_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_options_Params{Struct: s}) }
	}
	return WebSession_Options_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession) Patch(ctx context.Context, params func(WebSession_patch_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_patch_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type WebSession_Server interface {
	Get(WebSession_get) error

	Post(WebSession_post) error

	OpenWebSocket(WebSession_openWebSocket) error

	Put(WebSession_put) error

	Delete(WebSession_delete) error

	PostStreaming(WebSession_postStreaming) error

	PutStreaming(WebSession_putStreaming) error

	Propfind(WebSession_propfind) error

	Proppatch(WebSession_proppatch) error

	Mkcol(WebSession_mkcol) error

	Copy(WebSession_copy) error

	Move(WebSession_move) error

	Lock(WebSession_lock) error

	Unlock(WebSession_unlock) error

	Acl(WebSession_acl) error

	Report(WebSession_report) error

	Options(WebSession_options) error

	Patch(WebSession_patch) error
}

func WebSession_ServerToClient(s WebSession_Server) WebSession {
	c, _ := s.(server.Closer)
	return WebSession{Client: server.New(WebSession_Methods(nil, s), c)}
}

func WebSession_Methods(methods []server.Method, s WebSession_Server) []server.Method {
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
			call := WebSession_get{c, opts, WebSession_get_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_post{c, opts, WebSession_post_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_openWebSocket{c, opts, WebSession_openWebSocket_Params{Struct: p}, WebSession_openWebSocket_Results{Struct: r}}
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
			call := WebSession_put{c, opts, WebSession_put_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_delete{c, opts, WebSession_delete_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_postStreaming{c, opts, WebSession_postStreaming_Params{Struct: p}, WebSession_postStreaming_Results{Struct: r}}
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
			call := WebSession_putStreaming{c, opts, WebSession_putStreaming_Params{Struct: p}, WebSession_putStreaming_Results{Struct: r}}
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
			call := WebSession_propfind{c, opts, WebSession_propfind_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_proppatch{c, opts, WebSession_proppatch_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_mkcol{c, opts, WebSession_mkcol_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_copy{c, opts, WebSession_copy_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_move{c, opts, WebSession_move_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_lock{c, opts, WebSession_lock_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_unlock{c, opts, WebSession_unlock_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_acl{c, opts, WebSession_acl_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_report{c, opts, WebSession_report_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_options{c, opts, WebSession_options_Params{Struct: p}, WebSession_Options{Struct: r}}
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
			call := WebSession_patch{c, opts, WebSession_patch_Params{Struct: p}, WebSession_Response{Struct: r}}
			return s.Patch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	return methods
}

// WebSession_get holds the arguments for a server call to WebSession.get.
type WebSession_get struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_get_Params
	Results WebSession_Response
}

// WebSession_post holds the arguments for a server call to WebSession.post.
type WebSession_post struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_post_Params
	Results WebSession_Response
}

// WebSession_openWebSocket holds the arguments for a server call to WebSession.openWebSocket.
type WebSession_openWebSocket struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_openWebSocket_Params
	Results WebSession_openWebSocket_Results
}

// WebSession_put holds the arguments for a server call to WebSession.put.
type WebSession_put struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_put_Params
	Results WebSession_Response
}

// WebSession_delete holds the arguments for a server call to WebSession.delete.
type WebSession_delete struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_delete_Params
	Results WebSession_Response
}

// WebSession_postStreaming holds the arguments for a server call to WebSession.postStreaming.
type WebSession_postStreaming struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_postStreaming_Params
	Results WebSession_postStreaming_Results
}

// WebSession_putStreaming holds the arguments for a server call to WebSession.putStreaming.
type WebSession_putStreaming struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_putStreaming_Params
	Results WebSession_putStreaming_Results
}

// WebSession_propfind holds the arguments for a server call to WebSession.propfind.
type WebSession_propfind struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_propfind_Params
	Results WebSession_Response
}

// WebSession_proppatch holds the arguments for a server call to WebSession.proppatch.
type WebSession_proppatch struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_proppatch_Params
	Results WebSession_Response
}

// WebSession_mkcol holds the arguments for a server call to WebSession.mkcol.
type WebSession_mkcol struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_mkcol_Params
	Results WebSession_Response
}

// WebSession_copy holds the arguments for a server call to WebSession.copy.
type WebSession_copy struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_copy_Params
	Results WebSession_Response
}

// WebSession_move holds the arguments for a server call to WebSession.move.
type WebSession_move struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_move_Params
	Results WebSession_Response
}

// WebSession_lock holds the arguments for a server call to WebSession.lock.
type WebSession_lock struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_lock_Params
	Results WebSession_Response
}

// WebSession_unlock holds the arguments for a server call to WebSession.unlock.
type WebSession_unlock struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_unlock_Params
	Results WebSession_Response
}

// WebSession_acl holds the arguments for a server call to WebSession.acl.
type WebSession_acl struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_acl_Params
	Results WebSession_Response
}

// WebSession_report holds the arguments for a server call to WebSession.report.
type WebSession_report struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_report_Params
	Results WebSession_Response
}

// WebSession_options holds the arguments for a server call to WebSession.options.
type WebSession_options struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_options_Params
	Results WebSession_Options
}

// WebSession_patch holds the arguments for a server call to WebSession.patch.
type WebSession_patch struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_patch_Params
	Results WebSession_Response
}

type WebSession_Params struct{ capnp.Struct }

func NewWebSession_Params(s *capnp.Segment) (WebSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_Params{st}, err
}

func NewRootWebSession_Params(s *capnp.Segment) (WebSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_Params{st}, err
}

func ReadRootWebSession_Params(msg *capnp.Message) (WebSession_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_Params{root.Struct()}, err
}

func (s WebSession_Params) String() string {
	str, _ := text.Marshal(0xd7051b9757f6b096, s.Struct)
	return str
}

func (s WebSession_Params) BasePath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_Params) HasBasePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Params) BasePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_Params) SetBasePath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_Params) UserAgent() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Params) HasUserAgent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Params) UserAgentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Params) SetUserAgent(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Params) AcceptableLanguages() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSession_Params) HasAcceptableLanguages() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Params) SetAcceptableLanguages(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewAcceptableLanguages sets the acceptableLanguages field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSession_Params) NewAcceptableLanguages(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// WebSession_Params_List is a list of WebSession_Params.
type WebSession_Params_List struct{ capnp.List }

// NewWebSession_Params creates a new list of WebSession_Params.
func NewWebSession_Params_List(s *capnp.Segment, sz int32) (WebSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_Params_List{l}, err
}

func (s WebSession_Params_List) At(i int) WebSession_Params {
	return WebSession_Params{s.List.Struct(i)}
}

func (s WebSession_Params_List) Set(i int, v WebSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Params_Promise is a wrapper for a WebSession_Params promised by a client call.
type WebSession_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_Params_Promise) Struct() (WebSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Params{s}, err
}

type WebSession_Context struct{ capnp.Struct }
type WebSession_Context_eTagPrecondition WebSession_Context
type WebSession_Context_eTagPrecondition_Which uint16

const (
	WebSession_Context_eTagPrecondition_Which_none          WebSession_Context_eTagPrecondition_Which = 0
	WebSession_Context_eTagPrecondition_Which_exists        WebSession_Context_eTagPrecondition_Which = 1
	WebSession_Context_eTagPrecondition_Which_doesntExist   WebSession_Context_eTagPrecondition_Which = 4
	WebSession_Context_eTagPrecondition_Which_matchesOneOf  WebSession_Context_eTagPrecondition_Which = 2
	WebSession_Context_eTagPrecondition_Which_matchesNoneOf WebSession_Context_eTagPrecondition_Which = 3
)

func (w WebSession_Context_eTagPrecondition_Which) String() string {
	const s = "noneexistsdoesntExistmatchesOneOfmatchesNoneOf"
	switch w {
	case WebSession_Context_eTagPrecondition_Which_none:
		return s[0:4]
	case WebSession_Context_eTagPrecondition_Which_exists:
		return s[4:10]
	case WebSession_Context_eTagPrecondition_Which_doesntExist:
		return s[10:21]
	case WebSession_Context_eTagPrecondition_Which_matchesOneOf:
		return s[21:33]
	case WebSession_Context_eTagPrecondition_Which_matchesNoneOf:
		return s[33:46]

	}
	return "WebSession_Context_eTagPrecondition_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSession_Context{st}, err
}

func NewRootWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSession_Context{st}, err
}

func ReadRootWebSession_Context(msg *capnp.Message) (WebSession_Context, error) {
	root, err := msg.RootPtr()
	return WebSession_Context{root.Struct()}, err
}

func (s WebSession_Context) String() string {
	str, _ := text.Marshal(0xf5cae52becabc767, s.Struct)
	return str
}

func (s WebSession_Context) Cookies() (util.KeyValue_List, error) {
	p, err := s.Struct.Ptr(0)
	return util.KeyValue_List{List: p.List()}, err
}

func (s WebSession_Context) HasCookies() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetCookies(v util.KeyValue_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewCookies sets the cookies field to a newly
// allocated util.KeyValue_List, preferring placement in s's segment.
func (s WebSession_Context) NewCookies(n int32) (util.KeyValue_List, error) {
	l, err := util.NewKeyValue_List(s.Struct.Segment(), n)
	if err != nil {
		return util.KeyValue_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s WebSession_Context) ResponseStream() util.ByteStream {
	p, _ := s.Struct.Ptr(1)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s WebSession_Context) HasResponseStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetResponseStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s WebSession_Context) Accept() (WebSession_AcceptedType_List, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_AcceptedType_List{List: p.List()}, err
}

func (s WebSession_Context) HasAccept() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetAccept(v WebSession_AcceptedType_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewAccept sets the accept field to a newly
// allocated WebSession_AcceptedType_List, preferring placement in s's segment.
func (s WebSession_Context) NewAccept(n int32) (WebSession_AcceptedType_List, error) {
	l, err := NewWebSession_AcceptedType_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_AcceptedType_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s WebSession_Context) ETagPrecondition() WebSession_Context_eTagPrecondition {
	return WebSession_Context_eTagPrecondition(s)
}

func (s WebSession_Context_eTagPrecondition) Which() WebSession_Context_eTagPrecondition_Which {
	return WebSession_Context_eTagPrecondition_Which(s.Struct.Uint16(0))
}
func (s WebSession_Context_eTagPrecondition) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s WebSession_Context_eTagPrecondition) SetExists() {
	s.Struct.SetUint16(0, 1)

}

func (s WebSession_Context_eTagPrecondition) SetDoesntExist() {
	s.Struct.SetUint16(0, 4)

}

func (s WebSession_Context_eTagPrecondition) MatchesOneOf() (WebSession_ETag_List, error) {
	p, err := s.Struct.Ptr(4)
	return WebSession_ETag_List{List: p.List()}, err
}

func (s WebSession_Context_eTagPrecondition) HasMatchesOneOf() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSession_Context_eTagPrecondition) SetMatchesOneOf(v WebSession_ETag_List) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewMatchesOneOf sets the matchesOneOf field to a newly
// allocated WebSession_ETag_List, preferring placement in s's segment.
func (s WebSession_Context_eTagPrecondition) NewMatchesOneOf(n int32) (WebSession_ETag_List, error) {
	s.Struct.SetUint16(0, 2)
	l, err := NewWebSession_ETag_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_ETag_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s WebSession_Context_eTagPrecondition) MatchesNoneOf() (WebSession_ETag_List, error) {
	p, err := s.Struct.Ptr(4)
	return WebSession_ETag_List{List: p.List()}, err
}

func (s WebSession_Context_eTagPrecondition) HasMatchesNoneOf() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSession_Context_eTagPrecondition) SetMatchesNoneOf(v WebSession_ETag_List) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewMatchesNoneOf sets the matchesNoneOf field to a newly
// allocated WebSession_ETag_List, preferring placement in s's segment.
func (s WebSession_Context_eTagPrecondition) NewMatchesNoneOf(n int32) (WebSession_ETag_List, error) {
	s.Struct.SetUint16(0, 3)
	l, err := NewWebSession_ETag_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_ETag_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s WebSession_Context) AdditionalHeaders() (WebSession_Context_Header_List, error) {
	p, err := s.Struct.Ptr(3)
	return WebSession_Context_Header_List{List: p.List()}, err
}

func (s WebSession_Context) HasAdditionalHeaders() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetAdditionalHeaders(v WebSession_Context_Header_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewAdditionalHeaders sets the additionalHeaders field to a newly
// allocated WebSession_Context_Header_List, preferring placement in s's segment.
func (s WebSession_Context) NewAdditionalHeaders(n int32) (WebSession_Context_Header_List, error) {
	l, err := NewWebSession_Context_Header_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_Context_Header_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

// WebSession_Context_List is a list of WebSession_Context.
type WebSession_Context_List struct{ capnp.List }

// NewWebSession_Context creates a new list of WebSession_Context.
func NewWebSession_Context_List(s *capnp.Segment, sz int32) (WebSession_Context_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return WebSession_Context_List{l}, err
}

func (s WebSession_Context_List) At(i int) WebSession_Context {
	return WebSession_Context{s.List.Struct(i)}
}

func (s WebSession_Context_List) Set(i int, v WebSession_Context) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Context_Promise is a wrapper for a WebSession_Context promised by a client call.
type WebSession_Context_Promise struct{ *capnp.Pipeline }

func (p WebSession_Context_Promise) Struct() (WebSession_Context, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Context{s}, err
}

func (p WebSession_Context_Promise) ResponseStream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p WebSession_Context_Promise) ETagPrecondition() WebSession_Context_eTagPrecondition_Promise {
	return WebSession_Context_eTagPrecondition_Promise{p.Pipeline}
}

// WebSession_Context_eTagPrecondition_Promise is a wrapper for a WebSession_Context_eTagPrecondition promised by a client call.
type WebSession_Context_eTagPrecondition_Promise struct{ *capnp.Pipeline }

func (p WebSession_Context_eTagPrecondition_Promise) Struct() (WebSession_Context_eTagPrecondition, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Context_eTagPrecondition{s}, err
}

type WebSession_Context_Header struct{ capnp.Struct }

func NewWebSession_Context_Header(s *capnp.Segment) (WebSession_Context_Header, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_Context_Header{st}, err
}

func NewRootWebSession_Context_Header(s *capnp.Segment) (WebSession_Context_Header, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_Context_Header{st}, err
}

func ReadRootWebSession_Context_Header(msg *capnp.Message) (WebSession_Context_Header, error) {
	root, err := msg.RootPtr()
	return WebSession_Context_Header{root.Struct()}, err
}

func (s WebSession_Context_Header) String() string {
	str, _ := text.Marshal(0xb4e5f4cccb748429, s.Struct)
	return str
}

func (s WebSession_Context_Header) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_Context_Header) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Context_Header) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_Context_Header) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_Context_Header) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Context_Header) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Context_Header) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Context_Header) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// WebSession_Context_Header_List is a list of WebSession_Context_Header.
type WebSession_Context_Header_List struct{ capnp.List }

// NewWebSession_Context_Header creates a new list of WebSession_Context_Header.
func NewWebSession_Context_Header_List(s *capnp.Segment, sz int32) (WebSession_Context_Header_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSession_Context_Header_List{l}, err
}

func (s WebSession_Context_Header_List) At(i int) WebSession_Context_Header {
	return WebSession_Context_Header{s.List.Struct(i)}
}

func (s WebSession_Context_Header_List) Set(i int, v WebSession_Context_Header) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Context_Header_Promise is a wrapper for a WebSession_Context_Header promised by a client call.
type WebSession_Context_Header_Promise struct{ *capnp.Pipeline }

func (p WebSession_Context_Header_Promise) Struct() (WebSession_Context_Header, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Context_Header{s}, err
}

type WebSession_PostContent struct{ capnp.Struct }

func NewWebSession_PostContent(s *capnp.Segment) (WebSession_PostContent, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_PostContent{st}, err
}

func NewRootWebSession_PostContent(s *capnp.Segment) (WebSession_PostContent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_PostContent{st}, err
}

func ReadRootWebSession_PostContent(msg *capnp.Message) (WebSession_PostContent, error) {
	root, err := msg.RootPtr()
	return WebSession_PostContent{root.Struct()}, err
}

func (s WebSession_PostContent) String() string {
	str, _ := text.Marshal(0xb7d82eac416ab63e, s.Struct)
	return str
}

func (s WebSession_PostContent) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_PostContent) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_PostContent) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_PostContent) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_PostContent) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s WebSession_PostContent) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_PostContent) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s WebSession_PostContent) Encoding() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSession_PostContent) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_PostContent) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSession_PostContent) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// WebSession_PostContent_List is a list of WebSession_PostContent.
type WebSession_PostContent_List struct{ capnp.List }

// NewWebSession_PostContent creates a new list of WebSession_PostContent.
func NewWebSession_PostContent_List(s *capnp.Segment, sz int32) (WebSession_PostContent_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_PostContent_List{l}, err
}

func (s WebSession_PostContent_List) At(i int) WebSession_PostContent {
	return WebSession_PostContent{s.List.Struct(i)}
}

func (s WebSession_PostContent_List) Set(i int, v WebSession_PostContent) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_PostContent_Promise is a wrapper for a WebSession_PostContent promised by a client call.
type WebSession_PostContent_Promise struct{ *capnp.Pipeline }

func (p WebSession_PostContent_Promise) Struct() (WebSession_PostContent, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_PostContent{s}, err
}

type WebSession_PutContent struct{ capnp.Struct }

func NewWebSession_PutContent(s *capnp.Segment) (WebSession_PutContent, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_PutContent{st}, err
}

func NewRootWebSession_PutContent(s *capnp.Segment) (WebSession_PutContent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_PutContent{st}, err
}

func ReadRootWebSession_PutContent(msg *capnp.Message) (WebSession_PutContent, error) {
	root, err := msg.RootPtr()
	return WebSession_PutContent{root.Struct()}, err
}

func (s WebSession_PutContent) String() string {
	str, _ := text.Marshal(0xd7aff1fe39659132, s.Struct)
	return str
}

func (s WebSession_PutContent) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_PutContent) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_PutContent) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_PutContent) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_PutContent) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s WebSession_PutContent) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_PutContent) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s WebSession_PutContent) Encoding() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSession_PutContent) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_PutContent) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSession_PutContent) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// WebSession_PutContent_List is a list of WebSession_PutContent.
type WebSession_PutContent_List struct{ capnp.List }

// NewWebSession_PutContent creates a new list of WebSession_PutContent.
func NewWebSession_PutContent_List(s *capnp.Segment, sz int32) (WebSession_PutContent_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_PutContent_List{l}, err
}

func (s WebSession_PutContent_List) At(i int) WebSession_PutContent {
	return WebSession_PutContent{s.List.Struct(i)}
}

func (s WebSession_PutContent_List) Set(i int, v WebSession_PutContent) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_PutContent_Promise is a wrapper for a WebSession_PutContent promised by a client call.
type WebSession_PutContent_Promise struct{ *capnp.Pipeline }

func (p WebSession_PutContent_Promise) Struct() (WebSession_PutContent, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_PutContent{s}, err
}

type WebSession_ETag struct{ capnp.Struct }

func NewWebSession_ETag(s *capnp.Segment) (WebSession_ETag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_ETag{st}, err
}

func NewRootWebSession_ETag(s *capnp.Segment) (WebSession_ETag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_ETag{st}, err
}

func ReadRootWebSession_ETag(msg *capnp.Message) (WebSession_ETag, error) {
	root, err := msg.RootPtr()
	return WebSession_ETag{root.Struct()}, err
}

func (s WebSession_ETag) String() string {
	str, _ := text.Marshal(0xd22c0be5b9c16558, s.Struct)
	return str
}

func (s WebSession_ETag) Value() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_ETag) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_ETag) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_ETag) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_ETag) Weak() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_ETag) SetWeak(v bool) {
	s.Struct.SetBit(0, v)
}

// WebSession_ETag_List is a list of WebSession_ETag.
type WebSession_ETag_List struct{ capnp.List }

// NewWebSession_ETag creates a new list of WebSession_ETag.
func NewWebSession_ETag_List(s *capnp.Segment, sz int32) (WebSession_ETag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return WebSession_ETag_List{l}, err
}

func (s WebSession_ETag_List) At(i int) WebSession_ETag { return WebSession_ETag{s.List.Struct(i)} }

func (s WebSession_ETag_List) Set(i int, v WebSession_ETag) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_ETag_Promise is a wrapper for a WebSession_ETag promised by a client call.
type WebSession_ETag_Promise struct{ *capnp.Pipeline }

func (p WebSession_ETag_Promise) Struct() (WebSession_ETag, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_ETag{s}, err
}

type WebSession_Cookie struct{ capnp.Struct }
type WebSession_Cookie_expires WebSession_Cookie
type WebSession_Cookie_expires_Which uint16

const (
	WebSession_Cookie_expires_Which_none     WebSession_Cookie_expires_Which = 0
	WebSession_Cookie_expires_Which_absolute WebSession_Cookie_expires_Which = 1
	WebSession_Cookie_expires_Which_relative WebSession_Cookie_expires_Which = 2
)

func (w WebSession_Cookie_expires_Which) String() string {
	const s = "noneabsoluterelative"
	switch w {
	case WebSession_Cookie_expires_Which_none:
		return s[0:4]
	case WebSession_Cookie_expires_Which_absolute:
		return s[4:12]
	case WebSession_Cookie_expires_Which_relative:
		return s[12:20]

	}
	return "WebSession_Cookie_expires_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewWebSession_Cookie(s *capnp.Segment) (WebSession_Cookie, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 3})
	return WebSession_Cookie{st}, err
}

func NewRootWebSession_Cookie(s *capnp.Segment) (WebSession_Cookie, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 3})
	return WebSession_Cookie{st}, err
}

func ReadRootWebSession_Cookie(msg *capnp.Message) (WebSession_Cookie, error) {
	root, err := msg.RootPtr()
	return WebSession_Cookie{root.Struct()}, err
}

func (s WebSession_Cookie) String() string {
	str, _ := text.Marshal(0xa87d65bed9b60243, s.Struct)
	return str
}

func (s WebSession_Cookie) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_Cookie) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_Cookie) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_Cookie) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Cookie) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Cookie) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Cookie) Expires() WebSession_Cookie_expires { return WebSession_Cookie_expires(s) }

func (s WebSession_Cookie_expires) Which() WebSession_Cookie_expires_Which {
	return WebSession_Cookie_expires_Which(s.Struct.Uint16(0))
}
func (s WebSession_Cookie_expires) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s WebSession_Cookie_expires) Absolute() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s WebSession_Cookie_expires) SetAbsolute(v int64) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint64(8, uint64(v))
}

func (s WebSession_Cookie_expires) Relative() uint64 {
	return s.Struct.Uint64(8)
}

func (s WebSession_Cookie_expires) SetRelative(v uint64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, v)
}

func (s WebSession_Cookie) HttpOnly() bool {
	return s.Struct.Bit(16)
}

func (s WebSession_Cookie) SetHttpOnly(v bool) {
	s.Struct.SetBit(16, v)
}

func (s WebSession_Cookie) Path() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSession_Cookie) HasPath() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSession_Cookie) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// WebSession_Cookie_List is a list of WebSession_Cookie.
type WebSession_Cookie_List struct{ capnp.List }

// NewWebSession_Cookie creates a new list of WebSession_Cookie.
func NewWebSession_Cookie_List(s *capnp.Segment, sz int32) (WebSession_Cookie_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 3}, sz)
	return WebSession_Cookie_List{l}, err
}

func (s WebSession_Cookie_List) At(i int) WebSession_Cookie {
	return WebSession_Cookie{s.List.Struct(i)}
}

func (s WebSession_Cookie_List) Set(i int, v WebSession_Cookie) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Cookie_Promise is a wrapper for a WebSession_Cookie promised by a client call.
type WebSession_Cookie_Promise struct{ *capnp.Pipeline }

func (p WebSession_Cookie_Promise) Struct() (WebSession_Cookie, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Cookie{s}, err
}

func (p WebSession_Cookie_Promise) Expires() WebSession_Cookie_expires_Promise {
	return WebSession_Cookie_expires_Promise{p.Pipeline}
}

// WebSession_Cookie_expires_Promise is a wrapper for a WebSession_Cookie_expires promised by a client call.
type WebSession_Cookie_expires_Promise struct{ *capnp.Pipeline }

func (p WebSession_Cookie_expires_Promise) Struct() (WebSession_Cookie_expires, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Cookie_expires{s}, err
}

type WebSession_AcceptedType struct{ capnp.Struct }

func NewWebSession_AcceptedType(s *capnp.Segment) (WebSession_AcceptedType, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_AcceptedType{st}, err
}

func NewRootWebSession_AcceptedType(s *capnp.Segment) (WebSession_AcceptedType, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_AcceptedType{st}, err
}

func ReadRootWebSession_AcceptedType(msg *capnp.Message) (WebSession_AcceptedType, error) {
	root, err := msg.RootPtr()
	return WebSession_AcceptedType{root.Struct()}, err
}

func (s WebSession_AcceptedType) String() string {
	str, _ := text.Marshal(0xaaf9021b627cc1f9, s.Struct)
	return str
}

func (s WebSession_AcceptedType) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_AcceptedType) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_AcceptedType) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_AcceptedType) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_AcceptedType) QValue() float32 {
	return math.Float32frombits(s.Struct.Uint32(0) ^ 0x3f800000)
}

func (s WebSession_AcceptedType) SetQValue(v float32) {
	s.Struct.SetUint32(0, math.Float32bits(v)^0x3f800000)
}

// WebSession_AcceptedType_List is a list of WebSession_AcceptedType.
type WebSession_AcceptedType_List struct{ capnp.List }

// NewWebSession_AcceptedType creates a new list of WebSession_AcceptedType.
func NewWebSession_AcceptedType_List(s *capnp.Segment, sz int32) (WebSession_AcceptedType_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return WebSession_AcceptedType_List{l}, err
}

func (s WebSession_AcceptedType_List) At(i int) WebSession_AcceptedType {
	return WebSession_AcceptedType{s.List.Struct(i)}
}

func (s WebSession_AcceptedType_List) Set(i int, v WebSession_AcceptedType) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_AcceptedType_Promise is a wrapper for a WebSession_AcceptedType promised by a client call.
type WebSession_AcceptedType_Promise struct{ *capnp.Pipeline }

func (p WebSession_AcceptedType_Promise) Struct() (WebSession_AcceptedType, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_AcceptedType{s}, err
}

type WebSession_Response struct{ capnp.Struct }
type WebSession_Response_content WebSession_Response
type WebSession_Response_content_body WebSession_Response
type WebSession_Response_content_disposition WebSession_Response
type WebSession_Response_noContent WebSession_Response
type WebSession_Response_preconditionFailed WebSession_Response
type WebSession_Response_redirect WebSession_Response
type WebSession_Response_clientError WebSession_Response
type WebSession_Response_serverError WebSession_Response
type WebSession_Response_Which uint16

const (
	WebSession_Response_Which_content            WebSession_Response_Which = 1
	WebSession_Response_Which_noContent          WebSession_Response_Which = 4
	WebSession_Response_Which_preconditionFailed WebSession_Response_Which = 5
	WebSession_Response_Which_redirect           WebSession_Response_Which = 0
	WebSession_Response_Which_clientError        WebSession_Response_Which = 2
	WebSession_Response_Which_serverError        WebSession_Response_Which = 3
)

func (w WebSession_Response_Which) String() string {
	const s = "contentnoContentpreconditionFailedredirectclientErrorserverError"
	switch w {
	case WebSession_Response_Which_content:
		return s[0:7]
	case WebSession_Response_Which_noContent:
		return s[7:16]
	case WebSession_Response_Which_preconditionFailed:
		return s[16:34]
	case WebSession_Response_Which_redirect:
		return s[34:42]
	case WebSession_Response_Which_clientError:
		return s[42:53]
	case WebSession_Response_Which_serverError:
		return s[53:64]

	}
	return "WebSession_Response_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type WebSession_Response_content_body_Which uint16

const (
	WebSession_Response_content_body_Which_bytes  WebSession_Response_content_body_Which = 0
	WebSession_Response_content_body_Which_stream WebSession_Response_content_body_Which = 1
)

func (w WebSession_Response_content_body_Which) String() string {
	const s = "bytesstream"
	switch w {
	case WebSession_Response_content_body_Which_bytes:
		return s[0:5]
	case WebSession_Response_content_body_Which_stream:
		return s[5:11]

	}
	return "WebSession_Response_content_body_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type WebSession_Response_content_disposition_Which uint16

const (
	WebSession_Response_content_disposition_Which_normal   WebSession_Response_content_disposition_Which = 0
	WebSession_Response_content_disposition_Which_download WebSession_Response_content_disposition_Which = 1
)

func (w WebSession_Response_content_disposition_Which) String() string {
	const s = "normaldownload"
	switch w {
	case WebSession_Response_content_disposition_Which_normal:
		return s[0:6]
	case WebSession_Response_content_disposition_Which_download:
		return s[6:14]

	}
	return "WebSession_Response_content_disposition_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewWebSession_Response(s *capnp.Segment) (WebSession_Response, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 9})
	return WebSession_Response{st}, err
}

func NewRootWebSession_Response(s *capnp.Segment) (WebSession_Response, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 9})
	return WebSession_Response{st}, err
}

func ReadRootWebSession_Response(msg *capnp.Message) (WebSession_Response, error) {
	root, err := msg.RootPtr()
	return WebSession_Response{root.Struct()}, err
}

func (s WebSession_Response) String() string {
	str, _ := text.Marshal(0x8193ac6cb5429c83, s.Struct)
	return str
}

func (s WebSession_Response) Which() WebSession_Response_Which {
	return WebSession_Response_Which(s.Struct.Uint16(2))
}
func (s WebSession_Response) SetCookies() (WebSession_Cookie_List, error) {
	p, err := s.Struct.Ptr(0)
	return WebSession_Cookie_List{List: p.List()}, err
}

func (s WebSession_Response) HasSetCookies() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Response) SetSetCookies(v WebSession_Cookie_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSetCookies sets the setCookies field to a newly
// allocated WebSession_Cookie_List, preferring placement in s's segment.
func (s WebSession_Response) NewSetCookies(n int32) (WebSession_Cookie_List, error) {
	l, err := NewWebSession_Cookie_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_Cookie_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s WebSession_Response) CachePolicy() (WebSession_CachePolicy, error) {
	p, err := s.Struct.Ptr(6)
	return WebSession_CachePolicy{Struct: p.Struct()}, err
}

func (s WebSession_Response) HasCachePolicy() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s WebSession_Response) SetCachePolicy(v WebSession_CachePolicy) error {
	return s.Struct.SetPtr(6, v.Struct.ToPtr())
}

// NewCachePolicy sets the cachePolicy field to a newly
// allocated WebSession_CachePolicy struct, preferring placement in s's segment.
func (s WebSession_Response) NewCachePolicy() (WebSession_CachePolicy, error) {
	ss, err := NewWebSession_CachePolicy(s.Struct.Segment())
	if err != nil {
		return WebSession_CachePolicy{}, err
	}
	err = s.Struct.SetPtr(6, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_Response) Content() WebSession_Response_content {
	return WebSession_Response_content(s)
}

func (s WebSession_Response) SetContent() {
	s.Struct.SetUint16(2, 1)
}

func (s WebSession_Response_content) StatusCode() WebSession_Response_SuccessCode {
	return WebSession_Response_SuccessCode(s.Struct.Uint16(4))
}

func (s WebSession_Response_content) SetStatusCode(v WebSession_Response_SuccessCode) {
	s.Struct.SetUint16(4, uint16(v))
}

func (s WebSession_Response_content) Encoding() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Response_content) HasEncoding() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Response_content) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Response_content) Language() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSession_Response_content) HasLanguage() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) LanguageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSession_Response_content) SetLanguage(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s WebSession_Response_content) MimeType() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s WebSession_Response_content) HasMimeType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s WebSession_Response_content) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s WebSession_Response_content) ETag() (WebSession_ETag, error) {
	p, err := s.Struct.Ptr(7)
	return WebSession_ETag{Struct: p.Struct()}, err
}

func (s WebSession_Response_content) HasETag() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) SetETag(v WebSession_ETag) error {
	return s.Struct.SetPtr(7, v.Struct.ToPtr())
}

// NewETag sets the eTag field to a newly
// allocated WebSession_ETag struct, preferring placement in s's segment.
func (s WebSession_Response_content) NewETag() (WebSession_ETag, error) {
	ss, err := NewWebSession_ETag(s.Struct.Segment())
	if err != nil {
		return WebSession_ETag{}, err
	}
	err = s.Struct.SetPtr(7, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_Response_content) Body() WebSession_Response_content_body {
	return WebSession_Response_content_body(s)
}

func (s WebSession_Response_content_body) Which() WebSession_Response_content_body_Which {
	return WebSession_Response_content_body_Which(s.Struct.Uint16(0))
}
func (s WebSession_Response_content_body) Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return []byte(p.Data()), err
}

func (s WebSession_Response_content_body) HasBytes() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content_body) SetBytes(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, d.List.ToPtr())
}

func (s WebSession_Response_content_body) Stream() util.Handle {
	p, _ := s.Struct.Ptr(4)
	return util.Handle{Client: p.Interface().Client()}
}

func (s WebSession_Response_content_body) HasStream() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content_body) SetStream(v util.Handle) error {
	s.Struct.SetUint16(0, 1)
	if v.Client == nil {
		return s.Struct.SetPtr(4, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(4, in.ToPtr())
}

func (s WebSession_Response_content) Disposition() WebSession_Response_content_disposition {
	return WebSession_Response_content_disposition(s)
}

func (s WebSession_Response_content_disposition) Which() WebSession_Response_content_disposition_Which {
	return WebSession_Response_content_disposition_Which(s.Struct.Uint16(6))
}
func (s WebSession_Response_content_disposition) SetNormal() {
	s.Struct.SetUint16(6, 0)

}

func (s WebSession_Response_content_disposition) Download() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s WebSession_Response_content_disposition) HasDownload() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content_disposition) DownloadBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s WebSession_Response_content_disposition) SetDownload(v string) error {
	s.Struct.SetUint16(6, 1)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, t.List.ToPtr())
}

func (s WebSession_Response) NoContent() WebSession_Response_noContent {
	return WebSession_Response_noContent(s)
}

func (s WebSession_Response) SetNoContent() {
	s.Struct.SetUint16(2, 4)
}

func (s WebSession_Response_noContent) ShouldResetForm() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_Response_noContent) SetShouldResetForm(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_Response_noContent) ETag() (WebSession_ETag, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_ETag{Struct: p.Struct()}, err
}

func (s WebSession_Response_noContent) HasETag() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_noContent) SetETag(v WebSession_ETag) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewETag sets the eTag field to a newly
// allocated WebSession_ETag struct, preferring placement in s's segment.
func (s WebSession_Response_noContent) NewETag() (WebSession_ETag, error) {
	ss, err := NewWebSession_ETag(s.Struct.Segment())
	if err != nil {
		return WebSession_ETag{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_Response) PreconditionFailed() WebSession_Response_preconditionFailed {
	return WebSession_Response_preconditionFailed(s)
}

func (s WebSession_Response) SetPreconditionFailed() {
	s.Struct.SetUint16(2, 5)
}

func (s WebSession_Response_preconditionFailed) MatchingETag() (WebSession_ETag, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_ETag{Struct: p.Struct()}, err
}

func (s WebSession_Response_preconditionFailed) HasMatchingETag() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_preconditionFailed) SetMatchingETag(v WebSession_ETag) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewMatchingETag sets the matchingETag field to a newly
// allocated WebSession_ETag struct, preferring placement in s's segment.
func (s WebSession_Response_preconditionFailed) NewMatchingETag() (WebSession_ETag, error) {
	ss, err := NewWebSession_ETag(s.Struct.Segment())
	if err != nil {
		return WebSession_ETag{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_Response) Redirect() WebSession_Response_redirect {
	return WebSession_Response_redirect(s)
}

func (s WebSession_Response) SetRedirect() {
	s.Struct.SetUint16(2, 0)
}

func (s WebSession_Response_redirect) IsPermanent() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_Response_redirect) SetIsPermanent(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_Response_redirect) SwitchToGet() bool {
	return s.Struct.Bit(1)
}

func (s WebSession_Response_redirect) SetSwitchToGet(v bool) {
	s.Struct.SetBit(1, v)
}

func (s WebSession_Response_redirect) Location() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Response_redirect) HasLocation() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_redirect) LocationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Response_redirect) SetLocation(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Response) ClientError() WebSession_Response_clientError {
	return WebSession_Response_clientError(s)
}

func (s WebSession_Response) SetClientError() {
	s.Struct.SetUint16(2, 2)
}

func (s WebSession_Response_clientError) StatusCode() WebSession_Response_ClientErrorCode {
	return WebSession_Response_ClientErrorCode(s.Struct.Uint16(0))
}

func (s WebSession_Response_clientError) SetStatusCode(v WebSession_Response_ClientErrorCode) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s WebSession_Response_clientError) DescriptionHtml() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Response_clientError) HasDescriptionHtml() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_clientError) DescriptionHtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Response_clientError) SetDescriptionHtml(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Response) ServerError() WebSession_Response_serverError {
	return WebSession_Response_serverError(s)
}

func (s WebSession_Response) SetServerError() {
	s.Struct.SetUint16(2, 3)
}

func (s WebSession_Response_serverError) DescriptionHtml() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Response_serverError) HasDescriptionHtml() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_serverError) DescriptionHtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Response_serverError) SetDescriptionHtml(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_Response) AdditionalHeaders() (WebSession_Response_Header_List, error) {
	p, err := s.Struct.Ptr(8)
	return WebSession_Response_Header_List{List: p.List()}, err
}

func (s WebSession_Response) HasAdditionalHeaders() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s WebSession_Response) SetAdditionalHeaders(v WebSession_Response_Header_List) error {
	return s.Struct.SetPtr(8, v.List.ToPtr())
}

// NewAdditionalHeaders sets the additionalHeaders field to a newly
// allocated WebSession_Response_Header_List, preferring placement in s's segment.
func (s WebSession_Response) NewAdditionalHeaders(n int32) (WebSession_Response_Header_List, error) {
	l, err := NewWebSession_Response_Header_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_Response_Header_List{}, err
	}
	err = s.Struct.SetPtr(8, l.List.ToPtr())
	return l, err
}

// WebSession_Response_List is a list of WebSession_Response.
type WebSession_Response_List struct{ capnp.List }

// NewWebSession_Response creates a new list of WebSession_Response.
func NewWebSession_Response_List(s *capnp.Segment, sz int32) (WebSession_Response_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 9}, sz)
	return WebSession_Response_List{l}, err
}

func (s WebSession_Response_List) At(i int) WebSession_Response {
	return WebSession_Response{s.List.Struct(i)}
}

func (s WebSession_Response_List) Set(i int, v WebSession_Response) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Response_Promise is a wrapper for a WebSession_Response promised by a client call.
type WebSession_Response_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_Promise) Struct() (WebSession_Response, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response{s}, err
}

func (p WebSession_Response_Promise) CachePolicy() WebSession_CachePolicy_Promise {
	return WebSession_CachePolicy_Promise{Pipeline: p.Pipeline.GetPipeline(6)}
}

func (p WebSession_Response_Promise) Content() WebSession_Response_content_Promise {
	return WebSession_Response_content_Promise{p.Pipeline}
}

// WebSession_Response_content_Promise is a wrapper for a WebSession_Response_content promised by a client call.
type WebSession_Response_content_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_content_Promise) Struct() (WebSession_Response_content, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_content{s}, err
}

func (p WebSession_Response_content_Promise) ETag() WebSession_ETag_Promise {
	return WebSession_ETag_Promise{Pipeline: p.Pipeline.GetPipeline(7)}
}

func (p WebSession_Response_content_Promise) Body() WebSession_Response_content_body_Promise {
	return WebSession_Response_content_body_Promise{p.Pipeline}
}

// WebSession_Response_content_body_Promise is a wrapper for a WebSession_Response_content_body promised by a client call.
type WebSession_Response_content_body_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_content_body_Promise) Struct() (WebSession_Response_content_body, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_content_body{s}, err
}

func (p WebSession_Response_content_body_Promise) Stream() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(4).Client()}
}

func (p WebSession_Response_content_Promise) Disposition() WebSession_Response_content_disposition_Promise {
	return WebSession_Response_content_disposition_Promise{p.Pipeline}
}

// WebSession_Response_content_disposition_Promise is a wrapper for a WebSession_Response_content_disposition promised by a client call.
type WebSession_Response_content_disposition_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_content_disposition_Promise) Struct() (WebSession_Response_content_disposition, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_content_disposition{s}, err
}

func (p WebSession_Response_Promise) NoContent() WebSession_Response_noContent_Promise {
	return WebSession_Response_noContent_Promise{p.Pipeline}
}

// WebSession_Response_noContent_Promise is a wrapper for a WebSession_Response_noContent promised by a client call.
type WebSession_Response_noContent_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_noContent_Promise) Struct() (WebSession_Response_noContent, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_noContent{s}, err
}

func (p WebSession_Response_noContent_Promise) ETag() WebSession_ETag_Promise {
	return WebSession_ETag_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_Response_Promise) PreconditionFailed() WebSession_Response_preconditionFailed_Promise {
	return WebSession_Response_preconditionFailed_Promise{p.Pipeline}
}

// WebSession_Response_preconditionFailed_Promise is a wrapper for a WebSession_Response_preconditionFailed promised by a client call.
type WebSession_Response_preconditionFailed_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_preconditionFailed_Promise) Struct() (WebSession_Response_preconditionFailed, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_preconditionFailed{s}, err
}

func (p WebSession_Response_preconditionFailed_Promise) MatchingETag() WebSession_ETag_Promise {
	return WebSession_ETag_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_Response_Promise) Redirect() WebSession_Response_redirect_Promise {
	return WebSession_Response_redirect_Promise{p.Pipeline}
}

// WebSession_Response_redirect_Promise is a wrapper for a WebSession_Response_redirect promised by a client call.
type WebSession_Response_redirect_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_redirect_Promise) Struct() (WebSession_Response_redirect, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_redirect{s}, err
}

func (p WebSession_Response_Promise) ClientError() WebSession_Response_clientError_Promise {
	return WebSession_Response_clientError_Promise{p.Pipeline}
}

// WebSession_Response_clientError_Promise is a wrapper for a WebSession_Response_clientError promised by a client call.
type WebSession_Response_clientError_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_clientError_Promise) Struct() (WebSession_Response_clientError, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_clientError{s}, err
}

func (p WebSession_Response_Promise) ServerError() WebSession_Response_serverError_Promise {
	return WebSession_Response_serverError_Promise{p.Pipeline}
}

// WebSession_Response_serverError_Promise is a wrapper for a WebSession_Response_serverError promised by a client call.
type WebSession_Response_serverError_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_serverError_Promise) Struct() (WebSession_Response_serverError, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_serverError{s}, err
}

type WebSession_Response_SuccessCode uint16

// Values of WebSession_Response_SuccessCode.
const (
	WebSession_Response_SuccessCode_ok             WebSession_Response_SuccessCode = 0
	WebSession_Response_SuccessCode_created        WebSession_Response_SuccessCode = 1
	WebSession_Response_SuccessCode_accepted       WebSession_Response_SuccessCode = 2
	WebSession_Response_SuccessCode_noContent      WebSession_Response_SuccessCode = 3
	WebSession_Response_SuccessCode_partialContent WebSession_Response_SuccessCode = 4
	WebSession_Response_SuccessCode_multiStatus    WebSession_Response_SuccessCode = 5
	WebSession_Response_SuccessCode_notModified    WebSession_Response_SuccessCode = 6
)

// String returns the enum's constant name.
func (c WebSession_Response_SuccessCode) String() string {
	switch c {
	case WebSession_Response_SuccessCode_ok:
		return "ok"
	case WebSession_Response_SuccessCode_created:
		return "created"
	case WebSession_Response_SuccessCode_accepted:
		return "accepted"
	case WebSession_Response_SuccessCode_noContent:
		return "noContent"
	case WebSession_Response_SuccessCode_partialContent:
		return "partialContent"
	case WebSession_Response_SuccessCode_multiStatus:
		return "multiStatus"
	case WebSession_Response_SuccessCode_notModified:
		return "notModified"

	default:
		return ""
	}
}

// WebSession_Response_SuccessCodeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func WebSession_Response_SuccessCodeFromString(c string) WebSession_Response_SuccessCode {
	switch c {
	case "ok":
		return WebSession_Response_SuccessCode_ok
	case "created":
		return WebSession_Response_SuccessCode_created
	case "accepted":
		return WebSession_Response_SuccessCode_accepted
	case "noContent":
		return WebSession_Response_SuccessCode_noContent
	case "partialContent":
		return WebSession_Response_SuccessCode_partialContent
	case "multiStatus":
		return WebSession_Response_SuccessCode_multiStatus
	case "notModified":
		return WebSession_Response_SuccessCode_notModified

	default:
		return 0
	}
}

type WebSession_Response_SuccessCode_List struct{ capnp.List }

func NewWebSession_Response_SuccessCode_List(s *capnp.Segment, sz int32) (WebSession_Response_SuccessCode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return WebSession_Response_SuccessCode_List{l.List}, err
}

func (l WebSession_Response_SuccessCode_List) At(i int) WebSession_Response_SuccessCode {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_Response_SuccessCode(ul.At(i))
}

func (l WebSession_Response_SuccessCode_List) Set(i int, v WebSession_Response_SuccessCode) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type WebSession_Response_ClientErrorCode uint16

// Values of WebSession_Response_ClientErrorCode.
const (
	WebSession_Response_ClientErrorCode_badRequest            WebSession_Response_ClientErrorCode = 0
	WebSession_Response_ClientErrorCode_forbidden             WebSession_Response_ClientErrorCode = 1
	WebSession_Response_ClientErrorCode_notFound              WebSession_Response_ClientErrorCode = 2
	WebSession_Response_ClientErrorCode_methodNotAllowed      WebSession_Response_ClientErrorCode = 3
	WebSession_Response_ClientErrorCode_notAcceptable         WebSession_Response_ClientErrorCode = 4
	WebSession_Response_ClientErrorCode_conflict              WebSession_Response_ClientErrorCode = 5
	WebSession_Response_ClientErrorCode_gone                  WebSession_Response_ClientErrorCode = 6
	WebSession_Response_ClientErrorCode_preconditionFailed    WebSession_Response_ClientErrorCode = 11
	WebSession_Response_ClientErrorCode_requestEntityTooLarge WebSession_Response_ClientErrorCode = 7
	WebSession_Response_ClientErrorCode_requestUriTooLong     WebSession_Response_ClientErrorCode = 8
	WebSession_Response_ClientErrorCode_unsupportedMediaType  WebSession_Response_ClientErrorCode = 9
	WebSession_Response_ClientErrorCode_imATeapot             WebSession_Response_ClientErrorCode = 10
	WebSession_Response_ClientErrorCode_unprocessableEntity   WebSession_Response_ClientErrorCode = 12
)

// String returns the enum's constant name.
func (c WebSession_Response_ClientErrorCode) String() string {
	switch c {
	case WebSession_Response_ClientErrorCode_badRequest:
		return "badRequest"
	case WebSession_Response_ClientErrorCode_forbidden:
		return "forbidden"
	case WebSession_Response_ClientErrorCode_notFound:
		return "notFound"
	case WebSession_Response_ClientErrorCode_methodNotAllowed:
		return "methodNotAllowed"
	case WebSession_Response_ClientErrorCode_notAcceptable:
		return "notAcceptable"
	case WebSession_Response_ClientErrorCode_conflict:
		return "conflict"
	case WebSession_Response_ClientErrorCode_gone:
		return "gone"
	case WebSession_Response_ClientErrorCode_preconditionFailed:
		return "preconditionFailed"
	case WebSession_Response_ClientErrorCode_requestEntityTooLarge:
		return "requestEntityTooLarge"
	case WebSession_Response_ClientErrorCode_requestUriTooLong:
		return "requestUriTooLong"
	case WebSession_Response_ClientErrorCode_unsupportedMediaType:
		return "unsupportedMediaType"
	case WebSession_Response_ClientErrorCode_imATeapot:
		return "imATeapot"
	case WebSession_Response_ClientErrorCode_unprocessableEntity:
		return "unprocessableEntity"

	default:
		return ""
	}
}

// WebSession_Response_ClientErrorCodeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func WebSession_Response_ClientErrorCodeFromString(c string) WebSession_Response_ClientErrorCode {
	switch c {
	case "badRequest":
		return WebSession_Response_ClientErrorCode_badRequest
	case "forbidden":
		return WebSession_Response_ClientErrorCode_forbidden
	case "notFound":
		return WebSession_Response_ClientErrorCode_notFound
	case "methodNotAllowed":
		return WebSession_Response_ClientErrorCode_methodNotAllowed
	case "notAcceptable":
		return WebSession_Response_ClientErrorCode_notAcceptable
	case "conflict":
		return WebSession_Response_ClientErrorCode_conflict
	case "gone":
		return WebSession_Response_ClientErrorCode_gone
	case "preconditionFailed":
		return WebSession_Response_ClientErrorCode_preconditionFailed
	case "requestEntityTooLarge":
		return WebSession_Response_ClientErrorCode_requestEntityTooLarge
	case "requestUriTooLong":
		return WebSession_Response_ClientErrorCode_requestUriTooLong
	case "unsupportedMediaType":
		return WebSession_Response_ClientErrorCode_unsupportedMediaType
	case "imATeapot":
		return WebSession_Response_ClientErrorCode_imATeapot
	case "unprocessableEntity":
		return WebSession_Response_ClientErrorCode_unprocessableEntity

	default:
		return 0
	}
}

type WebSession_Response_ClientErrorCode_List struct{ capnp.List }

func NewWebSession_Response_ClientErrorCode_List(s *capnp.Segment, sz int32) (WebSession_Response_ClientErrorCode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return WebSession_Response_ClientErrorCode_List{l.List}, err
}

func (l WebSession_Response_ClientErrorCode_List) At(i int) WebSession_Response_ClientErrorCode {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_Response_ClientErrorCode(ul.At(i))
}

func (l WebSession_Response_ClientErrorCode_List) Set(i int, v WebSession_Response_ClientErrorCode) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type WebSession_Response_Header struct{ capnp.Struct }

func NewWebSession_Response_Header(s *capnp.Segment) (WebSession_Response_Header, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_Response_Header{st}, err
}

func NewRootWebSession_Response_Header(s *capnp.Segment) (WebSession_Response_Header, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_Response_Header{st}, err
}

func ReadRootWebSession_Response_Header(msg *capnp.Message) (WebSession_Response_Header, error) {
	root, err := msg.RootPtr()
	return WebSession_Response_Header{root.Struct()}, err
}

func (s WebSession_Response_Header) String() string {
	str, _ := text.Marshal(0xb4b873147ab5ce5e, s.Struct)
	return str
}

func (s WebSession_Response_Header) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_Response_Header) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_Header) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_Response_Header) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_Response_Header) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_Response_Header) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_Header) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_Response_Header) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// WebSession_Response_Header_List is a list of WebSession_Response_Header.
type WebSession_Response_Header_List struct{ capnp.List }

// NewWebSession_Response_Header creates a new list of WebSession_Response_Header.
func NewWebSession_Response_Header_List(s *capnp.Segment, sz int32) (WebSession_Response_Header_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSession_Response_Header_List{l}, err
}

func (s WebSession_Response_Header_List) At(i int) WebSession_Response_Header {
	return WebSession_Response_Header{s.List.Struct(i)}
}

func (s WebSession_Response_Header_List) Set(i int, v WebSession_Response_Header) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Response_Header_Promise is a wrapper for a WebSession_Response_Header promised by a client call.
type WebSession_Response_Header_Promise struct{ *capnp.Pipeline }

func (p WebSession_Response_Header_Promise) Struct() (WebSession_Response_Header, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Response_Header{s}, err
}

type WebSession_RequestStream struct{ Client capnp.Client }

func (c WebSession_RequestStream) GetResponse(ctx context.Context, params func(WebSession_RequestStream_getResponse_Params) error, opts ...capnp.CallOption) WebSession_Response_Promise {
	if c.Client == nil {
		return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x99ffc2f3f69a6a9f,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession.RequestStream",
			MethodName:    "getResponse",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_RequestStream_getResponse_Params{Struct: s}) }
	}
	return WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession_RequestStream) Write(ctx context.Context, params func(util.ByteStream_write_Params) error, opts ...capnp.CallOption) util.ByteStream_write_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_write_Params{Struct: s}) }
	}
	return util.ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession_RequestStream) Done(ctx context.Context, params func(util.ByteStream_done_Params) error, opts ...capnp.CallOption) util.ByteStream_done_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_done_Params{Struct: s}) }
	}
	return util.ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSession_RequestStream) ExpectSize(ctx context.Context, params func(util.ByteStream_expectSize_Params) error, opts ...capnp.CallOption) util.ByteStream_expectSize_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_expectSize_Params{Struct: s}) }
	}
	return util.ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type WebSession_RequestStream_Server interface {
	GetResponse(WebSession_RequestStream_getResponse) error

	Write(util.ByteStream_write) error

	Done(util.ByteStream_done) error

	ExpectSize(util.ByteStream_expectSize) error
}

func WebSession_RequestStream_ServerToClient(s WebSession_RequestStream_Server) WebSession_RequestStream {
	c, _ := s.(server.Closer)
	return WebSession_RequestStream{Client: server.New(WebSession_RequestStream_Methods(nil, s), c)}
}

func WebSession_RequestStream_Methods(methods []server.Method, s WebSession_RequestStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x99ffc2f3f69a6a9f,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession.RequestStream",
			MethodName:    "getResponse",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSession_RequestStream_getResponse{c, opts, WebSession_RequestStream_getResponse_Params{Struct: p}, WebSession_Response{Struct: r}}
			return s.GetResponse(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_write{c, opts, util.ByteStream_write_Params{Struct: p}, util.ByteStream_write_Results{Struct: r}}
			return s.Write(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_done{c, opts, util.ByteStream_done_Params{Struct: p}, util.ByteStream_done_Results{Struct: r}}
			return s.Done(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_expectSize{c, opts, util.ByteStream_expectSize_Params{Struct: p}, util.ByteStream_expectSize_Results{Struct: r}}
			return s.ExpectSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// WebSession_RequestStream_getResponse holds the arguments for a server call to WebSession_RequestStream.getResponse.
type WebSession_RequestStream_getResponse struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_RequestStream_getResponse_Params
	Results WebSession_Response
}

type WebSession_RequestStream_getResponse_Params struct{ capnp.Struct }

func NewWebSession_RequestStream_getResponse_Params(s *capnp.Segment) (WebSession_RequestStream_getResponse_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSession_RequestStream_getResponse_Params{st}, err
}

func NewRootWebSession_RequestStream_getResponse_Params(s *capnp.Segment) (WebSession_RequestStream_getResponse_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSession_RequestStream_getResponse_Params{st}, err
}

func ReadRootWebSession_RequestStream_getResponse_Params(msg *capnp.Message) (WebSession_RequestStream_getResponse_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_RequestStream_getResponse_Params{root.Struct()}, err
}

func (s WebSession_RequestStream_getResponse_Params) String() string {
	str, _ := text.Marshal(0xe9a02a3219bdbd70, s.Struct)
	return str
}

// WebSession_RequestStream_getResponse_Params_List is a list of WebSession_RequestStream_getResponse_Params.
type WebSession_RequestStream_getResponse_Params_List struct{ capnp.List }

// NewWebSession_RequestStream_getResponse_Params creates a new list of WebSession_RequestStream_getResponse_Params.
func NewWebSession_RequestStream_getResponse_Params_List(s *capnp.Segment, sz int32) (WebSession_RequestStream_getResponse_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSession_RequestStream_getResponse_Params_List{l}, err
}

func (s WebSession_RequestStream_getResponse_Params_List) At(i int) WebSession_RequestStream_getResponse_Params {
	return WebSession_RequestStream_getResponse_Params{s.List.Struct(i)}
}

func (s WebSession_RequestStream_getResponse_Params_List) Set(i int, v WebSession_RequestStream_getResponse_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_RequestStream_getResponse_Params_Promise is a wrapper for a WebSession_RequestStream_getResponse_Params promised by a client call.
type WebSession_RequestStream_getResponse_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_RequestStream_getResponse_Params_Promise) Struct() (WebSession_RequestStream_getResponse_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_RequestStream_getResponse_Params{s}, err
}

type WebSession_WebSocketStream struct{ Client capnp.Client }

func (c WebSession_WebSocketStream) SendBytes(ctx context.Context, params func(WebSession_WebSocketStream_sendBytes_Params) error, opts ...capnp.CallOption) WebSession_WebSocketStream_sendBytes_Results_Promise {
	if c.Client == nil {
		return WebSession_WebSocketStream_sendBytes_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf001fc1d5e574a07,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession.WebSocketStream",
			MethodName:    "sendBytes",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSession_WebSocketStream_sendBytes_Params{Struct: s}) }
	}
	return WebSession_WebSocketStream_sendBytes_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type WebSession_WebSocketStream_Server interface {
	SendBytes(WebSession_WebSocketStream_sendBytes) error
}

func WebSession_WebSocketStream_ServerToClient(s WebSession_WebSocketStream_Server) WebSession_WebSocketStream {
	c, _ := s.(server.Closer)
	return WebSession_WebSocketStream{Client: server.New(WebSession_WebSocketStream_Methods(nil, s), c)}
}

func WebSession_WebSocketStream_Methods(methods []server.Method, s WebSession_WebSocketStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf001fc1d5e574a07,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession.WebSocketStream",
			MethodName:    "sendBytes",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSession_WebSocketStream_sendBytes{c, opts, WebSession_WebSocketStream_sendBytes_Params{Struct: p}, WebSession_WebSocketStream_sendBytes_Results{Struct: r}}
			return s.SendBytes(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// WebSession_WebSocketStream_sendBytes holds the arguments for a server call to WebSession_WebSocketStream.sendBytes.
type WebSession_WebSocketStream_sendBytes struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSession_WebSocketStream_sendBytes_Params
	Results WebSession_WebSocketStream_sendBytes_Results
}

type WebSession_WebSocketStream_sendBytes_Params struct{ capnp.Struct }

func NewWebSession_WebSocketStream_sendBytes_Params(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_WebSocketStream_sendBytes_Params{st}, err
}

func NewRootWebSession_WebSocketStream_sendBytes_Params(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_WebSocketStream_sendBytes_Params{st}, err
}

func ReadRootWebSession_WebSocketStream_sendBytes_Params(msg *capnp.Message) (WebSession_WebSocketStream_sendBytes_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_WebSocketStream_sendBytes_Params{root.Struct()}, err
}

func (s WebSession_WebSocketStream_sendBytes_Params) String() string {
	str, _ := text.Marshal(0x9a712ce3fcad8cd8, s.Struct)
	return str
}

func (s WebSession_WebSocketStream_sendBytes_Params) Message() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s WebSession_WebSocketStream_sendBytes_Params) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_WebSocketStream_sendBytes_Params) SetMessage(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// WebSession_WebSocketStream_sendBytes_Params_List is a list of WebSession_WebSocketStream_sendBytes_Params.
type WebSession_WebSocketStream_sendBytes_Params_List struct{ capnp.List }

// NewWebSession_WebSocketStream_sendBytes_Params creates a new list of WebSession_WebSocketStream_sendBytes_Params.
func NewWebSession_WebSocketStream_sendBytes_Params_List(s *capnp.Segment, sz int32) (WebSession_WebSocketStream_sendBytes_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSession_WebSocketStream_sendBytes_Params_List{l}, err
}

func (s WebSession_WebSocketStream_sendBytes_Params_List) At(i int) WebSession_WebSocketStream_sendBytes_Params {
	return WebSession_WebSocketStream_sendBytes_Params{s.List.Struct(i)}
}

func (s WebSession_WebSocketStream_sendBytes_Params_List) Set(i int, v WebSession_WebSocketStream_sendBytes_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_WebSocketStream_sendBytes_Params_Promise is a wrapper for a WebSession_WebSocketStream_sendBytes_Params promised by a client call.
type WebSession_WebSocketStream_sendBytes_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_WebSocketStream_sendBytes_Params_Promise) Struct() (WebSession_WebSocketStream_sendBytes_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_WebSocketStream_sendBytes_Params{s}, err
}

type WebSession_WebSocketStream_sendBytes_Results struct{ capnp.Struct }

func NewWebSession_WebSocketStream_sendBytes_Results(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSession_WebSocketStream_sendBytes_Results{st}, err
}

func NewRootWebSession_WebSocketStream_sendBytes_Results(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSession_WebSocketStream_sendBytes_Results{st}, err
}

func ReadRootWebSession_WebSocketStream_sendBytes_Results(msg *capnp.Message) (WebSession_WebSocketStream_sendBytes_Results, error) {
	root, err := msg.RootPtr()
	return WebSession_WebSocketStream_sendBytes_Results{root.Struct()}, err
}

func (s WebSession_WebSocketStream_sendBytes_Results) String() string {
	str, _ := text.Marshal(0x82a3ee23aa0ae3a3, s.Struct)
	return str
}

// WebSession_WebSocketStream_sendBytes_Results_List is a list of WebSession_WebSocketStream_sendBytes_Results.
type WebSession_WebSocketStream_sendBytes_Results_List struct{ capnp.List }

// NewWebSession_WebSocketStream_sendBytes_Results creates a new list of WebSession_WebSocketStream_sendBytes_Results.
func NewWebSession_WebSocketStream_sendBytes_Results_List(s *capnp.Segment, sz int32) (WebSession_WebSocketStream_sendBytes_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSession_WebSocketStream_sendBytes_Results_List{l}, err
}

func (s WebSession_WebSocketStream_sendBytes_Results_List) At(i int) WebSession_WebSocketStream_sendBytes_Results {
	return WebSession_WebSocketStream_sendBytes_Results{s.List.Struct(i)}
}

func (s WebSession_WebSocketStream_sendBytes_Results_List) Set(i int, v WebSession_WebSocketStream_sendBytes_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_WebSocketStream_sendBytes_Results_Promise is a wrapper for a WebSession_WebSocketStream_sendBytes_Results promised by a client call.
type WebSession_WebSocketStream_sendBytes_Results_Promise struct{ *capnp.Pipeline }

func (p WebSession_WebSocketStream_sendBytes_Results_Promise) Struct() (WebSession_WebSocketStream_sendBytes_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_WebSocketStream_sendBytes_Results{s}, err
}

type WebSession_CachePolicy struct{ capnp.Struct }

func NewWebSession_CachePolicy(s *capnp.Segment) (WebSession_CachePolicy, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSession_CachePolicy{st}, err
}

func NewRootWebSession_CachePolicy(s *capnp.Segment) (WebSession_CachePolicy, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSession_CachePolicy{st}, err
}

func ReadRootWebSession_CachePolicy(msg *capnp.Message) (WebSession_CachePolicy, error) {
	root, err := msg.RootPtr()
	return WebSession_CachePolicy{root.Struct()}, err
}

func (s WebSession_CachePolicy) String() string {
	str, _ := text.Marshal(0xb37b21e300864885, s.Struct)
	return str
}

func (s WebSession_CachePolicy) WithCheck() WebSession_CachePolicy_Scope {
	return WebSession_CachePolicy_Scope(s.Struct.Uint16(0))
}

func (s WebSession_CachePolicy) SetWithCheck(v WebSession_CachePolicy_Scope) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s WebSession_CachePolicy) Permanent() WebSession_CachePolicy_Scope {
	return WebSession_CachePolicy_Scope(s.Struct.Uint16(2))
}

func (s WebSession_CachePolicy) SetPermanent(v WebSession_CachePolicy_Scope) {
	s.Struct.SetUint16(2, uint16(v))
}

func (s WebSession_CachePolicy) VariesOnCookie() bool {
	return s.Struct.Bit(32)
}

func (s WebSession_CachePolicy) SetVariesOnCookie(v bool) {
	s.Struct.SetBit(32, v)
}

func (s WebSession_CachePolicy) VariesOnAccept() bool {
	return s.Struct.Bit(33)
}

func (s WebSession_CachePolicy) SetVariesOnAccept(v bool) {
	s.Struct.SetBit(33, v)
}

// WebSession_CachePolicy_List is a list of WebSession_CachePolicy.
type WebSession_CachePolicy_List struct{ capnp.List }

// NewWebSession_CachePolicy creates a new list of WebSession_CachePolicy.
func NewWebSession_CachePolicy_List(s *capnp.Segment, sz int32) (WebSession_CachePolicy_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return WebSession_CachePolicy_List{l}, err
}

func (s WebSession_CachePolicy_List) At(i int) WebSession_CachePolicy {
	return WebSession_CachePolicy{s.List.Struct(i)}
}

func (s WebSession_CachePolicy_List) Set(i int, v WebSession_CachePolicy) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_CachePolicy_Promise is a wrapper for a WebSession_CachePolicy promised by a client call.
type WebSession_CachePolicy_Promise struct{ *capnp.Pipeline }

func (p WebSession_CachePolicy_Promise) Struct() (WebSession_CachePolicy, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_CachePolicy{s}, err
}

type WebSession_CachePolicy_Scope uint16

// Values of WebSession_CachePolicy_Scope.
const (
	WebSession_CachePolicy_Scope_none          WebSession_CachePolicy_Scope = 0
	WebSession_CachePolicy_Scope_perSession    WebSession_CachePolicy_Scope = 1
	WebSession_CachePolicy_Scope_perUser       WebSession_CachePolicy_Scope = 2
	WebSession_CachePolicy_Scope_perAppVersion WebSession_CachePolicy_Scope = 3
	WebSession_CachePolicy_Scope_universal     WebSession_CachePolicy_Scope = 4
)

// String returns the enum's constant name.
func (c WebSession_CachePolicy_Scope) String() string {
	switch c {
	case WebSession_CachePolicy_Scope_none:
		return "none"
	case WebSession_CachePolicy_Scope_perSession:
		return "perSession"
	case WebSession_CachePolicy_Scope_perUser:
		return "perUser"
	case WebSession_CachePolicy_Scope_perAppVersion:
		return "perAppVersion"
	case WebSession_CachePolicy_Scope_universal:
		return "universal"

	default:
		return ""
	}
}

// WebSession_CachePolicy_ScopeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func WebSession_CachePolicy_ScopeFromString(c string) WebSession_CachePolicy_Scope {
	switch c {
	case "none":
		return WebSession_CachePolicy_Scope_none
	case "perSession":
		return WebSession_CachePolicy_Scope_perSession
	case "perUser":
		return WebSession_CachePolicy_Scope_perUser
	case "perAppVersion":
		return WebSession_CachePolicy_Scope_perAppVersion
	case "universal":
		return WebSession_CachePolicy_Scope_universal

	default:
		return 0
	}
}

type WebSession_CachePolicy_Scope_List struct{ capnp.List }

func NewWebSession_CachePolicy_Scope_List(s *capnp.Segment, sz int32) (WebSession_CachePolicy_Scope_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return WebSession_CachePolicy_Scope_List{l.List}, err
}

func (l WebSession_CachePolicy_Scope_List) At(i int) WebSession_CachePolicy_Scope {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_CachePolicy_Scope(ul.At(i))
}

func (l WebSession_CachePolicy_Scope_List) Set(i int, v WebSession_CachePolicy_Scope) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type WebSession_Options struct{ capnp.Struct }

func NewWebSession_Options(s *capnp.Segment) (WebSession_Options, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_Options{st}, err
}

func NewRootWebSession_Options(s *capnp.Segment) (WebSession_Options, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_Options{st}, err
}

func ReadRootWebSession_Options(msg *capnp.Message) (WebSession_Options, error) {
	root, err := msg.RootPtr()
	return WebSession_Options{root.Struct()}, err
}

func (s WebSession_Options) String() string {
	str, _ := text.Marshal(0xe9ff06beec4e73d6, s.Struct)
	return str
}

func (s WebSession_Options) DavClass1() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_Options) SetDavClass1(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_Options) DavClass2() bool {
	return s.Struct.Bit(1)
}

func (s WebSession_Options) SetDavClass2(v bool) {
	s.Struct.SetBit(1, v)
}

func (s WebSession_Options) DavClass3() bool {
	return s.Struct.Bit(2)
}

func (s WebSession_Options) SetDavClass3(v bool) {
	s.Struct.SetBit(2, v)
}

func (s WebSession_Options) DavExtensions() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSession_Options) HasDavExtensions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Options) SetDavExtensions(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewDavExtensions sets the davExtensions field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSession_Options) NewDavExtensions(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// WebSession_Options_List is a list of WebSession_Options.
type WebSession_Options_List struct{ capnp.List }

// NewWebSession_Options creates a new list of WebSession_Options.
func NewWebSession_Options_List(s *capnp.Segment, sz int32) (WebSession_Options_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return WebSession_Options_List{l}, err
}

func (s WebSession_Options_List) At(i int) WebSession_Options {
	return WebSession_Options{s.List.Struct(i)}
}

func (s WebSession_Options_List) Set(i int, v WebSession_Options) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_Options_Promise is a wrapper for a WebSession_Options promised by a client call.
type WebSession_Options_Promise struct{ *capnp.Pipeline }

func (p WebSession_Options_Promise) Struct() (WebSession_Options, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_Options{s}, err
}

type WebSession_PropfindDepth uint16

// Values of WebSession_PropfindDepth.
const (
	WebSession_PropfindDepth_infinity WebSession_PropfindDepth = 0
	WebSession_PropfindDepth_zero     WebSession_PropfindDepth = 1
	WebSession_PropfindDepth_one      WebSession_PropfindDepth = 2
)

// String returns the enum's constant name.
func (c WebSession_PropfindDepth) String() string {
	switch c {
	case WebSession_PropfindDepth_infinity:
		return "infinity"
	case WebSession_PropfindDepth_zero:
		return "zero"
	case WebSession_PropfindDepth_one:
		return "one"

	default:
		return ""
	}
}

// WebSession_PropfindDepthFromString returns the enum value with a name,
// or the zero value if there's no such value.
func WebSession_PropfindDepthFromString(c string) WebSession_PropfindDepth {
	switch c {
	case "infinity":
		return WebSession_PropfindDepth_infinity
	case "zero":
		return WebSession_PropfindDepth_zero
	case "one":
		return WebSession_PropfindDepth_one

	default:
		return 0
	}
}

type WebSession_PropfindDepth_List struct{ capnp.List }

func NewWebSession_PropfindDepth_List(s *capnp.Segment, sz int32) (WebSession_PropfindDepth_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return WebSession_PropfindDepth_List{l.List}, err
}

func (l WebSession_PropfindDepth_List) At(i int) WebSession_PropfindDepth {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_PropfindDepth(ul.At(i))
}

func (l WebSession_PropfindDepth_List) Set(i int, v WebSession_PropfindDepth) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type WebSession_get_Params struct{ capnp.Struct }

func NewWebSession_get_Params(s *capnp.Segment) (WebSession_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return WebSession_get_Params{st}, err
}

func NewRootWebSession_get_Params(s *capnp.Segment) (WebSession_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return WebSession_get_Params{st}, err
}

func ReadRootWebSession_get_Params(msg *capnp.Message) (WebSession_get_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_get_Params{root.Struct()}, err
}

func (s WebSession_get_Params) String() string {
	str, _ := text.Marshal(0xcd94acddf4778328, s.Struct)
	return str
}

func (s WebSession_get_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_get_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_get_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_get_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_get_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_get_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_get_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_get_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_get_Params) IgnoreBody() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_get_Params) SetIgnoreBody(v bool) {
	s.Struct.SetBit(0, v)
}

// WebSession_get_Params_List is a list of WebSession_get_Params.
type WebSession_get_Params_List struct{ capnp.List }

// NewWebSession_get_Params creates a new list of WebSession_get_Params.
func NewWebSession_get_Params_List(s *capnp.Segment, sz int32) (WebSession_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return WebSession_get_Params_List{l}, err
}

func (s WebSession_get_Params_List) At(i int) WebSession_get_Params {
	return WebSession_get_Params{s.List.Struct(i)}
}

func (s WebSession_get_Params_List) Set(i int, v WebSession_get_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_get_Params_Promise is a wrapper for a WebSession_get_Params promised by a client call.
type WebSession_get_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_get_Params_Promise) Struct() (WebSession_get_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_get_Params{s}, err
}

func (p WebSession_get_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type WebSession_post_Params struct{ capnp.Struct }

func NewWebSession_post_Params(s *capnp.Segment) (WebSession_post_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_post_Params{st}, err
}

func NewRootWebSession_post_Params(s *capnp.Segment) (WebSession_post_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_post_Params{st}, err
}

func ReadRootWebSession_post_Params(msg *capnp.Message) (WebSession_post_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_post_Params{root.Struct()}, err
}

func (s WebSession_post_Params) String() string {
	str, _ := text.Marshal(0xaa6ef20a62c1cafd, s.Struct)
	return str
}

func (s WebSession_post_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_post_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_post_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_post_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_post_Params) Content() (WebSession_PostContent, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_PostContent{Struct: p.Struct()}, err
}

func (s WebSession_post_Params) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_post_Params) SetContent(v WebSession_PostContent) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContent sets the content field to a newly
// allocated WebSession_PostContent struct, preferring placement in s's segment.
func (s WebSession_post_Params) NewContent() (WebSession_PostContent, error) {
	ss, err := NewWebSession_PostContent(s.Struct.Segment())
	if err != nil {
		return WebSession_PostContent{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_post_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_post_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_post_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_post_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_post_Params_List is a list of WebSession_post_Params.
type WebSession_post_Params_List struct{ capnp.List }

// NewWebSession_post_Params creates a new list of WebSession_post_Params.
func NewWebSession_post_Params_List(s *capnp.Segment, sz int32) (WebSession_post_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_post_Params_List{l}, err
}

func (s WebSession_post_Params_List) At(i int) WebSession_post_Params {
	return WebSession_post_Params{s.List.Struct(i)}
}

func (s WebSession_post_Params_List) Set(i int, v WebSession_post_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_post_Params_Promise is a wrapper for a WebSession_post_Params promised by a client call.
type WebSession_post_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_post_Params_Promise) Struct() (WebSession_post_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_post_Params{s}, err
}

func (p WebSession_post_Params_Promise) Content() WebSession_PostContent_Promise {
	return WebSession_PostContent_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_post_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_openWebSocket_Params struct{ capnp.Struct }

func NewWebSession_openWebSocket_Params(s *capnp.Segment) (WebSession_openWebSocket_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_openWebSocket_Params{st}, err
}

func NewRootWebSession_openWebSocket_Params(s *capnp.Segment) (WebSession_openWebSocket_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_openWebSocket_Params{st}, err
}

func ReadRootWebSession_openWebSocket_Params(msg *capnp.Message) (WebSession_openWebSocket_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_openWebSocket_Params{root.Struct()}, err
}

func (s WebSession_openWebSocket_Params) String() string {
	str, _ := text.Marshal(0xc7c9c9b19d935e79, s.Struct)
	return str
}

func (s WebSession_openWebSocket_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_openWebSocket_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_openWebSocket_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_openWebSocket_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_openWebSocket_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_openWebSocket_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_openWebSocket_Params) Protocol() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSession_openWebSocket_Params) HasProtocol() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) SetProtocol(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewProtocol sets the protocol field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSession_openWebSocket_Params) NewProtocol(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s WebSession_openWebSocket_Params) ClientStream() WebSession_WebSocketStream {
	p, _ := s.Struct.Ptr(3)
	return WebSession_WebSocketStream{Client: p.Interface().Client()}
}

func (s WebSession_openWebSocket_Params) HasClientStream() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) SetClientStream(v WebSession_WebSocketStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(3, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(3, in.ToPtr())
}

// WebSession_openWebSocket_Params_List is a list of WebSession_openWebSocket_Params.
type WebSession_openWebSocket_Params_List struct{ capnp.List }

// NewWebSession_openWebSocket_Params creates a new list of WebSession_openWebSocket_Params.
func NewWebSession_openWebSocket_Params_List(s *capnp.Segment, sz int32) (WebSession_openWebSocket_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return WebSession_openWebSocket_Params_List{l}, err
}

func (s WebSession_openWebSocket_Params_List) At(i int) WebSession_openWebSocket_Params {
	return WebSession_openWebSocket_Params{s.List.Struct(i)}
}

func (s WebSession_openWebSocket_Params_List) Set(i int, v WebSession_openWebSocket_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_openWebSocket_Params_Promise is a wrapper for a WebSession_openWebSocket_Params promised by a client call.
type WebSession_openWebSocket_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_openWebSocket_Params_Promise) Struct() (WebSession_openWebSocket_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_openWebSocket_Params{s}, err
}

func (p WebSession_openWebSocket_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_openWebSocket_Params_Promise) ClientStream() WebSession_WebSocketStream {
	return WebSession_WebSocketStream{Client: p.Pipeline.GetPipeline(3).Client()}
}

type WebSession_openWebSocket_Results struct{ capnp.Struct }

func NewWebSession_openWebSocket_Results(s *capnp.Segment) (WebSession_openWebSocket_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_openWebSocket_Results{st}, err
}

func NewRootWebSession_openWebSocket_Results(s *capnp.Segment) (WebSession_openWebSocket_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_openWebSocket_Results{st}, err
}

func ReadRootWebSession_openWebSocket_Results(msg *capnp.Message) (WebSession_openWebSocket_Results, error) {
	root, err := msg.RootPtr()
	return WebSession_openWebSocket_Results{root.Struct()}, err
}

func (s WebSession_openWebSocket_Results) String() string {
	str, _ := text.Marshal(0xcc561276d31b392b, s.Struct)
	return str
}

func (s WebSession_openWebSocket_Results) Protocol() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSession_openWebSocket_Results) HasProtocol() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Results) SetProtocol(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewProtocol sets the protocol field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSession_openWebSocket_Results) NewProtocol(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s WebSession_openWebSocket_Results) ServerStream() WebSession_WebSocketStream {
	p, _ := s.Struct.Ptr(1)
	return WebSession_WebSocketStream{Client: p.Interface().Client()}
}

func (s WebSession_openWebSocket_Results) HasServerStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Results) SetServerStream(v WebSession_WebSocketStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// WebSession_openWebSocket_Results_List is a list of WebSession_openWebSocket_Results.
type WebSession_openWebSocket_Results_List struct{ capnp.List }

// NewWebSession_openWebSocket_Results creates a new list of WebSession_openWebSocket_Results.
func NewWebSession_openWebSocket_Results_List(s *capnp.Segment, sz int32) (WebSession_openWebSocket_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSession_openWebSocket_Results_List{l}, err
}

func (s WebSession_openWebSocket_Results_List) At(i int) WebSession_openWebSocket_Results {
	return WebSession_openWebSocket_Results{s.List.Struct(i)}
}

func (s WebSession_openWebSocket_Results_List) Set(i int, v WebSession_openWebSocket_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_openWebSocket_Results_Promise is a wrapper for a WebSession_openWebSocket_Results promised by a client call.
type WebSession_openWebSocket_Results_Promise struct{ *capnp.Pipeline }

func (p WebSession_openWebSocket_Results_Promise) Struct() (WebSession_openWebSocket_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_openWebSocket_Results{s}, err
}

func (p WebSession_openWebSocket_Results_Promise) ServerStream() WebSession_WebSocketStream {
	return WebSession_WebSocketStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

type WebSession_put_Params struct{ capnp.Struct }

func NewWebSession_put_Params(s *capnp.Segment) (WebSession_put_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_put_Params{st}, err
}

func NewRootWebSession_put_Params(s *capnp.Segment) (WebSession_put_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_put_Params{st}, err
}

func ReadRootWebSession_put_Params(msg *capnp.Message) (WebSession_put_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_put_Params{root.Struct()}, err
}

func (s WebSession_put_Params) String() string {
	str, _ := text.Marshal(0xf1c587295608596e, s.Struct)
	return str
}

func (s WebSession_put_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_put_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_put_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_put_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_put_Params) Content() (WebSession_PutContent, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_PutContent{Struct: p.Struct()}, err
}

func (s WebSession_put_Params) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_put_Params) SetContent(v WebSession_PutContent) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContent sets the content field to a newly
// allocated WebSession_PutContent struct, preferring placement in s's segment.
func (s WebSession_put_Params) NewContent() (WebSession_PutContent, error) {
	ss, err := NewWebSession_PutContent(s.Struct.Segment())
	if err != nil {
		return WebSession_PutContent{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_put_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_put_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_put_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_put_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_put_Params_List is a list of WebSession_put_Params.
type WebSession_put_Params_List struct{ capnp.List }

// NewWebSession_put_Params creates a new list of WebSession_put_Params.
func NewWebSession_put_Params_List(s *capnp.Segment, sz int32) (WebSession_put_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_put_Params_List{l}, err
}

func (s WebSession_put_Params_List) At(i int) WebSession_put_Params {
	return WebSession_put_Params{s.List.Struct(i)}
}

func (s WebSession_put_Params_List) Set(i int, v WebSession_put_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_put_Params_Promise is a wrapper for a WebSession_put_Params promised by a client call.
type WebSession_put_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_put_Params_Promise) Struct() (WebSession_put_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_put_Params{s}, err
}

func (p WebSession_put_Params_Promise) Content() WebSession_PutContent_Promise {
	return WebSession_PutContent_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_put_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_delete_Params struct{ capnp.Struct }

func NewWebSession_delete_Params(s *capnp.Segment) (WebSession_delete_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_delete_Params{st}, err
}

func NewRootWebSession_delete_Params(s *capnp.Segment) (WebSession_delete_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_delete_Params{st}, err
}

func ReadRootWebSession_delete_Params(msg *capnp.Message) (WebSession_delete_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_delete_Params{root.Struct()}, err
}

func (s WebSession_delete_Params) String() string {
	str, _ := text.Marshal(0xeba76bffb27b1975, s.Struct)
	return str
}

func (s WebSession_delete_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_delete_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_delete_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_delete_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_delete_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_delete_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_delete_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_delete_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_delete_Params_List is a list of WebSession_delete_Params.
type WebSession_delete_Params_List struct{ capnp.List }

// NewWebSession_delete_Params creates a new list of WebSession_delete_Params.
func NewWebSession_delete_Params_List(s *capnp.Segment, sz int32) (WebSession_delete_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSession_delete_Params_List{l}, err
}

func (s WebSession_delete_Params_List) At(i int) WebSession_delete_Params {
	return WebSession_delete_Params{s.List.Struct(i)}
}

func (s WebSession_delete_Params_List) Set(i int, v WebSession_delete_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_delete_Params_Promise is a wrapper for a WebSession_delete_Params promised by a client call.
type WebSession_delete_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_delete_Params_Promise) Struct() (WebSession_delete_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_delete_Params{s}, err
}

func (p WebSession_delete_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type WebSession_postStreaming_Params struct{ capnp.Struct }

func NewWebSession_postStreaming_Params(s *capnp.Segment) (WebSession_postStreaming_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_postStreaming_Params{st}, err
}

func NewRootWebSession_postStreaming_Params(s *capnp.Segment) (WebSession_postStreaming_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_postStreaming_Params{st}, err
}

func ReadRootWebSession_postStreaming_Params(msg *capnp.Message) (WebSession_postStreaming_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_postStreaming_Params{root.Struct()}, err
}

func (s WebSession_postStreaming_Params) String() string {
	str, _ := text.Marshal(0xd26a7affce43b1c0, s.Struct)
	return str
}

func (s WebSession_postStreaming_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_postStreaming_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_postStreaming_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_postStreaming_Params) MimeType() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_postStreaming_Params) HasMimeType() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_postStreaming_Params) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_postStreaming_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_postStreaming_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_postStreaming_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_postStreaming_Params) Encoding() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s WebSession_postStreaming_Params) HasEncoding() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s WebSession_postStreaming_Params) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

// WebSession_postStreaming_Params_List is a list of WebSession_postStreaming_Params.
type WebSession_postStreaming_Params_List struct{ capnp.List }

// NewWebSession_postStreaming_Params creates a new list of WebSession_postStreaming_Params.
func NewWebSession_postStreaming_Params_List(s *capnp.Segment, sz int32) (WebSession_postStreaming_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return WebSession_postStreaming_Params_List{l}, err
}

func (s WebSession_postStreaming_Params_List) At(i int) WebSession_postStreaming_Params {
	return WebSession_postStreaming_Params{s.List.Struct(i)}
}

func (s WebSession_postStreaming_Params_List) Set(i int, v WebSession_postStreaming_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_postStreaming_Params_Promise is a wrapper for a WebSession_postStreaming_Params promised by a client call.
type WebSession_postStreaming_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_postStreaming_Params_Promise) Struct() (WebSession_postStreaming_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_postStreaming_Params{s}, err
}

func (p WebSession_postStreaming_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_postStreaming_Results struct{ capnp.Struct }

func NewWebSession_postStreaming_Results(s *capnp.Segment) (WebSession_postStreaming_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_postStreaming_Results{st}, err
}

func NewRootWebSession_postStreaming_Results(s *capnp.Segment) (WebSession_postStreaming_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_postStreaming_Results{st}, err
}

func ReadRootWebSession_postStreaming_Results(msg *capnp.Message) (WebSession_postStreaming_Results, error) {
	root, err := msg.RootPtr()
	return WebSession_postStreaming_Results{root.Struct()}, err
}

func (s WebSession_postStreaming_Results) String() string {
	str, _ := text.Marshal(0xbf0e0653dc266205, s.Struct)
	return str
}

func (s WebSession_postStreaming_Results) Stream() WebSession_RequestStream {
	p, _ := s.Struct.Ptr(0)
	return WebSession_RequestStream{Client: p.Interface().Client()}
}

func (s WebSession_postStreaming_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Results) SetStream(v WebSession_RequestStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSession_postStreaming_Results_List is a list of WebSession_postStreaming_Results.
type WebSession_postStreaming_Results_List struct{ capnp.List }

// NewWebSession_postStreaming_Results creates a new list of WebSession_postStreaming_Results.
func NewWebSession_postStreaming_Results_List(s *capnp.Segment, sz int32) (WebSession_postStreaming_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSession_postStreaming_Results_List{l}, err
}

func (s WebSession_postStreaming_Results_List) At(i int) WebSession_postStreaming_Results {
	return WebSession_postStreaming_Results{s.List.Struct(i)}
}

func (s WebSession_postStreaming_Results_List) Set(i int, v WebSession_postStreaming_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_postStreaming_Results_Promise is a wrapper for a WebSession_postStreaming_Results promised by a client call.
type WebSession_postStreaming_Results_Promise struct{ *capnp.Pipeline }

func (p WebSession_postStreaming_Results_Promise) Struct() (WebSession_postStreaming_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_postStreaming_Results{s}, err
}

func (p WebSession_postStreaming_Results_Promise) Stream() WebSession_RequestStream {
	return WebSession_RequestStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSession_putStreaming_Params struct{ capnp.Struct }

func NewWebSession_putStreaming_Params(s *capnp.Segment) (WebSession_putStreaming_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_putStreaming_Params{st}, err
}

func NewRootWebSession_putStreaming_Params(s *capnp.Segment) (WebSession_putStreaming_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return WebSession_putStreaming_Params{st}, err
}

func ReadRootWebSession_putStreaming_Params(msg *capnp.Message) (WebSession_putStreaming_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_putStreaming_Params{root.Struct()}, err
}

func (s WebSession_putStreaming_Params) String() string {
	str, _ := text.Marshal(0xa1ece076a7105939, s.Struct)
	return str
}

func (s WebSession_putStreaming_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_putStreaming_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_putStreaming_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_putStreaming_Params) MimeType() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_putStreaming_Params) HasMimeType() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_putStreaming_Params) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_putStreaming_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_putStreaming_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_putStreaming_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_putStreaming_Params) Encoding() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s WebSession_putStreaming_Params) HasEncoding() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s WebSession_putStreaming_Params) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

// WebSession_putStreaming_Params_List is a list of WebSession_putStreaming_Params.
type WebSession_putStreaming_Params_List struct{ capnp.List }

// NewWebSession_putStreaming_Params creates a new list of WebSession_putStreaming_Params.
func NewWebSession_putStreaming_Params_List(s *capnp.Segment, sz int32) (WebSession_putStreaming_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return WebSession_putStreaming_Params_List{l}, err
}

func (s WebSession_putStreaming_Params_List) At(i int) WebSession_putStreaming_Params {
	return WebSession_putStreaming_Params{s.List.Struct(i)}
}

func (s WebSession_putStreaming_Params_List) Set(i int, v WebSession_putStreaming_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_putStreaming_Params_Promise is a wrapper for a WebSession_putStreaming_Params promised by a client call.
type WebSession_putStreaming_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_putStreaming_Params_Promise) Struct() (WebSession_putStreaming_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_putStreaming_Params{s}, err
}

func (p WebSession_putStreaming_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_putStreaming_Results struct{ capnp.Struct }

func NewWebSession_putStreaming_Results(s *capnp.Segment) (WebSession_putStreaming_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_putStreaming_Results{st}, err
}

func NewRootWebSession_putStreaming_Results(s *capnp.Segment) (WebSession_putStreaming_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSession_putStreaming_Results{st}, err
}

func ReadRootWebSession_putStreaming_Results(msg *capnp.Message) (WebSession_putStreaming_Results, error) {
	root, err := msg.RootPtr()
	return WebSession_putStreaming_Results{root.Struct()}, err
}

func (s WebSession_putStreaming_Results) String() string {
	str, _ := text.Marshal(0xc60d14bf989d4454, s.Struct)
	return str
}

func (s WebSession_putStreaming_Results) Stream() WebSession_RequestStream {
	p, _ := s.Struct.Ptr(0)
	return WebSession_RequestStream{Client: p.Interface().Client()}
}

func (s WebSession_putStreaming_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Results) SetStream(v WebSession_RequestStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSession_putStreaming_Results_List is a list of WebSession_putStreaming_Results.
type WebSession_putStreaming_Results_List struct{ capnp.List }

// NewWebSession_putStreaming_Results creates a new list of WebSession_putStreaming_Results.
func NewWebSession_putStreaming_Results_List(s *capnp.Segment, sz int32) (WebSession_putStreaming_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSession_putStreaming_Results_List{l}, err
}

func (s WebSession_putStreaming_Results_List) At(i int) WebSession_putStreaming_Results {
	return WebSession_putStreaming_Results{s.List.Struct(i)}
}

func (s WebSession_putStreaming_Results_List) Set(i int, v WebSession_putStreaming_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_putStreaming_Results_Promise is a wrapper for a WebSession_putStreaming_Results promised by a client call.
type WebSession_putStreaming_Results_Promise struct{ *capnp.Pipeline }

func (p WebSession_putStreaming_Results_Promise) Struct() (WebSession_putStreaming_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_putStreaming_Results{s}, err
}

func (p WebSession_putStreaming_Results_Promise) Stream() WebSession_RequestStream {
	return WebSession_RequestStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSession_propfind_Params struct{ capnp.Struct }

func NewWebSession_propfind_Params(s *capnp.Segment) (WebSession_propfind_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_propfind_Params{st}, err
}

func NewRootWebSession_propfind_Params(s *capnp.Segment) (WebSession_propfind_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_propfind_Params{st}, err
}

func ReadRootWebSession_propfind_Params(msg *capnp.Message) (WebSession_propfind_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_propfind_Params{root.Struct()}, err
}

func (s WebSession_propfind_Params) String() string {
	str, _ := text.Marshal(0xca2d58de88f0b32e, s.Struct)
	return str
}

func (s WebSession_propfind_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_propfind_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_propfind_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_propfind_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_propfind_Params) XmlContent() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_propfind_Params) HasXmlContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_propfind_Params) XmlContentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_propfind_Params) SetXmlContent(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_propfind_Params) Depth() WebSession_PropfindDepth {
	return WebSession_PropfindDepth(s.Struct.Uint16(0))
}

func (s WebSession_propfind_Params) SetDepth(v WebSession_PropfindDepth) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s WebSession_propfind_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_propfind_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_propfind_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_propfind_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_propfind_Params_List is a list of WebSession_propfind_Params.
type WebSession_propfind_Params_List struct{ capnp.List }

// NewWebSession_propfind_Params creates a new list of WebSession_propfind_Params.
func NewWebSession_propfind_Params_List(s *capnp.Segment, sz int32) (WebSession_propfind_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return WebSession_propfind_Params_List{l}, err
}

func (s WebSession_propfind_Params_List) At(i int) WebSession_propfind_Params {
	return WebSession_propfind_Params{s.List.Struct(i)}
}

func (s WebSession_propfind_Params_List) Set(i int, v WebSession_propfind_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_propfind_Params_Promise is a wrapper for a WebSession_propfind_Params promised by a client call.
type WebSession_propfind_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_propfind_Params_Promise) Struct() (WebSession_propfind_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_propfind_Params{s}, err
}

func (p WebSession_propfind_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_proppatch_Params struct{ capnp.Struct }

func NewWebSession_proppatch_Params(s *capnp.Segment) (WebSession_proppatch_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_proppatch_Params{st}, err
}

func NewRootWebSession_proppatch_Params(s *capnp.Segment) (WebSession_proppatch_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_proppatch_Params{st}, err
}

func ReadRootWebSession_proppatch_Params(msg *capnp.Message) (WebSession_proppatch_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_proppatch_Params{root.Struct()}, err
}

func (s WebSession_proppatch_Params) String() string {
	str, _ := text.Marshal(0x9e582e7e054088ae, s.Struct)
	return str
}

func (s WebSession_proppatch_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_proppatch_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_proppatch_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_proppatch_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_proppatch_Params) XmlContent() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_proppatch_Params) HasXmlContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_proppatch_Params) XmlContentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_proppatch_Params) SetXmlContent(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_proppatch_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_proppatch_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_proppatch_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_proppatch_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_proppatch_Params_List is a list of WebSession_proppatch_Params.
type WebSession_proppatch_Params_List struct{ capnp.List }

// NewWebSession_proppatch_Params creates a new list of WebSession_proppatch_Params.
func NewWebSession_proppatch_Params_List(s *capnp.Segment, sz int32) (WebSession_proppatch_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_proppatch_Params_List{l}, err
}

func (s WebSession_proppatch_Params_List) At(i int) WebSession_proppatch_Params {
	return WebSession_proppatch_Params{s.List.Struct(i)}
}

func (s WebSession_proppatch_Params_List) Set(i int, v WebSession_proppatch_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_proppatch_Params_Promise is a wrapper for a WebSession_proppatch_Params promised by a client call.
type WebSession_proppatch_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_proppatch_Params_Promise) Struct() (WebSession_proppatch_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_proppatch_Params{s}, err
}

func (p WebSession_proppatch_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_mkcol_Params struct{ capnp.Struct }

func NewWebSession_mkcol_Params(s *capnp.Segment) (WebSession_mkcol_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_mkcol_Params{st}, err
}

func NewRootWebSession_mkcol_Params(s *capnp.Segment) (WebSession_mkcol_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_mkcol_Params{st}, err
}

func ReadRootWebSession_mkcol_Params(msg *capnp.Message) (WebSession_mkcol_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_mkcol_Params{root.Struct()}, err
}

func (s WebSession_mkcol_Params) String() string {
	str, _ := text.Marshal(0xf64da2416445f8b6, s.Struct)
	return str
}

func (s WebSession_mkcol_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_mkcol_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_mkcol_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_mkcol_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_mkcol_Params) Content() (WebSession_PostContent, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_PostContent{Struct: p.Struct()}, err
}

func (s WebSession_mkcol_Params) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_mkcol_Params) SetContent(v WebSession_PostContent) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContent sets the content field to a newly
// allocated WebSession_PostContent struct, preferring placement in s's segment.
func (s WebSession_mkcol_Params) NewContent() (WebSession_PostContent, error) {
	ss, err := NewWebSession_PostContent(s.Struct.Segment())
	if err != nil {
		return WebSession_PostContent{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_mkcol_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_mkcol_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_mkcol_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_mkcol_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_mkcol_Params_List is a list of WebSession_mkcol_Params.
type WebSession_mkcol_Params_List struct{ capnp.List }

// NewWebSession_mkcol_Params creates a new list of WebSession_mkcol_Params.
func NewWebSession_mkcol_Params_List(s *capnp.Segment, sz int32) (WebSession_mkcol_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_mkcol_Params_List{l}, err
}

func (s WebSession_mkcol_Params_List) At(i int) WebSession_mkcol_Params {
	return WebSession_mkcol_Params{s.List.Struct(i)}
}

func (s WebSession_mkcol_Params_List) Set(i int, v WebSession_mkcol_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_mkcol_Params_Promise is a wrapper for a WebSession_mkcol_Params promised by a client call.
type WebSession_mkcol_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_mkcol_Params_Promise) Struct() (WebSession_mkcol_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_mkcol_Params{s}, err
}

func (p WebSession_mkcol_Params_Promise) Content() WebSession_PostContent_Promise {
	return WebSession_PostContent_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_mkcol_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_copy_Params struct{ capnp.Struct }

func NewWebSession_copy_Params(s *capnp.Segment) (WebSession_copy_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_copy_Params{st}, err
}

func NewRootWebSession_copy_Params(s *capnp.Segment) (WebSession_copy_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_copy_Params{st}, err
}

func ReadRootWebSession_copy_Params(msg *capnp.Message) (WebSession_copy_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_copy_Params{root.Struct()}, err
}

func (s WebSession_copy_Params) String() string {
	str, _ := text.Marshal(0x8139673a82bfe07d, s.Struct)
	return str
}

func (s WebSession_copy_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_copy_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_copy_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_copy_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_copy_Params) Destination() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_copy_Params) HasDestination() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_copy_Params) DestinationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_copy_Params) SetDestination(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_copy_Params) NoOverwrite() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_copy_Params) SetNoOverwrite(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_copy_Params) Shallow() bool {
	return s.Struct.Bit(1)
}

func (s WebSession_copy_Params) SetShallow(v bool) {
	s.Struct.SetBit(1, v)
}

func (s WebSession_copy_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_copy_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_copy_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_copy_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_copy_Params_List is a list of WebSession_copy_Params.
type WebSession_copy_Params_List struct{ capnp.List }

// NewWebSession_copy_Params creates a new list of WebSession_copy_Params.
func NewWebSession_copy_Params_List(s *capnp.Segment, sz int32) (WebSession_copy_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return WebSession_copy_Params_List{l}, err
}

func (s WebSession_copy_Params_List) At(i int) WebSession_copy_Params {
	return WebSession_copy_Params{s.List.Struct(i)}
}

func (s WebSession_copy_Params_List) Set(i int, v WebSession_copy_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_copy_Params_Promise is a wrapper for a WebSession_copy_Params promised by a client call.
type WebSession_copy_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_copy_Params_Promise) Struct() (WebSession_copy_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_copy_Params{s}, err
}

func (p WebSession_copy_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_move_Params struct{ capnp.Struct }

func NewWebSession_move_Params(s *capnp.Segment) (WebSession_move_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_move_Params{st}, err
}

func NewRootWebSession_move_Params(s *capnp.Segment) (WebSession_move_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_move_Params{st}, err
}

func ReadRootWebSession_move_Params(msg *capnp.Message) (WebSession_move_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_move_Params{root.Struct()}, err
}

func (s WebSession_move_Params) String() string {
	str, _ := text.Marshal(0x81f5066b5576a609, s.Struct)
	return str
}

func (s WebSession_move_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_move_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_move_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_move_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_move_Params) Destination() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_move_Params) HasDestination() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_move_Params) DestinationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_move_Params) SetDestination(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_move_Params) NoOverwrite() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_move_Params) SetNoOverwrite(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_move_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_move_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_move_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_move_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_move_Params_List is a list of WebSession_move_Params.
type WebSession_move_Params_List struct{ capnp.List }

// NewWebSession_move_Params creates a new list of WebSession_move_Params.
func NewWebSession_move_Params_List(s *capnp.Segment, sz int32) (WebSession_move_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return WebSession_move_Params_List{l}, err
}

func (s WebSession_move_Params_List) At(i int) WebSession_move_Params {
	return WebSession_move_Params{s.List.Struct(i)}
}

func (s WebSession_move_Params_List) Set(i int, v WebSession_move_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_move_Params_Promise is a wrapper for a WebSession_move_Params promised by a client call.
type WebSession_move_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_move_Params_Promise) Struct() (WebSession_move_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_move_Params{s}, err
}

func (p WebSession_move_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_lock_Params struct{ capnp.Struct }

func NewWebSession_lock_Params(s *capnp.Segment) (WebSession_lock_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_lock_Params{st}, err
}

func NewRootWebSession_lock_Params(s *capnp.Segment) (WebSession_lock_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return WebSession_lock_Params{st}, err
}

func ReadRootWebSession_lock_Params(msg *capnp.Message) (WebSession_lock_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_lock_Params{root.Struct()}, err
}

func (s WebSession_lock_Params) String() string {
	str, _ := text.Marshal(0x9398280f1359570a, s.Struct)
	return str
}

func (s WebSession_lock_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_lock_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_lock_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_lock_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_lock_Params) XmlContent() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_lock_Params) HasXmlContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_lock_Params) XmlContentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_lock_Params) SetXmlContent(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_lock_Params) Shallow() bool {
	return s.Struct.Bit(0)
}

func (s WebSession_lock_Params) SetShallow(v bool) {
	s.Struct.SetBit(0, v)
}

func (s WebSession_lock_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_lock_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_lock_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_lock_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_lock_Params_List is a list of WebSession_lock_Params.
type WebSession_lock_Params_List struct{ capnp.List }

// NewWebSession_lock_Params creates a new list of WebSession_lock_Params.
func NewWebSession_lock_Params_List(s *capnp.Segment, sz int32) (WebSession_lock_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return WebSession_lock_Params_List{l}, err
}

func (s WebSession_lock_Params_List) At(i int) WebSession_lock_Params {
	return WebSession_lock_Params{s.List.Struct(i)}
}

func (s WebSession_lock_Params_List) Set(i int, v WebSession_lock_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_lock_Params_Promise is a wrapper for a WebSession_lock_Params promised by a client call.
type WebSession_lock_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_lock_Params_Promise) Struct() (WebSession_lock_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_lock_Params{s}, err
}

func (p WebSession_lock_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_unlock_Params struct{ capnp.Struct }

func NewWebSession_unlock_Params(s *capnp.Segment) (WebSession_unlock_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_unlock_Params{st}, err
}

func NewRootWebSession_unlock_Params(s *capnp.Segment) (WebSession_unlock_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_unlock_Params{st}, err
}

func ReadRootWebSession_unlock_Params(msg *capnp.Message) (WebSession_unlock_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_unlock_Params{root.Struct()}, err
}

func (s WebSession_unlock_Params) String() string {
	str, _ := text.Marshal(0xd684c6a791b97dbc, s.Struct)
	return str
}

func (s WebSession_unlock_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_unlock_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_unlock_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_unlock_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_unlock_Params) LockToken() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_unlock_Params) HasLockToken() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_unlock_Params) LockTokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_unlock_Params) SetLockToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_unlock_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_unlock_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_unlock_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_unlock_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_unlock_Params_List is a list of WebSession_unlock_Params.
type WebSession_unlock_Params_List struct{ capnp.List }

// NewWebSession_unlock_Params creates a new list of WebSession_unlock_Params.
func NewWebSession_unlock_Params_List(s *capnp.Segment, sz int32) (WebSession_unlock_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_unlock_Params_List{l}, err
}

func (s WebSession_unlock_Params_List) At(i int) WebSession_unlock_Params {
	return WebSession_unlock_Params{s.List.Struct(i)}
}

func (s WebSession_unlock_Params_List) Set(i int, v WebSession_unlock_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_unlock_Params_Promise is a wrapper for a WebSession_unlock_Params promised by a client call.
type WebSession_unlock_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_unlock_Params_Promise) Struct() (WebSession_unlock_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_unlock_Params{s}, err
}

func (p WebSession_unlock_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_acl_Params struct{ capnp.Struct }

func NewWebSession_acl_Params(s *capnp.Segment) (WebSession_acl_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_acl_Params{st}, err
}

func NewRootWebSession_acl_Params(s *capnp.Segment) (WebSession_acl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_acl_Params{st}, err
}

func ReadRootWebSession_acl_Params(msg *capnp.Message) (WebSession_acl_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_acl_Params{root.Struct()}, err
}

func (s WebSession_acl_Params) String() string {
	str, _ := text.Marshal(0x9f79c33e20119e8d, s.Struct)
	return str
}

func (s WebSession_acl_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_acl_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_acl_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_acl_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_acl_Params) XmlContent() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSession_acl_Params) HasXmlContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_acl_Params) XmlContentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSession_acl_Params) SetXmlContent(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSession_acl_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_acl_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_acl_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_acl_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_acl_Params_List is a list of WebSession_acl_Params.
type WebSession_acl_Params_List struct{ capnp.List }

// NewWebSession_acl_Params creates a new list of WebSession_acl_Params.
func NewWebSession_acl_Params_List(s *capnp.Segment, sz int32) (WebSession_acl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_acl_Params_List{l}, err
}

func (s WebSession_acl_Params_List) At(i int) WebSession_acl_Params {
	return WebSession_acl_Params{s.List.Struct(i)}
}

func (s WebSession_acl_Params_List) Set(i int, v WebSession_acl_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_acl_Params_Promise is a wrapper for a WebSession_acl_Params promised by a client call.
type WebSession_acl_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_acl_Params_Promise) Struct() (WebSession_acl_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_acl_Params{s}, err
}

func (p WebSession_acl_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_report_Params struct{ capnp.Struct }

func NewWebSession_report_Params(s *capnp.Segment) (WebSession_report_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_report_Params{st}, err
}

func NewRootWebSession_report_Params(s *capnp.Segment) (WebSession_report_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_report_Params{st}, err
}

func ReadRootWebSession_report_Params(msg *capnp.Message) (WebSession_report_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_report_Params{root.Struct()}, err
}

func (s WebSession_report_Params) String() string {
	str, _ := text.Marshal(0xc0643ea68efc60ae, s.Struct)
	return str
}

func (s WebSession_report_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_report_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_report_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_report_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_report_Params) Content() (WebSession_PostContent, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_PostContent{Struct: p.Struct()}, err
}

func (s WebSession_report_Params) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_report_Params) SetContent(v WebSession_PostContent) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContent sets the content field to a newly
// allocated WebSession_PostContent struct, preferring placement in s's segment.
func (s WebSession_report_Params) NewContent() (WebSession_PostContent, error) {
	ss, err := NewWebSession_PostContent(s.Struct.Segment())
	if err != nil {
		return WebSession_PostContent{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_report_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_report_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_report_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_report_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_report_Params_List is a list of WebSession_report_Params.
type WebSession_report_Params_List struct{ capnp.List }

// NewWebSession_report_Params creates a new list of WebSession_report_Params.
func NewWebSession_report_Params_List(s *capnp.Segment, sz int32) (WebSession_report_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_report_Params_List{l}, err
}

func (s WebSession_report_Params_List) At(i int) WebSession_report_Params {
	return WebSession_report_Params{s.List.Struct(i)}
}

func (s WebSession_report_Params_List) Set(i int, v WebSession_report_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_report_Params_Promise is a wrapper for a WebSession_report_Params promised by a client call.
type WebSession_report_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_report_Params_Promise) Struct() (WebSession_report_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_report_Params{s}, err
}

func (p WebSession_report_Params_Promise) Content() WebSession_PostContent_Promise {
	return WebSession_PostContent_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_report_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type WebSession_options_Params struct{ capnp.Struct }

func NewWebSession_options_Params(s *capnp.Segment) (WebSession_options_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_options_Params{st}, err
}

func NewRootWebSession_options_Params(s *capnp.Segment) (WebSession_options_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSession_options_Params{st}, err
}

func ReadRootWebSession_options_Params(msg *capnp.Message) (WebSession_options_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_options_Params{root.Struct()}, err
}

func (s WebSession_options_Params) String() string {
	str, _ := text.Marshal(0xd2e47e8eac54ea7e, s.Struct)
	return str
}

func (s WebSession_options_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_options_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_options_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_options_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_options_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_options_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_options_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_options_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_options_Params_List is a list of WebSession_options_Params.
type WebSession_options_Params_List struct{ capnp.List }

// NewWebSession_options_Params creates a new list of WebSession_options_Params.
func NewWebSession_options_Params_List(s *capnp.Segment, sz int32) (WebSession_options_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSession_options_Params_List{l}, err
}

func (s WebSession_options_Params_List) At(i int) WebSession_options_Params {
	return WebSession_options_Params{s.List.Struct(i)}
}

func (s WebSession_options_Params_List) Set(i int, v WebSession_options_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_options_Params_Promise is a wrapper for a WebSession_options_Params promised by a client call.
type WebSession_options_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_options_Params_Promise) Struct() (WebSession_options_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_options_Params{s}, err
}

func (p WebSession_options_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type WebSession_patch_Params struct{ capnp.Struct }

func NewWebSession_patch_Params(s *capnp.Segment) (WebSession_patch_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_patch_Params{st}, err
}

func NewRootWebSession_patch_Params(s *capnp.Segment) (WebSession_patch_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return WebSession_patch_Params{st}, err
}

func ReadRootWebSession_patch_Params(msg *capnp.Message) (WebSession_patch_Params, error) {
	root, err := msg.RootPtr()
	return WebSession_patch_Params{root.Struct()}, err
}

func (s WebSession_patch_Params) String() string {
	str, _ := text.Marshal(0xadef95edc22ca880, s.Struct)
	return str
}

func (s WebSession_patch_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_patch_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_patch_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_patch_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_patch_Params) Content() (WebSession_PostContent, error) {
	p, err := s.Struct.Ptr(1)
	return WebSession_PostContent{Struct: p.Struct()}, err
}

func (s WebSession_patch_Params) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_patch_Params) SetContent(v WebSession_PostContent) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContent sets the content field to a newly
// allocated WebSession_PostContent struct, preferring placement in s's segment.
func (s WebSession_patch_Params) NewContent() (WebSession_PostContent, error) {
	ss, err := NewWebSession_PostContent(s.Struct.Segment())
	if err != nil {
		return WebSession_PostContent{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s WebSession_patch_Params) Context() (WebSession_Context, error) {
	p, err := s.Struct.Ptr(2)
	return WebSession_Context{Struct: p.Struct()}, err
}

func (s WebSession_patch_Params) HasContext() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_patch_Params) SetContext(v WebSession_Context) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewContext sets the context field to a newly
// allocated WebSession_Context struct, preferring placement in s's segment.
func (s WebSession_patch_Params) NewContext() (WebSession_Context, error) {
	ss, err := NewWebSession_Context(s.Struct.Segment())
	if err != nil {
		return WebSession_Context{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// WebSession_patch_Params_List is a list of WebSession_patch_Params.
type WebSession_patch_Params_List struct{ capnp.List }

// NewWebSession_patch_Params creates a new list of WebSession_patch_Params.
func NewWebSession_patch_Params_List(s *capnp.Segment, sz int32) (WebSession_patch_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return WebSession_patch_Params_List{l}, err
}

func (s WebSession_patch_Params_List) At(i int) WebSession_patch_Params {
	return WebSession_patch_Params{s.List.Struct(i)}
}

func (s WebSession_patch_Params_List) Set(i int, v WebSession_patch_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_patch_Params_Promise is a wrapper for a WebSession_patch_Params promised by a client call.
type WebSession_patch_Params_Promise struct{ *capnp.Pipeline }

func (p WebSession_patch_Params_Promise) Struct() (WebSession_patch_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_patch_Params{s}, err
}

func (p WebSession_patch_Params_Promise) Content() WebSession_PostContent_Promise {
	return WebSession_PostContent_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p WebSession_patch_Params_Promise) Context() WebSession_Context_Promise {
	return WebSession_Context_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

const schema_a8cb0f2f1a756b32 = "x\xda\xcc;\x0dtT\xe5\x95\xef{3\x93I\x900" +
	"y\xbc\xf8\x93\x84t\x12M\x80D&\x92\x84\xf4\x94\x9c" +
	"\xd5\x90\x84\x08d\x09d&\x89@V-o2\x8fd" +
	"\x92\x99y\xc3\xbc7$\x93\x8a\x91\xa8+Z\xb5\x05\xa1" +
	"\x15\x0e\xf8[\x05\x11\xd6_X\xedJ\x15W\xda\x86\x8a" +
	"\xae\xecZ\xeb\xee\xda\xd5c=\xbbZ\xa9?\xabV*" +
	"\xf8\xf6\xde\xef\xcd\xfbI\x1c&\x03\x9e\xed\xe1\x1c=I" +
	"\xeew\xbf\xfb\xdd\xef\xfe\xdf\xfb=\xe6\xd6\xe7-`\xab" +
	"\x1d#\x85\x0c\xd31\x8f8\xb2\xd4\xf5o??Z\xdf" +
	";\x7f\x03\xe3-&\x84a\x1c6'\xc3\xd4\x96\xe7\xb6" +
	"\x12\x86\xd4V\xe7\xba\xe1\x87:\xfb\xed\xdc/\x1c]\x1d" +
	"\x1b\x18\xae\x82\xfd\xf3\xcd\x8b\xff\xfe\xdd\xd2\x1f<\x05\xab" +
	"K\xa6\x0d\x13\xfe\x9ai\x80\xcf\xaf\x9a6\x02x7\xee" +
	"h\xda\x1f\xda{\x97FJ}ew]\xdb\xfd\x9c\xf3" +
	"!\xa6%\xc7\x99EH\xed\xb6i5\x84\x7fd\xda\"" +
	" \xff\xce4\x15\xe9\xe6<\xbc\xaek \xeb\xf3\xf1G" +
	"o\xc8\xc3\xa3\xf9;\xf2\xfe\x9b!\xff\xf5\xa3\x8d\xc7^" +
	"v\xf2\xa3\xdcL\xd6$\x0e'\x07\xb9\xcd\x84\xdf\xc0\xe1" +
	"\xc9\xeb\xb9\xc7\x80\xd2\x83\xefN\xd9s\xc9\x9f\x1e\x1ce" +
	"8\x0f\x10\xb2#\x1d\x0f\xff\x17\xc2\xd8\xd5W\xb7|\xc9" +
	"\xae=\xa0\xdc\xc8xg\x12\x0b\x87\xe7\xe78\x01\xb1\xb6" +
	"\x94\x1f\xc5\xc3<\xfc \x10i\xf9\xce\xc9\x7f\xff\xe1\xbf" +
	"}\xf8\xc3\xd4\xa8w\xf0\xc3\x88\xba\x8d\xc7\xf3\xc8r\xfb" +
	"\xd2\xff(z\xf0Nf\"g\x97\xe7?I\xf8\xae|" +
	"d\xc0\x9b\xdf\x80\x97\x9c\xb2b\x15\xef\x9a}\xf7]\xe3" +
	".\xf9\xfa\x85\xf4\x92\xef\\\x08\x97T\xaf\xc8?o\xe6" +
	"\x1b\xbb\x96\xff\x94\xf1\x96\xc3\xb9\xbd\xbfz\xf4\xc3K\xdf" +
	";\xf29\xd3\xe5p\x12\x07\xe0>{\xd1\x93\xa8\x8b\xb1" +
	"\x8b\xa8.\xee\xed\xdf\xfe\xc5\xff\xbe\xa8nc\xb8b\x9b" +
	")bXw\x14v\x13\xbe\xa0\x10er~\xe1\"\xbe" +
	"\xb1\xf0B\x86Q\x7fw\xfb\xbe\x93\xef\xceY\xbb]\x93" +
	"\x8b\x83\xe0\xd9\xd5\x85\x9f\xe2\xd9-\x85\x0d@\xee\x1f6" +
	".p\\_\xb5\xf2\x1e g0'\x16\xc6\x10am" +
	"!\xde4\xeb\x82C\x9b\xdbg\xee\x02\x84\xe9v\xb5f" +
	" ^x\x99\xeb7\xbb\xf1\xbc\xdc\xa2V\xc2\x97\x16\xc1" +
	"\x86\x8e\x19E6\xd21\xbb\x88\x05\x0a\xea\x1d\xf7p%" +
	"W\xfcs\xe2^+\xc1\x82\xa2&$XZ\x84\x04\xe7" +
	"\xaf\xca\xdb\xb5\xee\xed\x0f\xefO\"P]\x1d,\xba\x0d" +
	"\x11\x8e\x16\xa18\x8cKq\xd3m\xd6\x13\xf9Gf|" +
	"\xc0\xef\x9f\xf1\x06\xe0\x87\x8bo\xb1\xf3\xde\x12\xbca3" +
	"{\xe0\xcd_\x88\xebw\xa3xYs\xafvpc\xc9" +
	"\xc5\x04\xf0\xf0\xd7\xb6\x12*\xbeSG\x0e\xf9\xa7|\x1a" +
	"\xd9ce\xaf\xa5\x94*\xc3[\x8a\xec\x9d8t\x9d\xbf" +
	"\x88=\xb1g\x82\x0dk\xa2{\xbf\xd4G\xf8S\xa5(" +
	"\xe3\x13\xa5h2\xbb\xbfJ\x9cz\xfc\xea\x97\x1eMm" +
	"2\xab.FA\xd6\x8a\x17\xff\x08\x8f\xbea\xf7\x9c\x17" +
	"\x8fo\xfdh\x9f\xf5\xe8\xb1K|x\xf4\xeb\x97\xe0\xd1" +
	"\xb7l|44u\xca\xe2\xc7\x18\xeftG\xf6\xb8\x9b" +
	"\x87\xcb>\xe0\x13e(j\xa5\xcc\x06\x1bt'\x1c\xcf" +
	"#\x15\xa5X\x06jI\x94\x81h\xf8\x0de(\xce\xa2" +
	"\x9f=\xb5\xad\xe5\xa5\x1dO\xc1\xb1\xec8\x83\x11\xcb\xc1" +
	"`\x12\xe5x\x99x\xf9r\xc0\xbc\xf6\x95\xfd\xc3\xf9\xf2" +
	"3O\x83Q[\xee\xe2`\x91\xec\xfe\xf2~\xc2\x8fQ" +
	"\xe4\xc3\xe5x\xf3\x8a\x9b\x94\xdf\xbc\xfc\xd9{\x80\\n" +
	"\xb1Y\x0d\xb9t\xa6\x9f\xf0u3\x11\xb9z&\"_" +
	"q\xa0\xbfqo\xd5\xef\xfe\x11\xaf>QGw\xcc\x04" +
	"\x86\xef\xa7\xc8;g\xa2\x18\xfe\xd8seg\xf4\x7f\xf6" +
	"\xfe\x1c(\xdbM\xca\xe8[\xb3\xf6\x80o\xcdB1\xb4" +
	"\xcf\x02\x8b\x0b\xcd\xa2\x16\xf7\xfbB\xf9\xc1\x1b\xeb\xea\x9e" +
	"\x03\xb1\x814\x0c\xb1i\x1a\xdb?\xab\x9e\xf0\x87q\x0f" +
	"\x7fh\x16\xb22\xb8xh\xc7\x81\xbf\x1b8\xc8x\xcb" +
	"\xc0^\x0c\xeb\xe9\xb29\x89\x0d\xf0\x8bg\xfbQ!\x15" +
	"\xb3\x91\x13\x87\x7f\xe6\x7fvdM{>\xa91Jp" +
	"l\xf6vDxs6\xf5\x9e\xd5'\xef|\xf8\x8a\xc0" +
	"\x0bV\x95\x9e\x98\xdd\x8d\x08\x8e\x0a\xa4\xd0\xb9p\xe7\xdd" +
	"\xcf\xe7\xe7\xfe\xd2Jag\xc5fD\xd8W\x81\x14\x12" +
	"\xd7\xde\xb5\xf3\x89\xb1\xb1_Y\xbd\xe1\xa8\x86\xf0V\x05" +
	"\xaa\xaf\xea\xa9\x8f7\xfe~\xa5\xe7\xc8\xb8\xf01V\xd9" +
	"O\xcd\xa6\x121.\x9d_\xf4\xaf\xeb\xa6_\xf5r\x92" +
	"\x04\xd5\xc0\xbeK)\x97\x07/\xc5+\xcf\xbeq\xf0\xb3" +
	"\xb7\xf6n9\xaa\x93\xa0\x18\xd5s\xa8O\xce\x9f\x83l" +
	"\xae\x14\x0f=\xfb\xdeys\x8e\xa54\xfa\xd7\xe7L'" +
	"\xfc{sP\x84\xef\xccAz/<\xd1\xfc\x8a:\xdc" +
	"\x7f\xcc\xca\xf3\xe5\x1e\xcas\x9b\x079\xba\xfe\x83\xce\xbd" +
	"w^\xff\x87cV\x8e*\xaa\xa8`\xeb\xaa\x90\xc2s" +
	"\xeb\x9f\xdd\xb4\xeb\x977\xfd\xd6*\xb7MUTn;" +
	"\xab\x90\xa1\x9f<\xfe\xc5\x8a\x9f\x169\xdeHe0\xf3" +
	"/\x03\xa7^r\x19\xf2\xd3r\x19\"\xd7l\x12\xe7\x7f" +
	"\xfd\xc9c)\x91O]\xd6D\xf8\xdc\xb9\x88\x9c3\x17" +
	"\x91\x17>s\xc3u\x81\xc7G\xdeL\xed\xb1\x8f\xcc\xa5" +
	"\xd7xv.r\x19=x\xb0\xa0\xa6\xf2\xbe\xf7\x19\xee" +
	"R=\xa9\x14W\x7f\x8aI\xe5\xb7\xf2\xb2\x0f\x7f\x91\xa5" +
	"\xbe\x9fR^\\u%\xc4\xbaj<\xb2\xb8\x1a\xc5\x11" +
	"/\xf8\xc1\x93\xea\xc0\xae?Z\xc5\xc1\xd5\xd0\xdb\x16\xd7" +
	"\xe0A/\xde\xba\x83\x7f\x7f\xcam\xc7\x19o\x15\xb1\xc4" +
	"\x94.\xe0\x89\x05\xdcD\x0dU\xe6\xad\x14wt\xef\x9c" +
	"\x03\x8d-\x17~\x04Nj\x1f\x97y\xde\xab\x81\xccs" +
	"\xaa\x06\xbd\xe3\xcb\x1a\xf0\x8e\xfcZ\xea\x1d\xce\xd6\x15\xd7" +
	"\x16\x9f$\x1f\x7f#_\x94\xd6\x82G\xd7\xd5R'\xad" +
	"]\xc4_\x83\xbf\xa9cG\xb7\xe7\x10\xef\xa2\x8fSK" +
	"\xa7\xa5\x96J\xa7\xab\x16-7\xb2*\xfb\xaa\x8a[\x0e" +
	"\x7fb\xd5a\xa2\x96\x1a\xd5\x86Z\x94\xb4\xcb\xb5\xfa\xce" +
	"\x19\x1f9>KM\xab`\xde\x11\x9ay\xe7!-\xc3" +
	"\xcb'\xc8\xd3A\xf3\xe8<\x90\xa70o\x16\xd2\x9fG" +
	"\x83\xf8\x81/[\x02\x8d\x0f\xb4}a=\xfa\x89:\x1a" +
	"I\x0f\xd6\xe1\xd1\xab\xfe\xe9;\x97\x9f\xa8\xf5\xfc9\x95" +
	"@\x89\xad\xb6\xeb\xbb\xc7\x107\xf8\xddA\xe6\x98:(" +
	"\xfa=\xb2(\xcb\xf6\xa0\x14\xa9\xea\x11\xa2\x91h\xfd\x0a" +
	"\xd1\xdf\x01\x10\x0a\x90\xa2\x89\xb2v!&\x84e\x86\xf1" +
	"\xe6\xdb\xec`\x08p&\xb7\xbe\x12\xfe\x1c\xb2\x11\xefM" +
	",\xe1\x08\xc9G\xc7\xe26\xf8\x01x\x03\x00og\x09" +
	"a\xf3Qy\xdc\xad\x08\xdb\x08\xb0-\x80h\x03D\x88" +
	"2\xdc\xa6&\x00\xde\x0e\xc0\xbb\x01h\x07L \xcbm" +
	"E\xe0\x8f\x01\xb8\x83%\xae\xa8\xa0\xf4\x91\xa9\x0c\x0b\xff" +
	"\x135 \xcaJ0\"(\x8c\x13X2\xa0\x11i\xf9" +
	":16\x18\x03\xa8\"\xc2\xf9p&CF\xe4>!" +
	"\x14\x92\x06\x8d\xbf{\xa4\x88\"\x0e)$\xcf\x1aJI" +
	"\x1eF\xc2\xb4\x17o\x16z\xfa\xc4v)\x14\xecI4" +
	"Tu\x80\x18\xc4vB\xbc\xda\x9d\xea*\x91\x06\xe7\xe9" +
	"\x86\x1f,W\x01|\x13\x1bW\x1a\x83\x1fv\xae\xd8\xc7" +
	"0\xae\x88\x14\x11\xd5\xa8\x18\xa3\xc4\x18\x9b\x14\x19\x81?" +
	"\xbad1\x86\xc0\xc6h\xf4*\x91q\xc7\xf0\x185\x1e" +
	"\x09\xc2%d\x81!!\x83![J\x86|\xa2\xec\x8e" +
	"J\x11Y\xf4\xe6\x11\xb30,\xf0[j\xb1\x82QK" +
	"\x0e+\xa8\xb7\xf8\xca\xf9\xa3jG\xbc\xa7\x07H53" +
	"N) \xaa\xcd\xa1\xa0\x18QZb$&\xc5\x9a\x01" +
	"\xc04,\x16\x85\x000\xd8G\x7f\xac\xe8# \xd4P" +
	"PV O\x18J?\x0e\x17\xf6~\x08\xbe\x95\x0df" +
	"\x94\xebPUb)\x1ay\x07ie\xd8\\\xf6k\xd5" +
	"jr\xdc'M\x00\xcc:\x85@#\xf4po\xfa\x01" +
	"\xe8<\x89@\xc3\xe3\xb8\xc3\x08\xb4}\x85@\xa3h\xe5" +
	"\x9e\xf01,G\xb2\xf2Ac\x0c\xb7\x13mi\x07X" +
	"\xc8n8\xde\xfe\x17\xc44\x9c\x8c{h;`fg" +
	"\xe7\x93|\xb4\xa5\xcd\x80y7`>\xcd\x12U\x16\x95" +
	"fI\x1a\x00E\x882\x99\xc6\x90v\x1b\x18\x80\x99\xfa" +
	"@\x95\x00Tcb \x18\x13{\x14\x86a4\xa3\x89" +
	"(jORJ\x8c\x13\xc4\x04tb\xa0+\xfd\xaf\x88" +
	"\xd4L\xb1\x18\x02xIc\x01S\xecI\x90<\xb3G" +
	"\xa0\x86\x16\x05\xb2R$\x10$\x0a\xe8\xf1J!\x18\xb2" +
	"\x89\x01U\x08\x04\x82\xf87\x11B\x9a\xe4\x89\x857C" +
	"\x8bI\xde\xd2\x9bjXZ'\x9a>\x9ag\xa8K@" +
	"\x1f\xbd\x1a\x84\xd0g\xf1Q\x11e\x18\x00`\xd4\xf4\xd1" +
	"0\xc2B\x00\x1bB\x1fe5\x1f\x8d\xa3;F\x01x" +
	"\xdd\xb7q\xc7\xc9\xdd\xcfq:k\xa7\xc6^\xa5\x9b\xad" +
	"\x14 \xd4\x03gP\x8e\xdb\x0a\xb967\xb8\xe0\xda&" +
	"n\xad\x1b|pk+\xb7\xcd\x0dNx\xc8\xc7\x1dv" +
	"\x83\x17\x1e\x1f\xe6>q\x13\x07_@\xfc|1q\x93" +
	",\xbe\x05~[B\xdc6i\xc0k\x87\xc2\xc7\xac5" +
	"\xed,i\xbc\x08Yb\x80<\xf95G\x0am\xcb\xff" +
	"\x16m &\x0a\x8a\x18H\x8b=\xc6\x91\xa6\x91f\x0d" +
	"\x11\xb2\x87\x00\xacF\xe9\xafL\x8am3\x8cmG8" +
	"\xd2\xaa6\x1a\xc8\xc4jKi7\xbe\xcc\x91nu\x99" +
	"TB\xb1\x19[\x04\x8cU\x8d\x0a1%(\x84\x9a\x99" +
	"\x06\xcdh\xd3\x12x\x85#\xa3j\xbb\xb6\xa3\x84h\x87" +
	"\xa2\xc5\xab\xe1xH\x09v(\x02\xe3T\xe2rZ\x12" +
	"\xafr\xa4_mCtO\x87\xc2\xb8\x04\xc0\x87\xfd\x11" +
	"Ii\x93\x02\xc15\xa0\xfd\x9423\xf6\xdb\xe6\x12$" +
	"\xb0LRJp\x03\xe3Z\x13\xa4\xb2\xd3\xcd!+\xa5" +
	"9\xe0\xafR\xcf\x80\xa8t( \xedp\x95,F\x02" +
	"M\x09E\x94\xcb|\x0d\xa2\x0c\xcc\xc8\x99\xda\x93.j" +
	"\x08l\xd9\xe8)\xf9\xc4\x05\xf6T1\x0a*\x9b\x0d\xc6" +
	">/\xe9*<\x00\xab\xd1\x7f\xe6\x00\xf0{\x18D\xfa" +
	"\xa4x(\xe0\x13\x09D\x93+\xa5X\x98\xd1M\xdc%" +
	"v\x0a\xbd`\xdfF\x0d\x99Qz1\xf8\x89\x89\x0dZ" +
	"\xdc\xf1N\xd5\xd8A'mA\x87\\\x08'\xb7\x03;" +
	",\xb0s\x1e\xda|+\x00\x97\x02pe\x92\xc7\xa9\x00" +
	"\xecB\xccN\x00\xae\x06\x1e\x83r\xbb\x18\x0b\x0b\x11\xc6" +
	"\x09\xf7\xd3\xf9SCR\x8f\x80\x91\x06\xd4l8\xab<" +
	"\x18Tz\xfa:%\xc6\xb9H413\x14\xa1\x9e;" +
	"b\x92\x9b\xe6\x0et\xcby\xd4-\xb7v\xa3\x1b\x92\xa4" +
	"\x1b\xb2\xdc\xf1VtC\x1b\xb8\xe1m|)\xb8\xa1\x9d" +
	"o#1\xdeK\xd05\x13\xa4\x95_O]\xf3!R" +
	"\x09\xff\xbbI6?F\x9e\xe4_\x83\xdfrx\xc2n" +
	"\xe6sX7\x99\xc2\xd7\xb1{\xf8\xcb\xe1\xb7\xf3\xf8 " +
	"\xeb\xe3\xc3\xf0\x9b\x93\xdf\xcan\xe7w\xc2oS\xf91" +
	"\xf6\x01\xfe5\xd6\xad\xfa\x05\xd0\xcd\xda88\x85\x9c\xd6" +
	"\x03l?\x06\xd9\xf9\xd5&!P\x82\xf8 )\x19\x9d" +
	"h\x8d\x14\xf3\x07\x03\x90\x06I$\xed\xee\xbb`\xb7O" +
	"\xbd\xd2\xc4\xd6L\xffJ)\x1e\x99\xcc\xebm[\xe8^" +
	"4|Dg\x08Z}XT\xfa\xa4\xc02\x89(\x8d" +
	"X\xb5\xa0'\xa4\xa0Qb\xd0\xd8\x0a4\xb6\xabmt" +
	"W\xc92\x02\xb4\xe8>\x1bu!`\x84\x06\x15\xc6\xad" +
	"\x08\xfe\x90\x98\x96\x99\x9f\x00\xa1a\xca\x0cna\x1a\xa2" +
	"t\x0b\x10\x81\x80\xbd\x06\x8a\x1ee\xb2\xdbl#\x18\xc4" +
	"\x9a\x0dl\xe2\xea\x85\xa2']\xb4\xb4m\x87-\x95\xae" +
	"E\x80\x06\xe7\xc4P[\xb2\xd2B\"JPItJ" +
	"\x92{\xa9\x10\xebME\xa0\xcc \xb0\x13\x08\x1cQ}" +
	"\xda\xce\x12\xb6\x85n-\x81\xbd%t/M\xe2t\xad" +
	"\x8b\xc4\x82\x00^\x0a\xe9\xa97\xad<\xef\x01\x8a{t" +
	"\x8a\x1e\xd2\xe5[B\xc9\xb9`g/0\x19\x8f\xc8\xf1" +
	"hT\x8a\x11\x88\xd3mP\x1f\x08\xae\xceDTLK" +
	"\xf1^\xa0\xf8s\xb5\xcb\xb2\xb1\x04w6\x08%\xb8\x15" +
	"h\x06\xc3\x8d\x9d\xa2\x10\x95&\x89\xf5\xb6\x07h\x9c\\" +
	"2+\\\"\x94@\xb0w\xc1\x16\x1a\xeeS\x14\x13i" +
	"\x19\xda\x01t\x1eP\xdb\xad\xdbJ`\x9f3Dm&" +
	"\x1e\x89\xc6$H\xb0DF\xfd\x83D!{'\xd2\xd2" +
	"{\x98\x8a\xack\xdc\xbe\x92\x96\x88\x0bua\x89\xe3\xa9" +
	"\xc3\x1e\xc4\xa2\x81\xccK\x95\xee\x14\xa5\x0aV%}\x00" +
	"S,\xa5\xca\xda&\xb3~\x19_\xaa\x0c\x85CF\xae" +
	"\xd4\x81g\xd1\"\xa4\x0e\x88\xcd\xda\xbe*L\x00I\xf9" +
	"\xbaiM\x87}\xd2TU\xd5\xda\x1ak\xa7\x94K\xbe" +
	"\x06\xa8\x03[\xa5z\x80^\x07\xd0\x8d,)\xb6\x9dR" +
	"\xed\xf9$\x0b\xc07\xf7\x03\xf8&\x00\xdf\x07`\xfbI" +
	"\x04;\xb1\xf0\x8d%\x0b\xdfg\x80\x06\xfb\x15\xd0\xc8\x06" +
	"\xe8~\x8c\xfdO\x03\xf4\x05\xb86\xb6\x1cLV\x838" +
	"\x045\xbb\xccd\xa9a\x01\x02\xbc(/\x87fD\\" +
	"\xbe\xc6,-\xadi\x0aK\xcb$\xde2\xc6-\xa5G" +
	"\x0cH\xa2\x0cQ\x7f\x082<\x84\xcc\xacI\xf3\x1bu" +
	")L\xd76!\x8c\x19\xc2ns\xa0\xf9&g\x0aD" +
	"\xef\x879\x0e\xca~.\xc7\xa9\xf6\x8a\x0a\xcd0\x8c\x13" +
	"r\xcc\x02\xc0\x87\xea\xfe\xfb\xef\xbeZ1\xf8\xbd\x15G" +
	"\x99oQ \xb4\xbb\xa9\xb9\xc1\xf9\xba\xb1\xe5\xa2\xb9d" +
	"\x83\xdc\xf2Y2\x12\x06\x02B\xafHr\xc1\x18r'" +
	"M\xdb`\xf4\xd1(J\xcc ;\xd5 \xdb\x82\x8a^" +
	"\x00d\x97Zlx\x09\xda\xf0b\x00vb&O\x1a" +
	"\xb1\xb7\xc9\xcc\xe4\x19\xd8\xeb\x19\xb5\xb0}\x8a\x12\x85:" +
	"\x0e\xaa\xb2\xc6\x08$\x04\x9a\xf9\x97\x04\xb4\xd2-\x07\x08" +
	"\xe6\x98\x8e=\xa9\xb7\x0a=!\xddY\xcf\x8d{\xa6\xf6" +
	"\xc3h<\xa9\xf5`\xa47\xf3\xe0\xd2j\x06\x17\x83\xe1" +
	"\xf1\xd1\xc5\x96\x8c.\xad\xa7k\x84\xc2\xc1\xb0\xa8Eu" +
	"&\x83K\x88\x91\x1e(x1\xaf\x98u\x98~1v" +
	"\xe2\xc5\\x3\xa8\xaa\xac\xa3@\xa1\xde\xa4\xc9]\xd3" +
	"d\x99*\xaf\xf2[\x86\x80\xab\xbaM\xdf\xe5VU\x9a" +
	"\x8d.\xd7Uo\x8e\xf79o\xbf9\x93\xe2\xbc\xad\x96" +
	"\x87\x15o\xcc25\xf3\x8e\xea\xdd,\xe7\xf5\x9b\x83?" +
	"\xd0\xade\xb2\xde\x16k\xd0\xe4>\x92\x8c\x89j\xbb$" +
	"+\xf8;-E\xd5\xf6\xb8b\xa8\xdb\xd5\x02\xe1\xb2\x81" +
	"6\xe4\xa2\xd9\x04\xd1\xc4\xaa\xea%&\x9a\xa6\x1eA\xa0" +
	"\x9aA\xdd\xaa\xba\x87\xb3Ie3\xfa\x84\x06{\xee\x91" +
	"\xe5Q4t\x19\xf2\x9c\x14]\x13\x84b\xcc\xbd\x10\x08" +
	"\xf7\x81Ib\xd4\xd1'\xc0F\xd4\xe1\xab\xd9\x8b\x19\x96" +
	"/g\x9d\xc4|\x131W\x0b\xd8JX\xcd\x85U\xa7" +
	"1\xa1&\xfa\x9c\x19j\xd3\x18\xac\x9e N\xc8\x92\xfa" +
	"\x14\xd0\xdc{\x9c \xe5w`\xd5fL>\xcd\xd5\xd7" +
	"I=\xac\x8e\xc1\xaa\xc3\x98#\x13}\xce\xce\x1f$H" +
	"y?\xacf\x19\xefDD\x9f\xa1\xf3\x8f\x90~X\xbd" +
	"\x1fV\xb3\x8d\xb1\xb8Iy+\x8e_\xf8;`5\xc7" +
	"x\xd52W7\x10\x1f\xac&`u\x8a1>4W" +
	"\xc3\xa4\x06V\x05X=\xcfx\x0f5W\xbb\x08Jc" +
	"\x09\xacN5\x9e,\xcd\xd5\xcb\xe9j5\xac\xe6\x1ao" +
	"}\xe6j9]-\x80\xd5i\xc6\xd4\xdb\\\xcd\xa5\xd2" +
	" \xb0\xea2\x1e\xce\xcc\xcc\xf09\x08\x92{\xdf\x09\xae" +
	"\xa4\xbf3\x98ko\xc1F\xee5'\xe1\x8ca;\xd1" +
	"m\x93;\xdc\x04k\xcf:\x89\xddxr2\xf7\xed\x83" +
	"\x8br\xf7;\x9d\x90m\x16\x807\x83\x95. \xaa\x14" +
	"\x15#h^\x90\x031\x85, N\x08)\x0bHC" +
	"@\x0c\x89\x0a\xe4\"\x15\xf1\xd0\xea\x187\x0d2\x08I" +
	"\xc6\x1c\xc6\xa5\x03t\xc3c\x98\xe4_\x98'\x18\xd2\xb7" +
	"\x80\xb8\xc3\x03=R\x08\x8e\xc3a*\xfc\xc0y\x0d\xfc" +
	"\xc0Z\x08\x0e\x89G\xb4_\x9c\x10m\xe1\xcf\x98\x08e" +
	"#\x1c>\"i\x16\x0d\xdbi\xc2IfD\xe9\x82\xfd" +
	"+\xff\xd4X\xf3\x855x\xdbNS\x9dHNp1" +
	":\xb5\x9cdl[c\xd6\"\x04\xfe3\xdf\x88\xb8\x9b" +
	"Q\x94\xb6<\xadf\x09\xb7Z\xc6D8\xb5\xc5\x8a%" +
	"^i\x89\x8e\x11!,\xea\xb1\xcd\xbdN\x08\xc5\x8d\xbf" +
	"F\xc4\xa1(\xb4\xb9\xb2\x8a\xf9iy$\x94\xc0(\xa8" +
	"w\xd0\xe3\x82\xea$\x09\x184a\xc6\xf8\xc9\x92R\x93" +
	"\xb5\x8dN\xc6\xf8\xb6&3S\x8d;Y\x9f\xfa\x81\xbd" +
	"\x19\x91U\x0b\xddg=D\xd6\xc3\x1b\x047mz\x95" +
	"m0\\\xd1jN\x1c(\xbf\xd06W\xd7[\x06\x0e" +
	")\x92K\xc3\xda\xab\xa8L\xa70\xec\xa1)74d" +
	">d\xe8\x91\xdc\xdalh\x86\xcdN\xa5\x83\x82\xd8\xdf" +
	"j\x96\x8fT:\x98\xec\x0e\"\xf09\x00\xfe:\x99\x01" +
	"\xb1\x84=\x8c\xc0\x97\x00\xf8/\xc0\xab\x83X\x1el\xb8" +
	"\xa3\xe0\xdey\xb0w\x0a\x12\xec6\x09\x92,by\x86" +
	"\xe0\x0eb\x91gw\xe6\x13\x0e\xf0\x1eA]\xfd\x0c\xf0" +
	"\x1egSg\xc5\x90\x10\xe9\x8dCQf\x85\xa5\x10\x87" +
	"\xcb/\x05\x12\xaaL\x8b\x9df\x89\xb1\x05D\xe2\xd2\x07" +
	"\xe3\xa0\x1e\x17\x16\xadA\x90\x80\x1c\xd4F\x95g7\xa9" +
	"I\x96{hrD>\x87,\x8e\x9dX\xf3\xd1Q\x1c" +
	"\x96yy\xe6\x93\xf0\x99<s\xd0\x86\xdf\xfa\xb1\x0b\xa9" +
	"q\xd3\x87\x0fk)\xe5\x03\xac\xd5p\x95\x90\xe5\xd2A" +
	"\x9f\xa5jbK\xb4K\xaf\x1d6\xe3\x02g+\xd5\xac" +
	"+1l\x86\x1fu0\xa8\xf45\xf7\x89\x10#\x07 " +
	"\xfc\x1b\xc7&5\x17\xa5\x93,:\x12M\xb1\xbaN\x88" +
	"\x05\xa1\xb9\x890\xc9B\xc2\x18a\x99\x0b\x9a\xef}c" +
	"\xb6\x95Z\x0cz\xd9\x00E\x83M\xe9CO\x9dJo" +
	"Q\xdc\xca\x95\xe2@\xab\xb1\x92>\xf5\xcc\xbf\x18\x1b\xf8" +
	"\x08`\xd2fW\x9b\x90\xbc\xb2\xe9\x93\xaf\x13\xbb\x02\x9f" +
	"j\x8d2\x88\x04\xe7r\xc4g\xe2\x91\xef3\x8ckX" +
	"\x8cIN|\x16\xca\xd0e\x17\x8b.|\x15\x98\x104" +
	"\xd0\xe6\xca@~s-\xe2\xf7\xd4\x98\x91$M\x04\x9e" +
	"\xcc\x10\x92\xcd\xecbQp\xfe5\xcf\xd5\x0bE:\x87" +
	"\xa6\x827=\xac\xd5\xe2L\xfa\xa9Vg2\x1b\x0d\xc4" +
	"\x84\xb6\xd5{5\x9b\xa6*\x07G\xd3\xbb\xbct\xb5x" +
	"\xfaf_\x7f\x1f\x0b*\xa2\x0b\xdf\xc7\xd0\xe1\x92=3" +
	"\x12\x9a\x86\xba\xbf\x82+\xd8\xcc\x95\x8er\xe51\xae\xa2" +
	"\x9b\xf3\xf8\xb8\xean\xae\xce\xa7\x0eyd!\x12\x90\x15" +
	"\"\xc5\xc2\x1e!\x1a\xf5\x90JU\xea\xf1(\xd0\xa5\x85" +
	"XOH\x8c\xf4*}\x0c\x00z\xfa\xe2\x91\x01\xb7G" +
	"\x0e\x0e\x8b\xb0\x07\x00a%\xc8\xd8\xc2\"\"\xaf\x09\x86" +
	"\xc4 C\x02\xaa\x8e\xc7\xe0c\xd2\x90\xa7\xafW\x88\xf5" +
	"z\x18\xa0\x98\xa2*Xl\xf4\x84\x0bE\xb9\xa7!\x16" +
	"\x8c*\xd2D\x0d\x17\x9a\x1aN\xa5`[0@\x9cp" +
	"?'\xa8W\x09*\xa1\xcc\xcd\x0a\xdd\xb3\x0a\xb3\xbf\x13" +
	"\xd2?j\x97NF\xd8\x09\x11\xb4\x98|\xad&\x1f\x87" +
	"\x97\xb4\x9a\x1a.fO!\xd8>A\xc7\xc9i\x87*" +
	"\xf8e)\x14W\xa8\xaa\x1d\xc0\x8f\x83\x8e\x00C\xd0\xf4" +
	"\xae\xa30\xda\xf0N\xde?&\xab;\xad\x81\xf4\xe1{" +
	"\x82M\x1970\xa87\x07\x06\x0d2\xc5\x84\xc2\xd3h" +
	"\x98\xf0-xRAhu\x1d\xcd#\xb6\xf0\xb9\x94G" +
	"2\xe8\xa9\xa9H\x14r\xf6\"I}\x86^v\xd3\xa2" +
	"\xdb\xc8\xb0\x93\xf5\xedM\x96\x0c\xa4K&\xd8\x9a\xcc@" +
	"7Y\xfa\xf6\x0d\xfd\xe6\xd7\x08)\xc4\x95R(P\xb7" +
	"+\x12\x14\xebh<\xe3\xbc\x9a$\xdf\x82\xf1\xc9+y" +
	"Y\xa3G\xceH\xffz\x7f\x00\xf7t\xd1\xc1\xd1\xd9\x0c" +
	"?k\xbe\xf5\xf0\xd3\x1d\xc0\xb6\x182\xaa\xd1\xbck\x19" +
	"\xf5\xac\x0de\xbc\x12\x0d\xe7IQ\xe8.\xb4\xdc\xae\xb1" +
	"\xdf4\xfatB\xd7\x9e\xdc\xcfZ\xe8\xbd\xa2\x92\xf1\x00" +
	"\xcb\xe2q\xba\xbc\xdb\xbaO3\xbfJ#\xad`oD" +
	"\x8a\x89MX\x8e&\xbeQ|\xa4\xee\xd4Z:\x05\xd2" +
	";!\x1c\xd7\xa4\x08\xc7\x95f8\x1e\x9fb]\x83\xa2" +
	"0\x90\xe1+\xde\xf8`\x97\xa9\xd3\x9dS\xc32\xfbi" +
	"\x0c\x91\xf6\xcbx%'\xfa\xd7d\xd5K\x93\xa5z\xc9" +
	"P\xb7\xe9\x19\xd0\x1a\xf9\x8c#\xbc\xef\x8c\x06\xa6H\xba" +
	"S\x1a\xc0'\xc73`4\xb5\xbd%%\x94I\xb1\xe5" +
	"\xb3<9\xebLv=\x00\xc0\x95\x9a\xdeU\xbf \x8b" +
	"\xed\xc0\xa8UOq\xf0\xda\xc6^\xadt\xd7a\xda\xa7" +
	"\x11\x82\x9f\x84\xc4\xa5\xb4\xbbs\xf6\x9a\x1f\xe3\xe8\xde>" +
	"I\xa5\x187\x0a\xc5s\xb8N4{n\xe3\x95\x9c\xc4" +
	"\xf4O\x0d\xf0-\xa7\xa2{\xc2\xa7\x06\xf8\x94S\x8d\xdf" +
	"\x1f\xcc\x05\xe0\xdf\xe0\xa7\x06\xe3:Z\xf3\xfb.\xbd\xa7" +
	"\x85\x1a\x0eK8\x02\xc7-V\xc2\xa1or\x965\xd9" +
	"\x93\x8c\x10\xae\xd2\x9f[\x80U\xfd-c\x12\xb3Y\x1e" +
	"U\\\xe8a(\xfcd\xbc\x80\x18\xfc\xcd\x8e\x10\x80\xe3" +
	";B\x14>~\xab\xe33C\x03I\x86\x8bD,\xd9" +
	"\x10n\x81{\x07\x84u\xcd!A\x96\x19Rm\x043" +
	"\x13V\x93\x02Vk\x85\xb5\x0c\xe1\x174n\xe4\xf5\x0c" +
	"-K\x9b\xf3\x19\x9e\xfb\xff\x189&\xb3\x19\xcd\xfc\xaa" +
	"\xfc\x12d\x0f\xe4c\x86\x8aO\x80\x8e\x09\x19\x01\x8b\xe5" +
	"\xe4\x83\xa1\xa7\xde\x92\x14\xfc\xf8\xe0\xa5\xdb\xaeY\x95\xbd" +
	"\xdd\xb4z\xf5\xde\xb2\xcf\xee\xce\xac*3\x981;\x1d" +
	"\xb7x\xbaN\xe7\x02\xce\xb1\x99\xcb\xedN\xd9\xda\x8ck" +
	"]\xd2+\xc0\xf2zG\x13\xbd\xf9^\xa8\xff\x03\x0e\xa2" +
	"\xff\x0b\x17\x8e\xf3i\xef\x85\xfa\x0b\x1fC\xe4\x05\xa4\x9d" +
	"d|-\xfd+>\xea\x97P\xd0\xe6\x81Zs\xb0\xa2" +
	"\x1d\xa5\x03E\xe2\xbd\x88\xcd\xc4\xc7\xec\xa7+\x9b\xcf\xa6" +
	"\xe08\xa3\x12\xdfx\xe9\xc9\xb8\xc4\xcfJ/\x11\xfd\x93" +
	"\x02\xfd\x8b\x02\xed+\x14M2\xd3Q2\xfd\x16\xc9\xd0" +
	"\xf7b\x8c\x88\xf4\x1d'\xcd4\xedtS\xe9\x88\x0b\xb9" +
	"\xf5f\x13\xeb\xbf\xa1\xc8\xa9\xb7\xfc\xb3\x07\xc7\xe8\xe9?" +
	"F\xb5\x8c\xb2\x9b\xcc\xd0a\x08u\xd3\xb0\xf9\x0d\xb1!" +
	"\xd4m\xe8#[\xb4y#gOF\x9e}\xf8\x85\xe8" +
	"\xe3\x00|\x19\xa3\x11\xb1\xfc+$n\xec6\x86\x05\xa9" +
	"b\xe3j\xf9&\xb3l\xde\x13\x8f\xfea\xc3}[\xcc" +
	"\xefE\x93oU\x0d\x1d\xba\xa3\x99\xcf\xd5\xd4\xd1\x1a\xb4" +
	"\xa4g\x920\x1e\xdf\x92$\xd2~\x03j\x08'\x89\xac" +
	"\x7fe@tUM\xfa|K\x9f\x1d\xce\xc1\xc9eV" +
	"fAP\x9f\xdd\x06%\x12\xc1XH\xe7\x06\xb9\x18\x0a" +
	"\xeb\xc7\x87BG>\x0dF\x9eVsz\xde\x00Ex" +
	"X\x081Yj@\x1a\x84\xbaL\x08X2\xf8\xff\x05" +
	"\x00\x00\xff\xff\x921ve"

func init() {
	schemas.Register(schema_a8cb0f2f1a756b32,
		0x8139673a82bfe07d,
		0x81535505f60de028,
		0x8193ac6cb5429c83,
		0x81f5066b5576a609,
		0x82130007ccd2888f,
		0x82a3ee23aa0ae3a3,
		0x8374b67102f894cf,
		0x8becd48bdafc1e45,
		0x8ea31bdb4c044f01,
		0x9398280f1359570a,
		0x974fa7d7260b143e,
		0x99ffc2f3f69a6a9f,
		0x9a712ce3fcad8cd8,
		0x9e582e7e054088ae,
		0x9ea7265092c11606,
		0x9f79c33e20119e8d,
		0xa1ece076a7105939,
		0xa50711a14d35a8ce,
		0xa87d65bed9b60243,
		0xaa6ef20a62c1cafd,
		0xaaf9021b627cc1f9,
		0xabc45cb0fd79fba8,
		0xadef95edc22ca880,
		0xaf480a0c6cab8887,
		0xb37b21e300864885,
		0xb39cc44599b3a41b,
		0xb4b873147ab5ce5e,
		0xb4e5f4cccb748429,
		0xb7d82eac416ab63e,
		0xbaace870544663eb,
		0xbc353583a3731ade,
		0xbd6b5bb69c784877,
		0xbf0e0653dc266205,
		0xc0643ea68efc60ae,
		0xc60d14bf989d4454,
		0xc7c9c9b19d935e79,
		0xca2d58de88f0b32e,
		0xcc561276d31b392b,
		0xcd94acddf4778328,
		0xd22c0be5b9c16558,
		0xd26a7affce43b1c0,
		0xd2e47e8eac54ea7e,
		0xd684c6a791b97dbc,
		0xd7051b9757f6b096,
		0xd7aff1fe39659132,
		0xd97fb0647c80b844,
		0xe9a02a3219bdbd70,
		0xe9ff06beec4e73d6,
		0xeba76bffb27b1975,
		0xed8a0ae9139c89c2,
		0xef174541b62cac82,
		0xf001fc1d5e574a07,
		0xf0475101099acdc9,
		0xf1c587295608596e,
		0xf405ef1c8e600f0f,
		0xf5cae52becabc767,
		0xf64da2416445f8b6,
		0xf72d33f93d1ebb59)
}

var x_a8cb0f2f1a756b32 = []byte{
	0, 0, 0, 0, 23, 0, 0, 0,
	1, 0, 0, 0, 62, 0, 0, 0,
	25, 0, 0, 0, 146, 0, 0, 0,
	33, 0, 0, 0, 130, 0, 0, 0,
	37, 0, 0, 0, 114, 0, 0, 0,
	41, 0, 0, 0, 90, 0, 0, 0,
	45, 0, 0, 0, 82, 0, 0, 0,
	49, 0, 0, 0, 90, 0, 0, 0,
	53, 0, 0, 0, 82, 0, 0, 0,
	120, 45, 115, 97, 110, 100, 115, 116,
	111, 114, 109, 45, 97, 112, 112, 45,
	42, 0, 0, 0, 0, 0, 0, 0,
	111, 99, 45, 116, 111, 116, 97, 108,
	45, 108, 101, 110, 103, 116, 104, 0,
	111, 99, 45, 99, 104, 117, 110, 107,
	45, 115, 105, 122, 101, 0, 0, 0,
	120, 45, 111, 99, 45, 109, 116, 105,
	109, 101, 0, 0, 0, 0, 0, 0,
	111, 99, 45, 102, 105, 108, 101, 105,
	100, 0, 0, 0, 0, 0, 0, 0,
	111, 99, 45, 99, 104, 117, 110, 107,
	101, 100, 0, 0, 0, 0, 0, 0,
	120, 45, 104, 103, 97, 114, 103, 45,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 8, 0, 0, 0,
	1, 0, 0, 0, 22, 0, 0, 0,
	5, 0, 0, 0, 146, 0, 0, 0,
	13, 0, 0, 0, 90, 0, 0, 0,
	120, 45, 115, 97, 110, 100, 115, 116,
	111, 114, 109, 45, 97, 112, 112, 45,
	42, 0, 0, 0, 0, 0, 0, 0,
	120, 45, 111, 99, 45, 109, 116, 105,
	109, 101, 0, 0, 0, 0, 0, 0,
}
