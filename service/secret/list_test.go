package secret_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_storeclient"
	"github.com/langered/gonedrive/service/secret"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret Push Service", func() {
	Context("#List", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockStoreClient = NewMockStoreClient(mockCtrl)
		})

		It("Lists all names of credentials in the secret file", func() {
			expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

			secretList, err := secret.List(mockStoreClient, "fake-token", "correct-password", "test.gdsecret")
			Expect(err).NotTo(HaveOccurred())
			Expect(secretList).To(ConsistOf("first-secret", "top", "user@secret.com"))
		})

		Context("An error occurs", func() {
			It("Lists all names of credentials in the secret file", func() {
				expectClientGET("", errors.New("Failed to get credentials"), "fake-token", "test.gdsecret")

				secretList, err := secret.List(mockStoreClient, "fake-token", "correct-password", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(secretList).To(BeEmpty())
			})

			It("returns the error when the password is not valid", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

				secretList, err := secret.List(mockStoreClient, "fake-token", "wrong-password", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(secretList).To(BeEmpty())
			})

			It("returns the error when the password file is corrupt", func() {
				expectClientGET(wrongSecretContent, nil, "fake-token", "test.gdsecret")

				secretList, err := secret.List(mockStoreClient, "fake-token", "correct-password", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(secretList).To(BeEmpty())
			})
		})
	})
})
