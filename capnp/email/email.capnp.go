package email

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type EmailAddress struct{ capnp.Struct }

func NewEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailAddress{st}, err
}

func NewRootEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailAddress{st}, err
}

func ReadRootEmailAddress(msg *capnp.Message) (EmailAddress, error) {
	root, err := msg.RootPtr()
	return EmailAddress{root.Struct()}, err
}

func (s EmailAddress) String() string {
	str, _ := text.Marshal(0xacaddcee86563ee1, s.Struct)
	return str
}

func (s EmailAddress) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailAddress) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAddress) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailAddress) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s EmailAddress) Name() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailAddress) HasName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAddress) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailAddress) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// EmailAddress_List is a list of EmailAddress.
type EmailAddress_List struct{ capnp.List }

// NewEmailAddress creates a new list of EmailAddress.
func NewEmailAddress_List(s *capnp.Segment, sz int32) (EmailAddress_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailAddress_List{l}, err
}

func (s EmailAddress_List) At(i int) EmailAddress { return EmailAddress{s.List.Struct(i)} }

func (s EmailAddress_List) Set(i int, v EmailAddress) error { return s.List.SetStruct(i, v.Struct) }

// EmailAddress_Promise is a wrapper for a EmailAddress promised by a client call.
type EmailAddress_Promise struct{ *capnp.Pipeline }

func (p EmailAddress_Promise) Struct() (EmailAddress, error) {
	s, err := p.Pipeline.Struct()
	return EmailAddress{s}, err
}

type EmailAttachment struct{ capnp.Struct }

func NewEmailAttachment(s *capnp.Segment) (EmailAttachment, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return EmailAttachment{st}, err
}

func NewRootEmailAttachment(s *capnp.Segment) (EmailAttachment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return EmailAttachment{st}, err
}

func ReadRootEmailAttachment(msg *capnp.Message) (EmailAttachment, error) {
	root, err := msg.RootPtr()
	return EmailAttachment{root.Struct()}, err
}

func (s EmailAttachment) String() string {
	str, _ := text.Marshal(0xb309c51a9d28244f, s.Struct)
	return str
}

func (s EmailAttachment) ContentType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailAttachment) HasContentType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s EmailAttachment) ContentDisposition() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailAttachment) HasContentDisposition() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentDispositionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentDisposition(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s EmailAttachment) ContentId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s EmailAttachment) HasContentId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s EmailAttachment) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s EmailAttachment) HasContent() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

// EmailAttachment_List is a list of EmailAttachment.
type EmailAttachment_List struct{ capnp.List }

// NewEmailAttachment creates a new list of EmailAttachment.
func NewEmailAttachment_List(s *capnp.Segment, sz int32) (EmailAttachment_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return EmailAttachment_List{l}, err
}

func (s EmailAttachment_List) At(i int) EmailAttachment { return EmailAttachment{s.List.Struct(i)} }

func (s EmailAttachment_List) Set(i int, v EmailAttachment) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAttachment_Promise is a wrapper for a EmailAttachment promised by a client call.
type EmailAttachment_Promise struct{ *capnp.Pipeline }

func (p EmailAttachment_Promise) Struct() (EmailAttachment, error) {
	s, err := p.Pipeline.Struct()
	return EmailAttachment{s}, err
}

type EmailMessage struct{ capnp.Struct }

func NewEmailMessage(s *capnp.Segment) (EmailMessage, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12})
	return EmailMessage{st}, err
}

func NewRootEmailMessage(s *capnp.Segment) (EmailMessage, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12})
	return EmailMessage{st}, err
}

func ReadRootEmailMessage(msg *capnp.Message) (EmailMessage, error) {
	root, err := msg.RootPtr()
	return EmailMessage{root.Struct()}, err
}

func (s EmailMessage) String() string {
	str, _ := text.Marshal(0xcff459e769562d2f, s.Struct)
	return str
}

func (s EmailMessage) Date() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s EmailMessage) SetDate(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s EmailMessage) From() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailMessage) HasFrom() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetFrom(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFrom sets the from field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewFrom() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailMessage) To() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(1)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasTo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetTo(v EmailAddress_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewTo sets the to field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewTo(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Cc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(2)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasCc() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetCc(v EmailAddress_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewCc sets the cc field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewCc(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Bcc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(3)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasBcc() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetBcc(v EmailAddress_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewBcc sets the bcc field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewBcc(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) ReplyTo() (EmailAddress, error) {
	p, err := s.Struct.Ptr(4)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailMessage) HasReplyTo() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetReplyTo(v EmailAddress) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewReplyTo sets the replyTo field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewReplyTo() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailMessage) MessageId() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s EmailMessage) HasMessageId() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s EmailMessage) MessageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s EmailMessage) SetMessageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, t.List.ToPtr())
}

func (s EmailMessage) References() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(6)
	return capnp.TextList{List: p.List()}, err
}

func (s EmailMessage) HasReferences() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetReferences(v capnp.TextList) error {
	return s.Struct.SetPtr(6, v.List.ToPtr())
}

// NewReferences sets the references field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s EmailMessage) NewReferences(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(6, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) InReplyTo() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(7)
	return capnp.TextList{List: p.List()}, err
}

func (s EmailMessage) HasInReplyTo() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetInReplyTo(v capnp.TextList) error {
	return s.Struct.SetPtr(7, v.List.ToPtr())
}

// NewInReplyTo sets the inReplyTo field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s EmailMessage) NewInReplyTo(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(7, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Subject() (string, error) {
	p, err := s.Struct.Ptr(8)
	return p.Text(), err
}

func (s EmailMessage) HasSubject() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SubjectBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(8)
	return p.TextBytes(), err
}

func (s EmailMessage) SetSubject(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(8, t.List.ToPtr())
}

func (s EmailMessage) Text() (string, error) {
	p, err := s.Struct.Ptr(9)
	return p.Text(), err
}

func (s EmailMessage) HasText() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s EmailMessage) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(9)
	return p.TextBytes(), err
}

func (s EmailMessage) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(9, t.List.ToPtr())
}

func (s EmailMessage) Html() (string, error) {
	p, err := s.Struct.Ptr(10)
	return p.Text(), err
}

func (s EmailMessage) HasHtml() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s EmailMessage) HtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(10)
	return p.TextBytes(), err
}

func (s EmailMessage) SetHtml(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(10, t.List.ToPtr())
}

func (s EmailMessage) Attachments() (EmailAttachment_List, error) {
	p, err := s.Struct.Ptr(11)
	return EmailAttachment_List{List: p.List()}, err
}

func (s EmailMessage) HasAttachments() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetAttachments(v EmailAttachment_List) error {
	return s.Struct.SetPtr(11, v.List.ToPtr())
}

// NewAttachments sets the attachments field to a newly
// allocated EmailAttachment_List, preferring placement in s's segment.
func (s EmailMessage) NewAttachments(n int32) (EmailAttachment_List, error) {
	l, err := NewEmailAttachment_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAttachment_List{}, err
	}
	err = s.Struct.SetPtr(11, l.List.ToPtr())
	return l, err
}

