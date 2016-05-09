package supervisor

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	persistent "zenhack.net/go/sandstorm/capnp/capnp/persistent"
	grain "zenhack.net/go/sandstorm/capnp/sandstorm/grain"
	util "zenhack.net/go/sandstorm/capnp/sandstorm/util"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
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
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
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
	if err != nil {
		return Supervisor_WwwFileStatus_List{}, err
	}
	return Supervisor_WwwFileStatus_List{l.List}, nil
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
	if err != nil {
		return Supervisor_getMainView_Params{}, err
	}
	return Supervisor_getMainView_Params{st}, nil
}

func NewRootSupervisor_getMainView_Params(s *capnp.Segment) (Supervisor_getMainView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_getMainView_Params{}, err
	}
	return Supervisor_getMainView_Params{st}, nil
}

func ReadRootSupervisor_getMainView_Params(msg *capnp.Message) (Supervisor_getMainView_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getMainView_Params{}, err
	}
	return Supervisor_getMainView_Params{root.Struct()}, nil
}

// Supervisor_getMainView_Params_List is a list of Supervisor_getMainView_Params.
type Supervisor_getMainView_Params_List struct{ capnp.List }

// NewSupervisor_getMainView_Params creates a new list of Supervisor_getMainView_Params.
func NewSupervisor_getMainView_Params_List(s *capnp.Segment, sz int32) (Supervisor_getMainView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_getMainView_Params_List{}, err
	}
	return Supervisor_getMainView_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_getMainView_Results{}, err
	}
	return Supervisor_getMainView_Results{st}, nil
}

func NewRootSupervisor_getMainView_Results(s *capnp.Segment) (Supervisor_getMainView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Supervisor_getMainView_Results{}, err
	}
	return Supervisor_getMainView_Results{st}, nil
}

func ReadRootSupervisor_getMainView_Results(msg *capnp.Message) (Supervisor_getMainView_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getMainView_Results{}, err
	}
	return Supervisor_getMainView_Results{root.Struct()}, nil
}

func (s Supervisor_getMainView_Results) View() grain.UiView {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return grain.UiView{}
	}
	return grain.UiView{Client: p.Interface().Client()}
}

func (s Supervisor_getMainView_Results) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_getMainView_Results) SetView(v grain.UiView) error {

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

// Supervisor_getMainView_Results_List is a list of Supervisor_getMainView_Results.
type Supervisor_getMainView_Results_List struct{ capnp.List }

// NewSupervisor_getMainView_Results creates a new list of Supervisor_getMainView_Results.
func NewSupervisor_getMainView_Results_List(s *capnp.Segment, sz int32) (Supervisor_getMainView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Supervisor_getMainView_Results_List{}, err
	}
	return Supervisor_getMainView_Results_List{l}, nil
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_keepAlive_Params{}, err
	}
	return Supervisor_keepAlive_Params{st}, nil
}

func NewRootSupervisor_keepAlive_Params(s *capnp.Segment) (Supervisor_keepAlive_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_keepAlive_Params{}, err
	}
	return Supervisor_keepAlive_Params{st}, nil
}

func ReadRootSupervisor_keepAlive_Params(msg *capnp.Message) (Supervisor_keepAlive_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_keepAlive_Params{}, err
	}
	return Supervisor_keepAlive_Params{root.Struct()}, nil
}

// Supervisor_keepAlive_Params_List is a list of Supervisor_keepAlive_Params.
type Supervisor_keepAlive_Params_List struct{ capnp.List }

// NewSupervisor_keepAlive_Params creates a new list of Supervisor_keepAlive_Params.
func NewSupervisor_keepAlive_Params_List(s *capnp.Segment, sz int32) (Supervisor_keepAlive_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_keepAlive_Params_List{}, err
	}
	return Supervisor_keepAlive_Params_List{l}, nil
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

type Supervisor_keepAlive_Results struct{ capnp.Struct }

func NewSupervisor_keepAlive_Results(s *capnp.Segment) (Supervisor_keepAlive_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_keepAlive_Results{}, err
	}
	return Supervisor_keepAlive_Results{st}, nil
}

func NewRootSupervisor_keepAlive_Results(s *capnp.Segment) (Supervisor_keepAlive_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_keepAlive_Results{}, err
	}
	return Supervisor_keepAlive_Results{st}, nil
}

func ReadRootSupervisor_keepAlive_Results(msg *capnp.Message) (Supervisor_keepAlive_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_keepAlive_Results{}, err
	}
	return Supervisor_keepAlive_Results{root.Struct()}, nil
}

// Supervisor_keepAlive_Results_List is a list of Supervisor_keepAlive_Results.
type Supervisor_keepAlive_Results_List struct{ capnp.List }

// NewSupervisor_keepAlive_Results creates a new list of Supervisor_keepAlive_Results.
func NewSupervisor_keepAlive_Results_List(s *capnp.Segment, sz int32) (Supervisor_keepAlive_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_keepAlive_Results_List{}, err
	}
	return Supervisor_keepAlive_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_shutdown_Params{}, err
	}
	return Supervisor_shutdown_Params{st}, nil
}

func NewRootSupervisor_shutdown_Params(s *capnp.Segment) (Supervisor_shutdown_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_shutdown_Params{}, err
	}
	return Supervisor_shutdown_Params{st}, nil
}

func ReadRootSupervisor_shutdown_Params(msg *capnp.Message) (Supervisor_shutdown_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_shutdown_Params{}, err
	}
	return Supervisor_shutdown_Params{root.Struct()}, nil
}

// Supervisor_shutdown_Params_List is a list of Supervisor_shutdown_Params.
type Supervisor_shutdown_Params_List struct{ capnp.List }

