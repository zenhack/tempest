package websession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	math "math"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

const HttpStatus = uint64(0xaf480a0c6cab8887)

// Constants defined in web-session.capnp.
const (
	HttpStatusAnnotationId = uint64(12630356203439622279)
)

type HttpStatusDescriptor struct{ capnp.Struct }

func NewHttpStatusDescriptor(s *capnp.Segment) (HttpStatusDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return HttpStatusDescriptor{}, err
	}
	return HttpStatusDescriptor{st}, nil
}

func NewRootHttpStatusDescriptor(s *capnp.Segment) (HttpStatusDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return HttpStatusDescriptor{}, err
	}
	return HttpStatusDescriptor{st}, nil
}

func ReadRootHttpStatusDescriptor(msg *capnp.Message) (HttpStatusDescriptor, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return HttpStatusDescriptor{}, err
	}
	return HttpStatusDescriptor{root.Struct()}, nil
}
func (s HttpStatusDescriptor) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s HttpStatusDescriptor) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s HttpStatusDescriptor) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s HttpStatusDescriptor) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HttpStatusDescriptor) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return HttpStatusDescriptor_List{}, err
	}
	return HttpStatusDescriptor_List{l}, nil
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
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
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

type WebSession_Server interface {
	Get(WebSession_get) error

	Post(WebSession_post) error

	OpenWebSocket(WebSession_openWebSocket) error

	Put(WebSession_put) error

	Delete(WebSession_delete) error

	PostStreaming(WebSession_postStreaming) error

	PutStreaming(WebSession_putStreaming) error
}

func WebSession_ServerToClient(s WebSession_Server) WebSession {
	c, _ := s.(server.Closer)
	return WebSession{Client: server.New(WebSession_Methods(nil, s), c)}
}

func WebSession_Methods(methods []server.Method, s WebSession_Server) []server.Method {
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
			call := WebSession_get{c, opts, WebSession_get_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_post{c, opts, WebSession_post_Params{Struct: p}, WebSession_Response{Struct: r}}
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
			call := WebSession_delete{c, opts, WebSession_delete_Params{Struct: p}, WebSession_Response{Struct: r}}
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

type WebSession_Params struct{ capnp.Struct }

func NewWebSession_Params(s *capnp.Segment) (WebSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_Params{}, err
	}
	return WebSession_Params{st}, nil
}

func NewRootWebSession_Params(s *capnp.Segment) (WebSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_Params{}, err
	}
	return WebSession_Params{st}, nil
}

func ReadRootWebSession_Params(msg *capnp.Message) (WebSession_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_Params{}, err
	}
	return WebSession_Params{root.Struct()}, nil
}
func (s WebSession_Params) BasePath() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Params) HasBasePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Params) BasePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Params) HasUserAgent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Params) UserAgentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
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
	if err != nil {
		return WebSession_Params_List{}, err
	}
	return WebSession_Params_List{l}, nil
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

func NewWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{st}, nil
}

func NewRootWebSession_Context(s *capnp.Segment) (WebSession_Context, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{st}, nil
}

func ReadRootWebSession_Context(msg *capnp.Message) (WebSession_Context, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{root.Struct()}, nil
}
func (s WebSession_Context) Cookies() (util.KeyValue_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.KeyValue_List{}, err
	}
	return util.KeyValue_List{List: p.List()}, nil
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
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s WebSession_Context) HasResponseStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Context) SetResponseStream(v util.ByteStream) error {
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s WebSession_Context) Accept() (WebSession_AcceptedType_List, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return WebSession_AcceptedType_List{}, err
	}
	return WebSession_AcceptedType_List{List: p.List()}, nil
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

// WebSession_Context_List is a list of WebSession_Context.
type WebSession_Context_List struct{ capnp.List }

// NewWebSession_Context creates a new list of WebSession_Context.
func NewWebSession_Context_List(s *capnp.Segment, sz int32) (WebSession_Context_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return WebSession_Context_List{}, err
	}
	return WebSession_Context_List{l}, nil
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

