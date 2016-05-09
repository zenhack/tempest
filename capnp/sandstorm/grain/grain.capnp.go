package grain

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/sandstorm/util"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type PowerboxDescriptor struct{ capnp.Struct }

func NewPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	return PowerboxDescriptor{st}, nil
}

func NewRootPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	return PowerboxDescriptor{st}, nil
}

func ReadRootPowerboxDescriptor(msg *capnp.Message) (PowerboxDescriptor, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	return PowerboxDescriptor{root.Struct()}, nil
}

func (s PowerboxDescriptor) Tags() (PowerboxDescriptor_Tag_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return PowerboxDescriptor_Tag_List{}, err
	}

	return PowerboxDescriptor_Tag_List{List: p.List()}, nil

}

func (s PowerboxDescriptor) HasTags() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor) SetTags(v PowerboxDescriptor_Tag_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

func (s PowerboxDescriptor) Quality() PowerboxDescriptor_MatchQuality {
	return PowerboxDescriptor_MatchQuality(s.Struct.Uint16(0))
}

func (s PowerboxDescriptor) SetQuality(v PowerboxDescriptor_MatchQuality) {

	s.Struct.SetUint16(0, uint16(v))
}

// PowerboxDescriptor_List is a list of PowerboxDescriptor.
type PowerboxDescriptor_List struct{ capnp.List }

// NewPowerboxDescriptor creates a new list of PowerboxDescriptor.
func NewPowerboxDescriptor_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return PowerboxDescriptor_List{}, err
	}
	return PowerboxDescriptor_List{l}, nil
}

func (s PowerboxDescriptor_List) At(i int) PowerboxDescriptor {
	return PowerboxDescriptor{s.List.Struct(i)}
}
func (s PowerboxDescriptor_List) Set(i int, v PowerboxDescriptor) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Promise is a wrapper for a PowerboxDescriptor promised by a client call.
type PowerboxDescriptor_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Promise) Struct() (PowerboxDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor{s}, err
}

type PowerboxDescriptor_Tag struct{ capnp.Struct }

func NewPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return PowerboxDescriptor_Tag{}, err
	}
	return PowerboxDescriptor_Tag{st}, nil
}

func NewRootPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return PowerboxDescriptor_Tag{}, err
	}
	return PowerboxDescriptor_Tag{st}, nil
}

func ReadRootPowerboxDescriptor_Tag(msg *capnp.Message) (PowerboxDescriptor_Tag, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return PowerboxDescriptor_Tag{}, err
	}
	return PowerboxDescriptor_Tag{root.Struct()}, nil
}

func (s PowerboxDescriptor_Tag) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s PowerboxDescriptor_Tag) SetId(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s PowerboxDescriptor_Tag) Value() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s PowerboxDescriptor_Tag) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor_Tag) ValuePtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s PowerboxDescriptor_Tag) SetValue(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s PowerboxDescriptor_Tag) SetValuePtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// PowerboxDescriptor_Tag_List is a list of PowerboxDescriptor_Tag.
type PowerboxDescriptor_Tag_List struct{ capnp.List }

// NewPowerboxDescriptor_Tag creates a new list of PowerboxDescriptor_Tag.
func NewPowerboxDescriptor_Tag_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return PowerboxDescriptor_Tag_List{}, err
	}
	return PowerboxDescriptor_Tag_List{l}, nil
}

func (s PowerboxDescriptor_Tag_List) At(i int) PowerboxDescriptor_Tag {
	return PowerboxDescriptor_Tag{s.List.Struct(i)}
}
func (s PowerboxDescriptor_Tag_List) Set(i int, v PowerboxDescriptor_Tag) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Tag_Promise is a wrapper for a PowerboxDescriptor_Tag promised by a client call.
type PowerboxDescriptor_Tag_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Tag_Promise) Struct() (PowerboxDescriptor_Tag, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor_Tag{s}, err
}

func (p PowerboxDescriptor_Tag_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type PowerboxDescriptor_MatchQuality uint16

// Values of PowerboxDescriptor_MatchQuality.
const (
	PowerboxDescriptor_MatchQuality_preferred    PowerboxDescriptor_MatchQuality = 1
	PowerboxDescriptor_MatchQuality_acceptable   PowerboxDescriptor_MatchQuality = 0
	PowerboxDescriptor_MatchQuality_unacceptable PowerboxDescriptor_MatchQuality = 2
)

// String returns the enum's constant name.
func (c PowerboxDescriptor_MatchQuality) String() string {
	switch c {
	case PowerboxDescriptor_MatchQuality_preferred:
		return "preferred"
	case PowerboxDescriptor_MatchQuality_acceptable:
		return "acceptable"
	case PowerboxDescriptor_MatchQuality_unacceptable:
		return "unacceptable"

	default:
		return ""
	}
}

// PowerboxDescriptor_MatchQualityFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PowerboxDescriptor_MatchQualityFromString(c string) PowerboxDescriptor_MatchQuality {
	switch c {
	case "preferred":
		return PowerboxDescriptor_MatchQuality_preferred
	case "acceptable":
		return PowerboxDescriptor_MatchQuality_acceptable
	case "unacceptable":
		return PowerboxDescriptor_MatchQuality_unacceptable

	default:
		return 0
	}
}

type PowerboxDescriptor_MatchQuality_List struct{ capnp.List }

func NewPowerboxDescriptor_MatchQuality_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_MatchQuality_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return PowerboxDescriptor_MatchQuality_List{}, err
	}
	return PowerboxDescriptor_MatchQuality_List{l.List}, nil
}

func (l PowerboxDescriptor_MatchQuality_List) At(i int) PowerboxDescriptor_MatchQuality {
	ul := capnp.UInt16List{List: l.List}
	return PowerboxDescriptor_MatchQuality(ul.At(i))
}

func (l PowerboxDescriptor_MatchQuality_List) Set(i int, v PowerboxDescriptor_MatchQuality) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type PowerboxDisplayInfo struct{ capnp.Struct }

func NewPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	return PowerboxDisplayInfo{st}, nil
}

func NewRootPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	return PowerboxDisplayInfo{st}, nil
}

func ReadRootPowerboxDisplayInfo(msg *capnp.Message) (PowerboxDisplayInfo, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	return PowerboxDisplayInfo{root.Struct()}, nil
}

func (s PowerboxDisplayInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s PowerboxDisplayInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetTitle(v util.LocalizedText) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s PowerboxDisplayInfo) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetVerbPhrase(v util.LocalizedText) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewVerbPhrase() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s PowerboxDisplayInfo) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetDescription(v util.LocalizedText) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// PowerboxDisplayInfo_List is a list of PowerboxDisplayInfo.
type PowerboxDisplayInfo_List struct{ capnp.List }

// NewPowerboxDisplayInfo creates a new list of PowerboxDisplayInfo.
func NewPowerboxDisplayInfo_List(s *capnp.Segment, sz int32) (PowerboxDisplayInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return PowerboxDisplayInfo_List{}, err
	}
	return PowerboxDisplayInfo_List{l}, nil
}

func (s PowerboxDisplayInfo_List) At(i int) PowerboxDisplayInfo {
	return PowerboxDisplayInfo{s.List.Struct(i)}
}
func (s PowerboxDisplayInfo_List) Set(i int, v PowerboxDisplayInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDisplayInfo_Promise is a wrapper for a PowerboxDisplayInfo promised by a client call.
type PowerboxDisplayInfo_Promise struct{ *capnp.Pipeline }

func (p PowerboxDisplayInfo_Promise) Struct() (PowerboxDisplayInfo, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDisplayInfo{s}, err
}

func (p PowerboxDisplayInfo_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PowerboxDisplayInfo_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p PowerboxDisplayInfo_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SandstormApi struct{ Client capnp.Client }

func (c SandstormApi) DeprecatedPublish(ctx context.Context, params func(SandstormApi_deprecatedPublish_Params) error, opts ...capnp.CallOption) SandstormApi_deprecatedPublish_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deprecatedPublish_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      0,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedPublish",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deprecatedPublish_Params{Struct: s}) }
	}
	return SandstormApi_deprecatedPublish_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) DeprecatedRegisterAction(ctx context.Context, params func(SandstormApi_deprecatedRegisterAction_Params) error, opts ...capnp.CallOption) SandstormApi_deprecatedRegisterAction_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deprecatedRegisterAction_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      1,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedRegisterAction",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deprecatedRegisterAction_Params{Struct: s}) }
	}
	return SandstormApi_deprecatedRegisterAction_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) ShareCap(ctx context.Context, params func(SandstormApi_shareCap_Params) error, opts ...capnp.CallOption) SandstormApi_shareCap_Results_Promise {
	if c.Client == nil {
		return SandstormApi_shareCap_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      2,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareCap",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_shareCap_Params{Struct: s}) }
	}
	return SandstormApi_shareCap_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) ShareView(ctx context.Context, params func(SandstormApi_shareView_Params) error, opts ...capnp.CallOption) SandstormApi_shareView_Results_Promise {
	if c.Client == nil {
		return SandstormApi_shareView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      3,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_shareView_Params{Struct: s}) }
	}
	return SandstormApi_shareView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) Restore(ctx context.Context, params func(SandstormApi_restore_Params) error, opts ...capnp.CallOption) SandstormApi_restore_Results_Promise {
	if c.Client == nil {
		return SandstormApi_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      4,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_restore_Params{Struct: s}) }
	}
	return SandstormApi_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) Drop(ctx context.Context, params func(SandstormApi_drop_Params) error, opts ...capnp.CallOption) SandstormApi_drop_Results_Promise {
	if c.Client == nil {
		return SandstormApi_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      5,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_drop_Params{Struct: s}) }
	}
	return SandstormApi_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) Deleted(ctx context.Context, params func(SandstormApi_deleted_Params) error, opts ...capnp.CallOption) SandstormApi_deleted_Results_Promise {
	if c.Client == nil {
		return SandstormApi_deleted_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      6,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deleted",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_deleted_Params{Struct: s}) }
	}
	return SandstormApi_deleted_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) StayAwake(ctx context.Context, params func(SandstormApi_stayAwake_Params) error, opts ...capnp.CallOption) SandstormApi_stayAwake_Results_Promise {
	if c.Client == nil {
		return SandstormApi_stayAwake_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      7,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "stayAwake",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_stayAwake_Params{Struct: s}) }
	}
	return SandstormApi_stayAwake_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SandstormApi) Save(ctx context.Context, params func(SandstormApi_save_Params) error, opts ...capnp.CallOption) SandstormApi_save_Results_Promise {
	if c.Client == nil {
		return SandstormApi_save_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      8,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SandstormApi_save_Params{Struct: s}) }
	}
	return SandstormApi_save_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SandstormApi_Server interface {
	DeprecatedPublish(SandstormApi_deprecatedPublish) error

	DeprecatedRegisterAction(SandstormApi_deprecatedRegisterAction) error

	ShareCap(SandstormApi_shareCap) error

	ShareView(SandstormApi_shareView) error

	Restore(SandstormApi_restore) error

	Drop(SandstormApi_drop) error

	Deleted(SandstormApi_deleted) error

	StayAwake(SandstormApi_stayAwake) error

	Save(SandstormApi_save) error
}

func SandstormApi_ServerToClient(s SandstormApi_Server) SandstormApi {
	c, _ := s.(server.Closer)
	return SandstormApi{Client: server.New(SandstormApi_Methods(nil, s), c)}
}

func SandstormApi_Methods(methods []server.Method, s SandstormApi_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 9)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      0,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedPublish",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deprecatedPublish{c, opts, SandstormApi_deprecatedPublish_Params{Struct: p}, SandstormApi_deprecatedPublish_Results{Struct: r}}
			return s.DeprecatedPublish(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      1,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deprecatedRegisterAction",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deprecatedRegisterAction{c, opts, SandstormApi_deprecatedRegisterAction_Params{Struct: p}, SandstormApi_deprecatedRegisterAction_Results{Struct: r}}
			return s.DeprecatedRegisterAction(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      2,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareCap",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_shareCap{c, opts, SandstormApi_shareCap_Params{Struct: p}, SandstormApi_shareCap_Results{Struct: r}}
			return s.ShareCap(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      3,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "shareView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_shareView{c, opts, SandstormApi_shareView_Params{Struct: p}, SandstormApi_shareView_Results{Struct: r}}
			return s.ShareView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      4,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_restore{c, opts, SandstormApi_restore_Params{Struct: p}, SandstormApi_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      5,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_drop{c, opts, SandstormApi_drop_Params{Struct: p}, SandstormApi_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      6,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "deleted",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_deleted{c, opts, SandstormApi_deleted_Params{Struct: p}, SandstormApi_deleted_Results{Struct: r}}
			return s.Deleted(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      7,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "stayAwake",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_stayAwake{c, opts, SandstormApi_stayAwake_Params{Struct: p}, SandstormApi_stayAwake_Results{Struct: r}}
			return s.StayAwake(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xd2654fcf2a7002cb,
			MethodID:      8,
			InterfaceName: "grain.capnp:SandstormApi",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SandstormApi_save{c, opts, SandstormApi_save_Params{Struct: p}, SandstormApi_save_Results{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SandstormApi_deprecatedPublish holds the arguments for a server call to SandstormApi.deprecatedPublish.
type SandstormApi_deprecatedPublish struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deprecatedPublish_Params
	Results SandstormApi_deprecatedPublish_Results
}

// SandstormApi_deprecatedRegisterAction holds the arguments for a server call to SandstormApi.deprecatedRegisterAction.
type SandstormApi_deprecatedRegisterAction struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deprecatedRegisterAction_Params
	Results SandstormApi_deprecatedRegisterAction_Results
}

// SandstormApi_shareCap holds the arguments for a server call to SandstormApi.shareCap.
type SandstormApi_shareCap struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_shareCap_Params
	Results SandstormApi_shareCap_Results
}

// SandstormApi_shareView holds the arguments for a server call to SandstormApi.shareView.
type SandstormApi_shareView struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_shareView_Params
	Results SandstormApi_shareView_Results
}

// SandstormApi_restore holds the arguments for a server call to SandstormApi.restore.
type SandstormApi_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_restore_Params
	Results SandstormApi_restore_Results
}

// SandstormApi_drop holds the arguments for a server call to SandstormApi.drop.
type SandstormApi_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_drop_Params
	Results SandstormApi_drop_Results
}

// SandstormApi_deleted holds the arguments for a server call to SandstormApi.deleted.
type SandstormApi_deleted struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_deleted_Params
	Results SandstormApi_deleted_Results
}

// SandstormApi_stayAwake holds the arguments for a server call to SandstormApi.stayAwake.
type SandstormApi_stayAwake struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_stayAwake_Params
	Results SandstormApi_stayAwake_Results
}

// SandstormApi_save holds the arguments for a server call to SandstormApi.save.
type SandstormApi_save struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SandstormApi_save_Params
	Results SandstormApi_save_Results
}

type SandstormApi_deprecatedPublish_Params struct{ capnp.Struct }

func NewSandstormApi_deprecatedPublish_Params(s *capnp.Segment) (SandstormApi_deprecatedPublish_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedPublish_Params{}, err
	}
	return SandstormApi_deprecatedPublish_Params{st}, nil
}

func NewRootSandstormApi_deprecatedPublish_Params(s *capnp.Segment) (SandstormApi_deprecatedPublish_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedPublish_Params{}, err
	}
	return SandstormApi_deprecatedPublish_Params{st}, nil
}

func ReadRootSandstormApi_deprecatedPublish_Params(msg *capnp.Message) (SandstormApi_deprecatedPublish_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deprecatedPublish_Params{}, err
	}
	return SandstormApi_deprecatedPublish_Params{root.Struct()}, nil
}

// SandstormApi_deprecatedPublish_Params_List is a list of SandstormApi_deprecatedPublish_Params.
type SandstormApi_deprecatedPublish_Params_List struct{ capnp.List }

// NewSandstormApi_deprecatedPublish_Params creates a new list of SandstormApi_deprecatedPublish_Params.
func NewSandstormApi_deprecatedPublish_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedPublish_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_deprecatedPublish_Params_List{}, err
	}
	return SandstormApi_deprecatedPublish_Params_List{l}, nil
}

func (s SandstormApi_deprecatedPublish_Params_List) At(i int) SandstormApi_deprecatedPublish_Params {
	return SandstormApi_deprecatedPublish_Params{s.List.Struct(i)}
}
func (s SandstormApi_deprecatedPublish_Params_List) Set(i int, v SandstormApi_deprecatedPublish_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedPublish_Params_Promise is a wrapper for a SandstormApi_deprecatedPublish_Params promised by a client call.
type SandstormApi_deprecatedPublish_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedPublish_Params_Promise) Struct() (SandstormApi_deprecatedPublish_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedPublish_Params{s}, err
}

type SandstormApi_deprecatedPublish_Results struct{ capnp.Struct }

func NewSandstormApi_deprecatedPublish_Results(s *capnp.Segment) (SandstormApi_deprecatedPublish_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedPublish_Results{}, err
	}
	return SandstormApi_deprecatedPublish_Results{st}, nil
}

func NewRootSandstormApi_deprecatedPublish_Results(s *capnp.Segment) (SandstormApi_deprecatedPublish_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedPublish_Results{}, err
	}
	return SandstormApi_deprecatedPublish_Results{st}, nil
}

func ReadRootSandstormApi_deprecatedPublish_Results(msg *capnp.Message) (SandstormApi_deprecatedPublish_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deprecatedPublish_Results{}, err
	}
	return SandstormApi_deprecatedPublish_Results{root.Struct()}, nil
}

// SandstormApi_deprecatedPublish_Results_List is a list of SandstormApi_deprecatedPublish_Results.
type SandstormApi_deprecatedPublish_Results_List struct{ capnp.List }

// NewSandstormApi_deprecatedPublish_Results creates a new list of SandstormApi_deprecatedPublish_Results.
func NewSandstormApi_deprecatedPublish_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedPublish_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_deprecatedPublish_Results_List{}, err
	}
	return SandstormApi_deprecatedPublish_Results_List{l}, nil
}

func (s SandstormApi_deprecatedPublish_Results_List) At(i int) SandstormApi_deprecatedPublish_Results {
	return SandstormApi_deprecatedPublish_Results{s.List.Struct(i)}
}
func (s SandstormApi_deprecatedPublish_Results_List) Set(i int, v SandstormApi_deprecatedPublish_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedPublish_Results_Promise is a wrapper for a SandstormApi_deprecatedPublish_Results promised by a client call.
type SandstormApi_deprecatedPublish_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedPublish_Results_Promise) Struct() (SandstormApi_deprecatedPublish_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedPublish_Results{s}, err
}

type SandstormApi_deprecatedRegisterAction_Params struct{ capnp.Struct }

