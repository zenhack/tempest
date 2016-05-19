package ip

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	persistent "zombiezen.com/go/capnproto2/capnp/persistent"
	server "zombiezen.com/go/capnproto2/server"
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
	if err != nil {
		return IpNetwork_getRemoteHost_Params{}, err
	}
	return IpNetwork_getRemoteHost_Params{st}, nil
}

func NewRootIpNetwork_getRemoteHost_Params(s *capnp.Segment) (IpNetwork_getRemoteHost_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpNetwork_getRemoteHost_Params{}, err
	}
	return IpNetwork_getRemoteHost_Params{st}, nil
}

func ReadRootIpNetwork_getRemoteHost_Params(msg *capnp.Message) (IpNetwork_getRemoteHost_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpNetwork_getRemoteHost_Params{}, err
	}
	return IpNetwork_getRemoteHost_Params{root.Struct()}, nil
}
func (s IpNetwork_getRemoteHost_Params) Address() (IpAddress, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return IpAddress{}, err
	}
	return IpAddress{Struct: p.Struct()}, nil
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
	if err != nil {
		return IpNetwork_getRemoteHost_Params_List{}, err
	}
	return IpNetwork_getRemoteHost_Params_List{l}, nil
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
	if err != nil {
		return IpNetwork_getRemoteHost_Results{}, err
	}
	return IpNetwork_getRemoteHost_Results{st}, nil
}

func NewRootIpNetwork_getRemoteHost_Results(s *capnp.Segment) (IpNetwork_getRemoteHost_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpNetwork_getRemoteHost_Results{}, err
	}
	return IpNetwork_getRemoteHost_Results{st}, nil
}

func ReadRootIpNetwork_getRemoteHost_Results(msg *capnp.Message) (IpNetwork_getRemoteHost_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpNetwork_getRemoteHost_Results{}, err
	}
	return IpNetwork_getRemoteHost_Results{root.Struct()}, nil
}
func (s IpNetwork_getRemoteHost_Results) Host() IpRemoteHost {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return IpRemoteHost{}
	}
	return IpRemoteHost{Client: p.Interface().Client()}
}

func (s IpNetwork_getRemoteHost_Results) HasHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHost_Results) SetHost(v IpRemoteHost) error {
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

// IpNetwork_getRemoteHost_Results_List is a list of IpNetwork_getRemoteHost_Results.
type IpNetwork_getRemoteHost_Results_List struct{ capnp.List }

// NewIpNetwork_getRemoteHost_Results creates a new list of IpNetwork_getRemoteHost_Results.
func NewIpNetwork_getRemoteHost_Results_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHost_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpNetwork_getRemoteHost_Results_List{}, err
	}
	return IpNetwork_getRemoteHost_Results_List{l}, nil
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
	if err != nil {
		return IpNetwork_getRemoteHostByName_Params{}, err
	}
	return IpNetwork_getRemoteHostByName_Params{st}, nil
}

func NewRootIpNetwork_getRemoteHostByName_Params(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpNetwork_getRemoteHostByName_Params{}, err
	}
	return IpNetwork_getRemoteHostByName_Params{st}, nil
}

func ReadRootIpNetwork_getRemoteHostByName_Params(msg *capnp.Message) (IpNetwork_getRemoteHostByName_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpNetwork_getRemoteHostByName_Params{}, err
	}
	return IpNetwork_getRemoteHostByName_Params{root.Struct()}, nil
}
func (s IpNetwork_getRemoteHostByName_Params) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s IpNetwork_getRemoteHostByName_Params) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHostByName_Params) AddressBytes() ([]byte, error) {
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
	if err != nil {
		return IpNetwork_getRemoteHostByName_Params_List{}, err
	}
	return IpNetwork_getRemoteHostByName_Params_List{l}, nil
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
	if err != nil {
		return IpNetwork_getRemoteHostByName_Results{}, err
	}
	return IpNetwork_getRemoteHostByName_Results{st}, nil
}

func NewRootIpNetwork_getRemoteHostByName_Results(s *capnp.Segment) (IpNetwork_getRemoteHostByName_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpNetwork_getRemoteHostByName_Results{}, err
	}
	return IpNetwork_getRemoteHostByName_Results{st}, nil
}

func ReadRootIpNetwork_getRemoteHostByName_Results(msg *capnp.Message) (IpNetwork_getRemoteHostByName_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpNetwork_getRemoteHostByName_Results{}, err
	}
	return IpNetwork_getRemoteHostByName_Results{root.Struct()}, nil
}
func (s IpNetwork_getRemoteHostByName_Results) Host() IpRemoteHost {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return IpRemoteHost{}
	}
	return IpRemoteHost{Client: p.Interface().Client()}
}

