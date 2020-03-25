package secret

import (
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/langered/gonedrive/service/crypto"
)

func getDecryptedRemoteFileContent(storeClient service.StoreClient, accessToken string, password string, credFilePath string) (string, error) {
	secretContent, err := storeClient.Get(http.DefaultClient, accessToken, credFilePath)
	if err != nil {
		return secretContent, err
	}
	decryptedContent, err := crypto.Decrypt(secretContent, password)
	if err != nil {
		return "", err
	}
	return decryptedContent, nil
}
