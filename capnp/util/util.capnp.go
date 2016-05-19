package util

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type KeyValue struct{ capnp.Struct }

func NewKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return KeyValue{}, err
	}
	return KeyValue{st}, nil
}

func NewRootKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return KeyValue{}, err
	}
	return KeyValue{st}, nil
}

func ReadRootKeyValue(msg *capnp.Message) (KeyValue, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return KeyValue{}, err
	}
	return KeyValue{root.Struct()}, nil
}
func (s KeyValue) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s KeyValue) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s KeyValue) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s KeyValue) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s KeyValue) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return KeyValue_List{}, err
	}
	return KeyValue_List{l}, nil
}

func (s KeyValue_List) At(i int) KeyValue           { return KeyValue{s.List.Struct(i)} }
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
	if err != nil {
		return LocalizedText{}, err
	}
	return LocalizedText{st}, nil
}

func NewRootLocalizedText(s *capnp.Segment) (LocalizedText, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return LocalizedText{}, err
	}
	return LocalizedText{st}, nil
}

func ReadRootLocalizedText(msg *capnp.Message) (LocalizedText, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return LocalizedText{}, err
	}
	return LocalizedText{root.Struct()}, nil
}
func (s LocalizedText) DefaultText() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s LocalizedText) HasDefaultText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText) DefaultTextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return LocalizedText_Localization_List{}, err
	}
	return LocalizedText_Localization_List{List: p.List()}, nil
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
	if err != nil {
		return LocalizedText_List{}, err
	}
	return LocalizedText_List{l}, nil
}

func (s LocalizedText_List) At(i int) LocalizedText           { return LocalizedText{s.List.Struct(i)} }
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
	if err != nil {
		return LocalizedText_Localization{}, err
	}
	return LocalizedText_Localization{st}, nil
}

func NewRootLocalizedText_Localization(s *capnp.Segment) (LocalizedText_Localization, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return LocalizedText_Localization{}, err
	}
	return LocalizedText_Localization{st}, nil
}

func ReadRootLocalizedText_Localization(msg *capnp.Message) (LocalizedText_Localization, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return LocalizedText_Localization{}, err
	}
	return LocalizedText_Localization{root.Struct()}, nil
}
func (s LocalizedText_Localization) Locale() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s LocalizedText_Localization) HasLocale() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) LocaleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s LocalizedText_Localization) HasText() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
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
	if err != nil {
		return LocalizedText_Localization_List{}, err
	}
	return LocalizedText_Localization_List{l}, nil
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
	if err != nil {
		return ByteStream_write_Params{}, err
	}
	return ByteStream_write_Params{st}, nil
}

func NewRootByteStream_write_Params(s *capnp.Segment) (ByteStream_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return ByteStream_write_Params{}, err
	}
	return ByteStream_write_Params{st}, nil
}

func ReadRootByteStream_write_Params(msg *capnp.Message) (ByteStream_write_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_write_Params{}, err
	}
	return ByteStream_write_Params{root.Struct()}, nil
}
func (s ByteStream_write_Params) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
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
	if err != nil {
		return ByteStream_write_Params_List{}, err
	}
	return ByteStream_write_Params_List{l}, nil
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
	if err != nil {
		return ByteStream_write_Results{}, err
	}
	return ByteStream_write_Results{st}, nil
}

func NewRootByteStream_write_Results(s *capnp.Segment) (ByteStream_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ByteStream_write_Results{}, err
	}
	return ByteStream_write_Results{st}, nil
}

func ReadRootByteStream_write_Results(msg *capnp.Message) (ByteStream_write_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_write_Results{}, err
	}
	return ByteStream_write_Results{root.Struct()}, nil
}

// ByteStream_write_Results_List is a list of ByteStream_write_Results.
type ByteStream_write_Results_List struct{ capnp.List }

