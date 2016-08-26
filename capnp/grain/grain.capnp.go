package grain

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	activity "zenhack.net/go/sandstorm/capnp/activity"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	powerbox "zenhack.net/go/sandstorm/capnp/powerbox"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type SandstormApi struct{ Client capnp.Client }

func (c SandstormApi) DeprecatedPublish(ctx context.Context, params func(SandstormApi_deprecatedPublish_Params) error, opts ...capnp.CallOption) SandstormApi_deprecatedPublish_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deprecatedPublish_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      0,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedPublish",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deprecatedPublish_Params{Struct: s}) }
	}
	return SandstormApi_deprecatedPublish_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) DeprecatedRegisterAction(ctx context.Context, params func(SandstormApi_deprecatedRegisterAction_Params) error, opts ...capnp.CallOption) SandstormApi_deprecatedRegisterAction_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deprecatedRegisterAction_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      1,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedRegisterAction",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deprecatedRegisterAction_Params{Struct: s}) }
	}
	return SandstormApi_deprecatedRegisterAction_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) ShareCap(ctx context.Context, params func(SandstormApi_shareCap_Params) error, opts ...capnp.CallOption) SandstormApi_shareCap_Results_Promise {
	if c.Client == nil {
		return SandstormApi_shareCap_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      2,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareCap",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_shareCap_Params{Struct: s}) }
	}
	return SandstormApi_shareCap_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) ShareView(ctx context.Context, params func(SandstormApi_shareView_Params) error, opts ...capnp.CallOption) SandstormApi_shareView_Results_Promise {
	if c.Client == nil {
		return SandstormApi_shareView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      3,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_shareView_Params{Struct: s}) }
	}
	return SandstormApi_shareView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) Restore(ctx context.Context, params func(SandstormApi_restore_Params) error, opts ...capnp.CallOption) SandstormApi_restore_Results_Promise {
	if c.Client == nil {
		return SandstormApi_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      4,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_restore_Params{Struct: s}) }
	}
	return SandstormApi_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) Drop(ctx context.Context, params func(SandstormApi_drop_Params) error, opts ...capnp.CallOption) SandstormApi_drop_Results_Promise {
	if c.Client == nil {
		return SandstormApi_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      5,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_drop_Params{Struct: s}) }
	}
	return SandstormApi_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) Deleted(ctx context.Context, params func(SandstormApi_deleted_Params) error, opts ...capnp.CallOption) SandstormApi_deleted_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deleted_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      6,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deleted",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deleted_Params{Struct: s}) }
	}
	return SandstormApi_deleted_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) StayAwake(ctx context.Context, params func(SandstormApi_stayAwake_Params) error, opts ...capnp.CallOption) SandstormApi_stayAwake_Results_Promise {
	if c.Client == nil {
		return SandstormApi_stayAwake_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      7,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "stayAwake",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_stayAwake_Params{Struct: s}) }
	}
	return SandstormApi_stayAwake_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) Save(ctx context.Context, params func(SandstormApi_save_Params) error, opts ...capnp.CallOption) SandstormApi_save_Results_Promise {
	if c.Client == nil {
		return SandstormApi_save_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      8,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_save_Params{Struct: s}) }
	}
	return SandstormApi_save_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SandstormApi) BackgroundActivity(ctx context.Context, params func(SandstormApi_backgroundActivity_Params) error, opts ...capnp.CallOption) SandstormApi_backgroundActivity_Results_Promise {
	if c.Client == nil {
		return SandstormApi_backgroundActivity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      9,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "backgroundActivity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_backgroundActivity_Params{Struct: s}) }
	}
	return SandstormApi_backgroundActivity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormApi_Server interface {
	DeprecatedPublish(SandstormApi_deprecatedPublish) error

	DeprecatedRegisterAction(SandstormApi_deprecatedRegisterAction) error

	ShareCap(SandstormApi_shareCap) error

	ShareView(SandstormApi_shareView) error

	Restore(SandstormApi_restore) error

	Drop(SandstormApi_drop) error

	Deleted(SandstormApi_deleted) error

	StayAwake(SandstormApi_stayAwake) error

	Save(SandstormApi_save) error

	BackgroundActivity(SandstormApi_backgroundActivity) error
}

func SandstormApi_ServerToClient(s SandstormApi_Server) SandstormApi {
	c, _ := s.(server.Closer)
	return SandstormApi{Client: server.New(SandstormApi_Methods(nil, s), c)}
}

func SandstormApi_Methods(methods []server.Method, s SandstormApi_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 10)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      0,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedPublish",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deprecatedPublish{c, opts, SandstormApi_deprecatedPublish_Params{Struct: p}, SandstormApi_deprecatedPublish_Results{Struct: r}}
			return s.DeprecatedPublish(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      1,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedRegisterAction",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deprecatedRegisterAction{c, opts, SandstormApi_deprecatedRegisterAction_Params{Struct: p}, SandstormApi_deprecatedRegisterAction_Results{Struct: r}}
			return s.DeprecatedRegisterAction(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      2,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareCap",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_shareCap{c, opts, SandstormApi_shareCap_Params{Struct: p}, SandstormApi_shareCap_Results{Struct: r}}
			return s.ShareCap(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      3,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_shareView{c, opts, SandstormApi_shareView_Params{Struct: p}, SandstormApi_shareView_Results{Struct: r}}
			return s.ShareView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      4,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_restore{c, opts, SandstormApi_restore_Params{Struct: p}, SandstormApi_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      5,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_drop{c, opts, SandstormApi_drop_Params{Struct: p}, SandstormApi_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      6,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deleted",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deleted{c, opts, SandstormApi_deleted_Params{Struct: p}, SandstormApi_deleted_Results{Struct: r}}
			return s.Deleted(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      7,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "stayAwake",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_stayAwake{c, opts, SandstormApi_stayAwake_Params{Struct: p}, SandstormApi_stayAwake_Results{Struct: r}}
			return s.StayAwake(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      8,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_save{c, opts, SandstormApi_save_Params{Struct: p}, SandstormApi_save_Results{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      9,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "backgroundActivity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_backgroundActivity{c, opts, SandstormApi_backgroundActivity_Params{Struct: p}, SandstormApi_backgroundActivity_Results{Struct: r}}
			return s.BackgroundActivity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// SandstormApi_deprecatedPublish holds the arguments for a server call to SandstormApi.deprecatedPublish.
type SandstormApi_deprecatedPublish struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deprecatedPublish_Params
	Results SandstormApi_deprecatedPublish_Results
}

// SandstormApi_deprecatedRegisterAction holds the arguments for a server call to SandstormApi.deprecatedRegisterAction.
type SandstormApi_deprecatedRegisterAction struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deprecatedRegisterAction_Params
	Results SandstormApi_deprecatedRegisterAction_Results
}

// SandstormApi_shareCap holds the arguments for a server call to SandstormApi.shareCap.
type SandstormApi_shareCap struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_shareCap_Params
	Results SandstormApi_shareCap_Results
}

// SandstormApi_shareView holds the arguments for a server call to SandstormApi.shareView.
type SandstormApi_shareView struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_shareView_Params
	Results SandstormApi_shareView_Results
}

// SandstormApi_restore holds the arguments for a server call to SandstormApi.restore.
type SandstormApi_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_restore_Params
	Results SandstormApi_restore_Results
}

// SandstormApi_drop holds the arguments for a server call to SandstormApi.drop.
type SandstormApi_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_drop_Params
	Results SandstormApi_drop_Results
}

// SandstormApi_deleted holds the arguments for a server call to SandstormApi.deleted.
type SandstormApi_deleted struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deleted_Params
	Results SandstormApi_deleted_Results
}

// SandstormApi_stayAwake holds the arguments for a server call to SandstormApi.stayAwake.
type SandstormApi_stayAwake struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_stayAwake_Params
	Results SandstormApi_stayAwake_Results
}

// SandstormApi_save holds the arguments for a server call to SandstormApi.save.
type SandstormApi_save struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_save_Params
	Results SandstormApi_save_Results
}

// SandstormApi_backgroundActivity holds the arguments for a server call to SandstormApi.backgroundActivity.
type SandstormApi_backgroundActivity struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_backgroundActivity_Params
	Results SandstormApi_backgroundActivity_Results
}

type SandstormApi_deprecatedPublish_Params struct{ capnp.Struct }

func NewSandstormApi_deprecatedPublish_Params(s *capnp.Segment) (SandstormApi_deprecatedPublish_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedPublish_Params{st}, err
}

func NewRootSandstormApi_deprecatedPublish_Params(s *capnp.Segment) (SandstormApi_deprecatedPublish_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedPublish_Params{st}, err
}

func ReadRootSandstormApi_deprecatedPublish_Params(msg *capnp.Message) (SandstormApi_deprecatedPublish_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deprecatedPublish_Params{root.Struct()}, err
}

func (s SandstormApi_deprecatedPublish_Params) String() string {
	str, _ := text.Marshal(0xa2873a59df6d885c, s.Struct)
	return str
}

// SandstormApi_deprecatedPublish_Params_List is a list of SandstormApi_deprecatedPublish_Params.
type SandstormApi_deprecatedPublish_Params_List struct{ capnp.List }

// NewSandstormApi_deprecatedPublish_Params creates a new list of SandstormApi_deprecatedPublish_Params.
func NewSandstormApi_deprecatedPublish_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedPublish_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_deprecatedPublish_Params_List{l}, err
}

func (s SandstormApi_deprecatedPublish_Params_List) At(i int) SandstormApi_deprecatedPublish_Params {
	return SandstormApi_deprecatedPublish_Params{s.List.Struct(i)}
}

func (s SandstormApi_deprecatedPublish_Params_List) Set(i int, v SandstormApi_deprecatedPublish_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedPublish_Params_Promise is a wrapper for a SandstormApi_deprecatedPublish_Params promised by a client call.
type SandstormApi_deprecatedPublish_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedPublish_Params_Promise) Struct() (SandstormApi_deprecatedPublish_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedPublish_Params{s}, err
}

type SandstormApi_deprecatedPublish_Results struct{ capnp.Struct }

func NewSandstormApi_deprecatedPublish_Results(s *capnp.Segment) (SandstormApi_deprecatedPublish_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedPublish_Results{st}, err
}

func NewRootSandstormApi_deprecatedPublish_Results(s *capnp.Segment) (SandstormApi_deprecatedPublish_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedPublish_Results{st}, err
}

func ReadRootSandstormApi_deprecatedPublish_Results(msg *capnp.Message) (SandstormApi_deprecatedPublish_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deprecatedPublish_Results{root.Struct()}, err
}

func (s SandstormApi_deprecatedPublish_Results) String() string {
	str, _ := text.Marshal(0xb42ccfaaf45a3f7a, s.Struct)
	return str
}

// SandstormApi_deprecatedPublish_Results_List is a list of SandstormApi_deprecatedPublish_Results.
type SandstormApi_deprecatedPublish_Results_List struct{ capnp.List }

// NewSandstormApi_deprecatedPublish_Results creates a new list of SandstormApi_deprecatedPublish_Results.
func NewSandstormApi_deprecatedPublish_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedPublish_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_deprecatedPublish_Results_List{l}, err
}

func (s SandstormApi_deprecatedPublish_Results_List) At(i int) SandstormApi_deprecatedPublish_Results {
	return SandstormApi_deprecatedPublish_Results{s.List.Struct(i)}
}

func (s SandstormApi_deprecatedPublish_Results_List) Set(i int, v SandstormApi_deprecatedPublish_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedPublish_Results_Promise is a wrapper for a SandstormApi_deprecatedPublish_Results promised by a client call.
type SandstormApi_deprecatedPublish_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedPublish_Results_Promise) Struct() (SandstormApi_deprecatedPublish_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedPublish_Results{s}, err
}

type SandstormApi_deprecatedRegisterAction_Params struct{ capnp.Struct }

func NewSandstormApi_deprecatedRegisterAction_Params(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedRegisterAction_Params{st}, err
}

func NewRootSandstormApi_deprecatedRegisterAction_Params(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedRegisterAction_Params{st}, err
}

func ReadRootSandstormApi_deprecatedRegisterAction_Params(msg *capnp.Message) (SandstormApi_deprecatedRegisterAction_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deprecatedRegisterAction_Params{root.Struct()}, err
}

func (s SandstormApi_deprecatedRegisterAction_Params) String() string {
	str, _ := text.Marshal(0xd271034eec62b43b, s.Struct)
	return str
}

// SandstormApi_deprecatedRegisterAction_Params_List is a list of SandstormApi_deprecatedRegisterAction_Params.
type SandstormApi_deprecatedRegisterAction_Params_List struct{ capnp.List }

// NewSandstormApi_deprecatedRegisterAction_Params creates a new list of SandstormApi_deprecatedRegisterAction_Params.
func NewSandstormApi_deprecatedRegisterAction_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedRegisterAction_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_deprecatedRegisterAction_Params_List{l}, err
}

func (s SandstormApi_deprecatedRegisterAction_Params_List) At(i int) SandstormApi_deprecatedRegisterAction_Params {
	return SandstormApi_deprecatedRegisterAction_Params{s.List.Struct(i)}
}

func (s SandstormApi_deprecatedRegisterAction_Params_List) Set(i int, v SandstormApi_deprecatedRegisterAction_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedRegisterAction_Params_Promise is a wrapper for a SandstormApi_deprecatedRegisterAction_Params promised by a client call.
type SandstormApi_deprecatedRegisterAction_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedRegisterAction_Params_Promise) Struct() (SandstormApi_deprecatedRegisterAction_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedRegisterAction_Params{s}, err
}

type SandstormApi_deprecatedRegisterAction_Results struct{ capnp.Struct }

func NewSandstormApi_deprecatedRegisterAction_Results(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedRegisterAction_Results{st}, err
}

func NewRootSandstormApi_deprecatedRegisterAction_Results(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deprecatedRegisterAction_Results{st}, err
}

func ReadRootSandstormApi_deprecatedRegisterAction_Results(msg *capnp.Message) (SandstormApi_deprecatedRegisterAction_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deprecatedRegisterAction_Results{root.Struct()}, err
}

func (s SandstormApi_deprecatedRegisterAction_Results) String() string {
	str, _ := text.Marshal(0xb9d62f4beefefc29, s.Struct)
	return str
}

// SandstormApi_deprecatedRegisterAction_Results_List is a list of SandstormApi_deprecatedRegisterAction_Results.
type SandstormApi_deprecatedRegisterAction_Results_List struct{ capnp.List }

// NewSandstormApi_deprecatedRegisterAction_Results creates a new list of SandstormApi_deprecatedRegisterAction_Results.
func NewSandstormApi_deprecatedRegisterAction_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedRegisterAction_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_deprecatedRegisterAction_Results_List{l}, err
}

func (s SandstormApi_deprecatedRegisterAction_Results_List) At(i int) SandstormApi_deprecatedRegisterAction_Results {
	return SandstormApi_deprecatedRegisterAction_Results{s.List.Struct(i)}
}

func (s SandstormApi_deprecatedRegisterAction_Results_List) Set(i int, v SandstormApi_deprecatedRegisterAction_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedRegisterAction_Results_Promise is a wrapper for a SandstormApi_deprecatedRegisterAction_Results promised by a client call.
type SandstormApi_deprecatedRegisterAction_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedRegisterAction_Results_Promise) Struct() (SandstormApi_deprecatedRegisterAction_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedRegisterAction_Results{s}, err
}

type SandstormApi_shareCap_Params struct{ capnp.Struct }

func NewSandstormApi_shareCap_Params(s *capnp.Segment) (SandstormApi_shareCap_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareCap_Params{st}, err
}

func NewRootSandstormApi_shareCap_Params(s *capnp.Segment) (SandstormApi_shareCap_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareCap_Params{st}, err
}

func ReadRootSandstormApi_shareCap_Params(msg *capnp.Message) (SandstormApi_shareCap_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_shareCap_Params{root.Struct()}, err
}

func (s SandstormApi_shareCap_Params) String() string {
	str, _ := text.Marshal(0xeb3c29aff080ec3e, s.Struct)
	return str
}

func (s SandstormApi_shareCap_Params) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormApi_shareCap_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Params) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormApi_shareCap_Params) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_shareCap_Params) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_shareCap_Params) DisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(1)
	return powerbox.PowerboxDisplayInfo{Struct: p.Struct()}, err
}

func (s SandstormApi_shareCap_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Params) SetDisplayInfo(v powerbox.PowerboxDisplayInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated powerbox.PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SandstormApi_shareCap_Params) NewDisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	ss, err := powerbox.NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SandstormApi_shareCap_Params_List is a list of SandstormApi_shareCap_Params.
type SandstormApi_shareCap_Params_List struct{ capnp.List }

// NewSandstormApi_shareCap_Params creates a new list of SandstormApi_shareCap_Params.
func NewSandstormApi_shareCap_Params_List(s *capnp.Segment, sz int32) (SandstormApi_shareCap_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormApi_shareCap_Params_List{l}, err
}

func (s SandstormApi_shareCap_Params_List) At(i int) SandstormApi_shareCap_Params {
	return SandstormApi_shareCap_Params{s.List.Struct(i)}
}

func (s SandstormApi_shareCap_Params_List) Set(i int, v SandstormApi_shareCap_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareCap_Params_Promise is a wrapper for a SandstormApi_shareCap_Params promised by a client call.
type SandstormApi_shareCap_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareCap_Params_Promise) Struct() (SandstormApi_shareCap_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareCap_Params{s}, err
}

func (p SandstormApi_shareCap_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_shareCap_Params_Promise) DisplayInfo() powerbox.PowerboxDisplayInfo_Promise {
	return powerbox.PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormApi_shareCap_Results struct{ capnp.Struct }

func NewSandstormApi_shareCap_Results(s *capnp.Segment) (SandstormApi_shareCap_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareCap_Results{st}, err
}

func NewRootSandstormApi_shareCap_Results(s *capnp.Segment) (SandstormApi_shareCap_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareCap_Results{st}, err
}

func ReadRootSandstormApi_shareCap_Results(msg *capnp.Message) (SandstormApi_shareCap_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_shareCap_Results{root.Struct()}, err
}

func (s SandstormApi_shareCap_Results) String() string {
	str, _ := text.Marshal(0xb96fc5fb8137a705, s.Struct)
	return str
}

func (s SandstormApi_shareCap_Results) SharedCap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormApi_shareCap_Results) HasSharedCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Results) SharedCapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormApi_shareCap_Results) SetSharedCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_shareCap_Results) SetSharedCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_shareCap_Results) Link() SharingLink {
	p, _ := s.Struct.Ptr(1)
	return SharingLink{Client: p.Interface().Client()}
}

func (s SandstormApi_shareCap_Results) HasLink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Results) SetLink(v SharingLink) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_shareCap_Results_List is a list of SandstormApi_shareCap_Results.
type SandstormApi_shareCap_Results_List struct{ capnp.List }

// NewSandstormApi_shareCap_Results creates a new list of SandstormApi_shareCap_Results.
func NewSandstormApi_shareCap_Results_List(s *capnp.Segment, sz int32) (SandstormApi_shareCap_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormApi_shareCap_Results_List{l}, err
}

func (s SandstormApi_shareCap_Results_List) At(i int) SandstormApi_shareCap_Results {
	return SandstormApi_shareCap_Results{s.List.Struct(i)}
}

func (s SandstormApi_shareCap_Results_List) Set(i int, v SandstormApi_shareCap_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareCap_Results_Promise is a wrapper for a SandstormApi_shareCap_Results promised by a client call.
type SandstormApi_shareCap_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareCap_Results_Promise) Struct() (SandstormApi_shareCap_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareCap_Results{s}, err
}

