package webpublishing

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type WebSite struct{ Client capnp.Client }

func (c WebSite) GetUrl(ctx context.Context, params func(WebSite_getUrl_Params) error, opts ...capnp.CallOption) WebSite_getUrl_Results_Promise {
	if c.Client == nil {
		return WebSite_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      0,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getUrl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getUrl_Params{Struct: s}) }
	}
	return WebSite_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetEntities(ctx context.Context, params func(WebSite_getEntities_Params) error, opts ...capnp.CallOption) WebSite_getEntities_Results_Promise {
	if c.Client == nil {
		return WebSite_getEntities_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      1,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getEntities",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getEntities_Params{Struct: s}) }
	}
	return WebSite_getEntities_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetNotFoundEntities(ctx context.Context, params func(WebSite_getNotFoundEntities_Params) error, opts ...capnp.CallOption) WebSite_getNotFoundEntities_Results_Promise {
	if c.Client == nil {
		return WebSite_getNotFoundEntities_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      2,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getNotFoundEntities",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getNotFoundEntities_Params{Struct: s}) }
	}
	return WebSite_getNotFoundEntities_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) UploadBlob(ctx context.Context, params func(WebSite_uploadBlob_Params) error, opts ...capnp.CallOption) WebSite_uploadBlob_Results_Promise {
	if c.Client == nil {
		return WebSite_uploadBlob_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      3,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "uploadBlob",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_uploadBlob_Params{Struct: s}) }
	}
	return WebSite_uploadBlob_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetSubsite(ctx context.Context, params func(WebSite_getSubsite_Params) error, opts ...capnp.CallOption) WebSite_getSubsite_Results_Promise {
	if c.Client == nil {
		return WebSite_getSubsite_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      4,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getSubsite",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getSubsite_Params{Struct: s}) }
	}
	return WebSite_getSubsite_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) ListResources(ctx context.Context, params func(WebSite_listResources_Params) error, opts ...capnp.CallOption) WebSite_listResources_Results_Promise {
	if c.Client == nil {
		return WebSite_listResources_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      5,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "listResources",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_listResources_Params{Struct: s}) }
	}
	return WebSite_listResources_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type WebSite_Server interface {
	GetUrl(WebSite_getUrl) error

	GetEntities(WebSite_getEntities) error

	GetNotFoundEntities(WebSite_getNotFoundEntities) error

	UploadBlob(WebSite_uploadBlob) error

	GetSubsite(WebSite_getSubsite) error

	ListResources(WebSite_listResources) error
}

func WebSite_ServerToClient(s WebSite_Server) WebSite {
	c, _ := s.(server.Closer)
	return WebSite{Client: server.New(WebSite_Methods(nil, s), c)}
}

func WebSite_Methods(methods []server.Method, s WebSite_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      0,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getUrl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getUrl{c, opts, WebSite_getUrl_Params{Struct: p}, WebSite_getUrl_Results{Struct: r}}
			return s.GetUrl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      1,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getEntities",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getEntities{c, opts, WebSite_getEntities_Params{Struct: p}, WebSite_getEntities_Results{Struct: r}}
			return s.GetEntities(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      2,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getNotFoundEntities",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getNotFoundEntities{c, opts, WebSite_getNotFoundEntities_Params{Struct: p}, WebSite_getNotFoundEntities_Results{Struct: r}}
			return s.GetNotFoundEntities(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      3,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "uploadBlob",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_uploadBlob{c, opts, WebSite_uploadBlob_Params{Struct: p}, WebSite_uploadBlob_Results{Struct: r}}
			return s.UploadBlob(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      4,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getSubsite",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getSubsite{c, opts, WebSite_getSubsite_Params{Struct: p}, WebSite_getSubsite_Results{Struct: r}}
			return s.GetSubsite(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      5,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "listResources",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_listResources{c, opts, WebSite_listResources_Params{Struct: p}, WebSite_listResources_Results{Struct: r}}
			return s.ListResources(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// WebSite_getUrl holds the arguments for a server call to WebSite.getUrl.
type WebSite_getUrl struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getUrl_Params
	Results WebSite_getUrl_Results
}

// WebSite_getEntities holds the arguments for a server call to WebSite.getEntities.
type WebSite_getEntities struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getEntities_Params
	Results WebSite_getEntities_Results
}

// WebSite_getNotFoundEntities holds the arguments for a server call to WebSite.getNotFoundEntities.
type WebSite_getNotFoundEntities struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getNotFoundEntities_Params
	Results WebSite_getNotFoundEntities_Results
}

// WebSite_uploadBlob holds the arguments for a server call to WebSite.uploadBlob.
type WebSite_uploadBlob struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_uploadBlob_Params
	Results WebSite_uploadBlob_Results
}

// WebSite_getSubsite holds the arguments for a server call to WebSite.getSubsite.
type WebSite_getSubsite struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getSubsite_Params
	Results WebSite_getSubsite_Results
}

// WebSite_listResources holds the arguments for a server call to WebSite.listResources.
type WebSite_listResources struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_listResources_Params
	Results WebSite_listResources_Results
}

