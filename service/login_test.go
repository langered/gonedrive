package service_test

import (
	"github.com/langered/gonedrive/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login-Service", func() {

	Context("#AuthURL", func() {
		It("returns the authentication URL for this APP", func() {
			authURL := service.AuthURL()
			Expect(authURL).To(Equal("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=94126bd2-3582-4928-adb7-bf307c7d5135&scope=files.readwrite&response_type=token&redirect_uri=https%3A%2F%2Flangered.github.io%2Fgonedrive%2Fdoc%2Foauthcallbackhandler.html"))
		})
	})

	Context("#ValidateToken", func() {
		It("returns no error if token is valid", func() {
			err := service.ValidateToken("12345679")
			Expect(err).NotTo(HaveOccurred())
		})

		Context("Token is invalid", func() {
			It("returns error that token is invalid", func() {
				err := service.ValidateToken("")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
