package services

//GetPage GetPage
func (c *ContentService) SavePage(clientID string, page string) bool {
	var rtn = false

	return rtn
}

//GetPage GetPage
func (c *ContentService) GetPage(clientID string, page string) (*PageHead, *[]Content) {
	var rtn = make([]Content, 0)
	var pghead = new(PageHead)

	return pghead, &rtn
}

//ClearPage ClearPage
func (c *ContentService) ClearPage(clientID string, page string) bool {
	var rtn = false

	return rtn
}
