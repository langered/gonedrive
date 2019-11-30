package service

import (
	"fmt"
	"net/http"
	"path/filepath"

	"io/ioutil"
	"strings"

	"github.com/langered/gonedrive/httpclient"
)

//Upload will upload a string to a file
func Upload(httpClient httpclient.HttpClient, accessToken string, remoteFilePath string, content string) (bool, error) {
	uploadFileURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s:/content", filepath.Base(remoteFilePath))

	if filepath.Dir(remoteFilePath) != "." {
		parentFolderItem, err := itemByPath(httpClient, accessToken, filepath.ToSlash(filepath.Dir(remoteFilePath)))
		if err != nil {
			return false, err
		}
		uploadFileURL = fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/items/%s:/%s:/content", parentFolderItem.ID, filepath.Base(remoteFilePath))
	}

	request := putRequest(uploadFileURL, accessToken, content)

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
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
