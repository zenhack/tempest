package database

import (
	"crypto/rand"

	"zenhack.net/go/util"
)

func GenToken() []byte {
	var token [32]byte
	_, err := rand.Read(token[:])
	util.Chkfatal(err)
	return token[:]
}
