package util

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type KeyValue struct{ capnp.Struct }

func NewKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyValue{st}, err
}

func NewRootKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyValue{st}, err
}

func ReadRootKeyValue(msg *capnp.Message) (KeyValue, error) {
	root, err := msg.RootPtr()
	return KeyValue{root.Struct()}, err
}

func (s KeyValue) String() string {
	str, _ := text.Marshal(0x94a081e4abb13424, s.Struct)
	return str
}

func (s KeyValue) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s KeyValue) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s KeyValue) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s KeyValue) SetKey(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s KeyValue) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s KeyValue) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s KeyValue) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s KeyValue) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// KeyValue_List is a list of KeyValue.
type KeyValue_List struct{ capnp.List }

// NewKeyValue creates a new list of KeyValue.
func NewKeyValue_List(s *capnp.Segment, sz int32) (KeyValue_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return KeyValue_List{l}, err
}

func (s KeyValue_List) At(i int) KeyValue { return KeyValue{s.List.Struct(i)} }

func (s KeyValue_List) Set(i int, v KeyValue) error { return s.List.SetStruct(i, v.Struct) }

// KeyValue_Promise is a wrapper for a KeyValue promised by a client call.
type KeyValue_Promise struct{ *capnp.Pipeline }

func (p KeyValue_Promise) Struct() (KeyValue, error) {
	s, err := p.Pipeline.Struct()
	return KeyValue{s}, err
}

type LocalizedText struct{ capnp.Struct }

func NewLocalizedText(s *capnp.Segment) (LocalizedText, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText{st}, err
}

func NewRootLocalizedText(s *capnp.Segment) (LocalizedText, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText{st}, err
}

func ReadRootLocalizedText(msg *capnp.Message) (LocalizedText, error) {
	root, err := msg.RootPtr()
	return LocalizedText{root.Struct()}, err
}

func (s LocalizedText) String() string {
	str, _ := text.Marshal(0x8b5db772377be249, s.Struct)
	return str
}

func (s LocalizedText) DefaultText() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s LocalizedText) HasDefaultText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText) DefaultTextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s LocalizedText) SetDefaultText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s LocalizedText) Localizations() (LocalizedText_Localization_List, error) {
	p, err := s.Struct.Ptr(1)
	return LocalizedText_Localization_List{List: p.List()}, err
}

func (s LocalizedText) HasLocalizations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LocalizedText) SetLocalizations(v LocalizedText_Localization_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewLocalizations sets the localizations field to a newly
// allocated LocalizedText_Localization_List, preferring placement in s's segment.
func (s LocalizedText) NewLocalizations(n int32) (LocalizedText_Localization_List, error) {
	l, err := NewLocalizedText_Localization_List(s.Struct.Segment(), n)
	if err != nil {
		return LocalizedText_Localization_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// LocalizedText_List is a list of LocalizedText.
type LocalizedText_List struct{ capnp.List }

// NewLocalizedText creates a new list of LocalizedText.
func NewLocalizedText_List(s *capnp.Segment, sz int32) (LocalizedText_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return LocalizedText_List{l}, err
}

func (s LocalizedText_List) At(i int) LocalizedText { return LocalizedText{s.List.Struct(i)} }

func (s LocalizedText_List) Set(i int, v LocalizedText) error { return s.List.SetStruct(i, v.Struct) }

// LocalizedText_Promise is a wrapper for a LocalizedText promised by a client call.
type LocalizedText_Promise struct{ *capnp.Pipeline }

func (p LocalizedText_Promise) Struct() (LocalizedText, error) {
	s, err := p.Pipeline.Struct()
	return LocalizedText{s}, err
}

type LocalizedText_Localization struct{ capnp.Struct }

func NewLocalizedText_Localization(s *capnp.Segment) (LocalizedText_Localization, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText_Localization{st}, err
}

func NewRootLocalizedText_Localization(s *capnp.Segment) (LocalizedText_Localization, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText_Localization{st}, err
}

func ReadRootLocalizedText_Localization(msg *capnp.Message) (LocalizedText_Localization, error) {
	root, err := msg.RootPtr()
	return LocalizedText_Localization{root.Struct()}, err
}

func (s LocalizedText_Localization) String() string {
	str, _ := text.Marshal(0xa4f5ae06dd1b7791, s.Struct)
	return str
}

func (s LocalizedText_Localization) Locale() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s LocalizedText_Localization) HasLocale() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) LocaleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s LocalizedText_Localization) SetLocale(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s LocalizedText_Localization) Text() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s LocalizedText_Localization) HasText() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s LocalizedText_Localization) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// LocalizedText_Localization_List is a list of LocalizedText_Localization.
type LocalizedText_Localization_List struct{ capnp.List }

// NewLocalizedText_Localization creates a new list of LocalizedText_Localization.
func NewLocalizedText_Localization_List(s *capnp.Segment, sz int32) (LocalizedText_Localization_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return LocalizedText_Localization_List{l}, err
}

func (s LocalizedText_Localization_List) At(i int) LocalizedText_Localization {
	return LocalizedText_Localization{s.List.Struct(i)}
}

func (s LocalizedText_Localization_List) Set(i int, v LocalizedText_Localization) error {
	return s.List.SetStruct(i, v.Struct)
}

// LocalizedText_Localization_Promise is a wrapper for a LocalizedText_Localization promised by a client call.
type LocalizedText_Localization_Promise struct{ *capnp.Pipeline }

func (p LocalizedText_Localization_Promise) Struct() (LocalizedText_Localization, error) {
	s, err := p.Pipeline.Struct()
	return LocalizedText_Localization{s}, err
}

type Handle struct{ Client capnp.Client }

type Handle_Server interface {
}

func Handle_ServerToClient(s Handle_Server) Handle {
	c, _ := s.(server.Closer)
	return Handle{Client: server.New(Handle_Methods(nil, s), c)}
}

func Handle_Methods(methods []server.Method, s Handle_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type ByteStream struct{ Client capnp.Client }

func (c ByteStream) Write(ctx context.Context, params func(ByteStream_write_Params) error, opts ...capnp.CallOption) ByteStream_write_Results_Promise {
	if c.Client == nil {
		return ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_write_Params{Struct: s}) }
	}
	return ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ByteStream) Done(ctx context.Context, params func(ByteStream_done_Params) error, opts ...capnp.CallOption) ByteStream_done_Results_Promise {
	if c.Client == nil {
		return ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_done_Params{Struct: s}) }
	}
	return ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ByteStream) ExpectSize(ctx context.Context, params func(ByteStream_expectSize_Params) error, opts ...capnp.CallOption) ByteStream_expectSize_Results_Promise {
	if c.Client == nil {
		return ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_expectSize_Params{Struct: s}) }
	}
	return ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ByteStream_Server interface {
	Write(ByteStream_write) error

	Done(ByteStream_done) error

	ExpectSize(ByteStream_expectSize) error
}

func ByteStream_ServerToClient(s ByteStream_Server) ByteStream {
	c, _ := s.(server.Closer)
	return ByteStream{Client: server.New(ByteStream_Methods(nil, s), c)}
}

func ByteStream_Methods(methods []server.Method, s ByteStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_write{c, opts, ByteStream_write_Params{Struct: p}, ByteStream_write_Results{Struct: r}}
			return s.Write(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_done{c, opts, ByteStream_done_Params{Struct: p}, ByteStream_done_Results{Struct: r}}
			return s.Done(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_expectSize{c, opts, ByteStream_expectSize_Params{Struct: p}, ByteStream_expectSize_Results{Struct: r}}
			return s.ExpectSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// ByteStream_write holds the arguments for a server call to ByteStream.write.
type ByteStream_write struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_write_Params
	Results ByteStream_write_Results
}

// ByteStream_done holds the arguments for a server call to ByteStream.done.
type ByteStream_done struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_done_Params
	Results ByteStream_done_Results
}

// ByteStream_expectSize holds the arguments for a server call to ByteStream.expectSize.
type ByteStream_expectSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_expectSize_Params
	Results ByteStream_expectSize_Results
}

type ByteStream_write_Params struct{ capnp.Struct }

func NewByteStream_write_Params(s *capnp.Segment) (ByteStream_write_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ByteStream_write_Params{st}, err
}

func NewRootByteStream_write_Params(s *capnp.Segment) (ByteStream_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ByteStream_write_Params{st}, err
}

func ReadRootByteStream_write_Params(msg *capnp.Message) (ByteStream_write_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_write_Params{root.Struct()}, err
}

