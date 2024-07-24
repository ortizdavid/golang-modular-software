package apperrors

import (
	"net/http"
)

func NewUnprocessableEntityError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}
