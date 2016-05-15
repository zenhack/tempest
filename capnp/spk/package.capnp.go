package spk

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
)

const OsiInfo = uint64(0x9476412d0315d869)
const CategoryInfo = uint64(0x8d51dd236606d205)
const (
	Manifest_sizeLimitInWords = uint64(1048576)
)

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
	root, err := msg.Root()
	if err != nil {
		return PackageDefinition{}, err
	}
	st := capnp.ToStruct(root)
	return PackageDefinition{st}, nil
}

func (s PackageDefinition) Id() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s PackageDefinition) SetId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s PackageDefinition) Manifest() (Manifest, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Manifest{}, err
	}

	ss := capnp.ToStruct(p)

	return Manifest{Struct: ss}, nil
}

func (s PackageDefinition) SetManifest(v Manifest) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewManifest sets the manifest field to a newly
// allocated Manifest struct, preferring placement in s's segment.
func (s PackageDefinition) NewManifest() (Manifest, error) {

	ss, err := NewManifest(s.Struct.Segment())
	if err != nil {
		return Manifest{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s PackageDefinition) SourceMap() (SourceMap, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return SourceMap{}, err
	}

	ss := capnp.ToStruct(p)

	return SourceMap{Struct: ss}, nil
}

func (s PackageDefinition) SetSourceMap(v SourceMap) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewSourceMap sets the sourceMap field to a newly
// allocated SourceMap struct, preferring placement in s's segment.
func (s PackageDefinition) NewSourceMap() (SourceMap, error) {

	ss, err := NewSourceMap(s.Struct.Segment())
	if err != nil {
		return SourceMap{}, err
	}
	err = s.Struct.SetPointer(2, ss)
	return ss, err
}

func (s PackageDefinition) FileList() (string, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s PackageDefinition) SetFileList(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(3, t)
}

func (s PackageDefinition) AlwaysInclude() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s PackageDefinition) SetAlwaysInclude(v capnp.TextList) error {

	return s.Struct.SetPointer(4, v.List)
}

func (s PackageDefinition) BridgeConfig() (BridgeConfig, error) {
	p, err := s.Struct.Pointer(5)
	if err != nil {
		return BridgeConfig{}, err
	}

	ss := capnp.ToStruct(p)

	return BridgeConfig{Struct: ss}, nil
}

func (s PackageDefinition) SetBridgeConfig(v BridgeConfig) error {

	return s.Struct.SetPointer(5, v.Struct)
}

// NewBridgeConfig sets the bridgeConfig field to a newly
// allocated BridgeConfig struct, preferring placement in s's segment.
func (s PackageDefinition) NewBridgeConfig() (BridgeConfig, error) {

	ss, err := NewBridgeConfig(s.Struct.Segment())
	if err != nil {
		return BridgeConfig{}, err
	}
	err = s.Struct.SetPointer(5, ss)
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{st}, nil
}

func NewRootManifest(s *capnp.Segment) (Manifest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	if err != nil {
		return Manifest{}, err
	}
	return Manifest{st}, nil
}

func ReadRootManifest(msg *capnp.Message) (Manifest, error) {
	root, err := msg.Root()
	if err != nil {
		return Manifest{}, err
	}
	st := capnp.ToStruct(root)
	return Manifest{st}, nil
}

func (s Manifest) AppTitle() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Manifest) SetAppTitle(v util.LocalizedText) error {

	return s.Struct.SetPointer(3, v.Struct)
}

// NewAppTitle sets the appTitle field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest) NewAppTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(3, ss)
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
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Manifest) SetAppMarketingVersion(v util.LocalizedText) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewAppMarketingVersion sets the appMarketingVersion field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest) NewAppMarketingVersion() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(2, ss)
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
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return Metadata{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata{Struct: ss}, nil
}

func (s Manifest) SetMetadata(v Metadata) error {

	return s.Struct.SetPointer(4, v.Struct)
}

// NewMetadata sets the metadata field to a newly
// allocated Metadata struct, preferring placement in s's segment.
func (s Manifest) NewMetadata() (Metadata, error) {

	ss, err := NewMetadata(s.Struct.Segment())
	if err != nil {
		return Metadata{}, err
	}
	err = s.Struct.SetPointer(4, ss)
	return ss, err
}

