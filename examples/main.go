package main

import (
	"log"

	"github.com/andreiavrammsd/apixu-go"
)

func main() {
	config := apixu.Config{
		Version: "1",
		Format:  "json",
		APIKey:  "<your api key>",
	}

	a, err := apixu.New(config)
	if err != nil {
		e := err.(*apixu.Error)
		log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
	}

	// Current weather
	q := "Amsterdam"
	current, err := a.GetCurrent(q)

	if err != nil {
		e := err.(*apixu.Error)
		log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
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

	// Search
	q = "Bouscat"
	search, err := a.Search(q)

	if err != nil {
		e := err.(*apixu.Error)
		log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
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
	conditions, err := a.GetConditions()

	if err != nil {
		e := err.(*apixu.Error)
		log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
	}

	for _, c := range conditions {
		log.Println("ID:", c.Code)
		log.Println("Day:", c.Day)
		log.Println("Night:", c.Night)
		log.Println("Icon:", c.Icon)
		log.Println()
	}
}
