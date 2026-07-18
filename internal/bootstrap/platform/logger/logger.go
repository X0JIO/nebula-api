package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

func New(level string) (*zap.Logger, error) {

	cfg := zap.NewProductionConfig()

	switch level {

	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	}

	return cfg.Build()

}