func (s ByteStream_write_Params) String() string {
	str, _ := text.Marshal(0x97ed122121126ff2, s.Struct)
	return str
}

func (s ByteStream_write_Params) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s ByteStream_write_Params) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ByteStream_write_Params) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// ByteStream_write_Params_List is a list of ByteStream_write_Params.
type ByteStream_write_Params_List struct{ capnp.List }

// NewByteStream_write_Params creates a new list of ByteStream_write_Params.
func NewByteStream_write_Params_List(s *capnp.Segment, sz int32) (ByteStream_write_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ByteStream_write_Params_List{l}, err
}

func (s ByteStream_write_Params_List) At(i int) ByteStream_write_Params {
	return ByteStream_write_Params{s.List.Struct(i)}
}

func (s ByteStream_write_Params_List) Set(i int, v ByteStream_write_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_write_Params_Promise is a wrapper for a ByteStream_write_Params promised by a client call.
type ByteStream_write_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_write_Params_Promise) Struct() (ByteStream_write_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_write_Params{s}, err
}

type ByteStream_write_Results struct{ capnp.Struct }

func NewByteStream_write_Results(s *capnp.Segment) (ByteStream_write_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_write_Results{st}, err
}

func NewRootByteStream_write_Results(s *capnp.Segment) (ByteStream_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_write_Results{st}, err
}

func ReadRootByteStream_write_Results(msg *capnp.Message) (ByteStream_write_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_write_Results{root.Struct()}, err
}

func (s ByteStream_write_Results) String() string {
	str, _ := text.Marshal(0xecde2a9c6f3f84c9, s.Struct)
	return str
}

// ByteStream_write_Results_List is a list of ByteStream_write_Results.
type ByteStream_write_Results_List struct{ capnp.List }

// NewByteStream_write_Results creates a new list of ByteStream_write_Results.
func NewByteStream_write_Results_List(s *capnp.Segment, sz int32) (ByteStream_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_write_Results_List{l}, err
}

func (s ByteStream_write_Results_List) At(i int) ByteStream_write_Results {
	return ByteStream_write_Results{s.List.Struct(i)}
}

func (s ByteStream_write_Results_List) Set(i int, v ByteStream_write_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_write_Results_Promise is a wrapper for a ByteStream_write_Results promised by a client call.
type ByteStream_write_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_write_Results_Promise) Struct() (ByteStream_write_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_write_Results{s}, err
}

type ByteStream_done_Params struct{ capnp.Struct }

func NewByteStream_done_Params(s *capnp.Segment) (ByteStream_done_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Params{st}, err
}

func NewRootByteStream_done_Params(s *capnp.Segment) (ByteStream_done_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Params{st}, err
}

func ReadRootByteStream_done_Params(msg *capnp.Message) (ByteStream_done_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_done_Params{root.Struct()}, err
}

func (s ByteStream_done_Params) String() string {
	str, _ := text.Marshal(0xbc1426493658b76e, s.Struct)
	return str
}

// ByteStream_done_Params_List is a list of ByteStream_done_Params.
type ByteStream_done_Params_List struct{ capnp.List }

// NewByteStream_done_Params creates a new list of ByteStream_done_Params.
func NewByteStream_done_Params_List(s *capnp.Segment, sz int32) (ByteStream_done_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_done_Params_List{l}, err
}

func (s ByteStream_done_Params_List) At(i int) ByteStream_done_Params {
	return ByteStream_done_Params{s.List.Struct(i)}
}

func (s ByteStream_done_Params_List) Set(i int, v ByteStream_done_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_done_Params_Promise is a wrapper for a ByteStream_done_Params promised by a client call.
type ByteStream_done_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_done_Params_Promise) Struct() (ByteStream_done_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_done_Params{s}, err
}

type ByteStream_done_Results struct{ capnp.Struct }

func NewByteStream_done_Results(s *capnp.Segment) (ByteStream_done_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Results{st}, err
}

func NewRootByteStream_done_Results(s *capnp.Segment) (ByteStream_done_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Results{st}, err
}

func ReadRootByteStream_done_Results(msg *capnp.Message) (ByteStream_done_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_done_Results{root.Struct()}, err
}

func (s ByteStream_done_Results) String() string {
	str, _ := text.Marshal(0xd0d8d935ee30b219, s.Struct)
	return str
}

// ByteStream_done_Results_List is a list of ByteStream_done_Results.
type ByteStream_done_Results_List struct{ capnp.List }

// NewByteStream_done_Results creates a new list of ByteStream_done_Results.
func NewByteStream_done_Results_List(s *capnp.Segment, sz int32) (ByteStream_done_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_done_Results_List{l}, err
}

func (s ByteStream_done_Results_List) At(i int) ByteStream_done_Results {
	return ByteStream_done_Results{s.List.Struct(i)}
}

func (s ByteStream_done_Results_List) Set(i int, v ByteStream_done_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_done_Results_Promise is a wrapper for a ByteStream_done_Results promised by a client call.
type ByteStream_done_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_done_Results_Promise) Struct() (ByteStream_done_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_done_Results{s}, err
}

type ByteStream_expectSize_Params struct{ capnp.Struct }

func NewByteStream_expectSize_Params(s *capnp.Segment) (ByteStream_expectSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ByteStream_expectSize_Params{st}, err
}

func NewRootByteStream_expectSize_Params(s *capnp.Segment) (ByteStream_expectSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ByteStream_expectSize_Params{st}, err
}

func ReadRootByteStream_expectSize_Params(msg *capnp.Message) (ByteStream_expectSize_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_expectSize_Params{root.Struct()}, err
}

func (s ByteStream_expectSize_Params) String() string {
	str, _ := text.Marshal(0x8c9a3c7674c761d3, s.Struct)
	return str
}

func (s ByteStream_expectSize_Params) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s ByteStream_expectSize_Params) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// ByteStream_expectSize_Params_List is a list of ByteStream_expectSize_Params.
type ByteStream_expectSize_Params_List struct{ capnp.List }

// NewByteStream_expectSize_Params creates a new list of ByteStream_expectSize_Params.
func NewByteStream_expectSize_Params_List(s *capnp.Segment, sz int32) (ByteStream_expectSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return ByteStream_expectSize_Params_List{l}, err
}

func (s ByteStream_expectSize_Params_List) At(i int) ByteStream_expectSize_Params {
	return ByteStream_expectSize_Params{s.List.Struct(i)}
}

func (s ByteStream_expectSize_Params_List) Set(i int, v ByteStream_expectSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_expectSize_Params_Promise is a wrapper for a ByteStream_expectSize_Params promised by a client call.
type ByteStream_expectSize_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_expectSize_Params_Promise) Struct() (ByteStream_expectSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_expectSize_Params{s}, err
}

type ByteStream_expectSize_Results struct{ capnp.Struct }

func NewByteStream_expectSize_Results(s *capnp.Segment) (ByteStream_expectSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_expectSize_Results{st}, err
}

func NewRootByteStream_expectSize_Results(s *capnp.Segment) (ByteStream_expectSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_expectSize_Results{st}, err
}

func ReadRootByteStream_expectSize_Results(msg *capnp.Message) (ByteStream_expectSize_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_expectSize_Results{root.Struct()}, err
}

func (s ByteStream_expectSize_Results) String() string {
	str, _ := text.Marshal(0xf35749d82a51479b, s.Struct)
	return str
}

// ByteStream_expectSize_Results_List is a list of ByteStream_expectSize_Results.
type ByteStream_expectSize_Results_List struct{ capnp.List }

// NewByteStream_expectSize_Results creates a new list of ByteStream_expectSize_Results.
func NewByteStream_expectSize_Results_List(s *capnp.Segment, sz int32) (ByteStream_expectSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_expectSize_Results_List{l}, err
}

func (s ByteStream_expectSize_Results_List) At(i int) ByteStream_expectSize_Results {
	return ByteStream_expectSize_Results{s.List.Struct(i)}
}

func (s ByteStream_expectSize_Results_List) Set(i int, v ByteStream_expectSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_expectSize_Results_Promise is a wrapper for a ByteStream_expectSize_Results promised by a client call.
type ByteStream_expectSize_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_expectSize_Results_Promise) Struct() (ByteStream_expectSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_expectSize_Results{s}, err
}

type Blob struct{ Client capnp.Client }

