package service

import (
	"github.com/bufbuild/connect-go"
	"github.com/google/go-cmp/cmp"
	licensev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/license/v1beta1"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := map[string]struct {
		r           *licensev1beta1.CreateLicenseRequest
		validateAll bool
		expectedFVs []*errdetails.BadRequest_FieldViolation
	}{
		"validateAll - 1 bad field": {
			validateAll: true,
			r: &licensev1beta1.CreateLicenseRequest{
				Name:     "",
				FullName: "John Doe",
			},
			expectedFVs: []*errdetails.BadRequest_FieldViolation{
				newBRFV("Name", "value length must be at least 1 runes"),
			},
		},
		"validateAll - multiple bad fields": {
			validateAll: true,
			r: &licensev1beta1.CreateLicenseRequest{
				Name:     "",
				FullName: "",
			},
			expectedFVs: []*errdetails.BadRequest_FieldViolation{
				newBRFV("Name", "value length must be at least 1 runes"),
				newBRFV("FullName", "value length must be at least 1 runes"),
			},
		},
		"1 bad field": {
			r: &licensev1beta1.CreateLicenseRequest{
				Name:     "",
				FullName: "John Doe",
			},
			expectedFVs: []*errdetails.BadRequest_FieldViolation{
				newBRFV("Name", "value length must be at least 1 runes"),
			},
		},
		"multiple bad fields": {
			r: &licensev1beta1.CreateLicenseRequest{
				Name:     "",
				FullName: "",
			},
			expectedFVs: []*errdetails.BadRequest_FieldViolation{
				newBRFV("Name", "value length must be at least 1 runes"),
			},
		},
	}
	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			var err error
			if tc.validateAll {
				err = tc.r.ValidateAll()
			} else {
				err = tc.r.Validate()
			}

			// There should be an error
			require.Error(t, err)

			// Convert that error to a *connect.Error
			err = convert(err)
			cErr, ok := err.(*connect.Error)
			require.True(t, ok)

			// Get the error details
			details := cErr.Details()
			require.NotEmpty(t, details)

			// Check the first detail object
			expected := &errdetails.BadRequest{
				FieldViolations: tc.expectedFVs,
			}
			dtv, err := details[0].Value()
			require.NoError(t, err)

			diff := cmp.Diff(dtv, expected,
				protocmp.Transform(),
			)
			require.Empty(t, diff)
		})
	}
}

func newBRFV(field, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: desc,
	}
}
