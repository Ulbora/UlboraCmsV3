package services

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var sessionKey string

//Session a sesstion
type Session struct {
	Store   *sessions.CookieStore
	MaxAge  int
	Name    string
	session *sessions.Session
}

// CreateSessionStore creates a sesstion
func (s *Session) CreateSessionStore(res http.ResponseWriter, req *http.Request) {
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		sessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		sessionKey = "554dfgdffdd11dfgf1ff1f"
	}
	s.Store = sessions.NewCookieStore([]byte(sessionKey))
	s.Store.Options = &sessions.Options{
		MaxAge:   s.MaxAge,
		HttpOnly: true,
	}
	var err error
	s.session, err = s.Store.Get(req, s.Name)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//Get gets session variables
func (s *Session) Get(name string) interface{} {
	return s.session.Values[name]
}
