package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
)

// admin content handlers -------------------------------------------------

func handleAdminIndex(w http.ResponseWriter, r *http.Request) {

	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		var c services.ContentService
		c.Host = getContentHost()
		res := c.GetContentList(authCodeClient)
		templatesAdmin.ExecuteTemplate(w, "index.html", &res)
	}

}

func handleAddContent(res http.ResponseWriter, req *http.Request) {

	s.InitSessionStore(res, req)
	session, err := s.GetSession(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
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
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
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
		c.UserID = getHashedUser()
		c.Hashed = "true"
		c.Token = token.AccessToken
		c.Host = getContentHost()
		res := c.AddContent(&ct)
		fmt.Println(res)
		if res.Success == true {
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			http.Redirect(w, r, "/addContent", http.StatusFound)
		}

	}

}

// end admin content handlers---------------------------------------------------------
