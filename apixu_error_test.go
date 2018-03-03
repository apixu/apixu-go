package apixu

import (
	"errors"
	"testing"
	"github.com/andreiavrammsd/apixu-go/response"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	text := "Message (1000)"
	code := uint16(1000)
	message := "Message"

	err := Error{
		errors.New(text),
		response.ErrorResponse{
			Code:    code,
			Message: message,
		},
	}

	assert.Equal(t, text, err.Error())
	assert.Equal(t, err.Response().Code, code)
	assert.Equal(t, err.Response().Message, message)
}
