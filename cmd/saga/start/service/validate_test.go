package service

import (
	"testing"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	temporalv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/temporal/v1beta1"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	buildValid := func() *temporalv1beta1.CreateOnboardingWorkflowRequest {
		start := time.Date(2023, 6, 19, 12, 30, 0, 0, time.UTC)
		end := time.Date(2024, 6, 19, 12, 30, 0, 0, time.UTC)

		return &temporalv1beta1.CreateOnboardingWorkflowRequest{
			WorkflowOptions: &temporalv1beta1.WorkflowOptions{
				WorkflowId: uuid.New().String(),
			},
			License: &temporalv1beta1.License{
				Start: timestamppb.New(start),
				End:   timestamppb.New(end),
			},
			Org: &temporalv1beta1.Org{
				Name: "Name",
			},
			Profile: &temporalv1beta1.Profile{
				Name: "Name",
			},
		}
	}

	tests := map[string]struct {
		build  func() *temporalv1beta1.CreateOnboardingWorkflowRequest
		expect func(t *testing.T, err error)
	}{
		"OK": {
			build: buildValid,
			expect: func(t *testing.T, err error) {
				t.Helper()
				require.NoError(t, err)
			},
		},
		"Invalid - Workflow ID should be UUID": {
			build: func() *temporalv1beta1.CreateOnboardingWorkflowRequest {
				out := buildValid()
				out.WorkflowOptions.WorkflowId = "foobar"

				return out
			},
			expect: func(t *testing.T, err error) {
				t.Helper()
				require.EqualError(t, err, `invalid request: validation error:
 - workflow_options.workflow_id: value must be a valid UUID [string.uuid]`)
			},
		},
		"Invalid - License start should precede its end": {
			build: func() *temporalv1beta1.CreateOnboardingWorkflowRequest {
				out := buildValid()
				out.License.Start = timestamppb.New(out.License.End.AsTime().Add(time.Hour))

				return out
			},
			expect: func(t *testing.T, err error) {
				t.Helper()
				require.EqualError(t, err, `invalid request: validation error:
 - license: end time must be after start time [transaction.expiration_date]`)
			},
		},
	}

	validator, err := protovalidate.New()
	require.NoError(t, err)

	for testName, tc := range tests {
		// https://github.com/kunwardeep/paralleltest#tparallel-is-called-in-the-range-method-and-test-case-variable-tc-being-used-but-is-not-reinitialised-more-info
		// https://gist.github.com/kunwardeep/80c2e9f3d3256c894898bae82d9f75d0
		tc := tc

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			err = validate(validator, tc.build())
			tc.expect(t, err)
		})
	}
}
