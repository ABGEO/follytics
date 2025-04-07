package migrate

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/command/migrate/apply"
	"github.com/abgeo/follytics/internal/command/migrate/down"
	"github.com/abgeo/follytics/internal/command/migrate/generate"
	"github.com/abgeo/follytics/internal/command/migrate/status"
	"github.com/abgeo/follytics/internal/commander"
)

type Command struct {
	commander.DummyCommand

	cmd *cobra.Command
}

var _ commander.Commander = (*Command)(nil)

func New() (*Command, error) {
	com := &Command{
		cmd: &cobra.Command{
			Use:   "migrate",
			Short: "Commands to migrate the database",
		},
	}

	applyCmd, err := apply.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize apply command: %w", err)
	}

	downCmd, err := down.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize down command: %w", err)
	}

	generateCmd, err := generate.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize generate command: %w", err)
	}

	statusCmd, err := status.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize status command: %w", err)
	}

	commander.Init(
		com,
		commander.WithChildCommands(
			applyCmd,
			downCmd,
			generateCmd,
			statusCmd,
		),
	)

	return com, nil
}

func (c *Command) GetCmd() *cobra.Command {
	return c.cmd
}