type WebSession_PostContent struct{ capnp.Struct }

func NewWebSession_PostContent(s *capnp.Segment) (WebSession_PostContent, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_PostContent{}, err
	}
	return WebSession_PostContent{st}, nil
}

func NewRootWebSession_PostContent(s *capnp.Segment) (WebSession_PostContent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_PostContent{}, err
	}
	return WebSession_PostContent{st}, nil
}

func ReadRootWebSession_PostContent(msg *capnp.Message) (WebSession_PostContent, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_PostContent{}, err
	}
	return WebSession_PostContent{root.Struct()}, nil
}
func (s WebSession_PostContent) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_PostContent) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_PostContent) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_PostContent) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_PostContent) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_PostContent_List{}, err
	}
	return WebSession_PostContent_List{l}, nil
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
	if err != nil {
		return WebSession_PutContent{}, err
	}
	return WebSession_PutContent{st}, nil
}

func NewRootWebSession_PutContent(s *capnp.Segment) (WebSession_PutContent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_PutContent{}, err
	}
	return WebSession_PutContent{st}, nil
}

func ReadRootWebSession_PutContent(msg *capnp.Message) (WebSession_PutContent, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_PutContent{}, err
	}
	return WebSession_PutContent{root.Struct()}, nil
}
func (s WebSession_PutContent) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_PutContent) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_PutContent) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_PutContent) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_PutContent) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_PutContent_List{}, err
	}
	return WebSession_PutContent_List{l}, nil
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
	if err != nil {
		return WebSession_Cookie{}, err
	}
	return WebSession_Cookie{st}, nil
}

func NewRootWebSession_Cookie(s *capnp.Segment) (WebSession_Cookie, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 3})
	if err != nil {
		return WebSession_Cookie{}, err
	}
	return WebSession_Cookie{st}, nil
}

func ReadRootWebSession_Cookie(msg *capnp.Message) (WebSession_Cookie, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_Cookie{}, err
	}
	return WebSession_Cookie{root.Struct()}, nil
}
func (s WebSession_Cookie) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Cookie) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Cookie) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Cookie) HasPath() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Cookie) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Cookie_List{}, err
	}
	return WebSession_Cookie_List{l}, nil
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
	if err != nil {
		return WebSession_AcceptedType{}, err
	}
	return WebSession_AcceptedType{st}, nil
}

func NewRootWebSession_AcceptedType(s *capnp.Segment) (WebSession_AcceptedType, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return WebSession_AcceptedType{}, err
	}
	return WebSession_AcceptedType{st}, nil
}

func ReadRootWebSession_AcceptedType(msg *capnp.Message) (WebSession_AcceptedType, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_AcceptedType{}, err
	}
	return WebSession_AcceptedType{root.Struct()}, nil
}
func (s WebSession_AcceptedType) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_AcceptedType) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_AcceptedType) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_AcceptedType_List{}, err
	}
	return WebSession_AcceptedType_List{l}, nil
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
type WebSession_Response_redirect WebSession_Response
type WebSession_Response_clientError WebSession_Response
type WebSession_Response_serverError WebSession_Response
type WebSession_Response_Which uint16

const (
	WebSession_Response_Which_content     WebSession_Response_Which = 1
	WebSession_Response_Which_noContent   WebSession_Response_Which = 4
	WebSession_Response_Which_redirect    WebSession_Response_Which = 0
	WebSession_Response_Which_clientError WebSession_Response_Which = 2
	WebSession_Response_Which_serverError WebSession_Response_Which = 3
)

