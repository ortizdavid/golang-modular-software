package apperrors

import (
	"net/http"
)

func NewConflictError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusConflict,
	}
}
