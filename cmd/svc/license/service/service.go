// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"

	licenseConnect "buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	licensev1beta1 "buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go/license/v1beta1"
	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/sirupsen/logrus"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

var _ licenseConnect.LicenseServiceHandler = (*Service)(nil)

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

// CreateLicense - Creates a new license.
func (s *Service) CreateLicense(
	ctx context.Context,
	req *connect.Request[licensev1beta1.CreateLicenseRequest],
) (*connect.Response[licensev1beta1.CreateLicenseResponse], error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("unable to construct validator: %w", err),
		)
	}

	err = validator.Validate(req.Msg)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err = simulated.PossibleError(simulated.MediumHigh)
	if err != nil {
		return nil, err
	}

	license := models.License{
		ID:        0,
		StartTime: time.Time{},
		EndTime:   time.Time{},
		UserID:    0,
	}

	err = license.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("unable to insert record: %w", err)
	}

	res := &licensev1beta1.CreateLicenseResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating License")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
