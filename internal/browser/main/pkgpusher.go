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
	return exn.Try(func(throw exn.Thrower) Msg {
		// Copy over to a new message to avoid early release:
		dstPkg, err := cloneStruct(pkg)
		throw(err)

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

// cloneStruct clones the struct src into a new message, and sets the messages read
// limit to math.MaxUint64.
//
// TODO: maybe put this in go-capnp somewhere?
func cloneStruct[T ~capnp.StructKind](src T) (T, error) {
	return exn.Try(func(throw exn.Thrower) T {
		_, seg := capnp.NewSingleSegmentMessage(nil)
		dst, err := capnp.NewRootStruct(seg, capnp.Struct(src).Size())
		throw(err)
		throw(capnp.Struct(dst).CopyFrom(capnp.Struct(src)))
		dst.Message().ResetReadLimit(math.MaxUint64)
		return T(dst)
	})
}
