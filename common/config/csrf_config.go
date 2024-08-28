package config

import (
	"strconv"

	"github.com/ortizdavid/go-nopain/conversion"
)

func CsrfExpiration() int {
	return conversion.StringToInt(GetEnv("CSRF_EXPIRATION"))
}

func CsrfCookieSecure() bool {
	if GetEnv("ENV") == "production" {
		return true
	}
	secure, err := strconv.ParseBool(GetEnv("CSRF_COOKIE_SECURE"))
	if err != nil {
		return false
	}
	return secure
}
