package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
)

// user handlers-----------------------------------------------------
func handleContactSend(w http.ResponseWriter, r *http.Request) {
	ans := r.FormValue("answer")
	//fmt.Print("Answer: ")
	//fmt.Println(ans)

	key := r.FormValue("key")
	//fmt.Print("Key: ")
	//fmt.Println(key)

	fromEmail := r.FormValue("fromEmail")
	//fmt.Print("fromEmail: ")
	//fmt.Println(fromEmail)

	text := r.FormValue("text")
	//fmt.Print("text: ")
	//fmt.Println(text)
	var c services.ChallengeService
	c.Host = getChallengeHost()
	c.ClientID = getAuthCodeClient()
	c.APIKey = getGatewayAPIKey()
	var ch services.Challenge
	ch.Answer = ans
	ch.Key = key
	cres := c.SendChallenge(&ch)
	//fmt.Print("Challenge Res: ")
	//fmt.Println(cres)
	if cres.Success == true {
		// get client token
		getCredentialsToken()
		var m services.MailServerService
		m.ClientID = getAuthCodeClient()
		m.APIKey = getGatewayAPIKey()
		m.Token = credentialToken.AccessToken
		m.Host = getMailHost()
		var mm services.MailMessage
		mm.FromEmail = fromEmail
		mm.Subject = "Ulbora CMS V3 message"
		mm.TextMessage = text
		mres := m.SendMail(&mm)
		fmt.Print("Send Mail Res: ")
		fmt.Println(mres)
		if mres.Success != true {
			fmt.Println("Sending mail failed from " + fromEmail)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleContactForm(w http.ResponseWriter, r *http.Request) {
	var c services.ChallengeService
	c.Host = getChallengeHost()
	c.ClientID = getAuthCodeClient()
	c.APIKey = getGatewayAPIKey()
	res := c.GetChallenge("en_us")
	//fmt.Print("challenge: ")
	//fmt.Println(res)
	templates.ExecuteTemplate(w, "contact.html", &res)
}

//end user handlers------------------------------------------------
