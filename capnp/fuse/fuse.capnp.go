package fuse

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type Node struct{ Client capnp.Client }

func (c Node) Lookup(ctx context.Context, params func(Node_lookup_Params) error, opts ...capnp.CallOption) Node_lookup_Results_Promise {
	if c.Client == nil {
		return Node_lookup_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      0,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "lookup",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Node_lookup_Params{Struct: s}) }
	}
	return Node_lookup_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Node) GetAttributes(ctx context.Context, params func(Node_getAttributes_Params) error, opts ...capnp.CallOption) Node_getAttributes_Results_Promise {
	if c.Client == nil {
		return Node_getAttributes_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      1,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "getAttributes",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Node_getAttributes_Params{Struct: s}) }
	}
	return Node_getAttributes_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Node) OpenAsFile(ctx context.Context, params func(Node_openAsFile_Params) error, opts ...capnp.CallOption) Node_openAsFile_Results_Promise {
	if c.Client == nil {
		return Node_openAsFile_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      2,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "openAsFile",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Node_openAsFile_Params{Struct: s}) }
	}
	return Node_openAsFile_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Node) OpenAsDirectory(ctx context.Context, params func(Node_openAsDirectory_Params) error, opts ...capnp.CallOption) Node_openAsDirectory_Results_Promise {
	if c.Client == nil {
		return Node_openAsDirectory_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      3,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "openAsDirectory",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Node_openAsDirectory_Params{Struct: s}) }
	}
	return Node_openAsDirectory_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Node) Readlink(ctx context.Context, params func(Node_readlink_Params) error, opts ...capnp.CallOption) Node_readlink_Results_Promise {
	if c.Client == nil {
		return Node_readlink_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      4,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "readlink",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Node_readlink_Params{Struct: s}) }
	}
	return Node_readlink_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Node_Server interface {
	Lookup(Node_lookup) error

	GetAttributes(Node_getAttributes) error

	OpenAsFile(Node_openAsFile) error

	OpenAsDirectory(Node_openAsDirectory) error

	Readlink(Node_readlink) error
}

func Node_ServerToClient(s Node_Server) Node {
	c, _ := s.(server.Closer)
	return Node{Client: server.New(Node_Methods(nil, s), c)}
}

func Node_Methods(methods []server.Method, s Node_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 5)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      0,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "lookup",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Node_lookup{c, opts, Node_lookup_Params{Struct: p}, Node_lookup_Results{Struct: r}}
			return s.Lookup(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      1,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "getAttributes",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Node_getAttributes{c, opts, Node_getAttributes_Params{Struct: p}, Node_getAttributes_Results{Struct: r}}
			return s.GetAttributes(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      2,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "openAsFile",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Node_openAsFile{c, opts, Node_openAsFile_Params{Struct: p}, Node_openAsFile_Results{Struct: r}}
			return s.OpenAsFile(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      3,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "openAsDirectory",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Node_openAsDirectory{c, opts, Node_openAsDirectory_Params{Struct: p}, Node_openAsDirectory_Results{Struct: r}}
			return s.OpenAsDirectory(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9abfd6533326d4e4,
			MethodID:      4,
			InterfaceName: "fuse.capnp:Node",
			MethodName:    "readlink",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Node_readlink{c, opts, Node_readlink_Params{Struct: p}, Node_readlink_Results{Struct: r}}
			return s.Readlink(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Node_lookup holds the arguments for a server call to Node.lookup.
type Node_lookup struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Node_lookup_Params
	Results Node_lookup_Results
}

// Node_getAttributes holds the arguments for a server call to Node.getAttributes.
type Node_getAttributes struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Node_getAttributes_Params
	Results Node_getAttributes_Results
}

// Node_openAsFile holds the arguments for a server call to Node.openAsFile.
type Node_openAsFile struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Node_openAsFile_Params
	Results Node_openAsFile_Results
}

// Node_openAsDirectory holds the arguments for a server call to Node.openAsDirectory.
type Node_openAsDirectory struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Node_openAsDirectory_Params
	Results Node_openAsDirectory_Results
}

// Node_readlink holds the arguments for a server call to Node.readlink.
type Node_readlink struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Node_readlink_Params
	Results Node_readlink_Results
}

type Node_Type uint16

// Values of Node_Type.
const (
	Node_Type_unknown         Node_Type = 0
	Node_Type_blockDevice     Node_Type = 1
	Node_Type_characterDevice Node_Type = 2
	Node_Type_directory       Node_Type = 3
	Node_Type_fifo            Node_Type = 4
	Node_Type_symlink         Node_Type = 5
	Node_Type_regular         Node_Type = 6
	Node_Type_socket          Node_Type = 7
)

// String returns the enum's constant name.
func (c Node_Type) String() string {
	switch c {
	case Node_Type_unknown:
		return "unknown"
	case Node_Type_blockDevice:
		return "blockDevice"
	case Node_Type_characterDevice:
		return "characterDevice"
	case Node_Type_directory:
		return "directory"
	case Node_Type_fifo:
		return "fifo"
	case Node_Type_symlink:
		return "symlink"
	case Node_Type_regular:
		return "regular"
	case Node_Type_socket:
		return "socket"

	default:
		return ""
	}
}

