package ip

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
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

type IpNetwork_PowerboxTag struct{ capnp.Struct }

// IpNetwork_PowerboxTag_TypeID is the unique identifier for the type IpNetwork_PowerboxTag.
const IpNetwork_PowerboxTag_TypeID = 0xcf9e3f33950df819

func NewIpNetwork_PowerboxTag(s *capnp.Segment) (IpNetwork_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_PowerboxTag{st}, err
}

func NewRootIpNetwork_PowerboxTag(s *capnp.Segment) (IpNetwork_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IpNetwork_PowerboxTag{st}, err
}

func ReadRootIpNetwork_PowerboxTag(msg *capnp.Message) (IpNetwork_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return IpNetwork_PowerboxTag{root.Struct()}, err
}

func (s IpNetwork_PowerboxTag) String() string {
	str, _ := text.Marshal(0xcf9e3f33950df819, s.Struct)
	return str
}

func (s IpNetwork_PowerboxTag) Encryption() (IpNetwork_PowerboxTag_Encryption, error) {
	p, err := s.Struct.Ptr(0)
	return IpNetwork_PowerboxTag_Encryption{Struct: p.Struct()}, err
}

func (s IpNetwork_PowerboxTag) HasEncryption() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IpNetwork_PowerboxTag) SetEncryption(v IpNetwork_PowerboxTag_Encryption) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEncryption sets the encryption field to a newly
// allocated IpNetwork_PowerboxTag_Encryption struct, preferring placement in s's segment.
func (s IpNetwork_PowerboxTag) NewEncryption() (IpNetwork_PowerboxTag_Encryption, error) {
	ss, err := NewIpNetwork_PowerboxTag_Encryption(s.Struct.Segment())
	if err != nil {
		return IpNetwork_PowerboxTag_Encryption{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// IpNetwork_PowerboxTag_List is a list of IpNetwork_PowerboxTag.
type IpNetwork_PowerboxTag_List struct{ capnp.List }

// NewIpNetwork_PowerboxTag creates a new list of IpNetwork_PowerboxTag.
func NewIpNetwork_PowerboxTag_List(s *capnp.Segment, sz int32) (IpNetwork_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IpNetwork_PowerboxTag_List{l}, err
}

func (s IpNetwork_PowerboxTag_List) At(i int) IpNetwork_PowerboxTag {
	return IpNetwork_PowerboxTag{s.List.Struct(i)}
}

func (s IpNetwork_PowerboxTag_List) Set(i int, v IpNetwork_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_PowerboxTag_Promise is a wrapper for a IpNetwork_PowerboxTag promised by a client call.
type IpNetwork_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_PowerboxTag_Promise) Struct() (IpNetwork_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_PowerboxTag{s}, err
}

func (p IpNetwork_PowerboxTag_Promise) Encryption() IpNetwork_PowerboxTag_Encryption_Promise {
	return IpNetwork_PowerboxTag_Encryption_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type IpNetwork_PowerboxTag_Encryption struct{ capnp.Struct }
type IpNetwork_PowerboxTag_Encryption_Which uint16

const (
	IpNetwork_PowerboxTag_Encryption_Which_none IpNetwork_PowerboxTag_Encryption_Which = 0
	IpNetwork_PowerboxTag_Encryption_Which_tls  IpNetwork_PowerboxTag_Encryption_Which = 1
)

func (w IpNetwork_PowerboxTag_Encryption_Which) String() string {
	const s = "nonetls"
	switch w {
	case IpNetwork_PowerboxTag_Encryption_Which_none:
		return s[0:4]
	case IpNetwork_PowerboxTag_Encryption_Which_tls:
		return s[4:7]

	}
	return "IpNetwork_PowerboxTag_Encryption_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// IpNetwork_PowerboxTag_Encryption_TypeID is the unique identifier for the type IpNetwork_PowerboxTag_Encryption.
const IpNetwork_PowerboxTag_Encryption_TypeID = 0xe2d94cf90fe4078d

func NewIpNetwork_PowerboxTag_Encryption(s *capnp.Segment) (IpNetwork_PowerboxTag_Encryption, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpNetwork_PowerboxTag_Encryption{st}, err
}

func NewRootIpNetwork_PowerboxTag_Encryption(s *capnp.Segment) (IpNetwork_PowerboxTag_Encryption, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return IpNetwork_PowerboxTag_Encryption{st}, err
}

func ReadRootIpNetwork_PowerboxTag_Encryption(msg *capnp.Message) (IpNetwork_PowerboxTag_Encryption, error) {
	root, err := msg.RootPtr()
	return IpNetwork_PowerboxTag_Encryption{root.Struct()}, err
}

func (s IpNetwork_PowerboxTag_Encryption) String() string {
	str, _ := text.Marshal(0xe2d94cf90fe4078d, s.Struct)
	return str
}

func (s IpNetwork_PowerboxTag_Encryption) Which() IpNetwork_PowerboxTag_Encryption_Which {
	return IpNetwork_PowerboxTag_Encryption_Which(s.Struct.Uint16(0))
}
func (s IpNetwork_PowerboxTag_Encryption) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s IpNetwork_PowerboxTag_Encryption) SetTls() {
	s.Struct.SetUint16(0, 1)

}

// IpNetwork_PowerboxTag_Encryption_List is a list of IpNetwork_PowerboxTag_Encryption.
type IpNetwork_PowerboxTag_Encryption_List struct{ capnp.List }

// NewIpNetwork_PowerboxTag_Encryption creates a new list of IpNetwork_PowerboxTag_Encryption.
func NewIpNetwork_PowerboxTag_Encryption_List(s *capnp.Segment, sz int32) (IpNetwork_PowerboxTag_Encryption_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return IpNetwork_PowerboxTag_Encryption_List{l}, err
}

func (s IpNetwork_PowerboxTag_Encryption_List) At(i int) IpNetwork_PowerboxTag_Encryption {
	return IpNetwork_PowerboxTag_Encryption{s.List.Struct(i)}
}

func (s IpNetwork_PowerboxTag_Encryption_List) Set(i int, v IpNetwork_PowerboxTag_Encryption) error {
	return s.List.SetStruct(i, v.Struct)
}

// IpNetwork_PowerboxTag_Encryption_Promise is a wrapper for a IpNetwork_PowerboxTag_Encryption promised by a client call.
type IpNetwork_PowerboxTag_Encryption_Promise struct{ *capnp.Pipeline }

func (p IpNetwork_PowerboxTag_Encryption_Promise) Struct() (IpNetwork_PowerboxTag_Encryption, error) {
	s, err := p.Pipeline.Struct()
	return IpNetwork_PowerboxTag_Encryption{s}, err
}

type IpNetwork_getRemoteHost_Params struct{ capnp.Struct }

// IpNetwork_getRemoteHost_Params_TypeID is the unique identifier for the type IpNetwork_getRemoteHost_Params.
const IpNetwork_getRemoteHost_Params_TypeID = 0xdd1700c1eb725eb4

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

// IpNetwork_getRemoteHost_Results_TypeID is the unique identifier for the type IpNetwork_getRemoteHost_Results.
const IpNetwork_getRemoteHost_Results_TypeID = 0xb57bd5aef30c4b61

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

// IpNetwork_getRemoteHostByName_Params_TypeID is the unique identifier for the type IpNetwork_getRemoteHostByName_Params.
const IpNetwork_getRemoteHostByName_Params_TypeID = 0x9d5f1f6efcf7bbc4

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

// IpNetwork_getRemoteHostByName_Results_TypeID is the unique identifier for the type IpNetwork_getRemoteHostByName_Results.
const IpNetwork_getRemoteHostByName_Results_TypeID = 0xd14a2ec2bad45f69

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

// IpAddress_TypeID is the unique identifier for the type IpAddress.
const IpAddress_TypeID = 0xeeb98f9937d32c0b

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

// IpInterface_listenTcp_Params_TypeID is the unique identifier for the type IpInterface_listenTcp_Params.
const IpInterface_listenTcp_Params_TypeID = 0xfd226ae4c6bd2b1e

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

// IpInterface_listenTcp_Results_TypeID is the unique identifier for the type IpInterface_listenTcp_Results.
const IpInterface_listenTcp_Results_TypeID = 0x9381253786627ecf

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

// IpInterface_listenUdp_Params_TypeID is the unique identifier for the type IpInterface_listenUdp_Params.
const IpInterface_listenUdp_Params_TypeID = 0xa1d8815a262abc49

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

// IpInterface_listenUdp_Results_TypeID is the unique identifier for the type IpInterface_listenUdp_Results.
const IpInterface_listenUdp_Results_TypeID = 0xcb83a480981bc290

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

// IpRemoteHost_getTcpPort_Params_TypeID is the unique identifier for the type IpRemoteHost_getTcpPort_Params.
const IpRemoteHost_getTcpPort_Params_TypeID = 0xed10beb11e7383e9

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

// IpRemoteHost_getTcpPort_Results_TypeID is the unique identifier for the type IpRemoteHost_getTcpPort_Results.
const IpRemoteHost_getTcpPort_Results_TypeID = 0xd77df9f44cfcde33

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

// IpRemoteHost_getUdpPort_Params_TypeID is the unique identifier for the type IpRemoteHost_getUdpPort_Params.
const IpRemoteHost_getUdpPort_Params_TypeID = 0xb62b02486ebe26ed

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

// IpRemoteHost_getUdpPort_Results_TypeID is the unique identifier for the type IpRemoteHost_getUdpPort_Results.
const IpRemoteHost_getUdpPort_Results_TypeID = 0xf53aa3a93e49003b

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

// TcpPort_connect_Params_TypeID is the unique identifier for the type TcpPort_connect_Params.
const TcpPort_connect_Params_TypeID = 0x8a60e53250a32321

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

// TcpPort_connect_Results_TypeID is the unique identifier for the type TcpPort_connect_Results.
const TcpPort_connect_Results_TypeID = 0xcdd1222d14073645

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

// UdpPort_send_Params_TypeID is the unique identifier for the type UdpPort_send_Params.
const UdpPort_send_Params_TypeID = 0xc6ca13f7c8dbd102

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

// UdpPort_send_Results_TypeID is the unique identifier for the type UdpPort_send_Results.
const UdpPort_send_Results_TypeID = 0x8e43fd8e213b1811

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

// IpPortPowerboxMetadata_TypeID is the unique identifier for the type IpPortPowerboxMetadata.
const IpPortPowerboxMetadata_TypeID = 0x856e71a6a4f22bba

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

const schema_f44732d435305f86 = "x\xda\xacW{l\x14\xe5\x16?gv\xb7C\xf7\xd9" +
	"\xb9S^\xf7\x06z\xcb-\xdc\xdb\x02\x9b\x96\xf2J\xc9" +
	"\xa5\xa5H`\x11\xc8NS$6\xa9\xed\xb4;\x94\xda" +
	"\xee\xec03\xcd\x02\x8a\xb5\x12 TA\x881A\x12" +
	"L\x90G\x821&\"\x98\x88B\x02*\x11\x0dV\"" +
	"\x15\x89AQ\xa8I!5\x825\x08\xa1\x8e\xf9\xbe\xdd" +
	"\xd9\x99v\xb7\x05\xc4\xff\xe6\xdb9s\x9e\xbf\xdf\xef|" +
	"[\x1ctU0%.#\x17@\xe8we\x19\xc7\xa7" +
	"\xde:ph\xad\xbc\x19\x84lDcK]\xf1\xac\x0b" +
	"3\x16\xf7\x83\x0bY\x80R\xc9\xcd \xbf\xd6\xcd\x02\xf0" +
	"Qw\x1c\xd0\xc8\xff\xcf\xfe\xf0\x8c\x9e\xfaN\xe08\x84" +
	"\xa4Q\xb7\x9bA@\xfe\x92\xbb\x1c\xd0\xe0\xc6\xcd\xcb\xdf" +
	"1\xb0pG\xc2\xc0I\xbe\x1cp\xdf\x02\xa7Q\\\xf2" +
	"ba\xcb\xc5\xda\x9d\xc0e;\xac0\x80\xfc\x0f\xee#" +
	"|/\x0d\xd1\xe3\xde\xca\xff\xdf\xc3\x02\x18]\xcf5l" +
	"\x993\xb9\xe3\x15\xe0F\xa7\xc2L\xf6\x94\x910\xd3=" +
	"$\xcc\xc7\x1f\xdc\xbe'\xe7\xd5\xed\x05\x8eO\x19,\xf7" +
	"\xa8\xc4\xe0Ij\x10\xfa\xb0hJM\xc77\xfb@\x18" +
	"\x8d)\x8b\xf5\x9e\x19\xc4\xa2\xc3CJ\xb95\xf7J\xfd" +
	"\x13\xf9\xef\x1eLK\xe8\xb2\xe7:\xdf\xeb\xa1\x09yX" +
	"\xbe\xc7\xf3_\x00\xa3\xe2\xdfE\x1bZV\xbdp8\x83" +
	"\xf1\x1e\xbe\xc73\x16\x80\xbf\xe9\xd9\xca/\xf7\x92\xec\xc5" +
	"\xc7\xbd\xbf\xbe\xdd\xfd\xcc1{r\xb3\xbcKI\xe8\x05" +
	"^\x92\\\xdf\x94\x93\xf2\x12f\xea{ \x8c\xc1d\x97" +
	"JEo%1h\xa6\x06_\xe0\x97c\xff\x11\xcc?" +
	"\x93\x16n\x9b\xb7\x93\xdfE\x82\xf0\xdb\xbd\x8b\xf9c4" +
	"\x1cs\xfe\xdbOo\xf3\x9f\x9dI\xce\x84!o\xf7z" +
	"\xaf\x03\xf2\xfb\xbc\xa4\xd0\x9d\xa7\xff\xb5\xfb\xf9\x03\x9b>" +
	"\xb77\x13}\xb4\x99\xd9>\x12m\xd1l6w\xfa\xa4" +
	"\xf3\xe7\xecC-\xf4\xb9\x89A\x0958z\xf3\xd0\xc1" +
	"\x9fn,\xecJKg\xa5\xef./\xfaH\xc0Z\x1f" +
	"\xcb\xd7\xfaH\xab\xc6\xff\xee{\xb5\xb4\xfc\xf5.R\xbd" +
	"\xd57\xea5iO\xba\x15\xa5\x8e\x9b\xeb.\x1c?\x1d" +
	"\\z\xde\xde\xa9]\xbe\x0d$\xf2^jP\xfa\xdd\xbd" +
	"e\xfdw6^\x04nL\xca\xe0\x84\x8f\xb6\xf2,1" +
	"\xf8\xfe\xe8S\xea\x8dSc/\xdb\xbe\xef\xf5\xd1F\xde" +
	"\xa4\xdfog\xaf\x05\xee,\xbb\xf4#\x08y\x88Vj" +
	"\x8f\xb1\x0c@)\xe7\xafB~\xb2\x9f$\x96\xef'\x9d" +
	"\xba\xfb\xf5\xfc^9<\xedjZ\x9d\x1b\xfdo\xf2\x9b" +
	"\xa9a\x87\x7f+\x7f\x89<\x19G\xb3\xd6\xfd\xf2O\xff" +
	"\x91\xebi\xc6\xa7\xfc\x9d\xfcYj\xfc\x89\x7f1\xdfG" +
	"\x8d{7i\x13\xdf9\x99\xd3g\x9fx\xb7\x9f&z" +
	"\xd9O\x12\xf5L\xfbj\xcek/\xbf\xff3\xa1!c" +
	"\xb9K\x10\xc8\xbf\x87w\x05\xc8\x13\x06\xe2\x80\x03\xf3B" +
	"\xf3\x0f\xef/\xfb\xcd\xd6\x92\xda\x00mIs\x80\xb8\x9a" +
	"8\xf5\xc4\x99kOO\x1a\x18\x04\xfdm\x01\x0a\xfd]" +
	"\x818\x18\x107\x9a\x95`\xa3\xa8\xc8\x8a\xa3,\xa4\x84" +
	"c\xaa\x1e\x8e\xc5%\xb5!\xb6n\xb9\xa4\x8b\x11Q\x17" +
	"!\x8c(\x8cr8\x01\x9c\x08\xc0\x15v\x02\x08\xd3\x1c" +
	"(\xcce\x101\x97\xb8\xe5f\xa9\x00\xc2L\x07\x0a\x15" +
	"\x0c\x1a\x8a*\xad\x96TU\xc2\x08\xf1\xb6\xa2-\x0a\x80" +
	",0\xc8\x82\xf5\x0e\xf2\"Kb\x9a\x8e^`\xd0K" +
	"\x86\x9f\xca\xa1\xba\x91&\x11l\x8c\xc9\xb2\xd4\xa8\x17\x84" +
	"EU\x8cj 8S\x09\xf8j\x00\x04\xaf\x03\x85q" +
	"\x0c\x1a\x91X\\\xd6tU\x02\x87\x18E\xce\xa8\xbb\xda" +
	"U\x18\x9f\xbb\xea\x1c\x00\"g\xf3\xcb\x94\xad\x8c$\xfc" +
	"j\x92\x1c)\xa8\x92\xf2\xb4\xb6V]K\xbd\xc7\xb2\x90" +
	"R%Ecz\x9eD\xf2JT\xec\xb2\x0d\x0bM\xf4" +
	"q%5\xc0p\x85,Z\xd4\xc5y\x90\x9c\xc2\x04\xf2" +
	"n4k4I:\xad\x03\x1c\xaa^\x81\xe4H\xc3'" +
	"\x8ea\xc4A=\x0f\xc9\xba\xa4\xae\x16\x1b\xa5`k\xb3" +
	"\xa6Kru\xa3RPU.\xd1\x04\xede\x97\x01\x08" +
	"\xa3\x1c(\xe42X\xbeF\x94#\xad\x12r\xc6\x95\xca" +
	"\xfa\xfa\xb7\x0a\xfaw\x0f-\xd8Y\x16RVHz<" +
	"\xa6\xb6\x04\x9b$\x9d\xd6FK\xab\\\xbfB\x8cJ\x05" +
	"\xe1<\xdaW\xbb\xffJ\xcb\x7f\xbb\x18\x89\xa8\x92\xa6e" +
	"\x98Oz\xbe+#J\xca\x9d\x0d&\xc4]\x81\x03\x85" +
	"b\x0b&\xd3\x8b\x00\x84\xff9P\x98\xc9`\xbb\x92\x00" +
	"\x87\x09\x8d\x009#g\xc9]\xda\x04\xc3\x92\xaa\xd1\x80" +
	":--@j\x0b#\x86\x1d.a\x14\xda\xe4\x85\xf0" +
	"\xb1\xfb\xe4\xb3\x17\xfb^\xfa\x88<\xdbGL;\xe2P" +
	"[\x04\xa7]\x018l0L\xd0\x03[-6%f" +
	"o\xea\x89M\xc3KTs\xf6\xe6\xd2AS\xb6\xb8\x09" +
	"o\x00\xc3\x8d\xa7\xb3\xa7\xdd\x86\x04\x94\x12\xe3\xa7\xbf\xa0" +
	"9\x00V\x8cJ\x83q0\xec\xb8\x0a\xaa\x12@\x80A" +
	"\x0c(\xb2F\x15XCx\xc4Y;5\x03\x10,w" +
	"A\x13\x8ajFf\xd9!0d>\xb6>&\x1d\x00" +
	"\x15\x06'\xa5\x89\xb9w\xd0\xdc\xf9\x1cW\x04\x0c\xe7b" +
	"\x03\x84o\x83k\x1d\xc2\xc5\xb0\x18\x18\x01:\\\x0a;" +
	"5\x96\xec\xb4G%M\x13\x9b$\xf4\x01\x83>@C" +
	"\x95\xf46UN\x10l\x04\x0c\x0d\x87\xdeGf[\xba" +
	"le\x1c\xdbR\x9bp\xb5)D\xb6\xc4(\x00\x8c\xa8" +
	"[v\xd4\x87d\xbd<\x91\xbc\x0d\xf7\xa9E5\x0c\xee" +
	"\x19\x1b\xb4(\xca\xcb\x1bb\xeb\x08\xc8\x09\x07R;\x91" +
	"\xc3\x1ac\x91\xdc\xa8\xaeW\xf4fp\xc4\xe4\xe1\xf4V" +
	"\xb2\xd9`\x8e\xf5= \xe6<\x84\xfcd\xea\xf8#\x81" +
	":\xd9\xff\xfb\xd3%)3\xa9\x8d\xfd\xa0\xbay?\xba" +
	"\x98\x8a\x99c\xad\xef\x11zb\xcaM\xb5\xd8\x144\xdb" +
	"\x1eC9\xb9i\xbd\x86\x91 B\x91E\x04\x1f\xfea" +
	"$\x990\xc9R\xd1\x80\x1c\x93%\xc8b\xf5V\x0d\xb2" +
	"\x06I\x1d\xc5y \x89\x95\xe423o\x03h\xde\xa9" +
	"\xb9\x92*S\xd0\xccK2\x9aWDnBUb\x99" +
	"\x99[\x09P\xa9@\xc3d\x0d=\xd9i\x8d&\x07\xec" +
	"\xba`\xfeG@\xf3^\xc9q\x95T\x17\xda\x93DI" +
	"\x97\xc1\xccs\xfd\xcbb\x15R\x16\x90\xc984m\xc8" +
	"5&\xa3\xc8T\xda\x16T+\x99\xd1\xec\x99\x98\x0d\x0c" +
	"f\x03\xb6\xb7)\x8a\xfd\xfc\x00\x0a\xfb\xa0`|\x18\xbd" +
	"\"\xb7\x83\xbfo\xdb\x0e\xa1\xc1\x9f\x01\x00\x00\xff\xff\x0e" +
	"\x9d\xe44"

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
		0xcf9e3f33950df819,
		0xd14a2ec2bad45f69,
		0xd77df9f44cfcde33,
		0xdd1700c1eb725eb4,
		0xe2d94cf90fe4078d,
		0xe32c506ee93ed6fa,
		0xeab20e1af07806b4,
		0xed10beb11e7383e9,
		0xeeb98f9937d32c0b,
		0xf53aa3a93e49003b,
		0xfd226ae4c6bd2b1e)
}
