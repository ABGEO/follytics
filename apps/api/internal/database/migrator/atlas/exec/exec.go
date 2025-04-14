package exec

import (
	"context"
	"fmt"
	"net"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"

	"github.com/abgeo/follytics/internal/config"
)

type Exec interface {
	Apply(ctx context.Context) (*atlasexec.MigrateApply, error)
	Down(ctx context.Context) (*atlasexec.MigrateDown, error)
	Status(ctx context.Context) (*atlasexec.MigrateStatus, error)
	Shutdown() error
}

type Executor struct {
	config  *config.Config
	workdir *atlasexec.WorkingDir
	client  *atlasexec.Client
}

var _ Exec = (*Executor)(nil)

func New(conf *config.Config) (*Executor, error) {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS(conf.DatabaseMigrator.MigrationsPath),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize atlas working directory: %w", err)
	}

	client, err := atlasexec.NewClient(workdir.Path(), conf.DatabaseMigrator.AtlasBinaryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize atlas client: %w", err)
	}

	return &Executor{
		config:  conf,
		workdir: workdir,
		client:  client,
	}, nil
}

func (e Executor) Apply(ctx context.Context) (*atlasexec.MigrateApply, error) {
	opts := atlasexec.MigrateApplyParams{
		URL: e.getDatabaseURL(),
	}

	res, err := e.client.MigrateApply(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return res, nil
}

func (e Executor) Down(ctx context.Context) (*atlasexec.MigrateDown, error) {
	opts := atlasexec.MigrateDownParams{
		URL:    e.getDatabaseURL(),
		DevURL: "docker://postgres/17/dev?search_path=public",
	}

	res, err := e.client.MigrateDown(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to drop migrations: %w", err)
	}

	return res, nil
}

func (e Executor) Status(ctx context.Context) (*atlasexec.MigrateStatus, error) {
	opts := atlasexec.MigrateStatusParams{
		URL: e.getDatabaseURL(),
	}

	res, err := e.client.MigrateStatus(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get migrations status: %w", err)
	}

	return res, nil
}

func (e Executor) Shutdown() error {
	if err := e.workdir.Close(); err != nil {
		return fmt.Errorf("failed to close working directory: %w", err)
	}

	return nil
}

func (e Executor) getDatabaseURL() string {
	// @todo: compose params from config.
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?search_path=public&sslmode=disable",
		e.config.Database.User,
		e.config.Database.Password,
		net.JoinHostPort(e.config.Database.Host, e.config.Database.Port),
		e.config.Database.Database,
	)
}