func (p SandstormApi_shareCap_Results_Promise) SharedCap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_shareCap_Results_Promise) Link() SharingLink {
	return SharingLink{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_shareView_Params struct{ capnp.Struct }

func NewSandstormApi_shareView_Params(s *capnp.Segment) (SandstormApi_shareView_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_shareView_Params{st}, err
}

func NewRootSandstormApi_shareView_Params(s *capnp.Segment) (SandstormApi_shareView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_shareView_Params{st}, err
}

func ReadRootSandstormApi_shareView_Params(msg *capnp.Message) (SandstormApi_shareView_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_shareView_Params{root.Struct()}, err
}

func (s SandstormApi_shareView_Params) String() string {
	str, _ := text.Marshal(0xb1e3f6ac609eb4d7, s.Struct)
	return str
}

func (s SandstormApi_shareView_Params) View() UiView {
	p, _ := s.Struct.Ptr(0)
	return UiView{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Params) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Params) SetView(v UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormApi_shareView_Params_List is a list of SandstormApi_shareView_Params.
type SandstormApi_shareView_Params_List struct{ capnp.List }

// NewSandstormApi_shareView_Params creates a new list of SandstormApi_shareView_Params.
func NewSandstormApi_shareView_Params_List(s *capnp.Segment, sz int32) (SandstormApi_shareView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_shareView_Params_List{l}, err
}

func (s SandstormApi_shareView_Params_List) At(i int) SandstormApi_shareView_Params {
	return SandstormApi_shareView_Params{s.List.Struct(i)}
}

func (s SandstormApi_shareView_Params_List) Set(i int, v SandstormApi_shareView_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareView_Params_Promise is a wrapper for a SandstormApi_shareView_Params promised by a client call.
type SandstormApi_shareView_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareView_Params_Promise) Struct() (SandstormApi_shareView_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareView_Params{s}, err
}

func (p SandstormApi_shareView_Params_Promise) View() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormApi_shareView_Results struct{ capnp.Struct }

func NewSandstormApi_shareView_Results(s *capnp.Segment) (SandstormApi_shareView_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareView_Results{st}, err
}

func NewRootSandstormApi_shareView_Results(s *capnp.Segment) (SandstormApi_shareView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_shareView_Results{st}, err
}

func ReadRootSandstormApi_shareView_Results(msg *capnp.Message) (SandstormApi_shareView_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_shareView_Results{root.Struct()}, err
}

func (s SandstormApi_shareView_Results) String() string {
	str, _ := text.Marshal(0xe6abbf843a84f35d, s.Struct)
	return str
}

func (s SandstormApi_shareView_Results) SharedView() UiView {
	p, _ := s.Struct.Ptr(0)
	return UiView{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Results) HasSharedView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Results) SetSharedView(v UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s SandstormApi_shareView_Results) Link() ViewSharingLink {
	p, _ := s.Struct.Ptr(1)
	return ViewSharingLink{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Results) HasLink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Results) SetLink(v ViewSharingLink) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_shareView_Results_List is a list of SandstormApi_shareView_Results.
type SandstormApi_shareView_Results_List struct{ capnp.List }

// NewSandstormApi_shareView_Results creates a new list of SandstormApi_shareView_Results.
func NewSandstormApi_shareView_Results_List(s *capnp.Segment, sz int32) (SandstormApi_shareView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormApi_shareView_Results_List{l}, err
}

func (s SandstormApi_shareView_Results_List) At(i int) SandstormApi_shareView_Results {
	return SandstormApi_shareView_Results{s.List.Struct(i)}
}

func (s SandstormApi_shareView_Results_List) Set(i int, v SandstormApi_shareView_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareView_Results_Promise is a wrapper for a SandstormApi_shareView_Results promised by a client call.
type SandstormApi_shareView_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareView_Results_Promise) Struct() (SandstormApi_shareView_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareView_Results{s}, err
}

func (p SandstormApi_shareView_Results_Promise) SharedView() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p SandstormApi_shareView_Results_Promise) Link() ViewSharingLink {
	return ViewSharingLink{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_restore_Params struct{ capnp.Struct }

func NewSandstormApi_restore_Params(s *capnp.Segment) (SandstormApi_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_restore_Params{st}, err
}

func NewRootSandstormApi_restore_Params(s *capnp.Segment) (SandstormApi_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_restore_Params{st}, err
}

func ReadRootSandstormApi_restore_Params(msg *capnp.Message) (SandstormApi_restore_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_restore_Params{root.Struct()}, err
}

func (s SandstormApi_restore_Params) String() string {
	str, _ := text.Marshal(0xd29e9db5843719f0, s.Struct)
	return str
}

func (s SandstormApi_restore_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormApi_restore_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_restore_Params) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_restore_Params_List is a list of SandstormApi_restore_Params.
type SandstormApi_restore_Params_List struct{ capnp.List }

// NewSandstormApi_restore_Params creates a new list of SandstormApi_restore_Params.
func NewSandstormApi_restore_Params_List(s *capnp.Segment, sz int32) (SandstormApi_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_restore_Params_List{l}, err
}

func (s SandstormApi_restore_Params_List) At(i int) SandstormApi_restore_Params {
	return SandstormApi_restore_Params{s.List.Struct(i)}
}

func (s SandstormApi_restore_Params_List) Set(i int, v SandstormApi_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_restore_Params_Promise is a wrapper for a SandstormApi_restore_Params promised by a client call.
type SandstormApi_restore_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_restore_Params_Promise) Struct() (SandstormApi_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_restore_Params{s}, err
}

type SandstormApi_restore_Results struct{ capnp.Struct }

func NewSandstormApi_restore_Results(s *capnp.Segment) (SandstormApi_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_restore_Results{st}, err
}

func NewRootSandstormApi_restore_Results(s *capnp.Segment) (SandstormApi_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_restore_Results{st}, err
}

func ReadRootSandstormApi_restore_Results(msg *capnp.Message) (SandstormApi_restore_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_restore_Results{root.Struct()}, err
}

func (s SandstormApi_restore_Results) String() string {
	str, _ := text.Marshal(0xecf1f14c4209c731, s.Struct)
	return str
}

func (s SandstormApi_restore_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormApi_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_restore_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormApi_restore_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_restore_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SandstormApi_restore_Results_List is a list of SandstormApi_restore_Results.
type SandstormApi_restore_Results_List struct{ capnp.List }

// NewSandstormApi_restore_Results creates a new list of SandstormApi_restore_Results.
func NewSandstormApi_restore_Results_List(s *capnp.Segment, sz int32) (SandstormApi_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_restore_Results_List{l}, err
}

func (s SandstormApi_restore_Results_List) At(i int) SandstormApi_restore_Results {
	return SandstormApi_restore_Results{s.List.Struct(i)}
}

func (s SandstormApi_restore_Results_List) Set(i int, v SandstormApi_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_restore_Results_Promise is a wrapper for a SandstormApi_restore_Results promised by a client call.
type SandstormApi_restore_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_restore_Results_Promise) Struct() (SandstormApi_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_restore_Results{s}, err
}

func (p SandstormApi_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormApi_drop_Params struct{ capnp.Struct }

func NewSandstormApi_drop_Params(s *capnp.Segment) (SandstormApi_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_drop_Params{st}, err
}

func NewRootSandstormApi_drop_Params(s *capnp.Segment) (SandstormApi_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_drop_Params{st}, err
}

func ReadRootSandstormApi_drop_Params(msg *capnp.Message) (SandstormApi_drop_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_drop_Params{root.Struct()}, err
}

func (s SandstormApi_drop_Params) String() string {
	str, _ := text.Marshal(0xadac227f85285c65, s.Struct)
	return str
}

func (s SandstormApi_drop_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormApi_drop_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_drop_Params) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_drop_Params_List is a list of SandstormApi_drop_Params.
type SandstormApi_drop_Params_List struct{ capnp.List }

// NewSandstormApi_drop_Params creates a new list of SandstormApi_drop_Params.
func NewSandstormApi_drop_Params_List(s *capnp.Segment, sz int32) (SandstormApi_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_drop_Params_List{l}, err
}

func (s SandstormApi_drop_Params_List) At(i int) SandstormApi_drop_Params {
	return SandstormApi_drop_Params{s.List.Struct(i)}
}

func (s SandstormApi_drop_Params_List) Set(i int, v SandstormApi_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_drop_Params_Promise is a wrapper for a SandstormApi_drop_Params promised by a client call.
type SandstormApi_drop_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_drop_Params_Promise) Struct() (SandstormApi_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_drop_Params{s}, err
}

type SandstormApi_drop_Results struct{ capnp.Struct }

func NewSandstormApi_drop_Results(s *capnp.Segment) (SandstormApi_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_drop_Results{st}, err
}

func NewRootSandstormApi_drop_Results(s *capnp.Segment) (SandstormApi_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_drop_Results{st}, err
}

func ReadRootSandstormApi_drop_Results(msg *capnp.Message) (SandstormApi_drop_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_drop_Results{root.Struct()}, err
}

func (s SandstormApi_drop_Results) String() string {
	str, _ := text.Marshal(0xfbbc20367c72bc59, s.Struct)
	return str
}

// SandstormApi_drop_Results_List is a list of SandstormApi_drop_Results.
type SandstormApi_drop_Results_List struct{ capnp.List }

// NewSandstormApi_drop_Results creates a new list of SandstormApi_drop_Results.
func NewSandstormApi_drop_Results_List(s *capnp.Segment, sz int32) (SandstormApi_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_drop_Results_List{l}, err
}

func (s SandstormApi_drop_Results_List) At(i int) SandstormApi_drop_Results {
	return SandstormApi_drop_Results{s.List.Struct(i)}
}

func (s SandstormApi_drop_Results_List) Set(i int, v SandstormApi_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_drop_Results_Promise is a wrapper for a SandstormApi_drop_Results promised by a client call.
type SandstormApi_drop_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_drop_Results_Promise) Struct() (SandstormApi_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_drop_Results{s}, err
}

type SandstormApi_deleted_Params struct{ capnp.Struct }

func NewSandstormApi_deleted_Params(s *capnp.Segment) (SandstormApi_deleted_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_deleted_Params{st}, err
}

func NewRootSandstormApi_deleted_Params(s *capnp.Segment) (SandstormApi_deleted_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_deleted_Params{st}, err
}

func ReadRootSandstormApi_deleted_Params(msg *capnp.Message) (SandstormApi_deleted_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deleted_Params{root.Struct()}, err
}

func (s SandstormApi_deleted_Params) String() string {
	str, _ := text.Marshal(0x87d94955ce3c61dd, s.Struct)
	return str
}

func (s SandstormApi_deleted_Params) Ref() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormApi_deleted_Params) HasRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_deleted_Params) RefPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormApi_deleted_Params) SetRef(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_deleted_Params) SetRefPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SandstormApi_deleted_Params_List is a list of SandstormApi_deleted_Params.
type SandstormApi_deleted_Params_List struct{ capnp.List }

// NewSandstormApi_deleted_Params creates a new list of SandstormApi_deleted_Params.
func NewSandstormApi_deleted_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deleted_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_deleted_Params_List{l}, err
}

func (s SandstormApi_deleted_Params_List) At(i int) SandstormApi_deleted_Params {
	return SandstormApi_deleted_Params{s.List.Struct(i)}
}

func (s SandstormApi_deleted_Params_List) Set(i int, v SandstormApi_deleted_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deleted_Params_Promise is a wrapper for a SandstormApi_deleted_Params promised by a client call.
type SandstormApi_deleted_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deleted_Params_Promise) Struct() (SandstormApi_deleted_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deleted_Params{s}, err
}

func (p SandstormApi_deleted_Params_Promise) Ref() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormApi_deleted_Results struct{ capnp.Struct }

func NewSandstormApi_deleted_Results(s *capnp.Segment) (SandstormApi_deleted_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deleted_Results{st}, err
}

func NewRootSandstormApi_deleted_Results(s *capnp.Segment) (SandstormApi_deleted_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_deleted_Results{st}, err
}

func ReadRootSandstormApi_deleted_Results(msg *capnp.Message) (SandstormApi_deleted_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_deleted_Results{root.Struct()}, err
}

func (s SandstormApi_deleted_Results) String() string {
	str, _ := text.Marshal(0xf8fe6b4e94a960f7, s.Struct)
	return str
}

// SandstormApi_deleted_Results_List is a list of SandstormApi_deleted_Results.
type SandstormApi_deleted_Results_List struct{ capnp.List }

// NewSandstormApi_deleted_Results creates a new list of SandstormApi_deleted_Results.
func NewSandstormApi_deleted_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deleted_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_deleted_Results_List{l}, err
}

func (s SandstormApi_deleted_Results_List) At(i int) SandstormApi_deleted_Results {
	return SandstormApi_deleted_Results{s.List.Struct(i)}
}

func (s SandstormApi_deleted_Results_List) Set(i int, v SandstormApi_deleted_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deleted_Results_Promise is a wrapper for a SandstormApi_deleted_Results promised by a client call.
type SandstormApi_deleted_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deleted_Results_Promise) Struct() (SandstormApi_deleted_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deleted_Results{s}, err
}

type SandstormApi_stayAwake_Params struct{ capnp.Struct }

func NewSandstormApi_stayAwake_Params(s *capnp.Segment) (SandstormApi_stayAwake_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_stayAwake_Params{st}, err
}

func NewRootSandstormApi_stayAwake_Params(s *capnp.Segment) (SandstormApi_stayAwake_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_stayAwake_Params{st}, err
}

func ReadRootSandstormApi_stayAwake_Params(msg *capnp.Message) (SandstormApi_stayAwake_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_stayAwake_Params{root.Struct()}, err
}

func (s SandstormApi_stayAwake_Params) String() string {
	str, _ := text.Marshal(0xb469e5d523b89e1b, s.Struct)
	return str
}

func (s SandstormApi_stayAwake_Params) DisplayInfo() (activity.NotificationDisplayInfo, error) {
	p, err := s.Struct.Ptr(0)
	return activity.NotificationDisplayInfo{Struct: p.Struct()}, err
}

func (s SandstormApi_stayAwake_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Params) SetDisplayInfo(v activity.NotificationDisplayInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated activity.NotificationDisplayInfo struct, preferring placement in s's segment.
func (s SandstormApi_stayAwake_Params) NewDisplayInfo() (activity.NotificationDisplayInfo, error) {
	ss, err := activity.NewNotificationDisplayInfo(s.Struct.Segment())
	if err != nil {
		return activity.NotificationDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SandstormApi_stayAwake_Params) Notification() activity.OngoingNotification {
	p, _ := s.Struct.Ptr(1)
	return activity.OngoingNotification{Client: p.Interface().Client()}
}

func (s SandstormApi_stayAwake_Params) HasNotification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Params) SetNotification(v activity.OngoingNotification) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_stayAwake_Params_List is a list of SandstormApi_stayAwake_Params.
type SandstormApi_stayAwake_Params_List struct{ capnp.List }

// NewSandstormApi_stayAwake_Params creates a new list of SandstormApi_stayAwake_Params.
func NewSandstormApi_stayAwake_Params_List(s *capnp.Segment, sz int32) (SandstormApi_stayAwake_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormApi_stayAwake_Params_List{l}, err
}

func (s SandstormApi_stayAwake_Params_List) At(i int) SandstormApi_stayAwake_Params {
	return SandstormApi_stayAwake_Params{s.List.Struct(i)}
}

func (s SandstormApi_stayAwake_Params_List) Set(i int, v SandstormApi_stayAwake_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_stayAwake_Params_Promise is a wrapper for a SandstormApi_stayAwake_Params promised by a client call.
type SandstormApi_stayAwake_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_stayAwake_Params_Promise) Struct() (SandstormApi_stayAwake_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_stayAwake_Params{s}, err
}

func (p SandstormApi_stayAwake_Params_Promise) DisplayInfo() activity.NotificationDisplayInfo_Promise {
	return activity.NotificationDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p SandstormApi_stayAwake_Params_Promise) Notification() activity.OngoingNotification {
	return activity.OngoingNotification{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_stayAwake_Results struct{ capnp.Struct }

func NewSandstormApi_stayAwake_Results(s *capnp.Segment) (SandstormApi_stayAwake_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_stayAwake_Results{st}, err
}

func NewRootSandstormApi_stayAwake_Results(s *capnp.Segment) (SandstormApi_stayAwake_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_stayAwake_Results{st}, err
}

func ReadRootSandstormApi_stayAwake_Results(msg *capnp.Message) (SandstormApi_stayAwake_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_stayAwake_Results{root.Struct()}, err
}

func (s SandstormApi_stayAwake_Results) String() string {
	str, _ := text.Marshal(0x9fd40f92e1eb5d21, s.Struct)
	return str
}

func (s SandstormApi_stayAwake_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s SandstormApi_stayAwake_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormApi_stayAwake_Results_List is a list of SandstormApi_stayAwake_Results.
type SandstormApi_stayAwake_Results_List struct{ capnp.List }

// NewSandstormApi_stayAwake_Results creates a new list of SandstormApi_stayAwake_Results.
func NewSandstormApi_stayAwake_Results_List(s *capnp.Segment, sz int32) (SandstormApi_stayAwake_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_stayAwake_Results_List{l}, err
}

func (s SandstormApi_stayAwake_Results_List) At(i int) SandstormApi_stayAwake_Results {
	return SandstormApi_stayAwake_Results{s.List.Struct(i)}
}

func (s SandstormApi_stayAwake_Results_List) Set(i int, v SandstormApi_stayAwake_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_stayAwake_Results_Promise is a wrapper for a SandstormApi_stayAwake_Results promised by a client call.
type SandstormApi_stayAwake_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_stayAwake_Results_Promise) Struct() (SandstormApi_stayAwake_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_stayAwake_Results{s}, err
}

func (p SandstormApi_stayAwake_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormApi_save_Params struct{ capnp.Struct }

func NewSandstormApi_save_Params(s *capnp.Segment) (SandstormApi_save_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_save_Params{st}, err
}

func NewRootSandstormApi_save_Params(s *capnp.Segment) (SandstormApi_save_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SandstormApi_save_Params{st}, err
}

func ReadRootSandstormApi_save_Params(msg *capnp.Message) (SandstormApi_save_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_save_Params{root.Struct()}, err
}

func (s SandstormApi_save_Params) String() string {
	str, _ := text.Marshal(0xd692a643ba8a1f58, s.Struct)
	return str
}

func (s SandstormApi_save_Params) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SandstormApi_save_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Params) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SandstormApi_save_Params) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_save_Params) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_save_Params) Label() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s SandstormApi_save_Params) HasLabel() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Params) SetLabel(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s SandstormApi_save_Params) NewLabel() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SandstormApi_save_Params_List is a list of SandstormApi_save_Params.
type SandstormApi_save_Params_List struct{ capnp.List }

// NewSandstormApi_save_Params creates a new list of SandstormApi_save_Params.
func NewSandstormApi_save_Params_List(s *capnp.Segment, sz int32) (SandstormApi_save_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SandstormApi_save_Params_List{l}, err
}

func (s SandstormApi_save_Params_List) At(i int) SandstormApi_save_Params {
	return SandstormApi_save_Params{s.List.Struct(i)}
}

func (s SandstormApi_save_Params_List) Set(i int, v SandstormApi_save_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_save_Params_Promise is a wrapper for a SandstormApi_save_Params promised by a client call.
type SandstormApi_save_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_save_Params_Promise) Struct() (SandstormApi_save_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_save_Params{s}, err
}

func (p SandstormApi_save_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_save_Params_Promise) Label() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormApi_save_Results struct{ capnp.Struct }

func NewSandstormApi_save_Results(s *capnp.Segment) (SandstormApi_save_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_save_Results{st}, err
}

func NewRootSandstormApi_save_Results(s *capnp.Segment) (SandstormApi_save_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_save_Results{st}, err
}

func ReadRootSandstormApi_save_Results(msg *capnp.Message) (SandstormApi_save_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_save_Results{root.Struct()}, err
}

func (s SandstormApi_save_Results) String() string {
	str, _ := text.Marshal(0x9206caa8d3e3cc7e, s.Struct)
	return str
}

func (s SandstormApi_save_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormApi_save_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Results) SetToken(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_save_Results_List is a list of SandstormApi_save_Results.
type SandstormApi_save_Results_List struct{ capnp.List }

// NewSandstormApi_save_Results creates a new list of SandstormApi_save_Results.
func NewSandstormApi_save_Results_List(s *capnp.Segment, sz int32) (SandstormApi_save_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_save_Results_List{l}, err
}

func (s SandstormApi_save_Results_List) At(i int) SandstormApi_save_Results {
	return SandstormApi_save_Results{s.List.Struct(i)}
}

func (s SandstormApi_save_Results_List) Set(i int, v SandstormApi_save_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_save_Results_Promise is a wrapper for a SandstormApi_save_Results promised by a client call.
type SandstormApi_save_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_save_Results_Promise) Struct() (SandstormApi_save_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_save_Results{s}, err
}

type SandstormApi_backgroundActivity_Params struct{ capnp.Struct }

func NewSandstormApi_backgroundActivity_Params(s *capnp.Segment) (SandstormApi_backgroundActivity_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_backgroundActivity_Params{st}, err
}

func NewRootSandstormApi_backgroundActivity_Params(s *capnp.Segment) (SandstormApi_backgroundActivity_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_backgroundActivity_Params{st}, err
}

func ReadRootSandstormApi_backgroundActivity_Params(msg *capnp.Message) (SandstormApi_backgroundActivity_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_backgroundActivity_Params{root.Struct()}, err
}

func (s SandstormApi_backgroundActivity_Params) String() string {
	str, _ := text.Marshal(0xec8866df56873858, s.Struct)
	return str
}

func (s SandstormApi_backgroundActivity_Params) Event() (activity.ActivityEvent, error) {
	p, err := s.Struct.Ptr(0)
	return activity.ActivityEvent{Struct: p.Struct()}, err
}

func (s SandstormApi_backgroundActivity_Params) HasEvent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_backgroundActivity_Params) SetEvent(v activity.ActivityEvent) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEvent sets the event field to a newly
// allocated activity.ActivityEvent struct, preferring placement in s's segment.
func (s SandstormApi_backgroundActivity_Params) NewEvent() (activity.ActivityEvent, error) {
	ss, err := activity.NewActivityEvent(s.Struct.Segment())
	if err != nil {
		return activity.ActivityEvent{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// SandstormApi_backgroundActivity_Params_List is a list of SandstormApi_backgroundActivity_Params.
type SandstormApi_backgroundActivity_Params_List struct{ capnp.List }

// NewSandstormApi_backgroundActivity_Params creates a new list of SandstormApi_backgroundActivity_Params.
func NewSandstormApi_backgroundActivity_Params_List(s *capnp.Segment, sz int32) (SandstormApi_backgroundActivity_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_backgroundActivity_Params_List{l}, err
}

func (s SandstormApi_backgroundActivity_Params_List) At(i int) SandstormApi_backgroundActivity_Params {
	return SandstormApi_backgroundActivity_Params{s.List.Struct(i)}
}

func (s SandstormApi_backgroundActivity_Params_List) Set(i int, v SandstormApi_backgroundActivity_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_backgroundActivity_Params_Promise is a wrapper for a SandstormApi_backgroundActivity_Params promised by a client call.
type SandstormApi_backgroundActivity_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_backgroundActivity_Params_Promise) Struct() (SandstormApi_backgroundActivity_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_backgroundActivity_Params{s}, err
}

func (p SandstormApi_backgroundActivity_Params_Promise) Event() activity.ActivityEvent_Promise {
	return activity.ActivityEvent_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type SandstormApi_backgroundActivity_Results struct{ capnp.Struct }

func NewSandstormApi_backgroundActivity_Results(s *capnp.Segment) (SandstormApi_backgroundActivity_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_backgroundActivity_Results{st}, err
}

func NewRootSandstormApi_backgroundActivity_Results(s *capnp.Segment) (SandstormApi_backgroundActivity_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SandstormApi_backgroundActivity_Results{st}, err
}

func ReadRootSandstormApi_backgroundActivity_Results(msg *capnp.Message) (SandstormApi_backgroundActivity_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_backgroundActivity_Results{root.Struct()}, err
}

func (s SandstormApi_backgroundActivity_Results) String() string {
	str, _ := text.Marshal(0xa535ac09456b2870, s.Struct)
	return str
}

// SandstormApi_backgroundActivity_Results_List is a list of SandstormApi_backgroundActivity_Results.
type SandstormApi_backgroundActivity_Results_List struct{ capnp.List }

// NewSandstormApi_backgroundActivity_Results creates a new list of SandstormApi_backgroundActivity_Results.
func NewSandstormApi_backgroundActivity_Results_List(s *capnp.Segment, sz int32) (SandstormApi_backgroundActivity_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SandstormApi_backgroundActivity_Results_List{l}, err
}

func (s SandstormApi_backgroundActivity_Results_List) At(i int) SandstormApi_backgroundActivity_Results {
	return SandstormApi_backgroundActivity_Results{s.List.Struct(i)}
}

func (s SandstormApi_backgroundActivity_Results_List) Set(i int, v SandstormApi_backgroundActivity_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_backgroundActivity_Results_Promise is a wrapper for a SandstormApi_backgroundActivity_Results promised by a client call.
type SandstormApi_backgroundActivity_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_backgroundActivity_Results_Promise) Struct() (SandstormApi_backgroundActivity_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_backgroundActivity_Results{s}, err
}

type StaticAsset struct{ Client capnp.Client }

func (c StaticAsset) GetUrl(ctx context.Context, params func(StaticAsset_getUrl_Params) error, opts ...capnp.CallOption) StaticAsset_getUrl_Results_Promise {
	if c.Client == nil {
		return StaticAsset_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xfabb5e621fa9a23f,
			MethodID:      0,
			InterfaceName: "grain.capnp:StaticAsset",
			MethodName:    "getUrl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(StaticAsset_getUrl_Params{Struct: s}) }
	}
	return StaticAsset_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type StaticAsset_Server interface {
	GetUrl(StaticAsset_getUrl) error
}

func StaticAsset_ServerToClient(s StaticAsset_Server) StaticAsset {
	c, _ := s.(server.Closer)
	return StaticAsset{Client: server.New(StaticAsset_Methods(nil, s), c)}
}

func StaticAsset_Methods(methods []server.Method, s StaticAsset_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfabb5e621fa9a23f,
			MethodID:      0,
			InterfaceName: "grain.capnp:StaticAsset",
			MethodName:    "getUrl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := StaticAsset_getUrl{c, opts, StaticAsset_getUrl_Params{Struct: p}, StaticAsset_getUrl_Results{Struct: r}}
			return s.GetUrl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	return methods
}

// StaticAsset_getUrl holds the arguments for a server call to StaticAsset.getUrl.
type StaticAsset_getUrl struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  StaticAsset_getUrl_Params
	Results StaticAsset_getUrl_Results
}

type StaticAsset_Protocol uint16

// Values of StaticAsset_Protocol.
const (
	StaticAsset_Protocol_https StaticAsset_Protocol = 0
	StaticAsset_Protocol_http  StaticAsset_Protocol = 1
)

// String returns the enum's constant name.
func (c StaticAsset_Protocol) String() string {
	switch c {
	case StaticAsset_Protocol_https:
		return "https"
	case StaticAsset_Protocol_http:
		return "http"

	default:
		return ""
	}
}

// StaticAsset_ProtocolFromString returns the enum value with a name,
// or the zero value if there's no such value.
func StaticAsset_ProtocolFromString(c string) StaticAsset_Protocol {
	switch c {
	case "https":
		return StaticAsset_Protocol_https
	case "http":
		return StaticAsset_Protocol_http

	default:
		return 0
	}
}

type StaticAsset_Protocol_List struct{ capnp.List }

func NewStaticAsset_Protocol_List(s *capnp.Segment, sz int32) (StaticAsset_Protocol_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return StaticAsset_Protocol_List{l.List}, err
}

func (l StaticAsset_Protocol_List) At(i int) StaticAsset_Protocol {
	ul := capnp.UInt16List{List: l.List}
	return StaticAsset_Protocol(ul.At(i))
}

func (l StaticAsset_Protocol_List) Set(i int, v StaticAsset_Protocol) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type StaticAsset_getUrl_Params struct{ capnp.Struct }

func NewStaticAsset_getUrl_Params(s *capnp.Segment) (StaticAsset_getUrl_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StaticAsset_getUrl_Params{st}, err
}

func NewRootStaticAsset_getUrl_Params(s *capnp.Segment) (StaticAsset_getUrl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StaticAsset_getUrl_Params{st}, err
}

func ReadRootStaticAsset_getUrl_Params(msg *capnp.Message) (StaticAsset_getUrl_Params, error) {
	root, err := msg.RootPtr()
	return StaticAsset_getUrl_Params{root.Struct()}, err
}

func (s StaticAsset_getUrl_Params) String() string {
	str, _ := text.Marshal(0xa75ecf12570b2966, s.Struct)
	return str
}

// StaticAsset_getUrl_Params_List is a list of StaticAsset_getUrl_Params.
type StaticAsset_getUrl_Params_List struct{ capnp.List }

// NewStaticAsset_getUrl_Params creates a new list of StaticAsset_getUrl_Params.
func NewStaticAsset_getUrl_Params_List(s *capnp.Segment, sz int32) (StaticAsset_getUrl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return StaticAsset_getUrl_Params_List{l}, err
}

func (s StaticAsset_getUrl_Params_List) At(i int) StaticAsset_getUrl_Params {
	return StaticAsset_getUrl_Params{s.List.Struct(i)}
}

func (s StaticAsset_getUrl_Params_List) Set(i int, v StaticAsset_getUrl_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// StaticAsset_getUrl_Params_Promise is a wrapper for a StaticAsset_getUrl_Params promised by a client call.
type StaticAsset_getUrl_Params_Promise struct{ *capnp.Pipeline }

func (p StaticAsset_getUrl_Params_Promise) Struct() (StaticAsset_getUrl_Params, error) {
	s, err := p.Pipeline.Struct()
	return StaticAsset_getUrl_Params{s}, err
}

type StaticAsset_getUrl_Results struct{ capnp.Struct }

func NewStaticAsset_getUrl_Results(s *capnp.Segment) (StaticAsset_getUrl_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StaticAsset_getUrl_Results{st}, err
}

func NewRootStaticAsset_getUrl_Results(s *capnp.Segment) (StaticAsset_getUrl_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StaticAsset_getUrl_Results{st}, err
}

func ReadRootStaticAsset_getUrl_Results(msg *capnp.Message) (StaticAsset_getUrl_Results, error) {
	root, err := msg.RootPtr()
	return StaticAsset_getUrl_Results{root.Struct()}, err
}

func (s StaticAsset_getUrl_Results) String() string {
	str, _ := text.Marshal(0xa5c3aa75d6b648e2, s.Struct)
	return str
}

func (s StaticAsset_getUrl_Results) Protocol() StaticAsset_Protocol {
	return StaticAsset_Protocol(s.Struct.Uint16(0))
}

func (s StaticAsset_getUrl_Results) SetProtocol(v StaticAsset_Protocol) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s StaticAsset_getUrl_Results) HostPath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s StaticAsset_getUrl_Results) HasHostPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StaticAsset_getUrl_Results) HostPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s StaticAsset_getUrl_Results) SetHostPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// StaticAsset_getUrl_Results_List is a list of StaticAsset_getUrl_Results.
type StaticAsset_getUrl_Results_List struct{ capnp.List }

// NewStaticAsset_getUrl_Results creates a new list of StaticAsset_getUrl_Results.
func NewStaticAsset_getUrl_Results_List(s *capnp.Segment, sz int32) (StaticAsset_getUrl_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return StaticAsset_getUrl_Results_List{l}, err
}

func (s StaticAsset_getUrl_Results_List) At(i int) StaticAsset_getUrl_Results {
	return StaticAsset_getUrl_Results{s.List.Struct(i)}
}

func (s StaticAsset_getUrl_Results_List) Set(i int, v StaticAsset_getUrl_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// StaticAsset_getUrl_Results_Promise is a wrapper for a StaticAsset_getUrl_Results promised by a client call.
type StaticAsset_getUrl_Results_Promise struct{ *capnp.Pipeline }

func (p StaticAsset_getUrl_Results_Promise) Struct() (StaticAsset_getUrl_Results, error) {
	s, err := p.Pipeline.Struct()
	return StaticAsset_getUrl_Results{s}, err
}

type UiView struct{ Client capnp.Client }

func (c UiView) GetViewInfo(ctx context.Context, params func(UiView_getViewInfo_Params) error, opts ...capnp.CallOption) UiView_ViewInfo_Promise {
	if c.Client == nil {
		return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_getViewInfo_Params{Struct: s}) }
	}
	return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c UiView) NewSession(ctx context.Context, params func(UiView_newSession_Params) error, opts ...capnp.CallOption) UiView_newSession_Results_Promise {
	if c.Client == nil {
		return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newSession_Params{Struct: s}) }
	}
	return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c UiView) NewRequestSession(ctx context.Context, params func(UiView_newRequestSession_Params) error, opts ...capnp.CallOption) UiView_newRequestSession_Results_Promise {
	if c.Client == nil {
		return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 5}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newRequestSession_Params{Struct: s}) }
	}
	return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c UiView) NewOfferSession(ctx context.Context, params func(UiView_newOfferSession_Params) error, opts ...capnp.CallOption) UiView_newOfferSession_Results_Promise {
	if c.Client == nil {
		return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 6}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newOfferSession_Params{Struct: s}) }
	}
	return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type UiView_Server interface {
	GetViewInfo(UiView_getViewInfo) error

	NewSession(UiView_newSession) error

	NewRequestSession(UiView_newRequestSession) error

	NewOfferSession(UiView_newOfferSession) error
}

