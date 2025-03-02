package commander

import (
	"context"

	"github.com/spf13/cobra"

	domainErrors "github.com/abgeo/follytics/internal/domain/errors"
)

type Commander interface {
	GetCmd() *cobra.Command
	RegisterFlags()
	Validate(cmd *cobra.Command, args []string) error
	Run(ctx context.Context) error
}

type DummyCommand struct{}

var _ Commander = (*DummyCommand)(nil)

func (c DummyCommand) GetCmd() *cobra.Command {
	panic(domainErrors.ErrNotImplemented)
}

func (c DummyCommand) RegisterFlags() {}

func (c DummyCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (c DummyCommand) Run(_ context.Context) error {
	return domainErrors.ErrNotImplemented
}
