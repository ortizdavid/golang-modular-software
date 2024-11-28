package apperrors

import (
	"net/http"
)

func ConflictError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusConflict,
	}
}