// NewSupervisor_shutdown_Params creates a new list of Supervisor_shutdown_Params.
func NewSupervisor_shutdown_Params_List(s *capnp.Segment, sz int32) (Supervisor_shutdown_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_shutdown_Params_List{}, err
	}
	return Supervisor_shutdown_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_shutdown_Results{}, err
	}
	return Supervisor_shutdown_Results{st}, nil
}

func NewRootSupervisor_shutdown_Results(s *capnp.Segment) (Supervisor_shutdown_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_shutdown_Results{}, err
	}
	return Supervisor_shutdown_Results{st}, nil
}

func ReadRootSupervisor_shutdown_Results(msg *capnp.Message) (Supervisor_shutdown_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_shutdown_Results{}, err
	}
	return Supervisor_shutdown_Results{root.Struct()}, nil
}

// Supervisor_shutdown_Results_List is a list of Supervisor_shutdown_Results.
type Supervisor_shutdown_Results_List struct{ capnp.List }

// NewSupervisor_shutdown_Results creates a new list of Supervisor_shutdown_Results.
func NewSupervisor_shutdown_Results_List(s *capnp.Segment, sz int32) (Supervisor_shutdown_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_shutdown_Results_List{}, err
	}
	return Supervisor_shutdown_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_getGrainSize_Params{}, err
	}
	return Supervisor_getGrainSize_Params{st}, nil
}

func NewRootSupervisor_getGrainSize_Params(s *capnp.Segment) (Supervisor_getGrainSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_getGrainSize_Params{}, err
	}
	return Supervisor_getGrainSize_Params{st}, nil
}

func ReadRootSupervisor_getGrainSize_Params(msg *capnp.Message) (Supervisor_getGrainSize_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getGrainSize_Params{}, err
	}
	return Supervisor_getGrainSize_Params{root.Struct()}, nil
}

// Supervisor_getGrainSize_Params_List is a list of Supervisor_getGrainSize_Params.
type Supervisor_getGrainSize_Params_List struct{ capnp.List }

// NewSupervisor_getGrainSize_Params creates a new list of Supervisor_getGrainSize_Params.
func NewSupervisor_getGrainSize_Params_List(s *capnp.Segment, sz int32) (Supervisor_getGrainSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_getGrainSize_Params_List{}, err
	}
	return Supervisor_getGrainSize_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_getGrainSize_Results{}, err
	}
	return Supervisor_getGrainSize_Results{st}, nil
}

func NewRootSupervisor_getGrainSize_Results(s *capnp.Segment) (Supervisor_getGrainSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Supervisor_getGrainSize_Results{}, err
	}
	return Supervisor_getGrainSize_Results{st}, nil
}

func ReadRootSupervisor_getGrainSize_Results(msg *capnp.Message) (Supervisor_getGrainSize_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getGrainSize_Results{}, err
	}
	return Supervisor_getGrainSize_Results{root.Struct()}, nil
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
	if err != nil {
		return Supervisor_getGrainSize_Results_List{}, err
	}
	return Supervisor_getGrainSize_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Params{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Params{st}, nil
}

func NewRootSupervisor_getGrainSizeWhenDifferent_Params(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Params{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Params{st}, nil
}

func ReadRootSupervisor_getGrainSizeWhenDifferent_Params(msg *capnp.Message) (Supervisor_getGrainSizeWhenDifferent_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Params{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Params{root.Struct()}, nil
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
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Params_List{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Results{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Results{st}, nil
}

func NewRootSupervisor_getGrainSizeWhenDifferent_Results(s *capnp.Segment) (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Results{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Results{st}, nil
}

func ReadRootSupervisor_getGrainSizeWhenDifferent_Results(msg *capnp.Message) (Supervisor_getGrainSizeWhenDifferent_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Results{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Results{root.Struct()}, nil
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
	if err != nil {
		return Supervisor_getGrainSizeWhenDifferent_Results_List{}, err
	}
	return Supervisor_getGrainSizeWhenDifferent_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_restore_Params{}, err
	}
	return Supervisor_restore_Params{st}, nil
}

func NewRootSupervisor_restore_Params(s *capnp.Segment) (Supervisor_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return Supervisor_restore_Params{}, err
	}
	return Supervisor_restore_Params{st}, nil
}

func ReadRootSupervisor_restore_Params(msg *capnp.Message) (Supervisor_restore_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_restore_Params{}, err
	}
	return Supervisor_restore_Params{root.Struct()}, nil
}

func (s Supervisor_restore_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return SupervisorObjectId{}, err
	}

	return SupervisorObjectId{Struct: p.Struct()}, nil

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
	if err != nil {
		return MembraneRequirement_List{}, err
	}

	return MembraneRequirement_List{List: p.List()}, nil

}

func (s Supervisor_restore_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Supervisor_restore_Params) SetRequirements(v MembraneRequirement_List) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s Supervisor_restore_Params) ParentToken() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return Supervisor_restore_Params_List{}, err
	}
	return Supervisor_restore_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_restore_Results{}, err
	}
	return Supervisor_restore_Results{st}, nil
}

func NewRootSupervisor_restore_Results(s *capnp.Segment) (Supervisor_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Supervisor_restore_Results{}, err
	}
	return Supervisor_restore_Results{st}, nil
}

func ReadRootSupervisor_restore_Results(msg *capnp.Message) (Supervisor_restore_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_restore_Results{}, err
	}
	return Supervisor_restore_Results{root.Struct()}, nil
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
	if err != nil {
		return Supervisor_restore_Results_List{}, err
	}
	return Supervisor_restore_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_drop_Params{}, err
	}
	return Supervisor_drop_Params{st}, nil
}

func NewRootSupervisor_drop_Params(s *capnp.Segment) (Supervisor_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Supervisor_drop_Params{}, err
	}
	return Supervisor_drop_Params{st}, nil
}

