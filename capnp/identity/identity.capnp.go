package identity

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Identity struct{ Client capnp.Client }

func (c Identity) GetProfile(ctx context.Context, params func(Identity_getProfile_Params) error, opts ...capnp.CallOption) Identity_getProfile_Results_Promise {
	if c.Client == nil {
		return Identity_getProfile_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Identity_getProfile_Params{Struct: s}) }
	}
	return Identity_getProfile_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Identity_Server interface {
	GetProfile(Identity_getProfile) error
}

func Identity_ServerToClient(s Identity_Server) Identity {
	c, _ := s.(server.Closer)
	return Identity{Client: server.New(Identity_Methods(nil, s), c)}
}

func Identity_Methods(methods []server.Method, s Identity_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Identity_getProfile{c, opts, Identity_getProfile_Params{Struct: p}, Identity_getProfile_Results{Struct: r}}
			return s.GetProfile(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Identity_getProfile holds the arguments for a server call to Identity.getProfile.
type Identity_getProfile struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Identity_getProfile_Params
	Results Identity_getProfile_Results
}

type Identity_PowerboxTag struct{ capnp.Struct }

// Identity_PowerboxTag_TypeID is the unique identifier for the type Identity_PowerboxTag.
const Identity_PowerboxTag_TypeID = 0xf35052cfb1d3bbe8

func NewIdentity_PowerboxTag(s *capnp.Segment) (Identity_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_PowerboxTag{st}, err
}

func NewRootIdentity_PowerboxTag(s *capnp.Segment) (Identity_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_PowerboxTag{st}, err
}

func ReadRootIdentity_PowerboxTag(msg *capnp.Message) (Identity_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return Identity_PowerboxTag{root.Struct()}, err
}

func (s Identity_PowerboxTag) String() string {
	str, _ := text.Marshal(0xf35052cfb1d3bbe8, s.Struct)
	return str
}

func (s Identity_PowerboxTag) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.BitList{List: p.List()}, err
}

func (s Identity_PowerboxTag) HasPermissions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Identity_PowerboxTag) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s Identity_PowerboxTag) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Identity_PowerboxTag_List is a list of Identity_PowerboxTag.
type Identity_PowerboxTag_List struct{ capnp.List }

// NewIdentity_PowerboxTag creates a new list of Identity_PowerboxTag.
func NewIdentity_PowerboxTag_List(s *capnp.Segment, sz int32) (Identity_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Identity_PowerboxTag_List{l}, err
}

func (s Identity_PowerboxTag_List) At(i int) Identity_PowerboxTag {
	return Identity_PowerboxTag{s.List.Struct(i)}
}

func (s Identity_PowerboxTag_List) Set(i int, v Identity_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// Identity_PowerboxTag_Promise is a wrapper for a Identity_PowerboxTag promised by a client call.
type Identity_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p Identity_PowerboxTag_Promise) Struct() (Identity_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return Identity_PowerboxTag{s}, err
}

type Identity_getProfile_Params struct{ capnp.Struct }

// Identity_getProfile_Params_TypeID is the unique identifier for the type Identity_getProfile_Params.
const Identity_getProfile_Params_TypeID = 0xf32d79a7b575d94c

func NewIdentity_getProfile_Params(s *capnp.Segment) (Identity_getProfile_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Identity_getProfile_Params{st}, err
}

func NewRootIdentity_getProfile_Params(s *capnp.Segment) (Identity_getProfile_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Identity_getProfile_Params{st}, err
}

func ReadRootIdentity_getProfile_Params(msg *capnp.Message) (Identity_getProfile_Params, error) {
	root, err := msg.RootPtr()
	return Identity_getProfile_Params{root.Struct()}, err
}

func (s Identity_getProfile_Params) String() string {
	str, _ := text.Marshal(0xf32d79a7b575d94c, s.Struct)
	return str
}

// Identity_getProfile_Params_List is a list of Identity_getProfile_Params.
type Identity_getProfile_Params_List struct{ capnp.List }

// NewIdentity_getProfile_Params creates a new list of Identity_getProfile_Params.
func NewIdentity_getProfile_Params_List(s *capnp.Segment, sz int32) (Identity_getProfile_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Identity_getProfile_Params_List{l}, err
}

func (s Identity_getProfile_Params_List) At(i int) Identity_getProfile_Params {
	return Identity_getProfile_Params{s.List.Struct(i)}
}

