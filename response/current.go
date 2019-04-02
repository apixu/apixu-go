package response

import "github.com/apixu/apixu-go/v2/types"

// Current defines the current weather info
type Current struct {
	LastUpdatedEpoch    int              `json:"last_updated_epoch" xml:"last_updated_epoch"`
	LastUpdated         types.DateTime   `json:"last_updated" xml:"last_updated"`
	TempCelsius         float64          `json:"temp_c" xml:"temp_c"`
	TempFahrenheit      float64          `json:"temp_f" xml:"temp_f"`
	IsDay               types.Bool       `json:"is_day" xml:"is_day"`
	Condition           CurrentCondition `json:"condition" xml:"condition"`
	WindMPH             float64          `json:"wind_mph" xml:"wind_mph"`
	WindKPH             float64          `json:"wind_kph" xml:"wind_kph"`
	WindDegree          int              `json:"wind_degree" xml:"wind_degree"`
	WindDirection       string           `json:"wind_dir" xml:"wind_dir"`
	PressureMB          float64          `json:"pressure_mb" xml:"pressure_mb"`
	PressureIN          float64          `json:"pressure_in" xml:"pressure_in"`
	PrecipMM            float64          `json:"precip_mm" xml:"precip_mm"`
	PrecipIN            float64          `json:"precip_in" xml:"precip_in"`
	Humidity            int              `json:"humidity" xml:"humidity"`
	Cloud               int              `json:"cloud" xml:"cloud"`
	FeelsLikeCelsius    float64          `json:"feelslike_c" xml:"feelslike_c"`
	FeelsLikeFahrenheit float64          `json:"feelslike_f" xml:"feelslike_f"`
	VisKM               float64          `json:"vis_km" xml:"vis_km"`
	VisMiles            float64          `json:"vis_miles" xml:"vis_miles"`
	UV                  float64          `json:"uv" xml:"uv"`
	GustMPH             float64          `json:"gust_mph" xml:"gust_mph"`
	GustKPH             float64          `json:"gust_kph" xml:"gust_kph"`
}

// CurrentCondition defines the condition item for current weather response
type CurrentCondition struct {
	Text string `json:"text" xml:"text"`
	Icon string `json:"icon" xml:"icon"`
	Code int    `json:"code" xml:"code"`
}
