package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	profilev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/profile/v1beta1"
	"github.com/sirupsen/logrus"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateProfile(
	ctx context.Context,
	req *connect.Request[profilev1beta1.CreateProfileRequest],
) (*connect.Response[profilev1beta1.CreateProfileResponse], error) {
	res := &profilev1beta1.CreateProfileResponse{}
	//if err != nil {
	//	return nil, err
	//}
	logrus.Info("Creating Profile")
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}