func UiView_ServerToClient(s UiView_Server) UiView {
	c, _ := s.(server.Closer)
	return UiView{Client: server.New(UiView_Methods(nil, s), c)}
}

func UiView_Methods(methods []server.Method, s UiView_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_getViewInfo{c, opts, UiView_getViewInfo_Params{Struct: p}, UiView_ViewInfo{Struct: r}}
			return s.GetViewInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newSession{c, opts, UiView_newSession_Params{Struct: p}, UiView_newSession_Results{Struct: r}}
			return s.NewSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newRequestSession{c, opts, UiView_newRequestSession_Params{Struct: p}, UiView_newRequestSession_Results{Struct: r}}
			return s.NewRequestSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newOfferSession{c, opts, UiView_newOfferSession_Params{Struct: p}, UiView_newOfferSession_Results{Struct: r}}
			return s.NewOfferSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// UiView_getViewInfo holds the arguments for a server call to UiView.getViewInfo.
type UiView_getViewInfo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_getViewInfo_Params
	Results UiView_ViewInfo
}

// UiView_newSession holds the arguments for a server call to UiView.newSession.
type UiView_newSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newSession_Params
	Results UiView_newSession_Results
}

// UiView_newRequestSession holds the arguments for a server call to UiView.newRequestSession.
type UiView_newRequestSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newRequestSession_Params
	Results UiView_newRequestSession_Results
}

// UiView_newOfferSession holds the arguments for a server call to UiView.newOfferSession.
type UiView_newOfferSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newOfferSession_Params
	Results UiView_newOfferSession_Results
}

type UiView_ViewInfo struct{ capnp.Struct }

func NewUiView_ViewInfo(s *capnp.Segment) (UiView_ViewInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 8})
	return UiView_ViewInfo{st}, err
}

func NewRootUiView_ViewInfo(s *capnp.Segment) (UiView_ViewInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 8})
	return UiView_ViewInfo{st}, err
}

func ReadRootUiView_ViewInfo(msg *capnp.Message) (UiView_ViewInfo, error) {
	root, err := msg.RootPtr()
	return UiView_ViewInfo{root.Struct()}, err
}

func (s UiView_ViewInfo) String() string {
	str, _ := text.Marshal(0xbc5e354741a8e665, s.Struct)
	return str
}

func (s UiView_ViewInfo) Permissions() (PermissionDef_List, error) {
	p, err := s.Struct.Ptr(0)
	return PermissionDef_List{List: p.List()}, err
}

func (s UiView_ViewInfo) HasPermissions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetPermissions(v PermissionDef_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated PermissionDef_List, preferring placement in s's segment.
func (s UiView_ViewInfo) NewPermissions(n int32) (PermissionDef_List, error) {
	l, err := NewPermissionDef_List(s.Struct.Segment(), n)
	if err != nil {
		return PermissionDef_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s UiView_ViewInfo) Roles() (RoleDef_List, error) {
	p, err := s.Struct.Ptr(1)
	return RoleDef_List{List: p.List()}, err
}

func (s UiView_ViewInfo) HasRoles() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetRoles(v RoleDef_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRoles sets the roles field to a newly
// allocated RoleDef_List, preferring placement in s's segment.
func (s UiView_ViewInfo) NewRoles(n int32) (RoleDef_List, error) {
	l, err := NewRoleDef_List(s.Struct.Segment(), n)
	if err != nil {
		return RoleDef_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s UiView_ViewInfo) DeniedPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.BitList{List: p.List()}, err
}

func (s UiView_ViewInfo) HasDeniedPermissions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetDeniedPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewDeniedPermissions sets the deniedPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s UiView_ViewInfo) NewDeniedPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s UiView_ViewInfo) MatchRequests() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(3)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s UiView_ViewInfo) HasMatchRequests() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetMatchRequests(v powerbox.PowerboxDescriptor_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewMatchRequests sets the matchRequests field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s UiView_ViewInfo) NewMatchRequests(n int32) (powerbox.PowerboxDescriptor_List, error) {
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s UiView_ViewInfo) MatchOffers() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(4)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s UiView_ViewInfo) HasMatchOffers() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetMatchOffers(v powerbox.PowerboxDescriptor_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewMatchOffers sets the matchOffers field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s UiView_ViewInfo) NewMatchOffers(n int32) (powerbox.PowerboxDescriptor_List, error) {
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s UiView_ViewInfo) AppTitle() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(5)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s UiView_ViewInfo) HasAppTitle() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetAppTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewAppTitle sets the appTitle field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s UiView_ViewInfo) NewAppTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_ViewInfo) GrainIcon() StaticAsset {
	p, _ := s.Struct.Ptr(6)
	return StaticAsset{Client: p.Interface().Client()}
}

func (s UiView_ViewInfo) HasGrainIcon() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetGrainIcon(v StaticAsset) error {
	if v.Client == nil {
		return s.Struct.SetPtr(6, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(6, in.ToPtr())
}

func (s UiView_ViewInfo) EventTypes() (activity.ActivityTypeDef_List, error) {
	p, err := s.Struct.Ptr(7)
	return activity.ActivityTypeDef_List{List: p.List()}, err
}

func (s UiView_ViewInfo) HasEventTypes() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetEventTypes(v activity.ActivityTypeDef_List) error {
	return s.Struct.SetPtr(7, v.List.ToPtr())
}

// NewEventTypes sets the eventTypes field to a newly
// allocated activity.ActivityTypeDef_List, preferring placement in s's segment.
func (s UiView_ViewInfo) NewEventTypes(n int32) (activity.ActivityTypeDef_List, error) {
	l, err := activity.NewActivityTypeDef_List(s.Struct.Segment(), n)
	if err != nil {
		return activity.ActivityTypeDef_List{}, err
	}
	err = s.Struct.SetPtr(7, l.List.ToPtr())
	return l, err
}

// UiView_ViewInfo_List is a list of UiView_ViewInfo.
type UiView_ViewInfo_List struct{ capnp.List }

// NewUiView_ViewInfo creates a new list of UiView_ViewInfo.
func NewUiView_ViewInfo_List(s *capnp.Segment, sz int32) (UiView_ViewInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 8}, sz)
	return UiView_ViewInfo_List{l}, err
}

func (s UiView_ViewInfo_List) At(i int) UiView_ViewInfo { return UiView_ViewInfo{s.List.Struct(i)} }

func (s UiView_ViewInfo_List) Set(i int, v UiView_ViewInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_ViewInfo_Promise is a wrapper for a UiView_ViewInfo promised by a client call.
type UiView_ViewInfo_Promise struct{ *capnp.Pipeline }

func (p UiView_ViewInfo_Promise) Struct() (UiView_ViewInfo, error) {
	s, err := p.Pipeline.Struct()
	return UiView_ViewInfo{s}, err
}

func (p UiView_ViewInfo_Promise) AppTitle() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p UiView_ViewInfo_Promise) GrainIcon() StaticAsset {
	return StaticAsset{Client: p.Pipeline.GetPipeline(6).Client()}
}

type UiView_PowerboxTag struct{ capnp.Struct }

func NewUiView_PowerboxTag(s *capnp.Segment) (UiView_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_PowerboxTag{st}, err
}

func NewRootUiView_PowerboxTag(s *capnp.Segment) (UiView_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_PowerboxTag{st}, err
}

func ReadRootUiView_PowerboxTag(msg *capnp.Message) (UiView_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return UiView_PowerboxTag{root.Struct()}, err
}

func (s UiView_PowerboxTag) String() string {
	str, _ := text.Marshal(0x982790c08b1958ec, s.Struct)
	return str
}

func (s UiView_PowerboxTag) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s UiView_PowerboxTag) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_PowerboxTag) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s UiView_PowerboxTag) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// UiView_PowerboxTag_List is a list of UiView_PowerboxTag.
type UiView_PowerboxTag_List struct{ capnp.List }

// NewUiView_PowerboxTag creates a new list of UiView_PowerboxTag.
func NewUiView_PowerboxTag_List(s *capnp.Segment, sz int32) (UiView_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UiView_PowerboxTag_List{l}, err
}

func (s UiView_PowerboxTag_List) At(i int) UiView_PowerboxTag {
	return UiView_PowerboxTag{s.List.Struct(i)}
}

func (s UiView_PowerboxTag_List) Set(i int, v UiView_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_PowerboxTag_Promise is a wrapper for a UiView_PowerboxTag promised by a client call.
type UiView_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p UiView_PowerboxTag_Promise) Struct() (UiView_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return UiView_PowerboxTag{s}, err
}

type UiView_getViewInfo_Params struct{ capnp.Struct }

func NewUiView_getViewInfo_Params(s *capnp.Segment) (UiView_getViewInfo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return UiView_getViewInfo_Params{st}, err
}

func NewRootUiView_getViewInfo_Params(s *capnp.Segment) (UiView_getViewInfo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return UiView_getViewInfo_Params{st}, err
}

func ReadRootUiView_getViewInfo_Params(msg *capnp.Message) (UiView_getViewInfo_Params, error) {
	root, err := msg.RootPtr()
	return UiView_getViewInfo_Params{root.Struct()}, err
}

func (s UiView_getViewInfo_Params) String() string {
	str, _ := text.Marshal(0x8f2ef49549d64e86, s.Struct)
	return str
}

// UiView_getViewInfo_Params_List is a list of UiView_getViewInfo_Params.
type UiView_getViewInfo_Params_List struct{ capnp.List }

// NewUiView_getViewInfo_Params creates a new list of UiView_getViewInfo_Params.
func NewUiView_getViewInfo_Params_List(s *capnp.Segment, sz int32) (UiView_getViewInfo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return UiView_getViewInfo_Params_List{l}, err
}

func (s UiView_getViewInfo_Params_List) At(i int) UiView_getViewInfo_Params {
	return UiView_getViewInfo_Params{s.List.Struct(i)}
}

func (s UiView_getViewInfo_Params_List) Set(i int, v UiView_getViewInfo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_getViewInfo_Params_Promise is a wrapper for a UiView_getViewInfo_Params promised by a client call.
type UiView_getViewInfo_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_getViewInfo_Params_Promise) Struct() (UiView_getViewInfo_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_getViewInfo_Params{s}, err
}

type UiView_newSession_Params struct{ capnp.Struct }

func NewUiView_newSession_Params(s *capnp.Segment) (UiView_newSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return UiView_newSession_Params{st}, err
}

func NewRootUiView_newSession_Params(s *capnp.Segment) (UiView_newSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return UiView_newSession_Params{st}, err
}

func ReadRootUiView_newSession_Params(msg *capnp.Message) (UiView_newSession_Params, error) {
	root, err := msg.RootPtr()
	return UiView_newSession_Params{root.Struct()}, err
}

func (s UiView_newSession_Params) String() string {
	str, _ := text.Marshal(0xf87a2c5a9f996828, s.Struct)
	return str
}

func (s UiView_newSession_Params) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s UiView_newSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s UiView_newSession_Params) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newSession_Params) Context() SessionContext {
	p, _ := s.Struct.Ptr(1)
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SetContext(v SessionContext) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newSession_Params) SetSessionType(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s UiView_newSession_Params) SessionParams() (capnp.Pointer, error) {
	return s.Struct.Pointer(2)
}

func (s UiView_newSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SessionParamsPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(2)
}

func (s UiView_newSession_Params) SetSessionParams(v capnp.Pointer) error {
	return s.Struct.SetPointer(2, v)
}

func (s UiView_newSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(2, v)
}

func (s UiView_newSession_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s UiView_newSession_Params) HasTabId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SetTabId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

// UiView_newSession_Params_List is a list of UiView_newSession_Params.
type UiView_newSession_Params_List struct{ capnp.List }

// NewUiView_newSession_Params creates a new list of UiView_newSession_Params.
func NewUiView_newSession_Params_List(s *capnp.Segment, sz int32) (UiView_newSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return UiView_newSession_Params_List{l}, err
}

func (s UiView_newSession_Params_List) At(i int) UiView_newSession_Params {
	return UiView_newSession_Params{s.List.Struct(i)}
}

func (s UiView_newSession_Params_List) Set(i int, v UiView_newSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newSession_Params_Promise is a wrapper for a UiView_newSession_Params promised by a client call.
type UiView_newSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newSession_Params_Promise) Struct() (UiView_newSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newSession_Params{s}, err
}

func (p UiView_newSession_Params_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

type UiView_newSession_Results struct{ capnp.Struct }

func NewUiView_newSession_Results(s *capnp.Segment) (UiView_newSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newSession_Results{st}, err
}

func NewRootUiView_newSession_Results(s *capnp.Segment) (UiView_newSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newSession_Results{st}, err
}

func ReadRootUiView_newSession_Results(msg *capnp.Message) (UiView_newSession_Results, error) {
	root, err := msg.RootPtr()
	return UiView_newSession_Results{root.Struct()}, err
}

func (s UiView_newSession_Results) String() string {
	str, _ := text.Marshal(0xa8f4ff97289294c7, s.Struct)
	return str
}

func (s UiView_newSession_Results) Session() UiSession {
	p, _ := s.Struct.Ptr(0)
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Results) SetSession(v UiSession) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newSession_Results_List is a list of UiView_newSession_Results.
type UiView_newSession_Results_List struct{ capnp.List }

// NewUiView_newSession_Results creates a new list of UiView_newSession_Results.
func NewUiView_newSession_Results_List(s *capnp.Segment, sz int32) (UiView_newSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UiView_newSession_Results_List{l}, err
}

func (s UiView_newSession_Results_List) At(i int) UiView_newSession_Results {
	return UiView_newSession_Results{s.List.Struct(i)}
}

func (s UiView_newSession_Results_List) Set(i int, v UiView_newSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newSession_Results_Promise is a wrapper for a UiView_newSession_Results promised by a client call.
type UiView_newSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newSession_Results_Promise) Struct() (UiView_newSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newSession_Results{s}, err
}

func (p UiView_newSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiView_newRequestSession_Params struct{ capnp.Struct }

func NewUiView_newRequestSession_Params(s *capnp.Segment) (UiView_newRequestSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return UiView_newRequestSession_Params{st}, err
}

func NewRootUiView_newRequestSession_Params(s *capnp.Segment) (UiView_newRequestSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return UiView_newRequestSession_Params{st}, err
}

func ReadRootUiView_newRequestSession_Params(msg *capnp.Message) (UiView_newRequestSession_Params, error) {
	root, err := msg.RootPtr()
	return UiView_newRequestSession_Params{root.Struct()}, err
}

func (s UiView_newRequestSession_Params) String() string {
	str, _ := text.Marshal(0xbc193a4219598bcb, s.Struct)
	return str
}

func (s UiView_newRequestSession_Params) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s UiView_newRequestSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s UiView_newRequestSession_Params) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newRequestSession_Params) Context() SessionContext {
	p, _ := s.Struct.Ptr(1)
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newRequestSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetContext(v SessionContext) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newRequestSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newRequestSession_Params) SetSessionType(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s UiView_newRequestSession_Params) SessionParams() (capnp.Pointer, error) {
	return s.Struct.Pointer(2)
}

func (s UiView_newRequestSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SessionParamsPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(2)
}

func (s UiView_newRequestSession_Params) SetSessionParams(v capnp.Pointer) error {
	return s.Struct.SetPointer(2, v)
}

func (s UiView_newRequestSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(2, v)
}

func (s UiView_newRequestSession_Params) RequestInfo() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(3)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s UiView_newRequestSession_Params) HasRequestInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetRequestInfo(v powerbox.PowerboxDescriptor_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewRequestInfo sets the requestInfo field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s UiView_newRequestSession_Params) NewRequestInfo(n int32) (powerbox.PowerboxDescriptor_List, error) {
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s UiView_newRequestSession_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return []byte(p.Data()), err
}

func (s UiView_newRequestSession_Params) HasTabId() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetTabId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, d.List.ToPtr())
}

// UiView_newRequestSession_Params_List is a list of UiView_newRequestSession_Params.
type UiView_newRequestSession_Params_List struct{ capnp.List }

// NewUiView_newRequestSession_Params creates a new list of UiView_newRequestSession_Params.
func NewUiView_newRequestSession_Params_List(s *capnp.Segment, sz int32) (UiView_newRequestSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return UiView_newRequestSession_Params_List{l}, err
}

func (s UiView_newRequestSession_Params_List) At(i int) UiView_newRequestSession_Params {
	return UiView_newRequestSession_Params{s.List.Struct(i)}
}

func (s UiView_newRequestSession_Params_List) Set(i int, v UiView_newRequestSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newRequestSession_Params_Promise is a wrapper for a UiView_newRequestSession_Params promised by a client call.
type UiView_newRequestSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newRequestSession_Params_Promise) Struct() (UiView_newRequestSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newRequestSession_Params{s}, err
}

func (p UiView_newRequestSession_Params_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newRequestSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newRequestSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

type UiView_newRequestSession_Results struct{ capnp.Struct }

func NewUiView_newRequestSession_Results(s *capnp.Segment) (UiView_newRequestSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newRequestSession_Results{st}, err
}

func NewRootUiView_newRequestSession_Results(s *capnp.Segment) (UiView_newRequestSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newRequestSession_Results{st}, err
}

func ReadRootUiView_newRequestSession_Results(msg *capnp.Message) (UiView_newRequestSession_Results, error) {
	root, err := msg.RootPtr()
	return UiView_newRequestSession_Results{root.Struct()}, err
}

func (s UiView_newRequestSession_Results) String() string {
	str, _ := text.Marshal(0xa22a2d1cf9579778, s.Struct)
	return str
}

func (s UiView_newRequestSession_Results) Session() UiSession {
	p, _ := s.Struct.Ptr(0)
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newRequestSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Results) SetSession(v UiSession) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newRequestSession_Results_List is a list of UiView_newRequestSession_Results.
type UiView_newRequestSession_Results_List struct{ capnp.List }

// NewUiView_newRequestSession_Results creates a new list of UiView_newRequestSession_Results.
func NewUiView_newRequestSession_Results_List(s *capnp.Segment, sz int32) (UiView_newRequestSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UiView_newRequestSession_Results_List{l}, err
}

func (s UiView_newRequestSession_Results_List) At(i int) UiView_newRequestSession_Results {
	return UiView_newRequestSession_Results{s.List.Struct(i)}
}

func (s UiView_newRequestSession_Results_List) Set(i int, v UiView_newRequestSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newRequestSession_Results_Promise is a wrapper for a UiView_newRequestSession_Results promised by a client call.
type UiView_newRequestSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newRequestSession_Results_Promise) Struct() (UiView_newRequestSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newRequestSession_Results{s}, err
}

func (p UiView_newRequestSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiView_newOfferSession_Params struct{ capnp.Struct }

func NewUiView_newOfferSession_Params(s *capnp.Segment) (UiView_newOfferSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return UiView_newOfferSession_Params{st}, err
}

func NewRootUiView_newOfferSession_Params(s *capnp.Segment) (UiView_newOfferSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return UiView_newOfferSession_Params{st}, err
}

func ReadRootUiView_newOfferSession_Params(msg *capnp.Message) (UiView_newOfferSession_Params, error) {
	root, err := msg.RootPtr()
	return UiView_newOfferSession_Params{root.Struct()}, err
}

func (s UiView_newOfferSession_Params) String() string {
	str, _ := text.Marshal(0xa53aedb3ce8994df, s.Struct)
	return str
}

func (s UiView_newOfferSession_Params) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s UiView_newOfferSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s UiView_newOfferSession_Params) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newOfferSession_Params) Context() SessionContext {
	p, _ := s.Struct.Ptr(1)
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newOfferSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetContext(v SessionContext) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newOfferSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newOfferSession_Params) SetSessionType(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s UiView_newOfferSession_Params) SessionParams() (capnp.Pointer, error) {
	return s.Struct.Pointer(2)
}

