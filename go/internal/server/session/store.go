package session

type Store struct {
	aead capnpAEAD
}

func NewStore(key [32]byte) Store {
	return Store{aead: newCapnpAEAD(key)}
}