// Node_TypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Node_TypeFromString(c string) Node_Type {
	switch c {
	case "unknown":
		return Node_Type_unknown
	case "blockDevice":
		return Node_Type_blockDevice
	case "characterDevice":
		return Node_Type_characterDevice
	case "directory":
		return Node_Type_directory
	case "fifo":
		return Node_Type_fifo
	case "symlink":
		return Node_Type_symlink
	case "regular":
		return Node_Type_regular
	case "socket":
		return Node_Type_socket

	default:
		return 0
	}
}

type Node_Type_List struct{ capnp.List }

func NewNode_Type_List(s *capnp.Segment, sz int32) (Node_Type_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Node_Type_List{}, err
	}
	return Node_Type_List{l.List}, nil
}

func (l Node_Type_List) At(i int) Node_Type {
	ul := capnp.UInt16List{List: l.List}
	return Node_Type(ul.At(i))
}

func (l Node_Type_List) Set(i int, v Node_Type) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Node_Attributes struct{ capnp.Struct }

func NewNode_Attributes(s *capnp.Segment) (Node_Attributes, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 80, PointerCount: 0})
	if err != nil {
		return Node_Attributes{}, err
	}
	return Node_Attributes{st}, nil
}

func NewRootNode_Attributes(s *capnp.Segment) (Node_Attributes, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 80, PointerCount: 0})
	if err != nil {
		return Node_Attributes{}, err
	}
	return Node_Attributes{st}, nil
}

func ReadRootNode_Attributes(msg *capnp.Message) (Node_Attributes, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_Attributes{}, err
	}
	return Node_Attributes{root.Struct()}, nil
}
func (s Node_Attributes) InodeNumber() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node_Attributes) SetInodeNumber(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node_Attributes) Type() Node_Type {
	return Node_Type(s.Struct.Uint16(8))
}

func (s Node_Attributes) SetType(v Node_Type) {
	s.Struct.SetUint16(8, uint16(v))
}

func (s Node_Attributes) Permissions() uint32 {
	return s.Struct.Uint32(12)
}

func (s Node_Attributes) SetPermissions(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s Node_Attributes) LinkCount() uint32 {
	return s.Struct.Uint32(16)
}

func (s Node_Attributes) SetLinkCount(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s Node_Attributes) OwnerId() uint32 {
	return s.Struct.Uint32(20)
}

func (s Node_Attributes) SetOwnerId(v uint32) {
	s.Struct.SetUint32(20, v)
}

func (s Node_Attributes) GroupId() uint32 {
	return s.Struct.Uint32(24)
}

func (s Node_Attributes) SetGroupId(v uint32) {
	s.Struct.SetUint32(24, v)
}

func (s Node_Attributes) DeviceMajor() uint32 {
	return s.Struct.Uint32(28)
}

func (s Node_Attributes) SetDeviceMajor(v uint32) {
	s.Struct.SetUint32(28, v)
}

func (s Node_Attributes) DeviceMinor() uint32 {
	return s.Struct.Uint32(32)
}

func (s Node_Attributes) SetDeviceMinor(v uint32) {
	s.Struct.SetUint32(32, v)
}

func (s Node_Attributes) Size() uint64 {
	return s.Struct.Uint64(40)
}

func (s Node_Attributes) SetSize(v uint64) {
	s.Struct.SetUint64(40, v)
}

func (s Node_Attributes) BlockCount() uint64 {
	return s.Struct.Uint64(48)
}

func (s Node_Attributes) SetBlockCount(v uint64) {
	s.Struct.SetUint64(48, v)
}

func (s Node_Attributes) BlockSize() uint32 {
	return s.Struct.Uint32(36)
}

func (s Node_Attributes) SetBlockSize(v uint32) {
	s.Struct.SetUint32(36, v)
}

func (s Node_Attributes) LastAccessTime() int64 {
	return int64(s.Struct.Uint64(56))
}

func (s Node_Attributes) SetLastAccessTime(v int64) {
	s.Struct.SetUint64(56, uint64(v))
}

func (s Node_Attributes) LastModificationTime() int64 {
	return int64(s.Struct.Uint64(64))
}

func (s Node_Attributes) SetLastModificationTime(v int64) {
	s.Struct.SetUint64(64, uint64(v))
}

func (s Node_Attributes) LastStatusChangeTime() int64 {
	return int64(s.Struct.Uint64(72))
}

func (s Node_Attributes) SetLastStatusChangeTime(v int64) {
	s.Struct.SetUint64(72, uint64(v))
}

// Node_Attributes_List is a list of Node_Attributes.
type Node_Attributes_List struct{ capnp.List }

// NewNode_Attributes creates a new list of Node_Attributes.
func NewNode_Attributes_List(s *capnp.Segment, sz int32) (Node_Attributes_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 80, PointerCount: 0}, sz)
	if err != nil {
		return Node_Attributes_List{}, err
	}
	return Node_Attributes_List{l}, nil
}

