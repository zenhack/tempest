package sandstormhttpbridgeinternal

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	websession "zenhack.net/go/sandstorm/capnp/websession"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

// Constants defined in sandstorm-http-bridge-internal.capnp.
const (
	BridgeRequestSessionHtml = "<!DOCTYPE html>\n\n<html>\n  <head>\n    <style type=\"text/css\">\n      * {\n        box-sizing: border-box;\n      }\n      body {\n        font-family: sans-serif;\n        font-size: 20px;\n      }\n      button {\n        border: none;\n        font-size: inherit;\n        font-family: inherit;\n        font-weight: inherit;\n        text-decoration: inherit;\n        color: inherit;\n        line-height: inherit;\n        background-color: transparent;\n        text-align: inherit;\n        padding: 0;\n        cursor: pointer;\n        display: block;\n        overflow: hidden;\n        text-overflow: ellipsis;\n        white-space: nowrap;\n        width: 100%;\n        padding-left: 32px;\n        height: 31px;\n      }\n      li {\n        border: 1px solid #ddd;\n        border-bottom: none;\n        background-color: #eee;\n        vertical-align: middle;\n        height: 32px;\n      }\n      ul {\n        border-bottom: 1px solid #ddd;\n        padding: 0;\n        margin: 10px;\n        list-style-type: none;\n      }\n    </style>\n  </head>\n  <body>\n    <ul id=\"list\">\n    </ul>\n\n    <script type=\"text/javascript\">\n      var config = @CONFIG@;\n\n      function makeClickHandler(name) {\n        return function () {\n          var xhr = new XMLHttpRequest();\n          xhr.onload = function () {\n            if (xhr.status >= 400) {\n              alert(\"XHR returned status \" + xhr.status + \":\\n\" + xhr.responseText);\n            }\n          }\n          xhr.onerror = function(e) { alert(e); };\n          xhr.open(\"post\", \"/\");\n          xhr.send(name);\n        }\n      }\n\n      var list = document.getElementById(\"list\");\n      for (var i in config) {\n        var api = config[i];\n\n        var button = document.createElement(\"button\");\n        button.addEventListener(\"click\", makeClickHandler(api.name));\n        if (api.displayInfo && api.displayInfo.title && api.displayInfo.title.defaultText) {\n          button.textContent = api.displayInfo.title.defaultText;\n        } else {\n          button.textContent = \"Use this grain\";\n        }\n\n        var item = document.createElement(\"li\");\n        item.appendChild(button);\n\n        list.appendChild(item);\n      }\n    </script>\n  </body>\n</html>\n"
)

type BridgeObjectId struct{ capnp.Struct }
type BridgeObjectId_Which uint16

const (
	BridgeObjectId_Which_application BridgeObjectId_Which = 0
	BridgeObjectId_Which_httpApi     BridgeObjectId_Which = 1
)

func (w BridgeObjectId_Which) String() string {
	const s = "applicationhttpApi"
	switch w {
	case BridgeObjectId_Which_application:
		return s[0:11]
	case BridgeObjectId_Which_httpApi:
		return s[11:18]

	}
	return "BridgeObjectId_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// BridgeObjectId_TypeID is the unique identifier for the type BridgeObjectId.
const BridgeObjectId_TypeID = 0xde7c54260c265bb4

func NewBridgeObjectId(s *capnp.Segment) (BridgeObjectId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return BridgeObjectId{st}, err
}

func NewRootBridgeObjectId(s *capnp.Segment) (BridgeObjectId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return BridgeObjectId{st}, err
}

func ReadRootBridgeObjectId(msg *capnp.Message) (BridgeObjectId, error) {
	root, err := msg.RootPtr()
	return BridgeObjectId{root.Struct()}, err
}

func (s BridgeObjectId) String() string {
	str, _ := text.Marshal(0xde7c54260c265bb4, s.Struct)
	return str
}

func (s BridgeObjectId) Which() BridgeObjectId_Which {
	return BridgeObjectId_Which(s.Struct.Uint16(0))
}
func (s BridgeObjectId) Application() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s BridgeObjectId) HasApplication() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId) ApplicationPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s BridgeObjectId) SetApplication(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPointer(0, v)
}

func (s BridgeObjectId) SetApplicationPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v)
}

func (s BridgeObjectId) HttpApi() (BridgeObjectId_HttpApi, error) {
	p, err := s.Struct.Ptr(0)
	return BridgeObjectId_HttpApi{Struct: p.Struct()}, err
}

