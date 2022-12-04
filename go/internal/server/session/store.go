package session

import (
	"encoding/hex"
	"net/http"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"zenhack.net/go/util/exn"
)

type Store struct {
	aead capnpAEAD
}

func NewStore(key [32]byte) Store {
	return Store{aead: newCapnpAEAD(key)}
}

type Payload struct {
	CookieName string
	Data       string
}

func (p Payload) ToCookie() *http.Cookie {
	return &http.Cookie{
		Name:     p.CookieName,
		Value:    p.Data,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func unseal(val any, typeId uint64, store Store, payload Payload) error {
	return exn.Try0(func(throw func(error)) {
		buf, err := hex.DecodeString(payload.Data)
		throw(err)
		msg, err := store.aead.unsealCapnp(buf, typeId)
		throw(err)
		root, err := msg.Root()
		throw(err)
		throw(pogs.Extract(val, typeId, root.Struct()))
	})
}

func seal[T ~capnp.StructKind](
	val any,
	typeId uint64,
	NewRoot func(*capnp.Segment) (T, error),
	store Store,
) (string, error) {
	return exn.Try(func(throw func(error)) string {
		msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		throw(err)
		root, err := NewRoot(seg)
		throw(err)
		throw(pogs.Insert(typeId, capnp.Struct(root), &val))
		return hex.EncodeToString(store.aead.sealCapnp(msg, typeId))
	})
}
