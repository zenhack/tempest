package spk

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	apisession "zenhack.net/go/sandstorm/capnp/apisession"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	powerbox "zenhack.net/go/sandstorm/capnp/powerbox"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

const OsiInfo = uint64(0x9476412d0315d869)
const CategoryInfoAnnotation = uint64(0x8d51dd236606d205)

// Constants defined in package.capnp.
const (
	Manifest_sizeLimitInWords = uint64(1048576)
)

// Constants defined in package.capnp.
var (
	MagicNumber = []byte{143, 198, 205, 239, 69, 26, 234, 150}
)

type PackageDefinition struct{ capnp.Struct }

// PackageDefinition_TypeID is the unique identifier for the type PackageDefinition.
const PackageDefinition_TypeID = 0x9f149fa71489be0b

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

// Manifest_TypeID is the unique identifier for the type Manifest.
const Manifest_TypeID = 0x855f296a69e6e1ca

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

// Manifest_Command_TypeID is the unique identifier for the type Manifest_Command.
const Manifest_Command_TypeID = 0xc64951b2a02886cf

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

// Manifest_Action_TypeID is the unique identifier for the type Manifest_Action.
const Manifest_Action_TypeID = 0xe5c59b9296375a00

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

// SourceMap_TypeID is the unique identifier for the type SourceMap.
const SourceMap_TypeID = 0xe3d7ba482b2e470b

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

// SourceMap_Mapping_TypeID is the unique identifier for the type SourceMap_Mapping.
const SourceMap_Mapping_TypeID = 0x87dcf1b1edcb3eaf

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

// BridgeConfig_TypeID is the unique identifier for the type BridgeConfig.
const BridgeConfig_TypeID = 0xdd8c82383168c096

func NewBridgeConfig(s *capnp.Segment) (BridgeConfig, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return BridgeConfig{st}, err
}

func NewRootBridgeConfig(s *capnp.Segment) (BridgeConfig, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
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

func (s BridgeConfig) PowerboxApis() (BridgeConfig_PowerboxApi_List, error) {
	p, err := s.Struct.Ptr(2)
	return BridgeConfig_PowerboxApi_List{List: p.List()}, err
}

func (s BridgeConfig) HasPowerboxApis() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s BridgeConfig) SetPowerboxApis(v BridgeConfig_PowerboxApi_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewPowerboxApis sets the powerboxApis field to a newly
// allocated BridgeConfig_PowerboxApi_List, preferring placement in s's segment.
func (s BridgeConfig) NewPowerboxApis(n int32) (BridgeConfig_PowerboxApi_List, error) {
	l, err := NewBridgeConfig_PowerboxApi_List(s.Struct.Segment(), n)
	if err != nil {
		return BridgeConfig_PowerboxApi_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// BridgeConfig_List is a list of BridgeConfig.
type BridgeConfig_List struct{ capnp.List }

// NewBridgeConfig creates a new list of BridgeConfig.
func NewBridgeConfig_List(s *capnp.Segment, sz int32) (BridgeConfig_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
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

type BridgeConfig_PowerboxApi struct{ capnp.Struct }

// BridgeConfig_PowerboxApi_TypeID is the unique identifier for the type BridgeConfig_PowerboxApi.
const BridgeConfig_PowerboxApi_TypeID = 0xc9702c7dbfc6d7e4

func NewBridgeConfig_PowerboxApi(s *capnp.Segment) (BridgeConfig_PowerboxApi, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	return BridgeConfig_PowerboxApi{st}, err
}

func NewRootBridgeConfig_PowerboxApi(s *capnp.Segment) (BridgeConfig_PowerboxApi, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	return BridgeConfig_PowerboxApi{st}, err
}

func ReadRootBridgeConfig_PowerboxApi(msg *capnp.Message) (BridgeConfig_PowerboxApi, error) {
	root, err := msg.RootPtr()
	return BridgeConfig_PowerboxApi{root.Struct()}, err
}

func (s BridgeConfig_PowerboxApi) String() string {
	str, _ := text.Marshal(0xc9702c7dbfc6d7e4, s.Struct)
	return str
}

func (s BridgeConfig_PowerboxApi) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s BridgeConfig_PowerboxApi) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BridgeConfig_PowerboxApi) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s BridgeConfig_PowerboxApi) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s BridgeConfig_PowerboxApi) DisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(1)
	return powerbox.PowerboxDisplayInfo{Struct: p.Struct()}, err
}

func (s BridgeConfig_PowerboxApi) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BridgeConfig_PowerboxApi) SetDisplayInfo(v powerbox.PowerboxDisplayInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated powerbox.PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s BridgeConfig_PowerboxApi) NewDisplayInfo() (powerbox.PowerboxDisplayInfo, error) {
	ss, err := powerbox.NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return powerbox.PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s BridgeConfig_PowerboxApi) Path() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s BridgeConfig_PowerboxApi) HasPath() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s BridgeConfig_PowerboxApi) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s BridgeConfig_PowerboxApi) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s BridgeConfig_PowerboxApi) Tag() (apisession.ApiSession_PowerboxTag, error) {
	p, err := s.Struct.Ptr(3)
	return apisession.ApiSession_PowerboxTag{Struct: p.Struct()}, err
}