type WebSite_Entity struct{ capnp.Struct }
type WebSite_Entity_body WebSite_Entity
type WebSite_Entity_body_Which uint16

const (
	WebSite_Entity_body_Which_bytes WebSite_Entity_body_Which = 0
	WebSite_Entity_body_Which_blob  WebSite_Entity_body_Which = 1
)

func (w WebSite_Entity_body_Which) String() string {
	const s = "bytesblob"
	switch w {
	case WebSite_Entity_body_Which_bytes:
		return s[0:5]
	case WebSite_Entity_body_Which_blob:
		return s[5:9]

	}
	return "WebSite_Entity_body_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewWebSite_Entity(s *capnp.Segment) (WebSite_Entity, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSite_Entity{st}, err
}

func NewRootWebSite_Entity(s *capnp.Segment) (WebSite_Entity, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSite_Entity{st}, err
}

func ReadRootWebSite_Entity(msg *capnp.Message) (WebSite_Entity, error) {
	root, err := msg.RootPtr()
	return WebSite_Entity{root.Struct()}, err
}

func (s WebSite_Entity) String() string {
	str, _ := text.Marshal(0xd019dd3e3c0e7e68, s.Struct)
	return str
}

func (s WebSite_Entity) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_Entity) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSite_Entity) Language() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSite_Entity) HasLanguage() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) LanguageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetLanguage(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSite_Entity) Encoding() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSite_Entity) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s WebSite_Entity) Body() WebSite_Entity_body { return WebSite_Entity_body(s) }

func (s WebSite_Entity_body) Which() WebSite_Entity_body_Which {
	return WebSite_Entity_body_Which(s.Struct.Uint16(0))
}
func (s WebSite_Entity_body) Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s WebSite_Entity_body) HasBytes() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity_body) SetBytes(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

func (s WebSite_Entity_body) Blob() util.Blob {
	p, _ := s.Struct.Ptr(3)
	return util.Blob{Client: p.Interface().Client()}
}

func (s WebSite_Entity_body) HasBlob() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity_body) SetBlob(v util.Blob) error {
	s.Struct.SetUint16(0, 1)
	if v.Client == nil {
		return s.Struct.SetPtr(3, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(3, in.ToPtr())
}

func (s WebSite_Entity) RedirectTo() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s WebSite_Entity) HasRedirectTo() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) RedirectToBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetRedirectTo(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

// WebSite_Entity_List is a list of WebSite_Entity.
type WebSite_Entity_List struct{ capnp.List }

// NewWebSite_Entity creates a new list of WebSite_Entity.
func NewWebSite_Entity_List(s *capnp.Segment, sz int32) (WebSite_Entity_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return WebSite_Entity_List{l}, err
}

func (s WebSite_Entity_List) At(i int) WebSite_Entity { return WebSite_Entity{s.List.Struct(i)} }

func (s WebSite_Entity_List) Set(i int, v WebSite_Entity) error { return s.List.SetStruct(i, v.Struct) }

// WebSite_Entity_Promise is a wrapper for a WebSite_Entity promised by a client call.
type WebSite_Entity_Promise struct{ *capnp.Pipeline }

func (p WebSite_Entity_Promise) Struct() (WebSite_Entity, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_Entity{s}, err
}

func (p WebSite_Entity_Promise) Body() WebSite_Entity_body_Promise {
	return WebSite_Entity_body_Promise{p.Pipeline}
}

// WebSite_Entity_body_Promise is a wrapper for a WebSite_Entity_body promised by a client call.
type WebSite_Entity_body_Promise struct{ *capnp.Pipeline }

func (p WebSite_Entity_body_Promise) Struct() (WebSite_Entity_body, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_Entity_body{s}, err
}

func (p WebSite_Entity_body_Promise) Blob() util.Blob {
	return util.Blob{Client: p.Pipeline.GetPipeline(3).Client()}
}

type WebSite_getUrl_Params struct{ capnp.Struct }

func NewWebSite_getUrl_Params(s *capnp.Segment) (WebSite_getUrl_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getUrl_Params{st}, err
}

func NewRootWebSite_getUrl_Params(s *capnp.Segment) (WebSite_getUrl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getUrl_Params{st}, err
}

func ReadRootWebSite_getUrl_Params(msg *capnp.Message) (WebSite_getUrl_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getUrl_Params{root.Struct()}, err
}

func (s WebSite_getUrl_Params) String() string {
	str, _ := text.Marshal(0xdc072ae47872d061, s.Struct)
	return str
}

// WebSite_getUrl_Params_List is a list of WebSite_getUrl_Params.
type WebSite_getUrl_Params_List struct{ capnp.List }

// NewWebSite_getUrl_Params creates a new list of WebSite_getUrl_Params.
func NewWebSite_getUrl_Params_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_getUrl_Params_List{l}, err
}

