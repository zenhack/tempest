package email

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	util "zenhack.net/go/sandstorm/capnp/util"
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
	root, err := msg.Root()
	if err != nil {
		return EmailAddress{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAddress{st}, nil
}

func (s EmailAddress) Address() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailAddress) SetAddress(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s EmailAddress) Name() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailAddress) SetName(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
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
	root, err := msg.Root()
	if err != nil {
		return EmailAttachment{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAttachment{st}, nil
}

func (s EmailAttachment) ContentType() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailAttachment) SetContentType(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s EmailAttachment) ContentDisposition() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailAttachment) SetContentDisposition(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s EmailAttachment) ContentId() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailAttachment) SetContentId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

func (s EmailAttachment) Content() ([]byte, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s EmailAttachment) SetContent(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(3, d)
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
	root, err := msg.Root()
	if err != nil {
		return EmailMessage{}, err
	}
	st := capnp.ToStruct(root)
	return EmailMessage{st}, nil
}

func (s EmailMessage) Date() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s EmailMessage) SetDate(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

func (s EmailMessage) From() (EmailAddress, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EmailAddress{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailAddress{Struct: ss}, nil
}

func (s EmailMessage) SetFrom(v EmailAddress) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewFrom sets the from field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewFrom() (EmailAddress, error) {

	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s EmailMessage) To() (EmailAddress_List, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return EmailAddress_List{}, err
	}

	l := capnp.ToList(p)

	return EmailAddress_List{List: l}, nil
}

func (s EmailMessage) SetTo(v EmailAddress_List) error {

	return s.Struct.SetPointer(1, v.List)
}

func (s EmailMessage) Cc() (EmailAddress_List, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return EmailAddress_List{}, err
	}

	l := capnp.ToList(p)

	return EmailAddress_List{List: l}, nil
}

func (s EmailMessage) SetCc(v EmailAddress_List) error {

	return s.Struct.SetPointer(2, v.List)
}

func (s EmailMessage) Bcc() (EmailAddress_List, error) {
	p, err := s.Struct.Pointer(3)
	if err != nil {
		return EmailAddress_List{}, err
	}

	l := capnp.ToList(p)

	return EmailAddress_List{List: l}, nil
}

func (s EmailMessage) SetBcc(v EmailAddress_List) error {

	return s.Struct.SetPointer(3, v.List)
}

func (s EmailMessage) ReplyTo() (EmailAddress, error) {
	p, err := s.Struct.Pointer(4)
	if err != nil {
		return EmailAddress{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailAddress{Struct: ss}, nil
}

func (s EmailMessage) SetReplyTo(v EmailAddress) error {

	return s.Struct.SetPointer(4, v.Struct)
}

// NewReplyTo sets the replyTo field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewReplyTo() (EmailAddress, error) {

	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPointer(4, ss)
	return ss, err
}

func (s EmailMessage) MessageId() (string, error) {
	p, err := s.Struct.Pointer(5)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailMessage) SetMessageId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(5, t)
}

func (s EmailMessage) References() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(6)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s EmailMessage) SetReferences(v capnp.TextList) error {

	return s.Struct.SetPointer(6, v.List)
}

func (s EmailMessage) InReplyTo() (capnp.TextList, error) {
	p, err := s.Struct.Pointer(7)
	if err != nil {
		return capnp.TextList{}, err
	}

	l := capnp.ToList(p)

	return capnp.TextList{List: l}, nil
}

func (s EmailMessage) SetInReplyTo(v capnp.TextList) error {

	return s.Struct.SetPointer(7, v.List)
}

func (s EmailMessage) Subject() (string, error) {
	p, err := s.Struct.Pointer(8)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailMessage) SetSubject(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(8, t)
}

func (s EmailMessage) Text() (string, error) {
	p, err := s.Struct.Pointer(9)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailMessage) SetText(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(9, t)
}

func (s EmailMessage) Html() (string, error) {
	p, err := s.Struct.Pointer(10)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailMessage) SetHtml(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(10, t)
}

func (s EmailMessage) Attachments() (EmailAttachment_List, error) {
	p, err := s.Struct.Pointer(11)
	if err != nil {
		return EmailAttachment_List{}, err
	}

	l := capnp.ToList(p)

	return EmailAttachment_List{List: l}, nil
}

func (s EmailMessage) SetAttachments(v EmailAttachment_List) error {

	return s.Struct.SetPointer(11, v.List)
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
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	return EmailSendPort_PowerboxTag{st}, nil
}

