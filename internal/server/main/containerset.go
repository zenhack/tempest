package servermain

import (
	"context"

	"github.com/apex/log"
	"zenhack.net/go/tempest/capnp/grain"
	"zenhack.net/go/tempest/internal/database"
	"zenhack.net/go/tempest/internal/server/container"
)

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
	c, err := container.Start(ctx, lg, db, grainId, api)
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