func NewSandstormApi_deprecatedRegisterAction_Params(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Params{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Params{st}, nil
}

func NewRootSandstormApi_deprecatedRegisterAction_Params(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Params{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Params{st}, nil
}

func ReadRootSandstormApi_deprecatedRegisterAction_Params(msg *capnp.Message) (SandstormApi_deprecatedRegisterAction_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Params{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Params{root.Struct()}, nil
}

// SandstormApi_deprecatedRegisterAction_Params_List is a list of SandstormApi_deprecatedRegisterAction_Params.
type SandstormApi_deprecatedRegisterAction_Params_List struct{ capnp.List }

// NewSandstormApi_deprecatedRegisterAction_Params creates a new list of SandstormApi_deprecatedRegisterAction_Params.
func NewSandstormApi_deprecatedRegisterAction_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedRegisterAction_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Params_List{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Params_List{l}, nil
}

func (s SandstormApi_deprecatedRegisterAction_Params_List) At(i int) SandstormApi_deprecatedRegisterAction_Params {
	return SandstormApi_deprecatedRegisterAction_Params{s.List.Struct(i)}
}
func (s SandstormApi_deprecatedRegisterAction_Params_List) Set(i int, v SandstormApi_deprecatedRegisterAction_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedRegisterAction_Params_Promise is a wrapper for a SandstormApi_deprecatedRegisterAction_Params promised by a client call.
type SandstormApi_deprecatedRegisterAction_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedRegisterAction_Params_Promise) Struct() (SandstormApi_deprecatedRegisterAction_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedRegisterAction_Params{s}, err
}

type SandstormApi_deprecatedRegisterAction_Results struct{ capnp.Struct }

func NewSandstormApi_deprecatedRegisterAction_Results(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Results{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Results{st}, nil
}

func NewRootSandstormApi_deprecatedRegisterAction_Results(s *capnp.Segment) (SandstormApi_deprecatedRegisterAction_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Results{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Results{st}, nil
}

func ReadRootSandstormApi_deprecatedRegisterAction_Results(msg *capnp.Message) (SandstormApi_deprecatedRegisterAction_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Results{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Results{root.Struct()}, nil
}

// SandstormApi_deprecatedRegisterAction_Results_List is a list of SandstormApi_deprecatedRegisterAction_Results.
type SandstormApi_deprecatedRegisterAction_Results_List struct{ capnp.List }

// NewSandstormApi_deprecatedRegisterAction_Results creates a new list of SandstormApi_deprecatedRegisterAction_Results.
func NewSandstormApi_deprecatedRegisterAction_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deprecatedRegisterAction_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_deprecatedRegisterAction_Results_List{}, err
	}
	return SandstormApi_deprecatedRegisterAction_Results_List{l}, nil
}

func (s SandstormApi_deprecatedRegisterAction_Results_List) At(i int) SandstormApi_deprecatedRegisterAction_Results {
	return SandstormApi_deprecatedRegisterAction_Results{s.List.Struct(i)}
}
func (s SandstormApi_deprecatedRegisterAction_Results_List) Set(i int, v SandstormApi_deprecatedRegisterAction_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deprecatedRegisterAction_Results_Promise is a wrapper for a SandstormApi_deprecatedRegisterAction_Results promised by a client call.
type SandstormApi_deprecatedRegisterAction_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deprecatedRegisterAction_Results_Promise) Struct() (SandstormApi_deprecatedRegisterAction_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deprecatedRegisterAction_Results{s}, err
}

type SandstormApi_shareCap_Params struct{ capnp.Struct }

func NewSandstormApi_shareCap_Params(s *capnp.Segment) (SandstormApi_shareCap_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareCap_Params{}, err
	}
	return SandstormApi_shareCap_Params{st}, nil
}

func NewRootSandstormApi_shareCap_Params(s *capnp.Segment) (SandstormApi_shareCap_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareCap_Params{}, err
	}
	return SandstormApi_shareCap_Params{st}, nil
}

func ReadRootSandstormApi_shareCap_Params(msg *capnp.Message) (SandstormApi_shareCap_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_shareCap_Params{}, err
	}
	return SandstormApi_shareCap_Params{root.Struct()}, nil
}

func (s SandstormApi_shareCap_Params) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SandstormApi_shareCap_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Params) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SandstormApi_shareCap_Params) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_shareCap_Params) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_shareCap_Params) DisplayInfo() (PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}

	return PowerboxDisplayInfo{Struct: p.Struct()}, nil

}

func (s SandstormApi_shareCap_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Params) SetDisplayInfo(v PowerboxDisplayInfo) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SandstormApi_shareCap_Params) NewDisplayInfo() (PowerboxDisplayInfo, error) {

	ss, err := NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SandstormApi_shareCap_Params_List is a list of SandstormApi_shareCap_Params.
type SandstormApi_shareCap_Params_List struct{ capnp.List }

// NewSandstormApi_shareCap_Params creates a new list of SandstormApi_shareCap_Params.
func NewSandstormApi_shareCap_Params_List(s *capnp.Segment, sz int32) (SandstormApi_shareCap_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_shareCap_Params_List{}, err
	}
	return SandstormApi_shareCap_Params_List{l}, nil
}

func (s SandstormApi_shareCap_Params_List) At(i int) SandstormApi_shareCap_Params {
	return SandstormApi_shareCap_Params{s.List.Struct(i)}
}
func (s SandstormApi_shareCap_Params_List) Set(i int, v SandstormApi_shareCap_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareCap_Params_Promise is a wrapper for a SandstormApi_shareCap_Params promised by a client call.
type SandstormApi_shareCap_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareCap_Params_Promise) Struct() (SandstormApi_shareCap_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareCap_Params{s}, err
}

func (p SandstormApi_shareCap_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_shareCap_Params_Promise) DisplayInfo() PowerboxDisplayInfo_Promise {
	return PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormApi_shareCap_Results struct{ capnp.Struct }

func NewSandstormApi_shareCap_Results(s *capnp.Segment) (SandstormApi_shareCap_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareCap_Results{}, err
	}
	return SandstormApi_shareCap_Results{st}, nil
}

func NewRootSandstormApi_shareCap_Results(s *capnp.Segment) (SandstormApi_shareCap_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareCap_Results{}, err
	}
	return SandstormApi_shareCap_Results{st}, nil
}

func ReadRootSandstormApi_shareCap_Results(msg *capnp.Message) (SandstormApi_shareCap_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_shareCap_Results{}, err
	}
	return SandstormApi_shareCap_Results{root.Struct()}, nil
}

func (s SandstormApi_shareCap_Results) SharedCap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SandstormApi_shareCap_Results) HasSharedCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Results) SharedCapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SandstormApi_shareCap_Results) SetSharedCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_shareCap_Results) SetSharedCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_shareCap_Results) Link() SharingLink {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return SharingLink{}
	}
	return SharingLink{Client: p.Interface().Client()}
}

func (s SandstormApi_shareCap_Results) HasLink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareCap_Results) SetLink(v SharingLink) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_shareCap_Results_List is a list of SandstormApi_shareCap_Results.
type SandstormApi_shareCap_Results_List struct{ capnp.List }

// NewSandstormApi_shareCap_Results creates a new list of SandstormApi_shareCap_Results.
func NewSandstormApi_shareCap_Results_List(s *capnp.Segment, sz int32) (SandstormApi_shareCap_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_shareCap_Results_List{}, err
	}
	return SandstormApi_shareCap_Results_List{l}, nil
}

func (s SandstormApi_shareCap_Results_List) At(i int) SandstormApi_shareCap_Results {
	return SandstormApi_shareCap_Results{s.List.Struct(i)}
}
func (s SandstormApi_shareCap_Results_List) Set(i int, v SandstormApi_shareCap_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareCap_Results_Promise is a wrapper for a SandstormApi_shareCap_Results promised by a client call.
type SandstormApi_shareCap_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareCap_Results_Promise) Struct() (SandstormApi_shareCap_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareCap_Results{s}, err
}

func (p SandstormApi_shareCap_Results_Promise) SharedCap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_shareCap_Results_Promise) Link() SharingLink {
	return SharingLink{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_shareView_Params struct{ capnp.Struct }

func NewSandstormApi_shareView_Params(s *capnp.Segment) (SandstormApi_shareView_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_shareView_Params{}, err
	}
	return SandstormApi_shareView_Params{st}, nil
}

func NewRootSandstormApi_shareView_Params(s *capnp.Segment) (SandstormApi_shareView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_shareView_Params{}, err
	}
	return SandstormApi_shareView_Params{st}, nil
}

func ReadRootSandstormApi_shareView_Params(msg *capnp.Message) (SandstormApi_shareView_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_shareView_Params{}, err
	}
	return SandstormApi_shareView_Params{root.Struct()}, nil
}

func (s SandstormApi_shareView_Params) View() UiView {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiView{}
	}
	return UiView{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Params) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Params) SetView(v UiView) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormApi_shareView_Params_List is a list of SandstormApi_shareView_Params.
type SandstormApi_shareView_Params_List struct{ capnp.List }

// NewSandstormApi_shareView_Params creates a new list of SandstormApi_shareView_Params.
func NewSandstormApi_shareView_Params_List(s *capnp.Segment, sz int32) (SandstormApi_shareView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_shareView_Params_List{}, err
	}
	return SandstormApi_shareView_Params_List{l}, nil
}

func (s SandstormApi_shareView_Params_List) At(i int) SandstormApi_shareView_Params {
	return SandstormApi_shareView_Params{s.List.Struct(i)}
}
func (s SandstormApi_shareView_Params_List) Set(i int, v SandstormApi_shareView_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareView_Params_Promise is a wrapper for a SandstormApi_shareView_Params promised by a client call.
type SandstormApi_shareView_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareView_Params_Promise) Struct() (SandstormApi_shareView_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareView_Params{s}, err
}

func (p SandstormApi_shareView_Params_Promise) View() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormApi_shareView_Results struct{ capnp.Struct }

func NewSandstormApi_shareView_Results(s *capnp.Segment) (SandstormApi_shareView_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareView_Results{}, err
	}
	return SandstormApi_shareView_Results{st}, nil
}

func NewRootSandstormApi_shareView_Results(s *capnp.Segment) (SandstormApi_shareView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_shareView_Results{}, err
	}
	return SandstormApi_shareView_Results{st}, nil
}

func ReadRootSandstormApi_shareView_Results(msg *capnp.Message) (SandstormApi_shareView_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_shareView_Results{}, err
	}
	return SandstormApi_shareView_Results{root.Struct()}, nil
}

func (s SandstormApi_shareView_Results) SharedView() UiView {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiView{}
	}
	return UiView{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Results) HasSharedView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Results) SetSharedView(v UiView) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s SandstormApi_shareView_Results) Link() ViewSharingLink {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return ViewSharingLink{}
	}
	return ViewSharingLink{Client: p.Interface().Client()}
}

func (s SandstormApi_shareView_Results) HasLink() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_shareView_Results) SetLink(v ViewSharingLink) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_shareView_Results_List is a list of SandstormApi_shareView_Results.
type SandstormApi_shareView_Results_List struct{ capnp.List }

// NewSandstormApi_shareView_Results creates a new list of SandstormApi_shareView_Results.
func NewSandstormApi_shareView_Results_List(s *capnp.Segment, sz int32) (SandstormApi_shareView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_shareView_Results_List{}, err
	}
	return SandstormApi_shareView_Results_List{l}, nil
}

func (s SandstormApi_shareView_Results_List) At(i int) SandstormApi_shareView_Results {
	return SandstormApi_shareView_Results{s.List.Struct(i)}
}
func (s SandstormApi_shareView_Results_List) Set(i int, v SandstormApi_shareView_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_shareView_Results_Promise is a wrapper for a SandstormApi_shareView_Results promised by a client call.
type SandstormApi_shareView_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_shareView_Results_Promise) Struct() (SandstormApi_shareView_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_shareView_Results{s}, err
}

func (p SandstormApi_shareView_Results_Promise) SharedView() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p SandstormApi_shareView_Results_Promise) Link() ViewSharingLink {
	return ViewSharingLink{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_restore_Params struct{ capnp.Struct }

func NewSandstormApi_restore_Params(s *capnp.Segment) (SandstormApi_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_restore_Params{}, err
	}
	return SandstormApi_restore_Params{st}, nil
}

func NewRootSandstormApi_restore_Params(s *capnp.Segment) (SandstormApi_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_restore_Params{}, err
	}
	return SandstormApi_restore_Params{st}, nil
}

func ReadRootSandstormApi_restore_Params(msg *capnp.Message) (SandstormApi_restore_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_restore_Params{}, err
	}
	return SandstormApi_restore_Params{root.Struct()}, nil
}

func (s SandstormApi_restore_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

}

func (s SandstormApi_restore_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_restore_Params) SetToken(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s SandstormApi_restore_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s SandstormApi_restore_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_restore_Params) SetRequiredPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// SandstormApi_restore_Params_List is a list of SandstormApi_restore_Params.
type SandstormApi_restore_Params_List struct{ capnp.List }

// NewSandstormApi_restore_Params creates a new list of SandstormApi_restore_Params.
func NewSandstormApi_restore_Params_List(s *capnp.Segment, sz int32) (SandstormApi_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_restore_Params_List{}, err
	}
	return SandstormApi_restore_Params_List{l}, nil
}

func (s SandstormApi_restore_Params_List) At(i int) SandstormApi_restore_Params {
	return SandstormApi_restore_Params{s.List.Struct(i)}
}
func (s SandstormApi_restore_Params_List) Set(i int, v SandstormApi_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_restore_Params_Promise is a wrapper for a SandstormApi_restore_Params promised by a client call.
type SandstormApi_restore_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_restore_Params_Promise) Struct() (SandstormApi_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_restore_Params{s}, err
}

type SandstormApi_restore_Results struct{ capnp.Struct }

func NewSandstormApi_restore_Results(s *capnp.Segment) (SandstormApi_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_restore_Results{}, err
	}
	return SandstormApi_restore_Results{st}, nil
}

func NewRootSandstormApi_restore_Results(s *capnp.Segment) (SandstormApi_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_restore_Results{}, err
	}
	return SandstormApi_restore_Results{st}, nil
}

func ReadRootSandstormApi_restore_Results(msg *capnp.Message) (SandstormApi_restore_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_restore_Results{}, err
	}
	return SandstormApi_restore_Results{root.Struct()}, nil
}

func (s SandstormApi_restore_Results) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SandstormApi_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_restore_Results) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SandstormApi_restore_Results) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_restore_Results) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// SandstormApi_restore_Results_List is a list of SandstormApi_restore_Results.
type SandstormApi_restore_Results_List struct{ capnp.List }

// NewSandstormApi_restore_Results creates a new list of SandstormApi_restore_Results.
func NewSandstormApi_restore_Results_List(s *capnp.Segment, sz int32) (SandstormApi_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_restore_Results_List{}, err
	}
	return SandstormApi_restore_Results_List{l}, nil
}

func (s SandstormApi_restore_Results_List) At(i int) SandstormApi_restore_Results {
	return SandstormApi_restore_Results{s.List.Struct(i)}
}
func (s SandstormApi_restore_Results_List) Set(i int, v SandstormApi_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_restore_Results_Promise is a wrapper for a SandstormApi_restore_Results promised by a client call.
type SandstormApi_restore_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_restore_Results_Promise) Struct() (SandstormApi_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_restore_Results{s}, err
}

func (p SandstormApi_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormApi_drop_Params struct{ capnp.Struct }

func NewSandstormApi_drop_Params(s *capnp.Segment) (SandstormApi_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_drop_Params{}, err
	}
	return SandstormApi_drop_Params{st}, nil
}

func NewRootSandstormApi_drop_Params(s *capnp.Segment) (SandstormApi_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_drop_Params{}, err
	}
	return SandstormApi_drop_Params{st}, nil
}

func ReadRootSandstormApi_drop_Params(msg *capnp.Message) (SandstormApi_drop_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_drop_Params{}, err
	}
	return SandstormApi_drop_Params{root.Struct()}, nil
}

func (s SandstormApi_drop_Params) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

}

func (s SandstormApi_drop_Params) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_drop_Params) SetToken(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_drop_Params_List is a list of SandstormApi_drop_Params.
type SandstormApi_drop_Params_List struct{ capnp.List }

// NewSandstormApi_drop_Params creates a new list of SandstormApi_drop_Params.
func NewSandstormApi_drop_Params_List(s *capnp.Segment, sz int32) (SandstormApi_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_drop_Params_List{}, err
	}
	return SandstormApi_drop_Params_List{l}, nil
}

func (s SandstormApi_drop_Params_List) At(i int) SandstormApi_drop_Params {
	return SandstormApi_drop_Params{s.List.Struct(i)}
}
func (s SandstormApi_drop_Params_List) Set(i int, v SandstormApi_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_drop_Params_Promise is a wrapper for a SandstormApi_drop_Params promised by a client call.
type SandstormApi_drop_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_drop_Params_Promise) Struct() (SandstormApi_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_drop_Params{s}, err
}

type SandstormApi_drop_Results struct{ capnp.Struct }

func NewSandstormApi_drop_Results(s *capnp.Segment) (SandstormApi_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_drop_Results{}, err
	}
	return SandstormApi_drop_Results{st}, nil
}

func NewRootSandstormApi_drop_Results(s *capnp.Segment) (SandstormApi_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_drop_Results{}, err
	}
	return SandstormApi_drop_Results{st}, nil
}

func ReadRootSandstormApi_drop_Results(msg *capnp.Message) (SandstormApi_drop_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_drop_Results{}, err
	}
	return SandstormApi_drop_Results{root.Struct()}, nil
}

// SandstormApi_drop_Results_List is a list of SandstormApi_drop_Results.
type SandstormApi_drop_Results_List struct{ capnp.List }

// NewSandstormApi_drop_Results creates a new list of SandstormApi_drop_Results.
func NewSandstormApi_drop_Results_List(s *capnp.Segment, sz int32) (SandstormApi_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_drop_Results_List{}, err
	}
	return SandstormApi_drop_Results_List{l}, nil
}

func (s SandstormApi_drop_Results_List) At(i int) SandstormApi_drop_Results {
	return SandstormApi_drop_Results{s.List.Struct(i)}
}
func (s SandstormApi_drop_Results_List) Set(i int, v SandstormApi_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_drop_Results_Promise is a wrapper for a SandstormApi_drop_Results promised by a client call.
type SandstormApi_drop_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_drop_Results_Promise) Struct() (SandstormApi_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_drop_Results{s}, err
}

type SandstormApi_deleted_Params struct{ capnp.Struct }

func NewSandstormApi_deleted_Params(s *capnp.Segment) (SandstormApi_deleted_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_deleted_Params{}, err
	}
	return SandstormApi_deleted_Params{st}, nil
}

func NewRootSandstormApi_deleted_Params(s *capnp.Segment) (SandstormApi_deleted_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_deleted_Params{}, err
	}
	return SandstormApi_deleted_Params{st}, nil
}

func ReadRootSandstormApi_deleted_Params(msg *capnp.Message) (SandstormApi_deleted_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deleted_Params{}, err
	}
	return SandstormApi_deleted_Params{root.Struct()}, nil
}

func (s SandstormApi_deleted_Params) Ref() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SandstormApi_deleted_Params) HasRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_deleted_Params) RefPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SandstormApi_deleted_Params) SetRef(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_deleted_Params) SetRefPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// SandstormApi_deleted_Params_List is a list of SandstormApi_deleted_Params.
type SandstormApi_deleted_Params_List struct{ capnp.List }

// NewSandstormApi_deleted_Params creates a new list of SandstormApi_deleted_Params.
func NewSandstormApi_deleted_Params_List(s *capnp.Segment, sz int32) (SandstormApi_deleted_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_deleted_Params_List{}, err
	}
	return SandstormApi_deleted_Params_List{l}, nil
}

func (s SandstormApi_deleted_Params_List) At(i int) SandstormApi_deleted_Params {
	return SandstormApi_deleted_Params{s.List.Struct(i)}
}
func (s SandstormApi_deleted_Params_List) Set(i int, v SandstormApi_deleted_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deleted_Params_Promise is a wrapper for a SandstormApi_deleted_Params promised by a client call.
type SandstormApi_deleted_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deleted_Params_Promise) Struct() (SandstormApi_deleted_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deleted_Params{s}, err
}

func (p SandstormApi_deleted_Params_Promise) Ref() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SandstormApi_deleted_Results struct{ capnp.Struct }

func NewSandstormApi_deleted_Results(s *capnp.Segment) (SandstormApi_deleted_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deleted_Results{}, err
	}
	return SandstormApi_deleted_Results{st}, nil
}

func NewRootSandstormApi_deleted_Results(s *capnp.Segment) (SandstormApi_deleted_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SandstormApi_deleted_Results{}, err
	}
	return SandstormApi_deleted_Results{st}, nil
}

func ReadRootSandstormApi_deleted_Results(msg *capnp.Message) (SandstormApi_deleted_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_deleted_Results{}, err
	}
	return SandstormApi_deleted_Results{root.Struct()}, nil
}

// SandstormApi_deleted_Results_List is a list of SandstormApi_deleted_Results.
type SandstormApi_deleted_Results_List struct{ capnp.List }

// NewSandstormApi_deleted_Results creates a new list of SandstormApi_deleted_Results.
func NewSandstormApi_deleted_Results_List(s *capnp.Segment, sz int32) (SandstormApi_deleted_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SandstormApi_deleted_Results_List{}, err
	}
	return SandstormApi_deleted_Results_List{l}, nil
}

