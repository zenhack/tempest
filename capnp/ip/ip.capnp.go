package ip

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type IpNetwork struct{ Client capnp.Client }

func (c IpNetwork) GetRemoteHost(ctx context.Context, params func(IpNetwork_getRemoteHost_Params) error, opts ...capnp.CallOption) IpNetwork_getRemoteHost_Results_Promise {
	if c.Client == nil {
		return IpNetwork_getRemoteHost_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHost",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpNetwork_getRemoteHost_Params{Struct: s}) }
	}
	return IpNetwork_getRemoteHost_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c IpNetwork) GetRemoteHostByName(ctx context.Context, params func(IpNetwork_getRemoteHostByName_Params) error, opts ...capnp.CallOption) IpNetwork_getRemoteHostByName_Results_Promise {
	if c.Client == nil {
		return IpNetwork_getRemoteHostByName_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHostByName",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpNetwork_getRemoteHostByName_Params{Struct: s}) }
	}
	return IpNetwork_getRemoteHostByName_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type IpNetwork_Server interface {
	GetRemoteHost(IpNetwork_getRemoteHost) error

	GetRemoteHostByName(IpNetwork_getRemoteHostByName) error
}

func IpNetwork_ServerToClient(s IpNetwork_Server) IpNetwork {
	c, _ := s.(server.Closer)
	return IpNetwork{Client: server.New(IpNetwork_Methods(nil, s), c)}
}

func IpNetwork_Methods(methods []server.Method, s IpNetwork_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHost",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpNetwork_getRemoteHost{c, opts, IpNetwork_getRemoteHost_Params{Struct: p}, IpNetwork_getRemoteHost_Results{Struct: r}}
			return s.GetRemoteHost(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHostByName",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpNetwork_getRemoteHostByName{c, opts, IpNetwork_getRemoteHostByName_Params{Struct: p}, IpNetwork_getRemoteHostByName_Results{Struct: r}}
			return s.GetRemoteHostByName(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// IpNetwork_getRemoteHost holds the arguments for a server call to IpNetwork.getRemoteHost.
type IpNetwork_getRemoteHost struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpNetwork_getRemoteHost_Params
	Results IpNetwork_getRemoteHost_Results
}

// IpNetwork_getRemoteHostByName holds the arguments for a server call to IpNetwork.getRemoteHostByName.
type IpNetwork_getRemoteHostByName struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpNetwork_getRemoteHostByName_Params
	Results IpNetwork_getRemoteHostByName_Results
}

type IpNetwork_getRemoteHost_Params struct{ capnp.Struct }

func NewIpNetwork_getRemoteHost_Params(s *capnp.Segment) (IpNetwork_getRemoteHost_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHost_Params{st}, err
}

func NewRootIpNetwork_getRemoteHost_Params(s *capnp.Segment) (IpNetwork_getRemoteHost_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHost_Params{st}, err
}

func ReadRootIpNetwork_getRemoteHost_Params(msg *capnp.Message) (IpNetwork_getRemoteHost_Params, error) {
	root, err := msg.RootPtr()
	return IpNetwork_getRemoteHost_Params{root.Struct()}, err
}

func (s IpNetwork_getRemoteHost_Params) String() string {
	str, _ := text.Marshal(0xdd1700c1eb725eb4, s.Struct)
	return str
}

func (s IpNetwork_getRemoteHost_Params) Address() (IpAddress, error) {
	p, err := s.Struct.Ptr(0)
	return IpAddress{Struct: p.Struct()}, err
}

func (s IpNetwork_getRemoteHost_Params) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHost_Params) SetAddress(v IpAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAddress sets the address field to a newly
// allocated IpAddress struct, preferring placement in s's segment.
func (s IpNetwork_getRemoteHost_Params) NewAddress() (IpAddress, error) {
	ss, err := NewIpAddress(s.Struct.Segment())
	if err != nil {
		return IpAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// IpNetwork_getRemoteHost_Params_List is a list of IpNetwork_getRemoteHost_Params.
type IpNetwork_getRemoteHost_Params_List struct{ capnp.List }

// NewIpNetwork_getRemoteHost_Params creates a new list of IpNetwork_getRemoteHost_Params.
func NewIpNetwork_getRemoteHost_Params_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHost_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpNetwork_getRemoteHost_Params_List{l}, err
}

func (s IpNetwork_getRemoteHost_Params_List) At(i int) IpNetwork_getRemoteHost_Params {
	return IpNetwork_getRemoteHost_Params{s.List.Struct(i)}
}

func (s IpNetwork_getRemoteHost_Params_List) Set(i int, v IpNetwork_getRemoteHost_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_getRemoteHost_Params_Promise is a wrapper for a IpNetwork_getRemoteHost_Params promised by a client call.
type IpNetwork_getRemoteHost_Params_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_getRemoteHost_Params_Promise) Struct() (IpNetwork_getRemoteHost_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_getRemoteHost_Params{s}, err
}

func (p IpNetwork_getRemoteHost_Params_Promise) Address() IpAddress_Promise {
	return IpAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type IpNetwork_getRemoteHost_Results struct{ capnp.Struct }

func NewIpNetwork_getRemoteHost_Results(s *capnp.Segment) (IpNetwork_getRemoteHost_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHost_Results{st}, err
}

func NewRootIpNetwork_getRemoteHost_Results(s *capnp.Segment) (IpNetwork_getRemoteHost_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHost_Results{st}, err
}

func ReadRootIpNetwork_getRemoteHost_Results(msg *capnp.Message) (IpNetwork_getRemoteHost_Results, error) {
	root, err := msg.RootPtr()
	return IpNetwork_getRemoteHost_Results{root.Struct()}, err
}

func (s IpNetwork_getRemoteHost_Results) String() string {
	str, _ := text.Marshal(0xb57bd5aef30c4b61, s.Struct)
	return str
}

func (s IpNetwork_getRemoteHost_Results) Host() IpRemoteHost {
	p, _ := s.Struct.Ptr(0)
	return IpRemoteHost{Client: p.Interface().Client()}
}

func (s IpNetwork_getRemoteHost_Results) HasHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHost_Results) SetHost(v IpRemoteHost) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpNetwork_getRemoteHost_Results_List is a list of IpNetwork_getRemoteHost_Results.
type IpNetwork_getRemoteHost_Results_List struct{ capnp.List }

// NewIpNetwork_getRemoteHost_Results creates a new list of IpNetwork_getRemoteHost_Results.
func NewIpNetwork_getRemoteHost_Results_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHost_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpNetwork_getRemoteHost_Results_List{l}, err
}

func (s IpNetwork_getRemoteHost_Results_List) At(i int) IpNetwork_getRemoteHost_Results {
	return IpNetwork_getRemoteHost_Results{s.List.Struct(i)}
}

func (s IpNetwork_getRemoteHost_Results_List) Set(i int, v IpNetwork_getRemoteHost_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_getRemoteHost_Results_Promise is a wrapper for a IpNetwork_getRemoteHost_Results promised by a client call.
type IpNetwork_getRemoteHost_Results_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_getRemoteHost_Results_Promise) Struct() (IpNetwork_getRemoteHost_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_getRemoteHost_Results{s}, err
}

func (p IpNetwork_getRemoteHost_Results_Promise) Host() IpRemoteHost {
	return IpRemoteHost{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpNetwork_getRemoteHostByName_Params struct{ capnp.Struct }

func NewIpNetwork_getRemoteHostByName_Params(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHostByName_Params{st}, err
}

func NewRootIpNetwork_getRemoteHostByName_Params(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHostByName_Params{st}, err
}

func ReadRootIpNetwork_getRemoteHostByName_Params(msg *capnp.Message) (IpNetwork_getRemoteHostByName_Params, error) {
	root, err := msg.RootPtr()
	return IpNetwork_getRemoteHostByName_Params{root.Struct()}, err
}

func (s IpNetwork_getRemoteHostByName_Params) String() string {
	str, _ := text.Marshal(0x9d5f1f6efcf7bbc4, s.Struct)
	return str
}

func (s IpNetwork_getRemoteHostByName_Params) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s IpNetwork_getRemoteHostByName_Params) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHostByName_Params) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s IpNetwork_getRemoteHostByName_Params) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// IpNetwork_getRemoteHostByName_Params_List is a list of IpNetwork_getRemoteHostByName_Params.
type IpNetwork_getRemoteHostByName_Params_List struct{ capnp.List }

// NewIpNetwork_getRemoteHostByName_Params creates a new list of IpNetwork_getRemoteHostByName_Params.
func NewIpNetwork_getRemoteHostByName_Params_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHostByName_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpNetwork_getRemoteHostByName_Params_List{l}, err
}

