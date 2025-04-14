package wrapper

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

const slowSQLThreshold = 200 * time.Millisecond

type GormWrapper struct {
	logger *slog.Logger
	config gormlogger.Config
}

func NewGormWrapper(logger *slog.Logger) *GormWrapper {
	return &GormWrapper{
		logger: logger.With(
			slog.String("component", "gorm"),
		),
		config: gormlogger.Config{
			SlowThreshold:             slowSQLThreshold,
			LogLevel:                  gormlogger.Warn,
			IgnoreRecordNotFoundError: true,
		},
	}
}

func (w *GormWrapper) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	w.config.LogLevel = level

	return w
}

func (w *GormWrapper) Info(ctx context.Context, msg string, data ...interface{}) {
	w.logger.InfoContext(ctx, fmt.Sprintf(msg, data...))
}

func (w *GormWrapper) Warn(ctx context.Context, msg string, data ...interface{}) {
	w.logger.WarnContext(ctx, fmt.Sprintf(msg, data...))
}

func (w *GormWrapper) Error(ctx context.Context, msg string, data ...interface{}) {
	w.logger.ErrorContext(ctx, fmt.Sprintf(msg, data...))
}

func (w *GormWrapper) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	attrs := []slog.Attr{
		slog.Duration("elapsed", elapsed),
		slog.Group(
			"sql",
			slog.String("query", sql),
			slog.Int64("rows", rows),
		),
	}

	switch {
	case err != nil && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !w.config.IgnoreRecordNotFoundError):
		attrs = append(attrs, slog.Any("error", err))
		w.logger.LogAttrs(ctx, slog.LevelError, "database error occurred", attrs...)

	case elapsed > w.config.SlowThreshold && w.config.SlowThreshold != 0:
		attrs = append(attrs, slog.Duration("threshold", w.config.SlowThreshold))
		w.logger.LogAttrs(ctx, slog.LevelWarn, "slow SQL", attrs...)

	case w.config.LogLevel == gormlogger.Info:
		w.logger.LogAttrs(ctx, slog.LevelDebug, "sql tracing", attrs...)
	}
}
