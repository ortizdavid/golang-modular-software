package apperrors

import (
	"net/http"
)

func NotImplementedError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusNotImplemented,
	}
}
