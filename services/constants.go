package services

const (
	localOauthHost = "http://localhost:3000"
	localClientID  = "403"
	//auth codr authorize
	oauthAuthCodeAuthorizeURI1 = "/oauth/authorize?response_type=code&client_id="
	oauthAuthCodeAuthorizeURI2 = "&redirect_uri="
	localRedirectURI           = "http://localhost:8090/token"
	oauthAuthCodeAuthorizeURI3 = "&scope=write&state="
	state                      = "66ggh"

	//auth code token
	oauthAuthCodeTokenURI1 = "/oauth/token?client_id="
	oauthAuthCodeTokenURI2 = "&client_secret="
	oauthAuthCodeTokenURI3 = "&grant_type=authorization_code&code="
	oauthAuthCodeTokenURI4 = "&redirect_uri="
	url                    = "http://localhost:3000/oauth/token?client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&grant_type=authorization_code&code=i76y13e340akRn6Ipkdbii&redirect_uri=http://www.google.com"
)
