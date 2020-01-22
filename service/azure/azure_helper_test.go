package azure_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service/azure"
)

var (
	client         azure.AzureClient
	mockCtrl       *gomock.Controller
	mockHttpClient *MockHttpClient
)

func expectGETRequest(expectedURL string, expectedAccessToken string, respBodyString string, statusCode int, err error) {
	respBody := ioutil.NopCloser(strings.NewReader(respBodyString))

	req, _ := http.NewRequest("GET", expectedURL, nil)
	bearerToken := fmt.Sprintf("Bearer %s", expectedAccessToken)
	req.Header.Add("Authorization", bearerToken)

	mockHttpClient.
		EXPECT().
		Do(req).
		Return(&http.Response{
			StatusCode: statusCode,
			Body:       respBody,
		}, err)
}

func expectPUTRequest(expectedURL string, expectedAccessToken string, respBodyString string, data string, statusCode int, err error) {
	respBody := ioutil.NopCloser(strings.NewReader(respBodyString))

	req, _ := http.NewRequest("PUT", expectedURL, ioutil.NopCloser(strings.NewReader(data)))
	bearerToken := fmt.Sprintf("Bearer %s", expectedAccessToken)
	req.Header.Add("Authorization", bearerToken)

	mockHttpClient.
		EXPECT().
		Do(req).
		Return(&http.Response{
			StatusCode: statusCode,
			Body:       respBody,
		}, err)
}

func expectPOSTRequest(expectedURL string, expectedPayload url.Values, expectedBody string, expectedStatusCode int, expectedError error) {
	respBody := ioutil.NopCloser(strings.NewReader(expectedBody))
	req, _ := http.NewRequest("POST", expectedURL, bytes.NewBufferString(expectedPayload.Encode()))

	mockHttpClient.
		EXPECT().
		Do(MatchesRequest(req)).
		Return(&http.Response{
			StatusCode: expectedStatusCode,
			Body:       respBody,
		}, expectedError)
}

type wantRequest struct{ request *http.Request }

func MatchesRequest(request *http.Request) gomock.Matcher {
	return &wantRequest{request}
}

func (wr *wantRequest) Matches(x interface{}) bool {
	gotRequest, _ := httputil.DumpRequest(x.(*http.Request), true)
	wantRequest, _ := httputil.DumpRequest(wr.request, true)
	res := bytes.Compare(gotRequest, wantRequest)

	if res == 0 {
		return true
	}
	return false
}

func (wr *wantRequest) String() string {
	requestDump, _ := httputil.DumpRequest(wr.request, true)
	return "following attributes to be set:\n" + string(requestDump)
}
