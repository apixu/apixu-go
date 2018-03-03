package response

// Error offers extra error info
type Error struct {
	Error ErrorResponse `json:"error" xml:"error"`
}

// ErrorResponse provides error code and message
// for errors to be handled specifically
type ErrorResponse struct {
	Code    uint16 `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}
