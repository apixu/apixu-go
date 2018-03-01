package apixu

import (
	"errors"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	httpGet = func(string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	client := httpClient{}
	res, err := client.Get("https://domain.tld")

	assert.IsType(t, &http.Response{}, res)
	assert.Nil(t, err)
}

func TestGetFail(t *testing.T) {
	httpGet = func(string) (*http.Response, error) {
		return &http.Response{}, errors.New("")
	}
	client := httpClient{}
	res, err := client.Get("https://domain.tld")

	assert.IsType(t, &http.Response{}, res)
	assert.NotNil(t, err)
}
