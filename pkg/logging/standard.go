package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type Field = zap.Field

var stdLogger *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.Encoding = "console"
	cfg.Level.SetLevel(zap.DebugLevel)
	cfg.DisableCaller = false
	cfg.OutputPaths = []string{"stdout"}

	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Error to create a New L: %s", err.Error())
	}
	stdLogger = zapLogger
}

func Debug(msg string, additionalFields ...Field) {
	defer stdLogger.Sync()
	stdLogger.WithOptions(
		zap.AddCallerSkip(1),
	).Debug(msg, additionalFields...)
}

func Info(msg string, additionalFields ...Field) {
	defer stdLogger.Sync()
	stdLogger.WithOptions(
		zap.AddCallerSkip(1),
	).Info(msg, additionalFields...)
}

func Warn(msg string, additionalFields ...Field) {
	defer stdLogger.Sync()
	stdLogger.WithOptions(
		zap.AddCallerSkip(1),
	).Warn(msg, additionalFields...)
}

func Error(msg string, err error, additionalFields ...Field) {
	defer stdLogger.Sync()
	additionalFields = append(additionalFields, zap.Error(err))
	stdLogger.WithOptions(
		zap.AddCallerSkip(1),
	).Error(msg,
		additionalFields...,
	)
}

func Fatal(msg string, additionalFields ...Field) {
	defer stdLogger.Sync()
	stdLogger.WithOptions(
		zap.AddCallerSkip(1),
	).Fatal(msg, additionalFields...)
}

func AddField(key string, value interface{}) Field {
	return zap.Any(key, value)
}
