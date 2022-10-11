package worker

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func NewController(license, org, profile *grpc.ClientConn) *saga.Controller {
	return saga.NewController(license, org, profile)
}

func NewConnToLicense() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortLicense)
	return grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewConnToOrg() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortOrg)
	return grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewConnToProfile() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%d", servicePortProfile)
	return grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewWorker(lc fx.Lifecycle, c client.Client, ctrl *saga.Controller) worker.Worker {
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, saga.CreateLicenseTaskQueue, worker.Options{
		// worker.Start() only return errors on start, so we need to catch
		// errors during run
		OnFatalError: func(err error) {
			logrus.WithError(err).Error("Worker failed!")
		},
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the worker in a non-blocking fashion.
			return w.Start()
		},
		OnStop: func(ctx context.Context) error {
			w.Stop()
			return nil
		},
	})

	w.RegisterWorkflow(saga.TransferMoney)

	// RegisterActivity - register an activity function or a pointer to a
	// structure with the worker.
	// The activity struct is a structure with all its exported methods treated
	// as activities. The default name of each activity is the method name.
	w.RegisterActivity(ctrl)

	return w
}

func RegisterWorkflowAndActivities(w worker.Worker) {
}

// TODO dial grpc connections to upstream services;
//  use uber-go/fx;
//  create a service struct w/ 3 grpc connections;
//  register the struct itself which causes all its exported functions to become activities:
//  from the docs:
//  And activity struct is a structure with all its exported methods treated as activities.
