package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level string) (*zap.Logger, error) {

	cfg := zap.NewProductionConfig()

	cfg.Encoding = "console"

	switch strings.ToLower(level) {

	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)

	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return cfg.Build()

}
