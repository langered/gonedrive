package azure_test

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/golang/mock/gomock"
	. "github.com/langered/gonedrive/fakes/mock_browser"
	. "github.com/langered/gonedrive/fakes/mock_httpclient"
	"github.com/langered/gonedrive/service/azure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	mockBrowser     *MockBrowser
	expectedPayload url.Values
	expectedBody    string
)

const (
	tokenURL             = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
	authURL              = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=94126bd2-3582-4928-adb7-bf307c7d5135&scope=files.readwrite&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A8261%2Fauthcode"
	localhostAuthcodeURL = "http://localhost:8261/authcode?code=authcode123"
)

var _ = Describe("Login-Service", func() {

	Context("#Login", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockHttpClient = NewMockHttpClient(mockCtrl)
			mockBrowser = NewMockBrowser(mockCtrl)
			client = azure.AzureClient{}

			expectedPayload = url.Values{}
			expectedPayload.Set("client_id", "94126bd2-3582-4928-adb7-bf307c7d5135")
			expectedPayload.Add("scope", "files.readwrite")
			expectedPayload.Add("code", "authcode123")
			expectedPayload.Add("grant_type", "authorization_code")
			expectedPayload.Add("redirect_uri", "http://localhost:8261/authcode")

			expectedBody = `{
	"access_token": "token123"
}`
		})

		It("returns the auth token", func() {
			expectOpenAuthURL(authURL, localhostAuthcodeURL, nil)

			expectPOSTRequest(tokenURL, expectedPayload, expectedBody, 200, nil)

			authenticationToken, err := client.Login(mockHttpClient, mockBrowser)

			Expect(err).NotTo(HaveOccurred())
			Expect(authenticationToken).To(Equal("token123"))
		})

		Context("an error occurs", func() {
			It("returns an empty string and the error when the http client returns one one", func() {
				expectOpenAuthURL(authURL, localhostAuthcodeURL, nil)
				expectPOSTRequest(tokenURL, expectedPayload, "fail", 500, errors.New("Failed to sent post request"))

				authenticationToken, err := client.Login(mockHttpClient, mockBrowser)

				Expect(err).To(HaveOccurred())
				Expect(authenticationToken).To(Equal(""))
			})

			It("returns an empty string and the error when the body of the access token is invalid", func() {
				expectOpenAuthURL(authURL, localhostAuthcodeURL, nil)
				expectPOSTRequest(tokenURL, expectedPayload, "no valid body", 200, nil)

				authenticationToken, err := client.Login(mockHttpClient, mockBrowser)

				Expect(err).To(HaveOccurred())
				Expect(authenticationToken).To(Equal(""))
			})

			It("returns an empty string and the error when the redirect_uri does not get a auth code", func() {
				expectOpenAuthURL(authURL, "http://localhost:8261/authcode?fail=reason", nil)

				authenticationToken, err := client.Login(mockHttpClient, mockBrowser)

				Expect(err).To(HaveOccurred())
				Expect(authenticationToken).To(Equal(""))
			})
		})
	})
})

func expectOpenAuthURL(expectedURL string, openURL string, expectedError error) {
	mockBrowser.
		EXPECT().
		OpenURL(expectedURL).
		DoAndReturn(func(url string) error {
			request, _ := http.NewRequest("GET", openURL, nil)
			_, err := http.DefaultClient.Do(request)
			if err != nil {
				return err
			}
			return expectedError
		})
}
