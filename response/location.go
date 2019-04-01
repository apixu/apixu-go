package response

import "github.com/apixu/apixu-go/v2/types"

// Location provides info on the returned location
type Location struct {
	ID             int            `json:"id,omitempty" xml:"id,omitempty"`
	Name           string         `json:"name" xml:"name"`
	Region         string         `json:"region" xml:"region"`
	Country        string         `json:"country" xml:"country"`
	Lat            float64        `json:"lat" xml:"lat"`
	Lon            float64        `json:"lon" xml:"lon"`
	URL            string         `json:"url,omitempty" xml:"url,omitempty"`
	Timezone       string         `json:"tz_id,omitempty" xml:"tz_id,omitempty"`
	LocalTimeEpoch int            `json:"localtime_epoch,omitempty" xml:"localtime_epoch,omitempty"`
	LocalTime      types.DateTime `json:"localtime,omitempty" xml:"localtime,omitempty"`
}
