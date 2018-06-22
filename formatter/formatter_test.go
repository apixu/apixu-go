package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type jsonStruct struct {
	Key string `json:"key"`
}

func TestNewJsonFormatter(t *testing.T) {
	formatter, err := New("json")
	assert.IsType(t, &jsonFormatter{}, formatter)
	assert.Nil(t, err)
}

func TestNewXmlFormatter(t *testing.T) {
	formatter, err := New("xml")
	assert.IsType(t, &xmlFormatter{}, formatter)
	assert.Nil(t, err)
}

func TestNewWhenFormatIsUnknown(t *testing.T) {
	_, err := New("txt")
	assert.NotNil(t, err)
}

func TestUnmarshalJson(t *testing.T) {
	formatter := &jsonFormatter{}

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
	formatter := &jsonFormatter{}

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
	formatter := &xmlFormatter{}

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
	formatter := &xmlFormatter{}

	xmlString := `<?xml version="1.0" encoding="utf-8"?>
	invalid xml`

	x := xmlStruct{}

	err := formatter.Unmarshal([]byte(xmlString), &x)

	assert.NotNil(t, err)
}
