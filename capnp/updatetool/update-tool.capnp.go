package updatetool

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in update-tool.capnp.
var (
	UpdatePublicKeys = PublicSigningKey_List{List: capnp.MustUnmarshalRootPtr(x_96c3fff3f4beb8fe[0:56]).List()}
)

type PublicSigningKey struct{ capnp.Struct }

func NewPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey{st}, err
}

func NewRootPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey{st}, err
}

func ReadRootPublicSigningKey(msg *capnp.Message) (PublicSigningKey, error) {
	root, err := msg.RootPtr()
	return PublicSigningKey{root.Struct()}, err
}

func (s PublicSigningKey) String() string {
	str, _ := text.Marshal(0x9b54bbec5de99f09, s.Struct)
	return str
}

func (s PublicSigningKey) Key0() uint64 {
	return s.Struct.Uint64(0)
}

func (s PublicSigningKey) SetKey0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s PublicSigningKey) Key1() uint64 {
	return s.Struct.Uint64(8)
}

func (s PublicSigningKey) SetKey1(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s PublicSigningKey) Key2() uint64 {
	return s.Struct.Uint64(16)
}

func (s PublicSigningKey) SetKey2(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s PublicSigningKey) Key3() uint64 {
	return s.Struct.Uint64(24)
}

func (s PublicSigningKey) SetKey3(v uint64) {
	s.Struct.SetUint64(24, v)
}

// PublicSigningKey_List is a list of PublicSigningKey.
type PublicSigningKey_List struct{ capnp.List }

// NewPublicSigningKey creates a new list of PublicSigningKey.
func NewPublicSigningKey_List(s *capnp.Segment, sz int32) (PublicSigningKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0}, sz)
	return PublicSigningKey_List{l}, err
}

func (s PublicSigningKey_List) At(i int) PublicSigningKey { return PublicSigningKey{s.List.Struct(i)} }

func (s PublicSigningKey_List) Set(i int, v PublicSigningKey) error {
	return s.List.SetStruct(i, v.Struct)
}

// PublicSigningKey_Promise is a wrapper for a PublicSigningKey promised by a client call.
type PublicSigningKey_Promise struct{ *capnp.Pipeline }

func (p PublicSigningKey_Promise) Struct() (PublicSigningKey, error) {
	s, err := p.Pipeline.Struct()
	return PublicSigningKey{s}, err
}

type Signature struct{ capnp.Struct }

func NewSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature{st}, err
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature{st}, err
}

func ReadRootSignature(msg *capnp.Message) (Signature, error) {
	root, err := msg.RootPtr()
	return Signature{root.Struct()}, err
}

func (s Signature) String() string {
	str, _ := text.Marshal(0xc4d0558d33210637, s.Struct)
	return str
}

func (s Signature) Sig0() uint64 {
	return s.Struct.Uint64(0)
}

func (s Signature) SetSig0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Signature) Sig1() uint64 {
	return s.Struct.Uint64(8)
}

func (s Signature) SetSig1(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Signature) Sig2() uint64 {
	return s.Struct.Uint64(16)
}

func (s Signature) SetSig2(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Signature) Sig3() uint64 {
	return s.Struct.Uint64(24)
}

func (s Signature) SetSig3(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s Signature) Sig4() uint64 {
	return s.Struct.Uint64(32)
}

func (s Signature) SetSig4(v uint64) {
	s.Struct.SetUint64(32, v)
}

func (s Signature) Sig5() uint64 {
	return s.Struct.Uint64(40)
}

func (s Signature) SetSig5(v uint64) {
	s.Struct.SetUint64(40, v)
}

func (s Signature) Sig6() uint64 {
	return s.Struct.Uint64(48)
}

func (s Signature) SetSig6(v uint64) {
	s.Struct.SetUint64(48, v)
}

func (s Signature) Sig7() uint64 {
	return s.Struct.Uint64(56)
}

func (s Signature) SetSig7(v uint64) {
	s.Struct.SetUint64(56, v)
}

// Signature_List is a list of Signature.
type Signature_List struct{ capnp.List }

// NewSignature creates a new list of Signature.
func NewSignature_List(s *capnp.Segment, sz int32) (Signature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0}, sz)
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

type UpdateSignature struct{ capnp.Struct }

func NewUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature{st}, err
}

func NewRootUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature{st}, err
}

func ReadRootUpdateSignature(msg *capnp.Message) (UpdateSignature, error) {
	root, err := msg.RootPtr()
	return UpdateSignature{root.Struct()}, err
}

func (s UpdateSignature) String() string {
	str, _ := text.Marshal(0x9c805f76ef46e6c0, s.Struct)
	return str
}

func (s UpdateSignature) Signatures() (Signature_List, error) {
	p, err := s.Struct.Ptr(0)
	return Signature_List{List: p.List()}, err
}

