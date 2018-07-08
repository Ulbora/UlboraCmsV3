package main

import (
	services "UlboraCmsV3/services"
	"fmt"
	"net"
	"net/http"
)

// user handlers-----------------------------------------------------
func handleContactSend(w http.ResponseWriter, r *http.Request) {
	var proceed = false
	fromEmail := r.FormValue("fromEmail")
	//fmt.Print("fromEmail: ")
	//fmt.Println(fromEmail)

	text := r.FormValue("text")
	//fmt.Print("text: ")
	//fmt.Println(text)

	recaptchaResp := r.FormValue("g-recaptcha-response")
	fmt.Print("recaptchaResp: ")
	fmt.Println(recaptchaResp)
	if recaptchaResp != "" {
		// do recaptcha

		var ipAddr string

		addrs, _ := net.InterfaceAddrs()
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ipAddr = ipnet.IP.String()
					break
				}
			}
		}
		// fmt.Print("captcha secret: ")
		// fmt.Println(h.CaptchaSecret)

		// fmt.Print("client ip address: ")
		// fmt.Println(ipAddr)

		// fmt.Print("recaptchaResp: ")
		// fmt.Println(recaptchaResp)

		var s services.CaptchaService
		s.Host = getCaptchaHost()
		var cap services.Captcha
		cap.Remoteip = ipAddr
		cap.Secret = captchaSecret
		cap.Response = recaptchaResp
		res := s.SendCaptchaCall(cap)
		if res.Success {
			proceed = true
		}
	} else {
		var c services.ChallengeService
		c.Host = getChallengeHost()
		c.ClientID = getAuthCodeClient()
		c.APIKey = getGatewayAPIKey()

		var ch services.Challenge
		ans := r.FormValue("answer")
		//fmt.Print("Answer: ")
		//fmt.Println(ans)

		key := r.FormValue("key")
		//fmt.Print("Key: ")
		//fmt.Println(key)
		ch.Answer = ans
		ch.Key = key

		cres := c.SendChallenge(&ch)
		if cres.Success {
			proceed = true
		}
	}

	//fmt.Print("Challenge Res: ")
	//fmt.Println(cres)
	if proceed {
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
