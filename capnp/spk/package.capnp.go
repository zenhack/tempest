package spk

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
)

// Constants defined in package.capnp.
var (
	MagicNumber = []byte{143, 198, 205, 239, 69, 26, 234, 150}
)

type PackageDefinition struct{ capnp.Struct }

func NewPackageDefinition(s *capnp.Segment) (PackageDefinition, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	if err != nil {
		return PackageDefinition{}, err
	}
	return PackageDefinition{st}, nil
}

func NewRootPackageDefinition(s *capnp.Segment) (PackageDefinition, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	if err != nil {
		return PackageDefinition{}, err
	}
	return PackageDefinition{st}, nil
}

func ReadRootPackageDefinition(msg *capnp.Message) (PackageDefinition, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return PackageDefinition{}, err
	}
	return PackageDefinition{root.Struct()}, nil
}
func (s PackageDefinition) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s PackageDefinition) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s PackageDefinition) SetId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s PackageDefinition) Manifest() (Manifest, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{Struct: p.Struct()}, nil
}

func (s PackageDefinition) HasManifest() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) SetManifest(v Manifest) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewManifest sets the manifest field to a newly
// allocated Manifest struct, preferring placement in s's segment.
func (s PackageDefinition) NewManifest() (Manifest, error) {
	ss, err := NewManifest(s.Struct.Segment())
	if err != nil {
		return Manifest{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PackageDefinition) SourceMap() (SourceMap, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return SourceMap{}, err
	}
	return SourceMap{Struct: p.Struct()}, nil
}

func (s PackageDefinition) HasSourceMap() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) SetSourceMap(v SourceMap) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewSourceMap sets the sourceMap field to a newly
// allocated SourceMap struct, preferring placement in s's segment.
func (s PackageDefinition) NewSourceMap() (SourceMap, error) {
	ss, err := NewSourceMap(s.Struct.Segment())
	if err != nil {
		return SourceMap{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s PackageDefinition) FileList() (string, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s PackageDefinition) HasFileList() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) FileListBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s PackageDefinition) SetFileList(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s PackageDefinition) AlwaysInclude() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
}

func (s PackageDefinition) HasAlwaysInclude() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) SetAlwaysInclude(v capnp.TextList) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewAlwaysInclude sets the alwaysInclude field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s PackageDefinition) NewAlwaysInclude(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s PackageDefinition) BridgeConfig() (BridgeConfig, error) {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		return BridgeConfig{}, err
	}
	return BridgeConfig{Struct: p.Struct()}, nil
}

func (s PackageDefinition) HasBridgeConfig() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) SetBridgeConfig(v BridgeConfig) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewBridgeConfig sets the bridgeConfig field to a newly
// allocated BridgeConfig struct, preferring placement in s's segment.
func (s PackageDefinition) NewBridgeConfig() (BridgeConfig, error) {
	ss, err := NewBridgeConfig(s.Struct.Segment())
	if err != nil {
		return BridgeConfig{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

// PackageDefinition_List is a list of PackageDefinition.
type PackageDefinition_List struct{ capnp.List }

// NewPackageDefinition creates a new list of PackageDefinition.
func NewPackageDefinition_List(s *capnp.Segment, sz int32) (PackageDefinition_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6}, sz)
	if err != nil {
		return PackageDefinition_List{}, err
	}
	return PackageDefinition_List{l}, nil
}

func (s PackageDefinition_List) At(i int) PackageDefinition {
	return PackageDefinition{s.List.Struct(i)}
}
func (s PackageDefinition_List) Set(i int, v PackageDefinition) error {
	return s.List.SetStruct(i, v.Struct)
}

// PackageDefinition_Promise is a wrapper for a PackageDefinition promised by a client call.
type PackageDefinition_Promise struct{ *capnp.Pipeline }

func (p PackageDefinition_Promise) Struct() (PackageDefinition, error) {
	s, err := p.Pipeline.Struct()
	return PackageDefinition{s}, err
}

func (p PackageDefinition_Promise) Manifest() Manifest_Promise {
	return Manifest_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p PackageDefinition_Promise) SourceMap() SourceMap_Promise {
	return SourceMap_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p PackageDefinition_Promise) BridgeConfig() BridgeConfig_Promise {
	return BridgeConfig_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

type Manifest struct{ capnp.Struct }

func NewManifest(s *capnp.Segment) (Manifest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{st}, nil
}

func NewRootManifest(s *capnp.Segment) (Manifest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{st}, nil
}

func ReadRootManifest(msg *capnp.Message) (Manifest, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{root.Struct()}, nil
}
func (s Manifest) AppTitle() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return util.LocalizedText{}, err
	}
	return util.LocalizedText{Struct: p.Struct()}, nil
}