func (c Blob) GetSize(ctx context.Context, params func(Blob_getSize_Params) error, opts ...capnp.CallOption) Blob_getSize_Results_Promise {
	if c.Client == nil {
		return Blob_getSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      0,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_getSize_Params{Struct: s}) }
	}
	return Blob_getSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Blob) WriteTo(ctx context.Context, params func(Blob_writeTo_Params) error, opts ...capnp.CallOption) Blob_writeTo_Results_Promise {
	if c.Client == nil {
		return Blob_writeTo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      1,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "writeTo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_writeTo_Params{Struct: s}) }
	}
	return Blob_writeTo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Blob) GetSlice(ctx context.Context, params func(Blob_getSlice_Params) error, opts ...capnp.CallOption) Blob_getSlice_Results_Promise {
	if c.Client == nil {
		return Blob_getSlice_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      2,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSlice",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_getSlice_Params{Struct: s}) }
	}
	return Blob_getSlice_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Blob_Server interface {
	GetSize(Blob_getSize) error

	WriteTo(Blob_writeTo) error

	GetSlice(Blob_getSlice) error
}

func Blob_ServerToClient(s Blob_Server) Blob {
	c, _ := s.(server.Closer)
	return Blob{Client: server.New(Blob_Methods(nil, s), c)}
}

func Blob_Methods(methods []server.Method, s Blob_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      0,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_getSize{c, opts, Blob_getSize_Params{Struct: p}, Blob_getSize_Results{Struct: r}}
			return s.GetSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      1,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "writeTo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_writeTo{c, opts, Blob_writeTo_Params{Struct: p}, Blob_writeTo_Results{Struct: r}}
			return s.WriteTo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      2,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSlice",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_getSlice{c, opts, Blob_getSlice_Params{Struct: p}, Blob_getSlice_Results{Struct: r}}
			return s.GetSlice(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Blob_getSize holds the arguments for a server call to Blob.getSize.
type Blob_getSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_getSize_Params
	Results Blob_getSize_Results
}

// Blob_writeTo holds the arguments for a server call to Blob.writeTo.
type Blob_writeTo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_writeTo_Params
	Results Blob_writeTo_Results
}

// Blob_getSlice holds the arguments for a server call to Blob.getSlice.
type Blob_getSlice struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_getSlice_Params
	Results Blob_getSlice_Results
}

type Blob_getSize_Params struct{ capnp.Struct }

func NewBlob_getSize_Params(s *capnp.Segment) (Blob_getSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Blob_getSize_Params{st}, err
}

func NewRootBlob_getSize_Params(s *capnp.Segment) (Blob_getSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Blob_getSize_Params{st}, err
}

func ReadRootBlob_getSize_Params(msg *capnp.Message) (Blob_getSize_Params, error) {
	root, err := msg.RootPtr()
	return Blob_getSize_Params{root.Struct()}, err
}

func (s Blob_getSize_Params) String() string {
	str, _ := text.Marshal(0x8ee5f62e1fab915d, s.Struct)
	return str
}

// Blob_getSize_Params_List is a list of Blob_getSize_Params.
type Blob_getSize_Params_List struct{ capnp.List }

// NewBlob_getSize_Params creates a new list of Blob_getSize_Params.
func NewBlob_getSize_Params_List(s *capnp.Segment, sz int32) (Blob_getSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Blob_getSize_Params_List{l}, err
}

func (s Blob_getSize_Params_List) At(i int) Blob_getSize_Params {
	return Blob_getSize_Params{s.List.Struct(i)}
}

func (s Blob_getSize_Params_List) Set(i int, v Blob_getSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSize_Params_Promise is a wrapper for a Blob_getSize_Params promised by a client call.
type Blob_getSize_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_getSize_Params_Promise) Struct() (Blob_getSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSize_Params{s}, err
}

type Blob_getSize_Results struct{ capnp.Struct }

func NewBlob_getSize_Results(s *capnp.Segment) (Blob_getSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Blob_getSize_Results{st}, err
}

func NewRootBlob_getSize_Results(s *capnp.Segment) (Blob_getSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Blob_getSize_Results{st}, err
}

func ReadRootBlob_getSize_Results(msg *capnp.Message) (Blob_getSize_Results, error) {
	root, err := msg.RootPtr()
	return Blob_getSize_Results{root.Struct()}, err
}

func (s Blob_getSize_Results) String() string {
	str, _ := text.Marshal(0x8e48cb1497f3d6f4, s.Struct)
	return str
}

func (s Blob_getSize_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_getSize_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Blob_getSize_Results_List is a list of Blob_getSize_Results.
type Blob_getSize_Results_List struct{ capnp.List }

// NewBlob_getSize_Results creates a new list of Blob_getSize_Results.
func NewBlob_getSize_Results_List(s *capnp.Segment, sz int32) (Blob_getSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Blob_getSize_Results_List{l}, err
}

func (s Blob_getSize_Results_List) At(i int) Blob_getSize_Results {
	return Blob_getSize_Results{s.List.Struct(i)}
}

func (s Blob_getSize_Results_List) Set(i int, v Blob_getSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSize_Results_Promise is a wrapper for a Blob_getSize_Results promised by a client call.
type Blob_getSize_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_getSize_Results_Promise) Struct() (Blob_getSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSize_Results{s}, err
}

type Blob_writeTo_Params struct{ capnp.Struct }

func NewBlob_writeTo_Params(s *capnp.Segment) (Blob_writeTo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Blob_writeTo_Params{st}, err
}

func NewRootBlob_writeTo_Params(s *capnp.Segment) (Blob_writeTo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Blob_writeTo_Params{st}, err
}

func ReadRootBlob_writeTo_Params(msg *capnp.Message) (Blob_writeTo_Params, error) {
	root, err := msg.RootPtr()
	return Blob_writeTo_Params{root.Struct()}, err
}

func (s Blob_writeTo_Params) String() string {
	str, _ := text.Marshal(0x9f0719e9a9dccc4b, s.Struct)
	return str
}

func (s Blob_writeTo_Params) Stream() ByteStream {
	p, _ := s.Struct.Ptr(0)
	return ByteStream{Client: p.Interface().Client()}
}

func (s Blob_writeTo_Params) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_writeTo_Params) SetStream(v ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s Blob_writeTo_Params) StartAtOffset() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_writeTo_Params) SetStartAtOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Blob_writeTo_Params_List is a list of Blob_writeTo_Params.
type Blob_writeTo_Params_List struct{ capnp.List }

// NewBlob_writeTo_Params creates a new list of Blob_writeTo_Params.
func NewBlob_writeTo_Params_List(s *capnp.Segment, sz int32) (Blob_writeTo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Blob_writeTo_Params_List{l}, err
}

func (s Blob_writeTo_Params_List) At(i int) Blob_writeTo_Params {
	return Blob_writeTo_Params{s.List.Struct(i)}
}

func (s Blob_writeTo_Params_List) Set(i int, v Blob_writeTo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_writeTo_Params_Promise is a wrapper for a Blob_writeTo_Params promised by a client call.
type Blob_writeTo_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_writeTo_Params_Promise) Struct() (Blob_writeTo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_writeTo_Params{s}, err
}

