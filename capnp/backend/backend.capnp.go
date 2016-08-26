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
}

func Backend_ServerToClient(s Backend_Server) Backend {
	c, _ := s.(server.Closer)
	return Backend{Client: server.New(Backend_Methods(nil, s), c)}
}

func Backend_Methods(methods []server.Method, s Backend_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 15)
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

const schema_dcbc0d702b1b47a5 = "x\xda\xccX}PT\xd7\xd9?\xe7^\xd6\x15#\xca" +
	"}/\x89\xc6\xbc\xe9\x96\x05fe?\"\x8a\xc9\x10\xc6" +
	"\x14\xdc\x84\x10\x88D\x16Bm\x98Z\xb9\xb0\x17Xd" +
	"?\xbcwW\xd8tRc4\xd4\xa6fR;\xb1\xd6" +
	"\xb4FtL[\x1a\xa9\xc6I\x9c\xb4\x95\x99\x90ij" +
	"\xb5I\xd1Zul\xebL\xad\x89\x09N[\x8bS\xff" +
	"H\xad\xb9}\xce\xd9=w\xcf\x06V\xc0\xb43\x1d\xc7" +
	"\x81\xfb|\x9d\xe7\xeb<\xe7\xf7P\"\xe7T\x0a\x8b-" +
	"\xb6*\x84\x1a\x0bE\xcb\x0c#\x7f\xf93U%\xca\x97" +
	"7!\xe9\x8e,c\xe5\xe0p\xf9\xba\xd1\x83\xc3\x08\xe1" +
	"\xd2\x98}\x16\x96\x9f\xb5[A\xf0i\xbb\x88\x1b\xb7\xda" +
	"\x05\x8c\x90\xb1\xa3\xf8\x17\xd7\x02\xee\xfb\xfa@\x1e>\xb3" +
	"\x80]\xfa\xbc\xbd\x15\xa3,\xc3\xe59\xbb\xe3\xd4\xbc\xe3" +
	"\xc0q\x02\xc7\x82\x09+n\x7f\x1b#,\x7f\xc3^\x81" +
	"\xb0\x91\x7f\xfe\xb1Yk\x07~\xf0\xf5\x84\xaa\x85\xea\x0e" +
	"\xdak\x89\xc0O\xed\x1f\x82\xc0X\xd1\xa6\xe0\x87\x87\x8a" +
	"\xb7$\x05\xa8\x85m\x05]D`W\x01\xb1`\xfc\xf0" +
	"\x0bwZ\xad\xf1o\xf2\x02C\x05\xcdD\xe0\x18\x158" +
	"\xf7\xe0\xa6\x92\xcbG\xab\xb6q\xde\x8d\x16\xd4\x12\xef\xee" +
	"\xcd\xddQ9\xb6$\x98\xe4X\x04\xc2:]\xd0JT" +
	"/\x14\xf4\x10\xdb\x8at~\xe6\xcb\x03/q\xaa\xcb\x0b" +
	"\x1b\x88\xea\xcf\x7f7_\xdc}u\xcb\xae\xa4\xaaHX" +
	"\x1e\xc2\xc2\xf2\xfd\x85\x07@\xf5\xa97\x7f\xb6ys\xcd" +
	"\xb7^\xe6TO\x17v\x11\xd5\xd5\x1f\x95\xd5^\xfa\xc3" +
	"\x9d\xfd\xc9\x9cP\xd5\xe1\xc2\xe3D\xf54U}\xf4\xd5" +
	"\xb2#\x9bK\x8f\xeeA\xbe;0\xd3\x8d\x17}\x9b&" +
	"\xad\x88D\xb4\xf6\xb7\x0f\x9c\xf3\x9cxa\x80\xb3=P" +
	"D\xf3\xbd\xba\xef\xae\x92\xad\xb6+?\xe6\x93\xb1\xbd\xe8" +
	"9\xa2\xfa\x0aU\x1d\xe9]\xb9l\xe3\xaa\xed\xfby\x81" +
	"w\x8ah\xbeOR\x81\x97\xfe\xff\xad\x8f\xe7||}" +
	"0yx\xa2\"\xd7\x8a\xbc Qz\xa3h\x15\xfc\xb8" +
	"\x92s\x10\xffjp\xcb\x81\xd4\xe1u\x8efr\xf8\xb2" +
	"u{sV\x9f\xda\xf8:\xd8\x16\xd3\xda\xe6~G\x03" +
	"\x96\xeb\x1c *\xd78\xaa\xe5u\x8ey\xd04+w" +
	"\x1e\x8f\xf6\x0d\xdb\x0e\xf3\x9e(\x0eZ\xb7\xa0\x83\xba\xfa" +
	"Bs\xdf\x913\xbe#i\x95O\x08\xec\xa2\x02\xe6\x11" +
	"\xd2\x1c\xd1x\xa5\xfa.W$\xe7\xc8\x1f\xe1<y\xc8" +
	"qH~\xc7\xe1\x00\xf9\x0f\x1cGE9\xee\x04\xd5O" +
	"\xe6/>\xf9\xd0?\xbe\xf8\xeb\x94\xd7\x8a\xf3\xff\x88\xd7" +
	"\x1bOYZ\xfe6\xeb\xfdw\xf9J\xd68\xe91M" +
	"NR\x0eGo\xdb\x9aWw\x7f\xff=.\xdbcN" +
	"\xda\x04\xa5\xa7?\xf7H\xec;\xed'\xf8\xfe9\xef\xa4" +
	"\xc9\x1cu\x92\xfei\x9ds\xf6+]\x7f=s\x82\x0f" +
	"\xa1\xcaE\xbb\xc4\xe7\"!\xb4<5ot\xe87\x03" +
	"\x179\xdb1\xd7\x02b{d\xc1'\xc7\x0e\x1c8\xf5" +
	">o{\xb5k\x09QU]\xc4\xf6R-\xda\xa2n" +
	"\xdd\xff\x11\x92\x0aL\xdb\xc3\xae\xdf\xd36\xa2\xb6G\x7f" +
	"\xd2p\xfd\xbe\x85\xb7]\xe6\x0f\xbf\xe6*'\x027\xa8" +
	"\x80s\xf8\xd2\xc5\xcf\x1f\x8c]\xe6-\xe4\xbb\xffL\x04" +
	"\xeeu\x13\x81\xb3\x87\xa5\xaf\xf5?\xb3\xfb2\xe7]\x93" +
	"\x9b\xf6\xd9H\xc9\xd3\x963\xb6\xc2+|\xd2\x96\xbb\xe9" +
	"\xcd\xa9s\x93\xa4\xb9\x82E\xed\x15\xeeM\x7f\xfftm" +
	"J\xc7\xdc\x0b\xb0\x8c=\xa4\x17n\xb8\xabe\x0f\xf9\xcd" +
	"\xf8\xaa\xbd\x17\x1b\x7fy\xf7*\xef\xea\xed\x1e\x9a\xa7|" +
	"\x0f\xf1$\xef\x89\xa5\x03\x7f*\xd3\xfe\x99\x96H\x8f\x97" +
	"\x9eG\x05V\xe4z\xc3\xdb\xce\xad\xb8\xce\x0b\x04=\xb4" +
	"\x8aq\x100\xd2\xfe\x09F\xab\xd2\xb6V\x0d\xf9\xef\x11" +
	"\xdb\x94H(R\xeeM~\xeaa\xf8\x19\xadW\xa2\x9d" +
	"\x08\xd5c\x8cg#\x01\xfe#\x09_5\x16\xadW\xb4" +
	"E\xba\x12\x12\xfcz4\xac\x05\x17%D\x17\xb5*6" +
	"\xaajZ\xccJ\xb7\x18\xd5\x94\x90\xde\xaej\xd5\x9a\x12" +
	"\x08\x156\xa8z\xcc\xda\x1d\xd5M\xe9\x19\xe9\xd2\xf5\xf0" +
	"S\xe9P\x9b\"\xdda\xc5\xdf\x18\xd5T%x\x8f\xae" +
	"\xacW\x97\xeb\x85\xf5\x8a\xa6\x04u\xe4\xcb\x12\xb3\xa0\x14" +
	"\x10\xa3\x94\xd3\x80\x90o\xb6\x88}\xf3\x05lD\x12\x9a" +
	"5\x08\xfb\x93^\xe3L.\x11r,\x92p(i\x15" +
	"\xf9rM\xb3J-|\xb6\x80\xd9n\x01K\x18\xe7\x91" +
	") \x05\xbc@\xf4\x031\x02DA\xc8\xc3\x02\x10\x83" +
	"\x84\xd8\x09\xc4(\x10E1\x0f\x8b@\\\xe7\x04b7" +
	"\x10{\x85\x84\x07\xb1H\x8d\x1fA\xe5\x93^m\x08\xf7" +
	"\x84T\xad\xc6\xf4rC\x07q$\xf5=7\x10j\x0f" +
	"\xe3\\\xa3\xf3\xee\x8e\xc6\x0bo\xfd\xeb\x0d\xd0\xc4\xb9\x99" +
	"\x83\x09\x84\xf4\xa8\xd2\xdd\x9dL\x1cI\xf0\xdc\x18$\x98" +
	"OS984\x13\x1c\xca\x13p\x85Ns\x8a\xa5\xd4" +
	"\xac\x02\xf3Rf\xf3\x9aJ\xaa\xad\xa6\xaa\xd7-\xa6\x1b" +
	"w\xa6\x8cO\xd3s\xbf\xda\xadF\xd5&]\xd5\x12\x86" +
	"\xa3P\x86\x8c\xb2\x904\xd2\x12^\x9aOR5+\x94" +
	"\x0d\x0ef~\x14\x93\xa2-\x04?\x96rE[L\"" +
	"w\x03\xb1l\xe2R\xa4\xb2\xb1\xe6\xe2HqO\xd9\xaa" +
	"\xf7&\xc9F\xc2e.\x19Q\xacO)s\xb4\xcd\xb0" +
	"\x0e\xdd\xca\x1c\xae\"\x0e?\x04\xbe\xd5s\x0e\xd7\x91\x86" +
	"z\x04\x88\x8fs]\xe6#\xc4\x15@\xfc\xd2\xad5\xd4" +
	"\x14s\xca:\xe7\x16\xae\xa6Y>>\xbe%\xf0Y\x09" +
	"^\xaf\xe0\xe2\xab\xa9\x9d(\xbe\xb7\x81\xf88\x10[\x04" +
	"lS\"\x11\xce\xf3\xa0\x12\x0a\xb4C\x1eI\xb4\xb9\xc6" +
	"\xf1\x0b\x97\x02]\xc5k\x9ee}\xa5\xc4\xa2\x9da\xad" +
	"\xbeC\x88<\xaa\xc6\x1f\x0e\x84:T-\xa2\x05BQ" +
	"4.rKz(\x1dj\x94\xb4]#\xd4\x87D\xa4" +
	"'/\x0e-g\xa6\xde\xd6\x03O\xaa8\x1b\xecfO" +
	"\xf9.&\xbbt:^\x98\xe3(\xc3\xf5\x8d\xe9|\xa5" +
	"3\xf9\x01^hQ\xbeK?e\xb1\x99\x9b\x9bz," +
	"\xa2j\xeb\x03:\x12\xc3\x1a\\\x84\x92\xa6\xb2\xb6\xfca" +
	"\xdf\xceI.\x02w\x02\x9b\xcb\xf3M\xfb;I\xc3\xbe" +
	"\x08\xf6\xfb\xb9\xd2\xef\"\xc4\xef\x02q\x1fW\xfa=d" +
	"\x82\xf7\x03q?7@\x07\x88\xe4> \x1e\x140\xce" +
	"\xca\xc3`V\x1a$\xdd\xf4#\xa0\xbd\x0e\x82\x160i" +
	"\x01\xf7^#\x82\xfb\x81\xf8\xa60\xf9%\x98\xe0}\xd8" +
	"\xd0\x16\x0eB\x87\xf9\xa1\xb5F\xfa\x16\xf6\x1f\xf2\xd5\xfc" +
	"2\xd9Z\xb6\x80\xfe\x98\xda\x03~\x83\x07 \xe7W\xd7" +
	"\xd7\x85\xfd*\xfb\xbe\xf9tH]':+3\x09\x8f" +
	"\xbfN8\x08/.\x94\xc9\x02c\x90\xad\x0d\x98ae" +
	"I*G\x82d\xb1V$\xae\\%\x08b\xcc\x8d-" +
	"ts\xa7\xb8\x96T\xc4\xa0\xfe\x1fxDc\x91\xf4\xd9" +
	"1\xeea\x98\xf0\xd5\xc94g13\x9e\x00\x13\xbe\x99" +
	"\x10\\j\x09\xcbn\xe6\xa0u\xf6^\x83\xa1\x14$F" +
	";\x0d\x96H\xcc2i\x85\xa3|\xcbD\xe8\x10\x13\xcb" +
	"c\x86\xfa\xe5\x93\xb8\x19\x09\xf21l\xc5\x82\x8901" +
	"\x03\x8a\xf2\x10\xae\x05\xee\x1b\xc0\x15Ml\x8b\xd9\x0e$" +
	"\x0f\xe0V\xe0\xee\x01\xee\x0cs\x03\xc1l=\x93\xb7\xe3" +
	"'\x81\xfb<p\xad&\x18\xc3\x0c&\xca\x1b\xb1\x06\xdc" +
	"8pg\x9a\xc0\x1f3x)\x07)W\x05n\xb6\xb9" +
	"\x10b\x06\xbb\xe5'\xe8\xb9>\xe0\xce2\xd7.\xcc\xb6" +
	">\xb9\x0aw\x01\xf7\x01\xe0\xdef\xc2H\xccV\x07y" +
	"1\xe5\x16\x03w\xb6\xb9\xeda\xb6\x9a\xc9wS\x9fo" +
	"\x07n\x8e\x09\xd51\xd9n\x10\xac7r6\xd5\xc5\xc0" +
	"\x9dc\xeeU\x98\xedf\xd2\xb5\xbd\xd0\x90cV\x9cZ" +
	" 0[\x83\xa5\x0f \x1c\xe9\xbc\x15[L\xdc\x8a\xd9" +
	"\x12*\x9d\x84\x02H\xc7\xb8\x85\x84\xdb\x01\x86\x9c\xc0{" +
	"\xcdj\xb0\xf1\x82\xc4@\xa8\x12\x1b0)\xe9\xac\x816" +
	"\x87/\xf6\x08#\xabB\xb9l\xea\xa2\x8aD'\x00)" +
	"\xaa\xc5\xabi\x83\xd8\x18\x85]\x83\x14\x85\xa1@f\x86" +
	"\xbd\xd7h\xae\x96 \xb0\x0eGs\x89 \xb1\x91|0" +
	"QE\xa2\xedM\xb3\x9c\x0c\x1b\xea\x98Mu+\xf3(" +
	"\x01\x83\x91\xad:i\x9e\xc1\x1f$\xaaZ%\x9e\x1b\x81" +
	"\xf7\xab\x12\xc3\x04\xc8\x04\xcd\x89\x00\xc3\x11S\x03\xdb\xec" +
	"\xb2s\x8f\xb2w\xa2G\xd9\xcb!\x116\x99\xeb\x9aS" +
	"\xa0c\xd2\xd9\x1aR{V\x12\x11$\xd6L\x0b|O" +
	"\x06\xa1x\xb8e\xbe\x8c\x1c\xe6#\x8e\x17\x82\x8f%\\" +
	"4\x1eo\x0a\x08\xde22\xe2\x079\x83n\xdch\xab" +
	"\xe5F\xe6\x04\x80\xec\xa6\x15\x1c\xf7.\x88\xe3PA2" +
	"`\x1b\x8d\xf8\xbf\x12o\x12\xda5\xc2\xe3G\x97\xb9\x07" +
	"\xa1\xf1\x1fV\xda\xe0\xd78q \x8d>\xd1\x9b\xe1M" +
	"\xcd\xf6\x8cgd\x8a\xab\xa1\"\x01I># \x99V" +
	"\x08l\xed\xcc\x04\xee\xda@\x08\xce\x9a}\xf8\xc5\xbd\xbd" +
	"\xab\xf6}oJ[\x00\xb7q\xa5\xaf\xb4\xe3\xeedb" +
	"\x16\xa5K\xff\xaf\"e1SZE-\x9e\x82'\xec" +
	"O/\x98\xfd\x05E\x92\x9e\x83\xc9\x9dc5X\xea1" +
	"S&#\x9b\x9fj7\x83\x11\x9f\xe1\xaae\xde-'" +
	"\xf8\xbb\xc1\xf4\x11uz\x11\xa7\x8f\xa2\xfe\x1d\x00\x00\xff" +
	"\xff\xeb\x81E\xd6"

func init() {
	schemas.Register(schema_dcbc0d702b1b47a5,
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
