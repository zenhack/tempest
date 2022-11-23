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
	return Session{
		sess: sess,
	}
}

func (s Session) Id() string {
	sessionId, ok := s.sess.Values["session-id"]
	if !ok {
		return ""
	}
	ret, ok := sessionId.(string)
	if !ok {
		return ""
	}
	return ret
}

func (s Session) SetDevCredential(name string) {
	s.sess.Values["dev-credential-name"] = name
}

func (s Session) Save(req *http.Request, w http.ResponseWriter) {
	s.sess.Save(req, w)
}