func (s Node_Attributes_List) At(i int) Node_Attributes { return Node_Attributes{s.List.Struct(i)} }
func (s Node_Attributes_List) Set(i int, v Node_Attributes) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_Attributes_Promise is a wrapper for a Node_Attributes promised by a client call.
type Node_Attributes_Promise struct{ *capnp.Pipeline }

func (p Node_Attributes_Promise) Struct() (Node_Attributes, error) {
	s, err := p.Pipeline.Struct()
	return Node_Attributes{s}, err
}

type Node_lookup_Params struct{ capnp.Struct }

func NewNode_lookup_Params(s *capnp.Segment) (Node_lookup_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_lookup_Params{}, err
	}
	return Node_lookup_Params{st}, nil
}

func NewRootNode_lookup_Params(s *capnp.Segment) (Node_lookup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_lookup_Params{}, err
	}
	return Node_lookup_Params{st}, nil
}

func ReadRootNode_lookup_Params(msg *capnp.Message) (Node_lookup_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_lookup_Params{}, err
	}
	return Node_lookup_Params{root.Struct()}, nil
}
func (s Node_lookup_Params) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Node_lookup_Params) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_lookup_Params) NameBytes() ([]byte, error) {
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

func (s Node_lookup_Params) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Node_lookup_Params_List is a list of Node_lookup_Params.
type Node_lookup_Params_List struct{ capnp.List }

// NewNode_lookup_Params creates a new list of Node_lookup_Params.
func NewNode_lookup_Params_List(s *capnp.Segment, sz int32) (Node_lookup_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Node_lookup_Params_List{}, err
	}
	return Node_lookup_Params_List{l}, nil
}

func (s Node_lookup_Params_List) At(i int) Node_lookup_Params {
	return Node_lookup_Params{s.List.Struct(i)}
}
func (s Node_lookup_Params_List) Set(i int, v Node_lookup_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_lookup_Params_Promise is a wrapper for a Node_lookup_Params promised by a client call.
type Node_lookup_Params_Promise struct{ *capnp.Pipeline }

func (p Node_lookup_Params_Promise) Struct() (Node_lookup_Params, error) {
	s, err := p.Pipeline.Struct()
	return Node_lookup_Params{s}, err
}

type Node_lookup_Results struct{ capnp.Struct }

func NewNode_lookup_Results(s *capnp.Segment) (Node_lookup_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Node_lookup_Results{}, err
	}
	return Node_lookup_Results{st}, nil
}

func NewRootNode_lookup_Results(s *capnp.Segment) (Node_lookup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Node_lookup_Results{}, err
	}
	return Node_lookup_Results{st}, nil
}

func ReadRootNode_lookup_Results(msg *capnp.Message) (Node_lookup_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_lookup_Results{}, err
	}
	return Node_lookup_Results{root.Struct()}, nil
}
func (s Node_lookup_Results) Node() Node {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Node{}
	}
	return Node{Client: p.Interface().Client()}
}

func (s Node_lookup_Results) HasNode() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_lookup_Results) SetNode(v Node) error {
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

func (s Node_lookup_Results) Ttl() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node_lookup_Results) SetTtl(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Node_lookup_Results_List is a list of Node_lookup_Results.
type Node_lookup_Results_List struct{ capnp.List }

// NewNode_lookup_Results creates a new list of Node_lookup_Results.
func NewNode_lookup_Results_List(s *capnp.Segment, sz int32) (Node_lookup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Node_lookup_Results_List{}, err
	}
	return Node_lookup_Results_List{l}, nil
}

func (s Node_lookup_Results_List) At(i int) Node_lookup_Results {
	return Node_lookup_Results{s.List.Struct(i)}
}
func (s Node_lookup_Results_List) Set(i int, v Node_lookup_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_lookup_Results_Promise is a wrapper for a Node_lookup_Results promised by a client call.
type Node_lookup_Results_Promise struct{ *capnp.Pipeline }

func (p Node_lookup_Results_Promise) Struct() (Node_lookup_Results, error) {
	s, err := p.Pipeline.Struct()
	return Node_lookup_Results{s}, err
}

func (p Node_lookup_Results_Promise) Node() Node {
	return Node{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Node_getAttributes_Params struct{ capnp.Struct }

func NewNode_getAttributes_Params(s *capnp.Segment) (Node_getAttributes_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_getAttributes_Params{}, err
	}
	return Node_getAttributes_Params{st}, nil
}

func NewRootNode_getAttributes_Params(s *capnp.Segment) (Node_getAttributes_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_getAttributes_Params{}, err
	}
	return Node_getAttributes_Params{st}, nil
}

func ReadRootNode_getAttributes_Params(msg *capnp.Message) (Node_getAttributes_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_getAttributes_Params{}, err
	}
	return Node_getAttributes_Params{root.Struct()}, nil
}

// Node_getAttributes_Params_List is a list of Node_getAttributes_Params.
type Node_getAttributes_Params_List struct{ capnp.List }

// NewNode_getAttributes_Params creates a new list of Node_getAttributes_Params.
func NewNode_getAttributes_Params_List(s *capnp.Segment, sz int32) (Node_getAttributes_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Node_getAttributes_Params_List{}, err
	}
	return Node_getAttributes_Params_List{l}, nil
}

