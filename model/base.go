package model

type ErrorResponse struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
	Error   string `json:"error,omitempty"`
}