// NewByteStream_write_Results creates a new list of ByteStream_write_Results.
func NewByteStream_write_Results_List(s *capnp.Segment, sz int32) (ByteStream_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return ByteStream_write_Results_List{}, err
	}
	return ByteStream_write_Results_List{l}, nil
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
	if err != nil {
		return ByteStream_done_Params{}, err
	}
	return ByteStream_done_Params{st}, nil
}

func NewRootByteStream_done_Params(s *capnp.Segment) (ByteStream_done_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ByteStream_done_Params{}, err
	}
	return ByteStream_done_Params{st}, nil
}

func ReadRootByteStream_done_Params(msg *capnp.Message) (ByteStream_done_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_done_Params{}, err
	}
	return ByteStream_done_Params{root.Struct()}, nil
}

// ByteStream_done_Params_List is a list of ByteStream_done_Params.
type ByteStream_done_Params_List struct{ capnp.List }

// NewByteStream_done_Params creates a new list of ByteStream_done_Params.
func NewByteStream_done_Params_List(s *capnp.Segment, sz int32) (ByteStream_done_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return ByteStream_done_Params_List{}, err
	}
	return ByteStream_done_Params_List{l}, nil
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
	if err != nil {
		return ByteStream_done_Results{}, err
	}
	return ByteStream_done_Results{st}, nil
}

func NewRootByteStream_done_Results(s *capnp.Segment) (ByteStream_done_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ByteStream_done_Results{}, err
	}
	return ByteStream_done_Results{st}, nil
}

func ReadRootByteStream_done_Results(msg *capnp.Message) (ByteStream_done_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_done_Results{}, err
	}
	return ByteStream_done_Results{root.Struct()}, nil
}

// ByteStream_done_Results_List is a list of ByteStream_done_Results.
type ByteStream_done_Results_List struct{ capnp.List }

// NewByteStream_done_Results creates a new list of ByteStream_done_Results.
func NewByteStream_done_Results_List(s *capnp.Segment, sz int32) (ByteStream_done_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return ByteStream_done_Results_List{}, err
	}
	return ByteStream_done_Results_List{l}, nil
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
	if err != nil {
		return ByteStream_expectSize_Params{}, err
	}
	return ByteStream_expectSize_Params{st}, nil
}

func NewRootByteStream_expectSize_Params(s *capnp.Segment) (ByteStream_expectSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return ByteStream_expectSize_Params{}, err
	}
	return ByteStream_expectSize_Params{st}, nil
}

func ReadRootByteStream_expectSize_Params(msg *capnp.Message) (ByteStream_expectSize_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_expectSize_Params{}, err
	}
	return ByteStream_expectSize_Params{root.Struct()}, nil
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
	if err != nil {
		return ByteStream_expectSize_Params_List{}, err
	}
	return ByteStream_expectSize_Params_List{l}, nil
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
	if err != nil {
		return ByteStream_expectSize_Results{}, err
	}
	return ByteStream_expectSize_Results{st}, nil
}

func NewRootByteStream_expectSize_Results(s *capnp.Segment) (ByteStream_expectSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return ByteStream_expectSize_Results{}, err
	}
	return ByteStream_expectSize_Results{st}, nil
}

func ReadRootByteStream_expectSize_Results(msg *capnp.Message) (ByteStream_expectSize_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return ByteStream_expectSize_Results{}, err
	}
	return ByteStream_expectSize_Results{root.Struct()}, nil
}

// ByteStream_expectSize_Results_List is a list of ByteStream_expectSize_Results.
type ByteStream_expectSize_Results_List struct{ capnp.List }

// NewByteStream_expectSize_Results creates a new list of ByteStream_expectSize_Results.
func NewByteStream_expectSize_Results_List(s *capnp.Segment, sz int32) (ByteStream_expectSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return ByteStream_expectSize_Results_List{}, err
	}
	return ByteStream_expectSize_Results_List{l}, nil
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
	if err != nil {
		return Assignable_Getter_get_Params{}, err
	}
	return Assignable_Getter_get_Params{st}, nil
}

func NewRootAssignable_Getter_get_Params(s *capnp.Segment) (Assignable_Getter_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Assignable_Getter_get_Params{}, err
	}
	return Assignable_Getter_get_Params{st}, nil
}

