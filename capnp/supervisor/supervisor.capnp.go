package supervisor

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	activity "zenhack.net/go/sandstorm/capnp/activity"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type Supervisor struct{ Client capnp.Client }

func (c Supervisor) GetMainView(ctx context.Context, params func(Supervisor_getMainView_Params) error, opts ...capnp.CallOption) Supervisor_getMainView_Results_Promise {
	if c.Client == nil {
		return Supervisor_getMainView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getMainView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_getMainView_Params{Struct: s}) }
	}
	return Supervisor_getMainView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) KeepAlive(ctx context.Context, params func(Supervisor_keepAlive_Params) error, opts ...capnp.CallOption) Supervisor_keepAlive_Results_Promise {
	if c.Client == nil {
		return Supervisor_keepAlive_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      1,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "keepAlive",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_keepAlive_Params{Struct: s}) }
	}
	return Supervisor_keepAlive_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) Shutdown(ctx context.Context, params func(Supervisor_shutdown_Params) error, opts ...capnp.CallOption) Supervisor_shutdown_Results_Promise {
	if c.Client == nil {
		return Supervisor_shutdown_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      2,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "shutdown",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_shutdown_Params{Struct: s}) }
	}
	return Supervisor_shutdown_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) ObsoleteGetGrainSize(ctx context.Context, params func(Supervisor_obsoleteGetGrainSize_Params) error, opts ...capnp.CallOption) Supervisor_obsoleteGetGrainSize_Results_Promise {
	if c.Client == nil {
		return Supervisor_obsoleteGetGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      3,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "obsoleteGetGrainSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_obsoleteGetGrainSize_Params{Struct: s}) }
	}
	return Supervisor_obsoleteGetGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) ObsoleteGetGrainSizeWhenDifferent(ctx context.Context, params func(Supervisor_obsoleteGetGrainSizeWhenDifferent_Params) error, opts ...capnp.CallOption) Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise {
	if c.Client == nil {
		return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "obsoleteGetGrainSizeWhenDifferent",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{Struct: s})
		}
	}
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) Restore(ctx context.Context, params func(Supervisor_restore_Params) error, opts ...capnp.CallOption) Supervisor_restore_Results_Promise {
	if c.Client == nil {
		return Supervisor_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      5,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_restore_Params{Struct: s}) }
	}
	return Supervisor_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) Drop(ctx context.Context, params func(Supervisor_drop_Params) error, opts ...capnp.CallOption) Supervisor_drop_Results_Promise {
	if c.Client == nil {
		return Supervisor_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      6,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_drop_Params{Struct: s}) }
	}
	return Supervisor_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) WatchLog(ctx context.Context, params func(Supervisor_watchLog_Params) error, opts ...capnp.CallOption) Supervisor_watchLog_Results_Promise {
	if c.Client == nil {
		return Supervisor_watchLog_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      7,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "watchLog",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_watchLog_Params{Struct: s}) }
	}
	return Supervisor_watchLog_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) SyncStorage(ctx context.Context, params func(Supervisor_syncStorage_Params) error, opts ...capnp.CallOption) Supervisor_syncStorage_Results_Promise {
	if c.Client == nil {
		return Supervisor_syncStorage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      8,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "syncStorage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_syncStorage_Params{Struct: s}) }
	}
	return Supervisor_syncStorage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) GetWwwFileHack(ctx context.Context, params func(Supervisor_getWwwFileHack_Params) error, opts ...capnp.CallOption) Supervisor_getWwwFileHack_Results_Promise {
	if c.Client == nil {
		return Supervisor_getWwwFileHack_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      9,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getWwwFileHack",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_getWwwFileHack_Params{Struct: s}) }
	}
	return Supervisor_getWwwFileHack_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Supervisor_Server interface {
	GetMainView(Supervisor_getMainView) error

	KeepAlive(Supervisor_keepAlive) error

	Shutdown(Supervisor_shutdown) error

	ObsoleteGetGrainSize(Supervisor_obsoleteGetGrainSize) error

	ObsoleteGetGrainSizeWhenDifferent(Supervisor_obsoleteGetGrainSizeWhenDifferent) error

	Restore(Supervisor_restore) error

	Drop(Supervisor_drop) error

	WatchLog(Supervisor_watchLog) error

	SyncStorage(Supervisor_syncStorage) error

	GetWwwFileHack(Supervisor_getWwwFileHack) error
}

func Supervisor_ServerToClient(s Supervisor_Server) Supervisor {
	c, _ := s.(server.Closer)
	return Supervisor{Client: server.New(Supervisor_Methods(nil, s), c)}
}

func Supervisor_Methods(methods []server.Method, s Supervisor_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 10)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getMainView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_getMainView{c, opts, Supervisor_getMainView_Params{Struct: p}, Supervisor_getMainView_Results{Struct: r}}
			return s.GetMainView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      1,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "keepAlive",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_keepAlive{c, opts, Supervisor_keepAlive_Params{Struct: p}, Supervisor_keepAlive_Results{Struct: r}}
			return s.KeepAlive(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      2,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "shutdown",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_shutdown{c, opts, Supervisor_shutdown_Params{Struct: p}, Supervisor_shutdown_Results{Struct: r}}
			return s.Shutdown(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      3,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "obsoleteGetGrainSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_obsoleteGetGrainSize{c, opts, Supervisor_obsoleteGetGrainSize_Params{Struct: p}, Supervisor_obsoleteGetGrainSize_Results{Struct: r}}
			return s.ObsoleteGetGrainSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "obsoleteGetGrainSizeWhenDifferent",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_obsoleteGetGrainSizeWhenDifferent{c, opts, Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{Struct: p}, Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{Struct: r}}
			return s.ObsoleteGetGrainSizeWhenDifferent(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      5,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_restore{c, opts, Supervisor_restore_Params{Struct: p}, Supervisor_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      6,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_drop{c, opts, Supervisor_drop_Params{Struct: p}, Supervisor_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      7,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "watchLog",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_watchLog{c, opts, Supervisor_watchLog_Params{Struct: p}, Supervisor_watchLog_Results{Struct: r}}
			return s.WatchLog(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      8,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "syncStorage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_syncStorage{c, opts, Supervisor_syncStorage_Params{Struct: p}, Supervisor_syncStorage_Results{Struct: r}}
			return s.SyncStorage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      9,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getWwwFileHack",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_getWwwFileHack{c, opts, Supervisor_getWwwFileHack_Params{Struct: p}, Supervisor_getWwwFileHack_Results{Struct: r}}
			return s.GetWwwFileHack(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// Supervisor_getMainView holds the arguments for a server call to Supervisor.getMainView.
type Supervisor_getMainView struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_getMainView_Params
	Results Supervisor_getMainView_Results
}

// Supervisor_keepAlive holds the arguments for a server call to Supervisor.keepAlive.
type Supervisor_keepAlive struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_keepAlive_Params
	Results Supervisor_keepAlive_Results
}

// Supervisor_shutdown holds the arguments for a server call to Supervisor.shutdown.
type Supervisor_shutdown struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_shutdown_Params
	Results Supervisor_shutdown_Results
}

// Supervisor_obsoleteGetGrainSize holds the arguments for a server call to Supervisor.obsoleteGetGrainSize.
type Supervisor_obsoleteGetGrainSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_obsoleteGetGrainSize_Params
	Results Supervisor_obsoleteGetGrainSize_Results
}

// Supervisor_obsoleteGetGrainSizeWhenDifferent holds the arguments for a server call to Supervisor.obsoleteGetGrainSizeWhenDifferent.
type Supervisor_obsoleteGetGrainSizeWhenDifferent struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_obsoleteGetGrainSizeWhenDifferent_Params
	Results Supervisor_obsoleteGetGrainSizeWhenDifferent_Results
}

// Supervisor_restore holds the arguments for a server call to Supervisor.restore.
type Supervisor_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_restore_Params
	Results Supervisor_restore_Results
}

// Supervisor_drop holds the arguments for a server call to Supervisor.drop.
type Supervisor_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_drop_Params
	Results Supervisor_drop_Results
}

// Supervisor_watchLog holds the arguments for a server call to Supervisor.watchLog.
type Supervisor_watchLog struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_watchLog_Params
	Results Supervisor_watchLog_Results
}

// Supervisor_syncStorage holds the arguments for a server call to Supervisor.syncStorage.
type Supervisor_syncStorage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_syncStorage_Params
	Results Supervisor_syncStorage_Results
}

// Supervisor_getWwwFileHack holds the arguments for a server call to Supervisor.getWwwFileHack.
type Supervisor_getWwwFileHack struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_getWwwFileHack_Params
	Results Supervisor_getWwwFileHack_Results
}

type Supervisor_WwwFileStatus uint16

// Values of Supervisor_WwwFileStatus.
const (
	Supervisor_WwwFileStatus_file      Supervisor_WwwFileStatus = 0
	Supervisor_WwwFileStatus_directory Supervisor_WwwFileStatus = 1
	Supervisor_WwwFileStatus_notFound  Supervisor_WwwFileStatus = 2
)

// String returns the enum's constant name.
func (c Supervisor_WwwFileStatus) String() string {
	switch c {
	case Supervisor_WwwFileStatus_file:
		return "file"
	case Supervisor_WwwFileStatus_directory:
		return "directory"
	case Supervisor_WwwFileStatus_notFound:
		return "notFound"

	default:
		return ""
	}
}

// Supervisor_WwwFileStatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Supervisor_WwwFileStatusFromString(c string) Supervisor_WwwFileStatus {
	switch c {
	case "file":
		return Supervisor_WwwFileStatus_file
	case "directory":
		return Supervisor_WwwFileStatus_directory
	case "notFound":
		return Supervisor_WwwFileStatus_notFound

	default:
		return 0
	}
}

type Supervisor_WwwFileStatus_List struct{ capnp.List }

func NewSupervisor_WwwFileStatus_List(s *capnp.Segment, sz int32) (Supervisor_WwwFileStatus_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Supervisor_WwwFileStatus_List{l.List}, err
}

func (l Supervisor_WwwFileStatus_List) At(i int) Supervisor_WwwFileStatus {
	ul := capnp.UInt16List{List: l.List}
	return Supervisor_WwwFileStatus(ul.At(i))
}

func (l Supervisor_WwwFileStatus_List) Set(i int, v Supervisor_WwwFileStatus) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Supervisor_getMainView_Params struct{ capnp.Struct }

// Supervisor_getMainView_Params_TypeID is the unique identifier for the type Supervisor_getMainView_Params.
const Supervisor_getMainView_Params_TypeID = 0xba19fd491deeb222

func NewSupervisor_getMainView_Params(s *capnp.Segment) (Supervisor_getMainView_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_getMainView_Params{st}, err
}

func NewRootSupervisor_getMainView_Params(s *capnp.Segment) (Supervisor_getMainView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_getMainView_Params{st}, err
}

func ReadRootSupervisor_getMainView_Params(msg *capnp.Message) (Supervisor_getMainView_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_getMainView_Params{root.Struct()}, err
}

func (s Supervisor_getMainView_Params) String() string {
	str, _ := text.Marshal(0xba19fd491deeb222, s.Struct)
	return str
}

// Supervisor_getMainView_Params_List is a list of Supervisor_getMainView_Params.
type Supervisor_getMainView_Params_List struct{ capnp.List }

// NewSupervisor_getMainView_Params creates a new list of Supervisor_getMainView_Params.
func NewSupervisor_getMainView_Params_List(s *capnp.Segment, sz int32) (Supervisor_getMainView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_getMainView_Params_List{l}, err
}

func (s Supervisor_getMainView_Params_List) At(i int) Supervisor_getMainView_Params {
	return Supervisor_getMainView_Params{s.List.Struct(i)}
}

func (s Supervisor_getMainView_Params_List) Set(i int, v Supervisor_getMainView_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getMainView_Params_Promise is a wrapper for a Supervisor_getMainView_Params promised by a client call.
type Supervisor_getMainView_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getMainView_Params_Promise) Struct() (Supervisor_getMainView_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getMainView_Params{s}, err
}

type Supervisor_getMainView_Results struct{ capnp.Struct }

// Supervisor_getMainView_Results_TypeID is the unique identifier for the type Supervisor_getMainView_Results.
const Supervisor_getMainView_Results_TypeID = 0x88abdb347bc63d0f

func NewSupervisor_getMainView_Results(s *capnp.Segment) (Supervisor_getMainView_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_getMainView_Results{st}, err
}

func NewRootSupervisor_getMainView_Results(s *capnp.Segment) (Supervisor_getMainView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_getMainView_Results{st}, err
}

func ReadRootSupervisor_getMainView_Results(msg *capnp.Message) (Supervisor_getMainView_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_getMainView_Results{root.Struct()}, err
}

func (s Supervisor_getMainView_Results) String() string {
	str, _ := text.Marshal(0x88abdb347bc63d0f, s.Struct)
	return str
}

func (s Supervisor_getMainView_Results) View() grain.UiView {
	p, _ := s.Struct.Ptr(0)
	return grain.UiView{Client: p.Interface().Client()}
}

func (s Supervisor_getMainView_Results) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_getMainView_Results) SetView(v grain.UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Supervisor_getMainView_Results_List is a list of Supervisor_getMainView_Results.
type Supervisor_getMainView_Results_List struct{ capnp.List }

// NewSupervisor_getMainView_Results creates a new list of Supervisor_getMainView_Results.
func NewSupervisor_getMainView_Results_List(s *capnp.Segment, sz int32) (Supervisor_getMainView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Supervisor_getMainView_Results_List{l}, err
}

func (s Supervisor_getMainView_Results_List) At(i int) Supervisor_getMainView_Results {
	return Supervisor_getMainView_Results{s.List.Struct(i)}
}

func (s Supervisor_getMainView_Results_List) Set(i int, v Supervisor_getMainView_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getMainView_Results_Promise is a wrapper for a Supervisor_getMainView_Results promised by a client call.
type Supervisor_getMainView_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getMainView_Results_Promise) Struct() (Supervisor_getMainView_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getMainView_Results{s}, err
}

func (p Supervisor_getMainView_Results_Promise) View() grain.UiView {
	return grain.UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Supervisor_keepAlive_Params struct{ capnp.Struct }

// Supervisor_keepAlive_Params_TypeID is the unique identifier for the type Supervisor_keepAlive_Params.
const Supervisor_keepAlive_Params_TypeID = 0xe4a4f650ea454237

func NewSupervisor_keepAlive_Params(s *capnp.Segment) (Supervisor_keepAlive_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_keepAlive_Params{st}, err
}

func NewRootSupervisor_keepAlive_Params(s *capnp.Segment) (Supervisor_keepAlive_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_keepAlive_Params{st}, err
}

func ReadRootSupervisor_keepAlive_Params(msg *capnp.Message) (Supervisor_keepAlive_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_keepAlive_Params{root.Struct()}, err
}

func (s Supervisor_keepAlive_Params) String() string {
	str, _ := text.Marshal(0xe4a4f650ea454237, s.Struct)
	return str
}

func (s Supervisor_keepAlive_Params) Core() SandstormCore {
	p, _ := s.Struct.Ptr(0)
	return SandstormCore{Client: p.Interface().Client()}
}

func (s Supervisor_keepAlive_Params) HasCore() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_keepAlive_Params) SetCore(v SandstormCore) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Supervisor_keepAlive_Params_List is a list of Supervisor_keepAlive_Params.
type Supervisor_keepAlive_Params_List struct{ capnp.List }

// NewSupervisor_keepAlive_Params creates a new list of Supervisor_keepAlive_Params.
func NewSupervisor_keepAlive_Params_List(s *capnp.Segment, sz int32) (Supervisor_keepAlive_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Supervisor_keepAlive_Params_List{l}, err
}

func (s Supervisor_keepAlive_Params_List) At(i int) Supervisor_keepAlive_Params {
	return Supervisor_keepAlive_Params{s.List.Struct(i)}
}

func (s Supervisor_keepAlive_Params_List) Set(i int, v Supervisor_keepAlive_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_keepAlive_Params_Promise is a wrapper for a Supervisor_keepAlive_Params promised by a client call.
type Supervisor_keepAlive_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_keepAlive_Params_Promise) Struct() (Supervisor_keepAlive_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_keepAlive_Params{s}, err
}

func (p Supervisor_keepAlive_Params_Promise) Core() SandstormCore {
	return SandstormCore{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Supervisor_keepAlive_Results struct{ capnp.Struct }

// Supervisor_keepAlive_Results_TypeID is the unique identifier for the type Supervisor_keepAlive_Results.
const Supervisor_keepAlive_Results_TypeID = 0xa0b4085080573e77

func NewSupervisor_keepAlive_Results(s *capnp.Segment) (Supervisor_keepAlive_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_keepAlive_Results{st}, err
}

func NewRootSupervisor_keepAlive_Results(s *capnp.Segment) (Supervisor_keepAlive_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_keepAlive_Results{st}, err
}

func ReadRootSupervisor_keepAlive_Results(msg *capnp.Message) (Supervisor_keepAlive_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_keepAlive_Results{root.Struct()}, err
}

func (s Supervisor_keepAlive_Results) String() string {
	str, _ := text.Marshal(0xa0b4085080573e77, s.Struct)
	return str
}

// Supervisor_keepAlive_Results_List is a list of Supervisor_keepAlive_Results.
type Supervisor_keepAlive_Results_List struct{ capnp.List }

// NewSupervisor_keepAlive_Results creates a new list of Supervisor_keepAlive_Results.
func NewSupervisor_keepAlive_Results_List(s *capnp.Segment, sz int32) (Supervisor_keepAlive_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_keepAlive_Results_List{l}, err
}

func (s Supervisor_keepAlive_Results_List) At(i int) Supervisor_keepAlive_Results {
	return Supervisor_keepAlive_Results{s.List.Struct(i)}
}

func (s Supervisor_keepAlive_Results_List) Set(i int, v Supervisor_keepAlive_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_keepAlive_Results_Promise is a wrapper for a Supervisor_keepAlive_Results promised by a client call.
type Supervisor_keepAlive_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_keepAlive_Results_Promise) Struct() (Supervisor_keepAlive_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_keepAlive_Results{s}, err
}

type Supervisor_shutdown_Params struct{ capnp.Struct }

// Supervisor_shutdown_Params_TypeID is the unique identifier for the type Supervisor_shutdown_Params.
const Supervisor_shutdown_Params_TypeID = 0xd597c8d788fec5df

func NewSupervisor_shutdown_Params(s *capnp.Segment) (Supervisor_shutdown_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_shutdown_Params{st}, err
}

func NewRootSupervisor_shutdown_Params(s *capnp.Segment) (Supervisor_shutdown_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_shutdown_Params{st}, err
}

func ReadRootSupervisor_shutdown_Params(msg *capnp.Message) (Supervisor_shutdown_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_shutdown_Params{root.Struct()}, err
}

func (s Supervisor_shutdown_Params) String() string {
	str, _ := text.Marshal(0xd597c8d788fec5df, s.Struct)
	return str
}

// Supervisor_shutdown_Params_List is a list of Supervisor_shutdown_Params.
type Supervisor_shutdown_Params_List struct{ capnp.List }

// NewSupervisor_shutdown_Params creates a new list of Supervisor_shutdown_Params.
func NewSupervisor_shutdown_Params_List(s *capnp.Segment, sz int32) (Supervisor_shutdown_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_shutdown_Params_List{l}, err
}

func (s Supervisor_shutdown_Params_List) At(i int) Supervisor_shutdown_Params {
	return Supervisor_shutdown_Params{s.List.Struct(i)}
}

func (s Supervisor_shutdown_Params_List) Set(i int, v Supervisor_shutdown_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_shutdown_Params_Promise is a wrapper for a Supervisor_shutdown_Params promised by a client call.
type Supervisor_shutdown_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_shutdown_Params_Promise) Struct() (Supervisor_shutdown_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_shutdown_Params{s}, err
}

type Supervisor_shutdown_Results struct{ capnp.Struct }

// Supervisor_shutdown_Results_TypeID is the unique identifier for the type Supervisor_shutdown_Results.
const Supervisor_shutdown_Results_TypeID = 0xcb7ee0fa69cd6e70

func NewSupervisor_shutdown_Results(s *capnp.Segment) (Supervisor_shutdown_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_shutdown_Results{st}, err
}

func NewRootSupervisor_shutdown_Results(s *capnp.Segment) (Supervisor_shutdown_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_shutdown_Results{st}, err
}

func ReadRootSupervisor_shutdown_Results(msg *capnp.Message) (Supervisor_shutdown_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_shutdown_Results{root.Struct()}, err
}

func (s Supervisor_shutdown_Results) String() string {
	str, _ := text.Marshal(0xcb7ee0fa69cd6e70, s.Struct)
	return str
}

// Supervisor_shutdown_Results_List is a list of Supervisor_shutdown_Results.
type Supervisor_shutdown_Results_List struct{ capnp.List }

// NewSupervisor_shutdown_Results creates a new list of Supervisor_shutdown_Results.
func NewSupervisor_shutdown_Results_List(s *capnp.Segment, sz int32) (Supervisor_shutdown_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_shutdown_Results_List{l}, err
}

func (s Supervisor_shutdown_Results_List) At(i int) Supervisor_shutdown_Results {
	return Supervisor_shutdown_Results{s.List.Struct(i)}
}

func (s Supervisor_shutdown_Results_List) Set(i int, v Supervisor_shutdown_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_shutdown_Results_Promise is a wrapper for a Supervisor_shutdown_Results promised by a client call.
type Supervisor_shutdown_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_shutdown_Results_Promise) Struct() (Supervisor_shutdown_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_shutdown_Results{s}, err
}

type Supervisor_obsoleteGetGrainSize_Params struct{ capnp.Struct }

// Supervisor_obsoleteGetGrainSize_Params_TypeID is the unique identifier for the type Supervisor_obsoleteGetGrainSize_Params.
const Supervisor_obsoleteGetGrainSize_Params_TypeID = 0xf3e98c16ae117300

func NewSupervisor_obsoleteGetGrainSize_Params(s *capnp.Segment) (Supervisor_obsoleteGetGrainSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSize_Params{st}, err
}

func NewRootSupervisor_obsoleteGetGrainSize_Params(s *capnp.Segment) (Supervisor_obsoleteGetGrainSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSize_Params{st}, err
}

func ReadRootSupervisor_obsoleteGetGrainSize_Params(msg *capnp.Message) (Supervisor_obsoleteGetGrainSize_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_obsoleteGetGrainSize_Params{root.Struct()}, err
}

func (s Supervisor_obsoleteGetGrainSize_Params) String() string {
	str, _ := text.Marshal(0xf3e98c16ae117300, s.Struct)
	return str
}

// Supervisor_obsoleteGetGrainSize_Params_List is a list of Supervisor_obsoleteGetGrainSize_Params.
type Supervisor_obsoleteGetGrainSize_Params_List struct{ capnp.List }

// NewSupervisor_obsoleteGetGrainSize_Params creates a new list of Supervisor_obsoleteGetGrainSize_Params.
func NewSupervisor_obsoleteGetGrainSize_Params_List(s *capnp.Segment, sz int32) (Supervisor_obsoleteGetGrainSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_obsoleteGetGrainSize_Params_List{l}, err
}

func (s Supervisor_obsoleteGetGrainSize_Params_List) At(i int) Supervisor_obsoleteGetGrainSize_Params {
	return Supervisor_obsoleteGetGrainSize_Params{s.List.Struct(i)}
}

func (s Supervisor_obsoleteGetGrainSize_Params_List) Set(i int, v Supervisor_obsoleteGetGrainSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_obsoleteGetGrainSize_Params_Promise is a wrapper for a Supervisor_obsoleteGetGrainSize_Params promised by a client call.
type Supervisor_obsoleteGetGrainSize_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_obsoleteGetGrainSize_Params_Promise) Struct() (Supervisor_obsoleteGetGrainSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_obsoleteGetGrainSize_Params{s}, err
}

type Supervisor_obsoleteGetGrainSize_Results struct{ capnp.Struct }

// Supervisor_obsoleteGetGrainSize_Results_TypeID is the unique identifier for the type Supervisor_obsoleteGetGrainSize_Results.
const Supervisor_obsoleteGetGrainSize_Results_TypeID = 0xdc76071bd22f9a4b

func NewSupervisor_obsoleteGetGrainSize_Results(s *capnp.Segment) (Supervisor_obsoleteGetGrainSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSize_Results{st}, err
}

func NewRootSupervisor_obsoleteGetGrainSize_Results(s *capnp.Segment) (Supervisor_obsoleteGetGrainSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSize_Results{st}, err
}

func ReadRootSupervisor_obsoleteGetGrainSize_Results(msg *capnp.Message) (Supervisor_obsoleteGetGrainSize_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_obsoleteGetGrainSize_Results{root.Struct()}, err
}

func (s Supervisor_obsoleteGetGrainSize_Results) String() string {
	str, _ := text.Marshal(0xdc76071bd22f9a4b, s.Struct)
	return str
}

func (s Supervisor_obsoleteGetGrainSize_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_obsoleteGetGrainSize_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_obsoleteGetGrainSize_Results_List is a list of Supervisor_obsoleteGetGrainSize_Results.
type Supervisor_obsoleteGetGrainSize_Results_List struct{ capnp.List }

// NewSupervisor_obsoleteGetGrainSize_Results creates a new list of Supervisor_obsoleteGetGrainSize_Results.
func NewSupervisor_obsoleteGetGrainSize_Results_List(s *capnp.Segment, sz int32) (Supervisor_obsoleteGetGrainSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_obsoleteGetGrainSize_Results_List{l}, err
}

func (s Supervisor_obsoleteGetGrainSize_Results_List) At(i int) Supervisor_obsoleteGetGrainSize_Results {
	return Supervisor_obsoleteGetGrainSize_Results{s.List.Struct(i)}
}

func (s Supervisor_obsoleteGetGrainSize_Results_List) Set(i int, v Supervisor_obsoleteGetGrainSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_obsoleteGetGrainSize_Results_Promise is a wrapper for a Supervisor_obsoleteGetGrainSize_Results promised by a client call.
type Supervisor_obsoleteGetGrainSize_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_obsoleteGetGrainSize_Results_Promise) Struct() (Supervisor_obsoleteGetGrainSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_obsoleteGetGrainSize_Results{s}, err
}

type Supervisor_obsoleteGetGrainSizeWhenDifferent_Params struct{ capnp.Struct }

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_TypeID is the unique identifier for the type Supervisor_obsoleteGetGrainSizeWhenDifferent_Params.
const Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_TypeID = 0xc08fb7eab4fb0e05

func NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Params(s *capnp.Segment) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{st}, err
}

func NewRootSupervisor_obsoleteGetGrainSizeWhenDifferent_Params(s *capnp.Segment) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{st}, err
}

func ReadRootSupervisor_obsoleteGetGrainSizeWhenDifferent_Params(msg *capnp.Message) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{root.Struct()}, err
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Params) String() string {
	str, _ := text.Marshal(0xc08fb7eab4fb0e05, s.Struct)
	return str
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Params) OldSize() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Params) SetOldSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List is a list of Supervisor_obsoleteGetGrainSizeWhenDifferent_Params.
type Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List struct{ capnp.List }

// NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Params creates a new list of Supervisor_obsoleteGetGrainSizeWhenDifferent_Params.
func NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Params_List(s *capnp.Segment, sz int32) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List{l}, err
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List) At(i int) Supervisor_obsoleteGetGrainSizeWhenDifferent_Params {
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{s.List.Struct(i)}
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_List) Set(i int, v Supervisor_obsoleteGetGrainSizeWhenDifferent_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_Promise is a wrapper for a Supervisor_obsoleteGetGrainSizeWhenDifferent_Params promised by a client call.
type Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_obsoleteGetGrainSizeWhenDifferent_Params_Promise) Struct() (Supervisor_obsoleteGetGrainSizeWhenDifferent_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Params{s}, err
}

type Supervisor_obsoleteGetGrainSizeWhenDifferent_Results struct{ capnp.Struct }

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_TypeID is the unique identifier for the type Supervisor_obsoleteGetGrainSizeWhenDifferent_Results.
const Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_TypeID = 0xcce106c2fbaa9b04

func NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Results(s *capnp.Segment) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{st}, err
}

func NewRootSupervisor_obsoleteGetGrainSizeWhenDifferent_Results(s *capnp.Segment) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{st}, err
}

func ReadRootSupervisor_obsoleteGetGrainSizeWhenDifferent_Results(msg *capnp.Message) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{root.Struct()}, err
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Results) String() string {
	str, _ := text.Marshal(0xcce106c2fbaa9b04, s.Struct)
	return str
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List is a list of Supervisor_obsoleteGetGrainSizeWhenDifferent_Results.
type Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List struct{ capnp.List }

// NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Results creates a new list of Supervisor_obsoleteGetGrainSizeWhenDifferent_Results.
func NewSupervisor_obsoleteGetGrainSizeWhenDifferent_Results_List(s *capnp.Segment, sz int32) (Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List{l}, err
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List) At(i int) Supervisor_obsoleteGetGrainSizeWhenDifferent_Results {
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{s.List.Struct(i)}
}

func (s Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_List) Set(i int, v Supervisor_obsoleteGetGrainSizeWhenDifferent_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise is a wrapper for a Supervisor_obsoleteGetGrainSizeWhenDifferent_Results promised by a client call.
type Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_obsoleteGetGrainSizeWhenDifferent_Results_Promise) Struct() (Supervisor_obsoleteGetGrainSizeWhenDifferent_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_obsoleteGetGrainSizeWhenDifferent_Results{s}, err
}

type Supervisor_restore_Params struct{ capnp.Struct }

// Supervisor_restore_Params_TypeID is the unique identifier for the type Supervisor_restore_Params.
const Supervisor_restore_Params_TypeID = 0xaae54cb2386e60ab

func NewSupervisor_restore_Params(s *capnp.Segment) (Supervisor_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Supervisor_restore_Params{st}, err
}

func NewRootSupervisor_restore_Params(s *capnp.Segment) (Supervisor_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Supervisor_restore_Params{st}, err
}

func ReadRootSupervisor_restore_Params(msg *capnp.Message) (Supervisor_restore_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_restore_Params{root.Struct()}, err
}

func (s Supervisor_restore_Params) String() string {
	str, _ := text.Marshal(0xaae54cb2386e60ab, s.Struct)
	return str
}

func (s Supervisor_restore_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	return SupervisorObjectId{Struct: p.Struct()}, err
}

func (s Supervisor_restore_Params) HasRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_restore_Params) SetRef(v SupervisorObjectId) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRef sets the ref field to a newly
// allocated SupervisorObjectId struct, preferring placement in s's segment.
func (s Supervisor_restore_Params) NewRef() (SupervisorObjectId, error) {
	ss, err := NewSupervisorObjectId(s.Struct.Segment())
	if err != nil {
		return SupervisorObjectId{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Supervisor_restore_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(1)
	return MembraneRequirement_List{List: p.List()}, err
}

func (s Supervisor_restore_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Supervisor_restore_Params) SetRequirements(v MembraneRequirement_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequirements sets the requirements field to a newly
// allocated MembraneRequirement_List, preferring placement in s's segment.
func (s Supervisor_restore_Params) NewRequirements(n int32) (MembraneRequirement_List, error) {
	l, err := NewMembraneRequirement_List(s.Struct.Segment(), n)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Supervisor_restore_Params) ParentToken() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Supervisor_restore_Params) HasParentToken() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Supervisor_restore_Params) SetParentToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, d.List.ToPtr())
}

// Supervisor_restore_Params_List is a list of Supervisor_restore_Params.
type Supervisor_restore_Params_List struct{ capnp.List }

// NewSupervisor_restore_Params creates a new list of Supervisor_restore_Params.
func NewSupervisor_restore_Params_List(s *capnp.Segment, sz int32) (Supervisor_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Supervisor_restore_Params_List{l}, err
}

func (s Supervisor_restore_Params_List) At(i int) Supervisor_restore_Params {
	return Supervisor_restore_Params{s.List.Struct(i)}
}

func (s Supervisor_restore_Params_List) Set(i int, v Supervisor_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_restore_Params_Promise is a wrapper for a Supervisor_restore_Params promised by a client call.
type Supervisor_restore_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_restore_Params_Promise) Struct() (Supervisor_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_restore_Params{s}, err
}

