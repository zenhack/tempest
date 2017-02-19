package backend

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	spk "zenhack.net/go/sandstorm/capnp/spk"
	supervisor "zenhack.net/go/sandstorm/capnp/supervisor"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

// Constants defined in backend.capnp.
const (
	Backend_socketPath = "/var/sandstorm/socket/backend"
)

type Backend struct{ Client capnp.Client }

func (c Backend) StartGrain(ctx context.Context, params func(Backend_startGrain_Params) error, opts ...capnp.CallOption) Backend_startGrain_Results_Promise {
	if c.Client == nil {
		return Backend_startGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      0,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "startGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_startGrain_Params{Struct: s}) }
	}
	return Backend_startGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) GetGrain(ctx context.Context, params func(Backend_getGrain_Params) error, opts ...capnp.CallOption) Backend_getGrain_Results_Promise {
	if c.Client == nil {
		return Backend_getGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      1,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_getGrain_Params{Struct: s}) }
	}
	return Backend_getGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) DeleteGrain(ctx context.Context, params func(Backend_deleteGrain_Params) error, opts ...capnp.CallOption) Backend_deleteGrain_Results_Promise {
	if c.Client == nil {
		return Backend_deleteGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      2,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_deleteGrain_Params{Struct: s}) }
	}
	return Backend_deleteGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) InstallPackage(ctx context.Context, params func(Backend_installPackage_Params) error, opts ...capnp.CallOption) Backend_installPackage_Results_Promise {
	if c.Client == nil {
		return Backend_installPackage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      3,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "installPackage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_installPackage_Params{Struct: s}) }
	}
	return Backend_installPackage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) TryGetPackage(ctx context.Context, params func(Backend_tryGetPackage_Params) error, opts ...capnp.CallOption) Backend_tryGetPackage_Results_Promise {
	if c.Client == nil {
		return Backend_tryGetPackage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      4,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "tryGetPackage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_tryGetPackage_Params{Struct: s}) }
	}
	return Backend_tryGetPackage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) DeletePackage(ctx context.Context, params func(Backend_deletePackage_Params) error, opts ...capnp.CallOption) Backend_deletePackage_Results_Promise {
	if c.Client == nil {
		return Backend_deletePackage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      5,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deletePackage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_deletePackage_Params{Struct: s}) }
	}
	return Backend_deletePackage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) BackupGrain(ctx context.Context, params func(Backend_backupGrain_Params) error, opts ...capnp.CallOption) Backend_backupGrain_Results_Promise {
	if c.Client == nil {
		return Backend_backupGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      6,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "backupGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_backupGrain_Params{Struct: s}) }
	}
	return Backend_backupGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) RestoreGrain(ctx context.Context, params func(Backend_restoreGrain_Params) error, opts ...capnp.CallOption) Backend_restoreGrain_Results_Promise {
	if c.Client == nil {
		return Backend_restoreGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      7,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "restoreGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_restoreGrain_Params{Struct: s}) }
	}
	return Backend_restoreGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) UploadBackup(ctx context.Context, params func(Backend_uploadBackup_Params) error, opts ...capnp.CallOption) Backend_uploadBackup_Results_Promise {
	if c.Client == nil {
		return Backend_uploadBackup_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      8,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "uploadBackup",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_uploadBackup_Params{Struct: s}) }
	}
	return Backend_uploadBackup_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) DownloadBackup(ctx context.Context, params func(Backend_downloadBackup_Params) error, opts ...capnp.CallOption) Backend_downloadBackup_Results_Promise {
	if c.Client == nil {
		return Backend_downloadBackup_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      9,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "downloadBackup",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_downloadBackup_Params{Struct: s}) }
	}
	return Backend_downloadBackup_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) DeleteBackup(ctx context.Context, params func(Backend_deleteBackup_Params) error, opts ...capnp.CallOption) Backend_deleteBackup_Results_Promise {
	if c.Client == nil {
		return Backend_deleteBackup_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      10,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteBackup",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_deleteBackup_Params{Struct: s}) }
	}
	return Backend_deleteBackup_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) GetUserStorageUsage(ctx context.Context, params func(Backend_getUserStorageUsage_Params) error, opts ...capnp.CallOption) Backend_getUserStorageUsage_Results_Promise {
	if c.Client == nil {
		return Backend_getUserStorageUsage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      11,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getUserStorageUsage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_getUserStorageUsage_Params{Struct: s}) }
	}
	return Backend_getUserStorageUsage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) TransferGrain(ctx context.Context, params func(Backend_transferGrain_Params) error, opts ...capnp.CallOption) Backend_transferGrain_Results_Promise {
	if c.Client == nil {
		return Backend_transferGrain_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      12,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "transferGrain",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_transferGrain_Params{Struct: s}) }
	}
	return Backend_transferGrain_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) DeleteUser(ctx context.Context, params func(Backend_deleteUser_Params) error, opts ...capnp.CallOption) Backend_deleteUser_Results_Promise {
	if c.Client == nil {
		return Backend_deleteUser_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      13,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteUser",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_deleteUser_Params{Struct: s}) }
	}
	return Backend_deleteUser_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) Ping(ctx context.Context, params func(Backend_ping_Params) error, opts ...capnp.CallOption) Backend_ping_Results_Promise {
	if c.Client == nil {
		return Backend_ping_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      14,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "ping",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_ping_Params{Struct: s}) }
	}
	return Backend_ping_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend) GetGrainStorageUsage(ctx context.Context, params func(Backend_getGrainStorageUsage_Params) error, opts ...capnp.CallOption) Backend_getGrainStorageUsage_Results_Promise {
	if c.Client == nil {
		return Backend_getGrainStorageUsage_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      15,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getGrainStorageUsage",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_getGrainStorageUsage_Params{Struct: s}) }
	}
	return Backend_getGrainStorageUsage_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Backend_Server interface {
	StartGrain(Backend_startGrain) error

	GetGrain(Backend_getGrain) error

	DeleteGrain(Backend_deleteGrain) error

	InstallPackage(Backend_installPackage) error

	TryGetPackage(Backend_tryGetPackage) error

	DeletePackage(Backend_deletePackage) error

	BackupGrain(Backend_backupGrain) error

	RestoreGrain(Backend_restoreGrain) error

	UploadBackup(Backend_uploadBackup) error

	DownloadBackup(Backend_downloadBackup) error

	DeleteBackup(Backend_deleteBackup) error

	GetUserStorageUsage(Backend_getUserStorageUsage) error

	TransferGrain(Backend_transferGrain) error

	DeleteUser(Backend_deleteUser) error

	Ping(Backend_ping) error

	GetGrainStorageUsage(Backend_getGrainStorageUsage) error
}

func Backend_ServerToClient(s Backend_Server) Backend {
	c, _ := s.(server.Closer)
	return Backend{Client: server.New(Backend_Methods(nil, s), c)}
}

func Backend_Methods(methods []server.Method, s Backend_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 16)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      0,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "startGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_startGrain{c, opts, Backend_startGrain_Params{Struct: p}, Backend_startGrain_Results{Struct: r}}
			return s.StartGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      1,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_getGrain{c, opts, Backend_getGrain_Params{Struct: p}, Backend_getGrain_Results{Struct: r}}
			return s.GetGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      2,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_deleteGrain{c, opts, Backend_deleteGrain_Params{Struct: p}, Backend_deleteGrain_Results{Struct: r}}
			return s.DeleteGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      3,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "installPackage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_installPackage{c, opts, Backend_installPackage_Params{Struct: p}, Backend_installPackage_Results{Struct: r}}
			return s.InstallPackage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      4,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "tryGetPackage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_tryGetPackage{c, opts, Backend_tryGetPackage_Params{Struct: p}, Backend_tryGetPackage_Results{Struct: r}}
			return s.TryGetPackage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      5,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deletePackage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_deletePackage{c, opts, Backend_deletePackage_Params{Struct: p}, Backend_deletePackage_Results{Struct: r}}
			return s.DeletePackage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      6,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "backupGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_backupGrain{c, opts, Backend_backupGrain_Params{Struct: p}, Backend_backupGrain_Results{Struct: r}}
			return s.BackupGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      7,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "restoreGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_restoreGrain{c, opts, Backend_restoreGrain_Params{Struct: p}, Backend_restoreGrain_Results{Struct: r}}
			return s.RestoreGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      8,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "uploadBackup",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_uploadBackup{c, opts, Backend_uploadBackup_Params{Struct: p}, Backend_uploadBackup_Results{Struct: r}}
			return s.UploadBackup(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      9,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "downloadBackup",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_downloadBackup{c, opts, Backend_downloadBackup_Params{Struct: p}, Backend_downloadBackup_Results{Struct: r}}
			return s.DownloadBackup(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      10,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteBackup",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_deleteBackup{c, opts, Backend_deleteBackup_Params{Struct: p}, Backend_deleteBackup_Results{Struct: r}}
			return s.DeleteBackup(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      11,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getUserStorageUsage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_getUserStorageUsage{c, opts, Backend_getUserStorageUsage_Params{Struct: p}, Backend_getUserStorageUsage_Results{Struct: r}}
			return s.GetUserStorageUsage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      12,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "transferGrain",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_transferGrain{c, opts, Backend_transferGrain_Params{Struct: p}, Backend_transferGrain_Results{Struct: r}}
			return s.TransferGrain(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      13,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "deleteUser",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_deleteUser{c, opts, Backend_deleteUser_Params{Struct: p}, Backend_deleteUser_Results{Struct: r}}
			return s.DeleteUser(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      14,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "ping",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_ping{c, opts, Backend_ping_Params{Struct: p}, Backend_ping_Results{Struct: r}}
			return s.Ping(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc1b0e9713ac1ad4f,
			MethodID:      15,
			InterfaceName: "backend.capnp:Backend",
			MethodName:    "getGrainStorageUsage",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_getGrainStorageUsage{c, opts, Backend_getGrainStorageUsage_Params{Struct: p}, Backend_getGrainStorageUsage_Results{Struct: r}}
			return s.GetGrainStorageUsage(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// Backend_startGrain holds the arguments for a server call to Backend.startGrain.
type Backend_startGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_startGrain_Params
	Results Backend_startGrain_Results
}

// Backend_getGrain holds the arguments for a server call to Backend.getGrain.
type Backend_getGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_getGrain_Params
	Results Backend_getGrain_Results
}

// Backend_deleteGrain holds the arguments for a server call to Backend.deleteGrain.
type Backend_deleteGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_deleteGrain_Params
	Results Backend_deleteGrain_Results
}

// Backend_installPackage holds the arguments for a server call to Backend.installPackage.
type Backend_installPackage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_installPackage_Params
	Results Backend_installPackage_Results
}

// Backend_tryGetPackage holds the arguments for a server call to Backend.tryGetPackage.
type Backend_tryGetPackage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_tryGetPackage_Params
	Results Backend_tryGetPackage_Results
}

// Backend_deletePackage holds the arguments for a server call to Backend.deletePackage.
type Backend_deletePackage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_deletePackage_Params
	Results Backend_deletePackage_Results
}

// Backend_backupGrain holds the arguments for a server call to Backend.backupGrain.
type Backend_backupGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_backupGrain_Params
	Results Backend_backupGrain_Results
}

// Backend_restoreGrain holds the arguments for a server call to Backend.restoreGrain.
type Backend_restoreGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_restoreGrain_Params
	Results Backend_restoreGrain_Results
}

// Backend_uploadBackup holds the arguments for a server call to Backend.uploadBackup.
type Backend_uploadBackup struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_uploadBackup_Params
	Results Backend_uploadBackup_Results
}

// Backend_downloadBackup holds the arguments for a server call to Backend.downloadBackup.
type Backend_downloadBackup struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_downloadBackup_Params
	Results Backend_downloadBackup_Results
}

// Backend_deleteBackup holds the arguments for a server call to Backend.deleteBackup.
type Backend_deleteBackup struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_deleteBackup_Params
	Results Backend_deleteBackup_Results
}

// Backend_getUserStorageUsage holds the arguments for a server call to Backend.getUserStorageUsage.
type Backend_getUserStorageUsage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_getUserStorageUsage_Params
	Results Backend_getUserStorageUsage_Results
}

// Backend_transferGrain holds the arguments for a server call to Backend.transferGrain.
type Backend_transferGrain struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_transferGrain_Params
	Results Backend_transferGrain_Results
}

// Backend_deleteUser holds the arguments for a server call to Backend.deleteUser.
type Backend_deleteUser struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_deleteUser_Params
	Results Backend_deleteUser_Results
}

// Backend_ping holds the arguments for a server call to Backend.ping.
type Backend_ping struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_ping_Params
	Results Backend_ping_Results
}

// Backend_getGrainStorageUsage holds the arguments for a server call to Backend.getGrainStorageUsage.
type Backend_getGrainStorageUsage struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_getGrainStorageUsage_Params
	Results Backend_getGrainStorageUsage_Results
}

type Backend_PackageUploadStream struct{ Client capnp.Client }

func (c Backend_PackageUploadStream) SaveAs(ctx context.Context, params func(Backend_PackageUploadStream_saveAs_Params) error, opts ...capnp.CallOption) Backend_PackageUploadStream_saveAs_Results_Promise {
	if c.Client == nil {
		return Backend_PackageUploadStream_saveAs_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xb481d35d0da2713c,
			MethodID:      0,
			InterfaceName: "backend.capnp:Backend.PackageUploadStream",
			MethodName:    "saveAs",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Backend_PackageUploadStream_saveAs_Params{Struct: s}) }
	}
	return Backend_PackageUploadStream_saveAs_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend_PackageUploadStream) Write(ctx context.Context, params func(util.ByteStream_write_Params) error, opts ...capnp.CallOption) util.ByteStream_write_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_write_Params{Struct: s}) }
	}
	return util.ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend_PackageUploadStream) Done(ctx context.Context, params func(util.ByteStream_done_Params) error, opts ...capnp.CallOption) util.ByteStream_done_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_done_Params{Struct: s}) }
	}
	return util.ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Backend_PackageUploadStream) ExpectSize(ctx context.Context, params func(util.ByteStream_expectSize_Params) error, opts ...capnp.CallOption) util.ByteStream_expectSize_Results_Promise {
	if c.Client == nil {
		return util.ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(util.ByteStream_expectSize_Params{Struct: s}) }
	}
	return util.ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Backend_PackageUploadStream_Server interface {
	SaveAs(Backend_PackageUploadStream_saveAs) error

	Write(util.ByteStream_write) error

	Done(util.ByteStream_done) error

	ExpectSize(util.ByteStream_expectSize) error
}

