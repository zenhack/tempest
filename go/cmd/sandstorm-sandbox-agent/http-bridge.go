package main

import (
	"context"
	"errors"

	"zenhack.net/go/sandstorm-next/capnp/http"
)

type httpBridge struct {
	portNo int
}

func (httpBridge) Request(context.Context, http.Server_request) error {
	return errors.New("not yet implemented.")
}
