package featurekey

// AUTO GENERATED - DO NOT EDIT

import (
	updatetool "zenhack.net/go/sandstorm/capnp/updatetool"
	capnp "zombiezen.com/go/capnproto2"
)

var (
	FeatureKey_signingKey = updatetool.PublicSigningKey{Struct: capnp.ToStruct(capnp.MustUnmarshalRoot(x_d24581e9cd6a6772[0:48]))}
)

type FeatureKey struct{ capnp.Struct }
type FeatureKey_features FeatureKey

func NewFeatureKey(s *capnp.Segment) (FeatureKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return FeatureKey{}, err
	}
	return FeatureKey{st}, nil
}

func NewRootFeatureKey(s *capnp.Segment) (FeatureKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	if err != nil {
		return FeatureKey{}, err
	}
	return FeatureKey{st}, nil
}

func ReadRootFeatureKey(msg *capnp.Message) (FeatureKey, error) {
	root, err := msg.Root()
	if err != nil {
		return FeatureKey{}, err
	}
	st := capnp.ToStruct(root)
	return FeatureKey{st}, nil
}

func (s FeatureKey) Secret() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s FeatureKey) SetSecret(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s FeatureKey) Customer() (FeatureKey_Customer, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return FeatureKey_Customer{}, err
	}

	ss := capnp.ToStruct(p)

	return FeatureKey_Customer{Struct: ss}, nil
}

func (s FeatureKey) SetCustomer(v FeatureKey_Customer) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewCustomer sets the customer field to a newly
// allocated FeatureKey_Customer struct, preferring placement in s's segment.
func (s FeatureKey) NewCustomer() (FeatureKey_Customer, error) {

	ss, err := NewFeatureKey_Customer(s.Struct.Segment())
	if err != nil {
		return FeatureKey_Customer{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s FeatureKey) Issued() uint64 {
	return s.Struct.Uint64(0)
}

func (s FeatureKey) SetIssued(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s FeatureKey) Expires() uint64 {
	return s.Struct.Uint64(8)
}

func (s FeatureKey) SetExpires(v uint64) {

	s.Struct.SetUint64(8, v)
}

func (s FeatureKey) UserLimit() uint32 {
	return s.Struct.Uint32(16)
}

func (s FeatureKey) SetUserLimit(v uint32) {

	s.Struct.SetUint32(16, v)
}

func (s FeatureKey) IsElasticBilling() bool {
	return s.Struct.Bit(160)
}

func (s FeatureKey) SetIsElasticBilling(v bool) {

	s.Struct.SetBit(160, v)
}

func (s FeatureKey) IsTrial() bool {
	return s.Struct.Bit(161)
}

func (s FeatureKey) SetIsTrial(v bool) {

	s.Struct.SetBit(161, v)
}

func (s FeatureKey) IsForTesting() bool {
	return s.Struct.Bit(165)
}

func (s FeatureKey) SetIsForTesting(v bool) {

	s.Struct.SetBit(165, v)
}
func (s FeatureKey) Features() FeatureKey_features { return FeatureKey_features(s) }

func (s FeatureKey_features) Ldap() bool {
	return !s.Struct.Bit(162)
}

func (s FeatureKey_features) SetLdap(v bool) {

	s.Struct.SetBit(162, !v)
}

func (s FeatureKey_features) Saml() bool {
	return !s.Struct.Bit(163)
}

func (s FeatureKey_features) SetSaml(v bool) {

	s.Struct.SetBit(163, !v)
}

func (s FeatureKey_features) OrgManagement() bool {
	return !s.Struct.Bit(164)
}

func (s FeatureKey_features) SetOrgManagement(v bool) {

	s.Struct.SetBit(164, !v)
}

// FeatureKey_List is a list of FeatureKey.
type FeatureKey_List struct{ capnp.List }

// NewFeatureKey creates a new list of FeatureKey.
func NewFeatureKey_List(s *capnp.Segment, sz int32) (FeatureKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	if err != nil {
		return FeatureKey_List{}, err
	}
	return FeatureKey_List{l}, nil
}

func (s FeatureKey_List) At(i int) FeatureKey           { return FeatureKey{s.List.Struct(i)} }
func (s FeatureKey_List) Set(i int, v FeatureKey) error { return s.List.SetStruct(i, v.Struct) }

// FeatureKey_Promise is a wrapper for a FeatureKey promised by a client call.
type FeatureKey_Promise struct{ *capnp.Pipeline }

func (p FeatureKey_Promise) Struct() (FeatureKey, error) {
	s, err := p.Pipeline.Struct()
	return FeatureKey{s}, err
}

func (p FeatureKey_Promise) Customer() FeatureKey_Customer_Promise {
	return FeatureKey_Customer_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}
func (p FeatureKey_Promise) Features() FeatureKey_features_Promise {
	return FeatureKey_features_Promise{p.Pipeline}
}

// FeatureKey_features_Promise is a wrapper for a FeatureKey_features promised by a client call.
type FeatureKey_features_Promise struct{ *capnp.Pipeline }

func (p FeatureKey_features_Promise) Struct() (FeatureKey_features, error) {
	s, err := p.Pipeline.Struct()
	return FeatureKey_features{s}, err
}

type FeatureKey_Customer struct{ capnp.Struct }

func NewFeatureKey_Customer(s *capnp.Segment) (FeatureKey_Customer, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return FeatureKey_Customer{}, err
	}
	return FeatureKey_Customer{st}, nil
}

func NewRootFeatureKey_Customer(s *capnp.Segment) (FeatureKey_Customer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return FeatureKey_Customer{}, err
	}
	return FeatureKey_Customer{st}, nil
}

