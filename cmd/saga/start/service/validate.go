package service

import (
	temporalv1beta1 "buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go/temporal/v1beta1"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
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
