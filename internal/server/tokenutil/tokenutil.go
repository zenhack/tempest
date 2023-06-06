package tokenutil

import (
	"crypto/rand"
	"encoding/base64"

	"zenhack.net/go/util"
)

func GenToken() []byte {
	var token [32]byte
	_, err := rand.Read(token[:])
	util.Chkfatal(err)
	return token[:]
}

func Gen128() []byte {
	return GenToken()[:16]
}

func Gen128Base64() string {
	return base64.RawURLEncoding.EncodeToString(Gen128())
}
