package websession

import (
	"net/http"

	"zenhack.net/go/sandstorm/capnp/websession"
)

// Mapping from web-session success codes to http status code numbers
var successCodeStatuses = map[websession.WebSession_Response_SuccessCode]int{
	websession.WebSession_Response_SuccessCode_ok:          http.StatusOK,
	websession.WebSession_Response_SuccessCode_created:     http.StatusCreated,
	websession.WebSession_Response_SuccessCode_accepted:    http.StatusAccepted,
	websession.WebSession_Response_SuccessCode_noContent:   http.StatusNoContent,
	websession.WebSession_Response_SuccessCode_multiStatus: http.StatusMultiStatus,
	websession.WebSession_Response_SuccessCode_notModified: http.StatusNotModified,
}

// Mapping from web-session error codes to http status code numbers
var clientErrorCodeStatuses = map[websession.WebSession_Response_ClientErrorCode]int{
	websession.WebSession_Response_ClientErrorCode_badRequest:            http.StatusBadRequest,
	websession.WebSession_Response_ClientErrorCode_forbidden:             http.StatusForbidden,
	websession.WebSession_Response_ClientErrorCode_notFound:              http.StatusNotFound,
	websession.WebSession_Response_ClientErrorCode_methodNotAllowed:      http.StatusMethodNotAllowed,
	websession.WebSession_Response_ClientErrorCode_notAcceptable:         http.StatusNotAcceptable,
	websession.WebSession_Response_ClientErrorCode_conflict:              http.StatusConflict,
	websession.WebSession_Response_ClientErrorCode_gone:                  http.StatusGone,
	websession.WebSession_Response_ClientErrorCode_preconditionFailed:    http.StatusPreconditionFailed,
	websession.WebSession_Response_ClientErrorCode_requestEntityTooLarge: http.StatusRequestEntityTooLarge,
	websession.WebSession_Response_ClientErrorCode_unsupportedMediaType:  http.StatusUnsupportedMediaType,
	websession.WebSession_Response_ClientErrorCode_imATeapot:             http.StatusTeapot,
	websession.WebSession_Response_ClientErrorCode_unprocessableEntity:   http.StatusUnprocessableEntity,
}
