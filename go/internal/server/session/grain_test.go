package session

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrainSealUnseal(t *testing.T) {
	cases := []GrainSession{
		{},
		{
			GrainId:   "RfLg3jgaJMXq4cqMpGonwL",
			SessionId: []byte("1234"),
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testGrainSealUnseal(t, c)
		})
	}
}

// TODO(cleanup): this is copypasta from testUserSealUnseal; factor out the common bits.
// this is a bit fiddly to do.
func testGrainSealUnseal(t *testing.T, sess GrainSession) {
	store := randomStore()
	data, err := sess.Seal(store)
	assert.NoError(t, err)
	var sessOut GrainSession

	assert.NoError(t, sessOut.Unseal(store, Payload{
		Data: data,
		// At some point, either this argument should just be the string,
		// or we should fail unsealing if this is wrong. For now it's
		// just ignored though:
		CookieName: "TODO",
	}))
}
