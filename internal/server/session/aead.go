package session

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"

	"capnproto.org/go/capnp/v3"
	"golang.org/x/crypto/chacha20poly1305"
	"zenhack.net/go/util"
)

const (
	nonceSize = chacha20poly1305.NonceSizeX
)

var (
	ErrPayloadTooShort = errors.New("payload too short; not enough space for nonce")
)

// capnpAEAD is a higher-lvel wrapper around a cipher.AEAD, which is used to encrypt/authenticate
// capnproto values. The underlying AEAD is always XChaCha20-Poly1305, nonces are chosen randomly,
// and the capnpAEAD takes care of bundling & unbundling the ciphertext & nonce into a single
// []byte payload.
type capnpAEAD struct {
	aead cipher.AEAD
}

// newCapnpAEAD returns a new capnpAEAD using the provided secret key.
func newCapnpAEAD(key [32]byte) capnpAEAD {
	aead, err := chacha20poly1305.NewX(key[:])
	util.Chkfatal(err)
	return capnpAEAD{aead: aead}
}

// typeIDToAD converts a numeric type ID to a byte array, suitable for passing in as
// additional data.
func typeIDToAD(id uint64) [8]byte {
	var ret [8]byte
	binary.LittleEndian.PutUint64(ret[:], id)
	return ret
}

// sealCapnp encrypts & authenticates a capnproto message, whose root object's
// type is given as typeID.
func (a capnpAEAD) sealCapnp(msg *capnp.Message, typeID uint64) []byte {
	plaintext, err := msg.Marshal()
	util.Chkfatal(err)
	ad := typeIDToAD(typeID)
	var nonce [nonceSize]byte
	_, err = rand.Read(nonce[:])
	util.Chkfatal(err)
	ret := a.aead.Seal(nil, nonce[:], plaintext, ad[:])
	return append(nonce[:], ret...)
}

// unsealCapnp decryptes & verifies the payload as a capnproto's message, whose
// root object must have the given type id; verification will fail if this is
// different than what was passed in to sealCapnp.
func (a capnpAEAD) unsealCapnp(payload []byte, typeID uint64) (*capnp.Message, error) {
	if len(payload) < nonceSize {
		return nil, ErrPayloadTooShort
	}
	nonce, ciphertext := payload[:nonceSize], payload[nonceSize:]
	ad := typeIDToAD(typeID)
	plaintext, err := a.aead.Open(nil, nonce, ciphertext, ad[:])
	if err != nil {
		return nil, err
	}
	return capnp.Unmarshal(plaintext)
}
