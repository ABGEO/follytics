package worker

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/commander"
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
			Use:   "worker",
			Short: "Run Worker",
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

func (c *Command) RegisterFlags() {
	c.cmd.Flags().
		StringArrayP("jobs", "j", []string{"all"}, "Jobs to run")
}

func (c *Command) Run(ctx context.Context) error {
	flags := c.GetCmd().Flags()

	jobNames, err := flags.GetStringArray("jobs")
	if err != nil {
		return fmt.Errorf("failed to get 'jobs' flag: %w", err)
	}

	reg, err := registry.NewWorker(ctx, flags)
	if err != nil {
		return fmt.Errorf("failed to register worker command: %w", err)
	}

	defer func(reg registry.WorkerRegistry) {
		if !reg.GetConfig().Telemetry.Enabled {
			return
		}

		if err = reg.GetTelemetry().Shutdown(ctx); err != nil {
			reg.GetLogger().
				ErrorContext(
					ctx,
					"failed to shutdown telemetry service",
					slog.Any("error", err),
				)
		}
	}(reg)

	reg.GetLogger().InfoContext(ctx, "starting worker")

	if err = reg.GetWorker().Process(ctx, jobNames); err != nil {
		return fmt.Errorf("failed to process: %w", err)
	}

	return nil
}
