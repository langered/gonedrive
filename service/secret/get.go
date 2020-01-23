package secret

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/langered/gonedrive/service/crypto"
	"gopkg.in/yaml.v2"
)

//Get returns the value of a given credential name in a given credential file
func Get(storeClient service.StoreClient, accessToken string, password string, secretName string, credFilePath string) (string, error) {
	encryptedSecrets, err := storeClient.Get(http.DefaultClient, accessToken, credFilePath)
	if err != nil {
		return "", err
	}
	decryptedSecrets, err := crypto.Decrypt(encryptedSecrets, password)
	if err != nil {
		return "", err
	}
	secret, err := getSecret(decryptedSecrets, secretName)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func getSecret(secrets string, secretName string) (string, error) {
	var gdsecrets GDSecret
	err := yaml.Unmarshal([]byte(secrets), &gdsecrets)
	if err != nil {
		return "", err
	}
	for _, secret := range gdsecrets.Secrets {
		if secret.Name == secretName {
			return secret.Value, nil
		}
	}
	return "", fmt.Errorf("The credential with the name '%s' could not be found.", secretName)
}
