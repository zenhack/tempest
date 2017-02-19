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
func (c SandstormApi) GetIdentityId(ctx context.Context, params func(SandstormApi_getIdentityId_Params) error, opts ...capnp.CallOption) SandstormApi_getIdentityId_Results_Promise {
	if c.Client == nil {
		return SandstormApi_getIdentityId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      10,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "getIdentityId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_getIdentityId_Params{Struct: s}) }
	}
	return SandstormApi_getIdentityId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
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

	GetIdentityId(SandstormApi_getIdentityId) error
}

func SandstormApi_ServerToClient(s SandstormApi_Server) SandstormApi {
	c, _ := s.(server.Closer)
	return SandstormApi{Client: server.New(SandstormApi_Methods(nil, s), c)}
}

func SandstormApi_Methods(methods []server.Method, s SandstormApi_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 11)
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

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      10,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "getIdentityId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_getIdentityId{c, opts, SandstormApi_getIdentityId_Params{Struct: p}, SandstormApi_getIdentityId_Results{Struct: r}}
			return s.GetIdentityId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
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

// SandstormApi_getIdentityId holds the arguments for a server call to SandstormApi.getIdentityId.
type SandstormApi_getIdentityId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_getIdentityId_Params
	Results SandstormApi_getIdentityId_Results
}

type SandstormApi_deprecatedPublish_Params struct{ capnp.Struct }

// SandstormApi_deprecatedPublish_Params_TypeID is the unique identifier for the type SandstormApi_deprecatedPublish_Params.
const SandstormApi_deprecatedPublish_Params_TypeID = 0xa2873a59df6d885c

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

// SandstormApi_deprecatedPublish_Results_TypeID is the unique identifier for the type SandstormApi_deprecatedPublish_Results.
const SandstormApi_deprecatedPublish_Results_TypeID = 0xb42ccfaaf45a3f7a

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

// SandstormApi_deprecatedRegisterAction_Params_TypeID is the unique identifier for the type SandstormApi_deprecatedRegisterAction_Params.
const SandstormApi_deprecatedRegisterAction_Params_TypeID = 0xd271034eec62b43b

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

// SandstormApi_deprecatedRegisterAction_Results_TypeID is the unique identifier for the type SandstormApi_deprecatedRegisterAction_Results.
const SandstormApi_deprecatedRegisterAction_Results_TypeID = 0xb9d62f4beefefc29

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

// SandstormApi_shareCap_Params_TypeID is the unique identifier for the type SandstormApi_shareCap_Params.
const SandstormApi_shareCap_Params_TypeID = 0xeb3c29aff080ec3e

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

// SandstormApi_shareCap_Results_TypeID is the unique identifier for the type SandstormApi_shareCap_Results.
const SandstormApi_shareCap_Results_TypeID = 0xb96fc5fb8137a705

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

// SandstormApi_shareView_Params_TypeID is the unique identifier for the type SandstormApi_shareView_Params.
const SandstormApi_shareView_Params_TypeID = 0xb1e3f6ac609eb4d7

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

// SandstormApi_shareView_Results_TypeID is the unique identifier for the type SandstormApi_shareView_Results.
const SandstormApi_shareView_Results_TypeID = 0xe6abbf843a84f35d

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

// SandstormApi_restore_Params_TypeID is the unique identifier for the type SandstormApi_restore_Params.
const SandstormApi_restore_Params_TypeID = 0xd29e9db5843719f0

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

// SandstormApi_restore_Results_TypeID is the unique identifier for the type SandstormApi_restore_Results.
const SandstormApi_restore_Results_TypeID = 0xecf1f14c4209c731

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

// SandstormApi_drop_Params_TypeID is the unique identifier for the type SandstormApi_drop_Params.
const SandstormApi_drop_Params_TypeID = 0xadac227f85285c65

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

// SandstormApi_drop_Results_TypeID is the unique identifier for the type SandstormApi_drop_Results.
const SandstormApi_drop_Results_TypeID = 0xfbbc20367c72bc59

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

// SandstormApi_deleted_Params_TypeID is the unique identifier for the type SandstormApi_deleted_Params.
const SandstormApi_deleted_Params_TypeID = 0x87d94955ce3c61dd

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

// SandstormApi_deleted_Results_TypeID is the unique identifier for the type SandstormApi_deleted_Results.
const SandstormApi_deleted_Results_TypeID = 0xf8fe6b4e94a960f7

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

// SandstormApi_stayAwake_Params_TypeID is the unique identifier for the type SandstormApi_stayAwake_Params.
const SandstormApi_stayAwake_Params_TypeID = 0xb469e5d523b89e1b

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

// SandstormApi_stayAwake_Results_TypeID is the unique identifier for the type SandstormApi_stayAwake_Results.
const SandstormApi_stayAwake_Results_TypeID = 0x9fd40f92e1eb5d21

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

// SandstormApi_save_Params_TypeID is the unique identifier for the type SandstormApi_save_Params.
const SandstormApi_save_Params_TypeID = 0xd692a643ba8a1f58

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

// SandstormApi_save_Results_TypeID is the unique identifier for the type SandstormApi_save_Results.
const SandstormApi_save_Results_TypeID = 0x9206caa8d3e3cc7e

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

// SandstormApi_backgroundActivity_Params_TypeID is the unique identifier for the type SandstormApi_backgroundActivity_Params.
const SandstormApi_backgroundActivity_Params_TypeID = 0xec8866df56873858

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

// SandstormApi_backgroundActivity_Results_TypeID is the unique identifier for the type SandstormApi_backgroundActivity_Results.
const SandstormApi_backgroundActivity_Results_TypeID = 0xa535ac09456b2870

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

type SandstormApi_getIdentityId_Params struct{ capnp.Struct }

// SandstormApi_getIdentityId_Params_TypeID is the unique identifier for the type SandstormApi_getIdentityId_Params.
const SandstormApi_getIdentityId_Params_TypeID = 0xd76b6c6364d6bff5

func NewSandstormApi_getIdentityId_Params(s *capnp.Segment) (SandstormApi_getIdentityId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_getIdentityId_Params{st}, err
}

func NewRootSandstormApi_getIdentityId_Params(s *capnp.Segment) (SandstormApi_getIdentityId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_getIdentityId_Params{st}, err
}

func ReadRootSandstormApi_getIdentityId_Params(msg *capnp.Message) (SandstormApi_getIdentityId_Params, error) {
	root, err := msg.RootPtr()
	return SandstormApi_getIdentityId_Params{root.Struct()}, err
}

func (s SandstormApi_getIdentityId_Params) String() string {
	str, _ := text.Marshal(0xd76b6c6364d6bff5, s.Struct)
	return str
}

func (s SandstormApi_getIdentityId_Params) Identity() identity.Identity {
	p, _ := s.Struct.Ptr(0)
	return identity.Identity{Client: p.Interface().Client()}
}

