// Package pgv provides a function to convert protoc-gen-validate (PGV) errors
// to Connect Go errors.
package pgv

import (
	"fmt"

	"github.com/bufbuild/connect-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

// Error - A single validation error.
type Error interface {
	Field() string
	Reason() string
}

// MultiError - A list of validation errors.
type MultiError interface {
	AllErrors() []error
}

// Convert - Converts a PGV error to *connect.Error.
func Convert(err error) error {
	var fieldViolations []*errdetails.BadRequest_FieldViolation

	pgve, ok := err.(Error)
	if ok {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       pgve.Field(),
			Description: pgve.Reason(),
		})
	}

	pgves, ok := err.(MultiError)
	if ok {
		for _, p := range pgves.AllErrors() {
			pgve, ok := p.(Error)
			if ok {
				fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       pgve.Field(),
					Description: pgve.Reason(),
				})
			}
		}
	}

	cErr := connect.NewError(
		connect.CodeInvalidArgument,
		fmt.Errorf("invalid request: %w", err),
	)

	br := &errdetails.BadRequest{
		FieldViolations: fieldViolations,
	}
	if detail, detailErr := connect.NewErrorDetail(br); detailErr == nil {
		cErr.AddDetail(detail)
	}

	return cErr
}
