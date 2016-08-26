package activity

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type ActivityEvent struct{ capnp.Struct }

func NewActivityEvent(s *capnp.Segment) (ActivityEvent, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return ActivityEvent{st}, err
}

func NewRootActivityEvent(s *capnp.Segment) (ActivityEvent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return ActivityEvent{st}, err
}

func ReadRootActivityEvent(msg *capnp.Message) (ActivityEvent, error) {
	root, err := msg.RootPtr()
	return ActivityEvent{root.Struct()}, err
}

func (s ActivityEvent) String() string {
	str, _ := text.Marshal(0xa7c0b6eaa98c6c4b, s.Struct)
	return str
}

func (s ActivityEvent) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ActivityEvent) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ActivityEvent) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ActivityEvent) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ActivityEvent) Thread() (ActivityEvent_ThreadInfo, error) {
	p, err := s.Struct.Ptr(3)
	return ActivityEvent_ThreadInfo{Struct: p.Struct()}, err
}

func (s ActivityEvent) HasThread() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s ActivityEvent) SetThread(v ActivityEvent_ThreadInfo) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewThread sets the thread field to a newly
// allocated ActivityEvent_ThreadInfo struct, preferring placement in s's segment.
func (s ActivityEvent) NewThread() (ActivityEvent_ThreadInfo, error) {
	ss, err := NewActivityEvent_ThreadInfo(s.Struct.Segment())
	if err != nil {
		return ActivityEvent_ThreadInfo{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s ActivityEvent) Notification() (NotificationDisplayInfo, error) {
	p, err := s.Struct.Ptr(1)
	return NotificationDisplayInfo{Struct: p.Struct()}, err
}

func (s ActivityEvent) HasNotification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ActivityEvent) SetNotification(v NotificationDisplayInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewNotification sets the notification field to a newly
// allocated NotificationDisplayInfo struct, preferring placement in s's segment.
func (s ActivityEvent) NewNotification() (NotificationDisplayInfo, error) {
	ss, err := NewNotificationDisplayInfo(s.Struct.Segment())
	if err != nil {
		return NotificationDisplayInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s ActivityEvent) Type() uint16 {
	return s.Struct.Uint16(0)
}

func (s ActivityEvent) SetType(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s ActivityEvent) Users() (ActivityEvent_User_List, error) {
	p, err := s.Struct.Ptr(2)
	return ActivityEvent_User_List{List: p.List()}, err
}

func (s ActivityEvent) HasUsers() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ActivityEvent) SetUsers(v ActivityEvent_User_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewUsers sets the users field to a newly
// allocated ActivityEvent_User_List, preferring placement in s's segment.
func (s ActivityEvent) NewUsers(n int32) (ActivityEvent_User_List, error) {
	l, err := NewActivityEvent_User_List(s.Struct.Segment(), n)
	if err != nil {
		return ActivityEvent_User_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// ActivityEvent_List is a list of ActivityEvent.
type ActivityEvent_List struct{ capnp.List }

// NewActivityEvent creates a new list of ActivityEvent.
func NewActivityEvent_List(s *capnp.Segment, sz int32) (ActivityEvent_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return ActivityEvent_List{l}, err
}

func (s ActivityEvent_List) At(i int) ActivityEvent { return ActivityEvent{s.List.Struct(i)} }

func (s ActivityEvent_List) Set(i int, v ActivityEvent) error { return s.List.SetStruct(i, v.Struct) }

// ActivityEvent_Promise is a wrapper for a ActivityEvent promised by a client call.
type ActivityEvent_Promise struct{ *capnp.Pipeline }

func (p ActivityEvent_Promise) Struct() (ActivityEvent, error) {
	s, err := p.Pipeline.Struct()
	return ActivityEvent{s}, err
}

func (p ActivityEvent_Promise) Thread() ActivityEvent_ThreadInfo_Promise {
	return ActivityEvent_ThreadInfo_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p ActivityEvent_Promise) Notification() NotificationDisplayInfo_Promise {
	return NotificationDisplayInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type ActivityEvent_ThreadInfo struct{ capnp.Struct }

func NewActivityEvent_ThreadInfo(s *capnp.Segment) (ActivityEvent_ThreadInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ActivityEvent_ThreadInfo{st}, err
}

func NewRootActivityEvent_ThreadInfo(s *capnp.Segment) (ActivityEvent_ThreadInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ActivityEvent_ThreadInfo{st}, err
}

func ReadRootActivityEvent_ThreadInfo(msg *capnp.Message) (ActivityEvent_ThreadInfo, error) {
	root, err := msg.RootPtr()
	return ActivityEvent_ThreadInfo{root.Struct()}, err
}

func (s ActivityEvent_ThreadInfo) String() string {
	str, _ := text.Marshal(0xf3902b5a86ffba44, s.Struct)
	return str
}

func (s ActivityEvent_ThreadInfo) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ActivityEvent_ThreadInfo) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ActivityEvent_ThreadInfo) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ActivityEvent_ThreadInfo) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ActivityEvent_ThreadInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s ActivityEvent_ThreadInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ActivityEvent_ThreadInfo) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ActivityEvent_ThreadInfo) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// ActivityEvent_ThreadInfo_List is a list of ActivityEvent_ThreadInfo.
type ActivityEvent_ThreadInfo_List struct{ capnp.List }

// NewActivityEvent_ThreadInfo creates a new list of ActivityEvent_ThreadInfo.
func NewActivityEvent_ThreadInfo_List(s *capnp.Segment, sz int32) (ActivityEvent_ThreadInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return ActivityEvent_ThreadInfo_List{l}, err
}

func (s ActivityEvent_ThreadInfo_List) At(i int) ActivityEvent_ThreadInfo {
	return ActivityEvent_ThreadInfo{s.List.Struct(i)}
}

func (s ActivityEvent_ThreadInfo_List) Set(i int, v ActivityEvent_ThreadInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// ActivityEvent_ThreadInfo_Promise is a wrapper for a ActivityEvent_ThreadInfo promised by a client call.
type ActivityEvent_ThreadInfo_Promise struct{ *capnp.Pipeline }

func (p ActivityEvent_ThreadInfo_Promise) Struct() (ActivityEvent_ThreadInfo, error) {
	s, err := p.Pipeline.Struct()
	return ActivityEvent_ThreadInfo{s}, err
}

func (p ActivityEvent_ThreadInfo_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type ActivityEvent_User struct{ capnp.Struct }

func NewActivityEvent_User(s *capnp.Segment) (ActivityEvent_User, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ActivityEvent_User{st}, err
}

func NewRootActivityEvent_User(s *capnp.Segment) (ActivityEvent_User, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ActivityEvent_User{st}, err
}

func ReadRootActivityEvent_User(msg *capnp.Message) (ActivityEvent_User, error) {
	root, err := msg.RootPtr()
	return ActivityEvent_User{root.Struct()}, err
}

func (s ActivityEvent_User) String() string {
	str, _ := text.Marshal(0xbb4d119bbcd89677, s.Struct)
	return str
}

func (s ActivityEvent_User) Identity() identity.Identity {
	p, _ := s.Struct.Ptr(0)
	return identity.Identity{Client: p.Interface().Client()}
}

func (s ActivityEvent_User) HasIdentity() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ActivityEvent_User) SetIdentity(v identity.Identity) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s ActivityEvent_User) Mentioned() bool {
	return s.Struct.Bit(0)
}

func (s ActivityEvent_User) SetMentioned(v bool) {
	s.Struct.SetBit(0, v)
}

func (s ActivityEvent_User) Subscribed() bool {
	return s.Struct.Bit(2)
}

func (s ActivityEvent_User) SetSubscribed(v bool) {
	s.Struct.SetBit(2, v)
}

func (s ActivityEvent_User) CanView() bool {
	return s.Struct.Bit(1)
}

func (s ActivityEvent_User) SetCanView(v bool) {
	s.Struct.SetBit(1, v)
}

// ActivityEvent_User_List is a list of ActivityEvent_User.
type ActivityEvent_User_List struct{ capnp.List }

// NewActivityEvent_User creates a new list of ActivityEvent_User.
func NewActivityEvent_User_List(s *capnp.Segment, sz int32) (ActivityEvent_User_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return ActivityEvent_User_List{l}, err
}

func (s ActivityEvent_User_List) At(i int) ActivityEvent_User {
	return ActivityEvent_User{s.List.Struct(i)}
}

func (s ActivityEvent_User_List) Set(i int, v ActivityEvent_User) error {
	return s.List.SetStruct(i, v.Struct)
}

// ActivityEvent_User_Promise is a wrapper for a ActivityEvent_User promised by a client call.
type ActivityEvent_User_Promise struct{ *capnp.Pipeline }

func (p ActivityEvent_User_Promise) Struct() (ActivityEvent_User, error) {
	s, err := p.Pipeline.Struct()
	return ActivityEvent_User{s}, err
}

func (p ActivityEvent_User_Promise) Identity() identity.Identity {
	return identity.Identity{Client: p.Pipeline.GetPipeline(0).Client()}
}

type ActivityTypeDef struct{ capnp.Struct }
type ActivityTypeDef_requiredPermission ActivityTypeDef
type ActivityTypeDef_requiredPermission_Which uint16

const (
	ActivityTypeDef_requiredPermission_Which_everyone        ActivityTypeDef_requiredPermission_Which = 0
	ActivityTypeDef_requiredPermission_Which_permissionIndex ActivityTypeDef_requiredPermission_Which = 1
	ActivityTypeDef_requiredPermission_Which_explicitList    ActivityTypeDef_requiredPermission_Which = 2
)

func (w ActivityTypeDef_requiredPermission_Which) String() string {
	const s = "everyonepermissionIndexexplicitList"
	switch w {
	case ActivityTypeDef_requiredPermission_Which_everyone:
		return s[0:8]
	case ActivityTypeDef_requiredPermission_Which_permissionIndex:
		return s[8:23]
	case ActivityTypeDef_requiredPermission_Which_explicitList:
		return s[23:35]

	}
	return "ActivityTypeDef_requiredPermission_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewActivityTypeDef(s *capnp.Segment) (ActivityTypeDef, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return ActivityTypeDef{st}, err
}

func NewRootActivityTypeDef(s *capnp.Segment) (ActivityTypeDef, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return ActivityTypeDef{st}, err
}

func ReadRootActivityTypeDef(msg *capnp.Message) (ActivityTypeDef, error) {
	root, err := msg.RootPtr()
	return ActivityTypeDef{root.Struct()}, err
}

func (s ActivityTypeDef) String() string {
	str, _ := text.Marshal(0xe638de0ad7c89a2b, s.Struct)
	return str
}

func (s ActivityTypeDef) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ActivityTypeDef) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ActivityTypeDef) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ActivityTypeDef) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ActivityTypeDef) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s ActivityTypeDef) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ActivityTypeDef) SetVerbPhrase(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ActivityTypeDef) NewVerbPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s ActivityTypeDef) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s ActivityTypeDef) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ActivityTypeDef) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ActivityTypeDef) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s ActivityTypeDef) RequiredPermission() ActivityTypeDef_requiredPermission {
	return ActivityTypeDef_requiredPermission(s)
}

