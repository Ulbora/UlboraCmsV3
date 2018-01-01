package main

import (
	services "UlboraCmsV3/services"
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
var credentialToken *oauth2.Token

//var templateLoc = getTemplate()

var templates *template.Template

var templatesAdmin = template.Must(template.ParseFiles("./static/admin/index.html", "./static/admin/header.html",
	"./static/admin/footer.html", "./static/admin/navbar.html", "./static/admin/contentNavbar.html",
	"./static/admin/addContent.html", "./static/admin/images.html", "./static/admin/templates.html",
	"./static/admin/updateContent.html", "./static/admin/mailServer.html", "./static/admin/templateUpload.html",
	"./static/admin/imageUpload.html"))

var username string

func main() {
	setTemplate()
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	}
	router := mux.NewRouter()

	//Web Site
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/{content}", handleIndex)
	router.HandleFunc("/contact/form", handleContactForm)
	router.HandleFunc("/contact/send", handleContactSend)

	//token link
	router.HandleFunc("/admin/token", handleToken)

	//Admin site
	router.HandleFunc("/admin/main", handleAdminIndex)
	router.HandleFunc("/admin/main/", handleAdminIndex)
	router.HandleFunc("/admin/addContent", handleAddContent)
	router.HandleFunc("/admin/addContent/", handleAddContent)
	router.HandleFunc("/admin/newContent", handleNewContent)
	router.HandleFunc("/admin/getContent/{id}", handleGetContent)
	router.HandleFunc("/admin/updateContent", handleUpdateContent)
	router.HandleFunc("/admin/deleteContent/{id}/{cat}", handleDeleteContent)

	router.HandleFunc("/admin/mailServer", handleMailServer)
	router.HandleFunc("/admin/mailServerUpdate", handleMailServerUpdate)

	router.HandleFunc("/admin/addImage", handleAddImage)
	router.HandleFunc("/admin/uploadImage", handleImagerUpload)

	router.HandleFunc("/admin/images", handleImages)
	router.HandleFunc("/admin/deleteImage/{id}", handleDeleteImage)

	router.HandleFunc("/admin/templates", handleTemplates)
	router.HandleFunc("/admin/addTemplate", handleAddTemplate)
	router.HandleFunc("/admin/uploadTemplate", handleTemplateUpload)
	router.HandleFunc("/admin/templateActive/{id}", handleTemplateActive)
	router.HandleFunc("/admin/deleteTemplate/{id}/{name}", handleDeleteTemplate)

	router.HandleFunc("/admin/logout", handleLogout)
	router.HandleFunc("/admin/logout/", handleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("Ulbora CMS V3 running!")
	log.Println("Listening on :8090...")
	http.ListenAndServe(":8090", router)
}
func setTemplate() {
	var templateLoc = "default"
	var t services.TemplateService
	t.ClientID = getAuthCodeClient()
	t.APIKey = getGatewayAPIKey()
	t.Host = getTemplateHost()
	tmpl := t.GetTemplate("cms", getAuthCodeClient())
	if tmpl.Active {
		templateLoc = tmpl.Name
	}
	fmt.Println("using template " + templateLoc)
	templates = template.Must(template.ParseFiles("./static/templates/"+templateLoc+"/index.html", "./static/templates/"+templateLoc+"/header.html",
		"./static/templates/"+templateLoc+"/footer.html", "./static/templates/"+templateLoc+"/navbar.html",
		"./static/templates/"+templateLoc+"/contact.html"))
}
