package services

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//ImageService service
type ImageService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

// Image the image info
type Image struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	FileExtension string `json:"fileExtension"`
	ClientID      int64  `json:"clientId"`
	ImageURL      string `json:"imageUrl"`
}

// UploadedFile file
type UploadedFile struct {
	Name             string
	Size             int64
	OriginalFileName string
	FileData         []byte
}

type imageFile struct {
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	FileExtension string `json:"fileExtension"`
	FileData      string `json:"fileData"`
}

//ImageResponse res
type ImageResponse struct {
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//AddImage add image
func (i *ImageService) AddImage(image *UploadedFile) *ImageResponse {
	var rtn = new(ImageResponse)
	var addURL = i.Host + "/rs/image/add"

	s64 := b64.StdEncoding.EncodeToString(image.FileData)

	igf := new(imageFile)
	igf.FileData = s64
	igf.Name = stripSpace(image.Name)
	igf.Size = image.Size
	igf.FileExtension = getExt(stripSpace(image.OriginalFileName))
	aJSON, err := json.Marshal(igf)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+i.Token)
			req.Header.Set("clientId", i.ClientID)
			req.Header.Set("apiKey", i.APIKey)
			//req.Header.Set("userId", i.UserID)
			//req.Header.Set("hashed", i.Hashed)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Image Add err: ")
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

//GetList add image
func (i *ImageService) GetList() *[]Image {
	var rtn = make([]Image, 0)
	var gURL = i.Host + "/rs/image/list/100"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Authorization", "Bearer "+i.Token)
		req.Header.Set("clientId", i.ClientID)
		req.Header.Set("apiKey", i.APIKey)
		//req.Header.Set("userId", m.UserID)
		//req.Header.Set("hashed", m.Hashed)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("image list get err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			//fmt.Print("resp body: ")
			//fmt.Println(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return &rtn
}

// DeleteImage delete image
func (i *ImageService) DeleteImage(id string) *ImageResponse {
	var rtn = new(ImageResponse)
	var gURL = i.Host + "/rs/image/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+i.Token)
		req.Header.Set("clientId", i.ClientID)
		req.Header.Set("apiKey", i.APIKey)
		//req.Header.Set("userId", i.UserID)
		//req.Header.Set("hashed", i.Hashed)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Image Service delete err: ")
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

func getExt(name string) string {
	var rtn = ""
	i := strings.LastIndex(name, ".")
	rtn = name[i+1:]
	return rtn
}

func stripSpace(name string) string {
	var rtn = ""
	rtn = strings.Replace(name, " ", "", -1)
	return rtn
}
