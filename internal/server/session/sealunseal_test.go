package session

import (
	"crypto/rand"
)

func randomStore() Store {
	var key [32]byte
	rand.Read(key[:])
	return NewStore([][32]byte{key})
}
