package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type uploadError struct {
	UploadFailed bool
}

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
		t.ClientID = getAuthCodeClient()
		t.APIKey = getGatewayAPIKey()
		res := t.GetTemplateList(appType, getAuthCodeClient())
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
		//fmt.Print("template name: ")
		//fmt.Println(name)
		err := r.ParseMultipartForm(1000000)
		if err != nil {
			fmt.Println(err)
		}

		file, handler, err := r.FormFile("template")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		//fmt.Print("File name: ")
		//fmt.Println(handler.Filename)

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		var t services.TemplateService
		t.ClientID = getAuthCodeClient()
		t.APIKey = getGatewayAPIKey()
		t.Host = getTemplateHost()
		t.Token = token.AccessToken
		dtemp := t.GetTemplate("cms", getAuthCodeClient())
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
		var eres = false
		if res.Success == true {
			// untar file
			var ts services.TemplateFileService
			ts.OriginalFileName = handler.Filename
			ts.Destination = "./static/templates"
			ts.Name = name
			ts.FileData = data
			eres = ts.ExtractFile()
		}
		if res.Success == true && eres == true {
			if dtemp.Active == false {
				var tmpl services.Template
				tmpl.Name = "default"
				tmpl.Active = true
				tmpl.Application = "cms"
				t.AddTemplate(&tmpl)
			}
			http.Redirect(w, r, "/admin/templates", http.StatusFound)
		} else {
			fmt.Println("Template upload failed")
			//http.Redirect(w, r, "/admin/addTemplate", http.StatusFound)
			var uErr uploadError
			uErr.UploadFailed = true
			// add error handling here and pass error to page-----------------------------------
			templatesAdmin.ExecuteTemplate(w, "templateUpload.html", &uErr)
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
		t.ClientID = getAuthCodeClient()
		t.APIKey = getGatewayAPIKey()
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
		if res.Success == true {
			fmt.Println("Made template id " + idStr + " active")
			fmt.Print("code: ")
			fmt.Println(res.Code)
			setTemplate()
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
		name := vars["name"]
		var t services.TemplateService
		t.ClientID = getAuthCodeClient()
		t.APIKey = getGatewayAPIKey()
		t.Host = getTemplateHost()
		t.Token = token.AccessToken
		var res *services.TemplateResponse
		res = t.DeleteTemplate(id)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = t.DeleteTemplate(id)
		}
		var dres = false
		if res.Success == true {
			var ts services.TemplateFileService
			ts.Destination = "./static/templates"
			ts.Name = name
			dres = ts.DeleteTemplate()
		}
		if res.Success != true || dres != true {
			fmt.Println("Delete Template failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/admin/templates", http.StatusFound)
	}
}
