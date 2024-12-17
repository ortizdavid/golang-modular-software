package apperrors

import (
	"net/http"
)

func PaymentRequiredError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusPaymentRequired,
	}
}