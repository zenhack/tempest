package session

import (
	"zenhack.net/go/sandstorm-next/capnp/private/cookie"
)

type UserSession struct {
	SessionId  []byte
	Credential UserSessionCredential
}

type UserSessionCredential struct {
	Type     string
	ScopedId string
}

func (sess *UserSession) Unseal(store Store, payload Payload) error {
	return unseal(sess, cookie.UserSession_TypeID, store, payload)
}

func (sess UserSession) Seal(store Store) (string, error) {
	return seal(
		sess,
		cookie.UserSession_TypeID,
		cookie.NewRootUserSession,
		store,
	)
}

func (sess UserSession) CookieName() string {
	return "sandstorm-user-session"
}
