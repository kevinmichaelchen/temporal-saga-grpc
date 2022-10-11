package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	temporalv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/temporal/v1beta1"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
	"go.temporal.io/sdk/client"
	"log"
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

	options := client.StartWorkflowOptions{
		ID:        req.Msg.GetWorkflowOptions().GetWorkflowId(),
		TaskQueue: saga.CreateLicenseTaskQueue,
	}
	transferDetails := saga.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}
	we, err := c.ExecuteWorkflow(
		context.Background(),
		options,
		saga.TransferMoney,
		transferDetails)
	if err != nil {
		log.Fatalln("error starting TransferMoney workflow", err)
	}

	printResults(transferDetails, we.GetID(), we.GetRunID())

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

func printResults(transferDetails saga.TransferDetails, workflowID, runID string) {
	log.Printf(
		"\nTransfer of $%f from account %s to account %s is processing. ReferenceID: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}
