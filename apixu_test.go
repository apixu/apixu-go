package apixu

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/andreiavrammsd/apixu-go/response"
	"github.com/stretchr/testify/assert"
)

// TestNew
func TestNew(t *testing.T) {
	c := Config{
		Version: "1",
		Format:  "json",
		APIKey:  "apikey",
	}
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.NoError(t, err)
}

func TestNewWithMissingVersion(t *testing.T) {
	c := Config{}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithMissingAPIKey(t *testing.T) {
	c := Config{
		Version: "1",
	}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithUnknownFormat(t *testing.T) {
	c := Config{
		Version: "1",
		APIKey:  "apikey",
		Format:  "txt",
	}
	a, err := New(c)

	assert.Nil(t, a)
	assert.Error(t, err)
}

type httpClientMock struct {
}

func (*httpClientMock) Get(url string) (*http.Response, error) {
	return httpClientResponse, httpClientError
}

type bodyMock struct {
}

func (*bodyMock) Read(p []byte) (n int, err error) {
	return
}

func (*bodyMock) Close() error {
	return httpClientResponseBodyCloseError
}

var (
	httpClientResponse = &http.Response{
		StatusCode: 200,
		Body:       &bodyMock{},
	}
	httpClientError                  error
	httpClientResponseBodyCloseError error
)

type jsonFormatterMock struct {
}

func (*jsonFormatterMock) Unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, object)
}

type xmlFormatterMock struct {
}

func (*xmlFormatterMock) Unmarshal(data []byte, object interface{}) error {
	return xml.Unmarshal(data, object)
}

func TestApixu_Conditions(t *testing.T) {
	f := &jsonFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := loadData(t, "conditions")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Conditions{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Conditions()
	assert.Equal(t, *expected, res)
	assert.NoError(t, err)
}

func TestApixu_Current(t *testing.T) {
	f := &jsonFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := loadData(t, "current")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.CurrentWeather{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Current("query")
	assert.Equal(t, expected, res)
	assert.NoError(t, err)
}

func TestApixu_CurrentWithQueryError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	res, err := a.Current(" ")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_Forecast(t *testing.T) {
	f := &jsonFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := loadData(t, "forecast")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Forecast{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Forecast("query", 2)
	assert.Equal(t, expected, res)
	assert.NoError(t, err)
}

func TestApixu_ForecastWithQueryError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	res, err := a.Forecast(strings.Repeat("q", maxQueryLength+1), 1)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_Search(t *testing.T) {
	f := &xmlFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := loadData(t, "search")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Search{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Search("query")
	assert.Equal(t, *expected, res)
	assert.NoError(t, err)
}

func TestApixu_SearchWithQueryError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	res, err := a.Search("")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_History(t *testing.T) {
	f := &jsonFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := loadData(t, "history")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.History{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.History("query", time.Time{})
	assert.Equal(t, *expected, *res)
	assert.NoError(t, err)
}

func TestApixu_HistoryWithQueryError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	res, err := a.History("", time.Time{})
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_HttpClientGetError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}
	httpClientError = errors.New("error")

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_ReadResponseBodyError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	httpClientError = nil

	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte{}, errors.New("error")
	}

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_CloseResponseBodyError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	httpClientError = nil
	httpClientResponseBodyCloseError = errors.New("error")

	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte{}, nil
	}

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_APIErrorResponse(t *testing.T) {
	f := &jsonFormatterMock{}
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	data := loadData(t, "error")
	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.IsType(t, &Error{}, err)

	expectedErrorResponse := &response.Error{}
	if err := f.Unmarshal(data, expectedErrorResponse); err != nil {
		assert.Fail(t, err.Error())
	}
	expectedError := &Error{
		err: err.(*Error).err,
		res: expectedErrorResponse.Error,
	}
	assert.Equal(t, expectedError, err)
}

func TestApixu_APIInternalServerError(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	httpClientResponse.StatusCode = 501
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte{}, nil
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_APIMalformedErrorResponse(t *testing.T) {
	a := &apixu{
		config:     Config{},
		httpClient: &httpClientMock{},
		formatter:  &jsonFormatterMock{},
	}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	ioUtilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte(`{invalid json}`), nil
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{
		Version: "1",
		Format:  "xml",
		APIKey:  "apikey",
	}

	p := url.Values{}
	p.Set("q", "query")

	req := request{
		method: "history",
		params: p,
	}

	expected := fmt.Sprintf(
		apiURL,
		a.config.Version,
		req.method,
		a.config.Format,
		a.config.APIKey,
		p.Encode(),
	)
	result := a.getAPIURL(req)

	assert.Equal(t, expected, result)
}

func loadData(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