func (s IpNetwork_getRemoteHostByName_Results) HasHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_getRemoteHostByName_Results) SetHost(v IpRemoteHost) error {
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

// IpNetwork_getRemoteHostByName_Results_List is a list of IpNetwork_getRemoteHostByName_Results.
type IpNetwork_getRemoteHostByName_Results_List struct{ capnp.List }

// NewIpNetwork_getRemoteHostByName_Results creates a new list of IpNetwork_getRemoteHostByName_Results.
func NewIpNetwork_getRemoteHostByName_Results_List(s *capnp.Segment, sz int32) (IpNetwork_getRemoteHostByName_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpNetwork_getRemoteHostByName_Results_List{}, err
	}
	return IpNetwork_getRemoteHostByName_Results_List{l}, nil
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
	if err != nil {
		return IpAddress{}, err
	}
	return IpAddress{st}, nil
}

func NewRootIpAddress(s *capnp.Segment) (IpAddress, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return IpAddress{}, err
	}
	return IpAddress{st}, nil
}

func ReadRootIpAddress(msg *capnp.Message) (IpAddress, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpAddress{}, err
	}
	return IpAddress{root.Struct()}, nil
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
	if err != nil {
		return IpAddress_List{}, err
	}
	return IpAddress_List{l}, nil
}

func (s IpAddress_List) At(i int) IpAddress           { return IpAddress{s.List.Struct(i)} }
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
	if err != nil {
		return IpInterface_listenTcp_Params{}, err
	}
	return IpInterface_listenTcp_Params{st}, nil
}

func NewRootIpInterface_listenTcp_Params(s *capnp.Segment) (IpInterface_listenTcp_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return IpInterface_listenTcp_Params{}, err
	}
	return IpInterface_listenTcp_Params{st}, nil
}

func ReadRootIpInterface_listenTcp_Params(msg *capnp.Message) (IpInterface_listenTcp_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpInterface_listenTcp_Params{}, err
	}
	return IpInterface_listenTcp_Params{root.Struct()}, nil
}
func (s IpInterface_listenTcp_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpInterface_listenTcp_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpInterface_listenTcp_Params) Port() TcpPort {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return TcpPort{}
	}
	return TcpPort{Client: p.Interface().Client()}
}

func (s IpInterface_listenTcp_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenTcp_Params) SetPort(v TcpPort) error {
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

// IpInterface_listenTcp_Params_List is a list of IpInterface_listenTcp_Params.
type IpInterface_listenTcp_Params_List struct{ capnp.List }

// NewIpInterface_listenTcp_Params creates a new list of IpInterface_listenTcp_Params.
func NewIpInterface_listenTcp_Params_List(s *capnp.Segment, sz int32) (IpInterface_listenTcp_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return IpInterface_listenTcp_Params_List{}, err
	}
	return IpInterface_listenTcp_Params_List{l}, nil
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
	if err != nil {
		return IpInterface_listenTcp_Results{}, err
	}
	return IpInterface_listenTcp_Results{st}, nil
}

func NewRootIpInterface_listenTcp_Results(s *capnp.Segment) (IpInterface_listenTcp_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpInterface_listenTcp_Results{}, err
	}
	return IpInterface_listenTcp_Results{st}, nil
}

func ReadRootIpInterface_listenTcp_Results(msg *capnp.Message) (IpInterface_listenTcp_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpInterface_listenTcp_Results{}, err
	}
	return IpInterface_listenTcp_Results{root.Struct()}, nil
}
func (s IpInterface_listenTcp_Results) Handle() util.Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s IpInterface_listenTcp_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenTcp_Results) SetHandle(v util.Handle) error {
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

// IpInterface_listenTcp_Results_List is a list of IpInterface_listenTcp_Results.
type IpInterface_listenTcp_Results_List struct{ capnp.List }

// NewIpInterface_listenTcp_Results creates a new list of IpInterface_listenTcp_Results.
func NewIpInterface_listenTcp_Results_List(s *capnp.Segment, sz int32) (IpInterface_listenTcp_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpInterface_listenTcp_Results_List{}, err
	}
	return IpInterface_listenTcp_Results_List{l}, nil
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
	if err != nil {
		return IpInterface_listenUdp_Params{}, err
	}
	return IpInterface_listenUdp_Params{st}, nil
}

