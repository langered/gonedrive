package httpclient

import "net/http"

//HttpClient is an interface representation of the used functionality of net/http client
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
