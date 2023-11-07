package logging

import (
	"github.com/chattronics/chattronics/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

var stdLogger *zap.Logger

type Field = zap.Field

func InitializeStandardLogger() {
	if config.LogFolderPath == "" {
		log.Fatalf("LogFolderPath not set, can't initialize standard logger.")
	}

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.TimeKey = "ts"
	// cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.Encoding = "console"
	cfg.Level.SetLevel(zap.DebugLevel)
	cfg.DisableCaller = false

	stdLogsFile := config.LogFolderPath + "console.logs"

	f, err := os.OpenFile(stdLogsFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %s", err.Error())
	}
	defer f.Close()

	cfg.OutputPaths = []string{stdLogsFile}

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