func NewRootIpInterface_listenUdp_Params(s *capnp.Segment) (IpInterface_listenUdp_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return IpInterface_listenUdp_Params{}, err
	}
	return IpInterface_listenUdp_Params{st}, nil
}

func ReadRootIpInterface_listenUdp_Params(msg *capnp.Message) (IpInterface_listenUdp_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpInterface_listenUdp_Params{}, err
	}
	return IpInterface_listenUdp_Params{root.Struct()}, nil
}
func (s IpInterface_listenUdp_Params) PortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpInterface_listenUdp_Params) SetPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpInterface_listenUdp_Params) Port() UdpPort {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UdpPort{}
	}
	return UdpPort{Client: p.Interface().Client()}
}

func (s IpInterface_listenUdp_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenUdp_Params) SetPort(v UdpPort) error {
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

// IpInterface_listenUdp_Params_List is a list of IpInterface_listenUdp_Params.
type IpInterface_listenUdp_Params_List struct{ capnp.List }

// NewIpInterface_listenUdp_Params creates a new list of IpInterface_listenUdp_Params.
func NewIpInterface_listenUdp_Params_List(s *capnp.Segment, sz int32) (IpInterface_listenUdp_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return IpInterface_listenUdp_Params_List{}, err
	}
	return IpInterface_listenUdp_Params_List{l}, nil
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
	if err != nil {
		return IpInterface_listenUdp_Results{}, err
	}
	return IpInterface_listenUdp_Results{st}, nil
}

func NewRootIpInterface_listenUdp_Results(s *capnp.Segment) (IpInterface_listenUdp_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpInterface_listenUdp_Results{}, err
	}
	return IpInterface_listenUdp_Results{st}, nil
}

func ReadRootIpInterface_listenUdp_Results(msg *capnp.Message) (IpInterface_listenUdp_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpInterface_listenUdp_Results{}, err
	}
	return IpInterface_listenUdp_Results{root.Struct()}, nil
}
func (s IpInterface_listenUdp_Results) Handle() util.Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s IpInterface_listenUdp_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpInterface_listenUdp_Results) SetHandle(v util.Handle) error {
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

// IpInterface_listenUdp_Results_List is a list of IpInterface_listenUdp_Results.
type IpInterface_listenUdp_Results_List struct{ capnp.List }

// NewIpInterface_listenUdp_Results creates a new list of IpInterface_listenUdp_Results.
func NewIpInterface_listenUdp_Results_List(s *capnp.Segment, sz int32) (IpInterface_listenUdp_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpInterface_listenUdp_Results_List{}, err
	}
	return IpInterface_listenUdp_Results_List{l}, nil
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
	if err != nil {
		return IpRemoteHost_getTcpPort_Params{}, err
	}
	return IpRemoteHost_getTcpPort_Params{st}, nil
}

func NewRootIpRemoteHost_getTcpPort_Params(s *capnp.Segment) (IpRemoteHost_getTcpPort_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return IpRemoteHost_getTcpPort_Params{}, err
	}
	return IpRemoteHost_getTcpPort_Params{st}, nil
}

func ReadRootIpRemoteHost_getTcpPort_Params(msg *capnp.Message) (IpRemoteHost_getTcpPort_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpRemoteHost_getTcpPort_Params{}, err
	}
	return IpRemoteHost_getTcpPort_Params{root.Struct()}, nil
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
	if err != nil {
		return IpRemoteHost_getTcpPort_Params_List{}, err
	}
	return IpRemoteHost_getTcpPort_Params_List{l}, nil
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
	if err != nil {
		return IpRemoteHost_getTcpPort_Results{}, err
	}
	return IpRemoteHost_getTcpPort_Results{st}, nil
}

func NewRootIpRemoteHost_getTcpPort_Results(s *capnp.Segment) (IpRemoteHost_getTcpPort_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpRemoteHost_getTcpPort_Results{}, err
	}
	return IpRemoteHost_getTcpPort_Results{st}, nil
}

func ReadRootIpRemoteHost_getTcpPort_Results(msg *capnp.Message) (IpRemoteHost_getTcpPort_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpRemoteHost_getTcpPort_Results{}, err
	}
	return IpRemoteHost_getTcpPort_Results{root.Struct()}, nil
}
func (s IpRemoteHost_getTcpPort_Results) Port() TcpPort {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return TcpPort{}
	}
	return TcpPort{Client: p.Interface().Client()}
}

