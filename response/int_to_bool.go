package response

import (
	"encoding/xml"
	"fmt"
)

// IntToBool is used to convert int 1/0 to bool true/false
type IntToBool bool

// UnmarshalJSON converts int to bool from JSON
func (b *IntToBool) UnmarshalJSON(data []byte) (err error) {
	str := string(data)

	switch str {
	case "1":
		*b = true
	case "0":
		*b = false
	default:
		err = fmt.Errorf("invalid value for bool: %s", str)
	}
	return
}

// UnmarshalXML converts int to bool from XML
func (b *IntToBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var el *int
	if err := d.DecodeElement(&el, &start); err != nil {
		return err
	}

	switch *el {
	case 1:
		*b = true
	case 0:
		*b = false
	default:
		err = fmt.Errorf("invalid value for bool: %d", *el)
	}
	return
}
