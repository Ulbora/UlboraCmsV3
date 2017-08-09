package services

const (
	localOauthHost                     = "http://localhost:3000"
	oauthAuthCodeAuthorizeURI1         = "/oauth/authorize?response_type=code&client_id="
	localClientID                      = "403"
	oauthAuthCodeAuthorizeURI2         = "&redirect_uri="
	localRedirectURI                   = "http://localhost:8090/token"
	oauthAuthCodeAuthorizeURI3         = "&scope=write&state="
	state                              = "66ggh"
	oauth2AuthCodeAuthorizeURIComplete = localOauthHost + oauthAuthCodeAuthorizeURI1 + localClientID + oauthAuthCodeAuthorizeURI2 + localRedirectURI + oauthAuthCodeAuthorizeURI3 + state
)
