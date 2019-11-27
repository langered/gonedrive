package service_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	listRootURL    string
)

var _ = Describe("File-Service", func() {
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHttpClient = NewMockHttpClient(mockCtrl)
		listRootURL = "https://graph.microsoft.com/v1.0/me/drive/root/children"
	})

	Context("#ListItems", func() {
		Context("no path given", func() {
			BeforeEach(func() {
				respBodyString := `{
	"value": [
		{ "id": "123", "name": "file1.txt"},
		{ "id": "456", "name": "file2.yml"},
		{ "id": "789", "name": "folder1"}
	]
}`
				prepareHttpClient(listRootURL, respBodyString, 200, nil)
			})

			It("list items on root level", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).NotTo(HaveOccurred())
				Expect(items).To(ConsistOf("file1.txt", "file2.yml", "folder1"))
			})
		})

		Context("path is given", func() {
			var (
				respBodyStringRoot string = `{
	"value": [
		{ "id": "123", "name": "file1.txt"},
		{ "id": "456", "name": "file2.yml"},
		{ "id": "789", "name": "folder1"}
	]
}`

				respBodyStringLevel1 string = `{
	"value": [
		{ "id": "735", "name": "file.txt"},
		{ "id": "890", "name": "folder1_1"},
		{ "id": "734", "name": "folder1_2"}
	]
}`

				respBodyStringLevel2 string = `{
	"value": [
		{ "id": "908", "name": "file_in_dir1.txt"},
		{ "id": "456", "name": "file_in_dir2.yml"},
		{ "id": "789", "name": "folder_in_dir1"}
	]
}`
			)
			BeforeEach(func() {
				prepareHttpClient(listRootURL, respBodyStringRoot, 200, nil)
				prepareHttpClient("https://graph.microsoft.com/v1.0/me/drive/items/789/children", respBodyStringLevel1, 200, nil)
				prepareHttpClient("https://graph.microsoft.com/v1.0/me/drive/items/734/children", respBodyStringLevel2, 200, nil)
			})

			It("list items on the level of the given path", func() {
				items, err := service.ListItems(mockHttpClient, "/folder1/folder1_2", "abc123")

				Expect(err).NotTo(HaveOccurred())
				Expect(items).To(ConsistOf("file_in_dir1.txt", "file_in_dir2.yml", "folder_in_dir1"))
			})

			Context("Some error occurs when somewhere in the path", func() {
				BeforeEach(func() {
					mockHttpClient = NewMockHttpClient(mockCtrl)

					prepareHttpClient(listRootURL, respBodyStringRoot, 200, nil)
					prepareHttpClient("https://graph.microsoft.com/v1.0/me/drive/items/789/children", respBodyStringLevel1, 500, errors.New("no response"))
				})

				It("returns an empty list and the error", func() {
					items, err := service.ListItems(mockHttpClient, "/folder1/folder1_2", "abc123")

					Expect(err).To(HaveOccurred())
					Expect(items).To(Equal([]string{}))
				})
			})

			Context("The given path contains a element which does not exist", func() {
				BeforeEach(func() {
					mockHttpClient = NewMockHttpClient(mockCtrl)

					prepareHttpClient(listRootURL, respBodyStringRoot, 200, nil)
				})

				It("returns an empty list and the error", func() {
					items, err := service.ListItems(mockHttpClient, "/folder2/folder1_2", "abc123")

					Expect(err).To(HaveOccurred())
					Expect(items).To(Equal([]string{}))
				})
			})
		})

		Context("http client returns an error", func() {
			BeforeEach(func() {
				prepareHttpClient(listRootURL, "", 500, errors.New("Failed to sent the request"))
			})

			It("returns empty item list and the error", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})
		})

		Context("Unmarshalling response fails", func() {
			BeforeEach(func() {
				prepareHttpClient(listRootURL, "invalid body", 200, nil)
			})

			It("returns an empty item list and the error", func() {
				items, err := service.ListItems(mockHttpClient, "", "abc123")

				Expect(err).To(HaveOccurred())
				Expect(items).To(Equal([]string{}))
			})
		})
	})
})
