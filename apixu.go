// Package apixu provides interaction with Apixu Weather service
package apixu

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/andreiavrammsd/apixu-go/formatter"
	"github.com/andreiavrammsd/apixu-go/response"
)

const (
	apiURL                  = "https://api.apixu.com/v%s/%s.%s?key=%s&%s"
	docWeatherConditionsURL = "https://www.apixu.com/doc/Apixu_weather_conditions.%s"
	maxQueryLength          = 256
	httpTimeout             = time.Second * 20
)

var ioUtilReadAll = ioutil.ReadAll

// Apixu defines methods implemented by Apixu Weather API
type Apixu interface {
	Conditions() (response.Conditions, error)
	Current(q string) (*response.CurrentWeather, error)
	Forecast(q string, days int) (*response.Forecast, error)
	Search(q string) (response.Search, error)
	History(q string, since time.Time) (*response.History, error)
}

type apixu struct {
	config     Config
	httpClient httpClient
	formatter  formatter.Formatter
}

type request struct {
	method string
	params url.Values
}

// Conditions retrieves the weather conditions list
func (a *apixu) Conditions() (res response.Conditions, err error) {
	u := fmt.Sprintf(docWeatherConditionsURL, a.config.Format)
	err = a.call(u, &res)
	return
}

// Current retrieves realtime weather information by city name
func (a *apixu) Current(q string) (res *response.CurrentWeather, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)

	req := request{
		method: "current",
		params: p,
	}

	err = a.call(a.getAPIURL(req), &res)
	return
}

// Forecast retrieves weather forecast for up to next 10 days
func (a *apixu) Forecast(q string, days int) (res *response.Forecast, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)
	p.Set("days", strconv.Itoa(days))

	req := request{
		method: "forecast",
		params: p,
	}

	err = a.call(a.getAPIURL(req), &res)
	return
}

// Search finds cities and towns matching your query (autocomplete)
func (a *apixu) Search(q string) (res response.Search, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)

	req := request{
		method: "search",
		params: p,
	}

	err = a.call(a.getAPIURL(req), &res)
	return
}

// History retrieves historical weather information for a city and a date starting 2015-01-01
func (a *apixu) History(q string, since time.Time) (res *response.History, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)
	p.Set("dt", since.Format("2006-01-02"))

	req := request{
		method: "history",
		params: p,
	}

	err = a.call(a.getAPIURL(req), &res)
	return
}

// validateQuery checks the given query for possible issues
func validateQuery(q string) error {
	q = strings.TrimSpace(q)

	if q == "" {
		return errors.New("query is missing")
	}

	if len(q) > maxQueryLength {
		return fmt.Errorf("query exceeds maximum length (%d)", maxQueryLength)
	}

	return nil
}

// getApiUrl generates the full API url for each request
func (a *apixu) getAPIURL(req request) string {
	return fmt.Sprintf(
		apiURL,
		a.config.Version,
		req.method,
		a.config.Format,
		a.config.APIKey,
		req.params.Encode(),
	)
}

// call uses the HTTP client to call the REST service
func (a *apixu) call(url string, b interface{}) (err error) {
	res, err := a.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("cannot call service (%s)", err)
	}

	body, err := ioUtilReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("cannot read response body (%s)", err)
	}

	defer func() {
		if e := res.Body.Close(); e != nil {
			err = fmt.Errorf("cannot close response body (%s)", e)
		}
	}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode >= http.StatusInternalServerError {
			return fmt.Errorf("internal server error (code %d)", res.StatusCode)
		}

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

	err = a.formatter.Unmarshal(body, b)
	if err != nil {
		return fmt.Errorf("cannot read api response (%s)", err)
	}

	return
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
		config: c,
		httpClient: &http.Client{
			Timeout: httpTimeout,
		},
		formatter: f,
	}

	return a, nil
}