func (s SandstormApi_getIdentityId_Params) HasIdentity() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_getIdentityId_Params) SetIdentity(v identity.Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormApi_getIdentityId_Params_List is a list of SandstormApi_getIdentityId_Params.
type SandstormApi_getIdentityId_Params_List struct{ capnp.List }

// NewSandstormApi_getIdentityId_Params creates a new list of SandstormApi_getIdentityId_Params.
func NewSandstormApi_getIdentityId_Params_List(s *capnp.Segment, sz int32) (SandstormApi_getIdentityId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_getIdentityId_Params_List{l}, err
}

func (s SandstormApi_getIdentityId_Params_List) At(i int) SandstormApi_getIdentityId_Params {
	return SandstormApi_getIdentityId_Params{s.List.Struct(i)}
}

func (s SandstormApi_getIdentityId_Params_List) Set(i int, v SandstormApi_getIdentityId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_getIdentityId_Params_Promise is a wrapper for a SandstormApi_getIdentityId_Params promised by a client call.
type SandstormApi_getIdentityId_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_getIdentityId_Params_Promise) Struct() (SandstormApi_getIdentityId_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_getIdentityId_Params{s}, err
}

func (p SandstormApi_getIdentityId_Params_Promise) Identity() identity.Identity {
	return identity.Identity{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormApi_getIdentityId_Results struct{ capnp.Struct }

// SandstormApi_getIdentityId_Results_TypeID is the unique identifier for the type SandstormApi_getIdentityId_Results.
const SandstormApi_getIdentityId_Results_TypeID = 0x8c4a70a31703d35c

func NewSandstormApi_getIdentityId_Results(s *capnp.Segment) (SandstormApi_getIdentityId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_getIdentityId_Results{st}, err
}

func NewRootSandstormApi_getIdentityId_Results(s *capnp.Segment) (SandstormApi_getIdentityId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormApi_getIdentityId_Results{st}, err
}

func ReadRootSandstormApi_getIdentityId_Results(msg *capnp.Message) (SandstormApi_getIdentityId_Results, error) {
	root, err := msg.RootPtr()
	return SandstormApi_getIdentityId_Results{root.Struct()}, err
}

func (s SandstormApi_getIdentityId_Results) String() string {
	str, _ := text.Marshal(0x8c4a70a31703d35c, s.Struct)
	return str
}

func (s SandstormApi_getIdentityId_Results) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s SandstormApi_getIdentityId_Results) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_getIdentityId_Results) SetId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_getIdentityId_Results_List is a list of SandstormApi_getIdentityId_Results.
type SandstormApi_getIdentityId_Results_List struct{ capnp.List }

// NewSandstormApi_getIdentityId_Results creates a new list of SandstormApi_getIdentityId_Results.
func NewSandstormApi_getIdentityId_Results_List(s *capnp.Segment, sz int32) (SandstormApi_getIdentityId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormApi_getIdentityId_Results_List{l}, err
}

func (s SandstormApi_getIdentityId_Results_List) At(i int) SandstormApi_getIdentityId_Results {
	return SandstormApi_getIdentityId_Results{s.List.Struct(i)}
}

func (s SandstormApi_getIdentityId_Results_List) Set(i int, v SandstormApi_getIdentityId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_getIdentityId_Results_Promise is a wrapper for a SandstormApi_getIdentityId_Results promised by a client call.
type SandstormApi_getIdentityId_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_getIdentityId_Results_Promise) Struct() (SandstormApi_getIdentityId_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_getIdentityId_Results{s}, err
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

// UiView_ViewInfo_TypeID is the unique identifier for the type UiView_ViewInfo.
const UiView_ViewInfo_TypeID = 0xbc5e354741a8e665

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

func (s UiView_ViewInfo) GrainIcon() util.StaticAsset {
	p, _ := s.Struct.Ptr(6)
	return util.StaticAsset{Client: p.Interface().Client()}
}

func (s UiView_ViewInfo) HasGrainIcon() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetGrainIcon(v util.StaticAsset) error {
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

func (p UiView_ViewInfo_Promise) GrainIcon() util.StaticAsset {
	return util.StaticAsset{Client: p.Pipeline.GetPipeline(6).Client()}
}

type UiView_PowerboxTag struct{ capnp.Struct }

// UiView_PowerboxTag_TypeID is the unique identifier for the type UiView_PowerboxTag.
const UiView_PowerboxTag_TypeID = 0x982790c08b1958ec

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

// UiView_getViewInfo_Params_TypeID is the unique identifier for the type UiView_getViewInfo_Params.
const UiView_getViewInfo_Params_TypeID = 0x8f2ef49549d64e86

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

// UiView_newSession_Params_TypeID is the unique identifier for the type UiView_newSession_Params.
const UiView_newSession_Params_TypeID = 0xf87a2c5a9f996828

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

// UiView_newSession_Results_TypeID is the unique identifier for the type UiView_newSession_Results.
const UiView_newSession_Results_TypeID = 0xa8f4ff97289294c7

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

// UiView_newRequestSession_Params_TypeID is the unique identifier for the type UiView_newRequestSession_Params.
const UiView_newRequestSession_Params_TypeID = 0xbc193a4219598bcb

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

// UiView_newRequestSession_Results_TypeID is the unique identifier for the type UiView_newRequestSession_Results.
const UiView_newRequestSession_Results_TypeID = 0xa22a2d1cf9579778

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

// UiView_newOfferSession_Params_TypeID is the unique identifier for the type UiView_newOfferSession_Params.
const UiView_newOfferSession_Params_TypeID = 0xa53aedb3ce8994df

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

// UiView_newOfferSession_Results_TypeID is the unique identifier for the type UiView_newOfferSession_Results.
const UiView_newOfferSession_Results_TypeID = 0x9eb6708c01ec2079

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

// SessionContext_getSharedPermissions_Params_TypeID is the unique identifier for the type SessionContext_getSharedPermissions_Params.
const SessionContext_getSharedPermissions_Params_TypeID = 0xe96859cf77da6e6b

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

// SessionContext_getSharedPermissions_Results_TypeID is the unique identifier for the type SessionContext_getSharedPermissions_Results.
const SessionContext_getSharedPermissions_Results_TypeID = 0xb70bd877cecb7b88

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

// SessionContext_tieToUser_Params_TypeID is the unique identifier for the type SessionContext_tieToUser_Params.
const SessionContext_tieToUser_Params_TypeID = 0xc41e71e8d893086c

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

// SessionContext_tieToUser_Results_TypeID is the unique identifier for the type SessionContext_tieToUser_Results.
const SessionContext_tieToUser_Results_TypeID = 0xf6f911c4804ba7e5

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

// SessionContext_offer_Params_TypeID is the unique identifier for the type SessionContext_offer_Params.
const SessionContext_offer_Params_TypeID = 0xfb3d38da0c9eaee6

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

// SessionContext_offer_Results_TypeID is the unique identifier for the type SessionContext_offer_Results.
const SessionContext_offer_Results_TypeID = 0xfe7135f15d39bd5b

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

// SessionContext_request_Params_TypeID is the unique identifier for the type SessionContext_request_Params.
const SessionContext_request_Params_TypeID = 0xf63b8546288ee8e1

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

// SessionContext_request_Results_TypeID is the unique identifier for the type SessionContext_request_Results.
const SessionContext_request_Results_TypeID = 0xd42684f756e09afd

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

// SessionContext_fulfillRequest_Params_TypeID is the unique identifier for the type SessionContext_fulfillRequest_Params.
const SessionContext_fulfillRequest_Params_TypeID = 0x9f6c36ef490dfd92

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

// SessionContext_fulfillRequest_Results_TypeID is the unique identifier for the type SessionContext_fulfillRequest_Results.
const SessionContext_fulfillRequest_Results_TypeID = 0xb4ecd69ac97e2de8

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

// SessionContext_close_Params_TypeID is the unique identifier for the type SessionContext_close_Params.
const SessionContext_close_Params_TypeID = 0xf12c60ebc67984d4

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

// SessionContext_close_Results_TypeID is the unique identifier for the type SessionContext_close_Results.
const SessionContext_close_Results_TypeID = 0x9d4102fadb4f069c

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

// SessionContext_openView_Params_TypeID is the unique identifier for the type SessionContext_openView_Params.
const SessionContext_openView_Params_TypeID = 0xf37f5e08534c68aa

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

// SessionContext_openView_Results_TypeID is the unique identifier for the type SessionContext_openView_Results.
const SessionContext_openView_Results_TypeID = 0xf9d6c8c6d207c123

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

// SessionContext_claimRequest_Params_TypeID is the unique identifier for the type SessionContext_claimRequest_Params.
const SessionContext_claimRequest_Params_TypeID = 0xda13a4f2919ce2cf

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

// SessionContext_claimRequest_Results_TypeID is the unique identifier for the type SessionContext_claimRequest_Results.
const SessionContext_claimRequest_Results_TypeID = 0xefea656d4b56b756

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

// SessionContext_activity_Params_TypeID is the unique identifier for the type SessionContext_activity_Params.
const SessionContext_activity_Params_TypeID = 0x85e320f14a5d23e0

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

// SessionContext_activity_Results_TypeID is the unique identifier for the type SessionContext_activity_Results.
const SessionContext_activity_Results_TypeID = 0xa93eadc9671ea08b

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

// PermissionDef_TypeID is the unique identifier for the type PermissionDef.
const PermissionDef_TypeID = 0xf144a5e58889dafb

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

// RoleDef_TypeID is the unique identifier for the type RoleDef.
const RoleDef_TypeID = 0xcb3f7064eae4dc5a

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

// SharingLink_getPetname_Params_TypeID is the unique identifier for the type SharingLink_getPetname_Params.
const SharingLink_getPetname_Params_TypeID = 0xf0931856093654c1

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

// SharingLink_getPetname_Results_TypeID is the unique identifier for the type SharingLink_getPetname_Results.
const SharingLink_getPetname_Results_TypeID = 0x9ad927034671cad1

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

// ViewSharingLink_RoleAssignment_TypeID is the unique identifier for the type ViewSharingLink_RoleAssignment.
const ViewSharingLink_RoleAssignment_TypeID = 0xf020f2be35e8e2b5

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

// ViewSharingLink_getRoleAssignment_Params_TypeID is the unique identifier for the type ViewSharingLink_getRoleAssignment_Params.
const ViewSharingLink_getRoleAssignment_Params_TypeID = 0xb8083dd65a24c770

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

// ViewSharingLink_getRoleAssignment_Results_TypeID is the unique identifier for the type ViewSharingLink_getRoleAssignment_Results.
const ViewSharingLink_getRoleAssignment_Results_TypeID = 0x9d159666de73f39d

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

// GrainInfo_TypeID is the unique identifier for the type GrainInfo.
const GrainInfo_TypeID = 0xb5fcc0e153671d68

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

// AppPersistent_save_Params_TypeID is the unique identifier for the type AppPersistent_save_Params.
const AppPersistent_save_Params_TypeID = 0xf0136e14d8019d3c

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

// AppPersistent_save_Results_TypeID is the unique identifier for the type AppPersistent_save_Results.
const AppPersistent_save_Results_TypeID = 0xba36a34b4eeb483f

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

// MainView_restore_Params_TypeID is the unique identifier for the type MainView_restore_Params.
const MainView_restore_Params_TypeID = 0x9ad62de07dfc6419

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

// MainView_restore_Results_TypeID is the unique identifier for the type MainView_restore_Results.
const MainView_restore_Results_TypeID = 0x99efcebf23bbae35

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

// MainView_drop_Params_TypeID is the unique identifier for the type MainView_drop_Params.
const MainView_drop_Params_TypeID = 0x8c519e0dedc17d73

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

// MainView_drop_Results_TypeID is the unique identifier for the type MainView_drop_Results.
const MainView_drop_Results_TypeID = 0x9210d9e69d14fa35

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

const schema_c8d91463cfc4fb4a = "x\xda\xdc:kxT\xd5\xb5k\x9d3\xc9\x10/i" +
	"r<\x09\x84\x10\x0c\x86xMB\x12^\xa2B\xd1\xbc" +
	"\xa4\x98\x88\x9a\x93\x04\x0a\xb9\xe2\xc7I\xe6$\x19\x99\xcc" +
	"Lf&\x84P)\xe2\x85+T\xbc\x95\x14o\x81O" +
	"\xb0Es\x05\xaf\xd4j\xa1\x82\x95\x8a^-\x0f\x1fT" +
	"\x04\x04-U\xf0QP\xa8\xa2\xe2\x1b\xce\xfd\xd6\x9e\xb3" +
	"\xcf\xecL&/\xbe\xfb\xf5G\x7f\xf01\xd9g\xed\xb5" +
	"\xd7^\xef\xc7\x1ek$\x15K\xe3\xe2\xbe\xca\x05\xa8>" +
	"\x89q\xf1\xe6\xbb\xa3\xe6T\x9c\x1d\xf9\xde2P\x86#" +
	"@\x1c:\x01&\x1cO\xa9E@\xf5LJ\x11\xa0y" +
	"L\x9f\xf2\xda\x8c\xf2\xa3\xf7\x802,\x0c\x90\x8c\x13\x12" +
	"SK\x09 5\x95\x00n{C\x1e\xfa\xb0\xbfb\xa5" +
	"\x0801u!\x01\x940\x80\xe0\xa2\xe7\xcf$n\xd0" +
	"V\x82\x92j\x03\xe8\xa9\x97\x10\x80\x9b\x01\xfc\xc7-\x87" +
	"\xcb\x1f\xf8\xa2\xf0\xe7\xa0\xa8\x08\xe0 \x12\xeeK\x1d\x8f" +
	"\xe00\x7f\xfa\xca{ol\xda\x17\xdf!\xe2n\xa5O" +
	"\xa8.b['~\x9b\xb2\xfe\xc3\xa3\xc9\x1da\xdc\x0e" +
	"\xfa\xbe>\xf5R\xdazz\xd6\xb0{w\xdd\x7f\xe5\x1a" +
	"Bj\xfe\xedD\xe3Gk\xde\xdc\xfav\xf8\x82\xea\xb2" +
	"\xd4o\xd5U\xa9\xf4\xeb\xbe0\x96\xdf\xfca\xd4s\xaf" +
	"}\xb2V\xa4pKj\x1e\x1d\xb3\x8d\x01\x0cs}\xbf" +
	"\xe8\xdd\x82\xc3\xebD\x80C\xa9Y\x04p\x8c\x01\xbc\xbe" +
	"\xaf\xe5G\xf2\x95G\xd7\x81\x92f\xb3\x11\x8706&" +
	"\x0e!\x80\xf5\x9f\x07\xff\xda\xf0_\xa9\xebA\xc9\xb0\x01" +
	"\xe6\x0cy\x8a\x00\x9a\x19\xc0\x83\xf1\xb7\xbe\xfd\xadT\xb2" +
	">,\x08\xc6\x84\x9dC*\xe8&\xed#O\xe3J\xff" +
	"\xef7\x84\xd9\x13\xde\xba9\x8c{\x1b\xdb\xdaq>\xb1" +
	"\xfc\x93\xab=\x0fY2d{\x0f\x0d\xf9\x19\x01|0" +
	"\xe4o\x80\xe6\xe5s>>\xde\x91t\xf0!\x91\x8dG" +
	"\x872\x0c\x1f\x0c%\x0c\x0b~\xf9\xe3o2\x0a\xf26" +
	"\x8aG\xc4\xa5\xddA\x00J\x1a\x13\xf2\xf2\xe6wfO" +
	"\xbegc\x18\x03\xe3\xf3\xb8\xb4\x0e\xa2\xee\xe4\x85\xff\xbd" +
	"y\xfe\x98\xaf\x1f\x01e\xb0lV|\xf7\xe2\xfe\xfa\x94" +
	"\xa3{\x00P\xbd<\xed\x84Z\x906\x14@\x9d\x946" +
	"Mu\xd3/\xd3\x9f3oj\xc2\xe3\x13;\x054\xb3" +
	"\xd36\x12\x9awV\xafx\xedwg&w\x82\xa6\"" +
	"\x91\x10O$\x94\xa7U!\xe0\x84\x19i?G@s" +
	"\xf7\xea\x8e\x9c_\x9a_l\x12\x89<3\x8c)\xc3\xb9" +
	"aD\xe4\xbd\xbf\xba\xacq\xef\x96\xeb7\x0b,LM" +
	"\xaf#\xec\xc6m9\xcb\x16g=\xbeEd\x00\xa63" +
	"\x01'\xa4\xd3\xd6\xed\xd3\xee9\xb6n\xc1\xb7OD\xdd" +
	"\x02\x92Q\xcdM?\xa0NL'e\x19\x97>M]" +
	"\x92\xeeT\x97\xa4'\x99on\xdd0\xf7\xf1/\xdf{" +
	"R\xc4\xd8\x9eN\xe4\xaa\xcb\x18\xc6\x85E\xb5_<\xb6" +
	"?\x7f\xabp\xd5\xce\xf4uD\xcc\xf0\x0d\xdbG\x1d\xfa" +
	"\xc0m}\x89\x93\xe8\xd3\x03\xe1\xad\xbfNo\x034O" +
	"\x16\xfct\xef\xba\xc3\xa7\xb7\x0a\xf7\x88\x1b\xce\x98\xdd4" +
	"\xa2\xb1\xfa\xf8\xae\xef\xb7\x816\x181Bg\x9cD\xe4" +
	"\x9dK\x7fJ=\xcf\x08\xfd&\xfd\x09@s\xf9O^" +
	"~\xad\xed\xc8\xbf<-\x9a\xf6\x03\xc3_\xa0s:\x87" +
	"\x13\x89\xfe\xdd\xd9\xb5\x87\xaf\x1b\xb4=\xac\x93\xe1s2" +
	"\x1e\xa3s\xe2\x1e\xbdf\xc9w/\xf9v\x88$\x9e\x1b" +
	"\xceH\xc4\x0c\"1\xf7\xfb\x0b\x7f\xbfi\xcc\xe1\x1d\xc2" +
	"\xed\xf4\x8c\x03\xb4\xb5\xe8\xc6\x8fo\xb9\xe9\xe1\xab\x9f\x01" +
	"%\xdd\xde:#c2m\x9d\xc3\xb6\xbe|\xef\xeca" +
	"\xa5\x93\x87=\xcb%\x1d\xc7T=\xa3\x8e$\xbd7\xe3" +
	"\xc7$i\xe3\xc3M%\xd3&\xde\xfel\x94\xed\x0e\xa2" +
	"\xbbM\xbc\xec\x84Zr\x19\xed\xb9\xee\xb2\xdd\x04|\xf5" +
	"\xe7\xf5\xb5#\x8a\xaf\x7f\xae\x9b\x02.\x19\xf9\x96z\xdf" +
	"H\x82\\1r\x9a\xa4&f9\x01\xcc7*\xd6\xfe" +
	"v\xef\xc8\x19\xbb\xbaA\x9f\xbb\xfc\x05\xf5\xfc\xe5\x8c{" +
	"\x97OS\xaf`\xc0\xdf\xef9\x91w\xf7\xa9\xb6\x17b" +
	"hEb\xd6cj*\x01\xa9J\xd6=\xea\x93YC" +
	"\xd5\x1dYI\xa6g\xd0/\x8e\x9cl\xb9\xecE\x8b\xe5" +
	"2\x1d\xfed\x16]M\xdd\x99E2\xa9\xfd\xcb\xfb\x1f" +
	"\xb9\xfcE/G\x0b\x90\x98\xaf\xb6\x8c\xda\xa8\xb6\x8f\xa2" +
	"=\xad\xa3\x18\x1f^\x96\xfcy\xfbo5\x0e\xc48\xff" +
	"x\xf6>\xf5L6\xc1\x9e\xca\xde-\xa9\xcbr\x9c\xea" +
	"\xb2\x9c$\xf3\x87[\xebN\xdf\"\xb7\x1c\x10\xe4\xb2(" +
	"g\x1f\xc9\xe5\xd3a\xd7,\xdd\xb6~\xc3\x01Qa[" +
	"r\x98\x1fo\xcf!m8\xbf\xee\xdd\x99_-\xfd\xd7" +
	"\x83\x16\xed\xa4R\x13\xd6\xe60'\xd1\x99C\x82\x9b\x95" +
	"\xf9\xb3g\xca\xfe\xbb\xe3\xb0\xa8\x14q\xb9\xcc\x88\x12s" +
	"\x09\xe0\xdcs\x87]\xf5\x9eyo\x8aG\xb8s\x03\x04" +
	"\xd0\x9aKG\xec?\xf1\xe0\xaa\xcf\x1eQ\xdf\xearD" +
	".\x0b\x15\x9d\x0c\x83-\xebh\xe9\x8c\xc8[\xa7^\x91" +
	"w%\xc0\x84\x92<'\xaa\xe7\xf2H<s>_:" +
	"y\xe9s\xff\xf3\xa1H\xd0\xb1<F\xf1\xa9<B7" +
	"\xcf\xfbV\xdb\xfe\xd9M\xa7\x04C*\x1f\xfd\x0cq\xe3" +
	"\xfa\xd3w}\xfaD\xee\x94\x8f\xc5\xad\x93FW\xd0\xd6" +
	"\xa9\xa3\xd9e\xaf\xbdg\xe6;\x0d\xcbO\x8bw\xe9\x1c" +
	"\xbd\x8e9\xdd\xd1t\x97q\xbb\x13J\xa7\x9f=\xdb\x05" +
	"\xe0P\x18\xc3q\x060\xf3\xe9\x9975\x1b\x1f}\"" +
	"\x9a\x1f\xe6\xdf\xcd\xd8\x95O\x00S\xd6\xe3\x91\x14\xaf\xfa" +
	"i\xd8R\x98\xac\x0a\xf2Y\xd8\xdbv\xe2\xe4\xc4?~" +
	"6\xf2S\xd02\x10#\x1ev\xaa\xe4\x94\xc9\xa7\xe5\xd7" +
	"\xa2\x9a\x9bO\xf8\xae\xc8\xcf$=y\xbe\xe6\xea\x84\x99" +
	"i\xbf\xf84\x1c}\xd8=\x8d\x82*\xc2tpi\xfb" +
	"\x9f>\x9e\x9b\x7fV\xe0\x80VPJ_\xbe{k\xc5" +
	"\xf2\x0f:o8\x1b\xad\x89\xa4\xb3\xea\xa4\x82\x03\xea\xd4" +
	"\x02\xfaUR@\x01\xe4\xb1\xa6\xe9\xd5\x83n_\xfc9" +
	"h\xc3\xd1\x96\xdc\x88B\xc6\xea\xdcBRl\xdf\x90m" +
	"\xb3\xfe^2\xfe\xcbn\x92{\xa9\xf0)\xf5\xd5BB" +
	"\xb5\xb7\xd0I\xff\x00\xcc\xe3'\xff3\xe7G\xcb~\xf8" +
	"\xa5\xa8\x07O\x162\xf7\xb2\xb3\x90\xb8\xff\xc1\xa37\xdd" +
	"\xf5\xa2\xf2\xcd\x97\"\xefJ\xc6\xb0xt\xf3\x18\xe2]" +
	"N\xd3\xda\x87j\xf3\x17~\xcd\x9d\x08\xbbY\xf3\x18R" +
	"\xc6\x09\xadc\x18S\xbe\x9a\xbby\xf5-\xf3.|-" +
	"\x98\xc2\xaa\xb1,\xa0\x8ez\xdey\xe0O{\x0e\x7f#" +
	"0e\xd1X\x16'>\xfc\xcd\x86\xc1o]{\xddw" +
	"b$u\x8feF\xd22\x96\x181\xfb\xd9\xc0\x9dW" +
	"\x8f|\xf6;\x01\xa91\x8e\xc9\xec\xdfvN\x9asv" +
	"b\xcb\x05\x01\xe9\xcd\xe3*\x10L0a\xa2\xd9\x18\xd0" +
	"\xdd\xde\xc2z\xdd\xe1\xf7\xfa'W\x1b\xc1\xa0\xdb\xe7-" +
	"\xf3yC\xc6\x82P\xa1^\x1fr\xcfw\x87\xda\xb3+" +
	"\xf5\x80.7\x075\x87\xec\x00p \x80\x928\x1e@" +
	"\x1b$\xa3\x96\"a\xa61\xdf\xf0\x860\xd9\xbc\xc9\xb3" +
	"r\xf3G\xbf\xdf\xf5(\x00b2`\x14n\xdd\xeb\x0a" +
	"\x86|\x81\xe6\x12\xbf\xbb\xd0ex\x8c\x90\xe1b\x88\x9b" +
	"\x83 \"\xce\x8a v\x06\x8c\x06\xbc\x14\x05w\x03\x80" +
	"\x97\xf6\x8a\xb8\xd1\x08\x95\xbb\x0co\xc8\x1dj/we" +
	"W\x15\x19\xc1VO\xa8\x0b\xe1\xe9\x11\xfc\xb2\xdb\x85\x89" +
	" a\xa2\x80R&\x947\xebn\xefL\xb7\xd1V\xe8" +
	"\x0a\xf8\xfc\x9c\xc8.TV\x00h\x83e\xd4\xd2$4" +
	"}uw\x18\xf5\xa1r\x17\xa3\x0e1\xe2\x9c\xa3\xa8e" +
	"\xa8g\xb8\x19\xe2F#D\xff\x97{\x1b|\xd9\x95\x99" +
	"\xec\x80\xaep]n\x15\xd4\xe7\x1b\xd9UFf\xb7\xcb" +
	"\x88R\x08\xf9\xe6\x19\xde\xfe\xdc\xa7\x8a1\x05#\x07J" +
	"\x02a\x95\xbe6#PT\xe7[P\xa37V\"\xf6" +
	"x\x98;\xe41p0H8\xb8\xc7\xc3\x02\x06]\x80" +
	"\x08\x0f&E\x13.J\xb9^\xf7\xe3\xa5\x0e\x19\xb0;" +
	"\xb7\xba\xe1\xaa\xd4\x03N\xbd\xab&^\x8c(\xc2\x8a\xd3" +
	"\xa4\x07\xdc\xde\xc6\xe9n\xef<\x92G\xa5\x11\xf2\xea\xcd" +
	"F\x98=rWr\xf3,ro\x900\x89\x80P1" +
	"\xd7\xae\xcaZ\xb3u\xc6g\x1f\x01@1*\x98\xa99" +
	"$\x14\x17\x15\x1c\xaa9\x10\x11+e\xc4d\xb3\xfc\xc4" +
	"O\xae\x09<=\xe7^\xb2\x0dE\xa0$\x8e(\xa1+" +
	"FQS\xe5\xf3\x18%\xc1\xa0\xbb\xd1\xdblxC1" +
	"\x85\x9f\x15!\xca9_\x0f\x0c\x9c&\xdb\x9bG\xd1\x14" +
	"\xcb\x17\xd4{|A\x8b7\xa1 @WXKw\xbc" +
	"F\xdb\xad\x0d\x0dF\xc0\xda\x1a\x93\x93\xa5\x11\xc1/\x0e" +
	"\x86\xe1P\x898\xe9X\xdc\x89\xa2\xa4\xa1\xd5\xd3\xe0\xf6" +
	"x\xaa\x8c\x96V#\x18\x8aXg\xb2}\x88N\x9c\xb9" +
	"MF\xadIB\x051\x85\x1c\xb1bl\x04\xd0\x9ad" +
	"\xd4\x96J\xa8HR\x0aJ\x00\xca\x92Z\x00\xed.\x19" +
	"\xb5\x95\x12*\xb2\x9c\x822\x80\xb2\xa2\x0e@[.\xa3" +
	"\xb6\xba\x9br\x06\x8c\x96Vw\xc0pa\xa5\x11hv" +
	"\x07\x83n\xa7\xcf\x1b\xc4\x1f\x00c(\x82D?M\x97" +
	"\x11\xac\x0f\xb8\xfd!\x90}\x01L6ge\xec~\xee" +
	"\xe1\xa4\xfdg\xb8Wt\xb9\x83~\x8f\xde^\x0eNo" +
	"\x83\x0f\x93\xcd\xd3\xa3\xdf\xfbc\xc2\xc6\xea\xce\xbe\xbdf" +
	"0\xa4\xb7\x97\xb4\xe9\xf3b\xeb\xe8\xe4\x08g\x8b\x9at" +
	"\xaf\xcbCZ\xfan\xe9\xdc\xb9\x8fg\x7f\xb1&\xa6\x88" +
	"#b\xb3\x98)\x08\xae\x9b\xc5^\x8c\xe0\xba\xba|\x7f" +
	"\xc0\xa8\xd7C\x86\xab\xb2\xb5\xce\xe3\x0e6\x85%\x17\xed" +
	"\x86DSp\xba\xbd\xf3H_#\x8a\xaa\xe0B\x93[" +
	"\x06\x14\x85mCs\xc8qT\xc1Y5\x03\xf2\x82V" +
	"Q:@R\x12\x9d&\xb7&\xe4\xe6\x84\xa1bdx" +
	"\xed\x9c[P\xe8\xee\x94\xd7\xe9\xf5\xf3\x1a\x03\xbeV\xaf" +
	"\xab\x84\x07D\xe2\xbf\xd3\x13\x0a\xf6\xcb\x0c\xac{j\x19" +
	"63\xb7\x91\xcf\xda*\xa3\xb6KP\xd0\x9d\xc4\xe1\xed" +
	"2j/J\x88\x96~>O\xaa\xb8KF\xed\x15\xd2" +
	"O)\xac\x9f{\x03\x00\xda\x1e\x19\xb5\x83\x12*\x0e9" +
	"\x05\x1d\x00\xca\xeb\xe4\x9d_\x91Q;\"\xa1\x12\xe7H" +
	"\xc18\x00\xe5\x10\xa9\xf7A\x19\xb5w%T\xe2\xe3R" +
	"0\x1e@9F\x90Gd\xd4\xde\x97\xd0l\x0d\x1a\x01" +
	"\x0aB\xe4 \x93\xcd7\x949\xbf\xfb\xe4\xf5\x1d\xab-" +
	"U\\\\\x1f69T\"\xb5\x8c%gK\x03j\xc0" +
	"\xd9\xee70\x01$L\x88\xacVB8\xa2\xe1\xa5 " +
	"\x91\xd9d\xfa\x88\x1d\x82\x19\xf5n!\x99!\xbd\xae\xbc" +
	"\x87\xc8\x1caqDW\xbb{\xc6\x81\xe9j\xaf\xa9O" +
	"7Iw\x8f\xceV\x92\x90\x14\xd0\x9b\x07\x18\x9c\x99\xc6" +
	"\x97\xf8\xfd\x95F \xe8\x0e\x86H3\xc31\x97\x14\x9a" +
	"g\xe1\xc8\x0bWE\xc9\x03(\x19\x8c%\x19\xa8\x148" +
	"\x93()\xe0>\xde\xee\x120M\x8a\xb5X\x8c\x95\x88" +
	"%\x0eT\xb0\xce,\xf1\xfbo\xa5@\x09\xceP\xb9\xab" +
	"7\x7f\xd3\xa4\x07\x0c\xe2\xb7\xad\xc21Bb\x8a\x84I" +
	"\xf3\xddF\x1b*\x91\"\xe9\"\x9c\x01\xf7h\xfdr\x7f" +
	"\x9c\x9cA69\xb9d(92jW\x09\x165\xee" +
	"\x0e\x00m\xac\x8c\xda\x14\xa9\xbb\xdb\xfd\xf7\xaf\x0e\x96\xee" +
	"\xfbl\xc7\x1b\xdc\xedz}!w\x83\xbb^\x87\xa4P" +
	"X]^-^\xb1\xfd\xed\x11\xcb.\\DL\xea\x96" +
	"a!\xed\x99\x16\xd0\xdd\x99^\xb26\x12\xf2`\x9b\xf6" +
	"\xa9\xa4(\xc52j\xd3%\xe4\xa4\x97\x93\xe5\xde(\xa3" +
	"VC\xd1\x0a\xc3\xde@#\xc0\xe92j\xb3$\xcc\xd4" +
	"\xfd\xfer\x97\x9d\x81\xe9~\xffLR\"\x90}^\x1c" +
	"\x04\x12\x0e\x82\x9e\x92\xb4\xf8\x18\xf47\x1a!r\xba\x86" +
	"\x8b\xc75\x9f7h\x07\xfbX9G\x8d\x9dsd\x97" +
	"\x1e\x9e[\xfd\xe5gw\xf5+\xe7\x10\"\xe5E$A" +
	"\x96\x8d\xf5\xa5\xb1ez$\xc7\x15U\xa4*\x96\x8a\x90" +
	"\x1a\xe7\xcb\xa8]+\xa1\xc9v\xbb\xcat@!\xe6'" +
	"y\xdc\xdey\xa8\x08\xb1\xa2\xab:\xc4\xf7\xa2\xdcUF" +
	"#\x99u\x80\xe2\x06\xcf\x87D\xb5\x90\xa3\xed?d\xa5" +
	"\xfa\xbcn\x11\xa8\xaf\x88E\xfdx\x81\xfa\xe8\xdc\xd76" +
	"\xffp\xee\x9b\xe9\xd1\xeb\x0cOT&\x9a\xdc\xef\x94\x80" +
	"'\xdei6Ek\x89\xa252j\x8f\x08\x14\xfd\x9a" +
	"\\\xef\x832j\x9b\"A\xac\x93l\xf3\x11\x19\xb5\xdf" +
	"\x0aAl\x0b\x05\xb1\xc7e\xd4\xb6\x0bAl[\x9d\x15" +
	"\x17\xff,\x04\xb1W\xc7G\xc2\xdd?&^\xb14\xcf" +
	"\x08\x86,_a)m\xd78\xf5\x83\x1e\xe3\x94XM" +
	"\xb1\x1a\xcfiY\xfcH\x9bu\xaf\xd3=\xff,\xa3v" +
	"Z`\xdd)\xba\xe7\xfb2V\xa1\x90\xa0\x9e\xef\x00\xa8" +
	"B\x19\xab30\x92\xa1\xaa\xc30\x00P\x9dF\xebW" +
	"\xd1\xba\xc3\xc1\xf8\xa7\x8e\xc3:\x80\xea\xb1\xb4>\x9d\xd6" +
	"\xe3\xe2\x18\x0b\xd5r\xac\x00\xa8\xbe\x91\xd6kh=>" +
	"\x9e\xa5\x02\xaa\x86U\x00\xd5\x95\xb4~\x1b\xad;\x9d)" +
	"l\x0c1\x1bk\x01\xaag\xd1z\x08%4\xfd\x96W" +
	"\x001\xddM\x8etj,~\x04|\x1eC\xf8l\xb7" +
	"\x14\xc3\x9fM\x97\xe1u\x93\x87A\xeeb\xb0[\xea\xdc" +
	"\xac\x87\xea\x9b\xaa\x8c\x16\xc8$\xfe\x07{d}\x18\xf0" +
	"\xd6\x86\x06p\x1a\x81^\xc0t\xbf\xbf\x86\x9c`X]" +
	"bj~y\xbd\x0f\x90\xfc}\xd1\xc6\xcd\x99u\xb7\xff" +
	"\xe1[\xae2\xac\xabQ\xd3\xee\x07Y\xbc\xd4\xe8u{" +
	"\xde\xbc\xe4\xaf\xd7~\xc8O\xe8\"u\xd1\xab\xca\x0bX" +
	",\xcff\xb1\x9c\xf7\xfb\x90\xb7\xbeU\x05\x1f\x03IM" +
	"D'\xa2\xdd\x9cE\xde]R\x11\xab@R\xbeq\xa2" +
	"dw~\x90\xf7q\x943\xe3AR\x8e;Q\xb6\xdb" +
	"U\xc8[\xa4\xca\xa1R\x90\x94\xbdN\x8c\xb3g/\xc8" +
	"\x9b\xf6\xca\xce\x85 )\xdb\x9c\x18o7\xdf\x90\xcfv" +
	"\x94\xcd\x84s\xbd\x13\x9dvG\x0dy3JYU\x01" +
	"\x92\xb2\xc2\x89\x0e\xbbM\x8a\xbc\x85\xa8,\xba\x03$\xa5" +
	"\xd5\x89\x83\xecy\x1d\xf2a\x87\xe2\xa6}:\xcb\xbfY" +
	"pA\x1e]\x92(\xbc\x14\xa3\x19r\x1b5\xbe\x19A" +
	"\x030Pl%\x89\xc5\xb8\xd82\xbeb4yT\x85" +
	"\"\xbe\x94\xc9\xca\xd0b4}~\x835\x06X\xd41" +
	"\xeb=\xba\xbb\x99\xd4&\xc9\xda\xc9\xd37\xb0\x12\x9f(" +
	"1E\xa2\x0cD\xf2-\xde\xabD>2S\x94Z\x90" +
	"\x94\x04v\x01\xd6\x1c\x00\xb9\xd9\xe8\x8a\x0ey\x93\"\x89" +
	"\x88!\\\x83\x18.>\x96C>\xc0S\xc6\x95\x02\x94" +
	"\xe4c\xc9\x14Tn&\x99\xf3\xd9#\xf2A\xa1RB" +
	"\xd9\xdd\x14,\x99\x8e\xca\x1c\xe7b\xab\xe1\xc1\x03\xaa\xdd" +
	"\xc8\x10\x13\xbc.\x8b\xc5\x98D\x89\xe8@6\xb0\x1a(" +
	"\x92\xb4A?\xd2\xc3\xaei\x03\x17`\xc0\x0e\x0fBV" +
	"\x93\x15\xc9jl\x1fW\xbe\xd1\xca`\x9a\x04\x1fg\x90" +
	"7t\xc9\xa8\xf9/\xb2\xde\xeegA\xcd\x84E)\x85" +
	"\xf3\x06\xa3\x81d%\xc42\xf2\xbd\xabe\xd4~%\x10" +
	"\xbb\xbe6\x12\xcblb\xbb\x063\xabc\xb0\xa5\xce\x0a" +
	"f{$D\xf2\xc5\x88\xcaK\x14\x1e_\xe4\xb1\x0cS" +
	"0\x0eQy\xb54\x12\xcb\xac\xf4\xac\xbb[\x9ao\x04" +
	"\xea*\x9b\x02:\xc8\xc1X\x9f\xed\xe2\xc9I.\xb4\xfb" +
	"\xf7\x98\xce\x9a\xf3\xcaW\x17\xf4y\x8c\x10s\x88\xb4\x86" +
	"\x80\x8b]F\x83Ny\x89\xf5w\x94\xad\x08\x99\x0d0" +
	"k\xc9g\x1a\xce\xe7\xae\xc8\xc7\x89\xaa\x86\x1d\x00e5" +
	"\x88e.D\xb5\x95y6>\xf5A>\x96S\xdd\xb8" +
	"\x0f\xa0\xcc\x8fXv\x17\xa2\xba\x0a\xc9\xc5\xf1\x99\x07\xf2" +
	"\xe9\x9e\xba\x8c\x02V\xd9R\xc4\xb2\xd5\x88j'\x92\xb3" +
	"\xe3\x83M\xe4\xd3\x15u-\x85\xaf\xb25\x88e\x9b\x10" +
	"\xd5\x1dHn\x8fO\x93\x90\x8fA\xd4-X\x0a@\x10" +
	"e\xdb\x11\xd5\xbdH\x0e\x90\x8f]\x91\xb7\xcd\xd5\x9d\x98" +
	"\x07@\x10e{\x10\xd5\xa3H\xae\x90\xbf0@\xde\xb1" +
	"W_e\xb8\xf6 \x96\x1dATO!9>>5" +
	"E>\xccV\x8f1\xba\xfe\x82Xv\x1aQ=\x8f\xe4" +
	":\xf9\x8c\x0a\xf9\xcb\x01\xf5,;\xf14b\xd9\xf7\x88" +
	"j\xa2\xe4\xc4\x04{\xb8\x83|\x1e\xad\xa2\xb4\x0e\xa0\xcc" +
	"!aY\x8a\x84\xea\x15\x92\x13/\xb1\xc7Y\xc8_8" +
	"\xa8\xc3\xa4\x00@Y\x9a\x84e9\x12\xaa\x93$\xa7\xc9" +
	"\x93P\xe4%\x166q/`w\xd3E\xd7\xd0e\xb1" +
	"\x18\xed\xfdRt\x16\xdb\xc3\x96\x9e\xf0\xf0\x8c\x1c.f" +
	"\xdfL\xb7\x01\xd86\x90}\xd1~\xb3?{\xba\xf8\xcd" +
	"~\x1db\x8d/\x06v!\xabz\x05\x1c\x18qb\x99" +
	"\xdf\xaf\x83x\xbb\x0ay\xbfJ\x0e\xb5\x0f\x08\x01\x9f\xa1" +
	"@&\x9b\xa2\x0cdo_=\x86\x01\x14K\xbc\xaf\xdb" +
	"K\xb5'\x0c\x05\xa2\xa7H}\xf7^b\xc5\xb3\x80X" +
	"\xbf\xcb]\xeb/\x0ag\xd92jc\x85\x08QP\x1b" +
	"\xa9\xbf\xa2\"W\x1fM\xe0>f=\xbc\x9b\xd4\xd7\xf9" +
	"\xe3#Ea\xd7\xf3\xfbW\xed\xf52:\xb3\x86R=" +
	"\x0dZ\xdc\x16$\xc5\x11\xc5L;\xa6m^\xb8f\xe9" +
	"\xae\xfe\x0d\x12\xc2\xb9\x1actQXt\xe2=\xef\x88" +
	"U\xe7R\xdap\x95\x8cZ\xa5d\x17g5\x90\xc4\xe4" +
	"\xca\xdb\x1b\xfdJ\x15\xba\xa4\x02T\xa3\xc9F\x9b6\x08" +
	"\xc5\x87\x1d\x09\x15\xc2\x0b\xad\x84:\x93\x8f\xe9\x00\xc0d" +
	"\xc3\xb1:\xdf\x02p\xd6\xe8\x8dZ2\x0b\x84\xfc\x8d\x18" +
	"r\x14JK\x1dH\x8a\x9b\xc2\x1f\x1f\x06#\x7f!\xa4" +
	"\xcc\xa1\x94r\x06\x05=\xfe\xda\x04\xf9\x13'\xa5\xbc\x03" +
	"$e*\x85:\xfe\xe6\x08\xf9\x0b+e\xd2\xdd )" +
	"\xe3X*\xca\x08b\xb9N1\x9a\xbc\x11\x0a\xb2\xcf\x1b" +
	"\xfe\x93\xb1\x16y\xe1\x8e\xd6*kH\xa3%\x8a\xa8\x9c" +
	"\xb8\xb7\xc6_,K\xa8\xed_\x1fe\xa6\x1b\xe4Xm" +
	"A\xdeP\xb1_\x0f\xf4\xa3\xbf\x16\xb3?e)O\x7f" +
	"\xdaA\xf6\x90\xa8/\x83\xaa\xeb\xd9\xa0\xfb\x99e\xf6k" +
	"~\xf0\xff>P\x17f\xad\xe1\x01\xdd\xc0\x86\xadq}" +
	"\x19jo\xbd\xc0>f\xb81\xfa[\xd1\x13oG\xac" +
	"\xfe\x9f\xd8\xfc\x93\xbd\xac\x98N\x91\x1d\x83M\x93\x9d\xbe" +
	"\x88\x14nAx\x9c\x97\x88\x17\xcc\xb0\xfc\x96TE\xe6" +
	"y#\xa4\xf3\xa6\xd58]1\x19@[*\xa3v\xbf" +
	"\x84h\xe5\xec\xf7-\x04\xd0VZ\xd9\xbd\x03\xc3\x0d\xa8" +
	"\xce\x0e\x00mSx\x06\x93\xe4\xf5y\x0d\x887u\x8f" +
	"\xa7\xa4\xbe\xde\x08\x02\x06!\xbe(\xe0\xf3\x18\xe5.t" +
	"\x82\x84N@Sw\x85U\x12\x8a\xc2J\x19\xedp\x02" +
	"F\xb3o\xbe\xd1k\xcf\xa3?3\xe9\xe8Q\x99\xb3\xc7" +
	"\x09\xad\xa5\xe9\x1c\x90\xe0\"6s\x83\x81\xac\xf0I\x96" +
	"\x1d\x88\xe1Q\xe9xaT*\xf1Q\xa9P\x90\x11\xbb" +
	"$D\xa5\x99\xbc\xbfGFm\x81\x84\xc9V\x0f\xaf\x95" +
	"d\xe0\x97Q\xbb\xb3\x97j\xa6\xafr%FE\x12\x1e" +
	"\xb1G\xf7\xadc\x05\x13\xde\x06\xb0\x0dJ\xa8?\xf3\xa2" +
	"\xebOD\xa5\x9c\x16o\x08\x07\x12\xd6\x9eDT\xb4\xc9" +
	"\x91\xaez\x0f\x93\x8c$\xbf\x1ej\xb2\x08\x02\x05/\x01" +
	"(\xf2\x1am5z]\xb7\x9a\xc9\x0a,\xd5Ff\xb8" +
	"\xb3\x88X)\xc7uw\xb6\xb1\xd3\x8e\x18c\x8d\xf1\x96" +
	"\x93*\x16\x9c\xd4u\x1b\xad\x9b\xdd&afK\xab\x11" +
	"h\xef\xb9\x016\xa0\xa8\xd8{\x89\xdf\xd7P8\xe46" +
	"\\e=\xbb\x81\xee\x93;\x9e\xe4\xa4\xd8\xf8\x16\x91\x92" +
	"\xdd)\xa3\xb6\\\xb8\xee\xb2\xd2\x88M\xf3\x96\xb28\xa2" +
	"\xb7[\xca\xab\x02\x00\xda\xfd2j\x0f\x0a-e\xb1\xb2" +
	"\xff\x07M;c\xb7\x85{~\x04\xd5\xc3\xa3\x8a\x9eT" +
	"=\xf6\xec9\x1a\x9a\xe2\xbd\x9d\x18\xff\xb3=\x8d\xe8a" +
	"\x06\xcbG\xc1\xfd\xe0\x8c\xcd\xf3\xff\x0b\x00\x00\xff\xffx" +
	"_\xaa\xae"

func init() {
	schemas.Register(schema_c8d91463cfc4fb4a,
		0x85e320f14a5d23e0,
		0x87d94955ce3c61dd,
		0x8c4a70a31703d35c,
		0x8c519e0dedc17d73,
		0x8f2ef49549d64e86,
		0x9206caa8d3e3cc7e,
		0x9210d9e69d14fa35,
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
		0xd76b6c6364d6bff5,
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
		0xfb3d38da0c9eaee6,
		0xfbbc20367c72bc59,
		0xfe7135f15d39bd5b)
}
