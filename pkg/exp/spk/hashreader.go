package spk

import (
	"hash"
	"io"
)

// TODO: move this to somewhere more sensible

type hashReader struct {
	Hash   hash.Hash
	Reader io.Reader
}

func (hr hashReader) Read(p []byte) (n int, err error) {
	n, err = hr.Reader.Read(p)
	hr.Hash.Write(p[:n])
	return
}
