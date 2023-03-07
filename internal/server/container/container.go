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
	"github.com/apex/log"
	"golang.org/x/sys/unix"

	"zenhack.net/go/tempest/capnp/grain"
	"zenhack.net/go/tempest/internal/config"
	"zenhack.net/go/tempest/internal/server/database"
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
	Log log.Interface
	DB  database.DB

	// GrainID is the ID of the grain to start
	GrainID string

	// Api will be provided to the grain as our bootstrap interface.
	Api grain.SandstormApi

	// Args will be passed to the grain agent as extra arguments.
	Args []string
}

// Start starts the container. It will shut down when ctx is canceled or
// Kill() is called.
func (cmd Command) Start(ctx context.Context) (Container, error) {
	cmd.Log.WithFields(log.Fields{
		"grainId": cmd.GrainID,
	}).Info("Starting grain")
	return exn.Try(func(throw func(error)) Container {
		tx, err := cmd.DB.Begin()
		throw(err)
		defer tx.Rollback()
		pkgId, err := tx.GetGrainPackageId(cmd.GrainID)
		throw(err)
		throw(tx.Commit())
		ret, err := pkgCommand{
			Command: cmd,
			PkgID:   pkgId,
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
		cmd.GrainID,
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
		cmd.Log.WithFields(log.Fields{
			"grainId": cmd.GrainID,
			"error":   err,
		}).Error("Starting sandbox launcher failed")
		cmd.Api.Release()
		supervisorSock.Close()
		return Container{}, err
	}
	cmd.Log.WithFields(log.Fields{
		"launcher-pid": osCmd.Process.Pid,
		"grainId":      cmd.GrainID,
	}).Debug("Started launcher proccess")

	pidBuf, err := io.ReadAll(pidR)
	launcherPid := osCmd.Process.Pid
	if err != nil {
		cmd.Log.WithFields(log.Fields{
			"error":        err,
			"read":         string(pidBuf),
			"grainId":      cmd.GrainID,
			"launcher-pid": launcherPid,
		}).Error("Failed to read grain pid")
		return Container{}, err
	}

	grainPid, err := strconv.Atoi(string(pidBuf))
	if err != nil {
		cmd.Log.WithFields(log.Fields{
			"error":        err,
			"grainId":      cmd.GrainID,
			"launcher-pid": launcherPid,
			"bad-pid":      strconv.Quote(string(pidBuf)),
		}).Error("bug: sandbox-launcher returned invalid pid")
		supervisorSock.Close()
		util.Chkfatal(osCmd.Process.Kill())
		util.Must(osCmd.Process.Wait())
		return Container{}, err
	}
	grainProc, err := os.FindProcess(grainPid)
	util.Chkfatal(err) // Can't fail on unix
	cmd.Log.WithFields(log.Fields{
		"grainId":      cmd.GrainID,
		"packageId":    cmd.PkgID,
		"launcher-pid": launcherPid,
		"grain-pid":    grainPid,
	}).Debug("Started grain process")
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
			cmd.Log.WithFields(log.Fields{
				"error":        err,
				"grainId":      cmd.GrainID,
				"launcher-pid": launcherPid,
				"grain-pid":    grainPid,
			}).Fatal("Failed to kill grain")
		}
		cmd.Log.WithFields(log.Fields{"pid": grainPid}).Debug("Killed grain")
		if _, err := osCmd.Process.Wait(); err != nil {
			cmd.Log.WithFields(log.Fields{
				"error":        err,
				"grainId":      cmd.GrainID,
				"launcher-pid": launcherPid,
				"grain-pid":    grainPid,
			}).Fatal("Failed to wait() on launcher")
		}
		cmd.Log.WithFields(log.Fields{"pid": launcherPid}).Debug("Wait()ed for launcher")
		<-conn.Done()
		close(exited)
	}()
	return Container{
		Bootstrap: grainBootstrap,
		cancel:    cancel,
		exited:    exited,
	}, nil
}
