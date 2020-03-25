package secret

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/langered/gonedrive/service/crypto"
	"gopkg.in/yaml.v2"
)

//Delete deletes a credential in the .gdsecret file
func Delete(storeClient service.StoreClient, accessToken string, password string, secretName string, credFilePath string) error {
	decryptedContent, err := getDecryptedRemoteFileContent(storeClient, accessToken, password, credFilePath)
	if err != nil {
		return err
	}
	newSecretContent, err := deleteSecret(decryptedContent, secretName)
	if err != nil {
		return err
	}
	encryptedNewSecrets, _ := crypto.Encrypt(newSecretContent, password)
	_, err = storeClient.Upload(http.DefaultClient, accessToken, credFilePath, encryptedNewSecrets)
	if err != nil {
		return err
	}
	return nil
}

func deleteSecret(secrets string, secretName string) (string, error) {
	var gdsecrets GDSecret
	err := yaml.Unmarshal([]byte(secrets), &gdsecrets)
	if err != nil {
		return "", err
	}
	existingSecret := false
	for index, secret := range gdsecrets.Secrets {
		if secret.Name == secretName {
			existingSecret = true
			gdsecrets.Secrets = append(gdsecrets.Secrets[:index], gdsecrets.Secrets[index+1:]...)
		}
	}
	if !existingSecret {
		return "", fmt.Errorf("The credential with the name '%s' could not be found.", secretName)
	}
	newSecrets, _ := yaml.Marshal(gdsecrets)
	return string(newSecrets), nil
}
