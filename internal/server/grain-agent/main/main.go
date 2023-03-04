// `tempest-grain-agent` runs inside the grain's sandbox, and is the first
// program executed during grain startup. Its file descriptor #3 is a socket
// over which we can speak capnp to the sandstorm server outside the sandbox.
//
// Any APIs available to the grain which don't actually need privileges the grain
// doesn't have should ideally be implemented here; this helps us minimize attack
// surface.
package grainagentmain

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"github.com/apex/log"
	spk "zenhack.net/go/tempest/capnp/package"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
	"zenhack.net/go/util"
)

type Command struct {
	Args []string
	Env  []string
}

func (c Command) ToOsCmd() *exec.Cmd {
	ret := exec.Command(c.Args[0], c.Args[1:]...)
	ret.Env = c.Env
	return ret
}

func parseCmd(cmd spk.Manifest_Command) (Command, error) {
	argv, err := cmd.Argv()
	if err != nil {
		return Command{}, err
	}
	var args []string
	for i := 0; i < argv.Len(); i++ {
		arg, err := argv.At(i)
		if err != nil {
			return Command{}, err
		}
		args = append(args, arg)
	}
	if len(args) == 0 {
		return Command{}, fmt.Errorf("len(cmd.argv) == 0")
	}
	environ, err := cmd.Environ()
	if err != nil {
		return Command{}, err
	}
	var env []string
	for i := 0; i < environ.Len(); i++ {
		kv := environ.At(i)
		k, err := kv.Key()
		if err != nil {
			return Command{}, err
		}
		v, err := kv.Value()
		if err != nil {
			return Command{}, err
		}
		env = append(env, k+"="+v)
	}
	return Command{
		Args: args,
		Env:  env,
	}, nil

}

func Main() {
	log.SetLevel(log.DebugLevel)
	lg := log.Log

	data, err := ioutil.ReadFile("/sandstorm-manifest")
	util.Chkfatal(err)
	msg, err := capnp.Unmarshal(data)
	util.Chkfatal(err)
	manifest, err := spk.ReadRootManifest(msg)
	util.Chkfatal(err)
	appTitle, err := manifest.AppTitle()
	util.Chkfatal(err)
	appTitleText, err := appTitle.DefaultText()
	util.Chkfatal(err)

	if len(os.Args) < 2 {
		lg.Fatal("Too few arugments")
	}
	buf, err := base64.StdEncoding.DecodeString(os.Args[1])
	util.Chkfatal(err)
	launchMsg := &capnp.Message{Arena: capnp.SingleSegment(buf)}
	launchCmd, err := grainagent.ReadRootLaunchCommand(launchMsg)
	util.Chkfatal(err)

	switch launchCmd.Which() {
	case grainagent.LaunchCommand_Which_continueGrain:
		spkCmd, err := manifest.ContinueCommand()
		util.Chkfatal(err)
		cmd, err := parseCmd(spkCmd)
		util.Chkfatal(err)

		lg.WithFields(log.Fields{
			"appTitle": appTitleText,
			"command":  cmd.Args,
		}).Info("Starting up app")

		osCmd := cmd.ToOsCmd()

		// TODO: make direct these in a more structured way?
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr

		apiSocket := os.NewFile(3, "supervisor socket")
		osCmd.ExtraFiles = []*os.File{apiSocket}

		util.Chkfatal(osCmd.Start())
		defer os.Exit(1)
		util.Chkfatal(osCmd.Wait())
		lg.Info("App exited; shutting down grain.")
	default:
		lg.WithFields(log.Fields{
			"launch command": launchCmd.Which(),
		}).Error("BUG: Unrecognized launch command")
	}
}

func startCapnpApi(lg log.Interface) error {
	l, err := net.Listen("unix", "/tmp/sandstorm-api")
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				continue
			}
			// TODO: do something with the connection.
			lg.Info("Got a connection to the api socket.")
			conn.Close()
		}
	}()
	return nil
}
