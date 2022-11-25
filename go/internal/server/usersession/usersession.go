package usersession

import (
	"crypto/rand"
	"io/ioutil"
	"net/http"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/pogs"
	"github.com/gorilla/sessions"
	"zenhack.net/go/sandstorm-next/capnp/internals/cookie"
	"zenhack.net/go/sandstorm-next/go/internal/config"
	"zenhack.net/go/util"
)

type Session struct {
	sess *sessions.Session
	Data SessionData
}

type SessionData struct {
	SessionId  []byte
	Credential struct {
		Type     string
		ScopedId string
	}
}

func GetKeys() ([][]byte, error) {
	var (
		data []byte
		err  error
	)
	const path = config.Localstatedir + "/sandstorm/session-key"
	data, err = ioutil.ReadFile(config.Localstatedir + "/sandstorm/session-key")
	if os.IsNotExist(err) {
		data = make([]byte, 64)
		rand.Read(data)
		err = os.WriteFile(path, data, 0600)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return [][]byte{data[:32], data[32:]}, nil
}

func Get(s sessions.Store, req *http.Request) Session {
	sess, err := s.Get(req, "user-session")
	ret := Session{
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

func (s Session) Save(req *http.Request, w http.ResponseWriter) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	util.Chkfatal(err)
	root, err := cookie.NewRootUserSession(seg)
	util.Chkfatal(err)
	util.Chkfatal(pogs.Insert(cookie.UserSession_TypeID, capnp.Struct(root), &s.Data))
	s.sess.Values["data"] = seg.Data()
	util.Chkfatal(s.sess.Save(req, w))
}
