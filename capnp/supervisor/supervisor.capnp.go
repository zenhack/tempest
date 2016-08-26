package supervisor

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	activity "zenhack.net/go/sandstorm/capnp/activity"
	grain "zenhack.net/go/sandstorm/capnp/grain"
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
func (c Supervisor) GetGrainSize(ctx context.Context, params func(Supervisor_getGrainSize_Params) error, opts ...capnp.CallOption) Supervisor_getGrainSize_Results_Promise {
	if c.Client == nil {
		return Supervisor_getGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      3,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getGrainSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_getGrainSize_Params{Struct: s}) }
	}
	return Supervisor_getGrainSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Supervisor) GetGrainSizeWhenDifferent(ctx context.Context, params func(Supervisor_getGrainSizeWhenDifferent_Params) error, opts ...capnp.CallOption) Supervisor_getGrainSizeWhenDifferent_Results_Promise {
	if c.Client == nil {
		return Supervisor_getGrainSizeWhenDifferent_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getGrainSizeWhenDifferent",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Supervisor_getGrainSizeWhenDifferent_Params{Struct: s}) }
	}
	return Supervisor_getGrainSizeWhenDifferent_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
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

	GetGrainSize(Supervisor_getGrainSize) error

	GetGrainSizeWhenDifferent(Supervisor_getGrainSizeWhenDifferent) error

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
			MethodName:    "getGrainSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_getGrainSize{c, opts, Supervisor_getGrainSize_Params{Struct: p}, Supervisor_getGrainSize_Results{Struct: r}}
			return s.GetGrainSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x9951c12163385530,
			MethodID:      4,
			InterfaceName: "supervisor.capnp:Supervisor",
			MethodName:    "getGrainSizeWhenDifferent",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Supervisor_getGrainSizeWhenDifferent{c, opts, Supervisor_getGrainSizeWhenDifferent_Params{Struct: p}, Supervisor_getGrainSizeWhenDifferent_Results{Struct: r}}
			return s.GetGrainSizeWhenDifferent(call)
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

// Supervisor_getGrainSize holds the arguments for a server call to Supervisor.getGrainSize.
type Supervisor_getGrainSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_getGrainSize_Params
	Results Supervisor_getGrainSize_Results
}

// Supervisor_getGrainSizeWhenDifferent holds the arguments for a server call to Supervisor.getGrainSizeWhenDifferent.
type Supervisor_getGrainSizeWhenDifferent struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Supervisor_getGrainSizeWhenDifferent_Params
	Results Supervisor_getGrainSizeWhenDifferent_Results
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

type Supervisor_getGrainSize_Params struct{ capnp.Struct }

func NewSupervisor_getGrainSize_Params(s *capnp.Segment) (Supervisor_getGrainSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_getGrainSize_Params{st}, err
}

func NewRootSupervisor_getGrainSize_Params(s *capnp.Segment) (Supervisor_getGrainSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Supervisor_getGrainSize_Params{st}, err
}

func ReadRootSupervisor_getGrainSize_Params(msg *capnp.Message) (Supervisor_getGrainSize_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_getGrainSize_Params{root.Struct()}, err
}

func (s Supervisor_getGrainSize_Params) String() string {
	str, _ := text.Marshal(0xf3e98c16ae117300, s.Struct)
	return str
}

// Supervisor_getGrainSize_Params_List is a list of Supervisor_getGrainSize_Params.
type Supervisor_getGrainSize_Params_List struct{ capnp.List }

// NewSupervisor_getGrainSize_Params creates a new list of Supervisor_getGrainSize_Params.
func NewSupervisor_getGrainSize_Params_List(s *capnp.Segment, sz int32) (Supervisor_getGrainSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Supervisor_getGrainSize_Params_List{l}, err
}

func (s Supervisor_getGrainSize_Params_List) At(i int) Supervisor_getGrainSize_Params {
	return Supervisor_getGrainSize_Params{s.List.Struct(i)}
}

func (s Supervisor_getGrainSize_Params_List) Set(i int, v Supervisor_getGrainSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getGrainSize_Params_Promise is a wrapper for a Supervisor_getGrainSize_Params promised by a client call.
type Supervisor_getGrainSize_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getGrainSize_Params_Promise) Struct() (Supervisor_getGrainSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getGrainSize_Params{s}, err
}

type Supervisor_getGrainSize_Results struct{ capnp.Struct }

func NewSupervisor_getGrainSize_Results(s *capnp.Segment) (Supervisor_getGrainSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSize_Results{st}, err
}

func NewRootSupervisor_getGrainSize_Results(s *capnp.Segment) (Supervisor_getGrainSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSize_Results{st}, err
}

func ReadRootSupervisor_getGrainSize_Results(msg *capnp.Message) (Supervisor_getGrainSize_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_getGrainSize_Results{root.Struct()}, err
}

func (s Supervisor_getGrainSize_Results) String() string {
	str, _ := text.Marshal(0xdc76071bd22f9a4b, s.Struct)
	return str
}

