package apixu

import "net/http"

// HTTPClient defines methods needed for the HTTP client
// used to call the REST service
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type httpClient struct {
}

var httpGet = http.Get

func (c *httpClient) Get(url string) (*http.Response, error) {
	return httpGet(url)
}
