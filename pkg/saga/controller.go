// Package saga provides a controller for our Temporal workflow.
package saga

import "google.golang.org/grpc"

// Controller - An activity struct whose exported methods are treated as
// Temporal activities. The default name of each activity is the method name.
type Controller struct {
	connLicense *grpc.ClientConn
	connOrg     *grpc.ClientConn
	connProfile *grpc.ClientConn
}

// NewController - Creates a new Controller.
func NewController(
	connLicense *grpc.ClientConn,
	connOrg *grpc.ClientConn,
	connProfile *grpc.ClientConn,
) *Controller {
	return &Controller{
		connLicense: connLicense,
		connOrg:     connOrg,
		connProfile: connProfile,
	}
}
