package services

import (
	"fmt"
	"strconv"
	"testing"
)

var addID int64
var addID2 int64
var token = testToken

func TestContentService_AddContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.Category = "books1"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a book"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	res := c.AddContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_UpdateContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.ID = addID
	ct.Category = "music2"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a song"
	ct.Text = "some music text"
	ct.Title = "the best song ever"
	res := c.UpdateContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_UpdateContentHits(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.ID = addID
	ct.Hits = 500
	res := c.UpdateContentHits(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_AddAnotherContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token
	var ct Content
	ct.Category = "books1"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a book"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	res := c.AddContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	addID2 = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_GetContent(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	res := c.GetContent(strconv.FormatInt(addID, 10), "403")
	fmt.Print("res: ")
	fmt.Println(res)
	if res.ID != addID && res.Hits == 500 {
		t.Fail()
	}
}

func TestContentService_GetContentList(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	res := c.GetContentList("403")
	fmt.Print("res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestContentService_GetContentListCategory(t *testing.T) {
	var c ContentService
	c.Host = "http://localhost:3008"
	_, res := c.GetContentListCategory("403", "books1")
	fmt.Print("res category list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestContentService_DeleteContent(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token

	res := c.DeleteContent(strconv.FormatInt(addID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestContentService_DeleteContent2(t *testing.T) {
	var c ContentService
	c.ClientID = "403"
	c.Host = "http://localhost:3008"
	c.Token = token

	res := c.DeleteContent(strconv.FormatInt(addID2, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}