func ReadRootAssignable_Getter_get_Params(msg *capnp.Message) (Assignable_Getter_get_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Getter_get_Params{}, err
	}
	return Assignable_Getter_get_Params{root.Struct()}, nil
}

// Assignable_Getter_get_Params_List is a list of Assignable_Getter_get_Params.
type Assignable_Getter_get_Params_List struct{ capnp.List }

// NewAssignable_Getter_get_Params creates a new list of Assignable_Getter_get_Params.
func NewAssignable_Getter_get_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Assignable_Getter_get_Params_List{}, err
	}
	return Assignable_Getter_get_Params_List{l}, nil
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
	if err != nil {
		return Assignable_Getter_get_Results{}, err
	}
	return Assignable_Getter_get_Results{st}, nil
}

func NewRootAssignable_Getter_get_Results(s *capnp.Segment) (Assignable_Getter_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_Getter_get_Results{}, err
	}
	return Assignable_Getter_get_Results{st}, nil
}

func ReadRootAssignable_Getter_get_Results(msg *capnp.Message) (Assignable_Getter_get_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Getter_get_Results{}, err
	}
	return Assignable_Getter_get_Results{root.Struct()}, nil
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
	if err != nil {
		return Assignable_Getter_get_Results_List{}, err
	}
	return Assignable_Getter_get_Results_List{l}, nil
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
	if err != nil {
		return Assignable_Getter_subscribe_Params{}, err
	}
	return Assignable_Getter_subscribe_Params{st}, nil
}

func NewRootAssignable_Getter_subscribe_Params(s *capnp.Segment) (Assignable_Getter_subscribe_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_Getter_subscribe_Params{}, err
	}
	return Assignable_Getter_subscribe_Params{st}, nil
}

func ReadRootAssignable_Getter_subscribe_Params(msg *capnp.Message) (Assignable_Getter_subscribe_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Getter_subscribe_Params{}, err
	}
	return Assignable_Getter_subscribe_Params{root.Struct()}, nil
}
func (s Assignable_Getter_subscribe_Params) Setter() Assignable_Setter {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Assignable_Setter{}
	}
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Params) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Params) SetSetter(v Assignable_Setter) error {
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

// Assignable_Getter_subscribe_Params_List is a list of Assignable_Getter_subscribe_Params.
type Assignable_Getter_subscribe_Params_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Params creates a new list of Assignable_Getter_subscribe_Params.
func NewAssignable_Getter_subscribe_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Assignable_Getter_subscribe_Params_List{}, err
	}
	return Assignable_Getter_subscribe_Params_List{l}, nil
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
	if err != nil {
		return Assignable_Getter_subscribe_Results{}, err
	}
	return Assignable_Getter_subscribe_Results{st}, nil
}

func NewRootAssignable_Getter_subscribe_Results(s *capnp.Segment) (Assignable_Getter_subscribe_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_Getter_subscribe_Results{}, err
	}
	return Assignable_Getter_subscribe_Results{st}, nil
}

func ReadRootAssignable_Getter_subscribe_Results(msg *capnp.Message) (Assignable_Getter_subscribe_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Getter_subscribe_Results{}, err
	}
	return Assignable_Getter_subscribe_Results{root.Struct()}, nil
}
func (s Assignable_Getter_subscribe_Results) Handle() Handle {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Handle{}
	}
	return Handle{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Results) SetHandle(v Handle) error {
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

// Assignable_Getter_subscribe_Results_List is a list of Assignable_Getter_subscribe_Results.
type Assignable_Getter_subscribe_Results_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Results creates a new list of Assignable_Getter_subscribe_Results.
func NewAssignable_Getter_subscribe_Results_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Assignable_Getter_subscribe_Results_List{}, err
	}
	return Assignable_Getter_subscribe_Results_List{l}, nil
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
	if err != nil {
		return Assignable_Setter_set_Params{}, err
	}
	return Assignable_Setter_set_Params{st}, nil
}

