package azure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/langered/gonedrive/httpclient"
)

type listResponse struct {
	Value []item `json:"value"`
}

type item struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//List returns the folders and files in a given path as []string
func (client AzureClient) List(httpClient httpclient.HttpClient, accessToken string, remotePath string) ([]string, error) {
	childrenURL := "https://graph.microsoft.com/v1.0/me/drive/root/children"

	if remotePath != "" {
		parentFolderItem, err := itemByPath(httpClient, accessToken, remotePath)
		if err != nil {
			return []string{}, err
		}
		childrenURL = fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/items/%s/children", parentFolderItem.ID)
	}

	childItems, err := listItemsAsStruct(httpClient, accessToken, childrenURL)
	if err != nil {
		return []string{}, err
	}

	items := []string{}
	for _, item := range childItems.Value {
		items = append(items, item.Name)
	}
	return items, nil
}

func itemByPath(httpClient httpclient.HttpClient, accessToken string, path string) (item, error) {
	itemByPathURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s", path)

	request := getRequest(itemByPathURL, accessToken)
	response, err := httpClient.Do(request)
	if err != nil {
		return item{}, err
	}

	return unmarshallItemResponse(response)
}

func listItemsAsStruct(httpClient httpclient.HttpClient, accessToken string, url string) (listResponse, error) {
	request := getRequest(url, accessToken)
	response, err := httpClient.Do(request)
	if err != nil {
		return listResponse{}, err
	}
	return unmarshallListResponse(response)
}

func getRequest(url string, accessToken string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	bearerToken := "Bearer " + accessToken
	req.Header.Add("Authorization", bearerToken)
	return req
}

func unmarshallListResponse(response *http.Response) (listResponse, error) {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var unmarshalledResponse listResponse

	err := json.Unmarshal(body, &unmarshalledResponse)
	if err != nil {
		return listResponse{}, err
	}
	return unmarshalledResponse, nil
}

func unmarshallItemResponse(response *http.Response) (item, error) {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var unmarshalledResponse item

	err := json.Unmarshal(body, &unmarshalledResponse)
	if err != nil {
		return item{}, err
	}
	return unmarshalledResponse, nil
}
