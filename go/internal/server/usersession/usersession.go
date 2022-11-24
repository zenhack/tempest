package usersession

import (
	"crypto/rand"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"zenhack.net/go/sandstorm-next/go/internal/config"
)

type Session struct {
	sess *sessions.Session
	id   string
	// TODO: expiration date?
}

func GetKey() ([]byte, error) {
	const path = config.Localstatedir + "/sandstorm/session-key"
	data, err := ioutil.ReadFile(config.Localstatedir + "/sandstorm/session-key")
	if os.IsNotExist(err) {
		data := make([]byte, 32)
		rand.Read(data)
		return data, os.WriteFile(path, data, 0600)
	} else {
		return nil, err
	}
	return data, nil
}

func Get(s sessions.Store, req *http.Request) Session {
	sess, _ := s.Get(req, "user-session")
	ret := Session{
		sess: sess,
	}
	sessionId, ok := sess.Values["session-id"]
	if ok {
		ret.id, ok = sessionId.(string)
	}
	if !ok {
		var buf [32]byte
		rand.Read(buf[:])
		ret.id = string(buf[:])
	}
	return ret
}

func (s Session) Id() string {
	return s.id
}

func (s Session) Credential() (typ, scopedId string) {
	typ = s.sess.Values["credential-type"].(string)
	scopedId = s.sess.Values["credential-scoped-id"].(string)
	return
}

func (s Session) SetCredential(typ, scopedId string) {
	s.sess.Values["credential-type"] = typ
	s.sess.Values["credential-scoped-id"] = scopedId
}

func (s Session) Save(req *http.Request, w http.ResponseWriter) {
	s.sess.Save(req, w)
}