func NewRootEmailSendPort_PowerboxTag(s *capnp.Segment) (EmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	return EmailSendPort_PowerboxTag{st}, nil
}

func ReadRootEmailSendPort_PowerboxTag(msg *capnp.Message) (EmailSendPort_PowerboxTag, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	st := capnp.ToStruct(root)
	return EmailSendPort_PowerboxTag{st}, nil
}

func (s EmailSendPort_PowerboxTag) FromHint() (EmailAddress, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EmailAddress{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailAddress{Struct: ss}, nil
}

func (s EmailSendPort_PowerboxTag) SetFromHint(v EmailAddress) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewFromHint sets the fromHint field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_PowerboxTag) NewFromHint() (EmailAddress, error) {

	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s EmailSendPort_PowerboxTag) ListIdHint() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailSendPort_PowerboxTag) SetListIdHint(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

// EmailSendPort_PowerboxTag_List is a list of EmailSendPort_PowerboxTag.
type EmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewEmailSendPort_PowerboxTag creates a new list of EmailSendPort_PowerboxTag.
func NewEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (EmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return EmailSendPort_PowerboxTag_List{}, err
	}
	return EmailSendPort_PowerboxTag_List{l}, nil
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
	root, err := msg.Root()
	if err != nil {
		return EmailSendPort_send_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailSendPort_send_Params{st}, nil
}

func (s EmailSendPort_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EmailMessage{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailMessage{Struct: ss}, nil
}

func (s EmailSendPort_send_Params) SetEmail(v EmailMessage) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailSendPort_send_Params) NewEmail() (EmailMessage, error) {

	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPointer(0, ss)
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
	root, err := msg.Root()
	if err != nil {
		return EmailSendPort_send_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailSendPort_send_Results{st}, nil
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

type EmailSendPort_hintAddress_Params struct{ capnp.Struct }

func NewEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailSendPort_hintAddress_Params{}, err
	}
	return EmailSendPort_hintAddress_Params{st}, nil
}

func NewRootEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailSendPort_hintAddress_Params{}, err
	}
	return EmailSendPort_hintAddress_Params{st}, nil
}

func ReadRootEmailSendPort_hintAddress_Params(msg *capnp.Message) (EmailSendPort_hintAddress_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailSendPort_hintAddress_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailSendPort_hintAddress_Params{st}, nil
}

func (s EmailSendPort_hintAddress_Params) Address() (EmailAddress, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EmailAddress{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailAddress{Struct: ss}, nil
}

func (s EmailSendPort_hintAddress_Params) SetAddress(v EmailAddress) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewAddress sets the address field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_hintAddress_Params) NewAddress() (EmailAddress, error) {

	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// EmailSendPort_hintAddress_Params_List is a list of EmailSendPort_hintAddress_Params.
type EmailSendPort_hintAddress_Params_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Params creates a new list of EmailSendPort_hintAddress_Params.
func NewEmailSendPort_hintAddress_Params_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailSendPort_hintAddress_Params_List{}, err
	}
	return EmailSendPort_hintAddress_Params_List{l}, nil
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
	if err != nil {
		return EmailSendPort_hintAddress_Results{}, err
	}
	return EmailSendPort_hintAddress_Results{st}, nil
}

func NewRootEmailSendPort_hintAddress_Results(s *capnp.Segment) (EmailSendPort_hintAddress_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return EmailSendPort_hintAddress_Results{}, err
	}
	return EmailSendPort_hintAddress_Results{st}, nil
}

func ReadRootEmailSendPort_hintAddress_Results(msg *capnp.Message) (EmailSendPort_hintAddress_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailSendPort_hintAddress_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailSendPort_hintAddress_Results{st}, nil
}

// EmailSendPort_hintAddress_Results_List is a list of EmailSendPort_hintAddress_Results.
type EmailSendPort_hintAddress_Results_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Results creates a new list of EmailSendPort_hintAddress_Results.
func NewEmailSendPort_hintAddress_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return EmailSendPort_hintAddress_Results_List{}, err
	}
	return EmailSendPort_hintAddress_Results_List{l}, nil
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
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	return VerifiedEmail_PowerboxTag{st}, nil
}

func NewRootVerifiedEmail_PowerboxTag(s *capnp.Segment) (VerifiedEmail_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	return VerifiedEmail_PowerboxTag{st}, nil
}

