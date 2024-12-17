package apperrors

import (
	"net/http"
)

func TooManyRequestsError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusTooManyRequests,
	}
}
