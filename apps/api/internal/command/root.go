package command

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/abgeo/follytics/internal/commander"
	"github.com/abgeo/follytics/internal/version"
)

type Command struct {
	commander.DummyCommand

	cmd *cobra.Command
}

var _ commander.Commander = (*Command)(nil)

func New(
	serveCmd commander.Commander,
	workerCmd commander.Commander,
) (*Command, error) {
	com := &Command{
		cmd: &cobra.Command{
			Use:     "follytics",
			Version: fmt.Sprintf("%s (%s)\n%s", version.Version, version.BuildTime, version.GitCommit),
		},
	}

	commander.Init(
		com,
		commander.WithChildCommands(
			serveCmd,
			workerCmd,
		),
	)

	return com, nil
}

func (c *Command) GetCmd() *cobra.Command {
	return c.cmd
}

func (c *Command) RegisterFlags() {
	defaultConfig := filepath.Join(userHomeDir(), "follytics.yaml")

	c.cmd.PersistentFlags().
		StringP("config", "c", defaultConfig, "Path to config file")
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			return os.Getenv("USERPROFILE")
		}

		return home
	}

	return os.Getenv("HOME")
}
