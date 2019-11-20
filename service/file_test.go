package service_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	mockCtrl       *gomock.Controller
	mockHttpClient *MockHttpClient
)

var _ = Describe("File-Service", func() {
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHttpClient = NewMockHttpClient(mockCtrl)
	})

	Context("#ListItems", func() {
		Context("no path given", func() {
			BeforeEach(func() {
				respBodyString := `{
	"value": [
		{ "id": 123, "name": "file1.txt"},
		{ "id": 456, "name": "file2.yml"},
		{ "id": 789, "name": "folder1"}
	]
}`
				prepareHttpClient(respBodyString, 200, nil)
			})

			It("list items on root level", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).NotTo(HaveOccurred())
				Expect(items).To(ConsistOf("file1.txt", "file2.yml", "folder1"))
			})
		})

		Context("http client returns an error", func() {
			BeforeEach(func() {
				prepareHttpClient("", 500, errors.New("Failed to sent the request"))
			})

			It("returns empty item list and the error", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})
		})

		Context("Unmarshalling response fails", func() {
			BeforeEach(func() {
				prepareHttpClient("invalid body", 200, nil)
			})

			It("returns an empty item list and the error", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})
		})
	})
})

func prepareHttpClient(respBodyString string, statusCode int, err error) {
	respBody := ioutil.NopCloser(strings.NewReader(respBodyString))

	req, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me/drive/root/children", nil)
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
