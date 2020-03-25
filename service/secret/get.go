package secret

import (
	"fmt"

	"github.com/langered/gonedrive/service"
	"gopkg.in/yaml.v2"
)

//Get returns the value of a given credential name in a given credential file
func Get(storeClient service.StoreClient, accessToken string, password string, secretName string, credFilePath string) (string, error) {
	decryptedContent, err := getDecryptedRemoteFileContent(storeClient, accessToken, password, credFilePath)
	if err != nil {
		return "", err
	}
	secret, err := getSecret(decryptedContent, secretName)
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