func (s UiView_newOfferSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SessionParamsPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(2)
}

func (s UiView_newOfferSession_Params) SetSessionParams(v capnp.Pointer) error {
	return s.Struct.SetPointer(2, v)
}

func (s UiView_newOfferSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(2, v)
}

func (s UiView_newOfferSession_Params) Offer() (capnp.Pointer, error) {
	return s.Struct.Pointer(3)
}

func (s UiView_newOfferSession_Params) HasOffer() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) OfferPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(3)
}

func (s UiView_newOfferSession_Params) SetOffer(v capnp.Pointer) error {
	return s.Struct.SetPointer(3, v)
}

func (s UiView_newOfferSession_Params) SetOfferPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(3, v)
}

func (s UiView_newOfferSession_Params) Descriptor() (powerbox.PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(4)
	return powerbox.PowerboxDescriptor{Struct: p.Struct()}, err
}

func (s UiView_newOfferSession_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetDescriptor(v powerbox.PowerboxDescriptor) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated powerbox.PowerboxDescriptor struct, preferring placement in s's segment.
func (s UiView_newOfferSession_Params) NewDescriptor() (powerbox.PowerboxDescriptor, error) {
	ss, err := powerbox.NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newOfferSession_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return []byte(p.Data()), err
}

func (s UiView_newOfferSession_Params) HasTabId() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetTabId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, d.List.ToPtr())
}

// UiView_newOfferSession_Params_List is a list of UiView_newOfferSession_Params.
type UiView_newOfferSession_Params_List struct{ capnp.List }

// NewUiView_newOfferSession_Params creates a new list of UiView_newOfferSession_Params.
func NewUiView_newOfferSession_Params_List(s *capnp.Segment, sz int32) (UiView_newOfferSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return UiView_newOfferSession_Params_List{l}, err
}

func (s UiView_newOfferSession_Params_List) At(i int) UiView_newOfferSession_Params {
	return UiView_newOfferSession_Params{s.List.Struct(i)}
}

func (s UiView_newOfferSession_Params_List) Set(i int, v UiView_newOfferSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newOfferSession_Params_Promise is a wrapper for a UiView_newOfferSession_Params promised by a client call.
type UiView_newOfferSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newOfferSession_Params_Promise) Struct() (UiView_newOfferSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newOfferSession_Params{s}, err
}

func (p UiView_newOfferSession_Params_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newOfferSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newOfferSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

func (p UiView_newOfferSession_Params_Promise) Offer() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(3)
}

func (p UiView_newOfferSession_Params_Promise) Descriptor() powerbox.PowerboxDescriptor_Promise {
	return powerbox.PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

type UiView_newOfferSession_Results struct{ capnp.Struct }

func NewUiView_newOfferSession_Results(s *capnp.Segment) (UiView_newOfferSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newOfferSession_Results{st}, err
}

func NewRootUiView_newOfferSession_Results(s *capnp.Segment) (UiView_newOfferSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UiView_newOfferSession_Results{st}, err
}

func ReadRootUiView_newOfferSession_Results(msg *capnp.Message) (UiView_newOfferSession_Results, error) {
	root, err := msg.RootPtr()
	return UiView_newOfferSession_Results{root.Struct()}, err
}

func (s UiView_newOfferSession_Results) String() string {
	str, _ := text.Marshal(0x9eb6708c01ec2079, s.Struct)
	return str
}

func (s UiView_newOfferSession_Results) Session() UiSession {
	p, _ := s.Struct.Ptr(0)
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newOfferSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Results) SetSession(v UiSession) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newOfferSession_Results_List is a list of UiView_newOfferSession_Results.
type UiView_newOfferSession_Results_List struct{ capnp.List }

// NewUiView_newOfferSession_Results creates a new list of UiView_newOfferSession_Results.
func NewUiView_newOfferSession_Results_List(s *capnp.Segment, sz int32) (UiView_newOfferSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UiView_newOfferSession_Results_List{l}, err
}

func (s UiView_newOfferSession_Results_List) At(i int) UiView_newOfferSession_Results {
	return UiView_newOfferSession_Results{s.List.Struct(i)}
}

func (s UiView_newOfferSession_Results_List) Set(i int, v UiView_newOfferSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newOfferSession_Results_Promise is a wrapper for a UiView_newOfferSession_Results promised by a client call.
type UiView_newOfferSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newOfferSession_Results_Promise) Struct() (UiView_newOfferSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newOfferSession_Results{s}, err
}

func (p UiView_newOfferSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiSession struct{ Client capnp.Client }

type UiSession_Server interface {
}

func UiSession_ServerToClient(s UiSession_Server) UiSession {
	c, _ := s.(server.Closer)
	return UiSession{Client: server.New(UiSession_Methods(nil, s), c)}
}

func UiSession_Methods(methods []server.Method, s UiSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type SessionContext struct{ Client capnp.Client }

func (c SessionContext) GetSharedPermissions(ctx context.Context, params func(SessionContext_getSharedPermissions_Params) error, opts ...capnp.CallOption) SessionContext_getSharedPermissions_Results_Promise {
	if c.Client == nil {
		return SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_getSharedPermissions_Params{Struct: s}) }
	}
	return SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) TieToUser(ctx context.Context, params func(SessionContext_tieToUser_Params) error, opts ...capnp.CallOption) SessionContext_tieToUser_Results_Promise {
	if c.Client == nil {
		return SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_tieToUser_Params{Struct: s}) }
	}
	return SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) Offer(ctx context.Context, params func(SessionContext_offer_Params) error, opts ...capnp.CallOption) SessionContext_offer_Results_Promise {
	if c.Client == nil {
		return SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_offer_Params{Struct: s}) }
	}
	return SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) Request(ctx context.Context, params func(SessionContext_request_Params) error, opts ...capnp.CallOption) SessionContext_request_Results_Promise {
	if c.Client == nil {
		return SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_request_Params{Struct: s}) }
	}
	return SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) FulfillRequest(ctx context.Context, params func(SessionContext_fulfillRequest_Params) error, opts ...capnp.CallOption) SessionContext_fulfillRequest_Results_Promise {
	if c.Client == nil {
		return SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_fulfillRequest_Params{Struct: s}) }
	}
	return SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) Close(ctx context.Context, params func(SessionContext_close_Params) error, opts ...capnp.CallOption) SessionContext_close_Results_Promise {
	if c.Client == nil {
		return SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_close_Params{Struct: s}) }
	}
	return SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) OpenView(ctx context.Context, params func(SessionContext_openView_Params) error, opts ...capnp.CallOption) SessionContext_openView_Results_Promise {
	if c.Client == nil {
		return SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_openView_Params{Struct: s}) }
	}
	return SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) ClaimRequest(ctx context.Context, params func(SessionContext_claimRequest_Params) error, opts ...capnp.CallOption) SessionContext_claimRequest_Results_Promise {
	if c.Client == nil {
		return SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_claimRequest_Params{Struct: s}) }
	}
	return SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c SessionContext) Activity(ctx context.Context, params func(SessionContext_activity_Params) error, opts ...capnp.CallOption) SessionContext_activity_Results_Promise {
	if c.Client == nil {
		return SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_activity_Params{Struct: s}) }
	}
	return SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SessionContext_Server interface {
	GetSharedPermissions(SessionContext_getSharedPermissions) error

	TieToUser(SessionContext_tieToUser) error

	Offer(SessionContext_offer) error

	Request(SessionContext_request) error

	FulfillRequest(SessionContext_fulfillRequest) error

	Close(SessionContext_close) error

	OpenView(SessionContext_openView) error

	ClaimRequest(SessionContext_claimRequest) error

	Activity(SessionContext_activity) error
}

func SessionContext_ServerToClient(s SessionContext_Server) SessionContext {
	c, _ := s.(server.Closer)
	return SessionContext{Client: server.New(SessionContext_Methods(nil, s), c)}
}

func SessionContext_Methods(methods []server.Method, s SessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 9)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_getSharedPermissions{c, opts, SessionContext_getSharedPermissions_Params{Struct: p}, SessionContext_getSharedPermissions_Results{Struct: r}}
			return s.GetSharedPermissions(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_tieToUser{c, opts, SessionContext_tieToUser_Params{Struct: p}, SessionContext_tieToUser_Results{Struct: r}}
			return s.TieToUser(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_offer{c, opts, SessionContext_offer_Params{Struct: p}, SessionContext_offer_Results{Struct: r}}
			return s.Offer(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_request{c, opts, SessionContext_request_Params{Struct: p}, SessionContext_request_Results{Struct: r}}
			return s.Request(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_fulfillRequest{c, opts, SessionContext_fulfillRequest_Params{Struct: p}, SessionContext_fulfillRequest_Results{Struct: r}}
			return s.FulfillRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_close{c, opts, SessionContext_close_Params{Struct: p}, SessionContext_close_Results{Struct: r}}
			return s.Close(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_openView{c, opts, SessionContext_openView_Params{Struct: p}, SessionContext_openView_Results{Struct: r}}
			return s.OpenView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_claimRequest{c, opts, SessionContext_claimRequest_Params{Struct: p}, SessionContext_claimRequest_Results{Struct: r}}
			return s.ClaimRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_activity{c, opts, SessionContext_activity_Params{Struct: p}, SessionContext_activity_Results{Struct: r}}
			return s.Activity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// SessionContext_getSharedPermissions holds the arguments for a server call to SessionContext.getSharedPermissions.
type SessionContext_getSharedPermissions struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_getSharedPermissions_Params
	Results SessionContext_getSharedPermissions_Results
}

// SessionContext_tieToUser holds the arguments for a server call to SessionContext.tieToUser.
type SessionContext_tieToUser struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_tieToUser_Params
	Results SessionContext_tieToUser_Results
}

// SessionContext_offer holds the arguments for a server call to SessionContext.offer.
type SessionContext_offer struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_offer_Params
	Results SessionContext_offer_Results
}

// SessionContext_request holds the arguments for a server call to SessionContext.request.
type SessionContext_request struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_request_Params
	Results SessionContext_request_Results
}

// SessionContext_fulfillRequest holds the arguments for a server call to SessionContext.fulfillRequest.
type SessionContext_fulfillRequest struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_fulfillRequest_Params
	Results SessionContext_fulfillRequest_Results
}

// SessionContext_close holds the arguments for a server call to SessionContext.close.
type SessionContext_close struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_close_Params
	Results SessionContext_close_Results
}

// SessionContext_openView holds the arguments for a server call to SessionContext.openView.
type SessionContext_openView struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_openView_Params
	Results SessionContext_openView_Results
}

// SessionContext_claimRequest holds the arguments for a server call to SessionContext.claimRequest.
type SessionContext_claimRequest struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_claimRequest_Params
	Results SessionContext_claimRequest_Results
}

// SessionContext_activity holds the arguments for a server call to SessionContext.activity.
type SessionContext_activity struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_activity_Params
	Results SessionContext_activity_Results
}

type SessionContext_getSharedPermissions_Params struct{ capnp.Struct }

func NewSessionContext_getSharedPermissions_Params(s *capnp.Segment) (SessionContext_getSharedPermissions_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_getSharedPermissions_Params{st}, err
}

func NewRootSessionContext_getSharedPermissions_Params(s *capnp.Segment) (SessionContext_getSharedPermissions_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_getSharedPermissions_Params{st}, err
}

func ReadRootSessionContext_getSharedPermissions_Params(msg *capnp.Message) (SessionContext_getSharedPermissions_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_getSharedPermissions_Params{root.Struct()}, err
}

func (s SessionContext_getSharedPermissions_Params) String() string {
	str, _ := text.Marshal(0xe96859cf77da6e6b, s.Struct)
	return str
}

// SessionContext_getSharedPermissions_Params_List is a list of SessionContext_getSharedPermissions_Params.
type SessionContext_getSharedPermissions_Params_List struct{ capnp.List }

// NewSessionContext_getSharedPermissions_Params creates a new list of SessionContext_getSharedPermissions_Params.
func NewSessionContext_getSharedPermissions_Params_List(s *capnp.Segment, sz int32) (SessionContext_getSharedPermissions_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_getSharedPermissions_Params_List{l}, err
}

func (s SessionContext_getSharedPermissions_Params_List) At(i int) SessionContext_getSharedPermissions_Params {
	return SessionContext_getSharedPermissions_Params{s.List.Struct(i)}
}

func (s SessionContext_getSharedPermissions_Params_List) Set(i int, v SessionContext_getSharedPermissions_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_getSharedPermissions_Params_Promise is a wrapper for a SessionContext_getSharedPermissions_Params promised by a client call.
type SessionContext_getSharedPermissions_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_getSharedPermissions_Params_Promise) Struct() (SessionContext_getSharedPermissions_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_getSharedPermissions_Params{s}, err
}

type SessionContext_getSharedPermissions_Results struct{ capnp.Struct }

func NewSessionContext_getSharedPermissions_Results(s *capnp.Segment) (SessionContext_getSharedPermissions_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_getSharedPermissions_Results{st}, err
}

func NewRootSessionContext_getSharedPermissions_Results(s *capnp.Segment) (SessionContext_getSharedPermissions_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_getSharedPermissions_Results{st}, err
}

func ReadRootSessionContext_getSharedPermissions_Results(msg *capnp.Message) (SessionContext_getSharedPermissions_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_getSharedPermissions_Results{root.Struct()}, err
}

func (s SessionContext_getSharedPermissions_Results) String() string {
	str, _ := text.Marshal(0xb70bd877cecb7b88, s.Struct)
	return str
}

func (s SessionContext_getSharedPermissions_Results) Var() util.Assignable_Getter {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable_Getter{Client: p.Interface().Client()}
}

func (s SessionContext_getSharedPermissions_Results) HasVar() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_getSharedPermissions_Results) SetVar(v util.Assignable_Getter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SessionContext_getSharedPermissions_Results_List is a list of SessionContext_getSharedPermissions_Results.
type SessionContext_getSharedPermissions_Results_List struct{ capnp.List }

// NewSessionContext_getSharedPermissions_Results creates a new list of SessionContext_getSharedPermissions_Results.
func NewSessionContext_getSharedPermissions_Results_List(s *capnp.Segment, sz int32) (SessionContext_getSharedPermissions_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SessionContext_getSharedPermissions_Results_List{l}, err
}

func (s SessionContext_getSharedPermissions_Results_List) At(i int) SessionContext_getSharedPermissions_Results {
	return SessionContext_getSharedPermissions_Results{s.List.Struct(i)}
}

func (s SessionContext_getSharedPermissions_Results_List) Set(i int, v SessionContext_getSharedPermissions_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_getSharedPermissions_Results_Promise is a wrapper for a SessionContext_getSharedPermissions_Results promised by a client call.
type SessionContext_getSharedPermissions_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_getSharedPermissions_Results_Promise) Struct() (SessionContext_getSharedPermissions_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_getSharedPermissions_Results{s}, err
}

func (p SessionContext_getSharedPermissions_Results_Promise) Var() util.Assignable_Getter {
	return util.Assignable_Getter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SessionContext_tieToUser_Params struct{ capnp.Struct }

func NewSessionContext_tieToUser_Params(s *capnp.Segment) (SessionContext_tieToUser_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SessionContext_tieToUser_Params{st}, err
}

func NewRootSessionContext_tieToUser_Params(s *capnp.Segment) (SessionContext_tieToUser_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SessionContext_tieToUser_Params{st}, err
}

func ReadRootSessionContext_tieToUser_Params(msg *capnp.Message) (SessionContext_tieToUser_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_tieToUser_Params{root.Struct()}, err
}

func (s SessionContext_tieToUser_Params) String() string {
	str, _ := text.Marshal(0xc41e71e8d893086c, s.Struct)
	return str
}

func (s SessionContext_tieToUser_Params) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_tieToUser_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_tieToUser_Params) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_tieToUser_Params) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_tieToUser_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SessionContext_tieToUser_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SessionContext_tieToUser_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s SessionContext_tieToUser_Params) DisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(2)
	return powerbox.PowerboxDisplayInfo{Struct: p.Struct()}, err
}

func (s SessionContext_tieToUser_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) SetDisplayInfo(v powerbox.PowerboxDisplayInfo) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated powerbox.PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_tieToUser_Params) NewDisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	ss, err := powerbox.NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_tieToUser_Params_List is a list of SessionContext_tieToUser_Params.
type SessionContext_tieToUser_Params_List struct{ capnp.List }

// NewSessionContext_tieToUser_Params creates a new list of SessionContext_tieToUser_Params.
func NewSessionContext_tieToUser_Params_List(s *capnp.Segment, sz int32) (SessionContext_tieToUser_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return SessionContext_tieToUser_Params_List{l}, err
}

func (s SessionContext_tieToUser_Params_List) At(i int) SessionContext_tieToUser_Params {
	return SessionContext_tieToUser_Params{s.List.Struct(i)}
}

func (s SessionContext_tieToUser_Params_List) Set(i int, v SessionContext_tieToUser_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_tieToUser_Params_Promise is a wrapper for a SessionContext_tieToUser_Params promised by a client call.
type SessionContext_tieToUser_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_tieToUser_Params_Promise) Struct() (SessionContext_tieToUser_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_tieToUser_Params{s}, err
}

func (p SessionContext_tieToUser_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_tieToUser_Params_Promise) DisplayInfo() powerbox.PowerboxDisplayInfo_Promise {
	return powerbox.PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SessionContext_tieToUser_Results struct{ capnp.Struct }

func NewSessionContext_tieToUser_Results(s *capnp.Segment) (SessionContext_tieToUser_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_tieToUser_Results{st}, err
}

func NewRootSessionContext_tieToUser_Results(s *capnp.Segment) (SessionContext_tieToUser_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_tieToUser_Results{st}, err
}

func ReadRootSessionContext_tieToUser_Results(msg *capnp.Message) (SessionContext_tieToUser_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_tieToUser_Results{root.Struct()}, err
}

func (s SessionContext_tieToUser_Results) String() string {
	str, _ := text.Marshal(0xf6f911c4804ba7e5, s.Struct)
	return str
}

func (s SessionContext_tieToUser_Results) TiedCap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_tieToUser_Results) HasTiedCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Results) TiedCapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_tieToUser_Results) SetTiedCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_tieToUser_Results) SetTiedCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SessionContext_tieToUser_Results_List is a list of SessionContext_tieToUser_Results.
type SessionContext_tieToUser_Results_List struct{ capnp.List }

// NewSessionContext_tieToUser_Results creates a new list of SessionContext_tieToUser_Results.
func NewSessionContext_tieToUser_Results_List(s *capnp.Segment, sz int32) (SessionContext_tieToUser_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SessionContext_tieToUser_Results_List{l}, err
}

func (s SessionContext_tieToUser_Results_List) At(i int) SessionContext_tieToUser_Results {
	return SessionContext_tieToUser_Results{s.List.Struct(i)}
}

func (s SessionContext_tieToUser_Results_List) Set(i int, v SessionContext_tieToUser_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_tieToUser_Results_Promise is a wrapper for a SessionContext_tieToUser_Results promised by a client call.
type SessionContext_tieToUser_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_tieToUser_Results_Promise) Struct() (SessionContext_tieToUser_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_tieToUser_Results{s}, err
}

func (p SessionContext_tieToUser_Results_Promise) TiedCap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SessionContext_offer_Params struct{ capnp.Struct }

func NewSessionContext_offer_Params(s *capnp.Segment) (SessionContext_offer_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return SessionContext_offer_Params{st}, err
}

func NewRootSessionContext_offer_Params(s *capnp.Segment) (SessionContext_offer_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return SessionContext_offer_Params{st}, err
}

func ReadRootSessionContext_offer_Params(msg *capnp.Message) (SessionContext_offer_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_offer_Params{root.Struct()}, err
}

func (s SessionContext_offer_Params) String() string {
	str, _ := text.Marshal(0xfb3d38da0c9eaee6, s.Struct)
	return str
}

func (s SessionContext_offer_Params) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_offer_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_offer_Params) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_offer_Params) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_offer_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SessionContext_offer_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SessionContext_offer_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s SessionContext_offer_Params) Descriptor() (powerbox.PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(2)
	return powerbox.PowerboxDescriptor{Struct: p.Struct()}, err
}

func (s SessionContext_offer_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetDescriptor(v powerbox.PowerboxDescriptor) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated powerbox.PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_offer_Params) NewDescriptor() (powerbox.PowerboxDescriptor, error) {
	ss, err := powerbox.NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionContext_offer_Params) DisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(3)
	return powerbox.PowerboxDisplayInfo{Struct: p.Struct()}, err
}

func (s SessionContext_offer_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetDisplayInfo(v powerbox.PowerboxDisplayInfo) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated powerbox.PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_offer_Params) NewDisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	ss, err := powerbox.NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_offer_Params_List is a list of SessionContext_offer_Params.
type SessionContext_offer_Params_List struct{ capnp.List }

// NewSessionContext_offer_Params creates a new list of SessionContext_offer_Params.
func NewSessionContext_offer_Params_List(s *capnp.Segment, sz int32) (SessionContext_offer_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return SessionContext_offer_Params_List{l}, err
}

func (s SessionContext_offer_Params_List) At(i int) SessionContext_offer_Params {
	return SessionContext_offer_Params{s.List.Struct(i)}
}

func (s SessionContext_offer_Params_List) Set(i int, v SessionContext_offer_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_offer_Params_Promise is a wrapper for a SessionContext_offer_Params promised by a client call.
type SessionContext_offer_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_offer_Params_Promise) Struct() (SessionContext_offer_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_offer_Params{s}, err
}

