package saga

import (
	"context"
	"fmt"
	licensev1beta12 "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/idl/com/teachingstrategies/license/v1beta1"
	orgv1beta12 "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/idl/com/teachingstrategies/org/v1beta1"
	profilev1beta12 "github.com/kevinmichaelchen/temporal-saga-grpc/pkg/idl/com/teachingstrategies/profile/v1beta1"
)

func (c *Controller) CreateOrg(ctx context.Context, args CreateLicenseInputArgs) error {
	client := orgv1beta12.NewOrgServiceClient(c.connOrg)

	_, err := client.CreateOrg(ctx, &orgv1beta12.CreateOrgRequest{
		Name: args.OrgName,
	})
	if err != nil {
		return fmt.Errorf("failed to create org: %w", err)
	}

	return nil
}

func (c *Controller) CreateProfile(ctx context.Context, args CreateLicenseInputArgs) error {
	client := profilev1beta12.NewProfileServiceClient(c.connProfile)

	_, err := client.CreateProfile(ctx, &profilev1beta12.CreateProfileRequest{
		Name: args.ProfileName,
	})
	if err != nil {
		return fmt.Errorf("failed to create profile: %w", err)
	}

	return nil
}

func (c *Controller) CreateLicense(ctx context.Context, args CreateLicenseInputArgs) error {
	client := licensev1beta12.NewLicenseServiceClient(c.connLicense)

	_, err := client.CreateLicense(ctx, &licensev1beta12.CreateLicenseRequest{
		Name: args.LicenseName,
	})
	if err != nil {
		return fmt.Errorf("failed to create license: %w", err)
	}

	return nil
}
