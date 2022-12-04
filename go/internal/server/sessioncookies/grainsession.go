package sessioncookies

import (
	"net/http"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"github.com/gorilla/sessions"
	"zenhack.net/go/sandstorm-next/capnp/private/cookie"
	"zenhack.net/go/util"
)

type GrainSession struct {
	sess *sessions.Session
	Data GrainSessionData
}

type GrainSessionData struct {
}

func GetGrainSession(s sessions.Store, req *http.Request) (_ GrainSession, ok bool) {
	sess, err := s.Get(req, "grain-session")
	ret := GrainSession{
		sess: sess,
	}
	valAny, ok := sess.Values["data"]
	var (
		msg *capnp.Message
		val []byte
	)
	if ok {
		val, ok = valAny.([]byte)
	}
	if ok {
		msg = &capnp.Message{
			Arena: capnp.SingleSegment(val),
		}
		root, err := msg.Root()
		util.Chkfatal(err)
		util.Chkfatal(pogs.Extract(&ret.Data, cookie.GrainSession_TypeID, root.Struct()))
	} else {
		return GrainSession{}, false
	}
	util.Chkfatal(err)
	return ret, ok
}
