package apperrors

import (
	"net/http"
)

func NewUnauthorizedError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}
