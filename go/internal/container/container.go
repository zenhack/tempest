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
	"zenhack.net/go/sandstorm-next/go/internal/database"
	utilcp "zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/exp/util/handle"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

type Container struct {
	Bootstrap capnp.Client
	Handle    utilcp.Handle
}

func (c *Container) Release() {
	c.Bootstrap.Release()
	c.Handle.Release()
}

func StartDummy(ctx context.Context, db database.DB) (*Container, error) {
	return exn.Try(func(throw func(error)) *Container {
		grainId := os.Getenv("DUMMY_GRAIN_ID")
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()
		pkgId, err := tx.GetGrainPackageId(grainId)
		throw(err)
		throw(tx.Commit())

		spawner := container.Spawner_ServerToClient(Spawner{})
		defer spawner.Release()
		fut, rel := spawner.Spawn(ctx, func(p container.Spawner_spawn_Params) error {
			// TODO: bootstrap
			util.Chkfatal(p.SetPackageId(pkgId))
			util.Chkfatal(p.SetGrainId(grainId))
			return nil
		})
		defer rel()
		results, err := fut.Struct()
		throw(err)
		return &Container{
			Bootstrap: results.Bootstrap().AddRef(),
			Handle:    results.Handle().AddRef(),
		}
	})
}

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
	grainBootstrap, err := startContainer(ctx, supervisorBootstrap.AddRef(), packageId, grainId)
	if err != nil {
		cancel.Release()
		return err
	}

	results.SetBootstrap(grainBootstrap)
	results.SetHandle(cancel)
	return nil
}

func startContainer(
	ctx context.Context,
	supervisorBootstrap capnp.Client,
	packageId, grainId string,
) (capnp.Client, error) {
	fds, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil {
		supervisorBootstrap.Release()
		return capnp.Client{}, err
	}
	grainSock := os.NewFile(uintptr(fds[0]), "grain api socket")
	supervisorSock := os.NewFile(uintptr(fds[1]), "supervisor api socket")
	cmd := exec.Command(
		config.Libexecdir+"/sandstorm/sandstorm-sandbox-launcher",
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
		return capnp.Client{}, err
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
		_, err = cmd.Process.Wait()
		util.Chkfatal(err)
		<-conn.Done()
	}()
	return grainBootstrap, nil
}
