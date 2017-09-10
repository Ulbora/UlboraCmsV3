package main

import (
	services "UlboraCmsV3/services"
	"net/http"

	"github.com/gorilla/mux"
)

// user handlers-----------------------------------------------------
func handleContactSend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["content"]
	if page == "" {
		page = "home"
	}
	var c services.ContentService
	c.Host = getContentHost()
	h, res := c.GetContentListCategory(authCodeClient, page)
	var pg = new(pageContent)
	pg.Cont = res
	pg.MetaAuthor = h.MetaAuthor
	pg.MetaKeyWords = h.MetaKeyWords
	pg.MetaDesc = h.MetaDesc
	pg.Title = h.Title
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handleContactForm(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}

//end user handlers------------------------------------------------
