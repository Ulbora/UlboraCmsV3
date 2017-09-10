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
	Host string
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
	resp, err := http.Get(gURL)
	//fmt.Println(resp)
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			log.Println(error.Error())
		}
	}
	return rtn
}