// EmailMessage_List is a list of EmailMessage.
type EmailMessage_List struct{ capnp.List }

// NewEmailMessage creates a new list of EmailMessage.
func NewEmailMessage_List(s *capnp.Segment, sz int32) (EmailMessage_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12}, sz)
	return EmailMessage_List{l}, err
}

func (s EmailMessage_List) At(i int) EmailMessage { return EmailMessage{s.List.Struct(i)} }

func (s EmailMessage_List) Set(i int, v EmailMessage) error { return s.List.SetStruct(i, v.Struct) }

// EmailMessage_Promise is a wrapper for a EmailMessage promised by a client call.
type EmailMessage_Promise struct{ *capnp.Pipeline }

func (p EmailMessage_Promise) Struct() (EmailMessage, error) {
	s, err := p.Pipeline.Struct()
	return EmailMessage{s}, err
}

func (p EmailMessage_Promise) From() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p EmailMessage_Promise) ReplyTo() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

type EmailSendPort struct{ Client capnp.Client }

func (c EmailSendPort) Send(ctx context.Context, params func(EmailSendPort_send_Params) error, opts ...capnp.CallOption) EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_send_Params{Struct: s}) }
	}
	return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailSendPort) HintAddress(ctx context.Context, params func(EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailSendPort_Server interface {
	Send(EmailSendPort_send) error

	HintAddress(EmailSendPort_hintAddress) error
}

func EmailSendPort_ServerToClient(s EmailSendPort_Server) EmailSendPort {
	c, _ := s.(server.Closer)
	return EmailSendPort{Client: server.New(EmailSendPort_Methods(nil, s), c)}
}

func EmailSendPort_Methods(methods []server.Method, s EmailSendPort_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailSendPort_send{c, opts, EmailSendPort_send_Params{Struct: p}, EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailSendPort_hintAddress{c, opts, EmailSendPort_hintAddress_Params{Struct: p}, EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// EmailSendPort_send holds the arguments for a server call to EmailSendPort.send.
type EmailSendPort_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailSendPort_send_Params
	Results EmailSendPort_send_Results
}

// EmailSendPort_hintAddress holds the arguments for a server call to EmailSendPort.hintAddress.
type EmailSendPort_hintAddress struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailSendPort_hintAddress_Params
	Results EmailSendPort_hintAddress_Results
}

type EmailSendPort_PowerboxTag struct{ capnp.Struct }

func NewEmailSendPort_PowerboxTag(s *capnp.Segment) (EmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailSendPort_PowerboxTag{st}, err
}

func NewRootEmailSendPort_PowerboxTag(s *capnp.Segment) (EmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailSendPort_PowerboxTag{st}, err
}

func ReadRootEmailSendPort_PowerboxTag(msg *capnp.Message) (EmailSendPort_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_PowerboxTag{root.Struct()}, err
}

func (s EmailSendPort_PowerboxTag) String() string {
	str, _ := text.Marshal(0x90790c61fc899dd3, s.Struct)
	return str
}

func (s EmailSendPort_PowerboxTag) FromHint() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailSendPort_PowerboxTag) HasFromHint() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_PowerboxTag) SetFromHint(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFromHint sets the fromHint field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_PowerboxTag) NewFromHint() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailSendPort_PowerboxTag) ListIdHint() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailSendPort_PowerboxTag) HasListIdHint() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_PowerboxTag) ListIdHintBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailSendPort_PowerboxTag) SetListIdHint(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// EmailSendPort_PowerboxTag_List is a list of EmailSendPort_PowerboxTag.
type EmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewEmailSendPort_PowerboxTag creates a new list of EmailSendPort_PowerboxTag.
func NewEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (EmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailSendPort_PowerboxTag_List{l}, err
}

func (s EmailSendPort_PowerboxTag_List) At(i int) EmailSendPort_PowerboxTag {
	return EmailSendPort_PowerboxTag{s.List.Struct(i)}
}

func (s EmailSendPort_PowerboxTag_List) Set(i int, v EmailSendPort_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_PowerboxTag_Promise is a wrapper for a EmailSendPort_PowerboxTag promised by a client call.
type EmailSendPort_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_PowerboxTag_Promise) Struct() (EmailSendPort_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_PowerboxTag{s}, err
}

func (p EmailSendPort_PowerboxTag_Promise) FromHint() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_send_Params struct{ capnp.Struct }

func NewEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_send_Params{st}, err
}

func NewRootEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_send_Params{st}, err
}

func ReadRootEmailSendPort_send_Params(msg *capnp.Message) (EmailSendPort_send_Params, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_send_Params{root.Struct()}, err
}

func (s EmailSendPort_send_Params) String() string {
	str, _ := text.Marshal(0xa5adb72b4ccc59ee, s.Struct)
	return str
}

func (s EmailSendPort_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Ptr(0)
	return EmailMessage{Struct: p.Struct()}, err
}

func (s EmailSendPort_send_Params) HasEmail() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_send_Params) SetEmail(v EmailMessage) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailSendPort_send_Params) NewEmail() (EmailMessage, error) {
	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailSendPort_send_Params_List is a list of EmailSendPort_send_Params.
type EmailSendPort_send_Params_List struct{ capnp.List }

// NewEmailSendPort_send_Params creates a new list of EmailSendPort_send_Params.
func NewEmailSendPort_send_Params_List(s *capnp.Segment, sz int32) (EmailSendPort_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailSendPort_send_Params_List{l}, err
}

func (s EmailSendPort_send_Params_List) At(i int) EmailSendPort_send_Params {
	return EmailSendPort_send_Params{s.List.Struct(i)}
}

func (s EmailSendPort_send_Params_List) Set(i int, v EmailSendPort_send_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_send_Params_Promise is a wrapper for a EmailSendPort_send_Params promised by a client call.
type EmailSendPort_send_Params_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_send_Params_Promise) Struct() (EmailSendPort_send_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_send_Params{s}, err
}

func (p EmailSendPort_send_Params_Promise) Email() EmailMessage_Promise {
	return EmailMessage_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_send_Results struct{ capnp.Struct }

func NewEmailSendPort_send_Results(s *capnp.Segment) (EmailSendPort_send_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_send_Results{st}, err
}

func NewRootEmailSendPort_send_Results(s *capnp.Segment) (EmailSendPort_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_send_Results{st}, err
}

func ReadRootEmailSendPort_send_Results(msg *capnp.Message) (EmailSendPort_send_Results, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_send_Results{root.Struct()}, err
}

func (s EmailSendPort_send_Results) String() string {
	str, _ := text.Marshal(0xd063b4e6c91bf8d8, s.Struct)
	return str
}

// EmailSendPort_send_Results_List is a list of EmailSendPort_send_Results.
type EmailSendPort_send_Results_List struct{ capnp.List }

// NewEmailSendPort_send_Results creates a new list of EmailSendPort_send_Results.
func NewEmailSendPort_send_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailSendPort_send_Results_List{l}, err
}

func (s EmailSendPort_send_Results_List) At(i int) EmailSendPort_send_Results {
	return EmailSendPort_send_Results{s.List.Struct(i)}
}

func (s EmailSendPort_send_Results_List) Set(i int, v EmailSendPort_send_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_send_Results_Promise is a wrapper for a EmailSendPort_send_Results promised by a client call.
type EmailSendPort_send_Results_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_send_Results_Promise) Struct() (EmailSendPort_send_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_send_Results{s}, err
}

type EmailSendPort_hintAddress_Params struct{ capnp.Struct }

func NewEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_hintAddress_Params{st}, err
}

func NewRootEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_hintAddress_Params{st}, err
}

func ReadRootEmailSendPort_hintAddress_Params(msg *capnp.Message) (EmailSendPort_hintAddress_Params, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_hintAddress_Params{root.Struct()}, err
}

func (s EmailSendPort_hintAddress_Params) String() string {
	str, _ := text.Marshal(0x9c78c3c5de56e4d4, s.Struct)
	return str
}

func (s EmailSendPort_hintAddress_Params) Address() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailSendPort_hintAddress_Params) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_hintAddress_Params) SetAddress(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAddress sets the address field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_hintAddress_Params) NewAddress() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailSendPort_hintAddress_Params_List is a list of EmailSendPort_hintAddress_Params.
type EmailSendPort_hintAddress_Params_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Params creates a new list of EmailSendPort_hintAddress_Params.
func NewEmailSendPort_hintAddress_Params_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailSendPort_hintAddress_Params_List{l}, err
}

func (s EmailSendPort_hintAddress_Params_List) At(i int) EmailSendPort_hintAddress_Params {
	return EmailSendPort_hintAddress_Params{s.List.Struct(i)}
}

func (s EmailSendPort_hintAddress_Params_List) Set(i int, v EmailSendPort_hintAddress_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_hintAddress_Params_Promise is a wrapper for a EmailSendPort_hintAddress_Params promised by a client call.
type EmailSendPort_hintAddress_Params_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_hintAddress_Params_Promise) Struct() (EmailSendPort_hintAddress_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_hintAddress_Params{s}, err
}

func (p EmailSendPort_hintAddress_Params_Promise) Address() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_hintAddress_Results struct{ capnp.Struct }

func NewEmailSendPort_hintAddress_Results(s *capnp.Segment) (EmailSendPort_hintAddress_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_hintAddress_Results{st}, err
}

func NewRootEmailSendPort_hintAddress_Results(s *capnp.Segment) (EmailSendPort_hintAddress_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_hintAddress_Results{st}, err
}

func ReadRootEmailSendPort_hintAddress_Results(msg *capnp.Message) (EmailSendPort_hintAddress_Results, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_hintAddress_Results{root.Struct()}, err
}

func (s EmailSendPort_hintAddress_Results) String() string {
	str, _ := text.Marshal(0xbd727a009329aabc, s.Struct)
	return str
}

// EmailSendPort_hintAddress_Results_List is a list of EmailSendPort_hintAddress_Results.
type EmailSendPort_hintAddress_Results_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Results creates a new list of EmailSendPort_hintAddress_Results.
func NewEmailSendPort_hintAddress_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailSendPort_hintAddress_Results_List{l}, err
}

func (s EmailSendPort_hintAddress_Results_List) At(i int) EmailSendPort_hintAddress_Results {
	return EmailSendPort_hintAddress_Results{s.List.Struct(i)}
}

func (s EmailSendPort_hintAddress_Results_List) Set(i int, v EmailSendPort_hintAddress_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_hintAddress_Results_Promise is a wrapper for a EmailSendPort_hintAddress_Results promised by a client call.
type EmailSendPort_hintAddress_Results_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_hintAddress_Results_Promise) Struct() (EmailSendPort_hintAddress_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_hintAddress_Results{s}, err
}

type VerifiedEmail struct{ Client capnp.Client }

type VerifiedEmail_Server interface {
}

func VerifiedEmail_ServerToClient(s VerifiedEmail_Server) VerifiedEmail {
	c, _ := s.(server.Closer)
	return VerifiedEmail{Client: server.New(VerifiedEmail_Methods(nil, s), c)}
}

func VerifiedEmail_Methods(methods []server.Method, s VerifiedEmail_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type VerifiedEmail_PowerboxTag struct{ capnp.Struct }

func NewVerifiedEmail_PowerboxTag(s *capnp.Segment) (VerifiedEmail_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return VerifiedEmail_PowerboxTag{st}, err
}

func NewRootVerifiedEmail_PowerboxTag(s *capnp.Segment) (VerifiedEmail_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return VerifiedEmail_PowerboxTag{st}, err
}

func ReadRootVerifiedEmail_PowerboxTag(msg *capnp.Message) (VerifiedEmail_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return VerifiedEmail_PowerboxTag{root.Struct()}, err
}

func (s VerifiedEmail_PowerboxTag) String() string {
	str, _ := text.Marshal(0x97469291ac5bb892, s.Struct)
	return str
}

func (s VerifiedEmail_PowerboxTag) VerifierId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s VerifiedEmail_PowerboxTag) HasVerifierId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) SetVerifierId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s VerifiedEmail_PowerboxTag) Address() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s VerifiedEmail_PowerboxTag) HasAddress() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s VerifiedEmail_PowerboxTag) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s VerifiedEmail_PowerboxTag) Domain() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s VerifiedEmail_PowerboxTag) HasDomain() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) DomainBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s VerifiedEmail_PowerboxTag) SetDomain(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// VerifiedEmail_PowerboxTag_List is a list of VerifiedEmail_PowerboxTag.
type VerifiedEmail_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmail_PowerboxTag creates a new list of VerifiedEmail_PowerboxTag.
func NewVerifiedEmail_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmail_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return VerifiedEmail_PowerboxTag_List{l}, err
}

func (s VerifiedEmail_PowerboxTag_List) At(i int) VerifiedEmail_PowerboxTag {
	return VerifiedEmail_PowerboxTag{s.List.Struct(i)}
}

func (s VerifiedEmail_PowerboxTag_List) Set(i int, v VerifiedEmail_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerifiedEmail_PowerboxTag_Promise is a wrapper for a VerifiedEmail_PowerboxTag promised by a client call.
type VerifiedEmail_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p VerifiedEmail_PowerboxTag_Promise) Struct() (VerifiedEmail_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return VerifiedEmail_PowerboxTag{s}, err
}

type VerifiedEmailSendPort struct{ Client capnp.Client }

