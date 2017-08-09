package services

import "fmt"

//AuthCodeAuthorize auth code
type AuthCodeAuthorize struct {
	OauthHost                      string
	OauthAuthCodeAuthorizeURIPart1 string
	OauthAuthCodeAuthorizeURIPart2 string
	OauthAuthCodeAuthorizeURIPart3 string
	RedirectURI                    string
	ClientID                       string
	State                          string
}

//AuthCodeAuthorizeUser authorize a user with grant type code
func (a *AuthCodeAuthorize) AuthCodeAuthorizeUser() {
	var uri = a.OauthHost + a.OauthAuthCodeAuthorizeURIPart1 + a.ClientID + a.OauthAuthCodeAuthorizeURIPart2 +
		a.RedirectURI + a.OauthAuthCodeAuthorizeURIPart3 + a.State
	fmt.Print("AuthCode Authorize URI: ")
	fmt.Println(uri)

}
