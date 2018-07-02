package response

import "encoding/xml"

// Error offers error info received from the API
type Error struct {
	Error ErrorResponse `json:"error"`
}

// ErrorResponse provides error code and message for errors to be handled specifically
type ErrorResponse struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

// UnmarshalXML inserts the API error response into the error element
func (e *Error) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	el := &ErrorResponse{}
	if err = d.DecodeElement(&el, &start); err != nil {
		return
	}

	e.Error = *el

	return
}
