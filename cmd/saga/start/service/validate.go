package service

import (
	"fmt"

	temporalv1beta1 "buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go/temporal/v1beta1"
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
