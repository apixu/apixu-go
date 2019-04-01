package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apixu/apixu-go/v2"
)

func main() {
	config := apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := a.Conditions()

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	for _, c := range conditions {
		fmt.Println("\tCode:", c.Code)
		fmt.Println("\tDay:", c.Day)
		fmt.Println("\tNight:", c.Night)
		fmt.Println("\tIcon:", c.Icon)
		fmt.Println()
	}
}
