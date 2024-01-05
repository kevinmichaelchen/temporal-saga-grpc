// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"fmt"

	orgConnect "buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	orgPB "buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go/org/v1beta1"
	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/models"
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
	ctx context.Context,
	req *connect.Request[orgPB.CreateOrgRequest],
) (*connect.Response[orgPB.CreateOrgResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.Low)
	if err != nil {
		return nil, err
	}

	org := models.Org{
		ID:   req.Msg.GetId(),
		Name: null.StringFrom(req.Msg.GetName()),
	}

	err = org.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("unable to insert record: %w", err)
	}

	res := &orgPB.CreateOrgResponse{}

	logrus.
		WithField("id", req.Msg.GetId()).
		WithField("name", req.Msg.GetName()).
		Info("Creating Org")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
