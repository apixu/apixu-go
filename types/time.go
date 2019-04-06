package types

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// DateTime is used to convert string represented time to time.Time format
type DateTime time.Time

const (
	dateMarshalFormat = "2006-01-02 15:04"
	null              = "null"
)

// MarshalJSON converts time to string representation
func (t *DateTime) MarshalJSON() ([]byte, error) {
	dt := formatDate(t)

	res := null
	if dt != nil {
		res = fmt.Sprintf(`"%s"`, *dt)
	}

	return []byte(res), nil
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

// MarshalXML converts time to string representation
func (t *DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	dt := formatDate(t)
	if dt == nil {
		return nil
	}

	return e.EncodeElement(*dt, start)
}

func parseTime(value string) (dt time.Time, err error) {
	value = strings.Trim(value, `"`)

	// dateLayouts of supported time formats
	dateLayouts := []string{
		"2006-01-02 15:04",
		"2006-01-02",
	}

	for _, l := range dateLayouts {
		if dt, err = time.Parse(l, value); err == nil {
			return
		}
	}

	return
}

func formatDate(value *DateTime) *string {
	if value == nil {
		return nil
	}

	dt := time.Time(*value)
	if dt.IsZero() {
		return nil
	}

	formatted := dt.Format(dateMarshalFormat)

	return &formatted
}
