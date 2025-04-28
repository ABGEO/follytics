package status

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/commander"
	"github.com/abgeo/follytics/internal/database/migrator/atlas/exec"
	"github.com/abgeo/follytics/internal/registry"
)

type Command struct {
	commander.DummyCommand

	cmd *cobra.Command
}

var _ commander.Commander = (*Command)(nil)

func New() (*Command, error) {
	com := &Command{
		cmd: &cobra.Command{
			Use:   "status",
			Short: "Get migrations status",
		},
	}

	commander.Init(
		com,
		commander.WithRunner(),
	)

	return com, nil
}

func (c *Command) GetCmd() *cobra.Command {
	return c.cmd
}

func (c *Command) Run(ctx context.Context) error {
	reg, err := registry.NewBase(ctx, "db-migrator", c.GetCmd().Flags())
	if err != nil {
		return fmt.Errorf("failed to register command: %w", err)
	}

	atlasExecutor, err := exec.New(reg.GetConfig())
	if err != nil {
		return fmt.Errorf("failed to initialize atlas executor: %w", err)
	}
	defer func(atlasExecutor *exec.Executor) {
		if err = atlasExecutor.Shutdown(); err != nil {
			reg.GetLogger().
				ErrorContext(
					ctx,
					"failed to shutdown atlas executor",
					slog.Any("error", err),
				)
		}
	}(atlasExecutor)

	res, err := atlasExecutor.Status(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve migration status: %w", err)
	}

	reg.GetLogger().
		WithGroup("migration").
		With(
			slog.Int("available", len(res.Available)),
			slog.Int("pending", len(res.Pending)),
			slog.Int("applied", len(res.Applied)),
			slog.String("current_version", res.Current),
			slog.String("next_version", res.Next),
			slog.Int("count", res.Count),
			slog.Int("total", res.Total),
			slog.String("status", res.Status),
			slog.String("error", res.Error),
			slog.String("sql", res.SQL),
		).
		InfoContext(ctx, "migrations status")

	return nil
}
