package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

//Encrypt Encrypt a string with a given password
func Encrypt(text string, password string) (string, error) {
	hashedKey := create32ByteHash(password)
	data := []byte(text)
	block, _ := aes.NewCipher([]byte(hashedKey))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return string(ciphertext), nil
}

//Decrypt Decrypt a string with a given password
func Decrypt(encryptedText string, password string) (string, error) {
	hashedKey := create32ByteHash(password)
	data := []byte(encryptedText)
	key := []byte(hashedKey)
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New("Password is invalid")
	}
	return string(plaintext), nil
}

func create32ByteHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
