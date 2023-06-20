// Package service implements this service's API handlers.
package service

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/bufbuild/protovalidate-go"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"

	temporalv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/temporal/v1beta1"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
)

// Service - An HTTP API for starting Temporal workflows.
type Service struct {
	client    client.Client
	validator *protovalidate.Validator
}

// NewService - Constructs a new service for starting Temporal workflows.
func NewService(c client.Client, validator *protovalidate.Validator) *Service {
	return &Service{
		client:    c,
		validator: validator,
	}
}

// CreateOnboardingWorkflow - A handler for starting a new Temporal workflow
// that results in the creation of a license and associated objects.
func (s *Service) CreateOnboardingWorkflow(
	ctx context.Context,
	req *connect.Request[temporalv1beta1.CreateOnboardingWorkflowRequest],
) (*connect.Response[temporalv1beta1.CreateOnboardingWorkflowResponse], error) {
	temporalClient := s.client

	// Validate the request
	err := validate(s.validator, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// The business identifier of the workflow execution
	workflowID := req.Msg.GetWorkflowOptions().GetWorkflowId()

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: saga.CreateLicenseTaskQueue,
	}

	args := saga.CreateLicenseInputArgs{
		OrgName:     req.Msg.GetOrg().GetName(),
		ProfileName: req.Msg.GetProfile().GetName(),
		LicenseName: req.Msg.GetLicense().GetName(),
	}

	workflow, err := temporalClient.ExecuteWorkflow(
		ctx,
		options,
		saga.CreateLicense,
		args,
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf("unable to execute Temporal workflow: %w", err))
	}

	printResults(args, workflow.GetID(), workflow.GetRunID())

	res := &temporalv1beta1.CreateOnboardingWorkflowResponse{}

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}

func printResults(args saga.CreateLicenseInputArgs, workflowID, runID string) {
	logrus.WithFields(logrus.Fields{
		"org_name":             args.OrgName,
		"profile_name":         args.ProfileName,
		"license_name":         args.LicenseName,
		"temporal.workflow_id": workflowID,
		"temporal.run_id":      runID,
	}).Info("Successfully completed Workflow")
}
