package response

// Location provides info on the returned location
type Location struct {
	ID             uint32  `json:"id,omitempty" xml:"id,omitempty"`
	Name           string  `json:"name" xml:"name"`
	Region         string  `json:"region" xml:"region"`
	Country        string  `json:"country" xml:"country"`
	Lat            float32 `json:"lat" xml:"lat"`
	Lon            float32 `json:"lon" xml:"lon"`
	URL            string  `json:"url,omitempty" xml:"url,omitempty"`
	Timezone       string  `json:"tz_id,omitempty" xml:"tz_id,omitempty"`
	LocalTimeEpoch uint32  `json:"localtime_epoch,omitempty" xml:"localtime_epoch,omitempty"`
	LocalTime      string  `json:"localtime,omitempty" xml:"localtime,omitempty"`
}
