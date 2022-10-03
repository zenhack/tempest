// `sandstorm-sandbox-agent` runs inside the grain's sandbox, and is the first
// program executed during grain startup. Its file descriptor #3 is a socket
// over which we can speak capnp to the sandstorm server outside the sandbox.
//
// Any APIs available to the grain which don't actually need privileges the grain
// doesn't have should ideally be implemented here; this helps us minimize attack
// surface.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"golang.org/x/sys/unix"
	"zenhack.net/go/sandstorm-next/go/internal/util"
	"zenhack.net/go/sandstorm/capnp/spk"
)

func buildCmd(cmd spk.Manifest_Command) (*exec.Cmd, error) {
	argv, err := cmd.Argv()
	if err != nil {
		return nil, err
	}
	var args []string
	for i := 0; i < argv.Len(); i++ {
		arg, err := argv.At(i)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	if len(args) == 0 {
		return nil, fmt.Errorf("len(cmd.argv) == 0")
	}
	environ, err := cmd.Environ()
	if err != nil {
		return nil, err
	}
	var env []string
	for i := 0; i < environ.Len(); i++ {
		kv := environ.At(i)
		k, err := kv.Key()
		if err != nil {
			return nil, err
		}
		v, err := kv.Value()
		if err != nil {
			return nil, err
		}
		env = append(env, k+"="+v)
	}

	ret := exec.Command(args[0], args[1:]...)
	ret.Env = env
	return ret, nil
}

func main() {
	data, err := ioutil.ReadFile("/sandstorm-manifest")
	util.Chkfatal(err)
	msg, err := capnp.Unmarshal(data)
	util.Chkfatal(err)
	manifest, err := spk.ReadRootManifest(msg)
	util.Chkfatal(err)
	appTitle, err := manifest.AppTitle()
	util.Chkfatal(err)
	text, err := appTitle.DefaultText()
	util.Chkfatal(err)
	fmt.Println("App title: ", text)
	spkCmd, err := manifest.ContinueCommand()
	util.Chkfatal(err)
	cmd, err := buildCmd(spkCmd)
	util.Chkfatal(err)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	util.Chkfatal(err)
	parentSock := os.NewFile(uintptr(fds[0]), "Parent socket")
	defer parentSock.Close()
	childSock := os.NewFile(uintptr(fds[1]), "Child socket")
	cmd.ExtraFiles = []*os.File{childSock}
}
