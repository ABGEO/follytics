package commander

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init(command Commander, options ...func(command Commander)) {
	command.RegisterFlags()

	for _, option := range options {
		option(command)
	}
}

func WithRunner() func(command Commander) {
	return func(command Commander) {
		fn := func(cmd *cobra.Command, args []string) error {
			if err := command.Validate(cmd, args); err != nil {
				return fmt.Errorf("failed to validate command input: %w", err)
			}

			return command.Run(cmd.Context())
		}

		command.GetCmd().RunE = fn
	}
}

func WithChildCommands(cmds ...Commander) func(command Commander) {
	return func(command Commander) {
		for _, cmd := range cmds {
			command.GetCmd().AddCommand(cmd.GetCmd())
		}
	}
}
