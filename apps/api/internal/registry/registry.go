package registry

import (
	"fmt"
	"log/slog"

	"github.com/spf13/pflag"
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/config"
	"github.com/abgeo/follytics/internal/database"
	"github.com/abgeo/follytics/internal/logger"
	"github.com/abgeo/follytics/internal/repository"
)

type Registry interface {
	GetConfig() *config.Config
	GetDB() *gorm.DB
	GetTransactionManager() *repository.TransactionManager
	GetLogger() *slog.Logger
}

type base struct {
	config    *config.Config
	db        *gorm.DB
	txManager *repository.TransactionManager
	logger    *slog.Logger
}

var _ Registry = (*base)(nil)

func newBase(flags *pflag.FlagSet) (*base, error) {
	var err error

	reg := &base{}

	configPath, err := flags.GetString("config")
	if err != nil {
		return nil, fmt.Errorf("failed to get 'config' flag: %w", err)
	}

	reg.config, err = config.New(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", err)
	}

	reg.logger, err = logger.New(reg.GetConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	reg.db, err = database.New(reg.GetConfig(), reg.GetLogger())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	reg.txManager = repository.NewTransactionManager(reg.GetDB())

	return reg, nil
}

func (r *base) GetConfig() *config.Config {
	return r.config
}

func (r *base) GetDB() *gorm.DB {
	return r.db
}

func (r *base) GetTransactionManager() *repository.TransactionManager {
	return r.txManager
}

func (r *base) GetLogger() *slog.Logger {
	return r.logger
}
