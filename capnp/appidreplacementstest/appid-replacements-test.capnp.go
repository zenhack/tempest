package appidreplacementstest

// AUTO GENERATED - DO NOT EDIT

import (
	appidreplacements "zenhack.net/go/sandstorm/capnp/appidreplacements"
	capnp "zombiezen.com/go/capnproto2"
)

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

var (
	TestAppIdReplacementList = appidreplacements.AppIdReplacement_List{List: capnp.ToList(capnp.MustUnmarshalRoot(x_bee445adfb01a777[0:712]))}
)

type TestIds struct{ capnp.Struct }

func NewTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return TestIds{}, err
	}
	return TestIds{st}, nil
}

func NewRootTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return TestIds{}, err
	}
	return TestIds{st}, nil
}

func ReadRootTestIds(msg *capnp.Message) (TestIds, error) {
	root, err := msg.Root()
	if err != nil {
		return TestIds{}, err
	}
	st := capnp.ToStruct(root)
	return TestIds{st}, nil
}

// TestIds_List is a list of TestIds.
type TestIds_List struct{ capnp.List }

// NewTestIds creates a new list of TestIds.
func NewTestIds_List(s *capnp.Segment, sz int32) (TestIds_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return TestIds_List{}, err
	}
	return TestIds_List{l}, nil
}

func (s TestIds_List) At(i int) TestIds           { return TestIds{s.List.Struct(i)} }
func (s TestIds_List) Set(i int, v TestIds) error { return s.List.SetStruct(i, v.Struct) }

// TestIds_Promise is a wrapper for a TestIds promised by a client call.
type TestIds_Promise struct{ *capnp.Pipeline }

func (p TestIds_Promise) Struct() (TestIds, error) {
	s, err := p.Pipeline.Struct()
	return TestIds{s}, err
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
