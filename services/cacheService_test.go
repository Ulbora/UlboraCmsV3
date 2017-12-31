package services

import (
	"fmt"
	"testing"
)

var ccidStr = "111"
var ccid int64 = 111

func TestCacheService_CachePage(t *testing.T) {
	var c CacheService
	c.ClientID = ccidStr
	c.PageSize = 2
	var p PageCache

	var ct Content
	ct.Category = "main"
	ct.ID = 1
	ct.Hits = 99
	ct.ClientID = ccid
	clist := make([]Content, 0)
	clist = append(clist, ct)
	p.PageName = "main"
	p.PageContent = &clist

	c.CachePage(p)
	cp := c.ReadPage("main")
	//print("hits: ")
	//println(cp)
	if len(*cp.Hits) != 1 {
		t.Fail()
	}
	for _, page := range *cp.PageContent {
		fmt.Print("page name: ")
		fmt.Println(page.Category)
		fmt.Print("page id: ")
		fmt.Println(page.ID)
		fmt.Print("page hits: ")
		fmt.Println(page.Hits)
		if page.ID != 1 {
			t.Fail()
		}
	}
	for _, page := range *cp.Hits {
		fmt.Print("page hit to updated to database: ")
		fmt.Println(page.Hits)
		if page.Hits != 100 {
			t.Fail()
		}
	}
}

func TestCacheService_ReadPage(t *testing.T) {
	var c CacheService
	c.ClientID = ccidStr
	c.PageSize = 2
	cp := c.ReadPage("main")
	//print("hits: ")
	//println(cp)
	if cp.Hits != nil && len(*cp.Hits) != 0 {
		t.Fail()
	}
	for _, page := range *cp.PageContent {
		fmt.Print("page name: ")
		fmt.Println(page.Category)
		fmt.Print("page id: ")
		fmt.Println(page.ID)
		fmt.Print("page hits: ")
		fmt.Println(page.Hits)
		if page.ID != 1 || page.Hits != 1 {
			t.Fail()
		}
	}
}

func TestCacheService_RemovePage(t *testing.T) {
	var c CacheService
	c.ClientID = ccidStr
	c.PageSize = 2
	suc, h := c.RemovePage("main")
	for _, hit := range *h {
		fmt.Print("page hit to updated to database: ")
		fmt.Println(hit.Hits)
		if hit.Hits != 1 || suc != true {
			t.Fail()
		}
	}
}

func TestCacheService_DeletePage(t *testing.T) {
	var c CacheService
	c.ClientID = ccidStr
	c.PageSize = 2
	suc := c.DeletePage("main")
	cp := c.ReadPage("main")
	fmt.Print("page after delete: ")
	fmt.Println(cp)

	fmt.Print("page content after delete: ")
	fmt.Println(cp.PageContent)

	fmt.Print("page hits after delete: ")
	fmt.Println(cp.Hits)

	fmt.Print("page header after delete: ")
	fmt.Println(cp.PageHeader)
	if cp.PageHeader != nil || suc != true {
		t.Fail()
	}
}
