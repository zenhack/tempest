package email

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type EmailAddress struct{ capnp.Struct }

func NewEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return EmailAddress{}, err
	}
	return EmailAddress{st}, nil
}

func NewRootEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return EmailAddress{}, err
	}
	return EmailAddress{st}, nil
}

func ReadRootEmailAddress(msg *capnp.Message) (EmailAddress, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return EmailAddress{}, err
	}
	return EmailAddress{root.Struct()}, nil
}

func (s EmailAddress) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailAddress) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAddress) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailAddress) HasName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAddress) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return EmailAddress_List{}, err
	}
	return EmailAddress_List{l}, nil
}

func (s EmailAddress_List) At(i int) EmailAddress           { return EmailAddress{s.List.Struct(i)} }
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
	if err != nil {
		return EmailAttachment{}, err
	}
	return EmailAttachment{st}, nil
}

func NewRootEmailAttachment(s *capnp.Segment) (EmailAttachment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	if err != nil {
		return EmailAttachment{}, err
	}
	return EmailAttachment{st}, nil
}

func ReadRootEmailAttachment(msg *capnp.Message) (EmailAttachment, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return EmailAttachment{}, err
	}
	return EmailAttachment{root.Struct()}, nil
}

func (s EmailAttachment) ContentType() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailAttachment) HasContentType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailAttachment) HasContentDisposition() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentDispositionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailAttachment) HasContentId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return nil, err
	}

	return []byte(p.Data()), nil

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
	if err != nil {
		return EmailAttachment_List{}, err
	}
	return EmailAttachment_List{l}, nil
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
	if err != nil {
		return EmailMessage{}, err
	}
	return EmailMessage{st}, nil
}

func NewRootEmailMessage(s *capnp.Segment) (EmailMessage, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12})
	if err != nil {
		return EmailMessage{}, err
	}
	return EmailMessage{st}, nil
}

func ReadRootEmailMessage(msg *capnp.Message) (EmailMessage, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return EmailMessage{}, err
	}
	return EmailMessage{root.Struct()}, nil
}

func (s EmailMessage) Date() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s EmailMessage) SetDate(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s EmailMessage) From() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return EmailAddress{}, err
	}

	return EmailAddress{Struct: p.Struct()}, nil

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
	if err != nil {
		return EmailAddress_List{}, err
	}

	return EmailAddress_List{List: p.List()}, nil

}

func (s EmailMessage) HasTo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetTo(v EmailAddress_List) error {

	return s.Struct.SetPtr(1, v.List.ToPtr())
}

func (s EmailMessage) Cc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return EmailAddress_List{}, err
	}

	return EmailAddress_List{List: p.List()}, nil

}

func (s EmailMessage) HasCc() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetCc(v EmailAddress_List) error {

	return s.Struct.SetPtr(2, v.List.ToPtr())
}

func (s EmailMessage) Bcc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return EmailAddress_List{}, err
	}

	return EmailAddress_List{List: p.List()}, nil

}

func (s EmailMessage) HasBcc() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetBcc(v EmailAddress_List) error {

	return s.Struct.SetPtr(3, v.List.ToPtr())
}

func (s EmailMessage) ReplyTo() (EmailAddress, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return EmailAddress{}, err
	}

	return EmailAddress{Struct: p.Struct()}, nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailMessage) HasMessageId() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s EmailMessage) MessageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return capnp.TextList{}, err
	}

	return capnp.TextList{List: p.List()}, nil

}

func (s EmailMessage) HasReferences() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetReferences(v capnp.TextList) error {

	return s.Struct.SetPtr(6, v.List.ToPtr())
}

func (s EmailMessage) InReplyTo() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(7)
	if err != nil {
		return capnp.TextList{}, err
	}

	return capnp.TextList{List: p.List()}, nil

}

func (s EmailMessage) HasInReplyTo() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetInReplyTo(v capnp.TextList) error {

	return s.Struct.SetPtr(7, v.List.ToPtr())
}

func (s EmailMessage) Subject() (string, error) {
	p, err := s.Struct.Ptr(8)
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailMessage) HasSubject() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SubjectBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(8)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailMessage) HasText() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s EmailMessage) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(9)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return "", err
	}

	return p.Text(), nil

}

func (s EmailMessage) HasHtml() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s EmailMessage) HtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(10)
	if err != nil {
		return nil, err
	}

	return p.Data(), nil

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
	if err != nil {
		return EmailAttachment_List{}, err
	}

	return EmailAttachment_List{List: p.List()}, nil

}

func (s EmailMessage) HasAttachments() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetAttachments(v EmailAttachment_List) error {

	return s.Struct.SetPtr(11, v.List.ToPtr())
}

// EmailMessage_List is a list of EmailMessage.
type EmailMessage_List struct{ capnp.List }

// NewEmailMessage creates a new list of EmailMessage.
func NewEmailMessage_List(s *capnp.Segment, sz int32) (EmailMessage_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12}, sz)
	if err != nil {
		return EmailMessage_List{}, err
	}
	return EmailMessage_List{l}, nil
}

func (s EmailMessage_List) At(i int) EmailMessage           { return EmailMessage{s.List.Struct(i)} }
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

type EmailSendPort_Server interface {
	Send(EmailSendPort_send) error
}

func EmailSendPort_ServerToClient(s EmailSendPort_Server) EmailSendPort {
	c, _ := s.(server.Closer)
	return EmailSendPort{Client: server.New(EmailSendPort_Methods(nil, s), c)}
}

func EmailSendPort_Methods(methods []server.Method, s EmailSendPort_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
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

	return methods
}

// EmailSendPort_send holds the arguments for a server call to EmailSendPort.send.
type EmailSendPort_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailSendPort_send_Params
	Results EmailSendPort_send_Results
}

type EmailSendPort_send_Params struct{ capnp.Struct }

func NewEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailSendPort_send_Params{}, err
	}
	return EmailSendPort_send_Params{st}, nil
}

func NewRootEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailSendPort_send_Params{}, err
	}
	return EmailSendPort_send_Params{st}, nil
}

func ReadRootEmailSendPort_send_Params(msg *capnp.Message) (EmailSendPort_send_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return EmailSendPort_send_Params{}, err
	}
	return EmailSendPort_send_Params{root.Struct()}, nil
}

func (s EmailSendPort_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return EmailMessage{}, err
	}

	return EmailMessage{Struct: p.Struct()}, nil

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
	if err != nil {
		return EmailSendPort_send_Params_List{}, err
	}
	return EmailSendPort_send_Params_List{l}, nil
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
	if err != nil {
		return EmailSendPort_send_Results{}, err
	}
	return EmailSendPort_send_Results{st}, nil
}

func NewRootEmailSendPort_send_Results(s *capnp.Segment) (EmailSendPort_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return EmailSendPort_send_Results{}, err
	}
	return EmailSendPort_send_Results{st}, nil
}

func ReadRootEmailSendPort_send_Results(msg *capnp.Message) (EmailSendPort_send_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return EmailSendPort_send_Results{}, err
	}
	return EmailSendPort_send_Results{root.Struct()}, nil
}

// EmailSendPort_send_Results_List is a list of EmailSendPort_send_Results.
type EmailSendPort_send_Results_List struct{ capnp.List }

// NewEmailSendPort_send_Results creates a new list of EmailSendPort_send_Results.
func NewEmailSendPort_send_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return EmailSendPort_send_Results_List{}, err
	}
	return EmailSendPort_send_Results_List{l}, nil
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
