package crypto_test

import (
	"github.com/langered/gonedrive/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File-Service", func() {

	Context("#Encrypt", func() {
		It("Encrypts my text", func() {
			encryptedText, err := crypto.Encrypt("my top secret text", "myPassword")

			Expect(err).NotTo(HaveOccurred())
			Expect(encryptedText).To(HaveLen(46))
		})
	})

	Context("#Decrypt", func() {
		It("Decrypts my encrypted text", func() {
			encryptedText, err := crypto.Encrypt("my top secret text", "myPassword")
			Expect(err).NotTo(HaveOccurred())
			plainText, err := crypto.Decrypt(encryptedText, "myPassword")

			Expect(err).NotTo(HaveOccurred())
			Expect(plainText).To(Equal("my top secret text"))
		})

		It("returns an error and an empty text when password is incorrect", func() {
			encryptedText, err := crypto.Encrypt("my top secret text", "myPassword")
			Expect(err).NotTo(HaveOccurred())
			plainText, err := crypto.Decrypt(encryptedText, "myWrongPassword")

			Expect(err).To(HaveOccurred())
			Expect(plainText).To(Equal(""))
		})
	})
})