func ReadRootVerifiedEmail_PowerboxTag(msg *capnp.Message) (VerifiedEmail_PowerboxTag, error) {
	root, err := msg.Root()
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	st := capnp.ToStruct(root)
	return VerifiedEmail_PowerboxTag{st}, nil
}

func (s VerifiedEmail_PowerboxTag) VerifierId() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s VerifiedEmail_PowerboxTag) SetVerifierId(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s VerifiedEmail_PowerboxTag) Address() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s VerifiedEmail_PowerboxTag) SetAddress(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s VerifiedEmail_PowerboxTag) Domain() (string, error) {
	p, err := s.Struct.Pointer(2)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s VerifiedEmail_PowerboxTag) SetDomain(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(2, t)
}

// VerifiedEmail_PowerboxTag_List is a list of VerifiedEmail_PowerboxTag.
type VerifiedEmail_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmail_PowerboxTag creates a new list of VerifiedEmail_PowerboxTag.
func NewVerifiedEmail_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmail_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	if err != nil {
		return VerifiedEmail_PowerboxTag_List{}, err
	}
	return VerifiedEmail_PowerboxTag_List{l}, nil
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
	if err != nil {
		return VerifiedEmailSendPort_PowerboxTag{}, err
	}
	return VerifiedEmailSendPort_PowerboxTag{st}, nil
}

func NewRootVerifiedEmailSendPort_PowerboxTag(s *capnp.Segment) (VerifiedEmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return VerifiedEmailSendPort_PowerboxTag{}, err
	}
	return VerifiedEmailSendPort_PowerboxTag{st}, nil
}

func ReadRootVerifiedEmailSendPort_PowerboxTag(msg *capnp.Message) (VerifiedEmailSendPort_PowerboxTag, error) {
	root, err := msg.Root()
	if err != nil {
		return VerifiedEmailSendPort_PowerboxTag{}, err
	}
	st := capnp.ToStruct(root)
	return VerifiedEmailSendPort_PowerboxTag{st}, nil
}

func (s VerifiedEmailSendPort_PowerboxTag) Verification() (VerifiedEmail_PowerboxTag, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}

	ss := capnp.ToStruct(p)

	return VerifiedEmail_PowerboxTag{Struct: ss}, nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetVerification(v VerifiedEmail_PowerboxTag) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewVerification sets the verification field to a newly
// allocated VerifiedEmail_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewVerification() (VerifiedEmail_PowerboxTag, error) {

	ss, err := NewVerifiedEmail_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

func (s VerifiedEmailSendPort_PowerboxTag) Port() (EmailSendPort_PowerboxTag, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailSendPort_PowerboxTag{Struct: ss}, nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetPort(v EmailSendPort_PowerboxTag) error {

	return s.Struct.SetPointer(1, v.Struct)
}

// NewPort sets the port field to a newly
// allocated EmailSendPort_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewPort() (EmailSendPort_PowerboxTag, error) {

	ss, err := NewEmailSendPort_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// VerifiedEmailSendPort_PowerboxTag_List is a list of VerifiedEmailSendPort_PowerboxTag.
type VerifiedEmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmailSendPort_PowerboxTag creates a new list of VerifiedEmailSendPort_PowerboxTag.
func NewVerifiedEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return VerifiedEmailSendPort_PowerboxTag_List{}, err
	}
	return VerifiedEmailSendPort_PowerboxTag_List{l}, nil
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
	if err != nil {
		return EmailVerifier_getId_Params{}, err
	}
	return EmailVerifier_getId_Params{st}, nil
}

func NewRootEmailVerifier_getId_Params(s *capnp.Segment) (EmailVerifier_getId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return EmailVerifier_getId_Params{}, err
	}
	return EmailVerifier_getId_Params{st}, nil
}

func ReadRootEmailVerifier_getId_Params(msg *capnp.Message) (EmailVerifier_getId_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailVerifier_getId_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailVerifier_getId_Params{st}, nil
}

// EmailVerifier_getId_Params_List is a list of EmailVerifier_getId_Params.
type EmailVerifier_getId_Params_List struct{ capnp.List }

// NewEmailVerifier_getId_Params creates a new list of EmailVerifier_getId_Params.
func NewEmailVerifier_getId_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return EmailVerifier_getId_Params_List{}, err
	}
	return EmailVerifier_getId_Params_List{l}, nil
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
	if err != nil {
		return EmailVerifier_getId_Results{}, err
	}
	return EmailVerifier_getId_Results{st}, nil
}