func (p SessionContext_offer_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_offer_Params_Promise) Descriptor() powerbox.PowerboxDescriptor_Promise {
	return powerbox.PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p SessionContext_offer_Params_Promise) DisplayInfo() powerbox.PowerboxDisplayInfo_Promise {
	return powerbox.PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SessionContext_offer_Results struct{ capnp.Struct }

func NewSessionContext_offer_Results(s *capnp.Segment) (SessionContext_offer_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_offer_Results{st}, err
}

func NewRootSessionContext_offer_Results(s *capnp.Segment) (SessionContext_offer_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_offer_Results{st}, err
}

func ReadRootSessionContext_offer_Results(msg *capnp.Message) (SessionContext_offer_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_offer_Results{root.Struct()}, err
}

func (s SessionContext_offer_Results) String() string {
	str, _ := text.Marshal(0xfe7135f15d39bd5b, s.Struct)
	return str
}

// SessionContext_offer_Results_List is a list of SessionContext_offer_Results.
type SessionContext_offer_Results_List struct{ capnp.List }

// NewSessionContext_offer_Results creates a new list of SessionContext_offer_Results.
func NewSessionContext_offer_Results_List(s *capnp.Segment, sz int32) (SessionContext_offer_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_offer_Results_List{l}, err
}

func (s SessionContext_offer_Results_List) At(i int) SessionContext_offer_Results {
	return SessionContext_offer_Results{s.List.Struct(i)}
}

func (s SessionContext_offer_Results_List) Set(i int, v SessionContext_offer_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_offer_Results_Promise is a wrapper for a SessionContext_offer_Results promised by a client call.
type SessionContext_offer_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_offer_Results_Promise) Struct() (SessionContext_offer_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_offer_Results{s}, err
}

type SessionContext_request_Params struct{ capnp.Struct }

func NewSessionContext_request_Params(s *capnp.Segment) (SessionContext_request_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_request_Params{st}, err
}

func NewRootSessionContext_request_Params(s *capnp.Segment) (SessionContext_request_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_request_Params{st}, err
}

func ReadRootSessionContext_request_Params(msg *capnp.Message) (SessionContext_request_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_request_Params{root.Struct()}, err
}

func (s SessionContext_request_Params) String() string {
	str, _ := text.Marshal(0xf63b8546288ee8e1, s.Struct)
	return str
}

func (s SessionContext_request_Params) Query() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(0)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s SessionContext_request_Params) HasQuery() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Params) SetQuery(v powerbox.PowerboxDescriptor_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewQuery sets the query field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s SessionContext_request_Params) NewQuery(n int32) (powerbox.PowerboxDescriptor_List, error) {
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s SessionContext_request_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SessionContext_request_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SessionContext_request_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// SessionContext_request_Params_List is a list of SessionContext_request_Params.
type SessionContext_request_Params_List struct{ capnp.List }

// NewSessionContext_request_Params creates a new list of SessionContext_request_Params.
func NewSessionContext_request_Params_List(s *capnp.Segment, sz int32) (SessionContext_request_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SessionContext_request_Params_List{l}, err
}

func (s SessionContext_request_Params_List) At(i int) SessionContext_request_Params {
	return SessionContext_request_Params{s.List.Struct(i)}
}

func (s SessionContext_request_Params_List) Set(i int, v SessionContext_request_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_request_Params_Promise is a wrapper for a SessionContext_request_Params promised by a client call.
type SessionContext_request_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_request_Params_Promise) Struct() (SessionContext_request_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_request_Params{s}, err
}

type SessionContext_request_Results struct{ capnp.Struct }

func NewSessionContext_request_Results(s *capnp.Segment) (SessionContext_request_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_request_Results{st}, err
}

func NewRootSessionContext_request_Results(s *capnp.Segment) (SessionContext_request_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_request_Results{st}, err
}

func ReadRootSessionContext_request_Results(msg *capnp.Message) (SessionContext_request_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_request_Results{root.Struct()}, err
}

func (s SessionContext_request_Results) String() string {
	str, _ := text.Marshal(0xd42684f756e09afd, s.Struct)
	return str
}

func (s SessionContext_request_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_request_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_request_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_request_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_request_Results) Descriptor() (powerbox.PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(1)
	return powerbox.PowerboxDescriptor{Struct: p.Struct()}, err
}

func (s SessionContext_request_Results) HasDescriptor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Results) SetDescriptor(v powerbox.PowerboxDescriptor) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated powerbox.PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_request_Results) NewDescriptor() (powerbox.PowerboxDescriptor, error) {
	ss, err := powerbox.NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_request_Results_List is a list of SessionContext_request_Results.
type SessionContext_request_Results_List struct{ capnp.List }

// NewSessionContext_request_Results creates a new list of SessionContext_request_Results.
func NewSessionContext_request_Results_List(s *capnp.Segment, sz int32) (SessionContext_request_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SessionContext_request_Results_List{l}, err
}

func (s SessionContext_request_Results_List) At(i int) SessionContext_request_Results {
	return SessionContext_request_Results{s.List.Struct(i)}
}

func (s SessionContext_request_Results_List) Set(i int, v SessionContext_request_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_request_Results_Promise is a wrapper for a SessionContext_request_Results promised by a client call.
type SessionContext_request_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_request_Results_Promise) Struct() (SessionContext_request_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_request_Results{s}, err
}

func (p SessionContext_request_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_request_Results_Promise) Descriptor() powerbox.PowerboxDescriptor_Promise {
	return powerbox.PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SessionContext_fulfillRequest_Params struct{ capnp.Struct }

func NewSessionContext_fulfillRequest_Params(s *capnp.Segment) (SessionContext_fulfillRequest_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return SessionContext_fulfillRequest_Params{st}, err
}

func NewRootSessionContext_fulfillRequest_Params(s *capnp.Segment) (SessionContext_fulfillRequest_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return SessionContext_fulfillRequest_Params{st}, err
}

func ReadRootSessionContext_fulfillRequest_Params(msg *capnp.Message) (SessionContext_fulfillRequest_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_fulfillRequest_Params{root.Struct()}, err
}

func (s SessionContext_fulfillRequest_Params) String() string {
	str, _ := text.Marshal(0x9f6c36ef490dfd92, s.Struct)
	return str
}

func (s SessionContext_fulfillRequest_Params) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_fulfillRequest_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_fulfillRequest_Params) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_fulfillRequest_Params) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_fulfillRequest_Params) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_fulfillRequest_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SessionContext_fulfillRequest_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_fulfillRequest_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SessionContext_fulfillRequest_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s SessionContext_fulfillRequest_Params) Descriptor() (powerbox.PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(2)
	return powerbox.PowerboxDescriptor{Struct: p.Struct()}, err
}

func (s SessionContext_fulfillRequest_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_fulfillRequest_Params) SetDescriptor(v powerbox.PowerboxDescriptor) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated powerbox.PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_fulfillRequest_Params) NewDescriptor() (powerbox.PowerboxDescriptor, error) {
	ss, err := powerbox.NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionContext_fulfillRequest_Params) DisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(3)
	return powerbox.PowerboxDisplayInfo{Struct: p.Struct()}, err
}

func (s SessionContext_fulfillRequest_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s SessionContext_fulfillRequest_Params) SetDisplayInfo(v powerbox.PowerboxDisplayInfo) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated powerbox.PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_fulfillRequest_Params) NewDisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	ss, err := powerbox.NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_fulfillRequest_Params_List is a list of SessionContext_fulfillRequest_Params.
type SessionContext_fulfillRequest_Params_List struct{ capnp.List }

// NewSessionContext_fulfillRequest_Params creates a new list of SessionContext_fulfillRequest_Params.
func NewSessionContext_fulfillRequest_Params_List(s *capnp.Segment, sz int32) (SessionContext_fulfillRequest_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return SessionContext_fulfillRequest_Params_List{l}, err
}

func (s SessionContext_fulfillRequest_Params_List) At(i int) SessionContext_fulfillRequest_Params {
	return SessionContext_fulfillRequest_Params{s.List.Struct(i)}
}

func (s SessionContext_fulfillRequest_Params_List) Set(i int, v SessionContext_fulfillRequest_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_fulfillRequest_Params_Promise is a wrapper for a SessionContext_fulfillRequest_Params promised by a client call.
type SessionContext_fulfillRequest_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_fulfillRequest_Params_Promise) Struct() (SessionContext_fulfillRequest_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_fulfillRequest_Params{s}, err
}

func (p SessionContext_fulfillRequest_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_fulfillRequest_Params_Promise) Descriptor() powerbox.PowerboxDescriptor_Promise {
	return powerbox.PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p SessionContext_fulfillRequest_Params_Promise) DisplayInfo() powerbox.PowerboxDisplayInfo_Promise {
	return powerbox.PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SessionContext_fulfillRequest_Results struct{ capnp.Struct }

func NewSessionContext_fulfillRequest_Results(s *capnp.Segment) (SessionContext_fulfillRequest_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_fulfillRequest_Results{st}, err
}

func NewRootSessionContext_fulfillRequest_Results(s *capnp.Segment) (SessionContext_fulfillRequest_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_fulfillRequest_Results{st}, err
}

func ReadRootSessionContext_fulfillRequest_Results(msg *capnp.Message) (SessionContext_fulfillRequest_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_fulfillRequest_Results{root.Struct()}, err
}

func (s SessionContext_fulfillRequest_Results) String() string {
	str, _ := text.Marshal(0xb4ecd69ac97e2de8, s.Struct)
	return str
}

// SessionContext_fulfillRequest_Results_List is a list of SessionContext_fulfillRequest_Results.
type SessionContext_fulfillRequest_Results_List struct{ capnp.List }

// NewSessionContext_fulfillRequest_Results creates a new list of SessionContext_fulfillRequest_Results.
func NewSessionContext_fulfillRequest_Results_List(s *capnp.Segment, sz int32) (SessionContext_fulfillRequest_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_fulfillRequest_Results_List{l}, err
}

func (s SessionContext_fulfillRequest_Results_List) At(i int) SessionContext_fulfillRequest_Results {
	return SessionContext_fulfillRequest_Results{s.List.Struct(i)}
}

func (s SessionContext_fulfillRequest_Results_List) Set(i int, v SessionContext_fulfillRequest_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_fulfillRequest_Results_Promise is a wrapper for a SessionContext_fulfillRequest_Results promised by a client call.
type SessionContext_fulfillRequest_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_fulfillRequest_Results_Promise) Struct() (SessionContext_fulfillRequest_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_fulfillRequest_Results{s}, err
}

type SessionContext_close_Params struct{ capnp.Struct }

func NewSessionContext_close_Params(s *capnp.Segment) (SessionContext_close_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_close_Params{st}, err
}

func NewRootSessionContext_close_Params(s *capnp.Segment) (SessionContext_close_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_close_Params{st}, err
}

func ReadRootSessionContext_close_Params(msg *capnp.Message) (SessionContext_close_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_close_Params{root.Struct()}, err
}

func (s SessionContext_close_Params) String() string {
	str, _ := text.Marshal(0xf12c60ebc67984d4, s.Struct)
	return str
}

// SessionContext_close_Params_List is a list of SessionContext_close_Params.
type SessionContext_close_Params_List struct{ capnp.List }

// NewSessionContext_close_Params creates a new list of SessionContext_close_Params.
func NewSessionContext_close_Params_List(s *capnp.Segment, sz int32) (SessionContext_close_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_close_Params_List{l}, err
}

func (s SessionContext_close_Params_List) At(i int) SessionContext_close_Params {
	return SessionContext_close_Params{s.List.Struct(i)}
}

func (s SessionContext_close_Params_List) Set(i int, v SessionContext_close_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_close_Params_Promise is a wrapper for a SessionContext_close_Params promised by a client call.
type SessionContext_close_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_close_Params_Promise) Struct() (SessionContext_close_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_close_Params{s}, err
}

type SessionContext_close_Results struct{ capnp.Struct }

func NewSessionContext_close_Results(s *capnp.Segment) (SessionContext_close_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_close_Results{st}, err
}

func NewRootSessionContext_close_Results(s *capnp.Segment) (SessionContext_close_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_close_Results{st}, err
}

func ReadRootSessionContext_close_Results(msg *capnp.Message) (SessionContext_close_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_close_Results{root.Struct()}, err
}

func (s SessionContext_close_Results) String() string {
	str, _ := text.Marshal(0x9d4102fadb4f069c, s.Struct)
	return str
}

// SessionContext_close_Results_List is a list of SessionContext_close_Results.
type SessionContext_close_Results_List struct{ capnp.List }

// NewSessionContext_close_Results creates a new list of SessionContext_close_Results.
func NewSessionContext_close_Results_List(s *capnp.Segment, sz int32) (SessionContext_close_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_close_Results_List{l}, err
}

func (s SessionContext_close_Results_List) At(i int) SessionContext_close_Results {
	return SessionContext_close_Results{s.List.Struct(i)}
}

func (s SessionContext_close_Results_List) Set(i int, v SessionContext_close_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_close_Results_Promise is a wrapper for a SessionContext_close_Results promised by a client call.
type SessionContext_close_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_close_Results_Promise) Struct() (SessionContext_close_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_close_Results{s}, err
}

type SessionContext_openView_Params struct{ capnp.Struct }

func NewSessionContext_openView_Params(s *capnp.Segment) (SessionContext_openView_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return SessionContext_openView_Params{st}, err
}

func NewRootSessionContext_openView_Params(s *capnp.Segment) (SessionContext_openView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return SessionContext_openView_Params{st}, err
}

func ReadRootSessionContext_openView_Params(msg *capnp.Message) (SessionContext_openView_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_openView_Params{root.Struct()}, err
}

func (s SessionContext_openView_Params) String() string {
	str, _ := text.Marshal(0xf37f5e08534c68aa, s.Struct)
	return str
}

func (s SessionContext_openView_Params) View() UiView {
	p, _ := s.Struct.Ptr(0)
	return UiView{Client: p.Interface().Client()}
}

func (s SessionContext_openView_Params) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_openView_Params) SetView(v UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s SessionContext_openView_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s SessionContext_openView_Params) HasPath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_openView_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s SessionContext_openView_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s SessionContext_openView_Params) NewTab() bool {
	return s.Struct.Bit(0)
}

func (s SessionContext_openView_Params) SetNewTab(v bool) {
	s.Struct.SetBit(0, v)
}

// SessionContext_openView_Params_List is a list of SessionContext_openView_Params.
type SessionContext_openView_Params_List struct{ capnp.List }

// NewSessionContext_openView_Params creates a new list of SessionContext_openView_Params.
func NewSessionContext_openView_Params_List(s *capnp.Segment, sz int32) (SessionContext_openView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return SessionContext_openView_Params_List{l}, err
}

func (s SessionContext_openView_Params_List) At(i int) SessionContext_openView_Params {
	return SessionContext_openView_Params{s.List.Struct(i)}
}

func (s SessionContext_openView_Params_List) Set(i int, v SessionContext_openView_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_openView_Params_Promise is a wrapper for a SessionContext_openView_Params promised by a client call.
type SessionContext_openView_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_openView_Params_Promise) Struct() (SessionContext_openView_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_openView_Params{s}, err
}

func (p SessionContext_openView_Params_Promise) View() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SessionContext_openView_Results struct{ capnp.Struct }

func NewSessionContext_openView_Results(s *capnp.Segment) (SessionContext_openView_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_openView_Results{st}, err
}

func NewRootSessionContext_openView_Results(s *capnp.Segment) (SessionContext_openView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_openView_Results{st}, err
}

func ReadRootSessionContext_openView_Results(msg *capnp.Message) (SessionContext_openView_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_openView_Results{root.Struct()}, err
}

func (s SessionContext_openView_Results) String() string {
	str, _ := text.Marshal(0xf9d6c8c6d207c123, s.Struct)
	return str
}

// SessionContext_openView_Results_List is a list of SessionContext_openView_Results.
type SessionContext_openView_Results_List struct{ capnp.List }

// NewSessionContext_openView_Results creates a new list of SessionContext_openView_Results.
func NewSessionContext_openView_Results_List(s *capnp.Segment, sz int32) (SessionContext_openView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_openView_Results_List{l}, err
}

func (s SessionContext_openView_Results_List) At(i int) SessionContext_openView_Results {
	return SessionContext_openView_Results{s.List.Struct(i)}
}

func (s SessionContext_openView_Results_List) Set(i int, v SessionContext_openView_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_openView_Results_Promise is a wrapper for a SessionContext_openView_Results promised by a client call.
type SessionContext_openView_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_openView_Results_Promise) Struct() (SessionContext_openView_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_openView_Results{s}, err
}

type SessionContext_claimRequest_Params struct{ capnp.Struct }

func NewSessionContext_claimRequest_Params(s *capnp.Segment) (SessionContext_claimRequest_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_claimRequest_Params{st}, err
}

func NewRootSessionContext_claimRequest_Params(s *capnp.Segment) (SessionContext_claimRequest_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return SessionContext_claimRequest_Params{st}, err
}

func ReadRootSessionContext_claimRequest_Params(msg *capnp.Message) (SessionContext_claimRequest_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_claimRequest_Params{root.Struct()}, err
}

func (s SessionContext_claimRequest_Params) String() string {
	str, _ := text.Marshal(0xda13a4f2919ce2cf, s.Struct)
	return str
}

func (s SessionContext_claimRequest_Params) RequestToken() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SessionContext_claimRequest_Params) HasRequestToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_claimRequest_Params) RequestTokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SessionContext_claimRequest_Params) SetRequestToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s SessionContext_claimRequest_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s SessionContext_claimRequest_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_claimRequest_Params) SetRequiredPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s SessionContext_claimRequest_Params) NewRequiredPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// SessionContext_claimRequest_Params_List is a list of SessionContext_claimRequest_Params.
type SessionContext_claimRequest_Params_List struct{ capnp.List }

// NewSessionContext_claimRequest_Params creates a new list of SessionContext_claimRequest_Params.
func NewSessionContext_claimRequest_Params_List(s *capnp.Segment, sz int32) (SessionContext_claimRequest_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return SessionContext_claimRequest_Params_List{l}, err
}

func (s SessionContext_claimRequest_Params_List) At(i int) SessionContext_claimRequest_Params {
	return SessionContext_claimRequest_Params{s.List.Struct(i)}
}

func (s SessionContext_claimRequest_Params_List) Set(i int, v SessionContext_claimRequest_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_claimRequest_Params_Promise is a wrapper for a SessionContext_claimRequest_Params promised by a client call.
type SessionContext_claimRequest_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_claimRequest_Params_Promise) Struct() (SessionContext_claimRequest_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_claimRequest_Params{s}, err
}

type SessionContext_claimRequest_Results struct{ capnp.Struct }

func NewSessionContext_claimRequest_Results(s *capnp.Segment) (SessionContext_claimRequest_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_claimRequest_Results{st}, err
}

func NewRootSessionContext_claimRequest_Results(s *capnp.Segment) (SessionContext_claimRequest_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_claimRequest_Results{st}, err
}

func ReadRootSessionContext_claimRequest_Results(msg *capnp.Message) (SessionContext_claimRequest_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_claimRequest_Results{root.Struct()}, err
}

func (s SessionContext_claimRequest_Results) String() string {
	str, _ := text.Marshal(0xefea656d4b56b756, s.Struct)
	return str
}

func (s SessionContext_claimRequest_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s SessionContext_claimRequest_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_claimRequest_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s SessionContext_claimRequest_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_claimRequest_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// SessionContext_claimRequest_Results_List is a list of SessionContext_claimRequest_Results.
type SessionContext_claimRequest_Results_List struct{ capnp.List }

// NewSessionContext_claimRequest_Results creates a new list of SessionContext_claimRequest_Results.
func NewSessionContext_claimRequest_Results_List(s *capnp.Segment, sz int32) (SessionContext_claimRequest_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SessionContext_claimRequest_Results_List{l}, err
}

func (s SessionContext_claimRequest_Results_List) At(i int) SessionContext_claimRequest_Results {
	return SessionContext_claimRequest_Results{s.List.Struct(i)}
}

func (s SessionContext_claimRequest_Results_List) Set(i int, v SessionContext_claimRequest_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_claimRequest_Results_Promise is a wrapper for a SessionContext_claimRequest_Results promised by a client call.
type SessionContext_claimRequest_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_claimRequest_Results_Promise) Struct() (SessionContext_claimRequest_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_claimRequest_Results{s}, err
}

func (p SessionContext_claimRequest_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SessionContext_activity_Params struct{ capnp.Struct }

func NewSessionContext_activity_Params(s *capnp.Segment) (SessionContext_activity_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_activity_Params{st}, err
}

func NewRootSessionContext_activity_Params(s *capnp.Segment) (SessionContext_activity_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SessionContext_activity_Params{st}, err
}

func ReadRootSessionContext_activity_Params(msg *capnp.Message) (SessionContext_activity_Params, error) {
	root, err := msg.RootPtr()
	return SessionContext_activity_Params{root.Struct()}, err
}

func (s SessionContext_activity_Params) String() string {
	str, _ := text.Marshal(0x85e320f14a5d23e0, s.Struct)
	return str
}

func (s SessionContext_activity_Params) Event() (activity.ActivityEvent, error) {
	p, err := s.Struct.Ptr(0)
	return activity.ActivityEvent{Struct: p.Struct()}, err
}

func (s SessionContext_activity_Params) HasEvent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_activity_Params) SetEvent(v activity.ActivityEvent) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEvent sets the event field to a newly
// allocated activity.ActivityEvent struct, preferring placement in s's segment.
func (s SessionContext_activity_Params) NewEvent() (activity.ActivityEvent, error) {
	ss, err := activity.NewActivityEvent(s.Struct.Segment())
	if err != nil {
		return activity.ActivityEvent{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_activity_Params_List is a list of SessionContext_activity_Params.
type SessionContext_activity_Params_List struct{ capnp.List }

// NewSessionContext_activity_Params creates a new list of SessionContext_activity_Params.
func NewSessionContext_activity_Params_List(s *capnp.Segment, sz int32) (SessionContext_activity_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SessionContext_activity_Params_List{l}, err
}

func (s SessionContext_activity_Params_List) At(i int) SessionContext_activity_Params {
	return SessionContext_activity_Params{s.List.Struct(i)}
}

func (s SessionContext_activity_Params_List) Set(i int, v SessionContext_activity_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_activity_Params_Promise is a wrapper for a SessionContext_activity_Params promised by a client call.
type SessionContext_activity_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_activity_Params_Promise) Struct() (SessionContext_activity_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_activity_Params{s}, err
}

func (p SessionContext_activity_Params_Promise) Event() activity.ActivityEvent_Promise {
	return activity.ActivityEvent_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type SessionContext_activity_Results struct{ capnp.Struct }

func NewSessionContext_activity_Results(s *capnp.Segment) (SessionContext_activity_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_activity_Results{st}, err
}

func NewRootSessionContext_activity_Results(s *capnp.Segment) (SessionContext_activity_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SessionContext_activity_Results{st}, err
}

func ReadRootSessionContext_activity_Results(msg *capnp.Message) (SessionContext_activity_Results, error) {
	root, err := msg.RootPtr()
	return SessionContext_activity_Results{root.Struct()}, err
}

func (s SessionContext_activity_Results) String() string {
	str, _ := text.Marshal(0xa93eadc9671ea08b, s.Struct)
	return str
}

// SessionContext_activity_Results_List is a list of SessionContext_activity_Results.
type SessionContext_activity_Results_List struct{ capnp.List }

// NewSessionContext_activity_Results creates a new list of SessionContext_activity_Results.
func NewSessionContext_activity_Results_List(s *capnp.Segment, sz int32) (SessionContext_activity_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SessionContext_activity_Results_List{l}, err
}

func (s SessionContext_activity_Results_List) At(i int) SessionContext_activity_Results {
	return SessionContext_activity_Results{s.List.Struct(i)}
}

func (s SessionContext_activity_Results_List) Set(i int, v SessionContext_activity_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_activity_Results_Promise is a wrapper for a SessionContext_activity_Results promised by a client call.
type SessionContext_activity_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_activity_Results_Promise) Struct() (SessionContext_activity_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_activity_Results{s}, err
}

type PermissionDef struct{ capnp.Struct }

func NewPermissionDef(s *capnp.Segment) (PermissionDef, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return PermissionDef{st}, err
}

func NewRootPermissionDef(s *capnp.Segment) (PermissionDef, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return PermissionDef{st}, err
}

func ReadRootPermissionDef(msg *capnp.Message) (PermissionDef, error) {
	root, err := msg.RootPtr()
	return PermissionDef{root.Struct()}, err
}

func (s PermissionDef) String() string {
	str, _ := text.Marshal(0xf144a5e58889dafb, s.Struct)
	return str
}

func (s PermissionDef) Name() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s PermissionDef) HasName() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PermissionDef) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s PermissionDef) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s PermissionDef) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PermissionDef) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PermissionDef) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PermissionDef) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PermissionDef) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PermissionDef) HasDescription() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PermissionDef) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PermissionDef) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PermissionDef) Obsolete() bool {
	return s.Struct.Bit(0)
}

func (s PermissionDef) SetObsolete(v bool) {
	s.Struct.SetBit(0, v)
}

// PermissionDef_List is a list of PermissionDef.
type PermissionDef_List struct{ capnp.List }

// NewPermissionDef creates a new list of PermissionDef.
func NewPermissionDef_List(s *capnp.Segment, sz int32) (PermissionDef_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return PermissionDef_List{l}, err
}

func (s PermissionDef_List) At(i int) PermissionDef { return PermissionDef{s.List.Struct(i)} }

