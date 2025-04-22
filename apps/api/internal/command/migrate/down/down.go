package down

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
			Use:   "down",
			Short: "Drop the migrations from database",
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
	reg, err := registry.NewBase(ctx, c.GetCmd().Flags())
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

	res, err := atlasExecutor.Down(ctx)
	if err != nil {
		return fmt.Errorf("failed to drop migrations: %w", err)
	}

	reg.GetLogger().
		WithGroup("migration").
		With(
			slog.Int("planned", len(res.Planned)),
			slog.Int("reverted", len(res.Reverted)),
			slog.String("current_version", res.Current),
			slog.String("target_version", res.Target),
			slog.Int("total", res.Total),
			slog.Time("start", res.Start),
			slog.Time("end", res.End),
			slog.String("url", res.URL),
			slog.String("status", res.Status),
			slog.String("error", res.Error),
		).
		InfoContext(ctx, "migrations have been dropped successfully")

	return nil
}
