package powerbox

// AUTO GENERATED - DO NOT EDIT

import (
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type PowerboxDescriptor struct{ capnp.Struct }

func NewPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor{st}, err
}

func NewRootPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor{st}, err
}

func ReadRootPowerboxDescriptor(msg *capnp.Message) (PowerboxDescriptor, error) {
	root, err := msg.RootPtr()
	return PowerboxDescriptor{root.Struct()}, err
}

func (s PowerboxDescriptor) String() string {
	str, _ := text.Marshal(0xedcf0fa3bfc71c58, s.Struct)
	return str
}

func (s PowerboxDescriptor) Tags() (PowerboxDescriptor_Tag_List, error) {
	p, err := s.Struct.Ptr(0)
	return PowerboxDescriptor_Tag_List{List: p.List()}, err
}

func (s PowerboxDescriptor) HasTags() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor) SetTags(v PowerboxDescriptor_Tag_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTags sets the tags field to a newly
// allocated PowerboxDescriptor_Tag_List, preferring placement in s's segment.
func (s PowerboxDescriptor) NewTags(n int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := NewPowerboxDescriptor_Tag_List(s.Struct.Segment(), n)
	if err != nil {
		return PowerboxDescriptor_Tag_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s PowerboxDescriptor) Quality() PowerboxDescriptor_MatchQuality {
	return PowerboxDescriptor_MatchQuality(s.Struct.Uint16(0))
}

func (s PowerboxDescriptor) SetQuality(v PowerboxDescriptor_MatchQuality) {
	s.Struct.SetUint16(0, uint16(v))
}

// PowerboxDescriptor_List is a list of PowerboxDescriptor.
type PowerboxDescriptor_List struct{ capnp.List }

// NewPowerboxDescriptor creates a new list of PowerboxDescriptor.
func NewPowerboxDescriptor_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PowerboxDescriptor_List{l}, err
}

func (s PowerboxDescriptor_List) At(i int) PowerboxDescriptor {
	return PowerboxDescriptor{s.List.Struct(i)}
}

func (s PowerboxDescriptor_List) Set(i int, v PowerboxDescriptor) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Promise is a wrapper for a PowerboxDescriptor promised by a client call.
type PowerboxDescriptor_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Promise) Struct() (PowerboxDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor{s}, err
}

type PowerboxDescriptor_Tag struct{ capnp.Struct }

func NewPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag{st}, err
}

func NewRootPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag{st}, err
}

func ReadRootPowerboxDescriptor_Tag(msg *capnp.Message) (PowerboxDescriptor_Tag, error) {
	root, err := msg.RootPtr()
	return PowerboxDescriptor_Tag{root.Struct()}, err
}

func (s PowerboxDescriptor_Tag) String() string {
	str, _ := text.Marshal(0xbe3d16a8df03c418, s.Struct)
	return str
}

func (s PowerboxDescriptor_Tag) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s PowerboxDescriptor_Tag) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s PowerboxDescriptor_Tag) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s PowerboxDescriptor_Tag) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor_Tag) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s PowerboxDescriptor_Tag) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s PowerboxDescriptor_Tag) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// PowerboxDescriptor_Tag_List is a list of PowerboxDescriptor_Tag.
type PowerboxDescriptor_Tag_List struct{ capnp.List }

// NewPowerboxDescriptor_Tag creates a new list of PowerboxDescriptor_Tag.
func NewPowerboxDescriptor_Tag_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PowerboxDescriptor_Tag_List{l}, err
}

func (s PowerboxDescriptor_Tag_List) At(i int) PowerboxDescriptor_Tag {
	return PowerboxDescriptor_Tag{s.List.Struct(i)}
}

func (s PowerboxDescriptor_Tag_List) Set(i int, v PowerboxDescriptor_Tag) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Tag_Promise is a wrapper for a PowerboxDescriptor_Tag promised by a client call.
type PowerboxDescriptor_Tag_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Tag_Promise) Struct() (PowerboxDescriptor_Tag, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor_Tag{s}, err
}