func (s Manifest) HasAppTitle() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Manifest) SetAppTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewAppTitle sets the appTitle field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest) NewAppTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest) AppVersion() uint32 {
	return s.Struct.Uint32(8)
}

func (s Manifest) SetAppVersion(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Manifest) MinUpgradableAppVersion() uint32 {
	return s.Struct.Uint32(12)
}

func (s Manifest) SetMinUpgradableAppVersion(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s Manifest) AppMarketingVersion() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return util.LocalizedText{}, err
	}
	return util.LocalizedText{Struct: p.Struct()}, nil
}

func (s Manifest) HasAppMarketingVersion() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Manifest) SetAppMarketingVersion(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewAppMarketingVersion sets the appMarketingVersion field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest) NewAppMarketingVersion() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest) MinApiVersion() uint32 {
	return s.Struct.Uint32(0)
}

func (s Manifest) SetMinApiVersion(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Manifest) MaxApiVersion() uint32 {
	return s.Struct.Uint32(4)
}

func (s Manifest) SetMaxApiVersion(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Manifest) Actions() (Manifest_Action_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Manifest_Action_List{}, err
	}
	return Manifest_Action_List{List: p.List()}, nil
}

func (s Manifest) HasActions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Manifest) SetActions(v Manifest_Action_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewActions sets the actions field to a newly
// allocated Manifest_Action_List, preferring placement in s's segment.
func (s Manifest) NewActions(n int32) (Manifest_Action_List, error) {
	l, err := NewManifest_Action_List(s.Struct.Segment(), n)
	if err != nil {
		return Manifest_Action_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Manifest) ContinueCommand() (Manifest_Command, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Manifest_Command{}, err
	}
	return Manifest_Command{Struct: p.Struct()}, nil
}

func (s Manifest) HasContinueCommand() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Manifest) SetContinueCommand(v Manifest_Command) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewContinueCommand sets the continueCommand field to a newly
// allocated Manifest_Command struct, preferring placement in s's segment.
func (s Manifest) NewContinueCommand() (Manifest_Command, error) {
	ss, err := NewManifest_Command(s.Struct.Segment())
	if err != nil {
		return Manifest_Command{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Manifest_List is a list of Manifest.
type Manifest_List struct{ capnp.List }

// NewManifest creates a new list of Manifest.
func NewManifest_List(s *capnp.Segment, sz int32) (Manifest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	if err != nil {
		return Manifest_List{}, err
	}
	return Manifest_List{l}, nil
}

func (s Manifest_List) At(i int) Manifest           { return Manifest{s.List.Struct(i)} }
func (s Manifest_List) Set(i int, v Manifest) error { return s.List.SetStruct(i, v.Struct) }

// Manifest_Promise is a wrapper for a Manifest promised by a client call.
type Manifest_Promise struct{ *capnp.Pipeline }

func (p Manifest_Promise) Struct() (Manifest, error) {
	s, err := p.Pipeline.Struct()
	return Manifest{s}, err
}

func (p Manifest_Promise) AppTitle() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Manifest_Promise) AppMarketingVersion() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Manifest_Promise) ContinueCommand() Manifest_Command_Promise {
	return Manifest_Command_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Manifest_Command struct{ capnp.Struct }

func NewManifest_Command(s *capnp.Segment) (Manifest_Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return Manifest_Command{}, err
	}
	return Manifest_Command{st}, nil
}

func NewRootManifest_Command(s *capnp.Segment) (Manifest_Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return Manifest_Command{}, err
	}
	return Manifest_Command{st}, nil
}

func ReadRootManifest_Command(msg *capnp.Message) (Manifest_Command, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Manifest_Command{}, err
	}
	return Manifest_Command{root.Struct()}, nil
}
func (s Manifest_Command) Argv() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
}

