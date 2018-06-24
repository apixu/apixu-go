package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andreiavrammsd/apixu-go"
)

func main() {
	config := apixu.Config{
		Version: "1",
		Format:  "xml",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Fatal(err)
	}

	q := "London"
	since := time.Now()
	history, err := a.History(q, since)
	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	loc := history.Location
	fmt.Println("Location")
	fmt.Println("\tName:", loc.Name)
	fmt.Println("\tRegion:", loc.Region)
	fmt.Println("\tCountry:", loc.Country)
	fmt.Println("\tLat:", loc.Lat)
	fmt.Println("\tLon:", loc.Lon)
	fmt.Println("\tTimezone:", loc.Timezone)
	fmt.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\tLocaltime:", time.Time(loc.LocalTime).String())

	fmt.Println("\nForecast Day")
	for _, fc := range history.Forecast.ForecastDay {
		fmt.Println("\tDate:", time.Time(fc.Date).String())
		fmt.Println("\tDateEpoch:", fc.DateEpoch)
		fmt.Println("\tDay")
		fmt.Println("\t\tMaxTempCelsius:", fc.Day.MaxTempCelsius)
		fmt.Println("\t\tMaxTempFahrenheit:", fc.Day.MaxTempFahrenheit)
		fmt.Println("\t\tMinTempCelsius:", fc.Day.MinTempCelsius)
		fmt.Println("\t\tMinTempFahrenheit:", fc.Day.MinTempFahrenheit)
		fmt.Println("\t\tAvgTempCelsius:", fc.Day.AvgTempCelsius)
		fmt.Println("\t\tAvgTempFahrenheit:", fc.Day.AvgTempFahrenheit)
		fmt.Println("\t\tMaxWindMPH:", fc.Day.MaxWindMPH)
		fmt.Println("\t\tMaxWindKMH:", fc.Day.MaxWindKMH)
		fmt.Println("\t\tTotalPrecipMM:", fc.Day.TotalPrecipMM)
		fmt.Println("\t\tTotalPrecipIN:", fc.Day.TotalPrecipIN)
		fmt.Println("\t\tVisKM:", fc.Day.AvgVisKM)
		fmt.Println("\t\tVisMiles:", fc.Day.AvgVisMiles)
		fmt.Println("\t\tAvgHumidity:", fc.Day.AvgHumidity)
		fmt.Println("\t\tCondition")
		fmt.Println("\t\tText:", fc.Day.Condition.Text)
		fmt.Println("\t\tIcon:", fc.Day.Condition.Icon)
		fmt.Println("\t\tCode:", fc.Day.Condition.Code)
		fmt.Println("\t\tUV:", fc.Day.UV)
		fmt.Println("\t\tAstro")
		fmt.Println("\t\t\tSunrise:", fc.Astro.Sunrise)
		fmt.Println("\t\t\tSunset:", fc.Astro.Sunset)
		fmt.Println("\t\t\tMoonrise:", fc.Astro.Moonrise)
		fmt.Println("\t\t\tMoonset:", fc.Astro.Moonset)

		fmt.Println("\t\tHour:")
		for _, fch := range fc.Hour {
			fmt.Println("\t\t\tTimeEpoch:", fch.TimeEpoch)
			fmt.Println("\t\t\tTime:", time.Time(fch.Time).String())
			fmt.Println("\t\t\tTempCelsius:", fch.TempCelsius)
			fmt.Println("\t\t\tTempFahrenheit:", fch.TempFahrenheit)
			fmt.Println("\t\t\tIsDay:", fch.IsDay)
			fmt.Println("\t\t\tCondition")
			fmt.Println("\t\t\t\tText:", fch.Condition.Text)
			fmt.Println("\t\t\t\tIcon:", fch.Condition.Icon)
			fmt.Println("\t\t\t\tCode:", fch.Condition.Code)
			fmt.Println("\t\t\tWindMPH:", fch.WindMPH)
			fmt.Println("\t\t\tWindKMH:", fch.WindKMH)
			fmt.Println("\t\t\tWindDegree:", fch.WindDegree)
			fmt.Println("\t\t\tWindDirection:", fch.WindDirection)
			fmt.Println("\t\t\tPressureMB:", fch.PressureMB)
			fmt.Println("\t\t\tPressureIN:", fch.PressureIN)
			fmt.Println("\t\t\tPrecipMM:", fch.PrecipMM)
			fmt.Println("\t\t\tPrecipIN:", fch.PrecipIN)
			fmt.Println("\t\t\tHumidity:", fch.Humidity)
			fmt.Println("\t\t\tCloud:", fch.Cloud)
			fmt.Println("\t\t\tFeelsLikeCelsius:", fch.FeelsLikeCelsius)
			fmt.Println("\t\t\tFeelsLikeFahrenheit:", fch.FeelsLikeFahrenheit)
			fmt.Println("\t\t\tWindChillCelsius:", fch.WindChillCelsius)
			fmt.Println("\t\t\tWindChillFahrenheit:", fch.WindChillFahrenheit)
			fmt.Println("\t\t\tHeatIndexCelsius:", fch.HeatIndexCelsius)
			fmt.Println("\t\t\tHeatIndexFahrenheit:", fch.HeatIndexFahrenheit)
			fmt.Println("\t\t\tDewPointCelsius:", fch.DewPointCelsius)
			fmt.Println("\t\t\tDewPointFahrenheit:", fch.DewPointFahrenheit)
			fmt.Println("\t\t\tWillItRain:", fch.WillItRain)
			fmt.Println("\t\t\tChanceOfRain:", fch.ChanceOfRain)
			fmt.Println("\t\t\tWillItSnow:", fch.WillItSnow)
			fmt.Println("\t\t\tChanceOfSnow:", fch.ChanceOfSnow)
			fmt.Println("\t\t\tVisKM:", fch.VisKM)
			fmt.Println("\t\t\tVisMiles:", fch.VisMiles)
			fmt.Println()
		}

		fmt.Println()
	}
}
