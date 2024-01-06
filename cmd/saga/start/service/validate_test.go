package service

import (
	"testing"
	"time"

	temporalPB "buf.build/gen/go/kevinmichaelchen/temporalapis/protocolbuffers/go/temporal/v1beta1"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	buildValid := func() *temporalPB.CreateOnboardingWorkflowRequest {
		now := time.Now()

		return &temporalPB.CreateOnboardingWorkflowRequest{
			WorkflowOptions: &temporalPB.WorkflowOptions{
				WorkflowId: uuid.New().String(),
			},
			License: &temporalPB.License{
				Start: timestamppb.New(now),
				End:   timestamppb.New(now.Add(time.Hour * 24 * 30)),
			},
			Org: &temporalPB.Org{
				Name: "Name",
			},
			Profile: &temporalPB.Profile{
				FullName: "Name",
			},
		}
	}

	tests := map[string]struct {
		build  func() *temporalPB.CreateOnboardingWorkflowRequest
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
			build: func() *temporalPB.CreateOnboardingWorkflowRequest {
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
		"Invalid - License end date must be after start date": {
			build: func() *temporalPB.CreateOnboardingWorkflowRequest {
				out := buildValid()
				out.License.End = timestamppb.New(out.GetLicense().GetStart().AsTime().Add(-1 * time.Hour))

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
