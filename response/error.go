package response

// Error offers error info received from the API
type Error struct {
	Error ErrorResponse `json:"error"`
}

// ErrorResponse provides error code and message for errors to be handled specifically
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
