package services

//GetPage GetPage
func (c *ContentService) GetPage(clientID string, category string) (*PageHead, *[]Content) {
	var rtn = make([]Content, 0)
	var pghead = new(PageHead)

	return pghead, &rtn
}
