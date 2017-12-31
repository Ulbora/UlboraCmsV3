package services

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var addIDP int64
var addID2P int64

//var token = testToken

func TestContentPageService_AddContent(t *testing.T) {
	var c ContentService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	var ct Content
	ct.Category = "books1"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "a book"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Hits = 98
	res := c.AddContent(&ct)
	fmt.Print("res: ")
	fmt.Println(res)
	addIDP = res.ID
	if res.Success != true {
		t.Fail()
	}
}
func TestContentPageService_GetPage(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	h, pc := c.GetPage("books1")
	fmt.Print("found page: (99 in db, 0 in cache) ")
	fmt.Println(*pc)
	if h.MetaDesc != "a book" {
		t.Fail()
	}
	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 0 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_GetPage2(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	h, pc := c.GetPage("books1")
	fmt.Print("found page:  (99 in db, 1 in cache)")
	fmt.Println(*pc)
	if h.MetaDesc != "a book" {
		t.Fail()
	}
	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 1 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_GetPage3(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	h, pc := c.GetPage("books1")
	fmt.Print("found page: (101 in db, 0 in cache)")
	fmt.Println(*pc)
	if h.MetaDesc != "a book" {
		t.Fail()
	}
	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 0 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_GetPage4(t *testing.T) {
	time.Sleep(1000 * time.Millisecond)
	var c ContentService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken

	pc := c.GetContentList(c.ClientID)
	fmt.Print("found GetPage4: (101 in db, 0 in cache) ")
	fmt.Println(*pc)

	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 101 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_GetPage5(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	h, pc := c.GetPage("books1")
	fmt.Print("found page: (101 in db, 1 in cache) ")
	fmt.Println(*pc)
	if h.MetaDesc != "a book" {
		t.Fail()
	}
	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 1 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_ClearPage(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	res := c.ClearPage("books1")
	if res != true {
		t.Fail()
	}
}

func TestContentPageService_GetPage6(t *testing.T) {
	time.Sleep(1000 * time.Millisecond)
	var c ContentService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken

	pc := c.GetContentList(c.ClientID)
	fmt.Print("found GetPage6: (102 in db, o in cache)  ")
	fmt.Println(*pc)

	if pc != nil {
		for _, p := range *pc {
			if p.Text != "some book text" || p.Hits != 102 {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}

func TestContentPageService_DeletePage(t *testing.T) {
	var c ContentPageService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = testToken
	c.PageSize = 2
	suc := c.DeletePage("books1")
	fmt.Print("delete page success ")
	fmt.Println(suc)
	if suc != true {
		t.Fail()
	}
}

func TestContentPageService_DeleteContent(t *testing.T) {
	var c ContentService
	c.ClientID = "616"
	c.Host = "http://localhost:3008"
	c.Token = token

	res := c.DeleteContent(strconv.FormatInt(addIDP, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}