func (s WebSite_getUrl_Params_List) At(i int) WebSite_getUrl_Params {
	return WebSite_getUrl_Params{s.List.Struct(i)}
}

func (s WebSite_getUrl_Params_List) Set(i int, v WebSite_getUrl_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getUrl_Params_Promise is a wrapper for a WebSite_getUrl_Params promised by a client call.
type WebSite_getUrl_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getUrl_Params_Promise) Struct() (WebSite_getUrl_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getUrl_Params{s}, err
}

type WebSite_getUrl_Results struct{ capnp.Struct }

func NewWebSite_getUrl_Results(s *capnp.Segment) (WebSite_getUrl_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getUrl_Results{st}, err
}

func NewRootWebSite_getUrl_Results(s *capnp.Segment) (WebSite_getUrl_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getUrl_Results{st}, err
}

func ReadRootWebSite_getUrl_Results(msg *capnp.Message) (WebSite_getUrl_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getUrl_Results{root.Struct()}, err
}

func (s WebSite_getUrl_Results) String() string {
	str, _ := text.Marshal(0xb2855d483568639e, s.Struct)
	return str
}

func (s WebSite_getUrl_Results) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getUrl_Results) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getUrl_Results) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getUrl_Results) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getUrl_Results_List is a list of WebSite_getUrl_Results.
type WebSite_getUrl_Results_List struct{ capnp.List }

// NewWebSite_getUrl_Results creates a new list of WebSite_getUrl_Results.
func NewWebSite_getUrl_Results_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getUrl_Results_List{l}, err
}

func (s WebSite_getUrl_Results_List) At(i int) WebSite_getUrl_Results {
	return WebSite_getUrl_Results{s.List.Struct(i)}
}

func (s WebSite_getUrl_Results_List) Set(i int, v WebSite_getUrl_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getUrl_Results_Promise is a wrapper for a WebSite_getUrl_Results promised by a client call.
type WebSite_getUrl_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getUrl_Results_Promise) Struct() (WebSite_getUrl_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getUrl_Results{s}, err
}

type WebSite_getEntities_Params struct{ capnp.Struct }

func NewWebSite_getEntities_Params(s *capnp.Segment) (WebSite_getEntities_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Params{st}, err
}

func NewRootWebSite_getEntities_Params(s *capnp.Segment) (WebSite_getEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Params{st}, err
}

func ReadRootWebSite_getEntities_Params(msg *capnp.Message) (WebSite_getEntities_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getEntities_Params{root.Struct()}, err
}

func (s WebSite_getEntities_Params) String() string {
	str, _ := text.Marshal(0xfe1643d6b01e7cf4, s.Struct)
	return str
}

func (s WebSite_getEntities_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getEntities_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getEntities_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getEntities_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getEntities_Params_List is a list of WebSite_getEntities_Params.
type WebSite_getEntities_Params_List struct{ capnp.List }

// NewWebSite_getEntities_Params creates a new list of WebSite_getEntities_Params.
func NewWebSite_getEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getEntities_Params_List{l}, err
}

func (s WebSite_getEntities_Params_List) At(i int) WebSite_getEntities_Params {
	return WebSite_getEntities_Params{s.List.Struct(i)}
}

func (s WebSite_getEntities_Params_List) Set(i int, v WebSite_getEntities_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getEntities_Params_Promise is a wrapper for a WebSite_getEntities_Params promised by a client call.
type WebSite_getEntities_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getEntities_Params_Promise) Struct() (WebSite_getEntities_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getEntities_Params{s}, err
}

type WebSite_getEntities_Results struct{ capnp.Struct }

func NewWebSite_getEntities_Results(s *capnp.Segment) (WebSite_getEntities_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Results{st}, err
}

func NewRootWebSite_getEntities_Results(s *capnp.Segment) (WebSite_getEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Results{st}, err
}

func ReadRootWebSite_getEntities_Results(msg *capnp.Message) (WebSite_getEntities_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getEntities_Results{root.Struct()}, err
}

func (s WebSite_getEntities_Results) String() string {
	str, _ := text.Marshal(0xc3739b6f070fb5ea, s.Struct)
	return str
}

func (s WebSite_getEntities_Results) Entities() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s WebSite_getEntities_Results) HasEntities() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getEntities_Results) SetEntities(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getEntities_Results_List is a list of WebSite_getEntities_Results.
type WebSite_getEntities_Results_List struct{ capnp.List }

// NewWebSite_getEntities_Results creates a new list of WebSite_getEntities_Results.
func NewWebSite_getEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getEntities_Results_List{l}, err
}

