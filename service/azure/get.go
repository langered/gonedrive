package azure

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/langered/gonedrive/httpclient"
)

const contentOfFileEndpoint string = "https://graph.microsoft.com/v1.0/me/drive/root:/%s:/content"

//Get returns the content of a file given by a path
func (client AzureClient) Get(httpClient httpclient.HttpClient, accessToken string, remotePath string) (string, error) {
	url := fmt.Sprintf(contentOfFileEndpoint, remotePath)
	request := getRequest(url, accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return strconv.Itoa(response.StatusCode), err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return strconv.Itoa(response.StatusCode), fmt.Errorf("Getting the file failed. It returned the status code: %v", response.StatusCode)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}
