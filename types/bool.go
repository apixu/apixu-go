package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// Bool is used to convert int 1/0 to bool true/false
type Bool bool

// UnmarshalJSON converts int to bool from JSON
func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	value, err := parseBool(num)
	if err != nil {
		return err
	}

	*b = Bool(value)

	return
}

// UnmarshalXML converts int to bool from XML
func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var el *int
	if err = d.DecodeElement(&el, &start); err != nil {
		return
	}

	if el == nil {
		return nil
	}

	value, err := parseBool(*el)
	if err != nil {
		return
	}

	*b = Bool(value)

	return
}

func parseBool(value int) (b bool, err error) {
	switch value {
	case 1:
		b = true
	case 0:
		b = false
	default:
		err = fmt.Errorf("invalid value for bool: %d", value)
	}

	return
}