func (s ActivityTypeDef_requiredPermission) Which() ActivityTypeDef_requiredPermission_Which {
	return ActivityTypeDef_requiredPermission_Which(s.Struct.Uint16(0))
}
func (s ActivityTypeDef_requiredPermission) SetEveryone() {
	s.Struct.SetUint16(0, 0)

}

func (s ActivityTypeDef_requiredPermission) PermissionIndex() uint16 {
	return s.Struct.Uint16(2)
}

func (s ActivityTypeDef_requiredPermission) SetPermissionIndex(v uint16) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint16(2, v)
}

func (s ActivityTypeDef_requiredPermission) SetExplicitList() {
	s.Struct.SetUint16(0, 2)

}

func (s ActivityTypeDef) Obsolete() bool {
	return s.Struct.Bit(32)
}

func (s ActivityTypeDef) SetObsolete(v bool) {
	s.Struct.SetBit(32, v)
}

func (s ActivityTypeDef) NotifySubscribers() bool {
	return s.Struct.Bit(33)
}

func (s ActivityTypeDef) SetNotifySubscribers(v bool) {
	s.Struct.SetBit(33, v)
}

func (s ActivityTypeDef) AutoSubscribeToThread() bool {
	return s.Struct.Bit(34)
}

func (s ActivityTypeDef) SetAutoSubscribeToThread(v bool) {
	s.Struct.SetBit(34, v)
}

