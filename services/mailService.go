package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//MailServerService service
type MailServerService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//MailServer server
type MailServer struct {
	ID               int64  `json:"id"`
	MailServer       string `json:"mailServer"`
	SecureConnection bool   `json:"secureConnection"`
	Port             int    `json:"port"`
	Debug            bool   `json:"debug"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	FromAddress      string `json:"fromAddress"`
	ClientID         int64  `json:"clientId"`
}

//MailServerResponse res
type MailServerResponse struct {
	Success bool       `json:"success"`
	Code    int        `json:"code"`
	Server  MailServer `json:"mailServer"`
}

//MailResponse res
type MailResponse struct {
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//MailMessage to send
type MailMessage struct {
	ToEmail     string `json:"toEmail"`
	FromEmail   string `json:"fromEmail"`
	Subject     string `json:"subject"`
	TextMessage string `json:"text"`
	HTMLMessage string `json:"html"`
}

//AddMailServer add mail server
func (m *MailServerService) AddMailServer(mailServer *MailServer) *MailResponse {
	var rtn = new(MailResponse)
	var addURL = m.Host + "/rs/mailServer/add"
	aJSON, err := json.Marshal(mailServer)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+m.Token)
			req.Header.Set("clientId", m.ClientID)
			req.Header.Set("userId", m.UserID)
			req.Header.Set("hashed", m.Hashed)
			req.Header.Set("apiKey", m.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("MailServer Service Add err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

//UpdateMailServer update mail server
func (m *MailServerService) UpdateMailServer(mailServer *MailServer) *MailResponse {
	var rtn = new(MailResponse)
	var upURL = m.Host + "/rs/mailServer/update"
	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(mailServer)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+m.Token)
			req.Header.Set("clientId", m.ClientID)
			req.Header.Set("userId", m.UserID)
			req.Header.Set("hashed", m.Hashed)
			req.Header.Set("apiKey", m.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("mail server Update err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

// GetMailServer get mail server
func (m *MailServerService) GetMailServer() *MailServerResponse {
	var rtn = new(MailServerResponse)
	var gURL = m.Host + "/rs/mailServer/get"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Authorization", "Bearer "+m.Token)
		req.Header.Set("clientId", m.ClientID)
		req.Header.Set("apiKey", m.APIKey)
		//req.Header.Set("userId", m.UserID)
		//req.Header.Set("hashed", m.Hashed)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Mail Server get err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			rtn.Code = resp.StatusCode
		}
	}
	return rtn
}

//SendMail send mail
func (m *MailServerService) SendMail(mailMessage *MailMessage) *MailResponse {
	var rtn = new(MailResponse)
	var sendURL = m.Host + "/rs/mail/send"
	aJSON, err := json.Marshal(mailMessage)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", sendURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+m.Token)
			req.Header.Set("clientId", m.ClientID)
			req.Header.Set("userId", m.UserID)
			req.Header.Set("hashed", m.Hashed)
			req.Header.Set("apiKey", m.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Mail send err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}
