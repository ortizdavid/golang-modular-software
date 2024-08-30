package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidatorFormatErrors formats validation errors into user-friendly messages
func ValidatorFormatErrors(errs validator.ValidationErrors) error {
	var messages []string
	for _, e := range errs {
		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("Field '%s' is required.", e.Field()))
		case "min":
			messages = append(messages, fmt.Sprintf("Field '%s' must be at least %s characters long.", e.Field(), e.Param()))
		case "max":
			messages = append(messages, fmt.Sprintf("Field '%s' must be at most %s characters long.", e.Field(), e.Param()))
		case "len":
			messages = append(messages, fmt.Sprintf("Field '%s' must be exactly %s characters long.", e.Field(), e.Param()))
		case "email":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid email address.", e.Field()))
		case "url":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid URL.", e.Field()))
		case "eq":
			messages = append(messages, fmt.Sprintf("Field '%s' must be equal to '%s'.", e.Field(), e.Param()))
		case "ne":
			messages = append(messages, fmt.Sprintf("Field '%s' must not be equal to '%s'.", e.Field(), e.Param()))
		case "in":
			messages = append(messages, fmt.Sprintf("Field '%s' must be one of: %s.", e.Field(), e.Param()))
		case "notin":
			messages = append(messages, fmt.Sprintf("Field '%s' must not be one of: %s.", e.Field(), e.Param()))
		case "uuid":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid UUID.", e.Field()))
		case "alpha":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only alphabetic characters.", e.Field()))
		case "alphanum":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only alphanumeric characters.", e.Field()))
		case "numeric":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only numeric characters.", e.Field()))
		case "startswith":
			messages = append(messages, fmt.Sprintf("Field '%s' must start with '%s'.", e.Field(), e.Param()))
		case "oneof":
			messages = append(messages, fmt.Sprintf("Field '%s' must be one of: %s.", e.Field(), e.Param()))
		case "endswith":
			messages = append(messages, fmt.Sprintf("Field '%s' must end with '%s'.", e.Field(), e.Param()))
		case "contains":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain '%s'.", e.Field(), e.Param()))
		case "excludes":
			messages = append(messages, fmt.Sprintf("Field '%s' must not contain '%s'.", e.Field(), e.Param()))
		case "excludesrune":
			messages = append(messages, fmt.Sprintf("Field '%s' must not contain the rune '%s'.", e.Field(), e.Param()))
		case "gt":
			messages = append(messages, fmt.Sprintf("Field '%s' must be greater than %s.", e.Field(), e.Param()))
		case "gte":
			messages = append(messages, fmt.Sprintf("Field '%s' must be greater than or equal to %s.", e.Field(), e.Param()))
		case "lt":
			messages = append(messages, fmt.Sprintf("Field '%s' must be less than %s.", e.Field(), e.Param()))
		case "lte":
			messages = append(messages, fmt.Sprintf("Field '%s' must be less than or equal to %s.", e.Field(), e.Param()))
		case "base64":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid base64 string.", e.Field()))
		case "iso8601":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid ISO 8601 date.", e.Field()))
		case "datetime":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid datetime format.", e.Field()))
		default:
			messages = append(messages, fmt.Sprintf("Field '%s' is invalid.", e.Field()))
		}
	}
	return fmt.Errorf("validation errors:\n%s", strings.Join(messages, "\n"))
}

// ValidatorStartsWith checks if the field starts with the specified prefix
func ValidatorStartsWith(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), fl.Param())
}