func (s Manifest) Actions() (Manifest_Action_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Manifest_Action_List{}, err
	}

	l := capnp.ToList(p)

	return Manifest_Action_List{List: l}, nil
}

func (s Manifest) SetActions(v Manifest_Action_List) error {

	return s.Struct.SetPointer(0, v.List)
}

func (s Manifest) ContinueCommand() (Manifest_Command, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Manifest_Command{}, err
	}

	ss := capnp.ToStruct(p)

	return Manifest_Command{Struct: ss}, nil
}

func (s Manifest) SetContinueCommand(v Manifest_Command) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewContinueCommand sets the continueCommand field to a newly
// allocated Manifest_Command struct, preferring placement in s's segment.
func (s Manifest) NewContinueCommand() (Manifest_Command, error) {

	ss, err := NewManifest_Command(s.Struct.Segment())
	if err != nil {
		return Manifest_Command{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// Manifest_List is a list of Manifest.
type Manifest_List struct{ capnp.List }

// NewManifest creates a new list of Manifest.
func NewManifest_List(s *capnp.Segment, sz int32) (Manifest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5}, sz)
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

func (p Manifest_Promise) Metadata() Metadata_Promise {
	return Metadata_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
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
	root, err := msg.Root()
	if err != nil {
		return Manifest_Command{}, err
	}
	st := capnp.ToStruct(root)
	return Manifest_Command{st}, nil
}

func (s Manifest_Command) Argv() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s Manifest_Command) SetArgv(v capnp.TextList) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s Manifest_Command) Environ() (util.KeyValue_List, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return util.KeyValue_List{}, err
	}

	l := capnp.ToList(p)

	return util.KeyValue_List{List: l}, nil
}

func (s Manifest_Command) SetEnviron(v util.KeyValue_List) error {

	return s.Struct.SetPointer(2, v.List)
}

func (s Manifest_Command) DeprecatedExecutablePath() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Manifest_Command) SetDeprecatedExecutablePath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	if err != nil {
		return Manifest_Action{}, err
	}
	return Manifest_Action{st}, nil
}

func NewRootManifest_Action(s *capnp.Segment) (Manifest_Action, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	if err != nil {
		return Manifest_Action{}, err
	}
	return Manifest_Action{st}, nil
}

func ReadRootManifest_Action(msg *capnp.Message) (Manifest_Action, error) {
	root, err := msg.Root()
	if err != nil {
		return Manifest_Action{}, err
	}
	st := capnp.ToStruct(root)
	return Manifest_Action{st}, nil
}

func (s Manifest_Action) Input() Manifest_Action_input { return Manifest_Action_input(s) }

func (s Manifest_Action_input) Which() Manifest_Action_input_Which {
	return Manifest_Action_input_Which(s.Struct.Uint16(0))
}

func (s Manifest_Action_input) SetNone() {
	s.Struct.SetUint16(0, 0)
}

func (s Manifest_Action_input) Capability() (grain.PowerboxDescriptor_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return grain.PowerboxDescriptor_List{}, err
	}

	l := capnp.ToList(p)

	return grain.PowerboxDescriptor_List{List: l}, nil
}

func (s Manifest_Action_input) SetCapability(v grain.PowerboxDescriptor_List) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPointer(0, v.List)
}

func (s Manifest_Action) Command() (Manifest_Command, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Manifest_Command{}, err
	}

	ss := capnp.ToStruct(p)

	return Manifest_Command{Struct: ss}, nil
}

func (s Manifest_Action) SetCommand(v Manifest_Command) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewCommand sets the command field to a newly
// allocated Manifest_Command struct, preferring placement in s's segment.
func (s Manifest_Action) NewCommand() (Manifest_Command, error) {

	ss, err := NewManifest_Command(s.Struct.Segment())
	if err != nil {
		return Manifest_Command{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Manifest_Action) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Manifest_Action) SetTitle(v util.LocalizedText) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(2, ss)
	return ss, err
}

func (s Manifest_Action) NounPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Manifest_Action) SetNounPhrase(v util.LocalizedText) error {

	return s.Struct.SetPointer(4, v.Struct)
}

// NewNounPhrase sets the nounPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewNounPhrase() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(4, ss)
	return ss, err
}

func (s Manifest_Action) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Manifest_Action) SetDescription(v util.LocalizedText) error {

	return s.Struct.SetPointer(3, v.Struct)
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Manifest_Action) NewDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(3, ss)
	return ss, err
}