func (c VerifiedEmailSendPort) Send(ctx context.Context, params func(EmailSendPort_send_Params) error, opts ...capnp.CallOption) EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_send_Params{Struct: s}) }
	}
	return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c VerifiedEmailSendPort) HintAddress(ctx context.Context, params func(EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type VerifiedEmailSendPort_Server interface {
	Send(EmailSendPort_send) error

	HintAddress(EmailSendPort_hintAddress) error
}

func VerifiedEmailSendPort_ServerToClient(s VerifiedEmailSendPort_Server) VerifiedEmailSendPort {
	c, _ := s.(server.Closer)
	return VerifiedEmailSendPort{Client: server.New(VerifiedEmailSendPort_Methods(nil, s), c)}
}

func VerifiedEmailSendPort_Methods(methods []server.Method, s VerifiedEmailSendPort_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailSendPort_send{c, opts, EmailSendPort_send_Params{Struct: p}, EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailSendPort_hintAddress{c, opts, EmailSendPort_hintAddress_Params{Struct: p}, EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

type VerifiedEmailSendPort_PowerboxTag struct{ capnp.Struct }

func NewVerifiedEmailSendPort_PowerboxTag(s *capnp.Segment) (VerifiedEmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerifiedEmailSendPort_PowerboxTag{st}, err
}

func NewRootVerifiedEmailSendPort_PowerboxTag(s *capnp.Segment) (VerifiedEmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerifiedEmailSendPort_PowerboxTag{st}, err
}

func ReadRootVerifiedEmailSendPort_PowerboxTag(msg *capnp.Message) (VerifiedEmailSendPort_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return VerifiedEmailSendPort_PowerboxTag{root.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) String() string {
	str, _ := text.Marshal(0x8f555bd4141fbb3b, s.Struct)
	return str
}

func (s VerifiedEmailSendPort_PowerboxTag) Verification() (VerifiedEmail_PowerboxTag, error) {
	p, err := s.Struct.Ptr(0)
	return VerifiedEmail_PowerboxTag{Struct: p.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) HasVerification() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetVerification(v VerifiedEmail_PowerboxTag) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewVerification sets the verification field to a newly
// allocated VerifiedEmail_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewVerification() (VerifiedEmail_PowerboxTag, error) {
	ss, err := NewVerifiedEmail_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedEmailSendPort_PowerboxTag) Port() (EmailSendPort_PowerboxTag, error) {
	p, err := s.Struct.Ptr(1)
	return EmailSendPort_PowerboxTag{Struct: p.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) HasPort() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetPort(v EmailSendPort_PowerboxTag) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPort sets the port field to a newly
// allocated EmailSendPort_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewPort() (EmailSendPort_PowerboxTag, error) {
	ss, err := NewEmailSendPort_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerifiedEmailSendPort_PowerboxTag_List is a list of VerifiedEmailSendPort_PowerboxTag.
type VerifiedEmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmailSendPort_PowerboxTag creates a new list of VerifiedEmailSendPort_PowerboxTag.
func NewVerifiedEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return VerifiedEmailSendPort_PowerboxTag_List{l}, err
}

func (s VerifiedEmailSendPort_PowerboxTag_List) At(i int) VerifiedEmailSendPort_PowerboxTag {
	return VerifiedEmailSendPort_PowerboxTag{s.List.Struct(i)}
}

func (s VerifiedEmailSendPort_PowerboxTag_List) Set(i int, v VerifiedEmailSendPort_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerifiedEmailSendPort_PowerboxTag_Promise is a wrapper for a VerifiedEmailSendPort_PowerboxTag promised by a client call.
type VerifiedEmailSendPort_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Struct() (VerifiedEmailSendPort_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return VerifiedEmailSendPort_PowerboxTag{s}, err
}

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Verification() VerifiedEmail_PowerboxTag_Promise {
	return VerifiedEmail_PowerboxTag_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Port() EmailSendPort_PowerboxTag_Promise {
	return EmailSendPort_PowerboxTag_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type EmailVerifier struct{ Client capnp.Client }

func (c EmailVerifier) GetId(ctx context.Context, params func(EmailVerifier_getId_Params) error, opts ...capnp.CallOption) EmailVerifier_getId_Results_Promise {
	if c.Client == nil {
		return EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailVerifier_getId_Params{Struct: s}) }
	}
	return EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailVerifier) VerifyEmail(ctx context.Context, params func(EmailVerifier_verifyEmail_Params) error, opts ...capnp.CallOption) EmailVerifier_verifyEmail_Results_Promise {
	if c.Client == nil {
		return EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailVerifier_verifyEmail_Params{Struct: s}) }
	}
	return EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailVerifier_Server interface {
	GetId(EmailVerifier_getId) error

	VerifyEmail(EmailVerifier_verifyEmail) error
}

func EmailVerifier_ServerToClient(s EmailVerifier_Server) EmailVerifier {
	c, _ := s.(server.Closer)
	return EmailVerifier{Client: server.New(EmailVerifier_Methods(nil, s), c)}
}

func EmailVerifier_Methods(methods []server.Method, s EmailVerifier_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailVerifier_getId{c, opts, EmailVerifier_getId_Params{Struct: p}, EmailVerifier_getId_Results{Struct: r}}
			return s.GetId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailVerifier_verifyEmail{c, opts, EmailVerifier_verifyEmail_Params{Struct: p}, EmailVerifier_verifyEmail_Results{Struct: r}}
			return s.VerifyEmail(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// EmailVerifier_getId holds the arguments for a server call to EmailVerifier.getId.
type EmailVerifier_getId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailVerifier_getId_Params
	Results EmailVerifier_getId_Results
}

// EmailVerifier_verifyEmail holds the arguments for a server call to EmailVerifier.verifyEmail.
type EmailVerifier_verifyEmail struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailVerifier_verifyEmail_Params
	Results EmailVerifier_verifyEmail_Results
}

type EmailVerifier_getId_Params struct{ capnp.Struct }

func NewEmailVerifier_getId_Params(s *capnp.Segment) (EmailVerifier_getId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailVerifier_getId_Params{st}, err
}

func NewRootEmailVerifier_getId_Params(s *capnp.Segment) (EmailVerifier_getId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailVerifier_getId_Params{st}, err
}

func ReadRootEmailVerifier_getId_Params(msg *capnp.Message) (EmailVerifier_getId_Params, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_getId_Params{root.Struct()}, err
}

func (s EmailVerifier_getId_Params) String() string {
	str, _ := text.Marshal(0xe5927352f65eba5c, s.Struct)
	return str
}

// EmailVerifier_getId_Params_List is a list of EmailVerifier_getId_Params.
type EmailVerifier_getId_Params_List struct{ capnp.List }

// NewEmailVerifier_getId_Params creates a new list of EmailVerifier_getId_Params.
func NewEmailVerifier_getId_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailVerifier_getId_Params_List{l}, err
}

func (s EmailVerifier_getId_Params_List) At(i int) EmailVerifier_getId_Params {
	return EmailVerifier_getId_Params{s.List.Struct(i)}
}

func (s EmailVerifier_getId_Params_List) Set(i int, v EmailVerifier_getId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_getId_Params_Promise is a wrapper for a EmailVerifier_getId_Params promised by a client call.
type EmailVerifier_getId_Params_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_getId_Params_Promise) Struct() (EmailVerifier_getId_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_getId_Params{s}, err
}

type EmailVerifier_getId_Results struct{ capnp.Struct }

func NewEmailVerifier_getId_Results(s *capnp.Segment) (EmailVerifier_getId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_getId_Results{st}, err
}

func NewRootEmailVerifier_getId_Results(s *capnp.Segment) (EmailVerifier_getId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_getId_Results{st}, err
}

func ReadRootEmailVerifier_getId_Results(msg *capnp.Message) (EmailVerifier_getId_Results, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_getId_Results{root.Struct()}, err
}

func (s EmailVerifier_getId_Results) String() string {
	str, _ := text.Marshal(0xc7e287c5d3518c34, s.Struct)
	return str
}

func (s EmailVerifier_getId_Results) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s EmailVerifier_getId_Results) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_getId_Results) SetId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// EmailVerifier_getId_Results_List is a list of EmailVerifier_getId_Results.
type EmailVerifier_getId_Results_List struct{ capnp.List }

// NewEmailVerifier_getId_Results creates a new list of EmailVerifier_getId_Results.
func NewEmailVerifier_getId_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailVerifier_getId_Results_List{l}, err
}

func (s EmailVerifier_getId_Results_List) At(i int) EmailVerifier_getId_Results {
	return EmailVerifier_getId_Results{s.List.Struct(i)}
}

func (s EmailVerifier_getId_Results_List) Set(i int, v EmailVerifier_getId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_getId_Results_Promise is a wrapper for a EmailVerifier_getId_Results promised by a client call.
type EmailVerifier_getId_Results_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_getId_Results_Promise) Struct() (EmailVerifier_getId_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_getId_Results{s}, err
}

type EmailVerifier_verifyEmail_Params struct{ capnp.Struct }

func NewEmailVerifier_verifyEmail_Params(s *capnp.Segment) (EmailVerifier_verifyEmail_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailVerifier_verifyEmail_Params{st}, err
}

func NewRootEmailVerifier_verifyEmail_Params(s *capnp.Segment) (EmailVerifier_verifyEmail_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailVerifier_verifyEmail_Params{st}, err
}

func ReadRootEmailVerifier_verifyEmail_Params(msg *capnp.Message) (EmailVerifier_verifyEmail_Params, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_verifyEmail_Params{root.Struct()}, err
}

func (s EmailVerifier_verifyEmail_Params) String() string {
	str, _ := text.Marshal(0x93ee926bb1bd4eea, s.Struct)
	return str
}

func (s EmailVerifier_verifyEmail_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s EmailVerifier_verifyEmail_Params) HasTabId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Params) SetTabId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s EmailVerifier_verifyEmail_Params) Verification() VerifiedEmail {
	p, _ := s.Struct.Ptr(1)
	return VerifiedEmail{Client: p.Interface().Client()}
}

