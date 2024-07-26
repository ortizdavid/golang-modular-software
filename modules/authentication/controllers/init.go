package controllers

import (
	"github.com/ortizdavid/golang-modular-software/common/config"
)

var (
	userInfoLogger = config.NewLogger("user-info.log")
	userErrorLogger = config.NewLogger("user-error.log")
	authInfoLogger = config.NewLogger("auth-info.log")
	authErrorLgger = config.NewLogger("auth-info.log")
)
