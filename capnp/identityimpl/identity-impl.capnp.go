package identityimpl

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type PersistentIdentity struct{ Client capnp.Client }

func (c PersistentIdentity) GetProfile(ctx context.Context, params func(identity.Identity_getProfile_Params) error, opts ...capnp.CallOption) identity.Identity_getProfile_Results_Promise {
	if c.Client == nil {
		return identity.Identity_getProfile_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(identity.Identity_getProfile_Params{Struct: s}) }
	}
	return identity.Identity_getProfile_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentIdentity) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
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
func (c PersistentIdentity) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type PersistentIdentity_Server interface {
	GetProfile(identity.Identity_getProfile) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentIdentity_ServerToClient(s PersistentIdentity_Server) PersistentIdentity {
	c, _ := s.(server.Closer)
	return PersistentIdentity{Client: server.New(PersistentIdentity_Methods(nil, s), c)}
}

func PersistentIdentity_Methods(methods []server.Method, s PersistentIdentity_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := identity.Identity_getProfile{c, opts, identity.Identity_getProfile_Params{Struct: p}, identity.Identity_getProfile_Results{Struct: r}}
			return s.GetProfile(call)
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

const schema_acf10e4407376d3f = "x\xda2X\xce\xe8\xc0d\xc8\xba\x9f\x87\x81!p\x0a" +
	"+\xdb\xff\xaa\xbf\xe6\x07\x9fn{\xb3\x8aAP\x84\xf9" +
	"\xbf}\xae9\xbb\x0b\xdf\xc75\x0c\x0c\x8c\xc6\xaelV" +
	"\x8c\xc2\xa1l\xec\x0c\x0c\xc2\x81l\xec\xc2\x81l\xea\x0c" +
	"\xff\x19z\xfeg\xa6\xa4\xe6\x95d\x96T2\xebf\xe6" +
	"\x16\xe4\xe8%'\x16\xe4\x15X\x05\xa4\x16\x15g\x16\x97" +
	"\xa4\xe6\x95x\xa6\xd8C\xe4\x03\x18\x19\x03\x98Y\x039" +
	"\x18\x19\xffK\xdc\x0d\\Y5\xa3\xe5\x00\x03\x03\xc3\xff" +
	"-W\xf7\xd5\\\x7f\xdbs\x98\x81\x81\x01\x10\x00\x00\xff" +
	"\xff\xdc\x0f/\x04"

func init() {
	schemas.Register(schema_acf10e4407376d3f,
		0xaaecb6e5c137fd7a)
}
