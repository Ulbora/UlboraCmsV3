package services

import (
	"encoding/json"
	"fmt"
	"log"
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
	OauthHost    string
	RedirectURI  string
	ClientID     string
	Secret       string
	Code         string
	RefreshToken string
}

//Token the access token
type Token struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	TokenType     string `json:"token_type"`
	ExpiresIn     int    `json:"expires_in"`
	ErrorReturned string `json:"error"`
}

//AuthCodeToken auth code token
func (t *AuthCodeToken) AuthCodeToken() *Token {
	var url = t.OauthHost + oauthAuthCodeTokenURI1 + t.ClientID + oauthAuthCodeTokenURI2 + t.Secret +
		oauthAuthCodeTokenURI3 + t.Code + oauthAuthCodeTokenURI4 + t.RedirectURI
	fmt.Print("AuthCode Token URI: ")
	fmt.Println(url)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rtn := new(Token)
	if resp.StatusCode == 200 || resp.StatusCode == 401 {
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			log.Println(error.Error())
		}
	}
	return rtn
}

// AuthCodeRefreshToken get refresh token
func (t *AuthCodeToken) AuthCodeRefreshToken() *Token {
	var url = t.OauthHost + oauthAuthCodeRefreshTokenURI1 + t.ClientID + oauthAuthCodeRefreshTokenURI2 + t.Secret +
		oauthAuthCodeRefreshTokenURI3 + t.RefreshToken
	fmt.Print("AuthCode RefreshToken URI: ")
	fmt.Println(url)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	rtn := new(Token)
	if resp.StatusCode == 200 || resp.StatusCode == 401 {
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			log.Println(error.Error())
		}
	}
	return rtn
}