func (s BridgeObjectId) HasHttpApi() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId) SetHttpApi(v BridgeObjectId_HttpApi) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewHttpApi sets the httpApi field to a newly
// allocated BridgeObjectId_HttpApi struct, preferring placement in s's segment.
func (s BridgeObjectId) NewHttpApi() (BridgeObjectId_HttpApi, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewBridgeObjectId_HttpApi(s.Struct.Segment())
	if err != nil {
		return BridgeObjectId_HttpApi{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// BridgeObjectId_List is a list of BridgeObjectId.
type BridgeObjectId_List struct{ capnp.List }

// NewBridgeObjectId creates a new list of BridgeObjectId.
func NewBridgeObjectId_List(s *capnp.Segment, sz int32) (BridgeObjectId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return BridgeObjectId_List{l}, err
}

func (s BridgeObjectId_List) At(i int) BridgeObjectId { return BridgeObjectId{s.List.Struct(i)} }

func (s BridgeObjectId_List) Set(i int, v BridgeObjectId) error { return s.List.SetStruct(i, v.Struct) }

// BridgeObjectId_Promise is a wrapper for a BridgeObjectId promised by a client call.
type BridgeObjectId_Promise struct{ *capnp.Pipeline }

func (p BridgeObjectId_Promise) Struct() (BridgeObjectId, error) {
	s, err := p.Pipeline.Struct()
	return BridgeObjectId{s}, err
}

func (p BridgeObjectId_Promise) Application() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p BridgeObjectId_Promise) HttpApi() BridgeObjectId_HttpApi_Promise {
	return BridgeObjectId_HttpApi_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type BridgeObjectId_HttpApi struct{ capnp.Struct }

// BridgeObjectId_HttpApi_TypeID is the unique identifier for the type BridgeObjectId_HttpApi.
const BridgeObjectId_HttpApi_TypeID = 0x903896a2654fb12b

func NewBridgeObjectId_HttpApi(s *capnp.Segment) (BridgeObjectId_HttpApi, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return BridgeObjectId_HttpApi{st}, err
}

func NewRootBridgeObjectId_HttpApi(s *capnp.Segment) (BridgeObjectId_HttpApi, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return BridgeObjectId_HttpApi{st}, err
}

func ReadRootBridgeObjectId_HttpApi(msg *capnp.Message) (BridgeObjectId_HttpApi, error) {
	root, err := msg.RootPtr()
	return BridgeObjectId_HttpApi{root.Struct()}, err
}

func (s BridgeObjectId_HttpApi) String() string {
	str, _ := text.Marshal(0x903896a2654fb12b, s.Struct)
	return str
}

func (s BridgeObjectId_HttpApi) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s BridgeObjectId_HttpApi) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId_HttpApi) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s BridgeObjectId_HttpApi) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s BridgeObjectId_HttpApi) Path() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s BridgeObjectId_HttpApi) HasPath() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId_HttpApi) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s BridgeObjectId_HttpApi) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s BridgeObjectId_HttpApi) Permissions() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.BitList{List: p.List()}, err
}

func (s BridgeObjectId_HttpApi) HasPermissions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId_HttpApi) SetPermissions(v capnp.BitList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s BridgeObjectId_HttpApi) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s BridgeObjectId_HttpApi) IdentityId() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s BridgeObjectId_HttpApi) HasIdentityId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s BridgeObjectId_HttpApi) SetIdentityId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

// BridgeObjectId_HttpApi_List is a list of BridgeObjectId_HttpApi.
type BridgeObjectId_HttpApi_List struct{ capnp.List }

// NewBridgeObjectId_HttpApi creates a new list of BridgeObjectId_HttpApi.
func NewBridgeObjectId_HttpApi_List(s *capnp.Segment, sz int32) (BridgeObjectId_HttpApi_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return BridgeObjectId_HttpApi_List{l}, err
}

func (s BridgeObjectId_HttpApi_List) At(i int) BridgeObjectId_HttpApi {
	return BridgeObjectId_HttpApi{s.List.Struct(i)}
}

func (s BridgeObjectId_HttpApi_List) Set(i int, v BridgeObjectId_HttpApi) error {
	return s.List.SetStruct(i, v.Struct)
}

// BridgeObjectId_HttpApi_Promise is a wrapper for a BridgeObjectId_HttpApi promised by a client call.
type BridgeObjectId_HttpApi_Promise struct{ *capnp.Pipeline }