func (s ActivityTypeDef) AutoSubscribeToGrain() bool {
	return s.Struct.Bit(35)
}

func (s ActivityTypeDef) SetAutoSubscribeToGrain(v bool) {
	s.Struct.SetBit(35, v)
}

func (s ActivityTypeDef) SuppressUnread() bool {
	return s.Struct.Bit(36)
}

func (s ActivityTypeDef) SetSuppressUnread(v bool) {
	s.Struct.SetBit(36, v)
}

// ActivityTypeDef_List is a list of ActivityTypeDef.
type ActivityTypeDef_List struct{ capnp.List }

// NewActivityTypeDef creates a new list of ActivityTypeDef.
func NewActivityTypeDef_List(s *capnp.Segment, sz int32) (ActivityTypeDef_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return ActivityTypeDef_List{l}, err
}

func (s ActivityTypeDef_List) At(i int) ActivityTypeDef { return ActivityTypeDef{s.List.Struct(i)} }

func (s ActivityTypeDef_List) Set(i int, v ActivityTypeDef) error {
	return s.List.SetStruct(i, v.Struct)
}

// ActivityTypeDef_Promise is a wrapper for a ActivityTypeDef promised by a client call.
type ActivityTypeDef_Promise struct{ *capnp.Pipeline }

func (p ActivityTypeDef_Promise) Struct() (ActivityTypeDef, error) {
	s, err := p.Pipeline.Struct()
	return ActivityTypeDef{s}, err
}

