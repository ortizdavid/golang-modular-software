package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"go.uber.org/zap"
)

type RequestLoggerMiddleware struct {
	logger * zap.Logger
}

func NewRequestLoggerMiddleware() *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{
		logger: config.NewZapLogger("requests.log", zap.DebugLevel),
	}
}

func (mid *RequestLoggerMiddleware) Handle(ctx *fiber.Ctx) error {
	mid.logger.Info("Request",
		zap.String("Method", ctx.Method()),
		zap.String("Path", ctx.Path()),
		zap.String("StatusCode", fmt.Sprintf("%d", ctx.Response().StatusCode())),
	)
	return ctx.Next()
}
