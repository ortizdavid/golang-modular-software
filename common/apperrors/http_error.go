package apperrors

import "net/http"

type HttpError struct {
	Message    string
	StatusCode int
}

func (err *HttpError) Error() string {
	return http.StatusText(err.StatusCode) + ": " + err.Message
}