func (w WebSession_Response_Which) String() string {
	const s = "contentnoContentredirectclientErrorserverError"
	switch w {
	case WebSession_Response_Which_content:
		return s[0:7]
	case WebSession_Response_Which_noContent:
		return s[7:16]
	case WebSession_Response_Which_redirect:
		return s[16:24]
	case WebSession_Response_Which_clientError:
		return s[24:35]
	case WebSession_Response_Which_serverError:
		return s[35:46]

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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	if err != nil {
		return WebSession_Response{}, err
	}
	return WebSession_Response{st}, nil
}

func NewRootWebSession_Response(s *capnp.Segment) (WebSession_Response, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	if err != nil {
		return WebSession_Response{}, err
	}
	return WebSession_Response{st}, nil
}

func ReadRootWebSession_Response(msg *capnp.Message) (WebSession_Response, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_Response{}, err
	}
	return WebSession_Response{root.Struct()}, nil
}

func (s WebSession_Response) Which() WebSession_Response_Which {
	return WebSession_Response_Which(s.Struct.Uint16(2))
}
func (s WebSession_Response) SetCookies() (WebSession_Cookie_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return WebSession_Cookie_List{}, err
	}
	return WebSession_Cookie_List{List: p.List()}, nil
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
	if err != nil {
		return WebSession_CachePolicy{}, err
	}
	return WebSession_CachePolicy{Struct: p.Struct()}, nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_content) HasEncoding() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_content) HasLanguage() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) LanguageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_content) HasMimeType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s WebSession_Response_content) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s WebSession_Response_content) Body() WebSession_Response_content_body {
	return WebSession_Response_content_body(s)
}
func (s WebSession_Response_content_body) Which() WebSession_Response_content_body_Which {
	return WebSession_Response_content_body_Which(s.Struct.Uint16(0))
}
func (s WebSession_Response_content_body) Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	p, err := s.Struct.Ptr(4)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s WebSession_Response_content_body) HasStream() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content_body) SetStream(v util.Handle) error {
	s.Struct.SetUint16(0, 1)
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_content_disposition) HasDownload() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_content_disposition) DownloadBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_redirect) HasLocation() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_redirect) LocationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_clientError) HasDescriptionHtml() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_clientError) DescriptionHtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_Response_serverError) HasDescriptionHtml() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_Response_serverError) DescriptionHtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s WebSession_Response_serverError) SetDescriptionHtml(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// WebSession_Response_List is a list of WebSession_Response.
type WebSession_Response_List struct{ capnp.List }

// NewWebSession_Response creates a new list of WebSession_Response.
func NewWebSession_Response_List(s *capnp.Segment, sz int32) (WebSession_Response_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7}, sz)
	if err != nil {
		return WebSession_Response_List{}, err
	}
	return WebSession_Response_List{l}, nil
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
	WebSession_Response_SuccessCode_ok       WebSession_Response_SuccessCode = 0
	WebSession_Response_SuccessCode_created  WebSession_Response_SuccessCode = 1
	WebSession_Response_SuccessCode_accepted WebSession_Response_SuccessCode = 2
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

	default:
		return 0
	}
}

type WebSession_Response_SuccessCode_List struct{ capnp.List }

func NewWebSession_Response_SuccessCode_List(s *capnp.Segment, sz int32) (WebSession_Response_SuccessCode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return WebSession_Response_SuccessCode_List{}, err
	}
	return WebSession_Response_SuccessCode_List{l.List}, nil
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
	WebSession_Response_ClientErrorCode_requestEntityTooLarge WebSession_Response_ClientErrorCode = 7
	WebSession_Response_ClientErrorCode_requestUriTooLong     WebSession_Response_ClientErrorCode = 8
	WebSession_Response_ClientErrorCode_unsupportedMediaType  WebSession_Response_ClientErrorCode = 9
	WebSession_Response_ClientErrorCode_imATeapot             WebSession_Response_ClientErrorCode = 10
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
	case WebSession_Response_ClientErrorCode_requestEntityTooLarge:
		return "requestEntityTooLarge"
	case WebSession_Response_ClientErrorCode_requestUriTooLong:
		return "requestUriTooLong"
	case WebSession_Response_ClientErrorCode_unsupportedMediaType:
		return "unsupportedMediaType"
	case WebSession_Response_ClientErrorCode_imATeapot:
		return "imATeapot"

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
	case "requestEntityTooLarge":
		return WebSession_Response_ClientErrorCode_requestEntityTooLarge
	case "requestUriTooLong":
		return WebSession_Response_ClientErrorCode_requestUriTooLong
	case "unsupportedMediaType":
		return WebSession_Response_ClientErrorCode_unsupportedMediaType
	case "imATeapot":
		return WebSession_Response_ClientErrorCode_imATeapot

	default:
		return 0
	}
}

