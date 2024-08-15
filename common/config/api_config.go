package config

import "github.com/ortizdavid/go-nopain/conversion"


// ---- JWT 
func JwtSecretKey() string {
	return GetEnv("JWT_SECRET_KEY")
}

func JwtExpiration() int {
	return conversion.StringToInt(GetEnv("JWT_EXPIRATION"))
}

// ---- Requests
func RequestsPerSecond() int {
	return conversion.StringToInt(GetEnv("REQUESTS_PER_SECONDS"))
}

func RequestsExpiration() int {
	return conversion.StringToInt(GetEnv("REQUESTS_EXPIRATION"))
}

