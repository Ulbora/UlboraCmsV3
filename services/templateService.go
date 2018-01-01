package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//TemplateService service
type TemplateService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//Template template
type Template struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Application string `json:"application"`
	Active      bool   `json:"active"`
	ClientID    int64  `json:"clientId"`
}

//TemplateResponse res
type TemplateResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddTemplate add template
func (t *TemplateService) AddTemplate(tmpl *Template) *TemplateResponse {
	var rtn = new(TemplateResponse)
	var addURL = t.Host + "/rs/template/add"
	tmpl.Name = stripSpace(tmpl.Name)
	aJSON, err := json.Marshal(tmpl)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+t.Token)
			req.Header.Set("clientId", t.ClientID)
			req.Header.Set("userId", t.UserID)
			req.Header.Set("hashed", t.Hashed)
			req.Header.Set("apiKey", t.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Template Add err: ")
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

//UpdateTemplate update template
func (t *TemplateService) UpdateTemplate(tmpl *Template) *TemplateResponse {
	var rtn = new(TemplateResponse)
	var upURL = t.Host + "/rs/template/updateActive"

	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(tmpl)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+t.Token)
			req.Header.Set("clientId", t.ClientID)
			req.Header.Set("userId", t.UserID)
			req.Header.Set("hashed", t.Hashed)
			req.Header.Set("apiKey", t.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Template Service Update err: ")
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

// GetTemplate get template
func (t *TemplateService) GetTemplate(app string, clientID string) *Template {
	var rtn = new(Template)
	var gURL = t.Host + "/rs/template/get/" + app + "/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", t.ClientID)
		req.Header.Set("apiKey", t.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Template Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return rtn
}

// GetTemplateList get template list by client
func (t *TemplateService) GetTemplateList(app string, clientID string) *[]Template {
	var rtn = make([]Template, 0)
	var gURL = t.Host + "/rs/template/list/" + app + "/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", t.ClientID)
		req.Header.Set("apiKey", t.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Template Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}

	return &rtn
}

// DeleteTemplate delete template
func (t *TemplateService) DeleteTemplate(id string) *TemplateResponse {
	var rtn = new(TemplateResponse)
	var gURL = t.Host + "/rs/template/delete/" + id + "/" + t.ClientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+t.Token)
		req.Header.Set("clientId", t.ClientID)
		req.Header.Set("userId", t.UserID)
		req.Header.Set("hashed", t.Hashed)
		req.Header.Set("apiKey", t.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Template Service delete err: ")
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
