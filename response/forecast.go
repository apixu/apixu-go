package response

// Forecast defines the weather forecast response
type Forecast struct {
	Location Location        `json:"location" xml:"location"`
	Current  Current         `json:"current" xml:"current"`
	Forecast ForecastWeather `json:"forecast" xml:"forecast"`
}

// ForecastWeather holds the forecast data
type ForecastWeather struct {
	ForecastDay []struct {
		Date      string `json:"date" xml:"date"`
		DateEpoch int32  `json:"date_epoch" xml:"date_epoch"`
		Day       struct {
			MaxTempCelsius    float32          `json:"maxtemp_c" xml:"maxtemp_c"`
			MaxTempFahrenheit float32          `json:"maxtemp_f" xml:"maxtemp_f"`
			MinTempCelsius    float32          `json:"mintemp_c" xml:"mintemp_c"`
			MinTempFahrenheit float32          `json:"mintemp_f" xml:"mintemp_f"`
			AvgTempCelsius    float32          `json:"avgtemp_c" xml:"avgtemp_c"`
			AvgTempFahrenheit float32          `json:"avgtemp_f" xml:"avgtemp_f"`
			MaxWindMPH        float32          `json:"maxwind_mph" xml:"maxwind_mph"`
			MaxWindKMH        float32          `json:"maxwind_kph" xml:"maxwind_kph"`
			TotalPrecipMM     float32          `json:"totalprecip_mm" xml:"total_precip_mm"`
			TotalPrecipIN     float32          `json:"totalprecip_in" xml:"total_precip_in"`
			AvgVisKM          float32          `json:"avgvis_km" xml:"avgvis_km"`
			AvgVisMiles       float32          `json:"avgvis_miles" xml:"avgvis_miles"`
			AvgHumidity       float32          `json:"avghumidity" xml:"avghumidity"`
			Condition         CurrentCondition `json:"condition" xml:"condition"`
			UV                float32          `json:"uv" xml:"uv"`
		} `json:"day" xml:"day"`
		Astro struct {
			Sunrise  string `json:"sunrise" xml:"sunrise"`
			Sunset   string `json:"sunset" xml:"sunset"`
			Moonrise string `json:"moonrise" xml:"moonrise"`
			Moonset  string `json:"moonset" xml:"moonset"`
		} `json:"astro" xml:"astro"`
		Hour []struct {
			TimeEpoch           int32            `json:"time_epoch" xml:"time_epoch"`
			Time                string           `json:"time" xml:"time"`
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
			WindChillCelsius    float32          `json:"windchill_c" xml:"windchill_c"`
			WindChillFahrenheit float32          `json:"windchill_f" xml:"windchill_f"`
			HeatIndexCelsius    float32          `json:"heatindex_c" xml:"heatindex_c"`
			HeatIndexFahrenheit float32          `json:"heatindex_f" xml:"heatindex_f"`
			DewPointCelsius     float32          `json:"dewpoint_c" xml:"dewpoint_c"`
			DewPointFahrenheit  float32          `json:"dewpoint_f" xml:"dewpoint_f"`
			WillItRain          uint8            `json:"will_it_rain" xml:"will_it_rain"`
			ChanceOfRain        string           `json:"chance_of_rain" xml:"chance_of_rain"`
			WillItSnow          uint8            `json:"will_it_snow" xml:"will_it_snow"`
			ChanceOfSnow        string           `json:"chance_of_snow" xml:"chance_of_snow"`
			VisKM               float32          `json:"vis_km" xml:"vis_km"`
			VisMiles            float32          `json:"vis_miles" xml:"vis_miles"`
		} `json:"hour,omitempty" xml:"hour,omitempty"`
	} `json:"forecastday" xml:"forecastday"`
}
