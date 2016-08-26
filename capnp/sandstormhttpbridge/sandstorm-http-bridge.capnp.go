package sandstormhttpbridge

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type SandstormHttpBridge struct{ Client capnp.Client }

func (c SandstormHttpBridge) GetSandstormApi(ctx context.Context, params func(SandstormHttpBridge_getSandstormApi_Params) error, opts ...capnp.CallOption) SandstormHttpBridge_getSandstormApi_Results_Promise {
	if c.Client == nil {
		return SandstormHttpBridge_getSandstormApi_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      0,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSandstormApi",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormHttpBridge_getSandstormApi_Params{Struct: s}) }
	}
	return SandstormHttpBridge_getSandstormApi_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormHttpBridge) GetSessionContext(ctx context.Context, params func(SandstormHttpBridge_getSessionContext_Params) error, opts ...capnp.CallOption) SandstormHttpBridge_getSessionContext_Results_Promise {
	if c.Client == nil {
		return SandstormHttpBridge_getSessionContext_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      1,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSessionContext",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormHttpBridge_getSessionContext_Params{Struct: s}) }
	}
	return SandstormHttpBridge_getSessionContext_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormHttpBridge) GetSavedIdentity(ctx context.Context, params func(SandstormHttpBridge_getSavedIdentity_Params) error, opts ...capnp.CallOption) SandstormHttpBridge_getSavedIdentity_Results_Promise {
	if c.Client == nil {
		return SandstormHttpBridge_getSavedIdentity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      2,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSavedIdentity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormHttpBridge_getSavedIdentity_Params{Struct: s}) }
	}
	return SandstormHttpBridge_getSavedIdentity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormHttpBridge_Server interface {
	GetSandstormApi(SandstormHttpBridge_getSandstormApi) error

	GetSessionContext(SandstormHttpBridge_getSessionContext) error

	GetSavedIdentity(SandstormHttpBridge_getSavedIdentity) error
}

func SandstormHttpBridge_ServerToClient(s SandstormHttpBridge_Server) SandstormHttpBridge {
	c, _ := s.(server.Closer)
	return SandstormHttpBridge{Client: server.New(SandstormHttpBridge_Methods(nil, s), c)}
}

func SandstormHttpBridge_Methods(methods []server.Method, s SandstormHttpBridge_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      0,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSandstormApi",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormHttpBridge_getSandstormApi{c, opts, SandstormHttpBridge_getSandstormApi_Params{Struct: p}, SandstormHttpBridge_getSandstormApi_Results{Struct: r}}
			return s.GetSandstormApi(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      1,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSessionContext",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormHttpBridge_getSessionContext{c, opts, SandstormHttpBridge_getSessionContext_Params{Struct: p}, SandstormHttpBridge_getSessionContext_Results{Struct: r}}
			return s.GetSessionContext(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      2,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "getSavedIdentity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormHttpBridge_getSavedIdentity{c, opts, SandstormHttpBridge_getSavedIdentity_Params{Struct: p}, SandstormHttpBridge_getSavedIdentity_Results{Struct: r}}
			return s.GetSavedIdentity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SandstormHttpBridge_getSandstormApi holds the arguments for a server call to SandstormHttpBridge.getSandstormApi.
type SandstormHttpBridge_getSandstormApi struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormHttpBridge_getSandstormApi_Params
	Results SandstormHttpBridge_getSandstormApi_Results
}

// SandstormHttpBridge_getSessionContext holds the arguments for a server call to SandstormHttpBridge.getSessionContext.
type SandstormHttpBridge_getSessionContext struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormHttpBridge_getSessionContext_Params
	Results SandstormHttpBridge_getSessionContext_Results
}

// SandstormHttpBridge_getSavedIdentity holds the arguments for a server call to SandstormHttpBridge.getSavedIdentity.
type SandstormHttpBridge_getSavedIdentity struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormHttpBridge_getSavedIdentity_Params
	Results SandstormHttpBridge_getSavedIdentity_Results
}

type SandstormHttpBridge_getSandstormApi_Params struct{ capnp.Struct }

func NewSandstormHttpBridge_getSandstormApi_Params(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormHttpBridge_getSandstormApi_Params{st}, err
}