type WebSession_Response_ClientErrorCode_List struct{ capnp.List }

func NewWebSession_Response_ClientErrorCode_List(s *capnp.Segment, sz int32) (WebSession_Response_ClientErrorCode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return WebSession_Response_ClientErrorCode_List{}, err
	}
	return WebSession_Response_ClientErrorCode_List{l.List}, nil
}

func (l WebSession_Response_ClientErrorCode_List) At(i int) WebSession_Response_ClientErrorCode {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_Response_ClientErrorCode(ul.At(i))
}

func (l WebSession_Response_ClientErrorCode_List) Set(i int, v WebSession_Response_ClientErrorCode) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
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
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 7},
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
	if err != nil {
		return WebSession_RequestStream_getResponse_Params{}, err
	}
	return WebSession_RequestStream_getResponse_Params{st}, nil
}

func NewRootWebSession_RequestStream_getResponse_Params(s *capnp.Segment) (WebSession_RequestStream_getResponse_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return WebSession_RequestStream_getResponse_Params{}, err
	}
	return WebSession_RequestStream_getResponse_Params{st}, nil
}

func ReadRootWebSession_RequestStream_getResponse_Params(msg *capnp.Message) (WebSession_RequestStream_getResponse_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_RequestStream_getResponse_Params{}, err
	}
	return WebSession_RequestStream_getResponse_Params{root.Struct()}, nil
}

// WebSession_RequestStream_getResponse_Params_List is a list of WebSession_RequestStream_getResponse_Params.
type WebSession_RequestStream_getResponse_Params_List struct{ capnp.List }

// NewWebSession_RequestStream_getResponse_Params creates a new list of WebSession_RequestStream_getResponse_Params.
func NewWebSession_RequestStream_getResponse_Params_List(s *capnp.Segment, sz int32) (WebSession_RequestStream_getResponse_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return WebSession_RequestStream_getResponse_Params_List{}, err
	}
	return WebSession_RequestStream_getResponse_Params_List{l}, nil
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
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Params{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Params{st}, nil
}

func NewRootWebSession_WebSocketStream_sendBytes_Params(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Params{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Params{st}, nil
}

func ReadRootWebSession_WebSocketStream_sendBytes_Params(msg *capnp.Message) (WebSession_WebSocketStream_sendBytes_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Params{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Params{root.Struct()}, nil
}
func (s WebSession_WebSocketStream_sendBytes_Params) Message() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Params_List{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Params_List{l}, nil
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
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Results{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Results{st}, nil
}

func NewRootWebSession_WebSocketStream_sendBytes_Results(s *capnp.Segment) (WebSession_WebSocketStream_sendBytes_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Results{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Results{st}, nil
}

func ReadRootWebSession_WebSocketStream_sendBytes_Results(msg *capnp.Message) (WebSession_WebSocketStream_sendBytes_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Results{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Results{root.Struct()}, nil
}

// WebSession_WebSocketStream_sendBytes_Results_List is a list of WebSession_WebSocketStream_sendBytes_Results.
type WebSession_WebSocketStream_sendBytes_Results_List struct{ capnp.List }

// NewWebSession_WebSocketStream_sendBytes_Results creates a new list of WebSession_WebSocketStream_sendBytes_Results.
func NewWebSession_WebSocketStream_sendBytes_Results_List(s *capnp.Segment, sz int32) (WebSession_WebSocketStream_sendBytes_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return WebSession_WebSocketStream_sendBytes_Results_List{}, err
	}
	return WebSession_WebSocketStream_sendBytes_Results_List{l}, nil
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
	if err != nil {
		return WebSession_CachePolicy{}, err
	}
	return WebSession_CachePolicy{st}, nil
}

func NewRootWebSession_CachePolicy(s *capnp.Segment) (WebSession_CachePolicy, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return WebSession_CachePolicy{}, err
	}
	return WebSession_CachePolicy{st}, nil
}

func ReadRootWebSession_CachePolicy(msg *capnp.Message) (WebSession_CachePolicy, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_CachePolicy{}, err
	}
	return WebSession_CachePolicy{root.Struct()}, nil
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
	if err != nil {
		return WebSession_CachePolicy_List{}, err
	}
	return WebSession_CachePolicy_List{l}, nil
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
	if err != nil {
		return WebSession_CachePolicy_Scope_List{}, err
	}
	return WebSession_CachePolicy_Scope_List{l.List}, nil
}

func (l WebSession_CachePolicy_Scope_List) At(i int) WebSession_CachePolicy_Scope {
	ul := capnp.UInt16List{List: l.List}
	return WebSession_CachePolicy_Scope(ul.At(i))
}

func (l WebSession_CachePolicy_Scope_List) Set(i int, v WebSession_CachePolicy_Scope) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type WebSession_get_Params struct{ capnp.Struct }

func NewWebSession_get_Params(s *capnp.Segment) (WebSession_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return WebSession_get_Params{}, err
	}
	return WebSession_get_Params{st}, nil
}

func NewRootWebSession_get_Params(s *capnp.Segment) (WebSession_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return WebSession_get_Params{}, err
	}
	return WebSession_get_Params{st}, nil
}