func (s Node_getAttributes_Params_List) At(i int) Node_getAttributes_Params {
	return Node_getAttributes_Params{s.List.Struct(i)}
}
func (s Node_getAttributes_Params_List) Set(i int, v Node_getAttributes_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_getAttributes_Params_Promise is a wrapper for a Node_getAttributes_Params promised by a client call.
type Node_getAttributes_Params_Promise struct{ *capnp.Pipeline }

func (p Node_getAttributes_Params_Promise) Struct() (Node_getAttributes_Params, error) {
	s, err := p.Pipeline.Struct()
	return Node_getAttributes_Params{s}, err
}

type Node_getAttributes_Results struct{ capnp.Struct }

func NewNode_getAttributes_Results(s *capnp.Segment) (Node_getAttributes_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Node_getAttributes_Results{}, err
	}
	return Node_getAttributes_Results{st}, nil
}

func NewRootNode_getAttributes_Results(s *capnp.Segment) (Node_getAttributes_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Node_getAttributes_Results{}, err
	}
	return Node_getAttributes_Results{st}, nil
}

func ReadRootNode_getAttributes_Results(msg *capnp.Message) (Node_getAttributes_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_getAttributes_Results{}, err
	}
	return Node_getAttributes_Results{root.Struct()}, nil
}
func (s Node_getAttributes_Results) Attributes() (Node_Attributes, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Node_Attributes{}, err
	}
	return Node_Attributes{Struct: p.Struct()}, nil
}

func (s Node_getAttributes_Results) HasAttributes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_getAttributes_Results) SetAttributes(v Node_Attributes) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAttributes sets the attributes field to a newly
// allocated Node_Attributes struct, preferring placement in s's segment.
func (s Node_getAttributes_Results) NewAttributes() (Node_Attributes, error) {
	ss, err := NewNode_Attributes(s.Struct.Segment())
	if err != nil {
		return Node_Attributes{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Node_getAttributes_Results) Ttl() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node_getAttributes_Results) SetTtl(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Node_getAttributes_Results_List is a list of Node_getAttributes_Results.
type Node_getAttributes_Results_List struct{ capnp.List }

// NewNode_getAttributes_Results creates a new list of Node_getAttributes_Results.
func NewNode_getAttributes_Results_List(s *capnp.Segment, sz int32) (Node_getAttributes_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Node_getAttributes_Results_List{}, err
	}
	return Node_getAttributes_Results_List{l}, nil
}

func (s Node_getAttributes_Results_List) At(i int) Node_getAttributes_Results {
	return Node_getAttributes_Results{s.List.Struct(i)}
}
func (s Node_getAttributes_Results_List) Set(i int, v Node_getAttributes_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_getAttributes_Results_Promise is a wrapper for a Node_getAttributes_Results promised by a client call.
type Node_getAttributes_Results_Promise struct{ *capnp.Pipeline }

func (p Node_getAttributes_Results_Promise) Struct() (Node_getAttributes_Results, error) {
	s, err := p.Pipeline.Struct()
	return Node_getAttributes_Results{s}, err
}

func (p Node_getAttributes_Results_Promise) Attributes() Node_Attributes_Promise {
	return Node_Attributes_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Node_openAsFile_Params struct{ capnp.Struct }

func NewNode_openAsFile_Params(s *capnp.Segment) (Node_openAsFile_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_openAsFile_Params{}, err
	}
	return Node_openAsFile_Params{st}, nil
}

func NewRootNode_openAsFile_Params(s *capnp.Segment) (Node_openAsFile_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_openAsFile_Params{}, err
	}
	return Node_openAsFile_Params{st}, nil
}

func ReadRootNode_openAsFile_Params(msg *capnp.Message) (Node_openAsFile_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_openAsFile_Params{}, err
	}
	return Node_openAsFile_Params{root.Struct()}, nil
}

// Node_openAsFile_Params_List is a list of Node_openAsFile_Params.
type Node_openAsFile_Params_List struct{ capnp.List }

// NewNode_openAsFile_Params creates a new list of Node_openAsFile_Params.
func NewNode_openAsFile_Params_List(s *capnp.Segment, sz int32) (Node_openAsFile_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Node_openAsFile_Params_List{}, err
	}
	return Node_openAsFile_Params_List{l}, nil
}

func (s Node_openAsFile_Params_List) At(i int) Node_openAsFile_Params {
	return Node_openAsFile_Params{s.List.Struct(i)}
}
func (s Node_openAsFile_Params_List) Set(i int, v Node_openAsFile_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_openAsFile_Params_Promise is a wrapper for a Node_openAsFile_Params promised by a client call.
type Node_openAsFile_Params_Promise struct{ *capnp.Pipeline }

func (p Node_openAsFile_Params_Promise) Struct() (Node_openAsFile_Params, error) {
	s, err := p.Pipeline.Struct()
	return Node_openAsFile_Params{s}, err
}

type Node_openAsFile_Results struct{ capnp.Struct }

func NewNode_openAsFile_Results(s *capnp.Segment) (Node_openAsFile_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_openAsFile_Results{}, err
	}
	return Node_openAsFile_Results{st}, nil
}

