package logger

import (
	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	TimeKey             = "timestamp"
	LogEncoding         = flag.String("logenc", "json", "preferred logging format")
	LogOutputPaths      = flag.String("log", "stdout,access.log", "file paths to write logging output to")
	LogErrorOutputPaths = flag.String("logerror", "stderr,error.log", "path to write internal logger errors to.")
)

type Option func(*zap.Config)

func ConfigureTime(field, layout string) Option {
	return func(cfg *zap.Config) {
		cfg.EncoderConfig.TimeKey = field
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(layout)
	}
}

func ConfigureEncoding(encoding string) Option {
	return func(cfg *zap.Config) {
		cfg.Encoding = encoding
	}
}

func ConfigureOutput(paths []string) Option {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = paths
	}
}

func ConfigureErrorOutput(paths []string) Option {
	return func(cfg *zap.Config) {
		cfg.ErrorOutputPaths = paths
	}
}

func ConfigureTimeKey(key string) Option {
	return func(cfg *zap.Config) {
		cfg.EncoderConfig.TimeKey = key
	}
}

func NewConfig(opts ...Option) zap.Config {
	cfg := zap.NewProductionConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}
