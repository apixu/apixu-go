package main

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andreiavrammsd/apixu-go"
	"github.com/andreiavrammsd/apixu-go/response"
)

const idHashKey = "secretkey"

// MyApixuAPI defines an extension of Apixu API
type MyApixuAPI struct {
	apixu.Apixu
}

// MyConditions defines the custom Conditions list
type MyConditions []MyCondition

// MarshalXML handles formats XML string
func (c MyConditions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	res := struct {
		Items []MyCondition `xml:"condition"`
	}{Items: c}

	start.Name = xml.Name{
		Local: "MyConditionsList",
	}

	return e.EncodeElement(res, start)
}

// MyCondition defines the custom weather condition item
type MyCondition struct {
	response.Condition        // Embed main Conditions response
	ID                 string `json:"id" xml:"id"`                         // Add a new field
	Code               int    `json:"code,omitempty" xml:"code,omitempty"` // Overwrite existing field to hide it from xml and json representations
}

// Conditions retrieves the weather conditions list and adds extra information
func (a *MyApixuAPI) Conditions() (res MyConditions, err error) {
	conditions, err := a.Apixu.Conditions()

	for _, c := range conditions {
		c.Day = strings.ToUpper(c.Day)

		sha := sha256.New()
		sha.Write([]byte(fmt.Sprintf("%d.%s", c.Code, idHashKey)))
		id := fmt.Sprintf("%x", sha.Sum(nil))

		cond := MyCondition{
			Condition: c,
			ID:        id,
			Code:      c.Code,
		}
		res = append(res, cond)
	}
	return
}

func main() {
	config := apixu.Config{
		Version: "1",
		Format:  "xml",
		APIKey:  os.Getenv("APIXUKEY"),
	}

	a, err := apixu.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create custom API instance
	myApixuAPI := &MyApixuAPI{a}

	// Call overwritten method
	conditions, err := myApixuAPI.Conditions()

	if err != nil {
		if e, ok := err.(*apixu.Error); ok {
			log.Fatal(e.Error(), e.Response().Code, e.Response().Message)
		}
		log.Fatal(err)
	}

	b, err := xml.Marshal(conditions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", b)

	for _, c := range conditions {
		fmt.Println("\tCode:", c.Code)
		fmt.Println("\tDay:", c.Day)
		fmt.Println("\tNight:", c.Night)
		fmt.Println("\tIcon:", c.Icon)
		fmt.Println("\tID:", c.ID)
		fmt.Println()
	}

	// Call existing method
	q := "Kolno"
	search, err := myApixuAPI.Search(q)

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

	data, err := xml.Marshal(search)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s", data)
}
