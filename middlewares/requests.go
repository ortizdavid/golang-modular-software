package middlewares

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ortizdavid/golang-modular-software/config"
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


func limitRequestPerSecond(ctx *fiber.Ctx) fiber.Handler {
	return limiter.New(limiter.Config{
		Expiration: time.Duration(config.RequestsExpiration()) * time.Second,
		Max:      config.RequestsPerSecond(),
	})
}