func (p Supervisor_restore_Params_Promise) Ref() SupervisorObjectId_Promise {
	return SupervisorObjectId_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Supervisor_restore_Results struct{ capnp.Struct }

// Supervisor_restore_Results_TypeID is the unique identifier for the type Supervisor_restore_Results.
const Supervisor_restore_Results_TypeID = 0x96fb2fd9e320599f

func NewSupervisor_restore_Results(s *capnp.Segment) (Supervisor_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_restore_Results{st}, err
}

func NewRootSupervisor_restore_Results(s *capnp.Segment) (Supervisor_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_restore_Results{st}, err
}

func ReadRootSupervisor_restore_Results(msg *capnp.Message) (Supervisor_restore_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_restore_Results{root.Struct()}, err
}

func (s Supervisor_restore_Results) String() string {
	str, _ := text.Marshal(0x96fb2fd9e320599f, s.Struct)
	return str
}

func (s Supervisor_restore_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Supervisor_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_restore_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Supervisor_restore_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Supervisor_restore_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Supervisor_restore_Results_List is a list of Supervisor_restore_Results.
type Supervisor_restore_Results_List struct{ capnp.List }

// NewSupervisor_restore_Results creates a new list of Supervisor_restore_Results.
func NewSupervisor_restore_Results_List(s *capnp.Segment, sz int32) (Supervisor_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Supervisor_restore_Results_List{l}, err
}

func (s Supervisor_restore_Results_List) At(i int) Supervisor_restore_Results {
	return Supervisor_restore_Results{s.List.Struct(i)}
}

func (s Supervisor_restore_Results_List) Set(i int, v Supervisor_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_restore_Results_Promise is a wrapper for a Supervisor_restore_Results promised by a client call.
type Supervisor_restore_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_restore_Results_Promise) Struct() (Supervisor_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_restore_Results{s}, err
}

func (p Supervisor_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Supervisor_drop_Params struct{ capnp.Struct }

// Supervisor_drop_Params_TypeID is the unique identifier for the type Supervisor_drop_Params.
const Supervisor_drop_Params_TypeID = 0xaf3c0d4c9b788c3b

func NewSupervisor_drop_Params(s *capnp.Segment) (Supervisor_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_drop_Params{st}, err
}

func NewRootSupervisor_drop_Params(s *capnp.Segment) (Supervisor_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_drop_Params{st}, err
}

func ReadRootSupervisor_drop_Params(msg *capnp.Message) (Supervisor_drop_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_drop_Params{root.Struct()}, err
}

func (s Supervisor_drop_Params) String() string {
	str, _ := text.Marshal(0xaf3c0d4c9b788c3b, s.Struct)
	return str
}

func (s Supervisor_drop_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	return SupervisorObjectId{Struct: p.Struct()}, err
}

func (s Supervisor_drop_Params) HasRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_drop_Params) SetRef(v SupervisorObjectId) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRef sets the ref field to a newly
// allocated SupervisorObjectId struct, preferring placement in s's segment.
func (s Supervisor_drop_Params) NewRef() (SupervisorObjectId, error) {
	ss, err := NewSupervisorObjectId(s.Struct.Segment())
	if err != nil {
		return SupervisorObjectId{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Supervisor_drop_Params_List is a list of Supervisor_drop_Params.
type Supervisor_drop_Params_List struct{ capnp.List }

// NewSupervisor_drop_Params creates a new list of Supervisor_drop_Params.
func NewSupervisor_drop_Params_List(s *capnp.Segment, sz int32) (Supervisor_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Supervisor_drop_Params_List{l}, err
}

func (s Supervisor_drop_Params_List) At(i int) Supervisor_drop_Params {
	return Supervisor_drop_Params{s.List.Struct(i)}
}

func (s Supervisor_drop_Params_List) Set(i int, v Supervisor_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_drop_Params_Promise is a wrapper for a Supervisor_drop_Params promised by a client call.
type Supervisor_drop_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_drop_Params_Promise) Struct() (Supervisor_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_drop_Params{s}, err
}

func (p Supervisor_drop_Params_Promise) Ref() SupervisorObjectId_Promise {
	return SupervisorObjectId_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Supervisor_drop_Results struct{ capnp.Struct }

// Supervisor_drop_Results_TypeID is the unique identifier for the type Supervisor_drop_Results.
const Supervisor_drop_Results_TypeID = 0x80e7bfc1abd2efa7

func NewSupervisor_drop_Results(s *capnp.Segment) (Supervisor_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_drop_Results{st}, err
}

func NewRootSupervisor_drop_Results(s *capnp.Segment) (Supervisor_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_drop_Results{st}, err
}

func ReadRootSupervisor_drop_Results(msg *capnp.Message) (Supervisor_drop_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_drop_Results{root.Struct()}, err
}

func (s Supervisor_drop_Results) String() string {
	str, _ := text.Marshal(0x80e7bfc1abd2efa7, s.Struct)
	return str
}

// Supervisor_drop_Results_List is a list of Supervisor_drop_Results.
type Supervisor_drop_Results_List struct{ capnp.List }

// NewSupervisor_drop_Results creates a new list of Supervisor_drop_Results.
func NewSupervisor_drop_Results_List(s *capnp.Segment, sz int32) (Supervisor_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_drop_Results_List{l}, err
}

func (s Supervisor_drop_Results_List) At(i int) Supervisor_drop_Results {
	return Supervisor_drop_Results{s.List.Struct(i)}
}

func (s Supervisor_drop_Results_List) Set(i int, v Supervisor_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_drop_Results_Promise is a wrapper for a Supervisor_drop_Results promised by a client call.
type Supervisor_drop_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_drop_Results_Promise) Struct() (Supervisor_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_drop_Results{s}, err
}

type Supervisor_watchLog_Params struct{ capnp.Struct }

// Supervisor_watchLog_Params_TypeID is the unique identifier for the type Supervisor_watchLog_Params.
const Supervisor_watchLog_Params_TypeID = 0xc152ab1174b40c0a

func NewSupervisor_watchLog_Params(s *capnp.Segment) (Supervisor_watchLog_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Supervisor_watchLog_Params{st}, err
}

func NewRootSupervisor_watchLog_Params(s *capnp.Segment) (Supervisor_watchLog_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Supervisor_watchLog_Params{st}, err
}

func ReadRootSupervisor_watchLog_Params(msg *capnp.Message) (Supervisor_watchLog_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_watchLog_Params{root.Struct()}, err
}

func (s Supervisor_watchLog_Params) String() string {
	str, _ := text.Marshal(0xc152ab1174b40c0a, s.Struct)
	return str
}

func (s Supervisor_watchLog_Params) BacklogAmount() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_watchLog_Params) SetBacklogAmount(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Supervisor_watchLog_Params) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(0)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Supervisor_watchLog_Params) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_watchLog_Params) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Supervisor_watchLog_Params_List is a list of Supervisor_watchLog_Params.
type Supervisor_watchLog_Params_List struct{ capnp.List }

// NewSupervisor_watchLog_Params creates a new list of Supervisor_watchLog_Params.
func NewSupervisor_watchLog_Params_List(s *capnp.Segment, sz int32) (Supervisor_watchLog_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Supervisor_watchLog_Params_List{l}, err
}

func (s Supervisor_watchLog_Params_List) At(i int) Supervisor_watchLog_Params {
	return Supervisor_watchLog_Params{s.List.Struct(i)}
}

func (s Supervisor_watchLog_Params_List) Set(i int, v Supervisor_watchLog_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_watchLog_Params_Promise is a wrapper for a Supervisor_watchLog_Params promised by a client call.
type Supervisor_watchLog_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_watchLog_Params_Promise) Struct() (Supervisor_watchLog_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_watchLog_Params{s}, err
}

func (p Supervisor_watchLog_Params_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Supervisor_watchLog_Results struct{ capnp.Struct }

// Supervisor_watchLog_Results_TypeID is the unique identifier for the type Supervisor_watchLog_Results.
const Supervisor_watchLog_Results_TypeID = 0x98053037c12fa689

func NewSupervisor_watchLog_Results(s *capnp.Segment) (Supervisor_watchLog_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_watchLog_Results{st}, err
}

func NewRootSupervisor_watchLog_Results(s *capnp.Segment) (Supervisor_watchLog_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Supervisor_watchLog_Results{st}, err
}

func ReadRootSupervisor_watchLog_Results(msg *capnp.Message) (Supervisor_watchLog_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_watchLog_Results{root.Struct()}, err
}

func (s Supervisor_watchLog_Results) String() string {
	str, _ := text.Marshal(0x98053037c12fa689, s.Struct)
	return str
}

func (s Supervisor_watchLog_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s Supervisor_watchLog_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_watchLog_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Supervisor_watchLog_Results_List is a list of Supervisor_watchLog_Results.
type Supervisor_watchLog_Results_List struct{ capnp.List }

// NewSupervisor_watchLog_Results creates a new list of Supervisor_watchLog_Results.
func NewSupervisor_watchLog_Results_List(s *capnp.Segment, sz int32) (Supervisor_watchLog_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Supervisor_watchLog_Results_List{l}, err
}

func (s Supervisor_watchLog_Results_List) At(i int) Supervisor_watchLog_Results {
	return Supervisor_watchLog_Results{s.List.Struct(i)}
}

func (s Supervisor_watchLog_Results_List) Set(i int, v Supervisor_watchLog_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_watchLog_Results_Promise is a wrapper for a Supervisor_watchLog_Results promised by a client call.
type Supervisor_watchLog_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_watchLog_Results_Promise) Struct() (Supervisor_watchLog_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_watchLog_Results{s}, err
}

func (p Supervisor_watchLog_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Supervisor_syncStorage_Params struct{ capnp.Struct }

// Supervisor_syncStorage_Params_TypeID is the unique identifier for the type Supervisor_syncStorage_Params.
const Supervisor_syncStorage_Params_TypeID = 0xcf3e8fcfd0506bd0

func NewSupervisor_syncStorage_Params(s *capnp.Segment) (Supervisor_syncStorage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_syncStorage_Params{st}, err
}

func NewRootSupervisor_syncStorage_Params(s *capnp.Segment) (Supervisor_syncStorage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_syncStorage_Params{st}, err
}

func ReadRootSupervisor_syncStorage_Params(msg *capnp.Message) (Supervisor_syncStorage_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_syncStorage_Params{root.Struct()}, err
}

func (s Supervisor_syncStorage_Params) String() string {
	str, _ := text.Marshal(0xcf3e8fcfd0506bd0, s.Struct)
	return str
}

// Supervisor_syncStorage_Params_List is a list of Supervisor_syncStorage_Params.
type Supervisor_syncStorage_Params_List struct{ capnp.List }

// NewSupervisor_syncStorage_Params creates a new list of Supervisor_syncStorage_Params.
func NewSupervisor_syncStorage_Params_List(s *capnp.Segment, sz int32) (Supervisor_syncStorage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_syncStorage_Params_List{l}, err
}

func (s Supervisor_syncStorage_Params_List) At(i int) Supervisor_syncStorage_Params {
	return Supervisor_syncStorage_Params{s.List.Struct(i)}
}

func (s Supervisor_syncStorage_Params_List) Set(i int, v Supervisor_syncStorage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_syncStorage_Params_Promise is a wrapper for a Supervisor_syncStorage_Params promised by a client call.
type Supervisor_syncStorage_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_syncStorage_Params_Promise) Struct() (Supervisor_syncStorage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_syncStorage_Params{s}, err
}

type Supervisor_syncStorage_Results struct{ capnp.Struct }

// Supervisor_syncStorage_Results_TypeID is the unique identifier for the type Supervisor_syncStorage_Results.
const Supervisor_syncStorage_Results_TypeID = 0xba8b9f7f3a411a03

func NewSupervisor_syncStorage_Results(s *capnp.Segment) (Supervisor_syncStorage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_syncStorage_Results{st}, err
}

func NewRootSupervisor_syncStorage_Results(s *capnp.Segment) (Supervisor_syncStorage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_syncStorage_Results{st}, err
}

func ReadRootSupervisor_syncStorage_Results(msg *capnp.Message) (Supervisor_syncStorage_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_syncStorage_Results{root.Struct()}, err
}

func (s Supervisor_syncStorage_Results) String() string {
	str, _ := text.Marshal(0xba8b9f7f3a411a03, s.Struct)
	return str
}

// Supervisor_syncStorage_Results_List is a list of Supervisor_syncStorage_Results.
type Supervisor_syncStorage_Results_List struct{ capnp.List }

// NewSupervisor_syncStorage_Results creates a new list of Supervisor_syncStorage_Results.
func NewSupervisor_syncStorage_Results_List(s *capnp.Segment, sz int32) (Supervisor_syncStorage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_syncStorage_Results_List{l}, err
}

func (s Supervisor_syncStorage_Results_List) At(i int) Supervisor_syncStorage_Results {
	return Supervisor_syncStorage_Results{s.List.Struct(i)}
}

func (s Supervisor_syncStorage_Results_List) Set(i int, v Supervisor_syncStorage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_syncStorage_Results_Promise is a wrapper for a Supervisor_syncStorage_Results promised by a client call.
type Supervisor_syncStorage_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_syncStorage_Results_Promise) Struct() (Supervisor_syncStorage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_syncStorage_Results{s}, err
}

type Supervisor_getWwwFileHack_Params struct{ capnp.Struct }

// Supervisor_getWwwFileHack_Params_TypeID is the unique identifier for the type Supervisor_getWwwFileHack_Params.
const Supervisor_getWwwFileHack_Params_TypeID = 0xf9c6e362d6fcb22a

func NewSupervisor_getWwwFileHack_Params(s *capnp.Segment) (Supervisor_getWwwFileHack_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Supervisor_getWwwFileHack_Params{st}, err
}

func NewRootSupervisor_getWwwFileHack_Params(s *capnp.Segment) (Supervisor_getWwwFileHack_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Supervisor_getWwwFileHack_Params{st}, err
}

func ReadRootSupervisor_getWwwFileHack_Params(msg *capnp.Message) (Supervisor_getWwwFileHack_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_getWwwFileHack_Params{root.Struct()}, err
}

func (s Supervisor_getWwwFileHack_Params) String() string {
	str, _ := text.Marshal(0xf9c6e362d6fcb22a, s.Struct)
	return str
}

func (s Supervisor_getWwwFileHack_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Supervisor_getWwwFileHack_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_getWwwFileHack_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Supervisor_getWwwFileHack_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Supervisor_getWwwFileHack_Params) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(1)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Supervisor_getWwwFileHack_Params) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Supervisor_getWwwFileHack_Params) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// Supervisor_getWwwFileHack_Params_List is a list of Supervisor_getWwwFileHack_Params.
type Supervisor_getWwwFileHack_Params_List struct{ capnp.List }

// NewSupervisor_getWwwFileHack_Params creates a new list of Supervisor_getWwwFileHack_Params.
func NewSupervisor_getWwwFileHack_Params_List(s *capnp.Segment, sz int32) (Supervisor_getWwwFileHack_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Supervisor_getWwwFileHack_Params_List{l}, err
}

func (s Supervisor_getWwwFileHack_Params_List) At(i int) Supervisor_getWwwFileHack_Params {
	return Supervisor_getWwwFileHack_Params{s.List.Struct(i)}
}

func (s Supervisor_getWwwFileHack_Params_List) Set(i int, v Supervisor_getWwwFileHack_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getWwwFileHack_Params_Promise is a wrapper for a Supervisor_getWwwFileHack_Params promised by a client call.
type Supervisor_getWwwFileHack_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getWwwFileHack_Params_Promise) Struct() (Supervisor_getWwwFileHack_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getWwwFileHack_Params{s}, err
}

func (p Supervisor_getWwwFileHack_Params_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

type Supervisor_getWwwFileHack_Results struct{ capnp.Struct }

// Supervisor_getWwwFileHack_Results_TypeID is the unique identifier for the type Supervisor_getWwwFileHack_Results.
const Supervisor_getWwwFileHack_Results_TypeID = 0x902651d6de458996

func NewSupervisor_getWwwFileHack_Results(s *capnp.Segment) (Supervisor_getWwwFileHack_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getWwwFileHack_Results{st}, err
}

func NewRootSupervisor_getWwwFileHack_Results(s *capnp.Segment) (Supervisor_getWwwFileHack_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getWwwFileHack_Results{st}, err
}

func ReadRootSupervisor_getWwwFileHack_Results(msg *capnp.Message) (Supervisor_getWwwFileHack_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_getWwwFileHack_Results{root.Struct()}, err
}

func (s Supervisor_getWwwFileHack_Results) String() string {
	str, _ := text.Marshal(0x902651d6de458996, s.Struct)
	return str
}

func (s Supervisor_getWwwFileHack_Results) Status() Supervisor_WwwFileStatus {
	return Supervisor_WwwFileStatus(s.Struct.Uint16(0))
}

func (s Supervisor_getWwwFileHack_Results) SetStatus(v Supervisor_WwwFileStatus) {
	s.Struct.SetUint16(0, uint16(v))
}

// Supervisor_getWwwFileHack_Results_List is a list of Supervisor_getWwwFileHack_Results.
type Supervisor_getWwwFileHack_Results_List struct{ capnp.List }

// NewSupervisor_getWwwFileHack_Results creates a new list of Supervisor_getWwwFileHack_Results.
func NewSupervisor_getWwwFileHack_Results_List(s *capnp.Segment, sz int32) (Supervisor_getWwwFileHack_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_getWwwFileHack_Results_List{l}, err
}

func (s Supervisor_getWwwFileHack_Results_List) At(i int) Supervisor_getWwwFileHack_Results {
	return Supervisor_getWwwFileHack_Results{s.List.Struct(i)}
}

func (s Supervisor_getWwwFileHack_Results_List) Set(i int, v Supervisor_getWwwFileHack_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getWwwFileHack_Results_Promise is a wrapper for a Supervisor_getWwwFileHack_Results promised by a client call.
type Supervisor_getWwwFileHack_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getWwwFileHack_Results_Promise) Struct() (Supervisor_getWwwFileHack_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getWwwFileHack_Results{s}, err
}

type SandstormCore struct{ Client capnp.Client }

func (c SandstormCore) Restore(ctx context.Context, params func(SandstormCore_restore_Params) error, opts ...capnp.CallOption) SandstormCore_restore_Results_Promise {
	if c.Client == nil {
		return SandstormCore_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_restore_Params{Struct: s}) }
	}
	return SandstormCore_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) MakeToken(ctx context.Context, params func(SandstormCore_makeToken_Params) error, opts ...capnp.CallOption) SandstormCore_makeToken_Results_Promise {
	if c.Client == nil {
		return SandstormCore_makeToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      1,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "makeToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_makeToken_Params{Struct: s}) }
	}
	return SandstormCore_makeToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) GetOwnerNotificationTarget(ctx context.Context, params func(SandstormCore_getOwnerNotificationTarget_Params) error, opts ...capnp.CallOption) SandstormCore_getOwnerNotificationTarget_Results_Promise {
	if c.Client == nil {
		return SandstormCore_getOwnerNotificationTarget_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      2,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "getOwnerNotificationTarget",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_getOwnerNotificationTarget_Params{Struct: s}) }
	}
	return SandstormCore_getOwnerNotificationTarget_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) Drop(ctx context.Context, params func(SandstormCore_drop_Params) error, opts ...capnp.CallOption) SandstormCore_drop_Results_Promise {
	if c.Client == nil {
		return SandstormCore_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      3,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_drop_Params{Struct: s}) }
	}
	return SandstormCore_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) CheckRequirements(ctx context.Context, params func(SandstormCore_checkRequirements_Params) error, opts ...capnp.CallOption) SandstormCore_checkRequirements_Results_Promise {
	if c.Client == nil {
		return SandstormCore_checkRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "checkRequirements",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_checkRequirements_Params{Struct: s}) }
	}
	return SandstormCore_checkRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) MakeChildToken(ctx context.Context, params func(SandstormCore_makeChildToken_Params) error, opts ...capnp.CallOption) SandstormCore_makeChildToken_Results_Promise {
	if c.Client == nil {
		return SandstormCore_makeChildToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      5,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "makeChildToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_makeChildToken_Params{Struct: s}) }
	}
	return SandstormCore_makeChildToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) ClaimRequest(ctx context.Context, params func(SandstormCore_claimRequest_Params) error, opts ...capnp.CallOption) SandstormCore_claimRequest_Results_Promise {
	if c.Client == nil {
		return SandstormCore_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      6,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "claimRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_claimRequest_Params{Struct: s}) }
	}
	return SandstormCore_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) BackgroundActivity(ctx context.Context, params func(SandstormCore_backgroundActivity_Params) error, opts ...capnp.CallOption) SandstormCore_backgroundActivity_Results_Promise {
	if c.Client == nil {
		return SandstormCore_backgroundActivity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      7,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "backgroundActivity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_backgroundActivity_Params{Struct: s}) }
	}
	return SandstormCore_backgroundActivity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) ReportGrainSize(ctx context.Context, params func(SandstormCore_reportGrainSize_Params) error, opts ...capnp.CallOption) SandstormCore_reportGrainSize_Results_Promise {
	if c.Client == nil {
		return SandstormCore_reportGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      8,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "reportGrainSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_reportGrainSize_Params{Struct: s}) }
	}
	return SandstormCore_reportGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormCore) GetIdentityId(ctx context.Context, params func(SandstormCore_getIdentityId_Params) error, opts ...capnp.CallOption) SandstormCore_getIdentityId_Results_Promise {
	if c.Client == nil {
		return SandstormCore_getIdentityId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      9,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "getIdentityId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_getIdentityId_Params{Struct: s}) }
	}
	return SandstormCore_getIdentityId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormCore_Server interface {
	Restore(SandstormCore_restore) error

	MakeToken(SandstormCore_makeToken) error

	GetOwnerNotificationTarget(SandstormCore_getOwnerNotificationTarget) error

	Drop(SandstormCore_drop) error

	CheckRequirements(SandstormCore_checkRequirements) error

	MakeChildToken(SandstormCore_makeChildToken) error

	ClaimRequest(SandstormCore_claimRequest) error

	BackgroundActivity(SandstormCore_backgroundActivity) error

	ReportGrainSize(SandstormCore_reportGrainSize) error

	GetIdentityId(SandstormCore_getIdentityId) error
}

func SandstormCore_ServerToClient(s SandstormCore_Server) SandstormCore {
	c, _ := s.(server.Closer)
	return SandstormCore{Client: server.New(SandstormCore_Methods(nil, s), c)}
}

func SandstormCore_Methods(methods []server.Method, s SandstormCore_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 10)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_restore{c, opts, SandstormCore_restore_Params{Struct: p}, SandstormCore_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      1,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "makeToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_makeToken{c, opts, SandstormCore_makeToken_Params{Struct: p}, SandstormCore_makeToken_Results{Struct: r}}
			return s.MakeToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      2,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "getOwnerNotificationTarget",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_getOwnerNotificationTarget{c, opts, SandstormCore_getOwnerNotificationTarget_Params{Struct: p}, SandstormCore_getOwnerNotificationTarget_Results{Struct: r}}
			return s.GetOwnerNotificationTarget(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      3,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_drop{c, opts, SandstormCore_drop_Params{Struct: p}, SandstormCore_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "checkRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_checkRequirements{c, opts, SandstormCore_checkRequirements_Params{Struct: p}, SandstormCore_checkRequirements_Results{Struct: r}}
			return s.CheckRequirements(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      5,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "makeChildToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_makeChildToken{c, opts, SandstormCore_makeChildToken_Params{Struct: p}, SandstormCore_makeChildToken_Results{Struct: r}}
			return s.MakeChildToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      6,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "claimRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_claimRequest{c, opts, SandstormCore_claimRequest_Params{Struct: p}, SandstormCore_claimRequest_Results{Struct: r}}
			return s.ClaimRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      7,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "backgroundActivity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_backgroundActivity{c, opts, SandstormCore_backgroundActivity_Params{Struct: p}, SandstormCore_backgroundActivity_Results{Struct: r}}
			return s.BackgroundActivity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      8,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "reportGrainSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_reportGrainSize{c, opts, SandstormCore_reportGrainSize_Params{Struct: p}, SandstormCore_reportGrainSize_Results{Struct: r}}
			return s.ReportGrainSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9ba45778a294b60c,
			MethodID:      9,
			InterfaceName: "supervisor.capnp:SandstormCore",
			MethodName:    "getIdentityId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_getIdentityId{c, opts, SandstormCore_getIdentityId_Params{Struct: p}, SandstormCore_getIdentityId_Results{Struct: r}}
			return s.GetIdentityId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SandstormCore_restore holds the arguments for a server call to SandstormCore.restore.
type SandstormCore_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_restore_Params
	Results SandstormCore_restore_Results
}

// SandstormCore_makeToken holds the arguments for a server call to SandstormCore.makeToken.
type SandstormCore_makeToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_makeToken_Params
	Results SandstormCore_makeToken_Results
}

// SandstormCore_getOwnerNotificationTarget holds the arguments for a server call to SandstormCore.getOwnerNotificationTarget.
type SandstormCore_getOwnerNotificationTarget struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_getOwnerNotificationTarget_Params
	Results SandstormCore_getOwnerNotificationTarget_Results
}

// SandstormCore_drop holds the arguments for a server call to SandstormCore.drop.
type SandstormCore_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_drop_Params
	Results SandstormCore_drop_Results
}

// SandstormCore_checkRequirements holds the arguments for a server call to SandstormCore.checkRequirements.
type SandstormCore_checkRequirements struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_checkRequirements_Params
	Results SandstormCore_checkRequirements_Results
}

// SandstormCore_makeChildToken holds the arguments for a server call to SandstormCore.makeChildToken.
type SandstormCore_makeChildToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_makeChildToken_Params
	Results SandstormCore_makeChildToken_Results
}

// SandstormCore_claimRequest holds the arguments for a server call to SandstormCore.claimRequest.
type SandstormCore_claimRequest struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_claimRequest_Params
	Results SandstormCore_claimRequest_Results
}

// SandstormCore_backgroundActivity holds the arguments for a server call to SandstormCore.backgroundActivity.
type SandstormCore_backgroundActivity struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_backgroundActivity_Params
	Results SandstormCore_backgroundActivity_Results
}

// SandstormCore_reportGrainSize holds the arguments for a server call to SandstormCore.reportGrainSize.
type SandstormCore_reportGrainSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_reportGrainSize_Params
	Results SandstormCore_reportGrainSize_Results
}

// SandstormCore_getIdentityId holds the arguments for a server call to SandstormCore.getIdentityId.
type SandstormCore_getIdentityId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_getIdentityId_Params
	Results SandstormCore_getIdentityId_Results
}

type SandstormCore_RequirementObserver struct{ Client capnp.Client }

func (c SandstormCore_RequirementObserver) Observe(ctx context.Context, params func(SandstormCore_RequirementObserver_observe_Params) error, opts ...capnp.CallOption) SandstormCore_RequirementObserver_observe_Results_Promise {
	if c.Client == nil {
		return SandstormCore_RequirementObserver_observe_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa3eb8443f86b46f7,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SandstormCore.RequirementObserver",
			MethodName:    "observe",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCore_RequirementObserver_observe_Params{Struct: s}) }
	}
	return SandstormCore_RequirementObserver_observe_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormCore_RequirementObserver_Server interface {
	Observe(SandstormCore_RequirementObserver_observe) error
}

func SandstormCore_RequirementObserver_ServerToClient(s SandstormCore_RequirementObserver_Server) SandstormCore_RequirementObserver {
	c, _ := s.(server.Closer)
	return SandstormCore_RequirementObserver{Client: server.New(SandstormCore_RequirementObserver_Methods(nil, s), c)}
}

