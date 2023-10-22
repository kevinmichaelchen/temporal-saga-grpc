// Package service implements this service's API handlers.
package service

import (
	"context"
	"fmt"

	licenseConnect "buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	licensev1beta1 "buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go/license/v1beta1"
	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/sirupsen/logrus"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

var _ licenseConnect.LicenseServiceHandler = (*Service)(nil)

// Service - A controller for our business logic.
type Service struct{}

// NewService - Returns a new Service.
func NewService() *Service {
	return &Service{}
}

// CreateLicense - Creates a new license.
func (s *Service) CreateLicense(
	_ context.Context,
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

	res := &licensev1beta1.CreateLicenseResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating License")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
