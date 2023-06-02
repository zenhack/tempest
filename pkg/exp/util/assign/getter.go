package assign

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/tempest/capnp/util"
	"zenhack.net/go/tempest/pkg/exp/util/handle"
)

// FixedGetter returns a getter that always returns the argument. Its implementation
// of subscribe calls .set() on the passed setter once and then drops the setter.
// When the Getter is dropped, it will call ptr.Message().Release(), thereby dropping
// any capabilities within.
func FixedGetter(ptr capnp.Ptr) util.Getter {
	return util.Getter_ServerToClient(fixedGetter{ptr: ptr})
}

type fixedGetter struct {
	ptr capnp.Ptr
}

func (g fixedGetter) Get(ctx context.Context, p util.Getter_get) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	return results.SetValue(g.ptr)
}

func (g fixedGetter) Subscribe(ctx context.Context, p util.Getter_subscribe) error {
	p.Go()
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetHandle(handle.CallbackHandle(func() {}))
	setter := p.Args().Setter()
	_, rel := setter.Set(ctx, func(p util.Setter_set_Params) error {
		return p.SetValue(g.ptr)
	})
	rel()
	return nil
}

func (g fixedGetter) Shutdown() {
	g.ptr.Message().Release()
}
