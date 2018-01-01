package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
		//fmt.Print("name: ")
		//fmt.Println(name)
		err := r.ParseMultipartForm(2000000)
		if err != nil {
			fmt.Println(err)
		}
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		//fmt.Print("name: ")
		//fmt.Println(handler.Filename)

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

		//fmt.Print("file size: ")
		//fmt.Println(size)

		//sEnc := b64.StdEncoding.EncodeToString(data)
		//fmt.Print("file data: ")
		//fmt.Println(data)
		var i services.ImageService
		i.ClientID = getAuthCodeClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = token.AccessToken
		var img services.UploadedFile
		img.Name = name
		img.OriginalFileName = handler.Filename
		img.Size = size
		img.FileData = data
		var res *services.ImageResponse
		res = i.AddImage(&img)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = i.AddImage(&img)
		}
		//res := i.AddImage(&img)
		if res.Success == true {
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		} else {
			fmt.Println("Image upload failed")
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		}
	}
}

func handleImages(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		var i services.ImageService
		i.ClientID = getAuthCodeClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = token.AccessToken
		//i.UserID = getHashedUser()
		//i.Hashed = "true"

		//fmt.Println(token.AccessToken)
		//i.Host = getImageHost()

		res := i.GetList()

		templatesAdmin.ExecuteTemplate(w, "images.html", &res)
	}
}

func handleDeleteImage(w http.ResponseWriter, r *http.Request) {
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
		var i services.ImageService
		i.ClientID = getAuthCodeClient()
		i.APIKey = getGatewayAPIKey()
		i.Host = getImageHost()
		i.Token = token.AccessToken
		var res *services.ImageResponse
		res = i.DeleteImage(id)
		if res.Code == 401 {
			// get new token
			getRefreshToken(w, r)
			res = i.DeleteImage(id)
		}
		//res := i.DeleteImage(id)
		if res.Success != true {
			fmt.Println("Delete image failed on ID: " + id)
			fmt.Print("code: ")
			fmt.Println(res.Code)
		}
		http.Redirect(w, r, "/admin/images", http.StatusFound)
	}
}
