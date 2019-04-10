package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/apixu/apixu-go/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type handlers struct {
	apixu apixu.Apixu
}

type badRequest struct {
	message string
}

func (b *badRequest) Error() string {
	return b.message
}

type weatherResponse struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
}

func (s *handlers) CurrentWeather(c echo.Context) error {
	query := strings.TrimSpace(c.QueryParam("q"))
	if query == "" {
		return &badRequest{"Query is missing"}
	}

	data, err := s.apixu.Current(query)
	if err != nil {
		return err
	}

	res := weatherResponse{
		Location:    data.Location.Name,
		Temperature: data.Current.TempCelsius,
	}

	return c.JSON(http.StatusOK, res)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	a, err := apixu.New(apixu.Config{
		APIKey: os.Getenv("APIXUKEY"),
	})
	if err != nil {
		log.Fatalf("error at Apixu: %s", err)
	}

	server := echo.New()
	server.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	server.HTTPErrorHandler = func(err error, c echo.Context) {
		switch e := err.(type) {
		case *badRequest:
			err = echo.NewHTTPError(http.StatusBadRequest, e.Error())
		case *apixu.Error:
			err = echo.NewHTTPError(http.StatusNotFound, e.Response().Message)
		}

		c.Echo().DefaultHTTPErrorHandler(err, c)
	}

	handlers := &handlers{
		apixu: a,
	}
	server.GET("/weather/current", handlers.CurrentWeather)

	log.Fatal(server.Start(":8855"))
}
