package spk

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	powerbox "zenhack.net/go/sandstorm/capnp/powerbox"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

const OsiInfo = uint64(0x9476412d0315d869)
const CategoryInfo = uint64(0x8d51dd236606d205)

// Constants defined in package.capnp.
const (
	Manifest_sizeLimitInWords = uint64(1048576)
)

// Constants defined in package.capnp.
var (
	MagicNumber = []byte{143, 198, 205, 239, 69, 26, 234, 150}
)

type PackageDefinition struct{ capnp.Struct }

func NewPackageDefinition(s *capnp.Segment) (PackageDefinition, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return PackageDefinition{st}, err
}

func NewRootPackageDefinition(s *capnp.Segment) (PackageDefinition, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return PackageDefinition{st}, err
}

func ReadRootPackageDefinition(msg *capnp.Message) (PackageDefinition, error) {
	root, err := msg.RootPtr()
	return PackageDefinition{root.Struct()}, err
}

func (s PackageDefinition) String() string {
	str, _ := text.Marshal(0x9f149fa71489be0b, s.Struct)
	return str
}

func (s PackageDefinition) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PackageDefinition) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
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
	return Manifest{Struct: p.Struct()}, err
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
	return SourceMap{Struct: p.Struct()}, err
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
	return p.Text(), err
}

func (s PackageDefinition) HasFileList() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s PackageDefinition) FileListBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
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
	return capnp.TextList{List: p.List()}, err
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
	return BridgeConfig{Struct: p.Struct()}, err
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
	return PackageDefinition_List{l}, err
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	return Manifest{st}, err
}

func NewRootManifest(s *capnp.Segment) (Manifest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	return Manifest{st}, err
}

func ReadRootManifest(msg *capnp.Message) (Manifest, error) {
	root, err := msg.RootPtr()
	return Manifest{root.Struct()}, err
}

func (s Manifest) String() string {
	str, _ := text.Marshal(0x855f296a69e6e1ca, s.Struct)
	return str
}

func (s Manifest) AppTitle() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(3)
	return util.LocalizedText{Struct: p.Struct()}, err
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
	return util.LocalizedText{Struct: p.Struct()}, err
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

func (s Manifest) Metadata() (Metadata, error) {
	p, err := s.Struct.Ptr(4)
	return Metadata{Struct: p.Struct()}, err
}

func (s Manifest) HasMetadata() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Manifest) SetMetadata(v Metadata) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewMetadata sets the metadata field to a newly
// allocated Metadata struct, preferring placement in s's segment.
func (s Manifest) NewMetadata() (Metadata, error) {
	ss, err := NewMetadata(s.Struct.Segment())
	if err != nil {
		return Metadata{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest) Actions() (Manifest_Action_List, error) {
	p, err := s.Struct.Ptr(0)
	return Manifest_Action_List{List: p.List()}, err
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
	return Manifest_Command{Struct: p.Struct()}, err
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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5}, sz)
	return Manifest_List{l}, err
}

func (s Manifest_List) At(i int) Manifest { return Manifest{s.List.Struct(i)} }

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

func (p Manifest_Promise) Metadata() Metadata_Promise {
	return Metadata_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Manifest_Promise) ContinueCommand() Manifest_Command_Promise {
	return Manifest_Command_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Manifest_Command struct{ capnp.Struct }

func NewManifest_Command(s *capnp.Segment) (Manifest_Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Manifest_Command{st}, err
}

func NewRootManifest_Command(s *capnp.Segment) (Manifest_Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Manifest_Command{st}, err
}

func ReadRootManifest_Command(msg *capnp.Message) (Manifest_Command, error) {
	root, err := msg.RootPtr()
	return Manifest_Command{root.Struct()}, err
}

func (s Manifest_Command) String() string {
	str, _ := text.Marshal(0xc64951b2a02886cf, s.Struct)
	return str
}

func (s Manifest_Command) Argv() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
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
	return util.KeyValue_List{List: p.List()}, err
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
	return p.Text(), err
}

func (s Manifest_Command) HasDeprecatedExecutablePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Manifest_Command) DeprecatedExecutablePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
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
	return Manifest_Command_List{l}, err
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Manifest_Action{st}, err
}

func NewRootManifest_Action(s *capnp.Segment) (Manifest_Action, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return Manifest_Action{st}, err
}

func ReadRootManifest_Action(msg *capnp.Message) (Manifest_Action, error) {
	root, err := msg.RootPtr()
	return Manifest_Action{root.Struct()}, err
}

func (s Manifest_Action) String() string {
	str, _ := text.Marshal(0xe5c59b9296375a00, s.Struct)
	return str
}

func (s Manifest_Action) Input() Manifest_Action_input { return Manifest_Action_input(s) }

func (s Manifest_Action_input) Which() Manifest_Action_input_Which {
	return Manifest_Action_input_Which(s.Struct.Uint16(0))
}
func (s Manifest_Action_input) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s Manifest_Action_input) Capability() (powerbox.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(0)
	return powerbox.PowerboxDescriptor_List{List: p.List()}, err
}

func (s Manifest_Action_input) HasCapability() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Manifest_Action_input) SetCapability(v powerbox.PowerboxDescriptor_List) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewCapability sets the capability field to a newly
// allocated powerbox.PowerboxDescriptor_List, preferring placement in s's segment.
func (s Manifest_Action_input) NewCapability(n int32) (powerbox.PowerboxDescriptor_List, error) {
	s.Struct.SetUint16(0, 1)
	l, err := powerbox.NewPowerboxDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return powerbox.PowerboxDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Manifest_Action) Command() (Manifest_Command, error) {
	p, err := s.Struct.Ptr(1)
	return Manifest_Command{Struct: p.Struct()}, err
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
	return util.LocalizedText{Struct: p.Struct()}, err
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

func (s Manifest_Action) NounPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(4)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Manifest_Action) HasNounPhrase() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Manifest_Action) SetNounPhrase(v util.LocalizedText) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewNounPhrase sets the nounPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewNounPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Manifest_Action) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(3)
	return util.LocalizedText{Struct: p.Struct()}, err
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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return Manifest_Action_List{l}, err
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

func (p Manifest_Action_Promise) NounPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Manifest_Action_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SourceMap struct{ capnp.Struct }

func NewSourceMap(s *capnp.Segment) (SourceMap, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SourceMap{st}, err
}

func NewRootSourceMap(s *capnp.Segment) (SourceMap, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SourceMap{st}, err
}

func ReadRootSourceMap(msg *capnp.Message) (SourceMap, error) {
	root, err := msg.RootPtr()
	return SourceMap{root.Struct()}, err
}

func (s SourceMap) String() string {
	str, _ := text.Marshal(0xe3d7ba482b2e470b, s.Struct)
	return str
}

func (s SourceMap) SearchPath() (SourceMap_Mapping_List, error) {
	p, err := s.Struct.Ptr(0)
	return SourceMap_Mapping_List{List: p.List()}, err
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
	return SourceMap_List{l}, err
}

func (s SourceMap_List) At(i int) SourceMap { return SourceMap{s.List.Struct(i)} }

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
	return SourceMap_Mapping{st}, err
}

func NewRootSourceMap_Mapping(s *capnp.Segment) (SourceMap_Mapping, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return SourceMap_Mapping{st}, err
}

func ReadRootSourceMap_Mapping(msg *capnp.Message) (SourceMap_Mapping, error) {
	root, err := msg.RootPtr()
	return SourceMap_Mapping{root.Struct()}, err
}

func (s SourceMap_Mapping) String() string {
	str, _ := text.Marshal(0x87dcf1b1edcb3eaf, s.Struct)
	return str
}

func (s SourceMap_Mapping) PackagePath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SourceMap_Mapping) HasPackagePath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SourceMap_Mapping) PackagePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
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
	return p.Text(), err
}

func (s SourceMap_Mapping) HasSourcePath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SourceMap_Mapping) SourcePathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
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
	return capnp.TextList{List: p.List()}, err
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
	return SourceMap_Mapping_List{l}, err
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BridgeConfig{st}, err
}

func NewRootBridgeConfig(s *capnp.Segment) (BridgeConfig, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BridgeConfig{st}, err
}

func ReadRootBridgeConfig(msg *capnp.Message) (BridgeConfig, error) {
	root, err := msg.RootPtr()
	return BridgeConfig{root.Struct()}, err
}

func (s BridgeConfig) String() string {
	str, _ := text.Marshal(0xdd8c82383168c096, s.Struct)
	return str
}

func (s BridgeConfig) ViewInfo() (grain.UiView_ViewInfo, error) {
	p, err := s.Struct.Ptr(0)
	return grain.UiView_ViewInfo{Struct: p.Struct()}, err
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
	return p.Text(), err
}

func (s BridgeConfig) HasApiPath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BridgeConfig) ApiPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s BridgeConfig) SetApiPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s BridgeConfig) SaveIdentityCaps() bool {
	return s.Struct.Bit(0)
}

func (s BridgeConfig) SetSaveIdentityCaps(v bool) {
	s.Struct.SetBit(0, v)
}

// BridgeConfig_List is a list of BridgeConfig.
type BridgeConfig_List struct{ capnp.List }

// NewBridgeConfig creates a new list of BridgeConfig.
func NewBridgeConfig_List(s *capnp.Segment, sz int32) (BridgeConfig_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return BridgeConfig_List{l}, err
}

func (s BridgeConfig_List) At(i int) BridgeConfig { return BridgeConfig{s.List.Struct(i)} }

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

type Metadata struct{ capnp.Struct }
type Metadata_icons Metadata
type Metadata_license Metadata
type Metadata_author Metadata
type Metadata_license_Which uint16

const (
	Metadata_license_Which_none         Metadata_license_Which = 0
	Metadata_license_Which_openSource   Metadata_license_Which = 1
	Metadata_license_Which_proprietary  Metadata_license_Which = 2
	Metadata_license_Which_publicDomain Metadata_license_Which = 3
)

