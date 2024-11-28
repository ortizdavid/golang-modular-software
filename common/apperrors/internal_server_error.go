package apperrors

import "net/http"

func InternalServerError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}
