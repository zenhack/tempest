package browsermain

import (
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
)

var _ pusherHooks[types.GrainID, external.UiView] = grainPusher{}

type grainPusher struct {
}

func (gp grainPusher) Upsert(id types.GrainID, view external.UiView) (Msg, error) {
	return exn.Try(func(throw func(error)) Msg {
		title, err := view.Title()
		throw(err)
		sessionToken, err := view.SessionToken()
		throw(err)
		subdomain, err := view.Subdomain()
		throw(err)

		return UpsertGrain{
			ID: id,
			Grain: Grain{
				Title:        title,
				SessionToken: sessionToken,
				Subdomain:    subdomain,
				Controller:   view.Controller().AddRef(),
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
