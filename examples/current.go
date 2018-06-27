package main

import (
	"encoding/json"
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

	q := "Amsterdam"
	current, err := a.Current(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	loc := current.Location
	fmt.Println("Location")
	fmt.Println("\tName:", loc.Name)
	fmt.Println("\tRegion:", loc.Region)
	fmt.Println("\tCountry:", loc.Country)
	fmt.Println("\tLat:", loc.Lat)
	fmt.Println("\tLon:", loc.Lon)
	fmt.Println("\tTimezone:", loc.Timezone)
	fmt.Println("\tLocaltimeEpoch:", loc.LocalTimeEpoch)
	fmt.Println("\tLocaltime:", time.Time(loc.LocalTime).String())

	curr := current.Current
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

	data, err := json.Marshal(current)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s", data)

	return
}
