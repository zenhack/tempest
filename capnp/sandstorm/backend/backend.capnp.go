package backend

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	grain "zenhack.net/go/sandstorm/capnp/sandstorm/grain"
	spk "zenhack.net/go/sandstorm/capnp/sandstorm/spk"
	supervisor "zenhack.net/go/sandstorm/capnp/sandstorm/supervisor"
	util "zenhack.net/go/sandstorm/capnp/sandstorm/util"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

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
}

func Backend_ServerToClient(s Backend_Server) Backend {
	c, _ := s.(server.Closer)
	return Backend{Client: server.New(Backend_Methods(nil, s), c)}
}

func Backend_Methods(methods []server.Method, s Backend_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 12)
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
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
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
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
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
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Params{}, err
	}
	return Backend_PackageUploadStream_saveAs_Params{st}, nil
}

func NewRootBackend_PackageUploadStream_saveAs_Params(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Params{}, err
	}
	return Backend_PackageUploadStream_saveAs_Params{st}, nil
}

func ReadRootBackend_PackageUploadStream_saveAs_Params(msg *capnp.Message) (Backend_PackageUploadStream_saveAs_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Params{}, err
	}
	return Backend_PackageUploadStream_saveAs_Params{root.Struct()}, nil
}

func (s Backend_PackageUploadStream_saveAs_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_PackageUploadStream_saveAs_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Params_List{}, err
	}
	return Backend_PackageUploadStream_saveAs_Params_List{l}, nil
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Results{}, err
	}
	return Backend_PackageUploadStream_saveAs_Results{st}, nil
}

func NewRootBackend_PackageUploadStream_saveAs_Results(s *capnp.Segment) (Backend_PackageUploadStream_saveAs_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Results{}, err
	}
	return Backend_PackageUploadStream_saveAs_Results{st}, nil
}

func ReadRootBackend_PackageUploadStream_saveAs_Results(msg *capnp.Message) (Backend_PackageUploadStream_saveAs_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Results{}, err
	}
	return Backend_PackageUploadStream_saveAs_Results{root.Struct()}, nil
}

func (s Backend_PackageUploadStream_saveAs_Results) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_PackageUploadStream_saveAs_Results) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_PackageUploadStream_saveAs_Results) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return spk.Manifest{}, err
	}

	return spk.Manifest{Struct: p.Struct()}, nil

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

// Backend_PackageUploadStream_saveAs_Results_List is a list of Backend_PackageUploadStream_saveAs_Results.
type Backend_PackageUploadStream_saveAs_Results_List struct{ capnp.List }

// NewBackend_PackageUploadStream_saveAs_Results creates a new list of Backend_PackageUploadStream_saveAs_Results.
func NewBackend_PackageUploadStream_saveAs_Results_List(s *capnp.Segment, sz int32) (Backend_PackageUploadStream_saveAs_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Backend_PackageUploadStream_saveAs_Results_List{}, err
	}
	return Backend_PackageUploadStream_saveAs_Results_List{l}, nil
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
	if err != nil {
		return Backend_startGrain_Params{}, err
	}
	return Backend_startGrain_Params{st}, nil
}

func NewRootBackend_startGrain_Params(s *capnp.Segment) (Backend_startGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return Backend_startGrain_Params{}, err
	}
	return Backend_startGrain_Params{st}, nil
}

func ReadRootBackend_startGrain_Params(msg *capnp.Message) (Backend_startGrain_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_startGrain_Params{}, err
	}
	return Backend_startGrain_Params{root.Struct()}, nil
}

func (s Backend_startGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_startGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_startGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_startGrain_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return spk.Manifest_Command{}, err
	}

	return spk.Manifest_Command{Struct: p.Struct()}, nil

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
	if err != nil {
		return Backend_startGrain_Params_List{}, err
	}
	return Backend_startGrain_Params_List{l}, nil
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
	if err != nil {
		return Backend_startGrain_Results{}, err
	}
	return Backend_startGrain_Results{st}, nil
}

func NewRootBackend_startGrain_Results(s *capnp.Segment) (Backend_startGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_startGrain_Results{}, err
	}
	return Backend_startGrain_Results{st}, nil
}

