package services

import (
	"fmt"
	"testing"
)

var addMailID int64
var mailToken = testToken

func TestMailServerService_AddMailServer(t *testing.T) {
	var m MailServerService
	m.ClientID = "403"
	m.Host = "http://localhost:3002"
	m.Token = mailToken
	var ms MailServer
	ms.MailServer = "mail.some.com"
	ms.SecureConnection = true
	ms.Port = 465
	ms.Debug = false
	ms.Username = "ken"
	ms.Password = "ken"
	ms.FromAddress = "ken@ken.com"
	ms.ClientID = 403
	res := m.AddMailServer(&ms)
	fmt.Print("mail server add res: ")
	fmt.Println(res)
	addMailID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestMailServerService_UpdateMailServer(t *testing.T) {
	var m MailServerService
	m.ClientID = "403"
	m.Host = "http://localhost:3002"
	m.Token = mailToken
	var ms MailServer
	ms.ID = addMailID
	ms.MailServer = "mail.some.com"
	ms.SecureConnection = true
	ms.Port = 465
	ms.Debug = true
	ms.Username = "ken"
	ms.Password = "ken"
	ms.FromAddress = "ken1@ken.com"
	ms.ClientID = 403
	res := m.UpdateMailServer(&ms)
	fmt.Print("mail server add res: ")
	fmt.Println(res)
	addMailID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_GetMailServer(t *testing.T) {
	var m MailServerService
	m.ClientID = "403"
	m.Host = "http://localhost:3002"
	m.Token = mailToken
	res := m.GetMailServer()
	fmt.Print("res server get: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

// func TestMailServerService_SendMail(t *testing.T) {
// 	var m MailServerService
// 	m.ClientID = "403"
// 	m.Host = "http://localhost:3002"
// 	m.Token = token
// 	var ms MailMessage
// 	ms.FromEmail
// 	res := m.SendMail(&ms)
// 	fmt.Print("send mail res: ")
// 	fmt.Println(res)
// 	addMailID = res.ID
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }
