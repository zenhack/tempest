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
	"FaF6v\x06\x06\xe1\xbf\xac\xec@\xac\xce\xc0\xf0" +
	"\x7f\xb6\xbf\xd1[\xbb\xdbfO1T\x7f\x05\xa9\xe6\x04" +
	"\xabfec\x07bu\x86\xff\x0c\xd9\xffSs\x133" +
	"st3sY\x0ar\xf4\x92\x13\x0b\xf2\x0a\xac\x02R" +
	"\x8b\x8a3\x8bKR\xf3J\\Ara\xa9E\x99i" +
	"\x99\xa9E\x0c\x0c\x01\x8c\x8c\x01\xcc\xac\x81\x1c\x8c\x8c\xff" +
	"\xff\xaf\x94\x9e{\xea{\xc4\x15\x06\xa0\x8d[\xae\xee\xab" +
	"\xb9\xfe\xb6\xe70\x88\x8d\xcf4\xa8A)`SQL" +
	"\x8b\xfa\xe5\xeb\xc6\xf4\xb1\xfb\x07\x9ai\x80\x00\x00\x00\xff" +
	"\xffx\x05Z\x87"

func init() {
	schemas.Register(schema_92829022d203a580,
		0xd76bb6182f0aece3,
		0xe536db3eed324f9b)
}
