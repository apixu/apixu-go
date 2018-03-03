package apixu

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Formatter defines methods needed for the formatter
// used to convert the REST response to specified type
type Formatter interface {
	Unmarshal(data []byte, object interface{}) error
}

type jsonFormatter struct {
}

func (f *jsonFormatter) Unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, &object)
}

type xmlFormatter struct {
}

func (f *xmlFormatter) Unmarshal(data []byte, object interface{}) error {
	return xml.Unmarshal(data, &object)
}

// NewFormatter returns formatter instance based on
// specified format type (JSON or XML)
func NewFormatter(format string) (Formatter, error) {
	var  (
		formatter Formatter
		err error
	)

	switch format {
	case "json":
		formatter = &jsonFormatter{}
		break
	case "xml":
		formatter = &xmlFormatter{}
		break
	default:
		err = fmt.Errorf("Unknown format: %s", format)
		break
	}

	return formatter, err
}
