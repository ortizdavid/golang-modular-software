package apperrors

type HttpError struct {
	Message    string
	StatusCode int
}

func (err *HttpError) Error() string {
	return err.Message
}