func ReadRootWebSession_get_Params(msg *capnp.Message) (WebSession_get_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_get_Params{}, err
	}
	return WebSession_get_Params{root.Struct()}, nil
}
func (s WebSession_get_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_get_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_get_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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

// WebSession_get_Params_List is a list of WebSession_get_Params.
type WebSession_get_Params_List struct{ capnp.List }

// NewWebSession_get_Params creates a new list of WebSession_get_Params.
func NewWebSession_get_Params_List(s *capnp.Segment, sz int32) (WebSession_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return WebSession_get_Params_List{}, err
	}
	return WebSession_get_Params_List{l}, nil
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
	if err != nil {
		return WebSession_post_Params{}, err
	}
	return WebSession_post_Params{st}, nil
}

func NewRootWebSession_post_Params(s *capnp.Segment) (WebSession_post_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_post_Params{}, err
	}
	return WebSession_post_Params{st}, nil
}

func ReadRootWebSession_post_Params(msg *capnp.Message) (WebSession_post_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_post_Params{}, err
	}
	return WebSession_post_Params{root.Struct()}, nil
}
func (s WebSession_post_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_post_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_post_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_PostContent{}, err
	}
	return WebSession_PostContent{Struct: p.Struct()}, nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return WebSession_post_Params_List{}, err
	}
	return WebSession_post_Params_List{l}, nil
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
	if err != nil {
		return WebSession_openWebSocket_Params{}, err
	}
	return WebSession_openWebSocket_Params{st}, nil
}

func NewRootWebSession_openWebSocket_Params(s *capnp.Segment) (WebSession_openWebSocket_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return WebSession_openWebSocket_Params{}, err
	}
	return WebSession_openWebSocket_Params{st}, nil
}

func ReadRootWebSession_openWebSocket_Params(msg *capnp.Message) (WebSession_openWebSocket_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_openWebSocket_Params{}, err
	}
	return WebSession_openWebSocket_Params{root.Struct()}, nil
}
func (s WebSession_openWebSocket_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_openWebSocket_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
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
	p, err := s.Struct.Ptr(3)
	if err != nil {

		return WebSession_WebSocketStream{}
	}
	return WebSession_WebSocketStream{Client: p.Interface().Client()}
}

func (s WebSession_openWebSocket_Params) HasClientStream() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Params) SetClientStream(v WebSession_WebSocketStream) error {
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(3, in.ToPtr())
}