func (s SandstormApi_deleted_Results_List) At(i int) SandstormApi_deleted_Results {
	return SandstormApi_deleted_Results{s.List.Struct(i)}
}
func (s SandstormApi_deleted_Results_List) Set(i int, v SandstormApi_deleted_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_deleted_Results_Promise is a wrapper for a SandstormApi_deleted_Results promised by a client call.
type SandstormApi_deleted_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_deleted_Results_Promise) Struct() (SandstormApi_deleted_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_deleted_Results{s}, err
}

type SandstormApi_stayAwake_Params struct{ capnp.Struct }

func NewSandstormApi_stayAwake_Params(s *capnp.Segment) (SandstormApi_stayAwake_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_stayAwake_Params{}, err
	}
	return SandstormApi_stayAwake_Params{st}, nil
}

func NewRootSandstormApi_stayAwake_Params(s *capnp.Segment) (SandstormApi_stayAwake_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_stayAwake_Params{}, err
	}
	return SandstormApi_stayAwake_Params{st}, nil
}

func ReadRootSandstormApi_stayAwake_Params(msg *capnp.Message) (SandstormApi_stayAwake_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_stayAwake_Params{}, err
	}
	return SandstormApi_stayAwake_Params{root.Struct()}, nil
}

func (s SandstormApi_stayAwake_Params) DisplayInfo() (NotificationDisplayInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return NotificationDisplayInfo{}, err
	}

	return NotificationDisplayInfo{Struct: p.Struct()}, nil

}

func (s SandstormApi_stayAwake_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Params) SetDisplayInfo(v NotificationDisplayInfo) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated NotificationDisplayInfo struct, preferring placement in s's segment.
func (s SandstormApi_stayAwake_Params) NewDisplayInfo() (NotificationDisplayInfo, error) {

	ss, err := NewNotificationDisplayInfo(s.Struct.Segment())
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s SandstormApi_stayAwake_Params) Notification() OngoingNotification {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return OngoingNotification{}
	}
	return OngoingNotification{Client: p.Interface().Client()}
}

func (s SandstormApi_stayAwake_Params) HasNotification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Params) SetNotification(v OngoingNotification) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

// SandstormApi_stayAwake_Params_List is a list of SandstormApi_stayAwake_Params.
type SandstormApi_stayAwake_Params_List struct{ capnp.List }

// NewSandstormApi_stayAwake_Params creates a new list of SandstormApi_stayAwake_Params.
func NewSandstormApi_stayAwake_Params_List(s *capnp.Segment, sz int32) (SandstormApi_stayAwake_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_stayAwake_Params_List{}, err
	}
	return SandstormApi_stayAwake_Params_List{l}, nil
}

func (s SandstormApi_stayAwake_Params_List) At(i int) SandstormApi_stayAwake_Params {
	return SandstormApi_stayAwake_Params{s.List.Struct(i)}
}
func (s SandstormApi_stayAwake_Params_List) Set(i int, v SandstormApi_stayAwake_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_stayAwake_Params_Promise is a wrapper for a SandstormApi_stayAwake_Params promised by a client call.
type SandstormApi_stayAwake_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_stayAwake_Params_Promise) Struct() (SandstormApi_stayAwake_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_stayAwake_Params{s}, err
}

func (p SandstormApi_stayAwake_Params_Promise) DisplayInfo() NotificationDisplayInfo_Promise {
	return NotificationDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p SandstormApi_stayAwake_Params_Promise) Notification() OngoingNotification {
	return OngoingNotification{Client: p.Pipeline.GetPipeline(1).Client()}
}

type SandstormApi_stayAwake_Results struct{ capnp.Struct }

func NewSandstormApi_stayAwake_Results(s *capnp.Segment) (SandstormApi_stayAwake_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_stayAwake_Results{}, err
	}
	return SandstormApi_stayAwake_Results{st}, nil
}

func NewRootSandstormApi_stayAwake_Results(s *capnp.Segment) (SandstormApi_stayAwake_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_stayAwake_Results{}, err
	}
	return SandstormApi_stayAwake_Results{st}, nil
}

func ReadRootSandstormApi_stayAwake_Results(msg *capnp.Message) (SandstormApi_stayAwake_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_stayAwake_Results{}, err
	}
	return SandstormApi_stayAwake_Results{root.Struct()}, nil
}

func (s SandstormApi_stayAwake_Results) Handle() util.Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s SandstormApi_stayAwake_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_stayAwake_Results) SetHandle(v util.Handle) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SandstormApi_stayAwake_Results_List is a list of SandstormApi_stayAwake_Results.
type SandstormApi_stayAwake_Results_List struct{ capnp.List }

// NewSandstormApi_stayAwake_Results creates a new list of SandstormApi_stayAwake_Results.
func NewSandstormApi_stayAwake_Results_List(s *capnp.Segment, sz int32) (SandstormApi_stayAwake_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_stayAwake_Results_List{}, err
	}
	return SandstormApi_stayAwake_Results_List{l}, nil
}

func (s SandstormApi_stayAwake_Results_List) At(i int) SandstormApi_stayAwake_Results {
	return SandstormApi_stayAwake_Results{s.List.Struct(i)}
}
func (s SandstormApi_stayAwake_Results_List) Set(i int, v SandstormApi_stayAwake_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_stayAwake_Results_Promise is a wrapper for a SandstormApi_stayAwake_Results promised by a client call.
type SandstormApi_stayAwake_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_stayAwake_Results_Promise) Struct() (SandstormApi_stayAwake_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_stayAwake_Results{s}, err
}

func (p SandstormApi_stayAwake_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SandstormApi_save_Params struct{ capnp.Struct }

func NewSandstormApi_save_Params(s *capnp.Segment) (SandstormApi_save_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_save_Params{}, err
	}
	return SandstormApi_save_Params{st}, nil
}

func NewRootSandstormApi_save_Params(s *capnp.Segment) (SandstormApi_save_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SandstormApi_save_Params{}, err
	}
	return SandstormApi_save_Params{st}, nil
}

func ReadRootSandstormApi_save_Params(msg *capnp.Message) (SandstormApi_save_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_save_Params{}, err
	}
	return SandstormApi_save_Params{root.Struct()}, nil
}

func (s SandstormApi_save_Params) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SandstormApi_save_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Params) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SandstormApi_save_Params) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SandstormApi_save_Params) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SandstormApi_save_Params) Label() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s SandstormApi_save_Params) HasLabel() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Params) SetLabel(v util.LocalizedText) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s SandstormApi_save_Params) NewLabel() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SandstormApi_save_Params_List is a list of SandstormApi_save_Params.
type SandstormApi_save_Params_List struct{ capnp.List }

// NewSandstormApi_save_Params creates a new list of SandstormApi_save_Params.
func NewSandstormApi_save_Params_List(s *capnp.Segment, sz int32) (SandstormApi_save_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SandstormApi_save_Params_List{}, err
	}
	return SandstormApi_save_Params_List{l}, nil
}

func (s SandstormApi_save_Params_List) At(i int) SandstormApi_save_Params {
	return SandstormApi_save_Params{s.List.Struct(i)}
}
func (s SandstormApi_save_Params_List) Set(i int, v SandstormApi_save_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_save_Params_Promise is a wrapper for a SandstormApi_save_Params promised by a client call.
type SandstormApi_save_Params_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_save_Params_Promise) Struct() (SandstormApi_save_Params, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_save_Params{s}, err
}

func (p SandstormApi_save_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SandstormApi_save_Params_Promise) Label() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SandstormApi_save_Results struct{ capnp.Struct }

func NewSandstormApi_save_Results(s *capnp.Segment) (SandstormApi_save_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_save_Results{}, err
	}
	return SandstormApi_save_Results{st}, nil
}

func NewRootSandstormApi_save_Results(s *capnp.Segment) (SandstormApi_save_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SandstormApi_save_Results{}, err
	}
	return SandstormApi_save_Results{st}, nil
}

func ReadRootSandstormApi_save_Results(msg *capnp.Message) (SandstormApi_save_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SandstormApi_save_Results{}, err
	}
	return SandstormApi_save_Results{root.Struct()}, nil
}

func (s SandstormApi_save_Results) Token() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

}

func (s SandstormApi_save_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SandstormApi_save_Results) SetToken(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// SandstormApi_save_Results_List is a list of SandstormApi_save_Results.
type SandstormApi_save_Results_List struct{ capnp.List }

// NewSandstormApi_save_Results creates a new list of SandstormApi_save_Results.
func NewSandstormApi_save_Results_List(s *capnp.Segment, sz int32) (SandstormApi_save_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SandstormApi_save_Results_List{}, err
	}
	return SandstormApi_save_Results_List{l}, nil
}

func (s SandstormApi_save_Results_List) At(i int) SandstormApi_save_Results {
	return SandstormApi_save_Results{s.List.Struct(i)}
}
func (s SandstormApi_save_Results_List) Set(i int, v SandstormApi_save_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SandstormApi_save_Results_Promise is a wrapper for a SandstormApi_save_Results promised by a client call.
type SandstormApi_save_Results_Promise struct{ *capnp.Pipeline }

func (p SandstormApi_save_Results_Promise) Struct() (SandstormApi_save_Results, error) {
	s, err := p.Pipeline.Struct()
	return SandstormApi_save_Results{s}, err
}

type UiView struct{ Client capnp.Client }

func (c UiView) GetViewInfo(ctx context.Context, params func(UiView_getViewInfo_Params) error, opts ...capnp.CallOption) UiView_ViewInfo_Promise {
	if c.Client == nil {
		return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_getViewInfo_Params{Struct: s}) }
	}
	return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c UiView) NewSession(ctx context.Context, params func(UiView_newSession_Params) error, opts ...capnp.CallOption) UiView_newSession_Results_Promise {
	if c.Client == nil {
		return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newSession_Params{Struct: s}) }
	}
	return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c UiView) NewRequestSession(ctx context.Context, params func(UiView_newRequestSession_Params) error, opts ...capnp.CallOption) UiView_newRequestSession_Results_Promise {
	if c.Client == nil {
		return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newRequestSession_Params{Struct: s}) }
	}
	return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c UiView) NewOfferSession(ctx context.Context, params func(UiView_newOfferSession_Params) error, opts ...capnp.CallOption) UiView_newOfferSession_Results_Promise {
	if c.Client == nil {
		return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 5}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newOfferSession_Params{Struct: s}) }
	}
	return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type UiView_Server interface {
	GetViewInfo(UiView_getViewInfo) error

	NewSession(UiView_newSession) error

	NewRequestSession(UiView_newRequestSession) error

	NewOfferSession(UiView_newOfferSession) error
}

func UiView_ServerToClient(s UiView_Server) UiView {
	c, _ := s.(server.Closer)
	return UiView{Client: server.New(UiView_Methods(nil, s), c)}
}

func UiView_Methods(methods []server.Method, s UiView_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_getViewInfo{c, opts, UiView_getViewInfo_Params{Struct: p}, UiView_ViewInfo{Struct: r}}
			return s.GetViewInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 5},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newSession{c, opts, UiView_newSession_Params{Struct: p}, UiView_newSession_Results{Struct: r}}
			return s.NewSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newRequestSession{c, opts, UiView_newRequestSession_Params{Struct: p}, UiView_newRequestSession_Results{Struct: r}}
			return s.NewRequestSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newOfferSession{c, opts, UiView_newOfferSession_Params{Struct: p}, UiView_newOfferSession_Results{Struct: r}}
			return s.NewOfferSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// UiView_getViewInfo holds the arguments for a server call to UiView.getViewInfo.
type UiView_getViewInfo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_getViewInfo_Params
	Results UiView_ViewInfo
}

// UiView_newSession holds the arguments for a server call to UiView.newSession.
type UiView_newSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newSession_Params
	Results UiView_newSession_Results
}

// UiView_newRequestSession holds the arguments for a server call to UiView.newRequestSession.
type UiView_newRequestSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newRequestSession_Params
	Results UiView_newRequestSession_Results
}

// UiView_newOfferSession holds the arguments for a server call to UiView.newOfferSession.
type UiView_newOfferSession struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  UiView_newOfferSession_Params
	Results UiView_newOfferSession_Results
}

type UiView_ViewInfo struct{ capnp.Struct }

func NewUiView_ViewInfo(s *capnp.Segment) (UiView_ViewInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	if err != nil {
		return UiView_ViewInfo{}, err
	}
	return UiView_ViewInfo{st}, nil
}

func NewRootUiView_ViewInfo(s *capnp.Segment) (UiView_ViewInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	if err != nil {
		return UiView_ViewInfo{}, err
	}
	return UiView_ViewInfo{st}, nil
}

func ReadRootUiView_ViewInfo(msg *capnp.Message) (UiView_ViewInfo, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_ViewInfo{}, err
	}
	return UiView_ViewInfo{root.Struct()}, nil
}

func (s UiView_ViewInfo) Permissions() (PermissionDef_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return PermissionDef_List{}, err
	}

	return PermissionDef_List{List: p.List()}, nil

}

func (s UiView_ViewInfo) HasPermissions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetPermissions(v PermissionDef_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

func (s UiView_ViewInfo) Roles() (RoleDef_List, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return RoleDef_List{}, err
	}

	return RoleDef_List{List: p.List()}, nil

}

func (s UiView_ViewInfo) HasRoles() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetRoles(v RoleDef_List) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s UiView_ViewInfo) DeniedPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s UiView_ViewInfo) HasDeniedPermissions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetDeniedPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(2, v.List.ToPtr())
}

func (s UiView_ViewInfo) MatchRequests() (PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return PowerboxDescriptor_List{}, err
	}

	return PowerboxDescriptor_List{List: p.List()}, nil

}

func (s UiView_ViewInfo) HasMatchRequests() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetMatchRequests(v PowerboxDescriptor_List) error {

	return s.Struct.SetPtr(3, v.List.ToPtr())
}

func (s UiView_ViewInfo) MatchOffers() (PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return PowerboxDescriptor_List{}, err
	}

	return PowerboxDescriptor_List{List: p.List()}, nil

}

func (s UiView_ViewInfo) HasMatchOffers() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UiView_ViewInfo) SetMatchOffers(v PowerboxDescriptor_List) error {

	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// UiView_ViewInfo_List is a list of UiView_ViewInfo.
type UiView_ViewInfo_List struct{ capnp.List }

// NewUiView_ViewInfo creates a new list of UiView_ViewInfo.
func NewUiView_ViewInfo_List(s *capnp.Segment, sz int32) (UiView_ViewInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5}, sz)
	if err != nil {
		return UiView_ViewInfo_List{}, err
	}
	return UiView_ViewInfo_List{l}, nil
}

func (s UiView_ViewInfo_List) At(i int) UiView_ViewInfo { return UiView_ViewInfo{s.List.Struct(i)} }
func (s UiView_ViewInfo_List) Set(i int, v UiView_ViewInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_ViewInfo_Promise is a wrapper for a UiView_ViewInfo promised by a client call.
type UiView_ViewInfo_Promise struct{ *capnp.Pipeline }

func (p UiView_ViewInfo_Promise) Struct() (UiView_ViewInfo, error) {
	s, err := p.Pipeline.Struct()
	return UiView_ViewInfo{s}, err
}

type UiView_getViewInfo_Params struct{ capnp.Struct }

func NewUiView_getViewInfo_Params(s *capnp.Segment) (UiView_getViewInfo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return UiView_getViewInfo_Params{}, err
	}
	return UiView_getViewInfo_Params{st}, nil
}

func NewRootUiView_getViewInfo_Params(s *capnp.Segment) (UiView_getViewInfo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return UiView_getViewInfo_Params{}, err
	}
	return UiView_getViewInfo_Params{st}, nil
}

func ReadRootUiView_getViewInfo_Params(msg *capnp.Message) (UiView_getViewInfo_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_getViewInfo_Params{}, err
	}
	return UiView_getViewInfo_Params{root.Struct()}, nil
}

// UiView_getViewInfo_Params_List is a list of UiView_getViewInfo_Params.
type UiView_getViewInfo_Params_List struct{ capnp.List }

// NewUiView_getViewInfo_Params creates a new list of UiView_getViewInfo_Params.
func NewUiView_getViewInfo_Params_List(s *capnp.Segment, sz int32) (UiView_getViewInfo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return UiView_getViewInfo_Params_List{}, err
	}
	return UiView_getViewInfo_Params_List{l}, nil
}

func (s UiView_getViewInfo_Params_List) At(i int) UiView_getViewInfo_Params {
	return UiView_getViewInfo_Params{s.List.Struct(i)}
}
func (s UiView_getViewInfo_Params_List) Set(i int, v UiView_getViewInfo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_getViewInfo_Params_Promise is a wrapper for a UiView_getViewInfo_Params promised by a client call.
type UiView_getViewInfo_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_getViewInfo_Params_Promise) Struct() (UiView_getViewInfo_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_getViewInfo_Params{s}, err
}

type UiView_newSession_Params struct{ capnp.Struct }

func NewUiView_newSession_Params(s *capnp.Segment) (UiView_newSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return UiView_newSession_Params{}, err
	}
	return UiView_newSession_Params{st}, nil
}

func NewRootUiView_newSession_Params(s *capnp.Segment) (UiView_newSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return UiView_newSession_Params{}, err
	}
	return UiView_newSession_Params{st}, nil
}

func ReadRootUiView_newSession_Params(msg *capnp.Message) (UiView_newSession_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newSession_Params{}, err
	}
	return UiView_newSession_Params{root.Struct()}, nil
}

func (s UiView_newSession_Params) UserInfo() (UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return UserInfo{}, err
	}

	return UserInfo{Struct: p.Struct()}, nil

}

func (s UiView_newSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SetUserInfo(v UserInfo) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated UserInfo struct, preferring placement in s's segment.
func (s UiView_newSession_Params) NewUserInfo() (UserInfo, error) {

	ss, err := NewUserInfo(s.Struct.Segment())
	if err != nil {
		return UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newSession_Params) Context() SessionContext {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return SessionContext{}
	}
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SetContext(v SessionContext) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newSession_Params) SetSessionType(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s UiView_newSession_Params) SessionParams() (capnp.Pointer, error) {

	return s.Struct.Pointer(2)

}

func (s UiView_newSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Params) SessionParamsPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(2)

}

func (s UiView_newSession_Params) SetSessionParams(v capnp.Pointer) error {

	return s.Struct.SetPointer(2, v)
}

func (s UiView_newSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(2, v)
}

// UiView_newSession_Params_List is a list of UiView_newSession_Params.
type UiView_newSession_Params_List struct{ capnp.List }

// NewUiView_newSession_Params creates a new list of UiView_newSession_Params.
func NewUiView_newSession_Params_List(s *capnp.Segment, sz int32) (UiView_newSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return UiView_newSession_Params_List{}, err
	}
	return UiView_newSession_Params_List{l}, nil
}

func (s UiView_newSession_Params_List) At(i int) UiView_newSession_Params {
	return UiView_newSession_Params{s.List.Struct(i)}
}
func (s UiView_newSession_Params_List) Set(i int, v UiView_newSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newSession_Params_Promise is a wrapper for a UiView_newSession_Params promised by a client call.
type UiView_newSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newSession_Params_Promise) Struct() (UiView_newSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newSession_Params{s}, err
}

func (p UiView_newSession_Params_Promise) UserInfo() UserInfo_Promise {
	return UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

type UiView_newSession_Results struct{ capnp.Struct }

func NewUiView_newSession_Results(s *capnp.Segment) (UiView_newSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newSession_Results{}, err
	}
	return UiView_newSession_Results{st}, nil
}

func NewRootUiView_newSession_Results(s *capnp.Segment) (UiView_newSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newSession_Results{}, err
	}
	return UiView_newSession_Results{st}, nil
}

func ReadRootUiView_newSession_Results(msg *capnp.Message) (UiView_newSession_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newSession_Results{}, err
	}
	return UiView_newSession_Results{root.Struct()}, nil
}

func (s UiView_newSession_Results) Session() UiSession {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiSession{}
	}
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newSession_Results) SetSession(v UiSession) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newSession_Results_List is a list of UiView_newSession_Results.
type UiView_newSession_Results_List struct{ capnp.List }

