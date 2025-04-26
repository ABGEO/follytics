package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	healthcheck "github.com/tavsec/gin-healthcheck"
	healthcheckChecks "github.com/tavsec/gin-healthcheck/checks"
	healthcheckConfig "github.com/tavsec/gin-healthcheck/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"gorm.io/gorm"

	_ "github.com/abgeo/follytics/api/openapi" // import OpenAPI Definition
	"github.com/abgeo/follytics/internal/config"
	logwrapper "github.com/abgeo/follytics/internal/logger/wrapper"
	"github.com/abgeo/follytics/internal/route"
	"github.com/abgeo/follytics/internal/telemetry"
)

func NewRest(
	logger *slog.Logger,
	conf *config.Config,
	db *gorm.DB,
	telemetry telemetry.Provider,
	routes []route.Registerer,
) (*http.Server, error) {
	gin.DefaultWriter = logwrapper.NewGinWrapper(logger)

	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if conf.Env == "test" {
		gin.SetMode(gin.TestMode)
	}

	engine := gin.New()
	engine.ContextWithFallback = true

	if err := engine.SetTrustedProxies(conf.Server.TrustedProxies); err != nil {
		return nil, fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	registerMiddlewares(engine, conf, telemetry, logger)
	route.RegisterRoutes(engine, routes...)

	if conf.Healthcheck.Enabled {
		if err := registerHealthcheck(engine, conf, db); err != nil {
			return nil, err
		}
	}

	if conf.Swagger.Enabled {
		engine.GET(conf.Swagger.Path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return &http.Server{
		Addr:              net.JoinHostPort(conf.Server.ListenAddr, conf.Server.Port),
		Handler:           engine,
		ReadHeaderTimeout: 0,
	}, nil
}

func registerMiddlewares(engine *gin.Engine, conf *config.Config, telemetry telemetry.Provider, logger *slog.Logger) {
	engine.Use(gin.Recovery())

	if !conf.Telemetry.Enabled {
		engine.Use(gin.Logger())
	}

	if conf.CORS.Enabled {
		engine.Use(cors.New(cors.Config{
			AllowOrigins:     conf.CORS.AllowOrigins,
			AllowMethods:     conf.CORS.AllowMethods,
			AllowHeaders:     conf.CORS.AllowHeaders,
			ExposeHeaders:    conf.CORS.ExposeHeaders,
			AllowCredentials: conf.CORS.AllowCredentials,
			MaxAge:           conf.CORS.MaxAge,
		}))
	}

	if conf.Telemetry.Enabled {
		ignoredPaths := []string{
			conf.Healthcheck.Path,
		}

		engine.Use(otelgin.Middleware(
			"api",
			otelgin.WithTracerProvider(telemetry.TracerProvider()),
			otelgin.WithGinFilter(func(ctx *gin.Context) bool {
				for _, part := range ignoredPaths {
					if strings.Contains(ctx.Request.URL.Path, part) {
						return false
					}
				}

				return true
			}),
		))

		engine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
			WithUserAgent: true,
			WithRequestID: true,
			WithSpanID:    true,
			WithTraceID:   true,
			Filters: []sloggin.Filter{
				sloggin.IgnorePathContains(ignoredPaths...),
			},
		}))
	}
}

func registerHealthcheck(engine *gin.Engine, conf *config.Config, db *gorm.DB) error {
	var checks []healthcheckChecks.Check

	healthcheckConf := healthcheckConfig.DefaultConfig()
	healthcheckConf.HealthPath = conf.Healthcheck.Path

	dbInstance, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	checks = append(checks, healthcheckChecks.SqlCheck{Sql: dbInstance})

	if err = healthcheck.New(engine, healthcheckConf, checks); err != nil {
		return fmt.Errorf("failed to register healthchecks: %w", err)
	}

	return nil
}
