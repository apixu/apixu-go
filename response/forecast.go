package response

import "github.com/apixu/apixu-go/v2/types"

// Forecast defines the weather forecast response
type Forecast struct {
	Location Location        `json:"location" xml:"location"`
	Current  Current         `json:"current" xml:"current"`
	Forecast ForecastWeather `json:"forecast" xml:"forecast"`
}

// ForecastWeather holds the forecast data
type ForecastWeather struct {
	ForecastDay []struct {
		Date      types.DateTime `json:"date" xml:"date"`
		DateEpoch int            `json:"date_epoch" xml:"date_epoch"`
		Day       struct {
			MaxTempCelsius    float64          `json:"maxtemp_c" xml:"maxtemp_c"`
			MaxTempFahrenheit float64          `json:"maxtemp_f" xml:"maxtemp_f"`
			MinTempCelsius    float64          `json:"mintemp_c" xml:"mintemp_c"`
			MinTempFahrenheit float64          `json:"mintemp_f" xml:"mintemp_f"`
			AvgTempCelsius    float64          `json:"avgtemp_c" xml:"avgtemp_c"`
			AvgTempFahrenheit float64          `json:"avgtemp_f" xml:"avgtemp_f"`
			MaxWindMPH        float64          `json:"maxwind_mph" xml:"maxwind_mph"`
			MaxWindKPH        float64          `json:"maxwind_kph" xml:"maxwind_kph"`
			TotalPrecipMM     float64          `json:"totalprecip_mm" xml:"total_precip_mm"`
			TotalPrecipIN     float64          `json:"totalprecip_in" xml:"total_precip_in"`
			AvgVisKM          float64          `json:"avgvis_km" xml:"avgvis_km"`
			AvgVisMiles       float64          `json:"avgvis_miles" xml:"avgvis_miles"`
			AvgHumidity       float64          `json:"avghumidity" xml:"avghumidity"`
			Condition         CurrentCondition `json:"condition" xml:"condition"`
			UV                float64          `json:"uv" xml:"uv"`
		} `json:"day" xml:"day"`
		Astro struct {
			Sunrise          string `json:"sunrise" xml:"sunrise"`
			Sunset           string `json:"sunset" xml:"sunset"`
			Moonrise         string `json:"moonrise" xml:"moonrise"`
			Moonset          string `json:"moonset" xml:"moonset"`
			MoonPhase        string `json:"moon_phase" xml:"moon_phase"`
			MoonIllumination string `json:"moon_illumination" xml:"moon_illumination"`
		} `json:"astro" xml:"astro"`
		Hour []struct {
			Condition           CurrentCondition `json:"condition" xml:"condition"`
			Time                types.DateTime   `json:"time" xml:"time"`
			WindDirection       string           `json:"wind_dir" xml:"wind_dir"`
			ChanceOfRain        string           `json:"chance_of_rain" xml:"chance_of_rain"`
			ChanceOfSnow        string           `json:"chance_of_snow" xml:"chance_of_snow"`
			TimeEpoch           int              `json:"time_epoch" xml:"time_epoch"`
			TempCelsius         float64          `json:"temp_c" xml:"temp_c"`
			TempFahrenheit      float64          `json:"temp_f" xml:"temp_f"`
			WindMPH             float64          `json:"wind_mph" xml:"wind_mph"`
			WindKPH             float64          `json:"wind_kph" xml:"wind_kph"`
			WindDegree          int              `json:"wind_degree" xml:"wind_degree"`
			PressureMB          float64          `json:"pressure_mb" xml:"pressure_mb"`
			PressureIN          float64          `json:"pressure_in" xml:"pressure_in"`
			PrecipMM            float64          `json:"precip_mm" xml:"precip_mm"`
			PrecipIN            float64          `json:"precip_in" xml:"precip_in"`
			Humidity            int              `json:"humidity" xml:"humidity"`
			Cloud               int              `json:"cloud" xml:"cloud"`
			FeelsLikeCelsius    float64          `json:"feelslike_c" xml:"feelslike_c"`
			FeelsLikeFahrenheit float64          `json:"feelslike_f" xml:"feelslike_f"`
			WindChillCelsius    float64          `json:"windchill_c" xml:"windchill_c"`
			WindChillFahrenheit float64          `json:"windchill_f" xml:"windchill_f"`
			HeatIndexCelsius    float64          `json:"heatindex_c" xml:"heatindex_c"`
			HeatIndexFahrenheit float64          `json:"heatindex_f" xml:"heatindex_f"`
			DewPointCelsius     float64          `json:"dewpoint_c" xml:"dewpoint_c"`
			DewPointFahrenheit  float64          `json:"dewpoint_f" xml:"dewpoint_f"`
			VisKM               float64          `json:"vis_km" xml:"vis_km"`
			VisMiles            float64          `json:"vis_miles" xml:"vis_miles"`
			IsDay               types.Bool       `json:"is_day" xml:"is_day"`
			WillItRain          types.Bool       `json:"will_it_rain" xml:"will_it_rain"`
			WillItSnow          types.Bool       `json:"will_it_snow" xml:"will_it_snow"`
			GustMPH             float64          `json:"gust_mph" xml:"gust_mph"`
			GustKPH             float64          `json:"gust_kph" xml:"gust_kph"`
		} `json:"hour,omitempty" xml:"hour,omitempty"`
	} `json:"forecastday" xml:"forecastday"`
}
