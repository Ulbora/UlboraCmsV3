package services

import "sync"
import "fmt"
import "strconv"

//PageCache PageCache
type PageCache struct {
	PageHeader  *PageHead
	PageContent *[]Content
}

var pageCache = make(map[string]PageCache)
var mu sync.Mutex

//CachePage CachePage
func (c *ContentService) CachePage(clientID string, pageName string, page PageCache) {
	mu.Lock()
	defer mu.Unlock()
	key := clientID + ":" + pageName
	pageCache[key] = page
}

//ReadPage ReadPage
func (c *ContentService) ReadPage(clientID string, pageName string) (*PageHead, *[]Content) {
	mu.Lock()
	defer mu.Unlock()
	key := clientID + ":" + pageName
	cp := pageCache[key]

	var pghead = cp.PageHeader
	var cList = cp.PageContent
	rtn := make([]Content, 0)
	for _, pc := range *cList {
		pc.Hits++
		if pc.Hits >= 100 {
			go func(cont Content) {
				content := c.GetContent(strconv.FormatInt(cont.ID, 10), clientID)
				content.Hits += cont.Hits
				resp := c.UpdateContentHits(content)
				if resp.Success != true {
					fmt.Println("content hit update failed")
				}
			}(pc)
			pc.Hits = 0
		}
		rtn = append(rtn, pc)
	}
	cp.PageContent = &rtn
	pageCache[key] = cp

	return pghead, &rtn
}

//RemovePage RemovePage
func (c *ContentService) RemovePage(clientID string, pageName string) {
	mu.Lock()
	defer mu.Unlock()
	key := clientID + ":" + pageName
	cp := pageCache[key]
	var cachedList = cp.PageContent

	_, pageList := c.GetContentListCategory(clientID, pageName)

	for _, content := range *pageList {
		//var contentToSave *Content
		for _, pc := range *cachedList {
			if content.ID == pc.ID {
				content.Hits += pc.Hits
				break
			}
		}
		go func(cont Content) {
			resp := c.UpdateContent(&cont)
			fmt.Print("updating content ID: ")
			fmt.Println(cont.ID)
			if resp.Success != true {
				fmt.Println("content hit update failed")
			}
		}(content)
	}

	//key := clientID + ":" + pageName
	delete(pageCache, key)
}

//DeletePage DeletePage
func (c *ContentService) DeletePage(clientID string, pageName string) {
	mu.Lock()
	defer mu.Unlock()
	key := clientID + ":" + pageName
	delete(pageCache, key)
}