func (s WebSite_getEntities_Results_List) At(i int) WebSite_getEntities_Results {
	return WebSite_getEntities_Results{s.List.Struct(i)}
}

func (s WebSite_getEntities_Results_List) Set(i int, v WebSite_getEntities_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getEntities_Results_Promise is a wrapper for a WebSite_getEntities_Results promised by a client call.
type WebSite_getEntities_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getEntities_Results_Promise) Struct() (WebSite_getEntities_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getEntities_Results{s}, err
}

func (p WebSite_getEntities_Results_Promise) Entities() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_getNotFoundEntities_Params struct{ capnp.Struct }

func NewWebSite_getNotFoundEntities_Params(s *capnp.Segment) (WebSite_getNotFoundEntities_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getNotFoundEntities_Params{st}, err
}

func NewRootWebSite_getNotFoundEntities_Params(s *capnp.Segment) (WebSite_getNotFoundEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getNotFoundEntities_Params{st}, err
}

func ReadRootWebSite_getNotFoundEntities_Params(msg *capnp.Message) (WebSite_getNotFoundEntities_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getNotFoundEntities_Params{root.Struct()}, err
}

func (s WebSite_getNotFoundEntities_Params) String() string {
	str, _ := text.Marshal(0xc3db68ea10a823b6, s.Struct)
	return str
}

// WebSite_getNotFoundEntities_Params_List is a list of WebSite_getNotFoundEntities_Params.
type WebSite_getNotFoundEntities_Params_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Params creates a new list of WebSite_getNotFoundEntities_Params.
func NewWebSite_getNotFoundEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_getNotFoundEntities_Params_List{l}, err
}

func (s WebSite_getNotFoundEntities_Params_List) At(i int) WebSite_getNotFoundEntities_Params {
	return WebSite_getNotFoundEntities_Params{s.List.Struct(i)}
}

func (s WebSite_getNotFoundEntities_Params_List) Set(i int, v WebSite_getNotFoundEntities_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getNotFoundEntities_Params_Promise is a wrapper for a WebSite_getNotFoundEntities_Params promised by a client call.
type WebSite_getNotFoundEntities_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getNotFoundEntities_Params_Promise) Struct() (WebSite_getNotFoundEntities_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getNotFoundEntities_Params{s}, err
}

type WebSite_getNotFoundEntities_Results struct{ capnp.Struct }

func NewWebSite_getNotFoundEntities_Results(s *capnp.Segment) (WebSite_getNotFoundEntities_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getNotFoundEntities_Results{st}, err
}

func NewRootWebSite_getNotFoundEntities_Results(s *capnp.Segment) (WebSite_getNotFoundEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getNotFoundEntities_Results{st}, err
}

func ReadRootWebSite_getNotFoundEntities_Results(msg *capnp.Message) (WebSite_getNotFoundEntities_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getNotFoundEntities_Results{root.Struct()}, err
}

func (s WebSite_getNotFoundEntities_Results) String() string {
	str, _ := text.Marshal(0xff383cbabc241462, s.Struct)
	return str
}

func (s WebSite_getNotFoundEntities_Results) Entities() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s WebSite_getNotFoundEntities_Results) HasEntities() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getNotFoundEntities_Results) SetEntities(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getNotFoundEntities_Results_List is a list of WebSite_getNotFoundEntities_Results.
type WebSite_getNotFoundEntities_Results_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Results creates a new list of WebSite_getNotFoundEntities_Results.
func NewWebSite_getNotFoundEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getNotFoundEntities_Results_List{l}, err
}

func (s WebSite_getNotFoundEntities_Results_List) At(i int) WebSite_getNotFoundEntities_Results {
	return WebSite_getNotFoundEntities_Results{s.List.Struct(i)}
}

func (s WebSite_getNotFoundEntities_Results_List) Set(i int, v WebSite_getNotFoundEntities_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getNotFoundEntities_Results_Promise is a wrapper for a WebSite_getNotFoundEntities_Results promised by a client call.
type WebSite_getNotFoundEntities_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getNotFoundEntities_Results_Promise) Struct() (WebSite_getNotFoundEntities_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getNotFoundEntities_Results{s}, err
}

func (p WebSite_getNotFoundEntities_Results_Promise) Entities() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_uploadBlob_Params struct{ capnp.Struct }

func NewWebSite_uploadBlob_Params(s *capnp.Segment) (WebSite_uploadBlob_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_uploadBlob_Params{st}, err
}

func NewRootWebSite_uploadBlob_Params(s *capnp.Segment) (WebSite_uploadBlob_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_uploadBlob_Params{st}, err
}

func ReadRootWebSite_uploadBlob_Params(msg *capnp.Message) (WebSite_uploadBlob_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_uploadBlob_Params{root.Struct()}, err
}

