package response

import (
	"encoding/xml"
	"io"
)

// Search defines the search response list
type Search []Location

// UnmarshalXML inserts the Location elements into the list
func (s *Search) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
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
