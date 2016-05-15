package persistentuiview

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	capnp "zombiezen.com/go/capnproto2"
	persistent "zombiezen.com/go/capnproto2/capnp/persistent"
	server "zombiezen.com/go/capnproto2/server"
)

type PersistentUiView struct{ Client capnp.Client }

func (c PersistentUiView) GetViewInfo(ctx context.Context, params func(grain.UiView_getViewInfo_Params) error, opts ...capnp.CallOption) grain.UiView_ViewInfo_Promise {
	if c.Client == nil {
		return grain.UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.UiView_getViewInfo_Params{Struct: s}) }
	}
	return grain.UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c PersistentUiView) NewSession(ctx context.Context, params func(grain.UiView_newSession_Params) error, opts ...capnp.CallOption) grain.UiView_newSession_Results_Promise {
	if c.Client == nil {
		return grain.UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.UiView_newSession_Params{Struct: s}) }
	}
	return grain.UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c PersistentUiView) NewRequestSession(ctx context.Context, params func(grain.UiView_newRequestSession_Params) error, opts ...capnp.CallOption) grain.UiView_newRequestSession_Results_Promise {
	if c.Client == nil {
		return grain.UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.UiView_newRequestSession_Params{Struct: s}) }
	}
	return grain.UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c PersistentUiView) NewOfferSession(ctx context.Context, params func(grain.UiView_newOfferSession_Params) error, opts ...capnp.CallOption) grain.UiView_newOfferSession_Results_Promise {
	if c.Client == nil {
		return grain.UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.UiView_newOfferSession_Params{Struct: s}) }
	}
	return grain.UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c PersistentUiView) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
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

func (c PersistentUiView) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type PersistentUiView_Server interface {
	GetViewInfo(grain.UiView_getViewInfo) error

	NewSession(grain.UiView_newSession) error

	NewRequestSession(grain.UiView_newRequestSession) error

	NewOfferSession(grain.UiView_newOfferSession) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentUiView_ServerToClient(s PersistentUiView_Server) PersistentUiView {
	c, _ := s.(server.Closer)
	return PersistentUiView{Client: server.New(PersistentUiView_Methods(nil, s), c)}
}

func PersistentUiView_Methods(methods []server.Method, s PersistentUiView_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.UiView_getViewInfo{c, opts, grain.UiView_getViewInfo_Params{Struct: p}, grain.UiView_ViewInfo{Struct: r}}
			return s.GetViewInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 5},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.UiView_newSession{c, opts, grain.UiView_newSession_Params{Struct: p}, grain.UiView_newSession_Results{Struct: r}}
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
			call := grain.UiView_newRequestSession{c, opts, grain.UiView_newRequestSession_Params{Struct: p}, grain.UiView_newRequestSession_Results{Struct: r}}
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
			call := grain.UiView_newOfferSession{c, opts, grain.UiView_newOfferSession_Params{Struct: p}, grain.UiView_newOfferSession_Results{Struct: r}}
			return s.NewOfferSession(call)
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
