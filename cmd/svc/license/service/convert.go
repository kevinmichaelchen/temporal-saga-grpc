package service

import (
	"fmt"
	"github.com/bufbuild/connect-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type PGVError interface {
	Field() string
	Reason() string
}

type MultiError interface {
	AllErrors() []error
}

// convert PGV error to *connect.Error
func convert(err error) error {
	var fv []*errdetails.BadRequest_FieldViolation
	pgve, ok := err.(PGVError)
	if ok {
		fv = append(fv, &errdetails.BadRequest_FieldViolation{
			Field:       pgve.Field(),
			Description: pgve.Reason(),
		})
	}

	pgves, ok := err.(MultiError)
	if ok {
		for _, p := range pgves.AllErrors() {
			pgve, ok := p.(PGVError)
			if ok {
				fv = append(fv, &errdetails.BadRequest_FieldViolation{
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
		FieldViolations: fv,
	}
	if detail, detailErr := connect.NewErrorDetail(br); detailErr == nil {
		cErr.AddDetail(detail)
	}
	return cErr
}
