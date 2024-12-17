package apperrors

import (
	"net/http"
)

func ServiceUnavailableError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusServiceUnavailable,
	}
}