func (p PowerboxDescriptor_Tag_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type PowerboxDescriptor_MatchQuality uint16

// Values of PowerboxDescriptor_MatchQuality.
const (
	PowerboxDescriptor_MatchQuality_preferred    PowerboxDescriptor_MatchQuality = 1
	PowerboxDescriptor_MatchQuality_acceptable   PowerboxDescriptor_MatchQuality = 0
	PowerboxDescriptor_MatchQuality_unacceptable PowerboxDescriptor_MatchQuality = 2
)

// String returns the enum's constant name.
func (c PowerboxDescriptor_MatchQuality) String() string {
	switch c {
	case PowerboxDescriptor_MatchQuality_preferred:
		return "preferred"
	case PowerboxDescriptor_MatchQuality_acceptable:
		return "acceptable"
	case PowerboxDescriptor_MatchQuality_unacceptable:
		return "unacceptable"

	default:
		return ""
	}
}

// PowerboxDescriptor_MatchQualityFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PowerboxDescriptor_MatchQualityFromString(c string) PowerboxDescriptor_MatchQuality {
	switch c {
	case "preferred":
		return PowerboxDescriptor_MatchQuality_preferred
	case "acceptable":
		return PowerboxDescriptor_MatchQuality_acceptable
	case "unacceptable":
		return PowerboxDescriptor_MatchQuality_unacceptable

	default:
		return 0
	}
}

type PowerboxDescriptor_MatchQuality_List struct{ capnp.List }

func NewPowerboxDescriptor_MatchQuality_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_MatchQuality_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return PowerboxDescriptor_MatchQuality_List{l.List}, err
}

func (l PowerboxDescriptor_MatchQuality_List) At(i int) PowerboxDescriptor_MatchQuality {
	ul := capnp.UInt16List{List: l.List}
	return PowerboxDescriptor_MatchQuality(ul.At(i))
}

func (l PowerboxDescriptor_MatchQuality_List) Set(i int, v PowerboxDescriptor_MatchQuality) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type PowerboxDisplayInfo struct{ capnp.Struct }

func NewPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo{st}, err
}

func NewRootPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo{st}, err
}

func ReadRootPowerboxDisplayInfo(msg *capnp.Message) (PowerboxDisplayInfo, error) {
	root, err := msg.RootPtr()
	return PowerboxDisplayInfo{root.Struct()}, err
}

func (s PowerboxDisplayInfo) String() string {
	str, _ := text.Marshal(0xa553a209bee32bec, s.Struct)
	return str
}

