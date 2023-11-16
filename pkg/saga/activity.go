package saga

import (
	"context"
	"fmt"

	license "buf.build/gen/go/kevinmichaelchen/licenseapis/protocolbuffers/go/license/v1beta1"
	org "buf.build/gen/go/kevinmichaelchen/orgapis/protocolbuffers/go/org/v1beta1"
	profile "buf.build/gen/go/kevinmichaelchen/profileapis/protocolbuffers/go/profile/v1beta1"
	"connectrpc.com/connect"
)

// CreateOrg - A Temporal Activity for creating an Org.
func (c *Controller) CreateOrg(ctx context.Context, args CreateLicenseInputArgs) error {
	_, err := c.orgClient.CreateOrg(ctx, connect.NewRequest(
		&org.CreateOrgRequest{
			Name: args.OrgName,
		},
	))
	if err != nil {
		return fmt.Errorf("unable to create org: %w", err)
	}

	return nil
}

// CreateProfile - A Temporal Activity for creating a user Profile.
func (c *Controller) CreateProfile(ctx context.Context, args CreateLicenseInputArgs) error {
	_, err := c.profileClient.CreateProfile(ctx, connect.NewRequest(
		&profile.CreateProfileRequest{
			Name: args.ProfileName,
		},
	))
	if err != nil {
		return fmt.Errorf("unable to create profile: %w", err)
	}

	return nil
}

// CreateLicense - A Temporal Activity for creating a License.
func (c *Controller) CreateLicense(ctx context.Context, args CreateLicenseInputArgs) error {
	_, err := c.licenseClient.CreateLicense(ctx, connect.NewRequest(
		&license.CreateLicenseRequest{
			Name: args.LicenseName,
		},
	))
	if err != nil {
		return fmt.Errorf("unable to create license: %w", err)
	}

	return nil
}
