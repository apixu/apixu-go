package apixu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNew
func TestNew(t *testing.T) {
	c := Config{}
	c.Format = "json"
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.Nil(t, err)
}

func TestNewWithError(t *testing.T) {
	c := Config{}
	c.Format = "uknown format"
	a, err := New(c)

	assert.Implements(t, (*Apixu)(nil), a)
	assert.NotNil(t, err)
}

// TestGetApiUrl
func TestGetApiUrl(t *testing.T) {
	a := &apixu{}
	a.config = Config{"1", "xml", "apikey"}
	r := request{"GET", "query"}

	expected := fmt.Sprintf(
		API_URL,
		a.config.Version,
		r.method,
		a.config.Format,
		a.config.ApiKey,
		r.query,
	)
	result := a.getApiUrl(r)

	assert.Equal(t, expected, result)
}
