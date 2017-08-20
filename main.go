package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	services "UlboraCmsV3/services"

	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

const (
	userSession       = "user-session"
	sessingTimeToLive = (5 * 60) //five minutes

	//http
	schemeDefault = "http://"

	//OAuth Auth Code
	authCodeClient = "403"
	authCodeSecret = "554444vfg55ggfff22454sw2fff2dsfd"
	authCodeState  = "ghh66555h"
)

var s usession.Session
var token *oauth2.Token

var templates = template.Must(template.ParseFiles("./static/templates/default/index.html"))
var templatesAdmin = template.Must(template.ParseFiles("./static/admin/index.html", "./static/admin/header.html",
	"./static/admin/footer.html", "./static/admin/navbar.html", "./static/admin/addContent.html"))

var username string

func main() {
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	}
	router := mux.NewRouter()

	//Web Site
	router.HandleFunc("/", handleIndex)

	//token link
	router.HandleFunc("/token", handleToken)

	//Admin site
	router.HandleFunc("/admin", handleAdminIndex)
	router.HandleFunc("/admin/", handleAdminIndex)
	router.HandleFunc("/addContent", handleAddContent)
	router.HandleFunc("/addContent/", handleAddContent)
	router.HandleFunc("/newContent", handleNewContent)
	router.HandleFunc("/logout", handleLogout)
	router.HandleFunc("/logout/", handleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("Ulbora CMS V3 running!")
	log.Println("Listening on :8090...")
	http.ListenAndServe(":8090", router)
}

// user handlers-----------------------------------------------------
func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.TransferEncoding)
	fmt.Println(req.Host)
	templates.ExecuteTemplate(res, "index.html", nil)
}

//end user handlers------------------------------------------------

// admin handlers -------------------------------------------------

func handleAdminIndex(res http.ResponseWriter, req *http.Request) {

	s.InitSessionStore(res, req)
	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false {
		authorize(res, req)
	} else {
		templatesAdmin.ExecuteTemplate(res, "index.html", nil)
	}

}

func handleAddContent(res http.ResponseWriter, req *http.Request) {

	s.InitSessionStore(res, req)
	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false {
		authorize(res, req)
	} else {
		templatesAdmin.ExecuteTemplate(res, "addContent.html", nil)
	}
}

func handleNewContent(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false {
		authorize(w, r)
	} else {
		content := r.FormValue("content")
		fmt.Print("content: ")
		fmt.Println(content)

		title := r.FormValue("title")
		fmt.Print("title: ")
		fmt.Println(title)

		author := r.FormValue("author")
		fmt.Print("author: ")
		fmt.Println(author)

		category := r.FormValue("category")
		fmt.Print("category: ")
		fmt.Println(category)

		metaKeyWords := r.FormValue("metaKeyWords")
		fmt.Print("metaKeyWords: ")
		fmt.Println(metaKeyWords)

		desc := r.FormValue("desc")
		fmt.Print("desc: ")
		fmt.Println(desc)
		var ct services.Content
		ct.Text = content
		ct.Title = title
		ct.MetaAuthorName = author
		ct.Category = category
		ct.MetaKeyWords = metaKeyWords
		ct.MetaRobotKeyWords = metaKeyWords
		ct.MetaDesc = desc

		var c services.ContentService
		c.ClientID = authCodeClient
		c.Token = token.AccessToken
		c.Host = getContentHost()
		res := c.AddContent(&ct)
		fmt.Println(res)

	}

}

// end admin handlers---------------------------------------------------------

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

func authorize(res http.ResponseWriter, req *http.Request) {
	fmt.Println("in authorize")
	fmt.Println(schemeDefault)
	var a oauth2.AuthCodeAuthorize
	a.ClientID = getAuthCodeClient()
	a.OauthHost = getOauthHost()
	a.RedirectURI = getRedirectURI(req, "/token")
	a.Scope = "write"
	a.State = authCodeState
	a.Res = res
	a.Req = req
	resp := a.AuthCodeAuthorizeUser()
	if resp != true {
		fmt.Println("Authorize Failed")
	}
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
		tn.RedirectURI = getRedirectURI(req, "/token")
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
				http.Redirect(res, req, "/admin", http.StatusFound)

				// decode token and get user id
			}
		}
	}
}

func getAuthCodeClient() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_ID") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_ID")
	} else {
		rtn = authCodeClient
	}
	return rtn
}

func getAuthCodeSecret() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_SECRET") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_SECRET")
	} else {
		rtn = authCodeSecret
	}
	return rtn
}
func getOauthHost() string {
	var rtn = ""
	if os.Getenv("AUTH_HOST") != "" {
		rtn = os.Getenv("AUTH_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}
func getRedirectURI(req *http.Request, path string) string {
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String()
	} else {
		serverHost = schemeDefault + req.Host + path
	}
	return serverHost
}

func getContentHost() string {
	var rtn = ""
	if os.Getenv("CONTENT_HOST") != "" {
		rtn = os.Getenv("CONTENT_HOST")
	} else {
		rtn = "http://localhost:3008"
	}
	return rtn
}
