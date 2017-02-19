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

// ActivityEvent_TypeID is the unique identifier for the type ActivityEvent.
const ActivityEvent_TypeID = 0xa7c0b6eaa98c6c4b

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

// ActivityEvent_ThreadInfo_TypeID is the unique identifier for the type ActivityEvent_ThreadInfo.
const ActivityEvent_ThreadInfo_TypeID = 0xf3902b5a86ffba44

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

// ActivityEvent_User_TypeID is the unique identifier for the type ActivityEvent_User.
const ActivityEvent_User_TypeID = 0xbb4d119bbcd89677

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

// ActivityTypeDef_TypeID is the unique identifier for the type ActivityTypeDef.
const ActivityTypeDef_TypeID = 0xe638de0ad7c89a2b

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

// NotificationDisplayInfo_TypeID is the unique identifier for the type NotificationDisplayInfo.
const NotificationDisplayInfo_TypeID = 0xd3b9f2ca42d4f783

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

// NotificationTarget_addOngoing_Params_TypeID is the unique identifier for the type NotificationTarget_addOngoing_Params.
const NotificationTarget_addOngoing_Params_TypeID = 0xdaaca64376c9033a

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

// NotificationTarget_addOngoing_Results_TypeID is the unique identifier for the type NotificationTarget_addOngoing_Results.
const NotificationTarget_addOngoing_Results_TypeID = 0xee4cbafe4028d3c2

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

// OngoingNotification_cancel_Params_TypeID is the unique identifier for the type OngoingNotification_cancel_Params.
const OngoingNotification_cancel_Params_TypeID = 0xf1d1cb82a830948b

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

// OngoingNotification_cancel_Results_TypeID is the unique identifier for the type OngoingNotification_cancel_Results.
const OngoingNotification_cancel_Results_TypeID = 0xf81453d4b819e209

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

