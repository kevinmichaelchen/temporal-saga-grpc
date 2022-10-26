package connect

import (
	"context"
	"errors"
	"fmt"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/saga/start/service"
	"github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/temporal/v1beta1/temporalv1beta1connect"
	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/cors"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

var Module = fx.Module("grpc",
	fx.Provide(
		NewConfig,
		NewServer,
	),
	fx.Invoke(
		RegisterServer,
	),
)

type Config struct {
	ConnectConfig *NestedConfig `env:",prefix=GRPC_CONNECT"`
}

type NestedConfig struct {
	Host string `env:"HOST,default=localhost"`
	Port int    `env:"PORT,default=8081"`
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
			return srv.Shutdown(ctx)
		},
	})
	return mux
}

func RegisterServer(mux *http.ServeMux, svc *service.Service) {
	// Register our Connect-Go server
	path, handler := temporalv1beta1connect.NewTemporalServiceHandler(
		svc,
		connect.WithInterceptors(pkgConnect.UnaryInterceptors()...),
	)
	checker := grpchealth.NewStaticChecker(
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so we'd be able
		// to reference foov1beta1.FooService as opposed to foo.v1beta1.FooService.
		"com.teachingstrategies.temporalv1beta1.TemporalService",
	)
	mux.Handle(grpchealth.NewHandler(checker))
	mux.Handle(path, handler)
}
