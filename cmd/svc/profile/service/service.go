// Package service implements this service's API handlers.
package service

import (
	profileConnect "buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
	"context"

	profilePB "buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go/profile/v1beta1"
	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

var _ profileConnect.ProfileServiceHandler = (*Service)(nil)

// Service - A controller for our business logic.
type Service struct{}

// NewService - Returns a new Service.
func NewService() *Service {
	return &Service{}
}

// CreateProfile - Creates a new user profile.
func (s *Service) CreateProfile(
	_ context.Context,
	req *connect.Request[profilePB.CreateProfileRequest],
) (*connect.Response[profilePB.CreateProfileResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.MediumLow)
	if err != nil {
		return nil, err
	}

	res := &profilePB.CreateProfileResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating Profile")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
