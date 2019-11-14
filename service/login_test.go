package service_test

import (
	"github.com/langered/gonedrive/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login-Service", func() {
	Context("#Login", func() {
		It("returns the authorization uri", func() {
			auth_uri := service.Login()
			Expect(auth_uri).To(Equal("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=94126bd2-3582-4928-adb7-bf307c7d5135&scope=onedrive.readwrite&response_type=code&redirect_uri=https%3A%2F%2Flangered.github.io%2Fgonedrive%2Fdoc%2Foauthcallbackhandler.html"))
		})
	})
})
