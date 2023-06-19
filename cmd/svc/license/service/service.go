// Package service implements this service's API handlers.
package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"

	licensev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/license/v1beta1"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/connect/pgv"
	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateLicense(
	ctx context.Context,
	req *connect.Request[licensev1beta1.CreateLicenseRequest],
) (*connect.Response[licensev1beta1.CreateLicenseResponse], error) {
	err := req.Msg.Validate()
	if err != nil {
		return nil, pgv.Convert(err)
	}

	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err = simulated.PossibleError(simulated.MediumHigh)
	if err != nil {
		return nil, err
	}

	res := &licensev1beta1.CreateLicenseResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating License")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
