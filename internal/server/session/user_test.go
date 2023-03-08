package session

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"zenhack.net/go/tempest/internal/server/types"
)

func TestUserSealUnseal(t *testing.T) {
	cases := []UserSession{
		{},
		{
			SessionId: []byte("1234"),
			Credential: types.Credential{
				Type:     "dev",
				ScopedId: "Alice Dev Admin",
			},
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testUserSealUnseal(t, c)
		})
	}
}

func testUserSealUnseal(t *testing.T, sess UserSession) {
	store := randomStore()
	data, err := sess.Seal(store)
	assert.NoError(t, err)
	var sessOut UserSession

	assert.NoError(t, sessOut.Unseal(store, Payload{
		Data: data,
		// At some point, either this argument should just be the string,
		// or we should fail unsealing if this is wrong. For now it's
		// just ignored though:
		CookieName: "TODO",
	}))
}
