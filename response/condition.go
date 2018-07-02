package response

import (
	"encoding/xml"
	"io"
)

// Conditions defines Condition items list
type Conditions []Condition

// Condition defines the weather condition item
type Condition struct {
	Code  int    `json:"code" xml:"code"`
	Day   string `json:"day" xml:"day"`
	Night string `json:"night" xml:"night"`
	Icon  int    `json:"icon" xml:"icon"`
}

// MarshalXML converts Conditions list response to XML
func (c Conditions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	res := struct {
		Conditions []Condition `xml:"condition"`
	}{Conditions: c}

	start.Name = xml.Name{
		Local: "conditions",
	}

	return e.EncodeElement(res, start)
}

// UnmarshalXML inserts the Condition elements into the list
func (c *Conditions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	conditions := make(Conditions, 0)
	el := &Condition{}

	for {
		err := d.Decode(el)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		conditions = append(conditions, *el)
	}

	*c = conditions

	return nil
}