func ReadRootBackend_startGrain_Results(msg *capnp.Message) (Backend_startGrain_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_startGrain_Results{}, err
	}
	return Backend_startGrain_Results{root.Struct()}, nil
}

func (s Backend_startGrain_Results) Supervisor() supervisor.Supervisor {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return supervisor.Supervisor{}
	}
	return supervisor.Supervisor{Client: p.Interface().Client()}
}

func (s Backend_startGrain_Results) HasSupervisor() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_startGrain_Results) SetSupervisor(v supervisor.Supervisor) error {

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

// Backend_startGrain_Results_List is a list of Backend_startGrain_Results.
type Backend_startGrain_Results_List struct{ capnp.List }

// NewBackend_startGrain_Results creates a new list of Backend_startGrain_Results.
func NewBackend_startGrain_Results_List(s *capnp.Segment, sz int32) (Backend_startGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Backend_startGrain_Results_List{}, err
	}
	return Backend_startGrain_Results_List{l}, nil
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
	if err != nil {
		return Backend_getGrain_Params{}, err
	}
	return Backend_getGrain_Params{st}, nil
}

func NewRootBackend_getGrain_Params(s *capnp.Segment) (Backend_getGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_getGrain_Params{}, err
	}
	return Backend_getGrain_Params{st}, nil
}

func ReadRootBackend_getGrain_Params(msg *capnp.Message) (Backend_getGrain_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_getGrain_Params{}, err
	}
	return Backend_getGrain_Params{root.Struct()}, nil
}

func (s Backend_getGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_getGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_getGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_getGrain_Params_List{}, err
	}
	return Backend_getGrain_Params_List{l}, nil
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
	if err != nil {
		return Backend_getGrain_Results{}, err
	}
	return Backend_getGrain_Results{st}, nil
}

func NewRootBackend_getGrain_Results(s *capnp.Segment) (Backend_getGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_getGrain_Results{}, err
	}
	return Backend_getGrain_Results{st}, nil
}

func ReadRootBackend_getGrain_Results(msg *capnp.Message) (Backend_getGrain_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_getGrain_Results{}, err
	}
	return Backend_getGrain_Results{root.Struct()}, nil
}

func (s Backend_getGrain_Results) Supervisor() supervisor.Supervisor {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return supervisor.Supervisor{}
	}
	return supervisor.Supervisor{Client: p.Interface().Client()}
}

func (s Backend_getGrain_Results) HasSupervisor() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getGrain_Results) SetSupervisor(v supervisor.Supervisor) error {

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

// Backend_getGrain_Results_List is a list of Backend_getGrain_Results.
type Backend_getGrain_Results_List struct{ capnp.List }

// NewBackend_getGrain_Results creates a new list of Backend_getGrain_Results.
func NewBackend_getGrain_Results_List(s *capnp.Segment, sz int32) (Backend_getGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Backend_getGrain_Results_List{}, err
	}
	return Backend_getGrain_Results_List{l}, nil
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
	if err != nil {
		return Backend_deleteGrain_Params{}, err
	}
	return Backend_deleteGrain_Params{st}, nil
}

func NewRootBackend_deleteGrain_Params(s *capnp.Segment) (Backend_deleteGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_deleteGrain_Params{}, err
	}
	return Backend_deleteGrain_Params{st}, nil
}

func ReadRootBackend_deleteGrain_Params(msg *capnp.Message) (Backend_deleteGrain_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deleteGrain_Params{}, err
	}
	return Backend_deleteGrain_Params{root.Struct()}, nil
}

func (s Backend_deleteGrain_Params) OwnerId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_deleteGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deleteGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_deleteGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_deleteGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_deleteGrain_Params_List{}, err
	}
	return Backend_deleteGrain_Params_List{l}, nil
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
	if err != nil {
		return Backend_deleteGrain_Results{}, err
	}
	return Backend_deleteGrain_Results{st}, nil
}

func NewRootBackend_deleteGrain_Results(s *capnp.Segment) (Backend_deleteGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_deleteGrain_Results{}, err
	}
	return Backend_deleteGrain_Results{st}, nil
}

