package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	orgv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/com/teachingstrategies/org/v1beta1"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateOrg(
	ctx context.Context,
	req *connect.Request[orgv1beta1.CreateOrgRequest],
) (*connect.Response[orgv1beta1.CreateOrgResponse], error) {
	res := &orgv1beta1.CreateOrgResponse{}
	//if err != nil {
	//	return nil, err
	//}
	out := connect.NewResponse(res)
	out.Header().Set("API-Version", "v1beta1")
	return out, nil
}
