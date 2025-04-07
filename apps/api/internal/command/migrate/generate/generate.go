package generate

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/commander"
	"github.com/abgeo/follytics/internal/database/migrator/atlas/loader"
)

type Command struct {
	commander.DummyCommand

	cmd *cobra.Command
}

var _ commander.Commander = (*Command)(nil)

func New() (*Command, error) {
	com := &Command{
		cmd: &cobra.Command{
			Use:   "generate",
			Short: "Generate SQL for the database from current models",
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

func (c *Command) Run(_ context.Context) error {
	schemaLoader := loader.NewSchemaLoader()

	statements, err := schemaLoader.Load()
	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	_, err = io.WriteString(os.Stdout, statements)
	if err != nil {
		return fmt.Errorf("failed to write schema to stdout: %w", err)
	}

	return nil
}
