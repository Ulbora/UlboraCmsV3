package services

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//ContentService service
type ContentService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//Content content
type Content struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Category          string    `json:"category"`
	CreateDate        time.Time `json:"createDate"`
	ModifiedDate      time.Time `json:"modifiedDate"`
	UseModifiedDate   bool      `json:"useModifiedDate"`
	Hits              int64     `json:"hits"`
	MetaAuthorName    string    `json:"metaAuthorName"`
	MetaDesc          string    `json:"metaDesc"`
	MetaKeyWords      string    `json:"metaKeyWords"`
	MetaRobotKeyWords string    `json:"metaRobotKeyWords"`
	Text              string    `json:"text"`
	TextHTML          template.HTML
	SortOrder         int   `json:"sortOrder"`
	Archived          bool  `json:"archived"`
	ClientID          int64 `json:"clientId"`
}

// PageHead used for page head
type PageHead struct {
	Title        string
	MetaAuthor   string
	MetaDesc     string
	MetaKeyWords string
}

//Response res
type Response struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddContent add content
func (c *ContentService) AddContent(content *Content) *Response {
	var rtn = new(Response)
	var addURL = c.Host + "/rs/content/add"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			req.Header.Set("userId", c.UserID)
			req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", c.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Content Service Add err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

//UpdateContent update content
func (c *ContentService) UpdateContent(content *Content) *Response {
	//fmt.Print("Content Service at start: ")
	var rtn = new(Response)
	var upURL = c.Host + "/rs/content/update"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(content)
	if err != nil {
		fmt.Print("marchal err: ")
		fmt.Println(err)
	} else {
		//fmt.Print("Content Service before req: ")
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			req.Header.Set("userId", c.UserID)
			req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", c.APIKey)
			//fmt.Print("Content Service before rest call: ")
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Content Service Update err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					fmt.Print("Content Service Update decode err: ")
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	//fmt.Print("Content Service Update leaving: ")
	return rtn
}

//UpdateContentHits update content hits
func (c *ContentService) UpdateContentHits(content *Content) *Response {
	var rtn = new(Response)
	var upURL = c.Host + "/rs/content/hits"
	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			req.Header.Set("userId", c.UserID)
			req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", c.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Content Service Update err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

// GetContent get content
func (c *ContentService) GetContent(id string, clientID string) *Content {
	var rtn = new(Content)
	var gURL = c.Host + "/rs/content/get/" + id + "/" + clientID
	//fmt.Println(gURL)
	//resp, err := http.Get(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("apiKey", c.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Content Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			txt, err := b64.StdEncoding.DecodeString(rtn.Text)
			if err != nil {
				fmt.Println(err)
			} else {
				rtn.Text = string(txt)
			}
		}
	}
	return rtn
}

// GetContentList get content list by client
func (c *ContentService) GetContentList(clientID string) *[]Content {
	var rtn = make([]Content, 0)
	var gURL = c.Host + "/rs/content/list/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("apiKey", c.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Content Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			//var cont = new(Content)
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			for r := range rtn {
				txt, err := b64.StdEncoding.DecodeString(rtn[r].Text)
				if err != nil {
					fmt.Println(err)
				} else {
					rtn[r].Text = string(txt)
					//fmt.Println(rtn[r].Text)
				}
				//fmt.Println(rtn[r].ModifiedDate.Year())
				if rtn[r].ModifiedDate.Year() != 1 {
					rtn[r].UseModifiedDate = true
				}
			}
		}
	}
	return &rtn
}

// GetContentListCategory get content list by client
func (c *ContentService) GetContentListCategory(clientID string, category string) (*PageHead, *[]Content) {
	var rtn = make([]Content, 0)
	var pghead = new(PageHead)
	var gURL = c.Host + "/rs/content/list/" + clientID + "/" + category
	//fmt.Println(gURL)
	//resp, err := http.Get(gURL)
	//fmt.Print("get category list")
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("apiKey", c.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Content Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			//var cont = new(Content)
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			for r := range rtn {
				txt, err := b64.StdEncoding.DecodeString(rtn[r].Text)
				if err != nil {
					fmt.Println(err)
				} else {
					rtn[r].Text = string(txt)
					//fmt.Println(rtn[r].Text)
					rtn[r].TextHTML = template.HTML(rtn[r].Text)
					if r == 0 {
						pghead.MetaAuthor = rtn[r].MetaAuthorName
						pghead.MetaDesc = rtn[r].MetaDesc
						pghead.MetaKeyWords = rtn[r].MetaKeyWords
						pghead.Title = rtn[r].Title
					}
				}
				//fmt.Println(rtn[r].ModifiedDate.Year())
				if rtn[r].ModifiedDate.Year() != 1 {
					rtn[r].UseModifiedDate = true
				}
			}
		}
	}
	return pghead, &rtn
}

// DeleteContent delete content
func (c *ContentService) DeleteContent(id string) *Response {
	var rtn = new(Response)
	var gURL = c.Host + "/rs/content/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.Token)
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("userId", c.UserID)
		req.Header.Set("hashed", c.Hashed)
		req.Header.Set("apiKey", c.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Content Service delete err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			rtn.Code = resp.StatusCode
		}
	}
	return rtn
}