// Manifest_Action_List is a list of Manifest_Action.
type Manifest_Action_List struct{ capnp.List }

// NewManifest_Action creates a new list of Manifest_Action.
func NewManifest_Action_List(s *capnp.Segment, sz int32) (Manifest_Action_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
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

func (p Manifest_Action_Promise) NounPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
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
	root, err := msg.Root()
	if err != nil {
		return SourceMap{}, err
	}
	st := capnp.ToStruct(root)
	return SourceMap{st}, nil
}

func (s SourceMap) SearchPath() (SourceMap_Mapping_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return SourceMap_Mapping_List{}, err
	}

	l := capnp.ToList(p)

	return SourceMap_Mapping_List{List: l}, nil
}

func (s SourceMap) SetSearchPath(v SourceMap_Mapping_List) error {

	return s.Struct.SetPointer(0, v.List)
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
	root, err := msg.Root()
	if err != nil {
		return SourceMap_Mapping{}, err
	}
	st := capnp.ToStruct(root)
	return SourceMap_Mapping{st}, nil
}

func (s SourceMap_Mapping) PackagePath() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s SourceMap_Mapping) SetPackagePath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s SourceMap_Mapping) SourcePath() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s SourceMap_Mapping) SetSourcePath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s SourceMap_Mapping) HidePaths() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s SourceMap_Mapping) SetHidePaths(v capnp.TextList) error {

	return s.Struct.SetPointer(2, v.List)
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
	root, err := msg.Root()
	if err != nil {
		return BridgeConfig{}, err
	}
	st := capnp.ToStruct(root)
	return BridgeConfig{st}, nil
}

func (s BridgeConfig) ViewInfo() (grain.UiView_ViewInfo, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return grain.UiView_ViewInfo{}, err
	}

	ss := capnp.ToStruct(p)

	return grain.UiView_ViewInfo{Struct: ss}, nil
}

func (s BridgeConfig) SetViewInfo(v grain.UiView_ViewInfo) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewViewInfo sets the viewInfo field to a newly
// allocated grain.UiView_ViewInfo struct, preferring placement in s's segment.
func (s BridgeConfig) NewViewInfo() (grain.UiView_ViewInfo, error) {

	ss, err := grain.NewUiView_ViewInfo(s.Struct.Segment())
	if err != nil {
		return grain.UiView_ViewInfo{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s BridgeConfig) ApiPath() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s BridgeConfig) SetApiPath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
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
	if err != nil {
		return Metadata{}, err
	}
	return Metadata{st}, nil
}

func NewRootMetadata(s *capnp.Segment) (Metadata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 17})
	if err != nil {
		return Metadata{}, err
	}
	return Metadata{st}, nil
}

func ReadRootMetadata(msg *capnp.Message) (Metadata, error) {
	root, err := msg.Root()
	if err != nil {
		return Metadata{}, err
	}
	st := capnp.ToStruct(root)
	return Metadata{st}, nil
}

func (s Metadata) Icons() Metadata_icons { return Metadata_icons(s) }

func (s Metadata_icons) AppGrid() (Metadata_Icon, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Metadata_Icon{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata_Icon{Struct: ss}, nil
}

func (s Metadata_icons) SetAppGrid(v Metadata_Icon) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewAppGrid sets the appGrid field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewAppGrid() (Metadata_Icon, error) {

	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s Metadata_icons) Grain() (Metadata_Icon, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Metadata_Icon{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata_Icon{Struct: ss}, nil
}

func (s Metadata_icons) SetGrain(v Metadata_Icon) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewGrain sets the grain field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewGrain() (Metadata_Icon, error) {

	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Metadata_icons) Market() (Metadata_Icon, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return Metadata_Icon{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata_Icon{Struct: ss}, nil
}

func (s Metadata_icons) SetMarket(v Metadata_Icon) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewMarket sets the market field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewMarket() (Metadata_Icon, error) {

	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPointer(2, ss)
	return ss, err
}

func (s Metadata_icons) MarketBig() (Metadata_Icon, error) {
	p, err := s.Struct.Pointer(15)
	if err != nil {
		return Metadata_Icon{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata_Icon{Struct: ss}, nil
}

func (s Metadata_icons) SetMarketBig(v Metadata_Icon) error {

	return s.Struct.SetPointer(15, v.Struct)
}

// NewMarketBig sets the marketBig field to a newly
// allocated Metadata_Icon struct, preferring placement in s's segment.
func (s Metadata_icons) NewMarketBig() (Metadata_Icon, error) {

	ss, err := NewMetadata_Icon(s.Struct.Segment())
	if err != nil {
		return Metadata_Icon{}, err
	}
	err = s.Struct.SetPointer(15, ss)
	return ss, err
}

func (s Metadata) Website() (string, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Metadata) SetWebsite(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(3, t)
}

func (s Metadata) CodeUrl() (string, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Metadata) SetCodeUrl(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(4, t)
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
	p, err := s.Struct.Pointer(5)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata_license) SetProprietary(v util.LocalizedText) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPointer(5, v.Struct)
}

// NewProprietary sets the proprietary field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewProprietary() (util.LocalizedText, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(5, ss)
	return ss, err
}

func (s Metadata_license) PublicDomain() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(5)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata_license) SetPublicDomain(v util.LocalizedText) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPointer(5, v.Struct)
}

// NewPublicDomain sets the publicDomain field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewPublicDomain() (util.LocalizedText, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(5, ss)
	return ss, err
}

