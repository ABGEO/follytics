package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Logger struct {
	Level  string `default:"info" validate:"oneof=debug info warn error"`
	Format string `default:"json" validate:"oneof=text json"`
}

type Server struct {
	ListenAddr     string   `default:"0.0.0.0"              mapstructure:"address" validate:"ip"`
	Port           string   `default:"8000"                 mapstructure:"port"    validate:"numeric"`
	TrustedProxies []string `mapstructure:"trusted_proxies"`
}

type CORSConfig struct {
	Enabled          bool          `default:"false" mapstructure:"enabled"           validate:"boolean"`
	AllowOrigins     []string      `default:"[*]"   mapstructure:"allow_origins"`
	AllowMethods     []string      `default:"[*]"   mapstructure:"allow_methods"`
	AllowHeaders     []string      `default:"[*]"   mapstructure:"allow_headers"`
	ExposeHeaders    []string      `default:"[*]"   mapstructure:"expose_headers"`
	AllowCredentials bool          `default:"false" mapstructure:"allow_credentials"`
	MaxAge           time.Duration `default:"12h"   mapstructure:"max_age"`
}

type Swagger struct {
	Enabled bool   `default:"false"    mapstructure:"enabled" validate:"boolean"`
	Path    string `default:"/swagger" mapstructure:"path"`
}

type Healthcheck struct {
	Enabled bool   `default:"true"     mapstructure:"enabled" validate:"boolean"`
	Path    string `default:"/healthz" mapstructure:"path"`
}

type Database struct {
	Host                   string
	Port                   string `validate:"numeric"`
	User                   string
	Password               string
	Database               string
	BatchSize              int  `default:"100"  mapstructure:"batch_size"               validate:"numeric"`
	SkipDefaultTransaction bool `default:"true" mapstructure:"skip_default_transaction" validate:"boolean"`
}

type DatabaseMigrator struct {
	MigrationsPath  string `default:"/var/migrations" mapstructure:"migrations_path"`
	AtlasBinaryPath string `default:"atlas"           mapstructure:"atlas_binary_path"`
}

type GitHub struct {
	AppClientID       string `mapstructure:"app_client_id"`
	AppPrivateKeyPath string `mapstructure:"app_private_key_path"`
	AppInstallationID int64  `mapstructure:"app_installation_id"`
	JWTExpiration     int    `default:"1"                         mapstructure:"jwt_expiration" validate:"gt=0,lte=10"`
}

type SyncFollowersJob struct {
	BatchSize      int `default:"10"  mapstructure:"batch_size"       validate:"gt=0"`
	GitHubPageSize int `default:"100" mapstructure:"github_page_size" validate:"gt=0,lte=100"`
}

type Worker struct {
	Job struct {
		SyncFollowers SyncFollowersJob `mapstructure:"sync_followers"`
	}
}

type Config struct {
	Env string `default:"prod" validate:"oneof=dev test prod"`

	Logger           Logger
	Server           Server
	CORS             CORSConfig
	Swagger          Swagger
	Healthcheck      Healthcheck
	Database         Database
	DatabaseMigrator DatabaseMigrator `mapstructure:"database_migrator"`
	GitHub           GitHub
	Worker           Worker
}

func New(configPath string) (*Config, error) {
	conf := new(Config)
	configFile := filepath.Base(configPath)

	viperInstance := viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")),
	)

	viperInstance.AddConfigPath(".")
	viperInstance.AddConfigPath(filepath.Dir(configPath))
	viperInstance.SetConfigName(strings.TrimSuffix(configFile, filepath.Ext(configFile)))
	viperInstance.SetConfigType("yaml")

	viperInstance.AllowEmptyEnv(true)
	viperInstance.AutomaticEnv()

	defaults.SetDefaults(conf)

	if err := viperInstance.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viperInstance.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return conf, nil
}
