package spk

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"capnproto.org/go/capnp/v3"
	"github.com/ulikunitz/xz"
	"golang.org/x/sys/unix"
	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/util/exn"
)

var (
	ErrArchiveTooLarge = errors.New("spk archive is too large")
	ErrNoMagicNumber   = errors.New("spk file does not start with magic number")
)

type PackageHash [sha256.Size]byte

// ID returns the package ID based on the hash. This string is used in various places:
// as a directory name, as part of the package's URL in the app market, and others.
// The value is the first 128 bits, hex-encoded.
func (ph PackageHash) ID() string {
	return hex.EncodeToString(ph[:16])
}

// UnpackSpk reads an spk file from r and unpacks its contents to a newly created
// director at path, after verifying the package's signature. Returns the app's
// ID.
//
// This creates a tempory file under tmpDir, which is deleted before the function
// returns.
func UnpackSpk(path string, tmpDir string, r io.Reader) (AppID, PackageHash, error) {
	return exn.Try2(func(throw exn.Thrower) (AppID, PackageHash) {
		// For computing the package id:
		hr := hashReader{
			Hash:   sha256.New(),
			Reader: r,
		}
		var magic [8]byte
		_, err := io.ReadFull(hr, magic[:])
		throw(err)
		if !bytes.Equal(magic[:], spk.MagicNumber) {
			throw(ErrNoMagicNumber)
		}
		xr, err := xz.NewReader(hr)
		throw(err)
		dec := capnp.NewDecoder(xr)
		sigMsg, err := dec.Decode()
		throw(err)
		sig, err := spk.ReadRootSignature(sigMsg)
		throw(err)

		appID, signatureDigest, err := VerifySignature(sig)
		throw(err)

		h := sha512.New()

		tmpFile, err := os.CreateTemp(tmpDir, "spk-archive-*")
		throw(err)
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		archiveSize, err := io.Copy(io.MultiWriter(h, tmpFile), xr)
		throw(err)
		dataDigest := h.Sum(nil)
		if !bytes.Equal(signatureDigest, dataDigest) {
			throw(fmt.Errorf(
				"package verification failed: archive hash is sha512-%x but signature is for sha512-%x",
				dataDigest,
				signatureDigest,
			))
		}

		if archiveSize > math.MaxInt {
			throw(ErrArchiveTooLarge)
		}
		archiveBuf, err := unix.Mmap(int(tmpFile.Fd()), 0, int(archiveSize), unix.PROT_READ, unix.MAP_SHARED)
		throw(err)
		defer unix.Munmap(archiveBuf)

		archiveMsg, err := capnp.Unmarshal(archiveBuf)
		throw(err)

		// A couple times the size of the message; the default of 64MiB
		// is way too small:
		archiveMsg.ResetReadLimit(uint64(archiveSize) * 4)

		archive, err := spk.ReadRootArchive(archiveMsg)
		throw(unpackArchive(path, archive))
		pkgHash := ([sha256.Size]byte)(hr.Hash.Sum(nil))
		return appID, pkgHash
	})
}

func unpackArchive(path string, archive spk.Archive) error {
	files, err := archive.Files()
	if err != nil {
		return err
	}
	return unpackDirectory(path, files)
}

func unpackFile(path string, file spk.Archive_File) error {
	return exn.Try0(func(throw exn.Thrower) {
		switch file.Which() {
		case spk.Archive_File_Which_regular:
			data, err := file.Regular()
			throw(err)
			throw(os.WriteFile(path, data, 0644))
		case spk.Archive_File_Which_executable:
			data, err := file.Executable()
			throw(err)
			throw(os.WriteFile(path, data, 0755))
		case spk.Archive_File_Which_symlink:
			target, err := file.Symlink()
			throw(err)
			throw(os.Symlink(target, path))
		case spk.Archive_File_Which_directory:
			files, err := file.Directory()
			throw(err)
			throw(unpackDirectory(path, files))
		}
		if file.Which() != spk.Archive_File_Which_symlink {
			lastMod := time.Unix(0, file.LastModificationTimeNs())
			throw(os.Chtimes(path, lastMod, lastMod))
		}
	})
}

func unpackDirectory(path string, files spk.Archive_File_List) error {
	return exn.Try0(func(throw exn.Thrower) {
		throw(os.Mkdir(path, 0755))

		// Track what filenames we've seen, so we can reject duplicates:
		seen := make(map[string]struct{}, files.Len())

		for i := 0; i < files.Len(); i++ {
			file := files.At(i)
			name, err := file.Name()
			throw(err)
			if !isValidFileName(name) {
				throw(fmt.Errorf("invalid file name: %q", name))
			}
			if _, ok := seen[name]; ok {
				throw(fmt.Errorf(
					"name %q appears more than once in the same directory.",
					name,
				))
			}
			seen[name] = struct{}{}
			throw(unpackFile(filepath.Join(path, name), file))
		}
	})
}

func isValidFileName(name string) bool {
	return name != "" &&
		name != "." &&
		name != ".." &&
		!strings.Contains(name, "/")
}
