package servermain

import (
	"context"

	"capnproto.org/go/capnp/v3/exc"
	"zenhack.net/go/sandstorm/capnp/grain"
)

type sessionCtxImpl struct{}

func (sessionCtxImpl) GetSharedPermissions(context.Context, grain.SessionContext_getSharedPermissions) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) TieToUser(context.Context, grain.SessionContext_tieToUser) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) Offer(context.Context, grain.SessionContext_offer) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) Request(context.Context, grain.SessionContext_request) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) ClaimRequest(context.Context, grain.SessionContext_claimRequest) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) FulfillRequest(context.Context, grain.SessionContext_fulfillRequest) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) Close(context.Context, grain.SessionContext_close) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) OpenView(context.Context, grain.SessionContext_openView) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}

func (sessionCtxImpl) Activity(context.Context, grain.SessionContext_activity) error {
	return exc.New(exc.Unimplemented, "sessionCtxImpl", "TODO")
}
