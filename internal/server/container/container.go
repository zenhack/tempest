// Package container manages spawning containers/sandboxes
package container

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strconv"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"golang.org/x/exp/slog"
	"golang.org/x/sys/unix"

	"zenhack.net/go/tempest/capnp/grain"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/logging"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

// A Container is a reference to a running container/sandboxed grain.
type Container struct {
	Bootstrap capnp.Client       // Bootstrap interface for the Container.
	cancel    context.CancelFunc // cancel causes the container to shut down.
	exited    <-chan struct{}    // closed when the container has exited.
}

// Kill forcably shutds down the container. (Note: we do not provide a way
// to ask nicely via SIGTERM or such); apps are expected to be crash-only
// software.
//
// Does not wait for shutdown to complete; see Wait().
func (c Container) Kill() {
	c.Bootstrap.Release()
	c.cancel()
}

// Wait blocks until the container has shut down and then returns.
func (c Container) Wait() {
	<-c.exited
}

// A Command specifies a task to start in a container.
type Command struct {
	Log *slog.Logger
	DB  database.DB

	// GrainID is the ID of the grain to start
	GrainID types.GrainID

	// Api will be provided to the grain as our bootstrap interface.
	Api grain.SandstormApi

	// Args will be passed to the grain agent as extra arguments.
	Args []string
}

// Start starts the container. It will shut down when ctx is canceled or
// Kill() is called.
func (cmd Command) Start(ctx context.Context) (Container, error) {
	cmd.Log.Info("Starting grain",
		"grainID", cmd.GrainID,
	)
	return exn.Try(func(throw func(error)) Container {
		tx, err := cmd.DB.Begin()
		throw(err)
		defer tx.Rollback()
		pkgID, err := tx.GetGrainPackageID(cmd.GrainID)
		throw(err)
		throw(tx.Commit())
		ret, err := pkgCommand{
			Command: cmd,
			PkgID:   pkgID,
		}.Start(ctx)
		throw(err)
		return ret
	})
}

// pkgCommand is like Command, but it also includes the package ID, so looking that
// up in the database is unnecessary.
type pkgCommand struct {
	Command
	PkgID string
}

// Start is like Command.Start
func (cmd pkgCommand) Start(ctx context.Context) (Container, error) {
	// See the comments at the top of sandbox-launcher.c for the details
	// of how the sandbox launcher is supposed to be used.
	ctx, cancel := context.WithCancel(ctx)
	// RPC socket:
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		cmd.Api.Release()
		return Container{}, err
	}

	grainSock := os.NewFile(uintptr(fds[0]), "grain api socket")
	defer grainSock.Close()
	supervisorSock := os.NewFile(uintptr(fds[1]), "supervisor api socket")

	// Pipe to communicate the grain's PID:
	pidR, pidW, err := os.Pipe()
	if err != nil {
		supervisorSock.Close()
		return Container{}, err
	}
	defer pidR.Close()

	args := append([]string{
		cmd.PkgID,
		string(cmd.GrainID),
	}, cmd.Args...)
	osCmd := exec.Command(
		config.Libexecdir+"/tempest/tempest-sandbox-launcher",
		args...,
	)

	// TODO(soon) capture/log stdout/stderr
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr

	osCmd.ExtraFiles = []*os.File{grainSock, pidW}
	err = osCmd.Start()
	pidW.Close() // Close this now, so when the child closes it we hit EOF.
	if err != nil {
		cmd.Log.Error("Starting sandbox launcher failed", err,
			"grainID", cmd.GrainID,
		)
		cmd.Api.Release()
		supervisorSock.Close()
		return Container{}, err
	}
	cmd.Log.Debug("Started launcher proccess",
		"launcher-pid", osCmd.Process.Pid,
		"grainID", cmd.GrainID,
	)

	pidBuf, err := io.ReadAll(pidR)
	launcherPid := osCmd.Process.Pid
	if err != nil {
		cmd.Log.Error("Failed to read grain pid", err,
			"read", string(pidBuf),
			"grainID", cmd.GrainID,
			"launcher-pid", launcherPid,
		)
		return Container{}, err
	}

	grainPid, err := strconv.Atoi(string(pidBuf))
	if err != nil {
		cmd.Log.Error("bug: sandbox-launcher returned invalid pid", err,
			"grainID", cmd.GrainID,
			"launcher-pid", launcherPid,
			"bad-pid", strconv.Quote(string(pidBuf)),
		)
		supervisorSock.Close()
		util.Chkfatal(osCmd.Process.Kill())
		util.Must(osCmd.Process.Wait())
		return Container{}, err
	}
	grainProc, err := os.FindProcess(grainPid)
	util.Chkfatal(err) // Can't fail on unix
	cmd.Log.Debug("Started grain process",
		"grainID", cmd.GrainID,
		"packageId", cmd.PkgID,
		"launcher-pid", launcherPid,
		"grain-pid", grainPid,
	)
	trans := transport.NewStream(supervisorSock)
	options := &rpc.Options{
		BootstrapClient: capnp.Client(cmd.Api),
	}
	conn := rpc.NewConn(trans, options)
	grainBootstrap := conn.Bootstrap(ctx)
	exited := make(chan struct{})
	go func() {
		<-ctx.Done()
		// I(isd) don't see a sensible behavior if we fail to shut down the
		// container, so panic I guess.
		if err := grainProc.Kill(); err != nil {
			logging.Panic(cmd.Log, "Failed to kill grain",
				"error", err,
				"grainID", cmd.GrainID,
				"launcher-pid", launcherPid,
				"grain-pid", grainPid,
			)
		}
		cmd.Log.Debug("Killed grain",
			"pid", grainPid,
		)
		if _, err := osCmd.Process.Wait(); err != nil {
			logging.Panic(cmd.Log, "Failed to wait() on launcher",
				"error", err,
				"grainID", cmd.GrainID,
				"launcher-pid", launcherPid,
				"grain-pid", grainPid,
			)
		}
		cmd.Log.Debug("Wait()ed for launcher",
			"pid", launcherPid,
		)
		<-conn.Done()
		close(exited)
	}()
	return Container{
		Bootstrap: grainBootstrap,
		cancel:    cancel,
		exited:    exited,
	}, nil
}
