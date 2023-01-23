package browsermain

type Msg interface {
	Apply(Model) Model
}

type UpsertGrain struct {
	Id    ID[Grain]
	Grain Grain
}

type RemoveGrain struct {
	Id ID[Grain]
}

type ClearGrains struct{}

type FocusGrain struct {
	Id ID[Grain]
}

func (msg UpsertGrain) Apply(m Model) Model {
	m.Grains[msg.Id].Handle.Release()
	m.Grains[msg.Id] = msg.Grain
	return m
}

func (msg RemoveGrain) Apply(m Model) Model {
	m.Grains[msg.Id].Handle.Release()
	delete(m.Grains, msg.Id)
	return m
}

func (ClearGrains) Apply(m Model) Model {
	m.Grains = make(map[ID[Grain]]Grain)
	return m
}

func (msg FocusGrain) Apply(m Model) Model {
	m.FocusedGrain = msg.Id
	_, ok := m.OpenGrains[msg.Id]
	if !ok {
		m.OpenGrains[msg.Id] = OpenGrain{
			DomainNonce: newDomainNonce(),
		}
	}
	return m
}
