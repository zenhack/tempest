package session

import (
	"encoding/hex"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"zenhack.net/go/sandstorm-next/capnp/private/cookie"
	"zenhack.net/go/util/exn"
)

type Payload struct {
	CookieName string
	Data       string
}

type UserSession struct {
	SessionId  []byte
	Credential struct {
		Type     string
		ScopedId string
	}
}

func (sess *UserSession) Unseal(store Store, payload Payload) error {
	return exn.Try0(func(throw func(error)) {
		buf, err := hex.DecodeString(payload.Data)
		throw(err)
		var typeId uint64 = cookie.UserSession_TypeID
		msg, err := store.aead.unsealCapnp(buf, typeId)
		throw(err)
		root, err := msg.Root()
		throw(err)
		throw(pogs.Extract(sess, typeId, root.Struct()))
	})
}

func (sess UserSession) Seal(store Store) (string, error) {
	return exn.Try(func(throw func(error)) string {
		msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		throw(err)
		root, err := cookie.NewRootUserSession(seg)
		throw(err)
		var typeId uint64 = cookie.UserSession_TypeID
		pogs.Insert(typeId, capnp.Struct(root), &sess)
		return hex.EncodeToString(store.aead.sealCapnp(msg, typeId))
	})
}
