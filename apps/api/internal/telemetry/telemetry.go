package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/abgeo/follytics/internal/config"
)

type Provider interface {
	Resource() *resource.Resource
	TracerProvider() *sdktrace.TracerProvider
	MeterProvider() *sdkmetric.MeterProvider
	LoggerProvider() *sdklog.LoggerProvider
	Shutdown(ctx context.Context) error
}

type Telemetry struct {
	serviceName    string
	serviceVersion string

	config *config.Config

	grpcConnection *grpc.ClientConn

	resource *resource.Resource

	tracerProvider *sdktrace.TracerProvider
	meterProvider  *sdkmetric.MeterProvider
	loggerProvider *sdklog.LoggerProvider
}

var _ Provider = (*Telemetry)(nil)

const connectionTimeout = 10 * time.Second

func New(
	ctx context.Context,
	serviceName string,
	serviceVersion string,
	conf *config.Config,
) (*Telemetry, error) {
	var err error

	telemetry := &Telemetry{
		serviceName:    serviceName,
		serviceVersion: serviceVersion,

		config: conf,
	}

	telemetry.grpcConnection, err = initGRPCConnection(conf)
	if err != nil {
		return nil, err
	}

	telemetry.resource, err = initResource(serviceName, serviceVersion, conf.Env)
	if err != nil {
		return nil, err
	}

	telemetry.tracerProvider, err = initTracer(ctx, telemetry.grpcConnection, telemetry.resource)
	if err != nil {
		return nil, err
	}

	telemetry.meterProvider, err = initMeter(ctx, telemetry.grpcConnection, telemetry.resource)
	if err != nil {
		return nil, err
	}

	telemetry.loggerProvider, err = initLogger(ctx, telemetry.grpcConnection, telemetry.resource)
	if err != nil {
		return nil, err
	}

	initPropagation()

	return telemetry, nil
}

func (t *Telemetry) Resource() *resource.Resource {
	return t.resource
}

func (t *Telemetry) TracerProvider() *sdktrace.TracerProvider {
	return t.tracerProvider
}

func (t *Telemetry) MeterProvider() *sdkmetric.MeterProvider {
	return t.meterProvider
}

func (t *Telemetry) LoggerProvider() *sdklog.LoggerProvider {
	return t.loggerProvider
}

func (t *Telemetry) Shutdown(ctx context.Context) error {
	if err := t.meterProvider.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown Metrics Provider: %w", err)
	}

	if err := t.tracerProvider.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown Traces Provider: %w", err)
	}

	if err := t.loggerProvider.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown Log Provider: %w", err)
	}

	if err := t.grpcConnection.Close(); err != nil {
		return fmt.Errorf("failed to close gRPC Connection: %w", err)
	}

	return nil
}

func initGRPCConnection(conf *config.Config) (*grpc.ClientConn, error) {
	connection, err := grpc.NewClient(
		conf.Telemetry.CollectorURL,
		// @todo: setup from config.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return connection, nil
}

func withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, connectionTimeout)
}

func initResource(serviceName, serviceVersion, env string) (*resource.Resource, error) {
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
			semconv.DeploymentEnvironment(env),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to merge resources: %w", err)
	}

	return res, nil
}

func initTracer(
	ctx context.Context,
	connection *grpc.ClientConn,
	res *resource.Resource,
) (*sdktrace.TracerProvider, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(connection))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}

func initMeter(
	ctx context.Context,
	connection *grpc.ClientConn,
	res *resource.Resource,
) (*sdkmetric.MeterProvider, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(connection))
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider, nil
}

func initLogger(
	ctx context.Context,
	connection *grpc.ClientConn,
	res *resource.Resource,
) (*sdklog.LoggerProvider, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(connection))
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)),
	)

	return loggerProvider, nil
}

func initPropagation() {
	compositeTextPropagation := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(compositeTextPropagation)
}
