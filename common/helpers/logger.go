package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-modular-software/common/config"
	"go.uber.org/zap"
)

// Logger is a wrapper around zap.Logger for structured logging.
type Logger struct {
	logger *zap.Logger
}

type Field struct {
	Key   string
	Value interface{}
}

// NewInfoLogger creates a new Logger instance with Info level.
func NewInfoLogger(logFileName string) *Logger {
    return &Logger{logger: config.NewZapInfoLogger(logFileName)}
}

// NewErrorLogger creates a new Logger instance with Error level.
func NewErrorLogger(logFileName string) *Logger {
    return &Logger{logger: config.NewZapErrorLogger(logFileName)}
}

// NewDebugLogger creates a new Logger instance with Debug level.
func NewDebugLogger(logFileName string) *Logger {
    return &Logger{logger: config.NewZapDebugLogger(logFileName)}
}

// NewPanicLogger creates a new Logger instance with Panic level.
func NewPanicLogger(logFileName string) *Logger {
    return &Logger{logger: config.NewZapPanicLogger(logFileName)}
}

// Info logs an informational message with additional context.
func (l *Logger) Info(c *fiber.Ctx, message string, fields ...Field) {
	zapFields := convertFields(fields)
	zapFields = append(zapFields, zap.String("path", c.Path()))
	zapFields = append(zapFields, zap.String("method", c.Method()))
	zapFields = append(zapFields, zap.String("query", c.Context().QueryArgs().String()))
	l.logger.Info(message, zapFields...)
}

// Error logs an error message with additional context.
func (l *Logger) Error(c *fiber.Ctx, message string,  fields ...Field) {
	zapFields := convertFields(fields)
	zapFields = append(zapFields, zap.String("path", c.Path()))
	zapFields = append(zapFields, zap.String("method", c.Method()))
	zapFields = append(zapFields, zap.String("query", c.Context().QueryArgs().String()))
	l.logger.Error(message, zapFields...)
}

func convertFields(fields []Field) []zap.Field {
	var zapFields []zap.Field
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}