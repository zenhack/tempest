package sandstormhttpbridge

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	grain "zenhack.net/go/sandstorm/capnp/sandstorm/grain"
	capnp "zombiezen.com/go/capnproto2"
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

type SandstormHttpBridge_Server interface {
	GetSandstormApi(SandstormHttpBridge_getSandstormApi) error

	GetSessionContext(SandstormHttpBridge_getSessionContext) error
}

func SandstormHttpBridge_ServerToClient(s SandstormHttpBridge_Server) SandstormHttpBridge {
	c, _ := s.(server.Closer)
	return SandstormHttpBridge{Client: server.New(SandstormHttpBridge_Methods(nil, s), c)}
}

func SandstormHttpBridge_Methods(methods []server.Method, s SandstormHttpBridge_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
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

type SandstormHttpBridge_getSandstormApi_Params struct{ capnp.Struct }

func NewSandstormHttpBridge_getSandstormApi_Params(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Params{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Params{st}, nil
}

func NewRootSandstormHttpBridge_getSandstormApi_Params(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Params{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Params{st}, nil
}

func ReadRootSandstormHttpBridge_getSandstormApi_Params(msg *capnp.Message) (SandstormHttpBridge_getSandstormApi_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Params{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Params{root.Struct()}, nil
}

// SandstormHttpBridge_getSandstormApi_Params_List is a list of SandstormHttpBridge_getSandstormApi_Params.
type SandstormHttpBridge_getSandstormApi_Params_List struct{ capnp.List }

// NewSandstormHttpBridge_getSandstormApi_Params creates a new list of SandstormHttpBridge_getSandstormApi_Params.
func NewSandstormHttpBridge_getSandstormApi_Params_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSandstormApi_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Params_List{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Params_List{l}, nil
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
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Results{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Results{st}, nil
}

func NewRootSandstormHttpBridge_getSandstormApi_Results(s *capnp.Segment) (SandstormHttpBridge_getSandstormApi_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Results{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Results{st}, nil
}

func ReadRootSandstormHttpBridge_getSandstormApi_Results(msg *capnp.Message) (SandstormHttpBridge_getSandstormApi_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Results{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Results{root.Struct()}, nil
}

func (s SandstormHttpBridge_getSandstormApi_Results) Api() grain.SandstormApi {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return grain.SandstormApi{}
	}
	return grain.SandstormApi{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_getSandstormApi_Results) HasApi() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSandstormApi_Results) SetApi(v grain.SandstormApi) error {

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

// SandstormHttpBridge_getSandstormApi_Results_List is a list of SandstormHttpBridge_getSandstormApi_Results.
type SandstormHttpBridge_getSandstormApi_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_getSandstormApi_Results creates a new list of SandstormHttpBridge_getSandstormApi_Results.
func NewSandstormHttpBridge_getSandstormApi_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSandstormApi_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormHttpBridge_getSandstormApi_Results_List{}, err
	}
	return SandstormHttpBridge_getSandstormApi_Results_List{l}, nil
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
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Params{}, err
	}
	return SandstormHttpBridge_getSessionContext_Params{st}, nil
}

func NewRootSandstormHttpBridge_getSessionContext_Params(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Params{}, err
	}
	return SandstormHttpBridge_getSessionContext_Params{st}, nil
}

func ReadRootSandstormHttpBridge_getSessionContext_Params(msg *capnp.Message) (SandstormHttpBridge_getSessionContext_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Params{}, err
	}
	return SandstormHttpBridge_getSessionContext_Params{root.Struct()}, nil
}

func (s SandstormHttpBridge_getSessionContext_Params) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s SandstormHttpBridge_getSessionContext_Params) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSessionContext_Params) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Params_List{}, err
	}
	return SandstormHttpBridge_getSessionContext_Params_List{l}, nil
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
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Results{}, err
	}
	return SandstormHttpBridge_getSessionContext_Results{st}, nil
}

func NewRootSandstormHttpBridge_getSessionContext_Results(s *capnp.Segment) (SandstormHttpBridge_getSessionContext_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Results{}, err
	}
	return SandstormHttpBridge_getSessionContext_Results{st}, nil
}

func ReadRootSandstormHttpBridge_getSessionContext_Results(msg *capnp.Message) (SandstormHttpBridge_getSessionContext_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Results{}, err
	}
	return SandstormHttpBridge_getSessionContext_Results{root.Struct()}, nil
}

func (s SandstormHttpBridge_getSessionContext_Results) Context() grain.SessionContext {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return grain.SessionContext{}
	}
	return grain.SessionContext{Client: p.Interface().Client()}
}

func (s SandstormHttpBridge_getSessionContext_Results) HasContext() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormHttpBridge_getSessionContext_Results) SetContext(v grain.SessionContext) error {

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

// SandstormHttpBridge_getSessionContext_Results_List is a list of SandstormHttpBridge_getSessionContext_Results.
type SandstormHttpBridge_getSessionContext_Results_List struct{ capnp.List }

// NewSandstormHttpBridge_getSessionContext_Results creates a new list of SandstormHttpBridge_getSessionContext_Results.
func NewSandstormHttpBridge_getSessionContext_Results_List(s *capnp.Segment, sz int32) (SandstormHttpBridge_getSessionContext_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormHttpBridge_getSessionContext_Results_List{}, err
	}
	return SandstormHttpBridge_getSessionContext_Results_List{l}, nil
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
