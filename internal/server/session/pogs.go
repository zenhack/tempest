package session

import (
	"capnproto.org/go/capnp/v3/schemas"
	"zenhack.net/go/tempest/internal/capnp/cookie"
)

func init() {
	cookie.RegisterSchema(schemas.DefaultRegistry)
}
