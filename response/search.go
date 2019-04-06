package response

import (
	"encoding/xml"
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
