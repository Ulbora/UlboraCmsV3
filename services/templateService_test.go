package services

import (
	"fmt"
	"strconv"
	"testing"
)

var tempAddID int64
var tempAddID2 int64
var tempToken = testToken

func TestTemplateService_AddTemplate(t *testing.T) {
	var ts TemplateService
	ts.ClientID = "403"
	ts.Host = "http://localhost:3009"
	ts.Token = tempToken
	var tm Template
	tm.Name = "test temp1"
	tm.Application = "cms"
	tm.ClientID = 7
	res := ts.AddTemplate(&tm)
	fmt.Print("res: ")
	fmt.Println(res)
	tempAddID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestTemplateService_UpdateTemplate(t *testing.T) {
	var ts TemplateService
	ts.ClientID = "403"
	ts.Host = "http://localhost:3009"
	ts.Token = tempToken
	var tm Template
	tm.ID = tempAddID
	tm.Application = "cms"
	res := ts.UpdateTemplate(&tm)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestTemplateService_GetTemplate(t *testing.T) {
	var ts TemplateService
	ts.Host = "http://localhost:3009"
	res := ts.GetTemplate("cms", "403")
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Active != true {
		t.Fail()
	}
}

func TestTemplateService_GetTemplateList(t *testing.T) {
	var ts TemplateService
	ts.Host = "http://localhost:3009"
	res := ts.GetTemplateList("cms", "403")
	fmt.Print("res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestTemplateService_DeleteTemplate(t *testing.T) {
	var ts TemplateService
	ts.ClientID = "403"
	ts.Host = "http://localhost:3009"
	ts.Token = tempToken
	res := ts.DeleteTemplate(strconv.FormatInt(tempAddID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
