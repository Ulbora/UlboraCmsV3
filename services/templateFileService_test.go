package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestTemplateFileService_ExtractFile(t *testing.T) {
	tf, err := os.Open("./testFiles/upload/testTxt.tar.gz")
	if err != nil {
		fmt.Println("tar file not found!")
		os.Exit(1)
	}
	defer tf.Close()
	var ts TemplateFileService
	ts.OriginalFileName = tf.Name()
	ts.Destination = "./testFiles/downloaded"
	ts.Name = "newTemplate"
	data, err := ioutil.ReadAll(tf)
	if err != nil {
		fmt.Println(err)
	} else {
		ts.FileData = data
	}
	fmt.Print("file name: ")
	fmt.Println(ts.OriginalFileName)
	res := ts.ExtractFile()
	if res != true {
		t.Fail()
	}
}

func TestTemplateFileService_DeleteTemplate(t *testing.T) {
	var ts TemplateFileService
	ts.Destination = "./testFiles/downloaded"
	ts.Name = "newTemplate"
	res := ts.DeleteTemplate()
	if res != true {
		t.Fail()
	}
}
