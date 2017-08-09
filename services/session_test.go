package services

import (
	"fmt"
	"net/http"
	"testing"
)

var s Session

func TestSession_CreateSessionStore(t *testing.T) {
	s.MaxAge = 5 * 60
	s.Name = "user-session-test"
	var res http.ResponseWriter
	var req = new(http.Request)
	s.InitSessionStore(res, req)
	if s.store == nil {
		t.Fail()
	}
}
func TestSession(t *testing.T) {
	var req = new(http.Request)
	var res http.ResponseWriter
	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	session.Values["test"] = "test_2"
	test1 := session.Values["test"]
	fmt.Println(test1)
	if test1 != "test_2" {
		t.Fail()
	}
}
