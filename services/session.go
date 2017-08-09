package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var sessionKey string

//Session a sesstion
type Session struct {
	store  *sessions.CookieStore
	MaxAge int
	Name   string
}

//InitSessionStore initialize session store
func (s *Session) InitSessionStore(res http.ResponseWriter, req *http.Request) {
	if s.store == nil {
		s.createSessionStore(res, req)
	}
}

// CreateSessionStore creates a sesstion
func (s *Session) createSessionStore(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating Session Store")
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		sessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		sessionKey = "554dfgdffdd11dfgf1ff1f"
	}
	s.store = sessions.NewCookieStore([]byte(sessionKey))
	s.store.Options = &sessions.Options{
		MaxAge:   s.MaxAge,
		HttpOnly: true,
	}
}

//GetSession get session
func (s *Session) GetSession(req *http.Request) (*sessions.Session, error) {
	session, err := s.store.Get(req, s.Name)
	return session, err
}
