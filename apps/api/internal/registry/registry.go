package registry

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/pflag"
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/config"
	"github.com/abgeo/follytics/internal/database"
	"github.com/abgeo/follytics/internal/logger"
	"github.com/abgeo/follytics/internal/repository"
	"github.com/abgeo/follytics/internal/telemetry"
	"github.com/abgeo/follytics/internal/version"
)

type Registry interface {
	GetConfig() *config.Config
	GetDB() *gorm.DB
	GetTelemetry() *telemetry.Telemetry
	GetTransactionManager() *repository.TransactionManager
	GetLogger() *slog.Logger
}

type Base struct {
	config    *config.Config
	db        *gorm.DB
	telemetry *telemetry.Telemetry
	txManager *repository.TransactionManager
	logger    *slog.Logger
}

var _ Registry = (*Base)(nil)

func NewBase(ctx context.Context, flags *pflag.FlagSet) (*Base, error) {
	var err error

	reg := &Base{}

	configPath, err := flags.GetString("config")
	if err != nil {
		return nil, fmt.Errorf("failed to get 'config' flag: %w", err)
	}

	reg.config, err = config.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", err)
	}

	if reg.GetConfig().Telemetry.Enabled {
		reg.telemetry, err = telemetry.New(
			ctx,
			"api",
			version.Version,
			reg.GetConfig(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize telemetry service: %w", err)
		}
	}

	reg.logger, err = logger.New(reg.GetConfig(), reg.GetTelemetry())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	reg.db, err = database.New(reg.GetConfig(), reg.GetLogger(), reg.GetTelemetry())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	reg.txManager = repository.NewTransactionManager(reg.GetDB())

	return reg, nil
}

func (r *Base) GetConfig() *config.Config {
	return r.config
}

func (r *Base) GetDB() *gorm.DB {
	return r.db
}

func (r *Base) GetTelemetry() *telemetry.Telemetry {
	return r.telemetry
}

func (r *Base) GetTransactionManager() *repository.TransactionManager {
	return r.txManager
}

func (r *Base) GetLogger() *slog.Logger {
	return r.logger
}
