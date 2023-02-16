package servermain

import (
	"context"

	"capnproto.org/go/capnp/v3/exc"
	"zenhack.net/go/tempest/capnp/grain"
)

type sandstormApiImpl struct{}

func (sandstormApiImpl) DeprecatedPublish(context.Context, grain.SandstormApi_deprecatedPublish) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "unimplemented")
}

func (sandstormApiImpl) DeprecatedRegisterAction(context.Context, grain.SandstormApi_deprecatedRegisterAction) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "unimplemented")
}

func (sandstormApiImpl) ShareCap(context.Context, grain.SandstormApi_shareCap) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}

func (sandstormApiImpl) ShareView(context.Context, grain.SandstormApi_shareView) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) Save(context.Context, grain.SandstormApi_save) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) Restore(context.Context, grain.SandstormApi_restore) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) Drop(context.Context, grain.SandstormApi_drop) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) Deleted(context.Context, grain.SandstormApi_deleted) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) StayAwake(context.Context, grain.SandstormApi_stayAwake) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) BackgroundActivity(context.Context, grain.SandstormApi_backgroundActivity) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) GetIdentityId(context.Context, grain.SandstormApi_getIdentityId) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
func (sandstormApiImpl) Schedule(context.Context, grain.SandstormApi_schedule) error {
	return exc.New(exc.Unimplemented, "SandstormApi", "TODO")
}