func Backend_PackageUploadStream_ServerToClient(s Backend_PackageUploadStream_Server) Backend_PackageUploadStream {
	c, _ := s.(server.Closer)
	return Backend_PackageUploadStream{Client: server.New(Backend_PackageUploadStream_Methods(nil, s), c)}
}

func Backend_PackageUploadStream_Methods(methods []server.Method, s Backend_PackageUploadStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb481d35d0da2713c,
			MethodID:      0,
			InterfaceName: "backend.capnp:Backend.PackageUploadStream",
			MethodName:    "saveAs",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Backend_PackageUploadStream_saveAs{c, opts, Backend_PackageUploadStream_saveAs_Params{Struct: p}, Backend_PackageUploadStream_saveAs_Results{Struct: r}}
			return s.SaveAs(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_write{c, opts, util.ByteStream_write_Params{Struct: p}, util.ByteStream_write_Results{Struct: r}}
			return s.Write(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_done{c, opts, util.ByteStream_done_Params{Struct: p}, util.ByteStream_done_Results{Struct: r}}
			return s.Done(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := util.ByteStream_expectSize{c, opts, util.ByteStream_expectSize_Params{Struct: p}, util.ByteStream_expectSize_Results{Struct: r}}
			return s.ExpectSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Backend_PackageUploadStream_saveAs holds the arguments for a server call to Backend_PackageUploadStream.saveAs.
type Backend_PackageUploadStream_saveAs struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Backend_PackageUploadStream_saveAs_Params
	Results Backend_PackageUploadStream_saveAs_Results
}

type Backend_PackageUploadStream_saveAs_Params struct{ capnp.Struct }

// Backend_PackageUploadStream_saveAs_Params_TypeID is the unique identifier for the type Backend_PackageUploadStream_saveAs_Params.
const Backend_PackageUploadStream_saveAs_Params_TypeID = 0x86ca17d397d72d2b

func NewBackend_PackageUploadStream_saveAs_Params(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_PackageUploadStream_saveAs_Params{st}, err
}

func NewRootBackend_PackageUploadStream_saveAs_Params(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_PackageUploadStream_saveAs_Params{st}, err
}

func ReadRootBackend_PackageUploadStream_saveAs_Params(msg *capnp.Message) (Backend_PackageUploadStream_saveAs_Params, error) {
	root, err := msg.RootPtr()
	return Backend_PackageUploadStream_saveAs_Params{root.Struct()}, err
}

func (s Backend_PackageUploadStream_saveAs_Params) String() string {
	str, _ := text.Marshal(0x86ca17d397d72d2b, s.Struct)
	return str
}

func (s Backend_PackageUploadStream_saveAs_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_PackageUploadStream_saveAs_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_PackageUploadStream_saveAs_Params) SetPackageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_PackageUploadStream_saveAs_Params_List is a list of Backend_PackageUploadStream_saveAs_Params.
type Backend_PackageUploadStream_saveAs_Params_List struct{ capnp.List }

// NewBackend_PackageUploadStream_saveAs_Params creates a new list of Backend_PackageUploadStream_saveAs_Params.
func NewBackend_PackageUploadStream_saveAs_Params_List(s *capnp.Segment, sz int32) (Backend_PackageUploadStream_saveAs_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_PackageUploadStream_saveAs_Params_List{l}, err
}

func (s Backend_PackageUploadStream_saveAs_Params_List) At(i int) Backend_PackageUploadStream_saveAs_Params {
	return Backend_PackageUploadStream_saveAs_Params{s.List.Struct(i)}
}

func (s Backend_PackageUploadStream_saveAs_Params_List) Set(i int, v Backend_PackageUploadStream_saveAs_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_PackageUploadStream_saveAs_Params_Promise is a wrapper for a Backend_PackageUploadStream_saveAs_Params promised by a client call.
type Backend_PackageUploadStream_saveAs_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_PackageUploadStream_saveAs_Params_Promise) Struct() (Backend_PackageUploadStream_saveAs_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_PackageUploadStream_saveAs_Params{s}, err
}

type Backend_PackageUploadStream_saveAs_Results struct{ capnp.Struct }

// Backend_PackageUploadStream_saveAs_Results_TypeID is the unique identifier for the type Backend_PackageUploadStream_saveAs_Results.
const Backend_PackageUploadStream_saveAs_Results_TypeID = 0xa019dbe64a38e85d

func NewBackend_PackageUploadStream_saveAs_Results(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_PackageUploadStream_saveAs_Results{st}, err
}

func NewRootBackend_PackageUploadStream_saveAs_Results(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_PackageUploadStream_saveAs_Results{st}, err
}

func ReadRootBackend_PackageUploadStream_saveAs_Results(msg *capnp.Message) (Backend_PackageUploadStream_saveAs_Results, error) {
	root, err := msg.RootPtr()
	return Backend_PackageUploadStream_saveAs_Results{root.Struct()}, err
}

func (s Backend_PackageUploadStream_saveAs_Results) String() string {
	str, _ := text.Marshal(0xa019dbe64a38e85d, s.Struct)
	return str
}

func (s Backend_PackageUploadStream_saveAs_Results) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_PackageUploadStream_saveAs_Results) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Results) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_PackageUploadStream_saveAs_Results) SetAppId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_PackageUploadStream_saveAs_Results) Manifest() (spk.Manifest, error) {
	p, err := s.Struct.Ptr(1)
	return spk.Manifest{Struct: p.Struct()}, err
}

