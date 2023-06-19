// Package service implements this service's API handlers.
package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"
	orgv1beta1 "go.buf.build/bufbuild/connect-go/kevinmichaelchen/orgapis/org/v1beta1"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/simulated"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateOrg(
	_ context.Context,
	req *connect.Request[orgv1beta1.CreateOrgRequest],
) (*connect.Response[orgv1beta1.CreateOrgResponse], error) {
	// Sleep for a bit to simulate the latency of a database lookup
	simulated.Sleep()

	// Simulate a potential error to test retry logic
	err := simulated.PossibleError(simulated.Low)
	if err != nil {
		return nil, err
	}

	res := &orgv1beta1.CreateOrgResponse{}

	logrus.WithField("name", req.Msg.GetName()).Info("Creating Org")

	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")

	return out, nil
}