func NewRootSandstormHttpBridge_getSandstormApi_Params(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormHttpBridge_getSandstormApi_Params{st}, err
}

func ReadRootSandstormHttpBridge_getSandstormApi_Params(msg *capnp.Message) (SandstormHttpBridge_getSandstormApi_Params, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSandstormApi_Params{root.Struct()}, err
}

func (s SandstormHttpBridge_getSandstormApi_Params) String() string {
	str, _ := text.Marshal(0x86d490c59f64c564, s.Struct)
	return str
}

// SandstormHttpBridge_getSandstormApi_Params_List is a list of SandstormHttpBridge_getSandstormApi_Params.
type SandstormHttpBridge_getSandstormApi_Params_List struct{ capnp.List }

// NewSandstormHttpBridge_getSandstormApi_Params creates a new list of SandstormHttpBridge_getSandstormApi_Params.
func NewSandstormHttpBridge_getSandstormApi_Params_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSandstormApi_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormHttpBridge_getSandstormApi_Params_List{l}, err
}

func (s SandstormHttpBridge_getSandstormApi_Params_List) At(i int) SandstormHttpBridge_getSandstormApi_Params {
	return SandstormHttpBridge_getSandstormApi_Params{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSandstormApi_Params_List) Set(i int, v SandstormHttpBridge_getSandstormApi_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSandstormApi_Params_Promise is a wrapper for a SandstormHttpBridge_getSandstormApi_Params promised by a client call.
type SandstormHttpBridge_getSandstormApi_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSandstormApi_Params_Promise) Struct() (SandstormHttpBridge_getSandstormApi_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSandstormApi_Params{s}, err
}

type SandstormHttpBridge_getSandstormApi_Results struct{ capnp.Struct }

func NewSandstormHttpBridge_getSandstormApi_Results(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSandstormApi_Results{st}, err
}

func NewRootSandstormHttpBridge_getSandstormApi_Results(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSandstormApi_Results{st}, err
}

func ReadRootSandstormHttpBridge_getSandstormApi_Results(msg *capnp.Message) (SandstormHttpBridge_getSandstormApi_Results, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSandstormApi_Results{root.Struct()}, err
}

func (s SandstormHttpBridge_getSandstormApi_Results) String() string {
	str, _ := text.Marshal(0xc174fe273bd649ac, s.Struct)
	return str
}

func (s SandstormHttpBridge_getSandstormApi_Results) Api() grain.SandstormApi {
	p, _ := s.Struct.Ptr(0)
	return grain.SandstormApi{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_getSandstormApi_Results) HasApi() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSandstormApi_Results) SetApi(v grain.SandstormApi) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormHttpBridge_getSandstormApi_Results_List is a list of SandstormHttpBridge_getSandstormApi_Results.
type SandstormHttpBridge_getSandstormApi_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_getSandstormApi_Results creates a new list of SandstormHttpBridge_getSandstormApi_Results.
func NewSandstormHttpBridge_getSandstormApi_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSandstormApi_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_getSandstormApi_Results_List{l}, err
}

func (s SandstormHttpBridge_getSandstormApi_Results_List) At(i int) SandstormHttpBridge_getSandstormApi_Results {
	return SandstormHttpBridge_getSandstormApi_Results{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSandstormApi_Results_List) Set(i int, v SandstormHttpBridge_getSandstormApi_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSandstormApi_Results_Promise is a wrapper for a SandstormHttpBridge_getSandstormApi_Results promised by a client call.
type SandstormHttpBridge_getSandstormApi_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSandstormApi_Results_Promise) Struct() (SandstormHttpBridge_getSandstormApi_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSandstormApi_Results{s}, err
}

func (p SandstormHttpBridge_getSandstormApi_Results_Promise) Api() grain.SandstormApi {
	return grain.SandstormApi{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormHttpBridge_getSessionContext_Params struct{ capnp.Struct }

func NewSandstormHttpBridge_getSessionContext_Params(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSessionContext_Params{st}, err
}

func NewRootSandstormHttpBridge_getSessionContext_Params(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSessionContext_Params{st}, err
}

func ReadRootSandstormHttpBridge_getSessionContext_Params(msg *capnp.Message) (SandstormHttpBridge_getSessionContext_Params, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSessionContext_Params{root.Struct()}, err
}

func (s SandstormHttpBridge_getSessionContext_Params) String() string {
	str, _ := text.Marshal(0xc2a480fa3863a3fa, s.Struct)
	return str
}

func (s SandstormHttpBridge_getSessionContext_Params) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SandstormHttpBridge_getSessionContext_Params) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSessionContext_Params) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SandstormHttpBridge_getSessionContext_Params) SetId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// SandstormHttpBridge_getSessionContext_Params_List is a list of SandstormHttpBridge_getSessionContext_Params.
type SandstormHttpBridge_getSessionContext_Params_List struct{ capnp.List }

