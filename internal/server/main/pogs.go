package servermain

import (
	"capnproto.org/go/capnp/v3/schemas"
	websession "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/internal/capnp/cookie"
)

func init() {
	websession.RegisterSchema(schemas.DefaultRegistry)
	cookie.RegisterSchema(schemas.DefaultRegistry)
}
