package apixu

import "net/http"

type httpClient interface {
	Get(url string) (*http.Response, error)
}
