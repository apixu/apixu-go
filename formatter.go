package apixu

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
)

type Formatter interface {
	Unmarshal(data []byte, object interface{}) error
}

type JsonFormatter struct {
}

func (f *JsonFormatter) Unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, &object)
}

type XmlFormatter struct {
}

func (f *XmlFormatter) Unmarshal(data []byte, object interface{}) error {
	return xml.Unmarshal(data, &object)
}

func NewFormatter(format string) (Formatter, error) {
	var formatter Formatter
	var err error

	err = nil

	switch format {
	case "json":
		formatter = &JsonFormatter{}
		break
	case "xml":
		formatter = &XmlFormatter{}
		break
	default:
		err = errors.New(fmt.Sprintf("Unknown format: %s", format))
		break
	}

	return formatter, err
}