func (s PermissionDef_List) Set(i int, v PermissionDef) error { return s.List.SetStruct(i, v.Struct) }

// PermissionDef_Promise is a wrapper for a PermissionDef promised by a client call.
type PermissionDef_Promise struct{ *capnp.Pipeline }

func (p PermissionDef_Promise) Struct() (PermissionDef, error) {
	s, err := p.Pipeline.Struct()
	return PermissionDef{s}, err
}

func (p PermissionDef_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PermissionDef_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type RoleDef struct{ capnp.Struct }

func NewRoleDef(s *capnp.Segment) (RoleDef, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return RoleDef{st}, err
}

func NewRootRoleDef(s *capnp.Segment) (RoleDef, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return RoleDef{st}, err
}

func ReadRootRoleDef(msg *capnp.Message) (RoleDef, error) {
	root, err := msg.RootPtr()
	return RoleDef{root.Struct()}, err
}

func (s RoleDef) String() string {
	str, _ := text.Marshal(0xcb3f7064eae4dc5a, s.Struct)
	return str
}

func (s RoleDef) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s RoleDef) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s RoleDef) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetVerbPhrase(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewVerbPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s RoleDef) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.BitList{List: p.List()}, err
}

func (s RoleDef) HasPermissions() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s RoleDef) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s RoleDef) Obsolete() bool {
	return s.Struct.Bit(0)
}

func (s RoleDef) SetObsolete(v bool) {
	s.Struct.SetBit(0, v)
}

func (s RoleDef) Default() bool {
	return s.Struct.Bit(1)
}

func (s RoleDef) SetDefault(v bool) {
	s.Struct.SetBit(1, v)
}

// RoleDef_List is a list of RoleDef.
type RoleDef_List struct{ capnp.List }

// NewRoleDef creates a new list of RoleDef.
func NewRoleDef_List(s *capnp.Segment, sz int32) (RoleDef_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return RoleDef_List{l}, err
}

func (s RoleDef_List) At(i int) RoleDef { return RoleDef{s.List.Struct(i)} }

func (s RoleDef_List) Set(i int, v RoleDef) error { return s.List.SetStruct(i, v.Struct) }

// RoleDef_Promise is a wrapper for a RoleDef promised by a client call.
type RoleDef_Promise struct{ *capnp.Pipeline }

func (p RoleDef_Promise) Struct() (RoleDef, error) {
	s, err := p.Pipeline.Struct()
	return RoleDef{s}, err
}

func (p RoleDef_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p RoleDef_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p RoleDef_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SharingLink struct{ Client capnp.Client }

func (c SharingLink) GetPetname(ctx context.Context, params func(SharingLink_getPetname_Params) error, opts ...capnp.CallOption) SharingLink_getPetname_Results_Promise {
	if c.Client == nil {
		return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SharingLink_getPetname_Params{Struct: s}) }
	}
	return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SharingLink_Server interface {
	GetPetname(SharingLink_getPetname) error
}

func SharingLink_ServerToClient(s SharingLink_Server) SharingLink {
	c, _ := s.(server.Closer)
	return SharingLink{Client: server.New(SharingLink_Methods(nil, s), c)}
}

func SharingLink_Methods(methods []server.Method, s SharingLink_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SharingLink_getPetname{c, opts, SharingLink_getPetname_Params{Struct: p}, SharingLink_getPetname_Results{Struct: r}}
			return s.GetPetname(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SharingLink_getPetname holds the arguments for a server call to SharingLink.getPetname.
type SharingLink_getPetname struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SharingLink_getPetname_Params
	Results SharingLink_getPetname_Results
}

type SharingLink_getPetname_Params struct{ capnp.Struct }

func NewSharingLink_getPetname_Params(s *capnp.Segment) (SharingLink_getPetname_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SharingLink_getPetname_Params{st}, err
}

func NewRootSharingLink_getPetname_Params(s *capnp.Segment) (SharingLink_getPetname_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return SharingLink_getPetname_Params{st}, err
}

func ReadRootSharingLink_getPetname_Params(msg *capnp.Message) (SharingLink_getPetname_Params, error) {
	root, err := msg.RootPtr()
	return SharingLink_getPetname_Params{root.Struct()}, err
}

func (s SharingLink_getPetname_Params) String() string {
	str, _ := text.Marshal(0xf0931856093654c1, s.Struct)
	return str
}

// SharingLink_getPetname_Params_List is a list of SharingLink_getPetname_Params.
type SharingLink_getPetname_Params_List struct{ capnp.List }

// NewSharingLink_getPetname_Params creates a new list of SharingLink_getPetname_Params.
func NewSharingLink_getPetname_Params_List(s *capnp.Segment, sz int32) (SharingLink_getPetname_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return SharingLink_getPetname_Params_List{l}, err
}

func (s SharingLink_getPetname_Params_List) At(i int) SharingLink_getPetname_Params {
	return SharingLink_getPetname_Params{s.List.Struct(i)}
}

func (s SharingLink_getPetname_Params_List) Set(i int, v SharingLink_getPetname_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SharingLink_getPetname_Params_Promise is a wrapper for a SharingLink_getPetname_Params promised by a client call.
type SharingLink_getPetname_Params_Promise struct{ *capnp.Pipeline }

func (p SharingLink_getPetname_Params_Promise) Struct() (SharingLink_getPetname_Params, error) {
	s, err := p.Pipeline.Struct()
	return SharingLink_getPetname_Params{s}, err
}

type SharingLink_getPetname_Results struct{ capnp.Struct }

func NewSharingLink_getPetname_Results(s *capnp.Segment) (SharingLink_getPetname_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SharingLink_getPetname_Results{st}, err
}

func NewRootSharingLink_getPetname_Results(s *capnp.Segment) (SharingLink_getPetname_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SharingLink_getPetname_Results{st}, err
}

func ReadRootSharingLink_getPetname_Results(msg *capnp.Message) (SharingLink_getPetname_Results, error) {
	root, err := msg.RootPtr()
	return SharingLink_getPetname_Results{root.Struct()}, err
}

func (s SharingLink_getPetname_Results) String() string {
	str, _ := text.Marshal(0x9ad927034671cad1, s.Struct)
	return str
}

func (s SharingLink_getPetname_Results) Name() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s SharingLink_getPetname_Results) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SharingLink_getPetname_Results) SetName(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SharingLink_getPetname_Results_List is a list of SharingLink_getPetname_Results.
type SharingLink_getPetname_Results_List struct{ capnp.List }

// NewSharingLink_getPetname_Results creates a new list of SharingLink_getPetname_Results.
func NewSharingLink_getPetname_Results_List(s *capnp.Segment, sz int32) (SharingLink_getPetname_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SharingLink_getPetname_Results_List{l}, err
}

func (s SharingLink_getPetname_Results_List) At(i int) SharingLink_getPetname_Results {
	return SharingLink_getPetname_Results{s.List.Struct(i)}
}

func (s SharingLink_getPetname_Results_List) Set(i int, v SharingLink_getPetname_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SharingLink_getPetname_Results_Promise is a wrapper for a SharingLink_getPetname_Results promised by a client call.
type SharingLink_getPetname_Results_Promise struct{ *capnp.Pipeline }

func (p SharingLink_getPetname_Results_Promise) Struct() (SharingLink_getPetname_Results, error) {
	s, err := p.Pipeline.Struct()
	return SharingLink_getPetname_Results{s}, err
}

func (p SharingLink_getPetname_Results_Promise) Name() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type ViewSharingLink struct{ Client capnp.Client }

func (c ViewSharingLink) GetRoleAssignment(ctx context.Context, params func(ViewSharingLink_getRoleAssignment_Params) error, opts ...capnp.CallOption) ViewSharingLink_getRoleAssignment_Results_Promise {
	if c.Client == nil {
		return ViewSharingLink_getRoleAssignment_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa4f82f764dc3fee8,
			MethodID:      0,
			InterfaceName: "grain.capnp:ViewSharingLink",
			MethodName:    "getRoleAssignment",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ViewSharingLink_getRoleAssignment_Params{Struct: s}) }
	}
	return ViewSharingLink_getRoleAssignment_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ViewSharingLink) GetPetname(ctx context.Context, params func(SharingLink_getPetname_Params) error, opts ...capnp.CallOption) SharingLink_getPetname_Results_Promise {
	if c.Client == nil {
		return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SharingLink_getPetname_Params{Struct: s}) }
	}
	return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ViewSharingLink_Server interface {
	GetRoleAssignment(ViewSharingLink_getRoleAssignment) error

	GetPetname(SharingLink_getPetname) error
}

func ViewSharingLink_ServerToClient(s ViewSharingLink_Server) ViewSharingLink {
	c, _ := s.(server.Closer)
	return ViewSharingLink{Client: server.New(ViewSharingLink_Methods(nil, s), c)}
}

func ViewSharingLink_Methods(methods []server.Method, s ViewSharingLink_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa4f82f764dc3fee8,
			MethodID:      0,
			InterfaceName: "grain.capnp:ViewSharingLink",
			MethodName:    "getRoleAssignment",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ViewSharingLink_getRoleAssignment{c, opts, ViewSharingLink_getRoleAssignment_Params{Struct: p}, ViewSharingLink_getRoleAssignment_Results{Struct: r}}
			return s.GetRoleAssignment(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SharingLink_getPetname{c, opts, SharingLink_getPetname_Params{Struct: p}, SharingLink_getPetname_Results{Struct: r}}
			return s.GetPetname(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// ViewSharingLink_getRoleAssignment holds the arguments for a server call to ViewSharingLink.getRoleAssignment.
type ViewSharingLink_getRoleAssignment struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ViewSharingLink_getRoleAssignment_Params
	Results ViewSharingLink_getRoleAssignment_Results
}

type ViewSharingLink_RoleAssignment struct{ capnp.Struct }
type ViewSharingLink_RoleAssignment_Which uint16

const (
	ViewSharingLink_RoleAssignment_Which_none      ViewSharingLink_RoleAssignment_Which = 0
	ViewSharingLink_RoleAssignment_Which_allAccess ViewSharingLink_RoleAssignment_Which = 1
	ViewSharingLink_RoleAssignment_Which_roleId    ViewSharingLink_RoleAssignment_Which = 2
)

func (w ViewSharingLink_RoleAssignment_Which) String() string {
	const s = "noneallAccessroleId"
	switch w {
	case ViewSharingLink_RoleAssignment_Which_none:
		return s[0:4]
	case ViewSharingLink_RoleAssignment_Which_allAccess:
		return s[4:13]
	case ViewSharingLink_RoleAssignment_Which_roleId:
		return s[13:19]

	}
	return "ViewSharingLink_RoleAssignment_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewViewSharingLink_RoleAssignment(s *capnp.Segment) (ViewSharingLink_RoleAssignment, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return ViewSharingLink_RoleAssignment{st}, err
}

func NewRootViewSharingLink_RoleAssignment(s *capnp.Segment) (ViewSharingLink_RoleAssignment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return ViewSharingLink_RoleAssignment{st}, err
}

func ReadRootViewSharingLink_RoleAssignment(msg *capnp.Message) (ViewSharingLink_RoleAssignment, error) {
	root, err := msg.RootPtr()
	return ViewSharingLink_RoleAssignment{root.Struct()}, err
}

func (s ViewSharingLink_RoleAssignment) String() string {
	str, _ := text.Marshal(0xf020f2be35e8e2b5, s.Struct)
	return str
}

func (s ViewSharingLink_RoleAssignment) Which() ViewSharingLink_RoleAssignment_Which {
	return ViewSharingLink_RoleAssignment_Which(s.Struct.Uint16(0))
}
func (s ViewSharingLink_RoleAssignment) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s ViewSharingLink_RoleAssignment) SetAllAccess() {
	s.Struct.SetUint16(0, 1)

}

func (s ViewSharingLink_RoleAssignment) RoleId() uint16 {
	return s.Struct.Uint16(2)
}

func (s ViewSharingLink_RoleAssignment) SetRoleId(v uint16) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint16(2, v)
}

func (s ViewSharingLink_RoleAssignment) AddPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.BitList{List: p.List()}, err
}

func (s ViewSharingLink_RoleAssignment) HasAddPermissions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_RoleAssignment) SetAddPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewAddPermissions sets the addPermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s ViewSharingLink_RoleAssignment) NewAddPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s ViewSharingLink_RoleAssignment) RemovePermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.BitList{List: p.List()}, err
}

func (s ViewSharingLink_RoleAssignment) HasRemovePermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_RoleAssignment) SetRemovePermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRemovePermissions sets the removePermissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s ViewSharingLink_RoleAssignment) NewRemovePermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// ViewSharingLink_RoleAssignment_List is a list of ViewSharingLink_RoleAssignment.
type ViewSharingLink_RoleAssignment_List struct{ capnp.List }

// NewViewSharingLink_RoleAssignment creates a new list of ViewSharingLink_RoleAssignment.
func NewViewSharingLink_RoleAssignment_List(s *capnp.Segment, sz int32) (ViewSharingLink_RoleAssignment_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return ViewSharingLink_RoleAssignment_List{l}, err
}

func (s ViewSharingLink_RoleAssignment_List) At(i int) ViewSharingLink_RoleAssignment {
	return ViewSharingLink_RoleAssignment{s.List.Struct(i)}
}

func (s ViewSharingLink_RoleAssignment_List) Set(i int, v ViewSharingLink_RoleAssignment) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_RoleAssignment_Promise is a wrapper for a ViewSharingLink_RoleAssignment promised by a client call.
type ViewSharingLink_RoleAssignment_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_RoleAssignment_Promise) Struct() (ViewSharingLink_RoleAssignment, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_RoleAssignment{s}, err
}

type ViewSharingLink_getRoleAssignment_Params struct{ capnp.Struct }

func NewViewSharingLink_getRoleAssignment_Params(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ViewSharingLink_getRoleAssignment_Params{st}, err
}

func NewRootViewSharingLink_getRoleAssignment_Params(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ViewSharingLink_getRoleAssignment_Params{st}, err
}

func ReadRootViewSharingLink_getRoleAssignment_Params(msg *capnp.Message) (ViewSharingLink_getRoleAssignment_Params, error) {
	root, err := msg.RootPtr()
	return ViewSharingLink_getRoleAssignment_Params{root.Struct()}, err
}

func (s ViewSharingLink_getRoleAssignment_Params) String() string {
	str, _ := text.Marshal(0xb8083dd65a24c770, s.Struct)
	return str
}

// ViewSharingLink_getRoleAssignment_Params_List is a list of ViewSharingLink_getRoleAssignment_Params.
type ViewSharingLink_getRoleAssignment_Params_List struct{ capnp.List }

// NewViewSharingLink_getRoleAssignment_Params creates a new list of ViewSharingLink_getRoleAssignment_Params.
func NewViewSharingLink_getRoleAssignment_Params_List(s *capnp.Segment, sz int32) (ViewSharingLink_getRoleAssignment_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ViewSharingLink_getRoleAssignment_Params_List{l}, err
}

func (s ViewSharingLink_getRoleAssignment_Params_List) At(i int) ViewSharingLink_getRoleAssignment_Params {
	return ViewSharingLink_getRoleAssignment_Params{s.List.Struct(i)}
}

func (s ViewSharingLink_getRoleAssignment_Params_List) Set(i int, v ViewSharingLink_getRoleAssignment_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_getRoleAssignment_Params_Promise is a wrapper for a ViewSharingLink_getRoleAssignment_Params promised by a client call.
type ViewSharingLink_getRoleAssignment_Params_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_getRoleAssignment_Params_Promise) Struct() (ViewSharingLink_getRoleAssignment_Params, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_getRoleAssignment_Params{s}, err
}

type ViewSharingLink_getRoleAssignment_Results struct{ capnp.Struct }

func NewViewSharingLink_getRoleAssignment_Results(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ViewSharingLink_getRoleAssignment_Results{st}, err
}

func NewRootViewSharingLink_getRoleAssignment_Results(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ViewSharingLink_getRoleAssignment_Results{st}, err
}

func ReadRootViewSharingLink_getRoleAssignment_Results(msg *capnp.Message) (ViewSharingLink_getRoleAssignment_Results, error) {
	root, err := msg.RootPtr()
	return ViewSharingLink_getRoleAssignment_Results{root.Struct()}, err
}

func (s ViewSharingLink_getRoleAssignment_Results) String() string {
	str, _ := text.Marshal(0x9d159666de73f39d, s.Struct)
	return str
}

func (s ViewSharingLink_getRoleAssignment_Results) Var() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s ViewSharingLink_getRoleAssignment_Results) HasVar() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_getRoleAssignment_Results) SetVar(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// ViewSharingLink_getRoleAssignment_Results_List is a list of ViewSharingLink_getRoleAssignment_Results.
type ViewSharingLink_getRoleAssignment_Results_List struct{ capnp.List }

// NewViewSharingLink_getRoleAssignment_Results creates a new list of ViewSharingLink_getRoleAssignment_Results.
func NewViewSharingLink_getRoleAssignment_Results_List(s *capnp.Segment, sz int32) (ViewSharingLink_getRoleAssignment_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ViewSharingLink_getRoleAssignment_Results_List{l}, err
}

func (s ViewSharingLink_getRoleAssignment_Results_List) At(i int) ViewSharingLink_getRoleAssignment_Results {
	return ViewSharingLink_getRoleAssignment_Results{s.List.Struct(i)}
}

func (s ViewSharingLink_getRoleAssignment_Results_List) Set(i int, v ViewSharingLink_getRoleAssignment_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_getRoleAssignment_Results_Promise is a wrapper for a ViewSharingLink_getRoleAssignment_Results promised by a client call.
type ViewSharingLink_getRoleAssignment_Results_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_getRoleAssignment_Results_Promise) Struct() (ViewSharingLink_getRoleAssignment_Results, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_getRoleAssignment_Results{s}, err
}

func (p ViewSharingLink_getRoleAssignment_Results_Promise) Var() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type GrainInfo struct{ capnp.Struct }

func NewGrainInfo(s *capnp.Segment) (GrainInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return GrainInfo{st}, err
}

func NewRootGrainInfo(s *capnp.Segment) (GrainInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return GrainInfo{st}, err
}

func ReadRootGrainInfo(msg *capnp.Message) (GrainInfo, error) {
	root, err := msg.RootPtr()
	return GrainInfo{root.Struct()}, err
}

func (s GrainInfo) String() string {
	str, _ := text.Marshal(0xb5fcc0e153671d68, s.Struct)
	return str
}

func (s GrainInfo) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GrainInfo) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GrainInfo) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GrainInfo) SetAppId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s GrainInfo) AppVersion() uint32 {
	return s.Struct.Uint32(0)
}

func (s GrainInfo) SetAppVersion(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s GrainInfo) Title() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s GrainInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GrainInfo) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s GrainInfo) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// GrainInfo_List is a list of GrainInfo.
type GrainInfo_List struct{ capnp.List }

// NewGrainInfo creates a new list of GrainInfo.
func NewGrainInfo_List(s *capnp.Segment, sz int32) (GrainInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return GrainInfo_List{l}, err
}

func (s GrainInfo_List) At(i int) GrainInfo { return GrainInfo{s.List.Struct(i)} }

func (s GrainInfo_List) Set(i int, v GrainInfo) error { return s.List.SetStruct(i, v.Struct) }

// GrainInfo_Promise is a wrapper for a GrainInfo promised by a client call.
type GrainInfo_Promise struct{ *capnp.Pipeline }

func (p GrainInfo_Promise) Struct() (GrainInfo, error) {
	s, err := p.Pipeline.Struct()
	return GrainInfo{s}, err
}

type AppPersistent struct{ Client capnp.Client }

func (c AppPersistent) Save(ctx context.Context, params func(AppPersistent_save_Params) error, opts ...capnp.CallOption) AppPersistent_save_Results_Promise {
	if c.Client == nil {
		return AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaffa789add8747b8,
			MethodID:      0,
			InterfaceName: "grain.capnp:AppPersistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(AppPersistent_save_Params{Struct: s}) }
	}
	return AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type AppPersistent_Server interface {
	Save(AppPersistent_save) error
}

func AppPersistent_ServerToClient(s AppPersistent_Server) AppPersistent {
	c, _ := s.(server.Closer)
	return AppPersistent{Client: server.New(AppPersistent_Methods(nil, s), c)}
}

func AppPersistent_Methods(methods []server.Method, s AppPersistent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaffa789add8747b8,
			MethodID:      0,
			InterfaceName: "grain.capnp:AppPersistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := AppPersistent_save{c, opts, AppPersistent_save_Params{Struct: p}, AppPersistent_save_Results{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	return methods
}

// AppPersistent_save holds the arguments for a server call to AppPersistent.save.
type AppPersistent_save struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  AppPersistent_save_Params
	Results AppPersistent_save_Results
}

type AppPersistent_save_Params struct{ capnp.Struct }

func NewAppPersistent_save_Params(s *capnp.Segment) (AppPersistent_save_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return AppPersistent_save_Params{st}, err
}

func NewRootAppPersistent_save_Params(s *capnp.Segment) (AppPersistent_save_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return AppPersistent_save_Params{st}, err
}

func ReadRootAppPersistent_save_Params(msg *capnp.Message) (AppPersistent_save_Params, error) {
	root, err := msg.RootPtr()
	return AppPersistent_save_Params{root.Struct()}, err
}

func (s AppPersistent_save_Params) String() string {
	str, _ := text.Marshal(0xf0136e14d8019d3c, s.Struct)
	return str
}

// AppPersistent_save_Params_List is a list of AppPersistent_save_Params.
type AppPersistent_save_Params_List struct{ capnp.List }

// NewAppPersistent_save_Params creates a new list of AppPersistent_save_Params.
func NewAppPersistent_save_Params_List(s *capnp.Segment, sz int32) (AppPersistent_save_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return AppPersistent_save_Params_List{l}, err
}

func (s AppPersistent_save_Params_List) At(i int) AppPersistent_save_Params {
	return AppPersistent_save_Params{s.List.Struct(i)}
}

func (s AppPersistent_save_Params_List) Set(i int, v AppPersistent_save_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppPersistent_save_Params_Promise is a wrapper for a AppPersistent_save_Params promised by a client call.
type AppPersistent_save_Params_Promise struct{ *capnp.Pipeline }

func (p AppPersistent_save_Params_Promise) Struct() (AppPersistent_save_Params, error) {
	s, err := p.Pipeline.Struct()
	return AppPersistent_save_Params{s}, err
}

type AppPersistent_save_Results struct{ capnp.Struct }

func NewAppPersistent_save_Results(s *capnp.Segment) (AppPersistent_save_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return AppPersistent_save_Results{st}, err
}

func NewRootAppPersistent_save_Results(s *capnp.Segment) (AppPersistent_save_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return AppPersistent_save_Results{st}, err
}

func ReadRootAppPersistent_save_Results(msg *capnp.Message) (AppPersistent_save_Results, error) {
	root, err := msg.RootPtr()
	return AppPersistent_save_Results{root.Struct()}, err
}

func (s AppPersistent_save_Results) String() string {
	str, _ := text.Marshal(0xba36a34b4eeb483f, s.Struct)
	return str
}

func (s AppPersistent_save_Results) ObjectId() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s AppPersistent_save_Results) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s AppPersistent_save_Results) ObjectIdPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s AppPersistent_save_Results) SetObjectId(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s AppPersistent_save_Results) SetObjectIdPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s AppPersistent_save_Results) Label() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s AppPersistent_save_Results) HasLabel() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s AppPersistent_save_Results) SetLabel(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s AppPersistent_save_Results) NewLabel() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// AppPersistent_save_Results_List is a list of AppPersistent_save_Results.
type AppPersistent_save_Results_List struct{ capnp.List }

// NewAppPersistent_save_Results creates a new list of AppPersistent_save_Results.
func NewAppPersistent_save_Results_List(s *capnp.Segment, sz int32) (AppPersistent_save_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return AppPersistent_save_Results_List{l}, err
}

func (s AppPersistent_save_Results_List) At(i int) AppPersistent_save_Results {
	return AppPersistent_save_Results{s.List.Struct(i)}
}

