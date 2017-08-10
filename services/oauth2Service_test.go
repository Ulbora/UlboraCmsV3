package services

import "testing"

func TestAuthCodeAuthorize_AuthCodeAuthorizeUser(t *testing.T) {
	var a AuthCodeAuthorize
	a.ClientID = "211"
	a.OauthHost = "http://localhost:3000"
	a.RedirectURI = "http:/localhost/token"
	a.State = "12345"
	res := a.AuthCodeAuthorizeUser()
	if res != true {
		t.Fail()
	}

}

func TestAuthCodeToken(t *testing.T) {
	var tn AuthCodeToken
	tn.OauthHost = "http://localhost:3000"
	tn.ClientID = "211"
	tn.Secret = "2222222"
	tn.Code = "ldslkdslk"
	tn.RedirectURI = "http:/localhost/token"
	token := tn.AuthCodeToken()
	if token.AccessToken != "abcde" {
		t.Fail()
	}
}
