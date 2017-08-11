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

	//refresh token
	oauthAuthCodeRefreshTokenURI1 = "/oauth/token?grant_type=refresh_token&client_id="
	oauthAuthCodeRefreshTokenURI2 = "&client_secret="
	oauthAuthCodeRefreshTokenURI3 = "&refresh_token="
)