func (s AppPersistent_save_Results_List) Set(i int, v AppPersistent_save_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppPersistent_save_Results_Promise is a wrapper for a AppPersistent_save_Results promised by a client call.
type AppPersistent_save_Results_Promise struct{ *capnp.Pipeline }

func (p AppPersistent_save_Results_Promise) Struct() (AppPersistent_save_Results, error) {
	s, err := p.Pipeline.Struct()
	return AppPersistent_save_Results{s}, err
}

func (p AppPersistent_save_Results_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p AppPersistent_save_Results_Promise) Label() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type MainView struct{ Client capnp.Client }

func (c MainView) Restore(ctx context.Context, params func(MainView_restore_Params) error, opts ...capnp.CallOption) MainView_restore_Results_Promise {
	if c.Client == nil {
		return MainView_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      0,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(MainView_restore_Params{Struct: s}) }
	}
	return MainView_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c MainView) Drop(ctx context.Context, params func(MainView_drop_Params) error, opts ...capnp.CallOption) MainView_drop_Results_Promise {
	if c.Client == nil {
		return MainView_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      1,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(MainView_drop_Params{Struct: s}) }
	}
	return MainView_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c MainView) GetViewInfo(ctx context.Context, params func(UiView_getViewInfo_Params) error, opts ...capnp.CallOption) UiView_ViewInfo_Promise {
	if c.Client == nil {
		return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_getViewInfo_Params{Struct: s}) }
	}
	return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c MainView) NewSession(ctx context.Context, params func(UiView_newSession_Params) error, opts ...capnp.CallOption) UiView_newSession_Results_Promise {
	if c.Client == nil {
		return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newSession_Params{Struct: s}) }
	}
	return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c MainView) NewRequestSession(ctx context.Context, params func(UiView_newRequestSession_Params) error, opts ...capnp.CallOption) UiView_newRequestSession_Results_Promise {
	if c.Client == nil {
		return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 5}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newRequestSession_Params{Struct: s}) }
	}
	return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c MainView) NewOfferSession(ctx context.Context, params func(UiView_newOfferSession_Params) error, opts ...capnp.CallOption) UiView_newOfferSession_Results_Promise {
	if c.Client == nil {
		return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 6}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newOfferSession_Params{Struct: s}) }
	}
	return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type MainView_Server interface {
	Restore(MainView_restore) error

	Drop(MainView_drop) error

	GetViewInfo(UiView_getViewInfo) error

	NewSession(UiView_newSession) error

	NewRequestSession(UiView_newRequestSession) error

	NewOfferSession(UiView_newOfferSession) error
}

func MainView_ServerToClient(s MainView_Server) MainView {
	c, _ := s.(server.Closer)
	return MainView{Client: server.New(MainView_Methods(nil, s), c)}
}

func MainView_Methods(methods []server.Method, s MainView_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      0,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := MainView_restore{c, opts, MainView_restore_Params{Struct: p}, MainView_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      1,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := MainView_drop{c, opts, MainView_drop_Params{Struct: p}, MainView_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_getViewInfo{c, opts, UiView_getViewInfo_Params{Struct: p}, UiView_ViewInfo{Struct: r}}
			return s.GetViewInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 8},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newSession{c, opts, UiView_newSession_Params{Struct: p}, UiView_newSession_Results{Struct: r}}
			return s.NewSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newRequestSession{c, opts, UiView_newRequestSession_Params{Struct: p}, UiView_newRequestSession_Results{Struct: r}}
			return s.NewRequestSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newOfferSession{c, opts, UiView_newOfferSession_Params{Struct: p}, UiView_newOfferSession_Results{Struct: r}}
			return s.NewOfferSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// MainView_restore holds the arguments for a server call to MainView.restore.
type MainView_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  MainView_restore_Params
	Results MainView_restore_Results
}

// MainView_drop holds the arguments for a server call to MainView.drop.
type MainView_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  MainView_drop_Params
	Results MainView_drop_Results
}

type MainView_restore_Params struct{ capnp.Struct }

func NewMainView_restore_Params(s *capnp.Segment) (MainView_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_restore_Params{st}, err
}

func NewRootMainView_restore_Params(s *capnp.Segment) (MainView_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_restore_Params{st}, err
}

func ReadRootMainView_restore_Params(msg *capnp.Message) (MainView_restore_Params, error) {
	root, err := msg.RootPtr()
	return MainView_restore_Params{root.Struct()}, err
}

func (s MainView_restore_Params) String() string {
	str, _ := text.Marshal(0x9ad62de07dfc6419, s.Struct)
	return str
}

func (s MainView_restore_Params) ObjectId() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s MainView_restore_Params) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_restore_Params) ObjectIdPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s MainView_restore_Params) SetObjectId(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s MainView_restore_Params) SetObjectIdPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// MainView_restore_Params_List is a list of MainView_restore_Params.
type MainView_restore_Params_List struct{ capnp.List }

// NewMainView_restore_Params creates a new list of MainView_restore_Params.
func NewMainView_restore_Params_List(s *capnp.Segment, sz int32) (MainView_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return MainView_restore_Params_List{l}, err
}

func (s MainView_restore_Params_List) At(i int) MainView_restore_Params {
	return MainView_restore_Params{s.List.Struct(i)}
}

func (s MainView_restore_Params_List) Set(i int, v MainView_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_restore_Params_Promise is a wrapper for a MainView_restore_Params promised by a client call.
type MainView_restore_Params_Promise struct{ *capnp.Pipeline }

func (p MainView_restore_Params_Promise) Struct() (MainView_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return MainView_restore_Params{s}, err
}

func (p MainView_restore_Params_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_restore_Results struct{ capnp.Struct }

func NewMainView_restore_Results(s *capnp.Segment) (MainView_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_restore_Results{st}, err
}

func NewRootMainView_restore_Results(s *capnp.Segment) (MainView_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_restore_Results{st}, err
}

func ReadRootMainView_restore_Results(msg *capnp.Message) (MainView_restore_Results, error) {
	root, err := msg.RootPtr()
	return MainView_restore_Results{root.Struct()}, err
}

func (s MainView_restore_Results) String() string {
	str, _ := text.Marshal(0x99efcebf23bbae35, s.Struct)
	return str
}

func (s MainView_restore_Results) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s MainView_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_restore_Results) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s MainView_restore_Results) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s MainView_restore_Results) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// MainView_restore_Results_List is a list of MainView_restore_Results.
type MainView_restore_Results_List struct{ capnp.List }

// NewMainView_restore_Results creates a new list of MainView_restore_Results.
func NewMainView_restore_Results_List(s *capnp.Segment, sz int32) (MainView_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return MainView_restore_Results_List{l}, err
}

func (s MainView_restore_Results_List) At(i int) MainView_restore_Results {
	return MainView_restore_Results{s.List.Struct(i)}
}

func (s MainView_restore_Results_List) Set(i int, v MainView_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_restore_Results_Promise is a wrapper for a MainView_restore_Results promised by a client call.
type MainView_restore_Results_Promise struct{ *capnp.Pipeline }

func (p MainView_restore_Results_Promise) Struct() (MainView_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return MainView_restore_Results{s}, err
}

func (p MainView_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_drop_Params struct{ capnp.Struct }

func NewMainView_drop_Params(s *capnp.Segment) (MainView_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_drop_Params{st}, err
}

func NewRootMainView_drop_Params(s *capnp.Segment) (MainView_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return MainView_drop_Params{st}, err
}

func ReadRootMainView_drop_Params(msg *capnp.Message) (MainView_drop_Params, error) {
	root, err := msg.RootPtr()
	return MainView_drop_Params{root.Struct()}, err
}

func (s MainView_drop_Params) String() string {
	str, _ := text.Marshal(0x8c519e0dedc17d73, s.Struct)
	return str
}

func (s MainView_drop_Params) ObjectId() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s MainView_drop_Params) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_drop_Params) ObjectIdPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s MainView_drop_Params) SetObjectId(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s MainView_drop_Params) SetObjectIdPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// MainView_drop_Params_List is a list of MainView_drop_Params.
type MainView_drop_Params_List struct{ capnp.List }

// NewMainView_drop_Params creates a new list of MainView_drop_Params.
func NewMainView_drop_Params_List(s *capnp.Segment, sz int32) (MainView_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return MainView_drop_Params_List{l}, err
}

func (s MainView_drop_Params_List) At(i int) MainView_drop_Params {
	return MainView_drop_Params{s.List.Struct(i)}
}