func NewRootEmailVerifier_getId_Results(s *capnp.Segment) (EmailVerifier_getId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailVerifier_getId_Results{}, err
	}
	return EmailVerifier_getId_Results{st}, nil
}

func ReadRootEmailVerifier_getId_Results(msg *capnp.Message) (EmailVerifier_getId_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailVerifier_getId_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailVerifier_getId_Results{st}, nil
}

func (s EmailVerifier_getId_Results) Id() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s EmailVerifier_getId_Results) SetId(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

// EmailVerifier_getId_Results_List is a list of EmailVerifier_getId_Results.
type EmailVerifier_getId_Results_List struct{ capnp.List }

// NewEmailVerifier_getId_Results creates a new list of EmailVerifier_getId_Results.
func NewEmailVerifier_getId_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailVerifier_getId_Results_List{}, err
	}
	return EmailVerifier_getId_Results_List{l}, nil
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
	if err != nil {
		return EmailVerifier_verifyEmail_Params{}, err
	}
	return EmailVerifier_verifyEmail_Params{st}, nil
}

func NewRootEmailVerifier_verifyEmail_Params(s *capnp.Segment) (EmailVerifier_verifyEmail_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return EmailVerifier_verifyEmail_Params{}, err
	}
	return EmailVerifier_verifyEmail_Params{st}, nil
}

func ReadRootEmailVerifier_verifyEmail_Params(msg *capnp.Message) (EmailVerifier_verifyEmail_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailVerifier_verifyEmail_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailVerifier_verifyEmail_Params{st}, nil
}

func (s EmailVerifier_verifyEmail_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return nil, err
	}

	return []byte(capnp.ToData(p)), nil

}

func (s EmailVerifier_verifyEmail_Params) SetTabId(v []byte) error {

	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, d)
}

func (s EmailVerifier_verifyEmail_Params) Verification() VerifiedEmail {
	p, err := s.Struct.Pointer(1)
	if err != nil {

		return VerifiedEmail{}
	}
	c := capnp.ToInterface(p).Client()
	return VerifiedEmail{Client: c}
}

func (s EmailVerifier_verifyEmail_Params) SetVerification(v VerifiedEmail) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(1, in)
}

// EmailVerifier_verifyEmail_Params_List is a list of EmailVerifier_verifyEmail_Params.
type EmailVerifier_verifyEmail_Params_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Params creates a new list of EmailVerifier_verifyEmail_Params.
func NewEmailVerifier_verifyEmail_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return EmailVerifier_verifyEmail_Params_List{}, err
	}
	return EmailVerifier_verifyEmail_Params_List{l}, nil
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
	if err != nil {
		return EmailVerifier_verifyEmail_Results{}, err
	}
	return EmailVerifier_verifyEmail_Results{st}, nil
}

func NewRootEmailVerifier_verifyEmail_Results(s *capnp.Segment) (EmailVerifier_verifyEmail_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailVerifier_verifyEmail_Results{}, err
	}
	return EmailVerifier_verifyEmail_Results{st}, nil
}

func ReadRootEmailVerifier_verifyEmail_Results(msg *capnp.Message) (EmailVerifier_verifyEmail_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailVerifier_verifyEmail_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailVerifier_verifyEmail_Results{st}, nil
}

func (s EmailVerifier_verifyEmail_Results) Address() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s EmailVerifier_verifyEmail_Results) SetAddress(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// EmailVerifier_verifyEmail_Results_List is a list of EmailVerifier_verifyEmail_Results.
type EmailVerifier_verifyEmail_Results_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Results creates a new list of EmailVerifier_verifyEmail_Results.
func NewEmailVerifier_verifyEmail_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailVerifier_verifyEmail_Results_List{}, err
	}
	return EmailVerifier_verifyEmail_Results_List{l}, nil
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
	if err != nil {
		return EmailAgent_send_Params{}, err
	}
	return EmailAgent_send_Params{st}, nil
}

func NewRootEmailAgent_send_Params(s *capnp.Segment) (EmailAgent_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailAgent_send_Params{}, err
	}
	return EmailAgent_send_Params{st}, nil
}

func ReadRootEmailAgent_send_Params(msg *capnp.Message) (EmailAgent_send_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailAgent_send_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAgent_send_Params{st}, nil
}

func (s EmailAgent_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return EmailMessage{}, err
	}

	ss := capnp.ToStruct(p)

	return EmailMessage{Struct: ss}, nil
}

