package spk

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	spk "zenhack.net/go/tempest/capnp/package"
	"zenhack.net/go/util/exn"
)

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
		lastMod := time.Unix(0, file.LastModificationTimeNs())
		throw(os.Chtimes(path, lastMod, lastMod))
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