// NewUiView_newSession_Results creates a new list of UiView_newSession_Results.
func NewUiView_newSession_Results_List(s *capnp.Segment, sz int32) (UiView_newSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return UiView_newSession_Results_List{}, err
	}
	return UiView_newSession_Results_List{l}, nil
}

func (s UiView_newSession_Results_List) At(i int) UiView_newSession_Results {
	return UiView_newSession_Results{s.List.Struct(i)}
}
func (s UiView_newSession_Results_List) Set(i int, v UiView_newSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newSession_Results_Promise is a wrapper for a UiView_newSession_Results promised by a client call.
type UiView_newSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newSession_Results_Promise) Struct() (UiView_newSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newSession_Results{s}, err
}

func (p UiView_newSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiView_newRequestSession_Params struct{ capnp.Struct }

func NewUiView_newRequestSession_Params(s *capnp.Segment) (UiView_newRequestSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return UiView_newRequestSession_Params{}, err
	}
	return UiView_newRequestSession_Params{st}, nil
}

func NewRootUiView_newRequestSession_Params(s *capnp.Segment) (UiView_newRequestSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return UiView_newRequestSession_Params{}, err
	}
	return UiView_newRequestSession_Params{st}, nil
}

func ReadRootUiView_newRequestSession_Params(msg *capnp.Message) (UiView_newRequestSession_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newRequestSession_Params{}, err
	}
	return UiView_newRequestSession_Params{root.Struct()}, nil
}

func (s UiView_newRequestSession_Params) UserInfo() (UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return UserInfo{}, err
	}

	return UserInfo{Struct: p.Struct()}, nil

}

func (s UiView_newRequestSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetUserInfo(v UserInfo) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated UserInfo struct, preferring placement in s's segment.
func (s UiView_newRequestSession_Params) NewUserInfo() (UserInfo, error) {

	ss, err := NewUserInfo(s.Struct.Segment())
	if err != nil {
		return UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newRequestSession_Params) Context() SessionContext {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return SessionContext{}
	}
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newRequestSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetContext(v SessionContext) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newRequestSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newRequestSession_Params) SetSessionType(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s UiView_newRequestSession_Params) SessionParams() (capnp.Pointer, error) {

	return s.Struct.Pointer(2)

}

func (s UiView_newRequestSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SessionParamsPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(2)

}

func (s UiView_newRequestSession_Params) SetSessionParams(v capnp.Pointer) error {

	return s.Struct.SetPointer(2, v)
}

func (s UiView_newRequestSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(2, v)
}

func (s UiView_newRequestSession_Params) RequestInfo() (PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return PowerboxDescriptor_List{}, err
	}

	return PowerboxDescriptor_List{List: p.List()}, nil

}

func (s UiView_newRequestSession_Params) HasRequestInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Params) SetRequestInfo(v PowerboxDescriptor_List) error {

	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// UiView_newRequestSession_Params_List is a list of UiView_newRequestSession_Params.
type UiView_newRequestSession_Params_List struct{ capnp.List }

// NewUiView_newRequestSession_Params creates a new list of UiView_newRequestSession_Params.
func NewUiView_newRequestSession_Params_List(s *capnp.Segment, sz int32) (UiView_newRequestSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	if err != nil {
		return UiView_newRequestSession_Params_List{}, err
	}
	return UiView_newRequestSession_Params_List{l}, nil
}

func (s UiView_newRequestSession_Params_List) At(i int) UiView_newRequestSession_Params {
	return UiView_newRequestSession_Params{s.List.Struct(i)}
}
func (s UiView_newRequestSession_Params_List) Set(i int, v UiView_newRequestSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newRequestSession_Params_Promise is a wrapper for a UiView_newRequestSession_Params promised by a client call.
type UiView_newRequestSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newRequestSession_Params_Promise) Struct() (UiView_newRequestSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newRequestSession_Params{s}, err
}

func (p UiView_newRequestSession_Params_Promise) UserInfo() UserInfo_Promise {
	return UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newRequestSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newRequestSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

type UiView_newRequestSession_Results struct{ capnp.Struct }

func NewUiView_newRequestSession_Results(s *capnp.Segment) (UiView_newRequestSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newRequestSession_Results{}, err
	}
	return UiView_newRequestSession_Results{st}, nil
}

func NewRootUiView_newRequestSession_Results(s *capnp.Segment) (UiView_newRequestSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newRequestSession_Results{}, err
	}
	return UiView_newRequestSession_Results{st}, nil
}

func ReadRootUiView_newRequestSession_Results(msg *capnp.Message) (UiView_newRequestSession_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newRequestSession_Results{}, err
	}
	return UiView_newRequestSession_Results{root.Struct()}, nil
}

func (s UiView_newRequestSession_Results) Session() UiSession {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiSession{}
	}
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newRequestSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newRequestSession_Results) SetSession(v UiSession) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newRequestSession_Results_List is a list of UiView_newRequestSession_Results.
type UiView_newRequestSession_Results_List struct{ capnp.List }

// NewUiView_newRequestSession_Results creates a new list of UiView_newRequestSession_Results.
func NewUiView_newRequestSession_Results_List(s *capnp.Segment, sz int32) (UiView_newRequestSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return UiView_newRequestSession_Results_List{}, err
	}
	return UiView_newRequestSession_Results_List{l}, nil
}

func (s UiView_newRequestSession_Results_List) At(i int) UiView_newRequestSession_Results {
	return UiView_newRequestSession_Results{s.List.Struct(i)}
}
func (s UiView_newRequestSession_Results_List) Set(i int, v UiView_newRequestSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newRequestSession_Results_Promise is a wrapper for a UiView_newRequestSession_Results promised by a client call.
type UiView_newRequestSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newRequestSession_Results_Promise) Struct() (UiView_newRequestSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newRequestSession_Results{s}, err
}

func (p UiView_newRequestSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiView_newOfferSession_Params struct{ capnp.Struct }

func NewUiView_newOfferSession_Params(s *capnp.Segment) (UiView_newOfferSession_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	if err != nil {
		return UiView_newOfferSession_Params{}, err
	}
	return UiView_newOfferSession_Params{st}, nil
}

func NewRootUiView_newOfferSession_Params(s *capnp.Segment) (UiView_newOfferSession_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	if err != nil {
		return UiView_newOfferSession_Params{}, err
	}
	return UiView_newOfferSession_Params{st}, nil
}

func ReadRootUiView_newOfferSession_Params(msg *capnp.Message) (UiView_newOfferSession_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newOfferSession_Params{}, err
	}
	return UiView_newOfferSession_Params{root.Struct()}, nil
}

func (s UiView_newOfferSession_Params) UserInfo() (UserInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return UserInfo{}, err
	}

	return UserInfo{Struct: p.Struct()}, nil

}

func (s UiView_newOfferSession_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetUserInfo(v UserInfo) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated UserInfo struct, preferring placement in s's segment.
func (s UiView_newOfferSession_Params) NewUserInfo() (UserInfo, error) {

	ss, err := NewUserInfo(s.Struct.Segment())
	if err != nil {
		return UserInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s UiView_newOfferSession_Params) Context() SessionContext {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return SessionContext{}
	}
	return SessionContext{Client: p.Interface().Client()}
}

func (s UiView_newOfferSession_Params) HasContext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetContext(v SessionContext) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

func (s UiView_newOfferSession_Params) SessionType() uint64 {
	return s.Struct.Uint64(0)
}

func (s UiView_newOfferSession_Params) SetSessionType(v uint64) {

	s.Struct.SetUint64(0, v)
}

func (s UiView_newOfferSession_Params) SessionParams() (capnp.Pointer, error) {

	return s.Struct.Pointer(2)

}

func (s UiView_newOfferSession_Params) HasSessionParams() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SessionParamsPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(2)

}

func (s UiView_newOfferSession_Params) SetSessionParams(v capnp.Pointer) error {

	return s.Struct.SetPointer(2, v)
}

func (s UiView_newOfferSession_Params) SetSessionParamsPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(2, v)
}

func (s UiView_newOfferSession_Params) Offer() (capnp.Pointer, error) {

	return s.Struct.Pointer(3)

}

func (s UiView_newOfferSession_Params) HasOffer() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) OfferPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(3)

}

func (s UiView_newOfferSession_Params) SetOffer(v capnp.Pointer) error {

	return s.Struct.SetPointer(3, v)
}

func (s UiView_newOfferSession_Params) SetOfferPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(3, v)
}

func (s UiView_newOfferSession_Params) Descriptor() (PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return PowerboxDescriptor{}, err
	}

	return PowerboxDescriptor{Struct: p.Struct()}, nil

}

func (s UiView_newOfferSession_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Params) SetDescriptor(v PowerboxDescriptor) error {

	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated PowerboxDescriptor struct, preferring placement in s's segment.
func (s UiView_newOfferSession_Params) NewDescriptor() (PowerboxDescriptor, error) {

	ss, err := NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

// UiView_newOfferSession_Params_List is a list of UiView_newOfferSession_Params.
type UiView_newOfferSession_Params_List struct{ capnp.List }

// NewUiView_newOfferSession_Params creates a new list of UiView_newOfferSession_Params.
func NewUiView_newOfferSession_Params_List(s *capnp.Segment, sz int32) (UiView_newOfferSession_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	if err != nil {
		return UiView_newOfferSession_Params_List{}, err
	}
	return UiView_newOfferSession_Params_List{l}, nil
}

func (s UiView_newOfferSession_Params_List) At(i int) UiView_newOfferSession_Params {
	return UiView_newOfferSession_Params{s.List.Struct(i)}
}
func (s UiView_newOfferSession_Params_List) Set(i int, v UiView_newOfferSession_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newOfferSession_Params_Promise is a wrapper for a UiView_newOfferSession_Params promised by a client call.
type UiView_newOfferSession_Params_Promise struct{ *capnp.Pipeline }

func (p UiView_newOfferSession_Params_Promise) Struct() (UiView_newOfferSession_Params, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newOfferSession_Params{s}, err
}

func (p UiView_newOfferSession_Params_Promise) UserInfo() UserInfo_Promise {
	return UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p UiView_newOfferSession_Params_Promise) Context() SessionContext {
	return SessionContext{Client: p.Pipeline.GetPipeline(1).Client()}
}

func (p UiView_newOfferSession_Params_Promise) SessionParams() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

func (p UiView_newOfferSession_Params_Promise) Offer() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(3)
}

func (p UiView_newOfferSession_Params_Promise) Descriptor() PowerboxDescriptor_Promise {
	return PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

type UiView_newOfferSession_Results struct{ capnp.Struct }

func NewUiView_newOfferSession_Results(s *capnp.Segment) (UiView_newOfferSession_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newOfferSession_Results{}, err
	}
	return UiView_newOfferSession_Results{st}, nil
}

func NewRootUiView_newOfferSession_Results(s *capnp.Segment) (UiView_newOfferSession_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return UiView_newOfferSession_Results{}, err
	}
	return UiView_newOfferSession_Results{st}, nil
}

func ReadRootUiView_newOfferSession_Results(msg *capnp.Message) (UiView_newOfferSession_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UiView_newOfferSession_Results{}, err
	}
	return UiView_newOfferSession_Results{root.Struct()}, nil
}

func (s UiView_newOfferSession_Results) Session() UiSession {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiSession{}
	}
	return UiSession{Client: p.Interface().Client()}
}

func (s UiView_newOfferSession_Results) HasSession() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UiView_newOfferSession_Results) SetSession(v UiSession) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// UiView_newOfferSession_Results_List is a list of UiView_newOfferSession_Results.
type UiView_newOfferSession_Results_List struct{ capnp.List }

// NewUiView_newOfferSession_Results creates a new list of UiView_newOfferSession_Results.
func NewUiView_newOfferSession_Results_List(s *capnp.Segment, sz int32) (UiView_newOfferSession_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return UiView_newOfferSession_Results_List{}, err
	}
	return UiView_newOfferSession_Results_List{l}, nil
}

func (s UiView_newOfferSession_Results_List) At(i int) UiView_newOfferSession_Results {
	return UiView_newOfferSession_Results{s.List.Struct(i)}
}
func (s UiView_newOfferSession_Results_List) Set(i int, v UiView_newOfferSession_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// UiView_newOfferSession_Results_Promise is a wrapper for a UiView_newOfferSession_Results promised by a client call.
type UiView_newOfferSession_Results_Promise struct{ *capnp.Pipeline }

func (p UiView_newOfferSession_Results_Promise) Struct() (UiView_newOfferSession_Results, error) {
	s, err := p.Pipeline.Struct()
	return UiView_newOfferSession_Results{s}, err
}

func (p UiView_newOfferSession_Results_Promise) Session() UiSession {
	return UiSession{Client: p.Pipeline.GetPipeline(0).Client()}
}

type UiSession struct{ Client capnp.Client }

type UiSession_Server interface {
}

func UiSession_ServerToClient(s UiSession_Server) UiSession {
	c, _ := s.(server.Closer)
	return UiSession{Client: server.New(UiSession_Methods(nil, s), c)}
}

func UiSession_Methods(methods []server.Method, s UiSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type UserInfo struct{ capnp.Struct }

func NewUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return UserInfo{}, err
	}
	return UserInfo{st}, nil
}

func NewRootUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return UserInfo{}, err
	}
	return UserInfo{st}, nil
}

func ReadRootUserInfo(msg *capnp.Message) (UserInfo, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return UserInfo{}, err
	}
	return UserInfo{root.Struct()}, nil
}

func (s UserInfo) DisplayName() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

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

func (s UserInfo) DeprecatedPermissionsBlob() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s UserInfo) HasPermissions() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(3, v.List.ToPtr())
}

func (s UserInfo) UserId() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

}

func (s UserInfo) HasUserId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s UserInfo) SetUserId(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, d.List.ToPtr())
}

// UserInfo_List is a list of UserInfo.
type UserInfo_List struct{ capnp.List }

// NewUserInfo creates a new list of UserInfo.
func NewUserInfo_List(s *capnp.Segment, sz int32) (UserInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	if err != nil {
		return UserInfo_List{}, err
	}
	return UserInfo_List{l}, nil
}

func (s UserInfo_List) At(i int) UserInfo           { return UserInfo{s.List.Struct(i)} }
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

type SessionContext struct{ Client capnp.Client }

func (c SessionContext) GetSharedPermissions(ctx context.Context, params func(SessionContext_getSharedPermissions_Params) error, opts ...capnp.CallOption) SessionContext_getSharedPermissions_Results_Promise {
	if c.Client == nil {
		return SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_getSharedPermissions_Params{Struct: s}) }
	}
	return SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) TieToUser(ctx context.Context, params func(SessionContext_tieToUser_Params) error, opts ...capnp.CallOption) SessionContext_tieToUser_Results_Promise {
	if c.Client == nil {
		return SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_tieToUser_Params{Struct: s}) }
	}
	return SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) Offer(ctx context.Context, params func(SessionContext_offer_Params) error, opts ...capnp.CallOption) SessionContext_offer_Results_Promise {
	if c.Client == nil {
		return SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_offer_Params{Struct: s}) }
	}
	return SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) Request(ctx context.Context, params func(SessionContext_request_Params) error, opts ...capnp.CallOption) SessionContext_request_Results_Promise {
	if c.Client == nil {
		return SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_request_Params{Struct: s}) }
	}
	return SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) Provide(ctx context.Context, params func(SessionContext_provide_Params) error, opts ...capnp.CallOption) SessionContext_provide_Results_Promise {
	if c.Client == nil {
		return SessionContext_provide_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "provide",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_provide_Params{Struct: s}) }
	}
	return SessionContext_provide_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) Close(ctx context.Context, params func(SessionContext_close_Params) error, opts ...capnp.CallOption) SessionContext_close_Results_Promise {
	if c.Client == nil {
		return SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_close_Params{Struct: s}) }
	}
	return SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c SessionContext) OpenView(ctx context.Context, params func(SessionContext_openView_Params) error, opts ...capnp.CallOption) SessionContext_openView_Results_Promise {
	if c.Client == nil {
		return SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SessionContext_openView_Params{Struct: s}) }
	}
	return SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SessionContext_Server interface {
	GetSharedPermissions(SessionContext_getSharedPermissions) error

	TieToUser(SessionContext_tieToUser) error

	Offer(SessionContext_offer) error

	Request(SessionContext_request) error

	Provide(SessionContext_provide) error

	Close(SessionContext_close) error

	OpenView(SessionContext_openView) error
}

func SessionContext_ServerToClient(s SessionContext_Server) SessionContext {
	c, _ := s.(server.Closer)
	return SessionContext{Client: server.New(SessionContext_Methods(nil, s), c)}
}

func SessionContext_Methods(methods []server.Method, s SessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 7)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_getSharedPermissions{c, opts, SessionContext_getSharedPermissions_Params{Struct: p}, SessionContext_getSharedPermissions_Results{Struct: r}}
			return s.GetSharedPermissions(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_tieToUser{c, opts, SessionContext_tieToUser_Params{Struct: p}, SessionContext_tieToUser_Results{Struct: r}}
			return s.TieToUser(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_offer{c, opts, SessionContext_offer_Params{Struct: p}, SessionContext_offer_Results{Struct: r}}
			return s.Offer(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_request{c, opts, SessionContext_request_Params{Struct: p}, SessionContext_request_Results{Struct: r}}
			return s.Request(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "provide",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_provide{c, opts, SessionContext_provide_Params{Struct: p}, SessionContext_provide_Results{Struct: r}}
			return s.Provide(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_close{c, opts, SessionContext_close_Params{Struct: p}, SessionContext_close_Results{Struct: r}}
			return s.Close(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SessionContext_openView{c, opts, SessionContext_openView_Params{Struct: p}, SessionContext_openView_Results{Struct: r}}
			return s.OpenView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// SessionContext_getSharedPermissions holds the arguments for a server call to SessionContext.getSharedPermissions.
type SessionContext_getSharedPermissions struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_getSharedPermissions_Params
	Results SessionContext_getSharedPermissions_Results
}

// SessionContext_tieToUser holds the arguments for a server call to SessionContext.tieToUser.
type SessionContext_tieToUser struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_tieToUser_Params
	Results SessionContext_tieToUser_Results
}

// SessionContext_offer holds the arguments for a server call to SessionContext.offer.
type SessionContext_offer struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_offer_Params
	Results SessionContext_offer_Results
}

// SessionContext_request holds the arguments for a server call to SessionContext.request.
type SessionContext_request struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_request_Params
	Results SessionContext_request_Results
}

// SessionContext_provide holds the arguments for a server call to SessionContext.provide.
type SessionContext_provide struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_provide_Params
	Results SessionContext_provide_Results
}

// SessionContext_close holds the arguments for a server call to SessionContext.close.
type SessionContext_close struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_close_Params
	Results SessionContext_close_Results
}

// SessionContext_openView holds the arguments for a server call to SessionContext.openView.
type SessionContext_openView struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SessionContext_openView_Params
	Results SessionContext_openView_Results
}

type SessionContext_getSharedPermissions_Params struct{ capnp.Struct }

func NewSessionContext_getSharedPermissions_Params(s *capnp.Segment) (SessionContext_getSharedPermissions_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_getSharedPermissions_Params{}, err
	}
	return SessionContext_getSharedPermissions_Params{st}, nil
}

func NewRootSessionContext_getSharedPermissions_Params(s *capnp.Segment) (SessionContext_getSharedPermissions_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_getSharedPermissions_Params{}, err
	}
	return SessionContext_getSharedPermissions_Params{st}, nil
}

func ReadRootSessionContext_getSharedPermissions_Params(msg *capnp.Message) (SessionContext_getSharedPermissions_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_getSharedPermissions_Params{}, err
	}
	return SessionContext_getSharedPermissions_Params{root.Struct()}, nil
}

// SessionContext_getSharedPermissions_Params_List is a list of SessionContext_getSharedPermissions_Params.
type SessionContext_getSharedPermissions_Params_List struct{ capnp.List }

// NewSessionContext_getSharedPermissions_Params creates a new list of SessionContext_getSharedPermissions_Params.
func NewSessionContext_getSharedPermissions_Params_List(s *capnp.Segment, sz int32) (SessionContext_getSharedPermissions_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_getSharedPermissions_Params_List{}, err
	}
	return SessionContext_getSharedPermissions_Params_List{l}, nil
}

