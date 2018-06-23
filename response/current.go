package response

// Current defines the current weather info
type Current struct {
	LastUpdatedEpoch    int32            `json:"last_updated_epoch" xml:"last_updated_epoch"`
	LastUpdated         string           `json:"last_updated" xml:"last_updated"`
	TempCelsius         float32          `json:"temp_c" xml:"temp_c"`
	TempFahrenheit      float32          `json:"temp_f" xml:"temp_f"`
	IsDay               uint8            `json:"is_day" xml:"is_day"`
	Condition           CurrentCondition `json:"condition" xml:"condition"`
	WindMPH             float32          `json:"wind_mph" xml:"wind_mph"`
	WindKMH             float32          `json:"wind_kph" xml:"wind_kph"`
	WindDegree          uint32           `json:"wind_degree" xml:"wind_degree"`
	WindDirection       string           `json:"wind_dir" xml:"wind_dir"`
	PressureMB          float32          `json:"pressure_mb" xml:"pressure_mb"`
	PressureIN          float32          `json:"pressure_in" xml:"pressure_in"`
	PrecipMM            float32          `json:"precip_mm" xml:"precip_mm"`
	PrecipIN            float32          `json:"precip_in" xml:"precip_in"`
	Humidity            uint16           `json:"humidity" xml:"humidity"`
	Cloud               uint16           `json:"cloud" xml:"cloud"`
	FeelsLikeCelsius    float32          `json:"feelslike_c" xml:"feelslike_c"`
	FeelsLikeFahrenheit float32          `json:"feelslike_f" xml:"feelslike_f"`
	VisKM               float32          `json:"vis_km" xml:"vis_km"`
	VisMiles            float32          `json:"vis_miles" xml:"vis_miles"`
}

// CurrentCondition defines the condition item for current weather response
type CurrentCondition struct {
	Text string `json:"text" xml:"text"`
	Icon string `json:"icon" xml:"icon"`
	Code uint32 `json:"code" xml:"code"`
}
