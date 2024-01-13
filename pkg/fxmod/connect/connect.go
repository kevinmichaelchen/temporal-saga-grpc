// Package connect provides a function to build an FX module for Connect Go
// APIs.
package connect

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/vanguard"
	"github.com/bufbuild/protovalidate-go"
	"github.com/charmbracelet/log"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/cors"
)

// CreateModule - The primary function for building an FX module for Connect Go
// APIs.
func CreateModule(opts *ModuleOptions) fx.Option {
	return fx.Module("connect",
		fx.Provide(
			opts.HandlerProvider,
			func() *ModuleOptions {
				return opts
			},
			NewConfig,
			NewServer,
			NewValidator,
			fx.Annotate(
				NewTranscoder,
				fx.ResultTags(`name:"vanguard"`),
			),
		),
		fx.Invoke(
			RegisterHealthAndReflection,
			fx.Annotate(
				RegisterVanguardTranscoder,
				fx.ParamTags(``, `name:"vanguard"`),
			),
			RegisterConnectHandler,
		),
	)
}

// NewValidator - Creates a new protobuf validator.
func NewValidator() (*protovalidate.Validator, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, fmt.Errorf("unable to create protobuf validator: %w", err)
	}

	return v, nil
}

// HandlerOutput - The result of creating a new Connect Go HTTP handler.
type HandlerOutput struct {
	// Path - The path on which to mount the handler.
	Path string
	// Handler - The HTTP handler for the Connect Go service implementation.
	Handler http.Handler
}

// ModuleOptions - Information that the consumer of this FX module should
// "provide" (in the FX sense).
type ModuleOptions struct {
	// HandlerProvider - Provides a Connect Go HTTP handler and the path to
	// mount it on.
	HandlerProvider any
	// Service - Fully-qualified protobuf service names.
	// (For example, "acme.user.v1.UserService"). Generated Connect service
	// files have this declared as a constant.
	Service string
}

// Config - Contains env vars for our Connect Go server.
type Config struct {
	ConnectConfig *NestedConfig `env:",prefix=GRPC_CONNECT_"`
}

// NestedConfig - Contains env vars for our Connect Go server.
type NestedConfig struct {
	Host string `env:"HOST,default=localhost"`
	Port int    `env:"PORT,required"`
}

// NewConfig - Reads configs from the environment.
func NewConfig() (Config, error) {
	var cfg Config

	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unable to read environment variables: %w", err)
	}

	return cfg, nil
}

// NewServer - Creates a new HTTP request multiplexer for our Connect Go APIs.
func NewServer(lifecycle fx.Lifecycle, cfg Config) *http.ServeMux {
	mux := http.NewServeMux()
	address := fmt.Sprintf("%s:%d", cfg.ConnectConfig.Host, cfg.ConnectConfig.Port)
	srv := &http.Server{
		Addr: address,
		// Use h2c, so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(
			cors.NewCORS().Handler(mux),
			&http2.Server{},
		),
		// The maximum duration for reading the entire request, including the
		// body.
		ReadTimeout: 1 * time.Second,
		// The maximum duration before timing out writes of the response.
		WriteTimeout: 1 * time.Second,
		// The maximum amount of time to wait for the next request when
		// keep-alive is enabled.
		IdleTimeout: 30 * time.Second,
		// The amount of time allowed to read request headers.
		ReadHeaderTimeout: 2 * time.Second,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error("connect-go ListenAndServe failed", "err", err)
				}
			}()
			log.Info("Listening for connect-go", "address", address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			if err != nil {
				return fmt.Errorf("unable to shut down HTTP server: %w", err)
			}

			return nil
		},
	})

	return mux
}

// NewTranscoder - Creates a transcoder from the REST protocol to the Connect
// protocol.
func NewTranscoder(
	opts *ModuleOptions,
	handlerOutput HandlerOutput,
) (http.Handler, error) {
	log.Info("Creating new Vanguard transcoder", "service", opts.Service, "handler", handlerOutput.Handler)
	handler, err := vanguard.NewTranscoder(
		[]*vanguard.Service{
			vanguard.NewService(
				opts.Service,
				handlerOutput.Handler,
			),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create connect transcoder: %w", err)
	}

	return handler, nil
}

// RegisterHealthAndReflection - Registers health and reflection handlers.
func RegisterHealthAndReflection(
	opts *ModuleOptions,
	mux *http.ServeMux,
) {
	checker := grpchealth.NewStaticChecker(opts.Service)
	mux.Handle(grpchealth.NewHandler(checker))

	reflector := grpcreflect.NewStaticReflector(opts.Service)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	// Many tools still expect the older version of the server reflection API,
	// so most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
}

// RegisterVanguardTranscoder - Registers Vanguard transcoder.
func RegisterVanguardTranscoder(
	mux *http.ServeMux,
	transcoderHandler http.Handler,
) {
	log.Info("Registering Vanguard transcoder", "handler", transcoderHandler)
	// Register the vanguard transcoder as broadly as possible
	mux.Handle("/", transcoderHandler)
}

// RegisterConnectHandler - Registers Connect handler.
func RegisterConnectHandler(
	mux *http.ServeMux,
	handlerOutput HandlerOutput,
) {
	log.Info("Registering Connect handler", "path", handlerOutput.Path, "handler", handlerOutput.Handler)
	// Register the Connect handler on its own path
	mux.Handle(handlerOutput.Path, handlerOutput.Handler)
}
