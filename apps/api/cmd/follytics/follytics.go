package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/command"
	"github.com/abgeo/follytics/internal/command/migrate"
	"github.com/abgeo/follytics/internal/command/serve"
	"github.com/abgeo/follytics/internal/command/worker"
)

func main() {
	if err := execute(); err != nil {
		slog.Error("failed to execute command", slog.Any("error", err))
		os.Exit(1)
	}
}

func execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd, err := getRootCmd()
	if err != nil {
		return err
	}

	if err = cmd.ExecuteContext(ctx); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

func getRootCmd() (*cobra.Command, error) {
	migrateCmd, err := migrate.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migrate command: %w", err)
	}

	serveCmd, err := serve.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize serve command: %w", err)
	}

	workerCmd, err := worker.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize serve command: %w", err)
	}

	rootCmd, err := command.New(migrateCmd, serveCmd, workerCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize root command: %w", err)
	}

	return rootCmd.GetCmd(), nil
}