func (s Metadata_license) Notices() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(6)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata_license) SetNotices(v util.LocalizedText) error {

	return s.Struct.SetPointer(6, v.Struct)
}

// NewNotices sets the notices field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata_license) NewNotices() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(6, ss)
	return ss, err
}

func (s Metadata) Categories() (Category_List, error) {
	p, err := s.Struct.Pointer(7)
	if err != nil {
		return Category_List{}, err
	}

	l := capnp.ToList(p)

	return Category_List{List: l}, nil
}

func (s Metadata) SetCategories(v Category_List) error {

	return s.Struct.SetPointer(7, v.List)
}
func (s Metadata) Author() Metadata_author { return Metadata_author(s) }

func (s Metadata_author) UpstreamAuthor() (string, error) {
	p, err := s.Struct.Pointer(16)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Metadata_author) SetUpstreamAuthor(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(16, t)
}

func (s Metadata_author) ContactEmail() (string, error) {
	p, err := s.Struct.Pointer(8)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Metadata_author) SetContactEmail(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(8, t)
}

func (s Metadata_author) PgpSignature() ([]byte, error) {
	p, err := s.Struct.Pointer(9)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata_author) SetPgpSignature(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(9, d)
}

func (s Metadata) PgpKeyring() ([]byte, error) {
	p, err := s.Struct.Pointer(10)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata) SetPgpKeyring(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(10, d)
}

func (s Metadata) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(11)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata) SetDescription(v util.LocalizedText) error {

	return s.Struct.SetPointer(11, v.Struct)
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(11, ss)
	return ss, err
}

func (s Metadata) ShortDescription() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(12)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata) SetShortDescription(v util.LocalizedText) error {

	return s.Struct.SetPointer(12, v.Struct)
}

// NewShortDescription sets the shortDescription field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewShortDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(12, ss)
	return ss, err
}

func (s Metadata) Screenshots() (Metadata_Screenshot_List, error) {
	p, err := s.Struct.Pointer(13)
	if err != nil {
		return Metadata_Screenshot_List{}, err
	}

	l := capnp.ToList(p)

	return Metadata_Screenshot_List{List: l}, nil
}

func (s Metadata) SetScreenshots(v Metadata_Screenshot_List) error {

	return s.Struct.SetPointer(13, v.List)
}

func (s Metadata) ChangeLog() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(14)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s Metadata) SetChangeLog(v util.LocalizedText) error {

	return s.Struct.SetPointer(14, v.Struct)
}

// NewChangeLog sets the changeLog field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Metadata) NewChangeLog() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(14, ss)
	return ss, err
}

// Metadata_List is a list of Metadata.
type Metadata_List struct{ capnp.List }

// NewMetadata creates a new list of Metadata.
func NewMetadata_List(s *capnp.Segment, sz int32) (Metadata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 17}, sz)
	if err != nil {
		return Metadata_List{}, err
	}
	return Metadata_List{l}, nil
}

func (s Metadata_List) At(i int) Metadata           { return Metadata{s.List.Struct(i)} }
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
	if err != nil {
		return Metadata_Icon{}, err
	}
	return Metadata_Icon{st}, nil
}