func ReadRootSupervisor_drop_Params(msg *capnp.Message) (Supervisor_drop_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_drop_Params{}, err
	}
	return Supervisor_drop_Params{root.Struct()}, nil
}

func (s Supervisor_drop_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return SupervisorObjectId{}, err
	}

	return SupervisorObjectId{Struct: p.Struct()}, nil

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
	if err != nil {
		return Supervisor_drop_Params_List{}, err
	}
	return Supervisor_drop_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_drop_Results{}, err
	}
	return Supervisor_drop_Results{st}, nil
}

func NewRootSupervisor_drop_Results(s *capnp.Segment) (Supervisor_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_drop_Results{}, err
	}
	return Supervisor_drop_Results{st}, nil
}

func ReadRootSupervisor_drop_Results(msg *capnp.Message) (Supervisor_drop_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_drop_Results{}, err
	}
	return Supervisor_drop_Results{root.Struct()}, nil
}

// Supervisor_drop_Results_List is a list of Supervisor_drop_Results.
type Supervisor_drop_Results_List struct{ capnp.List }

// NewSupervisor_drop_Results creates a new list of Supervisor_drop_Results.
func NewSupervisor_drop_Results_List(s *capnp.Segment, sz int32) (Supervisor_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_drop_Results_List{}, err
	}
	return Supervisor_drop_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_watchLog_Params{}, err
	}
	return Supervisor_watchLog_Params{st}, nil
}

func NewRootSupervisor_watchLog_Params(s *capnp.Segment) (Supervisor_watchLog_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Supervisor_watchLog_Params{}, err
	}
	return Supervisor_watchLog_Params{st}, nil
}

func ReadRootSupervisor_watchLog_Params(msg *capnp.Message) (Supervisor_watchLog_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_watchLog_Params{}, err
	}
	return Supervisor_watchLog_Params{root.Struct()}, nil
}

func (s Supervisor_watchLog_Params) BacklogAmount() uint64 {
	return s.Struct.Uint64(0)
}

func (s Supervisor_watchLog_Params) SetBacklogAmount(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s Supervisor_watchLog_Params) Stream() util.ByteStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Supervisor_watchLog_Params) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_watchLog_Params) SetStream(v util.ByteStream) error {

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

// Supervisor_watchLog_Params_List is a list of Supervisor_watchLog_Params.
type Supervisor_watchLog_Params_List struct{ capnp.List }

// NewSupervisor_watchLog_Params creates a new list of Supervisor_watchLog_Params.
func NewSupervisor_watchLog_Params_List(s *capnp.Segment, sz int32) (Supervisor_watchLog_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Supervisor_watchLog_Params_List{}, err
	}
	return Supervisor_watchLog_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_watchLog_Results{}, err
	}
	return Supervisor_watchLog_Results{st}, nil
}

func NewRootSupervisor_watchLog_Results(s *capnp.Segment) (Supervisor_watchLog_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Supervisor_watchLog_Results{}, err
	}
	return Supervisor_watchLog_Results{st}, nil
}

func ReadRootSupervisor_watchLog_Results(msg *capnp.Message) (Supervisor_watchLog_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_watchLog_Results{}, err
	}
	return Supervisor_watchLog_Results{root.Struct()}, nil
}

func (s Supervisor_watchLog_Results) Handle() util.Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s Supervisor_watchLog_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_watchLog_Results) SetHandle(v util.Handle) error {

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

// Supervisor_watchLog_Results_List is a list of Supervisor_watchLog_Results.
type Supervisor_watchLog_Results_List struct{ capnp.List }

// NewSupervisor_watchLog_Results creates a new list of Supervisor_watchLog_Results.
func NewSupervisor_watchLog_Results_List(s *capnp.Segment, sz int32) (Supervisor_watchLog_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Supervisor_watchLog_Results_List{}, err
	}
	return Supervisor_watchLog_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_syncStorage_Params{}, err
	}
	return Supervisor_syncStorage_Params{st}, nil
}

func NewRootSupervisor_syncStorage_Params(s *capnp.Segment) (Supervisor_syncStorage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_syncStorage_Params{}, err
	}
	return Supervisor_syncStorage_Params{st}, nil
}

func ReadRootSupervisor_syncStorage_Params(msg *capnp.Message) (Supervisor_syncStorage_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_syncStorage_Params{}, err
	}
	return Supervisor_syncStorage_Params{root.Struct()}, nil
}

// Supervisor_syncStorage_Params_List is a list of Supervisor_syncStorage_Params.
type Supervisor_syncStorage_Params_List struct{ capnp.List }

// NewSupervisor_syncStorage_Params creates a new list of Supervisor_syncStorage_Params.
func NewSupervisor_syncStorage_Params_List(s *capnp.Segment, sz int32) (Supervisor_syncStorage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_syncStorage_Params_List{}, err
	}
	return Supervisor_syncStorage_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_syncStorage_Results{}, err
	}
	return Supervisor_syncStorage_Results{st}, nil
}

func NewRootSupervisor_syncStorage_Results(s *capnp.Segment) (Supervisor_syncStorage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Supervisor_syncStorage_Results{}, err
	}
	return Supervisor_syncStorage_Results{st}, nil
}

func ReadRootSupervisor_syncStorage_Results(msg *capnp.Message) (Supervisor_syncStorage_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_syncStorage_Results{}, err
	}
	return Supervisor_syncStorage_Results{root.Struct()}, nil
}

// Supervisor_syncStorage_Results_List is a list of Supervisor_syncStorage_Results.
type Supervisor_syncStorage_Results_List struct{ capnp.List }

// NewSupervisor_syncStorage_Results creates a new list of Supervisor_syncStorage_Results.
func NewSupervisor_syncStorage_Results_List(s *capnp.Segment, sz int32) (Supervisor_syncStorage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Supervisor_syncStorage_Results_List{}, err
	}
	return Supervisor_syncStorage_Results_List{l}, nil
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
	if err != nil {
		return Supervisor_getWwwFileHack_Params{}, err
	}
	return Supervisor_getWwwFileHack_Params{st}, nil
}

