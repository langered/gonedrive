package secret

import (
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/langered/gonedrive/service/crypto"
	"gopkg.in/yaml.v2"
)

//List returns a list of secret names
func List(storeClient service.StoreClient, accessToken string, password string, credFilePath string) ([]string, error) {
	encryptedSecrets, err := storeClient.Get(http.DefaultClient, accessToken, credFilePath)
	if err != nil {
		return []string{}, err
	}
	decryptedSecrets, err := crypto.Decrypt(encryptedSecrets, password)
	if err != nil {
		return []string{}, err
	}
	return secretNames(decryptedSecrets)
}

func secretNames(secrets string) ([]string, error) {
	var gdsecrets GDSecret
	err := yaml.Unmarshal([]byte(secrets), &gdsecrets)
	if err != nil {
		return []string{}, err
	}

	secretList := make([]string, 0)
	for _, secret := range gdsecrets.Secrets {
		secretList = append(secretList, secret.Name)
	}
	return secretList, nil
}
