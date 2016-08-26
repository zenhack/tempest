package featurekey

// AUTO GENERATED - DO NOT EDIT

import (
	updatetool "zenhack.net/go/sandstorm/capnp/updatetool"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in feature-key.capnp.
var (
	FeatureKey_signingKey = updatetool.PublicSigningKey{Struct: capnp.MustUnmarshalRootPtr(x_d24581e9cd6a6772[0:48]).Struct()}
)

type FeatureKey struct{ capnp.Struct }
type FeatureKey_features FeatureKey

func NewFeatureKey(s *capnp.Segment) (FeatureKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return FeatureKey{st}, err
}

func NewRootFeatureKey(s *capnp.Segment) (FeatureKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return FeatureKey{st}, err
}

func ReadRootFeatureKey(msg *capnp.Message) (FeatureKey, error) {
	root, err := msg.RootPtr()
	return FeatureKey{root.Struct()}, err
}

func (s FeatureKey) String() string {
	str, _ := text.Marshal(0x9548af5e959db169, s.Struct)
	return str
}

func (s FeatureKey) Secret() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s FeatureKey) HasSecret() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s FeatureKey) SetSecret(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s FeatureKey) Customer() (FeatureKey_Customer, error) {
	p, err := s.Struct.Ptr(1)
	return FeatureKey_Customer{Struct: p.Struct()}, err
}

