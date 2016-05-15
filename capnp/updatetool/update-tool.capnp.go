package updatetool

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
)

var (
	UpdatePublicKeys = PublicSigningKey_List{List: capnp.ToList(capnp.MustUnmarshalRoot(x_96c3fff3f4beb8fe[0:56]))}
)

type PublicSigningKey struct{ capnp.Struct }

func NewPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	if err != nil {
		return PublicSigningKey{}, err
	}
	return PublicSigningKey{st}, nil
}

func NewRootPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	if err != nil {
		return PublicSigningKey{}, err
	}
	return PublicSigningKey{st}, nil
}

func ReadRootPublicSigningKey(msg *capnp.Message) (PublicSigningKey, error) {
	root, err := msg.Root()
	if err != nil {
		return PublicSigningKey{}, err
	}
	st := capnp.ToStruct(root)
	return PublicSigningKey{st}, nil
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
	if err != nil {
		return PublicSigningKey_List{}, err
	}
	return PublicSigningKey_List{l}, nil
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
	if err != nil {
		return Signature{}, err
	}
	return Signature{st}, nil
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
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

type UpdateSignature struct{ capnp.Struct }

func NewUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UpdateSignature{}, err
	}
	return UpdateSignature{st}, nil
}

func NewRootUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UpdateSignature{}, err
	}
	return UpdateSignature{st}, nil
}

func ReadRootUpdateSignature(msg *capnp.Message) (UpdateSignature, error) {
	root, err := msg.Root()
	if err != nil {
		return UpdateSignature{}, err
	}
	st := capnp.ToStruct(root)
	return UpdateSignature{st}, nil
}

func (s UpdateSignature) Signatures() (Signature_List, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return Signature_List{}, err
	}

	l := capnp.ToList(p)

	return Signature_List{List: l}, nil
}

func (s UpdateSignature) SetSignatures(v Signature_List) error {

	return s.Struct.SetPointer(0, v.List)
}

// UpdateSignature_List is a list of UpdateSignature.
type UpdateSignature_List struct{ capnp.List }

// NewUpdateSignature creates a new list of UpdateSignature.
func NewUpdateSignature_List(s *capnp.Segment, sz int32) (UpdateSignature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return UpdateSignature_List{}, err
	}
	return UpdateSignature_List{l}, nil
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

var x_96c3fff3f4beb8fe = []byte{
	0, 0, 0, 0, 6, 0, 0, 0,
	1, 0, 0, 0, 39, 0, 0, 0,
	4, 0, 0, 0, 4, 0, 0, 0,
	119, 169, 123, 114, 158, 153, 45, 90,
	99, 207, 140, 112, 224, 166, 206, 131,
	188, 25, 190, 196, 237, 204, 112, 223,
	102, 115, 65, 219, 226, 126, 8, 129,
}