func NewRootAssignable_Setter_set_Params(s *capnp.Segment) (Assignable_Setter_set_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_Setter_set_Params{}, err
	}
	return Assignable_Setter_set_Params{st}, nil
}

func ReadRootAssignable_Setter_set_Params(msg *capnp.Message) (Assignable_Setter_set_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Setter_set_Params{}, err
	}
	return Assignable_Setter_set_Params{root.Struct()}, nil
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
	if err != nil {
		return Assignable_Setter_set_Params_List{}, err
	}
	return Assignable_Setter_set_Params_List{l}, nil
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
	if err != nil {
		return Assignable_Setter_set_Results{}, err
	}
	return Assignable_Setter_set_Results{st}, nil
}

func NewRootAssignable_Setter_set_Results(s *capnp.Segment) (Assignable_Setter_set_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Assignable_Setter_set_Results{}, err
	}
	return Assignable_Setter_set_Results{st}, nil
}

func ReadRootAssignable_Setter_set_Results(msg *capnp.Message) (Assignable_Setter_set_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_Setter_set_Results{}, err
	}
	return Assignable_Setter_set_Results{root.Struct()}, nil
}

// Assignable_Setter_set_Results_List is a list of Assignable_Setter_set_Results.
type Assignable_Setter_set_Results_List struct{ capnp.List }

// NewAssignable_Setter_set_Results creates a new list of Assignable_Setter_set_Results.
func NewAssignable_Setter_set_Results_List(s *capnp.Segment, sz int32) (Assignable_Setter_set_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Assignable_Setter_set_Results_List{}, err
	}
	return Assignable_Setter_set_Results_List{l}, nil
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
	if err != nil {
		return Assignable_get_Params{}, err
	}
	return Assignable_get_Params{st}, nil
}

func NewRootAssignable_get_Params(s *capnp.Segment) (Assignable_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Assignable_get_Params{}, err
	}
	return Assignable_get_Params{st}, nil
}

func ReadRootAssignable_get_Params(msg *capnp.Message) (Assignable_get_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_get_Params{}, err
	}
	return Assignable_get_Params{root.Struct()}, nil
}

// Assignable_get_Params_List is a list of Assignable_get_Params.
type Assignable_get_Params_List struct{ capnp.List }

// NewAssignable_get_Params creates a new list of Assignable_get_Params.
func NewAssignable_get_Params_List(s *capnp.Segment, sz int32) (Assignable_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Assignable_get_Params_List{}, err
	}
	return Assignable_get_Params_List{l}, nil
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
	if err != nil {
		return Assignable_get_Results{}, err
	}
	return Assignable_get_Results{st}, nil
}

func NewRootAssignable_get_Results(s *capnp.Segment) (Assignable_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	if err != nil {
		return Assignable_get_Results{}, err
	}
	return Assignable_get_Results{st}, nil
}

func ReadRootAssignable_get_Results(msg *capnp.Message) (Assignable_get_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_get_Results{}, err
	}
	return Assignable_get_Results{root.Struct()}, nil
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
	p, err := s.Struct.Ptr(1)
	if err != nil {

		return Assignable_Setter{}
	}
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_get_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Assignable_get_Results) SetSetter(v Assignable_Setter) error {
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

// Assignable_get_Results_List is a list of Assignable_get_Results.
type Assignable_get_Results_List struct{ capnp.List }

// NewAssignable_get_Results creates a new list of Assignable_get_Results.
func NewAssignable_get_Results_List(s *capnp.Segment, sz int32) (Assignable_get_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	if err != nil {
		return Assignable_get_Results_List{}, err
	}
	return Assignable_get_Results_List{l}, nil
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
	if err != nil {
		return Assignable_asGetter_Params{}, err
	}
	return Assignable_asGetter_Params{st}, nil
}

func NewRootAssignable_asGetter_Params(s *capnp.Segment) (Assignable_asGetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Assignable_asGetter_Params{}, err
	}
	return Assignable_asGetter_Params{st}, nil
}