func (s IpNetwork_getRemoteHostByName_Params_List) At(i int) IpNetwork_getRemoteHostByName_Params {
	return IpNetwork_getRemoteHostByName_Params{s.List.Struct(i)}
}

func (s IpNetwork_getRemoteHostByName_Params_List) Set(i int, v IpNetwork_getRemoteHostByName_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_getRemoteHostByName_Params_Promise is a wrapper for a IpNetwork_getRemoteHostByName_Params promised by a client call.
type IpNetwork_getRemoteHostByName_Params_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_getRemoteHostByName_Params_Promise) Struct() (IpNetwork_getRemoteHostByName_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_getRemoteHostByName_Params{s}, err
}

type IpNetwork_getRemoteHostByName_Results struct{ capnp.Struct }

func NewIpNetwork_getRemoteHostByName_Results(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHostByName_Results{st}, err
}

func NewRootIpNetwork_getRemoteHostByName_Results(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_getRemoteHostByName_Results{st}, err
}

func ReadRootIpNetwork_getRemoteHostByName_Results(msg *capnp.Message) (IpNetwork_getRemoteHostByName_Results, error) {
	root, err := msg.RootPtr()
	return IpNetwork_getRemoteHostByName_Results{root.Struct()}, err
}

func (s IpNetwork_getRemoteHostByName_Results) String() string {
	str, _ := text.Marshal(0xd14a2ec2bad45f69, s.Struct)
	return str
}

func (s IpNetwork_getRemoteHostByName_Results) Host() IpRemoteHost {
	p, _ := s.Struct.Ptr(0)
	return IpRemoteHost{Client: p.Interface().Client()}
}

func (s IpNetwork_getRemoteHostByName_Results) HasHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHostByName_Results) SetHost(v IpRemoteHost) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpNetwork_getRemoteHostByName_Results_List is a list of IpNetwork_getRemoteHostByName_Results.
type IpNetwork_getRemoteHostByName_Results_List struct{ capnp.List }

// NewIpNetwork_getRemoteHostByName_Results creates a new list of IpNetwork_getRemoteHostByName_Results.
func NewIpNetwork_getRemoteHostByName_Results_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHostByName_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpNetwork_getRemoteHostByName_Results_List{l}, err
}

func (s IpNetwork_getRemoteHostByName_Results_List) At(i int) IpNetwork_getRemoteHostByName_Results {
	return IpNetwork_getRemoteHostByName_Results{s.List.Struct(i)}
}

func (s IpNetwork_getRemoteHostByName_Results_List) Set(i int, v IpNetwork_getRemoteHostByName_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_getRemoteHostByName_Results_Promise is a wrapper for a IpNetwork_getRemoteHostByName_Results promised by a client call.
type IpNetwork_getRemoteHostByName_Results_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_getRemoteHostByName_Results_Promise) Struct() (IpNetwork_getRemoteHostByName_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_getRemoteHostByName_Results{s}, err
}