func NewRootSupervisor_getWwwFileHack_Params(s *capnp.Segment) (Supervisor_getWwwFileHack_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Supervisor_getWwwFileHack_Params{}, err
	}
	return Supervisor_getWwwFileHack_Params{st}, nil
}

func ReadRootSupervisor_getWwwFileHack_Params(msg *capnp.Message) (Supervisor_getWwwFileHack_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getWwwFileHack_Params{}, err
	}
	return Supervisor_getWwwFileHack_Params{root.Struct()}, nil
}

func (s Supervisor_getWwwFileHack_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Supervisor_getWwwFileHack_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Supervisor_getWwwFileHack_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s Supervisor_getWwwFileHack_Params) SetPath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Supervisor_getWwwFileHack_Params) Stream() util.ByteStream {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Supervisor_getWwwFileHack_Params) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Supervisor_getWwwFileHack_Params) SetStream(v util.ByteStream) error {

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

// Supervisor_getWwwFileHack_Params_List is a list of Supervisor_getWwwFileHack_Params.
type Supervisor_getWwwFileHack_Params_List struct{ capnp.List }

// NewSupervisor_getWwwFileHack_Params creates a new list of Supervisor_getWwwFileHack_Params.
func NewSupervisor_getWwwFileHack_Params_List(s *capnp.Segment, sz int32) (Supervisor_getWwwFileHack_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Supervisor_getWwwFileHack_Params_List{}, err
	}
	return Supervisor_getWwwFileHack_Params_List{l}, nil
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
	if err != nil {
		return Supervisor_getWwwFileHack_Results{}, err
	}
	return Supervisor_getWwwFileHack_Results{st}, nil
}

func NewRootSupervisor_getWwwFileHack_Results(s *capnp.Segment) (Supervisor_getWwwFileHack_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Supervisor_getWwwFileHack_Results{}, err
	}
	return Supervisor_getWwwFileHack_Results{st}, nil
}

func ReadRootSupervisor_getWwwFileHack_Results(msg *capnp.Message) (Supervisor_getWwwFileHack_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Supervisor_getWwwFileHack_Results{}, err
	}
	return Supervisor_getWwwFileHack_Results{root.Struct()}, nil
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
	if err != nil {
		return Supervisor_getWwwFileHack_Results_List{}, err
	}
	return Supervisor_getWwwFileHack_Results_List{l}, nil
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
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
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

type SandstormCore_Server interface {
	Restore(SandstormCore_restore) error

	MakeToken(SandstormCore_makeToken) error

	GetOwnerNotificationTarget(SandstormCore_getOwnerNotificationTarget) error

	Drop(SandstormCore_drop) error

	CheckRequirements(SandstormCore_checkRequirements) error

	MakeChildToken(SandstormCore_makeChildToken) error
}

func SandstormCore_ServerToClient(s SandstormCore_Server) SandstormCore {
	c, _ := s.(server.Closer)
	return SandstormCore{Client: server.New(SandstormCore_Methods(nil, s), c)}
}

func SandstormCore_Methods(methods []server.Method, s SandstormCore_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
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
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Params{}, err
	}
	return SandstormCore_RequirementObserver_observe_Params{st}, nil
}

func NewRootSandstormCore_RequirementObserver_observe_Params(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Params{}, err
	}
	return SandstormCore_RequirementObserver_observe_Params{st}, nil
}

func ReadRootSandstormCore_RequirementObserver_observe_Params(msg *capnp.Message) (SandstormCore_RequirementObserver_observe_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Params{}, err
	}
	return SandstormCore_RequirementObserver_observe_Params{root.Struct()}, nil
}

// SandstormCore_RequirementObserver_observe_Params_List is a list of SandstormCore_RequirementObserver_observe_Params.
type SandstormCore_RequirementObserver_observe_Params_List struct{ capnp.List }

// NewSandstormCore_RequirementObserver_observe_Params creates a new list of SandstormCore_RequirementObserver_observe_Params.
func NewSandstormCore_RequirementObserver_observe_Params_List(s *capnp.Segment, sz int32) (SandstormCore_RequirementObserver_observe_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Params_List{}, err
	}
	return SandstormCore_RequirementObserver_observe_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Results{}, err
	}
	return SandstormCore_RequirementObserver_observe_Results{st}, nil
}

func NewRootSandstormCore_RequirementObserver_observe_Results(s *capnp.Segment) (SandstormCore_RequirementObserver_observe_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Results{}, err
	}
	return SandstormCore_RequirementObserver_observe_Results{st}, nil
}

func ReadRootSandstormCore_RequirementObserver_observe_Results(msg *capnp.Message) (SandstormCore_RequirementObserver_observe_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Results{}, err
	}
	return SandstormCore_RequirementObserver_observe_Results{root.Struct()}, nil
}

// SandstormCore_RequirementObserver_observe_Results_List is a list of SandstormCore_RequirementObserver_observe_Results.
type SandstormCore_RequirementObserver_observe_Results_List struct{ capnp.List }

// NewSandstormCore_RequirementObserver_observe_Results creates a new list of SandstormCore_RequirementObserver_observe_Results.
func NewSandstormCore_RequirementObserver_observe_Results_List(s *capnp.Segment, sz int32) (SandstormCore_RequirementObserver_observe_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormCore_RequirementObserver_observe_Results_List{}, err
	}
	return SandstormCore_RequirementObserver_observe_Results_List{l}, nil
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormCore_restore_Params{}, err
	}
	return SandstormCore_restore_Params{st}, nil
}

func NewRootSandstormCore_restore_Params(s *capnp.Segment) (SandstormCore_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormCore_restore_Params{}, err
	}
	return SandstormCore_restore_Params{st}, nil
}

