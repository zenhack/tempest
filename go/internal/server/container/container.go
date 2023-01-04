package container

import (
	"context"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"golang.org/x/sys/unix"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/tempest/go/internal/config"
	"zenhack.net/go/tempest/go/internal/database"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

type Container struct {
	Bootstrap capnp.Client
	cancel    context.CancelFunc // cancel causes the container to shut down.
	exited    <-chan struct{}    // closed when the container has exited.
}

func (c Container) Release() {
	c.Bootstrap.Release()
	c.cancel()
	<-c.exited
}

func Start(ctx context.Context, db database.DB, grainId string, api grain.SandstormApi) (Container, error) {
	return exn.Try(func(throw func(error)) Container {
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()
		pkgId, err := tx.GetGrainPackageId(grainId)
		throw(err)
		throw(tx.Commit())
		ret, err := startContainer(ctx, capnp.Client(api), pkgId, grainId)
		throw(err)
		return ret
	})
}

func startContainer(
	ctx context.Context,
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
		util.Chkfatal(cmd.Process.Kill())
		util.Must(cmd.Process.Wait())
		<-conn.Done()
		close(exited)
	}()
	return Container{
		Bootstrap: grainBootstrap,
		cancel:    cancel,
		exited:    exited,
	}, nil
}