// NewSandstormHttpBridge_getSessionContext_Params creates a new list of SandstormHttpBridge_getSessionContext_Params.
func NewSandstormHttpBridge_getSessionContext_Params_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSessionContext_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_getSessionContext_Params_List{l}, err
}

func (s SandstormHttpBridge_getSessionContext_Params_List) At(i int) SandstormHttpBridge_getSessionContext_Params {
	return SandstormHttpBridge_getSessionContext_Params{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSessionContext_Params_List) Set(i int, v SandstormHttpBridge_getSessionContext_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSessionContext_Params_Promise is a wrapper for a SandstormHttpBridge_getSessionContext_Params promised by a client call.
type SandstormHttpBridge_getSessionContext_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSessionContext_Params_Promise) Struct() (SandstormHttpBridge_getSessionContext_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSessionContext_Params{s}, err
}

type SandstormHttpBridge_getSessionContext_Results struct{ capnp.Struct }

func NewSandstormHttpBridge_getSessionContext_Results(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSessionContext_Results{st}, err
}

func NewRootSandstormHttpBridge_getSessionContext_Results(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSessionContext_Results{st}, err
}

func ReadRootSandstormHttpBridge_getSessionContext_Results(msg *capnp.Message) (SandstormHttpBridge_getSessionContext_Results, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSessionContext_Results{root.Struct()}, err
}

func (s SandstormHttpBridge_getSessionContext_Results) String() string {
	str, _ := text.Marshal(0xb84ffa3322e48dbd, s.Struct)
	return str
}

func (s SandstormHttpBridge_getSessionContext_Results) Context() grain.SessionContext {
	p, _ := s.Struct.Ptr(0)
	return grain.SessionContext{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_getSessionContext_Results) HasContext() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSessionContext_Results) SetContext(v grain.SessionContext) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormHttpBridge_getSessionContext_Results_List is a list of SandstormHttpBridge_getSessionContext_Results.
type SandstormHttpBridge_getSessionContext_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_getSessionContext_Results creates a new list of SandstormHttpBridge_getSessionContext_Results.
func NewSandstormHttpBridge_getSessionContext_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSessionContext_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_getSessionContext_Results_List{l}, err
}

func (s SandstormHttpBridge_getSessionContext_Results_List) At(i int) SandstormHttpBridge_getSessionContext_Results {
	return SandstormHttpBridge_getSessionContext_Results{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSessionContext_Results_List) Set(i int, v SandstormHttpBridge_getSessionContext_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSessionContext_Results_Promise is a wrapper for a SandstormHttpBridge_getSessionContext_Results promised by a client call.
type SandstormHttpBridge_getSessionContext_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSessionContext_Results_Promise) Struct() (SandstormHttpBridge_getSessionContext_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSessionContext_Results{s}, err
}

func (p SandstormHttpBridge_getSessionContext_Results_Promise) Context() grain.SessionContext {
	return grain.SessionContext{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormHttpBridge_getSavedIdentity_Params struct{ capnp.Struct }

func NewSandstormHttpBridge_getSavedIdentity_Params(s *capnp.Segment) (SandstormHttpBridge_getSavedIdentity_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSavedIdentity_Params{st}, err
}

func NewRootSandstormHttpBridge_getSavedIdentity_Params(s *capnp.Segment) (SandstormHttpBridge_getSavedIdentity_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSavedIdentity_Params{st}, err
}

func ReadRootSandstormHttpBridge_getSavedIdentity_Params(msg *capnp.Message) (SandstormHttpBridge_getSavedIdentity_Params, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSavedIdentity_Params{root.Struct()}, err
}

func (s SandstormHttpBridge_getSavedIdentity_Params) String() string {
	str, _ := text.Marshal(0x930d8201947c86bf, s.Struct)
	return str
}

func (s SandstormHttpBridge_getSavedIdentity_Params) IdentityId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SandstormHttpBridge_getSavedIdentity_Params) HasIdentityId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSavedIdentity_Params) IdentityIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SandstormHttpBridge_getSavedIdentity_Params) SetIdentityId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// SandstormHttpBridge_getSavedIdentity_Params_List is a list of SandstormHttpBridge_getSavedIdentity_Params.
type SandstormHttpBridge_getSavedIdentity_Params_List struct{ capnp.List }

