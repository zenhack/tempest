package servermain

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/sandstorm-next/go/internal/database"
)

type externalApiImpl struct {
	db          database.DB
	userSession userSession
}

func (api externalApiImpl) GetLoginSession(ctx context.Context, p external.ExternalApi_getLoginSession) error {
	return capnp.Unimplemented(
		"TODO: implement getLoginSession() (session id: " + api.userSession.Id() + ")",
	)
}