func ReadRootAssignable_asGetter_Params(msg *capnp.Message) (Assignable_asGetter_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_asGetter_Params{}, err
	}
	return Assignable_asGetter_Params{root.Struct()}, nil
}

// Assignable_asGetter_Params_List is a list of Assignable_asGetter_Params.
type Assignable_asGetter_Params_List struct{ capnp.List }

// NewAssignable_asGetter_Params creates a new list of Assignable_asGetter_Params.
func NewAssignable_asGetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Assignable_asGetter_Params_List{}, err
	}
	return Assignable_asGetter_Params_List{l}, nil
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
	if err != nil {
		return Assignable_asGetter_Results{}, err
	}
	return Assignable_asGetter_Results{st}, nil
}

func NewRootAssignable_asGetter_Results(s *capnp.Segment) (Assignable_asGetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_asGetter_Results{}, err
	}
	return Assignable_asGetter_Results{st}, nil
}

func ReadRootAssignable_asGetter_Results(msg *capnp.Message) (Assignable_asGetter_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_asGetter_Results{}, err
	}
	return Assignable_asGetter_Results{root.Struct()}, nil
}
func (s Assignable_asGetter_Results) Getter() Assignable_Getter {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Assignable_Getter{}
	}
	return Assignable_Getter{Client: p.Interface().Client()}
}

func (s Assignable_asGetter_Results) HasGetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asGetter_Results) SetGetter(v Assignable_Getter) error {
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

// Assignable_asGetter_Results_List is a list of Assignable_asGetter_Results.
type Assignable_asGetter_Results_List struct{ capnp.List }

// NewAssignable_asGetter_Results creates a new list of Assignable_asGetter_Results.
func NewAssignable_asGetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Assignable_asGetter_Results_List{}, err
	}
	return Assignable_asGetter_Results_List{l}, nil
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
	if err != nil {
		return Assignable_asSetter_Params{}, err
	}
	return Assignable_asSetter_Params{st}, nil
}

func NewRootAssignable_asSetter_Params(s *capnp.Segment) (Assignable_asSetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Assignable_asSetter_Params{}, err
	}
	return Assignable_asSetter_Params{st}, nil
}

func ReadRootAssignable_asSetter_Params(msg *capnp.Message) (Assignable_asSetter_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_asSetter_Params{}, err
	}
	return Assignable_asSetter_Params{root.Struct()}, nil
}

// Assignable_asSetter_Params_List is a list of Assignable_asSetter_Params.
type Assignable_asSetter_Params_List struct{ capnp.List }

// NewAssignable_asSetter_Params creates a new list of Assignable_asSetter_Params.
func NewAssignable_asSetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Assignable_asSetter_Params_List{}, err
	}
	return Assignable_asSetter_Params_List{l}, nil
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
	if err != nil {
		return Assignable_asSetter_Results{}, err
	}
	return Assignable_asSetter_Results{st}, nil
}

func NewRootAssignable_asSetter_Results(s *capnp.Segment) (Assignable_asSetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Assignable_asSetter_Results{}, err
	}
	return Assignable_asSetter_Results{st}, nil
}

func ReadRootAssignable_asSetter_Results(msg *capnp.Message) (Assignable_asSetter_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Assignable_asSetter_Results{}, err
	}
	return Assignable_asSetter_Results{root.Struct()}, nil
}
func (s Assignable_asSetter_Results) Setter() Assignable_Setter {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Assignable_Setter{}
	}
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_asSetter_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asSetter_Results) SetSetter(v Assignable_Setter) error {
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

// Assignable_asSetter_Results_List is a list of Assignable_asSetter_Results.
type Assignable_asSetter_Results_List struct{ capnp.List }

// NewAssignable_asSetter_Results creates a new list of Assignable_asSetter_Results.
func NewAssignable_asSetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Assignable_asSetter_Results_List{}, err
	}
	return Assignable_asSetter_Results_List{l}, nil
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
