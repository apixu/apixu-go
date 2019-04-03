package apixu

import (
	"encoding/json"
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

	"github.com/apixu/apixu-go/v2/response"
	"github.com/stretchr/testify/assert"
)

// TestNew
func TestNew(t *testing.T) {
	c := Config{
		APIKey: "apikey",
	}
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.NoError(t, err)
}

func TestNewWithMissingAPIKey(t *testing.T) {
	c := Config{}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

type httpClientMock struct {
	response *http.Response
	err      error
}

func (c *httpClientMock) Get(string) (*http.Response, error) {
	return c.response, c.err
}

type bodyMock struct {
	closeErr error
}

func (*bodyMock) Read(p []byte) (n int, err error) {
	return
}

func (b *bodyMock) Close() error {
	return b.closeErr
}

func TestApixu_Conditions(t *testing.T) {
	data := loadData(t, "conditions")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	expected := &response.Conditions{}
	if err := json.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Conditions()
	assert.Equal(t, *expected, res)
	assert.NoError(t, err)
}

func TestApixu_Current(t *testing.T) {
	data := loadData(t, "current")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	expected := &response.CurrentWeather{}
	if err := json.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Current("query")
	assert.Equal(t, expected, res)
	assert.NoError(t, err)
}

func TestApixu_CurrentWithQueryError(t *testing.T) {
	a := &apixu{
		config: Config{},
	}

	res, err := a.Current(" ")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_Forecast(t *testing.T) {
	data := loadData(t, "forecast")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	expected := &response.Forecast{}
	if err := json.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	hour := 17
	res, err := a.Forecast("query", 2, &hour)
	assert.Equal(t, expected, res)
	assert.NoError(t, err)
}

func TestApixu_ForecastWithQueryError(t *testing.T) {
	a := &apixu{
		config: Config{},
	}

	res, err := a.Forecast(strings.Repeat("q", maxQueryLength+1), 1, nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_Search(t *testing.T) {
	data := loadData(t, "search")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	expected := &response.Search{}
	if err := json.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.Search("query")
	assert.Equal(t, *expected, res)
	assert.NoError(t, err)
}

func TestApixu_SearchWithQueryError(t *testing.T) {
	a := &apixu{
		config: Config{},
	}

	res, err := a.Search("")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_History(t *testing.T) {
	data := loadData(t, "history")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	expected := &response.History{}
	if err := json.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	res, err := a.History("query", time.Time{}, &time.Time{})
	assert.Equal(t, expected, res)
	assert.NoError(t, err)
}

func TestApixu_HistoryWithQueryError(t *testing.T) {
	a := &apixu{
		config: Config{},
	}

	res, err := a.History("", time.Time{}, nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_HttpClientGetError(t *testing.T) {
	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: errors.New("error"),
		},
	}

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_ReadResponseBodyError(t *testing.T) {
	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return []byte{}, errors.New("error")
		},
	}

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_CloseResponseBodyError(t *testing.T) {
	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body: &bodyMock{
					closeErr: errors.New("error"),
				},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return []byte{}, nil
		},
	}

	res, err := a.Search("query")
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_APIErrorResponse(t *testing.T) {
	data := loadData(t, "error")

	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return data, nil
		},
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.IsType(t, &Error{}, err)

	expectedErrorResponse := &response.Error{}
	if e := json.Unmarshal(data, expectedErrorResponse); e != nil {
		assert.Fail(t, e.Error())
	}
	expectedError := &Error{
		err: err.(*Error).err,
		res: expectedErrorResponse.Error,
	}
	assert.Equal(t, expectedError, err)
}

func TestApixu_APIInternalServerError(t *testing.T) {
	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusNotImplemented,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return []byte{}, nil
		},
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestApixu_APIMalformedErrorResponse(t *testing.T) {
	a := &apixu{
		config: Config{},
		httpClient: &httpClientMock{
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       &bodyMock{},
			},
			err: nil,
		},
		read: func(r io.Reader) ([]byte, error) {
			return []byte(`{invalid json}`), nil
		},
	}

	res, err := a.Search("query")

	assert.Nil(t, res)
	assert.Error(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{
		APIKey: "apikey",
	}

	p := url.Values{}
	p.Set("q", "query")

	req := request{
		method: "history",
		params: p,
	}

	expected := fmt.Sprintf(
		apiURL,
		req.method,
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
