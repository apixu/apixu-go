package response

type Error struct {
	Error ErrorResponse `json:"error" xml:"error"`
}

type ErrorResponse struct {
	Code    uint16 `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}
