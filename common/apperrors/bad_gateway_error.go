package apperrors

import (
	"net/http"
)

func BadGatewayError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusBadGateway,
	}
}