func (s Manifest_Command) HasArgv() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Manifest_Command) SetArgv(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewArgv sets the argv field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Manifest_Command) NewArgv(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Manifest_Command) Environ() (util.KeyValue_List, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return util.KeyValue_List{}, err
	}
	return util.KeyValue_List{List: p.List()}, nil
}

func (s Manifest_Command) HasEnviron() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Manifest_Command) SetEnviron(v util.KeyValue_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewEnviron sets the environ field to a newly
// allocated util.KeyValue_List, preferring placement in s's segment.
func (s Manifest_Command) NewEnviron(n int32) (util.KeyValue_List, error) {
	l, err := util.NewKeyValue_List(s.Struct.Segment(), n)
	if err != nil {
		return util.KeyValue_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Manifest_Command) DeprecatedExecutablePath() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Manifest_Command) HasDeprecatedExecutablePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Manifest_Command) DeprecatedExecutablePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Manifest_Command) SetDeprecatedExecutablePath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Manifest_Command_List is a list of Manifest_Command.
type Manifest_Command_List struct{ capnp.List }

// NewManifest_Command creates a new list of Manifest_Command.
func NewManifest_Command_List(s *capnp.Segment, sz int32) (Manifest_Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return Manifest_Command_List{}, err
	}
	return Manifest_Command_List{l}, nil
}

func (s Manifest_Command_List) At(i int) Manifest_Command { return Manifest_Command{s.List.Struct(i)} }
func (s Manifest_Command_List) Set(i int, v Manifest_Command) error {
	return s.List.SetStruct(i, v.Struct)
}

// Manifest_Command_Promise is a wrapper for a Manifest_Command promised by a client call.
type Manifest_Command_Promise struct{ *capnp.Pipeline }

func (p Manifest_Command_Promise) Struct() (Manifest_Command, error) {
	s, err := p.Pipeline.Struct()
	return Manifest_Command{s}, err
}

type Manifest_Action struct{ capnp.Struct }
type Manifest_Action_input Manifest_Action
type Manifest_Action_input_Which uint16

const (
	Manifest_Action_input_Which_none       Manifest_Action_input_Which = 0
	Manifest_Action_input_Which_capability Manifest_Action_input_Which = 1
)

func (w Manifest_Action_input_Which) String() string {
	const s = "nonecapability"
	switch w {
	case Manifest_Action_input_Which_none:
		return s[0:4]
	case Manifest_Action_input_Which_capability:
		return s[4:14]

	}
	return "Manifest_Action_input_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewManifest_Action(s *capnp.Segment) (Manifest_Action, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return Manifest_Action{}, err
	}
	return Manifest_Action{st}, nil
}

func NewRootManifest_Action(s *capnp.Segment) (Manifest_Action, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return Manifest_Action{}, err
	}
	return Manifest_Action{st}, nil
}

func ReadRootManifest_Action(msg *capnp.Message) (Manifest_Action, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Manifest_Action{}, err
	}
	return Manifest_Action{root.Struct()}, nil
}
func (s Manifest_Action) Input() Manifest_Action_input { return Manifest_Action_input(s) }
func (s Manifest_Action_input) Which() Manifest_Action_input_Which {
	return Manifest_Action_input_Which(s.Struct.Uint16(0))
}
func (s Manifest_Action_input) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s Manifest_Action_input) Capability() (grain.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return grain.PowerboxDescriptor_List{}, err
	}
	return grain.PowerboxDescriptor_List{List: p.List()}, nil
}

func (s Manifest_Action_input) HasCapability() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Manifest_Action_input) SetCapability(v grain.PowerboxDescriptor_List) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewCapability sets the capability field to a newly
// allocated grain.PowerboxDescriptor_List, preferring placement in s's segment.
func (s Manifest_Action_input) NewCapability(n int32) (grain.PowerboxDescriptor_List, error) {
	s.Struct.SetUint16(0, 1)
	l, err := grain.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return grain.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Manifest_Action) Command() (Manifest_Command, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Manifest_Command{}, err
	}
	return Manifest_Command{Struct: p.Struct()}, nil
}

