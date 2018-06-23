package response

// Forecast defines the weather forecast response
type Forecast struct {
	Location Location `json:"location" xml:"location"`
	Current  Current  `json:"current" xml:"current"`
	Forecast struct {
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
		} `json:"forecastday" xml:"forecastday"`
	} `json:"forecast" xml:"forecast"`
}
