package util

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"testing"
	"testing/quick"
	capnp "zenhack.net/go/sandstorm/capnp/util"
)

type writeNoopCloser struct {
	io.Writer
}

func (w writeNoopCloser) Close() error {
	return nil
}

// compute the md5 sum of p both directly and by writing it through
// the bytestream interface, returns true if the results are the same.
// We take a *testing.T for the logging facilities, but it is never
// flagged as failing.
func checkMd5(t *testing.T, p []byte) bool {
	if p == nil {
		// We don't want to deal with this
		return true
	}
	hash := md5.New()
	hash.Write(p)
	directDigest := fmt.Sprintf("%x", hash.Sum([]byte{}))

	hash.Reset()
	t.Logf("Direct digest: %q\n", directDigest)

	bsServer := &WriteCloserByteStream{writeNoopCloser{hash}}
	bsClient := ByteStreamWriteCloser{
		Ctx: context.Background(),
		Obj: capnp.ByteStream_ServerToClient(bsServer),
	}
	n, err := bsClient.Write(p)
	if err != nil || n != len(p) {
		t.Logf("Error: %v", err)
		return false
	}
	_, err = bsClient.Obj.Done(
		bsClient.Ctx,
		func(p capnp.ByteStream_done_Params) error {
			return nil
		}).Struct()
	if err != nil {
		t.Logf("Error: %v", err)
		return false
	}

	rpcDigest := fmt.Sprintf("%x", hash.Sum([]byte{}))
	t.Logf("RPC digest: %q\n", rpcDigest)

	return rpcDigest == directDigest
}

func TestMd5(t *testing.T) {
	err := quick.Check(func(p []byte) bool {
		return checkMd5(t, p)
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
}
