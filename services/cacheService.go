package services

import (
	"sync"
)

//CacheService CacheService
type CacheService struct {
	ClientID string
	PageSize int64
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
			if p.Hits >= c.PageSize {
				var h PageHit
				h.ID = p.ID
				h.Hits = p.Hits
				hits = append(hits, h)
				p.Hits = 0
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
func (c *CacheService) RemovePage(pageName string) (bool, *[]PageHit) {
	mu.Lock()
	defer mu.Unlock()
	key := c.ClientID + ":" + pageName
	//fmt.Print("key in remove: ")
	//fmt.Println(key)
	cp := pageCache[key]
	//fmt.Print("found cache in remove: ")
	//fmt.Println(*cp.PageContent)
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
	//fmt.Print("found cache hits in remove: ")
	//fmt.Println(hits)
	//key := clientID + ":" + pageName
	delete(pageCache, key)
	var rtn = false
	fcp := pageCache[key]
	if fcp.PageHeader == nil {
		rtn = true
	}
	return rtn, &hits
}

//DeletePage DeletePage
func (c *CacheService) DeletePage(pageName string) bool {
	mu.Lock()
	defer mu.Unlock()
	var rtn = false
	key := c.ClientID + ":" + pageName
	delete(pageCache, key)
	cp := pageCache[key]
	if cp.PageHeader == nil {
		rtn = true
	}
	return rtn
}