func (s Backend_PackageUploadStream_saveAs_Results) HasManifest() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Results) SetManifest(v spk.Manifest) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewManifest sets the manifest field to a newly
// allocated spk.Manifest struct, preferring placement in s's segment.
func (s Backend_PackageUploadStream_saveAs_Results) NewManifest() (spk.Manifest, error) {
	ss, err := spk.NewManifest(s.Struct.Segment())
	if err != nil {
		return spk.Manifest{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Backend_PackageUploadStream_saveAs_Results) AuthorPgpKeyFingerprint() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_PackageUploadStream_saveAs_Results) HasAuthorPgpKeyFingerprint() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Results) AuthorPgpKeyFingerprintBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_PackageUploadStream_saveAs_Results) SetAuthorPgpKeyFingerprint(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// Backend_PackageUploadStream_saveAs_Results_List is a list of Backend_PackageUploadStream_saveAs_Results.
type Backend_PackageUploadStream_saveAs_Results_List struct{ capnp.List }

// NewBackend_PackageUploadStream_saveAs_Results creates a new list of Backend_PackageUploadStream_saveAs_Results.
func NewBackend_PackageUploadStream_saveAs_Results_List(s *capnp.Segment, sz int32) (Backend_PackageUploadStream_saveAs_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Backend_PackageUploadStream_saveAs_Results_List{l}, err
}

func (s Backend_PackageUploadStream_saveAs_Results_List) At(i int) Backend_PackageUploadStream_saveAs_Results {
	return Backend_PackageUploadStream_saveAs_Results{s.List.Struct(i)}
}

func (s Backend_PackageUploadStream_saveAs_Results_List) Set(i int, v Backend_PackageUploadStream_saveAs_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_PackageUploadStream_saveAs_Results_Promise is a wrapper for a Backend_PackageUploadStream_saveAs_Results promised by a client call.
type Backend_PackageUploadStream_saveAs_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_PackageUploadStream_saveAs_Results_Promise) Struct() (Backend_PackageUploadStream_saveAs_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_PackageUploadStream_saveAs_Results{s}, err
}

func (p Backend_PackageUploadStream_saveAs_Results_Promise) Manifest() spk.Manifest_Promise {
	return spk.Manifest_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Backend_startGrain_Params struct{ capnp.Struct }

// Backend_startGrain_Params_TypeID is the unique identifier for the type Backend_startGrain_Params.
const Backend_startGrain_Params_TypeID = 0xadfbf90ef9c01c9a

func NewBackend_startGrain_Params(s *capnp.Segment) (Backend_startGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Backend_startGrain_Params{st}, err
}

func NewRootBackend_startGrain_Params(s *capnp.Segment) (Backend_startGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Backend_startGrain_Params{st}, err
}

func ReadRootBackend_startGrain_Params(msg *capnp.Message) (Backend_startGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_startGrain_Params{root.Struct()}, err
}

func (s Backend_startGrain_Params) String() string {
	str, _ := text.Marshal(0xadfbf90ef9c01c9a, s.Struct)
	return str
}

func (s Backend_startGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_startGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_startGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_startGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_startGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_startGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Backend_startGrain_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_startGrain_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_startGrain_Params) SetPackageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s Backend_startGrain_Params) Command() (spk.Manifest_Command, error) {
	p, err := s.Struct.Ptr(3)
	return spk.Manifest_Command{Struct: p.Struct()}, err
}

func (s Backend_startGrain_Params) HasCommand() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) SetCommand(v spk.Manifest_Command) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewCommand sets the command field to a newly
// allocated spk.Manifest_Command struct, preferring placement in s's segment.
func (s Backend_startGrain_Params) NewCommand() (spk.Manifest_Command, error) {
	ss, err := spk.NewManifest_Command(s.Struct.Segment())
	if err != nil {
		return spk.Manifest_Command{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Backend_startGrain_Params) IsNew() bool {
	return s.Struct.Bit(0)
}

func (s Backend_startGrain_Params) SetIsNew(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Backend_startGrain_Params) DevMode() bool {
	return s.Struct.Bit(1)
}

func (s Backend_startGrain_Params) SetDevMode(v bool) {
	s.Struct.SetBit(1, v)
}

func (s Backend_startGrain_Params) MountProc() bool {
	return s.Struct.Bit(2)
}

func (s Backend_startGrain_Params) SetMountProc(v bool) {
	s.Struct.SetBit(2, v)
}

// Backend_startGrain_Params_List is a list of Backend_startGrain_Params.
type Backend_startGrain_Params_List struct{ capnp.List }

// NewBackend_startGrain_Params creates a new list of Backend_startGrain_Params.
func NewBackend_startGrain_Params_List(s *capnp.Segment, sz int32) (Backend_startGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return Backend_startGrain_Params_List{l}, err
}

func (s Backend_startGrain_Params_List) At(i int) Backend_startGrain_Params {
	return Backend_startGrain_Params{s.List.Struct(i)}
}

func (s Backend_startGrain_Params_List) Set(i int, v Backend_startGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_startGrain_Params_Promise is a wrapper for a Backend_startGrain_Params promised by a client call.
type Backend_startGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_startGrain_Params_Promise) Struct() (Backend_startGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_startGrain_Params{s}, err
}

func (p Backend_startGrain_Params_Promise) Command() spk.Manifest_Command_Promise {
	return spk.Manifest_Command_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type Backend_startGrain_Results struct{ capnp.Struct }

// Backend_startGrain_Results_TypeID is the unique identifier for the type Backend_startGrain_Results.
const Backend_startGrain_Results_TypeID = 0xac9557813c4f78cf

func NewBackend_startGrain_Results(s *capnp.Segment) (Backend_startGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_startGrain_Results{st}, err
}

func NewRootBackend_startGrain_Results(s *capnp.Segment) (Backend_startGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_startGrain_Results{st}, err
}

func ReadRootBackend_startGrain_Results(msg *capnp.Message) (Backend_startGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_startGrain_Results{root.Struct()}, err
}

func (s Backend_startGrain_Results) String() string {
	str, _ := text.Marshal(0xac9557813c4f78cf, s.Struct)
	return str
}

func (s Backend_startGrain_Results) Supervisor() supervisor.Supervisor {
	p, _ := s.Struct.Ptr(0)
	return supervisor.Supervisor{Client: p.Interface().Client()}
}

func (s Backend_startGrain_Results) HasSupervisor() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Results) SetSupervisor(v supervisor.Supervisor) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Backend_startGrain_Results_List is a list of Backend_startGrain_Results.
type Backend_startGrain_Results_List struct{ capnp.List }

// NewBackend_startGrain_Results creates a new list of Backend_startGrain_Results.
func NewBackend_startGrain_Results_List(s *capnp.Segment, sz int32) (Backend_startGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_startGrain_Results_List{l}, err
}

func (s Backend_startGrain_Results_List) At(i int) Backend_startGrain_Results {
	return Backend_startGrain_Results{s.List.Struct(i)}
}

func (s Backend_startGrain_Results_List) Set(i int, v Backend_startGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_startGrain_Results_Promise is a wrapper for a Backend_startGrain_Results promised by a client call.
type Backend_startGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_startGrain_Results_Promise) Struct() (Backend_startGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_startGrain_Results{s}, err
}

func (p Backend_startGrain_Results_Promise) Supervisor() supervisor.Supervisor {
	return supervisor.Supervisor{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Backend_getGrain_Params struct{ capnp.Struct }

// Backend_getGrain_Params_TypeID is the unique identifier for the type Backend_getGrain_Params.
const Backend_getGrain_Params_TypeID = 0xe4d3afafc9fe1acf

func NewBackend_getGrain_Params(s *capnp.Segment) (Backend_getGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_getGrain_Params{st}, err
}

func NewRootBackend_getGrain_Params(s *capnp.Segment) (Backend_getGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_getGrain_Params{st}, err
}

func ReadRootBackend_getGrain_Params(msg *capnp.Message) (Backend_getGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_getGrain_Params{root.Struct()}, err
}

func (s Backend_getGrain_Params) String() string {
	str, _ := text.Marshal(0xe4d3afafc9fe1acf, s.Struct)
	return str
}

func (s Backend_getGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_getGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_getGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_getGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_getGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_getGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// Backend_getGrain_Params_List is a list of Backend_getGrain_Params.
type Backend_getGrain_Params_List struct{ capnp.List }

// NewBackend_getGrain_Params creates a new list of Backend_getGrain_Params.
func NewBackend_getGrain_Params_List(s *capnp.Segment, sz int32) (Backend_getGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Backend_getGrain_Params_List{l}, err
}

func (s Backend_getGrain_Params_List) At(i int) Backend_getGrain_Params {
	return Backend_getGrain_Params{s.List.Struct(i)}
}

func (s Backend_getGrain_Params_List) Set(i int, v Backend_getGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getGrain_Params_Promise is a wrapper for a Backend_getGrain_Params promised by a client call.
type Backend_getGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_getGrain_Params_Promise) Struct() (Backend_getGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getGrain_Params{s}, err
}

type Backend_getGrain_Results struct{ capnp.Struct }

// Backend_getGrain_Results_TypeID is the unique identifier for the type Backend_getGrain_Results.
const Backend_getGrain_Results_TypeID = 0xea0b2836fb52aee9

func NewBackend_getGrain_Results(s *capnp.Segment) (Backend_getGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_getGrain_Results{st}, err
}

func NewRootBackend_getGrain_Results(s *capnp.Segment) (Backend_getGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_getGrain_Results{st}, err
}

func ReadRootBackend_getGrain_Results(msg *capnp.Message) (Backend_getGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_getGrain_Results{root.Struct()}, err
}

func (s Backend_getGrain_Results) String() string {
	str, _ := text.Marshal(0xea0b2836fb52aee9, s.Struct)
	return str
}

func (s Backend_getGrain_Results) Supervisor() supervisor.Supervisor {
	p, _ := s.Struct.Ptr(0)
	return supervisor.Supervisor{Client: p.Interface().Client()}
}

func (s Backend_getGrain_Results) HasSupervisor() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Results) SetSupervisor(v supervisor.Supervisor) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Backend_getGrain_Results_List is a list of Backend_getGrain_Results.
type Backend_getGrain_Results_List struct{ capnp.List }

// NewBackend_getGrain_Results creates a new list of Backend_getGrain_Results.
func NewBackend_getGrain_Results_List(s *capnp.Segment, sz int32) (Backend_getGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_getGrain_Results_List{l}, err
}

func (s Backend_getGrain_Results_List) At(i int) Backend_getGrain_Results {
	return Backend_getGrain_Results{s.List.Struct(i)}
}

func (s Backend_getGrain_Results_List) Set(i int, v Backend_getGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getGrain_Results_Promise is a wrapper for a Backend_getGrain_Results promised by a client call.
type Backend_getGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_getGrain_Results_Promise) Struct() (Backend_getGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getGrain_Results{s}, err
}

func (p Backend_getGrain_Results_Promise) Supervisor() supervisor.Supervisor {
	return supervisor.Supervisor{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Backend_deleteGrain_Params struct{ capnp.Struct }

// Backend_deleteGrain_Params_TypeID is the unique identifier for the type Backend_deleteGrain_Params.
const Backend_deleteGrain_Params_TypeID = 0xd0669675481ed533

func NewBackend_deleteGrain_Params(s *capnp.Segment) (Backend_deleteGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_deleteGrain_Params{st}, err
}

func NewRootBackend_deleteGrain_Params(s *capnp.Segment) (Backend_deleteGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_deleteGrain_Params{st}, err
}

func ReadRootBackend_deleteGrain_Params(msg *capnp.Message) (Backend_deleteGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_deleteGrain_Params{root.Struct()}, err
}

func (s Backend_deleteGrain_Params) String() string {
	str, _ := text.Marshal(0xd0669675481ed533, s.Struct)
	return str
}

func (s Backend_deleteGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_deleteGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deleteGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_deleteGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_deleteGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_deleteGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_deleteGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_deleteGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// Backend_deleteGrain_Params_List is a list of Backend_deleteGrain_Params.
type Backend_deleteGrain_Params_List struct{ capnp.List }

// NewBackend_deleteGrain_Params creates a new list of Backend_deleteGrain_Params.
func NewBackend_deleteGrain_Params_List(s *capnp.Segment, sz int32) (Backend_deleteGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Backend_deleteGrain_Params_List{l}, err
}

func (s Backend_deleteGrain_Params_List) At(i int) Backend_deleteGrain_Params {
	return Backend_deleteGrain_Params{s.List.Struct(i)}
}

func (s Backend_deleteGrain_Params_List) Set(i int, v Backend_deleteGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteGrain_Params_Promise is a wrapper for a Backend_deleteGrain_Params promised by a client call.
type Backend_deleteGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteGrain_Params_Promise) Struct() (Backend_deleteGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteGrain_Params{s}, err
}

type Backend_deleteGrain_Results struct{ capnp.Struct }

// Backend_deleteGrain_Results_TypeID is the unique identifier for the type Backend_deleteGrain_Results.
const Backend_deleteGrain_Results_TypeID = 0x9aa99e08dd1161ff

func NewBackend_deleteGrain_Results(s *capnp.Segment) (Backend_deleteGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteGrain_Results{st}, err
}

func NewRootBackend_deleteGrain_Results(s *capnp.Segment) (Backend_deleteGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteGrain_Results{st}, err
}

func ReadRootBackend_deleteGrain_Results(msg *capnp.Message) (Backend_deleteGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_deleteGrain_Results{root.Struct()}, err
}

func (s Backend_deleteGrain_Results) String() string {
	str, _ := text.Marshal(0x9aa99e08dd1161ff, s.Struct)
	return str
}

// Backend_deleteGrain_Results_List is a list of Backend_deleteGrain_Results.
type Backend_deleteGrain_Results_List struct{ capnp.List }

// NewBackend_deleteGrain_Results creates a new list of Backend_deleteGrain_Results.
func NewBackend_deleteGrain_Results_List(s *capnp.Segment, sz int32) (Backend_deleteGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_deleteGrain_Results_List{l}, err
}

func (s Backend_deleteGrain_Results_List) At(i int) Backend_deleteGrain_Results {
	return Backend_deleteGrain_Results{s.List.Struct(i)}
}

func (s Backend_deleteGrain_Results_List) Set(i int, v Backend_deleteGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteGrain_Results_Promise is a wrapper for a Backend_deleteGrain_Results promised by a client call.
type Backend_deleteGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteGrain_Results_Promise) Struct() (Backend_deleteGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteGrain_Results{s}, err
}

type Backend_installPackage_Params struct{ capnp.Struct }

// Backend_installPackage_Params_TypeID is the unique identifier for the type Backend_installPackage_Params.
const Backend_installPackage_Params_TypeID = 0xa98fd02dd93dd26b

func NewBackend_installPackage_Params(s *capnp.Segment) (Backend_installPackage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_installPackage_Params{st}, err
}

func NewRootBackend_installPackage_Params(s *capnp.Segment) (Backend_installPackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_installPackage_Params{st}, err
}

func ReadRootBackend_installPackage_Params(msg *capnp.Message) (Backend_installPackage_Params, error) {
	root, err := msg.RootPtr()
	return Backend_installPackage_Params{root.Struct()}, err
}

func (s Backend_installPackage_Params) String() string {
	str, _ := text.Marshal(0xa98fd02dd93dd26b, s.Struct)
	return str
}

// Backend_installPackage_Params_List is a list of Backend_installPackage_Params.
type Backend_installPackage_Params_List struct{ capnp.List }

// NewBackend_installPackage_Params creates a new list of Backend_installPackage_Params.
func NewBackend_installPackage_Params_List(s *capnp.Segment, sz int32) (Backend_installPackage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_installPackage_Params_List{l}, err
}

func (s Backend_installPackage_Params_List) At(i int) Backend_installPackage_Params {
	return Backend_installPackage_Params{s.List.Struct(i)}
}

func (s Backend_installPackage_Params_List) Set(i int, v Backend_installPackage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_installPackage_Params_Promise is a wrapper for a Backend_installPackage_Params promised by a client call.
type Backend_installPackage_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_installPackage_Params_Promise) Struct() (Backend_installPackage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_installPackage_Params{s}, err
}

type Backend_installPackage_Results struct{ capnp.Struct }

// Backend_installPackage_Results_TypeID is the unique identifier for the type Backend_installPackage_Results.
const Backend_installPackage_Results_TypeID = 0x8829b2e76d8325f1

func NewBackend_installPackage_Results(s *capnp.Segment) (Backend_installPackage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_installPackage_Results{st}, err
}

func NewRootBackend_installPackage_Results(s *capnp.Segment) (Backend_installPackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_installPackage_Results{st}, err
}

func ReadRootBackend_installPackage_Results(msg *capnp.Message) (Backend_installPackage_Results, error) {
	root, err := msg.RootPtr()
	return Backend_installPackage_Results{root.Struct()}, err
}

func (s Backend_installPackage_Results) String() string {
	str, _ := text.Marshal(0x8829b2e76d8325f1, s.Struct)
	return str
}

func (s Backend_installPackage_Results) Stream() Backend_PackageUploadStream {
	p, _ := s.Struct.Ptr(0)
	return Backend_PackageUploadStream{Client: p.Interface().Client()}
}

func (s Backend_installPackage_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_installPackage_Results) SetStream(v Backend_PackageUploadStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Backend_installPackage_Results_List is a list of Backend_installPackage_Results.
type Backend_installPackage_Results_List struct{ capnp.List }

// NewBackend_installPackage_Results creates a new list of Backend_installPackage_Results.
func NewBackend_installPackage_Results_List(s *capnp.Segment, sz int32) (Backend_installPackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_installPackage_Results_List{l}, err
}

func (s Backend_installPackage_Results_List) At(i int) Backend_installPackage_Results {
	return Backend_installPackage_Results{s.List.Struct(i)}
}

func (s Backend_installPackage_Results_List) Set(i int, v Backend_installPackage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_installPackage_Results_Promise is a wrapper for a Backend_installPackage_Results promised by a client call.
type Backend_installPackage_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_installPackage_Results_Promise) Struct() (Backend_installPackage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_installPackage_Results{s}, err
}

func (p Backend_installPackage_Results_Promise) Stream() Backend_PackageUploadStream {
	return Backend_PackageUploadStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Backend_tryGetPackage_Params struct{ capnp.Struct }

// Backend_tryGetPackage_Params_TypeID is the unique identifier for the type Backend_tryGetPackage_Params.
const Backend_tryGetPackage_Params_TypeID = 0xfb4cd9916f42104c

func NewBackend_tryGetPackage_Params(s *capnp.Segment) (Backend_tryGetPackage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_tryGetPackage_Params{st}, err
}

func NewRootBackend_tryGetPackage_Params(s *capnp.Segment) (Backend_tryGetPackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_tryGetPackage_Params{st}, err
}

func ReadRootBackend_tryGetPackage_Params(msg *capnp.Message) (Backend_tryGetPackage_Params, error) {
	root, err := msg.RootPtr()
	return Backend_tryGetPackage_Params{root.Struct()}, err
}

func (s Backend_tryGetPackage_Params) String() string {
	str, _ := text.Marshal(0xfb4cd9916f42104c, s.Struct)
	return str
}

func (s Backend_tryGetPackage_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_tryGetPackage_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_tryGetPackage_Params) SetPackageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_tryGetPackage_Params_List is a list of Backend_tryGetPackage_Params.
type Backend_tryGetPackage_Params_List struct{ capnp.List }

// NewBackend_tryGetPackage_Params creates a new list of Backend_tryGetPackage_Params.
func NewBackend_tryGetPackage_Params_List(s *capnp.Segment, sz int32) (Backend_tryGetPackage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_tryGetPackage_Params_List{l}, err
}

func (s Backend_tryGetPackage_Params_List) At(i int) Backend_tryGetPackage_Params {
	return Backend_tryGetPackage_Params{s.List.Struct(i)}
}

func (s Backend_tryGetPackage_Params_List) Set(i int, v Backend_tryGetPackage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_tryGetPackage_Params_Promise is a wrapper for a Backend_tryGetPackage_Params promised by a client call.
type Backend_tryGetPackage_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_tryGetPackage_Params_Promise) Struct() (Backend_tryGetPackage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_tryGetPackage_Params{s}, err
}

type Backend_tryGetPackage_Results struct{ capnp.Struct }

// Backend_tryGetPackage_Results_TypeID is the unique identifier for the type Backend_tryGetPackage_Results.
const Backend_tryGetPackage_Results_TypeID = 0xef241fd6058030cf

func NewBackend_tryGetPackage_Results(s *capnp.Segment) (Backend_tryGetPackage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_tryGetPackage_Results{st}, err
}

func NewRootBackend_tryGetPackage_Results(s *capnp.Segment) (Backend_tryGetPackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_tryGetPackage_Results{st}, err
}

func ReadRootBackend_tryGetPackage_Results(msg *capnp.Message) (Backend_tryGetPackage_Results, error) {
	root, err := msg.RootPtr()
	return Backend_tryGetPackage_Results{root.Struct()}, err
}

func (s Backend_tryGetPackage_Results) String() string {
	str, _ := text.Marshal(0xef241fd6058030cf, s.Struct)
	return str
}

func (s Backend_tryGetPackage_Results) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_tryGetPackage_Results) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Results) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_tryGetPackage_Results) SetAppId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_tryGetPackage_Results) Manifest() (spk.Manifest, error) {
	p, err := s.Struct.Ptr(1)
	return spk.Manifest{Struct: p.Struct()}, err
}

func (s Backend_tryGetPackage_Results) HasManifest() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Results) SetManifest(v spk.Manifest) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewManifest sets the manifest field to a newly
// allocated spk.Manifest struct, preferring placement in s's segment.
func (s Backend_tryGetPackage_Results) NewManifest() (spk.Manifest, error) {
	ss, err := spk.NewManifest(s.Struct.Segment())
	if err != nil {
		return spk.Manifest{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Backend_tryGetPackage_Results) AuthorPgpKeyFingerprint() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_tryGetPackage_Results) HasAuthorPgpKeyFingerprint() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Results) AuthorPgpKeyFingerprintBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_tryGetPackage_Results) SetAuthorPgpKeyFingerprint(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// Backend_tryGetPackage_Results_List is a list of Backend_tryGetPackage_Results.
type Backend_tryGetPackage_Results_List struct{ capnp.List }

// NewBackend_tryGetPackage_Results creates a new list of Backend_tryGetPackage_Results.
func NewBackend_tryGetPackage_Results_List(s *capnp.Segment, sz int32) (Backend_tryGetPackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Backend_tryGetPackage_Results_List{l}, err
}

func (s Backend_tryGetPackage_Results_List) At(i int) Backend_tryGetPackage_Results {
	return Backend_tryGetPackage_Results{s.List.Struct(i)}
}

func (s Backend_tryGetPackage_Results_List) Set(i int, v Backend_tryGetPackage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_tryGetPackage_Results_Promise is a wrapper for a Backend_tryGetPackage_Results promised by a client call.
type Backend_tryGetPackage_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_tryGetPackage_Results_Promise) Struct() (Backend_tryGetPackage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_tryGetPackage_Results{s}, err
}

func (p Backend_tryGetPackage_Results_Promise) Manifest() spk.Manifest_Promise {
	return spk.Manifest_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Backend_deletePackage_Params struct{ capnp.Struct }

// Backend_deletePackage_Params_TypeID is the unique identifier for the type Backend_deletePackage_Params.
const Backend_deletePackage_Params_TypeID = 0xb61fc18674ca994f

func NewBackend_deletePackage_Params(s *capnp.Segment) (Backend_deletePackage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deletePackage_Params{st}, err
}

func NewRootBackend_deletePackage_Params(s *capnp.Segment) (Backend_deletePackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deletePackage_Params{st}, err
}

func ReadRootBackend_deletePackage_Params(msg *capnp.Message) (Backend_deletePackage_Params, error) {
	root, err := msg.RootPtr()
	return Backend_deletePackage_Params{root.Struct()}, err
}

func (s Backend_deletePackage_Params) String() string {
	str, _ := text.Marshal(0xb61fc18674ca994f, s.Struct)
	return str
}

func (s Backend_deletePackage_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_deletePackage_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deletePackage_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_deletePackage_Params) SetPackageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_deletePackage_Params_List is a list of Backend_deletePackage_Params.
type Backend_deletePackage_Params_List struct{ capnp.List }

// NewBackend_deletePackage_Params creates a new list of Backend_deletePackage_Params.
func NewBackend_deletePackage_Params_List(s *capnp.Segment, sz int32) (Backend_deletePackage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_deletePackage_Params_List{l}, err
}

func (s Backend_deletePackage_Params_List) At(i int) Backend_deletePackage_Params {
	return Backend_deletePackage_Params{s.List.Struct(i)}
}

func (s Backend_deletePackage_Params_List) Set(i int, v Backend_deletePackage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deletePackage_Params_Promise is a wrapper for a Backend_deletePackage_Params promised by a client call.
type Backend_deletePackage_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_deletePackage_Params_Promise) Struct() (Backend_deletePackage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deletePackage_Params{s}, err
}

type Backend_deletePackage_Results struct{ capnp.Struct }

// Backend_deletePackage_Results_TypeID is the unique identifier for the type Backend_deletePackage_Results.
const Backend_deletePackage_Results_TypeID = 0xea9f82a07e11b6d7

func NewBackend_deletePackage_Results(s *capnp.Segment) (Backend_deletePackage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deletePackage_Results{st}, err
}

func NewRootBackend_deletePackage_Results(s *capnp.Segment) (Backend_deletePackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deletePackage_Results{st}, err
}

func ReadRootBackend_deletePackage_Results(msg *capnp.Message) (Backend_deletePackage_Results, error) {
	root, err := msg.RootPtr()
	return Backend_deletePackage_Results{root.Struct()}, err
}

func (s Backend_deletePackage_Results) String() string {
	str, _ := text.Marshal(0xea9f82a07e11b6d7, s.Struct)
	return str
}

// Backend_deletePackage_Results_List is a list of Backend_deletePackage_Results.
type Backend_deletePackage_Results_List struct{ capnp.List }

// NewBackend_deletePackage_Results creates a new list of Backend_deletePackage_Results.
func NewBackend_deletePackage_Results_List(s *capnp.Segment, sz int32) (Backend_deletePackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_deletePackage_Results_List{l}, err
}

func (s Backend_deletePackage_Results_List) At(i int) Backend_deletePackage_Results {
	return Backend_deletePackage_Results{s.List.Struct(i)}
}

func (s Backend_deletePackage_Results_List) Set(i int, v Backend_deletePackage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deletePackage_Results_Promise is a wrapper for a Backend_deletePackage_Results promised by a client call.
type Backend_deletePackage_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_deletePackage_Results_Promise) Struct() (Backend_deletePackage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deletePackage_Results{s}, err
}

type Backend_backupGrain_Params struct{ capnp.Struct }

// Backend_backupGrain_Params_TypeID is the unique identifier for the type Backend_backupGrain_Params.
const Backend_backupGrain_Params_TypeID = 0x87a6a96b0a4edd21

func NewBackend_backupGrain_Params(s *capnp.Segment) (Backend_backupGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return Backend_backupGrain_Params{st}, err
}

func NewRootBackend_backupGrain_Params(s *capnp.Segment) (Backend_backupGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return Backend_backupGrain_Params{st}, err
}

func ReadRootBackend_backupGrain_Params(msg *capnp.Message) (Backend_backupGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_backupGrain_Params{root.Struct()}, err
}

func (s Backend_backupGrain_Params) String() string {
	str, _ := text.Marshal(0x87a6a96b0a4edd21, s.Struct)
	return str
}

func (s Backend_backupGrain_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_backupGrain_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_backupGrain_Params) SetBackupId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_backupGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_backupGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_backupGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Backend_backupGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_backupGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_backupGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s Backend_backupGrain_Params) Info() (grain.GrainInfo, error) {
	p, err := s.Struct.Ptr(3)
	return grain.GrainInfo{Struct: p.Struct()}, err
}

func (s Backend_backupGrain_Params) HasInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) SetInfo(v grain.GrainInfo) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewInfo sets the info field to a newly
// allocated grain.GrainInfo struct, preferring placement in s's segment.
func (s Backend_backupGrain_Params) NewInfo() (grain.GrainInfo, error) {
	ss, err := grain.NewGrainInfo(s.Struct.Segment())
	if err != nil {
		return grain.GrainInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// Backend_backupGrain_Params_List is a list of Backend_backupGrain_Params.
type Backend_backupGrain_Params_List struct{ capnp.List }

// NewBackend_backupGrain_Params creates a new list of Backend_backupGrain_Params.
func NewBackend_backupGrain_Params_List(s *capnp.Segment, sz int32) (Backend_backupGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return Backend_backupGrain_Params_List{l}, err
}

func (s Backend_backupGrain_Params_List) At(i int) Backend_backupGrain_Params {
	return Backend_backupGrain_Params{s.List.Struct(i)}
}

func (s Backend_backupGrain_Params_List) Set(i int, v Backend_backupGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_backupGrain_Params_Promise is a wrapper for a Backend_backupGrain_Params promised by a client call.
type Backend_backupGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_backupGrain_Params_Promise) Struct() (Backend_backupGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_backupGrain_Params{s}, err
}

func (p Backend_backupGrain_Params_Promise) Info() grain.GrainInfo_Promise {
	return grain.GrainInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type Backend_backupGrain_Results struct{ capnp.Struct }

// Backend_backupGrain_Results_TypeID is the unique identifier for the type Backend_backupGrain_Results.
const Backend_backupGrain_Results_TypeID = 0xcd9c9fab5f637827

func NewBackend_backupGrain_Results(s *capnp.Segment) (Backend_backupGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_backupGrain_Results{st}, err
}

func NewRootBackend_backupGrain_Results(s *capnp.Segment) (Backend_backupGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_backupGrain_Results{st}, err
}

func ReadRootBackend_backupGrain_Results(msg *capnp.Message) (Backend_backupGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_backupGrain_Results{root.Struct()}, err
}

func (s Backend_backupGrain_Results) String() string {
	str, _ := text.Marshal(0xcd9c9fab5f637827, s.Struct)
	return str
}

// Backend_backupGrain_Results_List is a list of Backend_backupGrain_Results.
type Backend_backupGrain_Results_List struct{ capnp.List }

// NewBackend_backupGrain_Results creates a new list of Backend_backupGrain_Results.
func NewBackend_backupGrain_Results_List(s *capnp.Segment, sz int32) (Backend_backupGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_backupGrain_Results_List{l}, err
}

func (s Backend_backupGrain_Results_List) At(i int) Backend_backupGrain_Results {
	return Backend_backupGrain_Results{s.List.Struct(i)}
}

func (s Backend_backupGrain_Results_List) Set(i int, v Backend_backupGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_backupGrain_Results_Promise is a wrapper for a Backend_backupGrain_Results promised by a client call.
type Backend_backupGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_backupGrain_Results_Promise) Struct() (Backend_backupGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_backupGrain_Results{s}, err
}

type Backend_restoreGrain_Params struct{ capnp.Struct }

// Backend_restoreGrain_Params_TypeID is the unique identifier for the type Backend_restoreGrain_Params.
const Backend_restoreGrain_Params_TypeID = 0x9d88f29f0318d4bb

func NewBackend_restoreGrain_Params(s *capnp.Segment) (Backend_restoreGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_restoreGrain_Params{st}, err
}

func NewRootBackend_restoreGrain_Params(s *capnp.Segment) (Backend_restoreGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_restoreGrain_Params{st}, err
}

func ReadRootBackend_restoreGrain_Params(msg *capnp.Message) (Backend_restoreGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_restoreGrain_Params{root.Struct()}, err
}

func (s Backend_restoreGrain_Params) String() string {
	str, _ := text.Marshal(0x9d88f29f0318d4bb, s.Struct)
	return str
}

func (s Backend_restoreGrain_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_restoreGrain_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_restoreGrain_Params) SetBackupId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_restoreGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_restoreGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_restoreGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Backend_restoreGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_restoreGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_restoreGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// Backend_restoreGrain_Params_List is a list of Backend_restoreGrain_Params.
type Backend_restoreGrain_Params_List struct{ capnp.List }

// NewBackend_restoreGrain_Params creates a new list of Backend_restoreGrain_Params.
func NewBackend_restoreGrain_Params_List(s *capnp.Segment, sz int32) (Backend_restoreGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Backend_restoreGrain_Params_List{l}, err
}

func (s Backend_restoreGrain_Params_List) At(i int) Backend_restoreGrain_Params {
	return Backend_restoreGrain_Params{s.List.Struct(i)}
}

func (s Backend_restoreGrain_Params_List) Set(i int, v Backend_restoreGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_restoreGrain_Params_Promise is a wrapper for a Backend_restoreGrain_Params promised by a client call.
type Backend_restoreGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_restoreGrain_Params_Promise) Struct() (Backend_restoreGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_restoreGrain_Params{s}, err
}

type Backend_restoreGrain_Results struct{ capnp.Struct }

// Backend_restoreGrain_Results_TypeID is the unique identifier for the type Backend_restoreGrain_Results.
const Backend_restoreGrain_Results_TypeID = 0x8b790707193ea7ff

func NewBackend_restoreGrain_Results(s *capnp.Segment) (Backend_restoreGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_restoreGrain_Results{st}, err
}

func NewRootBackend_restoreGrain_Results(s *capnp.Segment) (Backend_restoreGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_restoreGrain_Results{st}, err
}

func ReadRootBackend_restoreGrain_Results(msg *capnp.Message) (Backend_restoreGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_restoreGrain_Results{root.Struct()}, err
}

func (s Backend_restoreGrain_Results) String() string {
	str, _ := text.Marshal(0x8b790707193ea7ff, s.Struct)
	return str
}

func (s Backend_restoreGrain_Results) Info() (grain.GrainInfo, error) {
	p, err := s.Struct.Ptr(0)
	return grain.GrainInfo{Struct: p.Struct()}, err
}

func (s Backend_restoreGrain_Results) HasInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Results) SetInfo(v grain.GrainInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewInfo sets the info field to a newly
// allocated grain.GrainInfo struct, preferring placement in s's segment.
func (s Backend_restoreGrain_Results) NewInfo() (grain.GrainInfo, error) {
	ss, err := grain.NewGrainInfo(s.Struct.Segment())
	if err != nil {
		return grain.GrainInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Backend_restoreGrain_Results_List is a list of Backend_restoreGrain_Results.
type Backend_restoreGrain_Results_List struct{ capnp.List }

// NewBackend_restoreGrain_Results creates a new list of Backend_restoreGrain_Results.
func NewBackend_restoreGrain_Results_List(s *capnp.Segment, sz int32) (Backend_restoreGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_restoreGrain_Results_List{l}, err
}

func (s Backend_restoreGrain_Results_List) At(i int) Backend_restoreGrain_Results {
	return Backend_restoreGrain_Results{s.List.Struct(i)}
}

func (s Backend_restoreGrain_Results_List) Set(i int, v Backend_restoreGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_restoreGrain_Results_Promise is a wrapper for a Backend_restoreGrain_Results promised by a client call.
type Backend_restoreGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_restoreGrain_Results_Promise) Struct() (Backend_restoreGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_restoreGrain_Results{s}, err
}

func (p Backend_restoreGrain_Results_Promise) Info() grain.GrainInfo_Promise {
	return grain.GrainInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Backend_uploadBackup_Params struct{ capnp.Struct }

// Backend_uploadBackup_Params_TypeID is the unique identifier for the type Backend_uploadBackup_Params.
const Backend_uploadBackup_Params_TypeID = 0xf2ccecff0178227b

func NewBackend_uploadBackup_Params(s *capnp.Segment) (Backend_uploadBackup_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_uploadBackup_Params{st}, err
}

func NewRootBackend_uploadBackup_Params(s *capnp.Segment) (Backend_uploadBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_uploadBackup_Params{st}, err
}

func ReadRootBackend_uploadBackup_Params(msg *capnp.Message) (Backend_uploadBackup_Params, error) {
	root, err := msg.RootPtr()
	return Backend_uploadBackup_Params{root.Struct()}, err
}

func (s Backend_uploadBackup_Params) String() string {
	str, _ := text.Marshal(0xf2ccecff0178227b, s.Struct)
	return str
}

func (s Backend_uploadBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_uploadBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_uploadBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_uploadBackup_Params) SetBackupId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_uploadBackup_Params_List is a list of Backend_uploadBackup_Params.
type Backend_uploadBackup_Params_List struct{ capnp.List }

// NewBackend_uploadBackup_Params creates a new list of Backend_uploadBackup_Params.
func NewBackend_uploadBackup_Params_List(s *capnp.Segment, sz int32) (Backend_uploadBackup_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_uploadBackup_Params_List{l}, err
}

func (s Backend_uploadBackup_Params_List) At(i int) Backend_uploadBackup_Params {
	return Backend_uploadBackup_Params{s.List.Struct(i)}
}

func (s Backend_uploadBackup_Params_List) Set(i int, v Backend_uploadBackup_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_uploadBackup_Params_Promise is a wrapper for a Backend_uploadBackup_Params promised by a client call.
type Backend_uploadBackup_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_uploadBackup_Params_Promise) Struct() (Backend_uploadBackup_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_uploadBackup_Params{s}, err
}

type Backend_uploadBackup_Results struct{ capnp.Struct }

// Backend_uploadBackup_Results_TypeID is the unique identifier for the type Backend_uploadBackup_Results.
const Backend_uploadBackup_Results_TypeID = 0xbc51d6bc865a8fcf

func NewBackend_uploadBackup_Results(s *capnp.Segment) (Backend_uploadBackup_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_uploadBackup_Results{st}, err
}

func NewRootBackend_uploadBackup_Results(s *capnp.Segment) (Backend_uploadBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_uploadBackup_Results{st}, err
}

func ReadRootBackend_uploadBackup_Results(msg *capnp.Message) (Backend_uploadBackup_Results, error) {
	root, err := msg.RootPtr()
	return Backend_uploadBackup_Results{root.Struct()}, err
}

func (s Backend_uploadBackup_Results) String() string {
	str, _ := text.Marshal(0xbc51d6bc865a8fcf, s.Struct)
	return str
}

func (s Backend_uploadBackup_Results) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(0)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Backend_uploadBackup_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_uploadBackup_Results) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Backend_uploadBackup_Results_List is a list of Backend_uploadBackup_Results.
type Backend_uploadBackup_Results_List struct{ capnp.List }

// NewBackend_uploadBackup_Results creates a new list of Backend_uploadBackup_Results.
func NewBackend_uploadBackup_Results_List(s *capnp.Segment, sz int32) (Backend_uploadBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_uploadBackup_Results_List{l}, err
}

func (s Backend_uploadBackup_Results_List) At(i int) Backend_uploadBackup_Results {
	return Backend_uploadBackup_Results{s.List.Struct(i)}
}

func (s Backend_uploadBackup_Results_List) Set(i int, v Backend_uploadBackup_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_uploadBackup_Results_Promise is a wrapper for a Backend_uploadBackup_Results promised by a client call.
type Backend_uploadBackup_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_uploadBackup_Results_Promise) Struct() (Backend_uploadBackup_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_uploadBackup_Results{s}, err
}

func (p Backend_uploadBackup_Results_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Backend_downloadBackup_Params struct{ capnp.Struct }

// Backend_downloadBackup_Params_TypeID is the unique identifier for the type Backend_downloadBackup_Params.
const Backend_downloadBackup_Params_TypeID = 0x916d32f140971035

func NewBackend_downloadBackup_Params(s *capnp.Segment) (Backend_downloadBackup_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_downloadBackup_Params{st}, err
}

func NewRootBackend_downloadBackup_Params(s *capnp.Segment) (Backend_downloadBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_downloadBackup_Params{st}, err
}

func ReadRootBackend_downloadBackup_Params(msg *capnp.Message) (Backend_downloadBackup_Params, error) {
	root, err := msg.RootPtr()
	return Backend_downloadBackup_Params{root.Struct()}, err
}

func (s Backend_downloadBackup_Params) String() string {
	str, _ := text.Marshal(0x916d32f140971035, s.Struct)
	return str
}

func (s Backend_downloadBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_downloadBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_downloadBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_downloadBackup_Params) SetBackupId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_downloadBackup_Params) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(1)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Backend_downloadBackup_Params) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_downloadBackup_Params) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// Backend_downloadBackup_Params_List is a list of Backend_downloadBackup_Params.
type Backend_downloadBackup_Params_List struct{ capnp.List }

// NewBackend_downloadBackup_Params creates a new list of Backend_downloadBackup_Params.
func NewBackend_downloadBackup_Params_List(s *capnp.Segment, sz int32) (Backend_downloadBackup_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Backend_downloadBackup_Params_List{l}, err
}

func (s Backend_downloadBackup_Params_List) At(i int) Backend_downloadBackup_Params {
	return Backend_downloadBackup_Params{s.List.Struct(i)}
}

func (s Backend_downloadBackup_Params_List) Set(i int, v Backend_downloadBackup_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_downloadBackup_Params_Promise is a wrapper for a Backend_downloadBackup_Params promised by a client call.
type Backend_downloadBackup_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_downloadBackup_Params_Promise) Struct() (Backend_downloadBackup_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_downloadBackup_Params{s}, err
}

func (p Backend_downloadBackup_Params_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

type Backend_downloadBackup_Results struct{ capnp.Struct }

// Backend_downloadBackup_Results_TypeID is the unique identifier for the type Backend_downloadBackup_Results.
const Backend_downloadBackup_Results_TypeID = 0x9e90498484bab87d

func NewBackend_downloadBackup_Results(s *capnp.Segment) (Backend_downloadBackup_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_downloadBackup_Results{st}, err
}

func NewRootBackend_downloadBackup_Results(s *capnp.Segment) (Backend_downloadBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_downloadBackup_Results{st}, err
}

func ReadRootBackend_downloadBackup_Results(msg *capnp.Message) (Backend_downloadBackup_Results, error) {
	root, err := msg.RootPtr()
	return Backend_downloadBackup_Results{root.Struct()}, err
}

func (s Backend_downloadBackup_Results) String() string {
	str, _ := text.Marshal(0x9e90498484bab87d, s.Struct)
	return str
}

// Backend_downloadBackup_Results_List is a list of Backend_downloadBackup_Results.
type Backend_downloadBackup_Results_List struct{ capnp.List }

// NewBackend_downloadBackup_Results creates a new list of Backend_downloadBackup_Results.
func NewBackend_downloadBackup_Results_List(s *capnp.Segment, sz int32) (Backend_downloadBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_downloadBackup_Results_List{l}, err
}

func (s Backend_downloadBackup_Results_List) At(i int) Backend_downloadBackup_Results {
	return Backend_downloadBackup_Results{s.List.Struct(i)}
}

func (s Backend_downloadBackup_Results_List) Set(i int, v Backend_downloadBackup_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_downloadBackup_Results_Promise is a wrapper for a Backend_downloadBackup_Results promised by a client call.
type Backend_downloadBackup_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_downloadBackup_Results_Promise) Struct() (Backend_downloadBackup_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_downloadBackup_Results{s}, err
}

type Backend_deleteBackup_Params struct{ capnp.Struct }

// Backend_deleteBackup_Params_TypeID is the unique identifier for the type Backend_deleteBackup_Params.
const Backend_deleteBackup_Params_TypeID = 0xd0d6ed6a5ed70e62

func NewBackend_deleteBackup_Params(s *capnp.Segment) (Backend_deleteBackup_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deleteBackup_Params{st}, err
}

func NewRootBackend_deleteBackup_Params(s *capnp.Segment) (Backend_deleteBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deleteBackup_Params{st}, err
}

func ReadRootBackend_deleteBackup_Params(msg *capnp.Message) (Backend_deleteBackup_Params, error) {
	root, err := msg.RootPtr()
	return Backend_deleteBackup_Params{root.Struct()}, err
}

func (s Backend_deleteBackup_Params) String() string {
	str, _ := text.Marshal(0xd0d6ed6a5ed70e62, s.Struct)
	return str
}

func (s Backend_deleteBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_deleteBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deleteBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_deleteBackup_Params) SetBackupId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_deleteBackup_Params_List is a list of Backend_deleteBackup_Params.
type Backend_deleteBackup_Params_List struct{ capnp.List }

// NewBackend_deleteBackup_Params creates a new list of Backend_deleteBackup_Params.
func NewBackend_deleteBackup_Params_List(s *capnp.Segment, sz int32) (Backend_deleteBackup_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_deleteBackup_Params_List{l}, err
}

func (s Backend_deleteBackup_Params_List) At(i int) Backend_deleteBackup_Params {
	return Backend_deleteBackup_Params{s.List.Struct(i)}
}

func (s Backend_deleteBackup_Params_List) Set(i int, v Backend_deleteBackup_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteBackup_Params_Promise is a wrapper for a Backend_deleteBackup_Params promised by a client call.
type Backend_deleteBackup_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteBackup_Params_Promise) Struct() (Backend_deleteBackup_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteBackup_Params{s}, err
}

type Backend_deleteBackup_Results struct{ capnp.Struct }

// Backend_deleteBackup_Results_TypeID is the unique identifier for the type Backend_deleteBackup_Results.
const Backend_deleteBackup_Results_TypeID = 0xaf88ad00c801b00d

func NewBackend_deleteBackup_Results(s *capnp.Segment) (Backend_deleteBackup_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteBackup_Results{st}, err
}

func NewRootBackend_deleteBackup_Results(s *capnp.Segment) (Backend_deleteBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteBackup_Results{st}, err
}

func ReadRootBackend_deleteBackup_Results(msg *capnp.Message) (Backend_deleteBackup_Results, error) {
	root, err := msg.RootPtr()
	return Backend_deleteBackup_Results{root.Struct()}, err
}

func (s Backend_deleteBackup_Results) String() string {
	str, _ := text.Marshal(0xaf88ad00c801b00d, s.Struct)
	return str
}

// Backend_deleteBackup_Results_List is a list of Backend_deleteBackup_Results.
type Backend_deleteBackup_Results_List struct{ capnp.List }

// NewBackend_deleteBackup_Results creates a new list of Backend_deleteBackup_Results.
func NewBackend_deleteBackup_Results_List(s *capnp.Segment, sz int32) (Backend_deleteBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_deleteBackup_Results_List{l}, err
}

func (s Backend_deleteBackup_Results_List) At(i int) Backend_deleteBackup_Results {
	return Backend_deleteBackup_Results{s.List.Struct(i)}
}

func (s Backend_deleteBackup_Results_List) Set(i int, v Backend_deleteBackup_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteBackup_Results_Promise is a wrapper for a Backend_deleteBackup_Results promised by a client call.
type Backend_deleteBackup_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteBackup_Results_Promise) Struct() (Backend_deleteBackup_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteBackup_Results{s}, err
}

type Backend_getUserStorageUsage_Params struct{ capnp.Struct }

// Backend_getUserStorageUsage_Params_TypeID is the unique identifier for the type Backend_getUserStorageUsage_Params.
const Backend_getUserStorageUsage_Params_TypeID = 0xaaef1f8c301b865d

func NewBackend_getUserStorageUsage_Params(s *capnp.Segment) (Backend_getUserStorageUsage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_getUserStorageUsage_Params{st}, err
}

func NewRootBackend_getUserStorageUsage_Params(s *capnp.Segment) (Backend_getUserStorageUsage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_getUserStorageUsage_Params{st}, err
}

func ReadRootBackend_getUserStorageUsage_Params(msg *capnp.Message) (Backend_getUserStorageUsage_Params, error) {
	root, err := msg.RootPtr()
	return Backend_getUserStorageUsage_Params{root.Struct()}, err
}

func (s Backend_getUserStorageUsage_Params) String() string {
	str, _ := text.Marshal(0xaaef1f8c301b865d, s.Struct)
	return str
}

func (s Backend_getUserStorageUsage_Params) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_getUserStorageUsage_Params) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getUserStorageUsage_Params) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_getUserStorageUsage_Params) SetUserId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_getUserStorageUsage_Params_List is a list of Backend_getUserStorageUsage_Params.
type Backend_getUserStorageUsage_Params_List struct{ capnp.List }

// NewBackend_getUserStorageUsage_Params creates a new list of Backend_getUserStorageUsage_Params.
func NewBackend_getUserStorageUsage_Params_List(s *capnp.Segment, sz int32) (Backend_getUserStorageUsage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_getUserStorageUsage_Params_List{l}, err
}

func (s Backend_getUserStorageUsage_Params_List) At(i int) Backend_getUserStorageUsage_Params {
	return Backend_getUserStorageUsage_Params{s.List.Struct(i)}
}

func (s Backend_getUserStorageUsage_Params_List) Set(i int, v Backend_getUserStorageUsage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getUserStorageUsage_Params_Promise is a wrapper for a Backend_getUserStorageUsage_Params promised by a client call.
type Backend_getUserStorageUsage_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_getUserStorageUsage_Params_Promise) Struct() (Backend_getUserStorageUsage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getUserStorageUsage_Params{s}, err
}

type Backend_getUserStorageUsage_Results struct{ capnp.Struct }

// Backend_getUserStorageUsage_Results_TypeID is the unique identifier for the type Backend_getUserStorageUsage_Results.
const Backend_getUserStorageUsage_Results_TypeID = 0xa1c73384bc38ab4b

func NewBackend_getUserStorageUsage_Results(s *capnp.Segment) (Backend_getUserStorageUsage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Backend_getUserStorageUsage_Results{st}, err
}

func NewRootBackend_getUserStorageUsage_Results(s *capnp.Segment) (Backend_getUserStorageUsage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Backend_getUserStorageUsage_Results{st}, err
}

func ReadRootBackend_getUserStorageUsage_Results(msg *capnp.Message) (Backend_getUserStorageUsage_Results, error) {
	root, err := msg.RootPtr()
	return Backend_getUserStorageUsage_Results{root.Struct()}, err
}

func (s Backend_getUserStorageUsage_Results) String() string {
	str, _ := text.Marshal(0xa1c73384bc38ab4b, s.Struct)
	return str
}

func (s Backend_getUserStorageUsage_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Backend_getUserStorageUsage_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Backend_getUserStorageUsage_Results_List is a list of Backend_getUserStorageUsage_Results.
type Backend_getUserStorageUsage_Results_List struct{ capnp.List }

// NewBackend_getUserStorageUsage_Results creates a new list of Backend_getUserStorageUsage_Results.
func NewBackend_getUserStorageUsage_Results_List(s *capnp.Segment, sz int32) (Backend_getUserStorageUsage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Backend_getUserStorageUsage_Results_List{l}, err
}

func (s Backend_getUserStorageUsage_Results_List) At(i int) Backend_getUserStorageUsage_Results {
	return Backend_getUserStorageUsage_Results{s.List.Struct(i)}
}

func (s Backend_getUserStorageUsage_Results_List) Set(i int, v Backend_getUserStorageUsage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getUserStorageUsage_Results_Promise is a wrapper for a Backend_getUserStorageUsage_Results promised by a client call.
type Backend_getUserStorageUsage_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_getUserStorageUsage_Results_Promise) Struct() (Backend_getUserStorageUsage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getUserStorageUsage_Results{s}, err
}

type Backend_transferGrain_Params struct{ capnp.Struct }

// Backend_transferGrain_Params_TypeID is the unique identifier for the type Backend_transferGrain_Params.
const Backend_transferGrain_Params_TypeID = 0xcce40aee6005d381

func NewBackend_transferGrain_Params(s *capnp.Segment) (Backend_transferGrain_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_transferGrain_Params{st}, err
}

func NewRootBackend_transferGrain_Params(s *capnp.Segment) (Backend_transferGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Backend_transferGrain_Params{st}, err
}

func ReadRootBackend_transferGrain_Params(msg *capnp.Message) (Backend_transferGrain_Params, error) {
	root, err := msg.RootPtr()
	return Backend_transferGrain_Params{root.Struct()}, err
}

func (s Backend_transferGrain_Params) String() string {
	str, _ := text.Marshal(0xcce40aee6005d381, s.Struct)
	return str
}

func (s Backend_transferGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_transferGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_transferGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_transferGrain_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_transferGrain_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_transferGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_transferGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_transferGrain_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Backend_transferGrain_Params) NewOwnerId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Backend_transferGrain_Params) HasNewOwnerId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_transferGrain_Params) NewOwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Backend_transferGrain_Params) SetNewOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// Backend_transferGrain_Params_List is a list of Backend_transferGrain_Params.
type Backend_transferGrain_Params_List struct{ capnp.List }

// NewBackend_transferGrain_Params creates a new list of Backend_transferGrain_Params.
func NewBackend_transferGrain_Params_List(s *capnp.Segment, sz int32) (Backend_transferGrain_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Backend_transferGrain_Params_List{l}, err
}

func (s Backend_transferGrain_Params_List) At(i int) Backend_transferGrain_Params {
	return Backend_transferGrain_Params{s.List.Struct(i)}
}

func (s Backend_transferGrain_Params_List) Set(i int, v Backend_transferGrain_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_transferGrain_Params_Promise is a wrapper for a Backend_transferGrain_Params promised by a client call.
type Backend_transferGrain_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_transferGrain_Params_Promise) Struct() (Backend_transferGrain_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_transferGrain_Params{s}, err
}

type Backend_transferGrain_Results struct{ capnp.Struct }

// Backend_transferGrain_Results_TypeID is the unique identifier for the type Backend_transferGrain_Results.
const Backend_transferGrain_Results_TypeID = 0x86362c69f5c42997

func NewBackend_transferGrain_Results(s *capnp.Segment) (Backend_transferGrain_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_transferGrain_Results{st}, err
}

func NewRootBackend_transferGrain_Results(s *capnp.Segment) (Backend_transferGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_transferGrain_Results{st}, err
}

func ReadRootBackend_transferGrain_Results(msg *capnp.Message) (Backend_transferGrain_Results, error) {
	root, err := msg.RootPtr()
	return Backend_transferGrain_Results{root.Struct()}, err
}

func (s Backend_transferGrain_Results) String() string {
	str, _ := text.Marshal(0x86362c69f5c42997, s.Struct)
	return str
}

// Backend_transferGrain_Results_List is a list of Backend_transferGrain_Results.
type Backend_transferGrain_Results_List struct{ capnp.List }

// NewBackend_transferGrain_Results creates a new list of Backend_transferGrain_Results.
func NewBackend_transferGrain_Results_List(s *capnp.Segment, sz int32) (Backend_transferGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_transferGrain_Results_List{l}, err
}

func (s Backend_transferGrain_Results_List) At(i int) Backend_transferGrain_Results {
	return Backend_transferGrain_Results{s.List.Struct(i)}
}

func (s Backend_transferGrain_Results_List) Set(i int, v Backend_transferGrain_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_transferGrain_Results_Promise is a wrapper for a Backend_transferGrain_Results promised by a client call.
type Backend_transferGrain_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_transferGrain_Results_Promise) Struct() (Backend_transferGrain_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_transferGrain_Results{s}, err
}

type Backend_deleteUser_Params struct{ capnp.Struct }

// Backend_deleteUser_Params_TypeID is the unique identifier for the type Backend_deleteUser_Params.
const Backend_deleteUser_Params_TypeID = 0xfa7238e0a9345914

func NewBackend_deleteUser_Params(s *capnp.Segment) (Backend_deleteUser_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deleteUser_Params{st}, err
}

func NewRootBackend_deleteUser_Params(s *capnp.Segment) (Backend_deleteUser_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Backend_deleteUser_Params{st}, err
}

func ReadRootBackend_deleteUser_Params(msg *capnp.Message) (Backend_deleteUser_Params, error) {
	root, err := msg.RootPtr()
	return Backend_deleteUser_Params{root.Struct()}, err
}

func (s Backend_deleteUser_Params) String() string {
	str, _ := text.Marshal(0xfa7238e0a9345914, s.Struct)
	return str
}

func (s Backend_deleteUser_Params) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_deleteUser_Params) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deleteUser_Params) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_deleteUser_Params) SetUserId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Backend_deleteUser_Params_List is a list of Backend_deleteUser_Params.
type Backend_deleteUser_Params_List struct{ capnp.List }

