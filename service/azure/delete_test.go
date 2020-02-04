package azure_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service/azure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File-Service", func() {

	Context("#Delete", func() {
		var (
			respBodyParentItem string = `{
	"id": "1234",
	"name": "file-to-delete.yml"
}`
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockHttpClient = NewMockHttpClient(mockCtrl)
			client = azure.AzureClient{}
		})

		It("deletes the file by a path", func() {
			expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/file-to-delete.yml", "fake-token", respBodyParentItem, 200, nil)
			expectDELETERequest("https://graph.microsoft.com/v1.0/me/drive/items/1234", "fake-token", "Success", 200, nil)

			err := client.Delete(mockHttpClient, "fake-token", "folder_1/folder_2/file-to-delete.yml")
			Expect(err).NotTo(HaveOccurred())
		})

		It("runs successful when returning 404", func() {
			expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/file-to-delete.yml", "fake-token", respBodyParentItem, 200, nil)
			expectDELETERequest("https://graph.microsoft.com/v1.0/me/drive/items/1234", "fake-token", "not found", 404, nil)

			err := client.Delete(mockHttpClient, "fake-token", "folder_1/folder_2/file-to-delete.yml")
			Expect(err).NotTo(HaveOccurred())
		})

		Context("An error occurs", func() {
			It("returns the error when finding the item failed", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/file-to-delete.yml", "fake-token", "", 404, errors.New("Failed to find the item"))

				err := client.Delete(mockHttpClient, "fake-token", "folder_1/folder_2/file-to-delete.yml")
				Expect(err).To(HaveOccurred())
			})

			It("returns the error when deleting the item failed", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/file-to-delete.yml", "fake-token", respBodyParentItem, 200, nil)
				expectDELETERequest("https://graph.microsoft.com/v1.0/me/drive/items/1234", "fake-token", "fail", 500, errors.New("Failed to delete the item"))

				err := client.Delete(mockHttpClient, "fake-token", "folder_1/folder_2/file-to-delete.yml")
				Expect(err).To(HaveOccurred())
			})

			It("returns the error when deleting returned an unsuccessful code", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/file-to-delete.yml", "fake-token", respBodyParentItem, 200, nil)
				expectDELETERequest("https://graph.microsoft.com/v1.0/me/drive/items/1234", "fake-token", "fail", 500, nil)

				err := client.Delete(mockHttpClient, "fake-token", "folder_1/folder_2/file-to-delete.yml")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
