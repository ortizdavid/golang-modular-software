package apperrors

import (
	"net/http"
)

func GatewayTimeoutError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusGatewayTimeout,
	}
}
