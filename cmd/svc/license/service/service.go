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
	"github.com/charmbracelet/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"

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

	log.Info("Creating License...",
		"id", req.Msg.GetId(),
		"user_id", req.Msg.GetUserId(),
		"start", req.Msg.GetStart(),
		"end", req.Msg.GetEnd(),
	)

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

// GetLicense - Gets a License.
func (s *Service) GetLicense(
	ctx context.Context,
	req *connect.Request[licensev1beta1.GetLicenseRequest],
) (*connect.Response[licensev1beta1.GetLicenseResponse], error) {
	record, err := models.FindLicense(ctx, s.db, req.Msg.GetId())
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve item: %w", err)
	}

	return connect.NewResponse(&licensev1beta1.GetLicenseResponse{
		License: &licensev1beta1.License{
			Id:     record.ID,
			Start:  timestamppb.New(record.StartTime),
			End:    timestamppb.New(record.EndTime),
			UserId: record.UserID,
		},
	}), nil
}

// ListLicenses - Lists licenses.
func (s *Service) ListLicenses(
	ctx context.Context,
	req *connect.Request[licensev1beta1.ListLicensesRequest],
) (*connect.Response[licensev1beta1.ListLicensesResponse], error) {
	var mods []qm.QueryMod
	// TODO account for parent and page_token fields
	mods = append(mods, qm.Limit(max(1, int(req.Msg.GetPageSize()))))

	items, err := models.Licenses(mods...).All(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve items: %w", err)
	}

	var licenses []*licensev1beta1.License
	for _, item := range items {
		licenses = append(licenses, &licensev1beta1.License{
			Id:     item.ID,
			Start:  timestamppb.New(item.StartTime),
			End:    timestamppb.New(item.EndTime),
			UserId: item.UserID,
		})
	}

	return connect.NewResponse(&licensev1beta1.ListLicensesResponse{
		Licenses: licenses,
	}), nil
}
