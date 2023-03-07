package servermain

import (
	"capnproto.org/go/capnp/v3/schemas"
	websession "zenhack.net/go/tempest/capnp/web-session"
)

func init() {
	websession.RegisterSchema(schemas.DefaultRegistry)
}
