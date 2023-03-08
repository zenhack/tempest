package browsermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
)

type grainPusher struct {
	sendMsg func(Msg)
}

func (gp grainPusher) Upsert(ctx context.Context, p collection.Pusher_upsert) error {
	return exn.Try0(func(throw func(error)) {
		args := p.Args()
		key, err := args.Key()
		throw(err)
		val, err := args.Value()
		throw(err)
		grain := external.Grain{}.DecodeFromPtr(val)

		title, err := grain.Title()
		throw(err)
		sessionToken, err := grain.SessionToken()
		throw(err)

		gp.sendMsg(UpsertGrain{
			Id: types.ID[Grain](key.Text()),
			Grain: Grain{
				Title:        title,
				SessionToken: sessionToken,
				Handle:       grain.Handle().AddRef(),
			},
		})
	})
}

func (gp grainPusher) Remove(ctx context.Context, p collection.Pusher_remove) error {
	return exn.Try0(func(throw func(error)) {
		key, err := p.Args().Key()
		throw(err)
		gp.sendMsg(RemoveGrain{
			Id: types.ID[Grain](key.Text()),
		})
	})
}

func (gp grainPusher) Clear(context.Context, collection.Pusher_clear) error {
	gp.sendMsg(ClearGrains{})
	return nil
}

func (gp grainPusher) Ready(context.Context, collection.Pusher_ready) error {
	return nil
}
