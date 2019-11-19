package service

import (
	"errors"
	"fmt"
	"os"
)

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

func ValidateToken(token string) error {
	if len(token) < 1 {
		return errors.New("")
	}
	return nil
}

func SaveToken(path string, token string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	// defer file.Close()
	return nil
}
