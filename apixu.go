// Package apixu provides interaction with Apixu Weather service
package apixu

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/andreiavrammsd/apixu-go/formatter"
	"github.com/andreiavrammsd/apixu-go/response"
)

const apiURL = "https://api.apixu.com/v%s/%s.%s?key=%s&q=%s"
const docWeatherConditionsURL = "https://www.apixu.com/doc/Apixu_weather_conditions.%s"

// Apixu defines methods implemented by Apixu weather API
type Apixu interface {
	GetConditions() (response.Conditions, error)
	GetCurrent(q string) (response.CurrentWeather, error)
	Search(q string) (response.Search, error)
}

type apixu struct {
	config     Config
	httpClient HTTPClient
	formatter  formatter.Formatter
}

type request struct {
	method string
	query  string
}

// GetConditions retrieves the weather conditions list
func (a *apixu) GetConditions() (response.Conditions, error) {
	url := fmt.Sprintf(docWeatherConditionsURL, a.config.Format)
	c := response.Conditions{}

	err := a.call(url, &c)

	return c, err
}

// GetCurrent retrieves current weather data
func (a *apixu) GetCurrent(q string) (response.CurrentWeather, error) {
	r := request{
		"current",
		q,
	}
	url := a.getAPIURL(r)
	w := response.CurrentWeather{}

	err := a.call(url, &w)

	return w, err
}

// Search finds cities and towns matching your query (autocomplete)
func (a *apixu) Search(q string) (response.Search, error) {
	r := request{
		"search",
		q,
	}
	url := a.getAPIURL(r)

	s := response.Search{}

	err := a.call(url, &s)

	return s, err
}

var ioutilReadAll = ioutil.ReadAll

// call uses the HTTP Client to call the REST service
func (a *apixu) call(url string, b interface{}) (err error) {
	res, err := a.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("cannot call service (%s)", err)
	}

	body, err := ioutilReadAll(res.Body)
	if err != nil {
		return errors.New("cannot read response")
	}

	defer func() {
		if e := res.Body.Close(); e != nil {
			err = fmt.Errorf("cannot close response body (%s)", e)
		}
	}()

	if res.StatusCode >= 400 {
		apiError := response.Error{}
		err = a.formatter.Unmarshal(body, &apiError)
		if err != nil {
			return fmt.Errorf("malformed error response (%s)", err)
		}

		return &Error{
			fmt.Errorf(
				"%s (%d)",
				apiError.Error.Message,
				apiError.Error.Code,
			),
			apiError.Error,
		}
	}

	err = a.formatter.Unmarshal(body, &b)
	if err != nil {
		return fmt.Errorf("malformed response (%s)", err)
	}

	return
}

// getApiUrl forms the full API url for each request
func (a *apixu) getAPIURL(r request) string {
	return fmt.Sprintf(
		apiURL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.APIKey,
		r.query,
	)
}

// New creates an Apixu package instance
func New(c Config) (Apixu, error) {
	if c.Version == "" {
		return nil, errors.New("api version not specified")
	}

	if c.APIKey == "" {
		return nil, errors.New("api key not specified")
	}

	f, err := formatter.New(c.Format)
	if err != nil {
		return nil, err
	}

	a := &apixu{
		config:     c,
		httpClient: &httpClient{},
		formatter:  f,
	}

	return a, nil
}
