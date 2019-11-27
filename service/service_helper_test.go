package service_test

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
)

var (
	mockCtrl       *gomock.Controller
	mockHttpClient *MockHttpClient
)

func prepareHttpClient(expectedURL string, respBodyString string, statusCode int, err error) {
	respBody := ioutil.NopCloser(strings.NewReader(respBodyString))

	req, _ := http.NewRequest("GET", expectedURL, nil)
	bearerToken := "Bearer abc123"
	req.Header.Add("Authorization", bearerToken)

	mockHttpClient.
		EXPECT().
		Do(req).
		Return(&http.Response{
			StatusCode: statusCode,
			Body:       respBody,
		}, err)
}
