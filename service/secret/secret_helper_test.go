package secret_test

import (
	"errors"
	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_storeclient"
	"github.com/langered/gonedrive/httpclient"
	"github.com/langered/gonedrive/service/crypto"
)

var (
	mockCtrl        *gomock.Controller
	mockStoreClient *MockStoreClient

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

	wrongSecretFileContent = `
---
secrets:
- name: identation
	value: "is wrong"
`
	wrongSecretContent, _ = crypto.Encrypt(wrongSecretFileContent, "correct-password")
)

func expectCorrectSecretUpload(newSecrets string, decryptPassword string, accessToken string, secretFilePath string) {
	mockStoreClient.
		EXPECT().
		Upload(http.DefaultClient, accessToken, secretFilePath, gomock.Any()).
		DoAndReturn(func(client httpclient.HttpClient, token string, path string, content string) (bool, error) {
			decryptedContent, _ := crypto.Decrypt(content, decryptPassword)
			if decryptedContent != newSecrets {
				return false, errors.New("The uploaded secrets are different")
			}
			return true, nil
		})
}

func expectClientGET(expectedContent string, expectedError error, accessToken string, secretFilePath string) {
	mockStoreClient.
		EXPECT().
		Get(http.DefaultClient, accessToken, secretFilePath).
		Return(expectedContent, expectedError)
}
