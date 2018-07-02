package response

import (
	"encoding/xml"
	"io"
)

// Search defines the search response list
type Search []Location

// MarshalXML converts Search response to XML
func (s Search) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	res := struct {
		Search []Location `xml:"location"`
	}{Search: s}

	start.Name = xml.Name{
		Local: "Search",
	}

	return e.EncodeElement(res, start)
}

// UnmarshalXML inserts the Location elements into the list
func (s *Search) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	locations := make(Search, 0)
	el := &Location{}

	for {
		err := d.Decode(el)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		locations = append(locations, *el)
	}

	*s = locations

	return nil
}
