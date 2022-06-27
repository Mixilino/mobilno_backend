package util

import "net/http"

type RestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewRestErrInternalServerError() *RestError {
	return &RestError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	}
}

func NewRestErrBadRequest(message string) *RestError {
	return &RestError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewRestErrConflict(message string) *RestError {
	return &RestError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

func NewRestErrUnauthorized() *RestError {
	return &RestError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
	}
}
