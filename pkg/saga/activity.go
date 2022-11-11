package saga

import (
	"context"
	"fmt"
	licensev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/license/v1beta1"
	orgv1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/org/v1beta1"
	profilev1beta1 "github.com/kevinmichaelchen/temporal-saga-grpc/internal/idl/profile/v1beta1"
)

// TODO why isn't Remote Generation creating gRPC clients?

func (c *Controller) CreateOrg(ctx context.Context, args CreateLicenseInputArgs) error {
	client := orgv1beta1.NewOrgServiceClient(c.connOrg)

	_, err := client.CreateOrg(ctx, &orgv1beta1.CreateOrgRequest{
		Name: args.OrgName,
	})
	if err != nil {
		return fmt.Errorf("failed to create org: %w", err)
	}

	return nil
}

func (c *Controller) CreateProfile(ctx context.Context, args CreateLicenseInputArgs) error {
	client := profilev1beta1.NewProfileServiceClient(c.connProfile)

	_, err := client.CreateProfile(ctx, &profilev1beta1.CreateProfileRequest{
		Name: args.ProfileName,
	})
	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	return nil
}

func (c *Controller) CreateLicense(ctx context.Context, args CreateLicenseInputArgs) error {
	client := licensev1beta1.NewLicenseServiceClient(c.connLicense)

	_, err := client.CreateLicense(ctx, &licensev1beta1.CreateLicenseRequest{
		Name: args.LicenseName,
	})
	if err != nil {
		return fmt.Errorf("failed to create license: %w", err)
	}

	return nil
}
