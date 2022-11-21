package main

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/sandstorm-next/capnp/external"
)

type externalApiImpl struct {
}

func (externalApiImpl) GetLoginSession(ctx context.Context, p external.ExternalApi_getLoginSession) error {
	return capnp.Unimplemented("TODO: implement getLoginSession()")
}