// WebSession_openWebSocket_Params_List is a list of WebSession_openWebSocket_Params.
type WebSession_openWebSocket_Params_List struct{ capnp.List }

// NewWebSession_openWebSocket_Params creates a new list of WebSession_openWebSocket_Params.
func NewWebSession_openWebSocket_Params_List(s *capnp.Segment, sz int32) (WebSession_openWebSocket_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	if err != nil {
		return WebSession_openWebSocket_Params_List{}, err
	}
	return WebSession_openWebSocket_Params_List{l}, nil
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
	if err != nil {
		return WebSession_openWebSocket_Results{}, err
	}
	return WebSession_openWebSocket_Results{st}, nil
}

func NewRootWebSession_openWebSocket_Results(s *capnp.Segment) (WebSession_openWebSocket_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return WebSession_openWebSocket_Results{}, err
	}
	return WebSession_openWebSocket_Results{st}, nil
}

func ReadRootWebSession_openWebSocket_Results(msg *capnp.Message) (WebSession_openWebSocket_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_openWebSocket_Results{}, err
	}
	return WebSession_openWebSocket_Results{root.Struct()}, nil
}
func (s WebSession_openWebSocket_Results) Protocol() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
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
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return WebSession_WebSocketStream{}
	}
	return WebSession_WebSocketStream{Client: p.Interface().Client()}
}

func (s WebSession_openWebSocket_Results) HasServerStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_openWebSocket_Results) SetServerStream(v WebSession_WebSocketStream) error {
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

// WebSession_openWebSocket_Results_List is a list of WebSession_openWebSocket_Results.
type WebSession_openWebSocket_Results_List struct{ capnp.List }

// NewWebSession_openWebSocket_Results creates a new list of WebSession_openWebSocket_Results.
func NewWebSession_openWebSocket_Results_List(s *capnp.Segment, sz int32) (WebSession_openWebSocket_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return WebSession_openWebSocket_Results_List{}, err
	}
	return WebSession_openWebSocket_Results_List{l}, nil
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
	if err != nil {
		return WebSession_put_Params{}, err
	}
	return WebSession_put_Params{st}, nil
}

func NewRootWebSession_put_Params(s *capnp.Segment) (WebSession_put_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return WebSession_put_Params{}, err
	}
	return WebSession_put_Params{st}, nil
}

func ReadRootWebSession_put_Params(msg *capnp.Message) (WebSession_put_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_put_Params{}, err
	}
	return WebSession_put_Params{root.Struct()}, nil
}
func (s WebSession_put_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_put_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_put_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_PutContent{}, err
	}
	return WebSession_PutContent{Struct: p.Struct()}, nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return WebSession_put_Params_List{}, err
	}
	return WebSession_put_Params_List{l}, nil
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
	if err != nil {
		return WebSession_delete_Params{}, err
	}
	return WebSession_delete_Params{st}, nil
}

func NewRootWebSession_delete_Params(s *capnp.Segment) (WebSession_delete_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return WebSession_delete_Params{}, err
	}
	return WebSession_delete_Params{st}, nil
}

func ReadRootWebSession_delete_Params(msg *capnp.Message) (WebSession_delete_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_delete_Params{}, err
	}
	return WebSession_delete_Params{root.Struct()}, nil
}
func (s WebSession_delete_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_delete_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_delete_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return WebSession_delete_Params_List{}, err
	}
	return WebSession_delete_Params_List{l}, nil
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
	if err != nil {
		return WebSession_postStreaming_Params{}, err
	}
	return WebSession_postStreaming_Params{st}, nil
}

func NewRootWebSession_postStreaming_Params(s *capnp.Segment) (WebSession_postStreaming_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return WebSession_postStreaming_Params{}, err
	}
	return WebSession_postStreaming_Params{st}, nil
}