func (s FeatureKey) HasCustomer() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s FeatureKey) SetCustomer(v FeatureKey_Customer) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewCustomer sets the customer field to a newly
// allocated FeatureKey_Customer struct, preferring placement in s's segment.
func (s FeatureKey) NewCustomer() (FeatureKey_Customer, error) {
	ss, err := NewFeatureKey_Customer(s.Struct.Segment())
	if err != nil {
		return FeatureKey_Customer{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
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

func (s FeatureKey) IsFreeKey() bool {
	return s.Struct.Bit(166)
}

func (s FeatureKey) SetIsFreeKey(v bool) {
	s.Struct.SetBit(166, v)
}

// FeatureKey_List is a list of FeatureKey.
type FeatureKey_List struct{ capnp.List }

// NewFeatureKey creates a new list of FeatureKey.
func NewFeatureKey_List(s *capnp.Segment, sz int32) (FeatureKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	return FeatureKey_List{l}, err
}

func (s FeatureKey_List) At(i int) FeatureKey { return FeatureKey{s.List.Struct(i)} }

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
	return FeatureKey_Customer{st}, err
}

func NewRootFeatureKey_Customer(s *capnp.Segment) (FeatureKey_Customer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return FeatureKey_Customer{st}, err
}

func ReadRootFeatureKey_Customer(msg *capnp.Message) (FeatureKey_Customer, error) {
	root, err := msg.RootPtr()
	return FeatureKey_Customer{root.Struct()}, err
}

func (s FeatureKey_Customer) String() string {
	str, _ := text.Marshal(0x94066c58db8fb34e, s.Struct)
	return str
}

func (s FeatureKey_Customer) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s FeatureKey_Customer) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s FeatureKey_Customer) OrganizationName() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s FeatureKey_Customer) HasOrganizationName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s FeatureKey_Customer) OrganizationNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s FeatureKey_Customer) SetOrganizationName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s FeatureKey_Customer) ContactName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s FeatureKey_Customer) HasContactName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s FeatureKey_Customer) ContactNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s FeatureKey_Customer) SetContactName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s FeatureKey_Customer) ContactEmail() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s FeatureKey_Customer) HasContactEmail() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s FeatureKey_Customer) ContactEmailBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s FeatureKey_Customer) SetContactEmail(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// FeatureKey_Customer_List is a list of FeatureKey_Customer.
type FeatureKey_Customer_List struct{ capnp.List }

// NewFeatureKey_Customer creates a new list of FeatureKey_Customer.
func NewFeatureKey_Customer_List(s *capnp.Segment, sz int32) (FeatureKey_Customer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return FeatureKey_Customer_List{l}, err
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

const schema_d24581e9cd6a6772 = "x\xdal\x93Mh\xe4d\x18\xc7\xdf\xe7M23\xd1" +
	"\x19fb\xb2\xb0\xca\x96QT\xb6\xca\xae\xec\xb6~\xf6" +
	"\xb2\xa5:E\xfbEc\x8bT\xb1B\x9cy\x9d\xbe5" +
	"\xc9\x8cI\x06;*VA<\x08\xe2\xa5=I\xf1\xab" +
	"\xad\xa5 ED\x10=z\xf2\xe0\xc5\x9bx\xf0(\x15" +
	"oR\xc1S\xfc\xbf\xa9\x93)\xa5\x87\x81\xc9\xef\xf9\xca" +
	"\xff\xff<\xb9\xf1\x1dM\xf2\x9b\xc6U\x9d1\xf7\x86Q" +
	"H\x17\xbe\xf9\xe8\xb7\x15\xbf\xb0\xc5\xdc\x11\xa2T~\xbd" +
	"\xb3\xfd\xd2\xd1\xd3\xdb\xcc\xd0\x8a\x8c\x8d\xffKcd\x9b" +
	"\x1c\x7fm\x83\xff\xc1\xce\xc4\xdd;HK\xa3\xf6\xfa\xcf" +
	"\xc7\xef6~aF\x96r\xcc\xff\xb4O\xf8U\xd4Y" +
	"\xdac\x1c\xd9\xadX\xbesX\x8d\xf6\x985\xa2\x0fK" +
	"\x19\x8d\xcf\x1bSd\xaf\x1a(ZZ14Zj\x19" +
	"\x9c\x18K\xdf\xdc\xb8\xef\xcb\x9d\xfa\xe8\xbez\x17mX" +
	"p\x89\x17\x11\x1d\x17\xc6\x181\xb2\x03\xe3\x88=\x9f\xbe" +
	"\"\xbc\xa4\x17\x89\xeb\xda\xab\xa2\xffP\xd3\xeb\x86\xdd\x89" +
	"\xe9S4\x0b\xf0d/\xae'\x9d@D\x8bDnM" +
	"\x83V\x1d\x1d,\xef.\x88~Q#w\x8d\x13\x91C" +
	"\x8a\x89\x0f\xc0\xd6\xc0\x12N\x16\x07\xe4\x80\xaf\xbd\x0c\xd8" +
	"\x05|\x0bP\xe3\x0ei\x80\xfdu\xc0\x0d\xc0\xf78i" +
	"\xb2E&\xe3\xf8Q\xda\x89\xda^(\xdf\xf0(\x91\x9d" +
	"p\xc1\x0b\x04\x14\x96\x11+#\xd6\xec\x84\x89\xd7L\x16" +
	"X\x11\xfc<m\xb0j\xe0I?\xc7\x03M\xfc\xbc\xa6" +
	"\xaa\x12\xe5\x96\xb0\x9f|[\x969s\xc6^\xf3\x85\x14" +
	"\x923\xc5\x0c>\xc6\xb2\x1d\xca\xb0=\xcb4\x94\x8d\x0e" +
	"\xe4\xdb&M\xc0p\x9d`x\x8d \xec\x7f\x0b\xec\x0a" +
	"\xcd\x80\x97\x15\xbf\x0cN<3\xc1\xbe\x94\xa5\xd7\x14\xbe" +
	"\xa2\xd25\xca|\xb0\xef\xa4)pG\xf1\xbb\x15\xd7u" +
	"\x870\xc2\x1e\xa1g\xc1\xaf(>\xaa\xb8\xf1\xa9C\x06" +
	"\xf8\xfd\x04\x8b\x81\xc0\x1fV\xbc\xf0\x99C\x05\xf0\x9bY" +
	"\x9fk\x8a?\xae\xc6*}\xf9\x05\xd8\x8f\xe0\xa5\xb8U" +
	"\xdcs\xe86\"\xfb\x1e\x82\xf9\x18\x87\xdck\xaa\x87\xb9" +
	"\xef\xd0\xed\xe8\xf1@63\xef}+\x16\xcdH$T" +
	"\x81\xa3\x15e\xf4\xd0\x14\xaa\x0d\xcdcD5F\xb7d" +
	"\x1c\xf7D\xbe\xc7M\xb1\xd1\x95\x91\x88\xf3\xbd\xf6b\x11" +
	"\xcd\xc9@2J\xa8\x04VR\xf7\x1f7|/N$" +
	"5\xa7\xa4\xef\xc3c\xd5\x0b1\x1c\xe6\xa6\x8c\x97#\xe9" +
	"\xf9\x83\xe7\xc1:c\xb5\x11\x19Ow\xa2e\xfc\xaf&" +
	"(\xcaS\x80#\x81\xe52\xea\x9f/\xbb\xae_x\xd9" +
	"\x83\xd5\x8a>\xc3mC\x94\xf9\xc9\xf1\xea_?,\x7f" +
	"|*\x8a\xd5\xf4\xf4\xd1\xc5\x7f~\xfd\xf6\xe0\xab\xf7\x8b" +
	"\xf7\xfexX\xfa\xde|\xe2r\xe5\xb9\x9f\x96v?|" +
	"=8x{\xe2\xef\xb1\x93\xad\xdf\x07\x13.\x1c\x80`" +
	"={m\xb7\xac\xe9\xb5\xcf\x1d*\x12Y\x8d\x07q\xfa" +
	"\x938\xfd9u6_8jW\xd63\x0a>\x05\xb8" +
	"\xa8\xbe\x9c]\x87L\xc0yx\xed\xce\x01\xaep\xaa\xfa" +
	"-\xaf\xabt\x19JX5\xf6\x02?\x7fR_\xcd\xbc" +
	"\x17zmV\x17\x81\x08\x93<\xf0_\x00\x00\x00\xff\xff" +
	"\xcc\xa3\x15\x08"

func init() {
	schemas.Register(schema_d24581e9cd6a6772,
		0x94066c58db8fb34e,
		0x9548af5e959db169,
		0xa5720faa80697364,
		0xa6281f9da724787b)
}

var x_d24581e9cd6a6772 = []byte{
	0, 0, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 4, 0, 0, 0,
	54, 80, 246, 217, 181, 168, 173, 134,
	35, 195, 170, 8, 186, 9, 57, 24,
	13, 86, 201, 83, 164, 141, 119, 109,
	168, 126, 58, 243, 50, 245, 148, 223,
}
