package apixu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type jsonStruct struct {
	Key string `json:"key"`
}

func TestNewFormatterJson(t *testing.T) {
	formatter, err := NewFormatter("json")
	assert.IsType(t, &JsonFormatter{}, formatter)
	assert.Nil(t, err)
}

func TestNewFormatterXml(t *testing.T) {
	formatter, err := NewFormatter("xml")
	assert.IsType(t, &XmlFormatter{}, formatter)
	assert.Nil(t, err)
}

func TestNewFormatterWhenFormatIsUnknown(t *testing.T) {
	_, err := NewFormatter("txt")
	assert.NotNil(t, err)
}

func TestUnmarshalJson(t *testing.T) {
	formatter := &JsonFormatter{}

	jsonString := `{
		"key" : "value"
		}`

	j := jsonStruct{}

	expected := jsonStruct{
		Key: "value",
	}

	err := formatter.Unmarshal([]byte(jsonString), &j)

	assert.Equal(t, expected, j)
	assert.Nil(t, err)
}

func TestUnmarshalJsonWithInvalidInput(t *testing.T) {
	formatter := &JsonFormatter{}

	jsonString := `{
		invalid json
		}`

	j := jsonStruct{}

	err := formatter.Unmarshal([]byte(jsonString), &j)

	assert.NotNil(t, err)
}

type xmlStruct struct {
	Key string `xml:"key"`
}

func TestUnmarshalXml(t *testing.T) {
	formatter := &XmlFormatter{}

	xmlString := `<?xml version="1.0" encoding="utf-8"?>
	<root>
		<key>value</key>
	</root>`

	x := xmlStruct{}

	expected := xmlStruct{
		Key: "value",
	}

	err := formatter.Unmarshal([]byte(xmlString), &x)

	assert.Equal(t, expected, x)
	assert.Nil(t, err)
}

func TestUnmarshalXmlWithInvalidInput(t *testing.T) {
	formatter := &XmlFormatter{}

	xmlString := `<?xml version="1.0" encoding="utf-8"?>
	invalid xml`

	x := xmlStruct{}

	err := formatter.Unmarshal([]byte(xmlString), &x)

	assert.NotNil(t, err)
}
