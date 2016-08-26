package appidreplacementstest

// AUTO GENERATED - DO NOT EDIT

import (
	appidreplacements "zenhack.net/go/sandstorm/capnp/appidreplacements"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in appid-replacements-test.capnp.
const (
	TestIds_unusedApp = "6pm4ujs8f5f5wugc87uhuhvhs57he09u10rv8qd2jgdup9f69yzh"
	TestIds_app1      = "vjvekechd398fn1t1kn1dgdnmaekqq9jkjv3zsgzymc4z913ref0"
	TestIds_app2      = "wq95qmutckc0yfmecv0ky96cqxgp156up8sv81yxvmery58q87jh"
	TestIds_app3      = "302t6c6kf8hjer1kh3469d4ch10d936g7wkwtxcs12pwh9u5axqh"
	TestIds_app4      = "5ddk4uqnstnsqvp3thc2tyed41c7wp4x5ygt20zrh3u0tnv5jqd0"
	TestIds_app5      = "jkz6yhywhp4uk5sgkc5ugwnee57a5h5wu4rfmujtahny5r8g3ych"
	TestIds_app6      = "adk6syfj42fpp3xhgqrrheqgfxkhaw8e1t11vug44ys6pzaxqugh"
	TestIds_unusedPkg = "7300e3448dd2b53e075d0a8481c2bc06"
	TestIds_pkg1      = "b5bb9d8014a0f9b1d61e21e796d78dcc"
	TestIds_pkg2      = "8613a11b8ac365cb36775a6b8ca6176c"
	TestIds_pkg3      = "77c4f45aee83e376d31a5680cdb841a2"
)

// Constants defined in appid-replacements-test.capnp.
var (
	TestAppIdReplacementList = appidreplacements.AppIdReplacement_List{List: capnp.MustUnmarshalRootPtr(x_bee445adfb01a777[0:712]).List()}
)

type TestIds struct{ capnp.Struct }

func NewTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds{st}, err
}

func NewRootTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds{st}, err
}

func ReadRootTestIds(msg *capnp.Message) (TestIds, error) {
	root, err := msg.RootPtr()
	return TestIds{root.Struct()}, err
}

func (s TestIds) String() string {
	str, _ := text.Marshal(0x9440399ec56efc9b, s.Struct)
	return str
}

// TestIds_List is a list of TestIds.
type TestIds_List struct{ capnp.List }

// NewTestIds creates a new list of TestIds.
func NewTestIds_List(s *capnp.Segment, sz int32) (TestIds_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return TestIds_List{l}, err
}

func (s TestIds_List) At(i int) TestIds { return TestIds{s.List.Struct(i)} }

func (s TestIds_List) Set(i int, v TestIds) error { return s.List.SetStruct(i, v.Struct) }

// TestIds_Promise is a wrapper for a TestIds promised by a client call.
type TestIds_Promise struct{ *capnp.Pipeline }

func (p TestIds_Promise) Struct() (TestIds, error) {
	s, err := p.Pipeline.Struct()
	return TestIds{s}, err
}

