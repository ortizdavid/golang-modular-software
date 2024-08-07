package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/ortizdavid/go-nopain/conversion"
)

func NewZapLogger(logFileName string, logLevel zapcore.Level) *zap.Logger {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   LogRootPath() +"/"+logFileName,
		MaxSize:    LogMaxFileSize(),
		MaxBackups: LogMaxBackups(),
		MaxAge:     LogMaxAge(),
		Compress:   true,
	}
	// Create a zap core that writes logs to the lumberjack logger
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		logLevel,
	)
	// Create a logger with the zap core
	logger := zap.New(zapCore)
	return logger
}

func NewZapInfoLogger(logFileName string) *zap.Logger {
	return NewZapLogger(logFileName, zapcore.InfoLevel)
}

func NewZapErrorLogger(logFileName string) *zap.Logger {
	return NewZapLogger(logFileName, zapcore.ErrorLevel)
}

func NewZapDebugLogger(logFileName string) *zap.Logger {
	return NewZapLogger(logFileName, zapcore.DebugLevel)
}

func NewZapPanicLogger(logFileName string) *zap.Logger {
	return NewZapLogger(logFileName, zapcore.PanicLevel)
}

func LogRequestPath(ctx *fiber.Ctx) zap.Field {
	return zap.String("path", ctx.Path())
}

func LogRootPath() string {
	return GetEnv("LOG_ROOT_PATH")
}

func LogMaxFileSize() int {
	return conversion.StringToInt(GetEnv("LOG_MAX_SIZE"))
}

func LogMaxAge() int {
	return conversion.StringToInt(GetEnv("LOG_MAX_AGE"))
}

func LogMaxBackups() int {
	return conversion.StringToInt(GetEnv("LOG_MAX_BACKUPS"))
}


