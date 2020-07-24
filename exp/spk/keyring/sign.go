package keyring

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"os"

	"golang.org/x/crypto/nacl/sign"

	"zenhack.net/go/sandstorm/capnp/spk"
	"zombiezen.com/go/capnproto2"
)

var (
	ErrKeyNotFound = errors.New("Key not found in keyring")
)

// The contents of a sandstorm keyring, typically stored at ~/.sandstorm-keyring
// or ~/.sandstorm/sandstorm-keyring.
type Keyring struct {
	keys []spk.KeyFile
}

// A package signing key
type Key spk.KeyFile

// Generate a new signing key. It will be the root object of its own message.
// The argument is a cryptographic random number generator. Defaults to
// crypto/rand.Reader if nil.
func GenerateKey(r io.Reader) (Key, error) {
	if r == nil {
		r = rand.Reader
	}
	pubKey, privKey, err := sign.GenerateKey(r)
	if err != nil {
		return Key{}, err
	}
	_, firstSeg, err := capnp.NewMessage(capnp.SingleSegment([]byte{}))
	if err != nil {
		return Key{}, err
	}
	key, err := spk.NewRootKeyFile(firstSeg)
	if err != nil {
		return Key{}, err
	}
	err = key.SetPublicKey(pubKey[:])
	if err != nil {
		return Key{}, err
	}
	err = key.SetPrivateKey(privKey[:])
	return Key(key), err
}

// Get a key from the keyring. The argument is the public part of the desired key
// (which is also the appId, after base32 decoding).
func (k Keyring) GetKey(targetPubKey []byte) (spk.KeyFile, error) {
	// simple linear search.
	for _, keyFile := range k.keys {
		pubKey, err := keyFile.PublicKey()
		if err != nil {
			return spk.KeyFile{}, err
		}
		if bytes.Compare(targetPubKey, pubKey) == 0 {
			return keyFile, nil
		}
	}
	return spk.KeyFile{}, ErrKeyNotFound
}

// Load the sandstorm keyring from a named file.
func Load(filename string) (Keyring, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Keyring{}, err
	}
	defer file.Close()

	ret := Keyring{}
	dec := capnp.NewDecoder(file)

	for {
		msg, err := dec.Decode()
		switch err {
		case nil:
			keyFile, err := spk.ReadRootKeyFile(msg)
			if err != nil {
				return Keyring{}, err
			}
			ret.keys = append(ret.keys, keyFile)
		case io.EOF:
			return ret, nil
		default:
			return Keyring{}, err
		}
	}
}

// Compute the signature of a package, given the raw bytes of the archive
// message. Returns the raw bytes of the signature message.
func signatureMessage(key spk.KeyFile, archiveBytes []byte) ([]byte, error) {
	pubKey, err := key.PublicKey()
	if err != nil {
		return nil, err
	}

	privKey, err := key.PrivateKey()
	if err != nil {
		return nil, err
	}

	// the go nacl library expects an array, not a slice:
	var naclPrivKey [64]byte
	copy(naclPrivKey[:], privKey)

	sigMsg, sigSeg, err := capnp.NewMessage(capnp.SingleSegment([]byte{}))
	if err != nil {
		return nil, err
	}
	sig, err := spk.NewRootSignature(sigSeg)
	if err != nil {
		return nil, err
	}

	hash := sha512.Sum512(archiveBytes)
	sigbytes := sign.Sign([]byte{}, hash[:], &naclPrivKey)

	if err = sig.SetPublicKey(pubKey); err != nil {
		return nil, err
	}
	if err = sig.SetSignature(sigbytes); err != nil {
		return nil, err
	}

	return sigMsg.Marshal()
}