// NewBackend_deleteUser_Params creates a new list of Backend_deleteUser_Params.
func NewBackend_deleteUser_Params_List(s *capnp.Segment, sz int32) (Backend_deleteUser_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Backend_deleteUser_Params_List{l}, err
}

func (s Backend_deleteUser_Params_List) At(i int) Backend_deleteUser_Params {
	return Backend_deleteUser_Params{s.List.Struct(i)}
}

func (s Backend_deleteUser_Params_List) Set(i int, v Backend_deleteUser_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteUser_Params_Promise is a wrapper for a Backend_deleteUser_Params promised by a client call.
type Backend_deleteUser_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteUser_Params_Promise) Struct() (Backend_deleteUser_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteUser_Params{s}, err
}

type Backend_deleteUser_Results struct{ capnp.Struct }

// Backend_deleteUser_Results_TypeID is the unique identifier for the type Backend_deleteUser_Results.
const Backend_deleteUser_Results_TypeID = 0x9145c7ea308343d9

func NewBackend_deleteUser_Results(s *capnp.Segment) (Backend_deleteUser_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteUser_Results{st}, err
}

func NewRootBackend_deleteUser_Results(s *capnp.Segment) (Backend_deleteUser_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_deleteUser_Results{st}, err
}

func ReadRootBackend_deleteUser_Results(msg *capnp.Message) (Backend_deleteUser_Results, error) {
	root, err := msg.RootPtr()
	return Backend_deleteUser_Results{root.Struct()}, err
}

