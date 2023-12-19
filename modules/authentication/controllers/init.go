package controllers

import (
	"github.com/ortizdavid/golang-modular-software/config"
)


var userLogger = config.NewLogger("user.log")
var authLogger = config.NewLogger("autentication.log")