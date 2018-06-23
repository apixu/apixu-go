package response

// CurrentWeather defines the current weather response
type CurrentWeather struct {
	Location Location `json:"location" xml:"location"`
	Current  Current  `json:"current" xml:"current"`
}
