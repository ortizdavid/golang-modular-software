package apperrors

import (
	"net/http"
)

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusBadRequest,
	}
}