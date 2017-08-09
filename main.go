package main

import (
	"UlboraCmsV3/services"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	userSession       = "user-session"
	sessingTimeToLive = (5 * 60)
)

var s services.Session

var templates = template.Must(template.ParseFiles("./static/templates/default/index.html"))
var templatesAdmin = template.Must(template.ParseFiles("./static/admin/index.html"))

var username string

func main() {
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	router := mux.NewRouter()

	//Web Site
	router.HandleFunc("/", handleIndex)

	//Admin site
	router.HandleFunc("/admin", handleAdminIndex)
	router.HandleFunc("/admin/", handleAdminIndex)
	fmt.Println("Ulbora CMS V3 running!")
	log.Println("Listening on :8090...")
	http.ListenAndServe(":8090", router)

}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "index.html", nil)
}

func handleAdminIndex(res http.ResponseWriter, req *http.Request) {
	s.InitSessionStore(res, req)
	fmt.Println("inside admin index")

	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
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