func (s EmailVerifier_verifyEmail_Params) HasVerification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Params) SetVerification(v VerifiedEmail) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// EmailVerifier_verifyEmail_Params_List is a list of EmailVerifier_verifyEmail_Params.
type EmailVerifier_verifyEmail_Params_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Params creates a new list of EmailVerifier_verifyEmail_Params.
func NewEmailVerifier_verifyEmail_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailVerifier_verifyEmail_Params_List{l}, err
}

func (s EmailVerifier_verifyEmail_Params_List) At(i int) EmailVerifier_verifyEmail_Params {
	return EmailVerifier_verifyEmail_Params{s.List.Struct(i)}
}

func (s EmailVerifier_verifyEmail_Params_List) Set(i int, v EmailVerifier_verifyEmail_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_verifyEmail_Params_Promise is a wrapper for a EmailVerifier_verifyEmail_Params promised by a client call.
type EmailVerifier_verifyEmail_Params_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_verifyEmail_Params_Promise) Struct() (EmailVerifier_verifyEmail_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_verifyEmail_Params{s}, err
}

func (p EmailVerifier_verifyEmail_Params_Promise) Verification() VerifiedEmail {
	return VerifiedEmail{Client: p.Pipeline.GetPipeline(1).Client()}
}

type EmailVerifier_verifyEmail_Results struct{ capnp.Struct }

func NewEmailVerifier_verifyEmail_Results(s *capnp.Segment) (EmailVerifier_verifyEmail_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_verifyEmail_Results{st}, err
}

func NewRootEmailVerifier_verifyEmail_Results(s *capnp.Segment) (EmailVerifier_verifyEmail_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_verifyEmail_Results{st}, err
}

func ReadRootEmailVerifier_verifyEmail_Results(msg *capnp.Message) (EmailVerifier_verifyEmail_Results, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_verifyEmail_Results{root.Struct()}, err
}

func (s EmailVerifier_verifyEmail_Results) String() string {
	str, _ := text.Marshal(0xcc99614322e49040, s.Struct)
	return str
}

func (s EmailVerifier_verifyEmail_Results) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailVerifier_verifyEmail_Results) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Results) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailVerifier_verifyEmail_Results) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// EmailVerifier_verifyEmail_Results_List is a list of EmailVerifier_verifyEmail_Results.
type EmailVerifier_verifyEmail_Results_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Results creates a new list of EmailVerifier_verifyEmail_Results.
func NewEmailVerifier_verifyEmail_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailVerifier_verifyEmail_Results_List{l}, err
}

func (s EmailVerifier_verifyEmail_Results_List) At(i int) EmailVerifier_verifyEmail_Results {
	return EmailVerifier_verifyEmail_Results{s.List.Struct(i)}
}

func (s EmailVerifier_verifyEmail_Results_List) Set(i int, v EmailVerifier_verifyEmail_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_verifyEmail_Results_Promise is a wrapper for a EmailVerifier_verifyEmail_Results promised by a client call.
type EmailVerifier_verifyEmail_Results_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_verifyEmail_Results_Promise) Struct() (EmailVerifier_verifyEmail_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_verifyEmail_Results{s}, err
}

type EmailAgent struct{ Client capnp.Client }

func (c EmailAgent) Send(ctx context.Context, params func(EmailAgent_send_Params) error, opts ...capnp.CallOption) EmailAgent_send_Results_Promise {
	if c.Client == nil {
		return EmailAgent_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailAgent_send_Params{Struct: s}) }
	}
	return EmailAgent_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailAgent) AddReceiver(ctx context.Context, params func(EmailAgent_addReceiver_Params) error, opts ...capnp.CallOption) EmailAgent_addReceiver_Results_Promise {
	if c.Client == nil {
		return EmailAgent_addReceiver_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "addReceiver",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailAgent_addReceiver_Params{Struct: s}) }
	}
	return EmailAgent_addReceiver_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailAgent_Server interface {
	Send(EmailAgent_send) error

	AddReceiver(EmailAgent_addReceiver) error
}

func EmailAgent_ServerToClient(s EmailAgent_Server) EmailAgent {
	c, _ := s.(server.Closer)
	return EmailAgent{Client: server.New(EmailAgent_Methods(nil, s), c)}
}

