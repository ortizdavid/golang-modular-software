package controllers

import (
	"github.com/ortizdavid/golang-modular-software/config"
)


var loggerUser = config.NewLogger("user.log")
var loggerAuth = config.NewLogger("autentication.log")