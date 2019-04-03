// Package apixu provides interaction with Apixu Weather service
package apixu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/apixu/apixu-go/v2/response"
)

const (
	apiURL                  = "https://api.apixu.com/v1/%s.json?key=%s&%s"
	docWeatherConditionsURL = "https://www.apixu.com/doc/Apixu_weather_conditions.json"
	maxQueryLength          = 256
	httpTimeout             = time.Second * 20
	historyDateFormat       = "2006-01-02"
)

// Apixu Weather API methods
type Apixu interface {
	Conditions() (response.Conditions, error)
	Current(q string) (*response.CurrentWeather, error)
	Forecast(q string, days int, hour *int) (*response.Forecast, error)
	Search(q string) (response.Search, error)
	History(q string, since time.Time, until *time.Time) (*response.History, error)
}

type apixu struct {
	config     Config
	httpClient httpClient
	read       func(r io.Reader) ([]byte, error)
}

type request struct {
	method string
	params url.Values
}

// Conditions retrieves the weather conditions list
func (a *apixu) Conditions() (res response.Conditions, err error) {
	err = a.call(docWeatherConditionsURL, &res)
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

// Forecast retrieves weather forecast
// Hourly  forecast is available for paid licenses only
func (a *apixu) Forecast(q string, days int, hour *int) (res *response.Forecast, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)
	p.Set("days", strconv.Itoa(days))
	if hour != nil {
		p.Set("hour", strconv.Itoa(*hour))
	}

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
// With a paid license, you can request a time range with until parameter.
func (a *apixu) History(q string, since time.Time, until *time.Time) (res *response.History, err error) {
	if err = validateQuery(q); err != nil {
		return
	}

	p := url.Values{}
	p.Set("q", q)
	p.Set("dt", since.Format(historyDateFormat))
	if until != nil {
		p.Set("end_dt", until.Format(historyDateFormat))
	}

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

// getAPIURL generates the full API url for each request
func (a *apixu) getAPIURL(req request) string {
	return fmt.Sprintf(
		apiURL,
		req.method,
		a.config.APIKey,
		req.params.Encode(),
	)
}

// call uses the HTTP client to call the REST service
func (a *apixu) call(apiURL string, b interface{}) error {
	res, err := a.httpClient.Get(apiURL)
	if err != nil {
		return fmt.Errorf("cannot call service (%s)", err)
	}

	body, err := a.read(res.Body)
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
		if err = json.Unmarshal(body, &apiError); err != nil {
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

	if err = json.Unmarshal(body, &b); err != nil {
		return fmt.Errorf("cannot read api response (%s)", err)
	}

	return nil
}

// New creates an Apixu package instance
func New(c Config) (Apixu, error) {
	if strings.TrimSpace(c.APIKey) == "" {
		return nil, errors.New("api key not specified")
	}

	a := &apixu{
		config: c,
		httpClient: &http.Client{
			Timeout: httpTimeout,
		},
		read: ioutil.ReadAll,
	}

	return a, nil
}
