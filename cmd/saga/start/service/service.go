package service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	temporalv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/temporal/v1beta1"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/baggage"
	"go.temporal.io/sdk/client"
)

type Service struct {
	client client.Client
}

func NewService(c client.Client) *Service {
	return &Service{client: c}
}

func (s *Service) CreateLicense(
	ctx context.Context,
	req *connect.Request[temporalv1beta1.CreateLicenseRequest],
) (*connect.Response[temporalv1beta1.CreateLicenseResponse], error) {
	c := s.client

	// The business identifier of the workflow execution
	workflowID := req.Msg.GetWorkflowOptions().GetWorkflowId()

	// Inject the workflow ID as OTel baggage, so it appears on all spans
	bgMem, err := baggage.NewMember("temporal.workflowID", workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to create otel baggage member: %w", err)
	}
	bg, err := baggage.New(bgMem)
	if err != nil {
		return nil, fmt.Errorf("failed to create otel baggage: %w", err)
	}
	ctx = baggage.ContextWithBaggage(ctx, bg)

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: saga.CreateLicenseTaskQueue,
	}
	args := saga.CreateLicenseInputArgs{
		OrgName:     req.Msg.GetOrg().GetName(),
		ProfileName: req.Msg.GetProfile().GetName(),
		LicenseName: req.Msg.GetLicense().GetName(),
	}
	we, err := c.ExecuteWorkflow(
		ctx,
		options,
		saga.CreateLicense,
		args)
	if err != nil {
		logrus.WithError(err).Fatal("error starting CreateLicense workflow")
	}

	printResults(args, we.GetID(), we.GetRunID())

	// TODO edge case - what happens if server crashes before response gets sent?
	//  or if client crashes before they receive a response?
	//  in that case, we might experience duplicate workflows.

	res := &temporalv1beta1.CreateLicenseResponse{}
	if err != nil {
		return nil, err
	}
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
