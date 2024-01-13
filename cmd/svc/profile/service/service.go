// Package service implements this service's API handlers.
package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"

	profileConnect "buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
	profilePB "buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go/profile/v1beta1"
	"connectrpc.com/connect"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/kevinmichaelchen/temporal-saga-grpc/cmd/svc/profile/models"
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
	log.Info("Creating Profile...",
		"id", req.Msg.GetId(),
		"org_id", req.Msg.GetOrgId(),
		"name", req.Msg.GetFullName(),
	)

	profile := models.Profile{
		ID:       req.Msg.GetId(),
		FullName: null.StringFrom(req.Msg.GetFullName()),
		OrgID:    req.Msg.GetOrgId(),
	}

	err := profile.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("unable to insert record: %w", err)
	}

	res := &profilePB.CreateProfileResponse{}

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}

// GetProfile - Gets a profile.
func (s *Service) GetProfile(
	ctx context.Context,
	req *connect.Request[profilePB.GetProfileRequest],
) (*connect.Response[profilePB.GetProfileResponse], error) {
	record, err := models.FindProfile(ctx, s.db, req.Msg.GetId())
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve item: %w", err)
	}

	return connect.NewResponse(&profilePB.GetProfileResponse{
		Profile: &profilePB.Profile{
			Id:       record.ID,
			FullName: record.FullName.String,
			OrgId:    record.OrgID,
		},
	}), nil
}

// ListProfiles - Lists profiles.
func (s *Service) ListProfiles(
	ctx context.Context,
	req *connect.Request[profilePB.ListProfilesRequest],
) (*connect.Response[profilePB.ListProfilesResponse], error) {
	var mods []qm.QueryMod
	// TODO account for parent and page_token fields
	mods = append(mods, qm.Limit(max(1, int(req.Msg.GetPageSize()))))

	items, err := models.Profiles(mods...).All(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve items: %w", err)
	}

	var profiles []*profilePB.Profile
	for _, item := range items {
		profiles = append(profiles, &profilePB.Profile{
			Id:       item.ID,
			FullName: item.FullName.String,
			OrgId:    item.OrgID,
		})
	}

	return connect.NewResponse(&profilePB.ListProfilesResponse{
		Profiles: profiles,
	}), nil
}
