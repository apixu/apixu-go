package apixu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

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

var (
	httpClientResponse = &http.Response{
		StatusCode: 200,
		Body:       &bodyMock{},
	}
	httpClientError                  error
	httpClientResponseBodyCloseError error
)

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

type jsonFormatterMock struct {
}

func (*jsonFormatterMock) Unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, object)
}

func TestApixu_Conditions(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`[  
   		{  
      		"code":1000,
      		"day":"Sunny",
      		"night":"Clear",
      		"icon":113
   		},
   		{  
      		"code":1003,
      		"day":"Partly cloudy",
      		"night":"Partly cloudy",
      		"icon":116
   		}
		]`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Conditions{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Conditions()
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_Current(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
   "location":{  
      "name":"Amsterdam",
      "region":"North Holland",
      "country":"Netherlands",
      "lat":52.37,
      "lon":4.89,
      "tz_id":"Europe/Amsterdam",
      "localtime_epoch":1529746782,
      "localtime":"2018-06-23 11:39"
   },
   "current":{  
      "last_updated_epoch":1529746209,
      "last_updated":"2018-06-23 11:30",
      "temp_c":15.0,
      "temp_f":59.0,
      "is_day":1,
      "condition":{  
         "text":"Partly cloudy",
         "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
         "code":1003
      },
      "wind_mph":9.4,
      "wind_kph":15.1,
      "wind_degree":320,
      "wind_dir":"NW",
      "pressure_mb":1027.0,
      "pressure_in":30.8,
      "precip_mm":0.1,
      "precip_in":0.0,
      "humidity":72,
      "cloud":75,
      "feelslike_c":14.1,
      "feelslike_f":57.3,
      "vis_km":10.0,
      "vis_miles":6.0
   }
}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.CurrentWeather{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Current("query")
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_Search(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`[  
   {  
      "id":3332210,
      "name":"Amsterdam, North Holland, Netherlands",
      "region":"North Holland",
      "country":"Netherlands",
      "lat":52.37,
      "lon":4.89,
      "url":"amsterdam-north-holland-netherlands"
   },
   {  
      "id":3332149,
      "name":"De Wallen, North Holland, Netherlands",
      "region":"North Holland",
      "country":"Netherlands",
      "lat":52.37,
      "lon":4.9,
      "url":"de-wallen-north-holland-netherlands"
   }]`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Search{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Search("query")
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_HttpClientGetError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = errors.New("error")

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return nil, nil
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_ReadResponseBodyError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return nil, errors.New("error")
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_CloseResponseBodyError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = nil
	httpClientResponseBodyCloseError = errors.New("error")

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte(""), nil
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_APIErrorResponse(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
   		"error":{  
      	"code":1005,
      	"message":"API URL is invalid."
   		}
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	r, err := a.Search("query")

	expected := response.Search{}
	assert.Equal(t, expected, r)

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

func TestApixu_APIMalformedErrorResponse(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
   		invalid json
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	r, err := a.Search("query")

	expected := response.Search{}
	assert.Equal(t, expected, r)

	assert.Error(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{"1", "xml", "apikey"}
	r := request{"GET", "query"}

	expected := fmt.Sprintf(
		apiURL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.APIKey,
		r.query,
	)
	result := a.getAPIURL(r)

	assert.Equal(t, expected, result)
}
