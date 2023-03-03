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

type Container struct {
	Bootstrap capnp.Client
	cancel    context.CancelFunc // cancel causes the container to shut down.
	exited    <-chan struct{}    // closed when the container has exited.
}

func (c Container) Kill() {
	c.Bootstrap.Release()
	c.cancel()
}

func (c Container) Wait() {
	<-c.exited
}

func Start(ctx context.Context, lg log.Interface, db database.DB, grainId string, api grain.SandstormApi) (Container, error) {
	lg.WithFields(log.Fields{
		"grainId": grainId,
	}).Info("Starting grain")
	return exn.Try(func(throw func(error)) Container {
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()
		pkgId, err := tx.GetGrainPackageId(grainId)
		throw(err)
		throw(tx.Commit())
		ret, err := startContainer(ctx, lg, capnp.Client(api), pkgId, grainId)
		throw(err)
		return ret
	})
}

func startContainer(
	ctx context.Context,
	lg log.Interface,
	supervisorBootstrap capnp.Client,
	packageId, grainId string,
) (Container, error) {
	// See the comments at the top of sandbox-launcher.c for the details
	// of how the sandbox launcher is supposed to be used.
	ctx, cancel := context.WithCancel(ctx)
	// RPC socket:
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		supervisorBootstrap.Release()
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

	cmd := exec.Command(
		config.Libexecdir+"/tempest/tempest-sandbox-launcher",
		packageId,
		grainId,
	)
	// TODO(soon) capture/log stdout/stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.ExtraFiles = []*os.File{grainSock, pidW}
	err = cmd.Start()
	pidW.Close() // Close this now, so when the child closes it we hit EOF.
	if err != nil {
		lg.WithFields(log.Fields{
			"grainId": grainId,
			"error":   err,
		}).Error("Starting sandbox launcher failed")
		supervisorBootstrap.Release()
		supervisorSock.Close()
		return Container{}, err
	}
	lg.WithFields(log.Fields{
		"launcher-pid": cmd.Process.Pid,
		"grainId":      grainId,
	}).Debug("Started launcher proccess")

	pidBuf, err := io.ReadAll(pidR)
	if err != nil {
		lg.WithFields(log.Fields{
			"error":        err,
			"read":         string(pidBuf),
			"grainId":      grainId,
			"launcher-pid": cmd.Process.Pid,
		}).Error("Failed to read grain pid")
		return Container{}, err
	}

	grainPid, err := strconv.Atoi(string(pidBuf))
	if err != nil {
		lg.WithFields(log.Fields{
			"error":        err,
			"grainId":      grainId,
			"launcher-pid": cmd.Process.Pid,
			"bad-pid":      strconv.Quote(string(pidBuf)),
		}).Error("bug: sandbox-launcher returned invalid pid")
		supervisorSock.Close()
		util.Chkfatal(cmd.Process.Kill())
		util.Must(cmd.Process.Wait())
		return Container{}, err
	}
	grainProc, err := os.FindProcess(grainPid)
	util.Chkfatal(err) // Can't fail on unix
	lg.WithFields(log.Fields{
		"grainId":      grainId,
		"packageId":    packageId,
		"launcher-pid": cmd.Process.Pid,
		"grain-pid":    grainPid,
	}).Debug("Started grain process")
	trans := transport.NewStream(supervisorSock)
	var options *rpc.Options
	if (supervisorBootstrap != capnp.Client{}) {
		options = &rpc.Options{
			BootstrapClient: supervisorBootstrap,
		}
	}
	conn := rpc.NewConn(trans, options)
	grainBootstrap := conn.Bootstrap(ctx)
	exited := make(chan struct{})
	go func() {
		<-ctx.Done()
		// I(isd) don't see a sensible behavior if we fail to shut down the
		// container, so panic I guess.
		if err := grainProc.Kill(); err != nil {
			lg.WithFields(log.Fields{
				"error":        err,
				"grainId":      grainId,
				"launcher-pid": cmd.Process.Pid,
				"grain-pid":    grainPid,
			}).Fatal("Failed to kill grain")
		}
		lg.WithFields(log.Fields{"pid": grainPid}).Debug("Killed grain")
		if _, err := cmd.Process.Wait(); err != nil {
			lg.WithFields(log.Fields{
				"error":        err,
				"grainId":      grainId,
				"launcher-pid": cmd.Process.Pid,
				"grain-pid":    grainPid,
			}).Fatal("Failed to wait() on launcher")
		}
		lg.WithFields(log.Fields{"pid": cmd.Process.Pid}).Debug("Wait()ed for launcher")
		<-conn.Done()
		close(exited)
	}()
	return Container{
		Bootstrap: grainBootstrap,
		cancel:    cancel,
		exited:    exited,
	}, nil
}