func NewRootNode_openAsFile_Results(s *capnp.Segment) (Node_openAsFile_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_openAsFile_Results{}, err
	}
	return Node_openAsFile_Results{st}, nil
}

func ReadRootNode_openAsFile_Results(msg *capnp.Message) (Node_openAsFile_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_openAsFile_Results{}, err
	}
	return Node_openAsFile_Results{root.Struct()}, nil
}
func (s Node_openAsFile_Results) File() File {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return File{}
	}
	return File{Client: p.Interface().Client()}
}

func (s Node_openAsFile_Results) HasFile() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_openAsFile_Results) SetFile(v File) error {
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

// Node_openAsFile_Results_List is a list of Node_openAsFile_Results.
type Node_openAsFile_Results_List struct{ capnp.List }

// NewNode_openAsFile_Results creates a new list of Node_openAsFile_Results.
func NewNode_openAsFile_Results_List(s *capnp.Segment, sz int32) (Node_openAsFile_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Node_openAsFile_Results_List{}, err
	}
	return Node_openAsFile_Results_List{l}, nil
}

func (s Node_openAsFile_Results_List) At(i int) Node_openAsFile_Results {
	return Node_openAsFile_Results{s.List.Struct(i)}
}
func (s Node_openAsFile_Results_List) Set(i int, v Node_openAsFile_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_openAsFile_Results_Promise is a wrapper for a Node_openAsFile_Results promised by a client call.
type Node_openAsFile_Results_Promise struct{ *capnp.Pipeline }

func (p Node_openAsFile_Results_Promise) Struct() (Node_openAsFile_Results, error) {
	s, err := p.Pipeline.Struct()
	return Node_openAsFile_Results{s}, err
}

func (p Node_openAsFile_Results_Promise) File() File {
	return File{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Node_openAsDirectory_Params struct{ capnp.Struct }

func NewNode_openAsDirectory_Params(s *capnp.Segment) (Node_openAsDirectory_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_openAsDirectory_Params{}, err
	}
	return Node_openAsDirectory_Params{st}, nil
}

func NewRootNode_openAsDirectory_Params(s *capnp.Segment) (Node_openAsDirectory_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_openAsDirectory_Params{}, err
	}
	return Node_openAsDirectory_Params{st}, nil
}

func ReadRootNode_openAsDirectory_Params(msg *capnp.Message) (Node_openAsDirectory_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_openAsDirectory_Params{}, err
	}
	return Node_openAsDirectory_Params{root.Struct()}, nil
}

// Node_openAsDirectory_Params_List is a list of Node_openAsDirectory_Params.
type Node_openAsDirectory_Params_List struct{ capnp.List }

// NewNode_openAsDirectory_Params creates a new list of Node_openAsDirectory_Params.
func NewNode_openAsDirectory_Params_List(s *capnp.Segment, sz int32) (Node_openAsDirectory_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Node_openAsDirectory_Params_List{}, err
	}
	return Node_openAsDirectory_Params_List{l}, nil
}

func (s Node_openAsDirectory_Params_List) At(i int) Node_openAsDirectory_Params {
	return Node_openAsDirectory_Params{s.List.Struct(i)}
}
func (s Node_openAsDirectory_Params_List) Set(i int, v Node_openAsDirectory_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_openAsDirectory_Params_Promise is a wrapper for a Node_openAsDirectory_Params promised by a client call.
type Node_openAsDirectory_Params_Promise struct{ *capnp.Pipeline }

func (p Node_openAsDirectory_Params_Promise) Struct() (Node_openAsDirectory_Params, error) {
	s, err := p.Pipeline.Struct()
	return Node_openAsDirectory_Params{s}, err
}

type Node_openAsDirectory_Results struct{ capnp.Struct }

func NewNode_openAsDirectory_Results(s *capnp.Segment) (Node_openAsDirectory_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_openAsDirectory_Results{}, err
	}
	return Node_openAsDirectory_Results{st}, nil
}

func NewRootNode_openAsDirectory_Results(s *capnp.Segment) (Node_openAsDirectory_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_openAsDirectory_Results{}, err
	}
	return Node_openAsDirectory_Results{st}, nil
}

func ReadRootNode_openAsDirectory_Results(msg *capnp.Message) (Node_openAsDirectory_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_openAsDirectory_Results{}, err
	}
	return Node_openAsDirectory_Results{root.Struct()}, nil
}
func (s Node_openAsDirectory_Results) Directory() Directory {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Directory{}
	}
	return Directory{Client: p.Interface().Client()}
}

func (s Node_openAsDirectory_Results) HasDirectory() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_openAsDirectory_Results) SetDirectory(v Directory) error {
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

// Node_openAsDirectory_Results_List is a list of Node_openAsDirectory_Results.
type Node_openAsDirectory_Results_List struct{ capnp.List }

// NewNode_openAsDirectory_Results creates a new list of Node_openAsDirectory_Results.
func NewNode_openAsDirectory_Results_List(s *capnp.Segment, sz int32) (Node_openAsDirectory_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Node_openAsDirectory_Results_List{}, err
	}
	return Node_openAsDirectory_Results_List{l}, nil
}

