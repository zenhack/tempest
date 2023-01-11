package spk

import (
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	spk "zenhack.net/go/tempest/capnp/package"
)

const (
	// Path to the capnp schema that ship with Sandstorm, assuming sandstorm
	// is installed in /opt/sandstorm.
	SandstormCapnpPath = "/opt/sandstorm/latest/usr/include"
)

// Read the package definition from a textual pkgdef on disk.
// The `capnp` executable must be in PATH.
//
// Parameters:
//
// - file: The schema file to read
// - variable: the name of the variable in the file defining the package definition.
// - extraPaths: a list of extra directories to search for capnproto schema.
//
// A typical use of this would be:
//
// ReadPackageDefinition("sandstorm-pkgdef.capnp", "pkgdef", []string{SandstormCapnpPath})
func ReadPackageDefinition(file, variable string, extraPaths []string) (spk.PackageDefinition, error) {
	empty := spk.PackageDefinition{} // return on errors

	args := []string{"eval", "--binary"}
	for _, path := range extraPaths {
		args = append(args, "-I", path)
	}
	args = append(args, file, variable)
	cmd := exec.Command("capnp", args...)
	cmd.Stderr = os.Stderr
	pkgDefBytes, err := cmd.Output()
	if err != nil {
		return empty, err
	}
	msg, err := capnp.Unmarshal(pkgDefBytes)
	if err != nil {
		return empty, err
	}
	return spk.ReadRootPackageDefinition(msg)
}
