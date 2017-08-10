package services

import (
	"fmt"
	"net/http"
)

//AuthCodeAuthorize auth code
type AuthCodeAuthorize struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	State       string
}

//AuthCodeAuthorizeUser authorize a user with grant type code
func (a *AuthCodeAuthorize) AuthCodeAuthorizeUser() bool {
	var rtn = false
	var uri = a.OauthHost + oauthAuthCodeAuthorizeURI1 + a.ClientID + oauthAuthCodeAuthorizeURI2 +
		a.RedirectURI + oauthAuthCodeAuthorizeURI3 + a.State
	fmt.Print("AuthCode Authorize URI: ")
	fmt.Println(uri)
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	} else {
		rtn = true
	}
	defer resp.Body.Close()
	return rtn

}

//AuthCodeToken auth code token
type AuthCodeToken struct {
	OauthHost   string
	RedirectURI string
	ClientID    string
	Secret      string
	Code        string
}

//Token the access token
type Token struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int
}

//AuthCodeToken auth code token
func (t *AuthCodeToken) AuthCodeToken() *Token {
	var url = t.OauthHost + oauthAuthCodeTokenURI1 + t.ClientID + oauthAuthCodeTokenURI2 + t.Secret +
		oauthAuthCodeTokenURI3 + t.Code + oauthAuthCodeTokenURI4 + t.RedirectURI
	fmt.Print("AuthCode Token URI: ")
	fmt.Println(url)
	var token Token
	token.AccessToken = "abcde"
	//fmt.Errorf("sessions: invalid character in cookie name: %s", name)
	return &token
}
