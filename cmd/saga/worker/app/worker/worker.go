// Package worker provides an FX module for a Temporal worker.
package worker

import (
	"context"
	"fmt"
	"net/http"

	"buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	"buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	"buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
	"connectrpc.com/connect"
	"github.com/charmbracelet/log"
	"github.com/sethvargo/go-envconfig"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/fxmod/connect/interceptor"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal/workflow"
)

// Module - An FX module for a Temporal worker.
var Module = fx.Module("worker",
	fx.Options(interceptor.Module),
	fx.Provide(
		NewConfig,
		NewController,
		NewWorker,
		NewLicenseClient,
		NewOrgClient,
		NewProfileClient,
	),
	fx.Invoke(
		RegisterWorkflowAndActivities,
	),
)

// Config - Environment-based config.
type Config struct {
	PortAPILicense int `env:"PORT_API_LICENSE"`
	PortAPIOrg     int `env:"PORT_API_ORG"`
	PortAPIProfile int `env:"PORT_API_PROFILE"`
}

// NewConfig - Reads config from environment.
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to read config from environment: %w", err)
	}

	return &cfg, nil
}

// NewController - Returns a new controller for our Temporal workflow.
func NewController(
	licenseClient licensev1beta1connect.LicenseServiceClient,
	orgClient orgv1beta1connect.OrgServiceClient,
	profileClient profilev1beta1connect.ProfileServiceClient,
) *workflow.Controller {
	return workflow.NewController(licenseClient, orgClient, profileClient)
}

// NewLicenseClient - Returns a new Connect client for the License service.
func NewLicenseClient(
	cfg *Config,
	interceptors []connect.Interceptor,
) licensev1beta1connect.LicenseServiceClient {
	addr := fmt.Sprintf("http://localhost:%d", cfg.PortAPILicense)

	return licensev1beta1connect.NewLicenseServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(interceptors...),
	)
}

// NewOrgClient - Returns a new Connect client for the Org service.
func NewOrgClient(
	cfg *Config,
	interceptors []connect.Interceptor,
) orgv1beta1connect.OrgServiceClient {
	addr := fmt.Sprintf("http://localhost:%d", cfg.PortAPIOrg)

	return orgv1beta1connect.NewOrgServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(interceptors...),
	)
}

// NewProfileClient - Returns a new Connect client for the Profile service.
func NewProfileClient(
	cfg *Config,
	interceptors []connect.Interceptor,
) profilev1beta1connect.ProfileServiceClient {
	addr := fmt.Sprintf("http://localhost:%d", cfg.PortAPIProfile)

	return profilev1beta1connect.NewProfileServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(interceptors...),
	)
}

// NewWorker - Returns a new worker for our Temporal workflow.
func NewWorker(
	lifecycle fx.Lifecycle,
	c client.Client,
	activityController *workflow.Controller,
) (*worker.Worker, error) {
	// This worker hosts both Workflow and Activity functions
	temporalWorker := worker.New(c, workflow.TaskQueue, worker.Options{
		// worker.Start() only return errors on start, so we need to catch
		// errors during run
		OnFatalError: func(err error) {
			log.Error("Worker failed!", "err", err)
		},
	})

	temporalWorker.RegisterWorkflow(workflow.CreateLicense)

	// Per the RegisterActivity docs, the registered activity doesn't need to be
	// a function, it can also be a struct with receiver functions.
	temporalWorker.RegisterActivity(activityController)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start the worker in a non-blocking fashion.
			err := temporalWorker.Start()
			if err != nil {
				return fmt.Errorf("unable to start Temporal worker: %w", err)
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			temporalWorker.Stop()

			return nil
		},
	})

	return &temporalWorker, nil
}

// RegisterWorkflowAndActivities - Registers the workflow and activities onto
// the Temporal worker.
func RegisterWorkflowAndActivities(_ *worker.Worker) {
}
