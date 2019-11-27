package httpclient

import "net/http"

//HttpClient is an interface representation of the used functionality of net/http client
type HttpClient interface {
	Get(url string) (*http.Response, error)
	Do(request *http.Request) (*http.Response, error)
}
