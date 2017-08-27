package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"io/ioutil"
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

func handleImagerUpload(w http.ResponseWriter, r *http.Request) {
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

		cur, err := file.Seek(0, 1)
		size, err := file.Seek(0, 2)
		_, err1 := file.Seek(cur, 0)
		if err1 != nil {
			fmt.Println(err1)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(data)

		fmt.Print("file size: ")
		fmt.Println(size)

		//sEnc := b64.StdEncoding.EncodeToString(data)
		//fmt.Print("file data: ")
		//fmt.Println(data)
		var i services.ImageService
		i.ClientID = authCodeClient
		i.Host = getImageHost()
		i.Token = token.AccessToken
		var img services.UploadedFile
		img.Name = name
		img.OriginalFileName = handler.Filename
		img.Size = size
		img.FileData = data
		res := i.AddImage(&img)
		if res.Success == true {
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			fmt.Println("Mail Server update failed")
			http.Redirect(w, r, "/admin", http.StatusFound)
		}
	}
}
