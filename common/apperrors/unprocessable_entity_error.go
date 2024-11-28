package apperrors

import (
	"net/http"
)

func UnprocessableEntityError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}