func (s Node_openAsDirectory_Results_List) At(i int) Node_openAsDirectory_Results {
	return Node_openAsDirectory_Results{s.List.Struct(i)}
}
func (s Node_openAsDirectory_Results_List) Set(i int, v Node_openAsDirectory_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_openAsDirectory_Results_Promise is a wrapper for a Node_openAsDirectory_Results promised by a client call.
type Node_openAsDirectory_Results_Promise struct{ *capnp.Pipeline }

func (p Node_openAsDirectory_Results_Promise) Struct() (Node_openAsDirectory_Results, error) {
	s, err := p.Pipeline.Struct()
	return Node_openAsDirectory_Results{s}, err
}

func (p Node_openAsDirectory_Results_Promise) Directory() Directory {
	return Directory{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Node_readlink_Params struct{ capnp.Struct }

func NewNode_readlink_Params(s *capnp.Segment) (Node_readlink_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_readlink_Params{}, err
	}
	return Node_readlink_Params{st}, nil
}

func NewRootNode_readlink_Params(s *capnp.Segment) (Node_readlink_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Node_readlink_Params{}, err
	}
	return Node_readlink_Params{st}, nil
}

func ReadRootNode_readlink_Params(msg *capnp.Message) (Node_readlink_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_readlink_Params{}, err
	}
	return Node_readlink_Params{root.Struct()}, nil
}

// Node_readlink_Params_List is a list of Node_readlink_Params.
type Node_readlink_Params_List struct{ capnp.List }

// NewNode_readlink_Params creates a new list of Node_readlink_Params.
func NewNode_readlink_Params_List(s *capnp.Segment, sz int32) (Node_readlink_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Node_readlink_Params_List{}, err
	}
	return Node_readlink_Params_List{l}, nil
}

func (s Node_readlink_Params_List) At(i int) Node_readlink_Params {
	return Node_readlink_Params{s.List.Struct(i)}
}
func (s Node_readlink_Params_List) Set(i int, v Node_readlink_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_readlink_Params_Promise is a wrapper for a Node_readlink_Params promised by a client call.
type Node_readlink_Params_Promise struct{ *capnp.Pipeline }

func (p Node_readlink_Params_Promise) Struct() (Node_readlink_Params, error) {
	s, err := p.Pipeline.Struct()
	return Node_readlink_Params{s}, err
}

type Node_readlink_Results struct{ capnp.Struct }

func NewNode_readlink_Results(s *capnp.Segment) (Node_readlink_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_readlink_Results{}, err
	}
	return Node_readlink_Results{st}, nil
}

func NewRootNode_readlink_Results(s *capnp.Segment) (Node_readlink_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Node_readlink_Results{}, err
	}
	return Node_readlink_Results{st}, nil
}

func ReadRootNode_readlink_Results(msg *capnp.Message) (Node_readlink_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Node_readlink_Results{}, err
	}
	return Node_readlink_Results{root.Struct()}, nil
}
func (s Node_readlink_Results) Link() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Node_readlink_Results) HasLink() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node_readlink_Results) LinkBytes() ([]byte, error) {
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

func (s Node_readlink_Results) SetLink(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Node_readlink_Results_List is a list of Node_readlink_Results.
type Node_readlink_Results_List struct{ capnp.List }

// NewNode_readlink_Results creates a new list of Node_readlink_Results.
func NewNode_readlink_Results_List(s *capnp.Segment, sz int32) (Node_readlink_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Node_readlink_Results_List{}, err
	}
	return Node_readlink_Results_List{l}, nil
}

func (s Node_readlink_Results_List) At(i int) Node_readlink_Results {
	return Node_readlink_Results{s.List.Struct(i)}
}
func (s Node_readlink_Results_List) Set(i int, v Node_readlink_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Node_readlink_Results_Promise is a wrapper for a Node_readlink_Results promised by a client call.
type Node_readlink_Results_Promise struct{ *capnp.Pipeline }

func (p Node_readlink_Results_Promise) Struct() (Node_readlink_Results, error) {
	s, err := p.Pipeline.Struct()
	return Node_readlink_Results{s}, err
}

type File struct{ Client capnp.Client }

func (c File) Read(ctx context.Context, params func(File_read_Params) error, opts ...capnp.CallOption) File_read_Results_Promise {
	if c.Client == nil {
		return File_read_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xac0b41617511a23d,
			MethodID:      0,
			InterfaceName: "fuse.capnp:File",
			MethodName:    "read",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(File_read_Params{Struct: s}) }
	}
	return File_read_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type File_Server interface {
	Read(File_read) error
}

func File_ServerToClient(s File_Server) File {
	c, _ := s.(server.Closer)
	return File{Client: server.New(File_Methods(nil, s), c)}
}

func File_Methods(methods []server.Method, s File_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xac0b41617511a23d,
			MethodID:      0,
			InterfaceName: "fuse.capnp:File",
			MethodName:    "read",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := File_read{c, opts, File_read_Params{Struct: p}, File_read_Results{Struct: r}}
			return s.Read(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// File_read holds the arguments for a server call to File.read.
type File_read struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  File_read_Params
	Results File_read_Results
}

type File_read_Params struct{ capnp.Struct }

func NewFile_read_Params(s *capnp.Segment) (File_read_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return File_read_Params{}, err
	}
	return File_read_Params{st}, nil
}

func NewRootFile_read_Params(s *capnp.Segment) (File_read_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return File_read_Params{}, err
	}
	return File_read_Params{st}, nil
}

func ReadRootFile_read_Params(msg *capnp.Message) (File_read_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return File_read_Params{}, err
	}
	return File_read_Params{root.Struct()}, nil
}
func (s File_read_Params) Offset() uint64 {
	return s.Struct.Uint64(0)
}

