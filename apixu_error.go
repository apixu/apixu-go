package apixu

import "github.com/andreiavrammsd/apixu-go/response"

// Error adds error details based on the API error response
type Error struct {
	err error
	res response.ErrorResponse
}

func (ar *Error) Error() string {
	return ar.err.Error()
}

// Response provides extra error info
func (ar *Error) Response() response.ErrorResponse {
	return ar.res
}
