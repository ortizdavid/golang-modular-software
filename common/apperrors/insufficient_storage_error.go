package apperrors

import (
	"net/http"
)

func InsufficientStorageError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusInsufficientStorage,
	}
}
