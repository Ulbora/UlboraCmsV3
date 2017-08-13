package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

const (
	userSession       = "user-session"
	sessingTimeToLive = (5 * 60)

	//http
	schemeDefault = "http://"

	//OAuth Auth Code
	authCodeClient = "403"
	authCodeSecret = "554444vfg55ggfff22454sw2fff2dsfd"
	authCodeState  = "ghh66555h"
)

var s usession.Session

//var oauth2Auth oauth2.AuthCodeAuthorize

var templates = template.Must(template.ParseFiles("./static/templates/default/index.html"))
var templatesAdmin = template.Must(template.ParseFiles("./static/admin/index.html"))

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
	router.HandleFunc("/admin2", handleAdminIndex)
	router.HandleFunc("/admin2/", handleAdminIndex)
	fmt.Println("Ulbora CMS V3 running!")
	log.Println("Listening on :8090...")
	http.ListenAndServe(":8090", router)

}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.TransferEncoding)
	fmt.Println(req.Host)
	templates.ExecuteTemplate(res, "index.html", nil)
}

func handleAdminIndex(res http.ResponseWriter, req *http.Request) {

	s.InitSessionStore(res, req)
	fmt.Println("inside admin index")

	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) != false {
		authorize(res, req)
	} else {
		user := session.Values["username"]
		fmt.Print("user session: ")
		fmt.Println(user)
		var username = ""
		if user != nil {
			username = user.(string)
		}

		if username == "" {
			fmt.Println("creating new user id")
			tempUs := (rand.Float64() * 5) + 5
			fmt.Print("new random: ")
			fmt.Println(tempUs)
			session.Values["username"] = strconv.FormatFloat(tempUs, 'f', 15, 64)
			session.Save(req, res)
			user := session.Values["username"]
			username = user.(string)
		}

		fmt.Print("user: ")
		fmt.Println(username)

		templatesAdmin.ExecuteTemplate(res, "index.html", nil)
	}

}

func authorize(res http.ResponseWriter, req *http.Request) {
	fmt.Println("in authorize")
	fmt.Println(schemeDefault)
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String()
	} else {
		serverHost = schemeDefault + req.Host
	}
	var a oauth2.AuthCodeAuthorize
	a.ClientID = getAuthCodeClient()
	a.OauthHost = getOauthHost()
	a.RedirectURI = serverHost + "/token"
	fmt.Println(a.ClientID)
	fmt.Println(a.RedirectURI)
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
	fmt.Println(req.TransferEncoding)
	fmt.Println(req.Host)
	//templates.ExecuteTemplate(res, "index.html", nil)
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