func NewRootMetadata_Icon(s *capnp.Segment) (Metadata_Icon, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Metadata_Icon{}, err
	}
	return Metadata_Icon{st}, nil
}

func ReadRootMetadata_Icon(msg *capnp.Message) (Metadata_Icon, error) {
	root, err := msg.Root()
	if err != nil {
		return Metadata_Icon{}, err
	}
	st := capnp.ToStruct(root)
	return Metadata_Icon{st}, nil
}

func (s Metadata_Icon) Which() Metadata_Icon_Which {
	return Metadata_Icon_Which(s.Struct.Uint16(0))
}

func (s Metadata_Icon) SetUnknown() {
	s.Struct.SetUint16(0, 0)
}

func (s Metadata_Icon) Svg() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Metadata_Icon) SetSvg(v string) error {
	s.Struct.SetUint16(0, 1)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}
func (s Metadata_Icon) Png() Metadata_Icon_png { return Metadata_Icon_png(s) }

func (s Metadata_Icon) SetPng() { s.Struct.SetUint16(0, 2) }

func (s Metadata_Icon_png) Dpi1x() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata_Icon_png) SetDpi1x(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s Metadata_Icon_png) Dpi2x() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata_Icon_png) SetDpi2x(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
}

// Metadata_Icon_List is a list of Metadata_Icon.
type Metadata_Icon_List struct{ capnp.List }

// NewMetadata_Icon creates a new list of Metadata_Icon.
func NewMetadata_Icon_List(s *capnp.Segment, sz int32) (Metadata_Icon_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Metadata_Icon_List{}, err
	}
	return Metadata_Icon_List{l}, nil
}

func (s Metadata_Icon_List) At(i int) Metadata_Icon           { return Metadata_Icon{s.List.Struct(i)} }
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
	if err != nil {
		return Metadata_Screenshot{}, err
	}
	return Metadata_Screenshot{st}, nil
}

func NewRootMetadata_Screenshot(s *capnp.Segment) (Metadata_Screenshot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Metadata_Screenshot{}, err
	}
	return Metadata_Screenshot{st}, nil
}

func ReadRootMetadata_Screenshot(msg *capnp.Message) (Metadata_Screenshot, error) {
	root, err := msg.Root()
	if err != nil {
		return Metadata_Screenshot{}, err
	}
	st := capnp.ToStruct(root)
	return Metadata_Screenshot{st}, nil
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
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata_Screenshot) SetPng(v []byte) error {
	s.Struct.SetUint16(8, 1)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s Metadata_Screenshot) Jpeg() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Metadata_Screenshot) SetJpeg(v []byte) error {
	s.Struct.SetUint16(8, 2)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

// Metadata_Screenshot_List is a list of Metadata_Screenshot.
type Metadata_Screenshot_List struct{ capnp.List }

// NewMetadata_Screenshot creates a new list of Metadata_Screenshot.
func NewMetadata_Screenshot_List(s *capnp.Segment, sz int32) (Metadata_Screenshot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	if err != nil {
		return Metadata_Screenshot_List{}, err
	}
	return Metadata_Screenshot_List{l}, nil
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
	if err != nil {
		return OsiLicenseInfo{}, err
	}
	return OsiLicenseInfo{st}, nil
}

func NewRootOsiLicenseInfo(s *capnp.Segment) (OsiLicenseInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return OsiLicenseInfo{}, err
	}
	return OsiLicenseInfo{st}, nil
}

func ReadRootOsiLicenseInfo(msg *capnp.Message) (OsiLicenseInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return OsiLicenseInfo{}, err
	}
	st := capnp.ToStruct(root)
	return OsiLicenseInfo{st}, nil
}

func (s OsiLicenseInfo) Id() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s OsiLicenseInfo) SetId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s OsiLicenseInfo) Title() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s OsiLicenseInfo) SetTitle(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
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
	if err != nil {
		return OsiLicenseInfo_List{}, err
	}
	return OsiLicenseInfo_List{l}, nil
}

func (s OsiLicenseInfo_List) At(i int) OsiLicenseInfo           { return OsiLicenseInfo{s.List.Struct(i)} }
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
	if err != nil {
		return OpenSourceLicense_List{}, err
	}
	return OpenSourceLicense_List{l.List}, nil
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
	if err != nil {
		return AppId{}, err
	}
	return AppId{st}, nil
}

