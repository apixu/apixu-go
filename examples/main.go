package main

import (
	"fmt"
	"log"
	"os"

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

	// Current weather
	q := "Amsterdam"
	current, err := a.Current(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	fmt.Println("\nCurrent")
	fmt.Println()

	loc := current.Location
	fmt.Println("Location")
	fmt.Println("\tName:", loc.Name)
	fmt.Println("\tRegion:", loc.Region)
	fmt.Println("\tCountry:", loc.Country)
	fmt.Println("\tLat:", loc.Lat)
	fmt.Println("\tLon:", loc.Lon)
	fmt.Println("\tTimezone:", loc.Timezone)
	fmt.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\tLocaltime:", loc.LocalTime)

	curr := current.Current
	fmt.Println("Current")
	fmt.Println("\tLastUpdatedEpoch:", curr.LastUpdatedEpoch)
	fmt.Println("\tLastUpdated:", curr.LastUpdated)
	fmt.Println("\tTempCelsius:", curr.TempCelsius)
	fmt.Println("\tTempFahrenheit: ", curr.TempFahrenheit)
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

	// Weather forecast
	q = "Rome"
	forecast, err := a.Forecast(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	fmt.Println("\nForecast")
	fmt.Println()

	loc = forecast.Location
	fmt.Println("Location")
	fmt.Println("\tName:", loc.Name)
	fmt.Println("\tRegion:", loc.Region)
	fmt.Println("\tCountry:", loc.Country)
	fmt.Println("\tLat:", loc.Lat)
	fmt.Println("\tLon:", loc.Lon)
	fmt.Println("\tTimezone:", loc.Timezone)
	fmt.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\tLocaltime:", loc.LocalTime)

	curr = forecast.Current
	fmt.Println("Current")
	fmt.Println("\tLastUpdatedEpoch:", curr.LastUpdatedEpoch)
	fmt.Println("\tLastUpdated:", curr.LastUpdated)
	fmt.Println("\tTempCelsius:", curr.TempCelsius)
	fmt.Println("\tTempFahrenheit: ", curr.TempFahrenheit)
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
		fmt.Println("\tDate:", fc.Date)
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

	// Search
	q = "Bouscat"
	search, err := a.Search(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	fmt.Println("Search")
	for _, item := range search {
		fmt.Println("\tID:", item.ID)
		fmt.Println("\tName:", item.Name)
		fmt.Println("\tRegion:", item.Region)
		fmt.Println("\tCountry:", item.Country)
		fmt.Println("\tLat:", item.Lat)
		fmt.Println("\tLon:", item.Lon)
		fmt.Println("\tURL:", item.URL)
		fmt.Println()
	}

	// Conditions list
	conditions, err := a.Conditions()

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	fmt.Println("Conditions")
	for _, c := range conditions {
		fmt.Println("\tID:", c.Code)
		fmt.Println("\tDay:", c.Day)
		fmt.Println("\tNight:", c.Night)
		fmt.Println("\tIcon:", c.Icon)
		fmt.Println()
	}
}
