package service

import (
	"github.com/langered/gonedrive/browser"
	"github.com/langered/gonedrive/httpclient"
)

//StoreClient is the interface definition for required cloud endpoints
type StoreClient interface {
	Get(httpClient httpclient.HttpClient, accessToken string, remotePath string) (string, error)
	Upload(httpClient httpclient.HttpClient, accessToken string, remoteFilePath string, content string) (bool, error)
	Login(httpClient httpclient.HttpClient, browser browser.Browser) (string, error)
	List(httpClient httpclient.HttpClient, accessToken string, remotePath string) ([]string, error)
	Delete(httpClient httpclient.HttpClient, accessToken string, remotePath string) error
}
