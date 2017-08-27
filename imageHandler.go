package main

import (
	"fmt"
	"net/http"
)

func handleAddImage(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		templatesAdmin.ExecuteTemplate(w, "imageUpload.html", nil)
	}
}

func handleImagerUpdate(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {

		name := r.FormValue("name")
		fmt.Print("name: ")
		fmt.Println(name)
		err := r.ParseMultipartForm(200000)
		if err != nil {
			fmt.Println(err)
		}
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Print("name: ")
		fmt.Println(handler.Filename)
		fmt.Print("file len: ")
		fmt.Println(handler.Header)

		//if res.Success == true {
		http.Redirect(w, r, "/admin", http.StatusFound)
		//} else {
		//	fmt.Println("Mail Server update failed")
		//http.Redirect(w, r, "/admin", http.StatusFound)
		//}
	}
}
