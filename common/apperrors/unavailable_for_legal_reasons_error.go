package apperrors

import (
	"net/http"
)

func UnavailableForLegalReasonsError(message string) *HttpError {
	return &HttpError{
		Message: message,
		StatusCode: http.StatusUnavailableForLegalReasons,
	}
}