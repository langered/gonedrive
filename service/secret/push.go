package secret

import (
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/langered/gonedrive/service/crypto"
	"gopkg.in/yaml.v2"
)

//Push uploads a new credential to a .gdsecret file or returns an occurring error
func Push(storeClient service.StoreClient, accessToken string, password string, secretName string, secretValue string, credFilePath string) error {
	decryptedContent, err := getDecryptedRemoteFileContent(storeClient, accessToken, password, credFilePath)
	if err != nil {
		if decryptedContent != "404" {
			return err
		} else {
			decryptedContent = ""
		}
	}
	newSecretContent, err := addSecret(decryptedContent, secretName, secretValue)
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

func addSecret(secretValues string, secretName string, secretValue string) (string, error) {
	var gdsecrets GDSecret
	err := yaml.Unmarshal([]byte(secretValues), &gdsecrets)
	if err != nil {
		return "", err
	}
	existingSecret := false
	for index, secret := range gdsecrets.Secrets {
		if secret.Name == secretName {
			existingSecret = true
			gdsecrets.Secrets[index].Value = secretValue
		}
	}

	if !existingSecret {
		secretToAdd := Secret{
			Name:  secretName,
			Value: secretValue,
		}
		gdsecrets.Secrets = append(gdsecrets.Secrets, secretToAdd)
	}
	newSecrets, _ := yaml.Marshal(gdsecrets)
	return string(newSecrets), nil
}