func (s Manifest_Action) HasCommand() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Manifest_Action) SetCommand(v Manifest_Command) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewCommand sets the command field to a newly
// allocated Manifest_Command struct, preferring placement in s's segment.
func (s Manifest_Action) NewCommand() (Manifest_Command, error) {
	ss, err := NewManifest_Command(s.Struct.Segment())
	if err != nil {
		return Manifest_Command{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest_Action) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return util.LocalizedText{}, err
	}
	return util.LocalizedText{Struct: p.Struct()}, nil
}

func (s Manifest_Action) HasTitle() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Manifest_Action) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest_Action) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return util.LocalizedText{}, err
	}
	return util.LocalizedText{Struct: p.Struct()}, nil
}

func (s Manifest_Action) HasDescription() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Manifest_Action) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// Manifest_Action_List is a list of Manifest_Action.
type Manifest_Action_List struct{ capnp.List }

// NewManifest_Action creates a new list of Manifest_Action.
func NewManifest_Action_List(s *capnp.Segment, sz int32) (Manifest_Action_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	if err != nil {
		return Manifest_Action_List{}, err
	}
	return Manifest_Action_List{l}, nil
}

func (s Manifest_Action_List) At(i int) Manifest_Action { return Manifest_Action{s.List.Struct(i)} }
func (s Manifest_Action_List) Set(i int, v Manifest_Action) error {
	return s.List.SetStruct(i, v.Struct)
}

// Manifest_Action_Promise is a wrapper for a Manifest_Action promised by a client call.
type Manifest_Action_Promise struct{ *capnp.Pipeline }

func (p Manifest_Action_Promise) Struct() (Manifest_Action, error) {
	s, err := p.Pipeline.Struct()
	return Manifest_Action{s}, err
}

func (p Manifest_Action_Promise) Input() Manifest_Action_input_Promise {
	return Manifest_Action_input_Promise{p.Pipeline}
}

// Manifest_Action_input_Promise is a wrapper for a Manifest_Action_input promised by a client call.
type Manifest_Action_input_Promise struct{ *capnp.Pipeline }

func (p Manifest_Action_input_Promise) Struct() (Manifest_Action_input, error) {
	s, err := p.Pipeline.Struct()
	return Manifest_Action_input{s}, err
}

func (p Manifest_Action_Promise) Command() Manifest_Command_Promise {
	return Manifest_Command_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Manifest_Action_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Manifest_Action_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SourceMap struct{ capnp.Struct }

func NewSourceMap(s *capnp.Segment) (SourceMap, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SourceMap{}, err
	}
	return SourceMap{st}, nil
}

func NewRootSourceMap(s *capnp.Segment) (SourceMap, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SourceMap{}, err
	}
	return SourceMap{st}, nil
}

func ReadRootSourceMap(msg *capnp.Message) (SourceMap, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SourceMap{}, err
	}
	return SourceMap{root.Struct()}, nil
}
func (s SourceMap) SearchPath() (SourceMap_Mapping_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return SourceMap_Mapping_List{}, err
	}
	return SourceMap_Mapping_List{List: p.List()}, nil
}

func (s SourceMap) HasSearchPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SourceMap) SetSearchPath(v SourceMap_Mapping_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSearchPath sets the searchPath field to a newly
// allocated SourceMap_Mapping_List, preferring placement in s's segment.
func (s SourceMap) NewSearchPath(n int32) (SourceMap_Mapping_List, error) {
	l, err := NewSourceMap_Mapping_List(s.Struct.Segment(), n)
	if err != nil {
		return SourceMap_Mapping_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// SourceMap_List is a list of SourceMap.
type SourceMap_List struct{ capnp.List }

// NewSourceMap creates a new list of SourceMap.
func NewSourceMap_List(s *capnp.Segment, sz int32) (SourceMap_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SourceMap_List{}, err
	}
	return SourceMap_List{l}, nil
}

func (s SourceMap_List) At(i int) SourceMap           { return SourceMap{s.List.Struct(i)} }
func (s SourceMap_List) Set(i int, v SourceMap) error { return s.List.SetStruct(i, v.Struct) }

// SourceMap_Promise is a wrapper for a SourceMap promised by a client call.
type SourceMap_Promise struct{ *capnp.Pipeline }

func (p SourceMap_Promise) Struct() (SourceMap, error) {
	s, err := p.Pipeline.Struct()
	return SourceMap{s}, err
}

type SourceMap_Mapping struct{ capnp.Struct }

func NewSourceMap_Mapping(s *capnp.Segment) (SourceMap_Mapping, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SourceMap_Mapping{}, err
	}
	return SourceMap_Mapping{st}, nil
}

func NewRootSourceMap_Mapping(s *capnp.Segment) (SourceMap_Mapping, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SourceMap_Mapping{}, err
	}
	return SourceMap_Mapping{st}, nil
}

