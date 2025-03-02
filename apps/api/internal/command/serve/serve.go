package serve

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

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
			Use:   "serve",
			Short: "Run API server",
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
	// @todo: use these values.
	c.cmd.Flags().
		StringP("address", "a", "0.0.0.0", "Server address")
	c.cmd.Flags().
		StringP("port", "p", "8080", "Server port")
}

func (c *Command) Run(_ context.Context) error {
	reg, err := registry.NewServe(c.GetCmd().Flags())
	if err != nil {
		return fmt.Errorf("failed to register serve command: %w", err)
	}

	reg.GetLogger().Info(
		"starting HTTP server",
		slog.String("address", reg.GetConfig().Server.ListenAddr),
		slog.String("port", reg.GetConfig().Server.Port),
	)

	if err = reg.GetRestServer().ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start HTTP Server: %w", err)
	}

	return nil
}
