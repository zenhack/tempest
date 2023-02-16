package websession

import (
	"net/http"

	websession "zenhack.net/go/tempest/capnp/web-session"
)

// Mapping from web-session success codes to http status code numbers
var successCodeStatuses = map[websession.SuccessCode]int{
	websession.SuccessCode_ok:          http.StatusOK,
	websession.SuccessCode_created:     http.StatusCreated,
	websession.SuccessCode_accepted:    http.StatusAccepted,
	websession.SuccessCode_noContent:   http.StatusNoContent,
	websession.SuccessCode_multiStatus: http.StatusMultiStatus,
	websession.SuccessCode_notModified: http.StatusNotModified,
}

// Mapping from web-session error codes to http status code numbers
var clientErrorCodeStatuses = map[websession.ClientErrorCode]int{
	websession.ClientErrorCode_badRequest:            http.StatusBadRequest,
	websession.ClientErrorCode_forbidden:             http.StatusForbidden,
	websession.ClientErrorCode_notFound:              http.StatusNotFound,
	websession.ClientErrorCode_methodNotAllowed:      http.StatusMethodNotAllowed,
	websession.ClientErrorCode_notAcceptable:         http.StatusNotAcceptable,
	websession.ClientErrorCode_conflict:              http.StatusConflict,
	websession.ClientErrorCode_gone:                  http.StatusGone,
	websession.ClientErrorCode_preconditionFailed:    http.StatusPreconditionFailed,
	websession.ClientErrorCode_requestEntityTooLarge: http.StatusRequestEntityTooLarge,
	websession.ClientErrorCode_unsupportedMediaType:  http.StatusUnsupportedMediaType,
	websession.ClientErrorCode_imATeapot:             http.StatusTeapot,
	websession.ClientErrorCode_unprocessableEntity:   http.StatusUnprocessableEntity,
}
