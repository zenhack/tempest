package webpublishing

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
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
	if err != nil {
		return WebSite_Entity{}, err
	}
	return WebSite_Entity{st}, nil
}

func NewRootWebSite_Entity(s *capnp.Segment) (WebSite_Entity, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	if err != nil {
		return WebSite_Entity{}, err
	}
	return WebSite_Entity{st}, nil
}

func ReadRootWebSite_Entity(msg *capnp.Message) (WebSite_Entity, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_Entity{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_Entity{st}, nil
}

func (s WebSite_Entity) MimeType() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_Entity) SetMimeType(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s WebSite_Entity) Language() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_Entity) SetLanguage(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s WebSite_Entity) Encoding() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_Entity) SetEncoding(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}
func (s WebSite_Entity) Body() WebSite_Entity_body { return WebSite_Entity_body(s) }

func (s WebSite_Entity_body) Which() WebSite_Entity_body_Which {
	return WebSite_Entity_body_Which(s.Struct.Uint16(0))
}

func (s WebSite_Entity_body) Bytes() ([]byte, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s WebSite_Entity_body) SetBytes(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(3, d)
}

func (s WebSite_Entity_body) Blob() util.Blob {
	p, err := s.Struct.Pointer(3)
	if err != nil {

		return util.Blob{}
	}
	c := capnp.ToInterface(p).Client()
	return util.Blob{Client: c}
}

func (s WebSite_Entity_body) SetBlob(v util.Blob) error {
	s.Struct.SetUint16(0, 1)
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(3, in)
}

func (s WebSite_Entity) RedirectTo() (string, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_Entity) SetRedirectTo(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(4, t)
}

// WebSite_Entity_List is a list of WebSite_Entity.
type WebSite_Entity_List struct{ capnp.List }

// NewWebSite_Entity creates a new list of WebSite_Entity.
func NewWebSite_Entity_List(s *capnp.Segment, sz int32) (WebSite_Entity_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	if err != nil {
		return WebSite_Entity_List{}, err
	}
	return WebSite_Entity_List{l}, nil
}

func (s WebSite_Entity_List) At(i int) WebSite_Entity           { return WebSite_Entity{s.List.Struct(i)} }
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
	if err != nil {
		return WebSite_getUrl_Params{}, err
	}
	return WebSite_getUrl_Params{st}, nil
}

func NewRootWebSite_getUrl_Params(s *capnp.Segment) (WebSite_getUrl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return WebSite_getUrl_Params{}, err
	}
	return WebSite_getUrl_Params{st}, nil
}

func ReadRootWebSite_getUrl_Params(msg *capnp.Message) (WebSite_getUrl_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getUrl_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getUrl_Params{st}, nil
}

// WebSite_getUrl_Params_List is a list of WebSite_getUrl_Params.
type WebSite_getUrl_Params_List struct{ capnp.List }

// NewWebSite_getUrl_Params creates a new list of WebSite_getUrl_Params.
func NewWebSite_getUrl_Params_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return WebSite_getUrl_Params_List{}, err
	}
	return WebSite_getUrl_Params_List{l}, nil
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
	if err != nil {
		return WebSite_getUrl_Results{}, err
	}
	return WebSite_getUrl_Results{st}, nil
}

func NewRootWebSite_getUrl_Results(s *capnp.Segment) (WebSite_getUrl_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getUrl_Results{}, err
	}
	return WebSite_getUrl_Results{st}, nil
}

func ReadRootWebSite_getUrl_Results(msg *capnp.Message) (WebSite_getUrl_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getUrl_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getUrl_Results{st}, nil
}

func (s WebSite_getUrl_Results) Path() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_getUrl_Results) SetPath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// WebSite_getUrl_Results_List is a list of WebSite_getUrl_Results.
type WebSite_getUrl_Results_List struct{ capnp.List }

// NewWebSite_getUrl_Results creates a new list of WebSite_getUrl_Results.
func NewWebSite_getUrl_Results_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getUrl_Results_List{}, err
	}
	return WebSite_getUrl_Results_List{l}, nil
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
	if err != nil {
		return WebSite_getEntities_Params{}, err
	}
	return WebSite_getEntities_Params{st}, nil
}

func NewRootWebSite_getEntities_Params(s *capnp.Segment) (WebSite_getEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getEntities_Params{}, err
	}
	return WebSite_getEntities_Params{st}, nil
}

func ReadRootWebSite_getEntities_Params(msg *capnp.Message) (WebSite_getEntities_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getEntities_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getEntities_Params{st}, nil
}

func (s WebSite_getEntities_Params) Path() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_getEntities_Params) SetPath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// WebSite_getEntities_Params_List is a list of WebSite_getEntities_Params.
type WebSite_getEntities_Params_List struct{ capnp.List }

// NewWebSite_getEntities_Params creates a new list of WebSite_getEntities_Params.
func NewWebSite_getEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getEntities_Params_List{}, err
	}
	return WebSite_getEntities_Params_List{l}, nil
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
	if err != nil {
		return WebSite_getEntities_Results{}, err
	}
	return WebSite_getEntities_Results{st}, nil
}

