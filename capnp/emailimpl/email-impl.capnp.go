package emailimpl

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	email "zenhack.net/go/sandstorm/capnp/email"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
	persistent "zombiezen.com/go/capnproto2/std/capnp/persistent"
)

type PersistentEmailVerifier struct{ Client capnp.Client }

func (c PersistentEmailVerifier) GetId(ctx context.Context, params func(email.EmailVerifier_getId_Params) error, opts ...capnp.CallOption) email.EmailVerifier_getId_Results_Promise {
	if c.Client == nil {
		return email.EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailVerifier_getId_Params{Struct: s}) }
	}
	return email.EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentEmailVerifier) VerifyEmail(ctx context.Context, params func(email.EmailVerifier_verifyEmail_Params) error, opts ...capnp.CallOption) email.EmailVerifier_verifyEmail_Results_Promise {
	if c.Client == nil {
		return email.EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailVerifier_verifyEmail_Params{Struct: s}) }
	}
	return email.EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c PersistentEmailVerifier) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
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
func (c PersistentEmailVerifier) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type PersistentEmailVerifier_Server interface {
	GetId(email.EmailVerifier_getId) error

	VerifyEmail(email.EmailVerifier_verifyEmail) error

	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentEmailVerifier_ServerToClient(s PersistentEmailVerifier_Server) PersistentEmailVerifier {
	c, _ := s.(server.Closer)
	return PersistentEmailVerifier{Client: server.New(PersistentEmailVerifier_Methods(nil, s), c)}
}

func PersistentEmailVerifier_Methods(methods []server.Method, s PersistentEmailVerifier_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailVerifier_getId{c, opts, email.EmailVerifier_getId_Params{Struct: p}, email.EmailVerifier_getId_Results{Struct: r}}
			return s.GetId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailVerifier_verifyEmail{c, opts, email.EmailVerifier_verifyEmail_Params{Struct: p}, email.EmailVerifier_verifyEmail_Results{Struct: r}}
			return s.VerifyEmail(call)
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

type PersistentVerifiedEmail struct{ Client capnp.Client }

func (c PersistentVerifiedEmail) AddRequirements(ctx context.Context, params func(supervisor.SystemPersistent_addRequirements_Params) error, opts ...capnp.CallOption) supervisor.SystemPersistent_addRequirements_Results_Promise {
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
func (c PersistentVerifiedEmail) Save(ctx context.Context, params func(persistent.Persistent_SaveParams) error, opts ...capnp.CallOption) persistent.Persistent_SaveResults_Promise {
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

type PersistentVerifiedEmail_Server interface {
	AddRequirements(supervisor.SystemPersistent_addRequirements) error

	Save(persistent.Persistent_save) error
}

func PersistentVerifiedEmail_ServerToClient(s PersistentVerifiedEmail_Server) PersistentVerifiedEmail {
	c, _ := s.(server.Closer)
	return PersistentVerifiedEmail{Client: server.New(PersistentVerifiedEmail_Methods(nil, s), c)}
}

func PersistentVerifiedEmail_Methods(methods []server.Method, s PersistentVerifiedEmail_Server) []server.Method {
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

const schema_92829022d203a580 = "x\xda2\x98\xcf\xe8\xc0d\xc8\xca\xcf\xc3\xc0\x10\xd8\xc3" +
	"\xca\xf6\xff\xf1\x1b.}\x89m\xd9\xd7\x19\x04\x05\x99\xff" +
	"7,e\xbe\xa44\xa1i\x12\x03\x03\xa3\xf1[V/" +
	"FaF6v\x06\x06\xe1\xbf\xac\xec\xc2\x7fY\xd5\x19" +
	"\x18\xfe\xcf\xf67zkw\xdb\xec)\x86\xea\xaf \xd5" +
	"\x9c`\xd5\xacl\xec\xc2\xacl\xea\x0c\xff\x19\xb2\xff\xa7" +
	"\xe6&f\xe6\xe8f\xe6\xb2\x14\xe4\xe8%'\x16\xe4\x15" +
	"X\x05\xa4\x16\x15g\x16\x97\xa4\xe6\x95\xb8\x82\xe4\xc2R" +
	"\x8b2\xd32S\x8b\x18\x18\x02\x18\x19\x03\x98Y\x039" +
	"\x18\x19\xff\xff_)=\xf7\xd4\xf7\x88+\x0c\x0c\x0c\xff" +
	"\xb7\\\xddWs\xfdm\xcfa\x10\x1b\x9fiP\x83R" +
	"\xc0\xa6\xa2\x98\x16\xf5\xcb\xd7\x8d\xe9c\xf7\x0f4\xd3\x00" +
	"\x01\x00\x00\xff\xffx\x05Z\x87"

func init() {
	schemas.Register(schema_92829022d203a580,
		0xd76bb6182f0aece3,
		0xe536db3eed324f9b)
}
