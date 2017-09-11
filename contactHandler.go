package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net/http"
)

// user handlers-----------------------------------------------------
func handleContactSend(w http.ResponseWriter, r *http.Request) {
	ans := r.FormValue("answer")
	fmt.Print("Answer: ")
	fmt.Println(ans)

	key := r.FormValue("key")
	fmt.Print("Key: ")
	fmt.Println(key)

	fromEmail := r.FormValue("fromEmail")
	fmt.Print("fromEmail: ")
	fmt.Println(fromEmail)

	text := r.FormValue("text")
	fmt.Print("text: ")
	fmt.Println(text)
	var c services.ChallengeService
	c.Host = getChallengeHost()
	var ch services.Challenge
	ch.Answer = ans
	ch.Key = key
	cres := c.SendChallenge(&ch)
	fmt.Print("Challenge Res: ")
	fmt.Println(cres)
	if cres.Success == true {
		// get client token
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handleContactForm(w http.ResponseWriter, r *http.Request) {
	var c services.ChallengeService
	c.Host = getChallengeHost()
	res := c.GetChallenge("en_us")
	fmt.Print("challenge: ")
	fmt.Println(res)
	templates.ExecuteTemplate(w, "contact.html", &res)
}

//end user handlers------------------------------------------------