const schema_a4e001d4cbcf33fa = "x\xda\x94U]l\x14U\x14>\xdf\xdc\xdd\xce\x96-" +
	"n/;\x0fjl\x8a\xd0D~bS\xa8\x09\xa6\xd1" +
	"\xb4\xd4\x12R\xa4\xda\xcbn\x8d6i\xd2\xe9\xeem;" +
	"f;\xb3\xcc\xcc\x96.\x06ED#\x02\x09!j\x10" +
	"}\xd1\xa0\x02!1\x9a\xa0\x88>\x10\x9fTb\x88\x81" +
	"\x04\xffCxPC\xa2\xa1B\"\x10`\xcc\xdd\x7f\xda" +
	"b\xf4m\xf7\xce9\xdf\xfd\xcew\xbesn\xdb\xb3\xac" +
	"K[\x11\xd6\xa3Dbk\xb8.85|ih\x8d" +
	"\xfd\xda\x1b$\x9a\x81`\xf9\xfe/\xcf\xce\xfb\xe5\xc1_" +
	"i\x80\xe9`D\xedVx/\x08\xf1|\xf8\x03B\xf0" +
	"hf\xd7\xa1\x0b\x1f\x9fx\x9fD\x0c\x08\xae\xb5\x9f:" +
	"y\x06\xe7\x0eP8\xa4\x13\xc5\x9b\xea.\xc4\x97\xd6\xdd" +
	"G\xd4\xbe\xba\xae\x19\x84`\xd3\xeb\xdf}\xfe&\xef\xfb" +
	"\x8cD\x13jr\xc3\xd0\x89\xda\xa5\xbe\x00\xf1\x9c\xae\x12" +
	"7\xea\xbf\x11\x82\x17\xfe>\xd3\xfd\xf5_\x9f\x9e&\x1e" +
	"\xab\x85.\x04\x8bH\x07\xe2fD\x05\x0fE:\x09A" +
	"\x07\xfbj\xf2\x91\xf7\x8e\xfc@|\x11\x88\xc2\x9a\x0a\xda" +
	"\x12yG1\xdd\x1d\xd9D5u\xcc`\xca\x14\xc6t" +
	"\xe4Z\xfc\x86Bk\xbf\x1a\x09\x14\xd3/N/\xe9\xba" +
	"y|\xfd\x9f%\xb8\xc2\x9d\xf3\xa3\x87\x15\\ST\xdd" +
	"wqa\xdb\xd9U\xde\x95\x8b\xc4c\xac\x8aFh\x7f" +
	"8\xba\x00\xf1\xbe\xa8B\xed\x8d\xae\x8d\xe7\xd4\xaf`\xe7" +
	"\xabm\x07\xb7\x9d\xfcv\x9a\xf8b\x10)q\xda\x87\xa2" +
	"\xaf\x80BA\xcf\xf1\xe0\xa5\xc1\xe5{.\x11o\xaaU" +
	"\xa4\xc0\xbf/\xda\x8d\xf8P\x01\xea\xa9\xa8\xaa\xa1\xfe\xfc" +
	"]\xc7\xce$\x8c+50G\xa3{\x15\xcc7];" +
	"\x8e\xfd\xd8\xf4\xe2\xcdYt\xde\x8d\xde\x8d\xf8\xd1\x02\xc6" +
	"\x87\xd1\xb5\xf1\xef\xa3:\xfd\x14\x98)\xdf\x9a\xb4\xfc|" +
	"\xb85ef\xedl\xc7\xea\xd2\xffd>+{\xe4h" +
	"\xab+7\xe6,W\xa6\xfb\xa5;ay\x9e\xe5\xc0\x16" +
	"\x0d,\xd4\x10\x04\x86\xb2\x00_\xb3\x8eH\xf40\x88~" +
	"\x0dM\xb8\x19\xc0@\x88\x88\xf7m#\x12\xeb\x19\xc4\x93" +
	"\x1a\xe6k7\x02\x03a\">\xf04\x91H2\x88a" +
	"\x0d\x81\x9c\x94n\xde\xb1%\x91\x1edK\xe8p\xec^" +
	";-\xa7\x08:i\xd0\x09\x81\x9c\xcaf\xac\x94\xe5S" +
	"l\xbd\xe5\xf9TW!\xac\xcd \xbc&6)m_" +
	"D\x80\x1a!\xeb\x07\xabF\xe3\xf5\xcb\x82\xe4\xb8+\xcd" +
	"t\xafMl\xd4\x89\x0dx\xd2\x15\x06\x0b\x11\x85@\xc4" +
	"\xb7,#\x12S\x0cb\xbb\x06\xae\xc1\x80:|^1" +
	"\xde\xca vi\x003\xa0\x11\xf1\x1d*p;\x83\xd8" +
	"\xa3\x81\x87\xb4\xa2\x0e\xbbW\x12\x89\x97\x19\xc4\x01\x0d\\" +
	"E*\x15\xde\xee \x12o1\x88\x83\x1abY\xd3\x1f" +
	"G\x03ih \x04\xb6\xe3[\xa3V\xca\xa4\x98o9" +
	"6\x1a\xab\x0e'\xa0\x91\x10\xf3\xf3YY\x16\xa19\xe7" +
	"I\xd7\xc3\x1d\x84~\x064VK\"\xa8\xc3N\xbfP" +
	"\x15\x1a\xab\x85\x17A*Z\xb1\x99Z)\xa9Z\x07<" +
	"\x09\xb7\x1f\x10\x8d\x15\x0dL\xd5\xcda\x06\x91\xd1\x80\x92" +
	"\x04\xd6\x06\"1\xce |\x0d\x9c\xa1\xa8\xc1\xc6n\"" +
	"\x91a\x10SJ\xac\x92\x06\xb9A\"\xe13\x88\xad\x1a" +
	"\x02+-m\xdf\xf2\xf3D\x04\x1e\xdc\xf9\xb38\xb4y" +
	"\xdf\xf6\x13\x8a\x19'\x04\x13\xea\xa3j>\xd2\x00i\x00" +
	"\xe1\xb9\x94i?a\xc9M\xe5\xff\x81\x97\x1b\xf1R\xae" +
	"5BLV\x82f\x95\xf4XIH\x85\xd6cy\xd9" +
	"L\xa7\x99\xef\xb5G\x1dUW\xa8R\xd7|E7\xc2" +
	" \x0cM\xdd\x93-\x89\xde{\xfe\x99U\xee'C;" +
	"g\xea\x15\x9e\x03<i\xbac\xd2o5\xd3\xe9\xc7\xed" +
	"1\xc7\xb2\xc7Z\xfaMW7'<\x11\xa9\\\xb3t" +
	"\x84H,a\x10\x0f(\x13\x94\xf4[\xa1,\xd4\xc6 " +
	"\x1e\xd2\x10\xa4\x15E3\xdfK\xba=\xea\xccn\xfbL" +
	"_\xf0\xea4\x97\x84\xbb\x9d\xfd\x93\x9d\xc5\x81Uu\xb7" +
	"T\x08\xfd\xa1\xac\xfa;\x83\xb8\\ChZ\xb5\xe9\"" +
	"\x83\xb8^\xea\x9dj\xe8UE\xfd\x0aC\"\x04\xe5t" +
	"\xa0\xba\xfb\xe3\xc0~\xd2xh\xa1\x81:\x80O\xaf\xab" +
	"I\x0f\xdfk\xa8\x85\xc8\xaf\xee%\x12\xd7\x19\x12\x11h" +
	"\xe0u\x8b\x0cD\x88\xe2a|D\x94\x88\x80!a\xa8" +
	"s}\xb1\x81z\xa28\xc7a\xa2\x84\xa1\xce\x17\xaa\xf3" +
	"H\x8b\x81y\xea\x95\xc0f\xa2\xc4=\xea|\x094\xc4" +
	"lsBVFfR\xba#\xfd\xe3\xaeI\xcc\x93s" +
	"\xf4.-\x95Y\xb2>\xe9s\xf7\xb6\xbc\xc8P\xded" +
	"\xcc\xb1\x03g\xc4s2\xd2\x97\xca\xa4e\x87\x15Z\x90" +
	"O\xe4Pt\x9ft\xe1U\xdd\x97\xf3\x9dDn\xc4C" +
	"\xe1K\xd2I6\x17G\xefv\xdf\xd7\xc6\\\xd3\xb2k" +
	"\x1c\x9d\xcd\xba\xd2\xf3\xa8s\xc0\xbe5\xef\x7f\x19o\x83" +
	"\xf4b\xb9\x8c\xef\xd5\x1a\xbc\xa3j\xf0\xceq\xd3Ng" +
	"$xp\xae{x\xf8H\xcb\xe5}3\xcd\xc3nw" +
	"\x0d\xfc\xe2\xdc\x84\x89*\xcf(\xca\x0f \xe7\x83\xa4\xf1" +
	"z=(S!f\x8fu\xa1\x1f\xb3\xf9\x97\x98\xd6\xe2" +
	"\xb7\xa6L;%3jj\xcc\x09\x8f\xa8\x92\x13\x9as" +
	"9\x957\xf5\xa8C\x8aR\xcd\x8c)K\xb70\x88\xb6" +
	"\x1aK\xdf\xbf\xb2:x\xb7,\xdaf\xdf\xf23\xf2?" +
	"\x8c\xfa\xbf0\xde \xbd\\\xc6\x877K\xbeY9\xcc" +
	"\xb1\xab\xfa\x95\x1fz\x94\x9fj\xce;H\xe3a\xbd\xb3" +
	"\x88[P\xee\x9f\x00\x00\x00\xff\xff\xab\x1b\xa8\xbd"

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