func (s MainView_drop_Params_List) Set(i int, v MainView_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_drop_Params_Promise is a wrapper for a MainView_drop_Params promised by a client call.
type MainView_drop_Params_Promise struct{ *capnp.Pipeline }

func (p MainView_drop_Params_Promise) Struct() (MainView_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return MainView_drop_Params{s}, err
}

func (p MainView_drop_Params_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_drop_Results struct{ capnp.Struct }

func NewMainView_drop_Results(s *capnp.Segment) (MainView_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return MainView_drop_Results{st}, err
}

func NewRootMainView_drop_Results(s *capnp.Segment) (MainView_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return MainView_drop_Results{st}, err
}

func ReadRootMainView_drop_Results(msg *capnp.Message) (MainView_drop_Results, error) {
	root, err := msg.RootPtr()
	return MainView_drop_Results{root.Struct()}, err
}

func (s MainView_drop_Results) String() string {
	str, _ := text.Marshal(0x9210d9e69d14fa35, s.Struct)
	return str
}

// MainView_drop_Results_List is a list of MainView_drop_Results.
type MainView_drop_Results_List struct{ capnp.List }

// NewMainView_drop_Results creates a new list of MainView_drop_Results.
func NewMainView_drop_Results_List(s *capnp.Segment, sz int32) (MainView_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return MainView_drop_Results_List{l}, err
}

func (s MainView_drop_Results_List) At(i int) MainView_drop_Results {
	return MainView_drop_Results{s.List.Struct(i)}
}

func (s MainView_drop_Results_List) Set(i int, v MainView_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_drop_Results_Promise is a wrapper for a MainView_drop_Results promised by a client call.
type MainView_drop_Results_Promise struct{ *capnp.Pipeline }

func (p MainView_drop_Results_Promise) Struct() (MainView_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return MainView_drop_Results{s}, err
}

const schema_c8d91463cfc4fb4a = "x\xda\xdcz{xT\xd5\xb5\xf8\xd93\x99\x9c\xc4&" +
	"\x9d\x1c\xce\xc40<\x1a\x88\xe9\x8f$\xe6\x01A|\xa4" +
	"h^ \x06Q3y \xe4g\xfc8\x93\x9c$#" +
	"\x93\x99\xc9\xcc\x84\x10*E\xbdp\x81\xaa\xb7B\xf1\x0a" +
	"|\x06+\x95+p\xa5V\x0b\x0aV*Z\xad\x80\x0f" +
	"\x0a\"\"\x96\x0a(\x14\x11* \xa2\x820w\xad}" +
	"\xce>\xb3g2y\xf1\xdd\xaf\x7f\\\xbeo\xc8\xcc:" +
	"k\xaf\xbd\xf6z\xef\xb5\xce\xe8\"k\x89i\x8c\xa58" +
	"K\x10\xaa\xf7\x12K|\xe8\xd05\xf5\x93\xcf\x8c\xf8l" +
	"\x81 \x0d%\x82`!\xa2 \x8c\xb5\xd8\xea\x88@d" +
	"\xc9V,\x90\xd0Ae\xfc\xfb\xb5\x15\x1f/\x14$\xbb" +
	"\x86\x90B\xc6\x8e\xb1\x95!\xc2M\x14!0\xf7\xf5S" +
	"\xc9\xab\x1c\x8f\x08R\xaa\x810\xddv\x15\"(\x14\xe1" +
	"\xdf\xef\xdcW\xf1\xf8\xb9\xfc_\x09\x92\x0c\x08q\xb8\xc3" +
	"\x02[!\x11\xe2B\xbfx\xf7\xb3\x0f\xd6\xee\x8c_\xca" +
	"\xd3n\xc5GDn\xa7K\xc7]\xb0u\x1d\xfb8e" +
	"\xa9F;\x0e\x9f?n\x1b\x84KoJ\xda\x7fk{" +
	"\xb9\xed\x09A\x1al\x0a\x15\xaf^\x97\xee\xbc\xf7\x8f\x17" +
	"\x04\x81\x8c\x9d\x0b[\xcb\x8f\xda`\x17y\xb1\xed\x06\xa0" +
	"qr\x9a\xfd\xe1m\x8f\x8dZ\x8e\xdb\x87\xfeq\xa4\xf9" +
	"\xc4\xf2\x8f6~\xa2\x9d\x14\xf0.\xc8+(\xee\xe3\xda" +
	"~\xbf\xfb\xe35\xaf\xbd\xff\xd5\x0a\xfe,\x9bl9\xc8" +
	"\xd0V\x8a`o\xfca\xee\xa1\xbc}+y\x84\x83\xb6" +
	"\x0cD8J\x11v\xefl\xbb\xd5<\xeac@\x18l" +
	"\xc831\x95\xca35\x15\x11\xba\xbe\x0e\xfc\xbd\xe9?" +
	"S\xbb\x04i\x98\x81\xa0\xa6\xbeH\xcfL\x11\x9e\x8c\xbf" +
	"\xeb\x93\x0b\xa6\xd2.M#T\\o\xa5N\xc63w" +
	"\x8e8I\x1e\xf1\xbd\xb4J\x13\xa4\xb6\xf4\x05\x8d\xf6V" +
	"\xbat\xe9\xa5\xe4\x8a\xaf\xaew?\xa5+\x93\xae=\x98" +
	"\xfaKD8\x95\xfa\x0f@\x18Y\xff\xe5\xe1\xa5\xd6\xbd" +
	"O\xf1\x02?|5\xa5p\xeaj\xa40\xfb\x89\xbb\xbf" +
	"\x1f\x96\x97\xb3\x9a\xdf\"9\xed>D\xb0\xa7!\xc2=" +
	"\x8bZ?\x9d^\xb4p\xb5F\x81j\xe4\xa6\xb4\xa5\xc8" +
	"\xdd\xf1\xcb\x7f\xbecV\xc1w\xcf\x08R\x9294\xf9" +
	"\xe2\x9b\xbb\x1al\x1fo\x07\x8d\xc8\xd9iG\xe4qi" +
	"i \xe4\xd2\xb4Ir\x1b~\x0b\xf9\xb2fNL|" +
	"n\xdc\x1a\x8e\x8c\x92\xb6\x1a\xc9|\xbal\xf1\xfb\x7f8" +
	"U\xb4Fp\xc8\x04Y\x88G\x16\x1ciU\xc0\xc2\xd8" +
	"\xfa\xb4_\xc1\x9f\xd0\x91\xdb^\xda\xd7\xbe\xfe\xcf\x802" +
	"\x98\x18\\~3\xb8\x08\xb9\xbc4\xb8\x030\x9a\xb2\x7f" +
	"t\xf7\xa0]\xf7>\xabi\x81\xcaA\xb1S\x93{{" +
	"\xd9\xd2\xac'B\xe7\xd6\xf2\x07\xac\xb0S\x93s\xd8\xf1" +
	"\x80\x0f\xff\xe6'\xcd;6\xdc\xb2\x8e\x13\x7f\x9b\xdd\x89" +
	"K\xd5{\xb2\x16\xcc\xcbxn\x03/\xbcz;5\x0e" +
	"\x95.\xdd<i\xe1\xc1\x95\xb3/<\x1f%\x01!\x85" +
	"\xc8\x0b\xec{\xe4%vjr\xf6I\xf2A\xbb\x08\x1f" +
	"k\xe8\xa3\x8d\xabf<w\xfe\xb3\x17x\x8a\x1f\xda\xf1" +
	"\xa8\xf2aJqNq\xdd\xb9\xf5\xbbr7rb\"" +
	"CV\"3CWm\xbe\xe6\xc3\xa3.\xfd\x89\xc5\x84" +
	"\x8f\xcehK/\xd9Q\x04\xc7\xf3~\xb1c\xe5\xbe\x93" +
	"\x1b\xb9s(C\xa8\xa2Z\x867W\x1f\xde\xf6\xc3&" +
	"\xc1\x91DH\x98O\xa0\x01\xec9\x86\xbc(O\x1f\x82" +
	"\xdfj\x87<\x0fd\x16\xfd\xfc\x9d\xf7;\xf6\xff\xe8e" +
	">>\x9c\x19\xf2\x06\xeeC\x86\"\x8b\xbe\xb73\xeb\xf6" +
	"\xdd\x9c\xb0Y\xb3gm\x9f\xa1\xebq\x1f\xcb\xb37<" +
	"x\xf1-\xef\x16\x9eE\xc7P\xcab\xfdPd1\xfb" +
	"\x87\xcb\xff\xbc\xbd`\xdf\x16\xeet[\x87\xee\xc1\xa5\xc5" +
	"\xb7}y\xe7\xed\xbf\xbd\xfe\x15A\x1ab,}a(" +
	"U\xf0\x16\xba\xf4\x9d\x87\xa7\xdb\xcb\x8a\xec\xaf2+\xb1" +
	"\xe0\xbe\xf6aN\xb4\x92\x9f\x0e\xbb\x1b\xadD=\xb6\xb6" +
	"t\xd2\xb8{_\x8d\xf2\xfb\x04<\xdb\x92\xe1G\xe4\xae" +
	"\xe1\xb8f\xc5\xf0\xb7\x11\xf9\xfa\xaf\x1b\xea\x86\x97\xdc\xf2" +
	"Z7\xe3=\x98~@\xfe\"\x1d1\x8f\xa6O2\xc9" +
	"\xae\x91\xf05\xf4\xc1\xe4\x15\xbf\xdf1\xa2v[7l" +
	"\xc7\xc87\xe4\xe9#\xa9\xf4FN\x92\x1f\xa4\xc8?l" +
	"?\x92\xf3\xd0\x17\x1do\xc4\xb0\x0a\xd7\xc8\xf5r\x1bE" +
	"o\x1d\xb9PN\xceH\x93S3\xac!w\xc2\xaf\xf7" +
	"\x1fo\xfb\xc9\x9b\xba\xc8\xcd\xd4\x073\x9c\xd4\x073P" +
	"'u\x7f\xfb\xfcD\xa3\xaf\xf8\x9dh\x05\xa2\xf0\xe5\xf7" +
	"2V\xcb\x1ff\xe0\x9a\xdd\x19T\x0e\xef\x98|9\xbb" +
	"\xeeR\xf7\xc4\xd8\xff\xe6\xcc\x9drE&\xe2N\xcc\\" +
	"h\x92\x1f\x1d%\xc2\xc7\x1a\xfa\xd9F\xe7\xc9;\xcdm" +
	"{8\xbd,\x18\xb5\x13\xf5r\xda~\xc3\xfcM]\xab" +
	"\xf6\xf0\x06\xdb9\x8a&\x83\x07G\xa15\\Zyh" +
	"\xea\xb7\xf3\xff\xdf^\x9dw4\xa9\xb1O\x8f\xa2\x01f" +
	"\xc3(T\xdc\xb4\xf4_\xbeR\xfe_K\xf7\xf1F\x91" +
	"\x9cE\x9d(5\x0b\x11v\x1dyr\xc9\xd9g\xe4\x03" +
	"<\x85\xb6\xac9\x880\x97\"\x18\xaa\x8c\x16\xfe7Y" +
	"+\xe5KY\xa3\xd0\x10\xb2E\"\xbf\x97\x8d\xd2\xaf\xff" +
	"z~\xd1\xfc\xd7\xfe\xfb\x18\xbf\xdf\xa6l\xca\xd0\xeb\xd9" +
	"Hn\xa6\xe7@\xc7\xae\xe9-_p~22\xe7\x15" +
	"<\xec-'\x1f8\xfd|\xf6\xf8/\xf9\xa5R\xced" +
	"\\:<\x87\x9e\xe5\xc6\x85S?mZt\x92\x97\xc6" +
	"\xdc\x9c\x95\x88\xf0h\x0eJc\xcc\xdb\x89eS\xce\x9c" +
	"\x89@\xd8\xa0Q\xd8B\x11\xa6\xbe<\xf5\xf6V\xf5\xc4" +
	"W\xbcw}\x9c\xf3\x10M'\x14a|\x17\xd9o\xf3" +
	"\xc8\xa75G\xa0\xaa\xb0\\K\x03\xd9\xa6#\xc7\xc7\xfd" +
	"\xe9\xec\x88\xd3\x82c\x18\x18\x81\x11|'\x9aD3\x10" +
	"9\x95S\x07\x0ez-\xd2\xbb\x94\x93\x8ef\xf0z\xcd" +
	"\xf5\x89S\x07\xff\xfa4\x17\x12o\xce\xadBJ{\xe7" +
	"w\xfe\xe5\xcb\x19\xb9g8\x09d\xe7\x96\xe1\x93\x8b\x07" +
	"\x16/:\xbaf\xc2\x99hCC\x93\x94\xa5\xdc=\xf2" +
	"\xf0\\\xfcf\xcf\xc5\xdc\xb2\xbeeJu\xc2\xbd\xf3\xbe" +
	"\x16\x1cC\x89\xa1\xb9or\xa9\xa8I\x1e\xda\xad\xf7\xea" +
	"M\xd3\xfeYZx\xbe\x9b\xe6\xba\xf2^\x94\xd7\xe4!" +
	"\xa9\xa7\xf3D\xfc\x80\xe2\x0e\x1f\xff\x8f\xac[\x17\xfc\xec" +
	"<o\x07\x8b\xf3h\xf4x<\x0f\xa5\x7f\xf4\xd9\xdb\x1f" +
	"xS\xfa\xfe</;{>MU?\xcdG\xd9e" +
	"\xb5\xacx\xaa.w\xcew,F\xd0\x93M\xccG[" +
	"\x1b{G>\x15\xca\xb73\xd6-\xbbs\xe6\xe5\xef8" +
	"Kw\x15\xd0\\{\xcd\xeb\xe2\x9e\xbfl\xdf\xf7='" +
	"\x94\xda\x02\x9a\x06\x8cb#\xfa\x14\xa5\x05o\xc8\x15\x05" +
	"\x98\xe7j\x0b&\xc9s\x0b\xf0\x14\xc7~\xb7*\xe9\xc0" +
	"\x8d7_\xe43\xb2Z@\x1d\xa6\xb5\x00\xa56\xfdU" +
	"\xff\xfd\xd7\x8fx\xf5\"\x9f\x08GS\x05\xff\xff\xad7" +
	"\xd5\x9f\x19\xd7v\x99\xe3\xa0b4\xf0\x16\x12\x8e\x87\x9a" +
	"\xfd\x8a\xcb\x93\xdf\xa0\xc4\xf9<\xbe\xa2j5\x10py" +
	"=\xe5^OP\x9d\x1d\xccW\x1a\x82\xaeY\xae`g" +
	"f\xa5\xe2W\xcc\xad\x01G\x9c9\x0e\xd6\x03\x0d)\xb9" +
	"P\x10\x1c\x09f\xe2\xb0\x99H\xba:K\xf5\x04IJ" +
	"\xe8v\xf7#\xebN\xbc\xb4\xedY8\x02I\x01\x8e\"" +
	"i+\x9e\xc6@\xd0\xebo-\xf5\xb9\xf2\x1bU\xb7\x1a" +
	"T\x1b)\xe1\xd6\x80\xc0\x13\xce\x08\x13\x16\xfdj\x13\x19" +
	"D\xb8\xc0\x03\x94\x07q\x84\xcdH\xf8\x0e\xf8>\xd5\xa5" +
	"v\xe47\xfa\xbd>F1\x82\xe4d\xf8\x99\x04$\x07" +
	"\x9b\xc0r\x9c\xf7\xa9\x0d\xc1\x8aFJ\x0aH\x1b15" +
	"\x16\xe9Z\x17%\xdc\xac\x06\xf1o\x85\xa7\xc9\x9bY\x99" +
	"N7\x88\xc4\x8b8[@\x99\xa5fV\xa9\xe9\x81v" +
	"w\xb0G\x91\x05\xbd3U\x0fI\x16L\xf0\xe9\xf5<" +
	"U*\xd2!\xd1\x1b\x06\x95\xa0\xab\xa14\x10P\x83\xf9" +
	"\x95~o\xd0\xdb\xe0u\x0bB%!\x8e\x04b\x82\xbd" +
	"\xa4B\xd4\x81\x94\x98#\x08\xe9-\xc1\xa0/`\xc5\xff" +
	"\x0d\x1a&\xeep\x95\xde\x0e\xd5_\xec\xf4\xce\xaeQ\x9a" +
	"\x91@O\x0c\xbb\x82n\x95$\x01\xc3I=2\xecW" +
	"Q\x08x\xf8\x805\xfa\xf0\xbcZ\x1b\x14\x1f\x19\x14g" +
	"\x06\x0e\x07\xf5I\x0b\xf4)*\x91\xa6w%\xea\xd4L" +
	"\xb0E\xf1\xbb<\xcdS\\\x9e\x99\xa8\xd3J5\xe8Q" +
	"ZUM\xc4\xe6Hvstv'\x98\x88\x15\x91\x88" +
	"\x14Z\xb1$c\xf9\xc6\xda\xb3'\x80t\x09\x91H\xba" +
	"#\x0e\xf6\xe7\x80\x12Is\xc4\x11\xf8Wi\x06\xf3\x0f" +
	"U\x1c\xf9\xf9\x0d\xfe\x97\xeb\x1f\xa6\x8a\xe08\xb1 '" +
	"x\xc4(n\xaa\xbcn\x15\x14\xeaj\xf6\xb4\x827\xc5" +
	"4\xa0\x8c0S\xe2,\xc5?p\x9e\x8cH\x1f\xc5S" +
	",\xe7op{\x03\xbal\x82\xe0O\x91\xb8\xba\xedx" +
	"\xd4\x8e\xbb\x9a\x9aT\xbf\xbe4\xa6$\xcb\xc2\x8a\x9f\x17" +
	"\xd0\xf0\x80q#\x80\xc7\x92N\x14'M\xed\xee&\x97" +
	"\xdb]\xa5\xb6\xb5\x83Q\x84=<\xc5\xd8DA\xc9\xdc" +
	"\x03\x9b\xb4\x98@\x0a\xc4\x86AZRW\x03\xb0\x05\x80" +
	"\xf3\x01h2\xd9\xa8c<X\x07\xc0\x07\x00\xf8\x08\x00" +
	"\xcdf\x1b\x81\x14'-v\x02p\x11\x00\x97u3N" +
	"?l\xea\xf2\xab\x8d\xa4R\xf5\xb7\xba\x80-\xd1\xeb\x09" +
	"\x90\x1f\x0bT\xa0\x04\xbc\x01\xbe\x86\x1a\xd5@\x83\xdf\xe5" +
	"\x0b\x0af\xaf\x1f\xa4<m\xd8\xdb\xaf\xfd\xd6\xba\xeb\x14" +
	"\x0b\x83\x8d\xae\x80\xcf\xadtV\x08\"\x04\x10x~\xf2" +
	"\xda\xcf\xfe\x94\xb8\xbazM\xdfa2\x10T:K;" +
	"\x94\x99\xb1m\xb4(,\xd9\xe2\x16X\xe6F+=T" +
	"6c\xc6s\x99\xe7\x96\xc7TqXm\xba09\xc5" +
	"u\xf3\xd8+Q\\d\x8c\xf7\xf9\xd5\x06\x05\xc2|e" +
	"\xbb\xd3\xed\x0a\xb4h\x9a\xe3B\x99)\xda\x15D\xf0\x05" +
	"\xb4\xd7\xb0\xa1JdN\x88y\x86P\xac\xf9\x060i" +
	"\xc1\x8b\x9f~] \xec\x1e,IK\x05\x93\x94,\x86" +
	"\x987\x11\xe6N$XB(]\xa3\xdc\xe6\x0c\xba;" +
	"\xe7N\xa5af\xb3\xdf\xdb\xeei,e\x19\x10\xe5/" +
	"\x82\x80\xfa\xe5\x06\xfa9\x1d\xc3\x0can\xc2\x98\xb5\x11" +
	"\x84\xb9\x8d3\xd0\xad(\xe1\xcd\x00|\xd3D\x88n\x9f" +
	"\xaf\xa3)n\x03\xd8\xbbh\x9f&\xcd>w\xf8\x01\xb8" +
	"\x1d\x80{\x01\x18\x07F\x0bd\xa5\xdd\x18\x9d\xdf\x05\xe0" +
	"~\x00Z\xe2l\x04d\"}\x88\xe6\xbd\x17\x80\x87\x00" +
	"\x18o\xb1\x91x\x00\x1eD\xcc\xfd\x00\xfc\x1c\x82C{" +
	"@\xf5c\"\xc3\x00\x99\x12\xfa@\xaa\xff\xc3W\xbb\xb7" +
	",\xd3Mq^\x83\xe6r\xa0g\xe3\x1a\xa3\xebY\xb7" +
	"\x80\x1aA\xec\xf4\xa9$\x11,?1\x0c\xad\x14\xb4\xac" +
	"\x08.cB\xb7I\xf7\xa2887\xea\xddC\xd2\x83" +
	"\x8a\xb3\xa21v6\xe43\x1d\xe8\xb5\xd6\xef\xce\xac*" +
	"\xd6\x82\x12X'\x93o6\xca7\x0b\x8ex\x1d\x88R" +
	"\x17\xef\x18\x84\x8d\x06\xd8x8\xb6\xcfH\x92\x02\xb1\x86" +
	"[>\xb0\xbf\x15vl\xf1\x06\x82\x95J\xb0\x05\x9f\xc6" +
	"Lq1\xb8\x88Y\x08\x84\x0d\"\xecY\xdd\xe3\xf8\xc0" +
	"<\xab\xd7\xca\xac\x9b]v\xafG\xf4\xb2\xc8\xeaWZ" +
	"\x07X\x8eP\xff,\xf5\xf9 \xf6\x05\\\x81 \xfa\x91" +
	"V!\xa0\xfb\xb1\xfb\x04a7lI\x82\xa4Y\x9aD" +
	"J\x87\x11)O\xb4b\x19\xc42\x92\xd1\xce\xa0\x8a\x89" +
	"\x05,!@\xb94\x0e\x9c\xc3\x19\x82\x1d\xef\xc2\xb4." +
	"\x88\x90\xd8{\x8b\x8e\x105T\x94\xb7\xe1p1\x128" +
	"\x9c\xce:\x0bp@\xc0\xc6u\xef\x0aB\x17\x8b\xbf\xfd" +
	"\x0a\xd6\x8c\x1d\xce>\x9da\xfb4\xfc\x7f\xcc}\x9c\x81" +
	"F'\x89\x7f\xfbvo\xd9\xce\xb3[>`I\xc2\xe3" +
	"\x0d\xba\x9a\\\x0d\x8a`\x0dj\xe6\xf2^\xc9\xe2\xcd\x9f" +
	"\x0c_p\xf9\x0a2h\xb7\x9a\x92\xe0\x9aI\xf0#\xdd" +
	"\x83\xb1\x01\x95\x9cd\xf0>\x11\x0d\xa5\x04\xd8\x9c\x12\xf6" +
	"\xad\x0a\x8c3\xb7\x01\xac\x06s+\xd1b\x97\x03\x11\xa7" +
	"\x00p\x1aX\x94\xe2\xf3\x81K3g\x82_S\xd1\x88" +
	"\xc0\xff=$\x01\x80\x09BO%e|\x0c\xfe\xc1\xe5" +
	"0E\x80:\xf4,\x0cI\xd8(MbUH5F" +
	"\x85\x94Y\xb6oF\xf5\xf9\xb3\x0f\xf4\xabB\xe2\xf2\xfa" +
	"\x15\x94l\xba\x8f\xf5e\xb1\xe5J\xb8\xaa\xe7M\xa4*" +
	"\x96\x89\xa0\x19\xe7\x02\xf0F\xe0\x9a\xaen,W\x04\xc2" +
	"U(V70\x03\xe7\x0cg\xb6Hs\x88\xef\xc5\xb8" +
	"\xab\xd4ftk?f9V\xbdu\xbbj\xf0\xfe\x1f" +
	"\xd4/7\xbd\x07\xe00\xf7\x85\x1c\xf7\xd1\x95\xba\xe1\xfe" +
	"Z\xa5\x9e\xeeV\x9c\xaa;\xaanN\xe9w\x01\xc3\xae" +
	"\x09\x83\x0d\x8eV G\xcba\xf3g8\x8e\x9e\xc6\xd0" +
	"\xfb$\x00\xd7\x86S\xee\x1a\xf4\xcdg\x00\xf6{.\xe5" +
	"n\xc0\x94\xfb\x1c\x007s)w\x93S\xcf\xe2\x7f\xe5" +
	"R\xee{\x85\xe1\xe4\xfc\xaf\xc9\xae\xb4(\x85\xb3\xeb\xb1" +
	"B7\xda\xc8\xac\xfa\xe3\x1e\xb3*\x7f\xf7\xa3\xb7ZQ" +
	"\xf7\xf8\x11\x86\xe8v\xe39\xff\x0aG:\xc9\x89\xee\x0b" +
	"<\xe7\xe7fRE\xb8r\xfa\xd2RA\xa8\"fR" +
	"=\x8c\x84\xebi\xd9N@z\xd5\x83\x11~\x1d\xc2\xe3" +
	"\xe2\xa8\xfc\xe41\x04\x08W\x8fF\xf8\x14\x84[,T" +
	"\x84r\x05\x01]U\xdf\x86\xf0\x1a\x84\xc7\xc7\xd3\xc2E" +
	"v\x10\xf0\x89\xeaJ\x84\xdf\x83pQ\xb4\xd1Y\xcbt" +
	"\x02\xe1\xa7z\x1a\xc2\x83\x04\x93\xbb\x1e\x15\x04\xbe8O" +
	"\x09\xf7\x9cty\xf8\xc1]\xb9\xc7F\xefS{\x0c5" +
	"\x8a\xc7\x85\x11\x86\xb0\x10C\xba\x15\xfa\xadJ\xb0\xa1\x05" +
	",OHG\xf9\x07z\x14\xbd\x86\x085\xa1 \x82\xeb" +
	"\xf4\x8c\x06a\xb1\x06\x83\xa0f.1-\xbf\xa2\xc1+" +
	"\x10\x8c\xf7\xdcd\x8a\x9a\x0cm\xba\xd4t\xfa\x043\x7f" +
	"\xa8kWn\xff\xe8\xaa\xbf\xdfx\x8c\xed\x10\xa1u>" +
	"\xaa\x9ag\xd3\\\x9eIs9\xeb\\\x12\xd6\xa3\x97%" +
	"\xb2^0\xc9\xc9D\x04Oe]d\xc2\xfad2\x01" +
	"\xc5\x98\xa4\xefEb2\xdaR\x845\x99\xa4S\x85\xf0" +
	"\xec\xb0H\xccF\xe3\x8d\xb0^\xae\xf4a\x19<\xdb!" +
	"\x12\x8b1`\"l\xba m\x9d\x03\xcf6\x89$\xde" +
	"h#\x126\xc0\x92\xd6!\xcd.\x91\x88Fo\x90\xb0" +
	"\xb6\x9a\xb4d2<[,\x928\xa3\xe1KX3T" +
	"\x9a{\x1f<k\x17I\x821\x9d$l*#\xb9p" +
	"\x9dBo\x0b4\xb9\x10\x96]\xac\x98^JH(\xe8" +
	"Rk\xbc\xb5\x01U \xfe\x12\xbd\xa4-!\xf3t\xe7" +
	"\x83\xe7,\xab\x0a\xc5\x0c\x94N/\xcd\xf0\xc8\xebSi" +
	"\x1b\x83f\x9dP\x83[q\xb5\xa2\xd9X\xf5\x95\xac|" +
	"\x13\xf4\xc2'JM\xe1,#\x84\xeb-\xd6u%l" +
	".(Iup\x80Dz\x00\xda\xca\x10\xcc\xadj$" +
	"9\xc2Z*Vd\x86\xb6\x87(-6{$lJ" +
	")\x8d\x81\xc0X\x9aKJ\xc7\x13\xe9\x0e\xd49\x1b\xc5" +
	"\x1267\x95J\xb1\xba\x1bOJ\xa7\x10\xa9^\x9c\xa7" +
	"\xb7gXB5\xda.|\x81\x17\x01,!V,D" +
	"\x07\xb2\x80\xde\xd8\xc2E\x9b\xd0\x8f\xf20\xb2l`\x0a" +
	"\xf4\x1b\xe9\x81\xabj2\xc2U\x8d\x11\xe3*V\xeb\x15" +
	"L\x0b\x17\xe3T\x8c\x86\x8d\x00\xf4]aw\xa0\x9f\xd7" +
	"\x7f\xaa,,)\xc4\x09j\x13\xea\x8a\xcbe\x18{\x97" +
	"\x01\x0b\xbf\xe1\x98\xed\xaa\x0b\xe72\x83\xd9\xc8d\xa6\xf7" +
	"768\xf5d\xb6\x1d\xb2\x1e\xc6b\x88 oaz" +
	"|\x93\xe52 i\x01\xe0{e\xe1\\\xa6\x97g\xdd" +
	"\xc3\xd2,\xd5\xef\xacl\xf1\x83\xb9\x05b=6\xaez" +
	"\"\x86\xd0\xee\xcfc\x06k&+\xaf3\xe0\xc5\xce0" +
	"\x06D\x84\x11\xc8\x9a\x8dj\x93\x82u\x89\xfe;\xcaW" +
	"\xb8\xcaFk\x80fQ\x0bg\xc3e\xc2\xe6\x9e\x90|" +
	" W\x95_GH\xf9\x04B\xe4Z\x1a\xd9\xd8x\x8a" +
	"\xb0\xf9!\xa4\xa2\x9d\x80\x05t\xcag\x00V\x1b\xc1\x10" +
	"\xc7\xa67\x84\x8d!e\x15\x13Vy#`\x01[\xf2" +
	"\x02\x82\xc1\x8eM`\x09\x9b\x13\xc9\x9d\x98\xbe\xcag\x03" +
	"\xd6\"\xc0ZA0\xec\xb1\xb1\x17a\x03\x1d\xf9Q\x02" +
	"BG\x8c\xf2\xe5\x80\xb5\x8e`\x00d\xf3a\xc2z\xfa" +
	"r\x17\x01\xf7C\x8c\xf2\xb5\x80\xb5\x85`(d\xefS" +
	"\x106{\x907PZ\x80Q\xbe\x19\xb0v\x10\x0c|" +
	"l\xbcK\xd8\xc4^\xdeJ\xf9z\x15\xb0\xde\x05\xac\x83" +
	"\x04C'\x1b\xa6\x11\xf6\"\x85\xbc\x9b\xee\x08\x18\xe5\x7f" +
	"\x03\xacS\x80\x95h\x8c\xa9\x08\x1b\xba\xcb\x87\xc9J\xc0" +
	"\xfa\x1c\xb0\xce\x01\x96\xc5$\x86XyI\xd8\xe5\x89\xb4" +
	"0\xff6\xda\xf8\xbc\xd3G\x00K\x88\xb1\xde\x14]\x9f" +
	"\xf6\xb0\xa4':\xac\xd6\x16\xaed\x1d\x84K\x81t\x0c" +
	"d]tD\xec\xcf\x9a\x88\x88\xd8\xafM\xf4\xb9\xc9\xc0" +
	"\x0e\xa4\xdfK\x0520\xe6\xf8\x0b|\xbf6bm3" +
	"\xc2\xfaf\xe6`\xe7@\x08\xf4\xd5\x01\x18\xc0U\x86\xf5" +
	"\x88{\xb9\x8bq\x03\x86\xe8\x11T\xdf\x9d\x91X\xd9\xc6" +
	"\xcf\xdf\xae\xcd\x91\xb7#L6\x99@q4\x17\xbf\xf3" +
	"\xea\xc2\xb7\xa3\xa8\xbc\xd2GC\xb9\x8f\xd9\x13\xeb\xf5\xf4" +
	"\xb5\x7fa\xf8\xca\x16\xb9\x7f\xff\xeeb\xdd\xe6\x05Z\x91" +
	"CeP\xacI\x95g\xe1\xbeX\x17D\xcc\xb7\xd7\x01" +
	"\xb0\xd2d\xdcjj\x04+\x159\xeb\x0b\xf4+\xc7F" +
	"\xe4P\xbc\xdc\x98\xd5\x0eG\x02\xe1_\xddH\x9c\xcc\xbd" +
	"\xbf\x95\xe8\x0c\xb1\x89\x1e\xb6\x82\xe9\x0c\xcc\xe9\x9d-\x88" +
	"5J\xb3#\x85f\x10\xf6\xae\x19a$\xa46'\xd4" +
	"[.\xcc\x1bl\x1eL\xd8;@R=\xd6b\xb5\x98" +
	"-\xd8\xfb$\x84\xbd\x00%U`[z\"\xe6\x08\xf6" +
	"F\x12a\xef_I7=\x04\xcf\xc6\xd0\x1a\x8e2D" +
	"\x8b\x04\xf0%\xd6A\xc4\x86\x89\xf6\x93\x8a\x96\xb0\x1b/" +
	"\xd1\xa1\xb4\xefLtUD\x15\x93\xbdu\xccb\x19i" +
	"]\xff\x1a\x10S\x81\xa9X\xfd4\xd6\x890^ \xe8" +
	"Gc*fcG7\x9e\xfe\xf4Q\x8cYP_\xb6" +
	"\xee\xec\xd9\xd7\xfaY\x9e\xf5kL\xf0\xbf>(\xe7F" +
	"\xaa\xda\x1cn`3UK_\x8e\xda[\x13\xad\x8fQ" +
	"m\x8c\xc6PtO<.V\xe3\x8c\xef\x9a\x99=\xf4" +
	"\x16j3\xc7%\x85Bt\xf7\xb9hp\xb3\xb5\xa9]" +
	"2\xb9\x1c\xd2\xf4\xf7`Uxl7\xdct)\xa4w" +
	"\x1c\x17\xe3\x08l>\x80\x1f\x83\xd2P/v\x1f\x9d\x03" +
	"\xb0G\xf4\xb28\x8eh\x9d\x9b5P\xfb9\xd6j\xa3" +
	"\x16\xab\xc7\xebQ\x85\xf8\x90\xe2v\x9764\xa8\x01\x81" +
	"\x04\x84\xf8bl\x0eT4\x12\x11\x02\x8a\x887\xf3F" +
	"\xcd$\x85b\xcd(\xa3\x03\x8e_m\xf5\xceR{m" +
	"\x16\xf4g\xf4\x1c=\x11\x13{\x1c\xc4\xea\x96\xce\x10\x11" +
	"/\xec3\x13TBo\x0c)f(\xed\xb5\x89h!" +
	"7\x115\xb1\x89(w\x93Aq\xc1\xffR+\xde\x03" +
	"\xdc\x00\x9bm\")z\xf3\xab\x1du\xe0\x03\xd8\xfd\xbd" +
	"\\\x03\xfa\xaa\xf3c\x94\xf2\xda$=\xba\xe1\x1b+\x99" +
	"\xb0\xfb\xb3\xe1P\xdc\xc5-'\xfa\xe2\x06\xa7\xa8@\xe0" +
	"\x04-\x91\xd0\xbe\x1e\xc0\x1cE\xe1vt\x0f#\x00\xab" +
	"O\x09\xb6\xe8\x0cAUr\x95 \x14CL\xadQ\x9c" +
	"\xdd.\x1bzb\xa9V\xd3\xb5\x96\x1c\xf6\x8b-\xdd\x83" +
	"m\xec\x8a \xc6<\xa0P\x0fR%\\\x90\xbay\xb5" +
	"~\xb2{@\xec\xb0\xd2\xdf\xd9s\xe7h@Y\xb1\xf7" +
	"\xbbq_\xb3_\xc0l,\xef9\x0ct\x1fy\xb1\xfa" +
	"\xc3f\xd0\x9b\x8bFv?\xd0[\xc4\x1dwAY\xd8" +
	"\xa7Y/\x96\x9f\xc4\x1b\xbd\xd8%\xd8\x8b}\x0c\x80O" +
	"r\xbdX\xfeJ\xfc/\x1aj\xc6\xee\xa7\xf6\xfcrS" +
	"\x0f\xefN\xf4d\xea\xddFy\xa6\xa8\xc9\xa3@\x1b\"" +
	"\xe1\xd7\xd5\xc9\xe4P\xf8\xdd\x1fAo\x13\xb1\xf7\x95\x09" +
	"{\xb5Y\x92\x8a\xa0\xc4\xb0\x88\xc5\xda\xdc2V\x89\x10" +
	"\xc5\x10\x96\x14FY\xfc\x7f\xed%\x8b\x1e\xe6\xa3lL" +
	"\xdb\x0f\xc9\x18j\xfd\x9f\x00\x00\x00\xff\xff\xce\x14\xd4H"

func init() {
	schemas.Register(schema_c8d91463cfc4fb4a,
		0x85e320f14a5d23e0,
		0x87d94955ce3c61dd,
		0x8c519e0dedc17d73,
		0x8f2ef49549d64e86,
		0x9206caa8d3e3cc7e,
		0x9210d9e69d14fa35,
		0x9714437546d80c39,
		0x982790c08b1958ec,
		0x99efcebf23bbae35,
		0x9ad62de07dfc6419,
		0x9ad927034671cad1,
		0x9d159666de73f39d,
		0x9d4102fadb4f069c,
		0x9eb6708c01ec2079,
		0x9f6c36ef490dfd92,
		0x9fd40f92e1eb5d21,
		0xa22a2d1cf9579778,
		0xa2873a59df6d885c,
		0xa4f82f764dc3fee8,
		0xa535ac09456b2870,
		0xa53aedb3ce8994df,
		0xa5c3aa75d6b648e2,
		0xa75ecf12570b2966,
		0xa8f4ff97289294c7,
		0xa93eadc9671ea08b,
		0xadac227f85285c65,
		0xaffa789add8747b8,
		0xb1e3f6ac609eb4d7,
		0xb42ccfaaf45a3f7a,
		0xb469e5d523b89e1b,
		0xb4ecd69ac97e2de8,
		0xb5fcc0e153671d68,
		0xb70bd877cecb7b88,
		0xb8083dd65a24c770,
		0xb96fc5fb8137a705,
		0xb9d62f4beefefc29,
		0xba36a34b4eeb483f,
		0xbc193a4219598bcb,
		0xbc5e354741a8e665,
		0xbf3e401d5a63f336,
		0xc05520c9b0994ad3,
		0xc277e9822ae2c8fc,
		0xc41e71e8d893086c,
		0xcb3f7064eae4dc5a,
		0xd2654fcf2a7002cb,
		0xd271034eec62b43b,
		0xd29e9db5843719f0,
		0xd42684f756e09afd,
		0xd692a643ba8a1f58,
		0xda13a4f2919ce2cf,
		0xdbb4d798ea67e2e7,
		0xe6abbf843a84f35d,
		0xe96859cf77da6e6b,
		0xeb3c29aff080ec3e,
		0xec8866df56873858,
		0xecf1f14c4209c731,
		0xefea656d4b56b756,
		0xf0136e14d8019d3c,
		0xf020f2be35e8e2b5,
		0xf0931856093654c1,
		0xf12c60ebc67984d4,
		0xf144a5e58889dafb,
		0xf37f5e08534c68aa,
		0xf63241ee58b5166f,
		0xf63b8546288ee8e1,
		0xf6f911c4804ba7e5,
		0xf87a2c5a9f996828,
		0xf8fe6b4e94a960f7,
		0xf9d6c8c6d207c123,
		0xfabb5e621fa9a23f,
		0xfb3d38da0c9eaee6,
		0xfbbc20367c72bc59,
		0xfe7135f15d39bd5b)
}