func (s SessionContext_getSharedPermissions_Params_List) At(i int) SessionContext_getSharedPermissions_Params {
	return SessionContext_getSharedPermissions_Params{s.List.Struct(i)}
}
func (s SessionContext_getSharedPermissions_Params_List) Set(i int, v SessionContext_getSharedPermissions_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_getSharedPermissions_Params_Promise is a wrapper for a SessionContext_getSharedPermissions_Params promised by a client call.
type SessionContext_getSharedPermissions_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_getSharedPermissions_Params_Promise) Struct() (SessionContext_getSharedPermissions_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_getSharedPermissions_Params{s}, err
}

type SessionContext_getSharedPermissions_Results struct{ capnp.Struct }

func NewSessionContext_getSharedPermissions_Results(s *capnp.Segment) (SessionContext_getSharedPermissions_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_getSharedPermissions_Results{}, err
	}
	return SessionContext_getSharedPermissions_Results{st}, nil
}

func NewRootSessionContext_getSharedPermissions_Results(s *capnp.Segment) (SessionContext_getSharedPermissions_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_getSharedPermissions_Results{}, err
	}
	return SessionContext_getSharedPermissions_Results{st}, nil
}

func ReadRootSessionContext_getSharedPermissions_Results(msg *capnp.Message) (SessionContext_getSharedPermissions_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_getSharedPermissions_Results{}, err
	}
	return SessionContext_getSharedPermissions_Results{root.Struct()}, nil
}

func (s SessionContext_getSharedPermissions_Results) Var() util.Assignable_Getter {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Assignable_Getter{}
	}
	return util.Assignable_Getter{Client: p.Interface().Client()}
}

func (s SessionContext_getSharedPermissions_Results) HasVar() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_getSharedPermissions_Results) SetVar(v util.Assignable_Getter) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SessionContext_getSharedPermissions_Results_List is a list of SessionContext_getSharedPermissions_Results.
type SessionContext_getSharedPermissions_Results_List struct{ capnp.List }

// NewSessionContext_getSharedPermissions_Results creates a new list of SessionContext_getSharedPermissions_Results.
func NewSessionContext_getSharedPermissions_Results_List(s *capnp.Segment, sz int32) (SessionContext_getSharedPermissions_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SessionContext_getSharedPermissions_Results_List{}, err
	}
	return SessionContext_getSharedPermissions_Results_List{l}, nil
}

func (s SessionContext_getSharedPermissions_Results_List) At(i int) SessionContext_getSharedPermissions_Results {
	return SessionContext_getSharedPermissions_Results{s.List.Struct(i)}
}
func (s SessionContext_getSharedPermissions_Results_List) Set(i int, v SessionContext_getSharedPermissions_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_getSharedPermissions_Results_Promise is a wrapper for a SessionContext_getSharedPermissions_Results promised by a client call.
type SessionContext_getSharedPermissions_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_getSharedPermissions_Results_Promise) Struct() (SessionContext_getSharedPermissions_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_getSharedPermissions_Results{s}, err
}

func (p SessionContext_getSharedPermissions_Results_Promise) Var() util.Assignable_Getter {
	return util.Assignable_Getter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SessionContext_tieToUser_Params struct{ capnp.Struct }

func NewSessionContext_tieToUser_Params(s *capnp.Segment) (SessionContext_tieToUser_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SessionContext_tieToUser_Params{}, err
	}
	return SessionContext_tieToUser_Params{st}, nil
}

func NewRootSessionContext_tieToUser_Params(s *capnp.Segment) (SessionContext_tieToUser_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return SessionContext_tieToUser_Params{}, err
	}
	return SessionContext_tieToUser_Params{st}, nil
}

func ReadRootSessionContext_tieToUser_Params(msg *capnp.Message) (SessionContext_tieToUser_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_tieToUser_Params{}, err
	}
	return SessionContext_tieToUser_Params{root.Struct()}, nil
}

func (s SessionContext_tieToUser_Params) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SessionContext_tieToUser_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SessionContext_tieToUser_Params) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_tieToUser_Params) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_tieToUser_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s SessionContext_tieToUser_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) SetRequiredPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s SessionContext_tieToUser_Params) DisplayInfo() (PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}

	return PowerboxDisplayInfo{Struct: p.Struct()}, nil

}

func (s SessionContext_tieToUser_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Params) SetDisplayInfo(v PowerboxDisplayInfo) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_tieToUser_Params) NewDisplayInfo() (PowerboxDisplayInfo, error) {

	ss, err := NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_tieToUser_Params_List is a list of SessionContext_tieToUser_Params.
type SessionContext_tieToUser_Params_List struct{ capnp.List }

// NewSessionContext_tieToUser_Params creates a new list of SessionContext_tieToUser_Params.
func NewSessionContext_tieToUser_Params_List(s *capnp.Segment, sz int32) (SessionContext_tieToUser_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return SessionContext_tieToUser_Params_List{}, err
	}
	return SessionContext_tieToUser_Params_List{l}, nil
}

func (s SessionContext_tieToUser_Params_List) At(i int) SessionContext_tieToUser_Params {
	return SessionContext_tieToUser_Params{s.List.Struct(i)}
}
func (s SessionContext_tieToUser_Params_List) Set(i int, v SessionContext_tieToUser_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_tieToUser_Params_Promise is a wrapper for a SessionContext_tieToUser_Params promised by a client call.
type SessionContext_tieToUser_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_tieToUser_Params_Promise) Struct() (SessionContext_tieToUser_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_tieToUser_Params{s}, err
}

func (p SessionContext_tieToUser_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_tieToUser_Params_Promise) DisplayInfo() PowerboxDisplayInfo_Promise {
	return PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SessionContext_tieToUser_Results struct{ capnp.Struct }

func NewSessionContext_tieToUser_Results(s *capnp.Segment) (SessionContext_tieToUser_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_tieToUser_Results{}, err
	}
	return SessionContext_tieToUser_Results{st}, nil
}

func NewRootSessionContext_tieToUser_Results(s *capnp.Segment) (SessionContext_tieToUser_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_tieToUser_Results{}, err
	}
	return SessionContext_tieToUser_Results{st}, nil
}

func ReadRootSessionContext_tieToUser_Results(msg *capnp.Message) (SessionContext_tieToUser_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_tieToUser_Results{}, err
	}
	return SessionContext_tieToUser_Results{root.Struct()}, nil
}

func (s SessionContext_tieToUser_Results) TiedCap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SessionContext_tieToUser_Results) HasTiedCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_tieToUser_Results) TiedCapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SessionContext_tieToUser_Results) SetTiedCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_tieToUser_Results) SetTiedCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// SessionContext_tieToUser_Results_List is a list of SessionContext_tieToUser_Results.
type SessionContext_tieToUser_Results_List struct{ capnp.List }

// NewSessionContext_tieToUser_Results creates a new list of SessionContext_tieToUser_Results.
func NewSessionContext_tieToUser_Results_List(s *capnp.Segment, sz int32) (SessionContext_tieToUser_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SessionContext_tieToUser_Results_List{}, err
	}
	return SessionContext_tieToUser_Results_List{l}, nil
}

func (s SessionContext_tieToUser_Results_List) At(i int) SessionContext_tieToUser_Results {
	return SessionContext_tieToUser_Results{s.List.Struct(i)}
}
func (s SessionContext_tieToUser_Results_List) Set(i int, v SessionContext_tieToUser_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_tieToUser_Results_Promise is a wrapper for a SessionContext_tieToUser_Results promised by a client call.
type SessionContext_tieToUser_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_tieToUser_Results_Promise) Struct() (SessionContext_tieToUser_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_tieToUser_Results{s}, err
}

func (p SessionContext_tieToUser_Results_Promise) TiedCap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type SessionContext_offer_Params struct{ capnp.Struct }

func NewSessionContext_offer_Params(s *capnp.Segment) (SessionContext_offer_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return SessionContext_offer_Params{}, err
	}
	return SessionContext_offer_Params{st}, nil
}

func NewRootSessionContext_offer_Params(s *capnp.Segment) (SessionContext_offer_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return SessionContext_offer_Params{}, err
	}
	return SessionContext_offer_Params{st}, nil
}

func ReadRootSessionContext_offer_Params(msg *capnp.Message) (SessionContext_offer_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_offer_Params{}, err
	}
	return SessionContext_offer_Params{root.Struct()}, nil
}

func (s SessionContext_offer_Params) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SessionContext_offer_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SessionContext_offer_Params) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_offer_Params) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_offer_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s SessionContext_offer_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetRequiredPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s SessionContext_offer_Params) Descriptor() (PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return PowerboxDescriptor{}, err
	}

	return PowerboxDescriptor{Struct: p.Struct()}, nil

}

func (s SessionContext_offer_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetDescriptor(v PowerboxDescriptor) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_offer_Params) NewDescriptor() (PowerboxDescriptor, error) {

	ss, err := NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionContext_offer_Params) DisplayInfo() (PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}

	return PowerboxDisplayInfo{Struct: p.Struct()}, nil

}

func (s SessionContext_offer_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s SessionContext_offer_Params) SetDisplayInfo(v PowerboxDisplayInfo) error {

	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_offer_Params) NewDisplayInfo() (PowerboxDisplayInfo, error) {

	ss, err := NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_offer_Params_List is a list of SessionContext_offer_Params.
type SessionContext_offer_Params_List struct{ capnp.List }

// NewSessionContext_offer_Params creates a new list of SessionContext_offer_Params.
func NewSessionContext_offer_Params_List(s *capnp.Segment, sz int32) (SessionContext_offer_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	if err != nil {
		return SessionContext_offer_Params_List{}, err
	}
	return SessionContext_offer_Params_List{l}, nil
}

func (s SessionContext_offer_Params_List) At(i int) SessionContext_offer_Params {
	return SessionContext_offer_Params{s.List.Struct(i)}
}
func (s SessionContext_offer_Params_List) Set(i int, v SessionContext_offer_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_offer_Params_Promise is a wrapper for a SessionContext_offer_Params promised by a client call.
type SessionContext_offer_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_offer_Params_Promise) Struct() (SessionContext_offer_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_offer_Params{s}, err
}

func (p SessionContext_offer_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_offer_Params_Promise) Descriptor() PowerboxDescriptor_Promise {
	return PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p SessionContext_offer_Params_Promise) DisplayInfo() PowerboxDisplayInfo_Promise {
	return PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SessionContext_offer_Results struct{ capnp.Struct }

func NewSessionContext_offer_Results(s *capnp.Segment) (SessionContext_offer_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_offer_Results{}, err
	}
	return SessionContext_offer_Results{st}, nil
}

func NewRootSessionContext_offer_Results(s *capnp.Segment) (SessionContext_offer_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_offer_Results{}, err
	}
	return SessionContext_offer_Results{st}, nil
}

func ReadRootSessionContext_offer_Results(msg *capnp.Message) (SessionContext_offer_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_offer_Results{}, err
	}
	return SessionContext_offer_Results{root.Struct()}, nil
}

// SessionContext_offer_Results_List is a list of SessionContext_offer_Results.
type SessionContext_offer_Results_List struct{ capnp.List }

// NewSessionContext_offer_Results creates a new list of SessionContext_offer_Results.
func NewSessionContext_offer_Results_List(s *capnp.Segment, sz int32) (SessionContext_offer_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_offer_Results_List{}, err
	}
	return SessionContext_offer_Results_List{l}, nil
}

func (s SessionContext_offer_Results_List) At(i int) SessionContext_offer_Results {
	return SessionContext_offer_Results{s.List.Struct(i)}
}
func (s SessionContext_offer_Results_List) Set(i int, v SessionContext_offer_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_offer_Results_Promise is a wrapper for a SessionContext_offer_Results promised by a client call.
type SessionContext_offer_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_offer_Results_Promise) Struct() (SessionContext_offer_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_offer_Results{s}, err
}

type SessionContext_request_Params struct{ capnp.Struct }

func NewSessionContext_request_Params(s *capnp.Segment) (SessionContext_request_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_request_Params{}, err
	}
	return SessionContext_request_Params{st}, nil
}

func NewRootSessionContext_request_Params(s *capnp.Segment) (SessionContext_request_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SessionContext_request_Params{}, err
	}
	return SessionContext_request_Params{st}, nil
}

func ReadRootSessionContext_request_Params(msg *capnp.Message) (SessionContext_request_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_request_Params{}, err
	}
	return SessionContext_request_Params{root.Struct()}, nil
}

func (s SessionContext_request_Params) Query() (PowerboxDescriptor_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return PowerboxDescriptor_List{}, err
	}

	return PowerboxDescriptor_List{List: p.List()}, nil

}

func (s SessionContext_request_Params) HasQuery() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Params) SetQuery(v PowerboxDescriptor_List) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// SessionContext_request_Params_List is a list of SessionContext_request_Params.
type SessionContext_request_Params_List struct{ capnp.List }

// NewSessionContext_request_Params creates a new list of SessionContext_request_Params.
func NewSessionContext_request_Params_List(s *capnp.Segment, sz int32) (SessionContext_request_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SessionContext_request_Params_List{}, err
	}
	return SessionContext_request_Params_List{l}, nil
}

func (s SessionContext_request_Params_List) At(i int) SessionContext_request_Params {
	return SessionContext_request_Params{s.List.Struct(i)}
}
func (s SessionContext_request_Params_List) Set(i int, v SessionContext_request_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_request_Params_Promise is a wrapper for a SessionContext_request_Params promised by a client call.
type SessionContext_request_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_request_Params_Promise) Struct() (SessionContext_request_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_request_Params{s}, err
}

type SessionContext_request_Results struct{ capnp.Struct }

func NewSessionContext_request_Results(s *capnp.Segment) (SessionContext_request_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SessionContext_request_Results{}, err
	}
	return SessionContext_request_Results{st}, nil
}

func NewRootSessionContext_request_Results(s *capnp.Segment) (SessionContext_request_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return SessionContext_request_Results{}, err
	}
	return SessionContext_request_Results{st}, nil
}

func ReadRootSessionContext_request_Results(msg *capnp.Message) (SessionContext_request_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_request_Results{}, err
	}
	return SessionContext_request_Results{root.Struct()}, nil
}

func (s SessionContext_request_Results) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SessionContext_request_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Results) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SessionContext_request_Results) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_request_Results) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_request_Results) Descriptor() (PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return PowerboxDescriptor{}, err
	}

	return PowerboxDescriptor{Struct: p.Struct()}, nil

}

func (s SessionContext_request_Results) HasDescriptor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_request_Results) SetDescriptor(v PowerboxDescriptor) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_request_Results) NewDescriptor() (PowerboxDescriptor, error) {

	ss, err := NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_request_Results_List is a list of SessionContext_request_Results.
type SessionContext_request_Results_List struct{ capnp.List }

// NewSessionContext_request_Results creates a new list of SessionContext_request_Results.
func NewSessionContext_request_Results_List(s *capnp.Segment, sz int32) (SessionContext_request_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return SessionContext_request_Results_List{}, err
	}
	return SessionContext_request_Results_List{l}, nil
}

func (s SessionContext_request_Results_List) At(i int) SessionContext_request_Results {
	return SessionContext_request_Results{s.List.Struct(i)}
}
func (s SessionContext_request_Results_List) Set(i int, v SessionContext_request_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_request_Results_Promise is a wrapper for a SessionContext_request_Results promised by a client call.
type SessionContext_request_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_request_Results_Promise) Struct() (SessionContext_request_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_request_Results{s}, err
}

func (p SessionContext_request_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_request_Results_Promise) Descriptor() PowerboxDescriptor_Promise {
	return PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type SessionContext_provide_Params struct{ capnp.Struct }

func NewSessionContext_provide_Params(s *capnp.Segment) (SessionContext_provide_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return SessionContext_provide_Params{}, err
	}
	return SessionContext_provide_Params{st}, nil
}

func NewRootSessionContext_provide_Params(s *capnp.Segment) (SessionContext_provide_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return SessionContext_provide_Params{}, err
	}
	return SessionContext_provide_Params{st}, nil
}

func ReadRootSessionContext_provide_Params(msg *capnp.Message) (SessionContext_provide_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_provide_Params{}, err
	}
	return SessionContext_provide_Params{root.Struct()}, nil
}

func (s SessionContext_provide_Params) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s SessionContext_provide_Params) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_provide_Params) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s SessionContext_provide_Params) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s SessionContext_provide_Params) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s SessionContext_provide_Params) RequiredPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s SessionContext_provide_Params) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_provide_Params) SetRequiredPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s SessionContext_provide_Params) Descriptor() (PowerboxDescriptor, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return PowerboxDescriptor{}, err
	}

	return PowerboxDescriptor{Struct: p.Struct()}, nil

}

func (s SessionContext_provide_Params) HasDescriptor() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s SessionContext_provide_Params) SetDescriptor(v PowerboxDescriptor) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescriptor sets the descriptor field to a newly
// allocated PowerboxDescriptor struct, preferring placement in s's segment.
func (s SessionContext_provide_Params) NewDescriptor() (PowerboxDescriptor, error) {

	ss, err := NewPowerboxDescriptor(s.Struct.Segment())
	if err != nil {
		return PowerboxDescriptor{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s SessionContext_provide_Params) DisplayInfo() (PowerboxDisplayInfo, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}

	return PowerboxDisplayInfo{Struct: p.Struct()}, nil

}

func (s SessionContext_provide_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s SessionContext_provide_Params) SetDisplayInfo(v PowerboxDisplayInfo) error {

	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated PowerboxDisplayInfo struct, preferring placement in s's segment.
func (s SessionContext_provide_Params) NewDisplayInfo() (PowerboxDisplayInfo, error) {

	ss, err := NewPowerboxDisplayInfo(s.Struct.Segment())
	if err != nil {
		return PowerboxDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// SessionContext_provide_Params_List is a list of SessionContext_provide_Params.
type SessionContext_provide_Params_List struct{ capnp.List }

// NewSessionContext_provide_Params creates a new list of SessionContext_provide_Params.
func NewSessionContext_provide_Params_List(s *capnp.Segment, sz int32) (SessionContext_provide_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	if err != nil {
		return SessionContext_provide_Params_List{}, err
	}
	return SessionContext_provide_Params_List{l}, nil
}

func (s SessionContext_provide_Params_List) At(i int) SessionContext_provide_Params {
	return SessionContext_provide_Params{s.List.Struct(i)}
}
func (s SessionContext_provide_Params_List) Set(i int, v SessionContext_provide_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_provide_Params_Promise is a wrapper for a SessionContext_provide_Params promised by a client call.
type SessionContext_provide_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_provide_Params_Promise) Struct() (SessionContext_provide_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_provide_Params{s}, err
}

func (p SessionContext_provide_Params_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p SessionContext_provide_Params_Promise) Descriptor() PowerboxDescriptor_Promise {
	return PowerboxDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p SessionContext_provide_Params_Promise) DisplayInfo() PowerboxDisplayInfo_Promise {
	return PowerboxDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type SessionContext_provide_Results struct{ capnp.Struct }

func NewSessionContext_provide_Results(s *capnp.Segment) (SessionContext_provide_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_provide_Results{}, err
	}
	return SessionContext_provide_Results{st}, nil
}

func NewRootSessionContext_provide_Results(s *capnp.Segment) (SessionContext_provide_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_provide_Results{}, err
	}
	return SessionContext_provide_Results{st}, nil
}

func ReadRootSessionContext_provide_Results(msg *capnp.Message) (SessionContext_provide_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_provide_Results{}, err
	}
	return SessionContext_provide_Results{root.Struct()}, nil
}

// SessionContext_provide_Results_List is a list of SessionContext_provide_Results.
type SessionContext_provide_Results_List struct{ capnp.List }

// NewSessionContext_provide_Results creates a new list of SessionContext_provide_Results.
func NewSessionContext_provide_Results_List(s *capnp.Segment, sz int32) (SessionContext_provide_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_provide_Results_List{}, err
	}
	return SessionContext_provide_Results_List{l}, nil
}

func (s SessionContext_provide_Results_List) At(i int) SessionContext_provide_Results {
	return SessionContext_provide_Results{s.List.Struct(i)}
}
func (s SessionContext_provide_Results_List) Set(i int, v SessionContext_provide_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_provide_Results_Promise is a wrapper for a SessionContext_provide_Results promised by a client call.
type SessionContext_provide_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_provide_Results_Promise) Struct() (SessionContext_provide_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_provide_Results{s}, err
}

type SessionContext_close_Params struct{ capnp.Struct }

func NewSessionContext_close_Params(s *capnp.Segment) (SessionContext_close_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_close_Params{}, err
	}
	return SessionContext_close_Params{st}, nil
}