func ReadRootBackend_deleteGrain_Results(msg *capnp.Message) (Backend_deleteGrain_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deleteGrain_Results{}, err
	}
	return Backend_deleteGrain_Results{root.Struct()}, nil
}

// Backend_deleteGrain_Results_List is a list of Backend_deleteGrain_Results.
type Backend_deleteGrain_Results_List struct{ capnp.List }

// NewBackend_deleteGrain_Results creates a new list of Backend_deleteGrain_Results.
func NewBackend_deleteGrain_Results_List(s *capnp.Segment, sz int32) (Backend_deleteGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_deleteGrain_Results_List{}, err
	}
	return Backend_deleteGrain_Results_List{l}, nil
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
	if err != nil {
		return Backend_installPackage_Params{}, err
	}
	return Backend_installPackage_Params{st}, nil
}

func NewRootBackend_installPackage_Params(s *capnp.Segment) (Backend_installPackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_installPackage_Params{}, err
	}
	return Backend_installPackage_Params{st}, nil
}

func ReadRootBackend_installPackage_Params(msg *capnp.Message) (Backend_installPackage_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_installPackage_Params{}, err
	}
	return Backend_installPackage_Params{root.Struct()}, nil
}

// Backend_installPackage_Params_List is a list of Backend_installPackage_Params.
type Backend_installPackage_Params_List struct{ capnp.List }

// NewBackend_installPackage_Params creates a new list of Backend_installPackage_Params.
func NewBackend_installPackage_Params_List(s *capnp.Segment, sz int32) (Backend_installPackage_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_installPackage_Params_List{}, err
	}
	return Backend_installPackage_Params_List{l}, nil
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
	if err != nil {
		return Backend_installPackage_Results{}, err
	}
	return Backend_installPackage_Results{st}, nil
}

func NewRootBackend_installPackage_Results(s *capnp.Segment) (Backend_installPackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_installPackage_Results{}, err
	}
	return Backend_installPackage_Results{st}, nil
}

func ReadRootBackend_installPackage_Results(msg *capnp.Message) (Backend_installPackage_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_installPackage_Results{}, err
	}
	return Backend_installPackage_Results{root.Struct()}, nil
}

func (s Backend_installPackage_Results) Stream() Backend_PackageUploadStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Backend_PackageUploadStream{}
	}
	return Backend_PackageUploadStream{Client: p.Interface().Client()}
}

func (s Backend_installPackage_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_installPackage_Results) SetStream(v Backend_PackageUploadStream) error {

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

// Backend_installPackage_Results_List is a list of Backend_installPackage_Results.
type Backend_installPackage_Results_List struct{ capnp.List }

// NewBackend_installPackage_Results creates a new list of Backend_installPackage_Results.
func NewBackend_installPackage_Results_List(s *capnp.Segment, sz int32) (Backend_installPackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Backend_installPackage_Results_List{}, err
	}
	return Backend_installPackage_Results_List{l}, nil
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
	if err != nil {
		return Backend_tryGetPackage_Params{}, err
	}
	return Backend_tryGetPackage_Params{st}, nil
}

func NewRootBackend_tryGetPackage_Params(s *capnp.Segment) (Backend_tryGetPackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_tryGetPackage_Params{}, err
	}
	return Backend_tryGetPackage_Params{st}, nil
}

func ReadRootBackend_tryGetPackage_Params(msg *capnp.Message) (Backend_tryGetPackage_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_tryGetPackage_Params{}, err
	}
	return Backend_tryGetPackage_Params{root.Struct()}, nil
}

func (s Backend_tryGetPackage_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_tryGetPackage_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_tryGetPackage_Params_List{}, err
	}
	return Backend_tryGetPackage_Params_List{l}, nil
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_tryGetPackage_Results{}, err
	}
	return Backend_tryGetPackage_Results{st}, nil
}

func NewRootBackend_tryGetPackage_Results(s *capnp.Segment) (Backend_tryGetPackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_tryGetPackage_Results{}, err
	}
	return Backend_tryGetPackage_Results{st}, nil
}

func ReadRootBackend_tryGetPackage_Results(msg *capnp.Message) (Backend_tryGetPackage_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_tryGetPackage_Results{}, err
	}
	return Backend_tryGetPackage_Results{root.Struct()}, nil
}

