package session

import (
	"zenhack.net/go/tempest/internal/capnp/cookie"
	"zenhack.net/go/tempest/internal/common/types"
)

type GrainSession struct {
	GrainID   types.GrainID `capnp:"grainId"`
	SessionID []byte        `capnp:"sessionId"`
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