func (s Supervisor_getGrainSize_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_getGrainSize_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_getGrainSize_Results_List is a list of Supervisor_getGrainSize_Results.
type Supervisor_getGrainSize_Results_List struct{ capnp.List }

// NewSupervisor_getGrainSize_Results creates a new list of Supervisor_getGrainSize_Results.
func NewSupervisor_getGrainSize_Results_List(s *capnp.Segment, sz int32) (Supervisor_getGrainSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_getGrainSize_Results_List{l}, err
}

func (s Supervisor_getGrainSize_Results_List) At(i int) Supervisor_getGrainSize_Results {
	return Supervisor_getGrainSize_Results{s.List.Struct(i)}
}

func (s Supervisor_getGrainSize_Results_List) Set(i int, v Supervisor_getGrainSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getGrainSize_Results_Promise is a wrapper for a Supervisor_getGrainSize_Results promised by a client call.
type Supervisor_getGrainSize_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getGrainSize_Results_Promise) Struct() (Supervisor_getGrainSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getGrainSize_Results{s}, err
}

type Supervisor_getGrainSizeWhenDifferent_Params struct{ capnp.Struct }

func NewSupervisor_getGrainSizeWhenDifferent_Params(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSizeWhenDifferent_Params{st}, err
}

func NewRootSupervisor_getGrainSizeWhenDifferent_Params(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSizeWhenDifferent_Params{st}, err
}

func ReadRootSupervisor_getGrainSizeWhenDifferent_Params(msg *capnp.Message) (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	root, err := msg.RootPtr()
	return Supervisor_getGrainSizeWhenDifferent_Params{root.Struct()}, err
}

func (s Supervisor_getGrainSizeWhenDifferent_Params) String() string {
	str, _ := text.Marshal(0xc08fb7eab4fb0e05, s.Struct)
	return str
}

func (s Supervisor_getGrainSizeWhenDifferent_Params) OldSize() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_getGrainSizeWhenDifferent_Params) SetOldSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_getGrainSizeWhenDifferent_Params_List is a list of Supervisor_getGrainSizeWhenDifferent_Params.
type Supervisor_getGrainSizeWhenDifferent_Params_List struct{ capnp.List }

// NewSupervisor_getGrainSizeWhenDifferent_Params creates a new list of Supervisor_getGrainSizeWhenDifferent_Params.
func NewSupervisor_getGrainSizeWhenDifferent_Params_List(s *capnp.Segment, sz int32) (Supervisor_getGrainSizeWhenDifferent_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_getGrainSizeWhenDifferent_Params_List{l}, err
}

func (s Supervisor_getGrainSizeWhenDifferent_Params_List) At(i int) Supervisor_getGrainSizeWhenDifferent_Params {
	return Supervisor_getGrainSizeWhenDifferent_Params{s.List.Struct(i)}
}

func (s Supervisor_getGrainSizeWhenDifferent_Params_List) Set(i int, v Supervisor_getGrainSizeWhenDifferent_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getGrainSizeWhenDifferent_Params_Promise is a wrapper for a Supervisor_getGrainSizeWhenDifferent_Params promised by a client call.
type Supervisor_getGrainSizeWhenDifferent_Params_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getGrainSizeWhenDifferent_Params_Promise) Struct() (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getGrainSizeWhenDifferent_Params{s}, err
}

type Supervisor_getGrainSizeWhenDifferent_Results struct{ capnp.Struct }

func NewSupervisor_getGrainSizeWhenDifferent_Results(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSizeWhenDifferent_Results{st}, err
}

func NewRootSupervisor_getGrainSizeWhenDifferent_Results(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Supervisor_getGrainSizeWhenDifferent_Results{st}, err
}

func ReadRootSupervisor_getGrainSizeWhenDifferent_Results(msg *capnp.Message) (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	root, err := msg.RootPtr()
	return Supervisor_getGrainSizeWhenDifferent_Results{root.Struct()}, err
}

func (s Supervisor_getGrainSizeWhenDifferent_Results) String() string {
	str, _ := text.Marshal(0xcce106c2fbaa9b04, s.Struct)
	return str
}