func (w Metadata_license_Which) String() string {
	const s = "noneopenSourceproprietarypublicDomain"
	switch w {
	case Metadata_license_Which_none:
		return s[0:4]
	case Metadata_license_Which_openSource:
		return s[4:14]
	case Metadata_license_Which_proprietary:
		return s[14:25]
	case Metadata_license_Which_publicDomain:
		return s[25:37]

	}
	return "Metadata_license_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewMetadata(s *capnp.Segment) (Metadata, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 17})
	return Metadata{st}, err
}

func NewRootMetadata(s *capnp.Segment) (Metadata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 17})
	return Metadata{st}, err
}

func ReadRootMetadata(msg *capnp.Message) (Metadata, error) {
	root, err := msg.RootPtr()
	return Metadata{root.Struct()}, err
}

func (s Metadata) String() string {
	str, _ := text.Marshal(0xe0c5892a5448f4ee, s.Struct)
	return str
}

func (s Metadata) Icons() Metadata_icons { return Metadata_icons(s) }

func (s Metadata_icons) AppGrid() (Metadata_Icon, error) {
	p, err := s.Struct.Ptr(0)
	return Metadata_Icon{Struct: p.Struct()}, err
}

func (s Metadata_icons) HasAppGrid() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Metadata_icons) SetAppGrid(v Metadata_Icon) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAppGrid sets the appGrid field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewAppGrid() (Metadata_Icon, error) {
	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata_icons) Grain() (Metadata_Icon, error) {
	p, err := s.Struct.Ptr(1)
	return Metadata_Icon{Struct: p.Struct()}, err
}

func (s Metadata_icons) HasGrain() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Metadata_icons) SetGrain(v Metadata_Icon) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewGrain sets the grain field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewGrain() (Metadata_Icon, error) {
	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata_icons) Market() (Metadata_Icon, error) {
	p, err := s.Struct.Ptr(2)
	return Metadata_Icon{Struct: p.Struct()}, err
}

func (s Metadata_icons) HasMarket() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Metadata_icons) SetMarket(v Metadata_Icon) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewMarket sets the market field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewMarket() (Metadata_Icon, error) {
	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata_icons) MarketBig() (Metadata_Icon, error) {
	p, err := s.Struct.Ptr(15)
	return Metadata_Icon{Struct: p.Struct()}, err
}

func (s Metadata_icons) HasMarketBig() bool {
	p, err := s.Struct.Ptr(15)
	return p.IsValid() || err != nil
}

func (s Metadata_icons) SetMarketBig(v Metadata_Icon) error {
	return s.Struct.SetPtr(15, v.Struct.ToPtr())
}

// NewMarketBig sets the marketBig field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewMarketBig() (Metadata_Icon, error) {
	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPtr(15, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata) Website() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s Metadata) HasWebsite() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Metadata) WebsiteBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s Metadata) SetWebsite(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s Metadata) CodeUrl() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Metadata) HasCodeUrl() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Metadata) CodeUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Metadata) SetCodeUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

func (s Metadata) License() Metadata_license { return Metadata_license(s) }

func (s Metadata_license) Which() Metadata_license_Which {
	return Metadata_license_Which(s.Struct.Uint16(0))
}
func (s Metadata_license) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s Metadata_license) OpenSource() OpenSourceLicense {
	return OpenSourceLicense(s.Struct.Uint16(2))
}

func (s Metadata_license) SetOpenSource(v OpenSourceLicense) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint16(2, uint16(v))
}

func (s Metadata_license) Proprietary() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(5)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata_license) HasProprietary() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Metadata_license) SetProprietary(v util.LocalizedText) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewProprietary sets the proprietary field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewProprietary() (util.LocalizedText, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata_license) PublicDomain() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(5)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata_license) HasPublicDomain() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Metadata_license) SetPublicDomain(v util.LocalizedText) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewPublicDomain sets the publicDomain field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewPublicDomain() (util.LocalizedText, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata_license) Notices() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(6)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata_license) HasNotices() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Metadata_license) SetNotices(v util.LocalizedText) error {
	return s.Struct.SetPtr(6, v.Struct.ToPtr())
}

// NewNotices sets the notices field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewNotices() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(6, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata) Categories() (Category_List, error) {
	p, err := s.Struct.Ptr(7)
	return Category_List{List: p.List()}, err
}

func (s Metadata) HasCategories() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s Metadata) SetCategories(v Category_List) error {
	return s.Struct.SetPtr(7, v.List.ToPtr())
}

// NewCategories sets the categories field to a newly
// allocated Category_List, preferring placement in s's segment.
func (s Metadata) NewCategories(n int32) (Category_List, error) {
	l, err := NewCategory_List(s.Struct.Segment(), n)
	if err != nil {
		return Category_List{}, err
	}
	err = s.Struct.SetPtr(7, l.List.ToPtr())
	return l, err
}

func (s Metadata) Author() Metadata_author { return Metadata_author(s) }

func (s Metadata_author) UpstreamAuthor() (string, error) {
	p, err := s.Struct.Ptr(16)
	return p.Text(), err
}

func (s Metadata_author) HasUpstreamAuthor() bool {
	p, err := s.Struct.Ptr(16)
	return p.IsValid() || err != nil
}

func (s Metadata_author) UpstreamAuthorBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(16)
	return p.TextBytes(), err
}

func (s Metadata_author) SetUpstreamAuthor(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(16, t.List.ToPtr())
}

func (s Metadata_author) ContactEmail() (string, error) {
	p, err := s.Struct.Ptr(8)
	return p.Text(), err
}

func (s Metadata_author) HasContactEmail() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s Metadata_author) ContactEmailBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(8)
	return p.TextBytes(), err
}

func (s Metadata_author) SetContactEmail(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(8, t.List.ToPtr())
}

func (s Metadata_author) PgpSignature() ([]byte, error) {
	p, err := s.Struct.Ptr(9)
	return []byte(p.Data()), err
}

func (s Metadata_author) HasPgpSignature() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s Metadata_author) SetPgpSignature(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(9, d.List.ToPtr())
}

func (s Metadata) PgpKeyring() ([]byte, error) {
	p, err := s.Struct.Ptr(10)
	return []byte(p.Data()), err
}

func (s Metadata) HasPgpKeyring() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s Metadata) SetPgpKeyring(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(10, d.List.ToPtr())
}

func (s Metadata) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(11)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata) HasDescription() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s Metadata) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(11, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(11, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata) ShortDescription() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(12)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata) HasShortDescription() bool {
	p, err := s.Struct.Ptr(12)
	return p.IsValid() || err != nil
}

func (s Metadata) SetShortDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(12, v.Struct.ToPtr())
}

// NewShortDescription sets the shortDescription field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewShortDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(12, ss.Struct.ToPtr())
	return ss, err
}

func (s Metadata) Screenshots() (Metadata_Screenshot_List, error) {
	p, err := s.Struct.Ptr(13)
	return Metadata_Screenshot_List{List: p.List()}, err
}

func (s Metadata) HasScreenshots() bool {
	p, err := s.Struct.Ptr(13)
	return p.IsValid() || err != nil
}

func (s Metadata) SetScreenshots(v Metadata_Screenshot_List) error {
	return s.Struct.SetPtr(13, v.List.ToPtr())
}

// NewScreenshots sets the screenshots field to a newly
// allocated Metadata_Screenshot_List, preferring placement in s's segment.
func (s Metadata) NewScreenshots(n int32) (Metadata_Screenshot_List, error) {
	l, err := NewMetadata_Screenshot_List(s.Struct.Segment(), n)
	if err != nil {
		return Metadata_Screenshot_List{}, err
	}
	err = s.Struct.SetPtr(13, l.List.ToPtr())
	return l, err
}

func (s Metadata) ChangeLog() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(14)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Metadata) HasChangeLog() bool {
	p, err := s.Struct.Ptr(14)
	return p.IsValid() || err != nil
}

func (s Metadata) SetChangeLog(v util.LocalizedText) error {
	return s.Struct.SetPtr(14, v.Struct.ToPtr())
}

// NewChangeLog sets the changeLog field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewChangeLog() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(14, ss.Struct.ToPtr())
	return ss, err
}

// Metadata_List is a list of Metadata.
type Metadata_List struct{ capnp.List }

// NewMetadata creates a new list of Metadata.
func NewMetadata_List(s *capnp.Segment, sz int32) (Metadata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 17}, sz)
	return Metadata_List{l}, err
}

func (s Metadata_List) At(i int) Metadata { return Metadata{s.List.Struct(i)} }

func (s Metadata_List) Set(i int, v Metadata) error { return s.List.SetStruct(i, v.Struct) }

// Metadata_Promise is a wrapper for a Metadata promised by a client call.
type Metadata_Promise struct{ *capnp.Pipeline }

func (p Metadata_Promise) Struct() (Metadata, error) {
	s, err := p.Pipeline.Struct()
	return Metadata{s}, err
}

func (p Metadata_Promise) Icons() Metadata_icons_Promise { return Metadata_icons_Promise{p.Pipeline} }

// Metadata_icons_Promise is a wrapper for a Metadata_icons promised by a client call.
type Metadata_icons_Promise struct{ *capnp.Pipeline }

func (p Metadata_icons_Promise) Struct() (Metadata_icons, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_icons{s}, err
}

func (p Metadata_icons_Promise) AppGrid() Metadata_Icon_Promise {
	return Metadata_Icon_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Metadata_icons_Promise) Grain() Metadata_Icon_Promise {
	return Metadata_Icon_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Metadata_icons_Promise) Market() Metadata_Icon_Promise {
	return Metadata_Icon_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Metadata_icons_Promise) MarketBig() Metadata_Icon_Promise {
	return Metadata_Icon_Promise{Pipeline: p.Pipeline.GetPipeline(15)}
}

func (p Metadata_Promise) License() Metadata_license_Promise {
	return Metadata_license_Promise{p.Pipeline}
}

// Metadata_license_Promise is a wrapper for a Metadata_license promised by a client call.
type Metadata_license_Promise struct{ *capnp.Pipeline }

func (p Metadata_license_Promise) Struct() (Metadata_license, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_license{s}, err
}

func (p Metadata_license_Promise) Proprietary() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Metadata_license_Promise) PublicDomain() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Metadata_license_Promise) Notices() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(6)}
}

