package main

import (
	"fmt"
	"net/http"
	"os"

	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/dgrijalva/jwt-go"
)

func getTemplate() string {
	var rtn = ""
	if os.Getenv("TEMPLATE_LOC") != "" {
		rtn = os.Getenv("TEMPLATE_LOC")
	} else {
		rtn = "default"
	}
	return rtn
}

func getAuthCodeClient() string {
	var rtn = ""
	if os.Getenv("AUTH_CODE_CLIENT_ID") != "" {
		rtn = os.Getenv("AUTH_CODE_CLIENT_ID")
	} else {
		rtn = authCodeClient
	}
	return rtn
}

func getGatewayAPIKey() string {
	var rtn = ""
	if os.Getenv("GATEWAY_API_KEY") != "" {
		rtn = os.Getenv("GATEWAY_API_KEY")
	} else {
		rtn = apiGatewayKey
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
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("CONTENT_HOST") != "" {
		rtn = os.Getenv("CONTENT_HOST")
	} else {
		rtn = "http://localhost:3011/content"
	}
	return rtn
}

func getMailHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("MAIL_HOST") != "" {
		rtn = os.Getenv("MAIL_HOST")
	} else {
		rtn = "http://localhost:3011/mail"
	}
	return rtn
}

func getImageHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("IMAGE_HOST") != "" {
		rtn = os.Getenv("IMAGE_HOST")
	} else {
		rtn = "http://localhost:3011/image"
	}
	return rtn
}

func getTemplateHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("TEMPLATE_HOST") != "" {
		rtn = os.Getenv("TEMPLATE_HOST")
	} else {
		rtn = "http://localhost:3011/template"
	}
	return rtn
}

func getChallengeHost() string {
	var rtn = ""
	if os.Getenv("API_GATEWAY_HOST") != "" {
		rtn = os.Getenv("API_GATEWAY_HOST")
	} else if os.Getenv("CHALLENGE_HOST") != "" {
		rtn = os.Getenv("CHALLENGE_HOST")
	} else {
		rtn = "http://localhost:3011/challenge"
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
		eout := err.Error()
		if eout != "key is of invalid type" {
			fmt.Print("jwt error: ")
			fmt.Println(eout)
		}
	}
	if tk != nil {
		if claims, ok := tk.Claims.(jwt.MapClaims); ok {
			uid := claims["userId"]
			//fmt.Print("user: ")
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

func getRefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting refresh token")
	var tn oauth2.AuthCodeToken
	tn.OauthHost = getOauthHost()
	tn.ClientID = getAuthCodeClient()
	tn.Secret = getAuthCodeSecret()
	tn.RefreshToken = token.RefreshToken
	fmt.Print("refresh token being sent: ")
	fmt.Println(tn.RefreshToken)
	resp := tn.AuthCodeRefreshToken()
	fmt.Print("refresh token response: ")
	fmt.Println(resp)
	if resp != nil && resp.AccessToken != "" {
		fmt.Print("new token: ")
		fmt.Println(resp.AccessToken)
		token = resp
		session, err := s.GetSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			session.Values["userLoggenIn"] = true
			session.Save(r, w)
			//http.Redirect(res, req, "/admin/main", http.StatusFound)

			// decode token and get user id
		}
	}

}

func getCredentialsToken() {
	fmt.Println("getting Client Credentials token")
	var tn oauth2.ClientCredentialsToken
	tn.OauthHost = getOauthHost()
	tn.ClientID = getAuthCodeClient()
	tn.Secret = getAuthCodeSecret()
	resp := tn.ClientCredentialsToken()
	//fmt.Print("credentils token response: ")
	//fmt.Println(resp)
	if resp != nil && resp.AccessToken != "" {
		//fmt.Print("new credentials token: ")
		//fmt.Println(resp.AccessToken)
		credentialToken = resp
	}
}
