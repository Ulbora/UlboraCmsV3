package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func getAuthCodeClient() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_ID") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_ID")
	} else {
		rtn = authCodeClient
	}
	return rtn
}

func getAuthCodeSecret() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_SECRET") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_SECRET")
	} else {
		rtn = authCodeSecret
	}
	return rtn
}
func getOauthHost() string {
	var rtn = ""
	if os.Getenv("AUTH_HOST") != "" {
		rtn = os.Getenv("AUTH_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}
func getRedirectURI(req *http.Request, path string) string {
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String()
	} else {
		serverHost = schemeDefault + req.Host + path
	}
	return serverHost
}

func getContentHost() string {
	var rtn = ""
	if os.Getenv("CONTENT_HOST") != "" {
		rtn = os.Getenv("CONTENT_HOST")
	} else {
		rtn = "http://localhost:3008"
	}
	return rtn
}

func getMailHost() string {
	var rtn = ""
	if os.Getenv("MAIL_HOST") != "" {
		rtn = os.Getenv("MAIL_HOST")
	} else {
		rtn = "http://localhost:3002"
	}
	return rtn
}

func getImageHost() string {
	var rtn = ""
	if os.Getenv("IMAGE_HOST") != "" {
		rtn = os.Getenv("IMAGE_HOST")
	} else {
		rtn = "http://localhost:3007"
	}
	return rtn
}

func getHashedUser() string {
	var rtn string
	//fmt.Println(token.AccessToken)
	tk, err := jwt.Parse(token.AccessToken, func(parsedToken *jwt.Token) (interface{}, error) {
		return parsedToken, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if tk != nil {
		if claims, ok := tk.Claims.(jwt.MapClaims); ok {
			uid := claims["userId"]
			//fmt.Println(uid)
			if uid != nil {
				rtn = uid.(string)
			}
		}
	} else {
		rtn = ""
	}
	//fmt.Println(rtn)
	return rtn
}