func (p Blob_writeTo_Params_Promise) Stream() ByteStream {
	return ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Blob_writeTo_Results struct{ capnp.Struct }

func NewBlob_writeTo_Results(s *capnp.Segment) (Blob_writeTo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_writeTo_Results{st}, err
}

func NewRootBlob_writeTo_Results(s *capnp.Segment) (Blob_writeTo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_writeTo_Results{st}, err
}

func ReadRootBlob_writeTo_Results(msg *capnp.Message) (Blob_writeTo_Results, error) {
	root, err := msg.RootPtr()
	return Blob_writeTo_Results{root.Struct()}, err
}

func (s Blob_writeTo_Results) String() string {
	str, _ := text.Marshal(0xdb3152bd3bc2aa40, s.Struct)
	return str
}

func (s Blob_writeTo_Results) Handle() Handle {
	p, _ := s.Struct.Ptr(0)
	return Handle{Client: p.Interface().Client()}
}

func (s Blob_writeTo_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_writeTo_Results) SetHandle(v Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Blob_writeTo_Results_List is a list of Blob_writeTo_Results.
type Blob_writeTo_Results_List struct{ capnp.List }

// NewBlob_writeTo_Results creates a new list of Blob_writeTo_Results.
func NewBlob_writeTo_Results_List(s *capnp.Segment, sz int32) (Blob_writeTo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Blob_writeTo_Results_List{l}, err
}

func (s Blob_writeTo_Results_List) At(i int) Blob_writeTo_Results {
	return Blob_writeTo_Results{s.List.Struct(i)}
}

func (s Blob_writeTo_Results_List) Set(i int, v Blob_writeTo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_writeTo_Results_Promise is a wrapper for a Blob_writeTo_Results promised by a client call.
type Blob_writeTo_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_writeTo_Results_Promise) Struct() (Blob_writeTo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_writeTo_Results{s}, err
}

func (p Blob_writeTo_Results_Promise) Handle() Handle {
	return Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Blob_getSlice_Params struct{ capnp.Struct }

func NewBlob_getSlice_Params(s *capnp.Segment) (Blob_getSlice_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Blob_getSlice_Params{st}, err
}

func NewRootBlob_getSlice_Params(s *capnp.Segment) (Blob_getSlice_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Blob_getSlice_Params{st}, err
}

func ReadRootBlob_getSlice_Params(msg *capnp.Message) (Blob_getSlice_Params, error) {
	root, err := msg.RootPtr()
	return Blob_getSlice_Params{root.Struct()}, err
}

func (s Blob_getSlice_Params) String() string {
	str, _ := text.Marshal(0x8edb5f3937d96b8a, s.Struct)
	return str
}

func (s Blob_getSlice_Params) Offset() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_getSlice_Params) SetOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Blob_getSlice_Params) Size() uint32 {
	return s.Struct.Uint32(8)
}

func (s Blob_getSlice_Params) SetSize(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Blob_getSlice_Params_List is a list of Blob_getSlice_Params.
type Blob_getSlice_Params_List struct{ capnp.List }

// NewBlob_getSlice_Params creates a new list of Blob_getSlice_Params.
func NewBlob_getSlice_Params_List(s *capnp.Segment, sz int32) (Blob_getSlice_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return Blob_getSlice_Params_List{l}, err
}

func (s Blob_getSlice_Params_List) At(i int) Blob_getSlice_Params {
	return Blob_getSlice_Params{s.List.Struct(i)}
}

func (s Blob_getSlice_Params_List) Set(i int, v Blob_getSlice_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSlice_Params_Promise is a wrapper for a Blob_getSlice_Params promised by a client call.
type Blob_getSlice_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_getSlice_Params_Promise) Struct() (Blob_getSlice_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSlice_Params{s}, err
}

type Blob_getSlice_Results struct{ capnp.Struct }

func NewBlob_getSlice_Results(s *capnp.Segment) (Blob_getSlice_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_getSlice_Results{st}, err
}

func NewRootBlob_getSlice_Results(s *capnp.Segment) (Blob_getSlice_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_getSlice_Results{st}, err
}

func ReadRootBlob_getSlice_Results(msg *capnp.Message) (Blob_getSlice_Results, error) {
	root, err := msg.RootPtr()
	return Blob_getSlice_Results{root.Struct()}, err
}

func (s Blob_getSlice_Results) String() string {
	str, _ := text.Marshal(0xc65caf9a2d389078, s.Struct)
	return str
}

func (s Blob_getSlice_Results) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Blob_getSlice_Results) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_getSlice_Results) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// Blob_getSlice_Results_List is a list of Blob_getSlice_Results.
type Blob_getSlice_Results_List struct{ capnp.List }

// NewBlob_getSlice_Results creates a new list of Blob_getSlice_Results.
func NewBlob_getSlice_Results_List(s *capnp.Segment, sz int32) (Blob_getSlice_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Blob_getSlice_Results_List{l}, err
}

func (s Blob_getSlice_Results_List) At(i int) Blob_getSlice_Results {
	return Blob_getSlice_Results{s.List.Struct(i)}
}

func (s Blob_getSlice_Results_List) Set(i int, v Blob_getSlice_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSlice_Results_Promise is a wrapper for a Blob_getSlice_Results promised by a client call.
type Blob_getSlice_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_getSlice_Results_Promise) Struct() (Blob_getSlice_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSlice_Results{s}, err
}

type Assignable struct{ Client capnp.Client }

func (c Assignable) Get(ctx context.Context, params func(Assignable_get_Params) error, opts ...capnp.CallOption) Assignable_get_Results_Promise {
	if c.Client == nil {
		return Assignable_get_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_get_Params{Struct: s}) }
	}
	return Assignable_get_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable) AsGetter(ctx context.Context, params func(Assignable_asGetter_Params) error, opts ...capnp.CallOption) Assignable_asGetter_Results_Promise {
	if c.Client == nil {
		return Assignable_asGetter_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asGetter",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_asGetter_Params{Struct: s}) }
	}
	return Assignable_asGetter_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable) AsSetter(ctx context.Context, params func(Assignable_asSetter_Params) error, opts ...capnp.CallOption) Assignable_asSetter_Results_Promise {
	if c.Client == nil {
		return Assignable_asSetter_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      2,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asSetter",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_asSetter_Params{Struct: s}) }
	}
	return Assignable_asSetter_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Server interface {
	Get(Assignable_get) error

	AsGetter(Assignable_asGetter) error

	AsSetter(Assignable_asSetter) error
}

func Assignable_ServerToClient(s Assignable_Server) Assignable {
	c, _ := s.(server.Closer)
	return Assignable{Client: server.New(Assignable_Methods(nil, s), c)}
}

func Assignable_Methods(methods []server.Method, s Assignable_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_get{c, opts, Assignable_get_Params{Struct: p}, Assignable_get_Results{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asGetter",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_asGetter{c, opts, Assignable_asGetter_Params{Struct: p}, Assignable_asGetter_Results{Struct: r}}
			return s.AsGetter(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      2,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asSetter",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_asSetter{c, opts, Assignable_asSetter_Params{Struct: p}, Assignable_asSetter_Results{Struct: r}}
			return s.AsSetter(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Assignable_get holds the arguments for a server call to Assignable.get.
type Assignable_get struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_get_Params
	Results Assignable_get_Results
}

// Assignable_asGetter holds the arguments for a server call to Assignable.asGetter.
type Assignable_asGetter struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_asGetter_Params
	Results Assignable_asGetter_Results
}

// Assignable_asSetter holds the arguments for a server call to Assignable.asSetter.
type Assignable_asSetter struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_asSetter_Params
	Results Assignable_asSetter_Results
}

type Assignable_Getter struct{ Client capnp.Client }

func (c Assignable_Getter) Get(ctx context.Context, params func(Assignable_Getter_get_Params) error, opts ...capnp.CallOption) Assignable_Getter_get_Results_Promise {
	if c.Client == nil {
		return Assignable_Getter_get_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Getter_get_Params{Struct: s}) }
	}
	return Assignable_Getter_get_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable_Getter) Subscribe(ctx context.Context, params func(Assignable_Getter_subscribe_Params) error, opts ...capnp.CallOption) Assignable_Getter_subscribe_Results_Promise {
	if c.Client == nil {
		return Assignable_Getter_subscribe_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "subscribe",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Getter_subscribe_Params{Struct: s}) }
	}
	return Assignable_Getter_subscribe_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Getter_Server interface {
	Get(Assignable_Getter_get) error

	Subscribe(Assignable_Getter_subscribe) error
}

func Assignable_Getter_ServerToClient(s Assignable_Getter_Server) Assignable_Getter {
	c, _ := s.(server.Closer)
	return Assignable_Getter{Client: server.New(Assignable_Getter_Methods(nil, s), c)}
}

