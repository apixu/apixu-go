package types

import (
	"encoding/xml"
	"strings"
	"time"
)

// DateTime is used to convert string represented time to time.Time format
type DateTime time.Time

// layouts of supported time formats
var layouts = []string{
	"2006-01-02 15:04",
	"2006-01-02",
}

// UnmarshalJSON converts string represented time to time.Time from JSON
func (t *DateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "null" {
		return nil
	}

	dt, err := parseTime(str)
	if err != nil {
		return err
	}

	*t = DateTime(dt)
	return nil
}

// UnmarshalXML converts string represented time to time.Time from XML
func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var el *string

	if err := d.DecodeElement(&el, &start); err != nil {
		return err
	}

	if *el == "null" {
		return nil
	}

	dt, err := parseTime(*el)
	if err != nil {
		return err
	}

	*t = DateTime(dt)
	return nil
}

func parseTime(value string) (dt time.Time, err error) {
	value = strings.Trim(value, `"`)

	for _, l := range layouts {
		if dt, err = time.Parse(l, value); err == nil {
			return
		}
	}

	return
}