func (s IpRemoteHost_getTcpPort_Results) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpRemoteHost_getTcpPort_Results) SetPort(v TcpPort) error {
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

// IpRemoteHost_getTcpPort_Results_List is a list of IpRemoteHost_getTcpPort_Results.
type IpRemoteHost_getTcpPort_Results_List struct{ capnp.List }

// NewIpRemoteHost_getTcpPort_Results creates a new list of IpRemoteHost_getTcpPort_Results.
func NewIpRemoteHost_getTcpPort_Results_List(s *capnp.Segment, sz int32) (IpRemoteHost_getTcpPort_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpRemoteHost_getTcpPort_Results_List{}, err
	}
	return IpRemoteHost_getTcpPort_Results_List{l}, nil
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
	if err != nil {
		return IpRemoteHost_getUdpPort_Params{}, err
	}
	return IpRemoteHost_getUdpPort_Params{st}, nil
}

func NewRootIpRemoteHost_getUdpPort_Params(s *capnp.Segment) (IpRemoteHost_getUdpPort_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return IpRemoteHost_getUdpPort_Params{}, err
	}
	return IpRemoteHost_getUdpPort_Params{st}, nil
}

func ReadRootIpRemoteHost_getUdpPort_Params(msg *capnp.Message) (IpRemoteHost_getUdpPort_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpRemoteHost_getUdpPort_Params{}, err
	}
	return IpRemoteHost_getUdpPort_Params{root.Struct()}, nil
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
	if err != nil {
		return IpRemoteHost_getUdpPort_Params_List{}, err
	}
	return IpRemoteHost_getUdpPort_Params_List{l}, nil
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
	if err != nil {
		return IpRemoteHost_getUdpPort_Results{}, err
	}
	return IpRemoteHost_getUdpPort_Results{st}, nil
}

func NewRootIpRemoteHost_getUdpPort_Results(s *capnp.Segment) (IpRemoteHost_getUdpPort_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return IpRemoteHost_getUdpPort_Results{}, err
	}
	return IpRemoteHost_getUdpPort_Results{st}, nil
}

func ReadRootIpRemoteHost_getUdpPort_Results(msg *capnp.Message) (IpRemoteHost_getUdpPort_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpRemoteHost_getUdpPort_Results{}, err
	}
	return IpRemoteHost_getUdpPort_Results{root.Struct()}, nil
}
func (s IpRemoteHost_getUdpPort_Results) Port() UdpPort {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UdpPort{}
	}
	return UdpPort{Client: p.Interface().Client()}
}

func (s IpRemoteHost_getUdpPort_Results) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpRemoteHost_getUdpPort_Results) SetPort(v UdpPort) error {
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

// IpRemoteHost_getUdpPort_Results_List is a list of IpRemoteHost_getUdpPort_Results.
type IpRemoteHost_getUdpPort_Results_List struct{ capnp.List }

// NewIpRemoteHost_getUdpPort_Results creates a new list of IpRemoteHost_getUdpPort_Results.
func NewIpRemoteHost_getUdpPort_Results_List(s *capnp.Segment, sz int32) (IpRemoteHost_getUdpPort_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return IpRemoteHost_getUdpPort_Results_List{}, err
	}
	return IpRemoteHost_getUdpPort_Results_List{l}, nil
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
	if err != nil {
		return TcpPort_connect_Params{}, err
	}
	return TcpPort_connect_Params{st}, nil
}

func NewRootTcpPort_connect_Params(s *capnp.Segment) (TcpPort_connect_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return TcpPort_connect_Params{}, err
	}
	return TcpPort_connect_Params{st}, nil
}

func ReadRootTcpPort_connect_Params(msg *capnp.Message) (TcpPort_connect_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return TcpPort_connect_Params{}, err
	}
	return TcpPort_connect_Params{root.Struct()}, nil
}
func (s TcpPort_connect_Params) Downstream() util.ByteStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s TcpPort_connect_Params) HasDownstream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TcpPort_connect_Params) SetDownstream(v util.ByteStream) error {
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

// TcpPort_connect_Params_List is a list of TcpPort_connect_Params.
type TcpPort_connect_Params_List struct{ capnp.List }

// NewTcpPort_connect_Params creates a new list of TcpPort_connect_Params.
func NewTcpPort_connect_Params_List(s *capnp.Segment, sz int32) (TcpPort_connect_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return TcpPort_connect_Params_List{}, err
	}
	return TcpPort_connect_Params_List{l}, nil
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
	if err != nil {
		return TcpPort_connect_Results{}, err
	}
	return TcpPort_connect_Results{st}, nil
}

