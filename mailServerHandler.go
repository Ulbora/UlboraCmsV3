package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
	"strconv"
)

func handleMailServer(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		var m services.MailServerService
		m.ClientID = getAuthCodeClient()
		//m.UserID = getHashedUser()
		//m.Hashed = "true"
		m.Token = token.AccessToken
		m.Host = getMailHost()
		m.APIKey = getGatewayAPIKey()
		res := m.GetMailServer()
		if res.Server.ID == 0 {
			res.Server.Port = 465
		}
		//fmt.Print("mail res: ")
		//fmt.Println(res)
		//fmt.Print("mail server: ")
		//fmt.Println(m)
		templatesAdmin.ExecuteTemplate(w, "mailServer.html", &res.Server)
	}
}

func handleMailServerUpdate(w http.ResponseWriter, r *http.Request) {
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

		mailServer := r.FormValue("mailServer")
		//fmt.Print("mailServer: ")
		//fmt.Println(mailServer)

		secureConnection := r.FormValue("secureConnection")
		//fmt.Print("secureConnection: ")
		//fmt.Println(secureConnection)

		port := r.FormValue("port")
		if port == "" {
			port = "0"
		}
		//fmt.Print("port: ")
		//fmt.Println(port)

		debug := r.FormValue("debug")
		//fmt.Print("debug: ")
		//fmt.Println(debug)

		username := r.FormValue("username")
		//fmt.Print("username: ")
		//fmt.Println(username)

		password := r.FormValue("password")
		//fmt.Print("password: ")
		//fmt.Println(password)

		fromAddress := r.FormValue("fromAddress")
		//fmt.Print("fromAddress: ")
		//fmt.Println(fromAddress)

		var ms services.MailServer
		ms.ID = id
		ms.MailServer = mailServer
		//ms.MailServer = mailServer
		if secureConnection == "on" {
			ms.SecureConnection = true
		} else {
			ms.SecureConnection = false
		}
		ms.Port, err = strconv.Atoi(port)
		if err != nil {
			fmt.Println(err)
		}
		if debug == "on" {
			ms.Debug = true
		} else {
			ms.Debug = false
		}
		ms.Username = username
		ms.Password = password
		ms.FromAddress = fromAddress

		var m services.MailServerService
		m.ClientID = getAuthCodeClient()
		//m.UserID = getHashedUser()
		//m.Hashed = "true"
		m.Token = token.AccessToken
		m.Host = getMailHost()
		m.APIKey = getGatewayAPIKey()
		var res *services.MailResponse
		if id == 0 {
			res = m.AddMailServer(&ms)
		} else {
			res = m.UpdateMailServer(&ms)
		}
		fmt.Println(res)
		if res.Success == true {
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		} else {
			fmt.Println("Mail Server update failed")
			http.Redirect(w, r, "/admin/main", http.StatusFound)
		}
	}
}