func NewRootSessionContext_close_Params(s *capnp.Segment) (SessionContext_close_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_close_Params{}, err
	}
	return SessionContext_close_Params{st}, nil
}

func ReadRootSessionContext_close_Params(msg *capnp.Message) (SessionContext_close_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_close_Params{}, err
	}
	return SessionContext_close_Params{root.Struct()}, nil
}

// SessionContext_close_Params_List is a list of SessionContext_close_Params.
type SessionContext_close_Params_List struct{ capnp.List }

// NewSessionContext_close_Params creates a new list of SessionContext_close_Params.
func NewSessionContext_close_Params_List(s *capnp.Segment, sz int32) (SessionContext_close_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_close_Params_List{}, err
	}
	return SessionContext_close_Params_List{l}, nil
}

func (s SessionContext_close_Params_List) At(i int) SessionContext_close_Params {
	return SessionContext_close_Params{s.List.Struct(i)}
}
func (s SessionContext_close_Params_List) Set(i int, v SessionContext_close_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_close_Params_Promise is a wrapper for a SessionContext_close_Params promised by a client call.
type SessionContext_close_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_close_Params_Promise) Struct() (SessionContext_close_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_close_Params{s}, err
}

type SessionContext_close_Results struct{ capnp.Struct }

func NewSessionContext_close_Results(s *capnp.Segment) (SessionContext_close_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_close_Results{}, err
	}
	return SessionContext_close_Results{st}, nil
}

func NewRootSessionContext_close_Results(s *capnp.Segment) (SessionContext_close_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_close_Results{}, err
	}
	return SessionContext_close_Results{st}, nil
}

func ReadRootSessionContext_close_Results(msg *capnp.Message) (SessionContext_close_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_close_Results{}, err
	}
	return SessionContext_close_Results{root.Struct()}, nil
}

// SessionContext_close_Results_List is a list of SessionContext_close_Results.
type SessionContext_close_Results_List struct{ capnp.List }

// NewSessionContext_close_Results creates a new list of SessionContext_close_Results.
func NewSessionContext_close_Results_List(s *capnp.Segment, sz int32) (SessionContext_close_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_close_Results_List{}, err
	}
	return SessionContext_close_Results_List{l}, nil
}

func (s SessionContext_close_Results_List) At(i int) SessionContext_close_Results {
	return SessionContext_close_Results{s.List.Struct(i)}
}
func (s SessionContext_close_Results_List) Set(i int, v SessionContext_close_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_close_Results_Promise is a wrapper for a SessionContext_close_Results promised by a client call.
type SessionContext_close_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_close_Results_Promise) Struct() (SessionContext_close_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_close_Results{s}, err
}

type SessionContext_openView_Params struct{ capnp.Struct }

func NewSessionContext_openView_Params(s *capnp.Segment) (SessionContext_openView_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return SessionContext_openView_Params{}, err
	}
	return SessionContext_openView_Params{st}, nil
}

func NewRootSessionContext_openView_Params(s *capnp.Segment) (SessionContext_openView_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return SessionContext_openView_Params{}, err
	}
	return SessionContext_openView_Params{st}, nil
}

func ReadRootSessionContext_openView_Params(msg *capnp.Message) (SessionContext_openView_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_openView_Params{}, err
	}
	return SessionContext_openView_Params{root.Struct()}, nil
}

func (s SessionContext_openView_Params) View() UiView {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return UiView{}
	}
	return UiView{Client: p.Interface().Client()}
}

func (s SessionContext_openView_Params) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SessionContext_openView_Params) SetView(v UiView) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s SessionContext_openView_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s SessionContext_openView_Params) HasPath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s SessionContext_openView_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s SessionContext_openView_Params) SetPath(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s SessionContext_openView_Params) NewTab() bool {
	return s.Struct.Bit(0)
}

func (s SessionContext_openView_Params) SetNewTab(v bool) {

	s.Struct.SetBit(0, v)
}

// SessionContext_openView_Params_List is a list of SessionContext_openView_Params.
type SessionContext_openView_Params_List struct{ capnp.List }

// NewSessionContext_openView_Params creates a new list of SessionContext_openView_Params.
func NewSessionContext_openView_Params_List(s *capnp.Segment, sz int32) (SessionContext_openView_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return SessionContext_openView_Params_List{}, err
	}
	return SessionContext_openView_Params_List{l}, nil
}

func (s SessionContext_openView_Params_List) At(i int) SessionContext_openView_Params {
	return SessionContext_openView_Params{s.List.Struct(i)}
}
func (s SessionContext_openView_Params_List) Set(i int, v SessionContext_openView_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_openView_Params_Promise is a wrapper for a SessionContext_openView_Params promised by a client call.
type SessionContext_openView_Params_Promise struct{ *capnp.Pipeline }

func (p SessionContext_openView_Params_Promise) Struct() (SessionContext_openView_Params, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_openView_Params{s}, err
}

func (p SessionContext_openView_Params_Promise) View() UiView {
	return UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type SessionContext_openView_Results struct{ capnp.Struct }

func NewSessionContext_openView_Results(s *capnp.Segment) (SessionContext_openView_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_openView_Results{}, err
	}
	return SessionContext_openView_Results{st}, nil
}

func NewRootSessionContext_openView_Results(s *capnp.Segment) (SessionContext_openView_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SessionContext_openView_Results{}, err
	}
	return SessionContext_openView_Results{st}, nil
}

func ReadRootSessionContext_openView_Results(msg *capnp.Message) (SessionContext_openView_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SessionContext_openView_Results{}, err
	}
	return SessionContext_openView_Results{root.Struct()}, nil
}

// SessionContext_openView_Results_List is a list of SessionContext_openView_Results.
type SessionContext_openView_Results_List struct{ capnp.List }

// NewSessionContext_openView_Results creates a new list of SessionContext_openView_Results.
func NewSessionContext_openView_Results_List(s *capnp.Segment, sz int32) (SessionContext_openView_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SessionContext_openView_Results_List{}, err
	}
	return SessionContext_openView_Results_List{l}, nil
}

func (s SessionContext_openView_Results_List) At(i int) SessionContext_openView_Results {
	return SessionContext_openView_Results{s.List.Struct(i)}
}
func (s SessionContext_openView_Results_List) Set(i int, v SessionContext_openView_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SessionContext_openView_Results_Promise is a wrapper for a SessionContext_openView_Results promised by a client call.
type SessionContext_openView_Results_Promise struct{ *capnp.Pipeline }

func (p SessionContext_openView_Results_Promise) Struct() (SessionContext_openView_Results, error) {
	s, err := p.Pipeline.Struct()
	return SessionContext_openView_Results{s}, err
}

type PermissionDef struct{ capnp.Struct }

func NewPermissionDef(s *capnp.Segment) (PermissionDef, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return PermissionDef{}, err
	}
	return PermissionDef{st}, nil
}

func NewRootPermissionDef(s *capnp.Segment) (PermissionDef, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return PermissionDef{}, err
	}
	return PermissionDef{st}, nil
}

func ReadRootPermissionDef(msg *capnp.Message) (PermissionDef, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return PermissionDef{}, err
	}
	return PermissionDef{root.Struct()}, nil
}

func (s PermissionDef) Name() (string, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s PermissionDef) HasName() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PermissionDef) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s PermissionDef) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s PermissionDef) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s PermissionDef) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PermissionDef) SetTitle(v util.LocalizedText) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PermissionDef) NewTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PermissionDef) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s PermissionDef) HasDescription() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PermissionDef) SetDescription(v util.LocalizedText) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PermissionDef) NewDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PermissionDef) Obsolete() bool {
	return s.Struct.Bit(0)
}

func (s PermissionDef) SetObsolete(v bool) {

	s.Struct.SetBit(0, v)
}

// PermissionDef_List is a list of PermissionDef.
type PermissionDef_List struct{ capnp.List }

// NewPermissionDef creates a new list of PermissionDef.
func NewPermissionDef_List(s *capnp.Segment, sz int32) (PermissionDef_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return PermissionDef_List{}, err
	}
	return PermissionDef_List{l}, nil
}

func (s PermissionDef_List) At(i int) PermissionDef           { return PermissionDef{s.List.Struct(i)} }
func (s PermissionDef_List) Set(i int, v PermissionDef) error { return s.List.SetStruct(i, v.Struct) }

// PermissionDef_Promise is a wrapper for a PermissionDef promised by a client call.
type PermissionDef_Promise struct{ *capnp.Pipeline }

func (p PermissionDef_Promise) Struct() (PermissionDef, error) {
	s, err := p.Pipeline.Struct()
	return PermissionDef{s}, err
}

func (p PermissionDef_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PermissionDef_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type RoleDef struct{ capnp.Struct }

func NewRoleDef(s *capnp.Segment) (RoleDef, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return RoleDef{}, err
	}
	return RoleDef{st}, nil
}

func NewRootRoleDef(s *capnp.Segment) (RoleDef, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return RoleDef{}, err
	}
	return RoleDef{st}, nil
}

func ReadRootRoleDef(msg *capnp.Message) (RoleDef, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return RoleDef{}, err
	}
	return RoleDef{root.Struct()}, nil
}

func (s RoleDef) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s RoleDef) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetTitle(v util.LocalizedText) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewTitle() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s RoleDef) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetVerbPhrase(v util.LocalizedText) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewVerbPhrase() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s RoleDef) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetDescription(v util.LocalizedText) error {

	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s RoleDef) NewDescription() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s RoleDef) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s RoleDef) HasPermissions() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s RoleDef) SetPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(3, v.List.ToPtr())
}

func (s RoleDef) Obsolete() bool {
	return s.Struct.Bit(0)
}

func (s RoleDef) SetObsolete(v bool) {

	s.Struct.SetBit(0, v)
}

func (s RoleDef) Default() bool {
	return s.Struct.Bit(1)
}

func (s RoleDef) SetDefault(v bool) {

	s.Struct.SetBit(1, v)
}

// RoleDef_List is a list of RoleDef.
type RoleDef_List struct{ capnp.List }

// NewRoleDef creates a new list of RoleDef.
func NewRoleDef_List(s *capnp.Segment, sz int32) (RoleDef_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	if err != nil {
		return RoleDef_List{}, err
	}
	return RoleDef_List{l}, nil
}

func (s RoleDef_List) At(i int) RoleDef           { return RoleDef{s.List.Struct(i)} }
func (s RoleDef_List) Set(i int, v RoleDef) error { return s.List.SetStruct(i, v.Struct) }

// RoleDef_Promise is a wrapper for a RoleDef promised by a client call.
type RoleDef_Promise struct{ *capnp.Pipeline }

func (p RoleDef_Promise) Struct() (RoleDef, error) {
	s, err := p.Pipeline.Struct()
	return RoleDef{s}, err
}

func (p RoleDef_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p RoleDef_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p RoleDef_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type SharingLink struct{ Client capnp.Client }

func (c SharingLink) GetPetname(ctx context.Context, params func(SharingLink_getPetname_Params) error, opts ...capnp.CallOption) SharingLink_getPetname_Results_Promise {
	if c.Client == nil {
		return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SharingLink_getPetname_Params{Struct: s}) }
	}
	return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type SharingLink_Server interface {
	GetPetname(SharingLink_getPetname) error
}

func SharingLink_ServerToClient(s SharingLink_Server) SharingLink {
	c, _ := s.(server.Closer)
	return SharingLink{Client: server.New(SharingLink_Methods(nil, s), c)}
}

func SharingLink_Methods(methods []server.Method, s SharingLink_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SharingLink_getPetname{c, opts, SharingLink_getPetname_Params{Struct: p}, SharingLink_getPetname_Results{Struct: r}}
			return s.GetPetname(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// SharingLink_getPetname holds the arguments for a server call to SharingLink.getPetname.
type SharingLink_getPetname struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  SharingLink_getPetname_Params
	Results SharingLink_getPetname_Results
}

type SharingLink_getPetname_Params struct{ capnp.Struct }

func NewSharingLink_getPetname_Params(s *capnp.Segment) (SharingLink_getPetname_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SharingLink_getPetname_Params{}, err
	}
	return SharingLink_getPetname_Params{st}, nil
}

func NewRootSharingLink_getPetname_Params(s *capnp.Segment) (SharingLink_getPetname_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return SharingLink_getPetname_Params{}, err
	}
	return SharingLink_getPetname_Params{st}, nil
}

func ReadRootSharingLink_getPetname_Params(msg *capnp.Message) (SharingLink_getPetname_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SharingLink_getPetname_Params{}, err
	}
	return SharingLink_getPetname_Params{root.Struct()}, nil
}

// SharingLink_getPetname_Params_List is a list of SharingLink_getPetname_Params.
type SharingLink_getPetname_Params_List struct{ capnp.List }

// NewSharingLink_getPetname_Params creates a new list of SharingLink_getPetname_Params.
func NewSharingLink_getPetname_Params_List(s *capnp.Segment, sz int32) (SharingLink_getPetname_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return SharingLink_getPetname_Params_List{}, err
	}
	return SharingLink_getPetname_Params_List{l}, nil
}

func (s SharingLink_getPetname_Params_List) At(i int) SharingLink_getPetname_Params {
	return SharingLink_getPetname_Params{s.List.Struct(i)}
}
func (s SharingLink_getPetname_Params_List) Set(i int, v SharingLink_getPetname_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// SharingLink_getPetname_Params_Promise is a wrapper for a SharingLink_getPetname_Params promised by a client call.
type SharingLink_getPetname_Params_Promise struct{ *capnp.Pipeline }

func (p SharingLink_getPetname_Params_Promise) Struct() (SharingLink_getPetname_Params, error) {
	s, err := p.Pipeline.Struct()
	return SharingLink_getPetname_Params{s}, err
}

type SharingLink_getPetname_Results struct{ capnp.Struct }

func NewSharingLink_getPetname_Results(s *capnp.Segment) (SharingLink_getPetname_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SharingLink_getPetname_Results{}, err
	}
	return SharingLink_getPetname_Results{st}, nil
}

func NewRootSharingLink_getPetname_Results(s *capnp.Segment) (SharingLink_getPetname_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return SharingLink_getPetname_Results{}, err
	}
	return SharingLink_getPetname_Results{st}, nil
}

func ReadRootSharingLink_getPetname_Results(msg *capnp.Message) (SharingLink_getPetname_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return SharingLink_getPetname_Results{}, err
	}
	return SharingLink_getPetname_Results{root.Struct()}, nil
}

func (s SharingLink_getPetname_Results) Name() util.Assignable {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Assignable{}
	}
	return util.Assignable{Client: p.Interface().Client()}
}

func (s SharingLink_getPetname_Results) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SharingLink_getPetname_Results) SetName(v util.Assignable) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// SharingLink_getPetname_Results_List is a list of SharingLink_getPetname_Results.
type SharingLink_getPetname_Results_List struct{ capnp.List }

// NewSharingLink_getPetname_Results creates a new list of SharingLink_getPetname_Results.
func NewSharingLink_getPetname_Results_List(s *capnp.Segment, sz int32) (SharingLink_getPetname_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return SharingLink_getPetname_Results_List{}, err
	}
	return SharingLink_getPetname_Results_List{l}, nil
}

func (s SharingLink_getPetname_Results_List) At(i int) SharingLink_getPetname_Results {
	return SharingLink_getPetname_Results{s.List.Struct(i)}
}
func (s SharingLink_getPetname_Results_List) Set(i int, v SharingLink_getPetname_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// SharingLink_getPetname_Results_Promise is a wrapper for a SharingLink_getPetname_Results promised by a client call.
type SharingLink_getPetname_Results_Promise struct{ *capnp.Pipeline }

func (p SharingLink_getPetname_Results_Promise) Struct() (SharingLink_getPetname_Results, error) {
	s, err := p.Pipeline.Struct()
	return SharingLink_getPetname_Results{s}, err
}

func (p SharingLink_getPetname_Results_Promise) Name() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type ViewSharingLink struct{ Client capnp.Client }

func (c ViewSharingLink) GetRoleAssignment(ctx context.Context, params func(ViewSharingLink_getRoleAssignment_Params) error, opts ...capnp.CallOption) ViewSharingLink_getRoleAssignment_Results_Promise {
	if c.Client == nil {
		return ViewSharingLink_getRoleAssignment_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xa4f82f764dc3fee8,
			MethodID:      0,
			InterfaceName: "grain.capnp:ViewSharingLink",
			MethodName:    "getRoleAssignment",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ViewSharingLink_getRoleAssignment_Params{Struct: s}) }
	}
	return ViewSharingLink_getRoleAssignment_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c ViewSharingLink) GetPetname(ctx context.Context, params func(SharingLink_getPetname_Params) error, opts ...capnp.CallOption) SharingLink_getPetname_Results_Promise {
	if c.Client == nil {
		return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(SharingLink_getPetname_Params{Struct: s}) }
	}
	return SharingLink_getPetname_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ViewSharingLink_Server interface {
	GetRoleAssignment(ViewSharingLink_getRoleAssignment) error

	GetPetname(SharingLink_getPetname) error
}

func ViewSharingLink_ServerToClient(s ViewSharingLink_Server) ViewSharingLink {
	c, _ := s.(server.Closer)
	return ViewSharingLink{Client: server.New(ViewSharingLink_Methods(nil, s), c)}
}

func ViewSharingLink_Methods(methods []server.Method, s ViewSharingLink_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xa4f82f764dc3fee8,
			MethodID:      0,
			InterfaceName: "grain.capnp:ViewSharingLink",
			MethodName:    "getRoleAssignment",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ViewSharingLink_getRoleAssignment{c, opts, ViewSharingLink_getRoleAssignment_Params{Struct: p}, ViewSharingLink_getRoleAssignment_Results{Struct: r}}
			return s.GetRoleAssignment(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xc05520c9b0994ad3,
			MethodID:      0,
			InterfaceName: "grain.capnp:SharingLink",
			MethodName:    "getPetname",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := SharingLink_getPetname{c, opts, SharingLink_getPetname_Params{Struct: p}, SharingLink_getPetname_Results{Struct: r}}
			return s.GetPetname(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// ViewSharingLink_getRoleAssignment holds the arguments for a server call to ViewSharingLink.getRoleAssignment.
type ViewSharingLink_getRoleAssignment struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ViewSharingLink_getRoleAssignment_Params
	Results ViewSharingLink_getRoleAssignment_Results
}

type ViewSharingLink_RoleAssignment struct{ capnp.Struct }
type ViewSharingLink_RoleAssignment_Which uint16

const (
	ViewSharingLink_RoleAssignment_Which_none      ViewSharingLink_RoleAssignment_Which = 0
	ViewSharingLink_RoleAssignment_Which_allAccess ViewSharingLink_RoleAssignment_Which = 1
	ViewSharingLink_RoleAssignment_Which_roleId    ViewSharingLink_RoleAssignment_Which = 2
)

func (w ViewSharingLink_RoleAssignment_Which) String() string {
	const s = "noneallAccessroleId"
	switch w {
	case ViewSharingLink_RoleAssignment_Which_none:
		return s[0:4]
	case ViewSharingLink_RoleAssignment_Which_allAccess:
		return s[4:13]
	case ViewSharingLink_RoleAssignment_Which_roleId:
		return s[13:19]

	}
	return "ViewSharingLink_RoleAssignment_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewViewSharingLink_RoleAssignment(s *capnp.Segment) (ViewSharingLink_RoleAssignment, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return ViewSharingLink_RoleAssignment{}, err
	}
	return ViewSharingLink_RoleAssignment{st}, nil
}

func NewRootViewSharingLink_RoleAssignment(s *capnp.Segment) (ViewSharingLink_RoleAssignment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return ViewSharingLink_RoleAssignment{}, err
	}
	return ViewSharingLink_RoleAssignment{st}, nil
}

