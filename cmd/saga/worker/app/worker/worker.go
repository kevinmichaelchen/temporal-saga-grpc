package worker

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
)

const (
	servicePortLicense = 9090
	servicePortOrg     = 9091
	servicePortProfile = 9092
)

var Module = fx.Module("worker",
	fx.Provide(
		fx.Annotate(
			NewController,
			fx.ParamTags(
				`name:"license"`,
				`name:"org"`,
				`name:"profile"`,
			),
		),
		NewWorker,
		fx.Annotate(
			NewConnToLicense,
			fx.ResultTags(`name:"license"`),
		),
		fx.Annotate(
			NewConnToOrg,
			fx.ResultTags(`name:"org"`),
		),
		fx.Annotate(
			NewConnToProfile,
			fx.ResultTags(`name:"profile"`),
		),
	),
	fx.Invoke(
		RegisterWorkflowAndActivities,
	),
)

// NewController - Returns a new controller for our Temporal workflow.
func NewController(license, org, profile *grpc.ClientConn) *saga.Controller {
	return saga.NewController(license, org, profile)
}

// NewConnToLicense - Returns a new gRPC connection to the License service.
func NewConnToLicense() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortLicense)

	return dial(addr)
}

// NewConnToOrg - Returns a new gRPC connection to the Org service.
func NewConnToOrg() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortOrg)

	return dial(addr)
}

// NewConnToProfile - Returns a new gRPC connection to the Profile service.
func NewConnToProfile() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortProfile)

	return dial(addr)
}

func dial(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to dial gRPC connection: %w", err)
	}

	return conn, nil
}

// NewWorker - Returns a new worker for our Temporal workflow.
func NewWorker(lc fx.Lifecycle, c client.Client, ctrl *saga.Controller) (*worker.Worker, error) {
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, saga.CreateLicenseTaskQueue, worker.Options{
		// worker.Start() only return errors on start, so we need to catch
		// errors during run
		OnFatalError: func(err error) {
			logrus.WithError(err).Error("Worker failed!")
		},
	})

	w.RegisterWorkflow(saga.CreateLicense)

	// RegisterActivity - register an activity function or a pointer to a
	// structure with the worker.
	// The activity struct is a structure with all its exported methods treated
	// as activities. The default name of each activity is the method name.
	w.RegisterActivity(ctrl)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the worker in a non-blocking fashion.
			err := w.Start()
			if err != nil {
				return fmt.Errorf("unable to start Temporal worker: %w", err)
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			w.Stop()

			return nil
		},
	})

	return &w, nil
}

func RegisterWorkflowAndActivities(w *worker.Worker) {
}