func ReadRootWebSession_postStreaming_Params(msg *capnp.Message) (WebSession_postStreaming_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_postStreaming_Params{}, err
	}
	return WebSession_postStreaming_Params{root.Struct()}, nil
}
func (s WebSession_postStreaming_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_postStreaming_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_postStreaming_Params) HasMimeType() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_postStreaming_Params) HasEncoding() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Params) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_postStreaming_Params_List{}, err
	}
	return WebSession_postStreaming_Params_List{l}, nil
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
	if err != nil {
		return WebSession_postStreaming_Results{}, err
	}
	return WebSession_postStreaming_Results{st}, nil
}

func NewRootWebSession_postStreaming_Results(s *capnp.Segment) (WebSession_postStreaming_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSession_postStreaming_Results{}, err
	}
	return WebSession_postStreaming_Results{st}, nil
}

func ReadRootWebSession_postStreaming_Results(msg *capnp.Message) (WebSession_postStreaming_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_postStreaming_Results{}, err
	}
	return WebSession_postStreaming_Results{root.Struct()}, nil
}
func (s WebSession_postStreaming_Results) Stream() WebSession_RequestStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return WebSession_RequestStream{}
	}
	return WebSession_RequestStream{Client: p.Interface().Client()}
}

func (s WebSession_postStreaming_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_postStreaming_Results) SetStream(v WebSession_RequestStream) error {
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

// WebSession_postStreaming_Results_List is a list of WebSession_postStreaming_Results.
type WebSession_postStreaming_Results_List struct{ capnp.List }

// NewWebSession_postStreaming_Results creates a new list of WebSession_postStreaming_Results.
func NewWebSession_postStreaming_Results_List(s *capnp.Segment, sz int32) (WebSession_postStreaming_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSession_postStreaming_Results_List{}, err
	}
	return WebSession_postStreaming_Results_List{l}, nil
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
	if err != nil {
		return WebSession_putStreaming_Params{}, err
	}
	return WebSession_putStreaming_Params{st}, nil
}

func NewRootWebSession_putStreaming_Params(s *capnp.Segment) (WebSession_putStreaming_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return WebSession_putStreaming_Params{}, err
	}
	return WebSession_putStreaming_Params{st}, nil
}

func ReadRootWebSession_putStreaming_Params(msg *capnp.Message) (WebSession_putStreaming_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_putStreaming_Params{}, err
	}
	return WebSession_putStreaming_Params{root.Struct()}, nil
}
func (s WebSession_putStreaming_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_putStreaming_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_putStreaming_Params) HasMimeType() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_Context{}, err
	}
	return WebSession_Context{Struct: p.Struct()}, nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s WebSession_putStreaming_Params) HasEncoding() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Params) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return WebSession_putStreaming_Params_List{}, err
	}
	return WebSession_putStreaming_Params_List{l}, nil
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
	if err != nil {
		return WebSession_putStreaming_Results{}, err
	}
	return WebSession_putStreaming_Results{st}, nil
}

func NewRootWebSession_putStreaming_Results(s *capnp.Segment) (WebSession_putStreaming_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSession_putStreaming_Results{}, err
	}
	return WebSession_putStreaming_Results{st}, nil
}

func ReadRootWebSession_putStreaming_Results(msg *capnp.Message) (WebSession_putStreaming_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return WebSession_putStreaming_Results{}, err
	}
	return WebSession_putStreaming_Results{root.Struct()}, nil
}
func (s WebSession_putStreaming_Results) Stream() WebSession_RequestStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return WebSession_RequestStream{}
	}
	return WebSession_RequestStream{Client: p.Interface().Client()}
}

func (s WebSession_putStreaming_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSession_putStreaming_Results) SetStream(v WebSession_RequestStream) error {
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

// WebSession_putStreaming_Results_List is a list of WebSession_putStreaming_Results.
type WebSession_putStreaming_Results_List struct{ capnp.List }

// NewWebSession_putStreaming_Results creates a new list of WebSession_putStreaming_Results.
func NewWebSession_putStreaming_Results_List(s *capnp.Segment, sz int32) (WebSession_putStreaming_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSession_putStreaming_Results_List{}, err
	}
	return WebSession_putStreaming_Results_List{l}, nil
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