func (s PowerboxDisplayInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetVerbPhrase(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewVerbPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// PowerboxDisplayInfo_List is a list of PowerboxDisplayInfo.
type PowerboxDisplayInfo_List struct{ capnp.List }

// NewPowerboxDisplayInfo creates a new list of PowerboxDisplayInfo.
func NewPowerboxDisplayInfo_List(s *capnp.Segment, sz int32) (PowerboxDisplayInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return PowerboxDisplayInfo_List{l}, err
}

func (s PowerboxDisplayInfo_List) At(i int) PowerboxDisplayInfo {
	return PowerboxDisplayInfo{s.List.Struct(i)}
}

func (s PowerboxDisplayInfo_List) Set(i int, v PowerboxDisplayInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDisplayInfo_Promise is a wrapper for a PowerboxDisplayInfo promised by a client call.
type PowerboxDisplayInfo_Promise struct{ *capnp.Pipeline }

func (p PowerboxDisplayInfo_Promise) Struct() (PowerboxDisplayInfo, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDisplayInfo{s}, err
}

func (p PowerboxDisplayInfo_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PowerboxDisplayInfo_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p PowerboxDisplayInfo_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

const schema_f6c200ab14cd53e4 = "x\xdat\x92Mk\x13_\x18\xc5\x9fso\xf2O\x03" +
	"\x7f\x9dN&\x8a\x06%\x10\x14\xdf0\xf4\x15E(\x06" +
	"i\xc1\x80b\xc6v!\xa5.n\x92i;%d\xc6" +
	"\x9bIkt\xd1\x0f\xe0\xca\xaf\xa0\x08\x82\x0b\x97.\x84" +
	"\x0a\"\xb8\x12\xd1\x0f`\xf1\x0d\\\xe8\xd6\x85\x9b\xf1\x99" +
	"d\xda\x0c\xd2.\x02\xb9\xe79\xb9\xe7<\xbf\xdc\xb1i" +
	"T\xc4xz3EdO\xa5\xff\x0b\x7f\x9e\xfb\xba\x95" +
	"}4\xff\x84L\x03\xdb\xdf\xe6\xdf\xe5\x9f\xbd\xfe\x9d\x96" +
	"\x19\xa2\xc9\x93\xa2\x00kZ\xf0Wk\\<'\x84\x9f" +
	"~,}\x9c\xfbSxIfI\x84\xb7\x8e\xbd}\xf5" +
	"\xd8x\xff\x8b\x08\x93\x1f\xc4=X\xdf\xfb\xce\xcf\xe2\x06" +
	";\x8f\xbc\x91\xdbO\x0f\xcfl\x91]\x02\x86\xd64\xa2" +
	"{!'`\x99Q\x84u@nPbn\x1b\x18v" +
	"\xe8{]\x99\x83\xd5\x93\xa7\xd8\xfb\x90\xbd\xcd\xd0\xf76" +
	"\x1c]\xf7\xee\xcarC\xf9m\xffR->\xcf\xba\x1d" +
	"\xbf\xa5z\xd5\xb6\\\xf6j\x80\xfd\xbf\xe4\x05S 2" +
	"\xe7&x\xd3\x8a\x84}M\xc0\x04\xf2\x88\xc4\xea\"\x8b" +
	"WY\\`Q\x88<\x04\x8bv\x9d\xc5\x1a\x8bK\x02" +
	"\xc5\xc0\x0dZ\x0eF\xc3\xea\x97\xfb\x17\xf4\x8b\xdb\x0fx" +
	"O\x8cr\xd9u\x8e\xab\xadjE\xb2\xb3\xd7\xb8\xe9t" +
	"\x1a\xda\xf5\x03\xca\xb8^{\x8f\xf9N\xff\xd4\xbf\xfd\xe3" +
	"\xdfy\xba|]\x05\x8dU\xfbrW\xb5\xdc\xa0\xd7\xdf" +
	"\x05\x020\x8f/\x12\xcc\xa37\xf9\"a\x1eZ#\x0a" +
	"U\xa3\xe1\xf8\x81\xaa\x93l9\xa1\xaf\x9deGk\x87" +
	"\xd0\x0c\xbb\xed\xc1\x88\x0cU\x8fF\xfb!\xdb\x89,\xea" +
	"\xf2\x82Z\x89\x92Fv\xa9\x9d)0\x8b\x13\xccb\x8c" +
	"\xb3ch\xe7#\x92\xa7Y\x9b\x12\x90n\x13Y\x12\xfc" +
	"Aq]\xb5\xba\x0er|\xca%6\xdc7\x0e\xda\x1e" +
	"A\xe2\x8d\x98\xd9R\xe2i\xa5\xd72\\&\x1c@\xe8" +
	"*2\"\x0a\xc9bg\xe3b\x95\xb8\x18\xa3\x99\xb9\xc2" +
	"\xdaE\xd6f\x05\x8c@\xadtp\x90\xc0\x7f$\xf3\xdf" +
	"Ma\xfe,n\xde\x19`\x851\x8c\xe4\x89A\xf8\x1b" +
	"\x00\x00\xff\xff\xae&\xda\xdd"

func init() {
	schemas.Register(schema_f6c200ab14cd53e4,
		0xa553a209bee32bec,
		0xbb1afa45d25ce8de,
		0xbe3d16a8df03c418,
		0xedcf0fa3bfc71c58)
}
