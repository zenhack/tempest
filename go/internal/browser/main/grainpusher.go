package browsermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/util/exn"
)

type grainPusher struct {
	uiMsgs chan<- Msg
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

		k := ID[Grain](key.Text())
		v := Grain{
			Title:        title,
			SessionToken: sessionToken,
			Handle:       grain.Handle().AddRef(),
		}
		gp.uiMsgs <- func(m Model) Model {
			m.Grains[k].Handle.Release()
			m.Grains[k] = v
			return m
		}
	})
}

func (gp grainPusher) Remove(ctx context.Context, p collection.Pusher_remove) error {
	return exn.Try0(func(throw func(error)) {
		key, err := p.Args().Key()
		throw(err)
		k := ID[Grain](key.Text())
		gp.uiMsgs <- func(m Model) Model {
			m.Grains[k].Handle.Release()
			delete(m.Grains, k)
			return m
		}
	})
}

func (gp grainPusher) Clear(context.Context, collection.Pusher_clear) error {
	gp.uiMsgs <- func(m Model) Model {
		m.Grains = make(map[ID[Grain]]Grain)
		return m
	}
	return nil
}

func (gp grainPusher) Ready(context.Context, collection.Pusher_ready) error {
	return nil
}