func (s Backend_tryGetPackage_Results) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_tryGetPackage_Results) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_tryGetPackage_Results) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return spk.Manifest{}, err
	}

	return spk.Manifest{Struct: p.Struct()}, nil

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

// Backend_tryGetPackage_Results_List is a list of Backend_tryGetPackage_Results.
type Backend_tryGetPackage_Results_List struct{ capnp.List }

// NewBackend_tryGetPackage_Results creates a new list of Backend_tryGetPackage_Results.
func NewBackend_tryGetPackage_Results_List(s *capnp.Segment, sz int32) (Backend_tryGetPackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Backend_tryGetPackage_Results_List{}, err
	}
	return Backend_tryGetPackage_Results_List{l}, nil
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
	if err != nil {
		return Backend_deletePackage_Params{}, err
	}
	return Backend_deletePackage_Params{st}, nil
}

func NewRootBackend_deletePackage_Params(s *capnp.Segment) (Backend_deletePackage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_deletePackage_Params{}, err
	}
	return Backend_deletePackage_Params{st}, nil
}

func ReadRootBackend_deletePackage_Params(msg *capnp.Message) (Backend_deletePackage_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deletePackage_Params{}, err
	}
	return Backend_deletePackage_Params{root.Struct()}, nil
}

func (s Backend_deletePackage_Params) PackageId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_deletePackage_Params) HasPackageId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deletePackage_Params) PackageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_deletePackage_Params_List{}, err
	}
	return Backend_deletePackage_Params_List{l}, nil
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
	if err != nil {
		return Backend_deletePackage_Results{}, err
	}
	return Backend_deletePackage_Results{st}, nil
}

func NewRootBackend_deletePackage_Results(s *capnp.Segment) (Backend_deletePackage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_deletePackage_Results{}, err
	}
	return Backend_deletePackage_Results{st}, nil
}

func ReadRootBackend_deletePackage_Results(msg *capnp.Message) (Backend_deletePackage_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deletePackage_Results{}, err
	}
	return Backend_deletePackage_Results{root.Struct()}, nil
}

// Backend_deletePackage_Results_List is a list of Backend_deletePackage_Results.
type Backend_deletePackage_Results_List struct{ capnp.List }

// NewBackend_deletePackage_Results creates a new list of Backend_deletePackage_Results.
func NewBackend_deletePackage_Results_List(s *capnp.Segment, sz int32) (Backend_deletePackage_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_deletePackage_Results_List{}, err
	}
	return Backend_deletePackage_Results_List{l}, nil
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
	if err != nil {
		return Backend_backupGrain_Params{}, err
	}
	return Backend_backupGrain_Params{st}, nil
}

func NewRootBackend_backupGrain_Params(s *capnp.Segment) (Backend_backupGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return Backend_backupGrain_Params{}, err
	}
	return Backend_backupGrain_Params{st}, nil
}

func ReadRootBackend_backupGrain_Params(msg *capnp.Message) (Backend_backupGrain_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_backupGrain_Params{}, err
	}
	return Backend_backupGrain_Params{root.Struct()}, nil
}

func (s Backend_backupGrain_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_backupGrain_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_backupGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_backupGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_backupGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return grain.GrainInfo{}, err
	}

	return grain.GrainInfo{Struct: p.Struct()}, nil

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
	if err != nil {
		return Backend_backupGrain_Params_List{}, err
	}
	return Backend_backupGrain_Params_List{l}, nil
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
	if err != nil {
		return Backend_backupGrain_Results{}, err
	}
	return Backend_backupGrain_Results{st}, nil
}

func NewRootBackend_backupGrain_Results(s *capnp.Segment) (Backend_backupGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_backupGrain_Results{}, err
	}
	return Backend_backupGrain_Results{st}, nil
}

func ReadRootBackend_backupGrain_Results(msg *capnp.Message) (Backend_backupGrain_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_backupGrain_Results{}, err
	}
	return Backend_backupGrain_Results{root.Struct()}, nil
}

