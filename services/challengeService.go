package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ChallengeService service
type ChallengeService struct {
	Host     string
	ClientID string
	APIKey   string
}

//Challenge template
type Challenge struct {
	Question string `json:"question"`
	Key      string `json:"key"`
	Answer   string `json:"answer"`
}

//ChallengeResponse res
type ChallengeResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//SendChallenge challenge
func (c *ChallengeService) SendChallenge(chal *Challenge) *ChallengeResponse {
	var rtn = new(ChallengeResponse)
	var sURL = c.Host + "/rs/challenge"
	aJSON, err := json.Marshal(chal)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("clientId", c.ClientID)
			req.Header.Set("apiKey", c.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Challenge send err: ")
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

// GetChallenge get challenge
func (c *ChallengeService) GetChallenge(lan string) *Challenge {
	var rtn = new(Challenge)
	var gURL = c.Host + "/rs/challenge/" + lan
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("apiKey", c.APIKey)
		//req.Header.Set("userId", m.UserID)
		//req.Header.Set("hashed", m.Hashed)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("challange get err: ")
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
