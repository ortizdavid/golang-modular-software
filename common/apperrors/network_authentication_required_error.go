package apperrors

import (
	"net/http"
)

func NetworkAuthenticationRequiredError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusNetworkAuthenticationRequired,
	}
}