func (s WebSite_uploadBlob_Params) String() string {
	str, _ := text.Marshal(0xe2ae317a2a41a9f6, s.Struct)
	return str
}

// WebSite_uploadBlob_Params_List is a list of WebSite_uploadBlob_Params.
type WebSite_uploadBlob_Params_List struct{ capnp.List }

// NewWebSite_uploadBlob_Params creates a new list of WebSite_uploadBlob_Params.
func NewWebSite_uploadBlob_Params_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_uploadBlob_Params_List{l}, err
}

func (s WebSite_uploadBlob_Params_List) At(i int) WebSite_uploadBlob_Params {
	return WebSite_uploadBlob_Params{s.List.Struct(i)}
}

func (s WebSite_uploadBlob_Params_List) Set(i int, v WebSite_uploadBlob_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_uploadBlob_Params_Promise is a wrapper for a WebSite_uploadBlob_Params promised by a client call.
type WebSite_uploadBlob_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_uploadBlob_Params_Promise) Struct() (WebSite_uploadBlob_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_uploadBlob_Params{s}, err
}

type WebSite_uploadBlob_Results struct{ capnp.Struct }

func NewWebSite_uploadBlob_Results(s *capnp.Segment) (WebSite_uploadBlob_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSite_uploadBlob_Results{st}, err
}

func NewRootWebSite_uploadBlob_Results(s *capnp.Segment) (WebSite_uploadBlob_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSite_uploadBlob_Results{st}, err
}

func ReadRootWebSite_uploadBlob_Results(msg *capnp.Message) (WebSite_uploadBlob_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_uploadBlob_Results{root.Struct()}, err
}

func (s WebSite_uploadBlob_Results) String() string {
	str, _ := text.Marshal(0xc63c6c15dd189744, s.Struct)
	return str
}

func (s WebSite_uploadBlob_Results) Blob() util.Blob {
	p, _ := s.Struct.Ptr(0)
	return util.Blob{Client: p.Interface().Client()}
}

func (s WebSite_uploadBlob_Results) HasBlob() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_uploadBlob_Results) SetBlob(v util.Blob) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s WebSite_uploadBlob_Results) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(1)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s WebSite_uploadBlob_Results) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSite_uploadBlob_Results) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// WebSite_uploadBlob_Results_List is a list of WebSite_uploadBlob_Results.
type WebSite_uploadBlob_Results_List struct{ capnp.List }

// NewWebSite_uploadBlob_Results creates a new list of WebSite_uploadBlob_Results.
func NewWebSite_uploadBlob_Results_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSite_uploadBlob_Results_List{l}, err
}

func (s WebSite_uploadBlob_Results_List) At(i int) WebSite_uploadBlob_Results {
	return WebSite_uploadBlob_Results{s.List.Struct(i)}
}

func (s WebSite_uploadBlob_Results_List) Set(i int, v WebSite_uploadBlob_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_uploadBlob_Results_Promise is a wrapper for a WebSite_uploadBlob_Results promised by a client call.
type WebSite_uploadBlob_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_uploadBlob_Results_Promise) Struct() (WebSite_uploadBlob_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_uploadBlob_Results{s}, err
}

func (p WebSite_uploadBlob_Results_Promise) Blob() util.Blob {
	return util.Blob{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p WebSite_uploadBlob_Results_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

type WebSite_getSubsite_Params struct{ capnp.Struct }

func NewWebSite_getSubsite_Params(s *capnp.Segment) (WebSite_getSubsite_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Params{st}, err
}

func NewRootWebSite_getSubsite_Params(s *capnp.Segment) (WebSite_getSubsite_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Params{st}, err
}

func ReadRootWebSite_getSubsite_Params(msg *capnp.Message) (WebSite_getSubsite_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getSubsite_Params{root.Struct()}, err
}

func (s WebSite_getSubsite_Params) String() string {
	str, _ := text.Marshal(0xdcbe913db7496644, s.Struct)
	return str
}

func (s WebSite_getSubsite_Params) Prefix() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getSubsite_Params) HasPrefix() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getSubsite_Params) PrefixBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getSubsite_Params) SetPrefix(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getSubsite_Params_List is a list of WebSite_getSubsite_Params.
type WebSite_getSubsite_Params_List struct{ capnp.List }

// NewWebSite_getSubsite_Params creates a new list of WebSite_getSubsite_Params.
func NewWebSite_getSubsite_Params_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getSubsite_Params_List{l}, err
}

func (s WebSite_getSubsite_Params_List) At(i int) WebSite_getSubsite_Params {
	return WebSite_getSubsite_Params{s.List.Struct(i)}
}

func (s WebSite_getSubsite_Params_List) Set(i int, v WebSite_getSubsite_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getSubsite_Params_Promise is a wrapper for a WebSite_getSubsite_Params promised by a client call.
type WebSite_getSubsite_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getSubsite_Params_Promise) Struct() (WebSite_getSubsite_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getSubsite_Params{s}, err
}

