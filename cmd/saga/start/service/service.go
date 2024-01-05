// Package service implements this service's API handlers.
package service

import (
	"context"
	"fmt"

	temporalConnect "buf.build/gen/go/kevinmichaelchen/temporalapis/connectrpc/go/temporal/v1beta1/temporalv1beta1connect"
	temporalPB "buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go/temporal/v1beta1"
	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/client"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
)

var _ temporalConnect.TemporalServiceHandler = (*Service)(nil)

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
	req *connect.Request[temporalPB.CreateOnboardingWorkflowRequest],
) (*connect.Response[temporalPB.CreateOnboardingWorkflowResponse], error) {
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

	orgID := uuid.New().String()
	profileID := uuid.New().String()

	args := saga.CreateLicenseInputArgs{
		Org: saga.Org{
			ID:   orgID,
			Name: req.Msg.GetOrg().GetName(),
		},
		Profile: saga.Profile{
			ID:       profileID,
			FullName: req.Msg.GetProfile().GetFullName(),
		},
		License: saga.License{
			ID:     uuid.New().String(),
			Start:  req.Msg.GetLicense().GetStart().AsTime(),
			End:    req.Msg.GetLicense().GetEnd().AsTime(),
			UserID: profileID,
		},
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

	res := &temporalPB.CreateOnboardingWorkflowResponse{
		WorkflowId: workflow.GetID(),
	}

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}

func printResults(args saga.CreateLicenseInputArgs, workflowID, runID string) {
	logrus.WithFields(logrus.Fields{
		"org_name":             args.Org.Name,
		"profile_name":         args.Profile.FullName,
		"temporal.workflow_id": workflowID,
		"temporal.run_id":      runID,
	}).Info("Successfully completed Workflow")
}
