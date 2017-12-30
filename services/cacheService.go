package services

import (
	"sync"
)

//CacheService CacheService
type CacheService struct {
	ClientID string
}

//PageCache PageCache
type PageCache struct {
	PageName    string
	PageHeader  *PageHead
	PageContent *[]Content
	Hits        *[]PageHit
}

//PageHit PageHit for updating database
type PageHit struct {
	ID   int64
	Hits int64
}

var pageCache = make(map[string]PageCache)
var mu sync.Mutex

//CachePage CachePage
func (c *CacheService) CachePage(page PageCache) {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + page.PageName
	pageCache[key] = page
}

//ReadPage ReadPage
func (c *CacheService) ReadPage(pageName string) *PageCache {
	mu.Lock()
	defer mu.Unlock()
	var rtn PageCache
	hits := make([]PageHit, 0)
	key := c.ClientID + ":" + pageName
	cp := pageCache[key]
	//fmt.Print("page found from cache: ")
	//fmt.Println(cp)
	//var pghead = cp.PageHeader
	var cu = false
	var hu = false
	cont := make([]Content, 0)
	if cp.PageName != "" && cp.PageContent != nil {
		var cList = cp.PageContent
		for _, p := range *cList {
			p.Hits++
			if p.Hits >= 100 {
				var h PageHit
				h.ID = p.ID
				h.Hits = p.Hits
				hits = append(hits, h)
				p.Hits = 0
				// go func(cont Content) {
				// 	content := c.GetContent(strconv.FormatInt(cont.ID, 10), clientID)
				// 	content.Hits += cont.Hits
				// 	resp := c.UpdateContentHits(content)
				// 	if resp.Success != true {
				// 		fmt.Println("content hit update failed")
				// 	}
				// }(pc)
				if hu == false {
					hu = true
				}

			}
			cont = append(cont, p)
			if cu == false {
				cu = true
			}

		}
		cp.PageContent = &cont
		pageCache[key] = cp
	}

	//return values
	if hu == true {
		rtn.Hits = &hits
	}
	rtn.PageHeader = cp.PageHeader
	if cu == true {
		rtn.PageContent = &cont
	}
	return &rtn
}

//RemovePage RemovePage
func (c *CacheService) RemovePage(pageName string) *[]PageHit {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + pageName
	cp := pageCache[key]
	//var cachedList = cp.PageContent
	hits := make([]PageHit, 0)
	if cp.PageName != "" && cp.PageContent != nil {
		var cList = cp.PageContent
		for _, p := range *cList {
			//p.Hits >= 100 {
			var h PageHit
			h.ID = p.ID
			h.Hits = p.Hits
			hits = append(hits, h)
		}
	}

	// _, pageList := c.GetContentListCategory(clientID, pageName)

	// for _, content := range *pageList {
	// 	//var contentToSave *Content
	// 	for _, pc := range *cachedList {
	// 		if content.ID == pc.ID {
	// 			content.Hits += pc.Hits
	// 			break
	// 		}
	// 	}
	// 	// go func(cont Content) {
	// 	// 	resp := c.UpdateContent(&cont)
	// 	// 	fmt.Print("updating content ID: ")
	// 	// 	fmt.Println(cont.ID)
	// 	// 	if resp.Success != true {
	// 	// 		fmt.Println("content hit update failed")
	// 	// 	}
	// 	// }(content)
	// }

	//key := clientID + ":" + pageName
	delete(pageCache, key)
	return &hits
}

//DeletePage DeletePage
func (c *CacheService) DeletePage(pageName string) {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + pageName
	delete(pageCache, key)
}
