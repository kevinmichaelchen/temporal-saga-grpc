package saga

import (
	"context"
	"fmt"

	licenseGRPC "buf.build/gen/go/kevinmichaelchen/licenseapis/grpc/go/license/v1beta1/licensev1beta1grpc"
	license "buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go/license/v1beta1"
	orgGRPC "buf.build/gen/go/kevinmichaelchen/orgapis/grpc/go/org/v1beta1/orgv1beta1grpc"
	org "buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go/org/v1beta1"
	profileGRPC "buf.build/gen/go/kevinmichaelchen/profileapis/grpc/go/profile/v1beta1/profilev1beta1grpc"
	profile "buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go/profile/v1beta1"
)

// CreateOrg - A Temporal Activity for creating an Org.
func (c *Controller) CreateOrg(ctx context.Context, args CreateLicenseInputArgs) error {
	client := orgGRPC.NewOrgServiceClient(c.connOrg)

	_, err := client.CreateOrg(ctx, &org.CreateOrgRequest{
		Name: args.OrgName,
	})
	if err != nil {
		return fmt.Errorf("unable to create org: %w", err)
	}

	return nil
}

// CreateProfile - A Temporal Activity for creating a user Profile.
func (c *Controller) CreateProfile(ctx context.Context, args CreateLicenseInputArgs) error {
	client := profileGRPC.NewProfileServiceClient(c.connProfile)

	_, err := client.CreateProfile(ctx, &profile.CreateProfileRequest{
		Name: args.ProfileName,
	})
	if err != nil {
		return fmt.Errorf("unable to create profile: %w", err)
	}

	return nil
}

// CreateLicense - A Temporal Activity for creating a License.
func (c *Controller) CreateLicense(ctx context.Context, args CreateLicenseInputArgs) error {
	client := licenseGRPC.NewLicenseServiceClient(c.connLicense)

	_, err := client.CreateLicense(ctx, &license.CreateLicenseRequest{
		Name: args.LicenseName,
	})
	if err != nil {
		return fmt.Errorf("unable to create license: %w", err)
	}

	return nil
}
