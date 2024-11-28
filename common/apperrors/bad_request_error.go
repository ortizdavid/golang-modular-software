package apperrors

import (
	"net/http"
)

func BadRequestError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}
