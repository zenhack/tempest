package session

import "zenhack.net/go/tempest/internal/capnp/cookie"

type GrainSession struct {
	GrainId   string
	SessionId []byte
}

func (sess *GrainSession) Unseal(store Store, payload Payload) error {
	return unseal(sess, cookie.GrainSession_TypeID, store, payload)
}

func (sess GrainSession) Seal(store Store) (string, error) {
	return seal(
		sess,
		cookie.GrainSession_TypeID,
		cookie.NewRootGrainSession,
		store,
	)
}

func (sess GrainSession) CookieName() string {
	return "sandstorm-grain-session"
}