func SandstormCore_RequirementObserver_Methods(methods []server.Method, s SandstormCore_RequirementObserver_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa3eb8443f86b46f7,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SandstormCore.RequirementObserver",
			MethodName:    "observe",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCore_RequirementObserver_observe{c, opts, SandstormCore_RequirementObserver_observe_Params{Struct: p}, SandstormCore_RequirementObserver_observe_Results{Struct: r}}
			return s.Observe(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// SandstormCore_RequirementObserver_observe holds the arguments for a server call to SandstormCore_RequirementObserver.observe.
type SandstormCore_RequirementObserver_observe struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCore_RequirementObserver_observe_Params
	Results SandstormCore_RequirementObserver_observe_Results
}

type SandstormCore_RequirementObserver_observe_Params struct{ capnp.Struct }

// SandstormCore_RequirementObserver_observe_Params_TypeID is the unique identifier for the type SandstormCore_RequirementObserver_observe_Params.
const SandstormCore_RequirementObserver_observe_Params_TypeID = 0x99bc33fd5d97c13d

func NewSandstormCore_RequirementObserver_observe_Params(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_RequirementObserver_observe_Params{st}, err
}

func NewRootSandstormCore_RequirementObserver_observe_Params(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_RequirementObserver_observe_Params{st}, err
}

func ReadRootSandstormCore_RequirementObserver_observe_Params(msg *capnp.Message) (SandstormCore_RequirementObserver_observe_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_RequirementObserver_observe_Params{root.Struct()}, err
}

func (s SandstormCore_RequirementObserver_observe_Params) String() string {
	str, _ := text.Marshal(0x99bc33fd5d97c13d, s.Struct)
	return str
}

// SandstormCore_RequirementObserver_observe_Params_List is a list of SandstormCore_RequirementObserver_observe_Params.
type SandstormCore_RequirementObserver_observe_Params_List struct{ capnp.List }

// NewSandstormCore_RequirementObserver_observe_Params creates a new list of SandstormCore_RequirementObserver_observe_Params.
func NewSandstormCore_RequirementObserver_observe_Params_List(s *capnp.Segment, sz int32) (SandstormCore_RequirementObserver_observe_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_RequirementObserver_observe_Params_List{l}, err
}

func (s SandstormCore_RequirementObserver_observe_Params_List) At(i int) SandstormCore_RequirementObserver_observe_Params {
	return SandstormCore_RequirementObserver_observe_Params{s.List.Struct(i)}
}

func (s SandstormCore_RequirementObserver_observe_Params_List) Set(i int, v SandstormCore_RequirementObserver_observe_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_RequirementObserver_observe_Params_Promise is a wrapper for a SandstormCore_RequirementObserver_observe_Params promised by a client call.
type SandstormCore_RequirementObserver_observe_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_RequirementObserver_observe_Params_Promise) Struct() (SandstormCore_RequirementObserver_observe_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_RequirementObserver_observe_Params{s}, err
}

type SandstormCore_RequirementObserver_observe_Results struct{ capnp.Struct }

// SandstormCore_RequirementObserver_observe_Results_TypeID is the unique identifier for the type SandstormCore_RequirementObserver_observe_Results.
const SandstormCore_RequirementObserver_observe_Results_TypeID = 0xcc28367ccc71b3df

func NewSandstormCore_RequirementObserver_observe_Results(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_RequirementObserver_observe_Results{st}, err
}

func NewRootSandstormCore_RequirementObserver_observe_Results(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_RequirementObserver_observe_Results{st}, err
}

func ReadRootSandstormCore_RequirementObserver_observe_Results(msg *capnp.Message) (SandstormCore_RequirementObserver_observe_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_RequirementObserver_observe_Results{root.Struct()}, err
}

func (s SandstormCore_RequirementObserver_observe_Results) String() string {
	str, _ := text.Marshal(0xcc28367ccc71b3df, s.Struct)
	return str
}

// SandstormCore_RequirementObserver_observe_Results_List is a list of SandstormCore_RequirementObserver_observe_Results.
type SandstormCore_RequirementObserver_observe_Results_List struct{ capnp.List }

// NewSandstormCore_RequirementObserver_observe_Results creates a new list of SandstormCore_RequirementObserver_observe_Results.
func NewSandstormCore_RequirementObserver_observe_Results_List(s *capnp.Segment, sz int32) (SandstormCore_RequirementObserver_observe_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_RequirementObserver_observe_Results_List{l}, err
}

func (s SandstormCore_RequirementObserver_observe_Results_List) At(i int) SandstormCore_RequirementObserver_observe_Results {
	return SandstormCore_RequirementObserver_observe_Results{s.List.Struct(i)}
}

func (s SandstormCore_RequirementObserver_observe_Results_List) Set(i int, v SandstormCore_RequirementObserver_observe_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_RequirementObserver_observe_Results_Promise is a wrapper for a SandstormCore_RequirementObserver_observe_Results promised by a client call.
type SandstormCore_RequirementObserver_observe_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_RequirementObserver_observe_Results_Promise) Struct() (SandstormCore_RequirementObserver_observe_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_RequirementObserver_observe_Results{s}, err
}

type SandstormCore_restore_Params struct{ capnp.Struct }

// SandstormCore_restore_Params_TypeID is the unique identifier for the type SandstormCore_restore_Params.
const SandstormCore_restore_Params_TypeID = 0xf839f92f21f00b08

func NewSandstormCore_restore_Params(s *capnp.Segment) (SandstormCore_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_restore_Params{st}, err
}

func NewRootSandstormCore_restore_Params(s *capnp.Segment) (SandstormCore_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_restore_Params{st}, err
}

func ReadRootSandstormCore_restore_Params(msg *capnp.Message) (SandstormCore_restore_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_restore_Params{root.Struct()}, err
}

func (s SandstormCore_restore_Params) String() string {
	str, _ := text.Marshal(0xf839f92f21f00b08, s.Struct)
	return str
}

func (s SandstormCore_restore_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_restore_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_restore_Params) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormCore_restore_Params_List is a list of SandstormCore_restore_Params.
type SandstormCore_restore_Params_List struct{ capnp.List }

// NewSandstormCore_restore_Params creates a new list of SandstormCore_restore_Params.
func NewSandstormCore_restore_Params_List(s *capnp.Segment, sz int32) (SandstormCore_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_restore_Params_List{l}, err
}

func (s SandstormCore_restore_Params_List) At(i int) SandstormCore_restore_Params {
	return SandstormCore_restore_Params{s.List.Struct(i)}
}

func (s SandstormCore_restore_Params_List) Set(i int, v SandstormCore_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_restore_Params_Promise is a wrapper for a SandstormCore_restore_Params promised by a client call.
type SandstormCore_restore_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_restore_Params_Promise) Struct() (SandstormCore_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_restore_Params{s}, err
}

type SandstormCore_restore_Results struct{ capnp.Struct }

// SandstormCore_restore_Results_TypeID is the unique identifier for the type SandstormCore_restore_Results.
const SandstormCore_restore_Results_TypeID = 0x92e92771f2b6b2b7

func NewSandstormCore_restore_Results(s *capnp.Segment) (SandstormCore_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_restore_Results{st}, err
}

func NewRootSandstormCore_restore_Results(s *capnp.Segment) (SandstormCore_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_restore_Results{st}, err
}

func ReadRootSandstormCore_restore_Results(msg *capnp.Message) (SandstormCore_restore_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_restore_Results{root.Struct()}, err
}

func (s SandstormCore_restore_Results) String() string {
	str, _ := text.Marshal(0x92e92771f2b6b2b7, s.Struct)
	return str
}

func (s SandstormCore_restore_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormCore_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_restore_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormCore_restore_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormCore_restore_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SandstormCore_restore_Results_List is a list of SandstormCore_restore_Results.
type SandstormCore_restore_Results_List struct{ capnp.List }

// NewSandstormCore_restore_Results creates a new list of SandstormCore_restore_Results.
func NewSandstormCore_restore_Results_List(s *capnp.Segment, sz int32) (SandstormCore_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_restore_Results_List{l}, err
}

func (s SandstormCore_restore_Results_List) At(i int) SandstormCore_restore_Results {
	return SandstormCore_restore_Results{s.List.Struct(i)}
}

func (s SandstormCore_restore_Results_List) Set(i int, v SandstormCore_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_restore_Results_Promise is a wrapper for a SandstormCore_restore_Results promised by a client call.
type SandstormCore_restore_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_restore_Results_Promise) Struct() (SandstormCore_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_restore_Results{s}, err
}

func (p SandstormCore_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormCore_makeToken_Params struct{ capnp.Struct }

// SandstormCore_makeToken_Params_TypeID is the unique identifier for the type SandstormCore_makeToken_Params.
const SandstormCore_makeToken_Params_TypeID = 0xf0e8359b121f97d2

func NewSandstormCore_makeToken_Params(s *capnp.Segment) (SandstormCore_makeToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SandstormCore_makeToken_Params{st}, err
}

func NewRootSandstormCore_makeToken_Params(s *capnp.Segment) (SandstormCore_makeToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SandstormCore_makeToken_Params{st}, err
}

func ReadRootSandstormCore_makeToken_Params(msg *capnp.Message) (SandstormCore_makeToken_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_makeToken_Params{root.Struct()}, err
}

func (s SandstormCore_makeToken_Params) String() string {
	str, _ := text.Marshal(0xf0e8359b121f97d2, s.Struct)
	return str
}

func (s SandstormCore_makeToken_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	return SupervisorObjectId{Struct: p.Struct()}, err
}

func (s SandstormCore_makeToken_Params) HasRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeToken_Params) SetRef(v SupervisorObjectId) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRef sets the ref field to a newly
// allocated SupervisorObjectId struct, preferring placement in s's segment.
func (s SandstormCore_makeToken_Params) NewRef() (SupervisorObjectId, error) {
	ss, err := NewSupervisorObjectId(s.Struct.Segment())
	if err != nil {
		return SupervisorObjectId{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SandstormCore_makeToken_Params) Owner() (ApiTokenOwner, error) {
	p, err := s.Struct.Ptr(1)
	return ApiTokenOwner{Struct: p.Struct()}, err
}

func (s SandstormCore_makeToken_Params) HasOwner() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeToken_Params) SetOwner(v ApiTokenOwner) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewOwner sets the owner field to a newly
// allocated ApiTokenOwner struct, preferring placement in s's segment.
func (s SandstormCore_makeToken_Params) NewOwner() (ApiTokenOwner, error) {
	ss, err := NewApiTokenOwner(s.Struct.Segment())
	if err != nil {
		return ApiTokenOwner{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s SandstormCore_makeToken_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(2)
	return MembraneRequirement_List{List: p.List()}, err
}

func (s SandstormCore_makeToken_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeToken_Params) SetRequirements(v MembraneRequirement_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewRequirements sets the requirements field to a newly
// allocated MembraneRequirement_List, preferring placement in s's segment.
func (s SandstormCore_makeToken_Params) NewRequirements(n int32) (MembraneRequirement_List, error) {
	l, err := NewMembraneRequirement_List(s.Struct.Segment(), n)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// SandstormCore_makeToken_Params_List is a list of SandstormCore_makeToken_Params.
type SandstormCore_makeToken_Params_List struct{ capnp.List }

// NewSandstormCore_makeToken_Params creates a new list of SandstormCore_makeToken_Params.
func NewSandstormCore_makeToken_Params_List(s *capnp.Segment, sz int32) (SandstormCore_makeToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return SandstormCore_makeToken_Params_List{l}, err
}

func (s SandstormCore_makeToken_Params_List) At(i int) SandstormCore_makeToken_Params {
	return SandstormCore_makeToken_Params{s.List.Struct(i)}
}

func (s SandstormCore_makeToken_Params_List) Set(i int, v SandstormCore_makeToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_makeToken_Params_Promise is a wrapper for a SandstormCore_makeToken_Params promised by a client call.
type SandstormCore_makeToken_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_makeToken_Params_Promise) Struct() (SandstormCore_makeToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_makeToken_Params{s}, err
}

func (p SandstormCore_makeToken_Params_Promise) Ref() SupervisorObjectId_Promise {
	return SupervisorObjectId_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p SandstormCore_makeToken_Params_Promise) Owner() ApiTokenOwner_Promise {
	return ApiTokenOwner_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormCore_makeToken_Results struct{ capnp.Struct }

// SandstormCore_makeToken_Results_TypeID is the unique identifier for the type SandstormCore_makeToken_Results.
const SandstormCore_makeToken_Results_TypeID = 0x9f96d4b948521f91

func NewSandstormCore_makeToken_Results(s *capnp.Segment) (SandstormCore_makeToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_makeToken_Results{st}, err
}

func NewRootSandstormCore_makeToken_Results(s *capnp.Segment) (SandstormCore_makeToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_makeToken_Results{st}, err
}

func ReadRootSandstormCore_makeToken_Results(msg *capnp.Message) (SandstormCore_makeToken_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_makeToken_Results{root.Struct()}, err
}

func (s SandstormCore_makeToken_Results) String() string {
	str, _ := text.Marshal(0x9f96d4b948521f91, s.Struct)
	return str
}

func (s SandstormCore_makeToken_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_makeToken_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeToken_Results) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormCore_makeToken_Results_List is a list of SandstormCore_makeToken_Results.
type SandstormCore_makeToken_Results_List struct{ capnp.List }

// NewSandstormCore_makeToken_Results creates a new list of SandstormCore_makeToken_Results.
func NewSandstormCore_makeToken_Results_List(s *capnp.Segment, sz int32) (SandstormCore_makeToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_makeToken_Results_List{l}, err
}

func (s SandstormCore_makeToken_Results_List) At(i int) SandstormCore_makeToken_Results {
	return SandstormCore_makeToken_Results{s.List.Struct(i)}
}

func (s SandstormCore_makeToken_Results_List) Set(i int, v SandstormCore_makeToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_makeToken_Results_Promise is a wrapper for a SandstormCore_makeToken_Results promised by a client call.
type SandstormCore_makeToken_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_makeToken_Results_Promise) Struct() (SandstormCore_makeToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_makeToken_Results{s}, err
}

type SandstormCore_getOwnerNotificationTarget_Params struct{ capnp.Struct }

// SandstormCore_getOwnerNotificationTarget_Params_TypeID is the unique identifier for the type SandstormCore_getOwnerNotificationTarget_Params.
const SandstormCore_getOwnerNotificationTarget_Params_TypeID = 0xaf72d693dbf4bf54

func NewSandstormCore_getOwnerNotificationTarget_Params(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_getOwnerNotificationTarget_Params{st}, err
}

func NewRootSandstormCore_getOwnerNotificationTarget_Params(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_getOwnerNotificationTarget_Params{st}, err
}

func ReadRootSandstormCore_getOwnerNotificationTarget_Params(msg *capnp.Message) (SandstormCore_getOwnerNotificationTarget_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_getOwnerNotificationTarget_Params{root.Struct()}, err
}

func (s SandstormCore_getOwnerNotificationTarget_Params) String() string {
	str, _ := text.Marshal(0xaf72d693dbf4bf54, s.Struct)
	return str
}

// SandstormCore_getOwnerNotificationTarget_Params_List is a list of SandstormCore_getOwnerNotificationTarget_Params.
type SandstormCore_getOwnerNotificationTarget_Params_List struct{ capnp.List }

// NewSandstormCore_getOwnerNotificationTarget_Params creates a new list of SandstormCore_getOwnerNotificationTarget_Params.
func NewSandstormCore_getOwnerNotificationTarget_Params_List(s *capnp.Segment, sz int32) (SandstormCore_getOwnerNotificationTarget_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_getOwnerNotificationTarget_Params_List{l}, err
}

func (s SandstormCore_getOwnerNotificationTarget_Params_List) At(i int) SandstormCore_getOwnerNotificationTarget_Params {
	return SandstormCore_getOwnerNotificationTarget_Params{s.List.Struct(i)}
}

func (s SandstormCore_getOwnerNotificationTarget_Params_List) Set(i int, v SandstormCore_getOwnerNotificationTarget_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_getOwnerNotificationTarget_Params_Promise is a wrapper for a SandstormCore_getOwnerNotificationTarget_Params promised by a client call.
type SandstormCore_getOwnerNotificationTarget_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_getOwnerNotificationTarget_Params_Promise) Struct() (SandstormCore_getOwnerNotificationTarget_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_getOwnerNotificationTarget_Params{s}, err
}

type SandstormCore_getOwnerNotificationTarget_Results struct{ capnp.Struct }

// SandstormCore_getOwnerNotificationTarget_Results_TypeID is the unique identifier for the type SandstormCore_getOwnerNotificationTarget_Results.
const SandstormCore_getOwnerNotificationTarget_Results_TypeID = 0x888c6d95df2cc976

func NewSandstormCore_getOwnerNotificationTarget_Results(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getOwnerNotificationTarget_Results{st}, err
}

func NewRootSandstormCore_getOwnerNotificationTarget_Results(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getOwnerNotificationTarget_Results{st}, err
}

func ReadRootSandstormCore_getOwnerNotificationTarget_Results(msg *capnp.Message) (SandstormCore_getOwnerNotificationTarget_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_getOwnerNotificationTarget_Results{root.Struct()}, err
}

func (s SandstormCore_getOwnerNotificationTarget_Results) String() string {
	str, _ := text.Marshal(0x888c6d95df2cc976, s.Struct)
	return str
}

func (s SandstormCore_getOwnerNotificationTarget_Results) Owner() activity.NotificationTarget {
	p, _ := s.Struct.Ptr(0)
	return activity.NotificationTarget{Client: p.Interface().Client()}
}

func (s SandstormCore_getOwnerNotificationTarget_Results) HasOwner() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_getOwnerNotificationTarget_Results) SetOwner(v activity.NotificationTarget) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormCore_getOwnerNotificationTarget_Results_List is a list of SandstormCore_getOwnerNotificationTarget_Results.
type SandstormCore_getOwnerNotificationTarget_Results_List struct{ capnp.List }

// NewSandstormCore_getOwnerNotificationTarget_Results creates a new list of SandstormCore_getOwnerNotificationTarget_Results.
func NewSandstormCore_getOwnerNotificationTarget_Results_List(s *capnp.Segment, sz int32) (SandstormCore_getOwnerNotificationTarget_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_getOwnerNotificationTarget_Results_List{l}, err
}

func (s SandstormCore_getOwnerNotificationTarget_Results_List) At(i int) SandstormCore_getOwnerNotificationTarget_Results {
	return SandstormCore_getOwnerNotificationTarget_Results{s.List.Struct(i)}
}

func (s SandstormCore_getOwnerNotificationTarget_Results_List) Set(i int, v SandstormCore_getOwnerNotificationTarget_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_getOwnerNotificationTarget_Results_Promise is a wrapper for a SandstormCore_getOwnerNotificationTarget_Results promised by a client call.
type SandstormCore_getOwnerNotificationTarget_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_getOwnerNotificationTarget_Results_Promise) Struct() (SandstormCore_getOwnerNotificationTarget_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_getOwnerNotificationTarget_Results{s}, err
}

func (p SandstormCore_getOwnerNotificationTarget_Results_Promise) Owner() activity.NotificationTarget {
	return activity.NotificationTarget{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormCore_drop_Params struct{ capnp.Struct }

// SandstormCore_drop_Params_TypeID is the unique identifier for the type SandstormCore_drop_Params.
const SandstormCore_drop_Params_TypeID = 0xe03b8c8163d957c6

func NewSandstormCore_drop_Params(s *capnp.Segment) (SandstormCore_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_drop_Params{st}, err
}

func NewRootSandstormCore_drop_Params(s *capnp.Segment) (SandstormCore_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_drop_Params{st}, err
}

func ReadRootSandstormCore_drop_Params(msg *capnp.Message) (SandstormCore_drop_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_drop_Params{root.Struct()}, err
}

func (s SandstormCore_drop_Params) String() string {
	str, _ := text.Marshal(0xe03b8c8163d957c6, s.Struct)
	return str
}

func (s SandstormCore_drop_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_drop_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_drop_Params) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormCore_drop_Params_List is a list of SandstormCore_drop_Params.
type SandstormCore_drop_Params_List struct{ capnp.List }

// NewSandstormCore_drop_Params creates a new list of SandstormCore_drop_Params.
func NewSandstormCore_drop_Params_List(s *capnp.Segment, sz int32) (SandstormCore_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_drop_Params_List{l}, err
}

func (s SandstormCore_drop_Params_List) At(i int) SandstormCore_drop_Params {
	return SandstormCore_drop_Params{s.List.Struct(i)}
}

func (s SandstormCore_drop_Params_List) Set(i int, v SandstormCore_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_drop_Params_Promise is a wrapper for a SandstormCore_drop_Params promised by a client call.
type SandstormCore_drop_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_drop_Params_Promise) Struct() (SandstormCore_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_drop_Params{s}, err
}

type SandstormCore_drop_Results struct{ capnp.Struct }

// SandstormCore_drop_Results_TypeID is the unique identifier for the type SandstormCore_drop_Results.
const SandstormCore_drop_Results_TypeID = 0xce435c92a97c1b97

func NewSandstormCore_drop_Results(s *capnp.Segment) (SandstormCore_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_drop_Results{st}, err
}

func NewRootSandstormCore_drop_Results(s *capnp.Segment) (SandstormCore_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_drop_Results{st}, err
}

func ReadRootSandstormCore_drop_Results(msg *capnp.Message) (SandstormCore_drop_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_drop_Results{root.Struct()}, err
}

func (s SandstormCore_drop_Results) String() string {
	str, _ := text.Marshal(0xce435c92a97c1b97, s.Struct)
	return str
}

// SandstormCore_drop_Results_List is a list of SandstormCore_drop_Results.
type SandstormCore_drop_Results_List struct{ capnp.List }

// NewSandstormCore_drop_Results creates a new list of SandstormCore_drop_Results.
func NewSandstormCore_drop_Results_List(s *capnp.Segment, sz int32) (SandstormCore_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_drop_Results_List{l}, err
}

func (s SandstormCore_drop_Results_List) At(i int) SandstormCore_drop_Results {
	return SandstormCore_drop_Results{s.List.Struct(i)}
}

func (s SandstormCore_drop_Results_List) Set(i int, v SandstormCore_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_drop_Results_Promise is a wrapper for a SandstormCore_drop_Results promised by a client call.
type SandstormCore_drop_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_drop_Results_Promise) Struct() (SandstormCore_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_drop_Results{s}, err
}

type SandstormCore_checkRequirements_Params struct{ capnp.Struct }

// SandstormCore_checkRequirements_Params_TypeID is the unique identifier for the type SandstormCore_checkRequirements_Params.
const SandstormCore_checkRequirements_Params_TypeID = 0x8867ef4f53bc45c3

func NewSandstormCore_checkRequirements_Params(s *capnp.Segment) (SandstormCore_checkRequirements_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_checkRequirements_Params{st}, err
}

func NewRootSandstormCore_checkRequirements_Params(s *capnp.Segment) (SandstormCore_checkRequirements_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_checkRequirements_Params{st}, err
}

func ReadRootSandstormCore_checkRequirements_Params(msg *capnp.Message) (SandstormCore_checkRequirements_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_checkRequirements_Params{root.Struct()}, err
}

func (s SandstormCore_checkRequirements_Params) String() string {
	str, _ := text.Marshal(0x8867ef4f53bc45c3, s.Struct)
	return str
}

func (s SandstormCore_checkRequirements_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(0)
	return MembraneRequirement_List{List: p.List()}, err
}

func (s SandstormCore_checkRequirements_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_checkRequirements_Params) SetRequirements(v MembraneRequirement_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewRequirements sets the requirements field to a newly
// allocated MembraneRequirement_List, preferring placement in s's segment.
func (s SandstormCore_checkRequirements_Params) NewRequirements(n int32) (MembraneRequirement_List, error) {
	l, err := NewMembraneRequirement_List(s.Struct.Segment(), n)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// SandstormCore_checkRequirements_Params_List is a list of SandstormCore_checkRequirements_Params.
type SandstormCore_checkRequirements_Params_List struct{ capnp.List }

// NewSandstormCore_checkRequirements_Params creates a new list of SandstormCore_checkRequirements_Params.
func NewSandstormCore_checkRequirements_Params_List(s *capnp.Segment, sz int32) (SandstormCore_checkRequirements_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_checkRequirements_Params_List{l}, err
}

func (s SandstormCore_checkRequirements_Params_List) At(i int) SandstormCore_checkRequirements_Params {
	return SandstormCore_checkRequirements_Params{s.List.Struct(i)}
}

func (s SandstormCore_checkRequirements_Params_List) Set(i int, v SandstormCore_checkRequirements_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_checkRequirements_Params_Promise is a wrapper for a SandstormCore_checkRequirements_Params promised by a client call.
type SandstormCore_checkRequirements_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_checkRequirements_Params_Promise) Struct() (SandstormCore_checkRequirements_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_checkRequirements_Params{s}, err
}

type SandstormCore_checkRequirements_Results struct{ capnp.Struct }

// SandstormCore_checkRequirements_Results_TypeID is the unique identifier for the type SandstormCore_checkRequirements_Results.
const SandstormCore_checkRequirements_Results_TypeID = 0xca83e6f36908ed7f

func NewSandstormCore_checkRequirements_Results(s *capnp.Segment) (SandstormCore_checkRequirements_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_checkRequirements_Results{st}, err
}

func NewRootSandstormCore_checkRequirements_Results(s *capnp.Segment) (SandstormCore_checkRequirements_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_checkRequirements_Results{st}, err
}

func ReadRootSandstormCore_checkRequirements_Results(msg *capnp.Message) (SandstormCore_checkRequirements_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_checkRequirements_Results{root.Struct()}, err
}

func (s SandstormCore_checkRequirements_Results) String() string {
	str, _ := text.Marshal(0xca83e6f36908ed7f, s.Struct)
	return str
}

func (s SandstormCore_checkRequirements_Results) Observer() SandstormCore_RequirementObserver {
	p, _ := s.Struct.Ptr(0)
	return SandstormCore_RequirementObserver{Client: p.Interface().Client()}
}

func (s SandstormCore_checkRequirements_Results) HasObserver() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_checkRequirements_Results) SetObserver(v SandstormCore_RequirementObserver) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormCore_checkRequirements_Results_List is a list of SandstormCore_checkRequirements_Results.
type SandstormCore_checkRequirements_Results_List struct{ capnp.List }

// NewSandstormCore_checkRequirements_Results creates a new list of SandstormCore_checkRequirements_Results.
func NewSandstormCore_checkRequirements_Results_List(s *capnp.Segment, sz int32) (SandstormCore_checkRequirements_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_checkRequirements_Results_List{l}, err
}

func (s SandstormCore_checkRequirements_Results_List) At(i int) SandstormCore_checkRequirements_Results {
	return SandstormCore_checkRequirements_Results{s.List.Struct(i)}
}

func (s SandstormCore_checkRequirements_Results_List) Set(i int, v SandstormCore_checkRequirements_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_checkRequirements_Results_Promise is a wrapper for a SandstormCore_checkRequirements_Results promised by a client call.
type SandstormCore_checkRequirements_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_checkRequirements_Results_Promise) Struct() (SandstormCore_checkRequirements_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_checkRequirements_Results{s}, err
}

func (p SandstormCore_checkRequirements_Results_Promise) Observer() SandstormCore_RequirementObserver {
	return SandstormCore_RequirementObserver{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormCore_makeChildToken_Params struct{ capnp.Struct }

// SandstormCore_makeChildToken_Params_TypeID is the unique identifier for the type SandstormCore_makeChildToken_Params.
const SandstormCore_makeChildToken_Params_TypeID = 0x9b25c148edb2b020

func NewSandstormCore_makeChildToken_Params(s *capnp.Segment) (SandstormCore_makeChildToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SandstormCore_makeChildToken_Params{st}, err
}

func NewRootSandstormCore_makeChildToken_Params(s *capnp.Segment) (SandstormCore_makeChildToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SandstormCore_makeChildToken_Params{st}, err
}

func ReadRootSandstormCore_makeChildToken_Params(msg *capnp.Message) (SandstormCore_makeChildToken_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_makeChildToken_Params{root.Struct()}, err
}

func (s SandstormCore_makeChildToken_Params) String() string {
	str, _ := text.Marshal(0x9b25c148edb2b020, s.Struct)
	return str
}

func (s SandstormCore_makeChildToken_Params) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_makeChildToken_Params) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeChildToken_Params) SetParent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s SandstormCore_makeChildToken_Params) Owner() (ApiTokenOwner, error) {
	p, err := s.Struct.Ptr(1)
	return ApiTokenOwner{Struct: p.Struct()}, err
}

func (s SandstormCore_makeChildToken_Params) HasOwner() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeChildToken_Params) SetOwner(v ApiTokenOwner) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewOwner sets the owner field to a newly
// allocated ApiTokenOwner struct, preferring placement in s's segment.
func (s SandstormCore_makeChildToken_Params) NewOwner() (ApiTokenOwner, error) {
	ss, err := NewApiTokenOwner(s.Struct.Segment())
	if err != nil {
		return ApiTokenOwner{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s SandstormCore_makeChildToken_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(2)
	return MembraneRequirement_List{List: p.List()}, err
}

func (s SandstormCore_makeChildToken_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeChildToken_Params) SetRequirements(v MembraneRequirement_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewRequirements sets the requirements field to a newly
// allocated MembraneRequirement_List, preferring placement in s's segment.
func (s SandstormCore_makeChildToken_Params) NewRequirements(n int32) (MembraneRequirement_List, error) {
	l, err := NewMembraneRequirement_List(s.Struct.Segment(), n)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// SandstormCore_makeChildToken_Params_List is a list of SandstormCore_makeChildToken_Params.
type SandstormCore_makeChildToken_Params_List struct{ capnp.List }

// NewSandstormCore_makeChildToken_Params creates a new list of SandstormCore_makeChildToken_Params.
func NewSandstormCore_makeChildToken_Params_List(s *capnp.Segment, sz int32) (SandstormCore_makeChildToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return SandstormCore_makeChildToken_Params_List{l}, err
}

func (s SandstormCore_makeChildToken_Params_List) At(i int) SandstormCore_makeChildToken_Params {
	return SandstormCore_makeChildToken_Params{s.List.Struct(i)}
}

func (s SandstormCore_makeChildToken_Params_List) Set(i int, v SandstormCore_makeChildToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_makeChildToken_Params_Promise is a wrapper for a SandstormCore_makeChildToken_Params promised by a client call.
type SandstormCore_makeChildToken_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_makeChildToken_Params_Promise) Struct() (SandstormCore_makeChildToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_makeChildToken_Params{s}, err
}

func (p SandstormCore_makeChildToken_Params_Promise) Owner() ApiTokenOwner_Promise {
	return ApiTokenOwner_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormCore_makeChildToken_Results struct{ capnp.Struct }

// SandstormCore_makeChildToken_Results_TypeID is the unique identifier for the type SandstormCore_makeChildToken_Results.
const SandstormCore_makeChildToken_Results_TypeID = 0x9ea56a46fc87138a

func NewSandstormCore_makeChildToken_Results(s *capnp.Segment) (SandstormCore_makeChildToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_makeChildToken_Results{st}, err
}

func NewRootSandstormCore_makeChildToken_Results(s *capnp.Segment) (SandstormCore_makeChildToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_makeChildToken_Results{st}, err
}

func ReadRootSandstormCore_makeChildToken_Results(msg *capnp.Message) (SandstormCore_makeChildToken_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_makeChildToken_Results{root.Struct()}, err
}

func (s SandstormCore_makeChildToken_Results) String() string {
	str, _ := text.Marshal(0x9ea56a46fc87138a, s.Struct)
	return str
}

func (s SandstormCore_makeChildToken_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_makeChildToken_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeChildToken_Results) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormCore_makeChildToken_Results_List is a list of SandstormCore_makeChildToken_Results.
type SandstormCore_makeChildToken_Results_List struct{ capnp.List }

// NewSandstormCore_makeChildToken_Results creates a new list of SandstormCore_makeChildToken_Results.
func NewSandstormCore_makeChildToken_Results_List(s *capnp.Segment, sz int32) (SandstormCore_makeChildToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_makeChildToken_Results_List{l}, err
}

func (s SandstormCore_makeChildToken_Results_List) At(i int) SandstormCore_makeChildToken_Results {
	return SandstormCore_makeChildToken_Results{s.List.Struct(i)}
}

func (s SandstormCore_makeChildToken_Results_List) Set(i int, v SandstormCore_makeChildToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_makeChildToken_Results_Promise is a wrapper for a SandstormCore_makeChildToken_Results promised by a client call.
type SandstormCore_makeChildToken_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_makeChildToken_Results_Promise) Struct() (SandstormCore_makeChildToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_makeChildToken_Results{s}, err
}

type SandstormCore_claimRequest_Params struct{ capnp.Struct }

// SandstormCore_claimRequest_Params_TypeID is the unique identifier for the type SandstormCore_claimRequest_Params.
const SandstormCore_claimRequest_Params_TypeID = 0xeeaeb799e53e0b01

func NewSandstormCore_claimRequest_Params(s *capnp.Segment) (SandstormCore_claimRequest_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormCore_claimRequest_Params{st}, err
}

func NewRootSandstormCore_claimRequest_Params(s *capnp.Segment) (SandstormCore_claimRequest_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormCore_claimRequest_Params{st}, err
}

func ReadRootSandstormCore_claimRequest_Params(msg *capnp.Message) (SandstormCore_claimRequest_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_claimRequest_Params{root.Struct()}, err
}

func (s SandstormCore_claimRequest_Params) String() string {
	str, _ := text.Marshal(0xeeaeb799e53e0b01, s.Struct)
	return str
}

func (s SandstormCore_claimRequest_Params) RequestToken() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SandstormCore_claimRequest_Params) HasRequestToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_claimRequest_Params) RequestTokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SandstormCore_claimRequest_Params) SetRequestToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s SandstormCore_claimRequest_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SandstormCore_claimRequest_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormCore_claimRequest_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SandstormCore_claimRequest_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// SandstormCore_claimRequest_Params_List is a list of SandstormCore_claimRequest_Params.
type SandstormCore_claimRequest_Params_List struct{ capnp.List }

// NewSandstormCore_claimRequest_Params creates a new list of SandstormCore_claimRequest_Params.
func NewSandstormCore_claimRequest_Params_List(s *capnp.Segment, sz int32) (SandstormCore_claimRequest_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormCore_claimRequest_Params_List{l}, err
}

func (s SandstormCore_claimRequest_Params_List) At(i int) SandstormCore_claimRequest_Params {
	return SandstormCore_claimRequest_Params{s.List.Struct(i)}
}

func (s SandstormCore_claimRequest_Params_List) Set(i int, v SandstormCore_claimRequest_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_claimRequest_Params_Promise is a wrapper for a SandstormCore_claimRequest_Params promised by a client call.
type SandstormCore_claimRequest_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_claimRequest_Params_Promise) Struct() (SandstormCore_claimRequest_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_claimRequest_Params{s}, err
}

type SandstormCore_claimRequest_Results struct{ capnp.Struct }

// SandstormCore_claimRequest_Results_TypeID is the unique identifier for the type SandstormCore_claimRequest_Results.
const SandstormCore_claimRequest_Results_TypeID = 0xb91071e3d7b9ab13

func NewSandstormCore_claimRequest_Results(s *capnp.Segment) (SandstormCore_claimRequest_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_claimRequest_Results{st}, err
}

func NewRootSandstormCore_claimRequest_Results(s *capnp.Segment) (SandstormCore_claimRequest_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_claimRequest_Results{st}, err
}

func ReadRootSandstormCore_claimRequest_Results(msg *capnp.Message) (SandstormCore_claimRequest_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_claimRequest_Results{root.Struct()}, err
}

func (s SandstormCore_claimRequest_Results) String() string {
	str, _ := text.Marshal(0xb91071e3d7b9ab13, s.Struct)
	return str
}

func (s SandstormCore_claimRequest_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormCore_claimRequest_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_claimRequest_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormCore_claimRequest_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormCore_claimRequest_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SandstormCore_claimRequest_Results_List is a list of SandstormCore_claimRequest_Results.
type SandstormCore_claimRequest_Results_List struct{ capnp.List }

// NewSandstormCore_claimRequest_Results creates a new list of SandstormCore_claimRequest_Results.
func NewSandstormCore_claimRequest_Results_List(s *capnp.Segment, sz int32) (SandstormCore_claimRequest_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_claimRequest_Results_List{l}, err
}

func (s SandstormCore_claimRequest_Results_List) At(i int) SandstormCore_claimRequest_Results {
	return SandstormCore_claimRequest_Results{s.List.Struct(i)}
}

func (s SandstormCore_claimRequest_Results_List) Set(i int, v SandstormCore_claimRequest_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_claimRequest_Results_Promise is a wrapper for a SandstormCore_claimRequest_Results promised by a client call.
type SandstormCore_claimRequest_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_claimRequest_Results_Promise) Struct() (SandstormCore_claimRequest_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_claimRequest_Results{s}, err
}

func (p SandstormCore_claimRequest_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormCore_backgroundActivity_Params struct{ capnp.Struct }

// SandstormCore_backgroundActivity_Params_TypeID is the unique identifier for the type SandstormCore_backgroundActivity_Params.
const SandstormCore_backgroundActivity_Params_TypeID = 0x938e798cc0e3d6ac

func NewSandstormCore_backgroundActivity_Params(s *capnp.Segment) (SandstormCore_backgroundActivity_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_backgroundActivity_Params{st}, err
}

func NewRootSandstormCore_backgroundActivity_Params(s *capnp.Segment) (SandstormCore_backgroundActivity_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_backgroundActivity_Params{st}, err
}

func ReadRootSandstormCore_backgroundActivity_Params(msg *capnp.Message) (SandstormCore_backgroundActivity_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_backgroundActivity_Params{root.Struct()}, err
}

func (s SandstormCore_backgroundActivity_Params) String() string {
	str, _ := text.Marshal(0x938e798cc0e3d6ac, s.Struct)
	return str
}

func (s SandstormCore_backgroundActivity_Params) Event() (activity.ActivityEvent, error) {
	p, err := s.Struct.Ptr(0)
	return activity.ActivityEvent{Struct: p.Struct()}, err
}

func (s SandstormCore_backgroundActivity_Params) HasEvent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_backgroundActivity_Params) SetEvent(v activity.ActivityEvent) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEvent sets the event field to a newly
// allocated activity.ActivityEvent struct, preferring placement in s's segment.
func (s SandstormCore_backgroundActivity_Params) NewEvent() (activity.ActivityEvent, error) {
	ss, err := activity.NewActivityEvent(s.Struct.Segment())
	if err != nil {
		return activity.ActivityEvent{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// SandstormCore_backgroundActivity_Params_List is a list of SandstormCore_backgroundActivity_Params.
type SandstormCore_backgroundActivity_Params_List struct{ capnp.List }

// NewSandstormCore_backgroundActivity_Params creates a new list of SandstormCore_backgroundActivity_Params.
func NewSandstormCore_backgroundActivity_Params_List(s *capnp.Segment, sz int32) (SandstormCore_backgroundActivity_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_backgroundActivity_Params_List{l}, err
}

func (s SandstormCore_backgroundActivity_Params_List) At(i int) SandstormCore_backgroundActivity_Params {
	return SandstormCore_backgroundActivity_Params{s.List.Struct(i)}
}

func (s SandstormCore_backgroundActivity_Params_List) Set(i int, v SandstormCore_backgroundActivity_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_backgroundActivity_Params_Promise is a wrapper for a SandstormCore_backgroundActivity_Params promised by a client call.
type SandstormCore_backgroundActivity_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_backgroundActivity_Params_Promise) Struct() (SandstormCore_backgroundActivity_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_backgroundActivity_Params{s}, err
}

func (p SandstormCore_backgroundActivity_Params_Promise) Event() activity.ActivityEvent_Promise {
	return activity.ActivityEvent_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type SandstormCore_backgroundActivity_Results struct{ capnp.Struct }

// SandstormCore_backgroundActivity_Results_TypeID is the unique identifier for the type SandstormCore_backgroundActivity_Results.
const SandstormCore_backgroundActivity_Results_TypeID = 0x9d87019c48640d21

func NewSandstormCore_backgroundActivity_Results(s *capnp.Segment) (SandstormCore_backgroundActivity_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_backgroundActivity_Results{st}, err
}

func NewRootSandstormCore_backgroundActivity_Results(s *capnp.Segment) (SandstormCore_backgroundActivity_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_backgroundActivity_Results{st}, err
}

func ReadRootSandstormCore_backgroundActivity_Results(msg *capnp.Message) (SandstormCore_backgroundActivity_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_backgroundActivity_Results{root.Struct()}, err
}

func (s SandstormCore_backgroundActivity_Results) String() string {
	str, _ := text.Marshal(0x9d87019c48640d21, s.Struct)
	return str
}

// SandstormCore_backgroundActivity_Results_List is a list of SandstormCore_backgroundActivity_Results.
type SandstormCore_backgroundActivity_Results_List struct{ capnp.List }

// NewSandstormCore_backgroundActivity_Results creates a new list of SandstormCore_backgroundActivity_Results.
func NewSandstormCore_backgroundActivity_Results_List(s *capnp.Segment, sz int32) (SandstormCore_backgroundActivity_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_backgroundActivity_Results_List{l}, err
}

func (s SandstormCore_backgroundActivity_Results_List) At(i int) SandstormCore_backgroundActivity_Results {
	return SandstormCore_backgroundActivity_Results{s.List.Struct(i)}
}

func (s SandstormCore_backgroundActivity_Results_List) Set(i int, v SandstormCore_backgroundActivity_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_backgroundActivity_Results_Promise is a wrapper for a SandstormCore_backgroundActivity_Results promised by a client call.
type SandstormCore_backgroundActivity_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_backgroundActivity_Results_Promise) Struct() (SandstormCore_backgroundActivity_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_backgroundActivity_Results{s}, err
}

type SandstormCore_reportGrainSize_Params struct{ capnp.Struct }

// SandstormCore_reportGrainSize_Params_TypeID is the unique identifier for the type SandstormCore_reportGrainSize_Params.
const SandstormCore_reportGrainSize_Params_TypeID = 0xaec15e35d479f4f3

func NewSandstormCore_reportGrainSize_Params(s *capnp.Segment) (SandstormCore_reportGrainSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return SandstormCore_reportGrainSize_Params{st}, err
}

func NewRootSandstormCore_reportGrainSize_Params(s *capnp.Segment) (SandstormCore_reportGrainSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return SandstormCore_reportGrainSize_Params{st}, err
}

func ReadRootSandstormCore_reportGrainSize_Params(msg *capnp.Message) (SandstormCore_reportGrainSize_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_reportGrainSize_Params{root.Struct()}, err
}

func (s SandstormCore_reportGrainSize_Params) String() string {
	str, _ := text.Marshal(0xaec15e35d479f4f3, s.Struct)
	return str
}

func (s SandstormCore_reportGrainSize_Params) Bytes() uint64 {
	return s.Struct.Uint64(0)
}

func (s SandstormCore_reportGrainSize_Params) SetBytes(v uint64) {
	s.Struct.SetUint64(0, v)
}

// SandstormCore_reportGrainSize_Params_List is a list of SandstormCore_reportGrainSize_Params.
type SandstormCore_reportGrainSize_Params_List struct{ capnp.List }

// NewSandstormCore_reportGrainSize_Params creates a new list of SandstormCore_reportGrainSize_Params.
func NewSandstormCore_reportGrainSize_Params_List(s *capnp.Segment, sz int32) (SandstormCore_reportGrainSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return SandstormCore_reportGrainSize_Params_List{l}, err
}

func (s SandstormCore_reportGrainSize_Params_List) At(i int) SandstormCore_reportGrainSize_Params {
	return SandstormCore_reportGrainSize_Params{s.List.Struct(i)}
}

func (s SandstormCore_reportGrainSize_Params_List) Set(i int, v SandstormCore_reportGrainSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_reportGrainSize_Params_Promise is a wrapper for a SandstormCore_reportGrainSize_Params promised by a client call.
type SandstormCore_reportGrainSize_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_reportGrainSize_Params_Promise) Struct() (SandstormCore_reportGrainSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_reportGrainSize_Params{s}, err
}

type SandstormCore_reportGrainSize_Results struct{ capnp.Struct }

// SandstormCore_reportGrainSize_Results_TypeID is the unique identifier for the type SandstormCore_reportGrainSize_Results.
const SandstormCore_reportGrainSize_Results_TypeID = 0xc43d5a1430e113ca

func NewSandstormCore_reportGrainSize_Results(s *capnp.Segment) (SandstormCore_reportGrainSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_reportGrainSize_Results{st}, err
}

func NewRootSandstormCore_reportGrainSize_Results(s *capnp.Segment) (SandstormCore_reportGrainSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormCore_reportGrainSize_Results{st}, err
}

func ReadRootSandstormCore_reportGrainSize_Results(msg *capnp.Message) (SandstormCore_reportGrainSize_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_reportGrainSize_Results{root.Struct()}, err
}

func (s SandstormCore_reportGrainSize_Results) String() string {
	str, _ := text.Marshal(0xc43d5a1430e113ca, s.Struct)
	return str
}

// SandstormCore_reportGrainSize_Results_List is a list of SandstormCore_reportGrainSize_Results.
type SandstormCore_reportGrainSize_Results_List struct{ capnp.List }

// NewSandstormCore_reportGrainSize_Results creates a new list of SandstormCore_reportGrainSize_Results.
func NewSandstormCore_reportGrainSize_Results_List(s *capnp.Segment, sz int32) (SandstormCore_reportGrainSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormCore_reportGrainSize_Results_List{l}, err
}

func (s SandstormCore_reportGrainSize_Results_List) At(i int) SandstormCore_reportGrainSize_Results {
	return SandstormCore_reportGrainSize_Results{s.List.Struct(i)}
}

func (s SandstormCore_reportGrainSize_Results_List) Set(i int, v SandstormCore_reportGrainSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_reportGrainSize_Results_Promise is a wrapper for a SandstormCore_reportGrainSize_Results promised by a client call.
type SandstormCore_reportGrainSize_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_reportGrainSize_Results_Promise) Struct() (SandstormCore_reportGrainSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_reportGrainSize_Results{s}, err
}

type SandstormCore_getIdentityId_Params struct{ capnp.Struct }

// SandstormCore_getIdentityId_Params_TypeID is the unique identifier for the type SandstormCore_getIdentityId_Params.
const SandstormCore_getIdentityId_Params_TypeID = 0xf59063f154adea97

func NewSandstormCore_getIdentityId_Params(s *capnp.Segment) (SandstormCore_getIdentityId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getIdentityId_Params{st}, err
}

func NewRootSandstormCore_getIdentityId_Params(s *capnp.Segment) (SandstormCore_getIdentityId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getIdentityId_Params{st}, err
}

func ReadRootSandstormCore_getIdentityId_Params(msg *capnp.Message) (SandstormCore_getIdentityId_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCore_getIdentityId_Params{root.Struct()}, err
}

func (s SandstormCore_getIdentityId_Params) String() string {
	str, _ := text.Marshal(0xf59063f154adea97, s.Struct)
	return str
}

func (s SandstormCore_getIdentityId_Params) Identity() identity.Identity {
	p, _ := s.Struct.Ptr(0)
	return identity.Identity{Client: p.Interface().Client()}
}

func (s SandstormCore_getIdentityId_Params) HasIdentity() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_getIdentityId_Params) SetIdentity(v identity.Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormCore_getIdentityId_Params_List is a list of SandstormCore_getIdentityId_Params.
type SandstormCore_getIdentityId_Params_List struct{ capnp.List }

// NewSandstormCore_getIdentityId_Params creates a new list of SandstormCore_getIdentityId_Params.
func NewSandstormCore_getIdentityId_Params_List(s *capnp.Segment, sz int32) (SandstormCore_getIdentityId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_getIdentityId_Params_List{l}, err
}

func (s SandstormCore_getIdentityId_Params_List) At(i int) SandstormCore_getIdentityId_Params {
	return SandstormCore_getIdentityId_Params{s.List.Struct(i)}
}

func (s SandstormCore_getIdentityId_Params_List) Set(i int, v SandstormCore_getIdentityId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_getIdentityId_Params_Promise is a wrapper for a SandstormCore_getIdentityId_Params promised by a client call.
type SandstormCore_getIdentityId_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_getIdentityId_Params_Promise) Struct() (SandstormCore_getIdentityId_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_getIdentityId_Params{s}, err
}

func (p SandstormCore_getIdentityId_Params_Promise) Identity() identity.Identity {
	return identity.Identity{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormCore_getIdentityId_Results struct{ capnp.Struct }

// SandstormCore_getIdentityId_Results_TypeID is the unique identifier for the type SandstormCore_getIdentityId_Results.
const SandstormCore_getIdentityId_Results_TypeID = 0xcaa1479a3b9c719b

func NewSandstormCore_getIdentityId_Results(s *capnp.Segment) (SandstormCore_getIdentityId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getIdentityId_Results{st}, err
}

func NewRootSandstormCore_getIdentityId_Results(s *capnp.Segment) (SandstormCore_getIdentityId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCore_getIdentityId_Results{st}, err
}

func ReadRootSandstormCore_getIdentityId_Results(msg *capnp.Message) (SandstormCore_getIdentityId_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCore_getIdentityId_Results{root.Struct()}, err
}

func (s SandstormCore_getIdentityId_Results) String() string {
	str, _ := text.Marshal(0xcaa1479a3b9c719b, s.Struct)
	return str
}

func (s SandstormCore_getIdentityId_Results) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormCore_getIdentityId_Results) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_getIdentityId_Results) SetId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormCore_getIdentityId_Results_List is a list of SandstormCore_getIdentityId_Results.
type SandstormCore_getIdentityId_Results_List struct{ capnp.List }

// NewSandstormCore_getIdentityId_Results creates a new list of SandstormCore_getIdentityId_Results.
func NewSandstormCore_getIdentityId_Results_List(s *capnp.Segment, sz int32) (SandstormCore_getIdentityId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCore_getIdentityId_Results_List{l}, err
}

func (s SandstormCore_getIdentityId_Results_List) At(i int) SandstormCore_getIdentityId_Results {
	return SandstormCore_getIdentityId_Results{s.List.Struct(i)}
}

func (s SandstormCore_getIdentityId_Results_List) Set(i int, v SandstormCore_getIdentityId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCore_getIdentityId_Results_Promise is a wrapper for a SandstormCore_getIdentityId_Results promised by a client call.
type SandstormCore_getIdentityId_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCore_getIdentityId_Results_Promise) Struct() (SandstormCore_getIdentityId_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCore_getIdentityId_Results{s}, err
}

type MembraneRequirement struct{ capnp.Struct }
type MembraneRequirement_permissionsHeld MembraneRequirement
type MembraneRequirement_Which uint16

const (
	MembraneRequirement_Which_tokenValid      MembraneRequirement_Which = 0
	MembraneRequirement_Which_permissionsHeld MembraneRequirement_Which = 1
	MembraneRequirement_Which_userIsAdmin     MembraneRequirement_Which = 2
)

func (w MembraneRequirement_Which) String() string {
	const s = "tokenValidpermissionsHelduserIsAdmin"
	switch w {
	case MembraneRequirement_Which_tokenValid:
		return s[0:10]
	case MembraneRequirement_Which_permissionsHeld:
		return s[10:25]
	case MembraneRequirement_Which_userIsAdmin:
		return s[25:36]

	}
	return "MembraneRequirement_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type MembraneRequirement_permissionsHeld_Which uint16

const (
	MembraneRequirement_permissionsHeld_Which_identityId MembraneRequirement_permissionsHeld_Which = 0
	MembraneRequirement_permissionsHeld_Which_tokenId    MembraneRequirement_permissionsHeld_Which = 1
)

func (w MembraneRequirement_permissionsHeld_Which) String() string {
	const s = "identityIdtokenId"
	switch w {
	case MembraneRequirement_permissionsHeld_Which_identityId:
		return s[0:10]
	case MembraneRequirement_permissionsHeld_Which_tokenId:
		return s[10:17]

	}
	return "MembraneRequirement_permissionsHeld_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// MembraneRequirement_TypeID is the unique identifier for the type MembraneRequirement.
const MembraneRequirement_TypeID = 0x918db9a721f13886

func NewMembraneRequirement(s *capnp.Segment) (MembraneRequirement, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return MembraneRequirement{st}, err
}

func NewRootMembraneRequirement(s *capnp.Segment) (MembraneRequirement, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return MembraneRequirement{st}, err
}

func ReadRootMembraneRequirement(msg *capnp.Message) (MembraneRequirement, error) {
	root, err := msg.RootPtr()
	return MembraneRequirement{root.Struct()}, err
}

func (s MembraneRequirement) String() string {
	str, _ := text.Marshal(0x918db9a721f13886, s.Struct)
	return str
}

func (s MembraneRequirement) Which() MembraneRequirement_Which {
	return MembraneRequirement_Which(s.Struct.Uint16(0))
}
func (s MembraneRequirement) TokenValid() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s MembraneRequirement) HasTokenValid() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement) TokenValidBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s MembraneRequirement) SetTokenValid(v string) error {
	s.Struct.SetUint16(0, 0)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s MembraneRequirement) PermissionsHeld() MembraneRequirement_permissionsHeld {
	return MembraneRequirement_permissionsHeld(s)
}

func (s MembraneRequirement) SetPermissionsHeld() {
	s.Struct.SetUint16(0, 1)
}

func (s MembraneRequirement_permissionsHeld) Which() MembraneRequirement_permissionsHeld_Which {
	return MembraneRequirement_permissionsHeld_Which(s.Struct.Uint16(2))
}
func (s MembraneRequirement_permissionsHeld) IdentityId() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s MembraneRequirement_permissionsHeld) HasIdentityId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) IdentityIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s MembraneRequirement_permissionsHeld) SetIdentityId(v string) error {
	s.Struct.SetUint16(2, 0)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s MembraneRequirement_permissionsHeld) TokenId() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s MembraneRequirement_permissionsHeld) HasTokenId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s MembraneRequirement_permissionsHeld) SetTokenId(v string) error {
	s.Struct.SetUint16(2, 1)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s MembraneRequirement_permissionsHeld) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s MembraneRequirement_permissionsHeld) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s MembraneRequirement_permissionsHeld) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s MembraneRequirement_permissionsHeld) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.BitList{List: p.List()}, err
}

func (s MembraneRequirement_permissionsHeld) HasPermissions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s MembraneRequirement_permissionsHeld) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s MembraneRequirement_permissionsHeld) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s MembraneRequirement_permissionsHeld) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s MembraneRequirement_permissionsHeld) SetUserId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s MembraneRequirement) UserIsAdmin() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s MembraneRequirement) HasUserIsAdmin() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement) UserIsAdminBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s MembraneRequirement) SetUserIsAdmin(v string) error {
	s.Struct.SetUint16(0, 2)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// MembraneRequirement_List is a list of MembraneRequirement.
type MembraneRequirement_List struct{ capnp.List }

// NewMembraneRequirement creates a new list of MembraneRequirement.
func NewMembraneRequirement_List(s *capnp.Segment, sz int32) (MembraneRequirement_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return MembraneRequirement_List{l}, err
}

func (s MembraneRequirement_List) At(i int) MembraneRequirement {
	return MembraneRequirement{s.List.Struct(i)}
}

func (s MembraneRequirement_List) Set(i int, v MembraneRequirement) error {
	return s.List.SetStruct(i, v.Struct)
}

// MembraneRequirement_Promise is a wrapper for a MembraneRequirement promised by a client call.
type MembraneRequirement_Promise struct{ *capnp.Pipeline }

func (p MembraneRequirement_Promise) Struct() (MembraneRequirement, error) {
	s, err := p.Pipeline.Struct()
	return MembraneRequirement{s}, err
}

func (p MembraneRequirement_Promise) PermissionsHeld() MembraneRequirement_permissionsHeld_Promise {
	return MembraneRequirement_permissionsHeld_Promise{p.Pipeline}
}

// MembraneRequirement_permissionsHeld_Promise is a wrapper for a MembraneRequirement_permissionsHeld promised by a client call.
type MembraneRequirement_permissionsHeld_Promise struct{ *capnp.Pipeline }

func (p MembraneRequirement_permissionsHeld_Promise) Struct() (MembraneRequirement_permissionsHeld, error) {
	s, err := p.Pipeline.Struct()
	return MembraneRequirement_permissionsHeld{s}, err
}

type SystemPersistent struct{ Client capnp.Client }

func (c SystemPersistent) AddRequirements(ctx context.Context, params func(SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(SystemPersistent_addRequirements_Params{Struct: s}) }
	}
	return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SystemPersistent) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type SystemPersistent_Server interface {
	AddRequirements(SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func SystemPersistent_ServerToClient(s SystemPersistent_Server) SystemPersistent {
	c, _ := s.(server.Closer)
	return SystemPersistent{Client: server.New(SystemPersistent_Methods(nil, s), c)}
}

func SystemPersistent_Methods(methods []server.Method, s SystemPersistent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SystemPersistent_addRequirements{c, opts, SystemPersistent_addRequirements_Params{Struct: p}, SystemPersistent_addRequirements_Results{Struct: r}}
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

// SystemPersistent_addRequirements holds the arguments for a server call to SystemPersistent.addRequirements.
type SystemPersistent_addRequirements struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SystemPersistent_addRequirements_Params
	Results SystemPersistent_addRequirements_Results
}

type SystemPersistent_addRequirements_Params struct{ capnp.Struct }

// SystemPersistent_addRequirements_Params_TypeID is the unique identifier for the type SystemPersistent_addRequirements_Params.
const SystemPersistent_addRequirements_Params_TypeID = 0xbb5eb0bde1481587

func NewSystemPersistent_addRequirements_Params(s *capnp.Segment) (SystemPersistent_addRequirements_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SystemPersistent_addRequirements_Params{st}, err
}

func NewRootSystemPersistent_addRequirements_Params(s *capnp.Segment) (SystemPersistent_addRequirements_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SystemPersistent_addRequirements_Params{st}, err
}

func ReadRootSystemPersistent_addRequirements_Params(msg *capnp.Message) (SystemPersistent_addRequirements_Params, error) {
	root, err := msg.RootPtr()
	return SystemPersistent_addRequirements_Params{root.Struct()}, err
}

func (s SystemPersistent_addRequirements_Params) String() string {
	str, _ := text.Marshal(0xbb5eb0bde1481587, s.Struct)
	return str
}

func (s SystemPersistent_addRequirements_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(0)
	return MembraneRequirement_List{List: p.List()}, err
}

func (s SystemPersistent_addRequirements_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SystemPersistent_addRequirements_Params) SetRequirements(v MembraneRequirement_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewRequirements sets the requirements field to a newly
// allocated MembraneRequirement_List, preferring placement in s's segment.
func (s SystemPersistent_addRequirements_Params) NewRequirements(n int32) (MembraneRequirement_List, error) {
	l, err := NewMembraneRequirement_List(s.Struct.Segment(), n)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// SystemPersistent_addRequirements_Params_List is a list of SystemPersistent_addRequirements_Params.
type SystemPersistent_addRequirements_Params_List struct{ capnp.List }

// NewSystemPersistent_addRequirements_Params creates a new list of SystemPersistent_addRequirements_Params.
func NewSystemPersistent_addRequirements_Params_List(s *capnp.Segment, sz int32) (SystemPersistent_addRequirements_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SystemPersistent_addRequirements_Params_List{l}, err
}

func (s SystemPersistent_addRequirements_Params_List) At(i int) SystemPersistent_addRequirements_Params {
	return SystemPersistent_addRequirements_Params{s.List.Struct(i)}
}

func (s SystemPersistent_addRequirements_Params_List) Set(i int, v SystemPersistent_addRequirements_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SystemPersistent_addRequirements_Params_Promise is a wrapper for a SystemPersistent_addRequirements_Params promised by a client call.
type SystemPersistent_addRequirements_Params_Promise struct{ *capnp.Pipeline }

func (p SystemPersistent_addRequirements_Params_Promise) Struct() (SystemPersistent_addRequirements_Params, error) {
	s, err := p.Pipeline.Struct()
	return SystemPersistent_addRequirements_Params{s}, err
}

type SystemPersistent_addRequirements_Results struct{ capnp.Struct }

// SystemPersistent_addRequirements_Results_TypeID is the unique identifier for the type SystemPersistent_addRequirements_Results.
const SystemPersistent_addRequirements_Results_TypeID = 0x8488d5d569f6cffe

func NewSystemPersistent_addRequirements_Results(s *capnp.Segment) (SystemPersistent_addRequirements_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SystemPersistent_addRequirements_Results{st}, err
}

func NewRootSystemPersistent_addRequirements_Results(s *capnp.Segment) (SystemPersistent_addRequirements_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SystemPersistent_addRequirements_Results{st}, err
}

func ReadRootSystemPersistent_addRequirements_Results(msg *capnp.Message) (SystemPersistent_addRequirements_Results, error) {
	root, err := msg.RootPtr()
	return SystemPersistent_addRequirements_Results{root.Struct()}, err
}

func (s SystemPersistent_addRequirements_Results) String() string {
	str, _ := text.Marshal(0x8488d5d569f6cffe, s.Struct)
	return str
}

func (s SystemPersistent_addRequirements_Results) Cap() SystemPersistent {
	p, _ := s.Struct.Ptr(0)
	return SystemPersistent{Client: p.Interface().Client()}
}

func (s SystemPersistent_addRequirements_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SystemPersistent_addRequirements_Results) SetCap(v SystemPersistent) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SystemPersistent_addRequirements_Results_List is a list of SystemPersistent_addRequirements_Results.
type SystemPersistent_addRequirements_Results_List struct{ capnp.List }

// NewSystemPersistent_addRequirements_Results creates a new list of SystemPersistent_addRequirements_Results.
func NewSystemPersistent_addRequirements_Results_List(s *capnp.Segment, sz int32) (SystemPersistent_addRequirements_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SystemPersistent_addRequirements_Results_List{l}, err
}

func (s SystemPersistent_addRequirements_Results_List) At(i int) SystemPersistent_addRequirements_Results {
	return SystemPersistent_addRequirements_Results{s.List.Struct(i)}
}

func (s SystemPersistent_addRequirements_Results_List) Set(i int, v SystemPersistent_addRequirements_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SystemPersistent_addRequirements_Results_Promise is a wrapper for a SystemPersistent_addRequirements_Results promised by a client call.
type SystemPersistent_addRequirements_Results_Promise struct{ *capnp.Pipeline }

func (p SystemPersistent_addRequirements_Results_Promise) Struct() (SystemPersistent_addRequirements_Results, error) {
	s, err := p.Pipeline.Struct()
	return SystemPersistent_addRequirements_Results{s}, err
}

func (p SystemPersistent_addRequirements_Results_Promise) Cap() SystemPersistent {
	return SystemPersistent{Client: p.Pipeline.GetPipeline(0).Client()}
}

type PersistentHandle struct{ Client capnp.Client }

func (c PersistentHandle) AddRequirements(ctx context.Context, params func(SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(SystemPersistent_addRequirements_Params{Struct: s}) }
	}
	return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentHandle) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type PersistentHandle_Server interface {
	AddRequirements(SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentHandle_ServerToClient(s PersistentHandle_Server) PersistentHandle {
	c, _ := s.(server.Closer)
	return PersistentHandle{Client: server.New(PersistentHandle_Methods(nil, s), c)}
}

func PersistentHandle_Methods(methods []server.Method, s PersistentHandle_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SystemPersistent_addRequirements{c, opts, SystemPersistent_addRequirements_Params{Struct: p}, SystemPersistent_addRequirements_Results{Struct: r}}
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

type PersistentOngoingNotification struct{ Client capnp.Client }

func (c PersistentOngoingNotification) AddRequirements(ctx context.Context, params func(SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) SystemPersistent_addRequirements_Results_Promise {
	if c.Client == nil {
		return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(SystemPersistent_addRequirements_Params{Struct: s}) }
	}
	return SystemPersistent_addRequirements_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentOngoingNotification) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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
func (c PersistentOngoingNotification) Cancel(ctx context.Context, params func(activity.OngoingNotification_cancel_Params) error, opts ...capnp.CallOption) activity.OngoingNotification_cancel_Results_Promise {
	if c.Client == nil {
		return activity.OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xfe851ddbb88940cd,
			MethodID:      0,
			InterfaceName: "activity.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(activity.OngoingNotification_cancel_Params{Struct: s}) }
	}
	return activity.OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type PersistentOngoingNotification_Server interface {
	AddRequirements(SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error

	Cancel(activity.OngoingNotification_cancel) error
}

func PersistentOngoingNotification_ServerToClient(s PersistentOngoingNotification_Server) PersistentOngoingNotification {
	c, _ := s.(server.Closer)
	return PersistentOngoingNotification{Client: server.New(PersistentOngoingNotification_Methods(nil, s), c)}
}

func PersistentOngoingNotification_Methods(methods []server.Method, s PersistentOngoingNotification_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc38cedd77cbed5b4,
			MethodID:      0,
			InterfaceName: "supervisor.capnp:SystemPersistent",
			MethodName:    "addRequirements",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SystemPersistent_addRequirements{c, opts, SystemPersistent_addRequirements_Params{Struct: p}, SystemPersistent_addRequirements_Results{Struct: r}}
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

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfe851ddbb88940cd,
			MethodID:      0,
			InterfaceName: "activity.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := activity.OngoingNotification_cancel{c, opts, activity.OngoingNotification_cancel_Params{Struct: p}, activity.OngoingNotification_cancel_Results{Struct: r}}
			return s.Cancel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

type DenormalizedGrainMetadata struct{ capnp.Struct }
type DenormalizedGrainMetadata_icon DenormalizedGrainMetadata
type DenormalizedGrainMetadata_Which uint16

const (
	DenormalizedGrainMetadata_Which_icon  DenormalizedGrainMetadata_Which = 0
	DenormalizedGrainMetadata_Which_appId DenormalizedGrainMetadata_Which = 1
)

func (w DenormalizedGrainMetadata_Which) String() string {
	const s = "iconappId"
	switch w {
	case DenormalizedGrainMetadata_Which_icon:
		return s[0:4]
	case DenormalizedGrainMetadata_Which_appId:
		return s[4:9]

	}
	return "DenormalizedGrainMetadata_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// DenormalizedGrainMetadata_TypeID is the unique identifier for the type DenormalizedGrainMetadata.
const DenormalizedGrainMetadata_TypeID = 0xbdd9bea5585df6c5

func NewDenormalizedGrainMetadata(s *capnp.Segment) (DenormalizedGrainMetadata, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return DenormalizedGrainMetadata{st}, err
}

func NewRootDenormalizedGrainMetadata(s *capnp.Segment) (DenormalizedGrainMetadata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return DenormalizedGrainMetadata{st}, err
}

func ReadRootDenormalizedGrainMetadata(msg *capnp.Message) (DenormalizedGrainMetadata, error) {
	root, err := msg.RootPtr()
	return DenormalizedGrainMetadata{root.Struct()}, err
}

func (s DenormalizedGrainMetadata) String() string {
	str, _ := text.Marshal(0xbdd9bea5585df6c5, s.Struct)
	return str
}

func (s DenormalizedGrainMetadata) Which() DenormalizedGrainMetadata_Which {
	return DenormalizedGrainMetadata_Which(s.Struct.Uint16(0))
}
func (s DenormalizedGrainMetadata) AppTitle() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s DenormalizedGrainMetadata) HasAppTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DenormalizedGrainMetadata) SetAppTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAppTitle sets the appTitle field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s DenormalizedGrainMetadata) NewAppTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s DenormalizedGrainMetadata) Icon() DenormalizedGrainMetadata_icon {
	return DenormalizedGrainMetadata_icon(s)
}

func (s DenormalizedGrainMetadata) SetIcon() {
	s.Struct.SetUint16(0, 0)
}

func (s DenormalizedGrainMetadata_icon) Format() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s DenormalizedGrainMetadata_icon) HasFormat() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DenormalizedGrainMetadata_icon) FormatBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s DenormalizedGrainMetadata_icon) SetFormat(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s DenormalizedGrainMetadata_icon) AssetId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s DenormalizedGrainMetadata_icon) HasAssetId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s DenormalizedGrainMetadata_icon) AssetIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s DenormalizedGrainMetadata_icon) SetAssetId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s DenormalizedGrainMetadata_icon) AssetId2xDpi() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s DenormalizedGrainMetadata_icon) HasAssetId2xDpi() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s DenormalizedGrainMetadata_icon) AssetId2xDpiBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s DenormalizedGrainMetadata_icon) SetAssetId2xDpi(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s DenormalizedGrainMetadata) AppId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s DenormalizedGrainMetadata) HasAppId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DenormalizedGrainMetadata) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s DenormalizedGrainMetadata) SetAppId(v string) error {
	s.Struct.SetUint16(0, 1)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// DenormalizedGrainMetadata_List is a list of DenormalizedGrainMetadata.
type DenormalizedGrainMetadata_List struct{ capnp.List }

// NewDenormalizedGrainMetadata creates a new list of DenormalizedGrainMetadata.
func NewDenormalizedGrainMetadata_List(s *capnp.Segment, sz int32) (DenormalizedGrainMetadata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return DenormalizedGrainMetadata_List{l}, err
}

func (s DenormalizedGrainMetadata_List) At(i int) DenormalizedGrainMetadata {
	return DenormalizedGrainMetadata{s.List.Struct(i)}
}

func (s DenormalizedGrainMetadata_List) Set(i int, v DenormalizedGrainMetadata) error {
	return s.List.SetStruct(i, v.Struct)
}

// DenormalizedGrainMetadata_Promise is a wrapper for a DenormalizedGrainMetadata promised by a client call.
type DenormalizedGrainMetadata_Promise struct{ *capnp.Pipeline }

func (p DenormalizedGrainMetadata_Promise) Struct() (DenormalizedGrainMetadata, error) {
	s, err := p.Pipeline.Struct()
	return DenormalizedGrainMetadata{s}, err
}

func (p DenormalizedGrainMetadata_Promise) AppTitle() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p DenormalizedGrainMetadata_Promise) Icon() DenormalizedGrainMetadata_icon_Promise {
	return DenormalizedGrainMetadata_icon_Promise{p.Pipeline}
}