func (s BridgeConfig_PowerboxApi) HasTag() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s BridgeConfig_PowerboxApi) SetTag(v apisession.ApiSession_PowerboxTag) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewTag sets the tag field to a newly
// allocated apisession.ApiSession_PowerboxTag struct, preferring placement in s's segment.
func (s BridgeConfig_PowerboxApi) NewTag() (apisession.ApiSession_PowerboxTag, error) {
	ss, err := apisession.NewApiSession_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return apisession.ApiSession_PowerboxTag{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s BridgeConfig_PowerboxApi) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(4)
	return capnp.BitList{List: p.List()}, err
}

func (s BridgeConfig_PowerboxApi) HasPermissions() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s BridgeConfig_PowerboxApi) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s BridgeConfig_PowerboxApi) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

// BridgeConfig_PowerboxApi_List is a list of BridgeConfig_PowerboxApi.
type BridgeConfig_PowerboxApi_List struct{ capnp.List }

// NewBridgeConfig_PowerboxApi creates a new list of BridgeConfig_PowerboxApi.
func NewBridgeConfig_PowerboxApi_List(s *capnp.Segment, sz int32) (BridgeConfig_PowerboxApi_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5}, sz)
	return BridgeConfig_PowerboxApi_List{l}, err
}

func (s BridgeConfig_PowerboxApi_List) At(i int) BridgeConfig_PowerboxApi {
	return BridgeConfig_PowerboxApi{s.List.Struct(i)}
}

func (s BridgeConfig_PowerboxApi_List) Set(i int, v BridgeConfig_PowerboxApi) error {
	return s.List.SetStruct(i, v.Struct)
}

// BridgeConfig_PowerboxApi_Promise is a wrapper for a BridgeConfig_PowerboxApi promised by a client call.
type BridgeConfig_PowerboxApi_Promise struct{ *capnp.Pipeline }

func (p BridgeConfig_PowerboxApi_Promise) Struct() (BridgeConfig_PowerboxApi, error) {
	s, err := p.Pipeline.Struct()
	return BridgeConfig_PowerboxApi{s}, err
}

