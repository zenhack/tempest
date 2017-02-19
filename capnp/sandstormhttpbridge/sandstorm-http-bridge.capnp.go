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
func (c SandstormHttpBridge) SaveIdentity(ctx context.Context, params func(SandstormHttpBridge_saveIdentity_Params) error, opts ...capnp.CallOption) SandstormHttpBridge_saveIdentity_Results_Promise {
	if c.Client == nil {
		return SandstormHttpBridge_saveIdentity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      3,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "saveIdentity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormHttpBridge_saveIdentity_Params{Struct: s}) }
	}
	return SandstormHttpBridge_saveIdentity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormHttpBridge_Server interface {
	GetSandstormApi(SandstormHttpBridge_getSandstormApi) error

	GetSessionContext(SandstormHttpBridge_getSessionContext) error

	GetSavedIdentity(SandstormHttpBridge_getSavedIdentity) error

	SaveIdentity(SandstormHttpBridge_saveIdentity) error
}

func SandstormHttpBridge_ServerToClient(s SandstormHttpBridge_Server) SandstormHttpBridge {
	c, _ := s.(server.Closer)
	return SandstormHttpBridge{Client: server.New(SandstormHttpBridge_Methods(nil, s), c)}
}

func SandstormHttpBridge_Methods(methods []server.Method, s SandstormHttpBridge_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
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

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xad678f0d09bdd98a,
			MethodID:      3,
			InterfaceName: "sandstorm-http-bridge.capnp:SandstormHttpBridge",
			MethodName:    "saveIdentity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormHttpBridge_saveIdentity{c, opts, SandstormHttpBridge_saveIdentity_Params{Struct: p}, SandstormHttpBridge_saveIdentity_Results{Struct: r}}
			return s.SaveIdentity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
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

// SandstormHttpBridge_saveIdentity holds the arguments for a server call to SandstormHttpBridge.saveIdentity.
type SandstormHttpBridge_saveIdentity struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormHttpBridge_saveIdentity_Params
	Results SandstormHttpBridge_saveIdentity_Results
}

type SandstormHttpBridge_getSandstormApi_Params struct{ capnp.Struct }

// SandstormHttpBridge_getSandstormApi_Params_TypeID is the unique identifier for the type SandstormHttpBridge_getSandstormApi_Params.
const SandstormHttpBridge_getSandstormApi_Params_TypeID = 0x86d490c59f64c564

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

// SandstormHttpBridge_getSandstormApi_Results_TypeID is the unique identifier for the type SandstormHttpBridge_getSandstormApi_Results.
const SandstormHttpBridge_getSandstormApi_Results_TypeID = 0xc174fe273bd649ac

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

// SandstormHttpBridge_getSessionContext_Params_TypeID is the unique identifier for the type SandstormHttpBridge_getSessionContext_Params.
const SandstormHttpBridge_getSessionContext_Params_TypeID = 0xc2a480fa3863a3fa

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

// SandstormHttpBridge_getSessionContext_Results_TypeID is the unique identifier for the type SandstormHttpBridge_getSessionContext_Results.
const SandstormHttpBridge_getSessionContext_Results_TypeID = 0xb84ffa3322e48dbd

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

// SandstormHttpBridge_getSavedIdentity_Params_TypeID is the unique identifier for the type SandstormHttpBridge_getSavedIdentity_Params.
const SandstormHttpBridge_getSavedIdentity_Params_TypeID = 0x930d8201947c86bf

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

// SandstormHttpBridge_getSavedIdentity_Results_TypeID is the unique identifier for the type SandstormHttpBridge_getSavedIdentity_Results.
const SandstormHttpBridge_getSavedIdentity_Results_TypeID = 0xc2ba6a7e9ab2e369

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

type SandstormHttpBridge_saveIdentity_Params struct{ capnp.Struct }

// SandstormHttpBridge_saveIdentity_Params_TypeID is the unique identifier for the type SandstormHttpBridge_saveIdentity_Params.
const SandstormHttpBridge_saveIdentity_Params_TypeID = 0xd274ecd8fd2907d0

func NewSandstormHttpBridge_saveIdentity_Params(s *capnp.Segment) (SandstormHttpBridge_saveIdentity_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_saveIdentity_Params{st}, err
}

func NewRootSandstormHttpBridge_saveIdentity_Params(s *capnp.Segment) (SandstormHttpBridge_saveIdentity_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormHttpBridge_saveIdentity_Params{st}, err
}

func ReadRootSandstormHttpBridge_saveIdentity_Params(msg *capnp.Message) (SandstormHttpBridge_saveIdentity_Params, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_saveIdentity_Params{root.Struct()}, err
}

func (s SandstormHttpBridge_saveIdentity_Params) String() string {
	str, _ := text.Marshal(0xd274ecd8fd2907d0, s.Struct)
	return str
}

func (s SandstormHttpBridge_saveIdentity_Params) Identity() identity.Identity {
	p, _ := s.Struct.Ptr(0)
	return identity.Identity{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_saveIdentity_Params) HasIdentity() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_saveIdentity_Params) SetIdentity(v identity.Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormHttpBridge_saveIdentity_Params_List is a list of SandstormHttpBridge_saveIdentity_Params.
type SandstormHttpBridge_saveIdentity_Params_List struct{ capnp.List }

// NewSandstormHttpBridge_saveIdentity_Params creates a new list of SandstormHttpBridge_saveIdentity_Params.
func NewSandstormHttpBridge_saveIdentity_Params_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_saveIdentity_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormHttpBridge_saveIdentity_Params_List{l}, err
}

func (s SandstormHttpBridge_saveIdentity_Params_List) At(i int) SandstormHttpBridge_saveIdentity_Params {
	return SandstormHttpBridge_saveIdentity_Params{s.List.Struct(i)}
}

func (s SandstormHttpBridge_saveIdentity_Params_List) Set(i int, v SandstormHttpBridge_saveIdentity_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_saveIdentity_Params_Promise is a wrapper for a SandstormHttpBridge_saveIdentity_Params promised by a client call.
type SandstormHttpBridge_saveIdentity_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_saveIdentity_Params_Promise) Struct() (SandstormHttpBridge_saveIdentity_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_saveIdentity_Params{s}, err
}

func (p SandstormHttpBridge_saveIdentity_Params_Promise) Identity() identity.Identity {
	return identity.Identity{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormHttpBridge_saveIdentity_Results struct{ capnp.Struct }

// SandstormHttpBridge_saveIdentity_Results_TypeID is the unique identifier for the type SandstormHttpBridge_saveIdentity_Results.
const SandstormHttpBridge_saveIdentity_Results_TypeID = 0xb1b610269c6a2190

func NewSandstormHttpBridge_saveIdentity_Results(s *capnp.Segment) (SandstormHttpBridge_saveIdentity_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormHttpBridge_saveIdentity_Results{st}, err
}

func NewRootSandstormHttpBridge_saveIdentity_Results(s *capnp.Segment) (SandstormHttpBridge_saveIdentity_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormHttpBridge_saveIdentity_Results{st}, err
}

func ReadRootSandstormHttpBridge_saveIdentity_Results(msg *capnp.Message) (SandstormHttpBridge_saveIdentity_Results, error) {
	root, err := msg.RootPtr()
	return SandstormHttpBridge_saveIdentity_Results{root.Struct()}, err
}

func (s SandstormHttpBridge_saveIdentity_Results) String() string {
	str, _ := text.Marshal(0xb1b610269c6a2190, s.Struct)
	return str
}

// SandstormHttpBridge_saveIdentity_Results_List is a list of SandstormHttpBridge_saveIdentity_Results.
type SandstormHttpBridge_saveIdentity_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_saveIdentity_Results creates a new list of SandstormHttpBridge_saveIdentity_Results.
func NewSandstormHttpBridge_saveIdentity_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_saveIdentity_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormHttpBridge_saveIdentity_Results_List{l}, err
}

func (s SandstormHttpBridge_saveIdentity_Results_List) At(i int) SandstormHttpBridge_saveIdentity_Results {
	return SandstormHttpBridge_saveIdentity_Results{s.List.Struct(i)}
}

func (s SandstormHttpBridge_saveIdentity_Results_List) Set(i int, v SandstormHttpBridge_saveIdentity_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormHttpBridge_saveIdentity_Results_Promise is a wrapper for a SandstormHttpBridge_saveIdentity_Results promised by a client call.
type SandstormHttpBridge_saveIdentity_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormHttpBridge_saveIdentity_Results_Promise) Struct() (SandstormHttpBridge_saveIdentity_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormHttpBridge_saveIdentity_Results{s}, err
}

const schema_ac137d236832bb1e = "x\xda\xa4\x93Oh\x13O\x14\xc7\xdf\x9bM\xba\xfdA" +
	"K;$\x1b\xca\x0f\x15\xac\xd5\xd2B\xa5\xb5(R\xd1" +
	"\xc6z1\x85\xd2l{\xb2\xe0a\xcd.\xed\x86&Y" +
	"\xb3\xd3\xa2\xe2?\x82\x14,\x88E=\x88\x08\x1e\xf4\"" +
	"\xb4\x07=\x88(\xa1*\xe4\xe4A)\"\x06\xecA\xf4" +
	"\xe0\xc1\x8bx\x10\x03:2\xd9?\xd9\xa8\x07\x1b\xc9%" +
	"\xbc\xf7\xdd\xf7\x99y\xdf\xef\xf4w7\xc5\xc9@\xf8k" +
	"\x0f\xc0\xe4G\x0c7q\xbd\xa4\xdf*-\xbdZ\x00\xda" +
	"\x8f\x00!\x19`\xf0]t\x88@\x88\xaf.\x9c\xbe\x86" +
	"\x85\xd6\xabN'\x8c\xa2U\x8e\x8e\x10\xc0\xc8\xe7\xe80" +
	" _,\x17\xffk\xbd<\xbd\x02t\x93\xc4\xb7<\xde" +
	"5\xb3\xedLd\x19\x00\x07\x15\xa5\x80\x91\x1eE|\xb0" +
	"]\x911\xb2.\xfe\xf2\xa5\xad\xe9\x9b;\xda\x1f\xdc\x0b" +
	"\x90JJ\xaf \x15/}\xe8\x1c\xac\x8c?\x0c\x92\x8a" +
	"\xca\x84 \x95\x15AZN\xbc\xde\xd7\xfd\x83=\x0d\x0a" +
	"\xbe)\xd5\xa3\xd0\x98\x10Tn\xa7\xf6V\xce\xdfy\x16" +
	"\x14\x0c\xc4F\x85`\xac*0\xdf\xdf\xbfq6\xfd\xa8" +
	"N\x90q\x04\x17\xab\x82\x97r\xcf\xf77\x9f\xd8ZP" +
	"\xb0\x12\xeb\x14\x82Rl\x18\xb8\xfb{\xcbm-\xab\xdb" +
	",\x97\x973}3\x8cY}\xc7\xf2\xa6>m\xecL" +
	"iV\xd6\x1a\x9at\x9b\x99\xc3\x8cY#Ng\xda`" +
	"~\xf9\xa0ev\x0d'\xb5\xbc\x96\xb1\xfdA\xcd\x1b\x1a" +
	"4o\xe8\x09\xdd\xc82\x93\x9d\xecr\x06\x81\x1a\x92B" +
	"\x00!\x04\xa0\xadS\x00j\x8b\x84j\x07An\xba:" +
	"\x90\x12:\xb6\x00\xc1\x16@\x9f\x1a\xfeK*$\x11\xd5" +
	"v)\x0c\xe0G\x05=;\xe8\xf1\x02\x10j\xcaX3" +
	"\x00=/\xe9\xd1+@\xe8\x11\x19\x89\x1f$\xf4L\xa0" +
	"c\x8b@hBF\xc9_;z\xe9\xa0\xfb\xd3@\xe8" +
	"n\x99{kCwo\x10\xc7j\xcd\xb0m\x13s\xd9" +
	"C\xb9,3N s\xabb/\xe8-\x06\x84\xd6\xd6" +
	"\xe6\x0dQ\x806Q\x8ac\x12q\xe3\xdeyC\xaa\xeb" +
	"\x9e0\xec\xb6\xb9Y\xd6\xa0s\xe2\xdc\xde\xb1\x99\x985" +
	"7\xcb\xd0\x0ez7\x02\xa06K\xa8F\x09\x9eK9" +
	":\xa4|\xcf\x97\xd4\xd4\xe6\xf8\x81U\x00D\x1a0\xb0" +
	"\xb9\xd1\xfc9\xe8\xfa\xd8t\xd6\xd0\xb2f\x99H\xf9s" +
	"b\xf5\xbe\x187\xd6\xfe\x09[\x7fg7\xaeu\xe0\xff" +
	"k`\xc9\xfc=\xa5\x8d\xbf\x0d\xef\x96u\xb4\xd1?\xbd" +
	"\x0e\x00\xa4\xbcc]\xbd{\xea\xfa\x85'\xbf^\xb7\xb1" +
	"\xa4$\xb5\xbc\xace\xec\x86\xd0?\x03\x00\x00\xff\xffH" +
	"4\xc7b"

func init() {
	schemas.Register(schema_ac137d236832bb1e,
		0x86d490c59f64c564,
		0x930d8201947c86bf,
		0xad678f0d09bdd98a,
		0xb1b610269c6a2190,
		0xb84ffa3322e48dbd,
		0xc174fe273bd649ac,
		0xc2a480fa3863a3fa,
		0xc2ba6a7e9ab2e369,
		0xd274ecd8fd2907d0)
}
