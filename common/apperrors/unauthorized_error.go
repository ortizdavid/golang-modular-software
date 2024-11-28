package apperrors

import (
	"net/http"
)

func UnauthorizedError(message string) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}