func (s Supervisor_getGrainSizeWhenDifferent_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_getGrainSizeWhenDifferent_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Supervisor_getGrainSizeWhenDifferent_Results_List is a list of Supervisor_getGrainSizeWhenDifferent_Results.
type Supervisor_getGrainSizeWhenDifferent_Results_List struct{ capnp.List }

// NewSupervisor_getGrainSizeWhenDifferent_Results creates a new list of Supervisor_getGrainSizeWhenDifferent_Results.
func NewSupervisor_getGrainSizeWhenDifferent_Results_List(s *capnp.Segment, sz int32) (Supervisor_getGrainSizeWhenDifferent_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Supervisor_getGrainSizeWhenDifferent_Results_List{l}, err
}

func (s Supervisor_getGrainSizeWhenDifferent_Results_List) At(i int) Supervisor_getGrainSizeWhenDifferent_Results {
	return Supervisor_getGrainSizeWhenDifferent_Results{s.List.Struct(i)}
}

func (s Supervisor_getGrainSizeWhenDifferent_Results_List) Set(i int, v Supervisor_getGrainSizeWhenDifferent_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Supervisor_getGrainSizeWhenDifferent_Results_Promise is a wrapper for a Supervisor_getGrainSizeWhenDifferent_Results promised by a client call.
type Supervisor_getGrainSizeWhenDifferent_Results_Promise struct{ *capnp.Pipeline }

func (p Supervisor_getGrainSizeWhenDifferent_Results_Promise) Struct() (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	s, err := p.Pipeline.Struct()
	return Supervisor_getGrainSizeWhenDifferent_Results{s}, err
}

type Supervisor_restore_Params struct{ capnp.Struct }

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

type SandstormCore_Server interface {
	Restore(SandstormCore_restore) error

	MakeToken(SandstormCore_makeToken) error

	GetOwnerNotificationTarget(SandstormCore_getOwnerNotificationTarget) error

	Drop(SandstormCore_drop) error

	CheckRequirements(SandstormCore_checkRequirements) error

	MakeChildToken(SandstormCore_makeChildToken) error

	ClaimRequest(SandstormCore_claimRequest) error

	BackgroundActivity(SandstormCore_backgroundActivity) error
}

func SandstormCore_ServerToClient(s SandstormCore_Server) SandstormCore {
	c, _ := s.(server.Closer)
	return SandstormCore{Client: server.New(SandstormCore_Methods(nil, s), c)}
}

func SandstormCore_Methods(methods []server.Method, s SandstormCore_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 8)
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

const schema_c7205d6d32c7b040 = "x\xda\xacY\x0d|TUv\xbf\xf7\xcdL^\x10\xc2" +
	"\xe41\xa1m\xbe\x9c$$\xcaD\xbe\xa3\x05\xd9\xc5$" +
	"\x08JXh2\x01\xc4\xf5'-/3/\xc9K\xe6" +
	"\x8b7/\x09A)+\xd5\x0a\xbb\xb8\x16\xb7+\x0aZ" +
	"\x95]\xaa\x88\xaeJQ7\xad\xb8F\x17?aU\xba" +
	"\x81\xcdRW>\x96ZiYq\xb5VTx=\xe7" +
	"\xbew\xe7\xbdL&\x84\xac\xfe\xf8%d\xce;\xef\xdc" +
	"{\xce\xfd\x9fs\xfe\xe7\xce\xb4\xdd95\xc2t\xcfG" +
	"U\x84,\xb9Y\xf0d\x19\x8f~tpW\xef/>" +
	"\xf8\x1e\x91\x8a(!n\x91\x90\xaa\xf7\xf2\x17R\xe26" +
	"\xce\xbf\xfd\x99\xda\xd7\xb7\xe1v\"\x95\xc1\x13\x0f\xc5G" +
	"o\xe4\x1f\xa4\x84\xfa\xde\xcb\xaf&\xd4x\xe0\xccS\xbf" +
	"\xf4\x7f\xb2\xff\x0exU0\xa6-\x9b\x15*\xed\x0d\xde" +
	"O\x08\xad:\x97\xdfH}R\x01\xbc\xe0\xcb)\xa8\x07" +
	"\xcdW\xe6\xbf\xb0\xa4\xfe\xa3\x96\x0dD\xf2\xa7L\x05\x0a" +
	"^FSW\x17\xa0\xa9\xa0\xfe\xe6\xef\xbezq\xfd\x06" +
	"\x12\xf4Sj\xec\xdc\xf9\xd8\xf1\x99\x9e-\xbf%\xe3=" +
	"\"\xe8WE\x0b\xb6\xa2\xeaZ\xa6\xda\xf9\xc6\xa4\xf7\x7f" +
	"\x1c\xdd4\xc0\xd6#\x05\xe3\x04P\xe8a\x0a\xde9\xaf" +
	"\xder\xe5\x91]\x1bL\x8fL\x85\xfe\x82\xf5h\xe1$" +
	"S\xa8\xf9\xf9\x91\x99\xa2\xa2\xff\x90\x04%X\xac\xe6\xe9" +
	"\xd7fDW\x94\xbcF\xe6SQ\xc8\xa5U\x9e\xc22" +
	"\xea\x1b_\x88\x9b\x97\x0a\xbb\xe0\xf7]\x85^\xe3\xde\x8d" +
	"\xf3\x7fw(x\xd9?\x90`\x11\xe5a\xba\xa3\x90\xed" +
	"\xea\xc7\x85h\xf3\xefg}\\\xfah\xcf]\x9b\xd3m" +
	"\xbaE\x17\xa8>[XI}\xfb\x98\xcd\xde\xc2\xa7@" +
	"\xfd\xf9\xdd\xcf\xfdq\xd5\xe5\x1f\xde\xe3t\xa2\xaeh\x0d" +
	"\xda[V\x84\xf6\x9e8t\xe2\xa5M\xdd?\xfc\x91S" +
	"\xa1\xa3\xe8MT\xd8\xc8\x14\x1e\xfan\xc9\x89\xfe\xa9_" +
	"\xde\xeb\xf4rgQ\x13*<\xcb\x146\xfe\xf3\xd4\xde" +
	"\x99\xd3<\xf79\x15\xfa\x8a\xdaP\xe1\x18SH\x9d\x98" +
	"$\xb9\xec\x1d\xc3cZ|\xdc\x97S\xfc\xe7\xa0\x9f_" +
	"|\xa7\xe0{\xf7Rx\xd5\x98\xd3\xbbe\xc5\xb9\xaa\x17" +
	"@\xbb\x8a\x07\xa0\xe7R\x88\xba\xdb(yz\xf7\xe9\x05" +
	"\xbd\x15\xdb\xac\xad\xba\xf0\xd13\x97>\x8e\x0b\xf5^\x8a" +
	"\xce\x8ey\xee\x1f\xb7\xaf^\xfe\xd3m\x83\x16\x8a\xfa\xbf" +
	"\xf0u\xfbq\xa1\x8d~Q\xf0}\\\x82\x0b\x95\xe6\x84" +
	"\x17<@\xef|\xd04\xc7\x16\xea/9\x88\x80\xfc\xbe" +
	"\xef\xce\xaf\xaek\xdb\xf1O\xce\x98\x1c(\xd9\xcd\x00Y" +
	"\x82\x1em\xf67.\xe8\xf9\xf5\xbd\x0f9\x15\xce\x95|" +
	"\x1f\x15rJQ\xa1\xeb\x9a\xe5\xdfk\xc8\xde\xf3\xb0\x03" +
	"\xec\x93K5\xb4\xfd\x7f\xd7\xb5\x7f~\xed\xed\xff\xfd\x13" +
	"x\xd5eo\x18\xb0\x9c_\xba\x95\xfa\xa6\x97\xe2\xd1M" +
	".\xbd\xde\xb7\x0c\xff2v\xad\x8c\xcd\xda\xbd\xe8\xe4\xe3" +
	"Vl\x99\xcbsJo\xc2\x85\xeaJ\xd1\xe5omZ" +
	"\xbdmQ\xce\xb7\x9fr\x06\xffl\xe9\\T\xa0e\xb8" +
	"\x93\xa5\xbf\xf8\xf4\xc8\x8f\x0eiO9\xbc,-\xbb\x04" +
	"\xc3\xe9\xdb\xd5s\xf8\xc4\xaa\xdc\x1e\xa7\x13\xf9e\xdb\xf1" +
	"\xd5\x00{\xb5l\xf7\x1f\x8a\xeb\xce\xe5\xff\xab\xc3\x89\xba" +
	"\xb25\xe8\x84\xab\xa0v\xf6\xba\x87~\xe0|rU\xd9" +
	"z|r\xe7\xf8\x05\xc7\xf6>\xfd\xd7\xff\xe6\xcc\xe5\xd2" +
	"2\x06\xa7\xab\x98\xd1}\x9f\xad\xb8q\xc7\x8b\xfd{\x07" +
	"\xe3W\x00U\xb5\xec&H\xbf2\x0cBw\x19\xfa\xe7" +
	"\x19\xfb\xe5\x9eS\xcf\xdf\xfd\x923\x1f\xce\x96\x9db\x91" +
	"\x9e\x80\xf6.\x19\xb3G\x97v5\xf6Z\x0a\xe6\x8a\x93" +
	"'0|^=\xa1\x0b4\x84K6\xf5\x06\x8eOz" +
	"9\x1d\x15U\xf7O\x18G};'\xe0b;&\x88" +
	"\xf0s9D|O\xdf\x8b\xb7\x1e>\xbd\xe9\x95A\xda" +
	"\x0f\xa2\xf6\x93L{\xe7\x84\xeb}}\x13\x00N\xc6\x13" +
	"\xf3?Y{\xcbF\xe5\xb5\xcc\xe5$\xa7\x9c\x81\xb3\xb8" +
	"\x1c=Yw:[\xfd\xe4?\xff\xeeMg\xb8{\xcb" +
	"Yd\xfa\xca\xd1\x93\xd2pV\xebk\xee\xe4[\x99m" +
	"\x9d-\xafd\x87Z\xf1\x01\xa8&b\x07\xd4/\x8e\xfe" +
	"\xed[\x8e\xf8\x9f\xach\xc3\xf8\xbf\xff/\xab\xf6\xdf\xfa" +
	"\x97\x13\xf7;\xb2\xe7\xdd\x8a\x02<n\xf7\xb6\xc7\xbf|" +
	"9\xeb\xd8~g$\xdf\xa8\xf8#\x03u\x05\xae\xbf\xa5" +
	"\xf0\xd6\x9d\xf7\xdc|\xed\xaf\x1cH9W\xd1\x84F\xdf" +
	"iox\xe7\xed\xbb\xafy\xdb\xb1\xdc\x87\x15\x0c\x08\xef" +
	"\xef;\xbf\xe1\xf0\xeb[\xfa\x1cO\xfa\xccw~\xb2w" +
	"\\v\xeb}\x85\xfd$x\x05u\x9c\xfax7\xf3\xa6" +
	"\xb7\x82\x95\xc9\x03\x15\x18\x999\x87\xae?\xde\xf1J}" +
	"\x7ff\xc7\xd5\xcb\x00I\xb4j\xd5ew\xc3\x7f\xf6\xd3" +
	"4\xf8xDVG'~\xe1+\x9e\xc8\x80<\x91\xa9" +
	"/\xfb\xf7\xb57\xbe\xfe\xab\xbdG\x06\x9df]`\x0d" +
	"\xf5\xad\x08\xe0K\xdf\x0d\x88\xf0\x83g\xff\x9d\xadS\x0f" +
	"\x16\x8a\x9d\xff\xe1\x0cQ0\xc0\xd2Z\x0e`\x88^]" +
	"\xde\x1f\xbam\xd3\xb7\x8e:\xcf\xf0\xb6\x00K\xc7\xbb\x98" +
	"\xc2\xcc\xb9\xf3O5|\xf6\xd3\xdf;\xd3\xf1\xc9\x00\xab" +
	"\x85=L\x81\x8e\xbe\xe6\xe4\xfd\xcf\xff\xec\x0f\x96\x05\x81" +
	"U\x9d\x00\xab\xef\x1f\x06\x10\xad\x07\xb7\xf8\xc7m\xbb\xea" +
	"\xbf\xce8\x8b\x9c\\\xc9\xa2\x15\xad\x84h\x9dOJ?" +
	"\xfb\xb3M\x1f~b\x07<\xff\x0a\x96y\xd3^\xacn" +
	";=i\xd4\xe7$XA\x1dmb\x19\x04\\\xa0\xd0" +
	"m\xae@,V\x8d\xbf\xc2\x8fq\xc9\x1e}\xa6t\xea" +
	"\xd9\xab?w:2g\x92\xc6\xea\xca$\xdcg\xe5\xee" +
	"\xaf\x0e5\x9dx\xf5\xac\xe5\x08\xdb\xa7:\xe9\x1eT\xe8" +
	"\x9e\xd4E\x0c\xeb\xdfh#\xd9\x91P\xb4N5\xe9\x8e" +
	"kSBr\"\x96\x98\xbd\xc4\x92\x80 \xac\xc5\x13\xe5" +
	"\x8dJ\xb2#\xa2'!\xbe\\7\xcb\xd6\xedN\xeaJ" +
	"\xb4A\xd1\x92*\xfc\x11\xd3\xa7\xc8\xe1p\xa3\xb2\xaaC" +
	"\xd5\x94(|LZ/\xd3d\xd0\xedr\x83\xbf\xb0\x17" +
	")\xa7\x8c\x90`\xb6\x8b\x06\xf3\x04*\x82\x19*\xd9I" +
	"K(\x95`\xf7\x17\xdc\xd4\xf2\xae\xae\xeb\xd4\x88\xb2D" +
	"\x97\xf5\x0e\x9al\xa048\x86B\xed\x91\x8a+\xd9\xeb" +
	"\xe3\x1b\xe1?A\x92\x16\x12\xe2m\x06=#\x0c\x9b\x09" +
	"\xe9q\x8d\xd0n#\x16\xd7\xaf\x8bw\xc4\xc2$\xb37" +
	"r,\x9c\x04\xcd\xe8\xb5qM\x99\x12jUB\xed\x03" +
	"\x9ci\x9059\x9a$N_\xda\xc0\x971\xe0\xcbD" +
	"\x81\x1a\x9a\xa5K\xbc\xa8M\xc7\x12\xda\xe0\xa24\xd7>" +
	"K\xd8\xdeX\x87w\x9e\xd4\xc2\xb5\x09ui\xbc]\x89" +
	"\xd5w\xc5\x14\x90ET0\xd0\x10\xefR\xb4\xa6\xf8\xea" +
	"\xfa\xe6f\x97\xa2\x99k\xe6At`\xd1Fk\xd1\xbf" +
	"\x80E\x93J2\xa9\xc6cu\x84\x86\xe9\x18\"\xc0\x8f" +
	"\xbd\x808\x84g-\x8a\xceV\xfa\xab\xb8\xae6\xab!" +
	"Y\x07\x03Ke\x0d\xc4\x19\xcfk\x86}^\xfe8\xbe" +
	"\x07'v\xa6d\xda\xe1\x99\xc9\xcf\xcf\xa4\x9f\x98'\xd3" +
	"\x89\x81\xe1\xc5\xb2\x1a\xbbAU\xbaRhr.Pi" +
	"/\xe0\xed\x04%\xb0\xff\xc1\xf1\x96S\xf7\x1d\xdes$" +
	"\xdd\xbe+\x83\xfd\xfa\xa66%$\xeaua\xc4B\xb6" +
	"\xcb=\xc60\x98\xdd\xc0l\xb0[\x0ev\xa7\x09\xb4\x98" +
	"\x9e7h\x1eV\x05i\xf2\xe3 \x9e\x06\xe2o\x0b\xb4" +
	"ZN$\x1a\x95f:\x8e:\xd8\x1e\xac9\x0e9\x80" +
	"\xdc\xae,\x8a\x87\xda)\x0f\x93\x17\xe3D\xb3!\xc8\xd9" +
	"\x84\xd6\xba\xa9D\x9b\x8c\xdaD\x82-Op\xfda\xa3" +
	"`Aw\x81\x1cj7\x03\xe1\xd2\x07Dz\xb6\x1d\x88" +
	"\xea$\xe2;I\xbd6{\x86Px3\x86b\xb1\x12" +
	"m\xd2\xe4\x98b\x81\xd5\x8bheya\xc7b\xfeM" +
	"`z\x1e\x98n\x10h\x0e\xc6\x82\xda\x05GZ\xbc\x9e" +
	"\x089\xc29#\x8f\xc2N\xa4\xab\x9b@u\x16\xa8\xce" +
	"\x03x\xe9\x08\xcb\x1b\xe4\x08q\xa96\xbe`\xfd\xa8\x0a" +
	"\xb0\x13\xe2\xb1\xe4\x02%\x12&\x1dIE\xabK\xd6\x86" +
	"\xc5\xa8\x1a\x1b\x84B\xf7\x10(\xd4\x14\xfc[)o\xac" +
	"6!q\xa1\x121\xce\xed\x02\xef\xc79\xac\x0e\x95\xb5" +
	"M\x10\xdb\x16\x0ds\xbc6\xa4\xab\x9d\xaa\xde\xcd\xd3\x96" +
	"\x0c\x85i\xa5\x13\xe2\x05y\xfa\x9d\xc8\xa6\x9d\xa7\x9e{" +
	"\xe9Q\x0ct\xeepU(\xb5{\xd8\xbc8\xe2\xddg" +
	"4\xd9%\xeb\xa1\xd6E\xf1\x16\xb4\xe9M\x8f\x88\x13\x1a" +
	"\xad\xe0qD\x81,9:w\xe5\xca'\xca?\xbd/" +
	"=K\x84t\xf3\"\xd8\x0f\xba\xa9s\x14\xa3\x9a\xc1+" +
	")\xf13\xac\x05'\xba<P\x179m\xa4|>\xf2" +
	"M\xa7MD\xf0\x05\xa8H\xed\x0eI9E\xf6\x15\xd3" +
	"Fx:\x1e\x9e\xbaR|\x82r\x86\xe3\x1bE\x17\xc2" +
	"S\x0aO\xdd\xbc\xf5\xd9}Z\xfa\xdf6\"H\xa7E" +
	"\xeaI\x11E\xcay\x8et\xec <;&\xd2\xac\x14" +
	"\x8b\xa6|\x96\x91\xfa\xe6\xc2\xb37D*\xa6\x084\xe5" +
	"\xf3\xa9\xb4\xb7\x12\x9e=#\xd2\xec\x14\xb7\xa4|\xc4\x91" +
	"v\xc0f\xa4\x07\xa1\xa3\xa68\x11\xe5\\X\xda\x0cN" +
	"J\x1bE:*\xd5<)\x9f\xe6\xa4\xb5k\xe0Y\x87" +
	"h\xf0JFD(S5\xd4hW\x94DmD\xed" +
	"$T\x81O\xc9\xd6\x0e=\x0c\x15\x12*\x08|\x02\xdd" +
	"\xeb5P&\xde%\xea\x1a\xc5!\x10\xf0\xf3\xf2V%" +
	"6OmnV4%F\xf5\x1a\xba\xce\x02T\x0d\xf5" +
	"b\xd7\xad\xc1\xfac\xc2\xc1\xb4\x96\xec\x8e\x85\x96`\x17" +
	"\x13\xe5\x16\xcb\x18;>Rm\xd6\x93\x1a\x0a\x19?l" +
	"\xe5w\xb4\xb3\xfa&H\xdaNh7q\xf3\x0f3M" +
	"h2S\x11\x1b`#\x0au\xf1\xdaV5\x12f-" +
	"\x0b^\xf3bza\xb1\xb1\xb0:\x1f\xb1Z\x03X]" +
	"$@\x8d\xb4\xaan\xdd\x0c\xbb\x00I\x82\x90\xc7Z\xf6" +
	"bl\x9f\x8b@\xd8\x0a\xa8N\xc8\x1a\xe6a\x0e\x94\x8f" +
	"\x1c\xc2;M\xae\xcd\x18\xad\xac\x1ci\x97\x15\xd2\xdd\xa8" +
	"6\xfd`\xe9`Ost\xbb\xc1\x83CytD\x08" +
	"O\xb0\x84%\x05'\\\x94\x0f\xe4\xd2\xc7\x08\xc1\x93\x08" +
	"zN\xf9(\x1f+\xa5~\xc8\x08\xe9]\x845\x1f\xdf" +
	"(\xbf\x8d\x90\xf6\xfd\x16\x9e\xedC\x08r6J9s" +
	"\x97z\x10\xbaO\"\xe4\xf9E\x08\xe5c\x87\xf4\xc8=" +
	"&t\xed9\x9a\xf29W\xda\xbc\xc6\x84\xae\xcdO)" +
	"\x9f\x0e\xa5\xb5m\x0c\xba\x90*\xfc\xaa\x80\xf2\xc9YR" +
	"\xb7\xc23E\xb4\xb1g\xe0\xe1\xe2\xb9\x12\x1a3A\xc6" +
	"\x88\x82\x90\xce\x14\\\x8an\x03\x95\x13%\xca\xa1E\x93" +
	"\x96%\x84\x09\xa96\x81\x82\x8a\x11Y\x8d\x82\"\xf1v" +
	"\xc0\x82 \xe0\xc5\x9a\xf2j\xed\xd2\xbb\x07\x02y\x04e" +
	"\x9eS\x97\x91\x02\xb8Q\xf1\x0f\xea?\xce\xf6\xc0Z " +
	"\x87\xe5E\x19\xe7v-\xde<2\xc3\x19[\x03\xaf3" +
	"Jj\xb7\xc3ndP\xaa\x03\x8bDR\xe0fx\xe6" +
	"\xb74\x94\x0f\x9c\x92\x84x\xf6\x88\xeb\xacr0\xf0\x18" +
	".\xd8\x02\xb1p\xb8\x06V\x80\xb2L\x15\x00\x93}\x01" +
	"\x08\xc3\x8e\x0a #\xd9X\x09\xc2\x08\xf4J\x0d\x98X" +
	"\xae\x83\x87\x8d,\xe1\xcd\xfa\xb14N\xc4\x8b\x0e+\x9b" +
	"o2\xb0zg\x03\xcf\xbc\xa9\xafA\xb19\x1d\x19\xf6" +
	"\x04y\xba`\xb2|\x93<\xc3I\xc7\xab\xcd\xdd\xa4^" +
	"\x18\x93\xe9\x05\xde{\xa0\xf3\xa4P\x9d\xb4\x1f_\xe4(" +
	"\x98\x89\x86}\x9d\xf1\xc9\xf6m\x9e\x12\x83\x98\xc9\x11h" +
	"\xaba\xd6a\x17+\xba\x1c\x96]\xbal\xb1`\x0e\xcb" +
	"\x85N\x12l \x09N]sH\x8b\xa1\xf6\x16\x0bl" +
	"J@\x16|\xd5\x0c{J0`JX\xaa\xea\x11\x05" +
	"G\x83\\\xa3\xee\xf8-3\xb5\xe7W\xfc\xc0\xc2\x82W" +
	"\x0d\xc5c~P\xa9\x1b<\x81e\x0d\x11\x7f\xb6\xcbt" +
	"\x1a\xa0\xa7z\xa9#Bs\xed3^\x17\x8f\x84\xf1\x1d" +
	":\x0aV\x19u\xd1l\x12\xe2.\xa2\xd1\xec\x94\xd1\x80" +
	"\x06F'\x82\xd1+\x05\xca\xb3s:6\xedI \x9b" +
	"%\x98u9\x12o\xa9%\xfe(\x94X\x9d/\x083" +
	"\x89\xa6\x00G\x90\x8c\xbf9\xf1v\xa0k\xd6\xf2\x03C" +
	"\x8fg6\x10\x16 [\xa5\x0a\x1cE\x83\xcb\x13\xcc\x86" +
	"\xa0\xdb\xe3>qPX\x92i\xc8\x1b\x08*\xaa\xdb\x15" +
	"\x8c\xdfTR\xfe\xf5\x83$\xc1\x18#\x8d\x12\x0d\x0e<" +
	"j!\x0f\xe8\x13k\xf6\x9box\xe8\xc0\xd4\xd2\xb7^" +
	"\x87\xcf\x12\xf5\x07\xdd\x82C\x04\x11\xa0\x97\xe3\xe6\xd8\x08" +
	"\x89\xbfs\xe17\x16\x91\x01\x0cd\x84\xb3;l\x83\xb5" +
	":\x13\x83y`\x0f@8\xd7Q\x1b\x05\x88\xfeX\xac" +
	"\x8d[-\"t#\x16L\xa8\x8d^\x10.\xc39\x7f" +
	")\x08W\xc2\xd9\xb7 b\x1c\x08Sc\xba\x16\x0fw" +
	"\x84(L^aX\x12[\xa7\x0d\xbf\x0b\\\x0a\\\xfc" +
	"uG\xc6\x06\xb6\xd0q\xf7`\xb5\x0a\x0d\xd3B\xb2\xb9" +
	"\xd4\x90\x98\x18\x18\xaa\x16/\xba\x04qf\xa1A\xe32" +
	"\x86\xe6f\x93\x0e\xb2\xb6\x81\x1dB\xc1(\x84A\x98\x00" +
	"\xa1\x0bB\x03\x87/E\xd7\x800\x02\xc2\xd5\x18DW" +
	"\x1e\xa0\x93H\x1d\x18\xc4\xd5 \xbc=C\xbc\x92r\xa7" +
	"\xb2Hn\x02b\x13\x19\x9c\xc2\xa9h\x92jE[\x06" +
	"N]d\x98/\x94\x7f|\x18HMs_\x8b\x9c\xf3" +
	"\xdb\x82?\xb5\xb2d\"9\xcek\x97\xe40ue\xc0" +
	"FS\xd7\x82\xa2\xd3-\xf7p\xad#\xbd\xd7\\8l" +
	"V\xd9\xca\x90qC\x94{]\x9e\x82\x95\x18\xe1\xe9\xce" +
	"\xb5\xaa\xda\xc0Q\xc4\xe2\x1cus\x9d\xa3\x08\x80\xc7\xe5" +
	"\x1cE \x03\xab\x9b\xd1\xbe\xce\xcfy\x9d\x9cL*\xba" +
	"\x03K\xd6\xe7\x19\xc4\xbbz^B\x1d\x04\x87\xa1\x00\xdf" +
	"!\xc2Y\x06\x8b\\n\x0akf\xc1\x9a\xcf\xe2\xee\x9e" +
	"\x865_\xb0\xf0.\x82\xb0\x07\x1b\xcf\x1e\x10\xbed\xd1" +
	"\xa4l\x10\xee=\x08\xc2_\x82\xf0\x1d(H\xf0\xf6%" +
	" ;\x80W:\xfbA\xf6\x1bPt\xbb\xf3\xe8h\x10" +
	"\xf6ai\xff5\x08\x8f\x82\xd0\x93\x9b\x07\x1b#\xd2{" +
	"\xe8\xf0o@\xf8{\x10fIy\x00y\"\x1d[\x0f" +
	"\xc2\xa3 \xfc\x1fp\x98]\xe0\xa4\x1c\xf4\xeb\xd8\xecR" +
	"~\x85\xadx\x0b\xe9\x01\xa72\xa4R\xea\xab\x01\x9eJ" +
	",M\xf4n\xe2rD\xac#a\xb6\x0e\xe2_\xea4" +
	"\x0d\x83GL\x8e*a8+\xe8D\xacp)\xb1\xda" +
	"H\xa4\x96Z\xa4\x9e\xd8O\x06\x8dsVp\xab\xcd\xe8" +
	"bc(\xb2\xef\xbb\x9c\xa1\xb5\xee\xbbR\xdf\xcf`\x88" +
	"\x85\x1c7\xdew\xe1\xc9?\x82%\xedaP}\x02T" +
	"=_Y\xb7`;Q\xfa\x18H\xf7\x804\xebK4" +
	"\x90\xfa\x9eCz\xa6\x12/\xcc\xbe@a\xea\x1b$\xe9" +
	"\xfe\xdd t\x9dEa\xea[ji\xe3v\"Tw" +
	")M\xedJ7\xc9\xf2\xb3\xa2\x845E\xd1b\x8an" +
	"^4\x0a\x8c\xb05kq\x90\xe2\xc5\xb4\xe8\xc5\xc30" +
	"x#\xa1\xbc\x93\xf8\x19\x0fL\x97\xd77\x8b\x90\xe1\x19" +
	"R\xca\xee\xbf\xf5\xb1\x96\xb8\x1ak1\x99h\xb5IE" +
	"\x87h\xc6\x07j6\xfe\xfcH\xf1\x1d\xe7Ifv\x9a" +
	"\xb9\xd0d\xee\x13_\xaf\xb8\xf0\x91\xe2\x9b\x9c\x9c2\x90" +
	"+\xe7&C\xb08t1\xc7W\xbaC]l\x0fM" +
	"\xd2\xf9\xb6\x1dT\xab\xcd\xa6Z\xa9Ih\xfav\x10^" +
	"i\x16 F{\xe1\xdd\xa5\xc4\xcb\x9c\xe2\x19c\xb1\xe1" +
	"0m\xb0.[\xc5x,E\x8a1)\xc6\x8eh\x18" +
	"\xe5C\xcep3\xdap\xb74\x99\xc7\xa1o\xec\xb2f" +
	"\xf4\xb0`\xb3\x1c\x194~\xa4\xdd\x81c\x13\x9d\xc2\xef" +
	"\xa9!r^\xbc\xa7\x0e\xe6A\xedu\x9b\x9e\xae\x9dm" +
	"\x93\x05F\xc3\xd0\xd3\xdb\xb0P\xde\x0a\xc2\x0d\x16\xd7\xc0" +
	"\xfap\x07\x8e\xa8\xb7\x83\xf0a\x81\x16\x19\x86\xcbd " +
	"\x0fb\xf5}\x00\xa4\x8f\x99_.X%}\xc7\\\xbb" +
	"\x98\xa4U\xd5A\x94\x84o\x8fd:\xdaLet\x1d" +
	"\xc3}\x861c\xb8+\xf6\x06\xbf\x9c\x0e\xfd\x11\xddp" +
	"\x0c\xf9%\x86ug\xe8\x04|\xa5\xfd\xa5K\x0aV\x93" +
	"g\xdbY\xe0M\xc8z+\xf7`\xc8\x89\xe2\xff\x03\x00" +
	"\x00\xff\xff\xae\x15\xe0\xb0"

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
		0xc765897b7df345ac,
		0xca83e6f36908ed7f,
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
		0xf8092ced6a3fbe30,
		0xf839f92f21f00b08,
		0xf9c6e362d6fcb22a)
}
