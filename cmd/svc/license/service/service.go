package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	licensev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/license/v1beta1"
	"github.com/sirupsen/logrus"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateLicense(
	ctx context.Context,
	req *connect.Request[licensev1beta1.CreateLicenseRequest],
) (*connect.Response[licensev1beta1.CreateLicenseResponse], error) {
	res := &licensev1beta1.CreateLicenseResponse{}
	//if err != nil {
	//	return nil, err
	//}
	logrus.Info("Creating License")
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}