func (s Identity_getProfile_Params_List) Set(i int, v Identity_getProfile_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Identity_getProfile_Params_Promise is a wrapper for a Identity_getProfile_Params promised by a client call.
type Identity_getProfile_Params_Promise struct{ *capnp.Pipeline }

func (p Identity_getProfile_Params_Promise) Struct() (Identity_getProfile_Params, error) {
	s, err := p.Pipeline.Struct()
	return Identity_getProfile_Params{s}, err
}

type Identity_getProfile_Results struct{ capnp.Struct }

// Identity_getProfile_Results_TypeID is the unique identifier for the type Identity_getProfile_Results.
const Identity_getProfile_Results_TypeID = 0xcd7272b855c92f6d

func NewIdentity_getProfile_Results(s *capnp.Segment) (Identity_getProfile_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_getProfile_Results{st}, err
}

func NewRootIdentity_getProfile_Results(s *capnp.Segment) (Identity_getProfile_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_getProfile_Results{st}, err
}

func ReadRootIdentity_getProfile_Results(msg *capnp.Message) (Identity_getProfile_Results, error) {
	root, err := msg.RootPtr()
	return Identity_getProfile_Results{root.Struct()}, err
}

func (s Identity_getProfile_Results) String() string {
	str, _ := text.Marshal(0xcd7272b855c92f6d, s.Struct)
	return str
}

func (s Identity_getProfile_Results) Profile() (Profile, error) {
	p, err := s.Struct.Ptr(0)
	return Profile{Struct: p.Struct()}, err
}

func (s Identity_getProfile_Results) HasProfile() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Identity_getProfile_Results) SetProfile(v Profile) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewProfile sets the profile field to a newly
// allocated Profile struct, preferring placement in s's segment.
func (s Identity_getProfile_Results) NewProfile() (Profile, error) {
	ss, err := NewProfile(s.Struct.Segment())
	if err != nil {
		return Profile{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Identity_getProfile_Results_List is a list of Identity_getProfile_Results.
type Identity_getProfile_Results_List struct{ capnp.List }

// NewIdentity_getProfile_Results creates a new list of Identity_getProfile_Results.
func NewIdentity_getProfile_Results_List(s *capnp.Segment, sz int32) (Identity_getProfile_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Identity_getProfile_Results_List{l}, err
}

func (s Identity_getProfile_Results_List) At(i int) Identity_getProfile_Results {
	return Identity_getProfile_Results{s.List.Struct(i)}
}

func (s Identity_getProfile_Results_List) Set(i int, v Identity_getProfile_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Identity_getProfile_Results_Promise is a wrapper for a Identity_getProfile_Results promised by a client call.
type Identity_getProfile_Results_Promise struct{ *capnp.Pipeline }

func (p Identity_getProfile_Results_Promise) Struct() (Identity_getProfile_Results, error) {
	s, err := p.Pipeline.Struct()
	return Identity_getProfile_Results{s}, err
}

func (p Identity_getProfile_Results_Promise) Profile() Profile_Promise {
	return Profile_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Profile struct{ capnp.Struct }

// Profile_TypeID is the unique identifier for the type Profile.
const Profile_TypeID = 0xd3d0c34d7201fcef

func NewProfile(s *capnp.Segment) (Profile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Profile{st}, err
}

func NewRootProfile(s *capnp.Segment) (Profile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Profile{st}, err
}

func ReadRootProfile(msg *capnp.Message) (Profile, error) {
	root, err := msg.RootPtr()
	return Profile{root.Struct()}, err
}

func (s Profile) String() string {
	str, _ := text.Marshal(0xd3d0c34d7201fcef, s.Struct)
	return str
}

func (s Profile) DisplayName() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s Profile) HasDisplayName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Profile) SetDisplayName(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisplayName sets the displayName field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Profile) NewDisplayName() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Profile) PreferredHandle() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Profile) HasPreferredHandle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Profile) PreferredHandleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Profile) SetPreferredHandle(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Profile) Picture() util.StaticAsset {
	p, _ := s.Struct.Ptr(2)
	return util.StaticAsset{Client: p.Interface().Client()}
}

