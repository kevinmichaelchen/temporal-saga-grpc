// Package otel provides an FX module for OpenTelemetry tracing.
package otel

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sethvargo/go-envconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.uber.org/fx"
)

// ModuleOptions - Options for configuring this OTel module.
type ModuleOptions struct {
	ServiceName string
}

// CreateModule - Creates an FX module for telemetry.
func CreateModule(opts ModuleOptions) fx.Option {
	return fx.Module("tracing",
		fx.Provide(
			func() *ModuleOptions {
				return &opts
			},
			NewSpanExporter,
			NewTracerProvider,
			NewConfig,
		),
		fx.Invoke(Register),
	)
}

// Config - Config for OTel tracing.
type Config struct {
	TraceConfig *TraceConfig `env:",prefix=TRACE_"`
}

// TraceConfig - Config for OTel tracing.
type TraceConfig struct {
	URL string `env:"URL,default=http://localhost:14268/api/traces"`
	Env string `env:"URL,default=local"`
}

// NewConfig - Reads config for OTel.
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to load environment-based config: %w", err)
	}

	return &cfg, nil
}

// Register - Registers the tracer provider.
func Register(tp *trace.TracerProvider) {
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Set the same Trace Propagator that Temporal uses by default
	otel.SetTextMapPropagator(opentelemetry.DefaultTextMapPropagator)
}

// NewSpanExporter - Returns a new span exporter.
func NewSpanExporter(
	cfg *Config,
) (trace.SpanExporter, error) {
	exp, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint(cfg.TraceConfig.URL),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create span exporter: %w", err)
	}

	return exp, nil
}

// NewTracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func NewTracerProvider(
	lifecycle fx.Lifecycle,
	opts *ModuleOptions,
	cfg *Config,
	exp trace.SpanExporter,
) (*trace.TracerProvider, error) {
	tracerProvider := trace.NewTracerProvider(
		// Always sample traces.
		trace.WithSampler(trace.AlwaysSample()),
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.ServiceName),
			attribute.String("environment", cfg.TraceConfig.Env),
		)),
	)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Do not make the application hang when it is shutdown.
			ctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()

			log.Println("Shutting down TracerProvider...")

			err := tracerProvider.Shutdown(ctx)
			if err != nil {
				return fmt.Errorf("unable to shut down tracer provider: %w", err)
			}

			log.Println("Successfully shut down TracerProvider")

			return nil
		},
	})

	return tracerProvider, nil
}
