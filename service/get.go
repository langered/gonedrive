package service

import (
	"fmt"
	"io/ioutil"

	"github.com/langered/gonedrive/httpclient"
)

var (
	contentOfFileURL string = "https://graph.microsoft.com/v1.0/me/drive/root:/%s:/content"
)

//Get returns the content of a file given by a path
func Get(httpClient httpclient.HttpClient, accessToken string, remotePath string) (string, error) {
	url := fmt.Sprintf(contentOfFileURL, remotePath)
	request := getRequest(url, accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}
