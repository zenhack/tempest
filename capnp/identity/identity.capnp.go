package identity

// AUTO GENERATED - DO NOT EDIT

import (
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Identity struct{ Client capnp.Client }

type Identity_Server interface {
}

func Identity_ServerToClient(s Identity_Server) Identity {
	c, _ := s.(server.Closer)
	return Identity{Client: server.New(Identity_Methods(nil, s), c)}
}

func Identity_Methods(methods []server.Method, s Identity_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type UserInfo struct{ capnp.Struct }

func NewUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	return UserInfo{st}, err
}

func NewRootUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	return UserInfo{st}, err
}

func ReadRootUserInfo(msg *capnp.Message) (UserInfo, error) {
	root, err := msg.RootPtr()
	return UserInfo{root.Struct()}, err
}

func (s UserInfo) String() string {
	str, _ := text.Marshal(0x94b9d1efb35d11d3, s.Struct)
	return str
}

func (s UserInfo) DisplayName() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s UserInfo) HasDisplayName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetDisplayName(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisplayName sets the displayName field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s UserInfo) NewDisplayName() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UserInfo) PreferredHandle() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s UserInfo) HasPreferredHandle() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UserInfo) PreferredHandleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s UserInfo) SetPreferredHandle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

func (s UserInfo) PictureUrl() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s UserInfo) HasPictureUrl() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s UserInfo) PictureUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s UserInfo) SetPictureUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, t.List.ToPtr())
}

func (s UserInfo) Pronouns() UserInfo_Pronouns {
	return UserInfo_Pronouns(s.Struct.Uint16(0))
}

func (s UserInfo) SetPronouns(v UserInfo_Pronouns) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s UserInfo) DeprecatedPermissionsBlob() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s UserInfo) HasDeprecatedPermissionsBlob() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetDeprecatedPermissionsBlob(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s UserInfo) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.BitList{List: p.List()}, err
}

func (s UserInfo) HasPermissions() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s UserInfo) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s UserInfo) IdentityId() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s UserInfo) HasIdentityId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetIdentityId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, d.List.ToPtr())
}

func (s UserInfo) Identity() Identity {
	p, _ := s.Struct.Ptr(6)
	return Identity{Client: p.Interface().Client()}
}

func (s UserInfo) HasIdentity() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetIdentity(v Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(6, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(6, in.ToPtr())
}

// UserInfo_List is a list of UserInfo.
type UserInfo_List struct{ capnp.List }

// NewUserInfo creates a new list of UserInfo.
func NewUserInfo_List(s *capnp.Segment, sz int32) (UserInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7}, sz)
	return UserInfo_List{l}, err
}

func (s UserInfo_List) At(i int) UserInfo { return UserInfo{s.List.Struct(i)} }

func (s UserInfo_List) Set(i int, v UserInfo) error { return s.List.SetStruct(i, v.Struct) }

// UserInfo_Promise is a wrapper for a UserInfo promised by a client call.
type UserInfo_Promise struct{ *capnp.Pipeline }

func (p UserInfo_Promise) Struct() (UserInfo, error) {
	s, err := p.Pipeline.Struct()
	return UserInfo{s}, err
}

func (p UserInfo_Promise) DisplayName() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UserInfo_Promise) Identity() Identity {
	return Identity{Client: p.Pipeline.GetPipeline(6).Client()}
}

type UserInfo_Pronouns uint16

// Values of UserInfo_Pronouns.
const (
	UserInfo_Pronouns_neutral UserInfo_Pronouns = 0
	UserInfo_Pronouns_male    UserInfo_Pronouns = 1
	UserInfo_Pronouns_female  UserInfo_Pronouns = 2
	UserInfo_Pronouns_robot   UserInfo_Pronouns = 3
)

// String returns the enum's constant name.
func (c UserInfo_Pronouns) String() string {
	switch c {
	case UserInfo_Pronouns_neutral:
		return "neutral"
	case UserInfo_Pronouns_male:
		return "male"
	case UserInfo_Pronouns_female:
		return "female"
	case UserInfo_Pronouns_robot:
		return "robot"

	default:
		return ""
	}
}

// UserInfo_PronounsFromString returns the enum value with a name,
// or the zero value if there's no such value.
func UserInfo_PronounsFromString(c string) UserInfo_Pronouns {
	switch c {
	case "neutral":
		return UserInfo_Pronouns_neutral
	case "male":
		return UserInfo_Pronouns_male
	case "female":
		return UserInfo_Pronouns_female
	case "robot":
		return UserInfo_Pronouns_robot

	default:
		return 0
	}
}

type UserInfo_Pronouns_List struct{ capnp.List }