func NewRootWebSite_getEntities_Results(s *capnp.Segment) (WebSite_getEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getEntities_Results{}, err
	}
	return WebSite_getEntities_Results{st}, nil
}

func ReadRootWebSite_getEntities_Results(msg *capnp.Message) (WebSite_getEntities_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getEntities_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getEntities_Results{st}, nil
}

func (s WebSite_getEntities_Results) Entities() util.Assignable {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return util.Assignable{}
	}
	c := capnp.ToInterface(p).Client()
	return util.Assignable{Client: c}
}

func (s WebSite_getEntities_Results) SetEntities(v util.Assignable) error {

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

// WebSite_getEntities_Results_List is a list of WebSite_getEntities_Results.
type WebSite_getEntities_Results_List struct{ capnp.List }

// NewWebSite_getEntities_Results creates a new list of WebSite_getEntities_Results.
func NewWebSite_getEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getEntities_Results_List{}, err
	}
	return WebSite_getEntities_Results_List{l}, nil
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
	if err != nil {
		return WebSite_getNotFoundEntities_Params{}, err
	}
	return WebSite_getNotFoundEntities_Params{st}, nil
}

func NewRootWebSite_getNotFoundEntities_Params(s *capnp.Segment) (WebSite_getNotFoundEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return WebSite_getNotFoundEntities_Params{}, err
	}
	return WebSite_getNotFoundEntities_Params{st}, nil
}

func ReadRootWebSite_getNotFoundEntities_Params(msg *capnp.Message) (WebSite_getNotFoundEntities_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getNotFoundEntities_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getNotFoundEntities_Params{st}, nil
}

// WebSite_getNotFoundEntities_Params_List is a list of WebSite_getNotFoundEntities_Params.
type WebSite_getNotFoundEntities_Params_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Params creates a new list of WebSite_getNotFoundEntities_Params.
func NewWebSite_getNotFoundEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return WebSite_getNotFoundEntities_Params_List{}, err
	}
	return WebSite_getNotFoundEntities_Params_List{l}, nil
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
	if err != nil {
		return WebSite_getNotFoundEntities_Results{}, err
	}
	return WebSite_getNotFoundEntities_Results{st}, nil
}

func NewRootWebSite_getNotFoundEntities_Results(s *capnp.Segment) (WebSite_getNotFoundEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getNotFoundEntities_Results{}, err
	}
	return WebSite_getNotFoundEntities_Results{st}, nil
}

func ReadRootWebSite_getNotFoundEntities_Results(msg *capnp.Message) (WebSite_getNotFoundEntities_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getNotFoundEntities_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getNotFoundEntities_Results{st}, nil
}

func (s WebSite_getNotFoundEntities_Results) Entities() util.Assignable {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return util.Assignable{}
	}
	c := capnp.ToInterface(p).Client()
	return util.Assignable{Client: c}
}

func (s WebSite_getNotFoundEntities_Results) SetEntities(v util.Assignable) error {

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

// WebSite_getNotFoundEntities_Results_List is a list of WebSite_getNotFoundEntities_Results.
type WebSite_getNotFoundEntities_Results_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Results creates a new list of WebSite_getNotFoundEntities_Results.
func NewWebSite_getNotFoundEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getNotFoundEntities_Results_List{}, err
	}
	return WebSite_getNotFoundEntities_Results_List{l}, nil
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
	if err != nil {
		return WebSite_uploadBlob_Params{}, err
	}
	return WebSite_uploadBlob_Params{st}, nil
}

func NewRootWebSite_uploadBlob_Params(s *capnp.Segment) (WebSite_uploadBlob_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return WebSite_uploadBlob_Params{}, err
	}
	return WebSite_uploadBlob_Params{st}, nil
}

func ReadRootWebSite_uploadBlob_Params(msg *capnp.Message) (WebSite_uploadBlob_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_uploadBlob_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_uploadBlob_Params{st}, nil
}

// WebSite_uploadBlob_Params_List is a list of WebSite_uploadBlob_Params.
type WebSite_uploadBlob_Params_List struct{ capnp.List }

// NewWebSite_uploadBlob_Params creates a new list of WebSite_uploadBlob_Params.
func NewWebSite_uploadBlob_Params_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return WebSite_uploadBlob_Params_List{}, err
	}
	return WebSite_uploadBlob_Params_List{l}, nil
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
	if err != nil {
		return WebSite_uploadBlob_Results{}, err
	}
	return WebSite_uploadBlob_Results{st}, nil
}

func NewRootWebSite_uploadBlob_Results(s *capnp.Segment) (WebSite_uploadBlob_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return WebSite_uploadBlob_Results{}, err
	}
	return WebSite_uploadBlob_Results{st}, nil
}

func ReadRootWebSite_uploadBlob_Results(msg *capnp.Message) (WebSite_uploadBlob_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_uploadBlob_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_uploadBlob_Results{st}, nil
}

func (s WebSite_uploadBlob_Results) Blob() util.Blob {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return util.Blob{}
	}
	c := capnp.ToInterface(p).Client()
	return util.Blob{Client: c}
}