func (s Backend_deleteUser_Results) String() string {
	str, _ := text.Marshal(0x9145c7ea308343d9, s.Struct)
	return str
}

// Backend_deleteUser_Results_List is a list of Backend_deleteUser_Results.
type Backend_deleteUser_Results_List struct{ capnp.List }

// NewBackend_deleteUser_Results creates a new list of Backend_deleteUser_Results.
func NewBackend_deleteUser_Results_List(s *capnp.Segment, sz int32) (Backend_deleteUser_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_deleteUser_Results_List{l}, err
}

func (s Backend_deleteUser_Results_List) At(i int) Backend_deleteUser_Results {
	return Backend_deleteUser_Results{s.List.Struct(i)}
}

func (s Backend_deleteUser_Results_List) Set(i int, v Backend_deleteUser_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_deleteUser_Results_Promise is a wrapper for a Backend_deleteUser_Results promised by a client call.
type Backend_deleteUser_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_deleteUser_Results_Promise) Struct() (Backend_deleteUser_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_deleteUser_Results{s}, err
}

type Backend_ping_Params struct{ capnp.Struct }

// Backend_ping_Params_TypeID is the unique identifier for the type Backend_ping_Params.
const Backend_ping_Params_TypeID = 0xcb56f444d1311800

func NewBackend_ping_Params(s *capnp.Segment) (Backend_ping_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_ping_Params{st}, err
}