func (s Profile) HasPicture() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Profile) SetPicture(v util.StaticAsset) error {
	if v.Client == nil {
		return s.Struct.SetPtr(2, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(2, in.ToPtr())
}

func (s Profile) Pronouns() Profile_Pronouns {
	return Profile_Pronouns(s.Struct.Uint16(0))
}

func (s Profile) SetPronouns(v Profile_Pronouns) {
	s.Struct.SetUint16(0, uint16(v))
}

// Profile_List is a list of Profile.
type Profile_List struct{ capnp.List }

// NewProfile creates a new list of Profile.
func NewProfile_List(s *capnp.Segment, sz int32) (Profile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return Profile_List{l}, err
}

func (s Profile_List) At(i int) Profile { return Profile{s.List.Struct(i)} }

func (s Profile_List) Set(i int, v Profile) error { return s.List.SetStruct(i, v.Struct) }

// Profile_Promise is a wrapper for a Profile promised by a client call.
type Profile_Promise struct{ *capnp.Pipeline }

func (p Profile_Promise) Struct() (Profile, error) {
	s, err := p.Pipeline.Struct()
	return Profile{s}, err
}

func (p Profile_Promise) DisplayName() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Profile_Promise) Picture() util.StaticAsset {
	return util.StaticAsset{Client: p.Pipeline.GetPipeline(2).Client()}
}

type Profile_Pronouns uint16

// Values of Profile_Pronouns.
const (
	Profile_Pronouns_neutral Profile_Pronouns = 0
	Profile_Pronouns_male    Profile_Pronouns = 1
	Profile_Pronouns_female  Profile_Pronouns = 2
	Profile_Pronouns_robot   Profile_Pronouns = 3
)

// String returns the enum's constant name.
func (c Profile_Pronouns) String() string {
	switch c {
	case Profile_Pronouns_neutral:
		return "neutral"
	case Profile_Pronouns_male:
		return "male"
	case Profile_Pronouns_female:
		return "female"
	case Profile_Pronouns_robot:
		return "robot"

	default:
		return ""
	}
}

// Profile_PronounsFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Profile_PronounsFromString(c string) Profile_Pronouns {
	switch c {
	case "neutral":
		return Profile_Pronouns_neutral
	case "male":
		return Profile_Pronouns_male
	case "female":
		return Profile_Pronouns_female
	case "robot":
		return Profile_Pronouns_robot

	default:
		return 0
	}
}

type Profile_Pronouns_List struct{ capnp.List }

func NewProfile_Pronouns_List(s *capnp.Segment, sz int32) (Profile_Pronouns_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Profile_Pronouns_List{l.List}, err
}

func (l Profile_Pronouns_List) At(i int) Profile_Pronouns {
	ul := capnp.UInt16List{List: l.List}
	return Profile_Pronouns(ul.At(i))
}

func (l Profile_Pronouns_List) Set(i int, v Profile_Pronouns) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type UserInfo struct{ capnp.Struct }

// UserInfo_TypeID is the unique identifier for the type UserInfo.
const UserInfo_TypeID = 0x94b9d1efb35d11d3

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

func (s UserInfo) Pronouns() Profile_Pronouns {
	return Profile_Pronouns(s.Struct.Uint16(0))
}

