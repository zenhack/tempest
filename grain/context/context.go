package context

import (
	"context"
	"errors"

	"zenhack.net/go/sandstorm/capnp/grain"
)

type ctxKey int

type typeIdKey uint64

const (
	sessionCtxKey ctxKey = iota
)

var (
	WrongSessionType = errors.New("Wrong session type")
)

// Add `sessionCtx` to the values of `parent`, returning the new context.
//
// `sessionCtx` can be retrieved using `GetSessionContext`, below.
func WithSessionContext(parent context.Context, sessionCtx grain.SessionContext) context.Context {
	return context.WithValue(parent, sessionCtxKey, sessionCtx)
}

// Retrieve the SessionContext from `ctx`'s values, which must have been previously set via
// `WithSessionContext`.
//
// panics if there is no SessionContext.
func GetSessionContext(ctx context.Context) grain.SessionContext {
	val := ctx.Value(sessionCtxKey)
	if val == nil {
		panic("No Session Context")
	}
	return val.(grain.SessionContext)
}

// Add a New*Session method's paramters to the context's Values.
//
// `p` must be one of the grain.UiView_new*Session_Params types.
// If it is not, `parent` will be returned unchanged.
//
// the value may be accessed with the Get*SessionParams functions, below.
func WithParams(parent context.Context, p interface{}) context.Context {
	var key typeIdKey
	switch p.(type) {
	case grain.UiView_newSession_Params:
		key = grain.UiView_newSession_Params_TypeID
	case grain.UiView_newRequestSession_Params:
		key = grain.UiView_newRequestSession_Params_TypeID
	case grain.UiView_newOfferSession_Params:
		key = grain.UiView_newOfferSession_Params_TypeID
	default:
		return parent
	}
	return context.WithValue(parent, key, p)
}

func getParams(ctx context.Context, key typeIdKey) (interface{}, error) {
	val := ctx.Value(key)
	if val == nil {
		return nil, WrongSessionType
	}
	return val, nil
}

func GetSessionParams(ctx context.Context) (p grain.UiView_newSession_Params, err error) {
	val, err := getParams(ctx, grain.UiView_newSession_Params_TypeID)
	if err != nil {
		return p, err
	}
	return val.(grain.UiView_newSession_Params), nil
}

func GetRequestSessionParams(ctx context.Context) (p grain.UiView_newRequestSession_Params, err error) {
	val, err := getParams(ctx, grain.UiView_newRequestSession_Params_TypeID)
	if err != nil {
		return p, err
	}
	return val.(grain.UiView_newRequestSession_Params), nil
}

func GetOfferSessionParams(ctx context.Context) (p grain.UiView_newOfferSession_Params, err error) {
	val, err := getParams(ctx, grain.UiView_newOfferSession_Params_TypeID)
	if err != nil {
		return p, err
	}
	return val.(grain.UiView_newOfferSession_Params), nil
}