func ReadRootFeatureKey_Customer(msg *capnp.Message) (FeatureKey_Customer, error) {
	root, err := msg.Root()
	if err != nil {
		return FeatureKey_Customer{}, err
	}
	st := capnp.ToStruct(root)
	return FeatureKey_Customer{st}, nil
}

func (s FeatureKey_Customer) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s FeatureKey_Customer) SetId(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s FeatureKey_Customer) OrganizationName() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s FeatureKey_Customer) SetOrganizationName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s FeatureKey_Customer) ContactName() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s FeatureKey_Customer) SetContactName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s FeatureKey_Customer) ContactEmail() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s FeatureKey_Customer) SetContactEmail(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

// FeatureKey_Customer_List is a list of FeatureKey_Customer.
type FeatureKey_Customer_List struct{ capnp.List }

// NewFeatureKey_Customer creates a new list of FeatureKey_Customer.
func NewFeatureKey_Customer_List(s *capnp.Segment, sz int32) (FeatureKey_Customer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return FeatureKey_Customer_List{}, err
	}
	return FeatureKey_Customer_List{l}, nil
}

func (s FeatureKey_Customer_List) At(i int) FeatureKey_Customer {
	return FeatureKey_Customer{s.List.Struct(i)}
}
func (s FeatureKey_Customer_List) Set(i int, v FeatureKey_Customer) error {
	return s.List.SetStruct(i, v.Struct)
}

// FeatureKey_Customer_Promise is a wrapper for a FeatureKey_Customer promised by a client call.
type FeatureKey_Customer_Promise struct{ *capnp.Pipeline }

func (p FeatureKey_Customer_Promise) Struct() (FeatureKey_Customer, error) {
	s, err := p.Pipeline.Struct()
	return FeatureKey_Customer{s}, err
}

var x_d24581e9cd6a6772 = []byte{
	0, 0, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 4, 0, 0, 0,
	54, 80, 246, 217, 181, 168, 173, 134,
	35, 195, 170, 8, 186, 9, 57, 24,
	13, 86, 201, 83, 164, 141, 119, 109,
	168, 126, 58, 243, 50, 245, 148, 223,
}
