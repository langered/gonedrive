package service_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File-Service", func() {

	Context("#Upload", func() {
		var (
			respBodyParentItem string = `{
	"id": "1234",
	"name": "folder1_2"
}`
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockHttpClient = NewMockHttpClient(mockCtrl)
		})

		It("uploads a local file under the given path", func() {
			expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", respBodyParentItem, 201, nil)
			expectPUTRequest("https://graph.microsoft.com/v1.0/me/drive/items/1234:/test.txt:/content", "abc123", "successful upload", "test-content", 201, nil)

			success, err := service.Upload(mockHttpClient, "abc123", "folder1/folder1_2/test.txt", "test-content")

			Expect(err).NotTo(HaveOccurred())
			Expect(success).To(BeTrue())
		})

		It("uploads a local file under the root level", func() {
			expectPUTRequest("https://graph.microsoft.com/v1.0/me/drive/root:/test.txt:/content", "abc123", "successful upload", "test-content", 200, nil)

			success, err := service.Upload(mockHttpClient, "abc123", "test.txt", "test-content")

			Expect(err).NotTo(HaveOccurred())
			Expect(success).To(BeTrue())
		})

		Context("An error occurs", func() {
			BeforeEach(func() {
				mockHttpClient = NewMockHttpClient(mockCtrl)
			})

			It("returns false and the error of the http-client when looking for the parent folder", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", "fail to load parent", 500, errors.New("Failed to load the given path"))
				success, err := service.Upload(mockHttpClient, "abc123", "folder1/folder1_2/test.txt", "test-content")

				Expect(err).To(HaveOccurred())
				Expect(success).To(BeFalse())
			})

			It("returns false and the error of the http-client when uploading the content", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", respBodyParentItem, 201, nil)
				expectPUTRequest("https://graph.microsoft.com/v1.0/me/drive/items/1234:/test.txt:/content", "abc123", "failed to upload", "test-content", 500, errors.New("Endpoint not available"))
				success, err := service.Upload(mockHttpClient, "abc123", "folder1/folder1_2/test.txt", "test-content")

				Expect(err).To(HaveOccurred())
				Expect(success).To(BeFalse())
			})

			It("returns false and an error when uploading the content returned another status code than a successful one", func() {
				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder1/folder1_2", "abc123", respBodyParentItem, 201, nil)
				expectPUTRequest("https://graph.microsoft.com/v1.0/me/drive/items/1234:/test.txt:/content", "abc123", "failed to upload", "test-content", 500, nil)
				success, err := service.Upload(mockHttpClient, "abc123", "folder1/folder1_2/test.txt", "test-content")

				Expect(err).To(HaveOccurred())
				Expect(success).To(BeFalse())
			})
		})
	})
})
