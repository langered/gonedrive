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

func itemByPath(httpClient httpclient.HttpClient, accessToken string, path string) (item, error) {
	itemByPathURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/%s", path)

	request := getRequest(itemByPathURL, accessToken)
	response, err := httpClient.Do(request)
	if err != nil {
		return item{}, err
	}

	return unmarshallItemResponse(response)
}

func getRequest(url string, accessToken string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	bearerToken := "Bearer " + accessToken
	req.Header.Add("Authorization", bearerToken)
	return req
}

func deleteRequest(url string, accessToken string) *http.Request {
	req, _ := http.NewRequest("DELETE", url, nil)
	bearerToken := "Bearer " + accessToken
	req.Header.Add("Authorization", bearerToken)
	return req
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
