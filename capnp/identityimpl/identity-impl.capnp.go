package identityimpl

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type PersistentIdentity struct{ Client capnp.Client }

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
	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentIdentity_ServerToClient(s PersistentIdentity_Server) PersistentIdentity {
	c, _ := s.(server.Closer)
	return PersistentIdentity{Client: server.New(PersistentIdentity_Methods(nil, s), c)}
}

func PersistentIdentity_Methods(methods []server.Method, s PersistentIdentity_Server) []server.Method {
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

const schema_acf10e4407376d3f = "x\xda2\xf0gt`2d\xfd\xcf\xc9\xc0\x10X\xc2" +
	"\xca\xf6\xbf\xea\xaf\xf9\xc1\xa7\xdb\xde\xacb\x10\x14a\xfe" +
	"o\x9fk\xce\xee\xc2\xf7q\x0d\x03\x03\xa3\xf1S\x16+" +
	"F\xe1\xaf,\xec\x0c\x0c\xc2\x1fY\xd8\x81X\x9d\xe1?" +
	"\x83\xc9\xff\xcc\x94\xd4\xbc\x92\xcc\x92Jf\xdd\xcc\xdc\x82" +
	"\x1c\xbd\xe4\xc4\x82\xbc\x02\xab\x80\xd4\xa2\xe2\xcc\xe2\x12\xa0" +
	"\x8cg\x8a=D>\x80\x911\x80\x995\x90\x83\x91\xf1" +
	"\xbf\xc4\xdd\xc0\x95U3Z\x0e000\xfc\xdfru" +
	"_\xcd\xf5\xb7=\x87\x81l@\x00\x00\x00\xff\xff~\xbe" +
	"0\xe9"

func init() {
	schemas.Register(schema_acf10e4407376d3f,
		0xaaecb6e5c137fd7a)
}