// DenormalizedGrainMetadata_icon_Promise is a wrapper for a DenormalizedGrainMetadata_icon promised by a client call.
type DenormalizedGrainMetadata_icon_Promise struct{ *capnp.Pipeline }

func (p DenormalizedGrainMetadata_icon_Promise) Struct() (DenormalizedGrainMetadata_icon, error) {
	s, err := p.Pipeline.Struct()
	return DenormalizedGrainMetadata_icon{s}, err
}

type ApiTokenOwner struct{ capnp.Struct }
type ApiTokenOwner_grain ApiTokenOwner
type ApiTokenOwner_clientPowerboxRequest ApiTokenOwner
type ApiTokenOwner_clientPowerboxOffer ApiTokenOwner
type ApiTokenOwner_user ApiTokenOwner
type ApiTokenOwner_Which uint16

const (
	ApiTokenOwner_Which_webkey                ApiTokenOwner_Which = 0
	ApiTokenOwner_Which_grain                 ApiTokenOwner_Which = 1
	ApiTokenOwner_Which_clientPowerboxRequest ApiTokenOwner_Which = 5
	ApiTokenOwner_Which_clientPowerboxOffer   ApiTokenOwner_Which = 6
	ApiTokenOwner_Which_internet              ApiTokenOwner_Which = 2
	ApiTokenOwner_Which_frontend              ApiTokenOwner_Which = 3
	ApiTokenOwner_Which_user                  ApiTokenOwner_Which = 4
)

