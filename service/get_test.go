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

	Context("#Get", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockHttpClient = NewMockHttpClient(mockCtrl)

			expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/folder_1/folder_2/test.txt:/content", "abc123", "test-content", 200, nil)
		})

		It("gets the file by a path and returns the content", func() {
			content, err := service.Get(mockHttpClient, "abc123", "folder_1/folder_2/test.txt")

			Expect(err).NotTo(HaveOccurred())
			Expect(content).To(Equal("test-content"))
		})

		Context("An error occurs", func() {
			BeforeEach(func() {
				mockHttpClient = NewMockHttpClient(mockCtrl)

				expectGETRequest("https://graph.microsoft.com/v1.0/me/drive/root:/test.txt:/content", "abc123", "fail", 500, errors.New("http-error"))
			})
			It("returns the error of the http-client", func() {
				content, err := service.Get(mockHttpClient, "abc123", "test.txt")

				Expect(err).To(HaveOccurred())
				Expect(content).To(Equal(""))
			})
		})
	})
})
