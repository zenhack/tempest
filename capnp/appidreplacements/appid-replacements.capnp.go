package appidreplacements

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in appid-replacements.capnp.
var (
	AppIdReplacementList = AppIdReplacement_List{List: capnp.MustUnmarshalRootPtr(x_a53cae3f717a1676[0:344]).List()}
)

type AppIdReplacement struct{ capnp.Struct }

func NewAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return AppIdReplacement{st}, err
}

func NewRootAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return AppIdReplacement{st}, err
}

func ReadRootAppIdReplacement(msg *capnp.Message) (AppIdReplacement, error) {
	root, err := msg.RootPtr()
	return AppIdReplacement{root.Struct()}, err
}

func (s AppIdReplacement) String() string {
	str, _ := text.Marshal(0x888dcc6878baa07a, s.Struct)
	return str
}

func (s AppIdReplacement) Original() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s AppIdReplacement) HasOriginal() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) OriginalBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s AppIdReplacement) SetOriginal(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s AppIdReplacement) Replacement() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s AppIdReplacement) HasReplacement() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) ReplacementBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s AppIdReplacement) SetReplacement(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s AppIdReplacement) RevokeExceptPackageIds() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s AppIdReplacement) HasRevokeExceptPackageIds() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) SetRevokeExceptPackageIds(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewRevokeExceptPackageIds sets the revokeExceptPackageIds field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s AppIdReplacement) NewRevokeExceptPackageIds(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// AppIdReplacement_List is a list of AppIdReplacement.
type AppIdReplacement_List struct{ capnp.List }

// NewAppIdReplacement creates a new list of AppIdReplacement.
func NewAppIdReplacement_List(s *capnp.Segment, sz int32) (AppIdReplacement_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return AppIdReplacement_List{l}, err
}

func (s AppIdReplacement_List) At(i int) AppIdReplacement { return AppIdReplacement{s.List.Struct(i)} }

func (s AppIdReplacement_List) Set(i int, v AppIdReplacement) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppIdReplacement_Promise is a wrapper for a AppIdReplacement promised by a client call.
type AppIdReplacement_Promise struct{ *capnp.Pipeline }

func (p AppIdReplacement_Promise) Struct() (AppIdReplacement, error) {
	s, err := p.Pipeline.Struct()
	return AppIdReplacement{s}, err
}

const schema_a53cae3f717a1676 = "x\xda|\x90\xbdn\xd4@\x14\x85\xe7\xdak\x82P\x94" +
	"\x8d\x95\x00\xa2\"\x0de\x90'\xbb\xeb\x1f\x14)\x04\x91" +
	"\"(H\x9bDHIA1;\x1e\xdb\xeb\xc98\xe3" +
	"\xb5=\xb1\xfd\x0444\x14\xd444+\x01}^\x80" +
	"\x02J\x1e\x80\x8a\xb7@f\x16\x11\x05\x1a\x0aKG\x9f" +
	"\x8e|\xcf7\xab\xdf\x1e\x1b\xd8\xfal t\xb8a\xdd" +
	"\xe8\xdaw\x97u\xf2\xf5\xf5+d\xdf\x83N\xddi\xf3" +
	"\x9d\x8f\xdb\xef\x91e.!4\xd8\x87#X{\x09:" +
	"\xae\x9d\xc2'\x04\xdd\xf7\xbb??\xbc}\xf3\xe5\x87." +
	"\xf7\xae\xcb\x08\x06\xb61\x83\xb5\x07\x86n\x1eo\x18&" +
	"\x1co\x1b\x06\xa0'\x1d\x91r\x1an\xceXO\x9e\x11" +
	"\xca\x04\xcb\xca\xe2!%2\x93\x8fv\xa5\xdc\x0f\x8f\xd8" +
	"\x15\x87r\x0cp\xb8l\xf6\x10\xea\x01B\xf6\xde3=" +
	"\xef\xa9\x09\x87c\x03l\x80uX\xc0\xe7\x13\x0d\x0f4" +
	"<\xd1\xd00\xd6A;\xd8/.5<\xd1\xb04\xa0" +
	";\x9fM\xe3iF\xce\x90\xde\xb4\x8c\x0c\xfdA7\xfb" +
	"s\x04-\xe9\xf3\x7fQu\xce\xd9^\x0d\x94\xc9rL" +
	"(\xdf!1\xdb\x0f\x0bXA06\xe1wO\xc7\xff" +
	"\x18\x90\x7f\x0d\xee\x97\x07\xd3baq\xf5\x87\xd5\xeb\xa7" +
	"E\x0b\x88l\xf0vo\x9a\xf8\xf6\x1c\xf0\xe6\x1c\xec\xbd" +
	"\x15|\xaac>\xd7j\x9dJ\x15\xe3\x8c&V8\x08" +
	"\xfc(\xc3%\xe6\x19\x0e\xe30\x13\x84\xf1<\x0fR\x9e" +
	"\xaaA[\xc4m#\xe8\xb0\x0d\xf0\xa0?c\x91\xd3]" +
	"\xe4\xc1(\x17UiQN\x9d&\x12\x8c*\x877\x81" +
	"K\xf3:\x96x\xe4V\xd2/\x94\x8f\x9bZ\x096k" +
	"F~\xde\xf7\xbd4\xc1p\x0b\xba\xc9h2\x09B\xdf" +
	"1\xf1\x908Q0\xc1\xa1\x8b\xd9\x16f^\xe0\x86\x9e" +
	"\x1fR\xbaXU3\x9f\\\xd0\xdaRe\xea6\x95\xa3" +
	"\xe2T\x0a\\\x16\x84U^\xed+\xbf\x8c\xa4\x87\xcbF" +
	"e\xc2m8g2\xe9WA\xeet\x99O3\x0f\x0f" +
	"\x1d\xcb\xcb\x86\x82\xb5\x99'b\x87\x8f8\x17[8\xad" +
	"*\x99$\x8c\xd2\xada\x12F\x01\x8fFn\xdb\xd4}" +
	"1$\xc9\xaf\x00\x00\x00\xff\xff'\x0c\xc2\xe6"

func init() {
	schemas.Register(schema_a53cae3f717a1676,
		0x888dcc6878baa07a,
		0xe6cb9296adfd17e0)
}

var x_a53cae3f717a1676 = []byte{
	0, 0, 0, 0, 42, 0, 0, 0,
	1, 0, 0, 0, 55, 0, 0, 0,
	8, 0, 0, 0, 0, 0, 3, 0,
	21, 0, 0, 0, 170, 1, 0, 0,
	45, 0, 0, 0, 170, 1, 0, 0,
	69, 0, 0, 0, 14, 0, 0, 0,
	89, 0, 0, 0, 170, 1, 0, 0,
	113, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	118, 106, 118, 101, 107, 101, 99, 104,
	100, 51, 57, 56, 102, 110, 49, 116,
	49, 107, 110, 49, 100, 103, 100, 110,
	109, 97, 101, 107, 113, 113, 57, 106,
	107, 106, 118, 51, 122, 115, 103, 122,
	121, 109, 99, 52, 122, 57, 49, 51,
	114, 101, 102, 48, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	1, 0, 0, 0, 10, 1, 0, 0,
	98, 53, 98, 98, 57, 100, 56, 48,
	49, 52, 97, 48, 102, 57, 98, 49,
	100, 54, 49, 101, 50, 49, 101, 55,
	57, 54, 100, 55, 56, 100, 99, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	118, 120, 101, 56, 97, 119, 99, 120,
	118, 116, 106, 54, 121, 117, 48, 118,
	103, 106, 112, 109, 49, 116, 115, 97,
	101, 117, 55, 120, 56, 118, 56, 116,
	102, 112, 55, 49, 116, 121, 118, 110,
	109, 54, 121, 107, 107, 101, 112, 104,
	117, 57, 113, 48, 0, 0, 0, 0,
	110, 56, 99, 110, 55, 49, 52, 48,
	55, 110, 52, 109, 101, 122, 110, 55,
	109, 103, 48, 107, 53, 107, 107, 109,
	50, 49, 106, 117, 117, 112, 104, 104,
	101, 99, 99, 50, 52, 104, 100, 102,
	57, 107, 102, 53, 54, 122, 121, 120,
	109, 52, 97, 104, 0, 0, 0, 0,
}