func ReadRootSandstormCore_restore_Params(msg *capnp.Message) (SandstormCore_restore_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_restore_Params{}, err
	}
	return SandstormCore_restore_Params{root.Struct()}, nil
}

func (s SandstormCore_restore_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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

func (s SandstormCore_restore_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s SandstormCore_restore_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormCore_restore_Params) SetRequiredPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// SandstormCore_restore_Params_List is a list of SandstormCore_restore_Params.
type SandstormCore_restore_Params_List struct{ capnp.List }

// NewSandstormCore_restore_Params creates a new list of SandstormCore_restore_Params.
func NewSandstormCore_restore_Params_List(s *capnp.Segment, sz int32) (SandstormCore_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormCore_restore_Params_List{}, err
	}
	return SandstormCore_restore_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_restore_Results{}, err
	}
	return SandstormCore_restore_Results{st}, nil
}

func NewRootSandstormCore_restore_Results(s *capnp.Segment) (SandstormCore_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_restore_Results{}, err
	}
	return SandstormCore_restore_Results{st}, nil
}

func ReadRootSandstormCore_restore_Results(msg *capnp.Message) (SandstormCore_restore_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_restore_Results{}, err
	}
	return SandstormCore_restore_Results{root.Struct()}, nil
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
	if err != nil {
		return SandstormCore_restore_Results_List{}, err
	}
	return SandstormCore_restore_Results_List{l}, nil
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
	if err != nil {
		return SandstormCore_makeToken_Params{}, err
	}
	return SandstormCore_makeToken_Params{st}, nil
}

func NewRootSandstormCore_makeToken_Params(s *capnp.Segment) (SandstormCore_makeToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SandstormCore_makeToken_Params{}, err
	}
	return SandstormCore_makeToken_Params{st}, nil
}

func ReadRootSandstormCore_makeToken_Params(msg *capnp.Message) (SandstormCore_makeToken_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_makeToken_Params{}, err
	}
	return SandstormCore_makeToken_Params{root.Struct()}, nil
}

func (s SandstormCore_makeToken_Params) Ref() (SupervisorObjectId, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return SupervisorObjectId{}, err
	}

	return SupervisorObjectId{Struct: p.Struct()}, nil

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
	if err != nil {
		return ApiTokenOwner{}, err
	}

	return ApiTokenOwner{Struct: p.Struct()}, nil

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
	if err != nil {
		return MembraneRequirement_List{}, err
	}

	return MembraneRequirement_List{List: p.List()}, nil

}

func (s SandstormCore_makeToken_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeToken_Params) SetRequirements(v MembraneRequirement_List) error {

	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// SandstormCore_makeToken_Params_List is a list of SandstormCore_makeToken_Params.
type SandstormCore_makeToken_Params_List struct{ capnp.List }

// NewSandstormCore_makeToken_Params creates a new list of SandstormCore_makeToken_Params.
func NewSandstormCore_makeToken_Params_List(s *capnp.Segment, sz int32) (SandstormCore_makeToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return SandstormCore_makeToken_Params_List{}, err
	}
	return SandstormCore_makeToken_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_makeToken_Results{}, err
	}
	return SandstormCore_makeToken_Results{st}, nil
}

func NewRootSandstormCore_makeToken_Results(s *capnp.Segment) (SandstormCore_makeToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_makeToken_Results{}, err
	}
	return SandstormCore_makeToken_Results{st}, nil
}

func ReadRootSandstormCore_makeToken_Results(msg *capnp.Message) (SandstormCore_makeToken_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_makeToken_Results{}, err
	}
	return SandstormCore_makeToken_Results{root.Struct()}, nil
}

func (s SandstormCore_makeToken_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return SandstormCore_makeToken_Results_List{}, err
	}
	return SandstormCore_makeToken_Results_List{l}, nil
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
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Params{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Params{st}, nil
}