func (p Metadata_Promise) Author() Metadata_author_Promise { return Metadata_author_Promise{p.Pipeline} }

// Metadata_author_Promise is a wrapper for a Metadata_author promised by a client call.
type Metadata_author_Promise struct{ *capnp.Pipeline }

func (p Metadata_author_Promise) Struct() (Metadata_author, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_author{s}, err
}

func (p Metadata_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(11)}
}

func (p Metadata_Promise) ShortDescription() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(12)}
}

func (p Metadata_Promise) ChangeLog() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(14)}
}

type Metadata_Icon struct{ capnp.Struct }
type Metadata_Icon_png Metadata_Icon
type Metadata_Icon_Which uint16

const (
	Metadata_Icon_Which_unknown Metadata_Icon_Which = 0
	Metadata_Icon_Which_svg     Metadata_Icon_Which = 1
	Metadata_Icon_Which_png     Metadata_Icon_Which = 2
)

func (w Metadata_Icon_Which) String() string {
	const s = "unknownsvgpng"
	switch w {
	case Metadata_Icon_Which_unknown:
		return s[0:7]
	case Metadata_Icon_Which_svg:
		return s[7:10]
	case Metadata_Icon_Which_png:
		return s[10:13]

	}
	return "Metadata_Icon_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewMetadata_Icon(s *capnp.Segment) (Metadata_Icon, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Metadata_Icon{st}, err
}

func NewRootMetadata_Icon(s *capnp.Segment) (Metadata_Icon, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Metadata_Icon{st}, err
}

func ReadRootMetadata_Icon(msg *capnp.Message) (Metadata_Icon, error) {
	root, err := msg.RootPtr()
	return Metadata_Icon{root.Struct()}, err
}

func (s Metadata_Icon) String() string {
	str, _ := text.Marshal(0xe492a2981208ad0b, s.Struct)
	return str
}

func (s Metadata_Icon) Which() Metadata_Icon_Which {
	return Metadata_Icon_Which(s.Struct.Uint16(0))
}
func (s Metadata_Icon) SetUnknown() {
	s.Struct.SetUint16(0, 0)

}

func (s Metadata_Icon) Svg() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Metadata_Icon) HasSvg() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Metadata_Icon) SvgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Metadata_Icon) SetSvg(v string) error {
	s.Struct.SetUint16(0, 1)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Metadata_Icon) Png() Metadata_Icon_png { return Metadata_Icon_png(s) }

func (s Metadata_Icon) SetPng() {
	s.Struct.SetUint16(0, 2)
}

func (s Metadata_Icon_png) Dpi1x() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Metadata_Icon_png) HasDpi1x() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Metadata_Icon_png) SetDpi1x(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Metadata_Icon_png) Dpi2x() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Metadata_Icon_png) HasDpi2x() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Metadata_Icon_png) SetDpi2x(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

// Metadata_Icon_List is a list of Metadata_Icon.
type Metadata_Icon_List struct{ capnp.List }

// NewMetadata_Icon creates a new list of Metadata_Icon.
func NewMetadata_Icon_List(s *capnp.Segment, sz int32) (Metadata_Icon_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Metadata_Icon_List{l}, err
}

func (s Metadata_Icon_List) At(i int) Metadata_Icon { return Metadata_Icon{s.List.Struct(i)} }

func (s Metadata_Icon_List) Set(i int, v Metadata_Icon) error { return s.List.SetStruct(i, v.Struct) }

// Metadata_Icon_Promise is a wrapper for a Metadata_Icon promised by a client call.
type Metadata_Icon_Promise struct{ *capnp.Pipeline }

func (p Metadata_Icon_Promise) Struct() (Metadata_Icon, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_Icon{s}, err
}

func (p Metadata_Icon_Promise) Png() Metadata_Icon_png_Promise {
	return Metadata_Icon_png_Promise{p.Pipeline}
}

// Metadata_Icon_png_Promise is a wrapper for a Metadata_Icon_png promised by a client call.
type Metadata_Icon_png_Promise struct{ *capnp.Pipeline }

func (p Metadata_Icon_png_Promise) Struct() (Metadata_Icon_png, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_Icon_png{s}, err
}

type Metadata_Screenshot struct{ capnp.Struct }
type Metadata_Screenshot_Which uint16

const (
	Metadata_Screenshot_Which_unknown Metadata_Screenshot_Which = 0
	Metadata_Screenshot_Which_png     Metadata_Screenshot_Which = 1
	Metadata_Screenshot_Which_jpeg    Metadata_Screenshot_Which = 2
)

func (w Metadata_Screenshot_Which) String() string {
	const s = "unknownpngjpeg"
	switch w {
	case Metadata_Screenshot_Which_unknown:
		return s[0:7]
	case Metadata_Screenshot_Which_png:
		return s[7:10]
	case Metadata_Screenshot_Which_jpeg:
		return s[10:14]

	}
	return "Metadata_Screenshot_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewMetadata_Screenshot(s *capnp.Segment) (Metadata_Screenshot, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Metadata_Screenshot{st}, err
}

func NewRootMetadata_Screenshot(s *capnp.Segment) (Metadata_Screenshot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Metadata_Screenshot{st}, err
}

func ReadRootMetadata_Screenshot(msg *capnp.Message) (Metadata_Screenshot, error) {
	root, err := msg.RootPtr()
	return Metadata_Screenshot{root.Struct()}, err
}

func (s Metadata_Screenshot) String() string {
	str, _ := text.Marshal(0x8bc9f4365959348e, s.Struct)
	return str
}

func (s Metadata_Screenshot) Which() Metadata_Screenshot_Which {
	return Metadata_Screenshot_Which(s.Struct.Uint16(8))
}
func (s Metadata_Screenshot) Width() uint32 {
	return s.Struct.Uint32(0)
}

func (s Metadata_Screenshot) SetWidth(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Metadata_Screenshot) Height() uint32 {
	return s.Struct.Uint32(4)
}

func (s Metadata_Screenshot) SetHeight(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Metadata_Screenshot) SetUnknown() {
	s.Struct.SetUint16(8, 0)

}

func (s Metadata_Screenshot) Png() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Metadata_Screenshot) HasPng() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Metadata_Screenshot) SetPng(v []byte) error {
	s.Struct.SetUint16(8, 1)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Metadata_Screenshot) Jpeg() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Metadata_Screenshot) HasJpeg() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Metadata_Screenshot) SetJpeg(v []byte) error {
	s.Struct.SetUint16(8, 2)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// Metadata_Screenshot_List is a list of Metadata_Screenshot.
type Metadata_Screenshot_List struct{ capnp.List }

// NewMetadata_Screenshot creates a new list of Metadata_Screenshot.
func NewMetadata_Screenshot_List(s *capnp.Segment, sz int32) (Metadata_Screenshot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Metadata_Screenshot_List{l}, err
}

func (s Metadata_Screenshot_List) At(i int) Metadata_Screenshot {
	return Metadata_Screenshot{s.List.Struct(i)}
}

func (s Metadata_Screenshot_List) Set(i int, v Metadata_Screenshot) error {
	return s.List.SetStruct(i, v.Struct)
}

// Metadata_Screenshot_Promise is a wrapper for a Metadata_Screenshot promised by a client call.
type Metadata_Screenshot_Promise struct{ *capnp.Pipeline }

func (p Metadata_Screenshot_Promise) Struct() (Metadata_Screenshot, error) {
	s, err := p.Pipeline.Struct()
	return Metadata_Screenshot{s}, err
}

type OsiLicenseInfo struct{ capnp.Struct }

func NewOsiLicenseInfo(s *capnp.Segment) (OsiLicenseInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return OsiLicenseInfo{st}, err
}

func NewRootOsiLicenseInfo(s *capnp.Segment) (OsiLicenseInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return OsiLicenseInfo{st}, err
}

func ReadRootOsiLicenseInfo(msg *capnp.Message) (OsiLicenseInfo, error) {
	root, err := msg.RootPtr()
	return OsiLicenseInfo{root.Struct()}, err
}

func (s OsiLicenseInfo) String() string {
	str, _ := text.Marshal(0xb755d258845a4a8f, s.Struct)
	return str
}

func (s OsiLicenseInfo) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s OsiLicenseInfo) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s OsiLicenseInfo) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s OsiLicenseInfo) SetId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s OsiLicenseInfo) Title() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s OsiLicenseInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s OsiLicenseInfo) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s OsiLicenseInfo) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s OsiLicenseInfo) RequireSource() bool {
	return s.Struct.Bit(0)
}

func (s OsiLicenseInfo) SetRequireSource(v bool) {
	s.Struct.SetBit(0, v)
}

// OsiLicenseInfo_List is a list of OsiLicenseInfo.
type OsiLicenseInfo_List struct{ capnp.List }

// NewOsiLicenseInfo creates a new list of OsiLicenseInfo.
func NewOsiLicenseInfo_List(s *capnp.Segment, sz int32) (OsiLicenseInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return OsiLicenseInfo_List{l}, err
}

func (s OsiLicenseInfo_List) At(i int) OsiLicenseInfo { return OsiLicenseInfo{s.List.Struct(i)} }

func (s OsiLicenseInfo_List) Set(i int, v OsiLicenseInfo) error { return s.List.SetStruct(i, v.Struct) }

// OsiLicenseInfo_Promise is a wrapper for a OsiLicenseInfo promised by a client call.
type OsiLicenseInfo_Promise struct{ *capnp.Pipeline }

func (p OsiLicenseInfo_Promise) Struct() (OsiLicenseInfo, error) {
	s, err := p.Pipeline.Struct()
	return OsiLicenseInfo{s}, err
}

type OpenSourceLicense uint16

