// +build integration

package apixu_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/apixu/apixu-go/v2"
	"github.com/xeipuuv/gojsonschema"
)

func TestApixu_Conditions(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := a.Conditions()
	if err != nil {
		t.Error(err)
	}

	validate(t, res, "conditions")
}

func TestApixu_Current(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := a.Current("London")
	if err != nil {
		t.Error(err)
	}

	validate(t, res, "current")
}

func TestApixu_Forecast(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := a.Forecast("London", 1, nil)
	if err != nil {
		t.Error(err)
	}

	validate(t, res, "forecast")
}

func TestApixu_Search(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := a.Search("London")
	if err != nil {
		t.Error(err)
	}

	validate(t, res, "search")
}

func TestApixu_History(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := a.History("London", time.Now().Add(time.Hour*-24), nil)
	if err != nil {
		t.Error(err)
	}

	validate(t, res, "history")
}

func TestApixu_Error(t *testing.T) {
	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = a.History(" . ", time.Now().Add(time.Hour*-24), nil)
	if err == nil {
		t.Fatal("error expected")
	}

	res := err.(*apixu.Error).Response()

	validate(t, res, "error")
}

func validate(t *testing.T, data interface{}, schemaFile string) {
	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://./testdata/schema/%s.json", schemaFile))
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		t.Fatal(err)
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(data))
	if err != nil {
		t.Fatal(err)
	}

	if !result.Valid() {
		t.Error(schemaFile, result.Errors())
	}
}
