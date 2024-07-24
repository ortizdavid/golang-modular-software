package apperrors

import (
	"net/http"
)

func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusNotFound,
	}
}

