package secret_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_storeclient"
	"github.com/langered/gonedrive/service/crypto"
	"github.com/langered/gonedrive/service/secret"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret Push Service", func() {
	Context("#Get", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockStoreClient = NewMockStoreClient(mockCtrl)
		})

		It("Gets a credential by name", func() {
			expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

			returnedSecret, err := secret.Get(mockStoreClient, "fake-token", "correct-password", "top", "test.gdsecret")
			Expect(err).ToNot(HaveOccurred())
			Expect(returnedSecret).To(Equal("fake-secret123"))
		})

		Context("An error occurs", func() {
			It("returns the error when getting the file failed", func() {
				expectClientGET("", errors.New("Failed to get the file"), "fake-token", "test.gdsecret")

				returnedSecret, err := secret.Get(mockStoreClient, "fake-token", "correct-password", "top", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(returnedSecret).To(Equal(""))
			})

			It("returns the error when the password is not valid", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

				returnedSecret, err := secret.Get(mockStoreClient, "fake-token", "wrong-password", "top", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(returnedSecret).To(Equal(""))
			})

			It("returns the error when the password file is corrupt", func() {
				var (
					wrongSecretFileContent = `
---
secrets:
- name: identation
	value: "is wrong"
`
					wrongSecretContent, _ = crypto.Encrypt(wrongSecretFileContent, "correct-password")
				)

				expectClientGET(wrongSecretContent, nil, "fake-token", "test.gdsecret")

				returnedSecret, err := secret.Get(mockStoreClient, "fake-token", "correct-password", "top", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(returnedSecret).To(Equal(""))
			})

			It("returns an error if the credential by the given name does not exist", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

				returnedSecret, err := secret.Get(mockStoreClient, "fake-token", "correct-password", "not-existing-secret", "test.gdsecret")
				Expect(err).To(HaveOccurred())
				Expect(returnedSecret).To(Equal(""))
			})
		})
	})
})
