// Package service implements this service's API handlers.
package service

import (
	orgConnect "buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	orgPB "buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go/org/v1beta1"
	"connectrpc.com/connect"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

var _ orgConnect.OrgServiceHandler = (*Service)(nil)

// Service - A controller for our business logic.
type Service struct {
	db *sql.DB
}

// NewService - Returns a new Service.
func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// CreateOrg - Creates a new org.
func (s *Service) CreateOrg(
	_ context.Context,
	req *connect.Request[orgPB.CreateOrgRequest],
) (*connect.Response[orgPB.CreateOrgResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.Low)
	if err != nil {
		return nil, err
	}

	res := &orgPB.CreateOrgResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating Org")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