func ReadRootViewSharingLink_RoleAssignment(msg *capnp.Message) (ViewSharingLink_RoleAssignment, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ViewSharingLink_RoleAssignment{}, err
	}
	return ViewSharingLink_RoleAssignment{root.Struct()}, nil
}

func (s ViewSharingLink_RoleAssignment) Which() ViewSharingLink_RoleAssignment_Which {
	return ViewSharingLink_RoleAssignment_Which(s.Struct.Uint16(0))
}

func (s ViewSharingLink_RoleAssignment) SetNone() {
	s.Struct.SetUint16(0, 0)
}

func (s ViewSharingLink_RoleAssignment) SetAllAccess() {
	s.Struct.SetUint16(0, 1)
}

func (s ViewSharingLink_RoleAssignment) RoleId() uint16 {
	return s.Struct.Uint16(2)
}

func (s ViewSharingLink_RoleAssignment) SetRoleId(v uint16) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint16(2, v)
}

func (s ViewSharingLink_RoleAssignment) AddPermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s ViewSharingLink_RoleAssignment) HasAddPermissions() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_RoleAssignment) SetAddPermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(0, v.List.ToPtr())
}

func (s ViewSharingLink_RoleAssignment) RemovePermissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return capnp.BitList{}, err
	}

	return capnp.BitList{List: p.List()}, nil

}

func (s ViewSharingLink_RoleAssignment) HasRemovePermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_RoleAssignment) SetRemovePermissions(v capnp.BitList) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// ViewSharingLink_RoleAssignment_List is a list of ViewSharingLink_RoleAssignment.
type ViewSharingLink_RoleAssignment_List struct{ capnp.List }

// NewViewSharingLink_RoleAssignment creates a new list of ViewSharingLink_RoleAssignment.
func NewViewSharingLink_RoleAssignment_List(s *capnp.Segment, sz int32) (ViewSharingLink_RoleAssignment_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return ViewSharingLink_RoleAssignment_List{}, err
	}
	return ViewSharingLink_RoleAssignment_List{l}, nil
}

func (s ViewSharingLink_RoleAssignment_List) At(i int) ViewSharingLink_RoleAssignment {
	return ViewSharingLink_RoleAssignment{s.List.Struct(i)}
}
func (s ViewSharingLink_RoleAssignment_List) Set(i int, v ViewSharingLink_RoleAssignment) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_RoleAssignment_Promise is a wrapper for a ViewSharingLink_RoleAssignment promised by a client call.
type ViewSharingLink_RoleAssignment_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_RoleAssignment_Promise) Struct() (ViewSharingLink_RoleAssignment, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_RoleAssignment{s}, err
}

type ViewSharingLink_getRoleAssignment_Params struct{ capnp.Struct }

func NewViewSharingLink_getRoleAssignment_Params(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Params{}, err
	}
	return ViewSharingLink_getRoleAssignment_Params{st}, nil
}

func NewRootViewSharingLink_getRoleAssignment_Params(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Params{}, err
	}
	return ViewSharingLink_getRoleAssignment_Params{st}, nil
}

func ReadRootViewSharingLink_getRoleAssignment_Params(msg *capnp.Message) (ViewSharingLink_getRoleAssignment_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Params{}, err
	}
	return ViewSharingLink_getRoleAssignment_Params{root.Struct()}, nil
}

// ViewSharingLink_getRoleAssignment_Params_List is a list of ViewSharingLink_getRoleAssignment_Params.
type ViewSharingLink_getRoleAssignment_Params_List struct{ capnp.List }

// NewViewSharingLink_getRoleAssignment_Params creates a new list of ViewSharingLink_getRoleAssignment_Params.
func NewViewSharingLink_getRoleAssignment_Params_List(s *capnp.Segment, sz int32) (ViewSharingLink_getRoleAssignment_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Params_List{}, err
	}
	return ViewSharingLink_getRoleAssignment_Params_List{l}, nil
}

func (s ViewSharingLink_getRoleAssignment_Params_List) At(i int) ViewSharingLink_getRoleAssignment_Params {
	return ViewSharingLink_getRoleAssignment_Params{s.List.Struct(i)}
}
func (s ViewSharingLink_getRoleAssignment_Params_List) Set(i int, v ViewSharingLink_getRoleAssignment_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_getRoleAssignment_Params_Promise is a wrapper for a ViewSharingLink_getRoleAssignment_Params promised by a client call.
type ViewSharingLink_getRoleAssignment_Params_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_getRoleAssignment_Params_Promise) Struct() (ViewSharingLink_getRoleAssignment_Params, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_getRoleAssignment_Params{s}, err
}

type ViewSharingLink_getRoleAssignment_Results struct{ capnp.Struct }

func NewViewSharingLink_getRoleAssignment_Results(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Results{}, err
	}
	return ViewSharingLink_getRoleAssignment_Results{st}, nil
}

func NewRootViewSharingLink_getRoleAssignment_Results(s *capnp.Segment) (ViewSharingLink_getRoleAssignment_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Results{}, err
	}
	return ViewSharingLink_getRoleAssignment_Results{st}, nil
}

func ReadRootViewSharingLink_getRoleAssignment_Results(msg *capnp.Message) (ViewSharingLink_getRoleAssignment_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Results{}, err
	}
	return ViewSharingLink_getRoleAssignment_Results{root.Struct()}, nil
}

func (s ViewSharingLink_getRoleAssignment_Results) Var() util.Assignable {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Assignable{}
	}
	return util.Assignable{Client: p.Interface().Client()}
}

func (s ViewSharingLink_getRoleAssignment_Results) HasVar() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ViewSharingLink_getRoleAssignment_Results) SetVar(v util.Assignable) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// ViewSharingLink_getRoleAssignment_Results_List is a list of ViewSharingLink_getRoleAssignment_Results.
type ViewSharingLink_getRoleAssignment_Results_List struct{ capnp.List }

// NewViewSharingLink_getRoleAssignment_Results creates a new list of ViewSharingLink_getRoleAssignment_Results.
func NewViewSharingLink_getRoleAssignment_Results_List(s *capnp.Segment, sz int32) (ViewSharingLink_getRoleAssignment_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return ViewSharingLink_getRoleAssignment_Results_List{}, err
	}
	return ViewSharingLink_getRoleAssignment_Results_List{l}, nil
}

func (s ViewSharingLink_getRoleAssignment_Results_List) At(i int) ViewSharingLink_getRoleAssignment_Results {
	return ViewSharingLink_getRoleAssignment_Results{s.List.Struct(i)}
}
func (s ViewSharingLink_getRoleAssignment_Results_List) Set(i int, v ViewSharingLink_getRoleAssignment_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ViewSharingLink_getRoleAssignment_Results_Promise is a wrapper for a ViewSharingLink_getRoleAssignment_Results promised by a client call.
type ViewSharingLink_getRoleAssignment_Results_Promise struct{ *capnp.Pipeline }

func (p ViewSharingLink_getRoleAssignment_Results_Promise) Struct() (ViewSharingLink_getRoleAssignment_Results, error) {
	s, err := p.Pipeline.Struct()
	return ViewSharingLink_getRoleAssignment_Results{s}, err
}

func (p ViewSharingLink_getRoleAssignment_Results_Promise) Var() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type NotificationDisplayInfo struct{ capnp.Struct }

func NewNotificationDisplayInfo(s *capnp.Segment) (NotificationDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	return NotificationDisplayInfo{st}, nil
}

func NewRootNotificationDisplayInfo(s *capnp.Segment) (NotificationDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	return NotificationDisplayInfo{st}, nil
}

func ReadRootNotificationDisplayInfo(msg *capnp.Message) (NotificationDisplayInfo, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	return NotificationDisplayInfo{root.Struct()}, nil
}

func (s NotificationDisplayInfo) Caption() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s NotificationDisplayInfo) HasCaption() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NotificationDisplayInfo) SetCaption(v util.LocalizedText) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCaption sets the caption field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s NotificationDisplayInfo) NewCaption() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// NotificationDisplayInfo_List is a list of NotificationDisplayInfo.
type NotificationDisplayInfo_List struct{ capnp.List }

// NewNotificationDisplayInfo creates a new list of NotificationDisplayInfo.
func NewNotificationDisplayInfo_List(s *capnp.Segment, sz int32) (NotificationDisplayInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return NotificationDisplayInfo_List{}, err
	}
	return NotificationDisplayInfo_List{l}, nil
}

func (s NotificationDisplayInfo_List) At(i int) NotificationDisplayInfo {
	return NotificationDisplayInfo{s.List.Struct(i)}
}
func (s NotificationDisplayInfo_List) Set(i int, v NotificationDisplayInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// NotificationDisplayInfo_Promise is a wrapper for a NotificationDisplayInfo promised by a client call.
type NotificationDisplayInfo_Promise struct{ *capnp.Pipeline }

func (p NotificationDisplayInfo_Promise) Struct() (NotificationDisplayInfo, error) {
	s, err := p.Pipeline.Struct()
	return NotificationDisplayInfo{s}, err
}

func (p NotificationDisplayInfo_Promise) Caption() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type NotificationTarget struct{ Client capnp.Client }

func (c NotificationTarget) AddOngoing(ctx context.Context, params func(NotificationTarget_addOngoing_Params) error, opts ...capnp.CallOption) NotificationTarget_addOngoing_Results_Promise {
	if c.Client == nil {
		return NotificationTarget_addOngoing_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xf0f87337d73020f0,
			MethodID:      0,
			InterfaceName: "grain.capnp:NotificationTarget",
			MethodName:    "addOngoing",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(NotificationTarget_addOngoing_Params{Struct: s}) }
	}
	return NotificationTarget_addOngoing_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type NotificationTarget_Server interface {
	AddOngoing(NotificationTarget_addOngoing) error
}

func NotificationTarget_ServerToClient(s NotificationTarget_Server) NotificationTarget {
	c, _ := s.(server.Closer)
	return NotificationTarget{Client: server.New(NotificationTarget_Methods(nil, s), c)}
}

func NotificationTarget_Methods(methods []server.Method, s NotificationTarget_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xf0f87337d73020f0,
			MethodID:      0,
			InterfaceName: "grain.capnp:NotificationTarget",
			MethodName:    "addOngoing",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := NotificationTarget_addOngoing{c, opts, NotificationTarget_addOngoing_Params{Struct: p}, NotificationTarget_addOngoing_Results{Struct: r}}
			return s.AddOngoing(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// NotificationTarget_addOngoing holds the arguments for a server call to NotificationTarget.addOngoing.
type NotificationTarget_addOngoing struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  NotificationTarget_addOngoing_Params
	Results NotificationTarget_addOngoing_Results
}

type NotificationTarget_addOngoing_Params struct{ capnp.Struct }

func NewNotificationTarget_addOngoing_Params(s *capnp.Segment) (NotificationTarget_addOngoing_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return NotificationTarget_addOngoing_Params{}, err
	}
	return NotificationTarget_addOngoing_Params{st}, nil
}

func NewRootNotificationTarget_addOngoing_Params(s *capnp.Segment) (NotificationTarget_addOngoing_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return NotificationTarget_addOngoing_Params{}, err
	}
	return NotificationTarget_addOngoing_Params{st}, nil
}

func ReadRootNotificationTarget_addOngoing_Params(msg *capnp.Message) (NotificationTarget_addOngoing_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return NotificationTarget_addOngoing_Params{}, err
	}
	return NotificationTarget_addOngoing_Params{root.Struct()}, nil
}

func (s NotificationTarget_addOngoing_Params) DisplayInfo() (NotificationDisplayInfo, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return NotificationDisplayInfo{}, err
	}

	return NotificationDisplayInfo{Struct: p.Struct()}, nil

}

func (s NotificationTarget_addOngoing_Params) HasDisplayInfo() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NotificationTarget_addOngoing_Params) SetDisplayInfo(v NotificationDisplayInfo) error {

	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisplayInfo sets the displayInfo field to a newly
// allocated NotificationDisplayInfo struct, preferring placement in s's segment.
func (s NotificationTarget_addOngoing_Params) NewDisplayInfo() (NotificationDisplayInfo, error) {

	ss, err := NewNotificationDisplayInfo(s.Struct.Segment())
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s NotificationTarget_addOngoing_Params) Notification() OngoingNotification {
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return OngoingNotification{}
	}
	return OngoingNotification{Client: p.Interface().Client()}
}

func (s NotificationTarget_addOngoing_Params) HasNotification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s NotificationTarget_addOngoing_Params) SetNotification(v OngoingNotification) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(1, in.ToPtr())
}

// NotificationTarget_addOngoing_Params_List is a list of NotificationTarget_addOngoing_Params.
type NotificationTarget_addOngoing_Params_List struct{ capnp.List }

// NewNotificationTarget_addOngoing_Params creates a new list of NotificationTarget_addOngoing_Params.
func NewNotificationTarget_addOngoing_Params_List(s *capnp.Segment, sz int32) (NotificationTarget_addOngoing_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return NotificationTarget_addOngoing_Params_List{}, err
	}
	return NotificationTarget_addOngoing_Params_List{l}, nil
}

func (s NotificationTarget_addOngoing_Params_List) At(i int) NotificationTarget_addOngoing_Params {
	return NotificationTarget_addOngoing_Params{s.List.Struct(i)}
}
func (s NotificationTarget_addOngoing_Params_List) Set(i int, v NotificationTarget_addOngoing_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// NotificationTarget_addOngoing_Params_Promise is a wrapper for a NotificationTarget_addOngoing_Params promised by a client call.
type NotificationTarget_addOngoing_Params_Promise struct{ *capnp.Pipeline }

func (p NotificationTarget_addOngoing_Params_Promise) Struct() (NotificationTarget_addOngoing_Params, error) {
	s, err := p.Pipeline.Struct()
	return NotificationTarget_addOngoing_Params{s}, err
}

func (p NotificationTarget_addOngoing_Params_Promise) DisplayInfo() NotificationDisplayInfo_Promise {
	return NotificationDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p NotificationTarget_addOngoing_Params_Promise) Notification() OngoingNotification {
	return OngoingNotification{Client: p.Pipeline.GetPipeline(1).Client()}
}

type NotificationTarget_addOngoing_Results struct{ capnp.Struct }

func NewNotificationTarget_addOngoing_Results(s *capnp.Segment) (NotificationTarget_addOngoing_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return NotificationTarget_addOngoing_Results{}, err
	}
	return NotificationTarget_addOngoing_Results{st}, nil
}

func NewRootNotificationTarget_addOngoing_Results(s *capnp.Segment) (NotificationTarget_addOngoing_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return NotificationTarget_addOngoing_Results{}, err
	}
	return NotificationTarget_addOngoing_Results{st}, nil
}

func ReadRootNotificationTarget_addOngoing_Results(msg *capnp.Message) (NotificationTarget_addOngoing_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return NotificationTarget_addOngoing_Results{}, err
	}
	return NotificationTarget_addOngoing_Results{root.Struct()}, nil
}

func (s NotificationTarget_addOngoing_Results) Handle() util.Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return util.Handle{}
	}
	return util.Handle{Client: p.Interface().Client()}
}

func (s NotificationTarget_addOngoing_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NotificationTarget_addOngoing_Results) SetHandle(v util.Handle) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// NotificationTarget_addOngoing_Results_List is a list of NotificationTarget_addOngoing_Results.
type NotificationTarget_addOngoing_Results_List struct{ capnp.List }

// NewNotificationTarget_addOngoing_Results creates a new list of NotificationTarget_addOngoing_Results.
func NewNotificationTarget_addOngoing_Results_List(s *capnp.Segment, sz int32) (NotificationTarget_addOngoing_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return NotificationTarget_addOngoing_Results_List{}, err
	}
	return NotificationTarget_addOngoing_Results_List{l}, nil
}

func (s NotificationTarget_addOngoing_Results_List) At(i int) NotificationTarget_addOngoing_Results {
	return NotificationTarget_addOngoing_Results{s.List.Struct(i)}
}
func (s NotificationTarget_addOngoing_Results_List) Set(i int, v NotificationTarget_addOngoing_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// NotificationTarget_addOngoing_Results_Promise is a wrapper for a NotificationTarget_addOngoing_Results promised by a client call.
type NotificationTarget_addOngoing_Results_Promise struct{ *capnp.Pipeline }

func (p NotificationTarget_addOngoing_Results_Promise) Struct() (NotificationTarget_addOngoing_Results, error) {
	s, err := p.Pipeline.Struct()
	return NotificationTarget_addOngoing_Results{s}, err
}

func (p NotificationTarget_addOngoing_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type OngoingNotification struct{ Client capnp.Client }

func (c OngoingNotification) Cancel(ctx context.Context, params func(OngoingNotification_cancel_Params) error, opts ...capnp.CallOption) OngoingNotification_cancel_Results_Promise {
	if c.Client == nil {
		return OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xfe851ddbb88940cd,
			MethodID:      0,
			InterfaceName: "grain.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(OngoingNotification_cancel_Params{Struct: s}) }
	}
	return OngoingNotification_cancel_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type OngoingNotification_Server interface {
	Cancel(OngoingNotification_cancel) error
}

func OngoingNotification_ServerToClient(s OngoingNotification_Server) OngoingNotification {
	c, _ := s.(server.Closer)
	return OngoingNotification{Client: server.New(OngoingNotification_Methods(nil, s), c)}
}

func OngoingNotification_Methods(methods []server.Method, s OngoingNotification_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xfe851ddbb88940cd,
			MethodID:      0,
			InterfaceName: "grain.capnp:OngoingNotification",
			MethodName:    "cancel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := OngoingNotification_cancel{c, opts, OngoingNotification_cancel_Params{Struct: p}, OngoingNotification_cancel_Results{Struct: r}}
			return s.Cancel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// OngoingNotification_cancel holds the arguments for a server call to OngoingNotification.cancel.
type OngoingNotification_cancel struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  OngoingNotification_cancel_Params
	Results OngoingNotification_cancel_Results
}

type OngoingNotification_cancel_Params struct{ capnp.Struct }

func NewOngoingNotification_cancel_Params(s *capnp.Segment) (OngoingNotification_cancel_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return OngoingNotification_cancel_Params{}, err
	}
	return OngoingNotification_cancel_Params{st}, nil
}

func NewRootOngoingNotification_cancel_Params(s *capnp.Segment) (OngoingNotification_cancel_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return OngoingNotification_cancel_Params{}, err
	}
	return OngoingNotification_cancel_Params{st}, nil
}

func ReadRootOngoingNotification_cancel_Params(msg *capnp.Message) (OngoingNotification_cancel_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return OngoingNotification_cancel_Params{}, err
	}
	return OngoingNotification_cancel_Params{root.Struct()}, nil
}

// OngoingNotification_cancel_Params_List is a list of OngoingNotification_cancel_Params.
type OngoingNotification_cancel_Params_List struct{ capnp.List }

// NewOngoingNotification_cancel_Params creates a new list of OngoingNotification_cancel_Params.
func NewOngoingNotification_cancel_Params_List(s *capnp.Segment, sz int32) (OngoingNotification_cancel_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return OngoingNotification_cancel_Params_List{}, err
	}
	return OngoingNotification_cancel_Params_List{l}, nil
}

func (s OngoingNotification_cancel_Params_List) At(i int) OngoingNotification_cancel_Params {
	return OngoingNotification_cancel_Params{s.List.Struct(i)}
}
func (s OngoingNotification_cancel_Params_List) Set(i int, v OngoingNotification_cancel_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// OngoingNotification_cancel_Params_Promise is a wrapper for a OngoingNotification_cancel_Params promised by a client call.
type OngoingNotification_cancel_Params_Promise struct{ *capnp.Pipeline }

func (p OngoingNotification_cancel_Params_Promise) Struct() (OngoingNotification_cancel_Params, error) {
	s, err := p.Pipeline.Struct()
	return OngoingNotification_cancel_Params{s}, err
}

type OngoingNotification_cancel_Results struct{ capnp.Struct }

func NewOngoingNotification_cancel_Results(s *capnp.Segment) (OngoingNotification_cancel_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return OngoingNotification_cancel_Results{}, err
	}
	return OngoingNotification_cancel_Results{st}, nil
}