func (s WebSite_uploadBlob_Results) SetBlob(v util.Blob) error {

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

func (s WebSite_uploadBlob_Results) Stream() util.ByteStream {
	p, err := s.Struct.Pointer(1)
	if err != nil {

		return util.ByteStream{}
	}
	c := capnp.ToInterface(p).Client()
	return util.ByteStream{Client: c}
}

func (s WebSite_uploadBlob_Results) SetStream(v util.ByteStream) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(1, in)
}

// WebSite_uploadBlob_Results_List is a list of WebSite_uploadBlob_Results.
type WebSite_uploadBlob_Results_List struct{ capnp.List }

// NewWebSite_uploadBlob_Results creates a new list of WebSite_uploadBlob_Results.
func NewWebSite_uploadBlob_Results_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return WebSite_uploadBlob_Results_List{}, err
	}
	return WebSite_uploadBlob_Results_List{l}, nil
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
	if err != nil {
		return WebSite_getSubsite_Params{}, err
	}
	return WebSite_getSubsite_Params{st}, nil
}

func NewRootWebSite_getSubsite_Params(s *capnp.Segment) (WebSite_getSubsite_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getSubsite_Params{}, err
	}
	return WebSite_getSubsite_Params{st}, nil
}

func ReadRootWebSite_getSubsite_Params(msg *capnp.Message) (WebSite_getSubsite_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getSubsite_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getSubsite_Params{st}, nil
}

func (s WebSite_getSubsite_Params) Prefix() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s WebSite_getSubsite_Params) SetPrefix(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// WebSite_getSubsite_Params_List is a list of WebSite_getSubsite_Params.
type WebSite_getSubsite_Params_List struct{ capnp.List }

// NewWebSite_getSubsite_Params creates a new list of WebSite_getSubsite_Params.
func NewWebSite_getSubsite_Params_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getSubsite_Params_List{}, err
	}
	return WebSite_getSubsite_Params_List{l}, nil
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
	if err != nil {
		return WebSite_getSubsite_Results{}, err
	}
	return WebSite_getSubsite_Results{st}, nil
}

func NewRootWebSite_getSubsite_Results(s *capnp.Segment) (WebSite_getSubsite_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_getSubsite_Results{}, err
	}
	return WebSite_getSubsite_Results{st}, nil
}

func ReadRootWebSite_getSubsite_Results(msg *capnp.Message) (WebSite_getSubsite_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_getSubsite_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_getSubsite_Results{st}, nil
}

func (s WebSite_getSubsite_Results) Site() WebSite {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return WebSite{}
	}
	c := capnp.ToInterface(p).Client()
	return WebSite{Client: c}
}

func (s WebSite_getSubsite_Results) SetSite(v WebSite) error {

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

// WebSite_getSubsite_Results_List is a list of WebSite_getSubsite_Results.
type WebSite_getSubsite_Results_List struct{ capnp.List }

// NewWebSite_getSubsite_Results creates a new list of WebSite_getSubsite_Results.
func NewWebSite_getSubsite_Results_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_getSubsite_Results_List{}, err
	}
	return WebSite_getSubsite_Results_List{l}, nil
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
	if err != nil {
		return WebSite_listResources_Params{}, err
	}
	return WebSite_listResources_Params{st}, nil
}

func NewRootWebSite_listResources_Params(s *capnp.Segment) (WebSite_listResources_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return WebSite_listResources_Params{}, err
	}
	return WebSite_listResources_Params{st}, nil
}

func ReadRootWebSite_listResources_Params(msg *capnp.Message) (WebSite_listResources_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_listResources_Params{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_listResources_Params{st}, nil
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
	if err != nil {
		return WebSite_listResources_Params_List{}, err
	}
	return WebSite_listResources_Params_List{l}, nil
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
	if err != nil {
		return WebSite_listResources_Results{}, err
	}
	return WebSite_listResources_Results{st}, nil
}

func NewRootWebSite_listResources_Results(s *capnp.Segment) (WebSite_listResources_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return WebSite_listResources_Results{}, err
	}
	return WebSite_listResources_Results{st}, nil
}

func ReadRootWebSite_listResources_Results(msg *capnp.Message) (WebSite_listResources_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return WebSite_listResources_Results{}, err
	}
	st := capnp.ToStruct(root)
	return WebSite_listResources_Results{st}, nil
}

func (s WebSite_listResources_Results) Names() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s WebSite_listResources_Results) SetNames(v capnp.TextList) error {

	return s.Struct.SetPointer(0, v.List)
}

// WebSite_listResources_Results_List is a list of WebSite_listResources_Results.
type WebSite_listResources_Results_List struct{ capnp.List }

// NewWebSite_listResources_Results creates a new list of WebSite_listResources_Results.
func NewWebSite_listResources_Results_List(s *capnp.Segment, sz int32) (WebSite_listResources_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return WebSite_listResources_Results_List{}, err
	}
	return WebSite_listResources_Results_List{l}, nil
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
