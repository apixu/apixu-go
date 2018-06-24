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
	days := 5
	forecast, err := a.Forecast(q, days)

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

	// History
	history, err := a.History("London", time.Now())
	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	fmt.Println("History")
	fmt.Println()

	loc = history.Location
	fmt.Println("\tLocation")
	fmt.Println("\t\tName:", loc.Name)
	fmt.Println("\t\tRegion:", loc.Region)
	fmt.Println("\t\tCountry:", loc.Country)
	fmt.Println("\t\tLat:", loc.Lat)
	fmt.Println("\t\tLon:", loc.Lon)
	fmt.Println("\t\tTimezone:", loc.Timezone)
	fmt.Println("\t\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\t\tLocaltime:", loc.LocalTime)

	fmt.Println("\n\tForecast Day")
	for _, fc := range history.Forecast.ForecastDay {
		fmt.Println("\t\tDate:", fc.Date)
		fmt.Println("\t\tDateEpoch:", fc.DateEpoch)
		fmt.Println("\t\tDay")
		fmt.Println("\t\t\tMaxTempCelsius:", fc.Day.MaxTempCelsius)
		fmt.Println("\t\t\tMaxTempFahrenheit:", fc.Day.MaxTempFahrenheit)
		fmt.Println("\t\t\tMinTempCelsius:", fc.Day.MinTempCelsius)
		fmt.Println("\t\t\tMinTempFahrenheit:", fc.Day.MinTempFahrenheit)
		fmt.Println("\t\t\tAvgTempCelsius:", fc.Day.AvgTempCelsius)
		fmt.Println("\t\t\tAvgTempFahrenheit:", fc.Day.AvgTempFahrenheit)
		fmt.Println("\t\t\tMaxWindMPH:", fc.Day.MaxWindMPH)
		fmt.Println("\t\t\tMaxWindKMH:", fc.Day.MaxWindKMH)
		fmt.Println("\t\t\tTotalPrecipMM:", fc.Day.TotalPrecipMM)
		fmt.Println("\t\t\tTotalPrecipIN:", fc.Day.TotalPrecipIN)
		fmt.Println("\t\t\tVisKM:", fc.Day.AvgVisKM)
		fmt.Println("\t\t\tVisMiles:", fc.Day.AvgVisMiles)
		fmt.Println("\t\t\tAvgHumidity:", fc.Day.AvgHumidity)
		fmt.Println("\t\t\tCondition")
		fmt.Println("\t\t\t\tText:", fc.Day.Condition.Text)
		fmt.Println("\t\t\t\tIcon:", fc.Day.Condition.Icon)
		fmt.Println("\t\t\t\tCode:", fc.Day.Condition.Code)
		fmt.Println("\t\t\tUV:", fc.Day.UV)
		fmt.Println("\t\tAstro")
		fmt.Println("\t\t\tSunrise:", fc.Astro.Sunrise)
		fmt.Println("\t\t\tSunset:", fc.Astro.Sunset)
		fmt.Println("\t\t\tMoonrise:", fc.Astro.Moonrise)
		fmt.Println("\t\t\tMoonset:", fc.Astro.Moonset)

		fmt.Println("\t\tHour:")
		for _, fch := range fc.Hour {
			fmt.Println("\t\t\t\tTimeEpoch:", fch.TimeEpoch)
			fmt.Println("\t\t\t\tTime:", fch.Time)
			fmt.Println("\t\t\t\tTempCelsius:", fch.TempCelsius)
			fmt.Println("\t\t\t\tTempFahrenheit:", fch.TempFahrenheit)
			fmt.Println("\t\t\t\tIsDay:", fch.IsDay)
			fmt.Println("\t\t\t\tCondition")
			fmt.Println("\t\t\t\t\tText:", fch.Condition.Text)
			fmt.Println("\t\t\t\t\tIcon:", fch.Condition.Icon)
			fmt.Println("\t\t\t\t\tCode:", fch.Condition.Code)
			fmt.Println("\t\t\t\tWindMPH:", fch.WindMPH)
			fmt.Println("\t\t\t\tWindKMH:", fch.WindKMH)
			fmt.Println("\t\t\t\tWindDegree:", fch.WindDegree)
			fmt.Println("\t\t\t\tWindDirection:", fch.WindDirection)
			fmt.Println("\t\t\t\tPressureMB:", fch.PressureMB)
			fmt.Println("\t\t\t\tPressureIN:", fch.PressureIN)
			fmt.Println("\t\t\t\tPrecipMM:", fch.PrecipMM)
			fmt.Println("\t\t\t\tPrecipIN:", fch.PrecipIN)
			fmt.Println("\t\t\t\tHumidity:", fch.Humidity)
			fmt.Println("\t\t\t\tCloud:", fch.Cloud)
			fmt.Println("\t\t\t\tFeelsLikeCelsius:", fch.FeelsLikeCelsius)
			fmt.Println("\t\t\t\tFeelsLikeFahrenheit:", fch.FeelsLikeFahrenheit)
			fmt.Println("\t\t\t\tWindChillCelsius:", fch.WindChillCelsius)
			fmt.Println("\t\t\t\tWindChillFahrenheit:", fch.WindChillFahrenheit)
			fmt.Println("\t\t\t\tHeatIndexCelsius:", fch.HeatIndexCelsius)
			fmt.Println("\t\t\t\tHeatIndexFahrenheit:", fch.HeatIndexFahrenheit)
			fmt.Println("\t\t\t\tDewPointCelsius:", fch.DewPointCelsius)
			fmt.Println("\t\t\t\tDewPointFahrenheit:", fch.DewPointFahrenheit)
			fmt.Println("\t\t\t\tWillItRain:", fch.WillItRain)
			fmt.Println("\t\t\t\tChanceOfRain:", fch.ChanceOfRain)
			fmt.Println("\t\t\t\tWillItSnow:", fch.WillItSnow)
			fmt.Println("\t\t\t\tChanceOfSnow:", fch.ChanceOfSnow)
			fmt.Println("\t\t\t\tVisKM:", fch.VisKM)
			fmt.Println("\t\t\t\tVisMiles:", fch.VisMiles)
			fmt.Println()
		}

		fmt.Println()
	}

}