func NewRootAppId(s *capnp.Segment) (AppId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	if err != nil {
		return AppId{}, err
	}
	return AppId{st}, nil
}

func ReadRootAppId(msg *capnp.Message) (AppId, error) {
	root, err := msg.Root()
	if err != nil {
		return AppId{}, err
	}
	st := capnp.ToStruct(root)
	return AppId{st}, nil
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
	if err != nil {
		return AppId_List{}, err
	}
	return AppId_List{l}, nil
}

func (s AppId_List) At(i int) AppId           { return AppId{s.List.Struct(i)} }
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
	if err != nil {
		return PackageId{}, err
	}
	return PackageId{st}, nil
}

func NewRootPackageId(s *capnp.Segment) (PackageId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	if err != nil {
		return PackageId{}, err
	}
	return PackageId{st}, nil
}

func ReadRootPackageId(msg *capnp.Message) (PackageId, error) {
	root, err := msg.Root()
	if err != nil {
		return PackageId{}, err
	}
	st := capnp.ToStruct(root)
	return PackageId{st}, nil
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
	if err != nil {
		return PackageId_List{}, err
	}
	return PackageId_List{l}, nil
}

func (s PackageId_List) At(i int) PackageId           { return PackageId{s.List.Struct(i)} }
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
	if err != nil {
		return VerifiedInfo{}, err
	}
	return VerifiedInfo{st}, nil
}

func NewRootVerifiedInfo(s *capnp.Segment) (VerifiedInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	if err != nil {
		return VerifiedInfo{}, err
	}
	return VerifiedInfo{st}, nil
}

func ReadRootVerifiedInfo(msg *capnp.Message) (VerifiedInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return VerifiedInfo{}, err
	}
	st := capnp.ToStruct(root)
	return VerifiedInfo{st}, nil
}

func (s VerifiedInfo) AppId() (AppId, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return AppId{}, err
	}

	ss := capnp.ToStruct(p)

	return AppId{Struct: ss}, nil
}

func (s VerifiedInfo) SetAppId(v AppId) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewAppId sets the appId field to a newly
// allocated AppId struct, preferring placement in s's segment.
func (s VerifiedInfo) NewAppId() (AppId, error) {

	ss, err := NewAppId(s.Struct.Segment())
	if err != nil {
		return AppId{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s VerifiedInfo) PackageId() (PackageId, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return PackageId{}, err
	}

	ss := capnp.ToStruct(p)

	return PackageId{Struct: ss}, nil
}

func (s VerifiedInfo) SetPackageId(v PackageId) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewPackageId sets the packageId field to a newly
// allocated PackageId struct, preferring placement in s's segment.
func (s VerifiedInfo) NewPackageId() (PackageId, error) {

	ss, err := NewPackageId(s.Struct.Segment())
	if err != nil {
		return PackageId{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s VerifiedInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s VerifiedInfo) SetTitle(v util.LocalizedText) error {

	return s.Struct.SetPointer(2, v.Struct)
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s VerifiedInfo) NewTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(2, ss)
	return ss, err
}

func (s VerifiedInfo) Version() uint32 {
	return s.Struct.Uint32(0)
}

func (s VerifiedInfo) SetVersion(v uint32) {

	s.Struct.SetUint32(0, v)
}

func (s VerifiedInfo) MarketingVersion() (util.LocalizedText, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return util.LocalizedText{}, err
	}

	ss := capnp.ToStruct(p)

	return util.LocalizedText{Struct: ss}, nil
}

func (s VerifiedInfo) SetMarketingVersion(v util.LocalizedText) error {

	return s.Struct.SetPointer(3, v.Struct)
}

// NewMarketingVersion sets the marketingVersion field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s VerifiedInfo) NewMarketingVersion() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPointer(3, ss)
	return ss, err
}

func (s VerifiedInfo) AuthorPgpKeyFingerprint() (string, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s VerifiedInfo) SetAuthorPgpKeyFingerprint(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(4, t)
}

func (s VerifiedInfo) Metadata() (Metadata, error) {
	p, err := s.Struct.Pointer(5)
	if err != nil {
		return Metadata{}, err
	}

	ss := capnp.ToStruct(p)

	return Metadata{Struct: ss}, nil
}

func (s VerifiedInfo) SetMetadata(v Metadata) error {

	return s.Struct.SetPointer(5, v.Struct)
}