func (w ApiTokenOwner_Which) String() string {
	const s = "webkeygrainclientPowerboxRequestclientPowerboxOfferinternetfrontenduser"
	switch w {
	case ApiTokenOwner_Which_webkey:
		return s[0:6]
	case ApiTokenOwner_Which_grain:
		return s[6:11]
	case ApiTokenOwner_Which_clientPowerboxRequest:
		return s[11:32]
	case ApiTokenOwner_Which_clientPowerboxOffer:
		return s[32:51]
	case ApiTokenOwner_Which_internet:
		return s[51:59]
	case ApiTokenOwner_Which_frontend:
		return s[59:67]
	case ApiTokenOwner_Which_user:
		return s[67:71]

	}
	return "ApiTokenOwner_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// ApiTokenOwner_TypeID is the unique identifier for the type ApiTokenOwner.
const ApiTokenOwner_TypeID = 0xda970537e2a8a9a9

func NewApiTokenOwner(s *capnp.Segment) (ApiTokenOwner, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return ApiTokenOwner{st}, err
}

func NewRootApiTokenOwner(s *capnp.Segment) (ApiTokenOwner, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return ApiTokenOwner{st}, err
}

func ReadRootApiTokenOwner(msg *capnp.Message) (ApiTokenOwner, error) {
	root, err := msg.RootPtr()
	return ApiTokenOwner{root.Struct()}, err
}

func (s ApiTokenOwner) String() string {
	str, _ := text.Marshal(0xda970537e2a8a9a9, s.Struct)
	return str
}

func (s ApiTokenOwner) Which() ApiTokenOwner_Which {
	return ApiTokenOwner_Which(s.Struct.Uint16(0))
}
func (s ApiTokenOwner) SetWebkey() {
	s.Struct.SetUint16(0, 0)

}

func (s ApiTokenOwner) Grain() ApiTokenOwner_grain { return ApiTokenOwner_grain(s) }

func (s ApiTokenOwner) SetGrain() {
	s.Struct.SetUint16(0, 1)
}

func (s ApiTokenOwner_grain) GrainId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiTokenOwner_grain) HasGrainId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_grain) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiTokenOwner_grain) SaveLabel() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s ApiTokenOwner_grain) HasSaveLabel() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) SetSaveLabel(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewSaveLabel sets the saveLabel field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ApiTokenOwner_grain) NewSaveLabel() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s ApiTokenOwner_grain) IntroducerIdentity() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s ApiTokenOwner_grain) HasIntroducerIdentity() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) IntroducerIdentityBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_grain) SetIntroducerIdentity(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s ApiTokenOwner_grain) IntroducerUser() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s ApiTokenOwner_grain) HasIntroducerUser() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) IntroducerUserBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_grain) SetIntroducerUser(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s ApiTokenOwner) ClientPowerboxRequest() ApiTokenOwner_clientPowerboxRequest {
	return ApiTokenOwner_clientPowerboxRequest(s)
}

