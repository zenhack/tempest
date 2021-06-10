package spk

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"os"

	"golang.org/x/crypto/nacl/sign"

	"zenhack.net/go/sandstorm/capnp/spk"
	"capnproto.org/go/capnp/v3"
)

var (
	ErrKeyNotFound  = errors.New("Key not found in keyring")
	ErrMalformedKey = errors.New("Key is malformed")
)

// The contents of a sandstorm keyring, typically stored at ~/.sandstorm-keyring
// or ~/.sandstorm/sandstorm-keyring.
type Keyring struct {
	keys []spk.KeyFile
}

// A package signing key
type Key spk.KeyFile

func (k Key) AppId() (AppId, error) {
	var ret AppId
	pubKey, err := spk.KeyFile(k).PublicKey()
	if err != nil {
		return ret, err
	}
	err = (&ret).UnmarshalBinary(pubKey)
	return ret, err
}

// Add the key to the keyring in the file at `path`, which is created if it does
// not already exist.
func (k Key) AddToFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return capnp.NewEncoder(f).Encode(k.Struct.Message())
}

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
	_, firstSeg, err := capnp.NewMessage(capnp.SingleSegment(nil))
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

// Get the key for the given app id from the keyring.
func (k Keyring) GetKey(appId AppId) (Key, error) {
	// simple linear search.
	for _, keyFile := range k.keys {
		pubKey, err := keyFile.PublicKey()
		if err != nil {
			// TODO: Don't lose err here; change ErrMalformedKey into
			// something that wraps the underlying error.
			return Key{}, ErrMalformedKey
		}
		if bytes.Compare(appId[:], pubKey) == 0 {
			return Key(keyFile), nil
		}
	}
	return Key{}, ErrKeyNotFound
}

// Load the sandstorm keyring from a named file.
func LoadKeyring(filename string) (Keyring, error) {
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

func (key Key) getPrivateKey() ([64]byte, error) {
	keyFile := spk.KeyFile(key)

	var ret [64]byte

	privKey, err := keyFile.PrivateKey()
	if err != nil {
		return ret, err
	}
	if len(privKey) != len(ret) {
		return ret, ErrMalformedKey
	}
	copy(ret[:], privKey)
	return ret, nil
}

// Sign an archive, which must be the root of its message.
func (key Key) signArchive(archive spk.Archive) (spk.Signature, error) {
	// Return this on errors:
	empty := spk.Signature{}

	keyFile := spk.KeyFile(key)
	pubKey, err := keyFile.PublicKey()
	if err != nil {
		return empty, err
	}
	privKey, err := key.getPrivateKey()
	if err != nil {
		return empty, err
	}

	hash := sha512.New()
	err = capnp.NewEncoder(hash).Encode(archive.Segment().Message())
	if err != nil {
		return empty, err
	}
	digest := hash.Sum(nil)
	sigbytes := sign.Sign(nil, digest, &privKey)

	_, sigSeg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return empty, err
	}
	sig, err := spk.NewRootSignature(sigSeg)
	if err != nil {
		return empty, err
	}
	if err = sig.SetPublicKey(pubKey); err != nil {
		return empty, err
	}
	if err = sig.SetSignature(sigbytes); err != nil {
		return empty, err
	}
	return sig, nil
}
