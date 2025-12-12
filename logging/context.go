package logging

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type contextKey string

const loggerKey contextKey = "logger"

func CtxWithLog(ctx context.Context, logger *log.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LogFromContext(ctx context.Context) *log.Entry {
	if logger, ok := ctx.Value(loggerKey).(*log.Entry); ok {
		return logger
	}
	return log.NewEntry(log.StandardLogger())
}

func WithFields(ctx context.Context, fields log.Fields) (context.Context, *log.Entry) {
	logger := LogFromContext(ctx)
	logger = logger.WithFields(fields)
	return CtxWithLog(ctx, logger), logger
}

func WithField(ctx context.Context, key string, value interface{}) (context.Context, *log.Entry) {
	logger := LogFromContext(ctx)
	logger = logger.WithField(key, value)
	return CtxWithLog(ctx, logger), logger
}
