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
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"

	pkgConnect "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
)

const (
	servicePortLicense = 9090
	servicePortOrg     = 9091
	servicePortProfile = 9092
)

// Module - An FX module for a Temporal worker.
var Module = fx.Module("worker",
	fx.Provide(
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

// NewController - Returns a new controller for our Temporal workflow.
func NewController(
	licenseClient licensev1beta1connect.LicenseServiceClient,
	orgClient orgv1beta1connect.OrgServiceClient,
	profileClient profilev1beta1connect.ProfileServiceClient,
) *saga.Controller {
	return saga.NewController(licenseClient, orgClient, profileClient)
}

// NewLicenseClient - Returns a new Connect client for the License service.
//
//nolint:ireturn // Connect client is an interface
func NewLicenseClient() licensev1beta1connect.LicenseServiceClient {
	addr := fmt.Sprintf("localhost:%d", servicePortLicense)

	return licensev1beta1connect.NewLicenseServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(pkgConnect.Interceptors()...),
	)
}

// NewOrgClient - Returns a new Connect client for the Org service.
//
//nolint:ireturn // Connect client is an interface
func NewOrgClient() orgv1beta1connect.OrgServiceClient {
	addr := fmt.Sprintf("localhost:%d", servicePortOrg)

	return orgv1beta1connect.NewOrgServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(pkgConnect.Interceptors()...),
	)
}

// NewProfileClient - Returns a new Connect client for the Profile service.
func NewProfileClient() profilev1beta1connect.ProfileServiceClient {
	addr := fmt.Sprintf("localhost:%d", servicePortProfile)

	return profilev1beta1connect.NewProfileServiceClient(
		http.DefaultClient,
		addr,
		connect.WithInterceptors(pkgConnect.Interceptors()...),
	)
}

// NewWorker - Returns a new worker for our Temporal workflow.
func NewWorker(
	lifecycle fx.Lifecycle,
	c client.Client,
	ctrl *saga.Controller,
) (*worker.Worker, error) {
	// This worker hosts both Workflow and Activity functions
	temporalWorker := worker.New(c, saga.CreateLicenseTaskQueue, worker.Options{
		// worker.Start() only return errors on start, so we need to catch
		// errors during run
		OnFatalError: func(err error) {
			logrus.WithError(err).Error("Worker failed!")
		},
	})

	// TODO can we move calls to RegisterWorkflow and RegisterActivity into an fx.Invoke block

	temporalWorker.RegisterWorkflow(saga.CreateLicense)

	// RegisterActivity - register an activity function or a pointer to a
	// structure with the worker.
	// The activity struct is a structure with all its exported methods treated
	// as activities. The default name of each activity is the method name.
	temporalWorker.RegisterActivity(ctrl)

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
