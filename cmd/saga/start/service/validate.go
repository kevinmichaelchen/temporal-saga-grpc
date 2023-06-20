package service

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"

	temporalv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/temporal/v1beta1"
)

func validate(
	validator *protovalidate.Validator,
	req *temporalv1beta1.CreateOnboardingWorkflowRequest,
) error {
	err := validator.Validate(req)
	if err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	return nil
}
