package main

import (
	services "UlboraCmsV3/services"
	"net/http"
)

// user handlers-----------------------------------------------------
func handleIndex(w http.ResponseWriter, r *http.Request) {
	var c services.ContentService
	c.Host = getContentHost()
	res := c.GetContentListCategory(authCodeClient, "home")

	templates.ExecuteTemplate(w, "index.html", &res)
}

//end user handlers------------------------------------------------