func NewRootSandstormCore_getOwnerNotificationTarget_Params(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Params{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Params{st}, nil
}

func ReadRootSandstormCore_getOwnerNotificationTarget_Params(msg *capnp.Message) (SandstormCore_getOwnerNotificationTarget_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Params{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Params{root.Struct()}, nil
}

// SandstormCore_getOwnerNotificationTarget_Params_List is a list of SandstormCore_getOwnerNotificationTarget_Params.
type SandstormCore_getOwnerNotificationTarget_Params_List struct{ capnp.List }

// NewSandstormCore_getOwnerNotificationTarget_Params creates a new list of SandstormCore_getOwnerNotificationTarget_Params.
func NewSandstormCore_getOwnerNotificationTarget_Params_List(s *capnp.Segment, sz int32) (SandstormCore_getOwnerNotificationTarget_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Params_List{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Results{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Results{st}, nil
}

func NewRootSandstormCore_getOwnerNotificationTarget_Results(s *capnp.Segment) (SandstormCore_getOwnerNotificationTarget_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Results{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Results{st}, nil
}

func ReadRootSandstormCore_getOwnerNotificationTarget_Results(msg *capnp.Message) (SandstormCore_getOwnerNotificationTarget_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Results{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Results{root.Struct()}, nil
}

func (s SandstormCore_getOwnerNotificationTarget_Results) Owner() grain.NotificationTarget {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return grain.NotificationTarget{}
	}
	return grain.NotificationTarget{Client: p.Interface().Client()}
}

func (s SandstormCore_getOwnerNotificationTarget_Results) HasOwner() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_getOwnerNotificationTarget_Results) SetOwner(v grain.NotificationTarget) error {

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

// SandstormCore_getOwnerNotificationTarget_Results_List is a list of SandstormCore_getOwnerNotificationTarget_Results.
type SandstormCore_getOwnerNotificationTarget_Results_List struct{ capnp.List }

// NewSandstormCore_getOwnerNotificationTarget_Results creates a new list of SandstormCore_getOwnerNotificationTarget_Results.
func NewSandstormCore_getOwnerNotificationTarget_Results_List(s *capnp.Segment, sz int32) (SandstormCore_getOwnerNotificationTarget_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormCore_getOwnerNotificationTarget_Results_List{}, err
	}
	return SandstormCore_getOwnerNotificationTarget_Results_List{l}, nil
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

func (p SandstormCore_getOwnerNotificationTarget_Results_Promise) Owner() grain.NotificationTarget {
	return grain.NotificationTarget{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormCore_drop_Params struct{ capnp.Struct }

func NewSandstormCore_drop_Params(s *capnp.Segment) (SandstormCore_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_drop_Params{}, err
	}
	return SandstormCore_drop_Params{st}, nil
}

func NewRootSandstormCore_drop_Params(s *capnp.Segment) (SandstormCore_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_drop_Params{}, err
	}
	return SandstormCore_drop_Params{st}, nil
}

func ReadRootSandstormCore_drop_Params(msg *capnp.Message) (SandstormCore_drop_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_drop_Params{}, err
	}
	return SandstormCore_drop_Params{root.Struct()}, nil
}

func (s SandstormCore_drop_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return SandstormCore_drop_Params_List{}, err
	}
	return SandstormCore_drop_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_drop_Results{}, err
	}
	return SandstormCore_drop_Results{st}, nil
}

func NewRootSandstormCore_drop_Results(s *capnp.Segment) (SandstormCore_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormCore_drop_Results{}, err
	}
	return SandstormCore_drop_Results{st}, nil
}

func ReadRootSandstormCore_drop_Results(msg *capnp.Message) (SandstormCore_drop_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_drop_Results{}, err
	}
	return SandstormCore_drop_Results{root.Struct()}, nil
}

// SandstormCore_drop_Results_List is a list of SandstormCore_drop_Results.
type SandstormCore_drop_Results_List struct{ capnp.List }

// NewSandstormCore_drop_Results creates a new list of SandstormCore_drop_Results.
func NewSandstormCore_drop_Results_List(s *capnp.Segment, sz int32) (SandstormCore_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormCore_drop_Results_List{}, err
	}
	return SandstormCore_drop_Results_List{l}, nil
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
	if err != nil {
		return SandstormCore_checkRequirements_Params{}, err
	}
	return SandstormCore_checkRequirements_Params{st}, nil
}

func NewRootSandstormCore_checkRequirements_Params(s *capnp.Segment) (SandstormCore_checkRequirements_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_checkRequirements_Params{}, err
	}
	return SandstormCore_checkRequirements_Params{st}, nil
}

func ReadRootSandstormCore_checkRequirements_Params(msg *capnp.Message) (SandstormCore_checkRequirements_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_checkRequirements_Params{}, err
	}
	return SandstormCore_checkRequirements_Params{root.Struct()}, nil
}

func (s SandstormCore_checkRequirements_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return MembraneRequirement_List{}, err
	}

	return MembraneRequirement_List{List: p.List()}, nil

}

func (s SandstormCore_checkRequirements_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_checkRequirements_Params) SetRequirements(v MembraneRequirement_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// SandstormCore_checkRequirements_Params_List is a list of SandstormCore_checkRequirements_Params.
type SandstormCore_checkRequirements_Params_List struct{ capnp.List }

// NewSandstormCore_checkRequirements_Params creates a new list of SandstormCore_checkRequirements_Params.
func NewSandstormCore_checkRequirements_Params_List(s *capnp.Segment, sz int32) (SandstormCore_checkRequirements_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormCore_checkRequirements_Params_List{}, err
	}
	return SandstormCore_checkRequirements_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_checkRequirements_Results{}, err
	}
	return SandstormCore_checkRequirements_Results{st}, nil
}

func NewRootSandstormCore_checkRequirements_Results(s *capnp.Segment) (SandstormCore_checkRequirements_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_checkRequirements_Results{}, err
	}
	return SandstormCore_checkRequirements_Results{st}, nil
}

func ReadRootSandstormCore_checkRequirements_Results(msg *capnp.Message) (SandstormCore_checkRequirements_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_checkRequirements_Results{}, err
	}
	return SandstormCore_checkRequirements_Results{root.Struct()}, nil
}

func (s SandstormCore_checkRequirements_Results) Observer() SandstormCore_RequirementObserver {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return SandstormCore_RequirementObserver{}
	}
	return SandstormCore_RequirementObserver{Client: p.Interface().Client()}
}

func (s SandstormCore_checkRequirements_Results) HasObserver() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCore_checkRequirements_Results) SetObserver(v SandstormCore_RequirementObserver) error {

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

// SandstormCore_checkRequirements_Results_List is a list of SandstormCore_checkRequirements_Results.
type SandstormCore_checkRequirements_Results_List struct{ capnp.List }

// NewSandstormCore_checkRequirements_Results creates a new list of SandstormCore_checkRequirements_Results.
func NewSandstormCore_checkRequirements_Results_List(s *capnp.Segment, sz int32) (SandstormCore_checkRequirements_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormCore_checkRequirements_Results_List{}, err
	}
	return SandstormCore_checkRequirements_Results_List{l}, nil
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
	if err != nil {
		return SandstormCore_makeChildToken_Params{}, err
	}
	return SandstormCore_makeChildToken_Params{st}, nil
}

func NewRootSandstormCore_makeChildToken_Params(s *capnp.Segment) (SandstormCore_makeChildToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SandstormCore_makeChildToken_Params{}, err
	}
	return SandstormCore_makeChildToken_Params{st}, nil
}

