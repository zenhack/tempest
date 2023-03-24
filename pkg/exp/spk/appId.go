package spk

import (
	"encoding/base32"
	"errors"
)

var (
	// The base32 alphabet used by Sandstorm for app-ids/public keys.
	SandstormBase32Encoding = base32.NewEncoding("0123456789acdefghjkmnpqrstuvwxyz").
				WithPadding(base32.NoPadding)

	ErrBadKeyLength = errors.New("invalid app id: wrong length")
)

// An app id/public key
type AppId [32]byte

func (id AppId) String() string {
	return SandstormBase32Encoding.EncodeToString(id[:])
}

func (id *AppId) UnmarshalText(text []byte) error {
	n, err := SandstormBase32Encoding.Decode(id[:], text)
	if err != nil {
		return err
	}
	if n != len(id[:]) {
		return ErrBadKeyLength
	}
	return nil
}

func (id AppId) MarshalText() (text []byte, err error) {
	return []byte(id.String()), nil
}

func (id *AppId) UnmarshalBinary(data []byte) error {
	if len(data) != len(id[:]) {
		return ErrBadKeyLength
	}
	copy(id[:], data)
	return nil
}

func (id AppId) MarshalBinary() ([]byte, error) {
	return id[:], nil
}