func (p BridgeObjectId_HttpApi_Promise) Struct() (BridgeObjectId_HttpApi, error) {
	s, err := p.Pipeline.Struct()
	return BridgeObjectId_HttpApi{s}, err
}

type BridgeHttpSession struct{ Client capnp.Client }

func (c BridgeHttpSession) Get(ctx context.Context, params func(websession.WebSession_get_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_get_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Post(ctx context.Context, params func(websession.WebSession_post_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_post_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) OpenWebSocket(ctx context.Context, params func(websession.WebSession_openWebSocket_Params) error, opts ...capnp.CallOption) websession.WebSession_openWebSocket_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_openWebSocket_Params{Struct: s}) }
	}
	return websession.WebSession_openWebSocket_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Put(ctx context.Context, params func(websession.WebSession_put_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_put_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Delete(ctx context.Context, params func(websession.WebSession_delete_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_delete_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) PostStreaming(ctx context.Context, params func(websession.WebSession_postStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_postStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_postStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_postStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) PutStreaming(ctx context.Context, params func(websession.WebSession_putStreaming_Params) error, opts ...capnp.CallOption) websession.WebSession_putStreaming_Results_Promise {
	if c.Client == nil {
		return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_putStreaming_Params{Struct: s}) }
	}
	return websession.WebSession_putStreaming_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Propfind(ctx context.Context, params func(websession.WebSession_propfind_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_propfind_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Proppatch(ctx context.Context, params func(websession.WebSession_proppatch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_proppatch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Mkcol(ctx context.Context, params func(websession.WebSession_mkcol_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_mkcol_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Copy(ctx context.Context, params func(websession.WebSession_copy_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_copy_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Move(ctx context.Context, params func(websession.WebSession_move_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_move_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Lock(ctx context.Context, params func(websession.WebSession_lock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_lock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Unlock(ctx context.Context, params func(websession.WebSession_unlock_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_unlock_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Acl(ctx context.Context, params func(websession.WebSession_acl_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_acl_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Report(ctx context.Context, params func(websession.WebSession_report_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_report_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Options(ctx context.Context, params func(websession.WebSession_options_Params) error, opts ...capnp.CallOption) websession.WebSession_Options_Promise {
	if c.Client == nil {
		return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_options_Params{Struct: s}) }
	}
	return websession.WebSession_Options_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Patch(ctx context.Context, params func(websession.WebSession_patch_Params) error, opts ...capnp.CallOption) websession.WebSession_Response_Promise {
	if c.Client == nil {
		return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(websession.WebSession_patch_Params{Struct: s}) }
	}
	return websession.WebSession_Response_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c BridgeHttpSession) Save(ctx context.Context, params func(grain.AppPersistent_save_Params) error, opts ...capnp.CallOption) grain.AppPersistent_save_Results_Promise {
	if c.Client == nil {
		return grain.AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.AppPersistent_save_Params{Struct: s}) }
	}
	return grain.AppPersistent_save_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type BridgeHttpSession_Server interface {
	Get(websession.WebSession_get) error

	Post(websession.WebSession_post) error

	OpenWebSocket(websession.WebSession_openWebSocket) error

	Put(websession.WebSession_put) error

	Delete(websession.WebSession_delete) error

	PostStreaming(websession.WebSession_postStreaming) error

	PutStreaming(websession.WebSession_putStreaming) error

	Propfind(websession.WebSession_propfind) error

	Proppatch(websession.WebSession_proppatch) error

	Mkcol(websession.WebSession_mkcol) error

	Copy(websession.WebSession_copy) error

	Move(websession.WebSession_move) error

	Lock(websession.WebSession_lock) error

	Unlock(websession.WebSession_unlock) error

	Acl(websession.WebSession_acl) error

	Report(websession.WebSession_report) error

	Options(websession.WebSession_options) error

	Patch(websession.WebSession_patch) error

	Save(grain.AppPersistent_save) error
}

func BridgeHttpSession_ServerToClient(s BridgeHttpSession_Server) BridgeHttpSession {
	c, _ := s.(server.Closer)
	return BridgeHttpSession{Client: server.New(BridgeHttpSession_Methods(nil, s), c)}
}

func BridgeHttpSession_Methods(methods []server.Method, s BridgeHttpSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 19)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      0,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_get{c, opts, websession.WebSession_get_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      1,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "post",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_post{c, opts, websession.WebSession_post_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Post(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      2,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "openWebSocket",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_openWebSocket{c, opts, websession.WebSession_openWebSocket_Params{Struct: p}, websession.WebSession_openWebSocket_Results{Struct: r}}
			return s.OpenWebSocket(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      3,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "put",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_put{c, opts, websession.WebSession_put_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Put(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      4,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "delete",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_delete{c, opts, websession.WebSession_delete_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Delete(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      5,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "postStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_postStreaming{c, opts, websession.WebSession_postStreaming_Params{Struct: p}, websession.WebSession_postStreaming_Results{Struct: r}}
			return s.PostStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      6,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "putStreaming",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_putStreaming{c, opts, websession.WebSession_putStreaming_Params{Struct: p}, websession.WebSession_putStreaming_Results{Struct: r}}
			return s.PutStreaming(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      7,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "propfind",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_propfind{c, opts, websession.WebSession_propfind_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Propfind(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      8,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "proppatch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_proppatch{c, opts, websession.WebSession_proppatch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Proppatch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      9,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "mkcol",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_mkcol{c, opts, websession.WebSession_mkcol_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Mkcol(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      10,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "copy",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_copy{c, opts, websession.WebSession_copy_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Copy(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      11,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "move",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_move{c, opts, websession.WebSession_move_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Move(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      12,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "lock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_lock{c, opts, websession.WebSession_lock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Lock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      13,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "unlock",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_unlock{c, opts, websession.WebSession_unlock_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Unlock(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      14,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "acl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_acl{c, opts, websession.WebSession_acl_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Acl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      15,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "report",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_report{c, opts, websession.WebSession_report_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Report(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      16,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "options",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_options{c, opts, websession.WebSession_options_Params{Struct: p}, websession.WebSession_Options{Struct: r}}
			return s.Options(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa50711a14d35a8ce,
			MethodID:      17,
			InterfaceName: "web-session.capnp:WebSession",
			MethodName:    "patch",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := websession.WebSession_patch{c, opts, websession.WebSession_patch_Params{Struct: p}, websession.WebSession_Response{Struct: r}}
			return s.Patch(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 9},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaffa789add8747b8,
			MethodID:      0,
			InterfaceName: "grain.capnp:AppPersistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.AppPersistent_save{c, opts, grain.AppPersistent_save_Params{Struct: p}, grain.AppPersistent_save_Results{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	return methods
}

const schema_f963cc483d8f9e3a = "x\xda\x94V\xcfo$G\x15\xae\xd7mg\xa2\xd5\xae" +
	"\xe3\xc9\x8c\x94\x13\xa4\x98\xc8\xcbxwgl\xaf\xf7\xb0" +
	"\x1a{\xc2f\x9d%6JXw\xbcHK\xb2H\xd4" +
	"t\xbd\x99\xaelMUSUc{\x02\xab\x80@\xca" +
	"\x05\x09\xb8\x00\xe2\xc0!\x7f\x00(\x12\x12?\xfe\x00\xc8" +
	"\x81\xdc\xb8\xc2\x01\x0ep\xe6\x06B4\xaa\xe9\xb6\xa7=" +
	"3\xab\x04\xf9\xe0\x9e\xf7\xbdz\xef\xab\xef\xfdPm\x8a" +
	"\xe5;\xc1\xd6\xf2O>C\xc8\xd1CX~&\xfb\xc5" +
	"\xd8\xfd\xe3\xb7?}\xeb\xbb\xa4\xba\xb6\x94u~\xfe\x83" +
	"\xee\xfe\x1f\xe3\x7f\x11\x02\xdb\x1f\xae\xfc\x13j\x1f\xafT" +
	"\x089\xfah%\x84\xa3?\xad\x04@Hv\xfd\xc3\xfb" +
	"\xf8\xc1\x8fo\xff\x90ToA\xf6\xab\xb7\xaf^\xbe\xfa" +
	"\xe0[\x7f!\xcbK\x15B\xb6\xaf\xd4\xfe\x0a\xb5\xf5Z" +
	"\x85\x90\xdaZ\xed\xef\x04\xb2\xff\\\xbb\xf6\xe8G\xb7?" +
	"\xfbkR]\x0b/\x04\xaf\xd6\x7f\x07\xb5\xf5\xfa\xc4\xb3" +
	"^\xa9\xad\xd5?O\xc84Z\xb4\x060u\xbf\x07\x95" +
	"\x80\x90\xed7\xea\x1f@\x0d\xeb/\x10R\x1b\xd6OH" +
	"6\xf9{?\xb3Lq\xeb\xb4yf\xd8J\x9cK[" +
	"=#\xf8\x00[B94\x8a\xc9v\xccR\x95vr" +
	"\xeb\x9b\xf8\x8d\x11Zw\x84\xd6\x0a\xfd\xa2\xdawCy" +
	"\x08\x00\x97I\x00\x97\x09\xd9\x82/\xbd\x9a\xed~\xee\xd5" +
	"\xfb{\x0f\xbez\x98\xdd\xa3\x89\x1b\xca\x97/]\xda\xcd" +
	"\xffS\xba\x9b \xe3\xfe\x83\xd2]\xeb\xc6\x12\xa9\x1b\xa7" +
	"\xd8m8<u\x1b\xb1\xb5\x8d\x1c\xa3\xf4\x1a\xfdf\xf1" +
	"EiO\x9f\xb6\xacxW\xa8A\x87\xf6\xb4\xe1hZ" +
	"=}\xbaS\xe0O.\x9dy\xf1q\xe9P_+\xd7" +
	"\xea\xb3\xa1\x90\xe3\x0e\xb5L\xd9\x96E#\xfa;\x17\x1d" +
	"\xacx\x17;\xf4\xe6f:\x1fn\xe4\x9cV\x17X\xf8" +
	"\xcc\x1d\xaa\xb4\xc2\x85Q\x84J\xd0\x08\xb7\xb3\x98\xc2b" +
	"\xf4\x04\xc5 q\x0bP/H\x8bc\xac\x0dsB\xab" +
	"\x05\x1e\xb1\x96\xda,\xb0K\xa1\xb0\x95<-n\x8f\xc5" +
	"\x8f\x07F\x8f\x14o\x15\x01\x9ca\xca\xa6\xcc\xa0\x9aM" +
	"\xcf\xa4\x18,\xca\x9c2\xce'\xc5\xd8,\xb1\x19\x19\xeb" +
	"\xa3\xa5z\xd25S\x80\x0b\x9bJ6\xee\xd0\x9e\xd4\xf1" +
	"\xe3\xa9]\x1f\xa3\xe9K}\xd2\xa1\x89\xe0\x1c\xd5L\xee" +
	")\x8cR\x8a\xd4\x0a;u8I\x84\xc3\x96MY\x8c" +
	"\xbe\x1c'\x86\xa5%Pp\x97t\xe8\xd6\xe6\xe6\xda\x1c" +
	"\xe3\x96\xc4\xbe\xeb\xd0\xed\x9b\xd3rSz&\xd5\xf6\xd6" +
	"|\x13H\xb1\xa0\x01\xb6\xd2Sj\xb5\x14\x9c\xbe\xc49" +
	"\xdf\x99\xc1[=\xed\x9c\x1e\xce\xf6\xc9\xbc\xee/!\x96" +
	"\xf0c4N\xc4L\x9e\x89>\x14\x9cK\\\xc0\xf2\xe6" +
	"<\xcb\x91\x9ccy\xce\xe2id\x17\xd5p\xc8\xcc@" +
	"(\xaf]Y\x1f)\xackMf\xb5\xe5g\xf5\xe2\xc5" +
	"r\x0a\xbb\x1b\x13|2\xe0\x1bg\x13\xbe\xeb\x07\xb2\x18" +
	"\xf5\x91\xa4\x82w\x1b>T1\xe1\xbb\x1b#\xbf\x18\xf2" +
	"E\x10\x1b\x91\xba\xf2&x\x87\x1d\xb3\xdcz\xbe\x10\x8e" +
	"\x99\xa1\xb1V}1\xa0]zg\xef\xfe\x97\xbfx\xf0" +
	"\xda\x9d\x9dK\x05\xda\x1f\xa9\xd8\xcf\x09\x1d\xb2\xc7\xb8'" +
	"E\xfcx\x9f).\xd14\x15\x1b\xe2zI\x1f\x83n" +
	"d\xd4\xf4@\xb3\x0c\xe6iN\x13C\xbbT\xe1\x09}" +
	"\xf8\xc6\xeb\xfb\xce\xa5\xc5\xcek\xae\xef\x94<O\x13\xd3" +
	"\xd6Jj\xc6i\xf7\xa9\xe1(\x15}\xda\xf4\xae\xd61" +
	"7\xb2\xf4\xe5.\xbd\xb5\xb99\xebD)\x93h\\\xb3" +
	"\xf1p\xff\xcd\x82!rZ\x1ci\xd0\xeb\xb4\x14\xe1:" +
	"mt\x1e\xa93\xa3A\x9bje\xf1\x01\x9e\xba\x0b\xf4" +
	"\xa6\xed1\xfb\x9d\x13Gc\xb4)1oz\x95\x0a\x1a" +
	"\xb8\xbeC\x9f\xcc\xdd5E\xd5l\xa4\xda\xba\xc6\x0d\xda" +
	"\xd8h\xcc\x89aQ\xf1\\\xee)\xf2\xe4\xbcOJe" +
	"\xf4m@\xbb\x94\xebx4D\xe5\xda\x03t\xf7$\xfa" +
	"\xcf\xbb\xe3\x03\xde\xcc\xdb\xe4<H_\x1b\xda\xf4\xc7\x04" +
	"\x15\xaah\x81\xb2~\x1eb\xa9\xa0\xdd\x02{[|\xed" +
	"\xbc-r\xb4X\xe5\xa5\x8c\xb1A\xe6\xb0H\xdal\xe4" +
	"\x0e\xe5\x1b\xe5\x966\xe3\xfc\xde1*\xf7\xba\xb0\x0e\x15" +
	"\x9af#\xf6\xbd\xd5\xb81\xdfg,\x15\xed\xc9\xe5K" +
	"Q|\xed\xbd\xbd\xd8\x81\x07\xaa\xaf\xe9\xd5\xabt\xc6\xd4" +
	"v\xc2I|*\xd0\xe6\xd8g#\xe9&\x15\xbe\xd07" +
	"\x05I?1{Z9T^\xd5O\x8cQ*\x0eE" +
	"i\xf1\x93C6\xbeb\x91\xbaDX:0L\xa8F" +
	"\xb9\xbc\x17\x84\x16\x0e\x87^\xe6\xecL\xe7\xd5\x19\xa1\xa5" +
	"(\x8b\xec\xdd\xdb,MQ\xf1\xbdDH\xde\xcc\xb3\xaf" +
	"\x97\xca\xe7[\xe1\x82\x8b?\xb3>\xb7}&\xab\"_" +
	"?\xf9\xd6\xd9\xdd\xc8_\x1c\x84|\xca\xc7\xcd\xdd\x89\xf5" +
	"~\xef\x1d\x8c\xdd\x01o\xfb\xa9\xaf\xbc\x92\x8aC\x80h" +
	"5\\\"d\x09\x08\xa9\xb2k\x84D\x8fB\x88\x92\x00" +
	"\xaa\x00u\xf0F\xf4\xc6\xaf\x87\x10\xc9\x00\xaaAP\x87" +
	"\x80\x90\xaa\xe8\x11\x12%!D\xdf\x0b\xa0\x1a\x86u\x08" +
	"\x09\xa9~\xe7-B\xa2o\x87\x10}?\x80\xe7|\xaf" +
	"\x14\xaf&x.e.9\xfb\x91\xa5h\x86\xc2ZA" +
	"*ZYX!p\x18\x02\x00\x09\xfcg&8*'" +
	"\xdc\x98\x84\x07\x1c\xae\x90\x00\xae\x108\xbf\xe2\xf2\xa7\xb8" +
	"\xa2\xbf\xd9\xd1\x17&\xaf7u\x08p\x18.G\xcf\x02" +
	"d\x7f\xf8x\xed\xf7\xe3\xbf\x8d?\"\x84d\xbfy\xed" +
	"\xfd?\xff\xec\xf4\xdf\xbf$w\xa0\x0a/FK\x01L" +
	"M\x84T\xe1\x85h\x09\x00&\xb4V\xa7\x0f\xce\xff\x93" +
	"\xc7D\xea\x8a;\xe0>X\xe9E\x0cw\xdf\xf3\x14_" +
	"IE\xf4l\xb8t9\xcb&\xca\xaf{=\x9b!D" +
	"\xb7\x02\xb8\x02\xff\xcdr\xe9\xb7\xee\x12\x12\xdd\x08!\xba" +
	"\x1d@\xc6\xd2T\x8a\x989R\x11Z\xc1\xf3$\x80\xe7" +
	"\x09\xbc\x97\xe4\xb1`u\x9a\x82\x00\xac\x12\xf8_\x00\x00" +
	"\x00\xff\xff\x82\"W\xaf"

func init() {
	schemas.Register(schema_f963cc483d8f9e3a,
		0x835a98b9e87479ae,
		0x903896a2654fb12b,
		0xb71e38915c2a2afc,
		0xde7c54260c265bb4)
}
