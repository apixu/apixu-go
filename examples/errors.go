package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/andreiavrammsd/apixu-go"
)

func main() {
	// Unknown response format
	config := apixu.Config{
		Version: "1",
		Format:  "txt",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	// Missing query
	config = apixu.Config{
		Version: "1",
		Format:  "json",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err = apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	_, err = a.Forecast("", 1)
	if err != nil {
		log.Println(err)
	}

	// Query too long
	config = apixu.Config{
		Version: "1",
		Format:  "json",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err = apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	_, err = a.Forecast(strings.Repeat("a", 257), 1)
	if err != nil {
		log.Println(err)
	}

	// API error
	config = apixu.Config{
		Version: "1",
		Format:  "xml",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err = apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	date, err := time.Parse("2006-01-02", "2005-01-01")
	if err != nil {
		log.Println(err)
	}

	_, err = a.History("London", date)
	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}
}
