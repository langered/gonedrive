package azure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"net"
	"net/http"
	"net/url"

	"github.com/langered/gonedrive/browser"
	"github.com/langered/gonedrive/httpclient"
)

var (
	authCodeValue string
	server        *http.Server
)

const (
	loginServerPort    = ":8261"
	authCodeRoute      = "/authcode"
	codeURLParamName   = "code"
	encodedRedirectURI = "http%3A%2F%2Flocalhost%3A8261%2Fauthcode"
	scope              = "files.readwrite"
	clientID           = "94126bd2-3582-4928-adb7-bf307c7d5135"
)

type authtokenResponse struct {
	AccessToken string `json:"access_token"`
}

//Login will open the login page for microsoft and parse the access token
func (client AzureClient) Login(httpClient httpclient.HttpClient, browser browser.Browser) (string, error) {
	authCode(browser)
	if authCodeValue == "" {
		return "", errors.New("No auth code received")
	}
	token, err := accessToken(httpClient)
	return token, err
}

var wg sync.WaitGroup

func authCode(browser browser.Browser) {
	server = &http.Server{Addr: loginServerPort}
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc(authCodeRoute, parseAuthCode)

	listener, _ := net.Listen("tcp", "localhost"+loginServerPort)
	wg.Add(1)
	go server.Serve(listener)
	browser.OpenURL(authCodeURL())
	wg.Wait()
}

func parseAuthCode(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()[codeURLParamName]
	if !ok || len(keys[0]) < 1 {
		authCodeValue = ""
		server.Close()
		wg.Done()
		return
	}
	authCodeValue = keys[0]
	server.Close()
	wg.Done()
}

func authCodeURL() string {
	authorizeURI := "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	responseType := "code"

	return fmt.Sprintf("%s?client_id=%s&scope=%s&response_type=%s&redirect_uri=%s",
		authorizeURI,
		clientID,
		scope,
		responseType,
		encodedRedirectURI)
}

func accessToken(httpClient httpclient.HttpClient) (string, error) {
	req, _ := accessTokenRequest()
	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var unmarshalledResponse authtokenResponse

	err = json.Unmarshal(body, &unmarshalledResponse)
	if err != nil {
		return "", err
	}
	return unmarshalledResponse.AccessToken, nil
}

func accessTokenRequest() (*http.Request, error) {
	authorizeURL := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Add("scope", scope)
	data.Add("code", authCodeValue)
	data.Add("grant_type", "authorization_code")
	data.Add("redirect_uri", "http://localhost:8261/authcode")

	return http.NewRequest("POST", authorizeURL, bytes.NewBufferString(data.Encode()))
}
