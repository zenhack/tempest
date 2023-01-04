package container

import (
	"context"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"github.com/apex/log"
	"golang.org/x/sys/unix"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/tempest/go/internal/config"
	"zenhack.net/go/tempest/go/internal/database"
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
	ctx, cancel := context.WithCancel(ctx)
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		supervisorBootstrap.Release()
		return Container{}, err
	}
	grainSock := os.NewFile(uintptr(fds[0]), "grain api socket")
	supervisorSock := os.NewFile(uintptr(fds[1]), "supervisor api socket")
	cmd := exec.Command(
		config.Libexecdir+"/tempest/tempest-sandbox-launcher",
		packageId,
		grainId,
	)
	// TODO(soon) capture/log stdout/stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.ExtraFiles = []*os.File{grainSock}
	err = cmd.Start()
	if err != nil {
		supervisorBootstrap.Release()
		grainSock.Close()
		supervisorSock.Close()
		return Container{}, err
	}
	lg.WithFields(log.Fields{
		"grainId":   grainId,
		"packageId": packageId,
		"pid":       cmd.Process.Pid,
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
		if err := cmd.Process.Kill(); err != nil {
			lg.WithFields(log.Fields{
				"error":   err,
				"grainId": grainId,
				"pid":     cmd.Process.Pid,
			}).Fatal("Failed to kill grain")
		}
		if _, err := cmd.Process.Wait(); err != nil {
			lg.WithFields(log.Fields{
				"error":   err,
				"grainId": grainId,
				"pid":     cmd.Process.Pid,
			}).Fatal("Failed to wait() on grain")
		}
		<-conn.Done()
		close(exited)
	}()
	return Container{
		Bootstrap: grainBootstrap,
		cancel:    cancel,
		exited:    exited,
	}, nil
}