func Assignable_Getter_Methods(methods []server.Method, s Assignable_Getter_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Getter_get{c, opts, Assignable_Getter_get_Params{Struct: p}, Assignable_Getter_get_Results{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "subscribe",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Getter_subscribe{c, opts, Assignable_Getter_subscribe_Params{Struct: p}, Assignable_Getter_subscribe_Results{Struct: r}}
			return s.Subscribe(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Assignable_Getter_get holds the arguments for a server call to Assignable_Getter.get.
type Assignable_Getter_get struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Getter_get_Params
	Results Assignable_Getter_get_Results
}

// Assignable_Getter_subscribe holds the arguments for a server call to Assignable_Getter.subscribe.
type Assignable_Getter_subscribe struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Getter_subscribe_Params
	Results Assignable_Getter_subscribe_Results
}

type Assignable_Getter_get_Params struct{ capnp.Struct }

func NewAssignable_Getter_get_Params(s *capnp.Segment) (Assignable_Getter_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Getter_get_Params{st}, err
}

func NewRootAssignable_Getter_get_Params(s *capnp.Segment) (Assignable_Getter_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Getter_get_Params{st}, err
}

func ReadRootAssignable_Getter_get_Params(msg *capnp.Message) (Assignable_Getter_get_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_get_Params{root.Struct()}, err
}

func (s Assignable_Getter_get_Params) String() string {
	str, _ := text.Marshal(0xb19fdbd356844119, s.Struct)
	return str
}

// Assignable_Getter_get_Params_List is a list of Assignable_Getter_get_Params.
type Assignable_Getter_get_Params_List struct{ capnp.List }

// NewAssignable_Getter_get_Params creates a new list of Assignable_Getter_get_Params.
func NewAssignable_Getter_get_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_Getter_get_Params_List{l}, err
}

func (s Assignable_Getter_get_Params_List) At(i int) Assignable_Getter_get_Params {
	return Assignable_Getter_get_Params{s.List.Struct(i)}
}

func (s Assignable_Getter_get_Params_List) Set(i int, v Assignable_Getter_get_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_get_Params_Promise is a wrapper for a Assignable_Getter_get_Params promised by a client call.
type Assignable_Getter_get_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_get_Params_Promise) Struct() (Assignable_Getter_get_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_get_Params{s}, err
}

type Assignable_Getter_get_Results struct{ capnp.Struct }

func NewAssignable_Getter_get_Results(s *capnp.Segment) (Assignable_Getter_get_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_get_Results{st}, err
}

func NewRootAssignable_Getter_get_Results(s *capnp.Segment) (Assignable_Getter_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_get_Results{st}, err
}

func ReadRootAssignable_Getter_get_Results(msg *capnp.Message) (Assignable_Getter_get_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_get_Results{root.Struct()}, err
}

func (s Assignable_Getter_get_Results) String() string {
	str, _ := text.Marshal(0x97ef2da226123492, s.Struct)
	return str
}

func (s Assignable_Getter_get_Results) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_Getter_get_Results) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_get_Results) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_Getter_get_Results) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_Getter_get_Results) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Assignable_Getter_get_Results_List is a list of Assignable_Getter_get_Results.
type Assignable_Getter_get_Results_List struct{ capnp.List }

// NewAssignable_Getter_get_Results creates a new list of Assignable_Getter_get_Results.
func NewAssignable_Getter_get_Results_List(s *capnp.Segment, sz int32) (Assignable_Getter_get_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_get_Results_List{l}, err
}

func (s Assignable_Getter_get_Results_List) At(i int) Assignable_Getter_get_Results {
	return Assignable_Getter_get_Results{s.List.Struct(i)}
}

func (s Assignable_Getter_get_Results_List) Set(i int, v Assignable_Getter_get_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_get_Results_Promise is a wrapper for a Assignable_Getter_get_Results promised by a client call.
type Assignable_Getter_get_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_get_Results_Promise) Struct() (Assignable_Getter_get_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_get_Results{s}, err
}

func (p Assignable_Getter_get_Results_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Assignable_Getter_subscribe_Params struct{ capnp.Struct }

func NewAssignable_Getter_subscribe_Params(s *capnp.Segment) (Assignable_Getter_subscribe_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Params{st}, err
}

func NewRootAssignable_Getter_subscribe_Params(s *capnp.Segment) (Assignable_Getter_subscribe_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Params{st}, err
}

func ReadRootAssignable_Getter_subscribe_Params(msg *capnp.Message) (Assignable_Getter_subscribe_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_subscribe_Params{root.Struct()}, err
}

func (s Assignable_Getter_subscribe_Params) String() string {
	str, _ := text.Marshal(0xf02783ef982ecea9, s.Struct)
	return str
}

func (s Assignable_Getter_subscribe_Params) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Params) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Params) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_Getter_subscribe_Params_List is a list of Assignable_Getter_subscribe_Params.
type Assignable_Getter_subscribe_Params_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Params creates a new list of Assignable_Getter_subscribe_Params.
func NewAssignable_Getter_subscribe_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_subscribe_Params_List{l}, err
}

func (s Assignable_Getter_subscribe_Params_List) At(i int) Assignable_Getter_subscribe_Params {
	return Assignable_Getter_subscribe_Params{s.List.Struct(i)}
}

func (s Assignable_Getter_subscribe_Params_List) Set(i int, v Assignable_Getter_subscribe_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_subscribe_Params_Promise is a wrapper for a Assignable_Getter_subscribe_Params promised by a client call.
type Assignable_Getter_subscribe_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_subscribe_Params_Promise) Struct() (Assignable_Getter_subscribe_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_subscribe_Params{s}, err
}

func (p Assignable_Getter_subscribe_Params_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_Getter_subscribe_Results struct{ capnp.Struct }

func NewAssignable_Getter_subscribe_Results(s *capnp.Segment) (Assignable_Getter_subscribe_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Results{st}, err
}

func NewRootAssignable_Getter_subscribe_Results(s *capnp.Segment) (Assignable_Getter_subscribe_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Results{st}, err
}

func ReadRootAssignable_Getter_subscribe_Results(msg *capnp.Message) (Assignable_Getter_subscribe_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_subscribe_Results{root.Struct()}, err
}

func (s Assignable_Getter_subscribe_Results) String() string {
	str, _ := text.Marshal(0x84e0f802c9af605b, s.Struct)
	return str
}

func (s Assignable_Getter_subscribe_Results) Handle() Handle {
	p, _ := s.Struct.Ptr(0)
	return Handle{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Results) SetHandle(v Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_Getter_subscribe_Results_List is a list of Assignable_Getter_subscribe_Results.
type Assignable_Getter_subscribe_Results_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Results creates a new list of Assignable_Getter_subscribe_Results.
func NewAssignable_Getter_subscribe_Results_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_subscribe_Results_List{l}, err
}

func (s Assignable_Getter_subscribe_Results_List) At(i int) Assignable_Getter_subscribe_Results {
	return Assignable_Getter_subscribe_Results{s.List.Struct(i)}
}

func (s Assignable_Getter_subscribe_Results_List) Set(i int, v Assignable_Getter_subscribe_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_subscribe_Results_Promise is a wrapper for a Assignable_Getter_subscribe_Results promised by a client call.
type Assignable_Getter_subscribe_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_subscribe_Results_Promise) Struct() (Assignable_Getter_subscribe_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_subscribe_Results{s}, err
}

func (p Assignable_Getter_subscribe_Results_Promise) Handle() Handle {
	return Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_Setter struct{ Client capnp.Client }

func (c Assignable_Setter) Set(ctx context.Context, params func(Assignable_Setter_set_Params) error, opts ...capnp.CallOption) Assignable_Setter_set_Results_Promise {
	if c.Client == nil {
		return Assignable_Setter_set_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd5256a3f93589d2f,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Setter",
			MethodName:    "set",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Setter_set_Params{Struct: s}) }
	}
	return Assignable_Setter_set_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Setter_Server interface {
	Set(Assignable_Setter_set) error
}

func Assignable_Setter_ServerToClient(s Assignable_Setter_Server) Assignable_Setter {
	c, _ := s.(server.Closer)
	return Assignable_Setter{Client: server.New(Assignable_Setter_Methods(nil, s), c)}
}

func Assignable_Setter_Methods(methods []server.Method, s Assignable_Setter_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd5256a3f93589d2f,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Setter",
			MethodName:    "set",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Setter_set{c, opts, Assignable_Setter_set_Params{Struct: p}, Assignable_Setter_set_Results{Struct: r}}
			return s.Set(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Assignable_Setter_set holds the arguments for a server call to Assignable_Setter.set.
type Assignable_Setter_set struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Setter_set_Params
	Results Assignable_Setter_set_Results
}

type Assignable_Setter_set_Params struct{ capnp.Struct }

func NewAssignable_Setter_set_Params(s *capnp.Segment) (Assignable_Setter_set_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Setter_set_Params{st}, err
}

func NewRootAssignable_Setter_set_Params(s *capnp.Segment) (Assignable_Setter_set_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Setter_set_Params{st}, err
}

func ReadRootAssignable_Setter_set_Params(msg *capnp.Message) (Assignable_Setter_set_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Setter_set_Params{root.Struct()}, err
}

func (s Assignable_Setter_set_Params) String() string {
	str, _ := text.Marshal(0x98d0372787b787d1, s.Struct)
	return str
}

func (s Assignable_Setter_set_Params) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_Setter_set_Params) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Setter_set_Params) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_Setter_set_Params) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_Setter_set_Params) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Assignable_Setter_set_Params_List is a list of Assignable_Setter_set_Params.
type Assignable_Setter_set_Params_List struct{ capnp.List }

// NewAssignable_Setter_set_Params creates a new list of Assignable_Setter_set_Params.
func NewAssignable_Setter_set_Params_List(s *capnp.Segment, sz int32) (Assignable_Setter_set_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Setter_set_Params_List{l}, err
}

