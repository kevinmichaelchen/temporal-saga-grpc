// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"fmt"

	profileConnect "buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
	profilePB "buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go/profile/v1beta1"
	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/models"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

var _ profileConnect.ProfileServiceHandler = (*Service)(nil)

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

// CreateProfile - Creates a new user profile.
func (s *Service) CreateProfile(
	ctx context.Context,
	req *connect.Request[profilePB.CreateProfileRequest],
) (*connect.Response[profilePB.CreateProfileResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.MediumLow)
	if err != nil {
		return nil, err
	}

	profile := models.Profile{
		ID:       req.Msg.GetId(),
		FullName: null.StringFrom(req.Msg.GetFullName()),
		OrgID:    req.Msg.GetOrgId(),
	}

	err = profile.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("unable to insert record: %w", err)
	}

	res := &profilePB.CreateProfileResponse{}

	logrus.
		WithField("id", req.Msg.GetId()).
		WithField("org_id", req.Msg.GetOrgId()).
		WithField("name", req.Msg.GetFullName()).
		Info("Creating Profile")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