// Values of OpenSourceLicense.
const (
	OpenSourceLicense_invalid    OpenSourceLicense = 0
	OpenSourceLicense_mit        OpenSourceLicense = 1
	OpenSourceLicense_apache2    OpenSourceLicense = 2
	OpenSourceLicense_gpl3       OpenSourceLicense = 3
	OpenSourceLicense_agpl3      OpenSourceLicense = 4
	OpenSourceLicense_bsd3Clause OpenSourceLicense = 5
	OpenSourceLicense_bsd2Clause OpenSourceLicense = 6
	OpenSourceLicense_gpl2       OpenSourceLicense = 7
	OpenSourceLicense_lgpl2      OpenSourceLicense = 8
	OpenSourceLicense_lgpl3      OpenSourceLicense = 9
	OpenSourceLicense_isc        OpenSourceLicense = 10
	OpenSourceLicense_artistic2  OpenSourceLicense = 11
	OpenSourceLicense_python2    OpenSourceLicense = 12
	OpenSourceLicense_php3       OpenSourceLicense = 13
	OpenSourceLicense_mpl2       OpenSourceLicense = 14
	OpenSourceLicense_cddl       OpenSourceLicense = 15
	OpenSourceLicense_epl        OpenSourceLicense = 16
	OpenSourceLicense_cpal       OpenSourceLicense = 17
	OpenSourceLicense_zlib       OpenSourceLicense = 18
)

// String returns the enum's constant name.
func (c OpenSourceLicense) String() string {
	switch c {
	case OpenSourceLicense_invalid:
		return "invalid"
	case OpenSourceLicense_mit:
		return "mit"
	case OpenSourceLicense_apache2:
		return "apache2"
	case OpenSourceLicense_gpl3:
		return "gpl3"
	case OpenSourceLicense_agpl3:
		return "agpl3"
	case OpenSourceLicense_bsd3Clause:
		return "bsd3Clause"
	case OpenSourceLicense_bsd2Clause:
		return "bsd2Clause"
	case OpenSourceLicense_gpl2:
		return "gpl2"
	case OpenSourceLicense_lgpl2:
		return "lgpl2"
	case OpenSourceLicense_lgpl3:
		return "lgpl3"
	case OpenSourceLicense_isc:
		return "isc"
	case OpenSourceLicense_artistic2:
		return "artistic2"
	case OpenSourceLicense_python2:
		return "python2"
	case OpenSourceLicense_php3:
		return "php3"
	case OpenSourceLicense_mpl2:
		return "mpl2"
	case OpenSourceLicense_cddl:
		return "cddl"
	case OpenSourceLicense_epl:
		return "epl"
	case OpenSourceLicense_cpal:
		return "cpal"
	case OpenSourceLicense_zlib:
		return "zlib"

	default:
		return ""
	}
}

// OpenSourceLicenseFromString returns the enum value with a name,
// or the zero value if there's no such value.
func OpenSourceLicenseFromString(c string) OpenSourceLicense {
	switch c {
	case "invalid":
		return OpenSourceLicense_invalid
	case "mit":
		return OpenSourceLicense_mit
	case "apache2":
		return OpenSourceLicense_apache2
	case "gpl3":
		return OpenSourceLicense_gpl3
	case "agpl3":
		return OpenSourceLicense_agpl3
	case "bsd3Clause":
		return OpenSourceLicense_bsd3Clause
	case "bsd2Clause":
		return OpenSourceLicense_bsd2Clause
	case "gpl2":
		return OpenSourceLicense_gpl2
	case "lgpl2":
		return OpenSourceLicense_lgpl2
	case "lgpl3":
		return OpenSourceLicense_lgpl3
	case "isc":
		return OpenSourceLicense_isc
	case "artistic2":
		return OpenSourceLicense_artistic2
	case "python2":
		return OpenSourceLicense_python2
	case "php3":
		return OpenSourceLicense_php3
	case "mpl2":
		return OpenSourceLicense_mpl2
	case "cddl":
		return OpenSourceLicense_cddl
	case "epl":
		return OpenSourceLicense_epl
	case "cpal":
		return OpenSourceLicense_cpal
	case "zlib":
		return OpenSourceLicense_zlib

	default:
		return 0
	}
}

type OpenSourceLicense_List struct{ capnp.List }

func NewOpenSourceLicense_List(s *capnp.Segment, sz int32) (OpenSourceLicense_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return OpenSourceLicense_List{l.List}, err
}

func (l OpenSourceLicense_List) At(i int) OpenSourceLicense {
	ul := capnp.UInt16List{List: l.List}
	return OpenSourceLicense(ul.At(i))
}

func (l OpenSourceLicense_List) Set(i int, v OpenSourceLicense) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type AppId struct{ capnp.Struct }

func NewAppId(s *capnp.Segment) (AppId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return AppId{st}, err
}

func NewRootAppId(s *capnp.Segment) (AppId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return AppId{st}, err
}

func ReadRootAppId(msg *capnp.Message) (AppId, error) {
	root, err := msg.RootPtr()
	return AppId{root.Struct()}, err
}

func (s AppId) String() string {
	str, _ := text.Marshal(0x880c6c7782a33310, s.Struct)
	return str
}

func (s AppId) Id0() uint64 {
	return s.Struct.Uint64(0)
}

func (s AppId) SetId0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s AppId) Id1() uint64 {
	return s.Struct.Uint64(8)
}

func (s AppId) SetId1(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s AppId) Id2() uint64 {
	return s.Struct.Uint64(16)
}

func (s AppId) SetId2(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s AppId) Id3() uint64 {
	return s.Struct.Uint64(24)
}

func (s AppId) SetId3(v uint64) {
	s.Struct.SetUint64(24, v)
}

// AppId_List is a list of AppId.
type AppId_List struct{ capnp.List }

// NewAppId creates a new list of AppId.
func NewAppId_List(s *capnp.Segment, sz int32) (AppId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0}, sz)
	return AppId_List{l}, err
}

func (s AppId_List) At(i int) AppId { return AppId{s.List.Struct(i)} }

func (s AppId_List) Set(i int, v AppId) error { return s.List.SetStruct(i, v.Struct) }

// AppId_Promise is a wrapper for a AppId promised by a client call.
type AppId_Promise struct{ *capnp.Pipeline }

func (p AppId_Promise) Struct() (AppId, error) {
	s, err := p.Pipeline.Struct()
	return AppId{s}, err
}

type PackageId struct{ capnp.Struct }

func NewPackageId(s *capnp.Segment) (PackageId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return PackageId{st}, err
}

func NewRootPackageId(s *capnp.Segment) (PackageId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return PackageId{st}, err
}

func ReadRootPackageId(msg *capnp.Message) (PackageId, error) {
	root, err := msg.RootPtr()
	return PackageId{root.Struct()}, err
}

func (s PackageId) String() string {
	str, _ := text.Marshal(0xe2e344d346ffda6b, s.Struct)
	return str
}

func (s PackageId) Id0() uint64 {
	return s.Struct.Uint64(0)
}

func (s PackageId) SetId0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s PackageId) Id1() uint64 {
	return s.Struct.Uint64(8)
}

func (s PackageId) SetId1(v uint64) {
	s.Struct.SetUint64(8, v)
}

// PackageId_List is a list of PackageId.
type PackageId_List struct{ capnp.List }

// NewPackageId creates a new list of PackageId.
func NewPackageId_List(s *capnp.Segment, sz int32) (PackageId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return PackageId_List{l}, err
}

func (s PackageId_List) At(i int) PackageId { return PackageId{s.List.Struct(i)} }

func (s PackageId_List) Set(i int, v PackageId) error { return s.List.SetStruct(i, v.Struct) }

// PackageId_Promise is a wrapper for a PackageId promised by a client call.
type PackageId_Promise struct{ *capnp.Pipeline }

func (p PackageId_Promise) Struct() (PackageId, error) {
	s, err := p.Pipeline.Struct()
	return PackageId{s}, err
}

type VerifiedInfo struct{ capnp.Struct }

func NewVerifiedInfo(s *capnp.Segment) (VerifiedInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return VerifiedInfo{st}, err
}

func NewRootVerifiedInfo(s *capnp.Segment) (VerifiedInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return VerifiedInfo{st}, err
}

func ReadRootVerifiedInfo(msg *capnp.Message) (VerifiedInfo, error) {
	root, err := msg.RootPtr()
	return VerifiedInfo{root.Struct()}, err
}

func (s VerifiedInfo) String() string {
	str, _ := text.Marshal(0x987ef3040a0342a9, s.Struct)
	return str
}

func (s VerifiedInfo) AppId() (AppId, error) {
	p, err := s.Struct.Ptr(0)
	return AppId{Struct: p.Struct()}, err
}

