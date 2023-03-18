package servermain

import (
	"context"
	"encoding/base64"

	"capnproto.org/go/capnp/v3"
	"golang.org/x/exp/slog"
	"zenhack.net/go/tempest/capnp/grain"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/tempest/internal/server/container"
	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/util"
)

// The first argument to pass to the grain agent to start a pre-existing grain.
// We compute this once on startup.
var continueArg string

func init() {
	// Compute continueArg:
	_, seg := capnp.NewSingleSegmentMessage(nil)
	launchCmd, err := grainagent.NewRootLaunchCommand(seg)
	util.Chkfatal(err)
	launchCmd.SetContinueGrain()
	continueArg = base64.StdEncoding.EncodeToString(seg.Data())
}

type ContainerSet struct {
	// map of grain id to already-running container. TODO:
	//
	// - We really want some kind of weakmap semantics here; this doesn't
	//   give us a clear way to ever shut down containers.
	// - We need to think about detecting containers shutting down on
	//   their own.
	containersByGrainID map[types.GrainID]container.Container
}

func (cset *ContainerSet) Get(ctx context.Context, lg *slog.Logger, db database.DB, grainID types.GrainID) (container.Container, error) {
	c, ok := cset.containersByGrainID[grainID]
	if ok {
		return c, nil
	}
	api := grain.SandstormApi_ServerToClient(sandstormApiImpl{})
	c, err := container.Command{
		Log:     lg,
		DB:      db,
		GrainID: grainID,
		Api:     api,
		Args:    []string{continueArg},
	}.Start(ctx)
	if err == nil {
		cset.containersByGrainID[grainID] = c
	}
	return c, err
}

func (cset *ContainerSet) Release() {
	for _, c := range cset.containersByGrainID {
		c.Kill()
	}
	for _, c := range cset.containersByGrainID {
		c.Wait()
	}
}
