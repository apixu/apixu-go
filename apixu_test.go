package apixu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/andreiavrammsd/apixu-go/response"
	"github.com/stretchr/testify/assert"
)

// TestNew
func TestNew(t *testing.T) {
	c := Config{
		Version: "1",
		Format:  "json",
		APIKey:  "apikey",
	}
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.NoError(t, err)
}

func TestNewWithMissingVersion(t *testing.T) {
	c := Config{}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithMissingAPIKey(t *testing.T) {
	c := Config{
		Version: "1",
	}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithUnknownFormat(t *testing.T) {
	c := Config{
		Version: "1",
		APIKey:  "apikey",
		Format:  "txt",
	}
	a, err := New(c)

	assert.Nil(t, a)
	assert.Error(t, err)
}

type httpClientMock struct {
}

var (
	httpClientResponse = &http.Response{
		StatusCode: 200,
		Body:       &bodyMock{},
	}
	httpClientError                  error
	httpClientResponseBodyCloseError error
)

func (*httpClientMock) Get(url string) (*http.Response, error) {
	return httpClientResponse, httpClientError
}

type bodyMock struct {
}

func (*bodyMock) Read(p []byte) (n int, err error) {
	return
}

func (*bodyMock) Close() error {
	return httpClientResponseBodyCloseError
}

type jsonFormatterMock struct {
}

func (*jsonFormatterMock) Unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, object)
}

