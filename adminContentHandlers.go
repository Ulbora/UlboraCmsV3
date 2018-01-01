package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type contentAndImages struct {
	Cont *services.Content
	Img  *[]services.Image
}

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
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()
		c.Host = getContentHost()
		res := c.GetContentList(getAuthCodeClient())
		templatesAdmin.ExecuteTemplate(w, "index.html", &res)
	}
}

func handleAddContent(w http.ResponseWriter, r *http.Request) {

	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
		fmt.Print("auth success: ")
		fmt.Println(s)
	} else {
		var i services.ImageService
		i.ClientID = getAuthCodeClient()
		i.APIKey = getGatewayAPIKey()
		//i.UserID = getHashedUser()
		//i.Hashed = "true"
		i.Token = token.AccessToken
		//fmt.Println(token.AccessToken)
		i.Host = getImageHost()

		res := i.GetList()

		templatesAdmin.ExecuteTemplate(w, "addContent.html", &res)
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
		//fmt.Print("content: ")
		//fmt.Println(content)

		title := r.FormValue("title")
		//fmt.Print("title: ")
		//fmt.Println(title)

		author := r.FormValue("author")
		//fmt.Print("author: ")
		//fmt.Println(author)

		category := r.FormValue("category")
		category = strings.Replace(category, " ", "", -1)
		//fmt.Print("category: ")
		//fmt.Println(category)

		sortOrder := r.FormValue("sortOrder")
		if sortOrder == "" {
			sortOrder = "0"
		}
		//fmt.Print("sortOrder: ")
		//fmt.Println(sortOrder)

		metaKeyWords := r.FormValue("metaKeyWords")
		//fmt.Print("metaKeyWords: ")
		//fmt.Println(metaKeyWords)

		desc := r.FormValue("desc")
		//fmt.Print("desc: ")
		//fmt.Println(desc)
		var ct services.Content
		ct.Text = content
		ct.Title = title
		ct.MetaAuthorName = author
		ct.Category = category
		ct.MetaKeyWords = metaKeyWords
		ct.MetaRobotKeyWords = metaKeyWords
		ct.MetaDesc = desc
		ct.SortOrder, err = strconv.Atoi(sortOrder)
		if err != nil {
			fmt.Println(err)
		}
		var c services.ContentService
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()
		c.UserID = getHashedUser()
		c.Hashed = "true"
		c.Token = token.AccessToken
		c.Host = getContentHost()
		var res *services.Response
		res = c.AddContent(&ct)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = c.AddContent(&ct)
		}

		fmt.Println(res)
		if res.Success == true {
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		} else {
			http.Redirect(w, r, "/admin/addContent", http.StatusFound)
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
		//fmt.Print("id: ")
		//fmt.Println(id)

		content := r.FormValue("content")
		//fmt.Print("content: ")
		//fmt.Println(content)

		title := r.FormValue("title")
		//fmt.Print("title: ")
		//fmt.Println(title)

		author := r.FormValue("author")
		//fmt.Print("author: ")
		//fmt.Println(author)

		category := r.FormValue("category")
		category = strings.Replace(category, " ", "", -1)
		//fmt.Print("category: ")
		//fmt.Println(category)

		sortOrder := r.FormValue("sortOrder")
		if sortOrder == "" {
			sortOrder = "0"
		}
		//fmt.Print("sortOrder: ")
		//fmt.Println(sortOrder)

		metaKeyWords := r.FormValue("metaKeyWords")
		//fmt.Print("metaKeyWords: ")
		//fmt.Println(metaKeyWords)

		desc := r.FormValue("desc")
		//fmt.Print("desc: ")
		//fmt.Println(desc)

		archived := r.FormValue("archived")
		//fmt.Print("archived: ")
		//fmt.Println(archived)

		var ct services.Content
		ct.ID = id
		ct.Text = content
		ct.Title = title
		ct.MetaAuthorName = author
		ct.Category = category
		ct.MetaKeyWords = metaKeyWords
		ct.MetaRobotKeyWords = metaKeyWords
		ct.MetaDesc = desc
		ct.SortOrder, err = strconv.Atoi(sortOrder)
		if err != nil {
			fmt.Print("sortOrder conversion error: ")
			fmt.Println(err)
		}
		if archived == "on" {
			ct.Archived = true
		} else {
			ct.Archived = false
		}

		var c services.ContentService
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()
		c.UserID = getHashedUser()
		c.Hashed = "true"
		c.Token = token.AccessToken
		c.Host = getContentHost()

		var res *services.Response

		res = c.UpdateContent(&ct)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = c.UpdateContent(&ct)
		}
		//fmt.Println(res)
		if res.Success == true {
			var c services.ContentPageService
			c.ClientID = getAuthCodeClient()
			c.APIKey = getGatewayAPIKey()
			c.Token = token.AccessToken
			c.Host = getContentHost()
			c.PageSize = 100
			c.ClearPage(category)
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		} else {
			fmt.Println("Content update failed")
			http.Redirect(w, r, "/admin/main", http.StatusFound)
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
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()
		c.Host = getContentHost()
		res := c.GetContent(id, getAuthCodeClient())

		var i services.ImageService
		i.ClientID = getAuthCodeClient()
		i.APIKey = getGatewayAPIKey()
		//i.UserID = getHashedUser()
		//i.Hashed = "true"
		i.Token = token.AccessToken
		//fmt.Println(token.AccessToken)
		i.Host = getImageHost()

		ires := i.GetList()

		var ci = new(contentAndImages)
		ci.Cont = res
		ci.Img = ires

		templatesAdmin.ExecuteTemplate(w, "updateContent.html", &ci)
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
		page := vars["cat"]
		//fmt.Print("page: ")
		//fmt.Println(page)
		var c services.ContentService
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()
		c.UserID = getHashedUser()
		c.Hashed = "true"
		c.Token = token.AccessToken
		c.Host = getContentHost()
		//res := c.DeleteContent(id)
		var res *services.Response
		res = c.DeleteContent(id)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = c.DeleteContent(id)
		}
		if res.Success != true {
			fmt.Println("Delete content failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		} else {
			// add code to delete cached page====================================
			var c services.ContentPageService
			c.ClientID = getAuthCodeClient()
			c.APIKey = getGatewayAPIKey()
			c.Token = token.AccessToken
			c.Host = getContentHost()
			c.PageSize = 100
			//res2 := c.DeletePage(page)
			c.DeletePage(page)
			//fmt.Print("delete res: ")
			//fmt.Println(res2)
		}
		http.Redirect(w, r, "/admin/main", http.StatusFound)
	}
}

// end admin content handlers---------------------------------------------------------