type WebSite_getSubsite_Results struct{ capnp.Struct }

func NewWebSite_getSubsite_Results(s *capnp.Segment) (WebSite_getSubsite_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Results{st}, err
}

func NewRootWebSite_getSubsite_Results(s *capnp.Segment) (WebSite_getSubsite_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Results{st}, err
}

func ReadRootWebSite_getSubsite_Results(msg *capnp.Message) (WebSite_getSubsite_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getSubsite_Results{root.Struct()}, err
}

func (s WebSite_getSubsite_Results) String() string {
	str, _ := text.Marshal(0x90dec3f1a5d9b591, s.Struct)
	return str
}

func (s WebSite_getSubsite_Results) Site() WebSite {
	p, _ := s.Struct.Ptr(0)
	return WebSite{Client: p.Interface().Client()}
}

func (s WebSite_getSubsite_Results) HasSite() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getSubsite_Results) SetSite(v WebSite) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getSubsite_Results_List is a list of WebSite_getSubsite_Results.
type WebSite_getSubsite_Results_List struct{ capnp.List }

// NewWebSite_getSubsite_Results creates a new list of WebSite_getSubsite_Results.
func NewWebSite_getSubsite_Results_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getSubsite_Results_List{l}, err
}

func (s WebSite_getSubsite_Results_List) At(i int) WebSite_getSubsite_Results {
	return WebSite_getSubsite_Results{s.List.Struct(i)}
}

func (s WebSite_getSubsite_Results_List) Set(i int, v WebSite_getSubsite_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getSubsite_Results_Promise is a wrapper for a WebSite_getSubsite_Results promised by a client call.
type WebSite_getSubsite_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getSubsite_Results_Promise) Struct() (WebSite_getSubsite_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getSubsite_Results{s}, err
}

