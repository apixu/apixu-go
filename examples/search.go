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

	q := "Bouscat"
	search, err := a.Search(q)

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

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
}
