package session

import (
	"zenhack.net/go/tempest/internal/capnp/cookie"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/tokenutil"
)

type UserSession struct {
	SessionID  []byte `capnp:"sessionId"`
	Credential types.Credential
}

func GenSessionID() []byte {
	return tokenutil.GenToken()
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
