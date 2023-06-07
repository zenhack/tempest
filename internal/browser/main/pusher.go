package browsermain

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/util/exn"
)

type pusherHooks[K ~string, V any] interface {
	Upsert(K, V) (Msg, error)
	Remove(K) Msg
	Clear() Msg
}

type pusher[K ~string, V capnp.TypeParam[V]] struct {
	hooks   pusherHooks[K, V]
	sendMsg func(Msg)
}

func (p pusher[K, V]) Upsert(ctx context.Context, call collection.Pusher_upsert) error {
	return exn.Try0(func(throw exn.Thrower) {
		args := call.Args()
		keyPtr, err := args.Key()
		throw(err)
		valPtr, err := args.Value()
		throw(err)
		var val V
		val = val.DecodeFromPtr(valPtr)

		msg, err := p.hooks.Upsert(K(keyPtr.Text()), val)
		throw(err)
		p.sendMsg(msg)
	})
}

func (p pusher[K, V]) Remove(ctx context.Context, call collection.Pusher_remove) error {
	return exn.Try0(func(throw exn.Thrower) {
		key, err := call.Args().Key()
		throw(err)
		p.sendMsg(p.hooks.Remove(K(key.Text())))
	})
}

func (p pusher[K, V]) Clear(context.Context, collection.Pusher_clear) error {
	p.sendMsg(p.hooks.Clear())
	return nil
}

func (pusher[K, V]) Ready(context.Context, collection.Pusher_ready) error {
	return nil
}
