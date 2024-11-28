package apperrors

import (
	"net/http"
)

func ForbiddenError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}
