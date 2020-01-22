package azure_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service/azure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	listRootURL string
)

var _ = Describe("File-Service", func() {
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHttpClient = NewMockHttpClient(mockCtrl)
		client = azure.AzureClient{}
		listRootURL = "https://graph.microsoft.com/v1.0/me/drive/root/children"
	})

	Context("#ListItems", func() {
		var (
			respBodyParentItem string = `{
	"id": "1234",
	"name": "folder1_2"
}`
		)
		Context("no path given", func() {
			BeforeEach(func() {
				respBodyString := `{
	"value": [
		{ "id": "123", "name": "file1.txt"},
		{ "id": "456", "name": "file2.yml"},
		{ "id": "789", "name": "folder1"}
	]
}`
				expectGETRequest(listRootURL, "abc123", respBodyString, 200, nil)
			})

			It("list items on root level", func() {
				items, err := client.List(mockHttpClient, "abc123", "")

				Expect(err).NotTo(HaveOccurred())
				Expect(items).To(ConsistOf("file1.txt", "file2.yml", "folder1"))
			})
		})

		Context("path is given", func() {
			var (
				respBodyChildren string = `{
	"value": [
		{ "id": "908", "name": "file_in_dir1.txt"},
		{ "id": "456", "name": "file_in_dir2.yml"},
		{ "id": "789", "name": "folder_in_dir1"}
	]
}`
			)
			BeforeEach(func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", respBodyParentItem, 200, nil)
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/items/1234/children", "abc123", respBodyChildren, 200, nil)
			})

			It("list items on the level of the given path", func() {
				items, err := client.List(mockHttpClient, "abc123", "folder1/folder1_2")

				Expect(err).NotTo(HaveOccurred())
				Expect(items).To(ConsistOf("file_in_dir1.txt", "file_in_dir2.yml", "folder_in_dir1"))
			})
		})

		Context("The http-client returns an error", func() {
			It("returns empty item list and the error when getting the parent-folder", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", "error", 500, errors.New("Failed to get dir"))

				items, err := client.List(mockHttpClient, "abc123", "folder1/folder1_2")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})

			It("returns empty item list and the error when getting the children", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", respBodyParentItem, 200, nil)
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/items/1234/children", "abc123", "error", 500, errors.New("Could not get childs"))

				items, err := client.List(mockHttpClient, "abc123", "folder1/folder1_2")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})

			It("returns an empty item list and the error when unmarshalling a list response", func() {
				expectGETRequest(listRootURL, "abc123", "invalid body", 200, nil)

				items, err := client.List(mockHttpClient, "abc123", "")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})

			It("returns an empty item list and the error when unmarshalling a item response", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", "invalid url", 200, nil)

				items, err := client.List(mockHttpClient, "abc123", "folder1/folder1_2")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})
		})
	})
})