func (s UpdateSignature) HasSignatures() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UpdateSignature) SetSignatures(v Signature_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSignatures sets the signatures field to a newly
// allocated Signature_List, preferring placement in s's segment.
func (s UpdateSignature) NewSignatures(n int32) (Signature_List, error) {
	l, err := NewSignature_List(s.Struct.Segment(), n)
	if err != nil {
		return Signature_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// UpdateSignature_List is a list of UpdateSignature.
type UpdateSignature_List struct{ capnp.List }

// NewUpdateSignature creates a new list of UpdateSignature.
func NewUpdateSignature_List(s *capnp.Segment, sz int32) (UpdateSignature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UpdateSignature_List{l}, err
}

func (s UpdateSignature_List) At(i int) UpdateSignature { return UpdateSignature{s.List.Struct(i)} }

func (s UpdateSignature_List) Set(i int, v UpdateSignature) error {
	return s.List.SetStruct(i, v.Struct)
}

// UpdateSignature_Promise is a wrapper for a UpdateSignature promised by a client call.
type UpdateSignature_Promise struct{ *capnp.Pipeline }

func (p UpdateSignature_Promise) Struct() (UpdateSignature, error) {
	s, err := p.Pipeline.Struct()
	return UpdateSignature{s}, err
}

const schema_96c3fff3f4beb8fe = "x\xda|\xd3Mh\x13Q\x10\x07\xf0\x99\xf7\xdenZ" +
	"\xf0\xa3\x9b\xdd\x8b\x9e\xea\xc9\"T\x9a\xaeZ\xf0\xe2\xc7" +
	"\xc1K/\x09\xb1\x97\xa2\xd8\x98\xacm0\xa6K>," +
	"A\x10\x8bW#x\x11\xc5/\xbc\x08\x1e<\x08\x1e\x14" +
	"KQ,\x88P\xb5\x82B\xc5\x8a\xb1\xb4\xdaB\x8b\xa2" +
	"\xbd\x09\xae3\x89\xee>A<\x04\xf2\xff1\x997o" +
	"\xb2\xdb\x93\xc3\xbd\"a\x8cI\x80T\x97a\x06\xed7" +
	"\x96\x0f\xaf<<x\x05RqT\xc1\xcf\xfb\x93k\xdf" +
	"\x83'\x17A\xc5\x00\xdc\x9b\xb8\x19\xed\xbbH_\xed;" +
	"\xf8\x190x\xf4\xe9\xc0\x97\x93G\xce\\\x05+\x8eQ" +
	"\xad\xc1\x15\xee\xb8\x88\xa3}Apq]\xec\xa1\xe2>" +
	"s\x8b[\x1f\x98\x99\xe2\xcem\x7fw\xb6\x1f\x8by\xfb" +
	"9\xd7\xba\xcf\xc4S\xa4\xe2\x1f\xe7\x87\x16':\x1f|" +
	"\xa3\xce\xda\x14\x80n]\xd1\x0c\xd7\xf8G\xe9KJb" +
	"\xfa\x9e\x12\x08\xfb\x83\xaa\x9f\xcbT\xbc\xee\x8a\x1c\x1d-" +
	"l\xcff\xfc\xa2\xbf;Y=Z\xc8g\xd3\xf9\xe1b" +
	"\xbe8\xdc/\xbdZ\x121\xd5!\x15\x80B\x00+\xb3" +
	"\x8dn|HbjD\xa0\x85\xe8 \xa3\xc78DX" +
	" \x14\xc2AA\x98g\xcc\x11\xfa\x84R:H\xbb\xb2" +
	"N0\x8e\x10V\x04n<\xee\xd5z\xb0\x1d\x04}\x9a" +
	"!\xa1\x87^=\xb8\x7f\xc2\xbf&\x1eh\x12O\x9c\xa9" +
	"TK\xe8\xf1\xc0*\x1cx\xfd \x9d\xb8\x8eN\xec\x12" +
	"\x18\x94\x7f\x17\x81\xf4\xca\xb8\x010)\x11;\xa2\x15\x03" +
	"2\x86g\x88\xe8\x8cV\xf7X\xb5\xd4\xec\xde\x19v\x7f" +
	"\xc5\xf7\x99\xa6\xee\xb3\xda:\xde0\xce\x10\xcei\xebx" +
	"\xcb\xf8\x9a\xb0\xa1\xad\xe3=\xe3,\xe1\x02\xa1R\x0eR" +
	"_\xeb#\xe3\x1c\xe1\x12\xa1a8h\x10.26\x08" +
	"W\x08M\xd3A\x93p\x99q\x81\xf0+a,\xe6\xf0" +
	"\x13d\xad2.\x11\xae\xd1\x8a\xe9\xc2\xd1\x8a)$\xf4" +
	"\xd0\xab\x07W\x0f;\xf4\xb0S\x0f\xbb\xf4\xd0\xf7\xbf\xbf" +
	"\xa5E\xad\xc7\xa9\xdf\x93\xb52m.\xday\xf8\xc24" +
	"w\x0e\x16n\xa5\xfb\x07c\xb7O\x95\xae_\xee\x1e\x94" +
	"\xd9\x97\xe7\xfc\xc6\xad\x17g'6MN\xadN\xfb\x1f" +
	"\x8e\x95\xf7\xbd\x9b?\xdd6\xfe+\x00\x00\xff\xffZ\x94" +
	"\xdd\x02"

func init() {
	schemas.Register(schema_96c3fff3f4beb8fe,
		0x9b54bbec5de99f09,
		0x9c805f76ef46e6c0,
		0xc4d0558d33210637,
		0xf2b920bce5608efb)
}

var x_96c3fff3f4beb8fe = []byte{
	0, 0, 0, 0, 6, 0, 0, 0,
	1, 0, 0, 0, 39, 0, 0, 0,
	4, 0, 0, 0, 4, 0, 0, 0,
	119, 169, 123, 114, 158, 153, 45, 90,
	99, 207, 140, 112, 224, 166, 206, 131,
	188, 25, 190, 196, 237, 204, 112, 223,
	102, 115, 65, 219, 226, 126, 8, 129,
}