func (p IpNetwork_getRemoteHostByName_Results_Promise) Host() IpRemoteHost {
	return IpRemoteHost{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpAddress struct{ capnp.Struct }

func NewIpAddress(s *capnp.Segment) (IpAddress, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return IpAddress{st}, err
}

func NewRootIpAddress(s *capnp.Segment) (IpAddress, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return IpAddress{st}, err
}

func ReadRootIpAddress(msg *capnp.Message) (IpAddress, error) {
	root, err := msg.RootPtr()
	return IpAddress{root.Struct()}, err
}

func (s IpAddress) String() string {
	str, _ := text.Marshal(0xeeb98f9937d32c0b, s.Struct)
	return str
}

func (s IpAddress) Lower64() uint64 {
	return s.Struct.Uint64(0)
}

func (s IpAddress) SetLower64(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s IpAddress) Upper64() uint64 {
	return s.Struct.Uint64(8)
}

func (s IpAddress) SetUpper64(v uint64) {
	s.Struct.SetUint64(8, v)
}

// IpAddress_List is a list of IpAddress.
type IpAddress_List struct{ capnp.List }

// NewIpAddress creates a new list of IpAddress.
func NewIpAddress_List(s *capnp.Segment, sz int32) (IpAddress_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return IpAddress_List{l}, err
}

func (s IpAddress_List) At(i int) IpAddress { return IpAddress{s.List.Struct(i)} }

func (s IpAddress_List) Set(i int, v IpAddress) error { return s.List.SetStruct(i, v.Struct) }

// IpAddress_Promise is a wrapper for a IpAddress promised by a client call.
type IpAddress_Promise struct{ *capnp.Pipeline }

func (p IpAddress_Promise) Struct() (IpAddress, error) {
	s, err := p.Pipeline.Struct()
	return IpAddress{s}, err
}

type IpInterface struct{ Client capnp.Client }

func (c IpInterface) ListenTcp(ctx context.Context, params func(IpInterface_listenTcp_Params) error, opts ...capnp.CallOption) IpInterface_listenTcp_Results_Promise {
	if c.Client == nil {
		return IpInterface_listenTcp_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenTcp",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpInterface_listenTcp_Params{Struct: s}) }
	}
	return IpInterface_listenTcp_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c IpInterface) ListenUdp(ctx context.Context, params func(IpInterface_listenUdp_Params) error, opts ...capnp.CallOption) IpInterface_listenUdp_Results_Promise {
	if c.Client == nil {
		return IpInterface_listenUdp_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenUdp",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpInterface_listenUdp_Params{Struct: s}) }
	}
	return IpInterface_listenUdp_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type IpInterface_Server interface {
	ListenTcp(IpInterface_listenTcp) error

	ListenUdp(IpInterface_listenUdp) error
}

func IpInterface_ServerToClient(s IpInterface_Server) IpInterface {
	c, _ := s.(server.Closer)
	return IpInterface{Client: server.New(IpInterface_Methods(nil, s), c)}
}

func IpInterface_Methods(methods []server.Method, s IpInterface_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenTcp",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpInterface_listenTcp{c, opts, IpInterface_listenTcp_Params{Struct: p}, IpInterface_listenTcp_Results{Struct: r}}
			return s.ListenTcp(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenUdp",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpInterface_listenUdp{c, opts, IpInterface_listenUdp_Params{Struct: p}, IpInterface_listenUdp_Results{Struct: r}}
			return s.ListenUdp(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// IpInterface_listenTcp holds the arguments for a server call to IpInterface.listenTcp.
type IpInterface_listenTcp struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpInterface_listenTcp_Params
	Results IpInterface_listenTcp_Results
}

// IpInterface_listenUdp holds the arguments for a server call to IpInterface.listenUdp.
type IpInterface_listenUdp struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpInterface_listenUdp_Params
	Results IpInterface_listenUdp_Results
}

type IpInterface_listenTcp_Params struct{ capnp.Struct }

func NewIpInterface_listenTcp_Params(s *capnp.Segment) (IpInterface_listenTcp_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpInterface_listenTcp_Params{st}, err
}

func NewRootIpInterface_listenTcp_Params(s *capnp.Segment) (IpInterface_listenTcp_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpInterface_listenTcp_Params{st}, err
}

func ReadRootIpInterface_listenTcp_Params(msg *capnp.Message) (IpInterface_listenTcp_Params, error) {
	root, err := msg.RootPtr()
	return IpInterface_listenTcp_Params{root.Struct()}, err
}

func (s IpInterface_listenTcp_Params) String() string {
	str, _ := text.Marshal(0xfd226ae4c6bd2b1e, s.Struct)
	return str
}

func (s IpInterface_listenTcp_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpInterface_listenTcp_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpInterface_listenTcp_Params) Port() TcpPort {
	p, _ := s.Struct.Ptr(0)
	return TcpPort{Client: p.Interface().Client()}
}

func (s IpInterface_listenTcp_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenTcp_Params) SetPort(v TcpPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpInterface_listenTcp_Params_List is a list of IpInterface_listenTcp_Params.
type IpInterface_listenTcp_Params_List struct{ capnp.List }

// NewIpInterface_listenTcp_Params creates a new list of IpInterface_listenTcp_Params.
func NewIpInterface_listenTcp_Params_List(s *capnp.Segment, sz int32) (IpInterface_listenTcp_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return IpInterface_listenTcp_Params_List{l}, err
}

func (s IpInterface_listenTcp_Params_List) At(i int) IpInterface_listenTcp_Params {
	return IpInterface_listenTcp_Params{s.List.Struct(i)}
}

func (s IpInterface_listenTcp_Params_List) Set(i int, v IpInterface_listenTcp_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpInterface_listenTcp_Params_Promise is a wrapper for a IpInterface_listenTcp_Params promised by a client call.
type IpInterface_listenTcp_Params_Promise struct{ *capnp.Pipeline }

func (p IpInterface_listenTcp_Params_Promise) Struct() (IpInterface_listenTcp_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpInterface_listenTcp_Params{s}, err
}

func (p IpInterface_listenTcp_Params_Promise) Port() TcpPort {
	return TcpPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpInterface_listenTcp_Results struct{ capnp.Struct }

func NewIpInterface_listenTcp_Results(s *capnp.Segment) (IpInterface_listenTcp_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpInterface_listenTcp_Results{st}, err
}

func NewRootIpInterface_listenTcp_Results(s *capnp.Segment) (IpInterface_listenTcp_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpInterface_listenTcp_Results{st}, err
}

func ReadRootIpInterface_listenTcp_Results(msg *capnp.Message) (IpInterface_listenTcp_Results, error) {
	root, err := msg.RootPtr()
	return IpInterface_listenTcp_Results{root.Struct()}, err
}

func (s IpInterface_listenTcp_Results) String() string {
	str, _ := text.Marshal(0x9381253786627ecf, s.Struct)
	return str
}

func (s IpInterface_listenTcp_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s IpInterface_listenTcp_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenTcp_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpInterface_listenTcp_Results_List is a list of IpInterface_listenTcp_Results.
type IpInterface_listenTcp_Results_List struct{ capnp.List }

// NewIpInterface_listenTcp_Results creates a new list of IpInterface_listenTcp_Results.
func NewIpInterface_listenTcp_Results_List(s *capnp.Segment, sz int32) (IpInterface_listenTcp_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpInterface_listenTcp_Results_List{l}, err
}

func (s IpInterface_listenTcp_Results_List) At(i int) IpInterface_listenTcp_Results {
	return IpInterface_listenTcp_Results{s.List.Struct(i)}
}

func (s IpInterface_listenTcp_Results_List) Set(i int, v IpInterface_listenTcp_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpInterface_listenTcp_Results_Promise is a wrapper for a IpInterface_listenTcp_Results promised by a client call.
type IpInterface_listenTcp_Results_Promise struct{ *capnp.Pipeline }

func (p IpInterface_listenTcp_Results_Promise) Struct() (IpInterface_listenTcp_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpInterface_listenTcp_Results{s}, err
}

func (p IpInterface_listenTcp_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpInterface_listenUdp_Params struct{ capnp.Struct }

func NewIpInterface_listenUdp_Params(s *capnp.Segment) (IpInterface_listenUdp_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpInterface_listenUdp_Params{st}, err
}

func NewRootIpInterface_listenUdp_Params(s *capnp.Segment) (IpInterface_listenUdp_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpInterface_listenUdp_Params{st}, err
}

func ReadRootIpInterface_listenUdp_Params(msg *capnp.Message) (IpInterface_listenUdp_Params, error) {
	root, err := msg.RootPtr()
	return IpInterface_listenUdp_Params{root.Struct()}, err
}

func (s IpInterface_listenUdp_Params) String() string {
	str, _ := text.Marshal(0xa1d8815a262abc49, s.Struct)
	return str
}

func (s IpInterface_listenUdp_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpInterface_listenUdp_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpInterface_listenUdp_Params) Port() UdpPort {
	p, _ := s.Struct.Ptr(0)
	return UdpPort{Client: p.Interface().Client()}
}

func (s IpInterface_listenUdp_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenUdp_Params) SetPort(v UdpPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpInterface_listenUdp_Params_List is a list of IpInterface_listenUdp_Params.
type IpInterface_listenUdp_Params_List struct{ capnp.List }

// NewIpInterface_listenUdp_Params creates a new list of IpInterface_listenUdp_Params.
func NewIpInterface_listenUdp_Params_List(s *capnp.Segment, sz int32) (IpInterface_listenUdp_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return IpInterface_listenUdp_Params_List{l}, err
}

func (s IpInterface_listenUdp_Params_List) At(i int) IpInterface_listenUdp_Params {
	return IpInterface_listenUdp_Params{s.List.Struct(i)}
}

func (s IpInterface_listenUdp_Params_List) Set(i int, v IpInterface_listenUdp_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpInterface_listenUdp_Params_Promise is a wrapper for a IpInterface_listenUdp_Params promised by a client call.
type IpInterface_listenUdp_Params_Promise struct{ *capnp.Pipeline }

func (p IpInterface_listenUdp_Params_Promise) Struct() (IpInterface_listenUdp_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpInterface_listenUdp_Params{s}, err
}

func (p IpInterface_listenUdp_Params_Promise) Port() UdpPort {
	return UdpPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpInterface_listenUdp_Results struct{ capnp.Struct }

func NewIpInterface_listenUdp_Results(s *capnp.Segment) (IpInterface_listenUdp_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpInterface_listenUdp_Results{st}, err
}

func NewRootIpInterface_listenUdp_Results(s *capnp.Segment) (IpInterface_listenUdp_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpInterface_listenUdp_Results{st}, err
}

func ReadRootIpInterface_listenUdp_Results(msg *capnp.Message) (IpInterface_listenUdp_Results, error) {
	root, err := msg.RootPtr()
	return IpInterface_listenUdp_Results{root.Struct()}, err
}

func (s IpInterface_listenUdp_Results) String() string {
	str, _ := text.Marshal(0xcb83a480981bc290, s.Struct)
	return str
}

func (s IpInterface_listenUdp_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s IpInterface_listenUdp_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenUdp_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpInterface_listenUdp_Results_List is a list of IpInterface_listenUdp_Results.
type IpInterface_listenUdp_Results_List struct{ capnp.List }

// NewIpInterface_listenUdp_Results creates a new list of IpInterface_listenUdp_Results.
func NewIpInterface_listenUdp_Results_List(s *capnp.Segment, sz int32) (IpInterface_listenUdp_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpInterface_listenUdp_Results_List{l}, err
}

func (s IpInterface_listenUdp_Results_List) At(i int) IpInterface_listenUdp_Results {
	return IpInterface_listenUdp_Results{s.List.Struct(i)}
}

func (s IpInterface_listenUdp_Results_List) Set(i int, v IpInterface_listenUdp_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpInterface_listenUdp_Results_Promise is a wrapper for a IpInterface_listenUdp_Results promised by a client call.
type IpInterface_listenUdp_Results_Promise struct{ *capnp.Pipeline }

func (p IpInterface_listenUdp_Results_Promise) Struct() (IpInterface_listenUdp_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpInterface_listenUdp_Results{s}, err
}

func (p IpInterface_listenUdp_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpRemoteHost struct{ Client capnp.Client }

func (c IpRemoteHost) GetTcpPort(ctx context.Context, params func(IpRemoteHost_getTcpPort_Params) error, opts ...capnp.CallOption) IpRemoteHost_getTcpPort_Results_Promise {
	if c.Client == nil {
		return IpRemoteHost_getTcpPort_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x905dd76b298b3130,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpRemoteHost",
			MethodName:    "getTcpPort",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpRemoteHost_getTcpPort_Params{Struct: s}) }
	}
	return IpRemoteHost_getTcpPort_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c IpRemoteHost) GetUdpPort(ctx context.Context, params func(IpRemoteHost_getUdpPort_Params) error, opts ...capnp.CallOption) IpRemoteHost_getUdpPort_Results_Promise {
	if c.Client == nil {
		return IpRemoteHost_getUdpPort_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x905dd76b298b3130,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpRemoteHost",
			MethodName:    "getUdpPort",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpRemoteHost_getUdpPort_Params{Struct: s}) }
	}
	return IpRemoteHost_getUdpPort_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type IpRemoteHost_Server interface {
	GetTcpPort(IpRemoteHost_getTcpPort) error

	GetUdpPort(IpRemoteHost_getUdpPort) error
}

func IpRemoteHost_ServerToClient(s IpRemoteHost_Server) IpRemoteHost {
	c, _ := s.(server.Closer)
	return IpRemoteHost{Client: server.New(IpRemoteHost_Methods(nil, s), c)}
}

func IpRemoteHost_Methods(methods []server.Method, s IpRemoteHost_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x905dd76b298b3130,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpRemoteHost",
			MethodName:    "getTcpPort",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpRemoteHost_getTcpPort{c, opts, IpRemoteHost_getTcpPort_Params{Struct: p}, IpRemoteHost_getTcpPort_Results{Struct: r}}
			return s.GetTcpPort(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x905dd76b298b3130,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpRemoteHost",
			MethodName:    "getUdpPort",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpRemoteHost_getUdpPort{c, opts, IpRemoteHost_getUdpPort_Params{Struct: p}, IpRemoteHost_getUdpPort_Results{Struct: r}}
			return s.GetUdpPort(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// IpRemoteHost_getTcpPort holds the arguments for a server call to IpRemoteHost.getTcpPort.
type IpRemoteHost_getTcpPort struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpRemoteHost_getTcpPort_Params
	Results IpRemoteHost_getTcpPort_Results
}

// IpRemoteHost_getUdpPort holds the arguments for a server call to IpRemoteHost.getUdpPort.
type IpRemoteHost_getUdpPort struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IpRemoteHost_getUdpPort_Params
	Results IpRemoteHost_getUdpPort_Results
}

type IpRemoteHost_getTcpPort_Params struct{ capnp.Struct }

func NewIpRemoteHost_getTcpPort_Params(s *capnp.Segment) (IpRemoteHost_getTcpPort_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpRemoteHost_getTcpPort_Params{st}, err
}

func NewRootIpRemoteHost_getTcpPort_Params(s *capnp.Segment) (IpRemoteHost_getTcpPort_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpRemoteHost_getTcpPort_Params{st}, err
}

func ReadRootIpRemoteHost_getTcpPort_Params(msg *capnp.Message) (IpRemoteHost_getTcpPort_Params, error) {
	root, err := msg.RootPtr()
	return IpRemoteHost_getTcpPort_Params{root.Struct()}, err
}

func (s IpRemoteHost_getTcpPort_Params) String() string {
	str, _ := text.Marshal(0xed10beb11e7383e9, s.Struct)
	return str
}

func (s IpRemoteHost_getTcpPort_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpRemoteHost_getTcpPort_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

// IpRemoteHost_getTcpPort_Params_List is a list of IpRemoteHost_getTcpPort_Params.
type IpRemoteHost_getTcpPort_Params_List struct{ capnp.List }

// NewIpRemoteHost_getTcpPort_Params creates a new list of IpRemoteHost_getTcpPort_Params.
func NewIpRemoteHost_getTcpPort_Params_List(s *capnp.Segment, sz int32) (IpRemoteHost_getTcpPort_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return IpRemoteHost_getTcpPort_Params_List{l}, err
}

func (s IpRemoteHost_getTcpPort_Params_List) At(i int) IpRemoteHost_getTcpPort_Params {
	return IpRemoteHost_getTcpPort_Params{s.List.Struct(i)}
}

func (s IpRemoteHost_getTcpPort_Params_List) Set(i int, v IpRemoteHost_getTcpPort_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpRemoteHost_getTcpPort_Params_Promise is a wrapper for a IpRemoteHost_getTcpPort_Params promised by a client call.
type IpRemoteHost_getTcpPort_Params_Promise struct{ *capnp.Pipeline }

func (p IpRemoteHost_getTcpPort_Params_Promise) Struct() (IpRemoteHost_getTcpPort_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpRemoteHost_getTcpPort_Params{s}, err
}

type IpRemoteHost_getTcpPort_Results struct{ capnp.Struct }

func NewIpRemoteHost_getTcpPort_Results(s *capnp.Segment) (IpRemoteHost_getTcpPort_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpRemoteHost_getTcpPort_Results{st}, err
}

func NewRootIpRemoteHost_getTcpPort_Results(s *capnp.Segment) (IpRemoteHost_getTcpPort_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpRemoteHost_getTcpPort_Results{st}, err
}

func ReadRootIpRemoteHost_getTcpPort_Results(msg *capnp.Message) (IpRemoteHost_getTcpPort_Results, error) {
	root, err := msg.RootPtr()
	return IpRemoteHost_getTcpPort_Results{root.Struct()}, err
}

func (s IpRemoteHost_getTcpPort_Results) String() string {
	str, _ := text.Marshal(0xd77df9f44cfcde33, s.Struct)
	return str
}

func (s IpRemoteHost_getTcpPort_Results) Port() TcpPort {
	p, _ := s.Struct.Ptr(0)
	return TcpPort{Client: p.Interface().Client()}
}

func (s IpRemoteHost_getTcpPort_Results) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpRemoteHost_getTcpPort_Results) SetPort(v TcpPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpRemoteHost_getTcpPort_Results_List is a list of IpRemoteHost_getTcpPort_Results.
type IpRemoteHost_getTcpPort_Results_List struct{ capnp.List }

// NewIpRemoteHost_getTcpPort_Results creates a new list of IpRemoteHost_getTcpPort_Results.
func NewIpRemoteHost_getTcpPort_Results_List(s *capnp.Segment, sz int32) (IpRemoteHost_getTcpPort_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpRemoteHost_getTcpPort_Results_List{l}, err
}

func (s IpRemoteHost_getTcpPort_Results_List) At(i int) IpRemoteHost_getTcpPort_Results {
	return IpRemoteHost_getTcpPort_Results{s.List.Struct(i)}
}

func (s IpRemoteHost_getTcpPort_Results_List) Set(i int, v IpRemoteHost_getTcpPort_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpRemoteHost_getTcpPort_Results_Promise is a wrapper for a IpRemoteHost_getTcpPort_Results promised by a client call.
type IpRemoteHost_getTcpPort_Results_Promise struct{ *capnp.Pipeline }

func (p IpRemoteHost_getTcpPort_Results_Promise) Struct() (IpRemoteHost_getTcpPort_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpRemoteHost_getTcpPort_Results{s}, err
}

func (p IpRemoteHost_getTcpPort_Results_Promise) Port() TcpPort {
	return TcpPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type IpRemoteHost_getUdpPort_Params struct{ capnp.Struct }

func NewIpRemoteHost_getUdpPort_Params(s *capnp.Segment) (IpRemoteHost_getUdpPort_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpRemoteHost_getUdpPort_Params{st}, err
}

func NewRootIpRemoteHost_getUdpPort_Params(s *capnp.Segment) (IpRemoteHost_getUdpPort_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpRemoteHost_getUdpPort_Params{st}, err
}

func ReadRootIpRemoteHost_getUdpPort_Params(msg *capnp.Message) (IpRemoteHost_getUdpPort_Params, error) {
	root, err := msg.RootPtr()
	return IpRemoteHost_getUdpPort_Params{root.Struct()}, err
}

func (s IpRemoteHost_getUdpPort_Params) String() string {
	str, _ := text.Marshal(0xb62b02486ebe26ed, s.Struct)
	return str
}

func (s IpRemoteHost_getUdpPort_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpRemoteHost_getUdpPort_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

// IpRemoteHost_getUdpPort_Params_List is a list of IpRemoteHost_getUdpPort_Params.
type IpRemoteHost_getUdpPort_Params_List struct{ capnp.List }

// NewIpRemoteHost_getUdpPort_Params creates a new list of IpRemoteHost_getUdpPort_Params.
func NewIpRemoteHost_getUdpPort_Params_List(s *capnp.Segment, sz int32) (IpRemoteHost_getUdpPort_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return IpRemoteHost_getUdpPort_Params_List{l}, err
}

func (s IpRemoteHost_getUdpPort_Params_List) At(i int) IpRemoteHost_getUdpPort_Params {
	return IpRemoteHost_getUdpPort_Params{s.List.Struct(i)}
}

func (s IpRemoteHost_getUdpPort_Params_List) Set(i int, v IpRemoteHost_getUdpPort_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpRemoteHost_getUdpPort_Params_Promise is a wrapper for a IpRemoteHost_getUdpPort_Params promised by a client call.
type IpRemoteHost_getUdpPort_Params_Promise struct{ *capnp.Pipeline }

func (p IpRemoteHost_getUdpPort_Params_Promise) Struct() (IpRemoteHost_getUdpPort_Params, error) {
	s, err := p.Pipeline.Struct()
	return IpRemoteHost_getUdpPort_Params{s}, err
}

type IpRemoteHost_getUdpPort_Results struct{ capnp.Struct }

func NewIpRemoteHost_getUdpPort_Results(s *capnp.Segment) (IpRemoteHost_getUdpPort_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpRemoteHost_getUdpPort_Results{st}, err
}

func NewRootIpRemoteHost_getUdpPort_Results(s *capnp.Segment) (IpRemoteHost_getUdpPort_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpRemoteHost_getUdpPort_Results{st}, err
}

func ReadRootIpRemoteHost_getUdpPort_Results(msg *capnp.Message) (IpRemoteHost_getUdpPort_Results, error) {
	root, err := msg.RootPtr()
	return IpRemoteHost_getUdpPort_Results{root.Struct()}, err
}

func (s IpRemoteHost_getUdpPort_Results) String() string {
	str, _ := text.Marshal(0xf53aa3a93e49003b, s.Struct)
	return str
}

func (s IpRemoteHost_getUdpPort_Results) Port() UdpPort {
	p, _ := s.Struct.Ptr(0)
	return UdpPort{Client: p.Interface().Client()}
}

func (s IpRemoteHost_getUdpPort_Results) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpRemoteHost_getUdpPort_Results) SetPort(v UdpPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// IpRemoteHost_getUdpPort_Results_List is a list of IpRemoteHost_getUdpPort_Results.
type IpRemoteHost_getUdpPort_Results_List struct{ capnp.List }

// NewIpRemoteHost_getUdpPort_Results creates a new list of IpRemoteHost_getUdpPort_Results.
func NewIpRemoteHost_getUdpPort_Results_List(s *capnp.Segment, sz int32) (IpRemoteHost_getUdpPort_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpRemoteHost_getUdpPort_Results_List{l}, err
}

func (s IpRemoteHost_getUdpPort_Results_List) At(i int) IpRemoteHost_getUdpPort_Results {
	return IpRemoteHost_getUdpPort_Results{s.List.Struct(i)}
}

func (s IpRemoteHost_getUdpPort_Results_List) Set(i int, v IpRemoteHost_getUdpPort_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpRemoteHost_getUdpPort_Results_Promise is a wrapper for a IpRemoteHost_getUdpPort_Results promised by a client call.
type IpRemoteHost_getUdpPort_Results_Promise struct{ *capnp.Pipeline }

func (p IpRemoteHost_getUdpPort_Results_Promise) Struct() (IpRemoteHost_getUdpPort_Results, error) {
	s, err := p.Pipeline.Struct()
	return IpRemoteHost_getUdpPort_Results{s}, err
}

func (p IpRemoteHost_getUdpPort_Results_Promise) Port() UdpPort {
	return UdpPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type TcpPort struct{ Client capnp.Client }

func (c TcpPort) Connect(ctx context.Context, params func(TcpPort_connect_Params) error, opts ...capnp.CallOption) TcpPort_connect_Results_Promise {
	if c.Client == nil {
		return TcpPort_connect_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeab20e1af07806b4,
			MethodID:      0,
			InterfaceName: "ip.capnp:TcpPort",
			MethodName:    "connect",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(TcpPort_connect_Params{Struct: s}) }
	}
	return TcpPort_connect_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type TcpPort_Server interface {
	Connect(TcpPort_connect) error
}

func TcpPort_ServerToClient(s TcpPort_Server) TcpPort {
	c, _ := s.(server.Closer)
	return TcpPort{Client: server.New(TcpPort_Methods(nil, s), c)}
}

func TcpPort_Methods(methods []server.Method, s TcpPort_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeab20e1af07806b4,
			MethodID:      0,
			InterfaceName: "ip.capnp:TcpPort",
			MethodName:    "connect",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := TcpPort_connect{c, opts, TcpPort_connect_Params{Struct: p}, TcpPort_connect_Results{Struct: r}}
			return s.Connect(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// TcpPort_connect holds the arguments for a server call to TcpPort.connect.
type TcpPort_connect struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  TcpPort_connect_Params
	Results TcpPort_connect_Results
}

type TcpPort_connect_Params struct{ capnp.Struct }

func NewTcpPort_connect_Params(s *capnp.Segment) (TcpPort_connect_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TcpPort_connect_Params{st}, err
}

func NewRootTcpPort_connect_Params(s *capnp.Segment) (TcpPort_connect_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TcpPort_connect_Params{st}, err
}

func ReadRootTcpPort_connect_Params(msg *capnp.Message) (TcpPort_connect_Params, error) {
	root, err := msg.RootPtr()
	return TcpPort_connect_Params{root.Struct()}, err
}

func (s TcpPort_connect_Params) String() string {
	str, _ := text.Marshal(0x8a60e53250a32321, s.Struct)
	return str
}

func (s TcpPort_connect_Params) Downstream() util.ByteStream {
	p, _ := s.Struct.Ptr(0)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s TcpPort_connect_Params) HasDownstream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TcpPort_connect_Params) SetDownstream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// TcpPort_connect_Params_List is a list of TcpPort_connect_Params.
type TcpPort_connect_Params_List struct{ capnp.List }

// NewTcpPort_connect_Params creates a new list of TcpPort_connect_Params.
func NewTcpPort_connect_Params_List(s *capnp.Segment, sz int32) (TcpPort_connect_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TcpPort_connect_Params_List{l}, err
}

func (s TcpPort_connect_Params_List) At(i int) TcpPort_connect_Params {
	return TcpPort_connect_Params{s.List.Struct(i)}
}

func (s TcpPort_connect_Params_List) Set(i int, v TcpPort_connect_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// TcpPort_connect_Params_Promise is a wrapper for a TcpPort_connect_Params promised by a client call.
type TcpPort_connect_Params_Promise struct{ *capnp.Pipeline }

func (p TcpPort_connect_Params_Promise) Struct() (TcpPort_connect_Params, error) {
	s, err := p.Pipeline.Struct()
	return TcpPort_connect_Params{s}, err
}

func (p TcpPort_connect_Params_Promise) Downstream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type TcpPort_connect_Results struct{ capnp.Struct }

func NewTcpPort_connect_Results(s *capnp.Segment) (TcpPort_connect_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TcpPort_connect_Results{st}, err
}

func NewRootTcpPort_connect_Results(s *capnp.Segment) (TcpPort_connect_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TcpPort_connect_Results{st}, err
}

func ReadRootTcpPort_connect_Results(msg *capnp.Message) (TcpPort_connect_Results, error) {
	root, err := msg.RootPtr()
	return TcpPort_connect_Results{root.Struct()}, err
}

func (s TcpPort_connect_Results) String() string {
	str, _ := text.Marshal(0xcdd1222d14073645, s.Struct)
	return str
}

func (s TcpPort_connect_Results) Upstream() util.ByteStream {
	p, _ := s.Struct.Ptr(0)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s TcpPort_connect_Results) HasUpstream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TcpPort_connect_Results) SetUpstream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// TcpPort_connect_Results_List is a list of TcpPort_connect_Results.
type TcpPort_connect_Results_List struct{ capnp.List }

// NewTcpPort_connect_Results creates a new list of TcpPort_connect_Results.
func NewTcpPort_connect_Results_List(s *capnp.Segment, sz int32) (TcpPort_connect_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TcpPort_connect_Results_List{l}, err
}

func (s TcpPort_connect_Results_List) At(i int) TcpPort_connect_Results {
	return TcpPort_connect_Results{s.List.Struct(i)}
}

func (s TcpPort_connect_Results_List) Set(i int, v TcpPort_connect_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// TcpPort_connect_Results_Promise is a wrapper for a TcpPort_connect_Results promised by a client call.
type TcpPort_connect_Results_Promise struct{ *capnp.Pipeline }

func (p TcpPort_connect_Results_Promise) Struct() (TcpPort_connect_Results, error) {
	s, err := p.Pipeline.Struct()
	return TcpPort_connect_Results{s}, err
}

func (p TcpPort_connect_Results_Promise) Upstream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UdpPort struct{ Client capnp.Client }

func (c UdpPort) Send(ctx context.Context, params func(UdpPort_send_Params) error, opts ...capnp.CallOption) UdpPort_send_Results_Promise {
	if c.Client == nil {
		return UdpPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc6212e1217d001ce,
			MethodID:      0,
			InterfaceName: "ip.capnp:UdpPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UdpPort_send_Params{Struct: s}) }
	}
	return UdpPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type UdpPort_Server interface {
	Send(UdpPort_send) error
}

func UdpPort_ServerToClient(s UdpPort_Server) UdpPort {
	c, _ := s.(server.Closer)
	return UdpPort{Client: server.New(UdpPort_Methods(nil, s), c)}
}

func UdpPort_Methods(methods []server.Method, s UdpPort_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc6212e1217d001ce,
			MethodID:      0,
			InterfaceName: "ip.capnp:UdpPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UdpPort_send{c, opts, UdpPort_send_Params{Struct: p}, UdpPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// UdpPort_send holds the arguments for a server call to UdpPort.send.
type UdpPort_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UdpPort_send_Params
	Results UdpPort_send_Results
}

type UdpPort_send_Params struct{ capnp.Struct }

func NewUdpPort_send_Params(s *capnp.Segment) (UdpPort_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return UdpPort_send_Params{st}, err
}

func NewRootUdpPort_send_Params(s *capnp.Segment) (UdpPort_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return UdpPort_send_Params{st}, err
}

func ReadRootUdpPort_send_Params(msg *capnp.Message) (UdpPort_send_Params, error) {
	root, err := msg.RootPtr()
	return UdpPort_send_Params{root.Struct()}, err
}

func (s UdpPort_send_Params) String() string {
	str, _ := text.Marshal(0xc6ca13f7c8dbd102, s.Struct)
	return str
}

func (s UdpPort_send_Params) Message() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s UdpPort_send_Params) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UdpPort_send_Params) SetMessage(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s UdpPort_send_Params) ReturnPort() UdpPort {
	p, _ := s.Struct.Ptr(1)
	return UdpPort{Client: p.Interface().Client()}
}

func (s UdpPort_send_Params) HasReturnPort() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UdpPort_send_Params) SetReturnPort(v UdpPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// UdpPort_send_Params_List is a list of UdpPort_send_Params.
type UdpPort_send_Params_List struct{ capnp.List }

// NewUdpPort_send_Params creates a new list of UdpPort_send_Params.
func NewUdpPort_send_Params_List(s *capnp.Segment, sz int32) (UdpPort_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return UdpPort_send_Params_List{l}, err
}

func (s UdpPort_send_Params_List) At(i int) UdpPort_send_Params {
	return UdpPort_send_Params{s.List.Struct(i)}
}

func (s UdpPort_send_Params_List) Set(i int, v UdpPort_send_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UdpPort_send_Params_Promise is a wrapper for a UdpPort_send_Params promised by a client call.
type UdpPort_send_Params_Promise struct{ *capnp.Pipeline }

func (p UdpPort_send_Params_Promise) Struct() (UdpPort_send_Params, error) {
	s, err := p.Pipeline.Struct()
	return UdpPort_send_Params{s}, err
}

func (p UdpPort_send_Params_Promise) ReturnPort() UdpPort {
	return UdpPort{Client: p.Pipeline.GetPipeline(1).Client()}
}

type UdpPort_send_Results struct{ capnp.Struct }

func NewUdpPort_send_Results(s *capnp.Segment) (UdpPort_send_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return UdpPort_send_Results{st}, err
}

func NewRootUdpPort_send_Results(s *capnp.Segment) (UdpPort_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return UdpPort_send_Results{st}, err
}

func ReadRootUdpPort_send_Results(msg *capnp.Message) (UdpPort_send_Results, error) {
	root, err := msg.RootPtr()
	return UdpPort_send_Results{root.Struct()}, err
}

func (s UdpPort_send_Results) String() string {
	str, _ := text.Marshal(0x8e43fd8e213b1811, s.Struct)
	return str
}

// UdpPort_send_Results_List is a list of UdpPort_send_Results.
type UdpPort_send_Results_List struct{ capnp.List }

// NewUdpPort_send_Results creates a new list of UdpPort_send_Results.
func NewUdpPort_send_Results_List(s *capnp.Segment, sz int32) (UdpPort_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return UdpPort_send_Results_List{l}, err
}

func (s UdpPort_send_Results_List) At(i int) UdpPort_send_Results {
	return UdpPort_send_Results{s.List.Struct(i)}
}

func (s UdpPort_send_Results_List) Set(i int, v UdpPort_send_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UdpPort_send_Results_Promise is a wrapper for a UdpPort_send_Results promised by a client call.
type UdpPort_send_Results_Promise struct{ *capnp.Pipeline }

func (p UdpPort_send_Results_Promise) Struct() (UdpPort_send_Results, error) {
	s, err := p.Pipeline.Struct()
	return UdpPort_send_Results{s}, err
}

type IpPortPowerboxMetadata struct{ capnp.Struct }

func NewIpPortPowerboxMetadata(s *capnp.Segment) (IpPortPowerboxMetadata, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpPortPowerboxMetadata{st}, err
}

func NewRootIpPortPowerboxMetadata(s *capnp.Segment) (IpPortPowerboxMetadata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return IpPortPowerboxMetadata{st}, err
}

func ReadRootIpPortPowerboxMetadata(msg *capnp.Message) (IpPortPowerboxMetadata, error) {
	root, err := msg.RootPtr()
	return IpPortPowerboxMetadata{root.Struct()}, err
}

func (s IpPortPowerboxMetadata) String() string {
	str, _ := text.Marshal(0x856e71a6a4f22bba, s.Struct)
	return str
}

func (s IpPortPowerboxMetadata) PreferredPortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpPortPowerboxMetadata) SetPreferredPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpPortPowerboxMetadata) PreferredHost() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s IpPortPowerboxMetadata) HasPreferredHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpPortPowerboxMetadata) PreferredHostBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s IpPortPowerboxMetadata) SetPreferredHost(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// IpPortPowerboxMetadata_List is a list of IpPortPowerboxMetadata.
type IpPortPowerboxMetadata_List struct{ capnp.List }

// NewIpPortPowerboxMetadata creates a new list of IpPortPowerboxMetadata.
func NewIpPortPowerboxMetadata_List(s *capnp.Segment, sz int32) (IpPortPowerboxMetadata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return IpPortPowerboxMetadata_List{l}, err
}

func (s IpPortPowerboxMetadata_List) At(i int) IpPortPowerboxMetadata {
	return IpPortPowerboxMetadata{s.List.Struct(i)}
}

func (s IpPortPowerboxMetadata_List) Set(i int, v IpPortPowerboxMetadata) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpPortPowerboxMetadata_Promise is a wrapper for a IpPortPowerboxMetadata promised by a client call.
type IpPortPowerboxMetadata_Promise struct{ *capnp.Pipeline }

func (p IpPortPowerboxMetadata_Promise) Struct() (IpPortPowerboxMetadata, error) {
	s, err := p.Pipeline.Struct()
	return IpPortPowerboxMetadata{s}, err
}

type PersistentIpNetwork struct{ Client capnp.Client }

func (c PersistentIpNetwork) GetRemoteHost(ctx context.Context, params func(IpNetwork_getRemoteHost_Params) error, opts ...capnp.CallOption) IpNetwork_getRemoteHost_Results_Promise {
	if c.Client == nil {
		return IpNetwork_getRemoteHost_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHost",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpNetwork_getRemoteHost_Params{Struct: s}) }
	}
	return IpNetwork_getRemoteHost_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpNetwork) GetRemoteHostByName(ctx context.Context, params func(IpNetwork_getRemoteHostByName_Params) error, opts ...capnp.CallOption) IpNetwork_getRemoteHostByName_Results_Promise {
	if c.Client == nil {
		return IpNetwork_getRemoteHostByName_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHostByName",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpNetwork_getRemoteHostByName_Params{Struct: s}) }
	}
	return IpNetwork_getRemoteHostByName_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpNetwork) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(supervisor.SystemPersistent_addRequirements_Params{Struct: s})
		}
	}
	return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpNetwork) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
	if c.Client == nil {
		return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(persistent.Persistent_SaveParams{Struct: s}) }
	}
	return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type PersistentIpNetwork_Server interface {
	GetRemoteHost(IpNetwork_getRemoteHost) error

	GetRemoteHostByName(IpNetwork_getRemoteHostByName) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentIpNetwork_ServerToClient(s PersistentIpNetwork_Server) PersistentIpNetwork {
	c, _ := s.(server.Closer)
	return PersistentIpNetwork{Client: server.New(PersistentIpNetwork_Methods(nil, s), c)}
}

func PersistentIpNetwork_Methods(methods []server.Method, s PersistentIpNetwork_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHost",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpNetwork_getRemoteHost{c, opts, IpNetwork_getRemoteHost_Params{Struct: p}, IpNetwork_getRemoteHost_Results{Struct: r}}
			return s.GetRemoteHost(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa982576b7a2a2040,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpNetwork",
			MethodName:    "getRemoteHostByName",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpNetwork_getRemoteHostByName{c, opts, IpNetwork_getRemoteHostByName_Params{Struct: p}, IpNetwork_getRemoteHostByName_Results{Struct: r}}
			return s.GetRemoteHostByName(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := supervisor.SystemPersistent_addRequirements{c, opts, supervisor.SystemPersistent_addRequirements_Params{Struct: p}, supervisor.SystemPersistent_addRequirements_Results{Struct: r}}
			return s.AddRequirements(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := persistent.Persistent_save{c, opts, persistent.Persistent_SaveParams{Struct: p}, persistent.Persistent_SaveResults{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

type PersistentIpInterface struct{ Client capnp.Client }

func (c PersistentIpInterface) ListenTcp(ctx context.Context, params func(IpInterface_listenTcp_Params) error, opts ...capnp.CallOption) IpInterface_listenTcp_Results_Promise {
	if c.Client == nil {
		return IpInterface_listenTcp_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenTcp",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpInterface_listenTcp_Params{Struct: s}) }
	}
	return IpInterface_listenTcp_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpInterface) ListenUdp(ctx context.Context, params func(IpInterface_listenUdp_Params) error, opts ...capnp.CallOption) IpInterface_listenUdp_Results_Promise {
	if c.Client == nil {
		return IpInterface_listenUdp_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenUdp",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IpInterface_listenUdp_Params{Struct: s}) }
	}
	return IpInterface_listenUdp_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpInterface) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(supervisor.SystemPersistent_addRequirements_Params{Struct: s})
		}
	}
	return supervisor.SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIpInterface) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
	if c.Client == nil {
		return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(persistent.Persistent_SaveParams{Struct: s}) }
	}
	return persistent.Persistent_SaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type PersistentIpInterface_Server interface {
	ListenTcp(IpInterface_listenTcp) error

	ListenUdp(IpInterface_listenUdp) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentIpInterface_ServerToClient(s PersistentIpInterface_Server) PersistentIpInterface {
	c, _ := s.(server.Closer)
	return PersistentIpInterface{Client: server.New(PersistentIpInterface_Methods(nil, s), c)}
}

func PersistentIpInterface_Methods(methods []server.Method, s PersistentIpInterface_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      0,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenTcp",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpInterface_listenTcp{c, opts, IpInterface_listenTcp_Params{Struct: p}, IpInterface_listenTcp_Results{Struct: r}}
			return s.ListenTcp(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe32c506ee93ed6fa,
			MethodID:      1,
			InterfaceName: "ip.capnp:IpInterface",
			MethodName:    "listenUdp",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IpInterface_listenUdp{c, opts, IpInterface_listenUdp_Params{Struct: p}, IpInterface_listenUdp_Results{Struct: r}}
			return s.ListenUdp(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := supervisor.SystemPersistent_addRequirements{c, opts, supervisor.SystemPersistent_addRequirements_Params{Struct: p}, supervisor.SystemPersistent_addRequirements_Results{Struct: r}}
			return s.AddRequirements(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "capnp/persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := persistent.Persistent_save{c, opts, persistent.Persistent_SaveParams{Struct: p}, persistent.Persistent_SaveResults{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

const schema_f44732d435305f86 = "x\xda\xacVkl\x14U\x14\xbewv\x97\xa5\xbb3" +
	"]\xc6iyh\xa0..\xe8\x16\xd8\xb4\xbc-\x09\xad" +
	"K\x08\x14-\xee6*\xb1\x89\xb6CwxHwv" +
	"\x99\x99\xcd\x02\x8a\xb5\x12 TE\x88\x7f\x80\xc4\x1f\xc8" +
	"#\xc1\x18\x13\x10L@%\x01\x95\x88\xa6\"\x11Dc" +
	"P\x12\xc4\x04p\x8d \x0d)\x81\x8c\xe7\xde\xd9yl" +
	"w\xa1U\xfc1\xc9\xcc\xbdg\xce\xf9\xcew\xbes\xee" +
	"\xad9\xe3n`j=OW \x14\xff\xc33D?" +
	"2\xe1\xfa\xee\xbd+\xe5\xf5(^\x86\xb1\xbe\xa1\xb5f" +
	"\xda\x99\xc9\xf3n \x0f\xf6\"4\x85\xf31X\x18\xe5" +
	"\x83W\xa1\xd2\x97EX\x0f>\xb2+6\xf9R[7" +
	"\xe2y\x8c\xf2Fk\xc1\x08a\xa1\xcbW\x0f\x06\xfc\xc8" +
	"Y\xc1\xcdw\xe6l6\x0c\xdc\xe4\xcf=\xbe\xeb\xc8\xad" +
	"\xd7\xd4\xbe\x11^q\xee\x85-\x88/s\xd9a\xe0\xbf" +
	"M\xbe\x03\xc2V\x1a\xe2-\xdfF\xe1\x12y\xd3O\xbd" +
	"\xbax\xc3\x8cq]\xef \xbe\xd2\x0a\xd3\xe3\xab#a" +
	"\xce\xd20_|r\xf3\xb6\\\xd5\xfa.\xe2\x05\xcb\xa0" +
	"\xd7\xa7\x10\x03\xec'\x06\x8d\x9fV\x8fo\xe9\xfaq'" +
	"\x8aWb\xcb\"\xe8\x9fL,\xc2~\x92\xca\xf5\x99\x17" +
	"\xda\x9e\x0b~\xb4\xa7\x08\xd0z\xff\x15a\xab\x9f\x02\xf2" +
	"{\xe1y\x14\x005<\\\xbdf\xc5\xa2\xd7\xf7\x950" +
	"\xdeA\xcc\xc0x\x93\x1f\xd0\x937]|\x92\xfd\xfb\xc3" +
	"\xb3/\x1fr\x82\xeb\xf1/ \xa1\x7f\xa2\xe0r\xe3\x8f" +
	"\xca\xf3\x99\x09\x1f\xa3\xf8p\x9cgiJ\x9f?J\xd1" +
	"\xb3\xc4\xe0[\xfc\xdd\x88\x07\"\xc1\x13E\xe1\x82l\xb7" +
	"\x10fI\xb8q\xec<\xa1\x89\xbc\xe9\xcc\xe9\x9f\xbf\xba" +
	")|}\"_\x13\x86\xecNc\xaf\x80\xf5\xe3,I" +
	"t\xcb\xf1\x87\xb6\xbd\xb6{\xdd7N2\xb7\xb3\x94\xcc" +
	"\x9d4\xda\xdc\xe9\xde\x8aIcO\xf78\x8bz\x8c\xf5" +
	"\x11\x83\x93\xd4\xe0\xe0\xb5\xbd{~\xbf:\xe7T\x11\x9c" +
	"\x1c{K\xe8\xa3pzY/<\x84\xaa\xe5\xadg\x8e" +
	"\x1c\x8f,8\xed\xcc>\xc7\xae!\xde\xfa\xa8\xb7)\xbf" +
	"\xdc~\xeaF\xdf\xdas\x88\x1fn\x19\x8c\xe2(=\xe3" +
	"80\xf8\xf5\xe0\x8b\xca\xd5c#\xce;\xfe\x9f\xcbQ" +
	"r\x9a\xc8\xb6~\xeb\x87\xd9\x97\xe5\xd8\xc4\x8bEh\x92" +
	"\xdc\xfbB\x86#hVr\x1b\x85\x93\xe4M?8d" +
	"\xd5_\x0f\x96\x1f\xb8Rd\xbc\x9f\xeb\x16\x0eS\xe3C" +
	"\xdc<\xe1<5\xbe\xbcN\x1d\xb3\xff\xe8\xb0\x9c\xb3." +
	"_\x1a\xa1{hh\xff\xc4\xefgl\x7f\xfb\xf0\x9f\xa4" +
	"Y\x18\xdb\x1d\x95y\x8e\xdb!\xf4R\x87\xd78\xa0\xfd" +
	"\xce\xac\xc6\xd9\xfbv\xd5\xf5:\x92l*\xa7I>_" +
	"N\\\x8d\x99\xf0\xd9\x89\xdf^\x1a{\xa7@\xa0\xab\xcb" +
	"\xa9@\xbb\xca\xb3HGY}y:\xd2.\xa6\xe5\xb4" +
	"\xab\xae1\x1dK)Z,\x95\x95\x94\xc5\xa9UM\x92" +
	"&&DMD1\x8c\xe3C]n\x80\x0a.\xf8p" +
	"7\xf4\xf4D\x17\x8e\xcfd0\xc6\x15\xc4-?M\x81" +
	"\xb5\xa9\xb0\xd6\xc0`=\xadHK$E\x91p\x82x" +
	"[\x98I\x02\x11^\xc4\xc0c\xef\xa1\xaa\xc4\xfc\x94\xaa" +
	"a\x16\xd6YX\xb71<\xd3NAD\xdaS\xb2," +
	"\xb5k\xa1\x98\xa8\x88I\x15\xc5\xdd\x16\x00\xae\x05\x82\xb1" +
	"\x10l$\x04K\xa4\xb2\xb2\xaa\x81C\x97\x98\xc4\xbc\xde" +
	"z\xf1T8;sQ\x0f\x84\xc4\xbc\xc3/S\xf7l" +
	"\xc2\xf0\xabJr\"\xd4,U\xa9\x99\x0eM\xb5\xf61" +
	"\xe4\xde,%SZ\x95Dp\x19\x19{\x1c\xc5\xc2\xa6" +
	"\x9e\xf8\xda\x16\xc4\xf0a/\xb6\x1b\x0c\xcfB\xf9*\x8c" +
	"&{\x95^}\xa9\xa4\xd1<\x90K\xd1\x1a0\xf9\xa4" +
	"\xe1\x8dOp^\xc0y\xa3\xacI\xca\x12\xb1]\x8at" +
	",W5I\x86?C\xcd\xf5\x12\x05\xe8L\xbb\x0e\xd2" +
	"\x1e\x0aiW0\xb8~\x99(':$H\xf8B\xb4" +
	"\xad\xed\x83\xd0\x8dm\xfd\x13v\x83\xe3\x85\x92\x96M)" +
	"+\"\x10\x9e\xe6FS\x8b\xae^(&\xa5P\xac\x8a" +
	"\xf2\xea\xf4\x1f\xb5\xfdw\x8a\x89\x84\"\xa9j\x89\xfa\x14" +
	"\xe3\x85\xd4,w\x0e\x99\x10w!pWc\xcbdR" +
	"5\xac=\x06kS!D\xda\x10\x87)\x8d\x00\xf9\x86" +
	"\x84\xac\xa1TT\xc1\x98\xa4\xa84\xa0FS\x0b\x90\xdc" +
	"\x80\xcb\x98\xcb\x13\x1f\x0a\x94Z\xc3\x93\xf4\xe3\xd9\xa3\xaf" +
	"\x9c\xcb\xbd\xf99yw\x96\x982\xe2\xa2\xbf\x19\xf55" +
	"\xa7\x80c\x9a\xd6*f}\xcd\xf1\x8f\xcda\xc3\x8f~" +
	"\x0f\xf6F\xd1\xfaRF\x91!\x17\xa3\xc4t\x05\x9b$" +
	"{\x81\xe5\xc2Z\xdf\xb5$ GZlT\xa0\xf2j" +
	"\xbb\x1c\x81e\xa4Wx\xfbt+Ql\xdb]\xc4\x94" +
	"\x9bR\xb2{\x9ce\xeeW\x03\x07Wy\x07\x886\xbf" +
	"\x9b\xb6\x82y\x02`\xf3\xf4\xe5\xf9j\xa0\xc3\xe3\x0d\x90" +
	"\x9e*\xcc\xb5_\xbf\xc5\xc4\xc0=\xe4\xc1[\xfah\xb1" +
	"GKg\x12\xf4'.\x950\x07\xd88\xc0\xa6HZ" +
	"F\x91\x8d&\xba\x87N\xee\xa6\xd0\xfb\xee\xa8\xe2\xd1T" +
	"\xb2l\x0b\x1c\xc3)\x93&\xa3I\x84\x01\x88\xee9\x9b" +
	"\x9c\xca\x06\xec\xf5\x06x\x87\xb6\xad\xc3\xe8.\xda\x1e\xa8" +
	"\xdbK%\x7f_\xfa\xcaS1\xb0r\xf3]m\x1d\x90" +
	"\x83\x1dS\x03)\xd7\x1cP\xc3\xec\xd3\x12\\\x0f+\xd0" +
	"p^\x08\x81<\x99\xf9\x89n\x1e\x89\xd8\xbc\xfe\xf1\xb5" +
	"\xcdf\xc7\x9b\xf79l\xdef\xf8\xd1\xcd\xc6D7G" +
	"3\xc2i\xe8vSV\xf4\xcb\xa9{l\x8a\xc4\xd98" +
	"\xe6u\x16\x9bW \x9e\x8f\xd2\xc6\xe9\xcc+\xa9xN" +
	"\x94f\xfb?wsc\xfa\x09\xc2\x97KU\xfb\x9d\xe5" +
	"%\xbb0\xea\x98\xd2\x1d\xe420}*.\x03\x9fe" +
	"\x08wf\xd2i\xe7\xf7 F\xd0`%\xf2o\x1a\x9a" +
	"\x1c\x91\xff\xdf\x91\xd3O\x9c\xff\x04\x00\x00\xff\xffR\x89" +
	"\x80 "

func init() {
	schemas.Register(schema_f44732d435305f86,
		0x856e71a6a4f22bba,
		0x8a60e53250a32321,
		0x8e43fd8e213b1811,
		0x905dd76b298b3130,
		0x9381253786627ecf,
		0x9d5f1f6efcf7bbc4,
		0xa1d8815a262abc49,
		0xa5b3215660e038f2,
		0xa982576b7a2a2040,
		0xb57bd5aef30c4b61,
		0xb62b02486ebe26ed,
		0xc6212e1217d001ce,
		0xc6ca13f7c8dbd102,
		0xcb83a480981bc290,
		0xcdd1222d14073645,
		0xcf43ebe6a5a6f1b4,
		0xd14a2ec2bad45f69,
		0xd77df9f44cfcde33,
		0xdd1700c1eb725eb4,
		0xe32c506ee93ed6fa,
		0xeab20e1af07806b4,
		0xed10beb11e7383e9,
		0xeeb98f9937d32c0b,
		0xf53aa3a93e49003b,
		0xfd226ae4c6bd2b1e)
}