func (p ActivityTypeDef_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p ActivityTypeDef_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p ActivityTypeDef_Promise) RequiredPermission() ActivityTypeDef_requiredPermission_Promise {
	return ActivityTypeDef_requiredPermission_Promise{p.Pipeline}
}

// ActivityTypeDef_requiredPermission_Promise is a wrapper for a ActivityTypeDef_requiredPermission promised by a client call.
type ActivityTypeDef_requiredPermission_Promise struct{ *capnp.Pipeline }

func (p ActivityTypeDef_requiredPermission_Promise) Struct() (ActivityTypeDef_requiredPermission, error) {
	s, err := p.Pipeline.Struct()
	return ActivityTypeDef_requiredPermission{s}, err
}

type NotificationDisplayInfo struct{ capnp.Struct }

func NewNotificationDisplayInfo(s *capnp.Segment) (NotificationDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return NotificationDisplayInfo{st}, err
}

func NewRootNotificationDisplayInfo(s *capnp.Segment) (NotificationDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return NotificationDisplayInfo{st}, err
}

func ReadRootNotificationDisplayInfo(msg *capnp.Message) (NotificationDisplayInfo, error) {
	root, err := msg.RootPtr()
	return NotificationDisplayInfo{root.Struct()}, err
}

func (s NotificationDisplayInfo) String() string {
	str, _ := text.Marshal(0xd3b9f2ca42d4f783, s.Struct)
	return str
}

