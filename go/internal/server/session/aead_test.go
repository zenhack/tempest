package session

import (
	"crypto/rand"
	"testing"

	"capnproto.org/go/capnp/v3"
	"github.com/stretchr/testify/assert"
)

func TestSealUnsealEmptyMessage(t *testing.T) {
	msg, _, err := capnp.NewMessage(capnp.SingleSegment(nil))
	assert.NoError(t, err)
	testSealUnsealCapnp(t, msg, 4)
}

func testSealUnsealCapnp(t *testing.T, msg *capnp.Message, typeId uint64) {
	var key [32]byte
	rand.Read(key[:])
	aead := newCapnpAEAD(key)

	// Seal & unseal. There should be no errors, and the value
	// should be the same.
	payload := aead.sealCapnp(msg, typeId)
	msgOut, err := aead.unsealCapnp(payload, typeId)
	assert.NoError(t, err)

	data, err := msg.Marshal()
	assert.NoError(t, err)
	dataOut, err := msgOut.Marshal()
	assert.NoError(t, err)
	assert.Equal(t, data, dataOut)

	// Unsealing with the wrong type id should fail.
	_, err = aead.unsealCapnp(payload, typeId+1)
	assert.NotNil(t, err)
}