// NewMetadata sets the metadata field to a newly
// allocated Metadata struct, preferring placement in s's segment.
func (s VerifiedInfo) NewMetadata() (Metadata, error) {

	ss, err := NewMetadata(s.Struct.Segment())
	if err != nil {
		return Metadata{}, err
	}
	err = s.Struct.SetPointer(5, ss)
	return ss, err
}

// VerifiedInfo_List is a list of VerifiedInfo.
type VerifiedInfo_List struct{ capnp.List }

// NewVerifiedInfo creates a new list of VerifiedInfo.
func NewVerifiedInfo_List(s *capnp.Segment, sz int32) (VerifiedInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	if err != nil {
		return VerifiedInfo_List{}, err
	}
	return VerifiedInfo_List{l}, nil
}

func (s VerifiedInfo_List) At(i int) VerifiedInfo           { return VerifiedInfo{s.List.Struct(i)} }
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
	if err != nil {
		return CategoryInfo{}, err
	}
	return CategoryInfo{st}, nil
}

func NewRootCategoryInfo(s *capnp.Segment) (CategoryInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return CategoryInfo{}, err
	}
	return CategoryInfo{st}, nil
}

func ReadRootCategoryInfo(msg *capnp.Message) (CategoryInfo, error) {
	root, err := msg.Root()
	if err != nil {
		return CategoryInfo{}, err
	}
	st := capnp.ToStruct(root)
	return CategoryInfo{st}, nil
}

func (s CategoryInfo) Title() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s CategoryInfo) SetTitle(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// CategoryInfo_List is a list of CategoryInfo.
type CategoryInfo_List struct{ capnp.List }

// NewCategoryInfo creates a new list of CategoryInfo.
func NewCategoryInfo_List(s *capnp.Segment, sz int32) (CategoryInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return CategoryInfo_List{}, err
	}
	return CategoryInfo_List{l}, nil
}

func (s CategoryInfo_List) At(i int) CategoryInfo           { return CategoryInfo{s.List.Struct(i)} }
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
	if err != nil {
		return Category_List{}, err
	}
	return Category_List{l.List}, nil
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
	root, err := msg.Root()
	if err != nil {
		return KeyFile{}, err
	}
	st := capnp.ToStruct(root)
	return KeyFile{st}, nil
}

func (s KeyFile) PublicKey() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s KeyFile) SetPublicKey(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s KeyFile) PrivateKey() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s KeyFile) SetPrivateKey(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
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
	root, err := msg.Root()
	if err != nil {
		return Signature{}, err
	}
	st := capnp.ToStruct(root)
	return Signature{st}, nil
}

func (s Signature) PublicKey() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Signature) SetPublicKey(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s Signature) Signature() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Signature) SetSignature(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
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
	root, err := msg.Root()
	if err != nil {
		return Archive{}, err
	}
	st := capnp.ToStruct(root)
	return Archive{st}, nil
}

func (s Archive) Files() (Archive_File_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Archive_File_List{}, err
	}

	l := capnp.ToList(p)

	return Archive_File_List{List: l}, nil
}

func (s Archive) SetFiles(v Archive_File_List) error {

	return s.Struct.SetPointer(0, v.List)
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
	root, err := msg.Root()
	if err != nil {
		return Archive_File{}, err
	}
	st := capnp.ToStruct(root)
	return Archive_File{st}, nil
}

func (s Archive_File) Which() Archive_File_Which {
	return Archive_File_Which(s.Struct.Uint16(0))
}

func (s Archive_File) Name() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Archive_File) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Archive_File) LastModificationTimeNs() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Archive_File) SetLastModificationTimeNs(v int64) {

	s.Struct.SetUint64(8, uint64(v))
}

func (s Archive_File) Regular() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Archive_File) SetRegular(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
}

func (s Archive_File) Executable() ([]byte, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s Archive_File) SetExecutable(v []byte) error {
	s.Struct.SetUint16(0, 1)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, d)
}

func (s Archive_File) Symlink() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Archive_File) SetSymlink(v string) error {
	s.Struct.SetUint16(0, 2)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s Archive_File) Directory() (Archive_File_List, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Archive_File_List{}, err
	}

	l := capnp.ToList(p)

	return Archive_File_List{List: l}, nil
}

func (s Archive_File) SetDirectory(v Archive_File_List) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPointer(1, v.List)
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
