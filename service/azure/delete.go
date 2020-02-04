package azure

import (
	"fmt"

	"github.com/langered/gonedrive/httpclient"
)

const deleteEndpoint string = "https://graph.microsoft.com/v1.0/me/drive/items/%s"

//Delete deletes the file by a given path
func (client AzureClient) Delete(httpClient httpclient.HttpClient, accessToken string, remotePath string) error {
	item, err := itemByPath(httpClient, accessToken, remotePath)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(deleteEndpoint, item.ID)
	request := deleteRequest(url, accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	if (response.StatusCode < 200 || response.StatusCode > 299) && response.StatusCode != 404 {
		return fmt.Errorf("Deleting the file failed. It returned the status code: %v", response.StatusCode)
	}
	return nil
}
