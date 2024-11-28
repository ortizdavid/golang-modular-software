package apperrors

import (
	"net/http"
)

func NotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusNotFound,
	}
}

