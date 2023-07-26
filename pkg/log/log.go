package log

import (
	"gas-detect/internal/conf"
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewStdLogger(hostname, name, version string, configZap *conf.Zap) log.Logger {
	cores := Zap.GetZapCores(configZap)
	zLogger := zap.New(zapcore.NewTee(cores...))
	logger := log.With(
		kzap.NewLogger(zLogger),
		"timestamp", log.DefaultTimestamp,
		"service", name,
		"caller", log.DefaultCaller,
		"host", hostname,
		"trace", tracing.TraceID(),
		"span", tracing.SpanID(),
	)
	return logger
}