func ReadRootSourceMap_Mapping(msg *capnp.Message) (SourceMap_Mapping, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SourceMap_Mapping{}, err
	}
	return SourceMap_Mapping{root.Struct()}, nil
}
func (s SourceMap_Mapping) PackagePath() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s SourceMap_Mapping) HasPackagePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SourceMap_Mapping) PackagePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s SourceMap_Mapping) SetPackagePath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s SourceMap_Mapping) SourcePath() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s SourceMap_Mapping) HasSourcePath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SourceMap_Mapping) SourcePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s SourceMap_Mapping) SetSourcePath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s SourceMap_Mapping) HidePaths() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return capnp.TextList{}, err
	}
	return capnp.TextList{List: p.List()}, nil
}

func (s SourceMap_Mapping) HasHidePaths() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SourceMap_Mapping) SetHidePaths(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewHidePaths sets the hidePaths field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s SourceMap_Mapping) NewHidePaths(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// SourceMap_Mapping_List is a list of SourceMap_Mapping.
type SourceMap_Mapping_List struct{ capnp.List }

// NewSourceMap_Mapping creates a new list of SourceMap_Mapping.
func NewSourceMap_Mapping_List(s *capnp.Segment, sz int32) (SourceMap_Mapping_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return SourceMap_Mapping_List{}, err
	}
	return SourceMap_Mapping_List{l}, nil
}

func (s SourceMap_Mapping_List) At(i int) SourceMap_Mapping {
	return SourceMap_Mapping{s.List.Struct(i)}
}
func (s SourceMap_Mapping_List) Set(i int, v SourceMap_Mapping) error {
	return s.List.SetStruct(i, v.Struct)
}

// SourceMap_Mapping_Promise is a wrapper for a SourceMap_Mapping promised by a client call.
type SourceMap_Mapping_Promise struct{ *capnp.Pipeline }

func (p SourceMap_Mapping_Promise) Struct() (SourceMap_Mapping, error) {
	s, err := p.Pipeline.Struct()
	return SourceMap_Mapping{s}, err
}

type BridgeConfig struct{ capnp.Struct }

func NewBridgeConfig(s *capnp.Segment) (BridgeConfig, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return BridgeConfig{}, err
	}
	return BridgeConfig{st}, nil
}

func NewRootBridgeConfig(s *capnp.Segment) (BridgeConfig, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return BridgeConfig{}, err
	}
	return BridgeConfig{st}, nil
}

func ReadRootBridgeConfig(msg *capnp.Message) (BridgeConfig, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return BridgeConfig{}, err
	}
	return BridgeConfig{root.Struct()}, nil
}
func (s BridgeConfig) ViewInfo() (grain.UiView_ViewInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return grain.UiView_ViewInfo{}, err
	}
	return grain.UiView_ViewInfo{Struct: p.Struct()}, nil
}

func (s BridgeConfig) HasViewInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BridgeConfig) SetViewInfo(v grain.UiView_ViewInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewViewInfo sets the viewInfo field to a newly
// allocated grain.UiView_ViewInfo struct, preferring placement in s's segment.
func (s BridgeConfig) NewViewInfo() (grain.UiView_ViewInfo, error) {
	ss, err := grain.NewUiView_ViewInfo(s.Struct.Segment())
	if err != nil {
		return grain.UiView_ViewInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s BridgeConfig) ApiPath() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s BridgeConfig) HasApiPath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BridgeConfig) ApiPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s BridgeConfig) SetApiPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// BridgeConfig_List is a list of BridgeConfig.
type BridgeConfig_List struct{ capnp.List }

// NewBridgeConfig creates a new list of BridgeConfig.
func NewBridgeConfig_List(s *capnp.Segment, sz int32) (BridgeConfig_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return BridgeConfig_List{}, err
	}
	return BridgeConfig_List{l}, nil
}

