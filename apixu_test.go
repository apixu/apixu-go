package apixu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNew
func TestNew(t *testing.T) {
	c := Config{
		Version: "1",
		Format:  "json",
		APIKey:  "apikey",
	}
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.NoError(t, err)
}

func TestNewWithMissingVersion(t *testing.T) {
	c := Config{}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithMissingAPIKey(t *testing.T) {
	c := Config{
		Version: "1",
	}
	a, err := New(c)
	assert.Nil(t, a)
	assert.Error(t, err)
}

func TestNewWithUnknownFormat(t *testing.T) {
	c := Config{
		Version: "1",
		APIKey:  "apikey",
		Format:  "txt",
	}
	a, err := New(c)

	assert.Nil(t, a)
	assert.Error(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{"1", "xml", "apikey"}
	r := request{"GET", "query"}

	expected := fmt.Sprintf(
		apiURL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.APIKey,
		r.query,
	)
	result := a.getAPIURL(r)

	assert.Equal(t, expected, result)
}
