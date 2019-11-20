package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/langered/gonedrive/httpclient"
)

type ListResponse struct {
	Value []Item `json:value`
}

type Item struct {
	Name string `json:name`
}

func ListItems(httpClient httpclient.HttpClient, path string, accessToken string) ([]string, error) {
	request := getRequest("https://graph.microsoft.com/v1.0/me/drive/root/children", accessToken)
	response, err := httpClient.Do(request)
	if err != nil {
		return []string{}, err
	}
	listResponse, err := unmarshalResponse(response)
	if err != nil {
		return []string{}, err
	}

	items := []string{}
	for _, item := range listResponse.Value {
		items = append(items, item.Name)
	}
	return items, nil
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
		fmt.Println(err)
		return ListResponse{}, err
	}
	return listResponse, nil
}
