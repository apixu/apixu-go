package main

import (
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

	loc := current.Location
	log.Println("Location")
	log.Println("\tName:", loc.Name)
	log.Println("\tRegion:", loc.Region)
	log.Println("\tCountry:", loc.Country)
	log.Println("\tLat:", loc.Lat)
	log.Println("\tLon:", loc.Lon)
	log.Println("\tTimezone:", loc.Timezone)
	log.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	log.Println("\tLocaltime:", loc.LocalTime)

	curr := current.Current
	log.Println("Current")
	log.Println("\tLastUpdatedEpoch:", curr.LastUpdatedEpoch)
	log.Println("\tLastUpdated:", curr.LastUpdated)
	log.Println("\tTempCelsius:", curr.TempCelsius)
	log.Println("\tTempFahrenheit: ", curr.TempFahrenheit)
	log.Println("\tIsDay:", curr.IsDay)
	log.Println("\tCondition")
	log.Println("\t\tText:", curr.Condition.Text)
	log.Println("\t\tIcon:", curr.Condition.Icon)
	log.Println("\t\tCode:", curr.Condition.Code)
	log.Println("\tWindMPH:", curr.WindMPH)
	log.Println("\tWindKMH:", curr.WindKMH)
	log.Println("\tWindDegree:", curr.WindDegree)
	log.Println("\tWindDirection:", curr.WindDirection)
	log.Println("\tPressureMB:", curr.PressureMB)
	log.Println("\tPressureIN:", curr.PressureIN)
	log.Println("\tPrecipMM:", curr.PrecipMM)
	log.Println("\tPrecipIN:", curr.PrecipIN)
	log.Println("\tHumidity:", curr.Humidity)
	log.Println("\tCloud:", curr.Cloud)
	log.Println("\tFeelsLikeCelsius:", curr.FeelsLikeCelsius)
	log.Println("\tFeelsLikeFahrenheit:", curr.FeelsLikeFahrenheit)
	log.Println("\tVisKM:", curr.VisKM)
	log.Println("\tVisMiles:", curr.VisMiles)

	// Weather forecast
	q = "Rome"
	forecast, err := a.Forecast(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	loc = forecast.Location
	log.Println("Location")
	log.Println("\tName:", loc.Name)
	log.Println("\tRegion:", loc.Region)
	log.Println("\tCountry:", loc.Country)
	log.Println("\tLat:", loc.Lat)
	log.Println("\tLon:", loc.Lon)
	log.Println("\tTimezone:", loc.Timezone)
	log.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	log.Println("\tLocaltime:", loc.LocalTime)

	curr = forecast.Current
	log.Println("Current")
	log.Println("\tLastUpdatedEpoch:", curr.LastUpdatedEpoch)
	log.Println("\tLastUpdated:", curr.LastUpdated)
	log.Println("\tTempCelsius:", curr.TempCelsius)
	log.Println("\tTempFahrenheit: ", curr.TempFahrenheit)
	log.Println("\tIsDay:", curr.IsDay)
	log.Println("\tCondition")
	log.Println("\t\tText:", curr.Condition.Text)
	log.Println("\t\tIcon:", curr.Condition.Icon)
	log.Println("\t\tCode:", curr.Condition.Code)
	log.Println("\tWindMPH:", curr.WindMPH)
	log.Println("\tWindKMH:", curr.WindKMH)
	log.Println("\tWindDegree:", curr.WindDegree)
	log.Println("\tWindDirection:", curr.WindDirection)
	log.Println("\tPressureMB:", curr.PressureMB)
	log.Println("\tPressureIN:", curr.PressureIN)
	log.Println("\tPrecipMM:", curr.PrecipMM)
	log.Println("\tPrecipIN:", curr.PrecipIN)
	log.Println("\tHumidity:", curr.Humidity)
	log.Println("\tCloud:", curr.Cloud)
	log.Println("\tFeelsLikeCelsius:", curr.FeelsLikeCelsius)
	log.Println("\tFeelsLikeFahrenheit:", curr.FeelsLikeFahrenheit)
	log.Println("\tVisKM:", curr.VisKM)
	log.Println("\tVisMiles:", curr.VisMiles)

	fcast := forecast.Forecast.ForecastDay
	log.Println("Forecast Day")
	for _, fc := range fcast {
		log.Println("\tDate:", fc.Date)
		log.Println("\tDateEpoch:", fc.DateEpoch)
		log.Println("\tDay")
		log.Println("\t\tMaxTempCelsius:", fc.Day.MaxTempCelsius)
		log.Println("\t\tMaxTempFahrenheit:", fc.Day.MaxTempFahrenheit)
		log.Println("\t\tMinTempCelsius:", fc.Day.MinTempCelsius)
		log.Println("\t\tMinTempFahrenheit:", fc.Day.MinTempFahrenheit)
		log.Println("\t\tAvgTempCelsius:", fc.Day.AvgTempCelsius)
		log.Println("\t\tAvgTempFahrenheit:", fc.Day.AvgTempFahrenheit)
		log.Println("\t\tMaxWindMPH:", fc.Day.MaxWindMPH)
		log.Println("\t\tMaxWindKMH:", fc.Day.MaxWindKMH)
		log.Println("\t\tTotalPrecipMM:", fc.Day.TotalPrecipMM)
		log.Println("\t\tTotalPrecipIN:", fc.Day.TotalPrecipIN)
		log.Println("\t\tVisKM:", fc.Day.AvgVisKM)
		log.Println("\t\tVisMiles:", fc.Day.AvgVisMiles)
		log.Println("\t\tAvgHumidity:", fc.Day.AvgHumidity)
		log.Println("\t\tCondition")
		log.Println("\t\t\tText:", fc.Day.Condition.Text)
		log.Println("\t\t\tIcon:", fc.Day.Condition.Icon)
		log.Println("\t\t\tCode:", fc.Day.Condition.Code)
		log.Println("\t\tUV:", fc.Day.UV)
		log.Println("\tAstro")
		log.Println("\t\tSunrise:", fc.Astro.Sunrise)
		log.Println("\t\tSunset:", fc.Astro.Sunset)
		log.Println("\t\tMoonrise:", fc.Astro.Moonrise)
		log.Println("\t\tMoonset:", fc.Astro.Moonset)
		log.Println()
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

	for _, item := range search {
		log.Println("ID:", item.ID)
		log.Println("Name:", item.Name)
		log.Println("Region:", item.Region)
		log.Println("Country:", item.Country)
		log.Println("Lat:", item.Lat)
		log.Println("Lon:", item.Lon)
		log.Println("URL:", item.URL)
		log.Println()
	}

	// Conditions list
	conditions, err := a.Conditions()

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	for _, c := range conditions {
		log.Println("ID:", c.Code)
		log.Println("Day:", c.Day)
		log.Println("Night:", c.Night)
		log.Println("Icon:", c.Icon)
		log.Println()
	}
}
