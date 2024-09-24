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
		var message string
		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("Field '%s' is required.", e.Field())
		case "min":
			message = fmt.Sprintf("Field '%s' must be at least %s characters long.", e.Field(), e.Param())
		case "max":
			message = fmt.Sprintf("Field '%s' must be at most %s characters long.", e.Field(), e.Param())
		case "len":
			message = fmt.Sprintf("Field '%s' must be exactly %s characters long.", e.Field(), e.Param())
		case "email":
			message = fmt.Sprintf("Field '%s' must be a valid email address.", e.Field())
		case "url":
			message = fmt.Sprintf("Field '%s' must be a valid URL.", e.Field())
		case "eq":
			message = fmt.Sprintf("Field '%s' must be equal to '%s'.", e.Field(), e.Param())
		case "ne":
			message = fmt.Sprintf("Field '%s' must not be equal to '%s'.", e.Field(), e.Param())
		case "in":
			message = fmt.Sprintf("Field '%s' must be one of: %s.", e.Field(), e.Param())
		case "notin":
			message = fmt.Sprintf("Field '%s' must not be one of: %s.", e.Field(), e.Param())
		case "uuid":
			message = fmt.Sprintf("Field '%s' must be a valid UUID.", e.Field())
		case "alpha":
			message = fmt.Sprintf("Field '%s' must contain only alphabetic characters.", e.Field())
		case "alphanum":
			message = fmt.Sprintf("Field '%s' must contain only alphanumeric characters.", e.Field())
		case "numeric":
			message = fmt.Sprintf("Field '%s' must contain only numeric characters.", e.Field())
		case "startswith":
			message = fmt.Sprintf("Field '%s' must start with '%s'.", e.Field(), e.Param())
		case "oneof":
			message = fmt.Sprintf("Field '%s' must be one of: %s.", e.Field(), e.Param())
		case "endswith":
			message = fmt.Sprintf("Field '%s' must end with '%s'.", e.Field(), e.Param())
		case "contains":
			message = fmt.Sprintf("Field '%s' must contain '%s'.", e.Field(), e.Param())
		case "excludes":
			message = fmt.Sprintf("Field '%s' must not contain '%s'.", e.Field(), e.Param())
		case "excludesrune":
			message = fmt.Sprintf("Field '%s' must not contain the rune '%s'.", e.Field(), e.Param())
		case "gt":
			message = fmt.Sprintf("Field '%s' must be greater than %s.", e.Field(), e.Param())
		case "gte":
			message = fmt.Sprintf("Field '%s' must be greater than or equal to %s.", e.Field(), e.Param())
		case "lt":
			message = fmt.Sprintf("Field '%s' must be less than %s.", e.Field(), e.Param())
		case "lte":
			message = fmt.Sprintf("Field '%s' must be less than or equal to %s.", e.Field(), e.Param())
		case "base64":
			message = fmt.Sprintf("Field '%s' must be a valid base64 string.", e.Field())
		case "iso8601":
			message = fmt.Sprintf("Field '%s' must be a valid ISO 8601 date.", e.Field())
		case "datetime":
			message = fmt.Sprintf("Field '%s' must be a valid datetime format.", e.Field())
		default:
			message = fmt.Sprintf("Field '%s' is invalid.", e.Field())
		}
		// Prepend a hyphen to each message
		messages = append(messages, fmt.Sprintf("\t\t- %s", message))
	}
	return fmt.Errorf("validation errors:\n%s", strings.Join(messages, "\n"))
}

// ValidatorStartsWith checks if the field starts with the specified prefix
func ValidatorStartsWith(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), fl.Param())
}