func (s NotificationDisplayInfo) Caption() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
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
	return NotificationDisplayInfo_List{l}, err
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
			InterfaceName: "activity.capnp:NotificationTarget",
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
			InterfaceName: "activity.capnp:NotificationTarget",
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
	return NotificationTarget_addOngoing_Params{st}, err
}

func NewRootNotificationTarget_addOngoing_Params(s *capnp.Segment) (NotificationTarget_addOngoing_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return NotificationTarget_addOngoing_Params{st}, err
}

func ReadRootNotificationTarget_addOngoing_Params(msg *capnp.Message) (NotificationTarget_addOngoing_Params, error) {
	root, err := msg.RootPtr()
	return NotificationTarget_addOngoing_Params{root.Struct()}, err
}

func (s NotificationTarget_addOngoing_Params) String() string {
	str, _ := text.Marshal(0xdaaca64376c9033a, s.Struct)
	return str
}

func (s NotificationTarget_addOngoing_Params) DisplayInfo() (NotificationDisplayInfo, error) {
	p, err := s.Struct.Ptr(0)
	return NotificationDisplayInfo{Struct: p.Struct()}, err
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
	p, _ := s.Struct.Ptr(1)
	return OngoingNotification{Client: p.Interface().Client()}
}

func (s NotificationTarget_addOngoing_Params) HasNotification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s NotificationTarget_addOngoing_Params) SetNotification(v OngoingNotification) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// NotificationTarget_addOngoing_Params_List is a list of NotificationTarget_addOngoing_Params.
type NotificationTarget_addOngoing_Params_List struct{ capnp.List }

// NewNotificationTarget_addOngoing_Params creates a new list of NotificationTarget_addOngoing_Params.
func NewNotificationTarget_addOngoing_Params_List(s *capnp.Segment, sz int32) (NotificationTarget_addOngoing_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return NotificationTarget_addOngoing_Params_List{l}, err
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
	return NotificationTarget_addOngoing_Results{st}, err
}

func NewRootNotificationTarget_addOngoing_Results(s *capnp.Segment) (NotificationTarget_addOngoing_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return NotificationTarget_addOngoing_Results{st}, err
}

func ReadRootNotificationTarget_addOngoing_Results(msg *capnp.Message) (NotificationTarget_addOngoing_Results, error) {
	root, err := msg.RootPtr()
	return NotificationTarget_addOngoing_Results{root.Struct()}, err
}

func (s NotificationTarget_addOngoing_Results) String() string {
	str, _ := text.Marshal(0xee4cbafe4028d3c2, s.Struct)
	return str
}

