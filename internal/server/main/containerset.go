package servermain

import (
	"context"
	"encoding/base64"

	"capnproto.org/go/capnp/v3"
	"github.com/apex/log"
	"zenhack.net/go/tempest/capnp/grain"
	grainagent "zenhack.net/go/tempest/internal/capnp/grain-agent"
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
	containersByGrainId map[string]container.Container
}

func (cset *ContainerSet) Get(ctx context.Context, lg log.Interface, db database.DB, grainId string) (container.Container, error) {
	c, ok := cset.containersByGrainId[grainId]
	if ok {
		return c, nil
	}
	api := grain.SandstormApi_ServerToClient(sandstormApiImpl{})
	c, err := container.Command{
		Log:     lg,
		DB:      db,
		GrainID: grainId,
		Api:     api,
		Args:    []string{continueArg},
	}.Start(ctx)
	if err == nil {
		cset.containersByGrainId[grainId] = c
	}
	return c, err
}

func (cset *ContainerSet) Release() {
	for _, c := range cset.containersByGrainId {
		c.Kill()
	}
	for _, c := range cset.containersByGrainId {
		c.Wait()
	}
}
