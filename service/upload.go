package service

import (
	"fmt"
	"net/http"

	"io/ioutil"
	"strings"

	"github.com/langered/gonedrive/httpclient"
)

//Upload will upload a string to a file
func Upload(httpClient httpclient.HttpClient, accessToken string, remotePath string, remoteFilename string, content string) (bool, error) {
	parentFolderItem, err := itemByPath(httpClient, accessToken, remotePath)
	if err != nil {
		return false, err
	}

	uploadFileURLTemplate := "https://graph.microsoft.com/v1.0/me/drive/items/%s:/%s:/content"
	uploadFileURL := fmt.Sprintf(uploadFileURLTemplate, parentFolderItem.ID, remoteFilename)
	request := putRequest(uploadFileURL, accessToken, content)

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 201 {
		return false, fmt.Errorf("Uploading the content was not successful. It returned the status code: %v", response.StatusCode)
	}
	return true, nil
}

func putRequest(url string, accessToken string, data string) *http.Request {
	req, _ := http.NewRequest("PUT", url, ioutil.NopCloser(strings.NewReader(data)))
	bearerToken := "Bearer " + accessToken
	req.Header.Add("Authorization", bearerToken)
	return req
}