func EmailAgent_Methods(methods []server.Method, s EmailAgent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailAgent_send{c, opts, EmailAgent_send_Params{Struct: p}, EmailAgent_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "addReceiver",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailAgent_addReceiver{c, opts, EmailAgent_addReceiver_Params{Struct: p}, EmailAgent_addReceiver_Results{Struct: r}}
			return s.AddReceiver(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// EmailAgent_send holds the arguments for a server call to EmailAgent.send.
type EmailAgent_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailAgent_send_Params
	Results EmailAgent_send_Results
}

// EmailAgent_addReceiver holds the arguments for a server call to EmailAgent.addReceiver.
type EmailAgent_addReceiver struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailAgent_addReceiver_Params
	Results EmailAgent_addReceiver_Results
}

type EmailAgent_send_Params struct{ capnp.Struct }

func NewEmailAgent_send_Params(s *capnp.Segment) (EmailAgent_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_send_Params{st}, err
}

func NewRootEmailAgent_send_Params(s *capnp.Segment) (EmailAgent_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_send_Params{st}, err
}

func ReadRootEmailAgent_send_Params(msg *capnp.Message) (EmailAgent_send_Params, error) {
	root, err := msg.RootPtr()
	return EmailAgent_send_Params{root.Struct()}, err
}

func (s EmailAgent_send_Params) String() string {
	str, _ := text.Marshal(0xa8eb16da45ad8f97, s.Struct)
	return str
}

func (s EmailAgent_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Ptr(0)
	return EmailMessage{Struct: p.Struct()}, err
}

func (s EmailAgent_send_Params) HasEmail() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_send_Params) SetEmail(v EmailMessage) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailAgent_send_Params) NewEmail() (EmailMessage, error) {
	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailAgent_send_Params_List is a list of EmailAgent_send_Params.
type EmailAgent_send_Params_List struct{ capnp.List }

// NewEmailAgent_send_Params creates a new list of EmailAgent_send_Params.
func NewEmailAgent_send_Params_List(s *capnp.Segment, sz int32) (EmailAgent_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_send_Params_List{l}, err
}

func (s EmailAgent_send_Params_List) At(i int) EmailAgent_send_Params {
	return EmailAgent_send_Params{s.List.Struct(i)}
}

func (s EmailAgent_send_Params_List) Set(i int, v EmailAgent_send_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_send_Params_Promise is a wrapper for a EmailAgent_send_Params promised by a client call.
type EmailAgent_send_Params_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_send_Params_Promise) Struct() (EmailAgent_send_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_send_Params{s}, err
}

func (p EmailAgent_send_Params_Promise) Email() EmailMessage_Promise {
	return EmailMessage_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailAgent_send_Results struct{ capnp.Struct }

func NewEmailAgent_send_Results(s *capnp.Segment) (EmailAgent_send_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailAgent_send_Results{st}, err
}

func NewRootEmailAgent_send_Results(s *capnp.Segment) (EmailAgent_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailAgent_send_Results{st}, err
}

func ReadRootEmailAgent_send_Results(msg *capnp.Message) (EmailAgent_send_Results, error) {
	root, err := msg.RootPtr()
	return EmailAgent_send_Results{root.Struct()}, err
}

func (s EmailAgent_send_Results) String() string {
	str, _ := text.Marshal(0x81f33f1803485545, s.Struct)
	return str
}

// EmailAgent_send_Results_List is a list of EmailAgent_send_Results.
type EmailAgent_send_Results_List struct{ capnp.List }

// NewEmailAgent_send_Results creates a new list of EmailAgent_send_Results.
func NewEmailAgent_send_Results_List(s *capnp.Segment, sz int32) (EmailAgent_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailAgent_send_Results_List{l}, err
}

func (s EmailAgent_send_Results_List) At(i int) EmailAgent_send_Results {
	return EmailAgent_send_Results{s.List.Struct(i)}
}

func (s EmailAgent_send_Results_List) Set(i int, v EmailAgent_send_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_send_Results_Promise is a wrapper for a EmailAgent_send_Results promised by a client call.
type EmailAgent_send_Results_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_send_Results_Promise) Struct() (EmailAgent_send_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_send_Results{s}, err
}

type EmailAgent_addReceiver_Params struct{ capnp.Struct }

func NewEmailAgent_addReceiver_Params(s *capnp.Segment) (EmailAgent_addReceiver_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Params{st}, err
}

func NewRootEmailAgent_addReceiver_Params(s *capnp.Segment) (EmailAgent_addReceiver_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Params{st}, err
}

func ReadRootEmailAgent_addReceiver_Params(msg *capnp.Message) (EmailAgent_addReceiver_Params, error) {
	root, err := msg.RootPtr()
	return EmailAgent_addReceiver_Params{root.Struct()}, err
}

func (s EmailAgent_addReceiver_Params) String() string {
	str, _ := text.Marshal(0xfacf412b11767e9e, s.Struct)
	return str
}

func (s EmailAgent_addReceiver_Params) Port() EmailSendPort {
	p, _ := s.Struct.Ptr(0)
	return EmailSendPort{Client: p.Interface().Client()}
}

func (s EmailAgent_addReceiver_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_addReceiver_Params) SetPort(v EmailSendPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EmailAgent_addReceiver_Params_List is a list of EmailAgent_addReceiver_Params.
type EmailAgent_addReceiver_Params_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Params creates a new list of EmailAgent_addReceiver_Params.
func NewEmailAgent_addReceiver_Params_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_addReceiver_Params_List{l}, err
}

func (s EmailAgent_addReceiver_Params_List) At(i int) EmailAgent_addReceiver_Params {
	return EmailAgent_addReceiver_Params{s.List.Struct(i)}
}

func (s EmailAgent_addReceiver_Params_List) Set(i int, v EmailAgent_addReceiver_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_addReceiver_Params_Promise is a wrapper for a EmailAgent_addReceiver_Params promised by a client call.
type EmailAgent_addReceiver_Params_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_addReceiver_Params_Promise) Struct() (EmailAgent_addReceiver_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_addReceiver_Params{s}, err
}

func (p EmailAgent_addReceiver_Params_Promise) Port() EmailSendPort {
	return EmailSendPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type EmailAgent_addReceiver_Results struct{ capnp.Struct }

func NewEmailAgent_addReceiver_Results(s *capnp.Segment) (EmailAgent_addReceiver_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Results{st}, err
}

func NewRootEmailAgent_addReceiver_Results(s *capnp.Segment) (EmailAgent_addReceiver_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Results{st}, err
}

func ReadRootEmailAgent_addReceiver_Results(msg *capnp.Message) (EmailAgent_addReceiver_Results, error) {
	root, err := msg.RootPtr()
	return EmailAgent_addReceiver_Results{root.Struct()}, err
}

func (s EmailAgent_addReceiver_Results) String() string {
	str, _ := text.Marshal(0x8e8e3d68615d430c, s.Struct)
	return str
}

func (s EmailAgent_addReceiver_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s EmailAgent_addReceiver_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_addReceiver_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EmailAgent_addReceiver_Results_List is a list of EmailAgent_addReceiver_Results.
type EmailAgent_addReceiver_Results_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Results creates a new list of EmailAgent_addReceiver_Results.
func NewEmailAgent_addReceiver_Results_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_addReceiver_Results_List{l}, err
}

func (s EmailAgent_addReceiver_Results_List) At(i int) EmailAgent_addReceiver_Results {
	return EmailAgent_addReceiver_Results{s.List.Struct(i)}
}

func (s EmailAgent_addReceiver_Results_List) Set(i int, v EmailAgent_addReceiver_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_addReceiver_Results_Promise is a wrapper for a EmailAgent_addReceiver_Results promised by a client call.
type EmailAgent_addReceiver_Results_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_addReceiver_Results_Promise) Struct() (EmailAgent_addReceiver_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_addReceiver_Results{s}, err
}

func (p EmailAgent_addReceiver_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

const schema_dd10df585a82c6d8 = "x\xda\x9cWml\x14]\x15\xbegfgg\x17v" +
	"\xdd\x8e\xb3\xafh\x09ihJ\xe4\xc3V(\xfc\x10\x14" +
	"\xb7\x14J\x8a\x01\xdd\xdd\x96\xa6-`:\xdd\x1d\xda\xad" +
	"\xdb\xdd23\x85\xd6\x80\xb5*\x16\x83|\x95\x1f\xa8\x01" +
	"\x13\xa3$hZ\x91h\xa2A\x881Z\x91\x0a\xc1\x00" +
	"?\x1a\x03\xf2\x11\x10Q\xc0\x18A\x09\xd1\x8c\xe7\xde\xdd" +
	"\xf9\xe8t[\xc8\xfbc\x93\xd9\xe7>s\xces\xcf9" +
	"\xf7\xdc3+\xf3\xbe:n\x95p~\x1e!\x89\xfd\x82" +
	"\xdfl\xd8\xde\xc8\x7f8\xf6\xafa\"-\x00B|\"" +
	"!\xab\xd7\x0b\x95@|\xa6~\xe5\x0f}G\xde\xcb\x1f" +
	"&R\x887\xa7~\xf7\x95\xf6\xd6{ew\x09\x01y" +
	"\x89pQ\xae\x16\x90)/\x13F\xe4\x03\xf4\xc9\x0cm" +
	"\xdc\xa5t\xaf?z\xb4`F\x00jG\x15\xda\x01\xe9" +
	"{\x84\x18\x01\xf3\x93\xbf\xac\x88\xde\xde\xb1\xfd\x18\x91*" +
	"\xc1|:\xf5\xe3\x86\xe6C\xd7\xbeO\x04\x8e\x12O\x08" +
	"\x1a\xc8g\x99\xc5\xef\x09\xfb\x90|\xeb\xcc7\xfe\xab\x84" +
	"\x06\x8f\x13\xa9\x1c\xcc\xc9KW\xb7\xfej\xd1W\x9f\x15" +
	"\xc9\xe0\xaf\x05Y\xf2Sr\xd8O\xc9\x7f\xfb\xec\xe5\x0b" +
	"_\x18}q\x92\x92I\x91\xb4\xc7\xdfC]\x1f`\x84" +
	"\xd1_\xec\x18;1\xba\xf9\x14\xb3\xd6\xfef\xdbf\xee" +
	"\x9f\x87_\x13\x81\xa7\xc4\xbb\xd4\xdasf\xed\xa9\xff<" +
	"\x92o?j\xf9\xf3\xc4o\x06N\x17\xad\xb1\x8d\x1c\x14" +
	"\x99\xb5\x13\"\xdd\x88\xad\xdd\x13\x96\xd5\x17\xc4\x0f\x82\xfc" +
	"kq\x01\x9a\xba.\x8a\xf8\xfb(\xc6\xe5E\xdb\xb5\xad" +
	"+~>~\xd6mnB\xac\xa5\xe6\xae3s\xa7\x8e" +
	"\x8d7\xfc\xe9C\x7f?\xe7\x0e\xdcs\xb1\x9c\x12^1" +
	"\xc2\x83O\xb7|\xfd\xc5\x9d\xf11\xf4\x07\x8e?\xb6M" +
	"\xf9\xbd\xc0\xa4\xbc8@\x9f\x16\x05\xe8N?W\xb5\xf4" +
	"L\xf9D\xf0\xa7\x1e.\xcd\xaa<\x18x(\x1fd\xdc" +
	"\xe1\xc0\x13\x02\xff\xb8\xf4\xa3e'\xbf\xa8]f\xbaX" +
	"\xda\xd5\xa0F\xd3\xbe\xe6\x9b\x89[\x13#\x0f\xaf\xb8\x15" +
	"'\x82\xf5TP[\x90\x0a\xaa;\xfe\xa8r\xa3\xf2\xed" +
	"kn\xc2 }\x17\xe4\x83\x8c\xf0\xf1\xea\x96\xcc\x93\xb6" +
	"\x977H\"\x04n\x19!\xea\xfclpR\xbe\x10\xa4" +
	"\xef\x8c\x07\xefq\xc8\x9ez\xbd\xf0\xea_~\x96\xfa#" +
	"q\x94\x84\xc3\xeb\xa8\x12\xf3\x87\x0b\xcfL\xfe\xa7\xf5\xf6" +
	"\x8c\x02|\x15\xba)C\x98Z\xfb_hDN\xd0'" +
	"s\xe7\xc5\xcf\xff;\xa9\x8f>v\x99Y[0c\xd7" +
	"\x8f\xd7\xcc\xe2\xf0M\xb9:L\xf3\xb56<\"\x1fa" +
	"f\xec\xf2\xf0\x92\xfb\x91<\xcc\xc8H,\x92\xbf\xfb\xa5" +
	"\xbd\xd2\x8a\x0d7\xde\xb8s\xd7\x1fN\xd2H\x0c\x87c" +
	"\xa4\xcdT{\x95L\xb6&\xa5\xf0}\xb9\xbeu\x0d\xf4" +
	"\xcf\x86.5g\xd4\xe8j.]\x95T\xf5~1k" +
	"\xe86\x0b,V\xac@\x8b\x03$\x02\xbc\x80\x8e\xac\"" +
	"\x01\xeb\xb4J\xab\x96\x13NZ\"b|-\x11`\x1d" +
	"A\xe9#\x9d\xb8&\x89\x11\xea\xa5\x0eL%\x9dN\xaa" +
	")5C\xc4\xbd\xaaV\x07h\xd5\xf6\xe8\xf3\xe8\xb2\xa8" +
	"Hd\xf2\xb2\xbc\xa1'|\xbc\x0f\x03\x8a\x1b\x94\xc2\xeb" +
	"\xb0i\x04xHD9\x88u+\xb9tV\x05\xc9\xbc" +
	"_\xdf\xd11V\xf5\xf2[\x18%\x90\x88\xc7x\x8b\xaa" +
	"evg\xd44s\xd2\x84\x82\xe2y\xcd\xa8\x89\xe7\xf7" +
	"\xa9ZgE~\xa0Y\xe9*\xec\xd2r\xb1\xac\x07]" +
	",E\x17k8\x90\x00\xa2@A\xdc,I|\x0c\xc1" +
	"Op`\xeee\x16S\x0a\x89\x18\x99|\x0e\xca\x9c\xe3" +
	"\x8d\xfe\xcb\x08D\xfa\xd0\x03\xc2v\x0f)\xc0%r\xf1" +
	".r>SJN;\x82+\x11\xfc\x14\xca\xd9\xad\xe5" +
	"{\x1b39\x83\xa0\x9f2\xe7\xb0\x16\x9df3\xba\xb1" +
	"%\xdd\x98!|\xce\x80\x10\xe1\xf0W*\xfa\xc5(i" +
	"5ls\x83\x0c\xab\x8a+\x11M\xe9\xd5\xddjj\xd1" +
	"q\x15:^\xe9RS\xdd\xe3\x04\xa7\xc2P:\xb7\xa4" +
	"!\x8c\x8e\xc2dF\xa8$\xa7\xba=\xa9\xe2g\xa4j" +
	"FLB\xb6\x8a\x06\xba\xfdM\xe80\xeeR\xb1\xad\x1e" +
	"\xc1F\x04\x9b\x11\xe4\xb8(p\x08&h\xbdlE\xb0" +
	"\xd5\xce\x9b\xaa\x11\xdeQ8\x84\x05\xa7\xa9\xban\x85&" +
	"\x96\xce\xa3\xf3\xdc\x1c\x91\xb2s\xd6\x8d1\xdfPx\xdb" +
	"\x8e\x94\xabR\xeb\x9dJ\xb5\x9d\xccL\xcf\xec\xfb/\xfa" +
	"\x01#\xe1\x03\xf7\xe5\x05\x9df!2\xf9\x01\"\xd2\xd0" +
	"\xf0B\"\x00\xaek\x05O\xab\xddp\xf0y\x8e\xaac" +
	"M ^\xa1x\xa5\xd7:\xd2+\xd8\xeb(\xdcn\xa9" +
	"\xb3\x16\xb3\xab\xb1\xc4\xd1$\xff\xfemr\x8e\xcdB\xe0" +
	"\x08\xf1\x1c\x8a\xfaRe\xb8\xdc9)\xde\xb4FrJ" +
	"\xaf:#\xa9.?\x86\xa1\xa4\xba{\xc5b\xcf+\xb3" +
	"=)\x9dh\xb4\x03\x8df]\x9e2\xdfA0\x8b\xe0" +
	"\x80\xab\xd4\xfa\x93\x08\x1a\x08~\x19A\x9e\x8f\x02\x8f\xe0" +
	"\x01*t\x00\xc1\xafa\xfd\xa5\xf29\x03\x1d4\x13q" +
	"\xb0\xcf\x11SD7AF\xef\xcb\xeb\x19#\xc3\xe7s" +
	"\xde\xc5-\x04\xd2\x166T\xc4\xec3\xf6\x8e%\x9aT" +
	"+\xb0\x9d\xba\x9a}\x89\xc3\xdf\xa5b\xaf(\xf4]C" +
	"'\xee\xf4\x95;\xe9\xe33\xe99|\x97l$\x96\xef" +
	"\xb7\x9d\x8f\xd9\x13\xb4\x0d\x97\x95.\xb5P\x08k,+" +
	"\xf2.\xc0\xa47\xb5\x02\x0fMi\xe0\xa0\x98 Ya" +
	"\xf0N\x0aw\x03\xcd\x11\xb0\x1c\xc9*\xe0>\x9a:(" +
	"\xbe\x9f\xe2<\xc7\xd2$\x0f2\xdc\xa0\xf8q\x8a\xfb0" +
	"}>z\xd1B%\xe2\x87(\xfe\x03\x8a\x0b\xbe(\x08" +
	"tF\x04T\xdft\x9a\xe2\xe7(\xee\x17\xa2\xe0\xa7\xa3" +
	"\x05`\x11 \x15\xf1\x9fP\\\xf4G\xe9\x95,\x8f\x03" +
	"v\xac\xa61\x8a\xff\x9e\xe2\x011\x0a\x01\xc4'\x18\xff" +
	"\xb7\x14\xbfO\xf1` \x0aA\xc4\xef2\xfbS\x14\x7f" +
	"D\xf1y\xc1(\xe0\xb8,?`\xfb\xbaC\xf1\xbfR" +
	"|\xfe\xbc(\xccG\xfc1\xc3\xefS\xfc\x19\xc5C\xf3" +
	"\xa3\x18G\x1c(\x01\xcb\x17\xa9\x88\xfb8\x0e\"i\xc5" +
	"Pq\x03\x1c\xfe B\xef\x8d\x99\x1d\x897\xf2\xf0\x01" +
	"\x02q\x1e<k\x08\xf2\xa9\xd4\xackb\xe7\x1c\x8bC" +
	"\x9a\xda\x97\x1dl\xce\x97h\x80\xbd\x85\xc4\xba+\xdc\xd4" +
	"\xd4\xdd\xaa\xa6\xe6R\x84Wu\xcb&]\xc3G3\x93" +
	"K2[\x04\xf2\x9e\xa5!\xbd\xbf\xb3GM\xd9\xb7\\" +
	"\xc4P\x07\x9c?\xddFo\xd6v\xa0\x14\x8e\xbbJ\xc4" +
	"\x9c\xa1;\xaa\xed\xe1\xb5\xa0\xfa\xad\xad3\x19S\xa7\x1f" +
	")\xce{\x0c@sF(k>\x04k\xbe\x95V\xd5" +
	"Z#\x94\xf5\x05\x01\xd6hk\x8dP\x15\xecH\xd6\x15" +
	"//<NDD\xeb\xd3g(~\xb6\x83\x1c\x8b\xb3" +
	"\xd6^B\xde\xf4\xcb\xc5\xf9\xd8\xf1^.E\xe9\xd67" +
	"\x04X\xa3\xb2k\xfa\xb3>W\x80\x8e\xf3\x84\xce\xf3\x9e" +
	"\xe9\xcfjBD\xc4TOW\xceyo=\xc82A" +
	"\xce\xf7R\x89\xdb\xee]\xc7G\xb6y\x98\xd6s\x96;" +
	"=\xa70\xa4I\xae\xab\x92\x0d$\xff\x0f\x00\x00\xff\xff" +
	"\xbe\x8a\xefT"

func init() {
	schemas.Register(schema_dd10df585a82c6d8,
		0x81f33f1803485545,
		0x8b6f158d70cbc773,
		0x8e8e3d68615d430c,
		0x8f555bd4141fbb3b,
		0x90790c61fc899dd3,
		0x93ee926bb1bd4eea,
		0x97469291ac5bb892,
		0x9c78c3c5de56e4d4,
		0xa3cc885445aed8e9,
		0xa5adb72b4ccc59ee,
		0xa8eb16da45ad8f97,
		0xacaddcee86563ee1,
		0xb309c51a9d28244f,
		0xbd727a009329aabc,
		0xc7e287c5d3518c34,
		0xcc99614322e49040,
		0xcff459e769562d2f,
		0xd063b4e6c91bf8d8,
		0xd458f7ca9d1ba9ff,
		0xe5927352f65eba5c,
		0xec831dbf4cc9bcca,
		0xf88bf102464dfa5a,
		0xfacf412b11767e9e)
}
