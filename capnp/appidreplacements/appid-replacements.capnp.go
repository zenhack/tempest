package appidreplacements

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
)

var (
	AppIdReplacementList = AppIdReplacement_List{List: capnp.ToList(capnp.MustUnmarshalRoot(x_a53cae3f717a1676[0:344]))}
)

type AppIdReplacement struct{ capnp.Struct }

func NewAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return AppIdReplacement{}, err
	}
	return AppIdReplacement{st}, nil
}

func NewRootAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return AppIdReplacement{}, err
	}
	return AppIdReplacement{st}, nil
}

func ReadRootAppIdReplacement(msg *capnp.Message) (AppIdReplacement, error) {
	root, err := msg.Root()
	if err != nil {
		return AppIdReplacement{}, err
	}
	st := capnp.ToStruct(root)
	return AppIdReplacement{st}, nil
}

func (s AppIdReplacement) Original() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s AppIdReplacement) SetOriginal(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s AppIdReplacement) Replacement() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s AppIdReplacement) SetReplacement(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s AppIdReplacement) RevokeExceptPackageIds() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s AppIdReplacement) SetRevokeExceptPackageIds(v capnp.TextList) error {

	return s.Struct.SetPointer(2, v.List)
}

// AppIdReplacement_List is a list of AppIdReplacement.
type AppIdReplacement_List struct{ capnp.List }

// NewAppIdReplacement creates a new list of AppIdReplacement.
func NewAppIdReplacement_List(s *capnp.Segment, sz int32) (AppIdReplacement_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return AppIdReplacement_List{}, err
	}
	return AppIdReplacement_List{l}, nil
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
