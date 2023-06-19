// Package service implements this service's API handlers.
package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"
	profilev1beta1 "go.buf.build/bufbuild/connect-go/kevinmichaelchen/profileapis/profile/v1beta1"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

// Service - A controller for our business logic.
type Service struct{}

// NewService - Returns a new Service.
func NewService() *Service {
	return &Service{}
}

// CreateProfile - Creates a new user profile.
func (s *Service) CreateProfile(
	_ context.Context,
	req *connect.Request[profilev1beta1.CreateProfileRequest],
) (*connect.Response[profilev1beta1.CreateProfileResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.MediumLow)
	if err != nil {
		return nil, err
	}

	res := &profilev1beta1.CreateProfileResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating Profile")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
