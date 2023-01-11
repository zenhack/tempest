package spk

import (
	"io"

	"zenhack.net/go/sandstorm/capnp/spk"

	"github.com/ulikunitz/xz"
	"capnproto.org/go/capnp/v3"
)

// Write an .spk into `dest`, using `archive` as the contents and `key` for
// signing. The archive must already contain the manifest.
func PackInto(dest io.Writer, key Key, archive spk.Archive) error {
	sig, err := key.signArchive(archive)
	if err != nil {
		return err
	}

	if _, err = dest.Write(spk.MagicNumber); err != nil {
		return err
	}

	compressedDest, err := xz.NewWriter(dest)
	if err != nil {
		return err
	}

	enc := capnp.NewEncoder(compressedDest)
	if err = enc.Encode(sig.Segment().Message()); err != nil {
		return err
	}
	if err = enc.Encode(archive.Segment().Message()); err != nil {
		return err
	}
	return compressedDest.Close()
}