const schema_bee445adfb01a777 = "x\xda\xac\x95[\x88\x1cE\x17\xc7\xbb\xb6\xa7wCH" +
	"2\xd9o\xc9\xf7\x14\\\x0c(2\xb8\xb1\xab\xbb\xab/" +
	"\xa2d\xf3\xb0\xc8\xa2\x81\xbdD\x88\x92\x88=U\xd5\xdd" +
	"\x99v\x9a\xee\xe9\xcbL\xcfKH4h@A!\x88" +
	"\x17\x94\x08Jp\xe3\xa2y\x08\x091\x11\x1f\x8c\x12\x88" +
	"\x0f\xa2\xe2\x05\x02\x0ay\x12}\x10Q#\x09\x8cU3" +
	"\xd1\xb43\xa8;\x83\x0fs8\x0c?\xce\xbf\xce\xa9\x7f" +
	"\x9d\x96\xf7\x8d\xcd\x8eAi\xc7\x84 ,\xee\x96\xc6;" +
	"\xbf\xbc{\xf0\xc1\xf8\xc0\xe5\xc7\x85\xc9\xdbK\x9d\x97\xaf" +
	"\x07\x17^\xb5f\x8f\x0a\x02P\x0f\x8b\x0f\x83\xa9\x17E" +
	"\xc6-\x1f\x15E\xb0|L\x1c\x03\x82p\x13\x99\xbc\x05" +
	"t\x9a\xc7\xc1\xb5\xd5\xb9+\xef\x09%\x86\xa9\xab\xa2\x02" +
	"\xa6\xce\x8b\xa7\x05\xb1\xf3\xf3\xf4\x89{\xae\x9e\x9cx~" +
	"\xa0\xeam%V\xd5\xe2\xf8\xb2VbUgK\xdd\xaa" +
	"\xa7\xee\xcf\xc5\xc3/\x89\xaf\x0f\xf0\xbbJ\x87\xc0\x94\xdd" +
	"\xe5\xf7r\xde\xeb\xf1\x0f\xa1\xe0\x8dl\xe3]o\x0f\xf0" +
	")\xaf\x7f\xa4\xcb?\xc1\xf9g{\xbc\xb7\xfa]\xeb\xf4" +
	"\xb9'\x07\xf9W8\x7f\xb2\xcb\xbf\xc5\xf93=\xfe\xc4" +
	"Fyz\xf5\x85w\xce\x0d\xf0\x178\xffU\x97\xff\x8c" +
	"\xf3\xdf\xf4\xf8/\x9e\x9e\x8f\xa7\xfdC\xef\x0f\xf0?p" +
	"\x1eH\x9c\xbf\xce\xf9uR\x97_\xf7\xf5s[\xff\xf7" +
	"\xe8\xe7\x1f\x0e\xf0\x1b%\xd6\xef\xad]~\xab\xc4\xf8;" +
	"z\xfc\xb6+\xdfo\xba\xb6\xa5s\x91M\xbdts\xea" +
	"\x8c\x9f\x91\xce\x82\xa9\xb9.?\xcb\xf9\xbd=\xfe\xa7m" +
	"\x97v\xa7\xe2\xc1O\x07\xea\xbf9\xce\xces~\x9c\xf3" +
	"g\xc6\x19\xff\xc1x\x97\x7f\xe4\x13\xfb\xc0\xf2G\xf7\xfd" +
	":\xc0_\xe4\xfc\xe5.\xff%\xe7\xaf\xf4\xf8u~x" +
	"|Of\\\x1d\xe0\x7f\xe4\xbc4\xc1\xf8\xa5\x09\x86o" +
	"\x98`\xf8\x9e\x8e\x1d\x86\xfb\xc9L\x83\x96\xc2\xc7lL" +
	"\xeb4H\xe2\x99\x84\xc6\xc9vl\x87Ax\xf7n\x96" +
	"\xce\x93x\xbb\x1d\x8a!Z\x00\x00l\x10\xc6\xd8O\x80" +
	"`\x05tj~[\xcf\xbd\xbc)y\xa1\x96\xfa(v" +
	"}\x8cR\xb7\x19P\x8a\x0c\x1by\xa8\x99j\x0d\xa7\x9e" +
	"\xd6\x12\xdb\x0br\xd40\xdd\xb2\x9ac\xefOI\xf1\xef" +
	"$\xa7\xbb\x9a\x8bw\x02Pp\xdf\xe2R\xc1Z\x8b\x95" +
	"\xc2\\vU\x0a\xa6\x9e\xaf\x14\x1c5W)<\xa2\x9d" +
	"\x95\x82w\xee\xad\x14.\xdaZ*L\xcd\xaa\x14,\x83" +
	"*\x85\xfb\x82\x95N\x1a\xa41%;C\x01\x84e\xd6" +
	"\x06\xe4A\xe1A\xe5A\xe3\x01\xf1\xa0\xdf \x17|\x01" +
	"\xb8\xe5\xd0w!\x0f\x0a\x0f\xea0#W\xfbG\xae\xca" +
	"J\xa2c\xdd\x97\x1c\xd3\xab\xd1\x06\xf4=U\xd3-\xa2" +
	"a\x0f\xca\xc4Ru\xd7h\xfa\xcd\xa4\x85c\xa8\x84M" +
	"\xcfJ\xd9iZ\xd1\xcd\x91K\xff,\xf9G{\xa1\xd0" +
	"\xaf\xab\x87u-\xad\xc5\xa6\xe4 \x87\xdd\xab\x8bM#" +
	"\xf5R/\xf3bdxT\xb6R(723\"J" +
	"\xcd%ih9z\xd9\xca\xdb\xde0\xad\xc2~\xc9\xac" +
	"\x96Q\x9fbO\"\xaae:\x01L\xa0\x1f@\xe2\x92" +
	"\xa0nS?\x8a\xac\x9a_\xcb\xd4v\xec\xb6\xf3:\xd6" +
	"\xda\x16T\xcb\x0d\xea\xc8\xc3Hj\xfd\x92\x88\x10_K" +
	"\xa3@\x8a\x93 \x8e\xb2PM<\xac$9%\x1a\xc4" +
	"F3\xd4Z(w\x13En7<5\x95\x93 C" +
	"\xe5ZD\x86\x92\xd4\xfb%m\xe2\xebq\xee\xd4$M" +
	"q\x98\x8dZ\x9e\x1b5\x1a\x1e\x8d\\\xa7\xe5{v\xd3" +
	"\xa4\xaco\x98\xa5\xae\xa6\xe5\xb1\x1e\xb6\xedV9J\xdd" +
	"5\x0f6\xf4EW\xf9\x8b\xe4z\xd01u\xa8\xda\x10" +
	"VE\xd3\xc6\xaa\x8epU\xd5\x0d\x03\xd9z\xd5\xc4\xb6" +
	"\x0e\x0d\x1d\xb3-2\x94a\x16|W\xe8\x171TY" +
	"\xa6\xaa\xa6\x89&!J\x15\xa9T6\x10\x91mS3" +
	"!V\xaaX\xd6\xff]\x84\xa7\xcc\x8b\xf3d\x89\xde\x00" +
	"v\x04\xc9\x03\xfb\xe3\x84km\x12\xc0\x82\x08\xc0\xe6N" +
	"\xfb\xd8\xd9\x96w\xe9\x99\xa7\xd8\x8ac\x7f\x0a\x93\xc0\xdd" +
	"\xb9Y\x843+\x00\xce\xad\x80\xc9}\xff\x87GX\xfa" +
	"\xda\x0a[\x8d\xf0\x14\xcb>\xeef\xdf\xb2\xec\xb7\x150" +
	"5\x096\x8d\xea\xb3fd\xa1\xa8\x9e&\x12\xf6\xb1\x9c" +
	";u\x8a3\xd9\xcf-\x1dG-7\x84HOC3" +
	"\xceL\x98\xb7\xb2:m\xe4\xc8\x8c\xca\xa6Q\xf3\xa0\xb4" +
	"\x1e\xc0-l>UT\xadZ\xc4\x94E\xa8\xd9\xb2c" +
	"U!\xd1!U 5,\x9d\x18&\xc1\xfc\x12\xd6r" +
	"Q\xa3\x1dc\xd4=2\xe2\x03\x19\xf1C1\xea\xf7e" +
	"\xa4'\xd5s\xad\x815GC\xb6H\xa9\xa9R\xd5\xd0" +
	"\x89\x0am\xa4\x9b2&US\x83\xb6Rp\xed\x1a\x9e" +
	"\x9e:\xf0*\xfe\xbb\xfa|\x9b(\xfd\xdbdD7\x0c" +
	"\xd1\x12\xecoi\x0dF\xfe=\x00\x00\xff\xffR\xda\x85" +
	"p"

