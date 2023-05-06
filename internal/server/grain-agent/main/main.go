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
	"errors"
	"fmt"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"golang.org/x/exp/slog"
	spk "zenhack.net/go/tempest/capnp/package"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
	"zenhack.net/go/tempest/internal/server/logging"
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
	lg := logging.NewLogger()

	data, err := os.ReadFile("/sandstorm-manifest")
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
		panic("Too few arugments")
	}
	buf, err := base64.StdEncoding.DecodeString(os.Args[1])
	util.Chkfatal(err)
	launchMsg := &capnp.Message{Arena: capnp.SingleSegment(buf)}
	launchCmd, err := grainagent.ReadRootLaunchCommand(launchMsg)
	util.Chkfatal(err)

	switch launchCmd.Which() {
	case grainagent.LaunchCommand_Which_continueGrain:
		cmd, err := manifest.ContinueCommand()
		util.Chkfatal(err)
		spawnSpkCmd(lg, appTitleText, cmd)
	case grainagent.LaunchCommand_Which_initGrain:
		index := launchCmd.InitGrain()
		actions, err := manifest.Actions()
		util.Chkfatal(err)
		cmd, err := actions.At(int(index)).Command()
		util.Chkfatal(err)
		spawnSpkCmd(lg, appTitleText, cmd)
	default:
		err := errors.New("unrecognized launch command")
		lg.Error("BUG",
			"error", err,
			"launch command", launchCmd.Which(),
		)
		panic(err)
	}
}

func spawnSpkCmd(lg *slog.Logger, appTitle string, spkCmd spk.Manifest_Command) {
	cmd, err := parseCmd(spkCmd)
	util.Chkfatal(err)

	lg.Info("Starting up app",
		"appTitle", appTitle,
		"command", cmd.Args,
	)

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
}
