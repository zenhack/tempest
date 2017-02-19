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
	WebSession_Context_headerWhitelist  = capnp.TextList{List: capnp.MustUnmarshalRootPtr(x_a8cb0f2f1a756b32[0:248]).List()}
	WebSession_Response_headerWhitelist = capnp.TextList{List: capnp.MustUnmarshalRootPtr(x_a8cb0f2f1a756b32[248:320]).List()}
)

type HttpStatusDescriptor struct{ capnp.Struct }

// HttpStatusDescriptor_TypeID is the unique identifier for the type HttpStatusDescriptor.
const HttpStatusDescriptor_TypeID = 0xbc353583a3731ade

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

// WebSession_Params_TypeID is the unique identifier for the type WebSession_Params.
const WebSession_Params_TypeID = 0xd7051b9757f6b096

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

// WebSession_Context_TypeID is the unique identifier for the type WebSession_Context.
const WebSession_Context_TypeID = 0xf5cae52becabc767

func NewWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return WebSession_Context{st}, err
}

func NewRootWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
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

func (s WebSession_Context) AcceptEncoding() (WebSession_AcceptedEncoding_List, error) {
	p, err := s.Struct.Ptr(5)
	return WebSession_AcceptedEncoding_List{List: p.List()}, err
}

func (s WebSession_Context) HasAcceptEncoding() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetAcceptEncoding(v WebSession_AcceptedEncoding_List) error {
	return s.Struct.SetPtr(5, v.List.ToPtr())
}

// NewAcceptEncoding sets the acceptEncoding field to a newly
// allocated WebSession_AcceptedEncoding_List, preferring placement in s's segment.
func (s WebSession_Context) NewAcceptEncoding(n int32) (WebSession_AcceptedEncoding_List, error) {
	l, err := NewWebSession_AcceptedEncoding_List(s.Struct.Segment(), n)
	if err != nil {
		return WebSession_AcceptedEncoding_List{}, err
	}
	err = s.Struct.SetPtr(5, l.List.ToPtr())
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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
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

// WebSession_Context_Header_TypeID is the unique identifier for the type WebSession_Context_Header.
const WebSession_Context_Header_TypeID = 0xb4e5f4cccb748429

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

// WebSession_PostContent_TypeID is the unique identifier for the type WebSession_PostContent.
const WebSession_PostContent_TypeID = 0xb7d82eac416ab63e

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

// WebSession_PutContent_TypeID is the unique identifier for the type WebSession_PutContent.
const WebSession_PutContent_TypeID = 0xd7aff1fe39659132

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

// WebSession_ETag_TypeID is the unique identifier for the type WebSession_ETag.
const WebSession_ETag_TypeID = 0xd22c0be5b9c16558

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

// WebSession_Cookie_TypeID is the unique identifier for the type WebSession_Cookie.
const WebSession_Cookie_TypeID = 0xa87d65bed9b60243

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

// WebSession_AcceptedType_TypeID is the unique identifier for the type WebSession_AcceptedType.
const WebSession_AcceptedType_TypeID = 0xaaf9021b627cc1f9

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

type WebSession_AcceptedEncoding struct{ capnp.Struct }

// WebSession_AcceptedEncoding_TypeID is the unique identifier for the type WebSession_AcceptedEncoding.
const WebSession_AcceptedEncoding_TypeID = 0xbda585bffe1dc7e8

func NewWebSession_AcceptedEncoding(s *capnp.Segment) (WebSession_AcceptedEncoding, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_AcceptedEncoding{st}, err
}

func NewRootWebSession_AcceptedEncoding(s *capnp.Segment) (WebSession_AcceptedEncoding, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return WebSession_AcceptedEncoding{st}, err
}

func ReadRootWebSession_AcceptedEncoding(msg *capnp.Message) (WebSession_AcceptedEncoding, error) {
	root, err := msg.RootPtr()
	return WebSession_AcceptedEncoding{root.Struct()}, err
}

func (s WebSession_AcceptedEncoding) String() string {
	str, _ := text.Marshal(0xbda585bffe1dc7e8, s.Struct)
	return str
}

func (s WebSession_AcceptedEncoding) ContentCoding() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSession_AcceptedEncoding) HasContentCoding() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_AcceptedEncoding) ContentCodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSession_AcceptedEncoding) SetContentCoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSession_AcceptedEncoding) QValue() float32 {
	return math.Float32frombits(s.Struct.Uint32(0) ^ 0x3f800000)
}

func (s WebSession_AcceptedEncoding) SetQValue(v float32) {
	s.Struct.SetUint32(0, math.Float32bits(v)^0x3f800000)
}

// WebSession_AcceptedEncoding_List is a list of WebSession_AcceptedEncoding.
type WebSession_AcceptedEncoding_List struct{ capnp.List }

// NewWebSession_AcceptedEncoding creates a new list of WebSession_AcceptedEncoding.
func NewWebSession_AcceptedEncoding_List(s *capnp.Segment, sz int32) (WebSession_AcceptedEncoding_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return WebSession_AcceptedEncoding_List{l}, err
}

func (s WebSession_AcceptedEncoding_List) At(i int) WebSession_AcceptedEncoding {
	return WebSession_AcceptedEncoding{s.List.Struct(i)}
}

func (s WebSession_AcceptedEncoding_List) Set(i int, v WebSession_AcceptedEncoding) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSession_AcceptedEncoding_Promise is a wrapper for a WebSession_AcceptedEncoding promised by a client call.
type WebSession_AcceptedEncoding_Promise struct{ *capnp.Pipeline }

