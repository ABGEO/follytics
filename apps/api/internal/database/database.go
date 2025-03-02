package database

import (
	"fmt"
	"log/slog"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/abgeo/follytics/internal/config"
	logwrapper "github.com/abgeo/follytics/internal/logger/wrapper"
)

func New(conf *config.Config, logger *slog.Logger) (*gorm.DB, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		conf.Database.User,
		conf.Database.Password,
		net.JoinHostPort(conf.Database.Host, conf.Database.Port),
		conf.Database.Database,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger:                 logwrapper.NewGormWrapper(logger),
		CreateBatchSize:        conf.Database.BatchSize,
		SkipDefaultTransaction: conf.Database.SkipDefaultTransaction,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
