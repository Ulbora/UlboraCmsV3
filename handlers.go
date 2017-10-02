package main

import (
	services "UlboraCmsV3/services"
	"net/http"

	"github.com/gorilla/mux"
)

type pageContent struct {
	Title        string
	MetaAuthor   string
	MetaDesc     string
	MetaKeyWords string
	Cont         *[]services.Content
}

// user handlers-----------------------------------------------------
func handleIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["content"]
	if page == "" {
		page = "home"
	}
	var c services.ContentService
	c.ClientID = getAuthCodeClient()
	c.APIKey = getGatewayAPIKey()
	c.Host = getContentHost()
	h, res := c.GetContentListCategory(getAuthCodeClient(), page)
	var pg = new(pageContent)
	pg.Cont = res
	pg.MetaAuthor = h.MetaAuthor
	pg.MetaKeyWords = h.MetaKeyWords
	pg.MetaDesc = h.MetaDesc
	pg.Title = h.Title
	templates.ExecuteTemplate(w, "index.html", pg)
}

//end user handlers------------------------------------------------
