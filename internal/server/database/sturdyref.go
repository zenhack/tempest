package database

import (
	"crypto/rand"
	"crypto/sha256"

	"zenhack.net/go/util"
)

func GenToken() []byte {
	var token [32]byte
	_, err := rand.Read(token[:])
	util.Chkfatal(err)
	return token[:]
}

// Generate a random sturdyRef and return the hash of the token (losing the
// original token). Useful for stuff that goes in user keyrings, and thus
// doesn't need to have an actually-usable token.
func genLostToken() [sha256.Size]byte {
	return sha256.Sum256(GenToken())
}