func ReadRootSandstormCore_makeChildToken_Params(msg *capnp.Message) (SandstormCore_makeChildToken_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_makeChildToken_Params{}, err
	}
	return SandstormCore_makeChildToken_Params{root.Struct()}, nil
}

func (s SandstormCore_makeChildToken_Params) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return ApiTokenOwner{}, err
	}

	return ApiTokenOwner{Struct: p.Struct()}, nil

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
	if err != nil {
		return MembraneRequirement_List{}, err
	}

	return MembraneRequirement_List{List: p.List()}, nil

}

func (s SandstormCore_makeChildToken_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SandstormCore_makeChildToken_Params) SetRequirements(v MembraneRequirement_List) error {

	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// SandstormCore_makeChildToken_Params_List is a list of SandstormCore_makeChildToken_Params.
type SandstormCore_makeChildToken_Params_List struct{ capnp.List }

// NewSandstormCore_makeChildToken_Params creates a new list of SandstormCore_makeChildToken_Params.
func NewSandstormCore_makeChildToken_Params_List(s *capnp.Segment, sz int32) (SandstormCore_makeChildToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return SandstormCore_makeChildToken_Params_List{}, err
	}
	return SandstormCore_makeChildToken_Params_List{l}, nil
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
	if err != nil {
		return SandstormCore_makeChildToken_Results{}, err
	}
	return SandstormCore_makeChildToken_Results{st}, nil
}

func NewRootSandstormCore_makeChildToken_Results(s *capnp.Segment) (SandstormCore_makeChildToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCore_makeChildToken_Results{}, err
	}
	return SandstormCore_makeChildToken_Results{st}, nil
}

func ReadRootSandstormCore_makeChildToken_Results(msg *capnp.Message) (SandstormCore_makeChildToken_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCore_makeChildToken_Results{}, err
	}
	return SandstormCore_makeChildToken_Results{root.Struct()}, nil
}

func (s SandstormCore_makeChildToken_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return SandstormCore_makeChildToken_Results_List{}, err
	}
	return SandstormCore_makeChildToken_Results_List{l}, nil
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

func NewMembraneRequirement(s *capnp.Segment) (MembraneRequirement, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return MembraneRequirement{}, err
	}
	return MembraneRequirement{st}, nil
}

func NewRootMembraneRequirement(s *capnp.Segment) (MembraneRequirement, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return MembraneRequirement{}, err
	}
	return MembraneRequirement{st}, nil
}

func ReadRootMembraneRequirement(msg *capnp.Message) (MembraneRequirement, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return MembraneRequirement{}, err
	}
	return MembraneRequirement{root.Struct()}, nil
}

func (s MembraneRequirement) Which() MembraneRequirement_Which {
	return MembraneRequirement_Which(s.Struct.Uint16(0))
}

func (s MembraneRequirement) TokenValid() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s MembraneRequirement) HasTokenValid() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement) TokenValidBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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

func (s MembraneRequirement) SetPermissionsHeld() { s.Struct.SetUint16(0, 1) }

func (s MembraneRequirement_permissionsHeld) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s MembraneRequirement_permissionsHeld) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s MembraneRequirement_permissionsHeld) SetUserId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s MembraneRequirement_permissionsHeld) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s MembraneRequirement_permissionsHeld) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s MembraneRequirement_permissionsHeld) HasPermissions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement_permissionsHeld) SetPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(2, v.List.ToPtr())
}

func (s MembraneRequirement) UserIsAdmin() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s MembraneRequirement) HasUserIsAdmin() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MembraneRequirement) UserIsAdminBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return MembraneRequirement_List{}, err
	}
	return MembraneRequirement_List{l}, nil
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
	if err != nil {
		return SystemPersistent_addRequirements_Params{}, err
	}
	return SystemPersistent_addRequirements_Params{st}, nil
}

func NewRootSystemPersistent_addRequirements_Params(s *capnp.Segment) (SystemPersistent_addRequirements_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SystemPersistent_addRequirements_Params{}, err
	}
	return SystemPersistent_addRequirements_Params{st}, nil
}

func ReadRootSystemPersistent_addRequirements_Params(msg *capnp.Message) (SystemPersistent_addRequirements_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SystemPersistent_addRequirements_Params{}, err
	}
	return SystemPersistent_addRequirements_Params{root.Struct()}, nil
}

func (s SystemPersistent_addRequirements_Params) Requirements() (MembraneRequirement_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return MembraneRequirement_List{}, err
	}

	return MembraneRequirement_List{List: p.List()}, nil

}

func (s SystemPersistent_addRequirements_Params) HasRequirements() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SystemPersistent_addRequirements_Params) SetRequirements(v MembraneRequirement_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// SystemPersistent_addRequirements_Params_List is a list of SystemPersistent_addRequirements_Params.
type SystemPersistent_addRequirements_Params_List struct{ capnp.List }

// NewSystemPersistent_addRequirements_Params creates a new list of SystemPersistent_addRequirements_Params.
func NewSystemPersistent_addRequirements_Params_List(s *capnp.Segment, sz int32) (SystemPersistent_addRequirements_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SystemPersistent_addRequirements_Params_List{}, err
	}
	return SystemPersistent_addRequirements_Params_List{l}, nil
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
	if err != nil {
		return SystemPersistent_addRequirements_Results{}, err
	}
	return SystemPersistent_addRequirements_Results{st}, nil
}

func NewRootSystemPersistent_addRequirements_Results(s *capnp.Segment) (SystemPersistent_addRequirements_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SystemPersistent_addRequirements_Results{}, err
	}
	return SystemPersistent_addRequirements_Results{st}, nil
}

func ReadRootSystemPersistent_addRequirements_Results(msg *capnp.Message) (SystemPersistent_addRequirements_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SystemPersistent_addRequirements_Results{}, err
	}
	return SystemPersistent_addRequirements_Results{root.Struct()}, nil
}