func (p BridgeConfig_PowerboxApi_Promise) DisplayInfo() powerbox.PowerboxDisplayInfo_Promise {
	return powerbox.PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p BridgeConfig_PowerboxApi_Promise) Tag() apisession.ApiSession_PowerboxTag_Promise {
	return apisession.ApiSession_PowerboxTag_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
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

// Metadata_TypeID is the unique identifier for the type Metadata.
const Metadata_TypeID = 0xe0c5892a5448f4ee

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

// Metadata_Icon_TypeID is the unique identifier for the type Metadata_Icon.
const Metadata_Icon_TypeID = 0xe492a2981208ad0b

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

// Metadata_Screenshot_TypeID is the unique identifier for the type Metadata_Screenshot.
const Metadata_Screenshot_TypeID = 0x8bc9f4365959348e

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

// OsiLicenseInfo_TypeID is the unique identifier for the type OsiLicenseInfo.
const OsiLicenseInfo_TypeID = 0xb755d258845a4a8f

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

// AppId_TypeID is the unique identifier for the type AppId.
const AppId_TypeID = 0x880c6c7782a33310

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

// PackageId_TypeID is the unique identifier for the type PackageId.
const PackageId_TypeID = 0xe2e344d346ffda6b

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

// VerifiedInfo_TypeID is the unique identifier for the type VerifiedInfo.
const VerifiedInfo_TypeID = 0x987ef3040a0342a9

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

// CategoryInfo_TypeID is the unique identifier for the type CategoryInfo.
const CategoryInfo_TypeID = 0xb9d2951d34ca391c

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

// KeyFile_TypeID is the unique identifier for the type KeyFile.
const KeyFile_TypeID = 0xe47ce2b3aab90f74

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

// Signature_TypeID is the unique identifier for the type Signature.
const Signature_TypeID = 0xeca8b9277cb36488

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

// Archive_TypeID is the unique identifier for the type Archive.
const Archive_TypeID = 0xf153ba7dee1c9118

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

// Archive_File_TypeID is the unique identifier for the type Archive_File.
const Archive_File_TypeID = 0xd92313d72a1ab4d0

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

const schema_df9bc20172856a3a = "x\xda\x94Z\x0fp\x1c\xe5u\x7fo\xf7N'\x19\xc9" +
	"\xa7\xf5\x9e\x8b\x9d\xda]!\xc4\x1f\x0bKX'c\xc7" +
	"\xc2\xce\xa1\x93\x85\x91-\xc1\xad\xcf\x06\xec)-{w" +
	"\xab\xbbO\xbe\xdb]\xf6\xf6$\xcbEQ` \xe5O" +
	"\x08\x81PJ\x1a\x18\xc0\x09\x13hC\x89!\x9e`\x0f" +
	"N\x88'\xa6\x18\x92\x81\xa1\xa1\x80\xa7\x9e\xc4\xfcIg" +
	"\x18L\x87\x94\xb4i\x06\xb3\x9d\xb7\xbb\xb7\xb7:I\xc8" +
	"\x1d\xcfxv\xbf{\xfb\xbe\xf7\xff\xfd\xde\xf7i\xcd\xeb" +
	"\xadWp=\xe1\xc4W\x00\xd2\xbf\xe6\xc2\x0d\xf6\xf1\xaf" +
	"\x7f\xba\xba\xe5\xdf\x8e\xdc\x0a\xc2\xe2\x90\xdd7v\xbb\x89" +
	"G\xbf\xfb\x1b\x00\x14\xc3\xd2\x1b\xa2 E\x00\xd2\xcd\x12" +
	"\x8f\xe9e\x12\x87\x00\xf6+\xa7~\xc7\xc6V\xfd\xf5\xed" +
	" /F\xaeF\x1f\x0eG\x00\xc4\x16\xe9\xb0\xb8TZ" +
	"\x0f\xd0\xbbI\xb2\x11\xd0~\xe6+\xaf\x9e>\xf0\xc9\xbf" +
	"\xff-\x08\xcb\xd0>gK\xf7%W\x1d~\xeb=\x08" +
	"\xf3\x11\x80\xde\xcay\x1c\x8a\xb7\x9cG\xdfM\x9d\xf7\x0c" +
	"\xa0\xdd\xda\xfb\xbd['\x8a\xcdw\x10\xeb\x80(!\xa2" +
	"\xe8i\xdf/nh\xa7\xa7\xcb\xda\xff\x03\xd0\xfe\xe6\xda" +
	"]\xbb\xd6}z\xfcn\x90\xcfE\xce\xfe\xf8\xd3\xabv" +
	"t\xdey\xec\xb70\x88\x11\x1eC\xbdM\xe7/Aq" +
	"\xf9\xf9\xb4\xcb\xd2\xf3%\x12$\xfcF\xc3\xe8\xf9'\xe5" +
	"{@^\x1cn\x9c\xa1\xe6\x05\x1d'\xc4\x9e\x8e\x88\xd8" +
	"\xd3!\xa5\xaf\xef\xe0\x89\xf8\xc5S\x0d\x07\xd2_\xfd\xe6" +
	"\xb7AX\xcc\xcd\xa0\x1d\xea8,\xca\x1d\xc4v\xa4#" +
	"J\x94\xec\xed\xa5|W\xff\xf8\x03\xb3\xd9\xfe\xfc\xc2g" +
	"\xc5\xe3\x17\x92\xf5~q\xa1\xc3\xf4\xf9\x9b\xfe{\xc5\x81" +
	"W\x9f{\x90$\xc6\x9a\xc4K\x85\x08\x02\x88\x07/\xfc" +
	"=\xa0x\xe8B\xb2\xc3SI~Q\xe8\xbf\xbe\xfa\x10" +
	"\xd9\x01\x03&n \xf5o\xb8\xe8\x84\xc8.\"\x11\xd4" +
	"\x8b\xee%\xbe\xd7\xbf\xfaR[\xc3\x9e3\x0f\xcd\xcd7" +
	"\xbc\xeaC@\xb1i\x15\xd9\xec\xef_a\xdb\xc6\xaf\x99" +
	"\xf8\x87:\xca\x9dB\x04C\x00\xe2\xa9U\x7f\x02\xec\xfd" +
	"`\x95c\xafs~zg\xec\x07\x8f\xc6\x1e\x05aq" +
	"\xbd\x08\xbdx\x09\x87b\xcb%\xf4\xd8t\xc9uD}" +
	"\xef\xd6\xdd\xb7]\xff\xc6\xce\x9f\xd4K\xcc\x91\xc4\x93\xab" +
	"?\x14o_MO\xb7\xac&\xf5Vlxe\xed\xca" +
	"\xbf{\xe3P\x1dk$\x8a\x95]'\xc4U]\xf4t" +
	"AW\x02\xd0~\xed\xeb\x17?\xf6\xac<\xf4\x12\x08\xe7" +
	"b-\xf4\x9c\xf8\x11\x07\xbb\xfe$\xca\x0e\xedH\x17\xf1" +
	"}\xff\xad\x97~6\xb5\xda8\x0e\xc2\x9f\xa3\xfd\xe0\x8b" +
	"\x85\x9e/\xdf\xfa\x8d\x93n`\xf6\x0a\xdd}(\x9e\xd7" +
	"M\x8f+\xbb\x1d\x05\xad\xe8\xdd\xab?\x1d\xd0\xde\xae\xf7" +
	"q\xef\xe0\xa5\x1c\x8a\xf2\xa5\x8e\x93/}\x8dH_\xff" +
	"\xf1\x97:\xdf\x12\xcf\x7f\x07\xe4?C\xce^v\xdf\x8a" +
	"\x8f\xa7\x0e\xa7?\x81A.Bf{j\xed\x09\xf1\xe0" +
	"Z\"?\xb0\xd65\xc6}\x87\x1f9R8\xf2\x0e\x08" +
	"\xe7\x86j2\x03\xf6\xae\\\x97D\xb1k\x1d\xc5\xc4\xc5" +
	"\xebxL\xaf]\xe7d\x94/j\x9d\xf1\x1c%/X" +
	"wB\xecYw.\x80\xb8i\x1d\xf9\xd0\xf7Z\x1d\xb1" +
	"\xe0$\xc9\xfa\xc3\xe2\x86\xf5\x17\x01\xf4\xca\xebu\x0e\xd0" +
	"\xdes\xc2\xbe\xf2_7\xbf\xf7n]\xb2:\x19u`" +
	"\xc3Q\xf1\xd0\x06z:\xb8a\x02\x02\xd99\x87[\x84" +
	"\xbe\xa3\xe2\xf2>\x92\xe2\xbc\xbe\x84c\xbcC\xff\xf8\xdc" +
	"\xbb7\xbf_G\xeb\xb8{g\xdf\xb3\xe2\x0d}\xf4\xb4" +
	"\xab\xcf\xe1\xfbt\xe3\x92\x87\xf6\xdf\xff~]\xd4\x0dr" +
	"\x11\x1e@<\xd2\xf7\xaex\xdc\xa1>\xd6\xf7\x0c\xe0\xe7" +
	"\xbb\xd7?x\xffw\x8f}\xe0\xd0\xd6\xfc\xed\x14\x97\xfe" +
	"\xcb\x7f/\x8e\\N\x86\x1e\xba\xdcq\xe1\x1d\xb9\xe7n" +
	"\xbe\xe8\xd0\x93\x1f\xcd%\xc5\xe0\xc6\xa3\xe2\xc8Fz\x1a" +
	"\xdaHR\xf8^\x9bC\xbb\xa77>+\x1e\xdcH\xda" +
	"\x1d\xd9H\xda}9\x9a\xf9\xe7\x1f\xacx\xf8\x7fA\xfe" +
	"\x0b\xf4%\xda\x19\x8e \x07\xd0{zc;\x02\x8a\x7f" +
	"p\xb8>z\xf5\xa6u7\xed\xfb\xcd\x19\x90W`@" +
	"\xd3\xa5\x1c\xe5^\xaf\xba\x89#R\xb6i\x02\xec\xc0\xbf" +
	"\xb2m(\xd9=J^\xed\xe6\xb2\x8a\xa1\x19}%%" +
	"\xcf\xb2WWJ\x19\x15\xcd\x14\"\xb6\x00\x87-\x00\x02" +
	"&\xed{_\xfa\xd5\x7f\x0e~\xe9\xc3\x07\xb9\xea'\xee" +
	"\x17#JBc\xa3j\xd9\x92\x9b1\x18p\xc2]\x81" +
	"\x8cY\x9a\xac\xca.\x08}v\x99\xedS\x87Y\x89\xa1" +
	"5\xa4]\xa7\x9b\xb92\xc0\xf4\x80^*)Z.\xd1" +
	"\x9f\xb5\x98\xae\xc9\x1d|\x08C!\x04\x10N\x9b\x00\xf2" +
	"G<\xca\x7f\xe4P\x08c\x0ci\xf1\x0f\xb4\xf8)\x8f" +
	"\xdb\x91C\x8c\xc4\xc8\x1a\xc2\x99$\x80\xfcG\x1e\xd31" +
	"\xe4Ph\xc4\x18\x92S\x05\xbc\x15 \xdd\x8a<\xa6W" +
	"\xd0:r1\xa7\xc2,\xc7\xdd\x00\xe9e\xb4\xdeA\xeb" +
	"\x1c\x1f\xc30\x05\x15\x1e\x05Hw\xd0\xfa\x1aZ\xe7\xb9" +
	"\x186\x00\x88]\xb8\x1f \xbd\x86\xd67\"\x87\xad|" +
	"\xcc\xf1\xd7\x06\xdc\x0a\x90\xfe2-o&\xf2\x86P\x0c" +
	"\x1b)@\x9c\xf5+h}\x189\xb4KL\xeb7\xd8" +
	"\xb5*Hf\x99\xe9\x1a6\x02\x87\x8d\x80vI\xd9;" +
	"\xd7\xfa\xb4\xe2\x98\xa1\x8c\x8b\x01S<bk\xd5|\x80" +
	"\xb4dgu\xcdbZEE\xcfl\x80\xad5k\x03" +
	"b+\xa0\xad\x18\xc6\xb5\xaaYf\xc0\x07\xf7c\xdaN" +
	"#o*\\N\xc9\x14\xd5~\x8fD\xd7\xc0\xa7P\x0c" +
	"cD1\xf7\xa8h1-O\xbfFH\xaeV{\xe8" +
	"\xdd\xbfYo\xfe\xe4\x86\xbb\x03\xecw0\xab\xa8\x02\xc0" +
	"\x1c\xbf\x96TK\xc9)\x96\xe2\xfe\xeag\x9a\xf7k5" +
	"\x82x7\x84\xd2z\xc5\xcc\xaa#\x8a\xd1=\xa2\x18\x06" +
	"\xd3\xf2\x90B\x94\x9b\xf9\x10\x80\x13\x02\x83\x19\x00y3" +
	"\x8fr\x8a\x1c\xe8\x85\xc0\xc8n\x00y\x98G\xf9z\xf2" +
	"\x1e\xe7\xc6\xc0\xce\xed\x00\xf2\x0e\x1ee\x83\xf3wIA" +
	"D\xb1\x0a\xd8\x0c\x1c6\x03\xdaeg\xb3\x94\x02|`" +
	"\xb1\xc0rjJ\xb1\x0a\x80\xbe\xc9\xe9\xa7\xc5\x01Y\xd1" +
	"\x95\xb5\xdf\x88\x18C9\x92\xaf\xd5\x97Oi\x07\x90\xff" +
	"\x92G\xb9\x10\x90O\xa5\xc5\x1by\x94\x8b\x01\xf9\x18-" +
	"\xe6\\\xf9\x04\x9ewBT(\xd1b\x81G\xd9\xe20" +
	"\xc2rk\xb0\x098l\x02z\xee\x09<\xc7\x03\xcf\xbd" +
	"\xd5\xe7zS\x8exv\xefNgMU\xd5\xca\x05\x1d" +
	"-\x126\xe6\x0b;\x15\x07\x90\xf7\xf2(\xdf\x16\x10\xf6" +
	"\x96>\x00\xf9f\x1e\xe5;8l\xe1l\xdb\x95\xf6v" +
	"\xca\xa8\xaf\xf1(\x7f\x83\xc3\x16\xfes\xdb\x15\xf7N\x12" +
	"\xf76\x1e\xe5oq\xd8\x12:c;\xf9$\xdc\xd3\x09" +
	" \xdf\xc1\xa3\xfc\x00\x87\xd2\x04\xcbY\x85jD%\x0a" +
	"*\xcb\x17,?\xb4+\xda\x1eM\x9f\xd0\xa0!bh" +
	"y\xaf\xc0`t\xccP\xfd\x97\xfa\xa2\x94U,5\xaf" +
	"\x9b\x93C\x1a?\xaa\x936!\xe4\xecO\xee\xb9\xf4\xdc" +
	"%7\x1e:\x0ar\x88\xc3\xfee\xe40\x10\xf0\xb0=" +
	"\xe0\x11\xe3\x906\xaa\xf7k\x9a\x9e\xb0\x14\xca%'\x10" +
	"\xfd~?\xdb\xb1\x03J\xc2\xfd\x906X\x8d\x1c.\x12" +
	"\xa6\xe2\xc2\x94\x04 <2&<.!\x0a\xc7\xf6\x09" +
	"\xc7%\xe4\x84\xd3}\xc2i\x09y\xb1\x05MQ@\x09" +
	"C\xe2e\xd8'^\x86\x12\x86\xc5\x9d\xb8O\xdc\x85\x12" +
	"6\x88S\x98\x14\xa7P\xc2F\xf1\x11\x8c\x8b\x8f\xa0\x84" +
	"M\xe2\x11\x8c\x8bGP\xc2\x88x\x12\xb7\x8a\xa7P\x92" +
	"t\xab\xa0\x9a\x8eB5XH\x0a\xc5(S\xe0\x0a\x14" +
	"0.]C4\x00\xb6a\xea\xb9J\xd6b\x10\x1dg" +
	"\xd6\xe4\x1c\x1f-\xf3?\x1a\xb3S3\x88\x01\xec\xac^" +
	"*U4\x96\x85\x84c\x8f\xf2\x17~\xbf\xcf\x1e\xa8#" +
	"\x07H\x94\xf5,S\x8a_(l_\"\xed\x10\x01\xd8" +
	"\x13j&U\xc9\x14\x19H\xe5\x02\xd3\xf2\x0blw\x9d" +
	"\x9ai#rH0\x87\x1c \xa1\x8f\x8e\xb2\xac\xba\xc0" +
	"v\xd78D\x00vN\x1dW\x8b\xba\xa1B\xc2\xdc\xa1" +
	"\xeb\xc5/Vo\xab\xbdY\x1dw\xc8\x00p\xba\x9ce" +
	"\xaa\xb6\xc0N\xc9\xe9\xb4K\x05 \x95\xd4\x1cS\x16\xf0" +
	"\xd9\x08\xd1\x00Hy\xa5\xa4\x96\x17\xa0\xddB4\x00v" +
	"\xdeT\x8c\x02\xcb\x96\x01`\x01\xe9\xb7\xf8\x94\xb3\xe2X" +
	"/K\x8cB\x9f\xbawk\x0d\x0d\xcf\xce*\xbfT(" +
	"R\xc5*\xe8&\x15\\\x01\x1bcx\x0e\x95\xdc\xb1@" +
	"\xc9\xe5\x9abNv\x8d\x8c\xd5Jnkk\x0cE\xaa" +
	"\xb8\xfb\xbc\x8a{#\xe7\xb6%%k\x0dB\xb4\xa4\xb0" +
	"\xa2_]\x8d\xbc\x91fyM\x81\xa8U1U?\xcb" +
	"+F\xd92U\xa5\x04\x89~G\x80\x1a\xfdLA\xaf" +
	"UM6\xca\xd4\x9c\x9f\xfe+\xfcbv\x90\x8a\xd9\x8f" +
	"x\x94_\x08\x14\xb3C\xd4\x04\x9e\xe7Q\xfeE\xa0\xf2" +
	"\xfe\x9c(_\xe0Q~\x99C\xf4\x0a\xef1\xaao/" +
	"\xf2(\xff\x92C!\xc4\xbb\x85\xec\xf8]\x00\xf2/y" +
	"\x94\xdf&\xbc\x11rP\x81\xf0\xe6Q\x00\xf9m\x1e\xe5" +
	"\xf7\xa9\xc7\x87\x1dH \x9c\xda\x0a \xff\x96G\xf9#" +
	"\x0e%\xc50\x86r\xd8Z\x1b\x1cg\xb6\xba!@\xfa" +
	"\xd5\x07\xc1\xee\xaf\x92E\x1dtv\xfb\x9c\x1eW\xeba" +
	"\x82\xb9G\xb5\x98\x86\xf9j\xbf\x9e\xab\xe9*\x8e\x19S" +
	"y\xce\xd8\xa6N^\xc9\xb4\xbcj\x1a&\xd3,\xf0-" +
	"{vm\xb9>@X4K\xf5\"\xd0\xf0\x92s5" +
	"\xbc\xf8\\\x0d\xaf/\xd8\xf0\xa21\\B\x0d\x8f\x1cT" +
	"\xe4Q\xde\xcb\xe1\xb4b\x18[LF\xc6\xf11\xabg" +
	"\x9c\xbc\xa90m\xf6z\xc25\xc6\xec\x1f<+%\x19" +
	"`~\x8e_\xe7\xd3\xae\x98`YU+\xab\xd4#\x9b" +
	"\xa9\xf5\x91\xc3\xa7:k]r%~n\xa3\xeb\xf2[" +
	"v\xd7:\xe2J\xee\x8c\x1dvP\xa0pg\xa6\xd6\xfc" +
	"V\xf2\x9f\xd1r#\x80p\x1f\xe5\xcb\xb7x\x94\x1f\xa6" +
	"\xf8j\x88a\x13\x80\xf0\x1d\xb2\xdd\x03<\xca\x8fq\x18" +
	"\xd5tM\x85\x06[7T\x8d@\x10\xf0Y\x15\xa3\xb5" +
	"y\x10\x10\xa3$\xb9\xa9\x1b&S-\x88(\xe6\xe4\x1c" +
	"~7\xa8\xccf7\xebN\xda\xcd\x81\xd5\xa65\xddb" +
	"Y\xb5<\xd7\xa73\xe1C\xca}\xdd\xac\x8e2\x8d9" +
	"]\x93\xf2m\x99\xef\xf8\xef|\xa9&\xbc\xef\xf8G(" +
	"\x0f\x1e\xe6Q~2\xe0\xf8'\xc8\xc7\xdf\xe7Q\xfeQ" +
	"\x00\xe9<M\x94?\xe4Q~\x9e\x0c\x12r\x13\xee " +
	"a\xf9\x1f\xf3(\xbfN\x09\x17v\xed\xff\xab\xb1Z\x16" +
	"\xf2,W\x0b`\xc5\x1d6\xdc\x00\x0eL\xb4\x8e6e" +
	"\x0fH\x02\x1a\x14\x00\xd5)\xd2\xfbu\x94\x15\xd5a\xe6" +
	"~[\xe5\xa7\x14'\x94\xc9\xf2\x90\x06R\xb6X\xc9\xa9" +
	"\xf5\xe0/c\xb2\\^\x1d\xd0!\xaa\x8d2\x0a*\x7f" +
	"@\x9e;\xa8\xae)\xb3a7\x9a\xa2^\x0d\x0e\xc2X" +
	"2\xde\x15<\xca\xc3\x01\xe3\x0d\xc5k\x85\x96f\x13\x0e" +
	"Q\x181ku6\xa0\xbcW-\xaa\xa2\x9b\xeaM\x15" +
	"f\xaai\x90\x1c\xad\x11\x81C\x9c-\xd2\xc0,\xf0\xe4" +
	"\x0b\xd4B{7\xf2(\xc7\xb8z\xe6\xf5\xc9\xe2\x99\xbd" +
	"{ \xe1\x0e\x1c\x9ef\xc8\xb9\x9a\xbd\x12h\x0c\x9eb" +
	";)\x83R^9\xc0j9Hz\xe5\xe0k\x1c\xda" +
	"9\xd50\xd5\xacbqjnp\xaf\x9a\xadX4\x92" +
	"8(\xdc\xf7OT1\xf3\xe3uN\x99V\xb5qf" +
	"\xea\x9a?\x1b\xd9\x1dk\x0f\xfc\xd3\xfb\xb7<\xf6\x00x" +
	"\xd3Q]P'='\x92\x0b\xbbS\xfa\x84j&2" +
	":\x8d]u\xc0\xb8s.`\x9c\xa9%\xbc\x1f\xdbw" +
	"v\xd6\x10\xb0\x1f\xdb\xf7\xb4\xd7j\x80\x1f\xdb\xf7e\xbc" +
	"\x0a\xf0CJv\xa5T3p\x8e\x95\x8d\xa229\x04" +
	"\x11mT\xc7V\xfb\xa3K\xde\xfbi\xd3\xfe\xf4\x13^" +
	"`E\x8d\xc0\xcc\x12\xb1\x14\x8a\xbd\xf8\xc7S\xd7\xbe\xf2" +
	"?\xa7\xbe\xed\xc7\x9ej\x96X\xb9\xcc \x12\x98\x14\xb1" +
	"nl\xf1lp\x8dWa\xb2\xaa\x17\x9fNb\x0f;" +
	"\xea\x9cL\x12G\xe1\xcdv\xe1M\x0991\x8cI1" +
	"\x8c\x84o7a\xa7\xb8\xc9\xc1\xb7%\x8c\x8b%\x07\xdf" +
	">\x8e\xbb\xc5'\x1c|\xfb\x0e\xee\x16O:XV\xe0" +
	":E\x81#\xa4\xdb\xcf\xc5\xc5~\x8e\x90n\x85\x8b\x8b" +
	"\x15N\xc2E\xe2S\\\xbb\xf8\x14'\xe19\xe2;\xdc" +
	"v\xf1$'a\xb3\xb8\x94O\x8aKy\x09[\xc4\x11" +
	"\xbeS\x1c\xe1%\\,N\xf1\x9d\xe2\x14/aT<" +
	"\xc8w\x8a\x07y\x09[\xc5\x0f\xf8v\xf1\x03^BA" +
	"\\\x1e\xea\x14\x97\x87$\\\"\xee\x0au\x8a\xbbB\xd2" +
	"4\xd3\xc6\x95\"\xcbEJ\xccr\xd0R\xed\xfc\x93\xd0" +
	"R\x87\x83\x96R\x8e\x82\xe1v!\x9c\x89\x8c\x0c\xed\xb0" +
	"G\x86v\xb4\x91\x01 B&\x80i\xc5P\xb2\x055" +
	">\xc7\xf7\xab\x83\xdf\xef\x16\x9a\xee\xb7\xfb\x1d\xda\xae8" +
	"\xf0\xddk\xbc\x97\xb6at\xad\xd96\x8eq\x80h\xde" +
	"(\xf6~\x91,\x88B8)\x84wOoI\x0dw" +
	"\xf5v\xaf\xb1\xb7\\\xbd\xb3mKj\xb8\x0d\xf8\xf1^" +
	"\x00I\x99\x87\xc1\xc5A\x06[\x85\xa6\x8c\xdd\xefq\x00" +
	"\x00\x87\x09\xbdC\xa4\x8d\xb8\xd8\x99r\xaew\xa0\xa8T" +
	"\x80/\xab_\xc4\x8c4\x1b\x13\x9a\xc6\xecdzsW" +
	"o\xd7@\x11\xa2J\xa5\xac\xd2k[\xed\xd5e\x18\xff" +
	"\x7f3\x8c\xcfd\x18\x0f0$;\xc5\xcf\xdaN\xf1\x99" +
	"v\x8a\x03H\xc5y\x18\xd4\xdb\xc9\xb4\x87]\x0e=U" +
	";\xd1;Hm\xe3\xce\x92\xc3\xe7\xec\xec=\\g\xef" +
	"\xe1\x9a\xbd#\xac\x9c=\xab\x00\x1cJ\x0f\xd8C\xe9\x81" +
	"`\x00\xda\x8ai\xb1\xb2\xc5\xb2\x80\x0b\x06\xe1\x98\xd0\xb4" +
	"\xdf\xee\xaf\xd2G\x1d\xbbT_\xb1\xcdK\xebH\x1b\x19" +
	"h\xda\x98\xb4\x0a\xbavvq\x9drh\xbd\xb8v_" +
	"\xea\xe3\xda(\x18\xbd\x0b\xa9\x98\x14\xc2\xfb\xa6SW\xa5" +
	"\x9c\xb8N]\x95r\xd5L8Lz\x01\xa2\xa5\xb9]" +
	"\xb6\xba\xce\xe7oL\x8fx>\x1f\xd1\xf7\xb1bQi" +
	"\xe3\x9c)2[\xd5\xd0\x13)\x9b\xcb\x15\x17\x0a\xa1\xad" +
	"BS\xa7=\xb0y\xf3pW\x8f\xe3\xb9(=\x03D" +
	"T\xa3\xb8\x90\xd3\x93B\xf8\xf0\xf4`\xca\xf9\xd2\x1e\xcc" +
	"\x16\x99QV\xdb\xb0*\x89\x87\x18I\x0cC\x99\x8b\xd7" +
	"\xda\x19b\xf44-A{ \xd5_\x95\xc3\x99\xadu" +
	"\xad-\xc5\xbb\xfc\xfa-\xcbd\x99\x0a\xe1-OK$" +
	"\xde\xfb\x8a,\xb3P\xbau\x0a\xe1\xfd\xd1\xddE\x96\xb1" +
	"\xe9\xbfK\x8b,\x83\x86\x96w\xb8x\x11V\xd7\xc9\xfb" +
	"\xcdl\x81\x8d\xab\xddW2\xbe\xa8\xd6\x81\xbb\xce\x1a\xb8" +
	"[\xc9\xd9v\x15\xdeQ\xcf~\x88G\xf9\xfb\x04n\x1d" +
	"(L[?NP\xf81\xb7\xb5\xad\x0c\x9d\xb1\xdd\xf3" +
	"V\xe1\xa9d\x0d\xf7\xad\x0c\x7fF\xcb\xd4\x08\x9f\xde\xee" +
	"!\xbf\x97\xbd\xe6J \xef\xd8a\x00\xf9e\x1e\xe5_" +
	"\xd7u\xc7iS\xcdW\x8a\x8a\xe9\x8f\x8a\xaa\x87\x0f\x80" +
	"/\xfa\xf3\xe3ty\xb2Td\xda\x9e@K5\xd5\xac" +
	"\xa5\x9b\x80\x935l\xe0_\x98x\xd8\xa0\xa8\x94\xad\x11" +
	"=\xc7p\x94e\x9d\x83\x8d\xc4\x0eVR\xaf.c\x18" +
	"8\x0c\x07\x1ag\xa8\x0e\xf9T\x0f\xad\xfd3k\x1a\xb0" +
	"\x9dS\xb8P\xebl3\xd70\x07\xcf\xf2r\x08\x83W" +
	"B\x98\xb1\x1d\x14\x92\xd1\xf7B\xa4\xdf`\xc1\xc1jk" +
	"`\x86\xaab\x10\x1f4\x19\x1eD\xa4\x11\x8a\xa6T\x83" +
	"G\xf9f\xef\x8c\x9al?9\xe6!\x98\x078\xb4\xc7" +
	"\x99:A\x18\xd4E\xca\xea\xef\x9e\xec\xdfr\xd9_\xbd" +
	"P\x9d\x08\x14\x83\xa5f\x9c\x8a*\xe3\xeaPN\xd5\xd0" +
	"b\xd6\xe4\x80b\x94\x89\xb0\x0a(\xab\xc2F\xfb\x0dV" +
	";\x92\xaeiT\x07\xbb\xb0:`%\xdc\x09Kn\x0c" +
	"^HPN\xfa7\xa6Bxwt(\xabkv\xf5" +
	"\xb8\x12x\xdd\x92W\x93=0p\x9d(^\x80q\xe0" +
	"\x04on\x17[0\x09\x90n\xc4\xeaY?\xe7\xa2-" +
	"Qp\xd6\x9bi}\x19\xa5\x1e\x8f\x81\x8bFq)&" +
	"\x81\x13B\x91\x18.\x02\x10\xceP\xf8~V\xa5\x0cc" +
	"\xe0RT\\\x8a}\xc0\x09\x0d\x8bb\xce\xd5\x87O\xd9" +
	"H{E\xce\x89\xe1b\x001\x8c\x19\x80t\x88\xf6j" +
	"u\xee\x1b\x9ac\x18ud\xbbk\xc6}CSK\x8c" +
	"\xd2U\\\xee\xd0;\xf7\x0dki}\xd1\xe2\x18\x0a\x00" +
	"b\x0fn\x0f\xde+H\x8cF\xed\xe9\x095SfV" +
	"-\x1f\xb2zN\xddi\xfa'*\xd3E\xb7\x00\xd9\xde" +
	"\xc9(\x03^\xf5\x1d\x13\xad\xdd\x19\xbb\x8eI\xb8\xe7\x02" +
	"\xb6\x917\xb6\xa9\x93D\\;w\xb5sj9k2" +
	"\xc3\x82y\x8e\xf9\xcb\x05\xdd\xb46\xabet\xa8\xe6;" +
	"y(W\xdd\x17\xd1\xad@\x84\xf8~\xae^[\x14\x14" +
	"-\xaf\x0e\xeb\xee\x80>\xdf,\xca\xcd\x98E\x87r\x0e" +
	"Tm\xf4sd\x15\x01\xed\x0e\x1e\xe55\x81\x1c\xe9\xa2" +
	"\xc5\x8by\x94\xd7\xce\x7f\x86^\xbf\x81\x7f\xed\x00N\x82" +
	"\xd6\xfe>\x00\x93\xd3\xdeEDpX\xa20h\xe6Q" +
	"\xbe\x98C\xbb\xac*f\xb6\xe0]\"\xf8\xea\xfa\x1c\xe6" +
	"N\x88m\xaa4y%s\xabn@\x9d\xed5\xc9}" +
	"uzh\xb35<\xca\x1b\xb9\xeax\xbfM\xa5\xba\xe6" +
	"\x9f\x90\x9bl\\\xb1\xd4m\xc0\xab\x93\xf3\x1d\x9b\xfb'" +
	"\x1cC\x91\xac\xaey\x13[\xb3m\xbb#[\xb26\x8c" +
	"\xb6\xe0\xe7\xb67\x8d\xb6\xd7\xa6\xd1\x16\xee\x8c\x8d\x81{" +
	"Ea\xa4\x1d\xb8\xc0\xe9}y<\xef\x0f(\x86\x96\x9f" +
	"wd\xec\x97\x9c\xbb,\x7f\xd8\xc2\xc0\xad\xa60\xe5\xe4" +
	"\xb5\xd7R\xd4d\xdd\x09\x12\x155\x16\x0f\x9c U\x0f" +
	"\xe9J\x99\xda\x09\x92\xc0{\x87t\x152\x9a\xe5\x0e\x97" +
	"\x12\xd3\x8c\x8a5\x9du\xa7\xd5\xd9\xb7c\xf3\x9d\xbc-" +
	"\x98\x0f\x9a^\xd1R\x05S!,\xbcp\x0c;\xc7\x9e" +
	"V\xc5T\xe1l\x9c\xbe}\x01\xa7\x97=n\x80\xea," +
	"\x9fW\xaf\xa7L\xc9\xe9\xefND\xfb\xbdO\xc0\xce(" +
	"E\xde\\\xb3\x7f\x07\x87\xd2(+\xaa\xe5\xf9{f\xfd" +
	"\x1d\x93\xefY\xc7\xb1\xddL\x8b\x18\x15\x8b\xb4\xab\x06\xd7" +
	"\xaa\xceZ\x8a\xd6\x82\xab\x8b\xfc\xb3\x9aG\xf9\xaa\xda\xc9" +
	"WV1\x94\x0c+2\xe0\xad@\xd3\xbe~\xc5\xbf\xfc" +
	"\xec{\xd1\xd7N\xcf'\x80\x1f\xd8Y]\xeb6\xb4<" +
	"\xb8\xa6u\xa3hU\xbc\xae>P\x14u\xc5k\x06\x97" +
	"r\x06\xeb\xd9[\xb5 \xbd\xc5\xfd\xb7\xff\x0b\x00\x00\xff" +
	"\xff\xe0\x93,\x12"

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
		0xc9702c7dbfc6d7e4,
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