func (s VerifiedInfo) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) SetAppId(v AppId) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAppId sets the appId field to a newly
// allocated AppId struct, preferring placement in s's segment.
func (s VerifiedInfo) NewAppId() (AppId, error) {
	ss, err := NewAppId(s.Struct.Segment())
	if err != nil {
		return AppId{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedInfo) PackageId() (PackageId, error) {
	p, err := s.Struct.Ptr(1)
	return PackageId{Struct: p.Struct()}, err
}

func (s VerifiedInfo) HasPackageId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) SetPackageId(v PackageId) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPackageId sets the packageId field to a newly
// allocated PackageId struct, preferring placement in s's segment.
func (s VerifiedInfo) NewPackageId() (PackageId, error) {
	ss, err := NewPackageId(s.Struct.Segment())
	if err != nil {
		return PackageId{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s VerifiedInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s VerifiedInfo) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedInfo) Version() uint32 {
	return s.Struct.Uint32(0)
}

func (s VerifiedInfo) SetVersion(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s VerifiedInfo) MarketingVersion() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(3)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s VerifiedInfo) HasMarketingVersion() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) SetMarketingVersion(v util.LocalizedText) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewMarketingVersion sets the marketingVersion field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s VerifiedInfo) NewMarketingVersion() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedInfo) AuthorPgpKeyFingerprint() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s VerifiedInfo) HasAuthorPgpKeyFingerprint() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) AuthorPgpKeyFingerprintBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s VerifiedInfo) SetAuthorPgpKeyFingerprint(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

func (s VerifiedInfo) Metadata() (Metadata, error) {
	p, err := s.Struct.Ptr(5)
	return Metadata{Struct: p.Struct()}, err
}

func (s VerifiedInfo) HasMetadata() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s VerifiedInfo) SetMetadata(v Metadata) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewMetadata sets the metadata field to a newly
// allocated Metadata struct, preferring placement in s's segment.
func (s VerifiedInfo) NewMetadata() (Metadata, error) {
	ss, err := NewMetadata(s.Struct.Segment())
	if err != nil {
		return Metadata{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

// VerifiedInfo_List is a list of VerifiedInfo.
type VerifiedInfo_List struct{ capnp.List }

// NewVerifiedInfo creates a new list of VerifiedInfo.
func NewVerifiedInfo_List(s *capnp.Segment, sz int32) (VerifiedInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return VerifiedInfo_List{l}, err
}

func (s VerifiedInfo_List) At(i int) VerifiedInfo { return VerifiedInfo{s.List.Struct(i)} }

func (s VerifiedInfo_List) Set(i int, v VerifiedInfo) error { return s.List.SetStruct(i, v.Struct) }

// VerifiedInfo_Promise is a wrapper for a VerifiedInfo promised by a client call.
type VerifiedInfo_Promise struct{ *capnp.Pipeline }

func (p VerifiedInfo_Promise) Struct() (VerifiedInfo, error) {
	s, err := p.Pipeline.Struct()
	return VerifiedInfo{s}, err
}

func (p VerifiedInfo_Promise) AppId() AppId_Promise {
	return AppId_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerifiedInfo_Promise) PackageId() PackageId_Promise {
	return PackageId_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p VerifiedInfo_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p VerifiedInfo_Promise) MarketingVersion() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p VerifiedInfo_Promise) Metadata() Metadata_Promise {
	return Metadata_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

type CategoryInfo struct{ capnp.Struct }

func NewCategoryInfo(s *capnp.Segment) (CategoryInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CategoryInfo{st}, err
}

func NewRootCategoryInfo(s *capnp.Segment) (CategoryInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CategoryInfo{st}, err
}

func ReadRootCategoryInfo(msg *capnp.Message) (CategoryInfo, error) {
	root, err := msg.RootPtr()
	return CategoryInfo{root.Struct()}, err
}

func (s CategoryInfo) String() string {
	str, _ := text.Marshal(0xb9d2951d34ca391c, s.Struct)
	return str
}

func (s CategoryInfo) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s CategoryInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CategoryInfo) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CategoryInfo) SetTitle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// CategoryInfo_List is a list of CategoryInfo.
type CategoryInfo_List struct{ capnp.List }

// NewCategoryInfo creates a new list of CategoryInfo.
func NewCategoryInfo_List(s *capnp.Segment, sz int32) (CategoryInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return CategoryInfo_List{l}, err
}

func (s CategoryInfo_List) At(i int) CategoryInfo { return CategoryInfo{s.List.Struct(i)} }

func (s CategoryInfo_List) Set(i int, v CategoryInfo) error { return s.List.SetStruct(i, v.Struct) }

// CategoryInfo_Promise is a wrapper for a CategoryInfo promised by a client call.
type CategoryInfo_Promise struct{ *capnp.Pipeline }

func (p CategoryInfo_Promise) Struct() (CategoryInfo, error) {
	s, err := p.Pipeline.Struct()
	return CategoryInfo{s}, err
}

type Category uint16

// Values of Category.
const (
	Category_productivity   Category = 1
	Category_communications Category = 2
	Category_social         Category = 3
	Category_webPublishing  Category = 4
	Category_office         Category = 5
	Category_developerTools Category = 6
	Category_science        Category = 7
	Category_graphics       Category = 10
	Category_media          Category = 8
	Category_games          Category = 9
	Category_other          Category = 0
)

// String returns the enum's constant name.
func (c Category) String() string {
	switch c {
	case Category_productivity:
		return "productivity"
	case Category_communications:
		return "communications"
	case Category_social:
		return "social"
	case Category_webPublishing:
		return "webPublishing"
	case Category_office:
		return "office"
	case Category_developerTools:
		return "developerTools"
	case Category_science:
		return "science"
	case Category_graphics:
		return "graphics"
	case Category_media:
		return "media"
	case Category_games:
		return "games"
	case Category_other:
		return "other"

	default:
		return ""
	}
}

// CategoryFromString returns the enum value with a name,
// or the zero value if there's no such value.
func CategoryFromString(c string) Category {
	switch c {
	case "productivity":
		return Category_productivity
	case "communications":
		return Category_communications
	case "social":
		return Category_social
	case "webPublishing":
		return Category_webPublishing
	case "office":
		return Category_office
	case "developerTools":
		return Category_developerTools
	case "science":
		return Category_science
	case "graphics":
		return Category_graphics
	case "media":
		return Category_media
	case "games":
		return Category_games
	case "other":
		return Category_other

	default:
		return 0
	}
}

type Category_List struct{ capnp.List }

func NewCategory_List(s *capnp.Segment, sz int32) (Category_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Category_List{l.List}, err
}

func (l Category_List) At(i int) Category {
	ul := capnp.UInt16List{List: l.List}
	return Category(ul.At(i))
}

func (l Category_List) Set(i int, v Category) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type KeyFile struct{ capnp.Struct }

func NewKeyFile(s *capnp.Segment) (KeyFile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyFile{st}, err
}

func NewRootKeyFile(s *capnp.Segment) (KeyFile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyFile{st}, err
}

func ReadRootKeyFile(msg *capnp.Message) (KeyFile, error) {
	root, err := msg.RootPtr()
	return KeyFile{root.Struct()}, err
}

func (s KeyFile) String() string {
	str, _ := text.Marshal(0xe47ce2b3aab90f74, s.Struct)
	return str
}

func (s KeyFile) PublicKey() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
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
	return []byte(p.Data()), err
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
	return KeyFile_List{l}, err
}

func (s KeyFile_List) At(i int) KeyFile { return KeyFile{s.List.Struct(i)} }

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
	return Signature{st}, err
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Signature{st}, err
}

func ReadRootSignature(msg *capnp.Message) (Signature, error) {
	root, err := msg.RootPtr()
	return Signature{root.Struct()}, err
}

func (s Signature) String() string {
	str, _ := text.Marshal(0xeca8b9277cb36488, s.Struct)
	return str
}

func (s Signature) PublicKey() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
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
	return []byte(p.Data()), err
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
	return Signature_List{l}, err
}

func (s Signature_List) At(i int) Signature { return Signature{s.List.Struct(i)} }

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
	return Archive{st}, err
}

func NewRootArchive(s *capnp.Segment) (Archive, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Archive{st}, err
}

func ReadRootArchive(msg *capnp.Message) (Archive, error) {
	root, err := msg.RootPtr()
	return Archive{root.Struct()}, err
}

func (s Archive) String() string {
	str, _ := text.Marshal(0xf153ba7dee1c9118, s.Struct)
	return str
}

func (s Archive) Files() (Archive_File_List, error) {
	p, err := s.Struct.Ptr(0)
	return Archive_File_List{List: p.List()}, err
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
	return Archive_List{l}, err
}

func (s Archive_List) At(i int) Archive { return Archive{s.List.Struct(i)} }

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
	return Archive_File{st}, err
}

func NewRootArchive_File(s *capnp.Segment) (Archive_File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Archive_File{st}, err
}

func ReadRootArchive_File(msg *capnp.Message) (Archive_File, error) {
	root, err := msg.RootPtr()
	return Archive_File{root.Struct()}, err
}

func (s Archive_File) String() string {
	str, _ := text.Marshal(0xd92313d72a1ab4d0, s.Struct)
	return str
}

func (s Archive_File) Which() Archive_File_Which {
	return Archive_File_Which(s.Struct.Uint16(0))
}
func (s Archive_File) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Archive_File) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Archive_File) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
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
	return []byte(p.Data()), err
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
	return []byte(p.Data()), err
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
	return p.Text(), err
}

func (s Archive_File) HasSymlink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Archive_File) SymlinkBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
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
	return Archive_File_List{List: p.List()}, err
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
	return Archive_File_List{l}, err
}

func (s Archive_File_List) At(i int) Archive_File { return Archive_File{s.List.Struct(i)} }

func (s Archive_File_List) Set(i int, v Archive_File) error { return s.List.SetStruct(i, v.Struct) }

// Archive_File_Promise is a wrapper for a Archive_File promised by a client call.
type Archive_File_Promise struct{ *capnp.Pipeline }

func (p Archive_File_Promise) Struct() (Archive_File, error) {
	s, err := p.Pipeline.Struct()
	return Archive_File{s}, err
}

