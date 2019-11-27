package service

import (
	"errors"
	"fmt"
)

//AuthURL returns the URL for the authentication in microsoft graph
func AuthURL() string {
	authorizeURI := "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	clientID := "94126bd2-3582-4928-adb7-bf307c7d5135"
	scope := "files.readwrite"
	responseType := "token"
	redirectURI := "https%3A%2F%2Flangered.github.io%2Fgonedrive%2Fdoc%2Foauthcallbackhandler.html"

	return fmt.Sprintf("%s?client_id=%s&scope=%s&response_type=%s&redirect_uri=%s",
		authorizeURI,
		clientID,
		scope,
		responseType,
		redirectURI)
}

//ValidateToken is a function to check if the given token is correct
func ValidateToken(token string) error {
	if len(token) < 1 {
		return errors.New("")
	}
	return nil
}
