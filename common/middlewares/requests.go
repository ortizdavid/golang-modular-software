package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"go.uber.org/zap"
)


var requestLogger = config.NewLogger("requests.log")


func requestLoggerMiddleware(ctx *fiber.Ctx) error {
	requestLogger.Info("Request",
		zap.String("Method", ctx.Method()),
		zap.String("Path", ctx.Path()),
		zap.String("StatusCode", fmt.Sprintf("%d", ctx.Response().StatusCode())),
	)
	return ctx.Next()
}
