package secret_test

import (
	"errors"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/langered/gonedrive/crypto"
	. "github.com/langered/gonedrive/fakes/mock_storeclient"
	"github.com/langered/gonedrive/service/secret"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret Push Service", func() {
	Context("#Push", func() {

		var (
			existingSecrets = `
---
secrets:
- name: first-secret
  value: "totally-unknown"
- name: top
  value: "fake-secret123"
- name: user@secret.com
  value: "U-Cannot-hack-Me"
`

			existingSecretContent, _ = crypto.Encrypt(existingSecrets, "correct-password")
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockStoreClient = NewMockStoreClient(mockCtrl)
		})

		It("uploads a new credential", func() {
			newSecrets := `secrets:
- name: first-secret
  value: totally-unknown
- name: top
  value: fake-secret123
- name: user@secret.com
  value: U-Cannot-hack-Me
- name: new-secret
  value: my-new-secret
`
			expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")
			expectCorrectSecretUpload(newSecrets, "correct-password", "fake-token", "mysecret.gdsecret")

			err := secret.Push(mockStoreClient, "fake-token", "correct-password", "new-secret", "my-new-secret", "mysecret.gdsecret")
			Expect(err).ToNot(HaveOccurred())
		})

		It("changes an existing credential", func() {
			changedSecrets := `secrets:
- name: first-secret
  value: changed-value
- name: top
  value: fake-secret123
- name: user@secret.com
  value: U-Cannot-hack-Me
`
			expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")
			expectCorrectSecretUpload(changedSecrets, "correct-password", "fake-token", "mysecret.gdsecret")

			err := secret.Push(mockStoreClient, "fake-token", "correct-password", "first-secret", "changed-value", "mysecret.gdsecret")
			Expect(err).ToNot(HaveOccurred())
		})

		Context("An error occures", func() {
			It("returns the error when getting the existing credential file fails", func() {
				expectClientGET("", errors.New("Failed to get the file"), "fake-token", "mysecret.gdsecret")

				err := secret.Push(mockStoreClient, "fake-token", "correct-password", "test", "test-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})

			It("creates a new file when error for getting the credential file is 404", func() {
				secretsFromScratch := `secrets:
- name: test
  value: test-secret
`
				expectClientGET("404", errors.New("Failed to get the file"), "fake-token", "mysecret.gdsecret")
				expectCorrectSecretUpload(secretsFromScratch, "correct-password", "fake-token", "mysecret.gdsecret")

				err := secret.Push(mockStoreClient, "fake-token", "correct-password", "test", "test-secret", "mysecret.gdsecret")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns the error when the password is not valid", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")

				err := secret.Push(mockStoreClient, "fake-token", "wrong-password", "test", "test-secret", "mysecret.gdsecret")
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

				err := secret.Push(mockStoreClient, "fake-token", "correct-password", "test", "test-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})

			It("returns the error when the upload fails", func() {
				expectClientGET(existingSecretContent, nil, "fake-token", "mysecret.gdsecret")
				mockStoreClient.
					EXPECT().
					Upload(http.DefaultClient, "fake-token", "mysecret.gdsecret", gomock.Any()).
					Return(false, errors.New("Failed to upload the new secret"))

				err := secret.Push(mockStoreClient, "fake-token", "correct-password", "test", "test-secret", "mysecret.gdsecret")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
