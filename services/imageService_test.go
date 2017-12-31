package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

var imgID int64
var imgToken = testToken

func TestImageService_getExt(t *testing.T) {
	var name = "tester.jpg"
	ext := getExt(name)
	fmt.Print("ext: ")
	fmt.Println(ext)
	if ext != "jpg" {
		t.Fail()
	}
}

func TestImageService_stripSpace(t *testing.T) {
	var name = "tes ter .jpg"
	fn := stripSpace(name)
	fmt.Print("name: ")
	fmt.Println(fn)
	if fn != "tester.jpg" {
		t.Fail()
	}
}

func TestImageService_AddImage(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = imgToken

	imgfile, err := os.Open("./testFiles/upload/test.jpg")
	if err != nil {
		fmt.Println("jpg file not found!")
		os.Exit(1)
	}
	defer imgfile.Close()

	var ui UploadedFile
	ui.Name = "testfile"
	ui.OriginalFileName = imgfile.Name()
	data, err := ioutil.ReadAll(imgfile)
	if err != nil {
		fmt.Println(err)
	}

	cur, err := imgfile.Seek(0, 1)
	size, err := imgfile.Seek(0, 2)
	_, err1 := imgfile.Seek(cur, 0)
	if err1 != nil {
		fmt.Println(err1)
	}

	ui.Size = size
	ui.FileData = data

	res := i.AddImage(&ui)
	fmt.Print("res: ")
	fmt.Println(res)

	if res.Success != true {
		t.Fail()
	} else {
		imgID = res.ID
	}
}

func TestImageService_GetList(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = imgToken
	res := i.GetList()
	fmt.Print("res: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestImageService_GetList_DeleteImage(t *testing.T) {
	var i ImageService
	i.ClientID = "403"
	i.Host = "http://localhost:3007"
	i.Token = imgToken

	res := i.DeleteImage(strconv.FormatInt(imgID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	addID = res.ID
	if res.Success != true {
		t.Fail()
	}
}
