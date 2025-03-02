package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/abgeo/follytics/internal/config"
	"github.com/abgeo/follytics/internal/version"
)

func New(conf *config.Config) (*slog.Logger, error) {
	attrs := []slog.Attr{
		slog.String("env", conf.Env),
		slog.String("version", version.Version),
	}

	level, err := parseLogLevel(conf.Logger.Level)
	if err != nil {
		return nil, err
	}

	handler := getLogHandler(level, conf.Logger.Format, conf.Env)
	handler = handler.WithAttrs(attrs)

	return slog.New(handler), nil
}

func parseLogLevel(rawLevel string) (slog.Level, error) {
	var level slog.Level

	if err := level.UnmarshalText([]byte(rawLevel)); err != nil {
		return level, fmt.Errorf("failed to parse log level: %w", err)
	}

	return level, nil
}

func getLogHandler(level slog.Level, format string, env string) slog.Handler {
	commonOptions := &slog.HandlerOptions{
		Level:     level,
		AddSource: env == "dev",
	}

	switch format {
	case "text":
		return slog.NewTextHandler(os.Stdout, commonOptions)
	case "json":
		return slog.NewJSONHandler(os.Stdout, commonOptions)
	default:
		if format != "" {
			slog.Warn("unknown log format, defaulting to JSON", slog.Any("format", format))
		}

		return slog.NewJSONHandler(os.Stdout, commonOptions)
	}
}
