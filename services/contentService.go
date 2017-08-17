package services

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//ContentService service
type ContentService struct {
	Token    string
	ClientID string
	Host     string
}

//Content content
type Content struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Category          string    `json:"category"`
	CreateDate        time.Time `json:"createDate"`
	ModifiedDate      time.Time `json:"modifiedDate"`
	Hits              int64     `json:"hits"`
	MetaAuthorName    string    `json:"metaAuthorName"`
	MetaDesc          string    `json:"metaDesc"`
	MetaKeyWords      string    `json:"metaKeyWords"`
	MetaRobotKeyWords string    `json:"metaRobotKeyWords"`
	Text              string    `json:"text"`
	ClientID          int64     `json:"clientId"`
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
	var rtn = new(Response)
	var upURL = c.Host + "/rs/content/update"
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

// //UpdateContent update content
// func (c *ContentService) UpdateContentHits(content *Content) *Response {
// 	var rtn = new(Response)
// 	var upURL = c.Host + "/rs/content/update"
// 	content.Text = b64.StdEncoding.EncodeToString([]byte(content.Text))
// 	//fmt.Println(content.Text)
// 	aJSON, err := json.Marshal(content)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
// 		if rErr != nil {
// 			fmt.Print("request err: ")
// 			fmt.Println(rErr)
// 		} else {
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Header.Set("Authorization", "Bearer "+c.Token)
// 			req.Header.Set("clientId", c.ClientID)
// 			client := &http.Client{}
// 			resp, cErr := client.Do(req)
// 			if cErr != nil {
// 				fmt.Print("Content Service Update err: ")
// 				fmt.Println(cErr)
// 			} else {
// 				defer resp.Body.Close()
// 				decoder := json.NewDecoder(resp.Body)
// 				error := decoder.Decode(&rtn)
// 				if error != nil {
// 					log.Println(error.Error())
// 				}
// 				rtn.Code = resp.StatusCode
// 			}
// 		}
// 	}
// 	return rtn
// }

// GetContent get content
func (c *ContentService) GetContent(id string, clientID string) *Content {
	var rtn = new(Content)
	var gURL = c.Host + "/rs/content/get/" + id + "/" + clientID
	fmt.Println(gURL)
	resp, err := http.Get(gURL)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			log.Println(error.Error())
		}
		//base64 decode text
	}
	return rtn
}
