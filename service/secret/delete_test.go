package secret_test

import (
	"errors"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_storeclient"
	"github.com/langered/gonedrive/service/crypto"
	"github.com/langered/gonedrive/service/secret"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret Push Service", func() {
	Context("#Delete", func() {

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockStoreClient = NewMockStoreClient(mockCtrl)
		})

		It("deletes a credential", func() {
			newSecrets := `secrets:
- name: top
  value: fake-secret123
- name: user@secret.com
  value: U-Cannot-hack-Me
`
			expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")
			expectCorrectSecretUpload(newSecrets, "correct-password", "fake-token", "mysecret.gdsecret")

			err := secret.Delete(mockStoreClient, "fake-token", "correct-password", "first-secret", "mysecret.gdsecret")
			Expect(err).ToNot(HaveOccurred())
		})

		Context("An error occures", func() {
			It("returns the error when getting the existing credential file fails", func() {
				expectClientGET("", errors.New("Failed to get the file"), "fake-token", "mysecret.gdsecret")

				err := secret.Delete(mockStoreClient, "fake-token", "correct-password", "first-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})

			It("returns the error when the password is not valid", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")

				err := secret.Delete(mockStoreClient, "fake-token", "wrong-password", "first-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
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

				expectClientGET(wrongSecretContent, nil, "fake-token", "mysecret.gdsecret")

				err := secret.Delete(mockStoreClient, "fake-token", "correct-password", "test", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})

			It("returns the error when the upload fails", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")
				mockStoreClient.
					EXPECT().
					Upload(http.DefaultClient, "fake-token", "mysecret.gdsecret", gomock.Any()).
					Return(false, errors.New("Failed to upload the new secret"))

				err := secret.Delete(mockStoreClient, "fake-token", "correct-password", "first-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if the credential by the given name does not exist", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "test.gdsecret")

				err := secret.Delete(mockStoreClient, "fake-token", "correct-password", "not-existing-secret", "test.gdsecret")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