func (s File_read_Params) SetOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File_read_Params) Size() uint32 {
	return s.Struct.Uint32(8)
}

func (s File_read_Params) SetSize(v uint32) {
	s.Struct.SetUint32(8, v)
}

// File_read_Params_List is a list of File_read_Params.
type File_read_Params_List struct{ capnp.List }

// NewFile_read_Params creates a new list of File_read_Params.
func NewFile_read_Params_List(s *capnp.Segment, sz int32) (File_read_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	if err != nil {
		return File_read_Params_List{}, err
	}
	return File_read_Params_List{l}, nil
}

func (s File_read_Params_List) At(i int) File_read_Params { return File_read_Params{s.List.Struct(i)} }
func (s File_read_Params_List) Set(i int, v File_read_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// File_read_Params_Promise is a wrapper for a File_read_Params promised by a client call.
type File_read_Params_Promise struct{ *capnp.Pipeline }

func (p File_read_Params_Promise) Struct() (File_read_Params, error) {
	s, err := p.Pipeline.Struct()
	return File_read_Params{s}, err
}

type File_read_Results struct{ capnp.Struct }

func NewFile_read_Results(s *capnp.Segment) (File_read_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return File_read_Results{}, err
	}
	return File_read_Results{st}, nil
}

func NewRootFile_read_Results(s *capnp.Segment) (File_read_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return File_read_Results{}, err
	}
	return File_read_Results{st}, nil
}

func ReadRootFile_read_Results(msg *capnp.Message) (File_read_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return File_read_Results{}, err
	}
	return File_read_Results{root.Struct()}, nil
}
func (s File_read_Results) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s File_read_Results) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File_read_Results) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// File_read_Results_List is a list of File_read_Results.
type File_read_Results_List struct{ capnp.List }

// NewFile_read_Results creates a new list of File_read_Results.
func NewFile_read_Results_List(s *capnp.Segment, sz int32) (File_read_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return File_read_Results_List{}, err
	}
	return File_read_Results_List{l}, nil
}

func (s File_read_Results_List) At(i int) File_read_Results {
	return File_read_Results{s.List.Struct(i)}
}
func (s File_read_Results_List) Set(i int, v File_read_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// File_read_Results_Promise is a wrapper for a File_read_Results promised by a client call.
type File_read_Results_Promise struct{ *capnp.Pipeline }

func (p File_read_Results_Promise) Struct() (File_read_Results, error) {
	s, err := p.Pipeline.Struct()
	return File_read_Results{s}, err
}

type Directory struct{ Client capnp.Client }

func (c Directory) Read(ctx context.Context, params func(Directory_read_Params) error, opts ...capnp.CallOption) Directory_read_Results_Promise {
	if c.Client == nil {
		return Directory_read_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xb7579e3a9ec7fef6,
			MethodID:      0,
			InterfaceName: "fuse.capnp:Directory",
			MethodName:    "read",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Directory_read_Params{Struct: s}) }
	}
	return Directory_read_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Directory_Server interface {
	Read(Directory_read) error
}

func Directory_ServerToClient(s Directory_Server) Directory {
	c, _ := s.(server.Closer)
	return Directory{Client: server.New(Directory_Methods(nil, s), c)}
}

func Directory_Methods(methods []server.Method, s Directory_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb7579e3a9ec7fef6,
			MethodID:      0,
			InterfaceName: "fuse.capnp:Directory",
			MethodName:    "read",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Directory_read{c, opts, Directory_read_Params{Struct: p}, Directory_read_Results{Struct: r}}
			return s.Read(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Directory_read holds the arguments for a server call to Directory.read.
type Directory_read struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Directory_read_Params
	Results Directory_read_Results
}

type Directory_Entry struct{ capnp.Struct }

func NewDirectory_Entry(s *capnp.Segment) (Directory_Entry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	if err != nil {
		return Directory_Entry{}, err
	}
	return Directory_Entry{st}, nil
}

func NewRootDirectory_Entry(s *capnp.Segment) (Directory_Entry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	if err != nil {
		return Directory_Entry{}, err
	}
	return Directory_Entry{st}, nil
}

func ReadRootDirectory_Entry(msg *capnp.Message) (Directory_Entry, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Directory_Entry{}, err
	}
	return Directory_Entry{root.Struct()}, nil
}
func (s Directory_Entry) InodeNumber() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory_Entry) SetInodeNumber(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory_Entry) NextOffset() uint64 {
	return s.Struct.Uint64(8)
}

func (s Directory_Entry) SetNextOffset(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Directory_Entry) Type() Node_Type {
	return Node_Type(s.Struct.Uint16(16))
}

