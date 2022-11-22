package main

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/sandstorm-next/capnp/external"
	"zenhack.net/go/sandstorm-next/go/internal/database"
)

type externalApiImpl struct {
	db            database.DB
	userSessionId string // "" if there is no logged-in user session.
}

func (api externalApiImpl) GetLoginSession(ctx context.Context, p external.ExternalApi_getLoginSession) error {
	return capnp.Unimplemented(
		"TODO: implement getLoginSession() (session id: " + api.userSessionId + ")",
	)
}
