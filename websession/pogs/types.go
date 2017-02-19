// Package websession_pogs defines pogs-style types for the websesion schema.
// This is WIP.
package websession_pogs

// TODO: default values?

import (
	"zenhack.net/go/sandstorm/capnp/websession"
)

// TODO: this is from util, we should move it there.
type KeyValue struct {
	Key   string
	Value string
}

type AcceptedEncoding struct {
	ContentCoding string
	QValue        float32
}

type AcceptedType struct {
	MimeType string
	QValue   float32
}

type CachePolicy struct {
	Scope websession.WebSession_CachePolicy
}

type Context struct {
	Cookies          []KeyValue
	ResponseStream   struct{} // TODO: interface?
	Accept           []AcceptedType
	AcceptEncoding   []AcceptedEncoding
	ETagPrecondition struct {
		Which         websession.WebSession_Context_eTagPrecondition_Which
		MatchesOneOf  []ETag
		MatchesNoneOf []ETag
	}
	AdditionalHeaders []Header
}

type Cookie struct {
	Name    string
	Value   string
	Expires struct {
		Which    websession.WebSession_Cookie_expires_Which
		Absolute int64
		Relative uint64
	}
	HttpOnly bool
	Path     string
}

type ETag struct {
	Value string
	Weak  bool
}

type Header struct {
	Name  string
	Value string
}

type Response struct {
	SetCookies  []Cookie
	CachePolicy CachePolicy
	Which       websession.WebSession_Response_Which
	Content     struct {
		StatusCode websession.WebSession_Response_SuccessCode
		Encoding   string
		Language   string
		MimeType   string
		ETag       ETag
		Body       struct {
			Which  websession.WebSession_Response_content_body_Which
			Bytes  []byte
			Stream struct{} // TODO: how do we put interfaces here?
		}
		Disposition struct {
			Which    websession.WebSession_Response_content_disposition_Which
			Download string
		}
	}
	NoContent struct {
		ShouldResetForm bool
		ETag            ETag
	}
	PreconditionFailed struct {
		MatchingETag ETag
	}
	Redirect struct {
		IsPermanent bool
		SwitchToGet bool
		Location    string
	}
	ClientError struct {
		StatusCode      websession.WebSession_Response_ClientErrorCode
		DescriptionHtml string
	}
	ServerError struct {
		DescriptionHtml string
	}
	AdditionalHeaders []Header
}

type Get_args struct {
	Path       string
	Context    Context
	IgnoreBody bool
}
