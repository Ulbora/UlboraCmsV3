package main

import (
	services "UlboraCmsV3/services"
	"net/http"
)

func handleTemplates(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		var t services.TemplateService
		t.Host = getTemplateHost()
		res := t.GetTemplateList(appType, authCodeClient)
		templatesAdmin.ExecuteTemplate(w, "templates.html", &res)
	}
}
