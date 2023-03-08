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
	ID    types.GrainID
	Grain Grain
}

type RemoveGrain struct {
	ID types.GrainID
}

type ClearGrains struct{}

type UpsertPackage struct {
	ID  types.ID[external.Package]
	Pkg external.Package
}

type RemovePackage struct {
	ID types.ID[external.Package]
}

type ClearPackages struct{}

type ChangeFocus struct {
	NewFocus Focus
}

type FocusGrain struct {
	ID types.GrainID
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
	m.Grains[msg.ID].Handle.Release()
	m.Grains[msg.ID] = msg.Grain
	return m, nil
}

func (msg RemoveGrain) Update(m Model) (Model, Cmd) {
	m.Grains[msg.ID].Handle.Release()
	delete(m.Grains, msg.ID)
	return m, nil
}

func (ClearGrains) Update(m Model) (Model, Cmd) {
	m.Grains = make(map[types.GrainID]Grain)
	return m, nil
}

func (msg UpsertPackage) Update(m Model) (Model, Cmd) {
	m.Packages[msg.ID].Controller().Release()
	m.Packages[msg.ID] = msg.Pkg
	return m, nil
}

func (msg RemovePackage) Update(m Model) (Model, Cmd) {
	// TODO(perf): release the whole message?
	m.Packages[msg.ID].Controller().Release()
	delete(m.Packages, msg.ID)
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
	m.FocusedGrain = msg.ID
	_, ok := m.OpenGrains[msg.ID]
	if !ok {
		index := m.GrainDomOrder.Add(msg.ID)
		m.OpenGrains[msg.ID] = OpenGrain{
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
				return exn.Try0(func(throw exn.Thrower) {
					manifest, err := pkg.Manifest()
					throw(err)
					appTitle, err := manifest.AppTitle()
					throw(err)
					appTitleText, err := appTitle.DefaultText()
					throw(err)

					// TODO: provide a way to choose this:
					index := 0

					actions, err := manifest.Actions()
					throw(err)
					nounPhrase, err := actions.At(index).NounPhrase()
					throw(err)
					nounPhraseText, err := nounPhrase.DefaultText()
					throw(err)

					p.SetTitle("Untitled " + appTitleText + " " + nounPhraseText)
					p.SetActionIndex(uint32(index))
				})
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
				ID: types.GrainID(id),
				Grain: Grain{
					Title:        title,
					SessionToken: sessionToken,
					Handle:       grain.Handle().AddRef(),
				},
			})
			sendMsg(FocusGrain{ID: types.GrainID(id)})
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
		pusher := collection.Pusher_ServerToClient(pusher[types.ID[external.Package], external.Package]{
			sendMsg: sendMsg,
			hooks:   pkgPusher{},
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
