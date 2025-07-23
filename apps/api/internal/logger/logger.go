package logger

import (
	"fmt"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log"

	"github.com/abgeo/follytics/internal/config"
	"github.com/abgeo/follytics/internal/telemetry"
)

func New(conf *config.Config, telemetry telemetry.Provider) (*slog.Logger, error) {
	level, err := parseLogLevel(conf.Logger.Level)
	if err != nil {
		return nil, err
	}

	handler := getDefaultHandler(level, conf.Logger.Format, conf.Env)
	if conf.Telemetry.Enabled {
		handler = slogmulti.Fanout(
			handler,
			getTelemetryHandler(conf.Env, telemetry.LoggerProvider()),
		)
	}

	return slog.New(handler), nil
}

func parseLogLevel(rawLevel string) (slog.Level, error) {
	var level slog.Level

	if err := level.UnmarshalText([]byte(rawLevel)); err != nil {
		return level, fmt.Errorf("failed to parse log level: %w", err)
	}

	return level, nil
}

func getDefaultHandler(level slog.Level, format string, env string) slog.Handler {
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
			//nolint:noctx
			slog.Warn("unknown log format, defaulting to JSON", slog.Any("format", format))
		}

		return slog.NewJSONHandler(os.Stdout, commonOptions)
	}
}

func getTelemetryHandler(env string, provider log.LoggerProvider) *otelslog.Handler {
	return otelslog.NewHandler(
		"",
		otelslog.WithLoggerProvider(provider),
		otelslog.WithSource(env == "dev"),
	)
}
