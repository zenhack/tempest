package browsermain

import (
	"context"

	"zenhack.net/go/tempest/capnp/collection"
	"zenhack.net/go/tempest/capnp/external"
	"zenhack.net/go/tempest/internal/common/types"
	"zenhack.net/go/util/exn"
	"zenhack.net/go/util/maybe"
	"zenhack.net/go/util/orerr"
	"zenhack.net/go/vdom/tea"
)

type Cmd = func(context.Context, func(Msg))

type Msg = tea.Message[Model]

type NewError struct {
	Err error
}

type UpsertGrain struct {
	Id    types.ID[Grain]
	Grain Grain
}

type RemoveGrain struct {
	Id types.ID[Grain]
}

type ClearGrains struct{}

type UpsertPackage struct {
	Id  types.ID[external.Package]
	Pkg external.Package
}

type RemovePackage struct {
	Id types.ID[external.Package]
}

type ClearPackages struct{}

type ChangeFocus struct {
	NewFocus Focus
}

type FocusGrain struct {
	Id types.ID[Grain]
}

type SpawnGrain struct {
	PkgID types.ID[external.Package]
}

type LoginSessionResult struct {
	Result orerr.T[external.LoginSession]
}

func (msg NewError) Update(m Model) (Model, Cmd) {
	m.Error = msg.Err
	return m, nil
}

func (msg UpsertGrain) Update(m Model) (Model, Cmd) {
	m.Grains[msg.Id].Handle.Release()
	m.Grains[msg.Id] = msg.Grain
	return m, nil
}

func (msg RemoveGrain) Update(m Model) (Model, Cmd) {
	m.Grains[msg.Id].Handle.Release()
	delete(m.Grains, msg.Id)
	return m, nil
}

func (ClearGrains) Update(m Model) (Model, Cmd) {
	m.Grains = make(map[types.ID[Grain]]Grain)
	return m, nil
}

func (msg UpsertPackage) Update(m Model) (Model, Cmd) {
	m.Packages[msg.Id].Controller().Release()
	m.Packages[msg.Id] = msg.Pkg
	return m, nil
}

func (msg RemovePackage) Update(m Model) (Model, Cmd) {
	// TODO(perf): release the whole message?
	m.Packages[msg.Id].Controller().Release()
	delete(m.Packages, msg.Id)
	return m, nil
}

func (ClearPackages) Update(m Model) (Model, Cmd) {
	m.Packages = make(map[types.ID[external.Package]]external.Package)
	return m, nil
}

func (msg ChangeFocus) Update(m Model) (Model, Cmd) {
	m.CurrentFocus = msg.NewFocus
	return m, nil
}

func (msg FocusGrain) Update(m Model) (Model, Cmd) {
	m.CurrentFocus = FocusOpenGrain
	m.FocusedGrain = msg.Id
	_, ok := m.OpenGrains[msg.Id]
	if !ok {
		index := m.GrainDomOrder.Add(msg.Id)
		m.OpenGrains[msg.Id] = OpenGrain{
			DomainNonce: newDomainNonce(),
			DomIndex:    index,
		}
	}
	return m, nil
}

func (msg SpawnGrain) Update(m Model) (Model, Cmd) {
	pkg := m.Packages[msg.PkgID]

	ctrl := pkg.Controller().AddRef()

	return m, func(ctx context.Context, sendMsg func(Msg)) {
		err := exn.Try0(func(throw func(error)) {
			defer ctrl.Release()
			fut, rel := ctrl.Create(ctx, func(p external.Package_Controller_create_Params) error {
				// TODO: pick something better, looking at the action/manifest:
				p.SetTitle("Untitled Grain")
				// TODO: provide a way to choose this:
				p.SetActionIndex(0)
				return nil
			})
			defer rel()
			res, err := fut.Struct()
			throw(err)

			id, err := res.Id()
			throw(err)
			grain, err := res.Grain()
			throw(err)

			title, err := grain.Title()
			throw(err)
			sessionToken, err := grain.SessionToken()
			throw(err)

			sendMsg(UpsertGrain{
				Id: types.ID[Grain](id),
				Grain: Grain{
					Title:        title,
					SessionToken: sessionToken,
					Handle:       grain.Handle().AddRef(),
				},
			})
			sendMsg(FocusGrain{Id: types.ID[Grain](id)})
		})
		if err != nil {
			sendMsg(NewError{Err: err})
		}
	}
}

func (msg LoginSessionResult) Update(m Model) (Model, Cmd) {
	m.LoginSession = maybe.New(msg.Result)
	sess, err := msg.Result.Get()
	if err != nil {
		return m, nil
	}
	return m, func(ctx context.Context, sendMsg func(Msg)) {
		pusher := collection.Pusher_ServerToClient(pkgPusher{
			sendMsg: sendMsg,
		})
		ret, rel := sess.ListPackages(context.Background(), func(p external.LoginSession_listPackages_Params) error {
			p.SetInto(pusher)
			return nil
		})
		defer rel()
		_, err := ret.Struct()
		if err != nil {
			println("listPackages(): " + err.Error())
		}
	}
}
