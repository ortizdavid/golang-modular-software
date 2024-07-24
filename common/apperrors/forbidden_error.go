package apperrors

import (
	"net/http"
)

func NewForbiddenError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}
