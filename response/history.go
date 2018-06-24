package response

// History defines the historical weather info
type History struct {
	Location Location        `json:"location" xml:"location"`
	Forecast ForecastWeather `json:"forecast" xml:"forecast"`
}
