package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/langered/gonedrive/httpclient"
)

type ListResponse struct {
	Value []Item `json:value`
}

type Item struct {
	Name string `json:name`
	ID   string `json:id`
}

var (
	listRootURL string = "https://graph.microsoft.com/v1.0/me/drive/root/children"
	listURL     string = "https://graph.microsoft.com/v1.0/me/drive/items/%s/children"
)

func ListItems(httpClient httpclient.HttpClient, path string, accessToken string) ([]string, error) {
	pathItems := strings.Split(path, "/")
	listResponse, err := listItemsAsStruct(httpClient, listRootURL, accessToken)
	if err != nil {
		return []string{}, err
	}
	for _, pathItem := range pathItems {
		if pathItem != "" {
			id := getIDByName(listResponse, pathItem)
			nextURL := fmt.Sprintf(listURL, id)
			listResponse, err = listItemsAsStruct(httpClient, nextURL, accessToken)
			if err != nil {
				return []string{}, err
			}
		}
	}

	items := []string{}
	for _, item := range listResponse.Value {
		items = append(items, item.Name)
	}
	return items, nil
}

func getIDByName(response ListResponse, name string) string {
	for _, item := range response.Value {
		if item.Name == name {
			return item.ID
		}
	}
	return ""
}

func listItemsAsStruct(httpClient httpclient.HttpClient, url string, accessToken string) (ListResponse, error) {
	request := getRequest(url, accessToken)
	response, err := httpClient.Do(request)
	if err != nil {
		return ListResponse{}, err
	}
	return unmarshalResponse(response)
}

func getRequest(url string, accessToken string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	bearerToken := "Bearer " + accessToken
	req.Header.Add("Authorization", bearerToken)
	return req
}

func unmarshalResponse(response *http.Response) (ListResponse, error) {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var listResponse ListResponse

	err := json.Unmarshal(body, &listResponse)
	if err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}