func (s SystemPersistent_addRequirements_Results) Cap() SystemPersistent {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return SystemPersistent{}
	}
	return SystemPersistent{Client: p.Interface().Client()}
}

func (s SystemPersistent_addRequirements_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SystemPersistent_addRequirements_Results) SetCap(v SystemPersistent) error {

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

// SystemPersistent_addRequirements_Results_List is a list of SystemPersistent_addRequirements_Results.
type SystemPersistent_addRequirements_Results_List struct{ capnp.List }

// NewSystemPersistent_addRequirements_Results creates a new list of SystemPersistent_addRequirements_Results.
func NewSystemPersistent_addRequirements_Results_List(s *capnp.Segment, sz int32) (SystemPersistent_addRequirements_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SystemPersistent_addRequirements_Results_List{}, err
	}
	return SystemPersistent_addRequirements_Results_List{l}, nil
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

func (c PersistentOngoingNotification) Cancel(ctx context.Context, params func(grain.OngoingNotification_cancel_Params) error, opts ...capnp.CallOption) grain.OngoingNotification_cancel_Results_Promise {
	if c.Client == nil {
		return grain.OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xfe851ddbb88940cd,
			MethodID:      0,
			InterfaceName: "grain.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.OngoingNotification_cancel_Params{Struct: s}) }
	}
	return grain.OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type PersistentOngoingNotification_Server interface {
	AddRequirements(SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error

	Cancel(grain.OngoingNotification_cancel) error
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
			InterfaceName: "grain.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.OngoingNotification_cancel{c, opts, grain.OngoingNotification_cancel_Params{Struct: p}, grain.OngoingNotification_cancel_Results{Struct: r}}
			return s.Cancel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

type ApiTokenOwner struct{ capnp.Struct }
type ApiTokenOwner_grain ApiTokenOwner
type ApiTokenOwner_user ApiTokenOwner
type ApiTokenOwner_Which uint16

const (
	ApiTokenOwner_Which_webkey   ApiTokenOwner_Which = 0
	ApiTokenOwner_Which_grain    ApiTokenOwner_Which = 1
	ApiTokenOwner_Which_internet ApiTokenOwner_Which = 2
	ApiTokenOwner_Which_frontend ApiTokenOwner_Which = 3
	ApiTokenOwner_Which_user     ApiTokenOwner_Which = 4
)

func (w ApiTokenOwner_Which) String() string {
	const s = "webkeygraininternetfrontenduser"
	switch w {
	case ApiTokenOwner_Which_webkey:
		return s[0:6]
	case ApiTokenOwner_Which_grain:
		return s[6:11]
	case ApiTokenOwner_Which_internet:
		return s[11:19]
	case ApiTokenOwner_Which_frontend:
		return s[19:27]
	case ApiTokenOwner_Which_user:
		return s[27:31]

	}
	return "ApiTokenOwner_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewApiTokenOwner(s *capnp.Segment) (ApiTokenOwner, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return ApiTokenOwner{}, err
	}
	return ApiTokenOwner{st}, nil
}

func NewRootApiTokenOwner(s *capnp.Segment) (ApiTokenOwner, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return ApiTokenOwner{}, err
	}
	return ApiTokenOwner{st}, nil
}

func ReadRootApiTokenOwner(msg *capnp.Message) (ApiTokenOwner, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ApiTokenOwner{}, err
	}
	return ApiTokenOwner{root.Struct()}, nil
}

func (s ApiTokenOwner) Which() ApiTokenOwner_Which {
	return ApiTokenOwner_Which(s.Struct.Uint16(0))
}

func (s ApiTokenOwner) SetWebkey() {
	s.Struct.SetUint16(0, 0)
}
func (s ApiTokenOwner) Grain() ApiTokenOwner_grain { return ApiTokenOwner_grain(s) }

func (s ApiTokenOwner) SetGrain() { s.Struct.SetUint16(0, 1) }

func (s ApiTokenOwner_grain) GrainId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s ApiTokenOwner_grain) HasGrainId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

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

func (s ApiTokenOwner_grain) IntroducerUser() (string, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s ApiTokenOwner_grain) HasIntroducerUser() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_grain) IntroducerUserBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s ApiTokenOwner_grain) SetIntroducerUser(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
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

func (s ApiTokenOwner) SetUser() { s.Struct.SetUint16(0, 4) }

func (s ApiTokenOwner_user) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s ApiTokenOwner_user) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s ApiTokenOwner_user) SetUserId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ApiTokenOwner_user) Title() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s ApiTokenOwner_user) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ApiTokenOwner_user) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s ApiTokenOwner_user) SetTitle(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// ApiTokenOwner_List is a list of ApiTokenOwner.
type ApiTokenOwner_List struct{ capnp.List }

// NewApiTokenOwner creates a new list of ApiTokenOwner.
func NewApiTokenOwner_List(s *capnp.Segment, sz int32) (ApiTokenOwner_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return ApiTokenOwner_List{}, err
	}
	return ApiTokenOwner_List{l}, nil
}

func (s ApiTokenOwner_List) At(i int) ApiTokenOwner           { return ApiTokenOwner{s.List.Struct(i)} }
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
	if err != nil {
		return SupervisorObjectId{}, err
	}
	return SupervisorObjectId{st}, nil
}

func NewRootSupervisorObjectId(s *capnp.Segment) (SupervisorObjectId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return SupervisorObjectId{}, err
	}
	return SupervisorObjectId{st}, nil
}

func ReadRootSupervisorObjectId(msg *capnp.Message) (SupervisorObjectId, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SupervisorObjectId{}, err
	}
	return SupervisorObjectId{root.Struct()}, nil
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
	if err != nil {
		return SupervisorObjectId_List{}, err
	}
	return SupervisorObjectId_List{l}, nil
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
