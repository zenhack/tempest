package container

import (
	"context"
	"os"
	"os/exec"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"capnproto.org/go/capnp/v3/rpc/transport"
	"golang.org/x/sys/unix"

	"zenhack.net/go/sandstorm-next/capnp/container"
	"zenhack.net/go/sandstorm-next/go/internal/config"
	"zenhack.net/go/sandstorm-next/go/internal/util"
	"zenhack.net/go/sandstorm/exp/util/handle"
)

type Spawner struct {
}

func (Spawner) Spawn(_ context.Context, p container.Spawner_spawn) error {
	args := p.Args()
	packageId, err := args.PackageId()
	if err != nil {
		return err
	}
	grainId, err := args.GrainId()
	if err != nil {
		return err
	}

	supervisorBootstrap := args.Bootstrap()

	results, err := p.AllocResults()
	if err != nil {
		return err
	}

	ctx, cancel := handle.WithCancel(context.Background())
	spawnResult, err := startContainer(ctx, supervisorBootstrap.AddRef(), packageId, grainId)
	if err != nil {
		cancel.Release()
		return err
	}

	results.SetBootstrap(spawnResult.bootstrap)
	results.SetHandle(cancel)
	return nil
}

type spawnResult struct {
	cmd       *exec.Cmd
	bootstrap capnp.Client
	sock      *os.File
	conn      *rpc.Conn
}

func startContainer(
	ctx context.Context,
	supervisorBootstrap capnp.Client,
	packageId, grainId string,
) (spawnResult, error) {
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		supervisorBootstrap.Release()
		return spawnResult{}, err
	}
	grainSock := os.NewFile(uintptr(fds[0]), "grain api socket")
	supervisorSock := os.NewFile(uintptr(fds[1]), "supervisor api socket")
	cmd := exec.Command(
		config.Libexecdir+"/sandstorm-sandbox-launcher",
		packageId,
		grainId,
	)
	// TODO(soon) redirect stdio.
	cmd.ExtraFiles = []*os.File{grainSock}
	err = cmd.Start()
	if err != nil {
		supervisorBootstrap.Release()
		grainSock.Close()
		supervisorSock.Close()
		return spawnResult{}, err
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
	go func() {
		<-ctx.Done()
		// I(isd) don't see a sensible behavior if we fail to shut down the
		// container, so panic I guess.
		util.Chkfatal(cmd.Process.Kill())
		util.Chkfatal(cmd.Process.Wait())
		<-conn.Done()
	}()
	return spawnResult{
		cmd:       cmd,
		bootstrap: grainBootstrap,
		sock:      supervisorSock,
		conn:      conn,
	}, nil
}
