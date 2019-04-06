package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/apixu/apixu-go/v2"
)

func main() {
	// Missing query
	config := apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = a.Forecast("", 1, nil)
	if err != nil {
		log.Println(err)
	}

	// Query too long
	config = apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	}

	a, err = apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	_, err = a.Forecast(strings.Repeat("a", 257), 1, nil)
	if err != nil {
		log.Println(err)
	}

	// API error
	config = apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	}

	a, err = apixu.New(config)
	if err != nil {
		log.Println(err)
	}

	date, err := time.Parse("2006-01-02", "2005-01-01")
	if err != nil {
		log.Println(err)
	}

	_, err = a.History("London", date, nil)
	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}
}
