package models

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorTitle   string `json:"error_title"`
	ErrorMessage string `json:"error_message"`
}
