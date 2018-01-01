package services

import (
	"fmt"
	"strconv"
)

//ContentPageService ContentPageService
type ContentPageService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
	PageSize int64
}

//GetPage GetPage
func (c *ContentPageService) GetPage(page string) (*PageHead, *[]Content) {
	var rtnC *[]Content // = make([]Content, 0)
	var rtnH *PageHead  // = new(PageHead)
	var ch CacheService
	ch.ClientID = c.ClientID
	ch.PageSize = c.PageSize
	p := ch.ReadPage(page)
	var ccc ContentService
	ccc.ClientID = c.ClientID
	ccc.APIKey = c.APIKey
	ccc.Host = c.Host
	ccc.Token = c.Token
	//fmt.Print("page: ")
	//fmt.Println(page)
	//fmt.Print("page from cache: ")
	//fmt.Println(p)
	if p.PageHeader != nil && p.PageContent != nil {
		//fmt.Println("Reading from Cache")
		rtnH = p.PageHeader
		rtnC = p.PageContent
		if p.Hits != nil {
			//fmt.Print("page before hit update in page service: ")
			//fmt.Println(*p.Hits)
			for _, h := range *p.Hits {
				go func(hit PageHit) {
					//fmt.Print("page in go routine before update in page service: ")
					//fmt.Println(hit)
					content := ccc.GetContent(strconv.FormatInt(hit.ID, 10), c.ClientID)
					//fmt.Print("content in go routine before hit change in page service: ")
					//fmt.Println(content)
					content.Hits += hit.Hits
					//fmt.Print("content in go routine after hit change in page service: ")
					//fmt.Println(content)
					resp := ccc.UpdateContentHits(content)
					if resp.Success != true {
						fmt.Println("content hit update failed")
					}
					//fmt.Print("content update resp in go routine after update in page service: ")
					//fmt.Println(resp)
				}(h)
			}

		}
	} else {
		fmt.Println("Reading from DB")
		rtnH, rtnC = ccc.GetContentListCategory(c.ClientID, page)
		var pc PageCache
		//fmt.Print("page before hit update: ")
		//fmt.Println(*rtnC)
		clist := make([]Content, 0)
		pc.PageHeader = rtnH
		for _, fp := range *rtnC {
			fp.Hits = 0
			clist = append(clist, fp)
		}
		pc.PageContent = &clist
		pc.PageName = page
		rtnC = &clist
		ch.CachePage(pc)
	}

	return rtnH, rtnC
}

//ClearPage ClearPage
func (c *ContentPageService) ClearPage(pageName string) bool {
	var rtn = false
	var ch CacheService
	ch.ClientID = c.ClientID
	ch.PageSize = c.PageSize
	suc, hits := ch.RemovePage(pageName)
	if len(*hits) > 0 {
		var ccc ContentService
		ccc.ClientID = c.ClientID
		ccc.Host = c.Host
		ccc.Token = c.Token
		ccc.APIKey = c.APIKey
		pageList := ccc.GetContentList(c.ClientID)
		//fmt.Print("found page list in clearpage: ")
		//fmt.Println(pageList)
		for _, content := range *pageList {
			//var contentToSave *Content
			for _, h := range *hits {
				if content.ID == h.ID {
					content.Hits += h.Hits
					break
				}
			}
			go func(cont Content) {
				//fmt.Print("updating content: ")
				//fmt.Println(cont)
				resp := ccc.UpdateContentHits(&cont)
				//fmt.Print("updating content resp: ")
				//fmt.Println(resp)
				if resp.Success != true {
					fmt.Println("content hit update failed")
				}
			}(content)
		}
	}
	rtn = suc
	return rtn
}

//DeletePage DeletePage
func (c *ContentPageService) DeletePage(pageName string) bool {
	var ch CacheService
	ch.ClientID = c.ClientID
	ch.PageSize = c.PageSize
	suc := ch.DeletePage(pageName)
	return suc
}
