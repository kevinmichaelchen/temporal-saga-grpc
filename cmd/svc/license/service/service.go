// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"fmt"

	licenseConnect "buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	licensev1beta1 "buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go/license/v1beta1"
	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/license/models"
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

	logrus.
		WithField("id", req.Msg.GetId()).
		WithField("user_id", req.Msg.GetUserId()).
		WithField("start", req.Msg.GetStart()).
		WithField("end", req.Msg.GetEnd()).
		Info("Creating License...")

	license := models.License{
		ID:        req.Msg.GetId(),
		StartTime: req.Msg.GetStart().AsTime(),
		EndTime:   req.Msg.GetEnd().AsTime(),
		UserID:    req.Msg.GetUserId(),
	}

	err = license.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("unable to insert record: %w", err)
	}

	res := &licensev1beta1.CreateLicenseResponse{}

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
