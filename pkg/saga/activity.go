package saga

import (
	"context"
	"fmt"

	licensev1beta1 "go.buf.build/grpc/go/kevinmichaelchen/licenseapis/license/v1beta1"
	orgv1beta1 "go.buf.build/grpc/go/kevinmichaelchen/orgapis/org/v1beta1"
	profilev1beta1 "go.buf.build/grpc/go/kevinmichaelchen/profileapis/profile/v1beta1"
)

// CreateOrg - A Temporal Activity for creating an Org.
func (c *Controller) CreateOrg(ctx context.Context, args CreateLicenseInputArgs) error {
	client := orgv1beta1.NewOrgServiceClient(c.connOrg)

	_, err := client.CreateOrg(ctx, &orgv1beta1.CreateOrgRequest{
		Name: args.OrgName,
	})
	if err != nil {
		return fmt.Errorf("unable to create org: %w", err)
	}

	return nil
}

// CreateProfile - A Temporal Activity for creating a user Profile.
func (c *Controller) CreateProfile(ctx context.Context, args CreateLicenseInputArgs) error {
	client := profilev1beta1.NewProfileServiceClient(c.connProfile)

	_, err := client.CreateProfile(ctx, &profilev1beta1.CreateProfileRequest{
		Name: args.ProfileName,
	})
	if err != nil {
		return fmt.Errorf("unable to create profile: %w", err)
	}

	return nil
}

// CreateLicense - A Temporal Activity for creating a License.
func (c *Controller) CreateLicense(ctx context.Context, args CreateLicenseInputArgs) error {
	client := licensev1beta1.NewLicenseServiceClient(c.connLicense)

	_, err := client.CreateLicense(ctx, &licensev1beta1.CreateLicenseRequest{
		Name: args.LicenseName,
	})
	if err != nil {
		return fmt.Errorf("unable to create license: %w", err)
	}

	return nil
}