func (s NotificationTarget_addOngoing_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s NotificationTarget_addOngoing_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NotificationTarget_addOngoing_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// NotificationTarget_addOngoing_Results_List is a list of NotificationTarget_addOngoing_Results.
type NotificationTarget_addOngoing_Results_List struct{ capnp.List }

// NewNotificationTarget_addOngoing_Results creates a new list of NotificationTarget_addOngoing_Results.
func NewNotificationTarget_addOngoing_Results_List(s *capnp.Segment, sz int32) (NotificationTarget_addOngoing_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return NotificationTarget_addOngoing_Results_List{l}, err
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
			InterfaceName: "activity.capnp:OngoingNotification",
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
			InterfaceName: "activity.capnp:OngoingNotification",
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
	return OngoingNotification_cancel_Params{st}, err
}

func NewRootOngoingNotification_cancel_Params(s *capnp.Segment) (OngoingNotification_cancel_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return OngoingNotification_cancel_Params{st}, err
}

func ReadRootOngoingNotification_cancel_Params(msg *capnp.Message) (OngoingNotification_cancel_Params, error) {
	root, err := msg.RootPtr()
	return OngoingNotification_cancel_Params{root.Struct()}, err
}

func (s OngoingNotification_cancel_Params) String() string {
	str, _ := text.Marshal(0xf1d1cb82a830948b, s.Struct)
	return str
}

// OngoingNotification_cancel_Params_List is a list of OngoingNotification_cancel_Params.
type OngoingNotification_cancel_Params_List struct{ capnp.List }

// NewOngoingNotification_cancel_Params creates a new list of OngoingNotification_cancel_Params.
func NewOngoingNotification_cancel_Params_List(s *capnp.Segment, sz int32) (OngoingNotification_cancel_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return OngoingNotification_cancel_Params_List{l}, err
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
	return OngoingNotification_cancel_Results{st}, err
}

func NewRootOngoingNotification_cancel_Results(s *capnp.Segment) (OngoingNotification_cancel_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return OngoingNotification_cancel_Results{st}, err
}

func ReadRootOngoingNotification_cancel_Results(msg *capnp.Message) (OngoingNotification_cancel_Results, error) {
	root, err := msg.RootPtr()
	return OngoingNotification_cancel_Results{root.Struct()}, err
}

func (s OngoingNotification_cancel_Results) String() string {
	str, _ := text.Marshal(0xf81453d4b819e209, s.Struct)
	return str
}

// OngoingNotification_cancel_Results_List is a list of OngoingNotification_cancel_Results.
type OngoingNotification_cancel_Results_List struct{ capnp.List }

// NewOngoingNotification_cancel_Results creates a new list of OngoingNotification_cancel_Results.
func NewOngoingNotification_cancel_Results_List(s *capnp.Segment, sz int32) (OngoingNotification_cancel_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return OngoingNotification_cancel_Results_List{l}, err
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

const schema_a4e001d4cbcf33fa = "x\xda\x94UohUe\x18\x7f\x9fs\xce\xbdgg" +
	"\xcc\xae\xa7s\x84\x8c\xc6LG\x99\xe2\x98.0Fp" +
	"\xd7\xda\x183\xad]w\x171\x10vv\xef\xbb\xed\xc4" +
	"\xdd9\xd7s\xce\x9d\x9b\x11\x99\xad\xc840\xa9P\xab" +
	"\x0f\x85\x94\x8a\x11\xf4\xa1\xcc\x12\xa4O\x99\x84\x84\x82Q" +
	"\x18\xe6\x87\xfe\x08\x86K!\x15\xe7\xe9\xf7\xde\xff\xeej" +
	"\xd4\x87\xc1\xee\xef}\xde\xe7\xf9=\xbf\xf7\xf7<\xa7\xf5" +
	"\x01\xb9CZ\x19\xe9\xa9c,1\x14\x89\x86'\x87." +
	"o\xe8v\xde\xda\xc3\x12MD\xe1\xf2\xbd\xdf\x9c\xa9\xff" +
	"\xf9\x91_\xd9\x80\xac\x92\xccX\x9b\xa6\xec\"F\xc6B" +
	"\xe5\x13F\xe1\x13\x99\x1d\x07.|v\xec#\x96\x88!" +
	"\xf6z\xdb\xc9\x13\xa7\xe9\xdc>\x16QT\xc6\x8c\xa3\xca" +
	"\x05\xe3\xb8\xf2 \xee\xfc\xa24\xe1N\xb8\xe9\xed\x1f\xbe" +
	"zG_\xf7%K4R\xd5\xdd\x08!\xba-\x12\xbd" +
	"\x9b\x8c\x05QqQ\x8f\xfe\x86\xe8\x97\xfe>\xdd\xf9\xed" +
	"__\x9cbz\xac:u>x&\xdaN\x06\xa9\"" +
	"x6\x1aGp\xbb||\xe2\xf1\x0f\x0f\xfd\xc8\xf4\xc5" +
	"\xc4XD\x12A\x8d\xea\x07\x82\xe9\x0au\x13\xab\xeac" +
	"\x0eSY\xe4xC\xbdn\xbc'\xb2\xb5\xedQC\xc1" +
	"\xf4\xebSK;n\x1eY\xfbg1]\xbe\xe6\xc7\xda" +
	"A\x91\xee\xa8&\xea]Z\xd4zf\xb5\x7f\xf5\x12\xc8" +
	"\xc9\x95l\x8c\xda\xcejh\xe3\xa2&\xb2\xfe\xa1\xf5\x18" +
	"\x0b\xea\xf1_\xb8\xfd\xcd\xd6\xfd[O|?\xc3\xf4%" +
	"H'\xc4i\x9b\xd5^#\xa6\x84]G\xc2W\x06\x97" +
	"\xef\xbc\xcc\xf4\xc6jE\xf2\xfc/j\x9dd\xcc\xe6S" +
	"]\xd3D\x0f\xda\xf9\x85\x87O\xf7\x9bW\xab\xd2\x0c\xd4" +
	"\xef\x12i\xbe\xeb\xd8v\xf8\xa7\xc6\x97o\xd6\xd0\xe9\xae" +
	"\xbf\x97\x8c\x01A\xc2H\xd4\xf7\x18S\xf8o:\xb4R" +
	"\x81=a\x07S\x91\x96\x94\x95u\xb2\xed\x8f\x15\x7f'" +
	"\xa7\xb2\xbc\x8b\x8f\xb4x|c\xce\xf6x\xba\x8f{\xe3" +
	"\xb6\xef\xdb.9\x89\x06Yi\x08CSX@\xef^" +
	"\x03\xabt\xc9\x94\xe8\x93\xa8\x91n\x86d\x92\x02x\xdd" +
	"V\xc0k\x01?#\xd1<i\x16\xc1\x11\xa0\x03\xcf\x02" +
	"M\x02\x1d\x92(\xe4\x13\xdc\x9br\x1d\xce\x98\x1af\x8b" +
	"\xd9\xc9uz\x9d4\x9fdPY\xc2\x1f\xa2&\xb3\x19" +
	";e\x07,\xb6\xd6\xf6\x03\x16-\x13\x96\xe6\x10\xee\x8e" +
	"Mp'H\xd4\xe1E+Bj\x83\x15\xa3\xe9\xda\xb2" +
	"09\xe6q+\xdd\xeb0y\xc4\x8d\x0d\xf8\xdcK\x98" +
	"2\xe8*\x90P\x7f~\x19\xc8M\x82\xdc\xb4D\xba\x84" +
	">\x04\xf8\xa2`\xbc\x05\xe0\x0e\x89H6I\x02\xb6M" +
	"\x04N\x03\xdb\x89@E*\xe8\xf0\xfa*\x80\xaf\x02\xdc" +
	"\x07PD\x0a\x15\xdeo\x07\xf8.\xc0\xfd\x12\xc5\xb2V" +
	"0F\x0dh\xab\x01m9n`\x8f\xd8)\x8b\xc5\x02" +
	"\xdbuh~\xc5\xe1\x8ch>\xa3X\x00\xfdK\"4" +
	"\xe5\xc0\xd4\xa7\xbb\x18\xf5\xc98\xad\xb4\x84X\x80\xf1 " +
	"\xdf\x15\x0e\xca\x8d\x17\x92\x94\xb5\x92\xe7j%\xa4jA" +
	"\xff\xe4\xf5\x11%\xe6\x975\xb0\xc4k\x0e\x81p\x06\xed" +
	"\x16%\xb0\xd7\x03\x1b\x03\x16\xa03\x99\x0a\x1al\xec\x04" +
	"\x98\x018)\xc4*j\x90\x1b\x04\x18\x00\xdc\x82\xe7\xb5" +
	"\xd3\xa8\x81Z\x0cd\xf4\xf0\x9e\xb3\x89\x03\x9bwO\x1f" +
	"\x13\xcct0\x1b\x17\x87\xe2\xf1)\x8d\"(\xc6\xe8\x85" +
	"\x94\xe5<m\xf3M\xa5\xdf\xa1\x9f\x1b\xf6S\x9e=\xcc" +
	"d^\x0e\xaai\xe9\xc9\xa2\x90\"[\x97\xedg3q" +
	"k\xaa\xd7\x19qE_J\xb9\xafy\x82n\x1d\x98\x99" +
	"\x92\xa8\x93-\x8a\xde{\xfe\xb9\xd5\xde\xe7\x1b\xb6\xcf\xd5" +
	"+r\x9b\xe4I\xcb\x1b\xe5A\x8b\x95N?\xe5\x8c\xba" +
	"\xb63\xda\xdcgy\xaa5\xee#o\xa9\xccC\xc3(" +
	"\xb3\x14e\x1e\x16&(\xea\xb7RX\xa8\x15\xe0\xa3P" +
	"%-(\x82!S\xc1\xb1\xf6\xd9\xe7\xfaB\xafLs" +
	"Q\xb8;\xd9?\x19/\x0c\xac\xe8\xbb\xb9L\xe8\xa2\xb0" +
	"\xea\xef\xa8}\xa5\x8a\xd0\x8cx\xa6K\x00o\x14\xdfN" +
	"<\xe85A\xfd\xaaL\xfd\x0a\x09\xa7c\x8a\xca\xbb\xdf" +
	" \xda\xcb$]YdR\x14$f\xd6T]\x8f\xdc" +
	"o\x8a\x85\xa8_\xdb\x05\xf0\x06\xae\xd7\xe1\xba\x1e]l" +
	"\x12>!F\x84>e\x0c\x10pS\xe0\xea\x12\x934" +
	"\xb1\xd4\xe9 pS\xe0\x8b\x04^\xd7lR=\xf0F" +
	"\xda\x0c\xfc>\x81/\x05\x1es\xacq^\x1e\x19\xac\x8b" +
	"\xe1\xbe1\xcfb2|[\xfbvi.\xcc\x92\x0d\x98" +
	"z\xfb\xb7--2*m2\xd9uBw\xd8w3" +
	"<\xe0\xc2\xa4%\x87\xe5\x9f`\xaa?G\x05\xf7q\x8f" +
	"\xfc\x8a\xfbr\x81\xdb\x0f[R\xfe$\xe9&\x9b\x0a\xa3" +
	"w\xa7\xf3\x9e\x98g\xd9N\x95\xa3\xb3Y\x8f\xfb>\x8b" +
	"\x0f8\xb7\xde\xfb_\xc6[\xcf\xfdX.\x13\xf8\xd5\x06" +
	"o\xaf\x18<>f9\xe9\x0c\x87y\xceu\x0e\x0d\x1d" +
	"j\xbe\xb2{\xaey\xe4;\x95\xa1\xa007X\xd6\xe5" +
	"\xcf(\x95>\x80\xba>\x08\x1bhjX\xa2\xc2dg" +
	"\xb4\x83p\xa3\x86\x7f\x91iu~\x9c8)\x9e\x11S" +
	"\x83\xa1A\xfe\xd2\x1d\xe5\xb6\xcb\xa9\xb4\xa9G\\&(" +
	"U\xcd\x98\xb0t3:m\xad\xb2\xf4\x8aU\x95\xc1\xbb" +
	"e\xd16a\x05e\xf8\x7f\x18\xf5\x7fa\x0c\xb9\xa16" +
	"\xf95\xf2\xd5\xdc\x81\xa5*\xfa\x95>\xf4T\xfaT\xeb" +
	"z;\xf4\x8b\xa8\xf1B\xde\xbcr\xff\x04\x00\x00\xff\xff" +
	":F\xaa]"

func init() {
	schemas.Register(schema_a4e001d4cbcf33fa,
		0x99956e455df360cf,
		0xa7c0b6eaa98c6c4b,
		0xbb4d119bbcd89677,
		0xd3b9f2ca42d4f783,
		0xdaaca64376c9033a,
		0xe638de0ad7c89a2b,
		0xee4cbafe4028d3c2,
		0xf0f87337d73020f0,
		0xf1d1cb82a830948b,
		0xf3902b5a86ffba44,
		0xf81453d4b819e209,
		0xfe851ddbb88940cd)
}