const schema_df9bc20172856a3a = "x\xda\x949\x0dp\x1c\xe5u\xdf\xb7{\xa7\x93\x8c\xe4" +
	"\xd3j\xcf\xf8gdV8\x0e\xc8\xb2$,\xc9\x09\xc4" +
	"Cz\xe8$c\xcb\x96\xf0\x9de\x05\xec\xfeeu\xb7" +
	"\xba\xfb\xe4\xbb\xbd\xf5\xde\x9dd\xb98\x07\x0c\xa4\x98\x9f" +
	"\x10(Ph`0N\x98\xe06n0\x84\x09xp" +
	"K\x98\x9ab\x08C\xc6-\x0dx\xeaI \x90N\x19" +
	"\x9c\x0e\xa9\xe9$\x19\xcc\xf6\xbdo\xf7vW\xa7\x93\xe5" +
	"\x8cg\xe5o\xdf\xbd\xef}\xef{\xff\xef\xed\xba\xde\xa6" +
	"\xeb\x84\x9e\xe0\x95_$d\xf4\x19\x1a\xac\xb3N~\xf3" +
	"\\g\xd3\x7f\x1c\xbf\x8dH\x8b\x03\xd6\x86\xc9;L\xfa" +
	"\xcaw~A\x08\x95gZN\xc9w\xb4\x84\x00\xef\x96" +
	"\x16\x91\x8e\xde\xd3\"PB\xac\xd7\xdf\xfb5\x9b\\\xf3" +
	"\x97w\x90\xc4b*x\xf8\xc1  \xca\xb7\xb6\x1c\x93" +
	"\x0f\xb4\\MH\xdf\xd1\x16\x8b\x12j\xfd\xf0O\xde8" +
	"{\xf4\x93\xff\xfck\"-\xa3\xd6%\x9b\xba\xd7n>" +
	"\xf6\xf3_\x91\xa0\x08\xc8}\xff\x1d\x11\xa8\xfci\x04\xf7" +
	"}\x12\xf9! 7\xf7}\xf7\xb6\xe9l\xe3\x9dH\xda" +
	"\xc7J\x001\x9eZrH>\xb2\x04W\x87\x97\xfc\x17" +
	"\xe0~k\xfd\xce\x9d_>w\xf2n\x92X\x0al\xfc" +
	"\xe6\xdc\xe6\x1d\x1d\x07N\xfc\x92l\xa4!\x91\x06\xfa\xf6" +
	"_\xdaB\xe5{/\xc5S\x0e\\\xaa #\xc1Su" +
	"\x13_8\x93\xb8\x17h\x07\xebg]\xf3\xd1\xa5\xa7\xe5" +
	"\xa7\x96\xe25\x0f.\x15\x11\xf5\xe5\xf7\xea\x8e\x8e~\xe3" +
	"[\x7f\x03\x02\x11fa\xde\xbf\xf4\x18`#\xd1\x87\x96" +
	"\x86\x11\x93\xbd\xb3D\xec\xea\x9fzp.\xd1\x86\x15\xcf" +
	"\xca\xd2\x0a$\xda\xb8\x82\x13}a\xcf\xff\xb5\x1e}\xe3" +
	"\xb9\x87\x91_\xea\xf1\xbbD\x0a\x81P\xe5\xf3\xcb\x7f\x0b" +
	"\x9b\xe8\x0a\x94\xc2\xe1\x98\xb8(\xf0\xbf\xdfx\x04\xa5@" +
	"}\x02\xae\xe3bXqZ>\x8at\xfb\x8e\xac\xb8\x0f" +
	"\xe9\xde\xf4\xc6\xabmu\xbb\xcf?R\x9b\xeeH\xebG" +
	"@7\xd1\x8a\x12\xfb\xdb\xd7\xd9\xd6\xa9m\xd3\x7fW\x85" +
	"9\x06\x98\x01@]\xb3\xf2\x0f\x84\xf6u\xad\xe4\xd2\xba" +
	"\xe4\x9f\x0eD\xbe\xffD\xe4\x09\x10A5\x0b}C\x97" +
	"\x81\xda\xc6.\xc3e\xe2\xb2\x1b\x11\xfb\xbe-\xbbn\xbf" +
	"\xe9\xd4\xd8\x8f\xab9\x16\x90\xe3\x9f(\x1f\xc9o*\xb8" +
	":\xa9\xe0\xf5Z\xbf\xf2\xfa\xfa\x95\x0f\x9dz\xb1\x8a4" +
	"E\x0c\xad\xed\xb4\xbc\xa7\x0dW\xb9\xb6(\xe0\xbe\xf5\xcd" +
	"\xf6\x83\xcf&\x86^%\xd2R\xea\x19\x1e\xb7\x1e\xf9\xde" +
	"\xb6?\xc8\x8fr\xdc\x87\xda\x90n1|w\xe7\xb9\x01" +
	"\xfd\x9dj\xad\xf5\xed\xbc\x1c\xf8\xd5.G~\xd5\xcb\xdf" +
	"B~\x7f\xf6\xa3\x15\x1d?\x97\xbf\xf0.I\\\x0a\xb6" +
	"\xb3\xec\xfe\xd6\xdf\xec?6\xfa\x09\xd9(\x84P\x10\xc7" +
	"\xaf8-\x9f\xbc\x02\xd1O\\a_\xef\xfec\x8f\x1f" +
	"\xcf\x1c\x7f\x17\xb8\x08x\\\x00\xe5\x9e\xf6\x18\x95\xfb\xdb" +
	"Q\xcb\xd7\xb6\x83\x87ln\xe7\x1e\xf2\xf0\xcb\x99\x9ek" +
	"n\xbb\xe7LMq|\xa5\xfd\xb4\xbc\x11\xb7\xc0Fd" +
	"\xdb\xd5C\x15\xb2\x84(\x1f\xb6\x1f\x93\xcf\xb6_\x09\xac" +
	"\x04\xd7\xe4\x05\xc0\xde}\xda\xba\xfe\xdf\x06\x7f\xf5~\x95" +
	"\xf3q\x0f\xd1\xd6\xbe\"\xe7\xd6\xe2\x8a\xad\x9d&>o" +
	"\xab!\xe8\x13\x80\xfb\xe6\xda\xa5\xb0z{m\x94\x0b\xef" +
	"\xc5\xbf\x7f\xee\xfd\x9b?\xa8\xc2\xe5\x1c7t\x82)w" +
	"\xe2\xaa\xa9\x93\xd3=R\xdf\xf2\xc8\xa1\x07>\xa8\xb2#" +
	"\x90\x9e\x088{:\xdf\x97\xf7s\xec\x99N\xb8\xdf\xe7" +
	"\xbb\xae~\xf8\x81\xef\x9c\xf8\x90\xe3z\x1a\xe4\xc1\xe2\xd3" +
	"\xce\xdf\xca\xb4\x0b\x05}\xbe\x93[\xdd\x9d\xa9\xe7n\xbe" +
	"\xf2\xc5\xa7?\xae\xc5\xc5\xef\xbb^\x91i7\xae\xcew" +
	"!\x17\xae\xd6j\xdcN\xed~Vf\xddx\xbb=\xdd" +
	"x\xbbk\xc2\xe3\xff\xf8\xfd\xd6\xc7~O\x12\x97Q\x97" +
	"\xa3\xb1`\x88\x0ap\xf6\x93\xdd\xab\xe0l\xf9p7R" +
	"}\xe2\x86\xaf~y\xcf\xbe_\x9c'\x89V\xea\xbb\xe9" +
	"\x12\x01\xbd\xa9o\xf9U\x02\xa2\xae\xbcj\x9aX\xfc\xdf" +
	"\x9fZ\x86\x9a\xdc\xad\xa6\xb5n!\xa9\x1a\xba\xb1!\xa7" +
	"\xa6Y\xf2\x86Rn\\\xa3f\x9cR\xdaD\x04x\x88" +
	"Dc\xd6}\xaf\xbe\xf9?\x1bW|\xf4\xb0P\xd9b" +
	"\xef\x18Q\xa3:\x9b\xd0\x0a\xc5D#\xf5\x9b\x9at\x97" +
	"\xcf\xfa\x97\xc4*\\K\xd2\x06\xab\xc0\xf6i\xc3,\xc7" +
	"hqH\xbf1o\xa6\x0a\x84\x94\x07\xf2\xb9\x9c\xaa\xa7" +
	"\xa2\xfd\xc9\"\xcb\xeb\x89\xd5b\x80\x06\x02\xc0\xb3t\xd6" +
	"$$\xf1\xb1H\x13\xbf\x13\xa8\x14\xa4\x11\x8a\xc0O\x11" +
	"xN\xa4\xdb\xa9@i(\x82r\x90\xce\xc7\x00\xf6;" +
	"\xb0bx\xa5R=`\xa2:%z\x1bXw3\x05" +
	"x+\xc2\xa9\x10\xe1\xd1b9\xdd\x05\xf0e\x08_\x8d" +
	"pA\x8c\xd0 \xc0/\xa7\xaf\x00|5\xc2\xd7!\\" +
	"\x04\xfc:\x80w\xd1C\x00_\x87\xf0k\x01\xde\x0c\xe8" +
	"\xdc\x1b\xe8\x16\x00_\x83\xe0AD\xaf\x0bDh=\xfa" +
	"\x06\x87_\x87\xf0a\x80[9\xa6\xf7\x1b\xeck\x1aQ" +
	"\xcc\x02\xdc\x0fp\x04x\x00\xae\xee\xad\x05/\xab\\\x0c" +
	"\x05\xba\x98\xd0\xb8HisE|\x84\"\xc8J\xe6\xf5" +
	"\"\xd3K\x1au\xc4Fh\xb3'm\xc0i\x06\x1c\xd5" +
	"0\xbe\xa6\x01U\"\xfa\xcfc\xfa\x98\x916U!\xa5" +
	"\x8eg\xb5~\x07%\xaf\x13\x17\x03\xb6\x8d\xa8\xe6n\x8d" +
	"\xc2\x01i\xfc5\x84|5[C\xef\xff\xd5\xd5\xe6\x8f" +
	"\xff\xfcn\x1f\xf9\x1d\xac\x98\xd5\x08!5~\xcdiE" +
	"5\xa5\x16U\xfbW\xd7\xc7\x9c_+\x16$\xda&4" +
	"\x9a/\x99ImD5\xba\xe11\xe0X\x02\xa6\x97h" +
	"\x14AO\xdc\x046\x8e\x83f\x07\xc1\x04\xe2\xa8@\xc7" +
	"\x04F@\x7f\x89a\x00\xde\x84\xda\x13l\x1b\x18\xdb\x0e" +
	"\xc0\x1d\x004\x04\xf7\x948\x09\xa9\xc5\x0cm\x84\xfb5" +
	"\xc2\xd9\x05~X\\%\xa2\x0f\x98a)\x00\x153\x84" +
	"\xba\"\xc7\x9f\x16\xfbx\xa56\xaf\xfdF\xc8\x18J!" +
	"\x7f\xcd.\x7f\xea*8\xf5\xcf\xe0\xd4\x8c\x8f?\x0d\x81" +
	"_\x07`\xd6\xc7\x1fC`\xca\xe6O\x12En\xa2R" +
	"\x0e\x81\x19\x00\x16\x05\x1ab\xa9u\xb4\x01\x8en \xb8" +
	"\xee\xf1\xad{}\xeb\xbe\xca\xbaZ\x94#\x8e\xdc\xbbG" +
	"\x93\xa6\xa6\xe9\x85L\x9e\x16\x91\xd9\x88\xcb\xec\xfe^8" +
	"m/\x9cv\xbb\x8f\xd9[7\x00\xf0f\x00\xde\x09\xee" +
	".X\x96\xcd\xed\x1d\xe8Q\xb7\x00\xf4\x1e\x80\x8a\x9f[" +
	"6\xbb\x07\x90\xdd\xdb\x01\xfam\x80\x06\xce[\xdc\x9f\xa4" +
	"{;\x00z'@\x1f\x14\xa82\xcdR \\\xc7\xa2" +
	"\xa2\x19\x8d\xa53E\xd7\xb4K\xfan=?\xad\x93\xba" +
	"\x90\xa1\xa7\x9d\x00C\xc3\x93\x86\xe6\xbeT\x07\xa5\xa4Z" +
	"\xd4\xd2ysfH\x17'\xf2\x18\x95\x9a\xbd$<W" +
	"C\x03j\xd4F\xc7{w\x82\xf3-\x82;K\xfb\x15" +
	"`\xf2\xf1I\xe9I\x85R\xe9\xc4>\xe9\xa4B\x05\xe9" +
	"\xec\x06\xe9\xacBE\xb9\x89\x9a\x10(\x14\x1a\x90\xbfD" +
	"7\xc0\xa3\xd0\xa0<F\xf7\xc9;aU'\xef\xa71" +
	"x\x14Z/?N{\xe1Qh\x83|\x1cV\xc7a" +
	"\x15\x92\xcf\xd0-\xf2{TQ\xf2\xc5\x8cf&\x02\x90" +
	"\xd5\xbcJ- \xd0\xfe\x08\x9a<\xb9\x0ed\xdd\xabl" +
	"C\x1cH\xb0\x86\x99O\x95\xc0\xc3Ix\x8a\x15gj" +
	"lZ\xe6n\x9a\xb4\xe2\xb3\x90aw\x12|\xbe\xa4\xb3" +
	"$\x89\xaa<H\\p\xff>k\xa0\x0a\x9d\x90h!" +
	"\x9fdj\xf6\x82\xccn\x88\x8er$8oZ\x1b\x8f" +
	"\x97\xc6\xb3\x8c(\x85\x0c8\xe7\x02\xc7\xdd\xa8\x8d\xb7!" +
	":\x892\x8e\x0e\xc7\xe5'&XR[\xe0\xb8m\x1c" +
	"\x09\x8eKiSZ6oh$j\xee\xc8\xe7\xb3\x17" +
	"\xbe\xde\x16kP\x9b\xe2h`\x09\xe5B\x92i\xfa\x02" +
	"'\xc5\xca\xa36\x16!JNK1u\x01\x9d\x8d " +
	"\x0e\xe0\xa6\xd5\x9cVX\x00w\x13\xe2\xc0\x15 \xc0\x1a" +
	"\x19\x96\x84%Y\x80\xfbM.\xe6\x1c;\xce\x17\x146" +
	"\xa4W\x0c\xde-Q\xe7\xba\x87\xeb\xf3\xaaR*f\xf2" +
	"&FN\x89\xd6G\xe8%\x18;'}\xb1Sh\x88" +
	"@\\\x83\xd89\xe9\xc5\xce\xe6\xe6\x08\x951t\xees" +
	"B\xe7\xd7\x05;\xbf@\x0a\xdaH\xc29\x95e\xdd0" +
	"i\xa4\x8dQ\x96\xd6U\x12.\x96L\xcdu\xd7\x92Q" +
	"(\x9a\x9a\x9a#\xd1~\xce\x80\x87?\x9bQ\xc8&l" +
	"\x82i\xa9\x8a\x1f'Z\xdd\xa8\xf4<F\xa5g\xe0\xf4" +
	"\x97|Q\xe9E\x8c\xe6/\x00\xf0_|!\xf4'\x88" +
	"\xf9\x12\x00_\x83\xd4\xefD\xd0\x13\x18\xa8^\x06\xd8O" +
	"\x011 \xda\x11\xe9\xe4]\x00\xfc)\x00\xdf\xc1\xc2!" +
	"\xc0\xd3\xbb\xf46d\xf7\xc4;\x00\xfc\x00\x93u\x90\xe7" +
	"v\xe9=\xc8\xd5\x89_\x02\xf0c\x88]\x90\x81\x86R" +
	" p\xb7\x97\x9b\x9d\xb3\x86\x08\xc5_\xdd:\xd6\xfeU" +
	")b*\x9c\x9b\x07\xcbSZu\xbe\x87\xd4\x0a\x99\x95" +
	"\xa6+\x89\xb7V\xf6T\xb9\x18\xe3i\xc1\xd8\xaa\xcd\\" +
	"\x0fN\xa4\x99\x86\xc9\xf4\"q%{q\xf9\xb5\xda@" +
	"X8\x89\xf1\xc2\x97\xb9b\xb52Wo\xad\xcc\xb5\xc1" +
	"\x9f\xb9\xc2\x11\xda\x82\x99\x0b\x15\x94\x05\xe0^\x01J\x16" +
	"\xc3\xd8d2\x14\x8e[v:\xc2\x01w`\xfa\\x" +
	"\xd4\x16\xc6\xdc\x1f\x1c)\xc5\x18\xa1\xe9\x1a\xbf\xcew\xbb" +
	"l\x14\x02\x88^\xd00\xd95b\x0eC\x85\xef\xef\xf0" +
	"\xd2\xddJ\xfa\xb9Em\x95\xdf\xba\xcbKm+\x85\xf3" +
	"V\x90\x97s\xd2\x81q/\x8b\xad\x14?C0Ts" +
	"\xd2\xfd\xe8/\xdf\x06\xf0ch_u\x11\xc8\xbcDz" +
	"\x14e\xf7 \x00\x0f\x0a4\xac\xe7u\x8d\xd4Y\x10\xb6" +
	"t\xacf\x88\x98\xd4h\xd8k\xe9\x80\xf30rn\xe6" +
	"A\x8fZ\x11*\x12s\xa6\x86\xde\x0d\x0c\xb3\xc9\xc1<" +
	"w\xbb\x1aEWY\xcf\x17\xe1\x8e\x85Z[g\xd7\x01" +
	"q\xfbuP\x9b`:\xc3\xb0\xcfK\xaae\xae\xe2\x1f" +
	"]\xe11\xef*\xfeq\xf4\x83\xc7\x00\xf8\xb4O\xf1O" +
	"\xa1\x8e\xbf\x07\xc0g|%\xcb\x11\xc4\xfc\x01\x00_@" +
	"\x81\x04l\x87{\x1e\x8b\xf2\x1f\x01\xf0g\xe8pA[" +
	"\xfeoNz^(\x82u\xb8\x06\xac\xda]\x83m\xc0" +
	"\xbe\xa6\x94\xdf\xa6\xe0T\x84\x84\x1ah\x00\x95F\xd0\xf9" +
	"u\x82e\xa1\x7f\xb0\xf7V\xe8\xa9\xd9iu\xa60\xa4" +
	"\x13%\x99-\xa5\xb4\xea*n\x1c,3\xad\x0d\x80d" +
	"\xf5\x09\x86F\xe5\xf6\xb8\xb5\x8dj[\x81\x0d\xdb\xd6\x14" +
	"vb\xb0\xbf\x1eE\xe1]\x07w\x1a\xf6\x09o\xa8\xd7" +
	"\x0b\xb4\xd8d\xc0_i\xc4\xf4\xe2\xac\xef\xf2N\xb4\xa8" +
	"\xb0nj{J\xcc\xd4F\x89\xc2o\x0d\xd4`\xef\\" +
	"\x96\x06\xaa\xaa\xa0D\xc0e\xa8\x09\xcf\xae\x87s\"B" +
	"5\xf1jgq\xc4\xde=\x10\xb5;\x07\xe7fT\xb0" +
	"o\xf6\xba/18\x17\x1bC\x0f\x8a;\xe1\x80V\xc2" +
	"A\xcc\x09\x07\xb7@\xb6Hi\x86\xa9A\x91&h\xa9" +
	"\x8d{\xb5d\xa9\x88\xbd\x05/\xa7]\xfd\x84U3=" +
	"U\xa5\x94\xb2\xa6O1\x13\x82c\xa5\xc9\xb1V\xaf?" +
	"\xfa\x0f\x1f\xdcz\xf0A\xe2\xb49UF\xbd\xcd\xf1\xae" +
	"\xa4\xe6\xe8\x86\x1b\xf50g\xe8L\x0c7Io\xaf\x92" +
	"\xde\x86\xcaN\x0eB\xcd\x16\xa4X\xdb}\x95v\xc0\x83" +
	"\xb5]\x0ej\xb6\x1c\xaf\xed\x9e\xa4\xbb\xe4\xa7xm\xf7" +
	".\xac\xce\xf0:N\x12:\xe0\xc1*\xaf_\xe8\x85\x07" +
	"\xab\xbc\x12\xacJ\xb0Z$\x1f\x16V\xc1\xa3\xd0K\xe4" +
	"w\x85\xed\xf2\x19X5\xcaK\xc4\x18<\x0am\x92G" +
	"\xc4\x0ex\x14\xbaX\xde\x0f\xab\xfd\xb0\x0a\xcb\xcf\xc3\xea" +
	"yX5\xcb\x1f\x8a\xab\xe0Q\xa8$/\x0ft\xc0\xa3" +
	"\xd0\x16y'\xacv\x06\x942\xd3\xa7\xd4,K\x85r" +
	"\xac\xc8+\x05o \x87\x95\xc2j^)\xc4\xf9\x05\x83" +
	"\xab\xa4\xe0xhdh\x87\x05O\x1b\x0a\x80\x84P\x04" +
	"\x04\x02\xaf\x9a\xcch\xbd5\xf6w\xfa\xf7\xef\x92\x1a\x1e" +
	"\xb0\xfa9nW/\x11\xbb\xd79/m\xc3\xd4\x96f" +
	"\xdb\x14\x05;\x0a\xa7\x8dl\xdf\x85x\x01)\x07c@" +
	"\xae\xbc)>\xdc\xd5\x07d6\xdd0\xd6\x06\xeb6\"" +
	"N\xf5A\xa1\xa4\xceC\xa0\xddO`\x8b\xd40n\xf5" +
	";\x14\xc0\x8f9\x11|'\xa16\xa4b\x8d\x17R}" +
	"\x03Y\xb5D\xc4\x82v!bx\xb3I\xa9a\xd2\x8a" +
	"\x8d\x0ev\xf5u\x0ddIX-\x154|m\xf3^" +
	"m\x82\xbd\x7f4\xc1\xde\xd9\x04{}\x04QN\xbd\x17" +
	"-\xa7\xde\xd9r\x029+\xd9y\x08T\xcb\xc9\xb4\x86" +
	"m\x0a=\x159\xe1;Q\xda\xa68\x88\xd3\xb98y" +
	"\x0fW\xc9{\xd8\x93w\x88\x15\x92\x17e\x80C\xa3\x03" +
	"\x16<~\x03\xb4T\xb3\x08\xd1\x18:\x0c\xba\xa0\x11\x82" +
	"d\x0fY\xfd\x15\xfc0\x97K\xe5\x95\xb69n\x0d," +
	"\x81\x80\xca\xc6\x0cTA\xfa\xc5\xd9u\x9c\xe3:vm" +
	"\xbfT\xdb\xb5\x911\xfa\x16\xba\"\xe8k_9\xbe9" +
	"\xce\xed\x1a\xfe\xb7\xaf\x19\xe5D@J\xe1\\m\x95u" +
	"V\xe9\xfcTy\xc4\xd1\xf9H~\x1f\xcbf\xd56\x81" +
	"wP\xc9\xca\x0d\x1d\x96\x92\xa9Tv!\x13\x02\xcdu" +
	"X\x03\x83\x83\xc3]=\\sa\\\x83\xc24#\xbb" +
	"\x90\xd2\x81\x93c\xe5\x8dq\xbe\xd3\xda\x98\xcc2\x03N" +
	"\xa6\x15N\x9cj\x09\xd90\xd4Z\xb4\xd6\xcfb\xa3\xa7" +
	"\xa1\x85Z\x03\xf1\xfe\x0a\x1f\xbc\xaf\x04)\xc7E\x9b^" +
	"\x7f\xb1h\xb2\xf1\x12\xd6\x1a\xce-)\xd2\xde\x97e\xe3" +
	"\x0b\xb9[\x87\x14<\x14\xde\x05\x88\x16\xfe\xb9\x0a\x1ej" +
	"\xe8iN\xc5\xb1\xb0\xaa,\xd6o&3lJ\xeb\xbe" +
	"\x9e\x89\x90hf\x176\x1d^a\xb3R\xb0\xacJi" +
	"\x83\xf9\xea\x11\x00\x7f\x0f\x0b;^\x06\xe2\xd1Ob\x19" +
	"x\x10\xc0?\x00p\xe0\xbce\x0f\x0d\xa5\xc31\xaf\xe6" +
	"Y\x19\xfc\x0c\xc1X\xe0\x1c\xd9\xeeT=\xaf9y\x1f" +
	"\x0b\x9c\x13\xc7\x00\xf8\x1a\x00\xff\x1d\xcb@h\x00+)" +
	"\xaflj\xe9RV5\xdd6Isr#\x01\xa6+" +
	"\xc0ra&\x97e\xfan7_\xa7\xa0\x12H\x16\xf3" +
	"&\xa13^^t\xe7\xfdN^\xcc\xaa\x85\xe2H>" +
	"\xc5(\xb4\xcc\xbc\xa9\x8f\xee`9\xed\x86\x020$\xc0" +
	"\xe3%\xce@U\xd6\xafL^\xdd\xc1+6\x97|\x94" +
	"\x14h\x9e+\xe6\x98S4\xe9\x13\"KW\x95@[" +
	"j\x8d\xe4Pl\x9b\x01\xb8\xc3)\x81\x00\x96\xb8\xcb\xd7" +
	"VN1m\x1a\xcb)\xbb\xe8\xd3~\xfdt\xff\xa6/" +
	"\xfd\xc5K\x95\xe2V5X|\xd6\xa4N\x9d\xd2\x86R" +
	"\x9aN\xa1\xa2\x99\x19P\x0dl\x92\xabk#Z\xe9\x01" +
	"\xa2v\x13\x90\xa8\xf7\x8f\xbd\xd1u\xdc\xefl\x10)\xc2" +
	"C\xd0\xffX\x95\xd1\x18\x11\xf3\xc5D'^\x89\xfa>" +
	"C\xc9_\x04\xd7\x14$\xa7\xb5\x94\x9b(\\j\xb4\x9e" +
	"V\xe6\xca\x82]\xec\xca\x12\x877\"|\x19z\x08\xe8" +
	"\xc9\xfb@%/\x81_\x05)\x10\x8a\xd0E8\x99F" +
	"+\xfb\xac\x82\x19\xa4\xbe\x8fi\x80\xb9\x010\xeb\x16E" +
	"\xf8\x98\xdd\xc5\xac\xc7\xb3B\x97D@\xdd\x04\x0a\x19h" +
	"KF\x03xV3\x9fm7F\xa0\x9d@\xde\xee\x9a" +
	"5\xdbnh\x8a\xa0W\xc9\xcb9>\x9fm\xafG\xf8" +
	"\xa2\xc5\x11*\x01\xbc\x87n\xf7\xcf\xb0\x15\x86\xdd`y" +
	"Z\x1b/\xb0\xa2g\xb6\xc9|J\x1b3\xdd\xa6\xbf\x9c" +
	"\xb5\xe3\x84\xe5L\xe1\x18\x115wH\x1a\xf6\xbe5\xda" +
	"\xa6\x19\xb5[W\x9c\x14@\xeb\x8a\xc8\xde\x8c\x0f\xea\xc4" +
	"B\xd2d\x06\xb4?\xb5G\xca\x05\xd8Y\x1c\x04\xea\x1c" +
	"k\xbe\xe6\xb8PQ_(_\xf4\x06\xe4\x9e\x9e+#" +
	"\xf2\x8c\x0a}\xf3p\xde\xee!\xe7k\x97\x84Y\xed\xd2" +
	"P\x8aW\x94\xf5\xae\x99\xaf\xc1Y\xe7j\xb0\xdeu>" +
	"3\xefB`;\x00\xd7\xcf?\xaf\xad>\xc0\x1dq\xe3" +
	"@\xc8\xffU\x99\xc6\xca\xce\xd0\xdb_\xcf\xa3\x194\xc2" +
	"\x09\xed\xe04\x05M\x85p\xe7\x0c\xac\xdd\xeb\xba\x14\xaa" +
	"Je\xc7!\xb6j\xca\xcc\xf5\xcc\x0e\x8e\xbe\xebl\xf7" +
	"8w\xaf\xd3\x83\x87\xad\x03\xe0\xb5B\xa5\x03\xdd\xaaa" +
	"\xf8q\xa7\xb1&\x9b\x02\xddo\x05\xc5\xcf\xcc7\xa2u" +
	"\x9b\xf0\xa1\x10\x18\x95\x13+\xa0\x09\xb7\x83E\xcc\xeb\x97" +
	"\x9a(\xce\x91y\xc3\xb4\xca\x0b!M\xd0\x82S\xdf\xd7" +
	"+id\x15\x11|\x93\xe2\xc2T\xbab\x8f85\x9e" +
	"\xb7\xab\xe9W\xf8w\x13w\xe2M}\xdf\xcep\xf2-" +
	"\xf0[c\\\xd2bUC\x0e\x8c\xfb\xac\xd77\xe4\xa8" +
	"\xcc\x91r\xe3\xde\x90C\x12\x9d9R\x09\x85V\xb4\xfb" +
	"\x1f\x85\xe9F\xa9XN\xda\x0d\xd5\xdc/1\xf3\x0d\x87" +
	"\x16\xf4\x07=_\xd2\xe3\x19S\xc5\x92ua\x1b\xe6\x93" +
	"9\x9c\xcb\x91\x8bQ\xfa\xf6\x05\x94^p\xa8\x11\xaa\xcd" +
	"\xd1y\xe5S\x88\xa9\xf04\xcc-\xdaMQ\x12\xed\x08" +
	"\xa3\xe5\xd5jOW\x83\xb0\xb0y/\xcc\x9f\xda\xaa\xbf" +
	"g\xb8\x9a\xe5\x8a\xedfz\x08d\x8d\xb7\xab\x18\xd7\x9a" +
	"\x0e\xcfE=\xe3\xeaB\xfdt\x02t\xb37\x9c\x01\x8a" +
	"\xea8\xcbBL*\xfar\xebM\xad\xff\xfa\xcf\xdf\x0d" +
	"\xbfuv>\x06\\\xc3\x06\xbb\xee\x06\xdb#\xb6hm" +
	"+Z\xd3[\x15\x1f\xd0\x8a\xbaz=\x81+)\x83\xf5" +
	"\xec\xadH\x10\xdfz\xdd\xb7\xff\x0f\x00\x00\xff\xff\xd4\xfd" +
	"\xb9\x17"

func init() {
	schemas.Register(schema_df9bc20172856a3a,
		0x82bdd60d2cf486c9,
		0x855f296a69e6e1ca,
		0x87dcf1b1edcb3eaf,
		0x880c6c7782a33310,
		0x8bc9f4365959348e,
		0x8d51dd236606d205,
		0x938e7e53b106e1c0,
		0x9476412d0315d869,
		0x96b3cbb11cf671b8,
		0x987ef3040a0342a9,
		0x98fd6b0620c6cb58,
		0x9a774f764b69ca97,
		0x9f149fa71489be0b,
		0xb755d258845a4a8f,
		0xb9d2951d34ca391c,
		0xc64951b2a02886cf,
		0xd86e43f42c8b0f74,
		0xd92313d72a1ab4d0,
		0xd9bd68bd9dba918f,
		0xdd8c82383168c096,
		0xe0c5892a5448f4ee,
		0xe2e344d346ffda6b,
		0xe3d7ba482b2e470b,
		0xe47ce2b3aab90f74,
		0xe492a2981208ad0b,
		0xe5c59b9296375a00,
		0xeca8b9277cb36488,
		0xf153ba7dee1c9118,
		0xf99c1ca7ae620f38,
		0xfddf7a71363d4e9f)
}