func (s Assignable_Setter_set_Params_List) At(i int) Assignable_Setter_set_Params {
	return Assignable_Setter_set_Params{s.List.Struct(i)}
}

func (s Assignable_Setter_set_Params_List) Set(i int, v Assignable_Setter_set_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Setter_set_Params_Promise is a wrapper for a Assignable_Setter_set_Params promised by a client call.
type Assignable_Setter_set_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Setter_set_Params_Promise) Struct() (Assignable_Setter_set_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Setter_set_Params{s}, err
}

func (p Assignable_Setter_set_Params_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Assignable_Setter_set_Results struct{ capnp.Struct }

func NewAssignable_Setter_set_Results(s *capnp.Segment) (Assignable_Setter_set_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Setter_set_Results{st}, err
}

func NewRootAssignable_Setter_set_Results(s *capnp.Segment) (Assignable_Setter_set_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Setter_set_Results{st}, err
}

func ReadRootAssignable_Setter_set_Results(msg *capnp.Message) (Assignable_Setter_set_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Setter_set_Results{root.Struct()}, err
}

func (s Assignable_Setter_set_Results) String() string {
	str, _ := text.Marshal(0xdbfbb635d3e6abab, s.Struct)
	return str
}

// Assignable_Setter_set_Results_List is a list of Assignable_Setter_set_Results.
type Assignable_Setter_set_Results_List struct{ capnp.List }

// NewAssignable_Setter_set_Results creates a new list of Assignable_Setter_set_Results.
func NewAssignable_Setter_set_Results_List(s *capnp.Segment, sz int32) (Assignable_Setter_set_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_Setter_set_Results_List{l}, err
}

func (s Assignable_Setter_set_Results_List) At(i int) Assignable_Setter_set_Results {
	return Assignable_Setter_set_Results{s.List.Struct(i)}
}

func (s Assignable_Setter_set_Results_List) Set(i int, v Assignable_Setter_set_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Setter_set_Results_Promise is a wrapper for a Assignable_Setter_set_Results promised by a client call.
type Assignable_Setter_set_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Setter_set_Results_Promise) Struct() (Assignable_Setter_set_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Setter_set_Results{s}, err
}

type Assignable_get_Params struct{ capnp.Struct }

func NewAssignable_get_Params(s *capnp.Segment) (Assignable_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_get_Params{st}, err
}

func NewRootAssignable_get_Params(s *capnp.Segment) (Assignable_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_get_Params{st}, err
}

func ReadRootAssignable_get_Params(msg *capnp.Message) (Assignable_get_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_get_Params{root.Struct()}, err
}

func (s Assignable_get_Params) String() string {
	str, _ := text.Marshal(0xbbfd27b5d2515662, s.Struct)
	return str
}

// Assignable_get_Params_List is a list of Assignable_get_Params.
type Assignable_get_Params_List struct{ capnp.List }

// NewAssignable_get_Params creates a new list of Assignable_get_Params.
func NewAssignable_get_Params_List(s *capnp.Segment, sz int32) (Assignable_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_get_Params_List{l}, err
}

func (s Assignable_get_Params_List) At(i int) Assignable_get_Params {
	return Assignable_get_Params{s.List.Struct(i)}
}

func (s Assignable_get_Params_List) Set(i int, v Assignable_get_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_get_Params_Promise is a wrapper for a Assignable_get_Params promised by a client call.
type Assignable_get_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_get_Params_Promise) Struct() (Assignable_get_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_get_Params{s}, err
}

type Assignable_get_Results struct{ capnp.Struct }

func NewAssignable_get_Results(s *capnp.Segment) (Assignable_get_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Assignable_get_Results{st}, err
}

func NewRootAssignable_get_Results(s *capnp.Segment) (Assignable_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Assignable_get_Results{st}, err
}

func ReadRootAssignable_get_Results(msg *capnp.Message) (Assignable_get_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_get_Results{root.Struct()}, err
}

func (s Assignable_get_Results) String() string {
	str, _ := text.Marshal(0xb351b437cd426a4f, s.Struct)
	return str
}

func (s Assignable_get_Results) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_get_Results) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_get_Results) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_get_Results) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_get_Results) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s Assignable_get_Results) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(1)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_get_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Assignable_get_Results) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// Assignable_get_Results_List is a list of Assignable_get_Results.
type Assignable_get_Results_List struct{ capnp.List }

// NewAssignable_get_Results creates a new list of Assignable_get_Results.
func NewAssignable_get_Results_List(s *capnp.Segment, sz int32) (Assignable_get_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Assignable_get_Results_List{l}, err
}

func (s Assignable_get_Results_List) At(i int) Assignable_get_Results {
	return Assignable_get_Results{s.List.Struct(i)}
}

func (s Assignable_get_Results_List) Set(i int, v Assignable_get_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_get_Results_Promise is a wrapper for a Assignable_get_Results promised by a client call.
type Assignable_get_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_get_Results_Promise) Struct() (Assignable_get_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_get_Results{s}, err
}

func (p Assignable_get_Results_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Assignable_get_Results_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(1).Client()}
}

type Assignable_asGetter_Params struct{ capnp.Struct }

func NewAssignable_asGetter_Params(s *capnp.Segment) (Assignable_asGetter_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asGetter_Params{st}, err
}

func NewRootAssignable_asGetter_Params(s *capnp.Segment) (Assignable_asGetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asGetter_Params{st}, err
}

func ReadRootAssignable_asGetter_Params(msg *capnp.Message) (Assignable_asGetter_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_asGetter_Params{root.Struct()}, err
}

func (s Assignable_asGetter_Params) String() string {
	str, _ := text.Marshal(0xf907945b872b26cf, s.Struct)
	return str
}

// Assignable_asGetter_Params_List is a list of Assignable_asGetter_Params.
type Assignable_asGetter_Params_List struct{ capnp.List }

// NewAssignable_asGetter_Params creates a new list of Assignable_asGetter_Params.
func NewAssignable_asGetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_asGetter_Params_List{l}, err
}

func (s Assignable_asGetter_Params_List) At(i int) Assignable_asGetter_Params {
	return Assignable_asGetter_Params{s.List.Struct(i)}
}

func (s Assignable_asGetter_Params_List) Set(i int, v Assignable_asGetter_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asGetter_Params_Promise is a wrapper for a Assignable_asGetter_Params promised by a client call.
type Assignable_asGetter_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_asGetter_Params_Promise) Struct() (Assignable_asGetter_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asGetter_Params{s}, err
}

type Assignable_asGetter_Results struct{ capnp.Struct }

func NewAssignable_asGetter_Results(s *capnp.Segment) (Assignable_asGetter_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asGetter_Results{st}, err
}

func NewRootAssignable_asGetter_Results(s *capnp.Segment) (Assignable_asGetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asGetter_Results{st}, err
}

func ReadRootAssignable_asGetter_Results(msg *capnp.Message) (Assignable_asGetter_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_asGetter_Results{root.Struct()}, err
}

func (s Assignable_asGetter_Results) String() string {
	str, _ := text.Marshal(0x8c3d547ef2930e96, s.Struct)
	return str
}

func (s Assignable_asGetter_Results) Getter() Assignable_Getter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Getter{Client: p.Interface().Client()}
}

func (s Assignable_asGetter_Results) HasGetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asGetter_Results) SetGetter(v Assignable_Getter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_asGetter_Results_List is a list of Assignable_asGetter_Results.
type Assignable_asGetter_Results_List struct{ capnp.List }

// NewAssignable_asGetter_Results creates a new list of Assignable_asGetter_Results.
func NewAssignable_asGetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_asGetter_Results_List{l}, err
}

func (s Assignable_asGetter_Results_List) At(i int) Assignable_asGetter_Results {
	return Assignable_asGetter_Results{s.List.Struct(i)}
}

func (s Assignable_asGetter_Results_List) Set(i int, v Assignable_asGetter_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asGetter_Results_Promise is a wrapper for a Assignable_asGetter_Results promised by a client call.
type Assignable_asGetter_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_asGetter_Results_Promise) Struct() (Assignable_asGetter_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asGetter_Results{s}, err
}