// Backend_backupGrain_Results_List is a list of Backend_backupGrain_Results.
type Backend_backupGrain_Results_List struct{ capnp.List }

// NewBackend_backupGrain_Results creates a new list of Backend_backupGrain_Results.
func NewBackend_backupGrain_Results_List(s *capnp.Segment, sz int32) (Backend_backupGrain_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_backupGrain_Results_List{}, err
	}
	return Backend_backupGrain_Results_List{l}, nil
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
	if err != nil {
		return Backend_restoreGrain_Params{}, err
	}
	return Backend_restoreGrain_Params{st}, nil
}

func NewRootBackend_restoreGrain_Params(s *capnp.Segment) (Backend_restoreGrain_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return Backend_restoreGrain_Params{}, err
	}
	return Backend_restoreGrain_Params{st}, nil
}

func ReadRootBackend_restoreGrain_Params(msg *capnp.Message) (Backend_restoreGrain_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_restoreGrain_Params{}, err
	}
	return Backend_restoreGrain_Params{root.Struct()}, nil
}

func (s Backend_restoreGrain_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_restoreGrain_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_restoreGrain_Params) HasOwnerId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) OwnerIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_restoreGrain_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Backend_restoreGrain_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_restoreGrain_Params_List{}, err
	}
	return Backend_restoreGrain_Params_List{l}, nil
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
	if err != nil {
		return Backend_restoreGrain_Results{}, err
	}
	return Backend_restoreGrain_Results{st}, nil
}

func NewRootBackend_restoreGrain_Results(s *capnp.Segment) (Backend_restoreGrain_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_restoreGrain_Results{}, err
	}
	return Backend_restoreGrain_Results{st}, nil
}

func ReadRootBackend_restoreGrain_Results(msg *capnp.Message) (Backend_restoreGrain_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_restoreGrain_Results{}, err
	}
	return Backend_restoreGrain_Results{root.Struct()}, nil
}

func (s Backend_restoreGrain_Results) Info() (grain.GrainInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return grain.GrainInfo{}, err
	}

	return grain.GrainInfo{Struct: p.Struct()}, nil

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
	if err != nil {
		return Backend_restoreGrain_Results_List{}, err
	}
	return Backend_restoreGrain_Results_List{l}, nil
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
	if err != nil {
		return Backend_uploadBackup_Params{}, err
	}
	return Backend_uploadBackup_Params{st}, nil
}

func NewRootBackend_uploadBackup_Params(s *capnp.Segment) (Backend_uploadBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_uploadBackup_Params{}, err
	}
	return Backend_uploadBackup_Params{st}, nil
}

func ReadRootBackend_uploadBackup_Params(msg *capnp.Message) (Backend_uploadBackup_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_uploadBackup_Params{}, err
	}
	return Backend_uploadBackup_Params{root.Struct()}, nil
}

func (s Backend_uploadBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_uploadBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_uploadBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_uploadBackup_Params_List{}, err
	}
	return Backend_uploadBackup_Params_List{l}, nil
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
	if err != nil {
		return Backend_uploadBackup_Results{}, err
	}
	return Backend_uploadBackup_Results{st}, nil
}

func NewRootBackend_uploadBackup_Results(s *capnp.Segment) (Backend_uploadBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_uploadBackup_Results{}, err
	}
	return Backend_uploadBackup_Results{st}, nil
}

func ReadRootBackend_uploadBackup_Results(msg *capnp.Message) (Backend_uploadBackup_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_uploadBackup_Results{}, err
	}
	return Backend_uploadBackup_Results{root.Struct()}, nil
}

func (s Backend_uploadBackup_Results) Stream() util.ByteStream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Backend_uploadBackup_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_uploadBackup_Results) SetStream(v util.ByteStream) error {

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

// Backend_uploadBackup_Results_List is a list of Backend_uploadBackup_Results.
type Backend_uploadBackup_Results_List struct{ capnp.List }

// NewBackend_uploadBackup_Results creates a new list of Backend_uploadBackup_Results.
func NewBackend_uploadBackup_Results_List(s *capnp.Segment, sz int32) (Backend_uploadBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Backend_uploadBackup_Results_List{}, err
	}
	return Backend_uploadBackup_Results_List{l}, nil
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
	if err != nil {
		return Backend_downloadBackup_Params{}, err
	}
	return Backend_downloadBackup_Params{st}, nil
}