func (s Directory_Entry) SetType(v Node_Type) {
	s.Struct.SetUint16(16, uint16(v))
}

func (s Directory_Entry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Directory_Entry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory_Entry) NameBytes() ([]byte, error) {
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

func (s Directory_Entry) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Directory_Entry_List is a list of Directory_Entry.
type Directory_Entry_List struct{ capnp.List }

// NewDirectory_Entry creates a new list of Directory_Entry.
func NewDirectory_Entry_List(s *capnp.Segment, sz int32) (Directory_Entry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	if err != nil {
		return Directory_Entry_List{}, err
	}
	return Directory_Entry_List{l}, nil
}

func (s Directory_Entry_List) At(i int) Directory_Entry { return Directory_Entry{s.List.Struct(i)} }
func (s Directory_Entry_List) Set(i int, v Directory_Entry) error {
	return s.List.SetStruct(i, v.Struct)
}

// Directory_Entry_Promise is a wrapper for a Directory_Entry promised by a client call.
type Directory_Entry_Promise struct{ *capnp.Pipeline }

func (p Directory_Entry_Promise) Struct() (Directory_Entry, error) {
	s, err := p.Pipeline.Struct()
	return Directory_Entry{s}, err
}

type Directory_read_Params struct{ capnp.Struct }

func NewDirectory_read_Params(s *capnp.Segment) (Directory_read_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return Directory_read_Params{}, err
	}
	return Directory_read_Params{st}, nil
}

func NewRootDirectory_read_Params(s *capnp.Segment) (Directory_read_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return Directory_read_Params{}, err
	}
	return Directory_read_Params{st}, nil
}

func ReadRootDirectory_read_Params(msg *capnp.Message) (Directory_read_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Directory_read_Params{}, err
	}
	return Directory_read_Params{root.Struct()}, nil
}
func (s Directory_read_Params) Offset() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory_read_Params) SetOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory_read_Params) Count() uint32 {
	return s.Struct.Uint32(8)
}

func (s Directory_read_Params) SetCount(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Directory_read_Params_List is a list of Directory_read_Params.
type Directory_read_Params_List struct{ capnp.List }

// NewDirectory_read_Params creates a new list of Directory_read_Params.
func NewDirectory_read_Params_List(s *capnp.Segment, sz int32) (Directory_read_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	if err != nil {
		return Directory_read_Params_List{}, err
	}
	return Directory_read_Params_List{l}, nil
}

func (s Directory_read_Params_List) At(i int) Directory_read_Params {
	return Directory_read_Params{s.List.Struct(i)}
}
func (s Directory_read_Params_List) Set(i int, v Directory_read_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Directory_read_Params_Promise is a wrapper for a Directory_read_Params promised by a client call.
type Directory_read_Params_Promise struct{ *capnp.Pipeline }

func (p Directory_read_Params_Promise) Struct() (Directory_read_Params, error) {
	s, err := p.Pipeline.Struct()
	return Directory_read_Params{s}, err
}

type Directory_read_Results struct{ capnp.Struct }

func NewDirectory_read_Results(s *capnp.Segment) (Directory_read_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Directory_read_Results{}, err
	}
	return Directory_read_Results{st}, nil
}

func NewRootDirectory_read_Results(s *capnp.Segment) (Directory_read_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Directory_read_Results{}, err
	}
	return Directory_read_Results{st}, nil
}

func ReadRootDirectory_read_Results(msg *capnp.Message) (Directory_read_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Directory_read_Results{}, err
	}
	return Directory_read_Results{root.Struct()}, nil
}
func (s Directory_read_Results) Entries() (Directory_Entry_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Directory_Entry_List{}, err
	}
	return Directory_Entry_List{List: p.List()}, nil
}

func (s Directory_read_Results) HasEntries() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory_read_Results) SetEntries(v Directory_Entry_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewEntries sets the entries field to a newly
// allocated Directory_Entry_List, preferring placement in s's segment.
func (s Directory_read_Results) NewEntries(n int32) (Directory_Entry_List, error) {
	l, err := NewDirectory_Entry_List(s.Struct.Segment(), n)
	if err != nil {
		return Directory_Entry_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Directory_read_Results_List is a list of Directory_read_Results.
type Directory_read_Results_List struct{ capnp.List }

// NewDirectory_read_Results creates a new list of Directory_read_Results.
func NewDirectory_read_Results_List(s *capnp.Segment, sz int32) (Directory_read_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Directory_read_Results_List{}, err
	}
	return Directory_read_Results_List{l}, nil
}

func (s Directory_read_Results_List) At(i int) Directory_read_Results {
	return Directory_read_Results{s.List.Struct(i)}
}
func (s Directory_read_Results_List) Set(i int, v Directory_read_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Directory_read_Results_Promise is a wrapper for a Directory_read_Results promised by a client call.
type Directory_read_Results_Promise struct{ *capnp.Pipeline }

func (p Directory_read_Results_Promise) Struct() (Directory_read_Results, error) {
	s, err := p.Pipeline.Struct()
	return Directory_read_Results{s}, err
}
