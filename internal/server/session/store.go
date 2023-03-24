package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/util/exn"
)

type Store struct {
	aead capnpAEAD
}

func GetKeys() (keys [][32]byte, err error) {
	const path = config.Localstatedir + "/sandstorm/session-key"
	data, err := os.ReadFile(config.Localstatedir + "/sandstorm/session-key")
	if os.IsNotExist(err) {
		data = make([]byte, 64)
		rand.Read(data)
		err = os.WriteFile(path, data, 0600)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	ret := make([][32]byte, 2)
	copy(ret[0][:], data[:32])
	copy(ret[1][:], data[32:])
	return ret, nil
}

func NewStore(keys [][32]byte) Store {
	// TODO: use other keys for decryption, to allow rotation.
	return Store{aead: newCapnpAEAD(keys[0])}
}

type Payload struct {
	CookieName string
	Data       string
}

func (p Payload) ToCookie(isHttps bool) *http.Cookie {
	return &http.Cookie{
		Name:     p.CookieName,
		Value:    p.Data,
		Path:     "/",
		Secure:   isHttps,
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

func seal[T ~capnp.StructKind, U any](
	val U,
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

type CookieReader interface {
	Unseal(Store, Payload) error
	CookieName() string
}

type CookieWriter interface {
	Seal(Store) (string, error)
	CookieName() string
}

func ReadCookie[T CookieReader](store Store, req *http.Request, val T) error {
	c, err := req.Cookie(val.CookieName())
	if err != nil {
		return err
	}
	return val.Unseal(store, Payload{
		CookieName: c.Name,
		Data:       c.Value,
	})
}

func WriteCookie[T CookieWriter](store Store, req *http.Request, w http.ResponseWriter, val T) error {
	data, err := val.Seal(store)
	if err != nil {
		return err
	}
	http.SetCookie(w, Payload{
		CookieName: val.CookieName(),
		Data:       data,
	}.ToCookie(req.URL.Scheme == "https"))
	return nil
}