func NewRootBackend_downloadBackup_Params(s *capnp.Segment) (Backend_downloadBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Backend_downloadBackup_Params{}, err
	}
	return Backend_downloadBackup_Params{st}, nil
}

func ReadRootBackend_downloadBackup_Params(msg *capnp.Message) (Backend_downloadBackup_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_downloadBackup_Params{}, err
	}
	return Backend_downloadBackup_Params{root.Struct()}, nil
}

func (s Backend_downloadBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_downloadBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_downloadBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s Backend_downloadBackup_Params) SetBackupId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Backend_downloadBackup_Params) Stream() util.ByteStream {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return util.ByteStream{}
	}
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s Backend_downloadBackup_Params) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Backend_downloadBackup_Params) SetStream(v util.ByteStream) error {

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

// Backend_downloadBackup_Params_List is a list of Backend_downloadBackup_Params.
type Backend_downloadBackup_Params_List struct{ capnp.List }

// NewBackend_downloadBackup_Params creates a new list of Backend_downloadBackup_Params.
func NewBackend_downloadBackup_Params_List(s *capnp.Segment, sz int32) (Backend_downloadBackup_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Backend_downloadBackup_Params_List{}, err
	}
	return Backend_downloadBackup_Params_List{l}, nil
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
	if err != nil {
		return Backend_downloadBackup_Results{}, err
	}
	return Backend_downloadBackup_Results{st}, nil
}

func NewRootBackend_downloadBackup_Results(s *capnp.Segment) (Backend_downloadBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_downloadBackup_Results{}, err
	}
	return Backend_downloadBackup_Results{st}, nil
}

func ReadRootBackend_downloadBackup_Results(msg *capnp.Message) (Backend_downloadBackup_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_downloadBackup_Results{}, err
	}
	return Backend_downloadBackup_Results{root.Struct()}, nil
}

// Backend_downloadBackup_Results_List is a list of Backend_downloadBackup_Results.
type Backend_downloadBackup_Results_List struct{ capnp.List }

// NewBackend_downloadBackup_Results creates a new list of Backend_downloadBackup_Results.
func NewBackend_downloadBackup_Results_List(s *capnp.Segment, sz int32) (Backend_downloadBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_downloadBackup_Results_List{}, err
	}
	return Backend_downloadBackup_Results_List{l}, nil
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
	if err != nil {
		return Backend_deleteBackup_Params{}, err
	}
	return Backend_deleteBackup_Params{st}, nil
}

func NewRootBackend_deleteBackup_Params(s *capnp.Segment) (Backend_deleteBackup_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_deleteBackup_Params{}, err
	}
	return Backend_deleteBackup_Params{st}, nil
}

func ReadRootBackend_deleteBackup_Params(msg *capnp.Message) (Backend_deleteBackup_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deleteBackup_Params{}, err
	}
	return Backend_deleteBackup_Params{root.Struct()}, nil
}

func (s Backend_deleteBackup_Params) BackupId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_deleteBackup_Params) HasBackupId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_deleteBackup_Params) BackupIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_deleteBackup_Params_List{}, err
	}
	return Backend_deleteBackup_Params_List{l}, nil
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
	if err != nil {
		return Backend_deleteBackup_Results{}, err
	}
	return Backend_deleteBackup_Results{st}, nil
}

func NewRootBackend_deleteBackup_Results(s *capnp.Segment) (Backend_deleteBackup_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Backend_deleteBackup_Results{}, err
	}
	return Backend_deleteBackup_Results{st}, nil
}

func ReadRootBackend_deleteBackup_Results(msg *capnp.Message) (Backend_deleteBackup_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_deleteBackup_Results{}, err
	}
	return Backend_deleteBackup_Results{root.Struct()}, nil
}

// Backend_deleteBackup_Results_List is a list of Backend_deleteBackup_Results.
type Backend_deleteBackup_Results_List struct{ capnp.List }