func (s EmailAgent_send_Params) SetEmail(v EmailMessage) error {

	return s.Struct.SetPointer(0, v.Struct)
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailAgent_send_Params) NewEmail() (EmailMessage, error) {

	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPointer(0, ss)
	return ss, err
}

// EmailAgent_send_Params_List is a list of EmailAgent_send_Params.
type EmailAgent_send_Params_List struct{ capnp.List }

// NewEmailAgent_send_Params creates a new list of EmailAgent_send_Params.
func NewEmailAgent_send_Params_List(s *capnp.Segment, sz int32) (EmailAgent_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailAgent_send_Params_List{}, err
	}
	return EmailAgent_send_Params_List{l}, nil
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
	if err != nil {
		return EmailAgent_send_Results{}, err
	}
	return EmailAgent_send_Results{st}, nil
}

func NewRootEmailAgent_send_Results(s *capnp.Segment) (EmailAgent_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return EmailAgent_send_Results{}, err
	}
	return EmailAgent_send_Results{st}, nil
}

func ReadRootEmailAgent_send_Results(msg *capnp.Message) (EmailAgent_send_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailAgent_send_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAgent_send_Results{st}, nil
}

// EmailAgent_send_Results_List is a list of EmailAgent_send_Results.
type EmailAgent_send_Results_List struct{ capnp.List }

// NewEmailAgent_send_Results creates a new list of EmailAgent_send_Results.
func NewEmailAgent_send_Results_List(s *capnp.Segment, sz int32) (EmailAgent_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return EmailAgent_send_Results_List{}, err
	}
	return EmailAgent_send_Results_List{l}, nil
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
	if err != nil {
		return EmailAgent_addReceiver_Params{}, err
	}
	return EmailAgent_addReceiver_Params{st}, nil
}

func NewRootEmailAgent_addReceiver_Params(s *capnp.Segment) (EmailAgent_addReceiver_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailAgent_addReceiver_Params{}, err
	}
	return EmailAgent_addReceiver_Params{st}, nil
}

func ReadRootEmailAgent_addReceiver_Params(msg *capnp.Message) (EmailAgent_addReceiver_Params, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailAgent_addReceiver_Params{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAgent_addReceiver_Params{st}, nil
}

func (s EmailAgent_addReceiver_Params) Port() EmailSendPort {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return EmailSendPort{}
	}
	c := capnp.ToInterface(p).Client()
	return EmailSendPort{Client: c}
}

func (s EmailAgent_addReceiver_Params) SetPort(v EmailSendPort) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(0, in)
}

// EmailAgent_addReceiver_Params_List is a list of EmailAgent_addReceiver_Params.
type EmailAgent_addReceiver_Params_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Params creates a new list of EmailAgent_addReceiver_Params.
func NewEmailAgent_addReceiver_Params_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailAgent_addReceiver_Params_List{}, err
	}
	return EmailAgent_addReceiver_Params_List{l}, nil
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
	if err != nil {
		return EmailAgent_addReceiver_Results{}, err
	}
	return EmailAgent_addReceiver_Results{st}, nil
}

func NewRootEmailAgent_addReceiver_Results(s *capnp.Segment) (EmailAgent_addReceiver_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return EmailAgent_addReceiver_Results{}, err
	}
	return EmailAgent_addReceiver_Results{st}, nil
}

func ReadRootEmailAgent_addReceiver_Results(msg *capnp.Message) (EmailAgent_addReceiver_Results, error) {
	root, err := msg.Root()
	if err != nil {
		return EmailAgent_addReceiver_Results{}, err
	}
	st := capnp.ToStruct(root)
	return EmailAgent_addReceiver_Results{st}, nil
}

func (s EmailAgent_addReceiver_Results) Handle() util.Handle {
	p, err := s.Struct.Pointer(0)
	if err != nil {

		return util.Handle{}
	}
	c := capnp.ToInterface(p).Client()
	return util.Handle{Client: c}
}

func (s EmailAgent_addReceiver_Results) SetHandle(v util.Handle) error {

	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPointer(0, in)
}

// EmailAgent_addReceiver_Results_List is a list of EmailAgent_addReceiver_Results.
type EmailAgent_addReceiver_Results_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Results creates a new list of EmailAgent_addReceiver_Results.
func NewEmailAgent_addReceiver_Results_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return EmailAgent_addReceiver_Results_List{}, err
	}
	return EmailAgent_addReceiver_Results_List{l}, nil
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
