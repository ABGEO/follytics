package wrapper

import (
	"log/slog"
	"regexp"
	"strings"
)

type GinWrapper struct {
	logger *slog.Logger
}

func NewGinWrapper(logger *slog.Logger) *GinWrapper {
	return &GinWrapper{
		logger: logger.With(
			slog.String("component", "gin"),
		),
	}
}

func (w *GinWrapper) Write(data []byte) (int, error) {
	message := strings.TrimSuffix(string(data), "\n")

	switch {
	case strings.HasPrefix(message, "[GIN-debug]"):
		w.logGinDebug(strings.TrimPrefix(message, "[GIN-debug] "))

	case strings.HasPrefix(message, "[GIN-error]"):
		//nolint:noctx
		w.logger.Error(strings.TrimPrefix(message, "[GIN-error] "))

	default:
		//nolint:noctx
		w.logger.Info(strings.TrimPrefix(message, "[GIN] "))
	}

	return len(data), nil
}

func (w *GinWrapper) logGinDebug(message string) {
	re := regexp.MustCompile(`^(\w+)\s+(\S+)\s+-->\s+(\S+)\s+\((\d+)\s+handlers\)$`)
	if matches := re.FindStringSubmatch(message); len(matches) > 0 {
		//nolint:noctx
		w.logger.Debug(
			"new handler added",
			slog.String("method", matches[1]),
			slog.String("route", matches[2]),
			slog.String("handler", matches[3]),
			slog.String("handlers_count", matches[4]),
		)

		return
	}

	//nolint:noctx
	w.logger.Debug(message)
}
