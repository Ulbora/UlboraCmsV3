package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func handleUpdateContent(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		idStr := r.FormValue("id")
		id, errID := strconv.ParseInt(idStr, 10, 0)
		if errID != nil {
			fmt.Print(errID)
		}
		fmt.Print("id: ")
		fmt.Println(id)

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
		ct.ID = id
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
		res := c.UpdateContent(&ct)
		fmt.Println(res)
		if res.Success == true {
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			fmt.Println("Content update failed")
			http.Redirect(w, r, "/admin", http.StatusFound)
		}
	}
}

func handleGetContent(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		var c services.ContentService
		c.Host = getContentHost()
		res := c.GetContent(id, authCodeClient)
		templatesAdmin.ExecuteTemplate(w, "updateContent.html", &res)
	}
}

func handleDeleteContent(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		var c services.ContentService
		c.ClientID = authCodeClient
		c.UserID = getHashedUser()
		c.Hashed = "true"
		c.Token = token.AccessToken
		c.Host = getContentHost()
		res := c.DeleteContent(id)
		if res.Success != true {
			fmt.Println("Delete content failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}

// end admin content handlers---------------------------------------------------------
