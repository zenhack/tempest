package browsermain

import (
	"math"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
)

var _ pusherHooks[types.ID[external.Package], external.Package] = pkgPusher{}

type pkgPusher struct {
}

func (pp pkgPusher) Upsert(id types.ID[external.Package], pkg external.Package) (Msg, error) {
	return exn.Try(func(throw func(error)) Msg {
		// Copy over to a new message to avoid early release:
		_, seg := capnp.NewSingleSegmentMessage(nil)
		dstPkg, err := external.NewPackage(seg)
		throw(err)
		throw(capnp.Struct(dstPkg).CopyFrom(capnp.Struct(pkg)))
		dstPkg.Message().ResetReadLimit(math.MaxUint64)

		return UpsertPackage{
			ID:  id,
			Pkg: dstPkg,
		}
	})
}

func (pp pkgPusher) Remove(id types.ID[external.Package]) Msg {
	return RemovePackage{ID: id}
}

func (pp pkgPusher) Clear() Msg {
	return ClearPackages{}
}