func TestApixu_Conditions(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`[  
   		{  
      		"code":1000,
      		"day":"Sunny",
      		"night":"Clear",
      		"icon":113
   		},
   		{  
      		"code":1003,
      		"day":"Partly cloudy",
      		"night":"Partly cloudy",
      		"icon":116
   		}
	]`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Conditions{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Conditions()
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_Current(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
	   "location":{  
	      "name":"Amsterdam",
	      "region":"North Holland",
	      "country":"Netherlands",
	      "lat":52.37,
	      "lon":4.89,
	      "tz_id":"Europe/Amsterdam",
	      "localtime_epoch":1529746782,
	      "localtime":"2018-06-23 11:39"
	   },
	   "current":{  
	      "last_updated_epoch":1529746209,
	      "last_updated":"2018-06-23 11:30",
	      "temp_c":15.0,
	      "temp_f":59.0,
	      "is_day":1,
	      "condition":{  
		 "text":"Partly cloudy",
		 "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		 "code":1003
	      },
	      "wind_mph":9.4,
	      "wind_kph":15.1,
	      "wind_degree":320,
	      "wind_dir":"NW",
	      "pressure_mb":1027.0,
	      "pressure_in":30.8,
	      "precip_mm":0.1,
	      "precip_in":0.0,
	      "humidity":72,
	      "cloud":75,
	      "feelslike_c":14.1,
	      "feelslike_f":57.3,
	      "vis_km":10.0,
	      "vis_miles":6.0
	   }
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.CurrentWeather{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Current("query")
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_Forecast(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
	   "location":{  
	      "name":"Paris",
	      "region":"Ile-de-France",
	      "country":"France",
	      "lat":48.87,
	      "lon":2.33,
	      "tz_id":"Europe/Paris",
	      "localtime_epoch":1529782822,
	      "localtime":"2018-06-23 21:40"
	   },
	   "current":{  
	      "last_updated_epoch":1529782208,
	      "last_updated":"2018-06-23 21:30",
	      "temp_c":19.0,
	      "temp_f":66.2,
	      "is_day":1,
	      "condition":{  
		 "text":"Sunny",
		 "icon":"//cdn.apixu.com/weather/64x64/day/113.png",
		 "code":1000
	      },
	      "wind_mph":9.4,
	      "wind_kph":15.1,
	      "wind_degree":50,
	      "wind_dir":"NE",
	      "pressure_mb":1025.0,
	      "pressure_in":30.8,
	      "precip_mm":0.0,
	      "precip_in":0.0,
	      "humidity":49,
	      "cloud":0,
	      "feelslike_c":19.0,
	      "feelslike_f":66.2,
	      "vis_km":10.0,
	      "vis_miles":6.0
	   },
	   "forecast":{  
	      "forecastday":[  
		 {  
		    "date":"2018-06-23",
		    "date_epoch":1529712000,
		    "day":{  
		       "maxtemp_c":22.3,
		       "maxtemp_f":72.1,
		       "mintemp_c":14.8,
		       "mintemp_f":58.6,
		       "avgtemp_c":17.3,
		       "avgtemp_f":63.1,
		       "maxwind_mph":8.5,
		       "maxwind_kph":13.7,
		       "totalprecip_mm":0.0,
		       "totalprecip_in":0.0,
		       "avgvis_km":18.8,
		       "avgvis_miles":11.0,
		       "avghumidity":51.0,
		       "condition":{  
		          "text":"Partly cloudy",
		          "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		          "code":1003
		       },
		       "uv":7.5
		    },
		    "astro":{  
		       "sunrise":"05:48 AM",
		       "sunset":"09:58 PM",
		       "moonrise":"05:11 PM",
		       "moonset":"03:21 AM"
		    }
		 }
	      ]
	   }
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Forecast{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Forecast("query", 2)
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_Search(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`[  
	   {  
	      "id":3332210,
	      "name":"Amsterdam, North Holland, Netherlands",
	      "region":"North Holland",
	      "country":"Netherlands",
	      "lat":52.37,
	      "lon":4.89,
	      "url":"amsterdam-north-holland-netherlands"
	   },
	   {  
	      "id":3332149,
	      "name":"De Wallen, North Holland, Netherlands",
	      "region":"North Holland",
	      "country":"Netherlands",
	      "lat":52.37,
	      "lon":4.9,
	      "url":"de-wallen-north-holland-netherlands"
	   }]`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.Search{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.Search("query")
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_History(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
	   "location":{  
	      "name":"Paris",
	      "region":"Ile-de-France",
	      "country":"France",
	      "lat":48.87,
	      "lon":2.33,
	      "tz_id":"Europe/Paris",
	      "localtime_epoch":1529795407,
	      "localtime":"2018-06-24 1:10"
	   },
	   "forecast":{  
	      "forecastday":[  
		 {  
		    "date":"2018-06-22",
		    "date_epoch":1529625600,
		    "day":{  
		       "maxtemp_c":21.0,
		       "maxtemp_f":69.8,
		       "mintemp_c":14.9,
		       "mintemp_f":58.8,
		       "avgtemp_c":18.7,
		       "avgtemp_f":65.7,
		       "maxwind_mph":10.5,
		       "maxwind_kph":16.9,
		       "totalprecip_mm":0.0,
		       "totalprecip_in":0.0,
		       "avgvis_km":10.0,
		       "avgvis_miles":6.0,
		       "avghumidity":52.0,
		       "condition":{  
		          "text":"Partly cloudy",
		          "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		          "code":1003
		       },
		       "uv":0.0
		    },
		    "astro":{  
		       "sunrise":"05:47 AM",
		       "sunset":"09:58 PM",
		       "moonrise":"04:03 PM",
		       "moonset":"02:56 AM",
		       "moon_phase":"First Quarter",
		       "moon_illumination":"48"
		    },
		    "hour":[  
		       {  
		          "time_epoch":1529622000,
		          "time":"2018-06-22 00:00",
		          "temp_c":13.6,
		          "temp_f":56.5,
		          "is_day":0,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/night/116.png",
		             "code":1003
		          },
		          "wind_mph":10.5,
		          "wind_kph":16.9,
		          "wind_degree":30,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":61,
		          "cloud":1,
		          "feelslike_c":12.2,
		          "feelslike_f":54.0,
		          "windchill_c":12.2,
		          "windchill_f":54.0,
		          "heatindex_c":13.6,
		          "heatindex_f":56.5,
		          "dewpoint_c":6.2,
		          "dewpoint_f":43.2,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529625600,
		          "time":"2018-06-22 01:00",
		          "temp_c":13.7,
		          "temp_f":56.7,
		          "is_day":0,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/night/116.png",
		             "code":1003
		          },
		          "wind_mph":9.4,
		          "wind_kph":15.1,
		          "wind_degree":31,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":60,
		          "cloud":1,
		          "feelslike_c":12.5,
		          "feelslike_f":54.5,
		          "windchill_c":12.5,
		          "windchill_f":54.5,
		          "heatindex_c":13.7,
		          "heatindex_f":56.7,
		          "dewpoint_c":6.1,
		          "dewpoint_f":43.0,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529629200,
		          "time":"2018-06-22 02:00",
		          "temp_c":13.8,
		          "temp_f":56.8,
		          "is_day":0,
		          "condition":{  
		             "text":"Clear",
		             "icon":"//cdn.apixu.com/weather/64x64/night/113.png",
		             "code":1000
		          },
		          "wind_mph":8.3,
		          "wind_kph":13.3,
		          "wind_degree":32,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":60,
		          "cloud":0,
		          "feelslike_c":12.8,
		          "feelslike_f":55.0,
		          "windchill_c":12.8,
		          "windchill_f":55.0,
		          "heatindex_c":13.8,
		          "heatindex_f":56.8,
		          "dewpoint_c":6.1,
		          "dewpoint_f":42.9,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529632800,
		          "time":"2018-06-22 03:00",
		          "temp_c":13.9,
		          "temp_f":57.0,
		          "is_day":0,
		          "condition":{  
		             "text":"Clear",
		             "icon":"//cdn.apixu.com/weather/64x64/night/113.png",
		             "code":1000
		          },
		          "wind_mph":7.2,
		          "wind_kph":11.5,
		          "wind_degree":34,
		          "wind_dir":"NE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":59,
		          "cloud":0,
		          "feelslike_c":13.1,
		          "feelslike_f":55.6,
		          "windchill_c":13.1,
		          "windchill_f":55.6,
		          "heatindex_c":13.9,
		          "heatindex_f":57.0,
		          "dewpoint_c":6.0,
		          "dewpoint_f":42.8,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529636400,
		          "time":"2018-06-22 04:00",
		          "temp_c":14.2,
		          "temp_f":57.5,
		          "is_day":0,
		          "condition":{  
		             "text":"Clear",
		             "icon":"//cdn.apixu.com/weather/64x64/night/113.png",
		             "code":1000
		          },
		          "wind_mph":6.6,
		          "wind_kph":10.7,
		          "wind_degree":29,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":58,
		          "cloud":0,
		          "feelslike_c":13.5,
		          "feelslike_f":56.3,
		          "windchill_c":13.5,
		          "windchill_f":56.3,
		          "heatindex_c":14.2,
		          "heatindex_f":57.5,
		          "dewpoint_c":6.1,
		          "dewpoint_f":43.0,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529640000,
		          "time":"2018-06-22 05:00",
		          "temp_c":14.4,
		          "temp_f":58.0,
		          "is_day":0,
		          "condition":{  
		             "text":"Clear",
		             "icon":"//cdn.apixu.com/weather/64x64/night/113.png",
		             "code":1000
		          },
		          "wind_mph":6.1,
		          "wind_kph":9.8,
		          "wind_degree":25,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":58,
		          "cloud":1,
		          "feelslike_c":13.9,
		          "feelslike_f":57.0,
		          "windchill_c":13.9,
		          "windchill_f":57.0,
		          "heatindex_c":14.4,
		          "heatindex_f":58.0,
		          "dewpoint_c":6.2,
		          "dewpoint_f":43.2,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529643600,
		          "time":"2018-06-22 06:00",
		          "temp_c":14.7,
		          "temp_f":58.5,
		          "is_day":1,
		          "condition":{  
		             "text":"Sunny",
		             "icon":"//cdn.apixu.com/weather/64x64/day/113.png",
		             "code":1000
		          },
		          "wind_mph":5.6,
		          "wind_kph":9.0,
		          "wind_degree":20,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":57,
		          "cloud":1,
		          "feelslike_c":14.3,
		          "feelslike_f":57.7,
		          "windchill_c":14.3,
		          "windchill_f":57.7,
		          "heatindex_c":14.7,
		          "heatindex_f":58.5,
		          "dewpoint_c":6.3,
		          "dewpoint_f":43.3,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529647200,
		          "time":"2018-06-22 07:00",
		          "temp_c":15.5,
		          "temp_f":60.0,
		          "is_day":1,
		          "condition":{  
		             "text":"Sunny",
		             "icon":"//cdn.apixu.com/weather/64x64/day/113.png",
		             "code":1000
		          },
		          "wind_mph":6.0,
		          "wind_kph":9.7,
		          "wind_degree":19,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":55,
		          "cloud":2,
		          "feelslike_c":15.2,
		          "feelslike_f":59.4,
		          "windchill_c":15.2,
		          "windchill_f":59.4,
		          "heatindex_c":15.5,
		          "heatindex_f":60.0,
		          "dewpoint_c":6.4,
		          "dewpoint_f":43.5,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529650800,
		          "time":"2018-06-22 08:00",
		          "temp_c":16.4,
		          "temp_f":61.5,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":6.5,
		          "wind_kph":10.4,
		          "wind_degree":17,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":52,
		          "cloud":2,
		          "feelslike_c":16.2,
		          "feelslike_f":61.1,
		          "windchill_c":16.2,
		          "windchill_f":61.1,
		          "heatindex_c":16.4,
		          "heatindex_f":61.5,
		          "dewpoint_c":6.5,
		          "dewpoint_f":43.7,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529654400,
		          "time":"2018-06-22 09:00",
		          "temp_c":17.2,
		          "temp_f":63.0,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":6.9,
		          "wind_kph":11.2,
		          "wind_degree":16,
		          "wind_dir":"NNE",
		          "pressure_mb":1030.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":50,
		          "cloud":2,
		          "feelslike_c":17.1,
		          "feelslike_f":62.8,
		          "windchill_c":17.1,
		          "windchill_f":62.8,
		          "heatindex_c":17.2,
		          "heatindex_f":63.0,
		          "dewpoint_c":6.6,
		          "dewpoint_f":43.9,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529658000,
		          "time":"2018-06-22 10:00",
		          "temp_c":18.1,
		          "temp_f":64.6,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":7.4,
		          "wind_kph":11.9,
		          "wind_degree":15,
		          "wind_dir":"NNE",
		          "pressure_mb":1030.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":48,
		          "cloud":4,
		          "feelslike_c":18.0,
		          "feelslike_f":64.5,
		          "windchill_c":18.0,
		          "windchill_f":64.5,
		          "heatindex_c":18.1,
		          "heatindex_f":64.6,
		          "dewpoint_c":6.7,
		          "dewpoint_f":44.1,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529661600,
		          "time":"2018-06-22 11:00",
		          "temp_c":19.0,
		          "temp_f":66.2,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":7.8,
		          "wind_kph":12.6,
		          "wind_degree":14,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":45,
		          "cloud":6,
		          "feelslike_c":19.0,
		          "feelslike_f":66.1,
		          "windchill_c":19.0,
		          "windchill_f":66.1,
		          "heatindex_c":19.0,
		          "heatindex_f":66.2,
		          "dewpoint_c":6.8,
		          "dewpoint_f":44.2,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529665200,
		          "time":"2018-06-22 12:00",
		          "temp_c":19.9,
		          "temp_f":67.8,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":8.3,
		          "wind_kph":13.3,
		          "wind_degree":14,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":43,
		          "cloud":8,
		          "feelslike_c":19.9,
		          "feelslike_f":67.8,
		          "windchill_c":19.9,
		          "windchill_f":67.8,
		          "heatindex_c":19.9,
		          "heatindex_f":67.8,
		          "dewpoint_c":6.9,
		          "dewpoint_f":44.4,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529668800,
		          "time":"2018-06-22 13:00",
		          "temp_c":20.2,
		          "temp_f":68.4,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":8.8,
		          "wind_kph":14.2,
		          "wind_degree":13,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":43,
		          "cloud":11,
		          "feelslike_c":20.2,
		          "feelslike_f":68.4,
		          "windchill_c":20.2,
		          "windchill_f":68.4,
		          "heatindex_c":20.2,
		          "heatindex_f":68.4,
		          "dewpoint_c":7.3,
		          "dewpoint_f":45.1,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529672400,
		          "time":"2018-06-22 14:00",
		          "temp_c":20.6,
		          "temp_f":69.0,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":9.3,
		          "wind_kph":15.0,
		          "wind_degree":13,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":43,
		          "cloud":14,
		          "feelslike_c":20.6,
		          "feelslike_f":69.0,
		          "windchill_c":20.6,
		          "windchill_f":69.0,
		          "heatindex_c":20.6,
		          "heatindex_f":69.0,
		          "dewpoint_c":7.6,
		          "dewpoint_f":45.7,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529676000,
		          "time":"2018-06-22 15:00",
		          "temp_c":20.9,
		          "temp_f":69.6,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":9.8,
		          "wind_kph":15.8,
		          "wind_degree":13,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":44,
		          "cloud":17,
		          "feelslike_c":20.9,
		          "feelslike_f":69.6,
		          "windchill_c":20.9,
		          "windchill_f":69.6,
		          "heatindex_c":20.9,
		          "heatindex_f":69.6,
		          "dewpoint_c":8.0,
		          "dewpoint_f":46.4,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529679600,
		          "time":"2018-06-22 16:00",
		          "temp_c":20.8,
		          "temp_f":69.4,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":10.1,
		          "wind_kph":16.2,
		          "wind_degree":15,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":45,
		          "cloud":13,
		          "feelslike_c":20.8,
		          "feelslike_f":69.4,
		          "windchill_c":20.8,
		          "windchill_f":69.4,
		          "heatindex_c":20.8,
		          "heatindex_f":69.4,
		          "dewpoint_c":8.4,
		          "dewpoint_f":47.2,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529683200,
		          "time":"2018-06-22 17:00",
		          "temp_c":20.7,
		          "temp_f":69.3,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":10.3,
		          "wind_kph":16.6,
		          "wind_degree":18,
		          "wind_dir":"NNE",
		          "pressure_mb":1028.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":47,
		          "cloud":10,
		          "feelslike_c":20.7,
		          "feelslike_f":69.3,
		          "windchill_c":20.7,
		          "windchill_f":69.3,
		          "heatindex_c":20.7,
		          "heatindex_f":69.3,
		          "dewpoint_c":8.9,
		          "dewpoint_f":48.0,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529686800,
		          "time":"2018-06-22 18:00",
		          "temp_c":20.6,
		          "temp_f":69.1,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":10.5,
		          "wind_kph":16.9,
		          "wind_degree":20,
		          "wind_dir":"NNE",
		          "pressure_mb":1028.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":48,
		          "cloud":7,
		          "feelslike_c":20.6,
		          "feelslike_f":69.1,
		          "windchill_c":20.6,
		          "windchill_f":69.1,
		          "heatindex_c":20.6,
		          "heatindex_f":69.1,
		          "dewpoint_c":9.3,
		          "dewpoint_f":48.7,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529690400,
		          "time":"2018-06-22 19:00",
		          "temp_c":19.6,
		          "temp_f":67.2,
		          "is_day":1,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/day/116.png",
		             "code":1003
		          },
		          "wind_mph":10.4,
		          "wind_kph":16.7,
		          "wind_degree":22,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":50,
		          "cloud":5,
		          "feelslike_c":19.6,
		          "feelslike_f":67.2,
		          "windchill_c":19.6,
		          "windchill_f":67.2,
		          "heatindex_c":19.6,
		          "heatindex_f":67.2,
		          "dewpoint_c":8.9,
		          "dewpoint_f":48.0,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529694000,
		          "time":"2018-06-22 20:00",
		          "temp_c":18.5,
		          "temp_f":65.4,
		          "is_day":1,
		          "condition":{  
		             "text":"Sunny",
		             "icon":"//cdn.apixu.com/weather/64x64/day/113.png",
		             "code":1000
		          },
		          "wind_mph":10.2,
		          "wind_kph":16.4,
		          "wind_degree":23,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":52,
		          "cloud":3,
		          "feelslike_c":18.5,
		          "feelslike_f":65.4,
		          "windchill_c":18.5,
		          "windchill_f":65.4,
		          "heatindex_c":18.5,
		          "heatindex_f":65.4,
		          "dewpoint_c":8.4,
		          "dewpoint_f":47.2,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529697600,
		          "time":"2018-06-22 21:00",
		          "temp_c":17.5,
		          "temp_f":63.5,
		          "is_day":1,
		          "condition":{  
		             "text":"Sunny",
		             "icon":"//cdn.apixu.com/weather/64x64/day/113.png",
		             "code":1000
		          },
		          "wind_mph":10.1,
		          "wind_kph":16.2,
		          "wind_degree":25,
		          "wind_dir":"NNE",
		          "pressure_mb":1029.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":54,
		          "cloud":1,
		          "feelslike_c":17.5,
		          "feelslike_f":63.5,
		          "windchill_c":17.5,
		          "windchill_f":63.5,
		          "heatindex_c":17.5,
		          "heatindex_f":63.5,
		          "dewpoint_c":8.0,
		          "dewpoint_f":46.4,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529701200,
		          "time":"2018-06-22 22:00",
		          "temp_c":16.7,
		          "temp_f":62.1,
		          "is_day":0,
		          "condition":{  
		             "text":"Clear",
		             "icon":"//cdn.apixu.com/weather/64x64/night/113.png",
		             "code":1000
		          },
		          "wind_mph":9.6,
		          "wind_kph":15.5,
		          "wind_degree":27,
		          "wind_dir":"NNE",
		          "pressure_mb":1030.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":56,
		          "cloud":1,
		          "feelslike_c":16.5,
		          "feelslike_f":61.7,
		          "windchill_c":16.5,
		          "windchill_f":61.7,
		          "heatindex_c":16.7,
		          "heatindex_f":62.1,
		          "dewpoint_c":7.9,
		          "dewpoint_f":46.2,
		          "will_it_rain":0,
		          "chance_of_rain":"0",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       },
		       {  
		          "time_epoch":1529704800,
		          "time":"2018-06-22 23:00",
		          "temp_c":15.9,
		          "temp_f":60.6,
		          "is_day":0,
		          "condition":{  
		             "text":"Partly cloudy",
		             "icon":"//cdn.apixu.com/weather/64x64/night/116.png",
		             "code":1003
		          },
		          "wind_mph":9.2,
		          "wind_kph":14.8,
		          "wind_degree":28,
		          "wind_dir":"NNE",
		          "pressure_mb":1030.0,
		          "pressure_in":30.9,
		          "precip_mm":0.0,
		          "precip_in":0.0,
		          "humidity":59,
		          "cloud":1,
		          "feelslike_c":15.5,
		          "feelslike_f":59.9,
		          "windchill_c":15.5,
		          "windchill_f":59.9,
		          "heatindex_c":15.9,
		          "heatindex_f":60.6,
		          "dewpoint_c":7.7,
		          "dewpoint_f":45.9,
		          "will_it_rain":0,
		          "chance_of_rain":"1",
		          "will_it_snow":0,
		          "chance_of_snow":"0",
		          "vis_km":10.0,
		          "vis_miles":6.0
		       }
		    ]
		 }
	      ]
	   }
	}`)

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	expected := &response.History{}
	if err := f.Unmarshal(data, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	r, err := a.History("query", time.Time{})
	assert.Equal(t, *expected, r)
	assert.NoError(t, err)
}

func TestApixu_HttpClientGetError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = errors.New("error")

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return nil, nil
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_ReadResponseBodyError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return nil, errors.New("error")
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_CloseResponseBodyError(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientError = nil
	httpClientResponseBodyCloseError = errors.New("error")

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return []byte(""), nil
	}

	expected := response.Search{}

	r, err := a.Search("query")
	assert.Equal(t, expected, r)
	assert.Error(t, err)
}

func TestApixu_APIErrorResponse(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
   		"error":{  
      	"code":1005,
      	"message":"API URL is invalid."
   		}
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	r, err := a.Search("query")

	expected := response.Search{}
	assert.Equal(t, expected, r)

	assert.Error(t, err)
	assert.IsType(t, &Error{}, err)

	expectedErrorResponse := &response.Error{}
	if err := f.Unmarshal(data, expectedErrorResponse); err != nil {
		assert.Fail(t, err.Error())
	}
	expectedError := &Error{
		err: err.(*Error).err,
		res: expectedErrorResponse.Error,
	}
	assert.Equal(t, expectedError, err)
}

func TestApixu_APIMalformedErrorResponse(t *testing.T) {
	c := Config{}
	f := &jsonFormatterMock{}

	httpClientResponse.StatusCode = 400
	httpClientError = nil
	httpClientResponseBodyCloseError = nil

	a := &apixu{
		config:     c,
		httpClient: &httpClientMock{},
		formatter:  f,
	}

	data := []byte(`{  
   		invalid json
	}`)
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return data, nil
	}

	r, err := a.Search("query")

	expected := response.Search{}
	assert.Equal(t, expected, r)

	assert.Error(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{"1", "xml", "apikey"}
	r := request{"GET", "query"}

	expected := fmt.Sprintf(
		apiURL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.APIKey,
		r.query,
	)
	result := a.getAPIURL(r)

	assert.Equal(t, expected, result)
}