func (p WebSite_getSubsite_Results_Promise) Site() WebSite {
	return WebSite{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_listResources_Params struct{ capnp.Struct }

func NewWebSite_listResources_Params(s *capnp.Segment) (WebSite_listResources_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSite_listResources_Params{st}, err
}

func NewRootWebSite_listResources_Params(s *capnp.Segment) (WebSite_listResources_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSite_listResources_Params{st}, err
}

func ReadRootWebSite_listResources_Params(msg *capnp.Message) (WebSite_listResources_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_listResources_Params{root.Struct()}, err
}

func (s WebSite_listResources_Params) String() string {
	str, _ := text.Marshal(0x87a72be8e04db694, s.Struct)
	return str
}

func (s WebSite_listResources_Params) Shallow() bool {
	return s.Struct.Bit(0)
}

func (s WebSite_listResources_Params) SetShallow(v bool) {
	s.Struct.SetBit(0, v)
}

// WebSite_listResources_Params_List is a list of WebSite_listResources_Params.
type WebSite_listResources_Params_List struct{ capnp.List }

// NewWebSite_listResources_Params creates a new list of WebSite_listResources_Params.
func NewWebSite_listResources_Params_List(s *capnp.Segment, sz int32) (WebSite_listResources_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return WebSite_listResources_Params_List{l}, err
}

func (s WebSite_listResources_Params_List) At(i int) WebSite_listResources_Params {
	return WebSite_listResources_Params{s.List.Struct(i)}
}

func (s WebSite_listResources_Params_List) Set(i int, v WebSite_listResources_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_listResources_Params_Promise is a wrapper for a WebSite_listResources_Params promised by a client call.
type WebSite_listResources_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_listResources_Params_Promise) Struct() (WebSite_listResources_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_listResources_Params{s}, err
}

type WebSite_listResources_Results struct{ capnp.Struct }

func NewWebSite_listResources_Results(s *capnp.Segment) (WebSite_listResources_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_listResources_Results{st}, err
}

func NewRootWebSite_listResources_Results(s *capnp.Segment) (WebSite_listResources_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_listResources_Results{st}, err
}

func ReadRootWebSite_listResources_Results(msg *capnp.Message) (WebSite_listResources_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_listResources_Results{root.Struct()}, err
}

func (s WebSite_listResources_Results) String() string {
	str, _ := text.Marshal(0xf5e0920223f0b273, s.Struct)
	return str
}

func (s WebSite_listResources_Results) Names() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSite_listResources_Results) HasNames() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_listResources_Results) SetNames(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNames sets the names field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSite_listResources_Results) NewNames(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// WebSite_listResources_Results_List is a list of WebSite_listResources_Results.
type WebSite_listResources_Results_List struct{ capnp.List }

// NewWebSite_listResources_Results creates a new list of WebSite_listResources_Results.
func NewWebSite_listResources_Results_List(s *capnp.Segment, sz int32) (WebSite_listResources_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_listResources_Results_List{l}, err
}

func (s WebSite_listResources_Results_List) At(i int) WebSite_listResources_Results {
	return WebSite_listResources_Results{s.List.Struct(i)}
}

func (s WebSite_listResources_Results_List) Set(i int, v WebSite_listResources_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_listResources_Results_Promise is a wrapper for a WebSite_listResources_Results promised by a client call.
type WebSite_listResources_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_listResources_Results_Promise) Struct() (WebSite_listResources_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_listResources_Results{s}, err
}

const schema_d5d3e63bd0a552b6 = "x\xda\xc4Umh\x1cU\x14}wf\xb6\x93\xecn" +
	"\x92>\xb2\xb5\x0d\x18\x03\xcd\x82f%\xa1\xad\x16jL" +
	"\xbbik\xfd(T2mC5Pd&\x99f\x17" +
	"fw\xd6\x9dY\xb6\x1bZ?\x8a\xa1\"\x8a\xd9b1" +
	"U*(\xac\xb5?\xd4Zl\x04\xb1B-\x08-\x88" +
	"\x11Q\x08&R\xd4\xaa\x04\x04\x15\x15\xfc\xd3\xe7}\xb3" +
	";\x1fI\xbb\xa6\xf9\xd5B\xc9\xce}\xe7\xdd{\xde\xbd" +
	"\xe7\x9d\xb7\xaeU\xea\x17\xd6\x87~\xa1\x84(\xd3\xa1\x15" +
	"\xecL\xe9\xc2\xdb\x89\xed\xef\x1f!J\x1c\x80\xa5\x9ej" +
	"\xee\xdb2\xd76M\x06C2\x08\x84\xdc3\xd3\xb0\x13" +
	"\x08\xb4^m(\x12`\xafL\xed\xba\xf2\xeb\xdd\xa7\x8e" +
	"\x12\xa5\x1d\x80\x10IF\xc0\xd6\xc6c\x1c\xa04&\x11" +
	"P>7S\xf9\xe3\xe2\xf7\x13\x84\xb6\xe3z\x088\xe0" +
	"\xc9\xc6#\x1cp\xd8\x01\xbc1\x9c\xda\xf8\xf0\xfe\xf1\xb3" +
	"A\xc0\xc9F\x8d\x03N;\x80\xf9s-\xb2\xf9\xbau" +
	"1\x08\xb8\xd4\xf8\x02\x07\xcc8\x80\xa9\xcewV\xce\xa7" +
	"\xbe\xab\x01\x1c\x0a\x9b\xc3\x9f\x01\x91\xd8\x03\xaf\xae\x99[" +
	"e\xf4}^\xdb*\xf0\xa5\xee\xb0S\xfc\xbe0\xa7\xef" +
	"\x1d\x8e\xd3g\xc5\xce-\xcf<tyv\x8e\x84B\x1c" +
	"y<\xbc\x16Z+a\xfe\xf3\xcdp\x07nb\xeat" +
	"\xfe\xe0O\x09y6P\xa9\x12\x19r*\x1dx\xe4\xa3" +
	"\xcd\xe5Og\x83$\xcb\x911^\xe9d\x84\x93\xf4r" +
	"\xd3U\"\x9b\xda]\x99\xbe\xff\xe7\xaf\xbf!\xb8|>" +
	"2\xdfz)\xb2\x1a\xf1_E\x8e\xe2\x91\xa2\xb8\x95\xfd" +
	"szkbl\xfd{?\x04\x0a]\x88\x8e\xf1B\xd6" +
	"\xd9\xdf;\x85cW\xfe\x0e\x16z7\xfa\x1a/t>" +
	"\xca\x0b\xfdu\xe8\x8e3\xdfn\xbf\xedZ\x10\xf0[\xd4" +
	"9\xf3\xbf\x0e@\x8b\xc5?\xf9\xb8o\x13\x0b\x02\xda\x9a" +
	".s@wS\x9202\xc9\x8a\xba\xd6\x9d+h\x86" +
	"\x94\xb6R\xe9\xech\xcf\xb0\x9a\xcb\xe6z\xf7\xe9\xda\x9e" +
	"\xb4\xad\xf7\xec\xc8\xdai\xbb\xd4\xa3\x99#%TL\x83" +
	"(\xdd\xce\x98\x18\x03\x91\x10\xda\xb5\x01#q\x11\x94u" +
	"\x02\xb4\xc35\x1e\x960\xdc\x9d\xc0\xf0]\x18\xbeW\x80" +
	"\x0e\xadd\xeb\x164\x11\x01\xffC\x8bf\x98\x1aP\xf6" +
	"r\xdb\xc4\xfeSwn\xbc\x8a\x1d\x01\xca\xdbUc\x10" +
	"\xba!\x03#m\xd9\xbbu\xcb,\xe4\x87u+>\xa0" +
	"\xe6\xd5\x0cX\x8a$b-\x09\xcfD\x9b\xb69\xc4@" +
	"\x89\x09\xf0\xb4\x95R\x0d\xc3,\xa2<\x05\x80%S\x8f" +
	"\xea\xf6\x9e\x82f\xe1\xcf8V(\x18\xb6E\x82\x89\x13" +
	"~\xe2\x16\x0eB\xea\xded\x17Q\x97\xea\xe5\x1f\xcc\x1b" +
	"Nn\x19\x93\xd7\xcb\x9dS\xed\x14D\x91q\xf4f\x18" +
	";\x13Ic+\\\xca\x0b8\xef\xc4\xcf(\xe6}L" +
	"\x00\xa6\xd7\x90\x04\xd9Rv\xa2\xbcv\xf2\xc3\xc1?\xe7" +
	"\xf1\xab\x1f(t(\x12B\x02A\x0a\xab\x15\x09\xf0\xdf" +
	"\x80\x08\xd0L\x9c?+\xfdK\xb3\xe8\xc0\xd1z\xf4\x1e" +
	"5\xed\x07\xcdBv\xc4\xa3\xe9L\xcc\"\xee\xc6\x1b\xee" +
	"+\xe4\x0cS\x1d\xd9\x86\xf2\xf0\x07\xd1\xe0\x1d\xaa+\xe1" +
	"\x0b\x8d\x02\xc4\xb8\xf9\xd0\xee^_fu\x84\x95\xb4\xec" +
	"\xbc\x8ej\xa1\xec\x89\x1f\xbf\xec*n\xda\xf7\xc5\xe2S" +
	"\x88\xf54/\xa3\xe8\x07\x00\xc7\xe3\x918\xcc;{\x08" +
	"\xeb=\x1f 1\xce\x83\xcfap\x02\x83\x82\x10\xe3\x9e" +
	"I_\xe2\xc1\x1718\x89\x1a\xc4&\xfa\x1eK\x8f'" +
	"\x88@%)\x06!\x9eq\xc8\xcf\xc82\xe9\x8c\xbe\xb7" +
	"\x94\xd3\xf9\xac\\)\x18jv\xb4\xa0\x8e.\x88\xe9\xd9" +
	"as\x04)\x07b-\xfcr\xb2\xbc>\x92\xce\xeb\xc3" +
	"6\x11\xf7\x9a\xd7\x89\xe9\x7f\xe5\xc9\x07$f\xac\xa5\xc1" +
	"\xee]IVG\x1a\x94]\xaf/\xe7d.\xaf\x1fH" +
	"\x1f\xbc\x8e\x83\xb08m\x0b\xcf\xcb%\xe7\x8b\x8cBo" +
	"\xb2\xea8\xca\x1a\x11{\xe4\x990\xb8/\x07\xfd\xa0\x17" +
	"[X\x91\xc17?p\x1f\x0dzB\xc3\xb5\xb2\x0c\xa2" +
	"\xf7L\x80k\x80t\xfc-\\{V\x06\xc9\xf3[p" +
	"_\x0cZ\x18\xc2\xb5\x8c\x0c!\xcf\xda\xc1}\xca\xa8\xca" +
	"\xd7\x1e\xc7\xd7\xd0{\xfe\xc0\xb5e\xba+\x8fk;\xe4" +
	"d\xb5\x91\xfd\xc0\xdc\xebIdT~\xf5\xdb\xb9\x0f\xe0" +
	"^\x88Z\xdc\xd5;\x11M\xad\x0aszKD[\xc7" +
	"O\xd7\xf1H\x87\xe3y\xfd0\x00K\x8d2p\x81j" +
	"\xd3Y\x96\xab:\x97N\\\xe8P\x1bj#\x8d\xa3\x8d" +
	"g\xd5\x0c\xdax\xcd\x14\xf8`\x9b\x97\xe5T5\x0bX" +
	"\x96\xff\xad\xb8i\x83\xb9\x15>\xf8_\x00\x00\x00\xff\xff" +
	"#\xdb\xa0/"

func init() {
	schemas.Register(schema_d5d3e63bd0a552b6,
		0x82af432aa6c179b0,
		0x87a72be8e04db694,
		0x90dec3f1a5d9b591,
		0xb2855d483568639e,
		0xc3739b6f070fb5ea,
		0xc3db68ea10a823b6,
		0xc63c6c15dd189744,
		0xd019dd3e3c0e7e68,
		0xdc072ae47872d061,
		0xdcbe913db7496644,
		0xdddcca47803e2377,
		0xe2ae317a2a41a9f6,
		0xf5e0920223f0b273,
		0xfe1643d6b01e7cf4,
		0xff383cbabc241462)
}