// NewBackend_deleteBackup_Results creates a new list of Backend_deleteBackup_Results.
func NewBackend_deleteBackup_Results_List(s *capnp.Segment, sz int32) (Backend_deleteBackup_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Backend_deleteBackup_Results_List{}, err
	}
	return Backend_deleteBackup_Results_List{l}, nil
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
	if err != nil {
		return Backend_getUserStorageUsage_Params{}, err
	}
	return Backend_getUserStorageUsage_Params{st}, nil
}

func NewRootBackend_getUserStorageUsage_Params(s *capnp.Segment) (Backend_getUserStorageUsage_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Backend_getUserStorageUsage_Params{}, err
	}
	return Backend_getUserStorageUsage_Params{st}, nil
}

func ReadRootBackend_getUserStorageUsage_Params(msg *capnp.Message) (Backend_getUserStorageUsage_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_getUserStorageUsage_Params{}, err
	}
	return Backend_getUserStorageUsage_Params{root.Struct()}, nil
}

func (s Backend_getUserStorageUsage_Params) UserId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s Backend_getUserStorageUsage_Params) HasUserId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Backend_getUserStorageUsage_Params) UserIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return Backend_getUserStorageUsage_Params_List{}, err
	}
	return Backend_getUserStorageUsage_Params_List{l}, nil
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
	if err != nil {
		return Backend_getUserStorageUsage_Results{}, err
	}
	return Backend_getUserStorageUsage_Results{st}, nil
}

func NewRootBackend_getUserStorageUsage_Results(s *capnp.Segment) (Backend_getUserStorageUsage_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Backend_getUserStorageUsage_Results{}, err
	}
	return Backend_getUserStorageUsage_Results{st}, nil
}

func ReadRootBackend_getUserStorageUsage_Results(msg *capnp.Message) (Backend_getUserStorageUsage_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Backend_getUserStorageUsage_Results{}, err
	}
	return Backend_getUserStorageUsage_Results{root.Struct()}, nil
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
	if err != nil {
		return Backend_getUserStorageUsage_Results_List{}, err
	}
	return Backend_getUserStorageUsage_Results_List{l}, nil
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
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Params{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Params{st}, nil
}

func NewRootSandstormCoreFactory_getSandstormCore_Params(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Params{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Params{st}, nil
}

func ReadRootSandstormCoreFactory_getSandstormCore_Params(msg *capnp.Message) (SandstormCoreFactory_getSandstormCore_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Params{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Params{root.Struct()}, nil
}

func (s SandstormCoreFactory_getSandstormCore_Params) GrainId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s SandstormCoreFactory_getSandstormCore_Params) HasGrainId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCoreFactory_getSandstormCore_Params) GrainIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Params_List{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Params_List{l}, nil
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
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Results{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Results{st}, nil
}

func NewRootSandstormCoreFactory_getSandstormCore_Results(s *capnp.Segment) (SandstormCoreFactory_getSandstormCore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Results{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Results{st}, nil
}

func ReadRootSandstormCoreFactory_getSandstormCore_Results(msg *capnp.Message) (SandstormCoreFactory_getSandstormCore_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Results{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Results{root.Struct()}, nil
}

func (s SandstormCoreFactory_getSandstormCore_Results) Core() supervisor.SandstormCore {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return supervisor.SandstormCore{}
	}
	return supervisor.SandstormCore{Client: p.Interface().Client()}
}

func (s SandstormCoreFactory_getSandstormCore_Results) HasCore() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormCoreFactory_getSandstormCore_Results) SetCore(v supervisor.SandstormCore) error {

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

// SandstormCoreFactory_getSandstormCore_Results_List is a list of SandstormCoreFactory_getSandstormCore_Results.
type SandstormCoreFactory_getSandstormCore_Results_List struct{ capnp.List }

// NewSandstormCoreFactory_getSandstormCore_Results creates a new list of SandstormCoreFactory_getSandstormCore_Results.
func NewSandstormCoreFactory_getSandstormCore_Results_List(s *capnp.Segment, sz int32) (SandstormCoreFactory_getSandstormCore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormCoreFactory_getSandstormCore_Results_List{}, err
	}
	return SandstormCoreFactory_getSandstormCore_Results_List{l}, nil
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