// NewSandstormHttpBridge_getSavedIdentity_Params creates a new list of SandstormHttpBridge_getSavedIdentity_Params.
func NewSandstormHttpBridge_getSavedIdentity_Params_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSavedIdentity_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_getSavedIdentity_Params_List{l}, err
}

func (s SandstormHttpBridge_getSavedIdentity_Params_List) At(i int) SandstormHttpBridge_getSavedIdentity_Params {
	return SandstormHttpBridge_getSavedIdentity_Params{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSavedIdentity_Params_List) Set(i int, v SandstormHttpBridge_getSavedIdentity_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSavedIdentity_Params_Promise is a wrapper for a SandstormHttpBridge_getSavedIdentity_Params promised by a client call.
type SandstormHttpBridge_getSavedIdentity_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSavedIdentity_Params_Promise) Struct() (SandstormHttpBridge_getSavedIdentity_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSavedIdentity_Params{s}, err
}

type SandstormHttpBridge_getSavedIdentity_Results struct{ capnp.Struct }

func NewSandstormHttpBridge_getSavedIdentity_Results(s *capnp.Segment) (SandstormHttpBridge_getSavedIdentity_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSavedIdentity_Results{st}, err
}

func NewRootSandstormHttpBridge_getSavedIdentity_Results(s *capnp.Segment) (SandstormHttpBridge_getSavedIdentity_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_getSavedIdentity_Results{st}, err
}

func ReadRootSandstormHttpBridge_getSavedIdentity_Results(msg *capnp.Message) (SandstormHttpBridge_getSavedIdentity_Results, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_getSavedIdentity_Results{root.Struct()}, err
}

func (s SandstormHttpBridge_getSavedIdentity_Results) String() string {
	str, _ := text.Marshal(0xc2ba6a7e9ab2e369, s.Struct)
	return str
}

func (s SandstormHttpBridge_getSavedIdentity_Results) Identity() identity.Identity {
	p, _ := s.Struct.Ptr(0)
	return identity.Identity{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_getSavedIdentity_Results) HasIdentity() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSavedIdentity_Results) SetIdentity(v identity.Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormHttpBridge_getSavedIdentity_Results_List is a list of SandstormHttpBridge_getSavedIdentity_Results.
type SandstormHttpBridge_getSavedIdentity_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_getSavedIdentity_Results creates a new list of SandstormHttpBridge_getSavedIdentity_Results.
func NewSandstormHttpBridge_getSavedIdentity_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSavedIdentity_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_getSavedIdentity_Results_List{l}, err
}

func (s SandstormHttpBridge_getSavedIdentity_Results_List) At(i int) SandstormHttpBridge_getSavedIdentity_Results {
	return SandstormHttpBridge_getSavedIdentity_Results{s.List.Struct(i)}
}

func (s SandstormHttpBridge_getSavedIdentity_Results_List) Set(i int, v SandstormHttpBridge_getSavedIdentity_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_getSavedIdentity_Results_Promise is a wrapper for a SandstormHttpBridge_getSavedIdentity_Results promised by a client call.
type SandstormHttpBridge_getSavedIdentity_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_getSavedIdentity_Results_Promise) Struct() (SandstormHttpBridge_getSavedIdentity_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_getSavedIdentity_Results{s}, err
}

func (p SandstormHttpBridge_getSavedIdentity_Results_Promise) Identity() identity.Identity {
	return identity.Identity{Client: p.Pipeline.GetPipeline(0).Client()}
}

const schema_ac137d236832bb1e = "x\xda2X\xc3\xea\xc0d\xc8\xda\xaf\xce\xc0\x10|\x86" +
	"\x91\x95\xed\x7f\xca\xd1\x94\xf9G'\\ic\x104`" +
	"d``ag`0^+l\xc5\xc4\xc0\xf2\x7f\x7f" +
	"[\xcd\x14\xc6&\xde\xc9\x10\x19VF\x90\xd4Ra'" +
	"&\x06F\xe1\xbd\xc2\xf6\x0c\x8c\xff\xbbn\xee\xe5\xe4\xed" +
	"O_\xcb (\xc3\xfc_n\xb7Q\x86r\xad\xf0\x1a" +
	"\x06\x06F\xe3\x87\xc2M\x8c\xc2_\x85\x81\x1a\x84?\x0a" +
	"\x1f\x17\xae\x15\x01\xb2\xfe\xef\xed}\xa2d\xfc\xcb\x7f\x07" +
	"\xb2q\xa9\"A \xe3\x1aE@\xc6\xad\xf1\xbcf\xad" +
	"\xfe\xaf\xe4 \xb2\x82\x85\"`\xfbv\x82\x15\xfcZ\x9c" +
	"l\xf1\xaba\xc9!d\x057E\xbc@\x0a\xbe\x82\x15" +
	"d>\xde4\xab.k\x17\x8a\x02QQ\xb0\x02CQ" +
	"{\x86\xffP\xb8\xf4\x7fqb^JqI~\x11{" +
	"\xaenFII\x81nRQfJz\xaa^rb" +
	"A^\x81U0T2\xd7\x03(\xe5\x04\x91IO-" +
	"\x81\x0b;\x16d\xaa\xd8\x07$\x16%\xe6\x16\xc3\x0d\xe2" +
	" \xc9\xa0\xb2\xd4\x14\xcf\x94\xd4\xbc\x92\xcc\x92J\x15\x88" +
	"A\x0c\x81,\xcc,\xc0\xe0\x07:[\x907\x8a\x81!" +
	"\x90\x87\x991P\x82\x09\xe8%\xa8:\x06f\xcf\x14F" +
	"\x1e\x06& f\x84\xdb\xcaJ\xa4\xad\x0c\x01\x8c\x8c@" +
	"\x13Y\x81\xb1\x00\x8bnFXh\x0b\x06610\x09" +
	"z\xb23\"\xc2\x97\x11\x16U\x82\xb6\x93\x80r\x96\xec" +
	"\x8cL\xf0\xc4\xc0\x08\x0bcA\xdd.\xa0\x9c&\xfb\x7f" +
	"X\xd00B\xc3\x86\xc1\x81\x11,\x96Z\\\x9c\xc9\x98" +
	"\x9f\xe7\x9c\x9fW\x92Z\xc1X\x02\x15\x05\xf9\x9d\x11\xe6" +
	"y\x06\xa0Z\xa0\xd3\xc8\x0bE\x90\xf90\xe3KT\x82" +
	"R\x8bKsJ\x18\x8b\x91\xc3\xd1\x09\x18\x8e\x1c\xc0p" +
	"\x14ab\xacO\x86\xa8c\x14\xfco\xf699J\xd6" +
	"\xc1n?0\xa12\x0a2\x90i9rZ\x80X\x8d" +
	"\x1a\x85J\x08\xab\xd9\x13\x0b2\x81\xd6\x9ef*\xd0:" +
	"\xef\x9fz\x89\"kQ\xfd\x0cM:(\x16K!," +
	"f\xce\xc4L1\xe4\xa7S\x98/Ql\xf3\xc2\x96R" +
	"\x81>\x14\xfc/q7pe\xd5\x8c\x96\x03P\xef\x02" +
	"\x02\x00\x00\xff\xff\x1e)f\xef"

func init() {
	schemas.Register(schema_ac137d236832bb1e,
		0x86d490c59f64c564,
		0x930d8201947c86bf,
		0xad678f0d09bdd98a,
		0xb84ffa3322e48dbd,
		0xc174fe273bd649ac,
		0xc2a480fa3863a3fa,
		0xc2ba6a7e9ab2e369)
}
