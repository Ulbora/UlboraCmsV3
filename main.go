package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	//services "UlboraCmsV3/services"

	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

var s usession.Session
var token *oauth2.Token

var templates = template.Must(template.ParseFiles("./static/templates/default/index.html"))
var templatesAdmin = template.Must(template.ParseFiles("./static/admin/index.html", "./static/admin/header.html",
	"./static/admin/footer.html", "./static/admin/navbar.html", "./static/admin/contentNavbar.html",
	"./static/admin/addContent.html",
	"./static/admin/updateContent.html", "./static/admin/mailServer.html",
	"./static/admin/imageUpload.html"))

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
	router.HandleFunc("/getContent/{id}", handleGetContent)
	router.HandleFunc("/updateContent", handleUpdateContent)
	router.HandleFunc("/deleteContent/{id}", handleDeleteContent)

	router.HandleFunc("/mailServer", handleMailServer)
	router.HandleFunc("/mailServerUpdate", handleMailServerUpdate)

	router.HandleFunc("/addImage", handleAddImage)
	router.HandleFunc("/uploadImage", handleImagerUpdate)

	router.HandleFunc("/logout", handleLogout)
	router.HandleFunc("/logout/", handleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("Ulbora CMS V3 running!")
	log.Println("Listening on :8090...")
	http.ListenAndServe(":8090", router)
}
