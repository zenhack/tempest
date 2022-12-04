package sessioncookies

import (
	"crypto/rand"
	"net/http"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"github.com/gorilla/sessions"
	"zenhack.net/go/sandstorm-next/capnp/private/cookie"
	"zenhack.net/go/util"
)

type UserSession struct {
	sess *sessions.Session
	Data UserSessionData
}

type UserSessionData struct {
	SessionId  []byte
	Credential struct {
		Type     string
		ScopedId string
	}
}

func GetUserSession(s sessions.Store, req *http.Request) UserSession {
	sess, err := s.Get(req, "user-session")
	ret := UserSession{
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
		util.Chkfatal(pogs.Extract(&ret.Data, cookie.UserSession_TypeID, root.Struct()))
	} else {
		var buf [32]byte
		rand.Read(buf[:])
		ret.Data.SessionId = buf[:]
	}
	util.Chkfatal(err)
	return ret
}

func (s UserSession) Save(req *http.Request, w http.ResponseWriter) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	util.Chkfatal(err)
	root, err := cookie.NewRootUserSession(seg)
	util.Chkfatal(err)
	util.Chkfatal(pogs.Insert(cookie.UserSession_TypeID, capnp.Struct(root), &s.Data))
	s.sess.Values["data"] = seg.Data()
	util.Chkfatal(s.sess.Save(req, w))
}