func NewRootTcpPort_connect_Results(s *capnp.Segment) (TcpPort_connect_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return TcpPort_connect_Results{}, err
	}
	return TcpPort_connect_Results{st}, nil
}

func ReadRootTcpPort_connect_Results(msg *capnp.Message) (TcpPort_connect_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return TcpPort_connect_Results{}, err
	}
	return TcpPort_connect_Results{root.Struct()}, nil
}
func (s TcpPort_connect_Results) Upstream() util.ByteStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s TcpPort_connect_Results) HasUpstream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TcpPort_connect_Results) SetUpstream(v util.ByteStream) error {
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

// TcpPort_connect_Results_List is a list of TcpPort_connect_Results.
type TcpPort_connect_Results_List struct{ capnp.List }

// NewTcpPort_connect_Results creates a new list of TcpPort_connect_Results.
func NewTcpPort_connect_Results_List(s *capnp.Segment, sz int32) (TcpPort_connect_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return TcpPort_connect_Results_List{}, err
	}
	return TcpPort_connect_Results_List{l}, nil
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
	if err != nil {
		return UdpPort_send_Params{}, err
	}
	return UdpPort_send_Params{st}, nil
}

func NewRootUdpPort_send_Params(s *capnp.Segment) (UdpPort_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return UdpPort_send_Params{}, err
	}
	return UdpPort_send_Params{st}, nil
}

func ReadRootUdpPort_send_Params(msg *capnp.Message) (UdpPort_send_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UdpPort_send_Params{}, err
	}
	return UdpPort_send_Params{root.Struct()}, nil
}
func (s UdpPort_send_Params) Message() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return UdpPort{}
	}
	return UdpPort{Client: p.Interface().Client()}
}

func (s UdpPort_send_Params) HasReturnPort() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UdpPort_send_Params) SetReturnPort(v UdpPort) error {
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

// UdpPort_send_Params_List is a list of UdpPort_send_Params.
type UdpPort_send_Params_List struct{ capnp.List }

// NewUdpPort_send_Params creates a new list of UdpPort_send_Params.
func NewUdpPort_send_Params_List(s *capnp.Segment, sz int32) (UdpPort_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return UdpPort_send_Params_List{}, err
	}
	return UdpPort_send_Params_List{l}, nil
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
	if err != nil {
		return UdpPort_send_Results{}, err
	}
	return UdpPort_send_Results{st}, nil
}

func NewRootUdpPort_send_Results(s *capnp.Segment) (UdpPort_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return UdpPort_send_Results{}, err
	}
	return UdpPort_send_Results{st}, nil
}

func ReadRootUdpPort_send_Results(msg *capnp.Message) (UdpPort_send_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UdpPort_send_Results{}, err
	}
	return UdpPort_send_Results{root.Struct()}, nil
}

// UdpPort_send_Results_List is a list of UdpPort_send_Results.
type UdpPort_send_Results_List struct{ capnp.List }

// NewUdpPort_send_Results creates a new list of UdpPort_send_Results.
func NewUdpPort_send_Results_List(s *capnp.Segment, sz int32) (UdpPort_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return UdpPort_send_Results_List{}, err
	}
	return UdpPort_send_Results_List{l}, nil
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
	if err != nil {
		return IpPortPowerboxMetadata{}, err
	}
	return IpPortPowerboxMetadata{st}, nil
}

func NewRootIpPortPowerboxMetadata(s *capnp.Segment) (IpPortPowerboxMetadata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return IpPortPowerboxMetadata{}, err
	}
	return IpPortPowerboxMetadata{st}, nil
}

func ReadRootIpPortPowerboxMetadata(msg *capnp.Message) (IpPortPowerboxMetadata, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return IpPortPowerboxMetadata{}, err
	}
	return IpPortPowerboxMetadata{root.Struct()}, nil
}
func (s IpPortPowerboxMetadata) PreferredPortNum() uint16 {
	return s.Struct.Uint16(0)
}

func (s IpPortPowerboxMetadata) SetPreferredPortNum(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s IpPortPowerboxMetadata) PreferredHost() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s IpPortPowerboxMetadata) HasPreferredHost() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpPortPowerboxMetadata) PreferredHostBytes() ([]byte, error) {
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
	if err != nil {
		return IpPortPowerboxMetadata_List{}, err
	}
	return IpPortPowerboxMetadata_List{l}, nil
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
