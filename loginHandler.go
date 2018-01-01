package main

import (
	"fmt"
	"net/http"

	oauth2 "github.com/Ulbora/go-oauth2-client"
)

// login handler
func handleLogout(res http.ResponseWriter, req *http.Request) {
	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	} else {
		session.Values["userLoggenIn"] = false
		session.Save(req, res)
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func authorize(res http.ResponseWriter, req *http.Request) bool {
	fmt.Println("in authorize")
	fmt.Println(schemeDefault)
	var a oauth2.AuthCodeAuthorize
	a.ClientID = getAuthCodeClient()
	a.OauthHost = getOauthHost()
	a.RedirectURI = getRedirectURI(req, "/admin/token")
	a.Scope = "write"
	a.State = authCodeState
	a.Res = res
	a.Req = req
	resp := a.AuthCodeAuthorizeUser()
	if resp != true {
		fmt.Println("Authorize Failed")
	}
	//fmt.Print("Resp: ")
	//fmt.Println(resp)
	return resp
}

func handleToken(res http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	state := req.URL.Query().Get("state")
	if state == authCodeState {
		var tn oauth2.AuthCodeToken
		tn.OauthHost = getOauthHost()
		tn.ClientID = getAuthCodeClient()
		tn.Secret = getAuthCodeSecret()
		tn.Code = code
		tn.RedirectURI = getRedirectURI(req, "/admin/token")
		resp := tn.AuthCodeToken()
		if resp != nil && resp.AccessToken != "" {
			//fmt.Println(resp.AccessToken)
			token = resp
			session, err := s.GetSession(req)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			} else {
				session.Values["userLoggenIn"] = true
				session.Save(req, res)
				http.Redirect(res, req, "/admin/main", http.StatusFound)

				// decode token and get user id
			}
		}
	}
}
