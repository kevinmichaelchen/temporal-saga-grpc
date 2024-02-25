// Package workflow provides a controller for our Temporal workflow.
package workflow

import (
	"buf.build/gen/go/kevinmichaelchen/licenseapis/connectrpc/go/license/v1beta1/licensev1beta1connect"
	"buf.build/gen/go/kevinmichaelchen/orgapis/connectrpc/go/org/v1beta1/orgv1beta1connect"
	"buf.build/gen/go/kevinmichaelchen/profileapis/connectrpc/go/profile/v1beta1/profilev1beta1connect"
)

// Controller - An activity struct whose exported methods are treated as
// Temporal activities. The default name of each activity is the method name.
type Controller struct {
	licenseClient licensev1beta1connect.LicenseServiceClient
	orgClient     orgv1beta1connect.OrgServiceClient
	profileClient profilev1beta1connect.ProfileServiceClient
}

// NewController - Creates a new Controller.
func NewController(
	licenseClient licensev1beta1connect.LicenseServiceClient,
	orgClient orgv1beta1connect.OrgServiceClient,
	profileClient profilev1beta1connect.ProfileServiceClient,
) *Controller {
	return &Controller{
		licenseClient: licenseClient,
		orgClient:     orgClient,
		profileClient: profileClient,
	}
}
