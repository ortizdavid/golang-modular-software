package apperrors

import "net/http"

func NewInternalServerError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}