func (s BridgeConfig_List) At(i int) BridgeConfig           { return BridgeConfig{s.List.Struct(i)} }
func (s BridgeConfig_List) Set(i int, v BridgeConfig) error { return s.List.SetStruct(i, v.Struct) }

// BridgeConfig_Promise is a wrapper for a BridgeConfig promised by a client call.
type BridgeConfig_Promise struct{ *capnp.Pipeline }

func (p BridgeConfig_Promise) Struct() (BridgeConfig, error) {
	s, err := p.Pipeline.Struct()
	return BridgeConfig{s}, err
}

func (p BridgeConfig_Promise) ViewInfo() grain.UiView_ViewInfo_Promise {
	return grain.UiView_ViewInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type KeyFile struct{ capnp.Struct }

func NewKeyFile(s *capnp.Segment) (KeyFile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return KeyFile{}, err
	}
	return KeyFile{st}, nil
}

func NewRootKeyFile(s *capnp.Segment) (KeyFile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return KeyFile{}, err
	}
	return KeyFile{st}, nil
}

func ReadRootKeyFile(msg *capnp.Message) (KeyFile, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return KeyFile{}, err
	}
	return KeyFile{root.Struct()}, nil
}
func (s KeyFile) PublicKey() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s KeyFile) HasPublicKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s KeyFile) SetPublicKey(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s KeyFile) PrivateKey() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s KeyFile) HasPrivateKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s KeyFile) SetPrivateKey(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

// KeyFile_List is a list of KeyFile.
type KeyFile_List struct{ capnp.List }

// NewKeyFile creates a new list of KeyFile.
func NewKeyFile_List(s *capnp.Segment, sz int32) (KeyFile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return KeyFile_List{}, err
	}
	return KeyFile_List{l}, nil
}

func (s KeyFile_List) At(i int) KeyFile           { return KeyFile{s.List.Struct(i)} }
func (s KeyFile_List) Set(i int, v KeyFile) error { return s.List.SetStruct(i, v.Struct) }

// KeyFile_Promise is a wrapper for a KeyFile promised by a client call.
type KeyFile_Promise struct{ *capnp.Pipeline }

func (p KeyFile_Promise) Struct() (KeyFile, error) {
	s, err := p.Pipeline.Struct()
	return KeyFile{s}, err
}

type Signature struct{ capnp.Struct }

func NewSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Signature{}, err
	}
	return Signature{st}, nil
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Signature{}, err
	}
	return Signature{st}, nil
}

func ReadRootSignature(msg *capnp.Message) (Signature, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Signature{}, err
	}
	return Signature{root.Struct()}, nil
}
func (s Signature) PublicKey() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Signature) HasPublicKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Signature) SetPublicKey(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Signature) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Signature) HasSignature() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Signature) SetSignature(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

// Signature_List is a list of Signature.
type Signature_List struct{ capnp.List }

// NewSignature creates a new list of Signature.
func NewSignature_List(s *capnp.Segment, sz int32) (Signature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Signature_List{}, err
	}
	return Signature_List{l}, nil
}

func (s Signature_List) At(i int) Signature           { return Signature{s.List.Struct(i)} }
func (s Signature_List) Set(i int, v Signature) error { return s.List.SetStruct(i, v.Struct) }

// Signature_Promise is a wrapper for a Signature promised by a client call.
type Signature_Promise struct{ *capnp.Pipeline }

func (p Signature_Promise) Struct() (Signature, error) {
	s, err := p.Pipeline.Struct()
	return Signature{s}, err
}

type Archive struct{ capnp.Struct }

func NewArchive(s *capnp.Segment) (Archive, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Archive{}, err
	}
	return Archive{st}, nil
}

func NewRootArchive(s *capnp.Segment) (Archive, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Archive{}, err
	}
	return Archive{st}, nil
}

func ReadRootArchive(msg *capnp.Message) (Archive, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Archive{}, err
	}
	return Archive{root.Struct()}, nil
}
func (s Archive) Files() (Archive_File_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Archive_File_List{}, err
	}
	return Archive_File_List{List: p.List()}, nil
}

