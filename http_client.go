package apixu

import "net/http"

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type httpClient struct {
}

var httpGet = http.Get

func (c *httpClient) Get(url string) (*http.Response, error) {
	return httpGet(url)
}
