package service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/temporal/ctxpropagation"
	"github.com/sirupsen/logrus"
	temporalv1beta1 "go.buf.build/bufbuild/connect-go/kevinmichaelchen/temporalapis/temporal/v1beta1"
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

	ctx, err := injectBaggage(ctx, workflowID)
	if err != nil {
		return nil, err
	}

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

func injectBaggage(ctx context.Context, workflowID string) (context.Context, error) {
	// Inject the workflow ID as OTel baggage, so it appears on all spans
	bgMem, err := baggage.NewMember("foo", "bar")
	if err != nil {
		return nil, fmt.Errorf("failed to create otel baggage member: %w", err)
	}
	bg, err := baggage.New(bgMem)
	if err != nil {
		return nil, fmt.Errorf("failed to create otel baggage: %w", err)
	}

	// INJECT BAGGAGE!!
	//ctx = baggage.ContextWithBaggage(ctx, bg)
	ctx = context.WithValue(ctx, ctxpropagation.PropagateKey, bg)

	return ctx, nil
}