func (p WebSession_AcceptedEncoding_Promise) Struct() (WebSession_AcceptedEncoding, error) {
	s, err := p.Pipeline.Struct()
	return WebSession_AcceptedEncoding{s}, err
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

// WebSession_Response_TypeID is the unique identifier for the type WebSession_Response.
const WebSession_Response_TypeID = 0x8193ac6cb5429c83

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

// WebSession_Response_Header_TypeID is the unique identifier for the type WebSession_Response_Header.
const WebSession_Response_Header_TypeID = 0xb4b873147ab5ce5e

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

// WebSession_RequestStream_getResponse_Params_TypeID is the unique identifier for the type WebSession_RequestStream_getResponse_Params.
const WebSession_RequestStream_getResponse_Params_TypeID = 0xe9a02a3219bdbd70

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

// WebSession_WebSocketStream_sendBytes_Params_TypeID is the unique identifier for the type WebSession_WebSocketStream_sendBytes_Params.
const WebSession_WebSocketStream_sendBytes_Params_TypeID = 0x9a712ce3fcad8cd8

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

// WebSession_WebSocketStream_sendBytes_Results_TypeID is the unique identifier for the type WebSession_WebSocketStream_sendBytes_Results.
const WebSession_WebSocketStream_sendBytes_Results_TypeID = 0x82a3ee23aa0ae3a3

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

// WebSession_CachePolicy_TypeID is the unique identifier for the type WebSession_CachePolicy.
const WebSession_CachePolicy_TypeID = 0xb37b21e300864885

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

// WebSession_Options_TypeID is the unique identifier for the type WebSession_Options.
const WebSession_Options_TypeID = 0xe9ff06beec4e73d6

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

// WebSession_get_Params_TypeID is the unique identifier for the type WebSession_get_Params.
const WebSession_get_Params_TypeID = 0xcd94acddf4778328

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

// WebSession_post_Params_TypeID is the unique identifier for the type WebSession_post_Params.
const WebSession_post_Params_TypeID = 0xaa6ef20a62c1cafd

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

// WebSession_openWebSocket_Params_TypeID is the unique identifier for the type WebSession_openWebSocket_Params.
const WebSession_openWebSocket_Params_TypeID = 0xc7c9c9b19d935e79

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

// WebSession_openWebSocket_Results_TypeID is the unique identifier for the type WebSession_openWebSocket_Results.
const WebSession_openWebSocket_Results_TypeID = 0xcc561276d31b392b

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

// WebSession_put_Params_TypeID is the unique identifier for the type WebSession_put_Params.
const WebSession_put_Params_TypeID = 0xf1c587295608596e

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

// WebSession_delete_Params_TypeID is the unique identifier for the type WebSession_delete_Params.
const WebSession_delete_Params_TypeID = 0xeba76bffb27b1975

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

// WebSession_postStreaming_Params_TypeID is the unique identifier for the type WebSession_postStreaming_Params.
const WebSession_postStreaming_Params_TypeID = 0xd26a7affce43b1c0

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

// WebSession_postStreaming_Results_TypeID is the unique identifier for the type WebSession_postStreaming_Results.
const WebSession_postStreaming_Results_TypeID = 0xbf0e0653dc266205

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

// WebSession_putStreaming_Params_TypeID is the unique identifier for the type WebSession_putStreaming_Params.
const WebSession_putStreaming_Params_TypeID = 0xa1ece076a7105939

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

// WebSession_putStreaming_Results_TypeID is the unique identifier for the type WebSession_putStreaming_Results.
const WebSession_putStreaming_Results_TypeID = 0xc60d14bf989d4454

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

// WebSession_propfind_Params_TypeID is the unique identifier for the type WebSession_propfind_Params.
const WebSession_propfind_Params_TypeID = 0xca2d58de88f0b32e

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

// WebSession_proppatch_Params_TypeID is the unique identifier for the type WebSession_proppatch_Params.
const WebSession_proppatch_Params_TypeID = 0x9e582e7e054088ae

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

// WebSession_mkcol_Params_TypeID is the unique identifier for the type WebSession_mkcol_Params.
const WebSession_mkcol_Params_TypeID = 0xf64da2416445f8b6

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

// WebSession_copy_Params_TypeID is the unique identifier for the type WebSession_copy_Params.
const WebSession_copy_Params_TypeID = 0x8139673a82bfe07d

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

// WebSession_move_Params_TypeID is the unique identifier for the type WebSession_move_Params.
const WebSession_move_Params_TypeID = 0x81f5066b5576a609

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

// WebSession_lock_Params_TypeID is the unique identifier for the type WebSession_lock_Params.
const WebSession_lock_Params_TypeID = 0x9398280f1359570a

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

// WebSession_unlock_Params_TypeID is the unique identifier for the type WebSession_unlock_Params.
const WebSession_unlock_Params_TypeID = 0xd684c6a791b97dbc

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

// WebSession_acl_Params_TypeID is the unique identifier for the type WebSession_acl_Params.
const WebSession_acl_Params_TypeID = 0x9f79c33e20119e8d

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

// WebSession_report_Params_TypeID is the unique identifier for the type WebSession_report_Params.
const WebSession_report_Params_TypeID = 0xc0643ea68efc60ae

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

// WebSession_options_Params_TypeID is the unique identifier for the type WebSession_options_Params.
const WebSession_options_Params_TypeID = 0xd2e47e8eac54ea7e

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

// WebSession_patch_Params_TypeID is the unique identifier for the type WebSession_patch_Params.
const WebSession_patch_Params_TypeID = 0xadef95edc22ca880

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

const schema_a8cb0f2f1a756b32 = "x\xda\xcc{\x7fx\x15\xd5\x99\xff\xfb\xce\xdc\x9b!H" +
	"\xb89Nl%\x98\xde\xc07 \x89\xdcH\x12h%" +
	"\xcfW\xf3\xcb(f\x8dd\xf2C$[Z'\xf7\x1e" +
	"\x92!\xf7\xce\\f\xe6BB\xa5H\x94\xadZk\x0b" +
	"B+<h\xab\x96\xfa\x0b\xb6\xfe\xc2-\xddRKW" +
	"\xdab\xd5\xaenm\xeb\xee\xda\xb5\x8f\xf5ii\xa5U" +
	"Wmm\xc5\xd9\xe7\x9c\xb9\xf3#\xf1\x92D\xfa\xec>" +
	"\xfe\xc5\xcd\x99\xf7\xbc\xe7=\xef\xf9\xbc?\xcfa\xc9\xcd" +
	"\xa5\xcdB]\xf4\xe7\xe5\x00=\x17`\xb4\xc8\xd9\xfc\xd2" +
	"\xe3c\x8d\x83\xcb\xb7\x82R\x81\x08\x10\x15%\x80\x06\xa5" +
	"\xa4\x03\x01\x1b\xd6\x94\xc4\x11\xd0Y\xf4R\xc9\xdb\xd1\xbe" +
	"\x9e\xad@\xaa\x85?m[\xf1\x0f/\xcf\xfb\xcc#\x80" +
	"\x0d\xb9\xd9\x9bP\xbeq\xb6\x04 o\x9b\xbd\x05\xd0\xb9" +
	"no\xeb\xc1\xf4\xfe[]V\xce3\xf7.\xeb\xbc\x93" +
	"H\xfb\xa0\xbdX*Bl8:\xbb\x1e\xe5\xe7g_" +
	"\x0a\xd0P\x12s\x18\xdf\xe2ol\xe8\x1b.zk\xfc" +
	"\xd2\x07J\xd9\xd2\xf2\xa1\xd2\xdf\x00\xfe\xd7\x17ox\xee" +
	")I\x1e#\x0b\x85\x809`\xc3.\xb2\x03\xe5\x03\x84" +
	"\xad|\x1f\xf9&\xa0s\xf7\xcb3\xef\xff\x7f\x7f\xb8{" +
	"\x0cH\x02\x01\"\x8c\xcfj\xf9/\x08\x11\xe7';\xff" +
	",\xac\x7f\xcc\xbe\x0e\x94\x85\x18\x92\xf0\xacb\x09\x01\x1a" +
	":\xe51\xb6\xd8jy#\xa0\xd3\xfe\xb1w\xff\xfd\xf3" +
	"?}\xf5\xf3\x85I\x0f\xc9\x9b\x18\xe9Q\x99\xad\x87+" +
	"#\x97\xff\xc7\xdc\xbbo\x81\x89\x92ie\x0f\xa3\xbc\xb5" +
	"\x8c\x09\xb0\xb9\xac\x89mr\xe6\xaa\xd5rl\xd1m\xb7" +
	"\x8e\xdb$\x9e\xcd7Yr\xf6o\x00\x9d\x8b\xca\xceX" +
	"\xf8\xf3{V~\x05\x94\x05\x88\xce\xe0\x0f\x1fx\xf5\xbc" +
	"W\x9e|\x0b\xfa\x8a$\x8c\x024\xbcr\xf6\xc3\xec," +
	"\xde:\x9b\x9f\xc5W\xd7\xedy\xfb\xbf\xbf\xef\xec\x06R" +
	"!\x06*\x06l\xa8+\xefG\xb9\xbd\x9c\xe9\xa4\xa5\xfc" +
	"R9S\xfeQ\x00\xe7\x177\x1fx\xf7\xe5\xc5\xeb\xf7" +
	"\xb8z\x89\"[{M\xf9\x1bl\xed\xf5\xe5M\x80\xce" +
	"?\xde\xd0\x1c\xfdl\xedUw\x00\xa9\xf0\x85\xdb^n" +
	"2\x82\xdb\xcb\xd9N\x8b>rdG\xd7\xc2{\xee\x00" +
	"rf\xc4\xa9\x1f\xce\x95\x9f\x1f\xfb\xf1\xbdl\xbd\xe5s" +
	";P\xee\x9c+\x01\xf4\xac\x98+bO\xef\\\x01\x01" +
	"\x9c/\xdcA*/\xfa\x97\xd1\xaf\x86\x19\xb6\xcfme" +
	"\x0c;\xe72\x86\xcbW\x97\xde\xb3\xe1\xa5W\xef\xcc\x13" +
	"\xf0\xb3:>\xf7&F\xf0\xce\\\xa6\x0e\x7fS\xe4L" +
	"1\xbc\xa2\xfc\xfc9\xbf\x93\x7fu\xceo\x18F*>" +
	"\x17\x91\xb7W\xb2\x1d\xb6\x09\x8f\xbd\xf0]\xba\xf9^\xa6" +
	"^!\x98\xeb.\xbc\xb5r>\xca\xdb+\xd9\xcf/T" +
	"r\xf5\x9d|\xf2\xc8\xc0\xcc7\xf4\xfb\xc3\xe2m\x9b\xc7" +
	"\x0fc\xfb<&\xde;G\xae\x19\x98+\xbcs\xff\x04" +
	"\x0c\xbb\xaa\xab\x9e\xdf\x8d\xf2\x85\xf3\x99\x8e\x97\xcfg\x90" +
	"\xb9\xf7\xaf\xa3'\x1f\xfc\xe4\x13\x0f\x14\x86\xcc\xee\xf9L" +
	"\x91\x0d\xfb\xe6\x7f\x91-}\xed\xbd\x8b\xbf\x7fb\xd7\x1f" +
	"\x0f\x84\x97.\xae\xeafK\x9fU\xc5\x96\xfe\xdc\x0d\x0f" +
	"\xa4g\xcd\\\xf1MP\xce\x8c\xce\x18\xb7\xf3\x03U\xbf" +
	"\x93\x0fU1U?Z%\"\xa0g\x84\xe3e\xe4\xaa" +
	"\xdcW\xd5\x81\xf2\xa1\xaa\x8f\x02\xc8G\xaa\x98:\xe7~" +
	"\xfd\x91\xdd\xedO\xec}\x04H\x850\x0e0\xfb\x16\xf4" +
	"\xa3|h\x01\xdb\xcc\xc1\x05+\x01\x9dO=spS" +
	"\x99\xf5\xadG\x81,\x0c\xed%*0\xb6o-X\x87" +
	"r\xf1BF\x1c]\xc8v^}\xbd\xfd\xe3\xa7\xde|" +
	"\xe5Q \x0bB\x98u\x89\xd5\x85\x03(\xe78\xf1z" +
	"N|\xd1c\xebZ\xf6\xd7\xfe\xe2\x9f\xd8\xd6'\x9e\xd1" +
	"\xd3\x0b;P\xfe\x15'~q!S\xc3\xef\x93\x97\xf4" +
	"f\x7f\xbb\xff\xdb@\x16D\x02\xce\x80\x0d\x9b\xcf\xbd\x1f" +
	"\xe5]\xe725|\xe9\\\x11{\xf6\x9f\xcb\x11\xf7\xcb" +
	"r\xeb\xee\xeb\x96-\xfb\x0e(g\"\x06jsO\x8c" +
	",jDy\xde\"\xc6\xbdb\x11\x13e\xe3\x8a\x91\xbd" +
	"\x8f\xfd\xfd\xf0aP\xaaP\x08\xd0\xd3'J(\x024" +
	"\x8c.\x1a`\x07\xb2m\x11\x93\xe4\xb7?\xacx\xef\xf1" +
	"m\xfb\x0e\x17\xc4\xc2\x82j\x13\xe5\xe5\xd5\x8c\xf3\xb2j" +
	"\xc69:\xb0\xf0?{\x8af?\x9e?_N\xb4\xab" +
	"z\x0fc\xb7\xaf\x9a\xdb\xda\xd5\xef\xde\xf2\x8d\x8bR\xdf" +
	"\x0b\x03\xe0hu?#x\xb6\x9a\xad\xd7{\xf1\xed\xb7" +
	"=^V\xf2\x830\x07\xadf\x07#\x18\xada\x1cF" +
	"?u\xeb\xed\x0f\x1d;\xf6\xc3\xb0\xed\xecv\x09\xee\xab" +
	"a\x87]\xfb\xc8k7\xfc\xf2\xaa\xc4\x93\xe3\x9c\xcd\xae" +
	"\xf3\xd61\x8a;\xcfc\x14\xe7-\x9f\xfbo\x1b\xce\xbc" +
	"\xf2\xa9<\x0b~^\xa3\x8b\xb9\x947.f\xdbXt" +
	"\xdd\xc67_\xdc\xbf\xf3i\x8f\x05\xa78\xb9\x98[p" +
	"4\xc1\xc4\xbc\x8a\x1e9\xf4\xca\x19\x8b\x9f+\xa8\x96;" +
	"\x13g\xa2\xfcP\x82\xa9\xe5@\x82\xf1\xfb\xdeCm\xcf" +
	"8\x9b\xd6=\x17\x96\xb9\xb8\x96\xcb<\xa7\x96I\xf4\xd9" +
	"\xdf\xf5\xee\xbf\xe5\xb3\xbf~.,\xd1[\xb5\xfc\x18\xf0" +
	"|\xc6\xe1;\x9b\x0fm\xbf\xe7\x07\xd7\xff,\xac\xb75" +
	"\xe7s\xbdi\xe73\x81\xbe\xfc\xe0\xdb\xab\xbe27\xfa" +
	"\xf3B\xf0\x8a.\x99\x8f\xf2YK\x98<d\x09#\xae" +
	"\xdfN\x97\xbf\xf7\xfa7\x0b\x12\x1f[\xd2\x8a\xf2\x0b\x9c" +
	"\xf8yN|\xf1\xb7\xae\xbd&\xf5\xe0\x96\x17\x0a\xdbw" +
	"\xae\x8eoc[\x1d\x932{\xf8\xf0\x9c\xfa\x9a\xaf\x1d" +
	"\x07r\x9e\x17\x82\x8e\xd7\xbd\xc1B\xd0\xcf\xac+^\xfd" +
	"n\x91s\xbc\xa0\xbe^\xac\xabA\xf9D\x1d[\xf2x" +
	"\x1dSGn\xceg\x1ev\x86\xef\xf9}X\x1d/\xd6" +
	"\xf3\xdd\x1e\xafg\x0b}\xff\xc6\xbd\xf2\xf1\x997\x9d\x00" +
	"\xa5\x16C\x1e\xa8\xafXB\x819\xdb\x06~\x98}\x0d" +
	"\x8cvl\xff\xe2\xc7Z\xda?\xfaG \x0b#\xe3\xe2" +
	"\xd4C\x0d\x0f\xa3|\xac\x81\xd9\xd2\x13\x0d\"\xf6\xbc\xd4" +
	"\xc0mI\xeaX\xf5\xa9\x8aw\xf1\xb5\xf7E\x97\x13\x0d" +
	"\xebP\xc6\xa5L\xcc\x93\x0d\x97\xca\xd5\xec\x97s\xec\xe9" +
	"=\xc5\xa8\\\xfaZa\xed\x90\xa5\\;\xf3\x962\xe4" +
	"\xea\xabg\\Y\xfd\xb9\xa3\xaf\x8f\x0b\x0bK\xdd\xb0\xb0" +
	"\x94i:\x16\xbb\xfa\x96s\xfe\x18}\xb30\xafW\x96" +
	">\xc9\x03\x04\xe7\xe5\xfb\x84\x09\xfa,bL+\x96\xd5" +
	"\xa0\x9cXv.\xe3\xbfl\x15\xf3\xbb\x8f\xfd\xb9=\xd5" +
	"rW\xe7\xdb\xe1\xa5\x9f\xfe8\xf7\xbb/|\x9c-\xbd" +
	"\xfa\x9f?v\xe1;\x0d\x89?\x15R(\x8a\x0d\xeb?" +
	"\xf1\x1c?\xe5Ol\x84\xe7\x9c\x8dt aQ\xcb\x8a" +
	"h\x86^\x9bT\xb3z\xb6q\x15\x1d\xe8\xa1\x96\xc5\x07" +
	"\x8c\xechU\x97j\xaa\x19\x0b@)\x13#\x00\x11\x04" +
	" \x9bk\x00\x94\x11\x11\x95\xeb\x05$\x88e\xcc\xb0\xc8" +
	"\xd6\x01\x00\xe5Z\x11\x95\x9b\x05D\xa1\x8c\x1d\x1e\xb9\x91" +
	"\x8d\xdd \xa2\xb2S@\"b\x19\xf3Id{+\x80" +
	"r\xb3\x88\xcam\x02\x92\x88P\x86\x11\x00\xb2\x8b\x0d~" +
	"IDe\xaf\x80\xb1\xacj\x0f\xe1,\x10p\x16\xa0\x93" +
	"\xa2\x96\xad\xe9\xaa\x0d\x92f\xe8\xfe\xa8n\xac\xdc@\xcd" +
	"\x8d&H\x9aM\x11A@\x04\xdcb\x0d\xa9\xe9\xb4\xb1" +
	"\xd1\xff;i\xe86\x1d\xb1\xb14\xecx\xb1\x94\xf9\xcd" +
	"I7\xde\xa6&\x87h\x97\x91\xd6\x92\xa3M\xb5=I" +
	"#K\xbb\x10\x15wO\xcbj\x18\x0f\x92\xe8\x07@\x81" +
	"T\xb7\x02\xa0H\xe6\x99\x00\x18!\x15\xdd\x001\xdd\xd0" +
	"\xa9\x93\xa5&g\x06\xa2\xa1o\xc9R\xb3\xcf\xa2&\x1b" +
	"l\xc9f\xaf\xa4\x107\xd92NN\xd76P\xd3R" +
	"\x01\xd3\xbe@bA\x81\xba\xa9\x15\xcf\x1a\xbaE\x95R" +
	"\x0c\xd2\xc89\x03\xa1\xccm\xceX(\xe2\xcdi\x0c\xd9" +
	"\xcaYcNO.\x99\xa4\x96\xd5\x06\x92\x91\xa2N[" +
	"Z\xa3\xba\xddn\xa2i\x98mF\x8aB\xd3\x0a\xaa\xa6" +
	"\xa8\xe9\x0c\xf1\x7fV\x0d\xa1f\xd3\xb4f\xd9\xa0T\xf9" +
	"\x87~\xa2\x1f@yU\xc4\x9e\x19(`I\xd4q0" +
	"\x94b\xcaQ\xec\x00\xa1Dx\xcf\x09C\x8e\xbc\xde\x0a" +
	"BI\xd1I6\xe8\xbb\x1e\xf2\xc2\x00\x08%\xd2\xbbl" +
	"\xd0\xb78r\x94\x0d\x8a\x7fe\x83~\x8aK\x1e\xea\x06" +
	"\x81`Q\x19\x96\x02\x90\xdb\x19\x96\xf6\x8a\xa8\xdc+`" +
	"I\xe4/\x8c\xd272\xb2o\x0f\x08d\xc6\x8c2," +
	"cX\xda\x01\xa0\xdc&\xa2\xf2\xa8\x80\x8eE\xed6\xc3" +
	"\x18\xd6@\xa4\x16\xce\x06\xec\x12\x11K\x83@\x09\xc8\x06" +
	"\x1d\x93\xa64\x93&m\x00pA\xa3\xdbN2\xaf%" +
	"\x90L\xc3t,jn\xa0\xa6\xf7\x97n\xb4q*@" +
	"\xdbI\xe6\xc1\x02\x92\x96\x1c\xc5\xd2\xa0\xa2\xe0@\xcb\x9a" +
	"4i\xe8)\x0dm\xcd\xd0/Q\xb5\xb4HS\x8e\x9a" +
	"Ji\xecoT\xd3\xae\xe61$\x9b\x7f\x8ay\xd9&" +
	"\x87j\xc6\xd8@\x03\x1b-\xf5\x8fKe6\xfaI\x11" +
	"\x95\xa1\x90\x8dR\xa6\xc3\x94\x88J6\xb0\xd1\x0c\x1bK" +
	"\x8b\xa8\x8c0\x1b\x15\\\x1b\xcd1s\xcc\x8a\xa8\\\xf3" +
	"\xb7\x98\xe3\xd4\xe6\x17=\x15\xda9\xd8k=\xd8\x1a)" +
	"\xe4\x16x\x0e\x97\xb8\xb3\x9ct\xc6\x11\xc9\xfaV\xb2>" +
	"\x8e\x02\xd9\xd5Av\xc7Q$G\xba\xc9\xd18F\xc8" +
	"\x89M\xe4\xf58F\xe598 W`\x1c\x8b\xe4v" +
	"\x1c\x90/\xc3\xb8h\x0c+\x11\x14B\x99iD\xc0\x96" +
	"\xb3\x99H\xd0\x85\x88?\"X.\xae\xfc;\x86\x01\x93" +
	"\xaa6MMJ}\x8c`\xeb\x966\x97\x10\xc0Q\x93" +
	"I\x9a\xe5?\xa1\xc0\xb4s\xfciO\x12\xecpZ|" +
	"b\x0cci\xd2\x89O\x11\xecw\xae0*95\x88" +
	"\xba\x0d\xe0dU\xd3\xd6\xd4t\x1b4\xb9\xa0\x9d\x94\xc1" +
	"3\x04\xc7\x9c.wF%\xba\x8b2\xc4;\x99\\\xda" +
	"\xd6zl\x15$;gM\xca\xe2'\x04\xd79\x9d\x8c" +
	"<\xd1cCL\xb5s\x16\x80\xa3\x1bv\xa7\x91\xd2\xd6" +
	"\x82\xa4\x15\xd4\x99?_\\\x82\x8c\xc1\x15\x86]\xc9&" +
	"@l\xad\xc6u\xe7\xc1\xa1\xa8 \x1c\xd8O#9L" +
	"\xed\x1e\xdb\xa4j\xa6\xd6\xa2z\xaau\xd4\xa6VUw" +
	"\x13\xb5ri\xdb\x9a.\x9e<U\xdb\xa0\xcc`\x96R" +
	"\x861\x00R=\x06\xa0,\x12QY\x9a7\x15\x19\x80" +
	"\xd41\xfbY,\xa2r\x01s\"CF.\x9d\xea\xa6" +
	"hQ\xfb\x12\xc3\xcc\x80\x07\xf1\x18\xedU\x07\xb14\xc8" +
	"!\xa7\x15^|yL\xda\xe4\xfa\x1de\x96+\x0e3" +
	"\xd2vf\x90\x17\x8b\xa8t\x09H\x04,\xc33\x18\xe6" +
	";\x00\x94\xcbET\xae\xca\xcb8\x0b\x80\xf41\xca^" +
	"\x11\x95\xab\x05t4\xab\x8b\x9a\x19U\x07\x89\xea\xb6'" +
	"\x9f\x936\x92*\xf34\x00\xe0\x1b\xab\xb5Q\xb3\x93C" +
	"\xbd\x06H\x97\xd2\x80r\x9a*\xf4b\x87i\xc4y\xec" +
	"`f\xb9\x94\x9b\xe5\xae~f\x86\x987C\x81\x9c\xe8" +
	"`f(\xcas\xf0&y\x1e\xc61\"w\xa2)+" +
	"\xc8Ls\x14;\xe4\xcd\xdc4\xf7a\x8d\xbc\x0f\xe38" +
	"C>\x86\x0f\xcb\xcfb\x1c\x8be\x14v\xc8\xc5B\x1c" +
	"g\xca\xcb\x84\xfb\xe5\x0b\x858\x9e!kB\xb7\x9c\x11" +
	"\xe2(\xc9\xbb\x84=\xf2\xedB\x1cg\xc9\xc7\x84\xbb\xe4" +
	"g\x85\xb83\xa0\xa6\xba\xe9\xfa\x1c\x05\xd1\x9a\xd4\x02\xc4" +
	"/!\xc1\x01\xa7UMU2z\x90\xa8\xc5\x8ch\xad" +
	"a\x0eh\xa9\x14\x05\xd4'\x9d}+\x12\xecv.\x09" +
	"\xa8]\xe8_b\xe4\xf4\xa9\xac^\xdc\xc9\xe72\xe03" +
	"r@\x86\xfa\x0c\xb5\x87\x8c\xd4\x15\x06\xda-,ka" +
	"\x96P\x80G\xa5\xcfc\x17\x12\xdc\xe3t\xf2Y\x95W" +
	"\xa0aW\xf2y\"7!\xdd\xb0\xb9S\x81\xb8\xad\x0e" +
	"\xa4\xe9\xa4\xc2|\x19\x09n\xe2\xc2\xb0)\xd0\x94\xe5S" +
	"\x00\x9c\xa4\xa1\xafMk<\x0cN\xca`72'\xd6" +
	"\xe6Scl\xd0\xd0\xe9d\xdeR\xdc\x83\x04kb\x97" +
	"\x1a:[\xc7d\xa7e\xd9\xed\xa8\xdb\x9a=\xdak\x18" +
	"\xf1\xcbUs\xb0\x10\x83*\x9f\xc1\xedH\xf0I\xa7\xdb" +
	"\x9dY)\xb4\xf3\xa9\x95\xbd\x86Q\xc9\xe7\xf2 \xce\xbf" +
	"\xf5\xa1\xa9\xf5\x1a\xc6\xe5\x86\x8e\x83\x93\xea\xf3\x0e$x" +
	"\xbf\xc71\x81}\xdd\x97qv\xb1\xcb\x0d}\x10\xc0\xc9" +
	"\xe9V.\x9b5L\xb4i\xaa\x93\xa645\xd6;\x9a" +
	"\xa5\x93r\xfc*\x12\xfc\xb6\xd3\x17\x9aX\xc9f6\xa9" +
	"\x95l*\x80\xa3eZz\xa9\x9a5\xa6\xf0\xf5\xe2]" +
	"\xdcO^vn\xa6R\xad\xb4)\xc4\xd4\xac\xc1\xdd}" +
	"\x81dbR\x81\xf6\"\xc1\xbb\x9c\xae\xf0\xb4\xcaKT" +
	"MJs\xcc\xe4\xf4\xaci$\xa9\x85\x16;\xffv\xdd" +
	"\x964{tR~\xdf\xe0*\xeb\x1b7\xaf\xb2]\x8f" +
	"\xb1\xb3\x08\xf9\xf1\xc2n/m$\x87\xa7\x9f\xaa\xf4\x17" +
	"HUXV2$\xa2b\x87R\x95\xf5\xadA\xfe2" +
	">U\x19\xc9\xa4\xfdX\xe9\x0d\x9eF\x89P\xd8!\xb6" +
	"\xb9\xf3jY\x00\xc8\xeb7\xces:V'\xcdr\x1c" +
	"\xb7\xac\x09WJ%\xf8\x9eS\x86QV*5\x02(" +
	"\xd7\x88\xa8\xdc `\x85x\xd2\x89\x94a\x11\x00\xd9\xb6" +
	"\x0e@\xb9^D\xe5k\x02VD\xdee\xc3\x12K|" +
	"\xcd|\xe2\xfb-\x01K\x84\xbf:e8\x03\x80\x1cd" +
	"\xbe\xffQ\x11\x95\xef\x09\xc8K\x0e(j\xa2#\x9ae" +
	"[P\xe4dT;9D\xad\x95\x10\xd3\xe9\xca\xb5A" +
	"j\x19\x0eS,\xb5\xcc\xd3]\x01qcr\xc2\x94A" +
	"-\xddn\x1f\x01\x89U\x05ES\xc67nR,\\" +
	"\x8bj\x86E\x88\x88\x18e\xf0\xcd\xf7\x14\xd0\xab\x87\x09" +
	"\x19\x00\x81\x14K\xce \xb5y\x84\x01I\xb7h3*" +
	"\x11D\xe7\xd3/\xff\xa4z\xe3\x05\xab\x9e\x86\xbf!A" +
	"\xe8\x8as\xb8)\x11\x1fl%\x0c.3DT\xca\x04" +
	"\xdc\x92\xa1\x96\xa5\x0eR,\x01\x01K\xa6\x0c\xdbY\xd3" +
	"\xc8f\x99\xc6|\xb6\xb3|\xb6\xed\xec\xa0\x9bET." +
	"\x0fa\xf82\x86\xe1\x15\"*\xbd,\x92\xe7A\xac\xb4" +
	"\x06\x91|\x1ax\xfd@%\xec\x90mg{l\x96\x95" +
	"\xb5\xe8\xbaa\xf3\xc8\x7fY\xcaM\xdd\x8aA\xc0\xe2\xc0" +
	"\xb0\xa7\xb4V5\x99\xf6\x8c\xf5\xc3\xb1\xcf\xc2v\x98\xcd" +
	"\xe5O]\xd3\x07\xa7\xef\\:\x02\xe7\xe2\x0b<\xde\xbb" +
	"\x88y\xef\xd2q\xaaB(\xa3e\xa8\xeb\xd5a\x1a\x9b" +
	"\xa0z\xd2Hi,\xae\x04y\x98\xb71a\xe2\xc6b" +
	"lg\xca\x05\x18n\x05j\x8d\x01OB[C=h" +
	"u \xd4\x04T\xfb\x03\xdb%jMP\xe8\x925\x8d" +
	"\xc1e\x00Y\xbd.\xe8\x06\x93\xd57\x05\x0d*\xb2\xa6" +
	"#t'\xb3\xc6\x0c\xb5\xd0\xd6\x8cy\xa5-Y3\x10" +
	"t\x01\xc9\x9a\xd6PS~\xb5\xd9\xe4\x1e\xc2\x96\xbc\x83" +
	"t\xba\x0c\xcbf\xbfy^\xeat\xe5l\xff\xecc\xed" +
	"\xbd\xea`\x13\xaf\xceiP\x11\xf1(\xeb\xff\x89\xed\xbe" +
	"\xe2\x1c/\x07\x05\xfe\xdbu1\x10\xe7\x87\xefx.@" +
	"\xc8\xa3\x01\xbc\x16\x0e+\xca\xb7\xac\xcc2K\xb0\x9c." +
	"\xd3\xc8\xae\xd5\xf4\x14\xc4/\xa6Y{HY\xc1\xdd\x92" +
	"\xd7\"\xf6\xdd\x92\\'\xcc\x07A^ H\x18\\\xb1" +
	"\x04_\xe7\x085 \xc8%\x82\x84\x92\xdf\xc2F\xaf\x11" +
	"-\xa3`\x82 \xbf\x83\x12\x0a~\x9b0\x98{\x02\x19" +
	"\xe7_\xa1\x84\xa2\xdf\x1a\x0d\xbe>\x8f\x8d \xc8\xc7P" +
	"\xc2\xa8\xdfhF\xaf\x11/\x1fF\xc6\xf9 JX\xe4" +
	"_;\xa1\xd7d\x97\xef\xc3u \xc8w\xa2\x843\xfc" +
	"\xbey\xc0y\x17v\x80 \x7f\x01%,\xf6/\xc9\x82" +
	"\xaf[\xb1\x1b\x04y\x14%\x9c\xe9\xf7\x17\x83\xaf\x19\xac" +
	"\x07AVQ\xc23\xfc\xeb\xd5\xe0k\x1f2m\\\x86" +
	"\x12\xce\xf2o@\x83\xaf\x17\xf2\xafu(a\x89\x7fu" +
	"\x18|]\xc0\xbf\xceA\x09g\xfbm\xf1\xe0k\x09\xd7" +
	"\x06\xa2\x841\xff\x1e.\x08\x1do\xcd\x07\x81\x1c\x97\xb0" +
	"\xd4\xbf\x88\x08\xbe\xbd\xd8\x08\x02yVB\xe2w\xe3\xd1" +
	"\xc3+9\xda\x0a\x029$a\xc4\xbf\xc1\x0a\xe6\x1d\xa8" +
	"\x07\x81\xdc)I\x83\xd4n\xc6X\xd6\xb0\xecft\x8c" +
	",\xd5\x19\xbc \xcecL3J\xd9\x9c\xdd\x8cM)" +
	"\x9a\xa66mF\x87\xd11\xd4A\x9c{!6\x92w" +
	"J\x10\xf3\x06<\xe0\x01\xe4\xffb\x81\x04p\xa8\x19\xe3" +
	"\x99\xe1\xa4\x91n\xc6X\xd2\xc8\x8e6c,cl\xa0" +
	"\xcd\x18c\xc9R36\xe5t\xf7\x87\xa4&\xd3\xcd\xd8" +
	"d\xd2\xaca\xda\xcd\xb8\xc5p\x11\xdd\x8cq\x1e\x91\xf2" +
	"!\xd3\xf8\xc8\xc1\xab\xfe\xd0R\xffv\xd8\xbb\x8b\xa7H" +
	"_\x0ciXs\xdb\x9aS\xf4u\xeb\x83d\x05\x05\x0c" +
	"]9\x91mL\x95b\xa9\x9b\xd4d:B}\xa4\x88" +
	"\xe0\xa64\xb9\x9a\x90\xfb\xd4\xd5\x0c\xf5\x9c_|\x83\x9a" +
	"\xce\xf9\x7fm\xa1#Y\xcd\xa4\x96\xc3\x02\xd8J==" +
	"\xca\xdc\xa4Wb\x8f\xf3\xbaSDh\xc3\xb2\x83 0" +
	"U\xd4j\x0d\xd7\xd9\xf9 \xd0\xd9\x1a\x84\xb2q+{" +
	"mA,\x0d\\\xaf\xeb\xdbO\xbb\xcb\xec\xf9\xb8\xde\xd1" +
	"\xac\xdb\xde\x9a\xe1\x0b\\\xdd\x11\xb4$\xb8\xbc\x88\xa4\xae" +
	"1\xd4\x91(\x10}\x9a\xd6_\xc9u:\x13\x84#3" +
	"\xafm\x9a~\x17\"i\xc4\xdd\xe6\xd19b\x84k\x87" +
	")\xe2`G\x90_r\xed\xb0hx\x98\x0d~GD" +
	"\xe5G\xf9\x10\xc9r\xdc\xa3l\xf0\x09\x11\x95\x7f\x15\x10" +
	"\xa3\x18\xba\xd1!O\xd7\x80P*\x94\xe1L\xc6\xb0?" +
	"`\x88E\x18\xba\xa7 \x87Y\x16\x18\x91\xca\x90\x00\x90" +
	"\xfb\xd8Y}]D\xe5A\xa1p\xd8L\xab\xfa`N" +
	"\x1d\x0co\xbd\x90:b\x03Fj\xd4\xb1x6\xd4f" +
	"\x80\x98\xa2\x18\xf3:\xe7\x80\x18cY\xadfe\x0dK" +
	"s{\x99\xa7\xd7\xca\xc9\xe7\x83\x0crh}\x88\x10'" +
	"LL\x0ay\xaf\x8e\xe5\x81\xa5\xc1\x0d\xf3\x07\xb9\x07\xe1" +
	"\x1d\x81\xf0\xdb\x19\xac\x8f\xf3\x9b\x91p\xae\xd5\x0d\xa0\\" +
	"-\xa2\x92\x0emZ\xeb\x0e\xa5UB\xa5\xbb\xe9\xf5\x9b" +
	"\x02\xbf@\xc4y.\xbaF7\x05\xee\xc7\xd9\xa8\xd9C" +
	"mC4\x098\x8c\xb1`\xd9\xfc\xc9ey\xab\x8b\xf7" +
	"L\x0b|\xdd\xa0\x9a\x1a\xb5V\xea\x90O.\xfc\x1eW" +
	"\xf0\xc1\xb5\xbd\xf75\xbf\x0a\xab\xc1K\x1b.\xa6Y\xd1" +
	"\x1eb\x96:\x8b\xef\xa2\xa2\x83\xcc\x8b#\x92\x96\x1a~" +
	"\x17\xb4|>\xab\xf0\xf5\xb5\x9a\xce\xaba\xb7\x85\xf2\xcc" +
	"\xf6\xd7\xdf\x1b\xbd'\xf5\x86[I\x97!o\xdcaw" +
	"@\x87\x9f\x06\x88m\xa2\xa6!\x19:\x9d\xae\xc9\xae\xa0" +
	"15E\xcd\x09N\x83a\xaeJDeIH\xfd\x89" +
	"\xfa\xc0\x93L\xe2\x81\xa7\x02B\xbe\xda]AU\xe9\xff" +
	"r]/y\xe4\x8dj\xae\xf8\xc0\xc2:B\xc6\xe4\xad" +
	"\x1a6\xa6\xa0\x12a\x94]\"*\x9f\x14&I\xdbu" +
	"\xdb/\x03'K\xd6'\xef\x06x\x17h\x9aMci" +
	"\xcd\xb2\x99\xc1\xe5\x8bj\xc6h6;\xfb+\xc8\xbc\x1d" +
	"\xa4z\x8c$LR\xd7O\x96u\x93\xe5\xfd\xe4\xc2n" +
	"\xd22F\xdaorF\x12\x96\xaa\xa7,\x1b\x0d3\x93" +
	"P\xb3\xd9\x04\xd68F2a\x1b\xb6\x9a\x16\x12i\xaa" +
	"\x0f\xdaC`$\x13\xc9\xa1\x9c>\x1cOX\xda&\xea" +
	"\x8c$\x8cd\"ck f(#^\xab\xa5\xa9\x06" +
	"\x98r<:\x10i\xca\x19I\x0c\x0d\xaa\xe6`\x02\xb0" +
	"\xc6\x19Id\x87\xd4\x01S\x13\x93\xaam\x98\x89\x1a\x18" +
	"I\xe4\x9bg4\x95`\x96W8\x7fX\xe1\x97\x97\x17" +
	"S+\xd9djY\xdb\x98\x88\x85\xf2\x00\x0b\x85\xa0 " +
	"j)\x94@@\x090nkvz\xfa\x00d\x86\\" +
	"\xcb\xf2\x04\xc9\xa4\xdc\xd3\xf2&\x8b0\xc1\xd7V\xe0{" +
	"N\xfe\x9e\xf9\xb2\x8e\x00\x0b\x15\xc2I6\x1c\x99\x80\x86" +
	"|\xe3\xc4Q\x07,#\x9d\xb39(\xa2 `\x94w" +
	"\x13\xd3\xaa\xadm\xe0c\xbcv\x9ev<o\xd7\x93q" +
	"\x8e\x9f\x09\xaa1\xa7\x88\xe9y\x18\xb6\x81;{\xea\xb8" +
	"\x1e=e\x12\x14\x94\xc4\xdd\xd4\xca\xa5E{\\\x0b\xa4" +
	"1h\x814Y\x9c\x12IP\xf5\x01\"\x99r\xafn" +
	"\"\xca\x03\x9f\x98\xf90\x05\xbeit\x09\xb8Jl<" +
	"}\x95\x14^\xc3\xab\x13x\x95\xe0\xa7\x04Su\"Z" +
	"C!\xd3\xd3\x8c\xd6\x91\x0f\x99\xd7\x87:\x11[\xd7\x05" +
	"\xef+\x0a\xa8\xab\xa0R\xb2\xa6a\x1bI#\xcd0<" +
	"\xce\x0da\xfev\xbb\xc7\x86X~\xb3~\xa1?\xad\xf3" +
	"\xf7\x0a\x9a\xaa.5\xc6[a\xa7\xd3\xce\xad\xff\x9b\xdb" +
	"\xb9\xf1\x14\xab\xe31\x16t \xdc\x14\xe0\xb4\x812\xfe" +
	"\x10}\xe3)\x90\x99_\x1c\xda]\xcb\xba\x00\xf4\x93)" +
	"\xdd}Dp\xdaJ\x1f\xa4\xf6\xb4[r!\x8b\xf3\xf4" +
	"\xdd\xd9\x7f\x8a\x8e\xdc$\xda\xd2\x06u\xc3\xa4\xad,\x7f" +
	"\x1e}_\xb6T\xb8\xb4l\xefUq\xa2\xeb\xab/\x10" +
	"\x15j\x02w8>'\x88m\xa4\xea\xf04\xef%\xc7" +
	";\xbb\xe9\x1a\xdd\x87\xaa\xfd\x179\x05\x10y\x81\xcf\xb6" +
	"$1\xfb\x9a*\xddj\x0d\xa5[\xd3<\xdb\xc9\x05p" +
	";\x0f\xd3\xf6\xf0\xdd\x1f\xa8\x05\xccX\xf7\x1a\xc3\x14P" +
	"\xff\x00\x82\x16\xc6[^C\xd3\xc9\x0e\xbbC\x97\xe8\x9e" +
	"\x90}w\x01(W\xb9\xe7\xee\x0c\xa8\x16\xedRY\xfe" +
	"\x13\x9cS\xce\xa2f\xcb\xa0[kxc\xeec\x0fu" +
	"\x00\xd3\xf4r^\x8eJ\x83\xc1\xf3\"\xcf\xda\xa7Hm" +
	"s~f\xfb!Nl\x83&\x81\x7f\xef\x8f\xa6\xf7x" +
	"Bbh\xec\x9f\xf0xb\x06\x00\xa9\x1b\x03P\x96\x88" +
	"\xa8\xfc\x7f\x01'\x94\xe0\xc1\x8b5\xaf\x08\xa7V\x92e" +
	"\x92\xa8\x19\xfa\x0a;\x93~\xbfdES]2\xa9\x99" +
	"Z\xef\x02I\xb7\xa8w;3\x05lVf\xed\x18\xb3" +
	"0\xa6\xfc\xbc\xbf@,P\xc2\"N(a\x99\xf2\x11" +
	"\xc9\xfa\xee\xc05`\xde]\x8c\x9a\xf9\x0av\xa7\x80N" +
	"J\xdd\xd0\x96V-\x0b\xb0\xcewf\xc1X}\x81\xb1" +
	"\x86\xf0X\xfb\x88Mu\x883Y? \xb2\xdc\xc6\xa4" +
	"o\xb9\xff\x8b\x9ec*\xcc\xb8\xf0\xab\x1d0\xc4\xd4(" +
	"\x93\xe3\x1c\xc7\x89\xb8\xdd\xc2pD`9{\xfe\x0a4" +
	"\xd1\x18\x0a\x0a\x03\xa36\xb5<\xec\x06Y\xd9K\xadW" +
	"_\xbd\xbf\xea\xcd\xdb\xa6\x97\x95\xf9\xc2\x04\xa5Y\x9c\x9e" +
	"\xaa4\xfb\x08\x89\xee %\xfd\x05\xab\xb0qU\xd6\xe4" +
	"\x07\x10\xba\x8f\xe4\x81>\xb8\x01\xf5\xfe\x03\x0bz\xff\xc3" +
	"\x87\x90n\xf7\x06\xd4\xbb\xb3\x04\xb4\x9a\xb1\x0b\xa7\xbd-" +
	"\xef]\"\xb7\xcb\x88\x18)\xc52,f\x19\xed\x18\xef" +
	"\x80\xa2r\xb60\x1d\x1b\x8b\x9c*m>\x9d\x84\xe3\x03" +
	"\xa5\xf8\xfe\xdd\xd5\xb4S\xfc\xa2\xc95\xe2=\x92\xf0\xde" +
	"H\xb8\xefj\\\xcd\x9c\xc94\xb3.\xa4\x19~\x03\xce" +
	"<\"\xbf\x8c\x9a\xa4\xfdw\xaa6\xba\x1ec\xd2*3" +
	"0\xfc\x7fH\x8a\x1bC\xff\xed#:v\xea\xe7\xb5g" +
	"\xfbJ\xdd\xcdT\xb5\xd3m{\xfaJ=\xb0\x09@\xd9" +
	"\xef^\xfd\xfbJ=\xc8l\xe4A\x11\x95\xa7\x04$\xd1" +
	"\xbc\xe79\xb6\x03@yJD\xe5\xb7\x02b\x04C\xff" +
	"\x0b\x8b\xbcr\x13\x08D\x8c\xba\xb0x\x9eq\xfc\xa9\x88" +
	"\xcak\x02S5+\xaaCOO\xab\x96>\xf4\xc0\xaf" +
	"\xb7~mg\xf0,6\x7f\xe3\xd6\xd4\xe3Y_p+" +
	"\xcf\xad\xaf\xc9\x8d\x84\x01\x0b\xff\x8e1\xcfb\xd2\xa7\xae" +
	"\xbe\xc6\xf2\xc4\xdec\x0a\xf4\xce\xcf\x7fW\xd9\xaeC\x93" +
	"\x1b\xbb\x82\xd9\xfe\x0d\xe6\xf4\x1e\xca\x0e'\x8d\xf4\x87\xb0" +
	"S[4=\x1f\xea\xf5\xaa5\x03u\xe6Jy\xf7\xa3" +
	"\x84y\xd2\xc6\xf1\x9e4Z\xc6}Y\xa2#\xe8,4" +
	"\xe9\x86\x99Q\xd3P\xe4\xa4\x8c\x8dz\xdaPS\xa1\x04" +
	"\xe0\x7f\x02\x00\x00\xff\xff4\xd2\xe2\x15"

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
		0xbda585bffe1dc7e8,
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
	0, 0, 0, 0, 30, 0, 0, 0,
	1, 0, 0, 0, 78, 0, 0, 0,
	33, 0, 0, 0, 146, 0, 0, 0,
	41, 0, 0, 0, 130, 0, 0, 0,
	45, 0, 0, 0, 114, 0, 0, 0,
	49, 0, 0, 0, 90, 0, 0, 0,
	53, 0, 0, 0, 82, 0, 0, 0,
	57, 0, 0, 0, 90, 0, 0, 0,
	61, 0, 0, 0, 82, 0, 0, 0,
	65, 0, 0, 0, 130, 0, 0, 0,
	69, 0, 0, 0, 138, 0, 0, 0,
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
	120, 45, 112, 104, 97, 98, 114, 105,
	99, 97, 116, 111, 114, 45, 42, 0,
	120, 45, 114, 101, 113, 117, 101, 115,
	116, 101, 100, 45, 119, 105, 116, 104,
	0, 0, 0, 0, 0, 0, 0, 0,
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
