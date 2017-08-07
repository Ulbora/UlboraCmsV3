package services

import (
	"net/http"
	"testing"
)

var s Session

func TestSession_CreateSessionStore(t *testing.T) {
	s.MaxAge = 5 * 60
	s.Name = "user-session-test"
	var res http.ResponseWriter
	var req = new(http.Request)
	s.CreateSessionStore(res, req)
	if s.Store == nil {
		t.Fail()
	}
}
func TestGet(t *testing.T) {
	val := s.Get("test")
	if val != nil {
		t.Fail()
	}
}
