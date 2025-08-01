package database

import (
	"fmt"
	"log/slog"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"

	"github.com/abgeo/follytics/internal/config"
	logwrapper "github.com/abgeo/follytics/internal/logger/wrapper"
	"github.com/abgeo/follytics/internal/telemetry"
)

func New(conf *config.Config, logger *slog.Logger, telemetry telemetry.Provider) (*gorm.DB, error) {
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

	if conf.Telemetry.Enabled {
		plugin := tracing.NewPlugin(
			tracing.WithDBSystem(conf.Database.Database),
			tracing.WithoutQueryVariables(),
			tracing.WithRecordStackTrace(),
			tracing.WithTracerProvider(telemetry.TracerProvider()),
		)
		if err = db.Use(plugin); err != nil {
			return nil, fmt.Errorf("failed to setup tracing plugin: %w", err)
		}
	}

	return db, nil
}