func NewRootBackend_ping_Params(s *capnp.Segment) (Backend_ping_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_ping_Params{st}, err
}

func ReadRootBackend_ping_Params(msg *capnp.Message) (Backend_ping_Params, error) {
	root, err := msg.RootPtr()
	return Backend_ping_Params{root.Struct()}, err
}

func (s Backend_ping_Params) String() string {
	str, _ := text.Marshal(0xcb56f444d1311800, s.Struct)
	return str
}

// Backend_ping_Params_List is a list of Backend_ping_Params.
type Backend_ping_Params_List struct{ capnp.List }

// NewBackend_ping_Params creates a new list of Backend_ping_Params.
func NewBackend_ping_Params_List(s *capnp.Segment, sz int32) (Backend_ping_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_ping_Params_List{l}, err
}

func (s Backend_ping_Params_List) At(i int) Backend_ping_Params {
	return Backend_ping_Params{s.List.Struct(i)}
}

func (s Backend_ping_Params_List) Set(i int, v Backend_ping_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_ping_Params_Promise is a wrapper for a Backend_ping_Params promised by a client call.
type Backend_ping_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_ping_Params_Promise) Struct() (Backend_ping_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_ping_Params{s}, err
}

type Backend_ping_Results struct{ capnp.Struct }

// Backend_ping_Results_TypeID is the unique identifier for the type Backend_ping_Results.
const Backend_ping_Results_TypeID = 0xe3a9cebde9177d60

func NewBackend_ping_Results(s *capnp.Segment) (Backend_ping_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_ping_Results{st}, err
}

func NewRootBackend_ping_Results(s *capnp.Segment) (Backend_ping_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Backend_ping_Results{st}, err
}

func ReadRootBackend_ping_Results(msg *capnp.Message) (Backend_ping_Results, error) {
	root, err := msg.RootPtr()
	return Backend_ping_Results{root.Struct()}, err
}

func (s Backend_ping_Results) String() string {
	str, _ := text.Marshal(0xe3a9cebde9177d60, s.Struct)
	return str
}

// Backend_ping_Results_List is a list of Backend_ping_Results.
type Backend_ping_Results_List struct{ capnp.List }

// NewBackend_ping_Results creates a new list of Backend_ping_Results.
func NewBackend_ping_Results_List(s *capnp.Segment, sz int32) (Backend_ping_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Backend_ping_Results_List{l}, err
}

func (s Backend_ping_Results_List) At(i int) Backend_ping_Results {
	return Backend_ping_Results{s.List.Struct(i)}
}

func (s Backend_ping_Results_List) Set(i int, v Backend_ping_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_ping_Results_Promise is a wrapper for a Backend_ping_Results promised by a client call.
type Backend_ping_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_ping_Results_Promise) Struct() (Backend_ping_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_ping_Results{s}, err
}

type Backend_getGrainStorageUsage_Params struct{ capnp.Struct }

// Backend_getGrainStorageUsage_Params_TypeID is the unique identifier for the type Backend_getGrainStorageUsage_Params.
const Backend_getGrainStorageUsage_Params_TypeID = 0xe06fe4e0d4e93178

func NewBackend_getGrainStorageUsage_Params(s *capnp.Segment) (Backend_getGrainStorageUsage_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_getGrainStorageUsage_Params{st}, err
}

func NewRootBackend_getGrainStorageUsage_Params(s *capnp.Segment) (Backend_getGrainStorageUsage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Backend_getGrainStorageUsage_Params{st}, err
}

func ReadRootBackend_getGrainStorageUsage_Params(msg *capnp.Message) (Backend_getGrainStorageUsage_Params, error) {
	root, err := msg.RootPtr()
	return Backend_getGrainStorageUsage_Params{root.Struct()}, err
}

func (s Backend_getGrainStorageUsage_Params) String() string {
	str, _ := text.Marshal(0xe06fe4e0d4e93178, s.Struct)
	return str
}

func (s Backend_getGrainStorageUsage_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Backend_getGrainStorageUsage_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getGrainStorageUsage_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Backend_getGrainStorageUsage_Params) SetOwnerId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_getGrainStorageUsage_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Backend_getGrainStorageUsage_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_getGrainStorageUsage_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Backend_getGrainStorageUsage_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// Backend_getGrainStorageUsage_Params_List is a list of Backend_getGrainStorageUsage_Params.
type Backend_getGrainStorageUsage_Params_List struct{ capnp.List }

// NewBackend_getGrainStorageUsage_Params creates a new list of Backend_getGrainStorageUsage_Params.
func NewBackend_getGrainStorageUsage_Params_List(s *capnp.Segment, sz int32) (Backend_getGrainStorageUsage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Backend_getGrainStorageUsage_Params_List{l}, err
}

func (s Backend_getGrainStorageUsage_Params_List) At(i int) Backend_getGrainStorageUsage_Params {
	return Backend_getGrainStorageUsage_Params{s.List.Struct(i)}
}

