package browsermain

import (
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
)

var _ pusherHooks[types.GrainID, external.Grain] = grainPusher{}

type grainPusher struct {
}

func (gp grainPusher) Upsert(id types.GrainID, grain external.Grain) (Msg, error) {
	return exn.Try(func(throw func(error)) Msg {
		title, err := grain.Title()
		throw(err)
		sessionToken, err := grain.SessionToken()
		throw(err)
		subdomain, err := grain.Subdomain()
		throw(err)

		return UpsertGrain{
			ID: id,
			Grain: Grain{
				Title:        title,
				SessionToken: sessionToken,
				Subdomain:    subdomain,
				Handle:       grain.Handle().AddRef(),
			},
		}
	})
}

func (gp grainPusher) Remove(id types.GrainID) Msg {
	return RemoveGrain{ID: id}
}

func (gp grainPusher) Clear() Msg {
	return ClearGrains{}
}
