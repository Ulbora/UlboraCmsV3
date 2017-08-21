package main

import (
	"fmt"
	"net/http"
)

// user handlers-----------------------------------------------------
func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.TransferEncoding)
	fmt.Println(req.Host)
	templates.ExecuteTemplate(res, "index.html", nil)
}

//end user handlers------------------------------------------------