func (s Backend_getGrainStorageUsage_Params_List) Set(i int, v Backend_getGrainStorageUsage_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getGrainStorageUsage_Params_Promise is a wrapper for a Backend_getGrainStorageUsage_Params promised by a client call.
type Backend_getGrainStorageUsage_Params_Promise struct{ *capnp.Pipeline }

func (p Backend_getGrainStorageUsage_Params_Promise) Struct() (Backend_getGrainStorageUsage_Params, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getGrainStorageUsage_Params{s}, err
}

type Backend_getGrainStorageUsage_Results struct{ capnp.Struct }

// Backend_getGrainStorageUsage_Results_TypeID is the unique identifier for the type Backend_getGrainStorageUsage_Results.
const Backend_getGrainStorageUsage_Results_TypeID = 0x809d3d6d45c4c37d

func NewBackend_getGrainStorageUsage_Results(s *capnp.Segment) (Backend_getGrainStorageUsage_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Backend_getGrainStorageUsage_Results{st}, err
}

func NewRootBackend_getGrainStorageUsage_Results(s *capnp.Segment) (Backend_getGrainStorageUsage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Backend_getGrainStorageUsage_Results{st}, err
}

func ReadRootBackend_getGrainStorageUsage_Results(msg *capnp.Message) (Backend_getGrainStorageUsage_Results, error) {
	root, err := msg.RootPtr()
	return Backend_getGrainStorageUsage_Results{root.Struct()}, err
}

func (s Backend_getGrainStorageUsage_Results) String() string {
	str, _ := text.Marshal(0x809d3d6d45c4c37d, s.Struct)
	return str
}

func (s Backend_getGrainStorageUsage_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Backend_getGrainStorageUsage_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Backend_getGrainStorageUsage_Results_List is a list of Backend_getGrainStorageUsage_Results.
type Backend_getGrainStorageUsage_Results_List struct{ capnp.List }

// NewBackend_getGrainStorageUsage_Results creates a new list of Backend_getGrainStorageUsage_Results.
func NewBackend_getGrainStorageUsage_Results_List(s *capnp.Segment, sz int32) (Backend_getGrainStorageUsage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Backend_getGrainStorageUsage_Results_List{l}, err
}

func (s Backend_getGrainStorageUsage_Results_List) At(i int) Backend_getGrainStorageUsage_Results {
	return Backend_getGrainStorageUsage_Results{s.List.Struct(i)}
}

func (s Backend_getGrainStorageUsage_Results_List) Set(i int, v Backend_getGrainStorageUsage_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Backend_getGrainStorageUsage_Results_Promise is a wrapper for a Backend_getGrainStorageUsage_Results promised by a client call.
type Backend_getGrainStorageUsage_Results_Promise struct{ *capnp.Pipeline }

func (p Backend_getGrainStorageUsage_Results_Promise) Struct() (Backend_getGrainStorageUsage_Results, error) {
	s, err := p.Pipeline.Struct()
	return Backend_getGrainStorageUsage_Results{s}, err
}

type SandstormCoreFactory struct{ Client capnp.Client }

func (c SandstormCoreFactory) GetSandstormCore(ctx context.Context, params func(SandstormCoreFactory_getSandstormCore_Params) error, opts ...capnp.CallOption) SandstormCoreFactory_getSandstormCore_Results_Promise {
	if c.Client == nil {
		return SandstormCoreFactory_getSandstormCore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf0832c3f66256d2b,
			MethodID:      0,
			InterfaceName: "backend.capnp:SandstormCoreFactory",
			MethodName:    "getSandstormCore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormCoreFactory_getSandstormCore_Params{Struct: s}) }
	}
	return SandstormCoreFactory_getSandstormCore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormCoreFactory_Server interface {
	GetSandstormCore(SandstormCoreFactory_getSandstormCore) error
}

func SandstormCoreFactory_ServerToClient(s SandstormCoreFactory_Server) SandstormCoreFactory {
	c, _ := s.(server.Closer)
	return SandstormCoreFactory{Client: server.New(SandstormCoreFactory_Methods(nil, s), c)}
}

func SandstormCoreFactory_Methods(methods []server.Method, s SandstormCoreFactory_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf0832c3f66256d2b,
			MethodID:      0,
			InterfaceName: "backend.capnp:SandstormCoreFactory",
			MethodName:    "getSandstormCore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormCoreFactory_getSandstormCore{c, opts, SandstormCoreFactory_getSandstormCore_Params{Struct: p}, SandstormCoreFactory_getSandstormCore_Results{Struct: r}}
			return s.GetSandstormCore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SandstormCoreFactory_getSandstormCore holds the arguments for a server call to SandstormCoreFactory.getSandstormCore.
type SandstormCoreFactory_getSandstormCore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormCoreFactory_getSandstormCore_Params
	Results SandstormCoreFactory_getSandstormCore_Results
}

type SandstormCoreFactory_getSandstormCore_Params struct{ capnp.Struct }

// SandstormCoreFactory_getSandstormCore_Params_TypeID is the unique identifier for the type SandstormCoreFactory_getSandstormCore_Params.
const SandstormCoreFactory_getSandstormCore_Params_TypeID = 0xe8ac8c6560747234

func NewSandstormCoreFactory_getSandstormCore_Params(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCoreFactory_getSandstormCore_Params{st}, err
}

func NewRootSandstormCoreFactory_getSandstormCore_Params(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCoreFactory_getSandstormCore_Params{st}, err
}

func ReadRootSandstormCoreFactory_getSandstormCore_Params(msg *capnp.Message) (SandstormCoreFactory_getSandstormCore_Params, error) {
	root, err := msg.RootPtr()
	return SandstormCoreFactory_getSandstormCore_Params{root.Struct()}, err
}

func (s SandstormCoreFactory_getSandstormCore_Params) String() string {
	str, _ := text.Marshal(0xe8ac8c6560747234, s.Struct)
	return str
}

func (s SandstormCoreFactory_getSandstormCore_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SandstormCoreFactory_getSandstormCore_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCoreFactory_getSandstormCore_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SandstormCoreFactory_getSandstormCore_Params) SetGrainId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// SandstormCoreFactory_getSandstormCore_Params_List is a list of SandstormCoreFactory_getSandstormCore_Params.
type SandstormCoreFactory_getSandstormCore_Params_List struct{ capnp.List }

// NewSandstormCoreFactory_getSandstormCore_Params creates a new list of SandstormCoreFactory_getSandstormCore_Params.
func NewSandstormCoreFactory_getSandstormCore_Params_List(s *capnp.Segment, sz int32) (SandstormCoreFactory_getSandstormCore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCoreFactory_getSandstormCore_Params_List{l}, err
}

func (s SandstormCoreFactory_getSandstormCore_Params_List) At(i int) SandstormCoreFactory_getSandstormCore_Params {
	return SandstormCoreFactory_getSandstormCore_Params{s.List.Struct(i)}
}

func (s SandstormCoreFactory_getSandstormCore_Params_List) Set(i int, v SandstormCoreFactory_getSandstormCore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCoreFactory_getSandstormCore_Params_Promise is a wrapper for a SandstormCoreFactory_getSandstormCore_Params promised by a client call.
type SandstormCoreFactory_getSandstormCore_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormCoreFactory_getSandstormCore_Params_Promise) Struct() (SandstormCoreFactory_getSandstormCore_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCoreFactory_getSandstormCore_Params{s}, err
}

type SandstormCoreFactory_getSandstormCore_Results struct{ capnp.Struct }

// SandstormCoreFactory_getSandstormCore_Results_TypeID is the unique identifier for the type SandstormCoreFactory_getSandstormCore_Results.
const SandstormCoreFactory_getSandstormCore_Results_TypeID = 0xea75b020e3e6c12a

func NewSandstormCoreFactory_getSandstormCore_Results(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCoreFactory_getSandstormCore_Results{st}, err
}

func NewRootSandstormCoreFactory_getSandstormCore_Results(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SandstormCoreFactory_getSandstormCore_Results{st}, err
}

func ReadRootSandstormCoreFactory_getSandstormCore_Results(msg *capnp.Message) (SandstormCoreFactory_getSandstormCore_Results, error) {
	root, err := msg.RootPtr()
	return SandstormCoreFactory_getSandstormCore_Results{root.Struct()}, err
}

func (s SandstormCoreFactory_getSandstormCore_Results) String() string {
	str, _ := text.Marshal(0xea75b020e3e6c12a, s.Struct)
	return str
}

func (s SandstormCoreFactory_getSandstormCore_Results) Core() supervisor.SandstormCore {
	p, _ := s.Struct.Ptr(0)
	return supervisor.SandstormCore{Client: p.Interface().Client()}
}

func (s SandstormCoreFactory_getSandstormCore_Results) HasCore() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCoreFactory_getSandstormCore_Results) SetCore(v supervisor.SandstormCore) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormCoreFactory_getSandstormCore_Results_List is a list of SandstormCoreFactory_getSandstormCore_Results.
type SandstormCoreFactory_getSandstormCore_Results_List struct{ capnp.List }

// NewSandstormCoreFactory_getSandstormCore_Results creates a new list of SandstormCoreFactory_getSandstormCore_Results.
func NewSandstormCoreFactory_getSandstormCore_Results_List(s *capnp.Segment, sz int32) (SandstormCoreFactory_getSandstormCore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SandstormCoreFactory_getSandstormCore_Results_List{l}, err
}

func (s SandstormCoreFactory_getSandstormCore_Results_List) At(i int) SandstormCoreFactory_getSandstormCore_Results {
	return SandstormCoreFactory_getSandstormCore_Results{s.List.Struct(i)}
}

func (s SandstormCoreFactory_getSandstormCore_Results_List) Set(i int, v SandstormCoreFactory_getSandstormCore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormCoreFactory_getSandstormCore_Results_Promise is a wrapper for a SandstormCoreFactory_getSandstormCore_Results promised by a client call.
type SandstormCoreFactory_getSandstormCore_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormCoreFactory_getSandstormCore_Results_Promise) Struct() (SandstormCoreFactory_getSandstormCore_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormCoreFactory_getSandstormCore_Results{s}, err
}

func (p SandstormCoreFactory_getSandstormCore_Results_Promise) Core() supervisor.SandstormCore {
	return supervisor.SandstormCore{Client: p.Pipeline.GetPipeline(0).Client()}
}

const schema_dcbc0d702b1b47a5 = "x\xda\xccX}pTW\x15?\xe7\xdd]\x1e\xa1\x09" +
	"\xc9\xf3\xa5|\xb5\x1a\xd1\x0da!\xc9.\x04:\x91)" +
	"&\xa4\xa5!)\x94\xbc\xa4\x88fDxI\x1e\xc9B" +
	"\xf6#\xef\xbd%I\x1d\xa4\x14\x88X\xdbA\x1c\x91\x01" +
	"\xa5|\x8c\xd5I\xdb\x082\x85\x11\x0d3\xa4\x82\x05\x04" +
	"\x01\xa10\xd4f\x86\xcaG\x1b\xc6\x8a0\xfa\x07\xa2}" +
	"\xce}\xbb\xf7\xed\xddf\x17\x08\xb53\x0e\xff\x90\xfb\xce" +
	"9\xf7|\xdds~\xbf\xf5\xbf\x9bU.Lqo\x98" +
	"\x03P7\x97\xb8\x87Y+\x7fwxvp\xe6\xb6\xe7" +
	"A\x19\x85\x08\xe0\x12\x01J\x8a<[\x11P\x9e\xe9)" +
	"\x03\xb4\xc6\xcfza\xb6_\xfd\xe6\x1a\x90F\xb9\xac\xf9" +
	"=}3\xda\x06\xf6\xf4\x01`\xc9\"\xcf\x08\x94\x83\x1e" +
	"\x11\xa0\xae\xc5C\xb0\xce\xf4\x08\x08`m\xf6\x1e\xfeg" +
	"\xa0\xf0\xb1.\x90F1{QO\x03\x82\xcb\x9a\\t" +
	"a\xf3\xd9\xd1\xc7\xbb@\x9a\x84\x00n\xa4\x9fT\xcf[" +
	"\xf4\xaa\xb6\xd8U\xfd\xcf\x8cX\xde\xfd\xf3\xef\xc6T\xdd" +
	"\xb6\xee&O5\x15\xd8\xe9\xf9\x00\xd0\xba\x99\xbf&\xf8" +
	"\xc1^\xef\xfa\xb8\x80m\xa13\x7f\x19\x15X\x97O-" +
	"X\xbf\xf8\xeaXQ\xec\xfc>/\xf0j~=\x15\xf8" +
	"\x95-p\xf1\x895\xfe\xebo\xcf\xde\xc8yw&\xbf" +
	"\x9az7=gs\xf9\xcd\xa9\xc1\xf8\x17\xb7@?\x1d" +
	"\xcco\xa0\xaa\xc7\xf2\xdb\xa9mU\xea\x1f\xfeJ\xf7V" +
	"N\xd5;\xa1\x96\xaa\xfe\xf6\xdc\x18\xb2\xfd\xd6\xfamq" +
	"UB?=L?\xa1<~\xc2n@k\xe5\xaf\x7f" +
	"\xb3vm\xd5\x0f^\xe1T\x0fNXFU\x17}X" +
	"Z}\xed\xcfcw\xc4sb\xabvO8NU\x0f" +
	"\xda\xaaO\xbf^\xda\xbb\xb6\xe4\xed\x9d|}\xd4\x82\x1f" +
	"\xdaI+\xa0\x11-\xff\xd3\xcc\x8bE\xa77ts\xb6" +
	"7\x16\xd8\xf9^\xd4\xf5\x88\xff\xa5\xbc\x1b\xaf\xf1\xc9X" +
	"Y\xf0\"U}\xd9V=\xd51\xff\xf1\xd5\x0b7\xbd" +
	"\xc1\x0b\xf4\x14\xd8\xf9>`\x0bl}\xf4\xd0\xed\x91\xb7" +
	"\xef\xf4\xc4/\x8fU\xa4\xbf\xa0\x02\x01K\xae\x16l@" +
	"\xc0\x1bY{\xf0h\xcf\xfa\xdd\\\xb1\xbd\xf5\xf4\xf2\xc7" +
	"\xdbve-:\xbb\xfaM\x90F\x91\xa4\xb6Q\xbd\xb5" +
	"(G\xbd\"\x80\xdc\xe6\xad\x94\xb7xG\x03X\xf3\xb7" +
	"\x1c7\xbb\xfa\xf2\xf6\xf3\x9e|\xcfk\xd7m\x93\xd7v" +
	"uC}W\xefy\xa5\x97\x178\x10\x138b\x0b8" +
	"WH#\x89\xf5j\xe5#\x93#Y\xbd\xef\x01\xa0|" +
	"\xd5\xbbW\xfe\xc8[\x00P\x925It\xc9g\x0aE" +
	"\x80\x8f\xc7L9\xf3\xe4?\xbe\xf6\x87\x84\xd7\x07\x0a?" +
	"G\xbd^}\xd6\xbd\xe4o#\xae\x9c\xe0+\xb9\xb3\xd0" +
	"\xbe\xa6\xa7\x90\x96\xa3\xa0\xa3q\xf1\xeb\xdb\x7fz\x92\xcb" +
	"vU\x91\xdd\x04%\xef|aN\xf4\xc7KO\xf3\xfd" +
	"3\xbd\xc8N\xe6\xac\"\xda?\x0d#/|k\xd9G" +
	"\xe7O\xf3!l+\xb2\xbb\xa4\xbb\x88\x86\xd01e\xe0" +
	"\xdc\xa5+\xe1K\xbc\x85\x93Ev\xa9\xfbm\x0bKV" +
	"\x8e\x1e8\xf8\xc7\xee\xcb\xdc\xe5_)\x1eG/?5" +
	"\xee\xe3c\xbbw\x9f\xbd\xc2\xab\x8e/\x9eJU\xbd\xc5" +
	"Tu\x9an.\xd1^z\xe3C\x90\xbe\xec\\\xbe\xae" +
	"\xf8]*\xb0\xa5\x98^>\xf0\xcb\xda;\x8fM|\xe8" +
	"zR\x82\x8bgP\x81>[`R\xdf\xb5\xcb_\xdc" +
	"\x13\xbd\xce[\xb8Z\xfc\x17*p\xdb\x16\xb8\xb0_\xfa" +
	"\xce\x8e\x17\xb6_\xe7\xbc\x1b\xeb\xb3\x1b\xf1\x94\xffy\xf7" +
	"\xf9<\xcf\x0d>\xabn\x9f\xfd\xb4$\x1f\xcd\xea\xe4`" +
	"\xfe\xd2\xb2\xc25\x7f\xffd\xf1J\xf6\xf9\xc6\xa1|\xc4" +
	"G\x9b\xa5\xcfW)\xdf\xa4\xff\xb3\xbe\xfd\xa5\x0e\xb4\xfe" +
	"z\xe2\x16\xef\xeaE\x9f\x9d\xc8\xab>\xeaI\xee7\xa6" +
	"u_*\xd5\xff\xc5\x0bd\xf8+\xec\xfb\xfcT`n" +
	"NEx\xe3\xc5\xb9wx\x81)~\xbb\xcc3\xfde" +
	"`%\xfd\xab\xb7\x1a\xd4\xc6\xe5Z\xa8\xa9\xd8\xdd\xa8F" +
	"B\x91\x19\x15\xf1?\x9b5\xb3RW\x03\xa1:3\xac" +
	"\xab\xcd\xda\x02Cm\xd6<\xb5\x9a\x11m%\xa6\xa1\xb8" +
	"\x88\x0b\xc0\x85\x00R\xd6$\x00e8A%W\xc0l" +
	"#\xf0\x9c\x86\x19 `\x06m\x89\xb8%\x92l\xd8\x08" +
	"7.\xd7\xcc\x1a\xd5l\x01\xa8A\xc4L\x100\x13@" +
	"\xc2[\x96o\x85\xaa\xfb\x0c5$4\x19fX\x0f\xfa" +
	"b\xa2\xbe\x065\xcfVu,\xba\x92-\x9a\xba\x1a2" +
	"\x96j\xba\xed\xaf\xed\xa3\xd8j\x1a\x8e\xf4\xb0d\xe9\x1a" +
	"\xb5q9\x0d(\xd2\x1aV\x9b\xeaL]S\x83\xc5\x86" +
	"\xbaB\x9bexjT]\x0d\x1a\xc0GW\x0b\xa0d" +
	"\x12T\xc6\x08hEb\x9aU\x80Mq\xaf1\x9dK" +
	"\xf48\x1a\x899\x14\xb7\x0aJ\x8ecV\xad\x06P\x96" +
	"\x10TZ\x05\x94\x10s\xe9\xfc\x91\x02\x15\x00J\x13A" +
	"%\"\xa0$\x08\xb9(\x00HAz\xd8BP1\x05" +
	"\x94\x08\xc9E\x02 \xb5\xd1\x9c\xb7\x12T:\x84\x98\x07" +
	"\xd1HU\x13\x000\xafV\x85\xdbC\x9a^\xe5x\xb9" +
	"\xaa\x99:\x92\xf8;;\x10Z\x1a\xc6\x1c\xab\xe5\xf3\xcd" +
	"u\xef\x1f\xfa\xf7>\x00\xc4\x9c\xf4\xc1\x04B\x86\xa9\xb6" +
	"\xb6\xc6\x13G\x13\x9c\x1dmMn\x82\x19\x89&(3" +
	"\xec\x9c\xa2\x94\x98\x92\x80(\xa57\xafk\xb4\xdaZ\xa2" +
	"zw\xeb\xb0\xa1y\xde\xa4\xb5j\xa6\xb6\xc0\xd0\xf4\x98" +
	"a\xd3\x00H+\x1bn\x0f\xd1\x96\xa8\xb0\xf3I\xab&" +
	"\xaaAC\x19\xee\xf8\xe1\xa5E\x9bHP\x99\xc6\x15m" +
	"\x0a\x8d\xbc\x90\xa0R\x9a\xba\x14\x89l,\xbe|\xca\xdb" +
	"^\xba\xf0\xe4=\xb2\x11s\x99K\x86\x89\xc6}e\xce" +
	"n34\x94L\xc7\xe1\xd9\xd4\xe1'\x09*5\x9c\xc3" +
	"\xf3hC\xcd!\xa8<\xcbu\x99B\x0f\xe7\x12T\xbe" +
	"\xfe`\x0du\x9f9e\x9d\xf3\x00O\xd3)\x1f\x1f\xdf" +
	"T\x00\xa5\x9c\xa02\x97\x8b\xaf\xaa:U|o\x01(" +
	"\xcf\x12T\x96\x08\x98\xa7F\"\x9c\xe7A5\x14X\xaa" +
	"\x19&\x8d6\xc7:\xfe\xfe\xb5\xc02\xef\xe2u\xac\xaf" +
	"\xd4\xa8\xd9\x12\xd6k\x9a\x85\xc8\xd3Z\xe7S\x81P\xb3" +
	"\xa6G\xf4@\xc8\x84A\x91\x0f\x1e\x9f\xb4\xed\x06OO" +
	"\x13\x876=\xef\xfe\x16\xe3]:\x14/\x9cq\x94\xe6" +
	"\xf9F\x0d\xbe\xd2\xe9\xfc0LU7\xf9.\xfd\x84\xc5" +
	"znn\x1a\xd1\x88\xa6\xaf\x08\x18@\xc2:J\x96\x7f" +
	"Ai\xe3\xf8>e\xcb=\x1e\x02w\x03\x9b\xcb\x8f:" +
	"\xf6\xf7\xd1\x86\xddCP\xe9\xe5J\x7f\x80\x1e\xbeIP" +
	"9\xc4\x95\xfe \x9d\xe0\xbd\x04\x95\xa3\xdc\x00=B%" +
	"\x0f\x11TN\x08\x88\xae\\t\x01H\xc7h7\x1d&" +
	"\xa8\x9c\x16Prc.\xba\x11\xa5\x93T\xf0(A\xe5" +
	"\x9c\x80\xd20!\x17\x87!Jg\xa8\xc9\xd3\x04\x95\xf7" +
	"\x84{\xbf\x8c\x14KcUc8\x18TCM\x98c" +
	"\x9d\xea\x9a\xb8c\xafR\xf5\xfbx\xbf\xe5\x05\x8cg\xb4" +
	"vD\x10\x10\x01W5i+\xe6\x85\x9b4\xf6\xb7\x15" +
	"\x0cGCf\x8d\x1e\x06lt\xce\xee:F\x12\xef\xce" +
	"\x1e\xaa\xe9\x84\x07\xbf;\x0c\xd6 *.\xe2\x06p\x98" +
	"\x0d28/I3@\x90\xdcbY\xecm\x96\xa3\xe2" +
	"B\xe4\xe6\x1b\xdc\xdd)\xaewU\x124\xfe\x07\xdb6" +
	"\x1aI\x1e2\x836H\xca\xf5\x94n #3\x1eC" +
	"\x1d\xcap\xe4ybF=\x87\xfe3vY\x0c\xce\x00" +
	"1[,\x96Hd\x99\x1455\xa8\x94\x137&\xe8" +
	"\x062b\"\xdf\xc4z\x10\xe4\x01\x14Qp0.2" +
	"\xa8*\xf7c5\x08\xf2;(\"q\xe072\x9a&" +
	"\x1f\xc3\x06\x10\xe4>\x14q\x98C\x92\x901Hy\x1f" +
	">\x07\x82\xdc\x83\"\x8a\x0e\x1cD\x06T\xe5\x9d\xa8\x83" +
	" oA\x11\x87;\xdc\x04\x19\xc0\x95_\xb6\xbf\xaeC" +
	"\x113\x1c\xce\x8a\x8c\x19\xc8\x9d\xf6\xbdm(\xe2\x08\x87" +
	"\x19\"#\xa6\xb2\x86\xcb@\x90\x17\xa1\x88\x0f9@\x16" +
	"\x19\xbb\x91\x15\xfbk\x15\x8a\x98\xe9\x10Rd\xecQ\x9e" +
	"i\xfb<\x1dE\xccr\xd8\x04R\x02\x06=\xebw\xcb" +
	"^[w<\x8a8\xd2\xa1~\xc8\xe8\xa3\xfc0\xee\x02" +
	"A\x96P\xc4\x04\xcbA\xc6\xd5e7\x8dH\xfa\x8f\x88" +
	"n\x07<#\xa3\xca\xd2\xcdz\x10\xa4\x01\x8e6qD" +
	"\xa4\x7f\x12\x08\xd2\x19\x11\xb3\x1d\xf6\x82\xec\x17\x05\xe9\xc8" +
	"k H}\xa2\xc5\xc6\x14\x90@\xa8\x1c-\x06\x9b\x01" +
	"\xa0\x1c-\xb6\xccAT\xed\xaflzCY\xacQ\xca" +
	"\xd12\xf5\xceJ\xbb\x7f\xf2\xd8\x09{%\x89\x13\x86&" +
	"\x99\x19\xb6\xf7![\x8f\x1d\xb0\x07\x00\xd9T\x90\xda\x88" +
	"/^(\x8b\xbd\x0a\xc7,'\xc3\x96\x03\xb2\xed 2" +
	"\x8fbp\x1a\xf2*\xe3\xe6\x19\x8c\x02\xa2\xe9\xe5\x98\x1d" +
	"\x09\x84\x9a\xb9P\x91\xad\x97l\xc36P\x83i\xb1?" +
	"\xd5d@\xe5\xfe\xd0<\x1b\x12\xdc\xd6\xafH\xb5\xf5+" +
	"8\xa8\xc3F\xff\xbc\xfa\x04\xaa\xb9\xe7\x9c\x0ei\xed\xf3" +
	"\xa9\x08\x90\xaa!\xa1\xfb{a4\x1e\xcf9\xab\x97\x03" +
	"\x95\xd4q\x0fA\xc5\xcfEST\x91@\x9a\x0f\x0c\xbd" +
	"\xf8\x05\xc0\xb0!7\x12\xab\xb9Q\x9b\x02\xf1\x0d\x89\x16" +
	"2\xf3\x9fEX\xa9:h\xd0>#\xa9\x9d\xf4\xd4\xe4" +
	"\xd9\x19\xffL\x1c\x8bc\xd7:5\x14c\xabO\x84u" +
	"\xed)\xb5\xd1\x0c\xeb\x9d\xd4\x81\xa4\xf3T\xbb\xae\"\xb1" +
	"\x93\xee7x'\xae\xda\xb2\x18\xe6\xfa\x94\x88kH!" +
	"0^\x9d\x0e\xbd6\x86u\x0d%+s\xff\x8fvu" +
	",\xfc\xd9O\xee\x8b\xe6p\x942\x99\xb3\x0f\x9a\x09\xb1" +
	"!\x99,\xfd\xffJ\x05H\xba\xb4\x12\xbd3\x01\xab\xd8" +
	"\x8fV\xc8~{\x92\xa4\x17A\x90\xb2D\x8b\xa5\x1e\x99" +
	"2\xdd%\xfcT\xbd\x1b\xfc\xf9\x14O==yN\xf1" +
	"\xc3\xc8\xd0)Cr\x11\x87\x8e\xfe\xfe\x1b\x00\x00\xff\xff" +
	"\x8bZ\xaf\x02"

func init() {
	schemas.Register(schema_dcbc0d702b1b47a5,
		0x809d3d6d45c4c37d,
		0x835c613045824121,
		0x86362c69f5c42997,
		0x86ca17d397d72d2b,
		0x87a6a96b0a4edd21,
		0x8829b2e76d8325f1,
		0x8b790707193ea7ff,
		0x9145c7ea308343d9,
		0x916d32f140971035,
		0x9aa99e08dd1161ff,
		0x9d88f29f0318d4bb,
		0x9e90498484bab87d,
		0xa019dbe64a38e85d,
		0xa1c73384bc38ab4b,
		0xa98fd02dd93dd26b,
		0xaaef1f8c301b865d,
		0xac9557813c4f78cf,
		0xadfbf90ef9c01c9a,
		0xaf88ad00c801b00d,
		0xb481d35d0da2713c,
		0xb61fc18674ca994f,
		0xbc51d6bc865a8fcf,
		0xc1b0e9713ac1ad4f,
		0xcb56f444d1311800,
		0xcce40aee6005d381,
		0xcd9c9fab5f637827,
		0xd0669675481ed533,
		0xd0d6ed6a5ed70e62,
		0xe06fe4e0d4e93178,
		0xe3a9cebde9177d60,
		0xe4d3afafc9fe1acf,
		0xe8ac8c6560747234,
		0xea0b2836fb52aee9,
		0xea75b020e3e6c12a,
		0xea9f82a07e11b6d7,
		0xef241fd6058030cf,
		0xf0832c3f66256d2b,
		0xf2ccecff0178227b,
		0xfa7238e0a9345914,
		0xfb4cd9916f42104c)
}
