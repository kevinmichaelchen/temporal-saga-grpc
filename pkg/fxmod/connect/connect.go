package connect

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/cors"
)

func CreateModule(opts *ModuleOptions) fx.Option {
	return fx.Module("grpc",
		fx.Provide(
			opts.HandlerProvider,
			func() *ModuleOptions {
				return opts
			},
			NewConfig,
			NewServer,
		),
		fx.Invoke(
			Register,
		),
	)
}

type HandlerOutput struct {
	Path    string
	Handler http.Handler
}

type ModuleOptions struct {
	HandlerProvider any
	Services        []string
}

type Config struct {
	ConnectConfig *NestedConfig `env:",prefix=GRPC_CONNECT_"`
}

type NestedConfig struct {
	Host string `env:"HOST,default=localhost"`
	Port int    `env:"PORT,required"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process(context.Background(), &cfg)

	return
}

func NewServer(lc fx.Lifecycle, cfg Config) *http.ServeMux {
	mux := http.NewServeMux()
	address := fmt.Sprintf("%s:%d", cfg.ConnectConfig.Host, cfg.ConnectConfig.Port)
	srv := &http.Server{
		Addr: address,
		// Use h2c, so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(
			cors.NewCORS().Handler(mux),
			&http2.Server{},
		),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logrus.WithError(err).Error("connect-go ListenAndServe failed")
				}
			}()
			logrus.WithField("address", address).Info("Listening for connect-go")

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

func Register(opts *ModuleOptions, mux *http.ServeMux, h HandlerOutput) {
	checker := grpchealth.NewStaticChecker(
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so we'd be able
		// to reference foov1beta1.FooService as opposed to foo.v1beta1.FooService.
		opts.Services...,
	)
	mux.Handle(grpchealth.NewHandler(checker))
	mux.Handle(h.Path, h.Handler)
}
