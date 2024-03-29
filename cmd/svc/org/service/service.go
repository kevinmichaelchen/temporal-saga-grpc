// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	orgConnect "buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	orgPB "buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go/org/v1beta1"
	"connectrpc.com/connect"
	"github.com/charmbracelet/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/org/models"
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
	log.Info("Creating Org...",
		"id", req.Msg.GetId(),
		"name", req.Msg.GetName(),
	)

	org := models.Org{
		ID:   req.Msg.GetId(),
		Name: null.StringFrom(req.Msg.GetName()),
	}

	err := org.Upsert(ctx, s.db, true, []string{models.OrgColumns.ID}, boil.Infer(), boil.Infer())
	if err != nil {
		log.Error("Failed to create Org",
			"id", req.Msg.GetId(),
			"err", err,
		)

		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("unable to insert record: %w", err),
		)
	}

	log.Info("Successfully created Org.",
		"id", req.Msg.GetId(),
		"name", req.Msg.GetName(),
	)

	res := &orgPB.CreateOrgResponse{
		Org: &orgPB.Org{
			Id:   org.ID,
			Name: org.Name.String,
		},
	}

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}

// GetOrg - Gets an org.
func (s *Service) GetOrg(
	ctx context.Context,
	req *connect.Request[orgPB.GetOrgRequest],
) (*connect.Response[orgPB.GetOrgResponse], error) {
	log.Info("Retrieving Org...",
		"id", req.Msg.GetId(),
	)

	record, err := models.FindOrg(ctx, s.db, req.Msg.GetId())
	if err != nil {
		log.Error("Failed to retrieve Org",
			"id", req.Msg.GetId(),
			"err", err,
		)

		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(
				connect.CodeNotFound,
				fmt.Errorf("unable to retrieve item: %w", err),
			)
		}

		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("unable to retrieve item: %w", err),
		)
	}

	log.Info("Successfully retrieved Org...",
		"id", req.Msg.GetId(),
	)

	return connect.NewResponse(&orgPB.GetOrgResponse{
		Org: &orgPB.Org{
			Id:   record.ID,
			Name: record.Name.String,
		},
	}), nil
}

// ListOrgs - Lists orgs.
func (s *Service) ListOrgs(
	ctx context.Context,
	req *connect.Request[orgPB.ListOrgsRequest],
) (*connect.Response[orgPB.ListOrgsResponse], error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Limit(max(1, int(req.Msg.GetPageSize()))))

	items, err := models.Orgs(mods...).All(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(
				connect.CodeNotFound,
				fmt.Errorf("unable to retrieve items: %w", err),
			)
		}

		return nil, fmt.Errorf("unable to retrieve items: %w", err)
	}

	var orgs []*orgPB.Org
	for _, item := range items {
		orgs = append(orgs, &orgPB.Org{
			Id:   item.ID,
			Name: item.Name.String,
		})
	}

	return connect.NewResponse(&orgPB.ListOrgsResponse{
		Orgs: orgs,
	}), nil
}