func NewRootOngoingNotification_cancel_Results(s *capnp.Segment) (OngoingNotification_cancel_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return OngoingNotification_cancel_Results{}, err
	}
	return OngoingNotification_cancel_Results{st}, nil
}

func ReadRootOngoingNotification_cancel_Results(msg *capnp.Message) (OngoingNotification_cancel_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return OngoingNotification_cancel_Results{}, err
	}
	return OngoingNotification_cancel_Results{root.Struct()}, nil
}

// OngoingNotification_cancel_Results_List is a list of OngoingNotification_cancel_Results.
type OngoingNotification_cancel_Results_List struct{ capnp.List }

// NewOngoingNotification_cancel_Results creates a new list of OngoingNotification_cancel_Results.
func NewOngoingNotification_cancel_Results_List(s *capnp.Segment, sz int32) (OngoingNotification_cancel_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return OngoingNotification_cancel_Results_List{}, err
	}
	return OngoingNotification_cancel_Results_List{l}, nil
}

func (s OngoingNotification_cancel_Results_List) At(i int) OngoingNotification_cancel_Results {
	return OngoingNotification_cancel_Results{s.List.Struct(i)}
}
func (s OngoingNotification_cancel_Results_List) Set(i int, v OngoingNotification_cancel_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// OngoingNotification_cancel_Results_Promise is a wrapper for a OngoingNotification_cancel_Results promised by a client call.
type OngoingNotification_cancel_Results_Promise struct{ *capnp.Pipeline }

func (p OngoingNotification_cancel_Results_Promise) Struct() (OngoingNotification_cancel_Results, error) {
	s, err := p.Pipeline.Struct()
	return OngoingNotification_cancel_Results{s}, err
}

type GrainInfo struct{ capnp.Struct }

func NewGrainInfo(s *capnp.Segment) (GrainInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return GrainInfo{}, err
	}
	return GrainInfo{st}, nil
}

func NewRootGrainInfo(s *capnp.Segment) (GrainInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return GrainInfo{}, err
	}
	return GrainInfo{st}, nil
}

func ReadRootGrainInfo(msg *capnp.Message) (GrainInfo, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return GrainInfo{}, err
	}
	return GrainInfo{root.Struct()}, nil
}

func (s GrainInfo) AppId() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s GrainInfo) HasAppId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GrainInfo) AppIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s GrainInfo) SetAppId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s GrainInfo) AppVersion() uint32 {
	return s.Struct.Uint32(0)
}

func (s GrainInfo) SetAppVersion(v uint32) {

	s.Struct.SetUint32(0, v)
}

func (s GrainInfo) Title() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s GrainInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GrainInfo) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

}

func (s GrainInfo) SetTitle(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// GrainInfo_List is a list of GrainInfo.
type GrainInfo_List struct{ capnp.List }

// NewGrainInfo creates a new list of GrainInfo.
func NewGrainInfo_List(s *capnp.Segment, sz int32) (GrainInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return GrainInfo_List{}, err
	}
	return GrainInfo_List{l}, nil
}

func (s GrainInfo_List) At(i int) GrainInfo           { return GrainInfo{s.List.Struct(i)} }
func (s GrainInfo_List) Set(i int, v GrainInfo) error { return s.List.SetStruct(i, v.Struct) }

// GrainInfo_Promise is a wrapper for a GrainInfo promised by a client call.
type GrainInfo_Promise struct{ *capnp.Pipeline }

func (p GrainInfo_Promise) Struct() (GrainInfo, error) {
	s, err := p.Pipeline.Struct()
	return GrainInfo{s}, err
}

type AppPersistent struct{ Client capnp.Client }

func (c AppPersistent) Save(ctx context.Context, params func(AppPersistent_save_Params) error, opts ...capnp.CallOption) AppPersistent_save_Results_Promise {
	if c.Client == nil {
		return AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xaffa789add8747b8,
			MethodID:      0,
			InterfaceName: "grain.capnp:AppPersistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(AppPersistent_save_Params{Struct: s}) }
	}
	return AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type AppPersistent_Server interface {
	Save(AppPersistent_save) error
}

func AppPersistent_ServerToClient(s AppPersistent_Server) AppPersistent {
	c, _ := s.(server.Closer)
	return AppPersistent{Client: server.New(AppPersistent_Methods(nil, s), c)}
}

func AppPersistent_Methods(methods []server.Method, s AppPersistent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xaffa789add8747b8,
			MethodID:      0,
			InterfaceName: "grain.capnp:AppPersistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := AppPersistent_save{c, opts, AppPersistent_save_Params{Struct: p}, AppPersistent_save_Results{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	return methods
}

// AppPersistent_save holds the arguments for a server call to AppPersistent.save.
type AppPersistent_save struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  AppPersistent_save_Params
	Results AppPersistent_save_Results
}

type AppPersistent_save_Params struct{ capnp.Struct }

func NewAppPersistent_save_Params(s *capnp.Segment) (AppPersistent_save_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return AppPersistent_save_Params{}, err
	}
	return AppPersistent_save_Params{st}, nil
}

func NewRootAppPersistent_save_Params(s *capnp.Segment) (AppPersistent_save_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return AppPersistent_save_Params{}, err
	}
	return AppPersistent_save_Params{st}, nil
}

func ReadRootAppPersistent_save_Params(msg *capnp.Message) (AppPersistent_save_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return AppPersistent_save_Params{}, err
	}
	return AppPersistent_save_Params{root.Struct()}, nil
}

// AppPersistent_save_Params_List is a list of AppPersistent_save_Params.
type AppPersistent_save_Params_List struct{ capnp.List }

// NewAppPersistent_save_Params creates a new list of AppPersistent_save_Params.
func NewAppPersistent_save_Params_List(s *capnp.Segment, sz int32) (AppPersistent_save_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return AppPersistent_save_Params_List{}, err
	}
	return AppPersistent_save_Params_List{l}, nil
}

func (s AppPersistent_save_Params_List) At(i int) AppPersistent_save_Params {
	return AppPersistent_save_Params{s.List.Struct(i)}
}
func (s AppPersistent_save_Params_List) Set(i int, v AppPersistent_save_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppPersistent_save_Params_Promise is a wrapper for a AppPersistent_save_Params promised by a client call.
type AppPersistent_save_Params_Promise struct{ *capnp.Pipeline }

func (p AppPersistent_save_Params_Promise) Struct() (AppPersistent_save_Params, error) {
	s, err := p.Pipeline.Struct()
	return AppPersistent_save_Params{s}, err
}

type AppPersistent_save_Results struct{ capnp.Struct }

func NewAppPersistent_save_Results(s *capnp.Segment) (AppPersistent_save_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return AppPersistent_save_Results{}, err
	}
	return AppPersistent_save_Results{st}, nil
}

func NewRootAppPersistent_save_Results(s *capnp.Segment) (AppPersistent_save_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return AppPersistent_save_Results{}, err
	}
	return AppPersistent_save_Results{st}, nil
}

func ReadRootAppPersistent_save_Results(msg *capnp.Message) (AppPersistent_save_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return AppPersistent_save_Results{}, err
	}
	return AppPersistent_save_Results{root.Struct()}, nil
}

func (s AppPersistent_save_Results) ObjectId() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s AppPersistent_save_Results) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s AppPersistent_save_Results) ObjectIdPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s AppPersistent_save_Results) SetObjectId(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s AppPersistent_save_Results) SetObjectIdPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

func (s AppPersistent_save_Results) Label() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return util.LocalizedText{}, err
	}

	return util.LocalizedText{Struct: p.Struct()}, nil

}

func (s AppPersistent_save_Results) HasLabel() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s AppPersistent_save_Results) SetLabel(v util.LocalizedText) error {

	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s AppPersistent_save_Results) NewLabel() (util.LocalizedText, error) {

	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// AppPersistent_save_Results_List is a list of AppPersistent_save_Results.
type AppPersistent_save_Results_List struct{ capnp.List }

// NewAppPersistent_save_Results creates a new list of AppPersistent_save_Results.
func NewAppPersistent_save_Results_List(s *capnp.Segment, sz int32) (AppPersistent_save_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return AppPersistent_save_Results_List{}, err
	}
	return AppPersistent_save_Results_List{l}, nil
}

func (s AppPersistent_save_Results_List) At(i int) AppPersistent_save_Results {
	return AppPersistent_save_Results{s.List.Struct(i)}
}
func (s AppPersistent_save_Results_List) Set(i int, v AppPersistent_save_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppPersistent_save_Results_Promise is a wrapper for a AppPersistent_save_Results promised by a client call.
type AppPersistent_save_Results_Promise struct{ *capnp.Pipeline }

func (p AppPersistent_save_Results_Promise) Struct() (AppPersistent_save_Results, error) {
	s, err := p.Pipeline.Struct()
	return AppPersistent_save_Results{s}, err
}

func (p AppPersistent_save_Results_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p AppPersistent_save_Results_Promise) Label() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type MainView struct{ Client capnp.Client }

func (c MainView) Restore(ctx context.Context, params func(MainView_restore_Params) error, opts ...capnp.CallOption) MainView_restore_Results_Promise {
	if c.Client == nil {
		return MainView_restore_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      0,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "restore",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(MainView_restore_Params{Struct: s}) }
	}
	return MainView_restore_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c MainView) Drop(ctx context.Context, params func(MainView_drop_Params) error, opts ...capnp.CallOption) MainView_drop_Results_Promise {
	if c.Client == nil {
		return MainView_drop_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      1,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "drop",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(MainView_drop_Params{Struct: s}) }
	}
	return MainView_drop_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c MainView) GetViewInfo(ctx context.Context, params func(UiView_getViewInfo_Params) error, opts ...capnp.CallOption) UiView_ViewInfo_Promise {
	if c.Client == nil {
		return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_getViewInfo_Params{Struct: s}) }
	}
	return UiView_ViewInfo_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c MainView) NewSession(ctx context.Context, params func(UiView_newSession_Params) error, opts ...capnp.CallOption) UiView_newSession_Results_Promise {
	if c.Client == nil {
		return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newSession_Params{Struct: s}) }
	}
	return UiView_newSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c MainView) NewRequestSession(ctx context.Context, params func(UiView_newRequestSession_Params) error, opts ...capnp.CallOption) UiView_newRequestSession_Results_Promise {
	if c.Client == nil {
		return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newRequestSession_Params{Struct: s}) }
	}
	return UiView_newRequestSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

func (c MainView) NewOfferSession(ctx context.Context, params func(UiView_newOfferSession_Params) error, opts ...capnp.CallOption) UiView_newOfferSession_Results_Promise {
	if c.Client == nil {
		return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 5}
		call.ParamsFunc = func(s capnp.Struct) error { return params(UiView_newOfferSession_Params{Struct: s}) }
	}
	return UiView_newOfferSession_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type MainView_Server interface {
	Restore(MainView_restore) error

	Drop(MainView_drop) error

	GetViewInfo(UiView_getViewInfo) error

	NewSession(UiView_newSession) error

	NewRequestSession(UiView_newRequestSession) error

	NewOfferSession(UiView_newOfferSession) error
}

func MainView_ServerToClient(s MainView_Server) MainView {
	c, _ := s.(server.Closer)
	return MainView{Client: server.New(MainView_Methods(nil, s), c)}
}

func MainView_Methods(methods []server.Method, s MainView_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      0,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "restore",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := MainView_restore{c, opts, MainView_restore_Params{Struct: p}, MainView_restore_Results{Struct: r}}
			return s.Restore(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xc277e9822ae2c8fc,
			MethodID:      1,
			InterfaceName: "grain.capnp:MainView",
			MethodName:    "drop",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := MainView_drop{c, opts, MainView_drop_Params{Struct: p}, MainView_drop_Results{Struct: r}}
			return s.Drop(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      0,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "getViewInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_getViewInfo{c, opts, UiView_getViewInfo_Params{Struct: p}, UiView_ViewInfo{Struct: r}}
			return s.GetViewInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 5},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      1,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newSession{c, opts, UiView_newSession_Params{Struct: p}, UiView_newSession_Results{Struct: r}}
			return s.NewSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      2,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newRequestSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newRequestSession{c, opts, UiView_newRequestSession_Params{Struct: p}, UiView_newRequestSession_Results{Struct: r}}
			return s.NewRequestSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{

			InterfaceID:   0xdbb4d798ea67e2e7,
			MethodID:      3,
			InterfaceName: "grain.capnp:UiView",
			MethodName:    "newOfferSession",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := UiView_newOfferSession{c, opts, UiView_newOfferSession_Params{Struct: p}, UiView_newOfferSession_Results{Struct: r}}
			return s.NewOfferSession(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// MainView_restore holds the arguments for a server call to MainView.restore.
type MainView_restore struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  MainView_restore_Params
	Results MainView_restore_Results
}

// MainView_drop holds the arguments for a server call to MainView.drop.
type MainView_drop struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  MainView_drop_Params
	Results MainView_drop_Results
}

type MainView_restore_Params struct{ capnp.Struct }

func NewMainView_restore_Params(s *capnp.Segment) (MainView_restore_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_restore_Params{}, err
	}
	return MainView_restore_Params{st}, nil
}

func NewRootMainView_restore_Params(s *capnp.Segment) (MainView_restore_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_restore_Params{}, err
	}
	return MainView_restore_Params{st}, nil
}

func ReadRootMainView_restore_Params(msg *capnp.Message) (MainView_restore_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return MainView_restore_Params{}, err
	}
	return MainView_restore_Params{root.Struct()}, nil
}

func (s MainView_restore_Params) ObjectId() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s MainView_restore_Params) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_restore_Params) ObjectIdPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s MainView_restore_Params) SetObjectId(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s MainView_restore_Params) SetObjectIdPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// MainView_restore_Params_List is a list of MainView_restore_Params.
type MainView_restore_Params_List struct{ capnp.List }

// NewMainView_restore_Params creates a new list of MainView_restore_Params.
func NewMainView_restore_Params_List(s *capnp.Segment, sz int32) (MainView_restore_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return MainView_restore_Params_List{}, err
	}
	return MainView_restore_Params_List{l}, nil
}

func (s MainView_restore_Params_List) At(i int) MainView_restore_Params {
	return MainView_restore_Params{s.List.Struct(i)}
}
func (s MainView_restore_Params_List) Set(i int, v MainView_restore_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_restore_Params_Promise is a wrapper for a MainView_restore_Params promised by a client call.
type MainView_restore_Params_Promise struct{ *capnp.Pipeline }

func (p MainView_restore_Params_Promise) Struct() (MainView_restore_Params, error) {
	s, err := p.Pipeline.Struct()
	return MainView_restore_Params{s}, err
}

func (p MainView_restore_Params_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_restore_Results struct{ capnp.Struct }

func NewMainView_restore_Results(s *capnp.Segment) (MainView_restore_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_restore_Results{}, err
	}
	return MainView_restore_Results{st}, nil
}

func NewRootMainView_restore_Results(s *capnp.Segment) (MainView_restore_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_restore_Results{}, err
	}
	return MainView_restore_Results{st}, nil
}

func ReadRootMainView_restore_Results(msg *capnp.Message) (MainView_restore_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return MainView_restore_Results{}, err
	}
	return MainView_restore_Results{root.Struct()}, nil
}

func (s MainView_restore_Results) Cap() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s MainView_restore_Results) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_restore_Results) CapPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s MainView_restore_Results) SetCap(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s MainView_restore_Results) SetCapPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// MainView_restore_Results_List is a list of MainView_restore_Results.
type MainView_restore_Results_List struct{ capnp.List }

// NewMainView_restore_Results creates a new list of MainView_restore_Results.
func NewMainView_restore_Results_List(s *capnp.Segment, sz int32) (MainView_restore_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return MainView_restore_Results_List{}, err
	}
	return MainView_restore_Results_List{l}, nil
}

func (s MainView_restore_Results_List) At(i int) MainView_restore_Results {
	return MainView_restore_Results{s.List.Struct(i)}
}
func (s MainView_restore_Results_List) Set(i int, v MainView_restore_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_restore_Results_Promise is a wrapper for a MainView_restore_Results promised by a client call.
type MainView_restore_Results_Promise struct{ *capnp.Pipeline }

func (p MainView_restore_Results_Promise) Struct() (MainView_restore_Results, error) {
	s, err := p.Pipeline.Struct()
	return MainView_restore_Results{s}, err
}

func (p MainView_restore_Results_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_drop_Params struct{ capnp.Struct }

func NewMainView_drop_Params(s *capnp.Segment) (MainView_drop_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_drop_Params{}, err
	}
	return MainView_drop_Params{st}, nil
}

func NewRootMainView_drop_Params(s *capnp.Segment) (MainView_drop_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return MainView_drop_Params{}, err
	}
	return MainView_drop_Params{st}, nil
}

func ReadRootMainView_drop_Params(msg *capnp.Message) (MainView_drop_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return MainView_drop_Params{}, err
	}
	return MainView_drop_Params{root.Struct()}, nil
}

func (s MainView_drop_Params) ObjectId() (capnp.Pointer, error) {

	return s.Struct.Pointer(0)

}

func (s MainView_drop_Params) HasObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MainView_drop_Params) ObjectIdPtr() (capnp.Ptr, error) {

	return s.Struct.Ptr(0)

}

func (s MainView_drop_Params) SetObjectId(v capnp.Pointer) error {

	return s.Struct.SetPointer(0, v)
}

func (s MainView_drop_Params) SetObjectIdPtr(v capnp.Ptr) error {

	return s.Struct.SetPtr(0, v)
}

// MainView_drop_Params_List is a list of MainView_drop_Params.
type MainView_drop_Params_List struct{ capnp.List }

// NewMainView_drop_Params creates a new list of MainView_drop_Params.
func NewMainView_drop_Params_List(s *capnp.Segment, sz int32) (MainView_drop_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return MainView_drop_Params_List{}, err
	}
	return MainView_drop_Params_List{l}, nil
}

func (s MainView_drop_Params_List) At(i int) MainView_drop_Params {
	return MainView_drop_Params{s.List.Struct(i)}
}
func (s MainView_drop_Params_List) Set(i int, v MainView_drop_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_drop_Params_Promise is a wrapper for a MainView_drop_Params promised by a client call.
type MainView_drop_Params_Promise struct{ *capnp.Pipeline }

func (p MainView_drop_Params_Promise) Struct() (MainView_drop_Params, error) {
	s, err := p.Pipeline.Struct()
	return MainView_drop_Params{s}, err
}

func (p MainView_drop_Params_Promise) ObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type MainView_drop_Results struct{ capnp.Struct }

func NewMainView_drop_Results(s *capnp.Segment) (MainView_drop_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return MainView_drop_Results{}, err
	}
	return MainView_drop_Results{st}, nil
}

func NewRootMainView_drop_Results(s *capnp.Segment) (MainView_drop_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return MainView_drop_Results{}, err
	}
	return MainView_drop_Results{st}, nil
}

func ReadRootMainView_drop_Results(msg *capnp.Message) (MainView_drop_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return MainView_drop_Results{}, err
	}
	return MainView_drop_Results{root.Struct()}, nil
}

// MainView_drop_Results_List is a list of MainView_drop_Results.
type MainView_drop_Results_List struct{ capnp.List }

// NewMainView_drop_Results creates a new list of MainView_drop_Results.
func NewMainView_drop_Results_List(s *capnp.Segment, sz int32) (MainView_drop_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return MainView_drop_Results_List{}, err
	}
	return MainView_drop_Results_List{l}, nil
}

func (s MainView_drop_Results_List) At(i int) MainView_drop_Results {
	return MainView_drop_Results{s.List.Struct(i)}
}
func (s MainView_drop_Results_List) Set(i int, v MainView_drop_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// MainView_drop_Results_Promise is a wrapper for a MainView_drop_Results promised by a client call.
type MainView_drop_Results_Promise struct{ *capnp.Pipeline }

func (p MainView_drop_Results_Promise) Struct() (MainView_drop_Results, error) {
	s, err := p.Pipeline.Struct()
	return MainView_drop_Results{s}, err
}
