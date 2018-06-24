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
		Format:  "json",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Fatal(err)
	}

	q := "Rome"
	days := 5
	forecast, err := a.Forecast(q, days)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	loc := forecast.Location
	fmt.Println("Location")
	fmt.Println("\tName:", loc.Name)
	fmt.Println("\tRegion:", loc.Region)
	fmt.Println("\tCountry:", loc.Country)
	fmt.Println("\tLat:", loc.Lat)
	fmt.Println("\tLon:", loc.Lon)
	fmt.Println("\tTimezone:", loc.Timezone)
	fmt.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\tLocaltime:", time.Time(loc.LocalTime).String())

	curr := forecast.Current
	fmt.Println("Current")
	fmt.Println("\tLastUpdatedEpoch:", curr.LastUpdatedEpoch)
	fmt.Println("\tLastUpdated:", time.Time(curr.LastUpdated).String())
	fmt.Println("\tTempCelsius:", curr.TempCelsius)
	fmt.Println("\tTempFahrenheit:", curr.TempFahrenheit)
	fmt.Println("\tIsDay:", curr.IsDay)
	fmt.Println("\tCondition")
	fmt.Println("\t\tText:", curr.Condition.Text)
	fmt.Println("\t\tIcon:", curr.Condition.Icon)
	fmt.Println("\t\tCode:", curr.Condition.Code)
	fmt.Println("\tWindMPH:", curr.WindMPH)
	fmt.Println("\tWindKMH:", curr.WindKMH)
	fmt.Println("\tWindDegree:", curr.WindDegree)
	fmt.Println("\tWindDirection:", curr.WindDirection)
	fmt.Println("\tPressureMB:", curr.PressureMB)
	fmt.Println("\tPressureIN:", curr.PressureIN)
	fmt.Println("\tPrecipMM:", curr.PrecipMM)
	fmt.Println("\tPrecipIN:", curr.PrecipIN)
	fmt.Println("\tHumidity:", curr.Humidity)
	fmt.Println("\tCloud:", curr.Cloud)
	fmt.Println("\tFeelsLikeCelsius:", curr.FeelsLikeCelsius)
	fmt.Println("\tFeelsLikeFahrenheit:", curr.FeelsLikeFahrenheit)
	fmt.Println("\tVisKM:", curr.VisKM)
	fmt.Println("\tVisMiles:", curr.VisMiles)

	fcast := forecast.Forecast.ForecastDay
	fmt.Println("Forecast Day")
	for _, fc := range fcast {
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
		fmt.Println("\t\t\tText:", fc.Day.Condition.Text)
		fmt.Println("\t\t\tIcon:", fc.Day.Condition.Icon)
		fmt.Println("\t\t\tCode:", fc.Day.Condition.Code)
		fmt.Println("\t\tUV:", fc.Day.UV)
		fmt.Println("\tAstro")
		fmt.Println("\t\tSunrise:", fc.Astro.Sunrise)
		fmt.Println("\t\tSunset:", fc.Astro.Sunset)
		fmt.Println("\t\tMoonrise:", fc.Astro.Moonrise)
		fmt.Println("\t\tMoonset:", fc.Astro.Moonset)
		fmt.Println()
	}
}
