// Package apixu provides interaction with Apixu Weather service
package apixu

import (
	"errors"
	"fmt"
	"io/ioutil"
	. "github.com/andreiavrammsd/apixu-go/response"
)

const API_URL = "https://api.apixu.com/v%s/%s.%s?key=%s&q=%s"
const DOC_WEATHER_CONDITIONS_URL = "https://www.apixu.com/doc/Apixu_weather_conditions.%s"

type Apixu interface {
	GetConditions() (Conditions, error)
	GetCurrent(q string) (CurrentWeather, error)
	Search(q string) (Search, error)
}

type apixu struct {
	config     Config
	httpClient HttpClient
	formatter  Formatter
}

type request struct {
	method string
	query  string
}

// GetConditions retrieves the weather conditions list
func (a *apixu) GetConditions() (Conditions, error) {
	url := fmt.Sprintf(DOC_WEATHER_CONDITIONS_URL, a.config.Format)
	c := Conditions{}

	err := a.call(url, &c)

	return c, err
}

// GetCurrent retrieves current weather data
func (a *apixu) GetCurrent(q string) (CurrentWeather, error) {
	r := request{
		"current",
		q,
	}
	url := a.getApiUrl(r)
	w := CurrentWeather{}

	err := a.call(url, &w)

	return w, err
}

// Search finds cities and towns matching your query (autocomplete)
func (a *apixu) Search(q string) (Search, error) {
	r := request{
		"search",
		q,
	}
	url := a.getApiUrl(r)

	s := Search{}

	err := a.call(url, &s)

	return s, err
}

var ioutilReadAll = ioutil.ReadAll

// call uses the HTTP Client to call the REST service
func (a *apixu) call(url string, b interface{}) error {
	res, err := a.httpClient.Get(url)
	if err != nil {
		return &ApixuError{err, ErrorResponse{}}
	}

	body, err := ioutilReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return &ApixuError{err, ErrorResponse{}}
	}

	if res.StatusCode >= 400 {
		apiError := Error{}
		err = a.formatter.Unmarshal(body, &apiError)
		if err != nil {
			return &ApixuError{err, ErrorResponse{}}
		}

		return &ApixuError{
			errors.New(
				fmt.Sprintf(
					"%s (%d)",
					apiError.Error.Message,
					apiError.Error.Code,
				),
			),
			apiError.Error,
		}
	}

	err = a.formatter.Unmarshal(body, &b)
	if err != nil {
		return &ApixuError{err, ErrorResponse{}}
	}

	return nil
}

// getApiUrl forms the full API url for each request
func (a *apixu) getApiUrl(r request) string {
	return fmt.Sprintf(
		API_URL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.ApiKey,
		r.query,
	)
}

func New(c Config) (Apixu, error) {
	a := &apixu{}

	formatter, err := NewFormatter(c.Format)
	if err != nil {
		return a, err
	}

	a.config = c
	a.httpClient = &httpClient{}
	a.formatter = formatter

	return a, nil
}