func (s ApiTokenOwner) SetClientPowerboxRequest() {
	s.Struct.SetUint16(0, 5)
}

func (s ApiTokenOwner_clientPowerboxRequest) SessionId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) HasSessionId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_clientPowerboxRequest) SessionIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) SetSessionId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s ApiTokenOwner_clientPowerboxRequest) GrainId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) HasGrainId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_clientPowerboxRequest) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiTokenOwner_clientPowerboxRequest) IntroducerIdentity() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) HasIntroducerIdentity() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_clientPowerboxRequest) IntroducerIdentityBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_clientPowerboxRequest) SetIntroducerIdentity(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s ApiTokenOwner) ClientPowerboxOffer() ApiTokenOwner_clientPowerboxOffer {
	return ApiTokenOwner_clientPowerboxOffer(s)
}

func (s ApiTokenOwner) SetClientPowerboxOffer() {
	s.Struct.SetUint16(0, 6)
}

func (s ApiTokenOwner_clientPowerboxOffer) SessionId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiTokenOwner_clientPowerboxOffer) HasSessionId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_clientPowerboxOffer) SessionIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_clientPowerboxOffer) SetSessionId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiTokenOwner) Internet() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s ApiTokenOwner) HasInternet() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner) InternetPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s ApiTokenOwner) SetInternet(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPointer(0, v)
}

func (s ApiTokenOwner) SetInternetPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v)
}

func (s ApiTokenOwner) SetFrontend() {
	s.Struct.SetUint16(0, 3)

}

func (s ApiTokenOwner) User() ApiTokenOwner_user { return ApiTokenOwner_user(s) }

func (s ApiTokenOwner) SetUser() {
	s.Struct.SetUint16(0, 4)
}

func (s ApiTokenOwner_user) IdentityId() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s ApiTokenOwner_user) HasIdentityId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) IdentityIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_user) SetIdentityId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s ApiTokenOwner_user) Title() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s ApiTokenOwner_user) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_user) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s ApiTokenOwner_user) DenormalizedGrainMetadata() (DenormalizedGrainMetadata, error) {
	p, err := s.Struct.Ptr(2)
	return DenormalizedGrainMetadata{Struct: p.Struct()}, err
}

func (s ApiTokenOwner_user) HasDenormalizedGrainMetadata() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) SetDenormalizedGrainMetadata(v DenormalizedGrainMetadata) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDenormalizedGrainMetadata sets the denormalizedGrainMetadata field to a newly
// allocated DenormalizedGrainMetadata struct, preferring placement in s's segment.
func (s ApiTokenOwner_user) NewDenormalizedGrainMetadata() (DenormalizedGrainMetadata, error) {
	ss, err := NewDenormalizedGrainMetadata(s.Struct.Segment())
	if err != nil {
		return DenormalizedGrainMetadata{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s ApiTokenOwner_user) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ApiTokenOwner_user) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_user) SetUserId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiTokenOwner_user) UpstreamTitle() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s ApiTokenOwner_user) HasUpstreamTitle() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) UpstreamTitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s ApiTokenOwner_user) SetUpstreamTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

func (s ApiTokenOwner_user) Renamed() bool {
	return s.Struct.Bit(16)
}

func (s ApiTokenOwner_user) SetRenamed(v bool) {
	s.Struct.SetBit(16, v)
}

func (s ApiTokenOwner_user) SeenAllActivity() bool {
	return s.Struct.Bit(17)
}

func (s ApiTokenOwner_user) SetSeenAllActivity(v bool) {
	s.Struct.SetBit(17, v)
}

// ApiTokenOwner_List is a list of ApiTokenOwner.
type ApiTokenOwner_List struct{ capnp.List }

// NewApiTokenOwner creates a new list of ApiTokenOwner.
func NewApiTokenOwner_List(s *capnp.Segment, sz int32) (ApiTokenOwner_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return ApiTokenOwner_List{l}, err
}

func (s ApiTokenOwner_List) At(i int) ApiTokenOwner { return ApiTokenOwner{s.List.Struct(i)} }

func (s ApiTokenOwner_List) Set(i int, v ApiTokenOwner) error { return s.List.SetStruct(i, v.Struct) }

// ApiTokenOwner_Promise is a wrapper for a ApiTokenOwner promised by a client call.
type ApiTokenOwner_Promise struct{ *capnp.Pipeline }

func (p ApiTokenOwner_Promise) Struct() (ApiTokenOwner, error) {
	s, err := p.Pipeline.Struct()
	return ApiTokenOwner{s}, err
}

func (p ApiTokenOwner_Promise) Grain() ApiTokenOwner_grain_Promise {
	return ApiTokenOwner_grain_Promise{p.Pipeline}
}

// ApiTokenOwner_grain_Promise is a wrapper for a ApiTokenOwner_grain promised by a client call.
type ApiTokenOwner_grain_Promise struct{ *capnp.Pipeline }

func (p ApiTokenOwner_grain_Promise) Struct() (ApiTokenOwner_grain, error) {
	s, err := p.Pipeline.Struct()
	return ApiTokenOwner_grain{s}, err
}

func (p ApiTokenOwner_grain_Promise) SaveLabel() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p ApiTokenOwner_Promise) ClientPowerboxRequest() ApiTokenOwner_clientPowerboxRequest_Promise {
	return ApiTokenOwner_clientPowerboxRequest_Promise{p.Pipeline}
}

// ApiTokenOwner_clientPowerboxRequest_Promise is a wrapper for a ApiTokenOwner_clientPowerboxRequest promised by a client call.
type ApiTokenOwner_clientPowerboxRequest_Promise struct{ *capnp.Pipeline }

func (p ApiTokenOwner_clientPowerboxRequest_Promise) Struct() (ApiTokenOwner_clientPowerboxRequest, error) {
	s, err := p.Pipeline.Struct()
	return ApiTokenOwner_clientPowerboxRequest{s}, err
}

func (p ApiTokenOwner_Promise) ClientPowerboxOffer() ApiTokenOwner_clientPowerboxOffer_Promise {
	return ApiTokenOwner_clientPowerboxOffer_Promise{p.Pipeline}
}

// ApiTokenOwner_clientPowerboxOffer_Promise is a wrapper for a ApiTokenOwner_clientPowerboxOffer promised by a client call.
type ApiTokenOwner_clientPowerboxOffer_Promise struct{ *capnp.Pipeline }

func (p ApiTokenOwner_clientPowerboxOffer_Promise) Struct() (ApiTokenOwner_clientPowerboxOffer, error) {
	s, err := p.Pipeline.Struct()
	return ApiTokenOwner_clientPowerboxOffer{s}, err
}

func (p ApiTokenOwner_Promise) Internet() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p ApiTokenOwner_Promise) User() ApiTokenOwner_user_Promise {
	return ApiTokenOwner_user_Promise{p.Pipeline}
}

// ApiTokenOwner_user_Promise is a wrapper for a ApiTokenOwner_user promised by a client call.
type ApiTokenOwner_user_Promise struct{ *capnp.Pipeline }

func (p ApiTokenOwner_user_Promise) Struct() (ApiTokenOwner_user, error) {
	s, err := p.Pipeline.Struct()
	return ApiTokenOwner_user{s}, err
}

func (p ApiTokenOwner_user_Promise) DenormalizedGrainMetadata() DenormalizedGrainMetadata_Promise {
	return DenormalizedGrainMetadata_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SupervisorObjectId struct{ capnp.Struct }
type SupervisorObjectId_Which uint16

const (
	SupervisorObjectId_Which_appRef               SupervisorObjectId_Which = 0
	SupervisorObjectId_Which_wakeLockNotification SupervisorObjectId_Which = 1
)

func (w SupervisorObjectId_Which) String() string {
	const s = "appRefwakeLockNotification"
	switch w {
	case SupervisorObjectId_Which_appRef:
		return s[0:6]
	case SupervisorObjectId_Which_wakeLockNotification:
		return s[6:26]

	}
	return "SupervisorObjectId_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// SupervisorObjectId_TypeID is the unique identifier for the type SupervisorObjectId.
const SupervisorObjectId_TypeID = 0x8e74650737dbb840

func NewSupervisorObjectId(s *capnp.Segment) (SupervisorObjectId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SupervisorObjectId{st}, err
}

func NewRootSupervisorObjectId(s *capnp.Segment) (SupervisorObjectId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SupervisorObjectId{st}, err
}

func ReadRootSupervisorObjectId(msg *capnp.Message) (SupervisorObjectId, error) {
	root, err := msg.RootPtr()
	return SupervisorObjectId{root.Struct()}, err
}

func (s SupervisorObjectId) String() string {
	str, _ := text.Marshal(0x8e74650737dbb840, s.Struct)
	return str
}

func (s SupervisorObjectId) Which() SupervisorObjectId_Which {
	return SupervisorObjectId_Which(s.Struct.Uint16(0))
}
func (s SupervisorObjectId) AppRef() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SupervisorObjectId) HasAppRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SupervisorObjectId) AppRefPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SupervisorObjectId) SetAppRef(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPointer(0, v)
}

func (s SupervisorObjectId) SetAppRefPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v)
}

func (s SupervisorObjectId) WakeLockNotification() uint32 {
	return s.Struct.Uint32(4)
}

func (s SupervisorObjectId) SetWakeLockNotification(v uint32) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint32(4, v)
}

// SupervisorObjectId_List is a list of SupervisorObjectId.
type SupervisorObjectId_List struct{ capnp.List }

// NewSupervisorObjectId creates a new list of SupervisorObjectId.
func NewSupervisorObjectId_List(s *capnp.Segment, sz int32) (SupervisorObjectId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return SupervisorObjectId_List{l}, err
}

func (s SupervisorObjectId_List) At(i int) SupervisorObjectId {
	return SupervisorObjectId{s.List.Struct(i)}
}

func (s SupervisorObjectId_List) Set(i int, v SupervisorObjectId) error {
	return s.List.SetStruct(i, v.Struct)
}

// SupervisorObjectId_Promise is a wrapper for a SupervisorObjectId promised by a client call.
type SupervisorObjectId_Promise struct{ *capnp.Pipeline }

func (p SupervisorObjectId_Promise) Struct() (SupervisorObjectId, error) {
	s, err := p.Pipeline.Struct()
	return SupervisorObjectId{s}, err
}

func (p SupervisorObjectId_Promise) AppRef() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

