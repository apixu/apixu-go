package apixu

import . "github.com/andreiavrammsd/apixu-go/response"

// ApixuError adds error details based on the API error response
type ApixuError struct {
	err error
	res ErrorResponse
}

func (ar *ApixuError) Error() string {
	return ar.err.Error()
}

func (ar *ApixuError) Response() ErrorResponse {
	return ar.res
}