func (s UserInfo) SetPronouns(v Profile_Pronouns) {
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

const schema_c822108a5c3d7d25 = "x\xda\x8c\x94M\x88\x1c\xc5\x1b\xc6\x9f\xa7\xaa{&\x81" +
	"\xdd\xffL\xd1\xf3'\x1b\x10\xf2A\x04\x15\xb3~\xe4 " +
	"\x06d\xc2^t\x17\x95\xae\x1d\xf7\xe0b\xc4\x9e\x9d\xda" +
	"u\xa0gz\xac\xee\xc1\xac\"+!{P\x09x\xf0" +
	"\xa0W/\x12\xc4\x83\x88\xe8\x9a\x8bD\x82\xe6\x10\x92\xb8" +
	"\xa7\x05\x05\xf5\x92\x08J\x14=\x89\xd0R=_K\xc4" +
	"\x90\xd3\x0co=\xfd\xbco\xbf\xcf\xaf\xeb\xc1\xff\x8b\x13" +
	"\xe2!\xffz\x09\xd0\xa1_\xca\xf9\xe1\xa3\x9bW\x8e\xf6" +
	"\xcf@\xed\x13\xf9\xcd\xbfi\x9f\xfa\xea\xea6\xc0c\xcf" +
	"J\xc1\xc0\xc82\x10Dr\x0d\xcc\xb7\xd5\xc9On^" +
	"\xdbz\x07\xbaB\xe6w\xbf\xf6\xd8soV\x0f\x7f\x03" +
	"\xbf\xec$\x9b\xf2Bp\xd6\x89\x8f\xbd!\xbf&\x98\xcf" +
	"|\xaf\xcf\xbd\xf2\xee\x99/\xa1*r\"\x06\x83M\xff" +
	"Bp\xd6\xdf\x07\x04\xef\xf9\x8f\x07\x17\xfd2\x90w\x1e" +
	"\xb8\xb4\xf4\xb9\xb5\x97\xa1f\x08\xf8t>\x1f\xf9\xcb\x04" +
	"\x83-\xbf\x0eN\x06\xbb\xa5w1\xde\x8e\xffE\xf0c" +
	"\xe1\xf8\xab\x7f\x1d\xcc\x9f\xdc\xe9\x7f\xfa\xc1\xfa\xd1?\x06" +
	"n\x9e3\xbb\\Z$\xbc\xfc\xc6\xf9\xed\x8f\xaf,\x86" +
	"\xc5\xc9d\xc2A\xbf\xad\xd2a\x06\x97J\xce\xf0b\xa9" +
	"\x8e\xb7\xf3v\xcbt\xb3v\xb6.gW\xa2^\xb7w" +
	"<\xb4\xc9j;6\xb3\xa1M\xbaI\xbf\x9b\"$u" +
	"\x95\x02P\xf7\xce\x01\xa4:t\x1f@\xa1\xf6\x1f\x07(" +
	"\x95z\x18\xd8\xe8\x9a~f\xa3\xb8\xd2\x89bS_5" +
	"\xee\xe7\x80M\x9aI6\xb6\x17C\xfb\xa5\xd4\xd8\xf9\xee" +
	"jR\xd8\x1e\x94\x1e\xe0\x11P\xd7\x9a\x80\xbe*\xa9\xbf" +
	"\x13T\x1ekt\xc5\x9do\x01\xfd\x83\xa4\xfeEP\x95" +
	"D\xad\x18\xe2\xe7e@\xdf\x90\xd4\x7f\x0a*_\xd6(" +
	"\x01\xf5\xbb{\xfc7\xc9\xc6\x14\x05\x15\xbd\x1a= \xd8" +
	"\xcb\xd3@c\x0f%\x1b5\x0aR\xd6\xe8\x03\x81\xe2\x02" +
	"\xd0\xa8\xba\xf2]N.\xfc\x1aK@\xb0\x9f\xcb@c" +
	"\xc6\xd5\x8f\xb8z\xb9Ts+\x0b\x0e\x15\xfa\x83\xae~" +
	"?\x05\xf3V;\xed\xc5\xd1\xfa\xd3(G\x1d\xc3j>" +
	"\xff\xd3\xab\x8f\xd8\xcfN\xbe\xe5\x96S\x05\xf3\x96\xe9Y" +
	"\xb3\x12e\xc2\xb4Bc;\xed4m'\xddt.N" +
	"\xd8\xe44\x04\xa7\xc1\xf1Z \xe7[\xe3bo\xa8F" +
	"9\xe9\xa6\xfc\x1f\x18J\x92\x10\xeeo\xde\xb3f\xd5X" +
	"k\xd8z\"\xea\xb6b\x03NAp\xaa8\x19&\x05" +
	"\xb02!\x1dd\xc5\x9d\xb6W\xb2\xbe5K\x906\x1e" +
	"?2n\x0fPM\x10q\xe9\xee:\x1de6?R" +
	"k\x8f\xdcE\x17\x9by\x98\xbcll39\x85\xf23" +
	"\xd1\x9a\xf6\xa4\x0f\x8c\xc1\xe4\x88w\xa5\x96!\xd4\xder" +
	"\xbef\xb2\x02/\xc8\xd8\x9c`\xc8I+\xef\x96V\xb3" +
	"#il\x8e,\x9a\xb4\x1f\xcb,\xd5\xde\x18\x97\xe99" +
	"@\xef\x91\xd45\xc1\x8d\xde@\xc7\xea\xee\x0f\xbbHb" +
	"\xe4\xce\x11\xdb\xf5\x81\xb4x\x8f\xc9\x95\xc0\x85<\x9c\xec" +
	"PW\xc7m\"\x87\xd5\x0b\x92:vT\x0d\xa9l\x9f" +
	"\x06\xf4\x8b\x92:s\xec\x0c\xa9|\xc9\x0d\x14K\xeaS" +
	"\x03\xce\x1c\x94\xfd\x05@g\x92\xfa\xf5;\x80\xe6?\xe3" +
	"\xdd\x18\x06H\x95\xd7\xdf?w\xa0\xf9\xfc\xf9\xbfF9" +
	"\xdd>\xf8;Ym\x18\xd9\xa8\xc3\xf4_\x97\xc0X;" +
	"\xca\xd7\xc5\xeb\xbe\xd8]\x11\xb8\xddLI\xea{\xc4\xed" +
	"\xc9\xfd'\x00\x00\xff\xff\xa0\xc8\x7fS"

func init() {
	schemas.Register(schema_c822108a5c3d7d25,
		0x84752dcf8539ab01,
		0x94b9d1efb35d11d3,
		0xc084987aa951dd18,
		0xcd7272b855c92f6d,
		0xd3d0c34d7201fcef,
		0xf32d79a7b575d94c,
		0xf35052cfb1d3bbe8)
}
