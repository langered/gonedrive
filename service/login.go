package service

import "fmt"

func Login() string {
	authorizeURI := "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	clientID := "94126bd2-3582-4928-adb7-bf307c7d5135"
	scope := "onedrive.readwrite"
	responseType := "code"
	redirectURI := "https%3A%2F%2Flangered.github.io%2Fgonedrive%2Fdoc%2Foauthcallbackhandler.html"

	return fmt.Sprintf("%s?client_id=%s&scope=%s&response_type=%s&redirect_uri=%s",
		authorizeURI,
		clientID,
		scope,
		responseType,
		redirectURI)
}