func (p Assignable_asGetter_Results_Promise) Getter() Assignable_Getter {
	return Assignable_Getter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_asSetter_Params struct{ capnp.Struct }

func NewAssignable_asSetter_Params(s *capnp.Segment) (Assignable_asSetter_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asSetter_Params{st}, err
}

func NewRootAssignable_asSetter_Params(s *capnp.Segment) (Assignable_asSetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asSetter_Params{st}, err
}

func ReadRootAssignable_asSetter_Params(msg *capnp.Message) (Assignable_asSetter_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_asSetter_Params{root.Struct()}, err
}

func (s Assignable_asSetter_Params) String() string {
	str, _ := text.Marshal(0xa01f603357f3b349, s.Struct)
	return str
}

// Assignable_asSetter_Params_List is a list of Assignable_asSetter_Params.
type Assignable_asSetter_Params_List struct{ capnp.List }

// NewAssignable_asSetter_Params creates a new list of Assignable_asSetter_Params.
func NewAssignable_asSetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_asSetter_Params_List{l}, err
}

func (s Assignable_asSetter_Params_List) At(i int) Assignable_asSetter_Params {
	return Assignable_asSetter_Params{s.List.Struct(i)}
}

func (s Assignable_asSetter_Params_List) Set(i int, v Assignable_asSetter_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asSetter_Params_Promise is a wrapper for a Assignable_asSetter_Params promised by a client call.
type Assignable_asSetter_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_asSetter_Params_Promise) Struct() (Assignable_asSetter_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asSetter_Params{s}, err
}

type Assignable_asSetter_Results struct{ capnp.Struct }

func NewAssignable_asSetter_Results(s *capnp.Segment) (Assignable_asSetter_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asSetter_Results{st}, err
}

func NewRootAssignable_asSetter_Results(s *capnp.Segment) (Assignable_asSetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asSetter_Results{st}, err
}

func ReadRootAssignable_asSetter_Results(msg *capnp.Message) (Assignable_asSetter_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_asSetter_Results{root.Struct()}, err
}

func (s Assignable_asSetter_Results) String() string {
	str, _ := text.Marshal(0xc6cbc10181c4f397, s.Struct)
	return str
}

func (s Assignable_asSetter_Results) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_asSetter_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asSetter_Results) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_asSetter_Results_List is a list of Assignable_asSetter_Results.
type Assignable_asSetter_Results_List struct{ capnp.List }

// NewAssignable_asSetter_Results creates a new list of Assignable_asSetter_Results.
func NewAssignable_asSetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_asSetter_Results_List{l}, err
}

func (s Assignable_asSetter_Results_List) At(i int) Assignable_asSetter_Results {
	return Assignable_asSetter_Results{s.List.Struct(i)}
}

func (s Assignable_asSetter_Results_List) Set(i int, v Assignable_asSetter_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asSetter_Results_Promise is a wrapper for a Assignable_asSetter_Results promised by a client call.
type Assignable_asSetter_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_asSetter_Results_Promise) Struct() (Assignable_asSetter_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asSetter_Results{s}, err
}

func (p Assignable_asSetter_Results_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(0).Client()}
}