const schema_c7205d6d32c7b040 = "x\xda\xacZ\x0dt\x14U\x96\xbe\xf7U7\x95hB" +
	"\xe7\xd1ag4\x09\x0d!!\x90@\xf8\x89\x0c\x88\x83" +
	"I\x90\x08a`I\x05\x10\x973\xccR\xe9~I\x9a" +
	"\xf4_\xaa*\x09AYfXX`\x06\xc7\xc1\xd9\x95" +
	"?Y\x95\x19W\x03*\xca\xa2\x0e\xb3\xe2\x10\x15\x15\x85" +
	"QY\x03\xc3\xb0\x8c\xfc\x0c\xeb\xc2.#.\xe2\x8a\x0a" +
	"\xb5\xe7UWuU\x92\x0e!\x07\x8f\xe7xb\xd7\xad" +
	"\xfb\xde\xbd\xef\xbb\xdf\xfd\xee+\xc7\xd4\xf7/#c\xdd" +
	"\xf3\xbf\x070g\x09q\xf7\xd3\x9f\xfe\xf4\xf0\x8e\xf6\xdf" +
	"}\xf2c\xa0\xd9\x08\xe0\x12\x01Jh\xd6\x0c\x04\x97~" +
	"\xed\xfd/\x82\x1d\x1dkV\x02\xcdE\x007\xf2GW" +
	"o?\x8c\x80^\x9aU\x0a\xa8?vq\xe7\x9b\xbeK" +
	"\x07W\x01\xcd&\xfa\x98y\x13\xfdC\xda\xa5M\x00X" +
	"26\xab\x1a\xbd\x15Y\"\x80\xb7<k6\xa0\xfeF" +
	"\xc5\xabsf\x7fZ\xb7\x06\xa8/\xe1ja\xd6\xeb\xdc" +
	"U\xd8p%i\xef\xfe\xe9\x9b\xd7V\xac\x01\xc9\x87\xa8" +
	"\xb7\xb5=sz\x82{\xc3\x1fa\xa0[D\x80\x92'" +
	"\xb36s\xd3\x17\x0d\xd3\xe6\x03#?\xfe\xa7\xf0\xbaN" +
	"\xbe:\xb2\x06\x10@\xef\x05\xc3\xc03\xf9\xad\x07\xee8" +
	"\xbecM<\xa2\xb8Az\xf6\x0a\xee\xe1\xb6lnP" +
	"\xf6\x9b\xe3\x13D\xa6\xfd\x1c$\x8a\xa8\x97\xbd\xf0\xf6\xb8" +
	"\xf0\xc2\xc1oC\x05\x8a$\x03K\xee\xcc\xceEoe" +
	"6\xdf|Ev\x0b\x80\xb7=\xdb\xa3?\xba\xb6\xe2O" +
	"G\xa4a\xbf\x00)\x1b\xad4\xed\xc96vu\xc0\xf0" +
	"\xf9\x0f\x13?\x1b\xf2\xf4\x9e\x87\xd6w\xf5\xe9\x12\x05\x80" +
	"\x92s\xd9\x85\xe8\xbdb\xf8\xbc\x9c\xbd\x13P\x7fe\xd7" +
	"\xcb\xff\xdbXp\xee\x11g\x10\xcbr\x96r\x7fks" +
	"\xb8\xbfg\x8f\x9c\xd9\xb7\xae\xf5\xe7\xbft\x1a\xb4\xe5\xbc" +
	"\xcb\x0d\xf6\x1a\x06\x8f\xff\xcd\xe03\xc7F\x7f\xfd\xa83" +
	"\xca\x1395\xdc\xe0\x9ca\xb0\xf6_F\xb7O\x18\xe3" +
	"\xde\xe84H\x1d\xb4\x98\x1b\x0c\x1c\xc4\x0d\x12'F\xa9" +
	"`\xef\x18\xd0;~\xd0io\xf9\xa0\xef\x00\x94\xcc\x1a" +
	"\xb4\x9ax\xd3\x07\x8b\x00\xfa\xe4\xf6\x0d\x0b\xaf\x96\xbc\xba" +
	"\x09h\x89\x95\x80\xcb\xbe\x01\x04\\\xfa\xe0\x17v]\x98" +
	"\xde\x9e\xbf\xc5\xdc\xaa\xc0\x1f]\xf0m\xe7\x0b]\xf5\xf1" +
	"`\xd3^\xfe\xc7mK\xe6\xffzK\xb7\x85\xda\x06\x7f" +
	"\xe5}i0_h\xff\xe0\xd5\xc4\xbb*\x97/4$" +
	"=0\xfd1\\\xbd5\xee\xceX(\x9c{\x98\x03\xf2" +
	"\xa7\xde\xd5\xdf\xdc\xbb\xf8\xa9\x7fv\xe6D\xce\xdd\xc5\x17" +
	"j\xcc\xe5\x11\xad\xf7UO\xdf\xf3\xd1\xa3\x8f;\x0d\xd6" +
	"\xe7\xfe\x94\x1b<i\x18\xb4\xdc=\xff\xc7U)\xbb\x9f" +
	"p\x80\xbd=W\xe1\xbe\xff\xef\xde\x86/\xefY\xf9\xdf" +
	"\xbf\x02\xea\x13\xec\x0d\x03\x96<\x97\xbb\x19\xbd\xfb\xf9\xce" +
	"\xbc\xed\xb9\xd3\xbc\xe7\x8c=\xeeX\x14\x99\xb8k\xe6\xd9" +
	"\xedfn\x8d\x90?\xcc]\xc0\x17:\x91\xcbC\xbe\xf4" +
	"y\xebG\xe3\x7f\xd4\xfe\xbc\x81gs\xa5\x87\x86\x1a[" +
	"\xdd:\x94\xef\xe4\xaeuK\xb6\xccL\xff\xfeN\xe7\xe9" +
	"\xec\x1d:\x85\x1b\xec7\x0c\xe6\xfe\xee\xf3\xe3\xbf<\xa2" +
	"\xect\xa4\xe1\xec\xd0[x\xbe\xbd;\xf6\x1c=\xd3\x98" +
	"\xb1\xc7\x19\xe5\x89\xa1\xdb\xf8\xab\x17\x8cWsw\xfd%" +
	"\xa7\xf2\xeam\xbfuD\x99\x9e\xb7\x94G)\xdc^>" +
	"i\xf9\xe3?s>\xb92t\x05\x7f\xb2z\xe0\xf4S" +
	"{_\xf8\xd1\xbf9\x8b\xfd\xecP\x03oW\x0c\xa7\xfb" +
	"\xbfXx\xffS\xaf\x1d\xdb\xdb\x1d\xe0\x04\xa0dT\xde" +
	"\x02\xf4\x96\xe7\xf1,M\xce\xe3\x09p\xf7\xffz\xf7\xf9" +
	"W\x1e\xde\xe7,\x98\xbdy\x85\xbcJ;\xf2\xb8\xbf[" +
	"\xd2vktGu\xbbi\x10_\xf1r\x9e\x01`\xcc" +
	"o\x01\xd4\xc9-\xeb\xdaG\x9c\x1e\xf9zW\xd8\x94\xc8" +
	"\xf9\x03\xd0\xdb\x98\xcf\x17\x0b\xe7\x8b\xdep~\x01\x80\xbe" +
	"\xbb\xe3\xb5\x07\x8f^X\xf7F7k\xc6\xad\x9b\x0c\xeb" +
	"\xc6\xfci\xde\xad\xf9\xdf\x01\xd0\xdf\xf5\x9e\x1a\x93\xb9`" +
	"\xf2\x9b\x8e\xecv\xe4\xff\x96'\xe2\xd9\x8aK\xcb\x1eX" +
	"\xcb\xdeN\xceD\xfb\xf3\x0d\\w\xe4\xf3\x18\x97_H" +
	"\x09^\xfa\xcf\xbf\x7f\xd7y\x10\xad\xc3\x8c\x9c=4\x8c" +
	"\xc7\xb8\xa5\xf1\xb1\xbb6O{\xb2\x93\xc1\x8b\xc3\x0c\x0f" +
	"\xed\x86\xc1\x90@\xbf\xfa\xb7]\xea{\xc9\x17;5\xac" +
	"\xd0(\xe7a\x9f\x00\xea\xb1\xc8\xa1\xe0W'\xff\xee=" +
	"\xc7\xd1\x1d*X\xccw\xfc\xf1\xbf6\x1e|\xf0{\xc3" +
	"\x0f:*sO\xc1\xed\x1c)\xae-\xdb\xbf~\xbd\xdf" +
	"\xa9\x83\xceCx\xb1`\x1c?\x84\x03\x05|\xfd\x0dY" +
	"\x0f\xb6=\xf2\xc3{~\xefH\xc3\xb9\x82\x1a\xee\xf4\x83" +
	"\x86\xaa\x0f\xde\x7f\xf8\xee\xf7\x1d\xcbu\x14\x18\x18\xfax" +
	"\xff\xb55G\xdf\xd9\xd0\xe1\xac\xa1\xf8;\xbf\xda; " +
	"\xa5~c\xd61\x90\x8a\xd0\x01\x98\x81.#\x9a\xe7\x0a" +
	"\x0c\x0a\xdeS\xc0S7\xf9\xc8\xb4\xd3Mo\xcc>\x96" +
	"<\xf0y\xc3s\x11\xb0d\xe1\xf0\x87\x11\x1cO\xbb " +
	"\xcf-\xf23\xbd:\xe2+oj!\xdf\x85\xbb\xd00" +
	"\x9f\xf7\xef\xcb\xee\x7f\xe7\xf7{\x8fw\x03\xc2\xf8\xa2\xa5" +
	"\xe8\xad,2\x88\xbdH\xf4V\x14q\xd8\xfc`\xf3\xe8" +
	"\xc3Yb\xf3\x7f8S4\xb9\xc88C\xa9\x88\xa7\xe8" +
	"\xad\xf9\xc7\xfc?Yw\xd7I\xe7\x196\x16\x19\xa5\xbe" +
	"\xcc0\x980\xa5\xe2|\xd5\x17\xbf\xfe\xb3\xb3\x92\xb7\x16" +
	"\x19<\xdbf\x18\xe0\xadw\x9f\xdd\xf4\xca\xf3\x7f1=" +
	"\x10np\xa0\xc8\xe8\x1d\xc7\x8a8\xd0\x0fo\xf0\x0d\xd8" +
	"2\xfe\xbf.:\x09T\x1aidk\xe1\xc8\x9d\x80\xd7" +
	"T\xfa\xfc_\xad;w\xc9Nx\xea\xa8\xd7y\xc27" +
	"\x9c\x7fn\xeeg\xfe_\\v\xee\xed\xcaH\x83\x09R" +
	"G\x19\x14\xffZ\xe9\xe2\x0b#S\xbf\x04)\x1f\x1d=" +
	"j\x9eKD\x82X2j\x14\xc7b\xc9\xe4Q>\x9e" +
	"\xb8\x94[/\x0e\x19}\xe5\xce/\x9d\xdeX\xb1b\xd0" +
	"k1\xf7V\xb8\xeb\x9b#5g\xde\xbabFj\x04" +
	"\xb2\xbe\xf8\x11\x83^\x8b[@7\xff\x19\xa9\xabM1" +
	"\xa64\x07UWT)\xf6\xcb\xb1Hl\xd2\x1c\xf3\x97" +
	"\xa8R\x1cP\xa2\xb1\xbcj\xa66\x854\x15 a\xdb" +
	"\xcf\xb6mU5\x16\xaeb\x8a\x1aT5\x16\xd1\x8a\xe5" +
	"@\xa0\x9a56\x05\x15\x16f\x11M5_FUr" +
	"\x09.\x00\x17\x02\xd0\xf4\\\x00)E@)\x93\xa0\xe8" +
	"\x97cHmB\x00D\x0ax\xfdM\xcdoi\xb97" +
	"\x18bs4YkB\xb5\x0aQJC\x02@s\x0a" +
	"\x8d\xd7\x07V\x03 \xa1t\x06\x80\xa76\x18bz " +
	"\xa80\xbf\x16U\x00[\xf5HT\xbb7\xda\x14\x09@" +
	"\xf2h\xe4H@\xd5\xa2J\xf8\x9e\xa8\xc2\x8a\xfd\xf5\xcc" +
	"\xdf\xd0)\x98*Y\x91\xc3*8cY\x0c \xa5\x09" +
	"(\x0d'\xa8+\xa6-x\xb85\xf6\x07\xac\x12\x103" +
	"\xec\xb3\x04\xe4?&\x16v'\x16.\x8f\x05\xe7F\x1b" +
	"XdvK\x84)\xc5\xfeP\x90E\xb4\xaah\x0bS" +
	"j\xa2Kf\xd7\xd6\x0aL\x89\xaf\x99\x89\x94/Zm" +
	".\xfa]\x82\xba\xcaT5\x18\x8dT\x02\x060\x0d\x08" +
	"\xa69\x16\x10{\x88\xac\x8ei\xc6J\x7f\x1d\xd5\x82\xb5" +
	"A\xbf\xac\x05\xa3\x91\xb9\xb2R\xc7\xb4\xa4\xe75\xce>" +
	"/_\x94\xbf\x87T\xbf8x\xcc\xd1\x09\xea\x97\x17\xbb" +
	"\x9e\x98;\xd9\x89\xd51m\x96\x1c\x8c\xdc\x17d-\x09" +
	"49\x17(\xb4\x17\xf04\x07Y\x0bR\xfd\x93\xd3u" +
	"\xe77\x1e\xdd}\xbc\xab\x7f!\x89\xff\xd95\x8b\x99_" +
	"\xd4*\x03\x1c\x0b)\x82+M\xd7\x0d\xbf#&\x01H" +
	"y\x02Jc\x08\xe6\xe05\x1d39m\xd0Q\xdb\x01" +
	"\xa41\x02J\xdf'X*\xc7b\xd5\xac\x16\x07\xa0C" +
	"j\x02\xe0\x00.@\xe4\x0663\xeao@+M\x1e" +
	"\x9e'L\x01\x82)\x80\xe5.\xa4X\xa3\x97\xc7b\xc6" +
	"\xf2\xc0\xd7\xef5\x0b&t\xa7\xcb\xfe\x86x\"\x04\xad" +
	"S\xa6'\xd9\x89(U9\xbeU\xf4\xd8\xd2\x1d\x10=" +
	"IS1\x8b\x85k\x149\xc2L\xb0z8Z\x8d\xba" +
	"\xb0sQ\xb1\x00@\x9a*\xa0TE0\x9d\xe7\xc2A" +
	"8t\xd6\x0a \xe9\xe4\xaa\x9e\x89.\x00zg\x0d\x80" +
	"4Q@i*A]\xe3\xb0\xbcO\x0e\x81\x10\xb4\xf1" +
	"\x15cJ8\xa8\xaaA\x12\x8d\xa8\xd3Y(\x00M*" +
	"S*\xd5\xf2\x80\x18\x0eF\xba\xa1\xd0\xd5\x03\x0a\x15\xc6" +
	"\xfffy\xd5\xa5qH\\\x8f\"\x06\xb8\x04@\xe3T" +
	"z\xab\xda\x1a\xd9\xdfP\xa7\xf0\x1a/\xf7k\xc1\xe6\xa0" +
	"\xd6j\x95-\xf4\x84i\xd6\xcc\"\x1af\xe8?\x08\xad" +
	"k;\xff\xf2\xbe\xa7y\xa23zc\xa1\xc4\xee\x99\xda" +
	"$\xf6y\xf7I]\xb6\xc8\x9a\xbf~f\xb4\x8e\xfb\xf4" +
	"t\xcd\x88\x13\x1a\xf5r$\x10bH\xf5\x93S\x16-" +
	"z6\xef\xf3\x8d]\xab\x84tu/\xaaQEr\xa1" +
	"s\x0eDE\xb7\x98\x14|\x06\xd6\xa4\xe1\x82\x1b !" +
	"I\xd1\x1a\xce\xbcc\xb1\x06\x88w\x04\x8ah\xb7P\xb4" +
	"\xf4\xb97\x07\xab\x81x\x07\xa2\x88BBp\xa0%\x81" +
	"\xbc\xa98\x03\x88\x17QD\x97\xd5\x1b\xedFN/o" +
	"\x07B?\x13\xd1\x9d\x10\xa1h\x09\xa1\xb1g\x07 \x10" +
	"zN\xc4~\x09\x0d\x8f\xd6$EOL\x01B?\x14" +
	"QL\xa8s\xb4\xa6c\xba\xbf\x10\x08\xdd#bJB" +
	"\xb8\xa25`\xd1\xe7f\x00\xa1O\x89H\x12\xaa\x09-" +
	"\xa1M7\xd5\x00\xa1\xebELMtO\xb4fI\xba" +
	"j)\x10\xbaL\xd4-*\x031\xc8Z\xcaPo`" +
	",V\x1e\x0a6\x03\xb22\xd4\xd5\xfa&-\x10m\x89" +
	"\x00@\x19\xea\xd1\x1a5\x1ab\x1a\xc3iL\x9b\xa6\xc8" +
	"\xc1\x88gNp)s<\x10\xac\x07\xfc\xf7\xf9\xf5," +
	"25X[\xcb\x14\x16A\xad\x0c\x97\x9b\x08+C\x0f" +
	"o\xc3e\x9c\x90\xe2\xf8\x88{W[#\xfe9\xbc\xad" +
	"\x89r\x1dwj\xd1\x0b\x94\xc6\x09\xa6\x0c\xab\xb0\xf7V" +
	"\xe0\xe8o\xb3kT\xa643\xa58\x1a\xff#^7" +
	"\xa8&c\xb5N>\xc2r\x03\xbb\xa7>\x18\x0a\x18=" +
	",\xafJ\xf6\xf0z\xe3\xecc\x82\xb7\x82\x83\xb7L@" +
	"i&A\x8a&\x0dW\x8e\xb3\x19\x89\x12\x92i\xf4\xf0" +
	"Y\xbc\x9f\xce\x14P\xaa'X\x1a\x93\x15^\x98\xe9@" +
	"0\x1d\xac\xd6\x93akL\xb3L\xfb\xdavI\xd70" +
	"J\xe3q\x18\xf5a\xcf\x96\xb8M\xb7\x92\x83Vv\xc4" +
	"f\xa6\x98Ub)0\xb4\xae\x07\xbccq\x0a\x10o" +
	"\xbeQ\x07\x96LDk\xcc\xf5\xdefT\x09E\x8eu" +
	"k^D\xeb~\xc4\xeb\xc6?\x02\xf1\xba\x91\x03\xd3R" +
	"\xb1h)~z\x85\x03\xfa\x02/\x04\xebr\x06\xady" +
	"\x86\x9ez\x04\x08=!\xa2=\xdb\xa35{\xd3\x0f9" +
	"h\x0f\xf0\xaa\xb5t-Z\x03)\xdd\xbb\x18\x08}\x89" +
	"\x17\x90u}\x81\xd64O\xdb6\xc7\x8b$%1\x1b" +
	"\xa35\x84\xd1M+\xac\"\xb1\x14-Z\xa3\x13]\xa5" +
	"\x18Eb#W\xe7\xd0\xe0\xa8\x00\x8c\xc4!j\xe8\x0e" +
	"\xd2Ux\x08L\xb3an\xe9.\xb4\x80\x89\xaa\xe9\x89" +
	"\x83\x0cJ\xe30\xe3\x86!9\x18\xaef\x8d\xe0ib" +
	"\xaaV\x86\xba\xc5\xfdh\x91\xbf\xa0\xb5\x96q\x80\xc4\xa2" +
	"\x8a6MA\xb3\xd2 \xbe\x95\xca\x00\x8bh\xe0\x0bj" +
	"\xad\x95\x81\xce\xc5\xd2\x87\xdeb\xe9\xa5\xbe\x16I5\xf3" +
	"ukz\xce\x9ed\xf4]\x0b\xfa7\xe4\xdc\xf2k\x8a" +
	"\xf5\xbe9N\xda\x8f,nc\x89\xdd\xf6\xba\x91nt" +
	"\"0\x85+\x11\x97Q3\xd6\xbd\x14Zc0\xa5\x9c" +
	"\xc6\xdd\xe2r\x93r:\x1f\xc3u\xfb.''\xa13" +
	"\xcb\xe4&c\x19N(\xd3\x05\x94\x02\x0e\x96\x91\xb9\xc2" +
	"Y$\xa0\x14\"(*\xac\x163\x1c\xe2\xafo\xa4\x12" +
	"\xe7\xa8\xb9Q\x10\xfbr^\x16\"M@\xe6U\xf9\x0c" +
	"\x89\xd2\xd3\x99\xd5\xb4jL\xc5T \x98\xda\xdb\x99\x19" +
	"\x13[\x929\xc5)I\x92G|\x13C\x83%\xb0z" +
	"\x8d\xdb\xaaX^\xaf\xdf\xa6rr\x0e\x18\xa5\xf1\xdd$" +
	"^HK\xf6\x82\xd5<\xe5:\x96(\x19\xd5~|\x83" +
	"\xc3m2ay3\x03\xa1\x1d\xdbT\x16\x89*a9" +
	"\x14\\\xca\x02\x06Lf1M\x0e\xc8\x82&\x9b\xba\xde" +
	"\xc2\xfc\x0c\xa7\xac\xd7\xb9\xacO\xdc\xec\xd0Y\x85@r" +
	"\x881\xf7p]?~\x9c=\xf7\xe8r,67\xa8" +
	"\x85\x18\x1fv2\xf4\xca\xd3\x0fLP^Y\xf83\x13" +
	"\x0b\x9e\xa0?\x1a\xf1\xc9\xb1X\xe5ugJ;_\x96" +
	"\xa0\xe9I\xcfh\x09Q\xe0\xc8\xd4\x14\xfb\xac\x97GC" +
	"\x01\xfe\xce\x8da<\xa1\x93\xabdE\xe4NS\x12N" +
	"G(\x00\xd2p\x01\xa5;\x08Z\x140\x96\xab\x8f\x91" +
	"\x02J\x13I\xbcE\x84\xa2u\xe5\xe0\x0bG\x9b\"\x9a" +
	"\xb5`\xa9\xaa)L\x0e#\xd5\xff\xf6\xcc\xfb#Z&" +
	"\xce?\xd4\xf3\xe0i\x03b:\xd7\xe1\xc8\xaa\x10\xab\x04" +
	"\xb7\x94\x82\xe8\xb8\xc8\x00\x878\x87d\xe3kgp\xa1" +
	"f\xd3\xa4u\xbf\x8b\xd6W\x1dJy\xd3M\x15u\x0b" +
	"\x80h\"\x10\xca\xd0P-\xeb\xef{\xfc\xd0\xe8!\xef" +
	"\xbd\x03eH\xd1'\xb9\x88\xe3'\x00\x8a\x05|s\xc6" +
	"p\xcc\xff\x9d!\xa0\xc1T\x9d\xa4T\x9f9\xcb\x1a\xdb" +
	"\xfax\x9dQ\xcd\x1a\x8dv\xcdA\xccO(\x9d\xa3x" +
	"\x8a\x83\xb9\x09fb\x7f\xce\xdc\x9bM)x?\xc1\x0c" +
	"\x92\x89\x1e\x00:\xaf\x1a@\x9a+\xa0\xb4\x88\xe0\xf2:" +
	"\xbe\x15\x07B\x83\x11M\x89\x06\x9a\xfc\xc8\x14\xa3\xb5\xf3" +
	"\xeeo\xc3\xf7:\xd7$7~\x01\x94\xb4\xbb\xcep\xdc" +
	"\xc6\x98}L\xe1eEm1\xd9\xf3%IW\x8a\x8d" +
	"\xef\x9cK\x92\xa4S\xe0\xedv\xcd\xf0\x81\xbck\xbf\x11" +
	"z8\x82:\x0f\xcf\x95\x94\x11\xbfC\xe2\x9ed\x9e\xf2" +
	"\x1f\xc6\x85\xb6\xd1,y_d<\xbd\x01\x01\xa5\x18A" +
	"*\x90Lt\x03\xd0\xf0R\x00)$\xa0\xb4\x84\x1f\x8e" +
	"\x90\x89\xa9\x00\xb4\x89\x1f\xce\x12\x01\xa5\x95I\x0eB\x95" +
	"\x9b\xd9L\xb9\x86\x01\x86\xbasK\xe2\x98\xa0\x94)\xf3" +
	"T\xa6\xdc\xe0\xf9]\x8f\x10\xac\xb1+\x91\xb2\x9b\x1a{" +
	"\xac\x8b\x99\x9b\xa5\xbcd\x12\xcfy\xd3\xa5\xf6Bx\x9d" +
	"6\x9c\xb8\x89\x15\x9d\xe1\xb9z\xebm]\x9b\xe1\xf5\xd3" +
	"g\xf2i\x12\x9c\xf6\xd0\x8f4\xb9\x98\xb7\x0a\x8e\x7fW" +
	"\x86I\xb7\x9d\x87=SqUNq\x0e{B&\x0a" +
	"\xcea\xef~\x82\xa5\xb5\xdc\xbff\x9d\xf7rYUy" +
	"1$\xce\xdf\xfc\xefq\xe0Y25\x16\xec\x06\x8b\x9e" +
	"\x80\xdf$\xaaL\x91\xb2\x05\x17\x0a\x99\xd8\x0f\x80\xbe\xc4" +
	"w\xf7\x82\x80\xd2\xab&\xeeE\x00\xba\x87w\xc6\xdd\x02" +
	"J\xfbL\x91\x98\x02@\xf7\x1e\x06\x90\xde\x14P\xfa\x80" +
	"`\x86\x90\x89\xb7\x00\xd0C\x0b\x00\xa4\x83\x02J\x7f " +
	"H]\xaeL\xbc\x15\x80v\xf0\x9e\xf3\x91\x80\xd2I\x82" +
	"\xd4\x9d\x91\x89i\x00\xf4\x04\x0f\xf8\x0f\x02J\x7f&H" +
	"\xfb\xd1L\xcc\x00\xa0\xa7V\x00H'\x05\x94\xfe\x87`" +
	"\xa9qg\x96\x08\xd0\xa7\xf1n\x9c\x88+`\xe6\x9bt" +
	"M8\xca\x98a\x7f\xae\xb1J\xca$\x0d\x10\x1c\x19k" +
	"\x8a\xc5{\x1a\xf8\xe6:]/WXD\x0e\xb3\x00\"" +
	"\x10D\x83\x19Y\xa4<\x14*Gs\xa4\x01\xfbI\xb7" +
	"\x81\xd9Lni<\xbb\xbcce\xdbW\x8c\xce\xd4\x9a" +
	"W\x8c\x89of<\xc5$\xdduU\x8f\x9f\xfc\x93\x9c" +
	"3\x9f\x10Pz\x96`\xba\xfb\x1b\xf3\xe2\xb1\x8d\xff\xfa" +
	"\x8c\x80\xd2n\x82\xe9\xfd\xbe\xe6\x0e\x12\xdf\x9e\xe8\x8b\x85" +
	"@\xd2\xc9W\xfc\xc7\xc4g?\xbai\x17\x90t\xe1\x0a" +
	"\xff1\xf1\x7f%\xd0\xb5\xdb\x80\x94\xb6\xb0\x9a\x06\xd6\x0a" +
	"\xfd|\x069qnaJ\x84i\xf1\xbb]b(\xca" +
	"Z%\x1a\xd1X$\x00 z\xf8a\xe8V\xa3B\xab" +
	"S\xf9\x0c\xa1\xda\xf5\xf7\xd9\xb5b-S\x92\x94\x94-" +
	"\x0cfG\xea\xa2\xc1H]\\*\x97\xc6\xb5r\x0f*" +
	"\xe1P\xd9\xda\xdf\x1c\xcfYu\xad\x87/\x12\xd7'\x9c" +
	"\xe4\x0d\xe9\xe6H\xc6\x1a\xac\xbe\xcd\xf91\x89\xfasn" +
	"\xd2\x1fU\x18R\xe7\xa7\xfc\x1bj\x97\x9d\xa6\x09k\xdb" +
	"\x0e-\xb8\xd8\xd6\x82\x89yp\xec6\x00\xe9\x8e8\x11" +
	"\x19\xfa\x9c\xa9\xda\\\xf0\x18AY\x95c\xca\xf6\x00V" +
	"\x99\xf7\xdcb4\x92P\xef\xbc8\xfa\xf7i$\xb7\xa6" +
	"\xb1\xde&\xd5\xde\xee\xc3\x92\xcfm\xdf\xda\xb5X\xda\x0d" +
	"\x83\xce\x0c\xa8\xfb\xbct\x1d1cI\xf5\x1etS\x82" +
	"\xc1\x0c\xdd\xf4\xdd\x13R\xdb\xd2\x8d+\xf7\xf5\x0c\x84." +
	"_<x\x1f/\xb6\xbeJD#\xaag:\x0b\x05\xa4" +
	"L\xc1\x85\xaexr\x97M\xb2\xf5\x8a\xa10yr\x7f" +
	"\xc29\xfaA\x01\xa55\xa6\xdc\xe1\xd4\xb4\xaa\x06@Z" +
	")\xa0\xf4\x04\xc1l]\x17\xe2\"h+'\xfe\xc7\x04" +
	"\x94\x9e\x89\x7fJ2\xbb\xc9SSl\x1e\xebB\xe8\xdd" +
	"T\x91\xb5=H\x86\xa6d\x0c\xbe\xdc(\xb5$#X" +
	"o\x1fTz\xb9M\xe8\xf5j\xa9\xc7OV\xe6\x85\xb0" +
	"\xb3\xc6\x0a\xedOl\x09$\x8f\x9ad\x17\x9e'&k" +
	"\xf5V\x04=NY\xff\x1f\x00\x00\xff\xff8\xba\xaa8"

func init() {
	schemas.Register(schema_c7205d6d32c7b040,
		0x80e7bfc1abd2efa7,
		0x8488d5d569f6cffe,
		0x85ccf31fc4aff09c,
		0x8867ef4f53bc45c3,
		0x8882befcdeca7451,
		0x888c6d95df2cc976,
		0x88abdb347bc63d0f,
		0x8e74650737dbb840,
		0x902651d6de458996,
		0x918db9a721f13886,
		0x92e92771f2b6b2b7,
		0x938e798cc0e3d6ac,
		0x96fb2fd9e320599f,
		0x98053037c12fa689,
		0x9951c12163385530,
		0x99bc33fd5d97c13d,
		0x9b25c148edb2b020,
		0x9ba45778a294b60c,
		0x9d87019c48640d21,
		0x9ea56a46fc87138a,
		0x9f96d4b948521f91,
		0xa0b4085080573e77,
		0xa3eb8443f86b46f7,
		0xaae54cb2386e60ab,
		0xaec15e35d479f4f3,
		0xaf3c0d4c9b788c3b,
		0xaf72d693dbf4bf54,
		0xb91071e3d7b9ab13,
		0xba19fd491deeb222,
		0xba8b9f7f3a411a03,
		0xbb5eb0bde1481587,
		0xbdd9bea5585df6c5,
		0xc08fb7eab4fb0e05,
		0xc152ab1174b40c0a,
		0xc22ce229c18c0a02,
		0xc38cedd77cbed5b4,
		0xc43d5a1430e113ca,
		0xc765897b7df345ac,
		0xca83e6f36908ed7f,
		0xcaa1479a3b9c719b,
		0xcb7304c768066421,
		0xcb7ee0fa69cd6e70,
		0xcc28367ccc71b3df,
		0xcce106c2fbaa9b04,
		0xce435c92a97c1b97,
		0xcf3e8fcfd0506bd0,
		0xd597c8d788fec5df,
		0xd91b98680812bda3,
		0xd94fc375e247d63d,
		0xda970537e2a8a9a9,
		0xdbbdcec8587dd355,
		0xdc76071bd22f9a4b,
		0xe03b8c8163d957c6,
		0xe4a4f650ea454237,
		0xeeaeb799e53e0b01,
		0xf0e8359b121f97d2,
		0xf3e98c16ae117300,
		0xf59063f154adea97,
		0xf8092ced6a3fbe30,
		0xf839f92f21f00b08,
		0xf9c6e362d6fcb22a)
}