func NewUserInfo_Pronouns_List(s *capnp.Segment, sz int32) (UserInfo_Pronouns_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return UserInfo_Pronouns_List{l.List}, err
}

func (l UserInfo_Pronouns_List) At(i int) UserInfo_Pronouns {
	ul := capnp.UInt16List{List: l.List}
	return UserInfo_Pronouns(ul.At(i))
}

func (l UserInfo_Pronouns_List) Set(i int, v UserInfo_Pronouns) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_c822108a5c3d7d25 = "x\xdad\x921\x88\xd4@\x14\x86\xdf?\x93l<\xb9" +
	"s7$\xe2.\x08\xde\x89\"\x8a\x1c\xa7'\x88\x07\x87" +
	"z\x95k!\x1b\x96k\xc4+\xb2\x9bY\x09d\x93e" +
	"\x92-N\x11A\xac\xb4\x13\x0b\xb5\xb3\xb4\xb6\x10;\xb9" +
	"FKO\xb1\x10\x14D\x10T\x04\x15\xb4\xb0\x8ao\x16" +
	"\xdd=\xb4\x98L\xf2\xbd\xff\x7f\xf3\xe6\xbd,\x1c\xc3)" +
	"q\xc4>`\x11\x05\x0bv\xa5|\xe9\xae=\xfc\xba\xf9" +
	"\xf86\x05U\xa0\xdc\x7fe\xf9\xc2\x8d\xda\xdegd;" +
	"\x0e\x91\xf7\x0b\x1b\x1e\xc4.\xa2\xc5\x19\xf1\x14\x84\xf2\xfe" +
	"\x87\xcf\xaf\xee\xfe\xbcu\x8f\xdc\xba\x98X\x09\x8bS\xd6" +
	"vx\x0d\xcb\x98vZ\x17YY\x7f\x1b<\xb8t\xe7" +
	"\xfa\x13r\xabr\x92\x96\xe0\x9d\xb06\xbc\xd3#\xe1\xb2" +
	"\xe5\x98E\xdd2\x8eTZ\xc4\xc5\xba\x98\xef\x86\x83t" +
	"\xb0\xb4\x9a+\xddL{\x19\x05\x16\xb6\x1e\x8a\xb3eK" +
	"gi6Ls\xe2\xfag%\xdf\xc2\x02\x91\xbb\xd9\xe1" +
	"\xcf\xe7\x12\xc1\x1b\x01\xd7\x82\x0f\x03_\xbf`\xf8\x8e\xe1" +
	"\x17\x86\x15\xe1C0\xfct\x9e\xe1G\x86?\x18\xda\xd2" +
	"\x87d\xf8\xdd\xd8\xbfI\xb4\xa7\xc1\x14\x96\x0fN\xecM" +
	"\xe1\x1aQ{\x1b\x98\xb3\x17`\xb1\xcd\x98\xab`\\3" +
	"x\xb7\x91\x0b\xdbG\x85y\x03\x9c\xba]7|\x9f\xe1" +
	"N\xc5\x87\xb9\xe6\xdcH?k\xf8a\xe6e\x14\xe7\x83" +
	"$\\?GN\xd8W\xa8\x95\xcd\xf7\x97\x8f\xebGk" +
	"7\xb99\xa8q\xe7\"5\xd0\xaa\x1b\x16BE-\xa5" +
	"\xfbq\x9e\xc7Y\x9a\xaf$\x19:\x98!\xc1\x0b\xe3\x86" +
	"\x91lFc8\xf8\xa3&\x87\xf5\xd8AhIp#" +
	"\x84y-9gOi\xad\x10\x9d\x09\xd3(Q\x84i" +
	"\x8eL\x8f\"\xe3\x8e\xa2:i6WS5\xd1\xb8[" +
	"\x0c\xb5Z%\xa9\x93\xb1e|<\xab\xdc\xc9\xa8\xd9\xe2" +
	"n\x89\xca\x7f\xa69?\x99]\x0b\x08j\xa3y\x1c\\" +
	"\x19\xd9\xe6\x0e\xf1&\xdc\xc6\x12o\xd2u\x8f\x12]M" +
	"\xd5\xb0\xd0aR\xed\x87\x89:\xd9Sf\xdb\xa3\xb3N" +
	"V\xfc\xf7\xbb4\xff\x96\xc3i[\xd2\xe6\xe7\xef\x00\x00" +
	"\x00\xff\xff\xae\xdf\xbdb"

func init() {
	schemas.Register(schema_c822108a5c3d7d25,
		0x94b9d1efb35d11d3,
		0x9a92f599d5eae5a1,
		0xc084987aa951dd18)
}