func init() {
	schemas.Register(schema_bee445adfb01a777,
		0x83dd7f735581bbf6,
		0x9440399ec56efc9b,
		0x9607b1f83cab1ff5,
		0xa4039a8503794bb5,
		0xaf2f0d76a56e3559,
		0xaf87bcb778eaad68,
		0xbcb098ad1f300dab,
		0xc0826b1f73498cd7,
		0xc6d560121c91da08,
		0xc9ff15fb0eece422,
		0xd381037554cc22f3,
		0xf747c7537f61d15e,
		0xf8377658a7706b08)
}

var x_bee445adfb01a777 = []byte{
	0, 0, 0, 0, 88, 0, 0, 0,
	1, 0, 0, 0, 103, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 3, 0,
	45, 0, 0, 0, 170, 1, 0, 0,
	69, 0, 0, 0, 170, 1, 0, 0,
	93, 0, 0, 0, 22, 0, 0, 0,
	137, 0, 0, 0, 170, 1, 0, 0,
	161, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	181, 0, 0, 0, 170, 1, 0, 0,
	205, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	225, 0, 0, 0, 170, 1, 0, 0,
	249, 0, 0, 0, 170, 1, 0, 0,
	17, 1, 0, 0, 14, 0, 0, 0,
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
	5, 0, 0, 0, 10, 1, 0, 0,
	21, 0, 0, 0, 10, 1, 0, 0,
	98, 53, 98, 98, 57, 100, 56, 48,
	49, 52, 97, 48, 102, 57, 98, 49,
	100, 54, 49, 101, 50, 49, 101, 55,
	57, 54, 100, 55, 56, 100, 99, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	56, 54, 49, 51, 97, 49, 49, 98,
	56, 97, 99, 51, 54, 53, 99, 98,
	51, 54, 55, 55, 53, 97, 54, 98,
	56, 99, 97, 54, 49, 55, 54, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	51, 48, 50, 116, 54, 99, 54, 107,
	102, 56, 104, 106, 101, 114, 49, 107,
	104, 51, 52, 54, 57, 100, 52, 99,
	104, 49, 48, 100, 57, 51, 54, 103,
	55, 119, 107, 119, 116, 120, 99, 115,
	49, 50, 112, 119, 104, 57, 117, 53,
	97, 120, 113, 104, 0, 0, 0, 0,
	53, 100, 100, 107, 52, 117, 113, 110,
	115, 116, 110, 115, 113, 118, 112, 51,
	116, 104, 99, 50, 116, 121, 101, 100,
	52, 49, 99, 55, 119, 112, 52, 120,
	53, 121, 103, 116, 50, 48, 122, 114,
	104, 51, 117, 48, 116, 110, 118, 53,
	106, 113, 100, 48, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	97, 100, 107, 54, 115, 121, 102, 106,
	52, 50, 102, 112, 112, 51, 120, 104,
	103, 113, 114, 114, 104, 101, 113, 103,
	102, 120, 107, 104, 97, 119, 56, 101,
	49, 116, 49, 49, 118, 117, 103, 52,
	52, 121, 115, 54, 112, 122, 97, 120,
	113, 117, 103, 104, 0, 0, 0, 0,
	1, 0, 0, 0, 10, 1, 0, 0,
	55, 55, 99, 52, 102, 52, 53, 97,
	101, 101, 56, 51, 101, 51, 55, 54,
	100, 51, 49, 97, 53, 54, 56, 48,
	99, 100, 98, 56, 52, 49, 97, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}
