package servermain

import (
	"context"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/tempest/go/internal/database"
	"zenhack.net/go/tempest/go/internal/server/container"
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

func (cset *ContainerSet) Get(ctx context.Context, db database.DB, grainId string) (container.Container, error) {
	c, ok := cset.containersByGrainId[grainId]
	if ok {
		return c, nil
	}
	api := grain.SandstormApi_ServerToClient(sandstormApiImpl{})
	c, err := container.Start(ctx, db, grainId, api)
	if err == nil {
		cset.containersByGrainId[grainId] = c
	}
	return c, err
}
