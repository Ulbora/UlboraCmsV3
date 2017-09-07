package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func handleAddTemplate(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		templatesAdmin.ExecuteTemplate(w, "templateUpload.html", nil)
	}
}

func handleTemplateUpload(w http.ResponseWriter, r *http.Request) {
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
		fmt.Print("template name: ")
		fmt.Println(name)
		err := r.ParseMultipartForm(1000000)
		if err != nil {
			fmt.Println(err)
		}

		file, handler, err := r.FormFile("template")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		fmt.Print("File name: ")
		fmt.Println(handler.Filename)

		// cur, err := file.Seek(0, 1)
		// size, err := file.Seek(0, 2)
		// _, err1 := file.Seek(cur, 0)
		// if err1 != nil {
		// 	fmt.Println(err1)
		// }

		// data, err := ioutil.ReadAll(file)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Println(data)

		// fmt.Print("file size: ")
		// fmt.Println(size)

		//sEnc := b64.StdEncoding.EncodeToString(data)
		//fmt.Print("file data: ")
		//fmt.Println(data)
		var t services.TemplateService
		t.ClientID = authCodeClient
		t.Host = getTemplateHost()
		t.Token = token.AccessToken
		var tmpl services.Template
		tmpl.Name = name
		tmpl.Active = false
		tmpl.Application = "cms"
		var res *services.TemplateResponse
		res = t.AddTemplate(&tmpl)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = t.AddTemplate(&tmpl)
		}
		//res := i.AddImage(&img)
		if res.Success == true {
			http.Redirect(w, r, "/admin/templates", http.StatusFound)
		} else {
			fmt.Println("Template upload failed")
			//http.Redirect(w, r, "/admin/addTemplate", http.StatusFound)

			// add error handling here and pass error to page-----------------------------------
			templatesAdmin.ExecuteTemplate(w, "templateUpload.html", nil)
		}
	}
}

func handleTemplateActive(w http.ResponseWriter, r *http.Request) {
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
		idStr := vars["id"]
		id, errID := strconv.ParseInt(idStr, 10, 0)
		if errID != nil {
			fmt.Print(errID)
		}
		var t services.TemplateService
		t.ClientID = authCodeClient
		t.Host = getTemplateHost()
		t.Token = token.AccessToken
		var tm services.Template
		tm.ID = id
		tm.Application = "cms"
		var res *services.TemplateResponse
		res = t.UpdateTemplate(&tm)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = t.UpdateTemplate(&tm)
		}
		//res := i.DeleteImage(id)
		if res.Success != true {
			fmt.Println("Made template id " + idStr + " active")
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/admin/templates", http.StatusFound)
	}
}

func handleDeleteTemplate(w http.ResponseWriter, r *http.Request) {
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
		var t services.TemplateService
		t.ClientID = authCodeClient
		t.Host = getTemplateHost()
		t.Token = token.AccessToken
		var res *services.TemplateResponse
		res = t.DeleteTemplate(id)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = t.DeleteTemplate(id)
		}
		//res := i.DeleteImage(id)
		if res.Success != true {
			fmt.Println("Delete Template failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/admin/templates", http.StatusFound)
	}
}
