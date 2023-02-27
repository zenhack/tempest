package browsermain

import (
	"context"
	"math"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/util/exn"
)

type pkgPusher struct {
	uiMsgs chan<- Msg
}

func (pp pkgPusher) Upsert(ctx context.Context, p collection.Pusher_upsert) error {
	return exn.Try0(func(throw func(error)) {
		args := p.Args()
		key, err := args.Key()
		throw(err)
		val, err := args.Value()
		throw(err)
		srcPkg := external.Package{}.DecodeFromPtr(val)

		_, seg := capnp.NewSingleSegmentMessage(nil)
		dstPkg, err := external.NewPackage(seg)
		throw(err)
		throw(capnp.Struct(dstPkg).CopyFrom(capnp.Struct(srcPkg)))
		dstPkg.Message().ResetReadLimit(math.MaxUint64)

		pp.uiMsgs <- UpsertPackage{
			Id:  ID[external.Package](key.Text()),
			Pkg: dstPkg,
		}
	})
}

func (pp pkgPusher) Remove(ctx context.Context, p collection.Pusher_remove) error {
	return exn.Try0(func(throw func(error)) {
		key, err := p.Args().Key()
		throw(err)
		pp.uiMsgs <- RemovePackage{
			Id: ID[external.Package](key.Text()),
		}
	})
}

func (pp pkgPusher) Clear(context.Context, collection.Pusher_clear) error {
	pp.uiMsgs <- ClearPackages{}
	return nil
}

func (pkgPusher) Ready(context.Context, collection.Pusher_ready) error {
	return nil
}