const schema_ecd50d792c3d9992 = "x\xda\xacW}lTY\x15\xbfgf\xda7_\xa5" +
	"}}\xdd5]%\x936]\xba[\x99\xb1\x1f`\xa1" +
	"n3\x9dIVh\x97\xcdv\xda.\xac\xee6\xf4\xb5" +
	"}\xe0\xc0\xb4S\xe6\xbdR\x8aQ\xac6iT\x82\x94" +
	"\x0f\x0b\x08\x81\x10\x8db\x8d\x84/\x09Q\xfe\xd0\x08\xa4" +
	"\"B%\x10H\xc5\x80Z\x13\x08\x8a\xb6H\x04c\xc6" +
	"s\xef\xbc\xfb\xe6M;\xd3Z\xf4\x8fIf\xe6\xfd\xde" +
	"9\xbf\xf3;\x1f\xf7\xdc\xf2\x1f\xd8\xea,\x15Y\xd1\x1c" +
	"BB\x87\xb3\xb2\xe3%\xc1\xdbm\xcd\xcf\xa6\xbeB\xc4" +
	"W\xad\xf1\x83\xc3\xc5\x07\xce\xbe?\xf5\x88\x10\x92\x07\xd2" +
	"\x91\xecG\xd2\x89l\x81\x10\xe9{\xd9CR\xbd\x80\xdf" +
	"\xe2\x1f\xb6\x9d\x1c\xb3\xfc\xf3\xfe \x11\x17\x03!Y " +
	"\xe4AU\x85\xb0\x1d\x08H\xb5\x82\x9f@\xbc\xfe\x0f_" +
	"\xac\x8e\x9do\xfd&\x11]\x10\xdfs\xb0vi\x7f\xce" +
	"\xad\xc7$\xcbB\xcd\xb4\x0a\xbf\x92\xc2\xc2\xc7\xf0[\xaf" +
	"\xd0\x87\xd8o/\xda;\xf5\xe5\x96\xda\x9d\xe8\xdb0\xf6" +
	"P\xa8\xa1\xc6\xfe\xce\x8c\xdd\x94\xafh[\xdf:\xb4\x93" +
	"\x84^\x05D\xd8\xd0H\xd5\xeb\xf6 \x05x\xed\x14\xf0" +
	"\xf4\xf6\xf4H\xc1\xd5\xd5\xbbH(\xcf\x00\xbck\xb7P" +
	"@\x88\x01\xbe\xb1\xf9n\xf5\xca\xf5\x13\x0c`\xd1\x01[" +
	"\x12\x80^;\xe5\xd0:<\xea\xf1=\x9b\xdcE\xc4<" +
	"\xdd\x804n\x7fAl\xf1\x92e\xa7F\xff4pl" +
	"_\xba@\xce\xd9\x8fK\x17\xed\xf4\xdb\x05fd*\x9a" +
	"_T\x94\xff\x97\x11#\x10\xf4\xf2\x8a\xe35\xeae\xb1" +
	"\x83\xd2\xd8\xb3,\x7f\xc9q\xef\x93\x11\xb3l\xb5\x8e\x06" +
	"\x0a\xa8g\x80\xf1\xa1\xf3C\xa5\xd57\x0e\x98\x01a\x07" +
	"\x8bt\x0b\x03\xdc\x0f\xb6\xb5\xfd\xa8\xe4)\x02\\\xd6$" +
	"\x1d|<\xec\xd8#\x1dtP.\xfb\x1d\x02\xfd`\x96" +
	"\xde\xf9\xf5\xefN<,\x14\x8e\xea\xb20B\xd2\x97\x1c" +
	"/\x10>\xe0\xa0\x84\xeb\xcfL\xaf\xabj\xf3\x1cK\x10" +
	"\xb6Qo\x0f\x1c\x95\x80a\x0f\xf7}\xfc^\xf6\x8f\xff" +
	"\xf1]\"\x16\x9a\x92\xc9\xc2\xae\x1aC\x88t\x97\xf9\xba" +
	"\xc5\xcc\x14\x06\x06\xd7\xde\x9c8z*\xc1\x9a\x99\xa9p" +
	"\x06\xa9\x99\xf76\x05\xafU\x9f\x0d\x9d\xd1\x15\xb1\xd0G" +
	"\x85\xce|\x1aO\x91\x93\xbe\xda\xbe6\xf4\xdbs\xa5\xff" +
	"\xfe\xa9\x89\xc1\xb0\xd3I_\xed>\xff\xc1\xa7\xeb\x97\x14" +
	"\xfc\x8c?A\xc7\xfd\xf4M[|\xdb\xee\x15\xdeC'" +
	"?\xba\x9c\xc8UBf\x99\xbe\x04R\xd8IE\x1a\x99" +
	"\xfe\xe5\x00\xfc\xfc\xeaesA}\xdd\xc9\x0aj\x98\x01" +
	"\xd6\xff\xf1\xfa\x9b}+\xd6]\x9b\xa5\xe2%\xe7i\xe9" +
	"\x9a\x93F6\xe6\xbc\"\xbd\xed\xa2*\x16\x9e.\xff\xeb" +
	"\xf2\xbbwn\x98\x88x]\xafQ\"\x9f:\xf2\xc1^" +
	"\xff\xa6\xd7o\xa5i\x9aW\\\x8f\xa4\"\xfa\xba\xb4\xd8" +
	"\xb5J\xeab\x86\xea~\xf8\x8b\xcf\\l\xaa\x980\xf3" +
	"~\xdf\xc5\x8a\xf0s.Jkt\xf4\xcf7\x97\xff\xe4" +
	"_\x13&\x1d\xfb]\x0d\xd4\xd3\xb7\x0aw\xb7~\xbft" +
	"\xf9\xe4,\xc2\x8a\xeb\xabR\x98\xf9Q\\W\xa4\x07\xcc" +
	"\x8f\xc1$\x15Li\x8d\xb9NK\xe3\xaeR\x84O\"" +
	"\xfc\x9e[\xc0On|l\xd0\x1f=\\\xf6\xfb\xc7\xa6" +
	"\x08\xc7\xdd\xc5\xd4\xef\x89\xdf\xf8\x0e<\xf9Z\xe9\xdf\xcc" +
	"\xf5x\xc1\x1d\xa3\x94/\xb9)\xe5\xef\xac\x0a\x95\xdd\xa9" +
	"_7mz\xf5\xb9\x9bQ\xbe\xbe\xe4\x93C\x1f\xee\x13" +
	"\x9e\x9b2;\xe9\xc6\xda\xfal\xbcW\x0bG|\x1dr" +
	"\x8f\xa5\xbb\xa7&\xa0\xaa\xe1\x8d\xddr{D\xf1\xadR" +
	"r5M\x895\x02\x84\xec\xd6,*\xbc^V\xc0\xdb" +
	"F\xac(&$\xb0\x14\x02o\x81\xf8\xae\x00`\xd0\x03" +
	">\x8f\xc4@\x13\"\xea \xd0\x08\xa2,\x08\x1b\x15M" +
	"\x04O\xc8f\x81\xa4(\xb4\x13\xd2\xfeY\x07q\xb5\xb7" +
	"]\xed\x88\x85\xdb\x09(\x0by\x0f)\x1bA\xd9f\x06" +
	"Ec\xf2q\xc3JI\x93_Q{#\x9a\x1a\xb2Y" +
	"m(\x0b5\x9aS\x83C\xd8n\x85P\x81\x05\xfc_" +
	"\x90\xbb;#\x0a\x88\xc9NG'\"\x81\x14\xd5\xd6D" +
	";\xe4Hx\xbb\xd2\xd9\xa2l\xd3\xf0e\x1b\x12Hv" +
	",l\x8a\xeb\x00\x99\xe4j\xe1h7\x1a\xe7\xbe\xdel" +
	"G\xf8\x1b\xe8k\x99\x05D\x80\x02\x16UE\x0c\xff," +
	"\xc7?\xd7`l\x9d\xca\x06\x19\x09\xb6\x10\x01m\x83\x9b" +
	"X\xf0\x03\xf1\x087\xe8\xa1\x06UXD\xa0\xd1\x0a\x90" +
	"\x97t\x8b4\x17\x99hZSu\x90\xd5\x84\x12s\xc6" +
	"\xff\x06\xc6\xbf\x91\xc10~\xe3D\xa2\x0a\xa7\x15>E" +
	"\x16\xaa{\xb0_S\x9a\xb5\x98\"w\xf9\x94m=J" +
	"\x87\xd6\x8c\x1a\x954\xca1\xb9K%f\x87eI\xc1" +
	"sU\xc4\x80\x03\xc3t\xcc`\x1f\x8cD\xdb}H\x87" +
	"\x19iJ\xb0\xfe\x1f\xacD\xc2\x1dI.\xa6\x84\xd0\xe0" +
	"K\xd0J9M\x88%\x91\x10oY2K\xfe\xe8\x86" +
	"\x0d\xaa\xa2q\xe3\x09Ov\xfca\x9fQ\x14)|\xfd" +
	"\x09O\x06\x00\x10\xf0\x8e\xd2\xbfV\x16\"\xbdJ\xa2\xc5" +
	"\x0c\x06\xc5f\x06zIx+\x93\x0c\x84\xcdJ?/" +
	"\x04\xcfV\x19\x0d\x18e\x91\x12iR\xfd\xbeXXK" +
	"\x04k\xedR3I\xd6)k2\xe4\xa0\xa1\x1c2_" +
	"\xfb`TF\x06RRP\x99\xb4\xa73\xcb\x070\x8f" +
	"b\xc8\xcfl\xbcY\xefM4\x9e\xa6F\x16f\x9a\xea" +
	"\xbb\x9a\xb6.Pu\x1b\xadY\xe6\x99`d\x87\xe9\xd2" +
	"\x12\xe5\xd9\xc9P\x06,\x07\xd8\xf4^\xda\x96K\xf1\xbf" +
	"\x15X\x05*S\x16\xfb\xc28\xbb\xf4\xb9\xa0jrL" +
	"\x0bh\xef\x11OJ\x9ddn\xc4D\xd8%\x8d\x9e\xd4" +
	"\x0a\xb1\xce\x9c+>\xfd\x97G\xa6-?\xa3fj\xd2" +
	"\xd5\x8c\xb9j\xd9\xc00\x0a%W3\x0f\x13\xee\xd2\x9e" +
	"1\xd9z>80\x15\x97\xac\x06H\x91\xb02\x1d\xa7" +
	"\x1a\x9d\xd3\xea\xb9\xd2\xe8W\xf9\xd41\x8e\xf4\xffn\xea" +
	"Xg3\xd3\xa9\x93L\xbd\xd1\x19\xed\xd6\xe7\x00\xa8s" +
	"L\x8a\xb4\xe5>o\xfbdJ\xf6|S\xf7e\xe3\x07" +
	"\x1e\x9b\x87\x05G\x8b\xc4\xcd\xcen\xbe\x0a\x03\xdf-\xc4" +
	"P%\xb1\x88o\xd3S\x9bov\xc07+qe\x19" +
	">\xf3\x0a`1V}\xe0\x8b\x85X\xf4y|V(" +
	"xX\xef\xd4a\xdc\xa8\x1f\x1e\xd5|\xc0\x13\xebv%" +
	"\xf5\x04N\xa77S\xd3\xaa\xa9\x99v\x8f\xe6\xe4\xeea" +
	"c\xfc\xf9\"\x0e|'\x13E\xba{\xb8!\xf0\x09l" +
	"LAP\x17\xb6Y\xcc\xe2g\x9e\x06\xe9\xce\x96\x05\xad" +
	"\x04\x99\xe7\x9aQE\xa9\x19C\xef$\x99+~\xf7\x01" +
	"~\x8d\x12CA\x9e+~\x85\x00\xbe\xbc\x8a+\x83<" +
	"W\xfcR\x05|!\x17\x8b\x1aX\xaev\xe8\xa7P\x1d" +
	"\xec\xd0c\xc4\x84\xf1\xca&3\xf5\x00\xce\xde\xc3\xe8\x87" +
	"\xec\xf8,y\x1du\xd4\x98\xd6\xec\xac\x1a\x7fbH\xf8" +
	"\x131\xea\x01\xf0K\x04\xf0\xeb\x86\x18\xa2\xc9Z\x03\x81" +
	"\x8f@\xec\xa2a\xf0e\x14\xf8]S\x94\x1b\x10\xd1\x06" +
	"\x81\x1e\x10\x07h0\xfc*\x04\xfc\xf2 \xf6R\x84\x06" +
	"\x81A\x10\xf7/|\x95\xe4\x0b\x0f\xc9\x00\xc9\xfc^\xf3" +
	"\xcb\xbc\x87\x82\x06l\xd8\xb0\xf9\xd02\xf7\xa1L+B" +
	"\x88\x98:a\xbe\x85U?'\xfe\xdf\x93#\xf3\xbe6" +
	"\xbbf3\xed\x92:\xb5\xff\x04\x00\x00\xff\xff\xacu\xbd" +
	"\xee"

func init() {
	schemas.Register(schema_ecd50d792c3d9992,
		0x80f2f65360d64224,
		0x84e0f802c9af605b,
		0x8b5db772377be249,
		0x8c3d547ef2930e96,
		0x8c9a3c7674c761d3,
		0x8e48cb1497f3d6f4,
		0x8edb5f3937d96b8a,
		0x8ee5f62e1fab915d,
		0x94a081e4abb13424,
		0x97ed122121126ff2,
		0x97ef2da226123492,
		0x98d0372787b787d1,
		0x98f424ac606042e0,
		0x9f0719e9a9dccc4b,
		0xa01f603357f3b349,
		0xa4f5ae06dd1b7791,
		0xb19fdbd356844119,
		0xb351b437cd426a4f,
		0xbbfd27b5d2515662,
		0xbc1426493658b76e,
		0xc65caf9a2d389078,
		0xc6cbc10181c4f397,
		0xcd57387729cfe35f,
		0xd0d8d935ee30b219,
		0xd5256a3f93589d2f,
		0xdb3152bd3bc2aa40,
		0xdbfbb635d3e6abab,
		0xe53527a75d90198f,
		0xeaf255b498229199,
		0xecde2a9c6f3f84c9,
		0xf02783ef982ecea9,
		0xf35749d82a51479b,
		0xf907945b872b26cf)
}