func (s Archive) HasFiles() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Archive) SetFiles(v Archive_File_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewFiles sets the files field to a newly
// allocated Archive_File_List, preferring placement in s's segment.
func (s Archive) NewFiles(n int32) (Archive_File_List, error) {
	l, err := NewArchive_File_List(s.Struct.Segment(), n)
	if err != nil {
		return Archive_File_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Archive_List is a list of Archive.
type Archive_List struct{ capnp.List }

// NewArchive creates a new list of Archive.
func NewArchive_List(s *capnp.Segment, sz int32) (Archive_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Archive_List{}, err
	}
	return Archive_List{l}, nil
}

func (s Archive_List) At(i int) Archive           { return Archive{s.List.Struct(i)} }
func (s Archive_List) Set(i int, v Archive) error { return s.List.SetStruct(i, v.Struct) }

// Archive_Promise is a wrapper for a Archive promised by a client call.
type Archive_Promise struct{ *capnp.Pipeline }

func (p Archive_Promise) Struct() (Archive, error) {
	s, err := p.Pipeline.Struct()
	return Archive{s}, err
}

type Archive_File struct{ capnp.Struct }
type Archive_File_Which uint16

const (
	Archive_File_Which_regular    Archive_File_Which = 0
	Archive_File_Which_executable Archive_File_Which = 1
	Archive_File_Which_symlink    Archive_File_Which = 2
	Archive_File_Which_directory  Archive_File_Which = 3
)

func (w Archive_File_Which) String() string {
	const s = "regularexecutablesymlinkdirectory"
	switch w {
	case Archive_File_Which_regular:
		return s[0:7]
	case Archive_File_Which_executable:
		return s[7:17]
	case Archive_File_Which_symlink:
		return s[17:24]
	case Archive_File_Which_directory:
		return s[24:33]

	}
	return "Archive_File_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewArchive_File(s *capnp.Segment) (Archive_File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	if err != nil {
		return Archive_File{}, err
	}
	return Archive_File{st}, nil
}

func NewRootArchive_File(s *capnp.Segment) (Archive_File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	if err != nil {
		return Archive_File{}, err
	}
	return Archive_File{st}, nil
}

func ReadRootArchive_File(msg *capnp.Message) (Archive_File, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Archive_File{}, err
	}
	return Archive_File{root.Struct()}, nil
}

func (s Archive_File) Which() Archive_File_Which {
	return Archive_File_Which(s.Struct.Uint16(0))
}
func (s Archive_File) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Archive_File) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Archive_File) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Archive_File) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Archive_File) LastModificationTimeNs() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Archive_File) SetLastModificationTimeNs(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s Archive_File) Regular() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Archive_File) HasRegular() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Archive_File) SetRegular(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s Archive_File) Executable() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Archive_File) HasExecutable() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Archive_File) SetExecutable(v []byte) error {
	s.Struct.SetUint16(0, 1)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s Archive_File) Symlink() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Archive_File) HasSymlink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Archive_File) SymlinkBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Archive_File) SetSymlink(v string) error {
	s.Struct.SetUint16(0, 2)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Archive_File) Directory() (Archive_File_List, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Archive_File_List{}, err
	}
	return Archive_File_List{List: p.List()}, nil
}

func (s Archive_File) HasDirectory() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Archive_File) SetDirectory(v Archive_File_List) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Archive_File_List, preferring placement in s's segment.
func (s Archive_File) NewDirectory(n int32) (Archive_File_List, error) {
	s.Struct.SetUint16(0, 3)
	l, err := NewArchive_File_List(s.Struct.Segment(), n)
	if err != nil {
		return Archive_File_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Archive_File_List is a list of Archive_File.
type Archive_File_List struct{ capnp.List }

// NewArchive_File creates a new list of Archive_File.
func NewArchive_File_List(s *capnp.Segment, sz int32) (Archive_File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	if err != nil {
		return Archive_File_List{}, err
	}
	return Archive_File_List{l}, nil
}

func (s Archive_File_List) At(i int) Archive_File           { return Archive_File{s.List.Struct(i)} }
func (s Archive_File_List) Set(i int, v Archive_File) error { return s.List.SetStruct(i, v.Struct) }

// Archive_File_Promise is a wrapper for a Archive_File promised by a client call.
type Archive_File_Promise struct{ *capnp.Pipeline }

func (p Archive_File_Promise) Struct() (Archive_File, error) {
	s, err := p.Pipeline.Struct()
	return Archive_File{s}, err
}
