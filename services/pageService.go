package services

//SavePage SavePage
func (c *ContentService) SavePage(page string) bool {
	var rtn = false

	return rtn
}

//GetPage GetPage
func (c *ContentService) GetPage(page string) (*PageHead, *[]Content) {
	var rtn = make([]Content, 0)
	var pghead = new(PageHead)
	var ch CacheService
	ch.ClientID = c.ClientID
	p := ch.ReadPage(page)
	if p.PageHeader != nil {

	}

	return pghead, &rtn
}

//ClearPage ClearPage
func (c *ContentService) ClearPage(page string) bool {
	var rtn = false

	return rtn
}

//DeletePage DeletePage
func (c *ContentService) DeletePage(page string) bool {
	var rtn = false

	return rtn
}
